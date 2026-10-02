package apify

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
)

// Status-code thresholds used to classify responses. Named constants so the retry
// policy has no magic numbers (see the language-agnostic coding rules).
const (
	// rateLimitExceededStatusCode is returned when the per-resource rate limit is hit.
	rateLimitExceededStatusCode = 429
	// minServerErrorStatusCode and above are treated as retryable internal errors.
	minServerErrorStatusCode = 500
	// maxSuccessStatusCode: responses with status < this are treated as success.
	maxSuccessStatusCode = 300
	// backoffFactor is the multiplier applied to the inter-retry delay after each
	// attempt (exponential backoff). Matches the reference client's async-retry factor of 2.
	backoffFactor = 2
	// minCompressRequestBytes is the smallest request body worth compressing. Below this
	// size the CPU cost and Content-Encoding header overhead outweigh the bandwidth saved,
	// so the body is sent as-is. Matches the reference client's MIN_COMPRESS_BYTES threshold.
	minCompressRequestBytes = 1024
)

// Content-Encoding values for compressed request bodies. The reference client prefers brotli
// and falls back to gzip, so this client supports both: brotli (best ratio) is tried first and
// gzip is the fallback when brotli is unavailable or fails to shrink the payload.
const (
	// contentEncodingBrotli is the preferred Content-Encoding (Brotli, RFC 7932).
	contentEncodingBrotli = "br"
	// contentEncodingGzip is the fallback Content-Encoding (gzip, RFC 1952).
	contentEncodingGzip = "gzip"
)

