package apify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
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
	// Webhooks are ad-hoc webhooks to attach to this run. They are serialized to
	// base64-encoded JSON as the `webhooks` query parameter, matching the reference clients.
	Webhooks []any
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

func (o ActorStartOptions) contentTypeOrDefault() string {
	if o.ContentType != nil && *o.ContentType != "" {
		return *o.ContentType
	}
	return contentTypeJSON
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
// input is any JSON-serializable value (or nil for no input).
func (c *ActorClient) Start(ctx context.Context, input any, options ActorStartOptions) (ActorRun, error) {
	params := NewQueryParams()
	options.apply(params)
	body, err := marshalInput(input)
	if err != nil {
		return ActorRun{}, err
	}
	return postWithBody[ActorRun](ctx, c.ctx, "runs", params, body, options.contentTypeOrDefault())
}

// Call starts the Actor and waits (client-side polling) for it to finish.
//
// waitSecs bounds the wait; nil waits indefinitely. It returns the finished run (or the
// still-running run if the wait budget was exhausted).
func (c *ActorClient) Call(ctx context.Context, input any, options ActorStartOptions, waitSecs *int64) (ActorRun, error) {
	run, err := c.Start(ctx, input, options)
	if err != nil {
		return ActorRun{}, err
	}
	// Use the root client's run client so polling targets the canonical run route.
	return c.root.Run(run.ID).WaitForFinish(ctx, waitSecs)
}

// Build builds the given version of the Actor and returns the created build.
func (c *ActorClient) Build(ctx context.Context, versionNumber string, options ActorBuildOptions) (Build, error) {
	params := NewQueryParams()
	params.AddString("version", &versionNumber).
		AddBool("betaPackages", options.BetaPackages).
		AddString("tag", options.Tag).
		AddBool("useCache", options.UseCache).
		AddInt("waitForFinish", options.WaitForFinish)
	return postWithBody[Build](ctx, c.ctx, "builds", params, nil, contentTypeJSON)
}

// DefaultBuild resolves the Actor's default build and returns a client for it.
//
// waitForFinish optionally bounds how long (in seconds) the API waits for the build to
// finish before responding, matching the reference client's defaultBuild(options).
func (c *ActorClient) DefaultBuild(ctx context.Context, waitForFinish *int64) (*BuildClient, error) {
	params := NewQueryParams()
	params.AddInt("waitForFinish", waitForFinish)
	url := params.applyToURL(c.ctx.subURL("builds/default"))
	resp, err := c.ctx.http.call(ctx, http.MethodGet, url, nil, "", defaultRequestTimeout)
	if err != nil {
		return nil, err
	}
	build, err := parseDataEnvelope[Build](resp.body)
	if err != nil {
		return nil, err
	}
	return newBuildClient(c.ctx.http, c.baseURL, build.ID), nil
}

// ValidateInput validates the given input against the Actor's input schema.
//
// It omits the build parameter, so the API validates against the input schema of the build
// tagged "latest". To validate against a specific build, use [ActorClient.ValidateInputForBuild].
//
// On success it returns the raw JSON validation result from the API (a JSON object reporting
// whether the input is valid and, if not, the schema violations).
func (c *ActorClient) ValidateInput(ctx context.Context, input any) (json.RawMessage, error) {
	return c.ValidateInputForBuild(ctx, input, "")
}

// ValidateInputForBuild validates the given input against the input schema of a specific
// Actor build, identified by its tag or number (e.g. "latest", "0.1.2"). An empty build
// omits the parameter, so the API validates against the build tagged "latest" (per the
// OpenAPI specification), equivalent to [ActorClient.ValidateInput].
//
// On success it returns the raw JSON validation result from the API (a JSON object reporting
// whether the input is valid and, if not, the schema violations).
func (c *ActorClient) ValidateInputForBuild(ctx context.Context, input any, build string) (json.RawMessage, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	params := NewQueryParams()
	if build != "" {
		params.AddString("build", &build)
	}
	url := params.applyToURL(c.ctx.subURL("validate-input"))
	resp, err := c.ctx.http.call(ctx, http.MethodPost, url, body, contentTypeJSON, defaultRequestTimeout)
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

// LastRunWithOptions returns a client for the last run of this Actor, optionally filtered by
// status and/or origin. See LastRunOptions. Mirrors the reference client's
// lastRun({ status, origin }).
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
