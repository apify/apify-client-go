package apify

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// serverAuthorizationHeader carries the API token of a request to an Actor's server. The
// platform consumes it before the request reaches the Actor, so the standard Authorization
// header stays free for the Actor's own authentication.
const serverAuthorizationHeader = "X-Apify-Authorization"

// Status codes a request to an Actor's server is retried on. The platform answers with them when
// it cannot pass a request to a server run, so a retry usually does not run the Actor's code
// again. An Actor that returns one itself gets the request again.
const (
	serverBusyStatusCode        = rateLimitExceededStatusCode
	serverUnavailableStatusCode = 503
)

// noRedirectsKey marks a request context whose request must not follow redirects.
type noRedirectsKey struct{}

// withoutRedirects returns a context that tells the HTTP backend not to follow redirects. The
// request to an Actor's server carries the API token, which a redirect target picked by the
// Actor's code must never receive.
func withoutRedirects(ctx context.Context) context.Context {
	return context.WithValue(ctx, noRedirectsKey{}, true)
}

// FollowsRedirects reports whether the HTTP backend may follow redirects for req. It is false
// for the requests that carry credentials meant for one host only, such as those sent by
// [ActorServerClient]: such a request must get a redirect response back as it is. A custom
// [HTTPBackend] that follows redirects has to check it, and the [DefaultHTTPBackend] does.
func FollowsRedirects(req *http.Request) bool {
	return req.Context().Value(noRedirectsKey{}) == nil
}

// serverCall describes one request to an Actor's server.
type serverCall struct {
	method string
	url    string
	body   []byte
	// contentType is the Content-Type of the body, set before the caller's headers.
	contentType string
	// headers are the caller's headers. They are applied last, so they win over every header
	// the client sets, the Authorization header included.
	headers http.Header
	// compressors compress the body when it is large enough; empty sends it as it is.
	compressors []requestCompressor
	// timeout is the base timeout of an attempt. A stream bounds only the wait for the response
	// headers.
	timeout time.Duration
	// stream returns the response body unread.
	stream bool
}

// serverResponse is the response of an Actor's server, whatever its status.
type serverResponse struct {
	statusCode int
	header     http.Header
	// body is the buffered body, nil for a stream.
	body []byte
	// stream is the unread body of a streamed response. The caller closes it.
	stream io.ReadCloser
}

// callServer sends call to an Actor's server and returns its response whatever the status.
//
// It differs from [httpClient.call] where the server is not the API: an error status is
// returned instead of raised, only 429 and 503 are retried, a transport error (a timeout
// included) is never retried since the Actor may have handled the request already, the token
// goes in X-Apify-Authorization, and redirects are not followed.
func (c *httpClient) callServer(ctx context.Context, call serverCall) (*serverResponse, error) {
	sendBody, contentEncoding := call.body, ""
	if call.headers.Get("Content-Encoding") == "" {
		sendBody, contentEncoding = maybeCompressRequestBody(call.body, call.contentType, call.compressors)
	}

	delay := c.retry.minDelayBetweenRetries
	maxAttempts := c.retry.maxRetries + 1
	for attempt := 1; ; attempt++ {
		resp, err := c.serverAttempt(withoutRedirects(ctx), call, sendBody, contentEncoding, c.attemptTimeout(call.timeout, attempt))
		if err != nil {
			return nil, err
		}
		retryable := resp.statusCode == serverBusyStatusCode || resp.statusCode == serverUnavailableStatusCode
		if !retryable || attempt == maxAttempts {
			return resp, nil
		}
		if resp.stream != nil {
			_ = resp.stream.Close()
		}
		if !sleepWithContext(ctx, randomizedDelay(delay)) {
			return nil, ctx.Err()
		}
		delay = minDuration(delay*backoffFactor, c.retry.timeoutMax)
	}
}

// serverAttempt performs a single round-trip of call.
func (c *httpClient) serverAttempt(ctx context.Context, call serverCall, body []byte, contentEncoding string, timeout time.Duration) (*serverResponse, error) {
	attemptCtx, cancel := context.WithCancelCause(ctx)
	// A buffered response is bounded as a whole, a stream only until its headers arrive.
	var timer *time.Timer
	if timeout != noRequestTimeout {
		timer = time.AfterFunc(timeout, func() { cancel(context.DeadlineExceeded) })
	}
	stopTimer := func() {
		if timer != nil {
			timer.Stop()
		}
	}

	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(attemptCtx, call.method, call.url, reader)
	if err != nil {
		stopTimer()
		cancel(nil)
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	if c.token != "" {
		req.Header.Set(serverAuthorizationHeader, "Bearer "+c.token)
	}
	if call.contentType != "" {
		req.Header.Set("Content-Type", call.contentType)
	}
	if contentEncoding != "" {
		req.Header.Set("Content-Encoding", contentEncoding)
	}
	for name, values := range call.headers {
		req.Header[http.CanonicalHeaderKey(name)] = values
	}

	resp, err := c.backend.Do(req)
	if err != nil {
		stopTimer()
		err = attemptError(attemptCtx, err)
		cancel(nil)
		return nil, err
	}

	if call.stream {
		stopTimer()
		return &serverResponse{statusCode: resp.StatusCode, header: resp.Header, stream: &cancelOnClose{ReadCloser: resp.Body, cancel: func() { cancel(nil) }}}, nil
	}
	defer cancel(nil)
	defer stopTimer()
	defer func() { _ = resp.Body.Close() }()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, attemptError(attemptCtx, err)
	}
	return &serverResponse{statusCode: resp.StatusCode, header: resp.Header, body: data}, nil
}

// attemptError reports a timeout of the attempt as [context.DeadlineExceeded], the error a
// request bounded by [context.WithTimeout] gives, and any other failure as it is.
func attemptError(attemptCtx context.Context, err error) error {
	if errors.Is(context.Cause(attemptCtx), context.DeadlineExceeded) && !errors.Is(err, context.DeadlineExceeded) {
		return fmt.Errorf("%w: %w", context.DeadlineExceeded, err)
	}
	return err
}

// cancelOnClose releases the context of a streamed attempt once its body is closed.
type cancelOnClose struct {
	io.ReadCloser
	cancel func()
}

func (b *cancelOnClose) Close() error {
	err := b.ReadCloser.Close()
	b.cancel()
	return err
}
