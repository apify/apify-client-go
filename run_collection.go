package apify

import "context"

// RunListOptions adds run-specific filters on top of [ListOptions] for
// [RunCollectionClient.List]. The startedAfter/startedBefore filters are only honoured by
// the Actor-scoped and task-scoped run collections.
type RunListOptions struct {
	// Status filters by one or more run statuses (e.g. "SUCCEEDED", "RUNNING"). Sent as a
	// comma-separated list, as the API accepts.
	Status []string
	// StartedAfter filters to runs started after this ISO-8601 timestamp.
	StartedAfter *string
	// StartedBefore filters to runs started before this ISO-8601 timestamp.
	StartedBefore *string
}

func (o RunListOptions) apply(q *QueryParams) {
	q.AddCSV("status", o.Status).
		AddString("startedAfter", o.StartedAfter).
		AddString("startedBefore", o.StartedBefore)
}

// RunCollectionClient is a client for a run collection: the account-wide collection
// (GET /v2/actor-runs), an Actor's runs (GET /v2/actors/{id}/runs), or a task's runs
// (GET /v2/actor-tasks/{id}/runs).
type RunCollectionClient struct {
	ctx *resourceContext
}

func newRunCollectionClient(hc *httpClient, baseURL, resourcePath string) *RunCollectionClient {
	return &RunCollectionClient{ctx: newCollectionContext(hc, baseURL, resourcePath)}
}

// List lists runs, applying the standard pagination and the run-specific filters.
func (c *RunCollectionClient) List(ctx context.Context, options ListOptions, filter RunListOptions) (PaginationList[ActorRun], error) {
	params := NewQueryParams()
	options.apply(params)
	filter.apply(params)
	return listResource[ActorRun](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over the runs matching the options and filter, fetching
// pages on demand. The options' Limit caps the total number of runs yielded (unset means all);
// the per-page size is chunkSize (nil for the server default). Mirrors the reference client's
// iterable list().
func (c *RunCollectionClient) Iterate(options ListOptions, filter RunListOptions, chunkSize *int64) *ListIterator[ActorRun] {
	return newListIterator(options.Limit, chunkSize, func(ctx context.Context, offset, limit int64) (PaginationList[ActorRun], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts, filter)
	})
}
