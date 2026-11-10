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

// GetBuffer fetches the whole log as bytes, or returns (nil, false, nil) if there is none. It is
// [LogClient.Get] without the decoding, for example to write the log to a file.
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

// Stream returns a reader over the live log, or (nil, nil) if there is no log. The caller closes
// it. The request is not retried.
//
// A 404 gives (nil, nil) only for client.Log(id). For run.Log() and build.Log() it is an error,
// since the run or build itself may be missing.
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
