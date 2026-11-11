package apify

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
)

// Status codes and limits of the retry and compression policy.
const (
	// rateLimitExceededStatusCode is the status of an exceeded rate limit.
	rateLimitExceededStatusCode = 429
	// minServerErrorStatusCode is the lowest status that is retried as a server error.
	minServerErrorStatusCode = 500
	// minNonSuccessStatusCode is the lowest status that is not a success.
	minNonSuccessStatusCode = 300
	// backoffFactor multiplies the delay between retries after each attempt.
	backoffFactor = 2
	// minCompressRequestBytes is the smallest request body that is compressed.
	minCompressRequestBytes = 1024
)

// Content-Encoding values of compressed request bodies.
const (
	contentEncodingBrotli = "br"
	contentEncodingGzip   = "gzip"
	contentEncodingZstd   = "zstd"
)

// HTTPBackend is the replaceable transport of the client. It sends one request and returns the
// raw response. Authentication, the User-Agent header, retries and decoding are done by the client.
//
// A non-2xx status is a normal *http.Response. Return an error only for a transport failure such
// as a refused connection, a DNS error or a timeout.
//
// A backend that follows redirects must not follow them for a request that [FollowsRedirects]
// reports false for: such a request carries credentials meant for its original host only.
type HTTPBackend interface {
	// Do sends a single HTTP request and returns the response.
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPBackend is the default [HTTPBackend], backed by an *http.Client. The client applies
// the timeout of each attempt through the request context, so the backend sets none.
type DefaultHTTPBackend struct {
	client *http.Client
}

// NewDefaultHTTPBackend creates a [DefaultHTTPBackend] with a default *http.Client.
func NewDefaultHTTPBackend() *DefaultHTTPBackend {
	return &DefaultHTTPBackend{client: &http.Client{}}
}

// NewHTTPBackendWithClient creates a [DefaultHTTPBackend] around client, for example to share a
// connection pool or to set proxy and TLS options.
func NewHTTPBackendWithClient(client *http.Client) *DefaultHTTPBackend {
	return &DefaultHTTPBackend{client: client}
}

// Do implements [HTTPBackend]. It follows redirects, except for a request that [FollowsRedirects]
// reports false for, whose redirect response is returned as it is.
func (b *DefaultHTTPBackend) Do(req *http.Request) (*http.Response, error) {
	if FollowsRedirects(req) {
		return b.client.Do(req)
	}
	// The copy shares the transport, so the connection pool is shared.
	client := *b.client
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	return client.Do(req)
}

// retryConfig is the retry policy of the httpClient.
type retryConfig struct {
	// maxRetries is the number of retries, so a request is attempted up to maxRetries+1 times.
	maxRetries int
	// minDelayBetweenRetries is the first delay between attempts. It doubles on each retry.
	minDelayBetweenRetries time.Duration
	// timeoutMax caps the timeout of one attempt and the delay between attempts.
	timeoutMax time.Duration
}

// timeoutTiers are the default request timeouts. Each method uses the tier that fits its
// request: short for a metadata read or write, medium for a list, batch or trigger call, long for
// a download, upload or stream. A context deadline can shorten the tier but not extend it.
type timeoutTiers struct {
	short  time.Duration
	medium time.Duration
	long   time.Duration
}

// noRequestTimeout leaves a request bounded by the caller's context alone. The polling behind
// WaitForFinish and Call uses it, since the server holds those requests open on purpose.
const noRequestTimeout time.Duration = 0

// httpClient is the HTTP client shared by every resource client. It applies the token, the
// User-Agent, the retry policy and the timeouts to each request. It is safe for concurrent use.
type httpClient struct {
	backend     HTTPBackend
	token       string
	userAgent   string
	retry       retryConfig
	tiers       timeoutTiers
	compressors []requestCompressor
	// doNotRetryTimeouts fails an attempt that timed out instead of retrying it. Set it for a
	// request that is not idempotent, where the attempt may have taken effect already.
	doNotRetryTimeouts bool
}

// withoutTimeoutRetry returns a copy of c that does not retry a timed-out attempt. Other
// retryable failures are still retried.
func (c *httpClient) withoutTimeoutRetry() *httpClient {
	clone := *c
	clone.doNotRetryTimeouts = true
	return &clone
}

// apiResponse is the status, headers and body of a response.
type apiResponse struct {
	statusCode int
	header     http.Header
	body       []byte
}

// call sends a request with the token, the User-Agent header and the retry policy, and returns
// the first successful response or the final error. baseTimeout grows with each attempt, up to
// the client's maximum.
func (c *httpClient) call(ctx context.Context, method, url string, body []byte, contentType string, baseTimeout time.Duration) (*apiResponse, error) {
	return c.callWithHeaders(ctx, method, url, body, contentType, nil, baseTimeout)
}

// callWithHeaders is [httpClient.call] with extra request headers, set on every attempt.
func (c *httpClient) callWithHeaders(ctx context.Context, method, url string, body []byte, contentType string, extraHeaders map[string]string, baseTimeout time.Duration) (*apiResponse, error) {
	// The body is the same on every attempt, so it is compressed once.
	sendBody, contentEncoding := maybeCompressRequestBody(body, contentType, c.compressors)
	return c.doRetryLoop(ctx, method, url, sendBody, contentType, contentEncoding, extraHeaders, baseTimeout)
}

// callPreEncoded is [httpClient.call] for a body that is already compressed under contentEncoding.
// The body is sent as it is, and the client's own compression is skipped even when contentEncoding
// is empty.
func (c *httpClient) callPreEncoded(ctx context.Context, method, url string, body []byte, contentType, contentEncoding string, baseTimeout time.Duration) (*apiResponse, error) {
	return c.doRetryLoop(ctx, method, url, body, contentType, contentEncoding, nil, baseTimeout)
}

// doRetryLoop sends the request and retries retryable failures with exponential backoff.
func (c *httpClient) doRetryLoop(ctx context.Context, method, url string, sendBody []byte, contentType, contentEncoding string, extraHeaders map[string]string, baseTimeout time.Duration) (*apiResponse, error) {
	delay := c.retry.minDelayBetweenRetries
	maxAttempts := c.retry.maxRetries + 1
	path := extractPath(url)

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err := c.doAttempt(ctx, method, url, sendBody, contentType, contentEncoding, extraHeaders, c.attemptTimeout(baseTimeout, attempt))

		var retryable bool
		if err != nil {
			lastErr = err
			retryable = isTransportErrorRetryable(err) && (!c.doNotRetryTimeouts || !isTimeoutError(err))
		} else if resp.statusCode < minNonSuccessStatusCode {
			return resp, nil
		} else {
			lastErr = buildAPIError(resp.statusCode, resp.body, attempt, method, path)
			retryable = isStatusRetryable(resp.statusCode)
		}

		if !retryable || attempt == maxAttempts {
			return nil, lastErr
		}

		if !c.backoff(ctx, &delay) {
			return nil, ctx.Err()
		}
	}

	return nil, lastErr
}

