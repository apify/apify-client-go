package apify

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	// timeout is the overall per-request timeout budget. Each attempt's timeout grows
	// but is capped here.
	timeout time.Duration
}

// httpClient is the orchestrating HTTP client shared by every resource client. It owns
// the backend, the optional API token, the User-Agent, and the retry/timeout policy, and
// applies them to every request. It is safe for concurrent use and cheap to copy.
type httpClient struct {
	backend   HTTPBackend
	token     string
	userAgent string
	retry     retryConfig
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

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err := c.doAttempt(ctx, method, url, body, contentType, extraHeaders, c.attemptTimeout(baseTimeout, attempt))

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

		// Sleep with randomized exponential backoff before the next attempt. The backoff
		// doubles each retry (matching the reference client's async-retry factor of 2) and
		// is capped at the overall request timeout.
		if !sleepWithContext(ctx, randomizedDelay(delay)) {
			return nil, ctx.Err()
		}
		delay = minDuration(delay*backoffFactor, c.retry.timeout)
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
func (c *httpClient) doAttempt(ctx context.Context, method, url string, body []byte, contentType string, extraHeaders map[string]string, timeout time.Duration) (*apiResponse, error) {
	attemptCtx, cancel := context.WithTimeout(ctx, timeout)
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

// attemptTimeout returns min(overall, base * 2^(attempt-1)).
//
// The first attempt uses the per-endpoint base timeout; each retry doubles it so a
// slow-but-progressing connection gets more time, while never exceeding the overall budget.
func (c *httpClient) attemptTimeout(base time.Duration, attempt int) time.Duration {
	scaled := base
	for i := 1; i < attempt; i++ {
		scaled *= 2
		if scaled >= c.retry.timeout {
			return c.retry.timeout
		}
	}
	return minDuration(scaled, c.retry.timeout)
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
