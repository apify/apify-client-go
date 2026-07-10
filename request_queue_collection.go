package apify

import "context"

// RequestQueueCollectionClient is a client for the request queue collection
// (GET/POST /v2/request-queues).
type RequestQueueCollectionClient struct {
	ctx *resourceContext
}

func newRequestQueueCollectionClient(hc *httpClient, baseURL string) *RequestQueueCollectionClient {
	return &RequestQueueCollectionClient{ctx: newCollectionContext(hc, baseURL, "request-queues")}
}

// List lists request queues.
func (c *RequestQueueCollectionClient) List(ctx context.Context, options StorageListOptions) (PaginationList[RequestQueue], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[RequestQueue](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over all request queues matching the options, fetching
// pages on demand. The options' Limit (if set) is used as the per-page size. Mirrors the
// reference client's iterable list().
func (c *RequestQueueCollectionClient) Iterate(options StorageListOptions) *ListIterator[RequestQueue] {
	return newListIterator(func(ctx context.Context, offset int64) (PaginationList[RequestQueue], error) {
		opts := options
		opts.Offset = &offset
		return c.List(ctx, opts)
	})
}

// GetOrCreate gets the queue with the given name, creating it if it does not exist. An
// empty name creates a new unnamed queue.
func (c *RequestQueueCollectionClient) GetOrCreate(ctx context.Context, name string) (RequestQueue, error) {
	return getOrCreateNamed[RequestQueue](ctx, c.ctx, name)
}