// backoff sleeps before the next attempt and doubles delay. It returns false when ctx ended during
// the sleep. The sleep is random in [delay, 2*delay), as the Apify API documentation prescribes,
// and the delay is capped at the maximum request timeout.
func (c *httpClient) backoff(ctx context.Context, delay *time.Duration) bool {
	if !sleepWithContext(ctx, randomizedDelay(*delay)) {
		return false
	}
	*delay = minDuration(*delay*backoffFactor, c.retry.timeoutMax)
	return true
}

// applyAuthHeaders sets the User-Agent header and, if a token is configured, the Authorization
// header.
func (c *httpClient) applyAuthHeaders(req *http.Request) {
	c.applyAuthHeadersAs(req, "Authorization")
}

// applyAuthHeadersAs is [httpClient.applyAuthHeaders] with the token in tokenHeader.
func (c *httpClient) applyAuthHeadersAs(req *http.Request, tokenHeader string) {
	req.Header.Set("User-Agent", c.userAgent)
	if c.token != "" {
		req.Header.Set(tokenHeader, "Bearer "+c.token)
	}
}

// newRequest builds a request with the User-Agent header, the token in tokenHeader and the
// content headers. A non-empty contentEncoding says that body is compressed with it.
func (c *httpClient) newRequest(ctx context.Context, method, url string, body []byte, contentType, contentEncoding, tokenHeader string) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, err
	}
	c.applyAuthHeadersAs(req, tokenHeader)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if contentEncoding != "" {
		req.Header.Set("Content-Encoding", contentEncoding)
	}
	return req, nil
}

// doAttempt performs one round-trip with the given timeout.
func (c *httpClient) doAttempt(ctx context.Context, method, url string, body []byte, contentType, contentEncoding string, extraHeaders map[string]string, timeout time.Duration) (*apiResponse, error) {
	attemptCtx := ctx
	cancel := func() {}
	if timeout != noRequestTimeout {
		attemptCtx, cancel = context.WithTimeout(ctx, timeout)
	}
	defer cancel()

	req, err := c.newRequest(attemptCtx, method, url, body, contentType, contentEncoding, "Authorization")
	if err != nil {
		// An invalid request is a caller error, so it is not retried.
		return nil, &APIError{StatusCode: 0, Type: "invalid-request", Message: err.Error(), Attempt: 1, HTTPMethod: method, Path: extractPath(url)}
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
		// A failed read is a retryable transport error.
		return nil, err
	}

	return &apiResponse{statusCode: resp.StatusCode, header: resp.Header, body: data}, nil
}

