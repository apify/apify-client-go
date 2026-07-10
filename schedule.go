package apify

import "context"

// ScheduleCollectionClient is a client for the schedule collection
// (GET/POST /v2/schedules).
type ScheduleCollectionClient struct {
	ctx *resourceContext
}

func newScheduleCollectionClient(hc *httpClient, baseURL string) *ScheduleCollectionClient {
	return &ScheduleCollectionClient{ctx: newCollectionContext(hc, baseURL, "schedules")}
}

// List lists the account's schedules.
func (c *ScheduleCollectionClient) List(ctx context.Context, options ListOptions) (PaginationList[Schedule], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[Schedule](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over the schedules matching the options, fetching pages on
// demand. The options' Limit caps the total number of schedules yielded (unset means all); the
// per-page size is chunkSize (nil for the server default). Mirrors the reference client's
// iterable list().
func (c *ScheduleCollectionClient) Iterate(options ListOptions, chunkSize *int64) *ListIterator[Schedule] {
	return newListIterator(options.Limit, chunkSize, offsetVal(options.Offset), func(ctx context.Context, offset, limit int64) (PaginationList[Schedule], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts)
	})
}

// Create creates a new schedule. schedule is any JSON-serializable schedule definition.
func (c *ScheduleCollectionClient) Create(ctx context.Context, schedule any) (Schedule, error) {
	return createResource[Schedule](ctx, c.ctx, NewQueryParams(), schedule)
}

// ScheduleClient is a client for a specific schedule (/v2/schedules/{scheduleId}).
type ScheduleClient struct {
	ctx *resourceContext
}

func newScheduleClient(hc *httpClient, baseURL, id string) *ScheduleClient {
	return &ScheduleClient{ctx: newSingleContext(hc, baseURL, "schedules", id)}
}

// Get fetches the schedule. The bool reports whether it exists.
func (c *ScheduleClient) Get(ctx context.Context) (Schedule, bool, error) {
	return getResource[Schedule](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the schedule with the given fields and returns the updated object.
func (c *ScheduleClient) Update(ctx context.Context, newFields any) (Schedule, error) {
	return updateResource[Schedule](ctx, c.ctx, "", newFields)
}

// Delete deletes the schedule.
func (c *ScheduleClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// GetLog fetches the schedule's invocation log as text, or (\"\", false, nil) if absent.
func (c *ScheduleClient) GetLog(ctx context.Context) (string, bool, error) {
	resp, err := getRaw(ctx, c.ctx, "log", NewQueryParams())
	if err != nil {
		return "", false, err
	}
	if resp == nil {
		return "", false, nil
	}
	return string(resp.body), true, nil
}