// HTTPBackend is the replaceable transport contract of the client.
//
// Implementations are responsible only for sending a single request and returning the
// raw response. Authentication, the User-Agent header, retries and (de)serialization are
// handled by httpClient, so a backend only needs to perform one network round-trip.
//
// A non-2xx HTTP status is NOT an error at this layer — return it as a normal
// *http.Response. Only transport-level failures (connection refused, DNS, timeout)
// should be returned as an error.
type HTTPBackend interface {
	// Do sends a single HTTP request and returns the response.
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPBackend is the default HTTPBackend implementation, backed by an
// *http.Client. The per-attempt timeout is applied via the request context by the
// orchestrating client, so this backend uses no client-level timeout of its own.
type DefaultHTTPBackend struct {
	client *http.Client
}

// NewDefaultHTTPBackend creates a backend with a sensible default *http.Client that
// keeps connections alive for connection-pool reuse.
func NewDefaultHTTPBackend() *DefaultHTTPBackend {
	return &DefaultHTTPBackend{client: &http.Client{}}
}

// NewHTTPBackendWithClient wraps a caller-provided *http.Client, useful for sharing a
// connection pool or customizing proxy/TLS settings.
func NewHTTPBackendWithClient(client *http.Client) *DefaultHTTPBackend {
	return &DefaultHTTPBackend{client: client}
}

// Do implements HTTPBackend.
func (b *DefaultHTTPBackend) Do(req *http.Request) (*http.Response, error) {
	return b.client.Do(req)
}

// retryConfig configures the retry/timeout behaviour of the httpClient.
type retryConfig struct {
	// maxRetries is the maximum number of retries (the request is attempted up to
	// maxRetries+1 times).
	maxRetries int
	// minDelayBetweenRetries is doubled on each subsequent retry (exponential backoff).
	minDelayBetweenRetries time.Duration
	// timeoutMax caps the timeout of any single request attempt. It bounds the growth of a
	// tier's timeout across retries (see attemptTimeout) and the inter-retry delay alike.
	timeoutMax time.Duration
}

// timeoutTiers holds the configurable default request timeouts. Every exported method is
// assigned the tier that fits the expected duration of its request — short for simple metadata
// reads/writes, medium for listing/batch/trigger operations, long for downloads, uploads and
// streaming — matching the reference client's timeout tiers. A caller that needs a different
// bound for one call already has full control via the context.Context every method takes (see
// the "Cancellation" section of the README): canceling it, or giving it an earlier deadline via
// context.WithTimeout, overrides whichever tier the method would otherwise use, since
// doAttempt's context.WithTimeout always resolves to the earlier of the two deadlines.
type timeoutTiers struct {
	short  time.Duration
	medium time.Duration
	long   time.Duration
}

// noRequestTimeout means "no client-imposed timeout for this request" — only the caller's
// context bounds it. Used for the polling requests behind WaitForFinish/Call, which must not be
// aborted by a fixed per-request budget while legitimately waiting for a job to finish.
const noRequestTimeout time.Duration = 0

// httpClient is the orchestrating HTTP client shared by every resource client. It owns
// the backend, the optional API token, the User-Agent, and the retry/timeout policy, and
// applies them to every request. It is safe for concurrent use and cheap to copy.
type httpClient struct {
	backend     HTTPBackend
	token       string
	userAgent   string
	retry       retryConfig
	tiers       timeoutTiers
	compressors []requestCompressor
}

// apiResponse is the parsed result of a single API call: the status, headers and the
// fully-buffered response body.
type apiResponse struct {
	statusCode int
	header     http.Header
	body       []byte
}

// call sends req with authentication, the User-Agent header and the retry policy applied,
// returning the first successful response or the final error.
//
// body, if non-nil, is the raw request body; it is re-sent on each retry (we buffer it).
// contentType, if non-empty, sets the Content-Type header. baseTimeout is the per-endpoint
// base timeout; it grows with each attempt up to the client's overall timeout budget.
func (c *httpClient) call(ctx context.Context, method, url string, body []byte, contentType string, baseTimeout time.Duration) (*apiResponse, error) {
	return c.callWithHeaders(ctx, method, url, body, contentType, nil, baseTimeout)
}

// callWithHeaders is like call but additionally sets the given extra request headers (e.g.
// the charge idempotency key). They are applied on every attempt.
func (c *httpClient) callWithHeaders(ctx context.Context, method, url string, body []byte, contentType string, extraHeaders map[string]string, baseTimeout time.Duration) (*apiResponse, error) {
	delay := c.retry.minDelayBetweenRetries
	maxAttempts := c.retry.maxRetries + 1
	path := extractPath(url)

	// Compress the body once (not per attempt): it is identical on every retry. contentEncoding
	// is "" when the body was left uncompressed (too small, already compressed, or compression
	// did not shrink it).
	sendBody, contentEncoding := maybeCompressRequestBody(body, contentType, c.compressors)

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err := c.doAttempt(ctx, method, url, sendBody, contentType, contentEncoding, extraHeaders, c.attemptTimeout(baseTimeout, attempt))

		var retryable bool
		if err != nil {
			lastErr = err
			retryable = isTransportErrorRetryable(err)
		} else if resp.statusCode < maxSuccessStatusCode {
			return resp, nil
		} else {
			lastErr = buildAPIError(resp.statusCode, resp.body, attempt, method, path)
			retryable = isStatusRetryable(resp.statusCode)
		}

		// Give up immediately on non-retryable errors or after the last attempt.
		if !retryable || attempt == maxAttempts {
			return nil, lastErr
		}

		// Sleep with exponential backoff before the next attempt. The base delay doubles each
		// retry (matching the reference client's async-retry factor of 2). Two intentional and
		// documented divergences from the reference JS client:
		//   - Jitter: the actual sleep is drawn uniformly from [delay, 2*delay). The reference
		//     does not jitter, but the Apify API docs explicitly prescribe it ("wait for a period
		//     of time chosen randomly from the interval DELAY to 2*DELAY milliseconds") to avoid
		//     synchronized retries (thundering herd), so we follow the API guidance here.
		//   - Cap: the base delay is capped at the overall request timeout as a safety net (the
		//     reference leaves it uncapped). With the default settings this cap is never reached
		//     within maxRetries, so it does not change observable behaviour.
		if !sleepWithContext(ctx, randomizedDelay(delay)) {
			return nil, ctx.Err()
		}
		delay = minDuration(delay*backoffFactor, c.retry.timeoutMax)
	}

	return nil, lastErr
}

