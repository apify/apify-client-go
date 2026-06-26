package apify

import (
	"context"
	"encoding/json"
)

// TaskClient is a client for a specific Actor task.
//
// Tasks are pre-configured Actor runs with stored input. The client provides CRUD methods
// plus convenience helpers to start/call the task and access its input, runs and webhooks.
type TaskClient struct {
	root *ApifyClient
	ctx  *resourceContext
}

func newTaskClient(root *ApifyClient, hc *httpClient, baseURL, id string) *TaskClient {
	return &TaskClient{root: root, ctx: newSingleContext(hc, baseURL, "actor-tasks", id)}
}

// Get fetches the task object. The bool reports whether it exists.
func (c *TaskClient) Get(ctx context.Context) (Task, bool, error) {
	return getResource[Task](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the task with the given fields and returns the updated object.
func (c *TaskClient) Update(ctx context.Context, newFields any) (Task, error) {
	return updateResource[Task](ctx, c.ctx, "", newFields)
}

// Delete deletes the task.
func (c *TaskClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// TaskStartOptions configures starting a task run ([TaskClient.Start]/[TaskClient.Call]).
//
// It mirrors [ActorStartOptions] but omits the fields the task run endpoint does not accept
// (the Actor-only `contentType` and `forcePermissionLevel`), matching the reference client.
type TaskStartOptions struct {
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
	// RestartOnError, if true, restarts the run if it fails.
	RestartOnError *bool
	// Webhooks are ad-hoc webhooks to attach to this run (base64-encoded JSON).
	Webhooks []any
}

func (o TaskStartOptions) apply(q *QueryParams) {
	q.AddString("build", o.Build).
		AddInt("memory", o.MemoryMbytes).
		AddInt("timeout", o.TimeoutSecs).
		AddInt("waitForFinish", o.WaitForFinish).
		AddInt("maxItems", o.MaxItems).
		AddFloat("maxTotalChargeUsd", o.MaxTotalChargeUsd).
		AddBool("restartOnError", o.RestartOnError)
	if encoded := encodeWebhooks(o.Webhooks); encoded != nil {
		q.AddString("webhooks", encoded)
	}
}

// Start starts the task and returns immediately with the created run. input optionally
// overrides the task's stored input (nil to use the stored input).
func (c *TaskClient) Start(ctx context.Context, input any, options TaskStartOptions) (ActorRun, error) {
	params := NewQueryParams()
	options.apply(params)
	body, err := marshalInput(input)
	if err != nil {
		return ActorRun{}, err
	}
	return postWithBody[ActorRun](ctx, c.ctx, "runs", params, body, contentTypeJSON)
}

// Call starts the task and waits (client-side polling) for it to finish. waitSecs bounds the
// wait; nil waits indefinitely.
func (c *TaskClient) Call(ctx context.Context, input any, options TaskStartOptions, waitSecs *int64) (ActorRun, error) {
	run, err := c.Start(ctx, input, options)
	if err != nil {
		return ActorRun{}, err
	}
	return c.root.Run(run.ID).WaitForFinish(ctx, waitSecs)
}

// GetInput fetches the task's stored input, or (nil, false, nil) if none is set.
func (c *TaskClient) GetInput(ctx context.Context) (json.RawMessage, bool, error) {
	resp, err := getRaw(ctx, c.ctx, "input", NewQueryParams())
	if err != nil {
		return nil, false, err
	}
	if resp == nil {
		return nil, false, nil
	}
	return resp.body, true, nil
}

// UpdateInput replaces the task's stored input and returns the updated input.
func (c *TaskClient) UpdateInput(ctx context.Context, input any) (json.RawMessage, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	url := c.ctx.subURL("input")
	resp, err := c.ctx.http.call(ctx, "PUT", url, data, contentTypeJSON, defaultRequestTimeout)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

// LastRun returns a client for the last run of this task, optionally filtered by status
// (e.g. "SUCCEEDED"). Pass an empty status for no filter.
//
// To also filter by run origin, use LastRunWithOptions.
func (c *TaskClient) LastRun(status string) *RunClient {
	return c.LastRunWithOptions(LastRunOptions{Status: status})
}

// LastRunWithOptions returns a client for the last run of this task, optionally filtered by
// status and/or origin. See LastRunOptions. Mirrors the reference client's
// lastRun({ status, origin }).
func (c *TaskClient) LastRunWithOptions(options LastRunOptions) *RunClient {
	client := newRunClient(c.root, c.ctx.http, c.ctx.subURL(""), "runs", "last")
	client.setLastRunParams(options)
	return client
}

// Runs returns a client for this task's run collection.
func (c *TaskClient) Runs() *RunCollectionClient {
	return newRunCollectionClient(c.ctx.http, c.ctx.subURL(""), "runs")
}

// Webhooks returns a client for this task's webhook collection.
func (c *TaskClient) Webhooks() *WebhookCollectionClient {
	return newWebhookCollectionClientWithBase(c.ctx.http, c.ctx.subURL(""))
}
