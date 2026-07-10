package apify

import "context"

// WebhookDispatchCollectionClient is a client for a webhook dispatch collection: the
// account-wide collection (GET /v2/webhook-dispatches) or dispatches nested under a webhook.
type WebhookDispatchCollectionClient struct {
	ctx *resourceContext
}

func newWebhookDispatchCollectionClient(hc *httpClient, baseURL string) *WebhookDispatchCollectionClient {
	return &WebhookDispatchCollectionClient{ctx: newCollectionContext(hc, baseURL, "webhook-dispatches")}
}

// newWebhookDispatchCollectionClientWithBase builds a dispatch collection nested under a
// webhook, using base as the parent URL.
func newWebhookDispatchCollectionClientWithBase(hc *httpClient, base string) *WebhookDispatchCollectionClient {
	return &WebhookDispatchCollectionClient{ctx: newCollectionContext(hc, base, "dispatches")}
}

// List lists webhook dispatches.
func (c *WebhookDispatchCollectionClient) List(ctx context.Context, options ListOptions) (PaginationList[WebhookDispatch], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[WebhookDispatch](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over the webhook dispatches matching the options, fetching
// pages on demand. The options' Limit caps the total number of dispatches yielded (unset means
// all); the per-page size is chunkSize (nil for the server default). Mirrors the reference
// client's iterable list().
func (c *WebhookDispatchCollectionClient) Iterate(options ListOptions, chunkSize *int64) *ListIterator[WebhookDispatch] {
	return newListIterator(options.Limit, chunkSize, func(ctx context.Context, offset, limit int64) (PaginationList[WebhookDispatch], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts)
	})
}

// WebhookDispatchClient is a client for a specific webhook dispatch
// (/v2/webhook-dispatches/{dispatchId}).
type WebhookDispatchClient struct {
	ctx *resourceContext
}

func newWebhookDispatchClient(hc *httpClient, baseURL, id string) *WebhookDispatchClient {
	return &WebhookDispatchClient{ctx: newSingleContext(hc, baseURL, "webhook-dispatches", id)}
}

// Get fetches the dispatch. The bool reports whether it exists.
func (c *WebhookDispatchClient) Get(ctx context.Context) (WebhookDispatch, bool, error) {
	return getResource[WebhookDispatch](ctx, c.ctx, "", NewQueryParams())
}
