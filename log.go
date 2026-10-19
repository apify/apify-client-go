package apify

import (
	"context"
	"io"
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
//
// A 404 resolves to "absent" only when this client names the log by its own id
// (client.Log(id)): a chained client without one, such as run.Log() or build.Log(), throws
// instead, since the run or build itself may be what is missing.
func (c *LogClient) Get(ctx context.Context) (string, bool, error) {
	return c.GetWithOptions(ctx, LogOptions{})
}

// GetWithOptions fetches the log with explicit options (raw, download). See [LogClient.Get]
// for how a 404 is handled.
func (c *LogClient) GetWithOptions(ctx context.Context, options LogOptions) (string, bool, error) {
	body, ok, err := c.GetBufferWithOptions(ctx, options)
	return string(body), ok, err
}

// GetBuffer fetches the entire log as raw bytes, or (nil, false, nil) if the log does not exist.
//
// It is the byte-oriented counterpart of [LogClient.Get]: both fetch the whole log in one
// response (for incremental reading, use [LogClient.Stream]), and GetBuffer only changes the
// return type to the undecoded bytes, e.g. for writing the log straight to a file. See
// [LogClient.Get] for how a 404 is handled.
func (c *LogClient) GetBuffer(ctx context.Context) ([]byte, bool, error) {
	return c.GetBufferWithOptions(ctx, LogOptions{})
}

// GetBufferWithOptions fetches the log as raw bytes with explicit options (raw, download). See
// [LogClient.GetBuffer] and [LogClient.Get] for how a 404 is handled.
func (c *LogClient) GetBufferWithOptions(ctx context.Context, options LogOptions) ([]byte, bool, error) {
	params := NewQueryParams()
	options.apply(params)
	resp, err := getRawRequired(ctx, c.ctx, "", params, c.ctx.longTimeout())
	if err != nil {
		if c.ctx.hasOwnID && isNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return resp.body, true, nil
}

// Stream opens a live, streaming connection to the log and returns a reader over the log
// bytes, or (nil, nil) if the log does not exist. The caller is responsible for closing the
// returned [io.ReadCloser].
//
// Unlike [LogClient.Get], this bypasses the buffered/retrying transport so the log can be
// followed in real time as the run produces it (the `stream=1` query parameter). Because the
// response is consumed incrementally, it is not retried; transient failures surface to the
// caller. This mirrors the reference clients' streamed-log behaviour used for log redirection.
//
// A 404 resolves to (nil, nil) only when this client names the log by its own id
// (client.Log(id)): a chained client without one, such as run.Log() or build.Log(), returns an
// error instead, since the run or build itself may be what is missing.
func (c *LogClient) Stream(ctx context.Context) (io.ReadCloser, error) {
	return c.StreamWithOptions(ctx, LogOptions{})
}

// StreamWithOptions opens a live log stream with explicit options (raw, download). See
// [LogClient.Stream] for the streaming semantics and how a 404 is handled.
func (c *LogClient) StreamWithOptions(ctx context.Context, options LogOptions) (io.ReadCloser, error) {
	params := NewQueryParams()
	streamOn := true
	params.AddBool("stream", &streamOn)
	options.apply(params)
	stream, _, err := streamRaw(ctx, c.ctx, "", params, c.ctx.hasOwnID)
	return stream, err
}