// applyAuthHeaders sets the User-Agent and (if a token is configured) Authorization headers
// that every request to the API must carry. It is the single place those headers are set, so
// the buffered request path and the raw streaming path (LogClient.Stream) stay in sync.
func (c *httpClient) applyAuthHeaders(req *http.Request) {
	req.Header.Set("User-Agent", c.userAgent)
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
}

// doAttempt performs a single HTTP round-trip with the given per-attempt timeout.
//
// contentEncoding, when non-empty, is set as the Content-Encoding header (the body is already
// compressed with that encoding by the caller).
func (c *httpClient) doAttempt(ctx context.Context, method, url string, body []byte, contentType, contentEncoding string, extraHeaders map[string]string, timeout time.Duration) (*apiResponse, error) {
	attemptCtx := ctx
	cancel := func() {}
	if timeout != noRequestTimeout {
		attemptCtx, cancel = context.WithTimeout(ctx, timeout)
	}
	defer cancel()

	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(attemptCtx, method, url, reader)
	if err != nil {
		// A malformed request is a programming/argument error, not a transport error.
		return nil, &APIError{StatusCode: 0, Type: "invalid-request", Message: err.Error(), Attempt: 1, HTTPMethod: method, Path: extractPath(url)}
	}

	c.applyAuthHeaders(req)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if contentEncoding != "" {
		req.Header.Set("Content-Encoding", contentEncoding)
	}
	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.backend.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		// Reading the body failed mid-stream; treat as a retryable transport error.
		return nil, err
	}

	return &apiResponse{statusCode: resp.StatusCode, header: resp.Header, body: data}, nil
}

// requestCompressor pairs a Content-Encoding value with the function that produces a body
// in that encoding. Splitting the encodings this way lets maybeCompressRequestBody try them
// in preference order and keeps each codec independently testable.
type requestCompressor struct {
	// encoding is the Content-Encoding header value this codec produces (e.g. "br", "gzip").
	encoding string
	// compress returns the compressed form of body, or an error if that codec is unavailable.
	compress func(body []byte) ([]byte, error)
}

// RequestCompression selects which codec compresses request bodies (see [WithRequestCompression]).
type RequestCompression int

const (
	// RequestCompressionAuto prefers brotli (typically a better ratio) and falls back to gzip
	// when brotli errors or fails to shrink the body; this is the default. Matches the
	// reference client's brotli-then-gzip preference.
	RequestCompressionAuto RequestCompression = iota
	// RequestCompressionBrotli always uses brotli; a body it does not shrink is sent
	// uncompressed rather than falling back to gzip.
	RequestCompressionBrotli
	// RequestCompressionGzip always uses gzip; a body it does not shrink is sent uncompressed.
	RequestCompressionGzip
)

// compressorsFor returns the codec(s) maybeCompressRequestBody tries, in preference order, for
// the given [RequestCompression] choice.
func compressorsFor(algo RequestCompression) []requestCompressor {
	switch algo {
	case RequestCompressionBrotli:
		return []requestCompressor{{encoding: contentEncodingBrotli, compress: brotliCompress}}
	case RequestCompressionGzip:
		return []requestCompressor{{encoding: contentEncodingGzip, compress: gzipCompress}}
	default:
		return []requestCompressor{
			{encoding: contentEncodingBrotli, compress: brotliCompress},
			{encoding: contentEncodingGzip, compress: gzipCompress},
		}
	}
}