// requestCompressor is a codec: the Content-Encoding value it produces and its compress function.
type requestCompressor struct {
	encoding string
	compress func(body []byte) ([]byte, error)
}

// RequestCompression selects which codec compresses request bodies (see [WithRequestCompression]).
type RequestCompression int

const (
	// RequestCompressionAuto uses brotli and falls back to gzip when brotli fails or does not
	// shrink the body. It is the default.
	RequestCompressionAuto RequestCompression = iota
	// RequestCompressionBrotli uses brotli only. A body it does not shrink is sent uncompressed.
	RequestCompressionBrotli
	// RequestCompressionGzip uses gzip only. A body it does not shrink is sent uncompressed.
	RequestCompressionGzip
	// RequestCompressionZstd uses zstd only, at its default level. A body it does not shrink is
	// sent uncompressed.
	RequestCompressionZstd
)

// compressorsFor returns the codecs of algo in the order they are tried.
func compressorsFor(algo RequestCompression) []requestCompressor {
	switch algo {
	case RequestCompressionBrotli:
		return []requestCompressor{{encoding: contentEncodingBrotli, compress: brotliCompress}}
	case RequestCompressionGzip:
		return []requestCompressor{{encoding: contentEncodingGzip, compress: gzipCompress}}
	case RequestCompressionZstd:
		return []requestCompressor{{encoding: contentEncodingZstd, compress: zstdCompress}}
	default:
		return []requestCompressor{
			{encoding: contentEncodingBrotli, compress: brotliCompress},
			{encoding: contentEncodingGzip, compress: gzipCompress},
		}
	}
}

// compressBody compresses body with the writer that newWriter creates. The writer must flush on
// Close.
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

// zstdEncoder compresses every zstd request body with EncodeAll, which is safe for concurrent use
// and cheaper than a new encoder for each request.
var zstdEncoder = newZstdEncoder()

func newZstdEncoder() *zstd.Encoder {
	encoder, err := zstd.NewWriter(nil)
	if err != nil {
		// Only invalid options fail, and none are passed.
		panic(err)
	}
	return encoder
}

// zstdCompress compresses body with zstd (Content-Encoding: zstd).
func zstdCompress(body []byte) ([]byte, error) {
	return zstdEncoder.EncodeAll(body, nil), nil
}

// maybeCompressRequestBody compresses body if it is at least [minCompressRequestBytes] long and
// its content type is compressible. It tries the compressors in order and uses the first that
// shrinks the body. It returns the bytes to send and the Content-Encoding value, which is empty
// when the body is sent as it is.
func maybeCompressRequestBody(body []byte, contentType string, compressors []requestCompressor) ([]byte, string) {
	if len(body) < minCompressRequestBytes {
		return body, ""
	}
	if !isCompressibleContentType(contentType) {
		return body, ""
	}
	return compressWith(body, compressors)
}

// alreadyCompressedContentTypePrefixes are the media types whose payloads are already compressed.
var alreadyCompressedContentTypePrefixes = []string{"audio/", "image/", "video/"}

// alreadyCompressedContentTypes are further media types whose payloads are already compressed.
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

// compressibleContentTypes are uncompressed media types under an already-compressed prefix.
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

// compressibleContentTypeSuffixes mark a media type as text, for example image/svg+xml.
var compressibleContentTypeSuffixes = []string{"+json", "+xml"}

// isCompressibleContentType reports whether a body of the given Content-Type is worth compressing.
// Images, audio, video and archives are not, except for raw formats such as image/bmp. An empty
// Content-Type is compressible.
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
			continue
		}
		if len(out) < len(body) {
			return out, c.encoding
		}
	}
	return body, ""
}

// attemptTimeout returns base doubled for each retry, capped at the maximum timeout. A base of
// [noRequestTimeout] stays unbounded.
func (c *httpClient) attemptTimeout(base time.Duration, attempt int) time.Duration {
	if base == noRequestTimeout {
		return noRequestTimeout
	}
	scaled := base
	for i := 1; i < attempt; i++ {
		scaled *= backoffFactor
		if scaled >= c.retry.timeoutMax {
			return c.retry.timeoutMax
		}
	}
	return minDuration(scaled, c.retry.timeoutMax)
}

// isStatusRetryable reports whether a status is retried: 429 and 5xx.
func isStatusRetryable(status int) bool {
	return status == rateLimitExceededStatusCode || status >= minServerErrorStatusCode
}

// isTransportErrorRetryable reports whether a transport error is retried. An *APIError, which
// comes from an invalid request, is not.
func isTransportErrorRetryable(err error) bool {
	if _, ok := AsAPIError(err); ok {
		return false
	}
	return true
}

// isTimeoutError reports whether err means a request attempt ran out of time: the attempt's own
// context deadline, or a timeout reported by the transport.
func isTimeoutError(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

// randomizedDelay returns a random delay in [delay, 2*delay).
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
