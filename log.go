package apify

import (
	"context"
	"io"
	"net/http"
)

// LogClient is a client for accessing the log of an Actor build or run
// (/v2/logs/{buildOrRunId}, or the run/build-nested .../log).
type LogClient struct {
	ctx *resourceContext
}

func newLogClient(hc *httpClient, baseURL, resourcePath, id string) *LogClient {
	return &LogClient{ctx: newSingleContext(hc, baseURL, resourcePath, id)}
}

// newNestedLogClient creates a log client for a run's or build's nested log endpoint
// (e.g. /v2/actor-runs/{runId}/log).
func newNestedLogClient(hc *httpClient, base string) *LogClient {
	return &LogClient{ctx: newCollectionContext(hc, base, "log")}
}

// LogOptions configures log retrieval/streaming.
type LogOptions struct {
	// Raw, if true, returns the unprocessed log content (no platform post-processing).
	Raw *bool
	// Download, if true, sets Content-Disposition so the log is served as a download.
	Download *bool
}

func (o LogOptions) apply(q *QueryParams) {
	q.AddBool("raw", o.Raw).AddBool("download", o.Download)
}

// Get fetches the entire log as text, or ("", false, nil) if the log does not exist.
func (c *LogClient) Get(ctx context.Context) (string, bool, error) {
	return c.GetWithOptions(ctx, LogOptions{})
}

// GetWithOptions fetches the log with explicit options (raw, download).
func (c *LogClient) GetWithOptions(ctx context.Context, options LogOptions) (string, bool, error) {
	params := NewQueryParams()
	options.apply(params)
	resp, err := getRaw(ctx, c.ctx, "", params)
	if err != nil {
		return "", false, err
	}
	if resp == nil {
		return "", false, nil
	}
	return string(resp.body), true, nil
}

// Stream opens a live, streaming connection to the log and returns a reader over the log
// bytes. The caller is responsible for closing the returned [io.ReadCloser].
//
// Unlike [LogClient.Get], this bypasses the buffered/retrying transport so the log can be
// followed in real time as the run produces it (the `stream=1` query parameter). Because the
// response is consumed incrementally, it is not retried; transient failures surface to the
// caller. This mirrors the reference clients' streamed-log behaviour used for log redirection.
func (c *LogClient) Stream(ctx context.Context) (io.ReadCloser, error) {
	return c.StreamWithOptions(ctx, LogOptions{})
}

// StreamWithOptions opens a live log stream with explicit options (raw, download). See
// [LogClient.Stream] for the streaming semantics.
func (c *LogClient) StreamWithOptions(ctx context.Context, options LogOptions) (io.ReadCloser, error) {
	params := NewQueryParams()
	streamOn := true
	params.AddBool("stream", &streamOn)
	options.apply(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL(""))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// Reuse the shared auth/User-Agent header setup so the streaming path can never drift from
	// the buffered request path.
	c.ctx.http.applyAuthHeaders(req)

	resp, err := c.ctx.http.backend.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= maxSuccessStatusCode {
		// Consume and classify the error body, then close.
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, buildAPIError(resp.StatusCode, body, 1, http.MethodGet, extractPath(url))
	}
	return resp.Body, nil
}
