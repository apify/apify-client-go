package apify

import (
	"context"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// RunResurrectOptions configures [RunClient.Resurrect].
type RunResurrectOptions struct {
	// Build is the tag or number of the build to resurrect with.
	Build *string
	// MemoryMbytes is the memory in megabytes to allocate.
	MemoryMbytes *int64
	// TimeoutSecs is the run timeout in seconds.
	TimeoutSecs *int64
	// MaxItems is the maximum number of dataset items to charge (pay-per-result Actors).
	MaxItems *int64
	// MaxTotalChargeUsd is the maximum total charge in USD (pay-per-event Actors).
	MaxTotalChargeUsd *float64
	// RestartOnError, if true, restarts the run if it fails.
	RestartOnError *bool
}

func (o RunResurrectOptions) apply(q *QueryParams) {
	q.AddString("build", o.Build).
		AddInt("memory", o.MemoryMbytes).
		AddInt("timeout", o.TimeoutSecs).
		AddInt("maxItems", o.MaxItems).
		AddFloat("maxTotalChargeUsd", o.MaxTotalChargeUsd).
		AddBool("restartOnError", o.RestartOnError)
}

// RunChargeOptions configures [RunClient.Charge].
type RunChargeOptions struct {
	// EventName is the name of the event to charge for. Required.
	EventName string
	// Count is the number of times to charge the event (defaults to 1).
	Count *int64
	// IdempotencyKey deduplicates the charge across retries. If empty, one is auto-generated
	// as "{runId}-{eventName}-{timestampMillis}-{random}", matching the reference client.
	IdempotencyKey string
}

// RunClient is a client for a specific Actor run.
//
// It provides CRUD methods plus convenience helpers (abort, metamorph, reboot, resurrect,
// charge, wait-for-finish) and accessors for the run's default storages and log.
type RunClient struct {
	root *ApifyClient
	ctx  *resourceContext
	id   string
}

func newRunClient(root *ApifyClient, hc *httpClient, baseURL, resourcePath, id string) *RunClient {
	return &RunClient{
		root: root,
		ctx:  newSingleContext(hc, baseURL, resourcePath, id),
		id:   id,
	}
}

// setStatusParam pins a `status` query parameter inherited by all calls on this client.
// Used by ActorClient.LastRun / TaskClient.LastRun to filter the "last" run by status.
func (c *RunClient) setStatusParam(status string) {
	c.ctx.baseParams.addRaw("status", status)
}

// Get fetches the run object. The bool reports whether it exists.
func (c *RunClient) Get(ctx context.Context) (ActorRun, bool, error) {
	return c.GetWithWait(ctx, nil)
}

// GetWithWait fetches the run, optionally asking the API to wait up to waitForFinishSecs
// seconds (max 60) for the run to reach a terminal state before responding. Pass nil for an
// immediate fetch. Mirrors the reference client's get({ waitForFinish }).
func (c *RunClient) GetWithWait(ctx context.Context, waitForFinishSecs *int64) (ActorRun, bool, error) {
	params := NewQueryParams()
	params.AddInt("waitForFinish", waitForFinishSecs)
	return getResource[ActorRun](ctx, c.ctx, "", params)
}

// Update updates the run with the given fields and returns the updated object.
func (c *RunClient) Update(ctx context.Context, newFields any) (ActorRun, error) {
	return updateResource[ActorRun](ctx, c.ctx, "", newFields)
}

// Delete deletes the run.
func (c *RunClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// Abort aborts the run. If gracefully points to true, the run is sent a signal so it can
// finish the current request before terminating; if false it is aborted immediately. Pass
// nil to omit the parameter entirely and let the server apply its default (immediate abort),
// matching the reference client's optional `gracefully` option.
func (c *RunClient) Abort(ctx context.Context, gracefully *bool) (ActorRun, error) {
	params := NewQueryParams()
	params.AddBool("gracefully", gracefully)
	return postWithBody[ActorRun](ctx, c.ctx, "abort", params, nil, "")
}

// MetamorphOptions configures [RunClient.Metamorph].
type MetamorphOptions struct {
	// Build optionally pins the target Actor's build (empty for default).
	Build string
	// ContentType is the content type of the input body. Defaults to application/json.
	ContentType string
}

// Metamorph transforms the run into a run of another Actor with a new input.
//
// targetActorID is the Actor to metamorph into. input is the new input (nil for none).
func (c *RunClient) Metamorph(ctx context.Context, targetActorID string, input any, options MetamorphOptions) (ActorRun, error) {
	params := NewQueryParams()
	params.AddString("targetActorId", &targetActorID)
	if options.Build != "" {
		params.AddString("build", &options.Build)
	}
	body, err := marshalInput(input)
	if err != nil {
		return ActorRun{}, err
	}
	contentType := options.ContentType
	if contentType == "" {
		contentType = contentTypeJSON
	}
	return postWithBody[ActorRun](ctx, c.ctx, "metamorph", params, body, contentType)
}

// Reboot reboots the run (restarts its container while keeping the same run).
func (c *RunClient) Reboot(ctx context.Context) (ActorRun, error) {
	return postWithBody[ActorRun](ctx, c.ctx, "reboot", NewQueryParams(), nil, "")
}

// Resurrect resurrects a finished run, starting it again from the beginning.
func (c *RunClient) Resurrect(ctx context.Context, options RunResurrectOptions) (ActorRun, error) {
	params := NewQueryParams()
	options.apply(params)
	return postWithBody[ActorRun](ctx, c.ctx, "resurrect", params, nil, "")
}

// Charge charges for a pay-per-event Actor run: it records occurrences of a named event.
// Only meaningful for runs of pay-per-event Actors.
//
// An idempotency key is always sent (auto-generated if not provided), so a charge that is
// retried by the transport is applied at most once, matching the reference client.
func (c *RunClient) Charge(ctx context.Context, options RunChargeOptions) error {
	count := int64(1)
	if options.Count != nil {
		count = *options.Count
	}
	idempotencyKey := options.IdempotencyKey
	if idempotencyKey == "" {
		idempotencyKey = c.generateIdempotencyKey(options.EventName)
	}
	body := mustMarshal(map[string]any{"eventName": options.EventName, "count": count})
	url := c.ctx.subURL("charge")
	headers := map[string]string{chargeIdempotencyHeader: idempotencyKey}
	_, err := c.ctx.http.callWithHeaders(ctx, "POST", url, body, contentTypeJSON, headers, defaultRequestTimeout)
	return err
}

// chargeIdempotencyHeader is the header the API uses to deduplicate charge requests.
const chargeIdempotencyHeader = "idempotency-key"

// generateIdempotencyKey builds a per-charge idempotency key of the form
// "{runId}-{eventName}-{timestampMillis}-{random}", matching the reference client. It need
// not be cryptographically secure — only unique enough to avoid collisions within a request.
func (c *RunClient) generateIdempotencyKey(eventName string) string {
	return c.id + "-" + eventName + "-" +
		strconv.FormatInt(time.Now().UnixMilli(), 10) + "-" +
		strconv.FormatInt(rand.Int63n(1_000_000), 10)
}

// WaitForFinish polls until the run reaches a terminal state or waitSecs elapses (nil waits
// indefinitely). It returns the latest run.
func (c *RunClient) WaitForFinish(ctx context.Context, waitSecs *int64) (ActorRun, error) {
	return waitForFinish[ActorRun](ctx, c.ctx, waitSecs, "run", func(r *ActorRun) bool { return r.IsTerminal() })
}

// Dataset returns a client for this run's default dataset.
func (c *RunClient) Dataset() *DatasetClient {
	return newDatasetNestedClient(c.ctx.http, c.ctx.subURL(""), "dataset")
}

// KeyValueStore returns a client for this run's default key-value store.
func (c *RunClient) KeyValueStore() *KeyValueStoreClient {
	return newKeyValueStoreNestedClient(c.ctx.http, c.ctx.subURL(""), "key-value-store")
}

// RequestQueue returns a client for this run's default request queue.
func (c *RunClient) RequestQueue() *RequestQueueClient {
	return newRequestQueueNestedClient(c.ctx.http, c.ctx.subURL(""), "request-queue")
}

// Log returns a client for accessing this run's log.
func (c *RunClient) Log() *LogClient {
	return newNestedLogClient(c.ctx.http, c.ctx.subURL(""))
}

// GetStreamedLog opens a live stream of this run's raw log, for convenient log redirection.
//
// It is a convenience wrapper over Log().StreamWithOptions with raw=true (matching the
// reference client's getStreamedLog, which streams raw log content). The caller must close
// the returned reader.
func (c *RunClient) GetStreamedLog(ctx context.Context) (io.ReadCloser, error) {
	raw := true
	return c.Log().StreamWithOptions(ctx, LogOptions{Raw: &raw})
}
