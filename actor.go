package apify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// ActorStartOptions configures starting an Actor or task run
// ([ActorClient.Start]/[ActorClient.Call] and the task equivalents).
type ActorStartOptions struct {
	// Build is the tag or number of the build to run (e.g. "latest", "0.1.2").
	Build *string
	// MemoryMbytes is the memory in megabytes allocated for the run.
	MemoryMbytes *int64
	// TimeoutSecs is the timeout for the run in seconds (0 means no timeout).
	TimeoutSecs *int64
	// WaitForFinish is the maximum seconds to wait server-side for the run to finish (max 60).
	WaitForFinish *int64
	// MaxItems is the maximum number of dataset items to charge (pay-per-result Actors).
	MaxItems *int64
	// MaxTotalChargeUsd is the maximum total charge in USD (pay-per-event Actors).
	MaxTotalChargeUsd *float64
	// ContentType is the content type of the input body. Defaults to application/json.
	ContentType *string
	// RestartOnError, if true, restarts the run if it fails.
	RestartOnError *bool
	// ForcePermissionLevel overrides the Actor's permission level for this run.
	ForcePermissionLevel *string
	// Webhooks are ad-hoc webhooks to attach to this run. They are serialized to base64-encoded
	// JSON as the `webhooks` query parameter.
	Webhooks []any
	// WaitForResources retries the start every 10 seconds while the API rejects it for lack of
	// resources (an *APIError of type "actor-memory-limit-exceeded" or
	// "concurrent-runs-limit-exceeded"). Nil does not retry. Use WaitForResourcesIndefinitely to
	// retry without a time bound, or a number of seconds after which the last error is returned.
	// Any other error is returned at once.
	//
	// A run that requests more memory than the account's limit is rejected with
	// actor-memory-limit-exceeded and never starts, so WaitForResourcesIndefinitely retries it forever.
	WaitForResources *int64
	// Log configures the redirection of the run's log that [ActorClient.Call] and
	// [ActorClient.CallRaw] do while they wait (see [StreamedLog]). Nil redirects it to os.Stderr
	// with the default prefix; set Log.Disabled to turn the redirection off. Start and StartRaw
	// ignore it.
	Log *StreamedLogOptions
}

func (o ActorStartOptions) apply(q *QueryParams) {
	q.AddString("build", o.Build).
		AddInt("memory", o.MemoryMbytes).
		AddInt("timeout", o.TimeoutSecs).
		AddInt("waitForFinish", o.WaitForFinish).
		AddInt("maxItems", o.MaxItems).
		AddFloat("maxTotalChargeUsd", o.MaxTotalChargeUsd).
		AddBool("restartOnError", o.RestartOnError).
		AddString("forcePermissionLevel", o.ForcePermissionLevel)
	if encoded := encodeWebhooks(o.Webhooks); encoded != nil {
		q.AddString("webhooks", encoded)
	}
}