// compressBody runs body through the streaming writer built by newWriter and returns the
// compressed bytes. newWriter's WriteCloser must flush all buffered output on Close. This is the
// shared body of the per-codec helpers, which differ only in the writer they construct.
func compressBody(body []byte, newWriter func(io.Writer) io.WriteCloser) ([]byte, error) {
	var buf bytes.Buffer
	w := newWriter(&buf)
	if _, err := w.Write(body); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// brotliCompress compresses body with Brotli (Content-Encoding: br).
func brotliCompress(body []byte) ([]byte, error) {
	return compressBody(body, func(w io.Writer) io.WriteCloser { return brotli.NewWriter(w) })
}

// gzipCompress compresses body with gzip (Content-Encoding: gzip).
func gzipCompress(body []byte) ([]byte, error) {
	return compressBody(body, func(w io.Writer) io.WriteCloser { return gzip.NewWriter(w) })
}

// maybeCompressRequestBody compresses body when it is large enough to be worth it and its
// content type is not already compressed, returning the bytes to send and the Content-Encoding
// value ("" when the body is sent uncompressed).
//
// This mirrors the reference client, which compresses request payloads above a 1 KiB threshold
// to save upload bandwidth. compressors is the configured codec list (see
// [RequestCompression]/[WithRequestCompression]), tried in order; if a codec errors or does not
// actually shrink the body (e.g. an already-compressed payload the content type did not flag),
// the next one is tried, and if none help the original body is sent uncompressed rather than
// failing the request.
func maybeCompressRequestBody(body []byte, contentType string, compressors []requestCompressor) ([]byte, string) {
	if len(body) < minCompressRequestBytes {
		return body, ""
	}
	if !isCompressibleContentType(contentType) {
		return body, ""
	}
	return compressWith(body, compressors)
}

// Media-type prefixes whose payloads already carry their own compression, so running them
// through brotli/gzip burns CPU and memory for a result that is usually no smaller — and the
// request also keeps the intended Content-Type, rather than becoming an unreadable
// Content-Encoding: br blob the destination may not expect for these formats. Matches the
// reference client's ALREADY_COMPRESSED_MEDIA_TYPE_PREFIXES.
var alreadyCompressedContentTypePrefixes = []string{"audio/", "image/", "video/"}

// Exact media types (outside the prefixes above) whose payloads already carry their own
// compression. Matches the reference client's ALREADY_COMPRESSED_MEDIA_TYPES.
var alreadyCompressedContentTypes = map[string]bool{
	"application/epub+zip":                    true,
	"application/gzip":                        true,
	"application/java-archive":                true,
	"application/vnd.android.package-archive": true,
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": true,
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         true,
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   true,
	"application/vnd.rar":          true,
	"application/x-7z-compressed":  true,
	"application/x-bzip":           true,
	"application/x-bzip2":          true,
	"application/x-gzip":           true,
	"application/x-rar-compressed": true,
	"application/x-xz":             true,
	"application/x-zip-compressed": true,
	"application/zip":              true,
	"application/zstd":             true,
	"font/woff":                    true,
	"font/woff2":                   true,
}

// Uncompressed media types that sit under an already-compressed prefix above, so compressing
// them still pays off. Matches the reference client's COMPRESSIBLE_MEDIA_TYPES.
var compressibleContentTypes = map[string]bool{
	"audio/aiff":                true,
	"audio/basic":               true,
	"audio/l16":                 true,
	"audio/l24":                 true,
	"audio/midi":                true,
	"audio/vnd.wave":            true,
	"audio/wav":                 true,
	"audio/wave":                true,
	"audio/x-aiff":              true,
	"audio/x-wav":               true,
	"image/bmp":                 true,
	"image/tiff":                true,
	"image/vnd.adobe.photoshop": true,
	"image/vnd.microsoft.icon":  true,
	"image/x-icon":              true,
	"image/x-ms-bmp":            true,
}

// Structured-syntax suffixes that mark a media type as text even under an already-compressed
// prefix (e.g. image/svg+xml). Matches the reference client's COMPRESSIBLE_MEDIA_TYPE_SUFFIXES.
var compressibleContentTypeSuffixes = []string{"+json", "+xml"}

// isCompressibleContentType decides whether a request body with the given Content-Type is
// worth compressing.
//
// Images, audio, video and archives already carry their own compression: running them through
// brotli or gzip burns CPU, holds a second full copy of the body in memory, and usually produces
// output no smaller than the input (sometimes larger). A format that is raw despite such a media
// type, e.g. image/bmp or audio/wav, is still compressed. An empty Content-Type is assumed
// compressible. Matches the reference client's isCompressibleContentType.
func isCompressibleContentType(contentType string) bool {
	if contentType == "" {
		return true
	}
	// Content-Type is case-insensitive and may carry parameters, e.g. "text/plain; charset=utf-8".
	mediaType := contentType
	if i := strings.IndexByte(mediaType, ';'); i >= 0 {
		mediaType = mediaType[:i]
	}
	mediaType = strings.ToLower(strings.TrimSpace(mediaType))

	if compressibleContentTypes[mediaType] {
		return true
	}
	for _, suffix := range compressibleContentTypeSuffixes {
		if strings.HasSuffix(mediaType, suffix) {
			return true
		}
	}
	if alreadyCompressedContentTypes[mediaType] {
		return false
	}
	for _, prefix := range alreadyCompressedContentTypePrefixes {
		if strings.HasPrefix(mediaType, prefix) {
			return false
		}
	}
	return true
}

// compressWith applies the given codecs in order and returns the first result that is smaller
// than body, along with its encoding. If no codec shrinks the body (or all error), it returns
// the original body with an empty encoding. It is separated from maybeCompressRequestBody so the
// preference-and-fallback logic can be tested with custom codec lists (e.g. a failing brotli).
func compressWith(body []byte, compressors []requestCompressor) ([]byte, string) {
	for _, c := range compressors {
		out, err := c.compress(body)
		if err != nil {
			continue // codec unavailable/failed: fall back to the next one
		}
		if len(out) < len(body) {
			return out, c.encoding
		}
	}
	return body, ""
}

// attemptTimeout returns min(overall, base * 2^(attempt-1)).
//
// The first attempt uses the per-endpoint base timeout; each retry doubles it so a
// slow-but-progressing connection gets more time, while never exceeding the overall budget.
// base == noRequestTimeout is returned unchanged: such a request grows no timeout of its own
// across retries, relying solely on the caller's context.
func (c *httpClient) attemptTimeout(base time.Duration, attempt int) time.Duration {
	if base == noRequestTimeout {
		return noRequestTimeout
	}
	scaled := base
	for i := 1; i < attempt; i++ {
		scaled *= 2
		if scaled >= c.retry.timeoutMax {
			return c.retry.timeoutMax
		}
	}
	return minDuration(scaled, c.retry.timeoutMax)
}

// isStatusRetryable reports whether a non-success status should be retried. We retry 429
// (rate limit) and 5xx (internal server errors); other 4xx statuses are caller errors.
func isStatusRetryable(status int) bool {
	return status == rateLimitExceededStatusCode || status >= minServerErrorStatusCode
}

// isTransportErrorRetryable reports whether a transport-level error should be retried.
// Network failures and timeouts are retryable; an *APIError raised from a malformed
// request (invalid argument) is not.
func isTransportErrorRetryable(err error) bool {
	if _, ok := AsAPIError(err); ok {
		return false
	}
	return true
}

// randomizedDelay returns a delay chosen randomly from [delay, 2*delay), matching the
// exponential-backoff-with-jitter algorithm described in the Apify API docs.
func randomizedDelay(delay time.Duration) time.Duration {
	if delay <= 0 {
		return delay
	}
	return delay + time.Duration(rand.Int63n(int64(delay)))
}

// sleepWithContext sleeps for d, returning false if ctx is cancelled first.
func sleepWithContext(ctx context.Context, d time.Duration) bool {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-timer.C:
		return true
	case <-ctx.Done():
		return false
	}
}

func minDuration(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

// extractPath returns the path+query portion of a URL, for error reporting.
func extractPath(url string) string {
	rest := url
	if i := strings.Index(rest, "://"); i >= 0 {
		rest = rest[i+3:]
	}
	if i := strings.IndexByte(rest, '/'); i >= 0 {
		return rest[i:]
	}
	return ""
}
