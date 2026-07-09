package apify

import "context"

// RunListOptions adds run-specific filters on top of [ListOptions] for
// [RunCollectionClient.List]. The startedAfter/startedBefore filters are only honoured by
// the Actor-scoped and task-scoped run collections.
type RunListOptions struct {
	// Status filters by one or more run statuses (the ActorJobStatus constants). Sent as a
	// comma-separated list, as the API accepts.
	Status []ActorJobStatus
	// StartedAfter filters to runs started after this ISO-8601 timestamp.
	StartedAfter *string
	// StartedBefore filters to runs started before this ISO-8601 timestamp.
	StartedBefore *string
}

func (o RunListOptions) apply(q *QueryParams) {
	statuses := make([]string, len(o.Status))
	for i, s := range o.Status {
		statuses[i] = string(s)
	}
	q.AddCSV("status", statuses).
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