// encodeWebhooks encodes an ad-hoc webhooks slice as base64-encoded JSON, as the API's
// `webhooks` query parameter requires. Returns nil for a nil slice. Shared by the Actor and
// task start options (DRY).
func encodeWebhooks(webhooks []any) *string {
	if webhooks == nil {
		return nil
	}
	data, err := json.Marshal(webhooks)
	if err != nil {
		return nil
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return &encoded
}

// ActorBuildOptions configures [ActorClient.Build].
type ActorBuildOptions struct {
	// BetaPackages, if true, uses beta versions of Apify packages.
	BetaPackages *bool
	// Tag is the tag to apply to the build (e.g. "latest").
	Tag *string
	// UseCache, if set, controls whether to use the Docker build cache (default true).
	UseCache *bool
	// WaitForFinish is the maximum seconds to wait server-side for the build (max 60).
	WaitForFinish *int64
}

// ActorClient is a client for a specific Actor.
//
// It provides CRUD methods plus convenience helpers to start/call the Actor, build it, and
// access its runs, builds, versions and webhooks.
type ActorClient struct {
	root    *ApifyClient
	ctx     *resourceContext
	baseURL string
	id      string

	serverOnce sync.Once
	server     *ActorServerClient
}

func newActorClient(root *ApifyClient, hc *httpClient, baseURL, id string) *ActorClient {
	return &ActorClient{
		root:    root,
		ctx:     newSingleContext(hc, baseURL, "actors", id),
		baseURL: baseURL,
		id:      id,
	}
}

// ID returns the Actor's ID (or username~name) as provided.
func (c *ActorClient) ID() string { return c.id }

// Get fetches the Actor object. The bool reports whether the Actor exists.
func (c *ActorClient) Get(ctx context.Context) (Actor, bool, error) {
	return getResource[Actor](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the Actor with the given fields and returns the updated object.
func (c *ActorClient) Update(ctx context.Context, newFields any) (Actor, error) {
	return updateResource[Actor](ctx, c.ctx, "", newFields)
}

// Delete deletes the Actor.
func (c *ActorClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// Start starts the Actor and returns immediately with the created run.
//
// input is any JSON-serializable value (or nil for no input). To send a non-JSON input (e.g. a
// ZIP archive) as raw bytes instead, use [ActorClient.StartRaw].
func (c *ActorClient) Start(ctx context.Context, input any, options ActorStartOptions) (ActorRun, error) {
	body, err := marshalInput(input)
	if err != nil {
		return ActorRun{}, err
	}
	return c.startWithBody(ctx, body, contentTypeJSON, options, c.startTimeout(options))
}

// StartRaw starts the Actor with a raw request body, for a non-JSON input such as a ZIP archive
// with options.ContentType set to "application/zip". For an object or array input, use
// [ActorClient.Start].
func (c *ActorClient) StartRaw(ctx context.Context, input []byte, options ActorStartOptions) (ActorRun, error) {
	return c.startWithBody(ctx, input, "application/octet-stream", options, c.startTimeout(options))
}

// startTimeout is the timeout of a start request: the "medium" tier, extended to cover the time
// the API may hold the response for options.WaitForFinish.
func (c *ActorClient) startTimeout(options ActorStartOptions) time.Duration {
	return extendForServerWait(c.ctx.mediumTimeout(), options.WaitForFinish)
}

// startWithBody starts the Actor with body, which has defaultContentType unless
// options.ContentType is set. A start request that times out is not retried, since it may have
// created a billed run already.
func (c *ActorClient) startWithBody(ctx context.Context, body []byte, defaultContentType string, options ActorStartOptions, timeout time.Duration) (ActorRun, error) {
	params := NewQueryParams()
	options.apply(params)
	contentType := defaultContentType
	if options.ContentType != nil && *options.ContentType != "" {
		contentType = *options.ContentType
	}
	return startWaitingForResources(ctx, options.WaitForResources, func() (ActorRun, error) {
		return postWithBody[ActorRun](ctx, c.ctx.withoutTimeoutRetry(), "runs", params, body, contentType, timeout)
	})
}

// Call starts the Actor and waits for it to finish. waitSecs bounds the wait, and nil waits
// indefinitely. It returns the finished run, or the running run if the wait ran out.
//
// While it waits, it redirects the run's log to os.Stderr or options.Log.Output, prefixed with the
// Actor name and the run ID (see [StreamedLog]). Set options.Log.Disabled to turn that off.
//
// The start request and the polls have no timeout of their own, so ctx bounds the whole call.
func (c *ActorClient) Call(ctx context.Context, input any, options ActorStartOptions, waitSecs *int64) (ActorRun, error) {
	body, err := marshalInput(input)
	if err != nil {
		return ActorRun{}, err
	}
	return c.callWithBody(ctx, body, contentTypeJSON, options, waitSecs)
}

// CallRaw starts the Actor with a raw request body ([ActorClient.StartRaw]) and waits for it to
// finish, exactly like [ActorClient.Call] but for a non-JSON input.
func (c *ActorClient) CallRaw(ctx context.Context, input []byte, options ActorStartOptions, waitSecs *int64) (ActorRun, error) {
	return c.callWithBody(ctx, input, "application/octet-stream", options, waitSecs)
}

// callWithBody is the shared implementation of Call and CallRaw.
func (c *ActorClient) callWithBody(ctx context.Context, body []byte, defaultContentType string, options ActorStartOptions, waitSecs *int64) (ActorRun, error) {
	run, err := c.startWithBody(ctx, body, defaultContentType, options, noRequestTimeout)
	if err != nil {
		return ActorRun{}, err
	}
	// Use the root client's run client so polling targets the canonical run route.
	runClient := c.root.Run(run.ID)
	streamed := startLogRedirection(ctx, runClient, options.Log)
	finished, err := runClient.WaitForFinish(ctx, waitSecs)
	stopLogRedirection(ctx, streamed, options.Log)
	return finished, err
}

// Build builds the given version of the Actor and returns the created build.
func (c *ActorClient) Build(ctx context.Context, versionNumber string, options ActorBuildOptions) (Build, error) {
	params := NewQueryParams()
	params.AddString("version", &versionNumber).
		AddBool("betaPackages", options.BetaPackages).
		AddString("tag", options.Tag).
		AddBool("useCache", options.UseCache).
		AddInt("waitForFinish", options.WaitForFinish)
	timeout := extendForServerWait(c.ctx.mediumTimeout(), options.WaitForFinish)
	// A build that timed out is not retried: it may already have been created.
	return postWithBody[Build](ctx, c.ctx.withoutTimeoutRetry(), "builds", params, nil, contentTypeJSON, timeout)
}

// DefaultBuild resolves the default build of the Actor and returns a client for it. waitForFinish
// is how many seconds the API may wait for the build to finish before responding.
func (c *ActorClient) DefaultBuild(ctx context.Context, waitForFinish *int64) (*BuildClient, error) {
	if c.ctx.idErr != nil {
		return nil, c.ctx.idErr
	}
	params := NewQueryParams()
	params.AddInt("waitForFinish", waitForFinish)
	url := params.applyToURL(c.ctx.subURL("builds/default"))
	timeout := extendForServerWait(c.ctx.shortTimeout(), waitForFinish)
	resp, err := c.ctx.http.call(ctx, http.MethodGet, url, nil, "", timeout)
	if err != nil {
		return nil, err
	}
	build, err := parseDataEnvelope[Build](resp.body)
	if err != nil {
		return nil, err
	}
	return newBuildClient(c.ctx.http, c.baseURL, build.ID), nil
}

// ValidateInput validates input against the input schema of the build tagged "latest" and returns
// the validation result of the API. To validate against another build, use
// [ActorClient.ValidateInputForBuild].
func (c *ActorClient) ValidateInput(ctx context.Context, input any) (json.RawMessage, error) {
	return c.ValidateInputForBuild(ctx, input, "")
}

// ValidateInputForBuild validates input against the input schema of a build, given by its tag or
// number such as "latest" or "0.1.2", and returns the validation result of the API. An empty
// build is the same as [ActorClient.ValidateInput].
func (c *ActorClient) ValidateInputForBuild(ctx context.Context, input any, build string) (json.RawMessage, error) {
	if c.ctx.idErr != nil {
		return nil, c.ctx.idErr
	}
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	params := NewQueryParams()
	if build != "" {
		params.AddString("build", &build)
	}
	url := params.applyToURL(c.ctx.subURL("validate-input"))
	resp, err := c.ctx.http.call(ctx, http.MethodPost, url, body, contentTypeJSON, c.ctx.shortTimeout())
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

// LastRun returns a client for the last run of this Actor, optionally filtered by status
// (e.g. "SUCCEEDED"). Pass an empty status for no filter.
//
// To also filter by run origin, use LastRunWithOptions.
func (c *ActorClient) LastRun(status string) *RunClient {
	return c.LastRunWithOptions(LastRunOptions{Status: status})
}

// LastRunWithOptions returns a client for the last run of this Actor, optionally filtered by status
// and/or origin. See LastRunOptions.
func (c *ActorClient) LastRunWithOptions(options LastRunOptions) *RunClient {
	client := newRunClient(c.root, c.ctx.http, c.ctx.subURL(""), "runs", "last")
	client.setLastRunParams(options)
	return client
}

// Builds returns a client for this Actor's build collection.
func (c *ActorClient) Builds() *BuildCollectionClient {
	return newBuildCollectionClientWithBase(c.ctx.http, c.ctx.subURL(""), "builds")
}

// Runs returns a client for this Actor's run collection.
func (c *ActorClient) Runs() *RunCollectionClient {
	return newRunCollectionClient(c.ctx.http, c.ctx.subURL(""), "runs")
}

// Version returns a client for a specific version of this Actor.
func (c *ActorClient) Version(versionNumber string) *ActorVersionClient {
	return newActorVersionClient(c.ctx.http, c.ctx.subURL(""), versionNumber)
}

// Versions returns a client for this Actor's version collection.
func (c *ActorClient) Versions() *ActorVersionCollectionClient {
	return newActorVersionCollectionClient(c.ctx.http, c.ctx.subURL(""))
}

// Webhooks returns a client for this Actor's webhook collection.
func (c *ActorClient) Webhooks() *WebhookCollectionClient {
	return newWebhookCollectionClientWithBase(c.ctx.http, c.ctx.subURL(""))
}

// marshalInput serializes a non-nil input to JSON, returning nil bytes for a nil input.
func marshalInput(input any) ([]byte, error) {
	if input == nil {
		return nil, nil
	}
	return json.Marshal(input)
}
