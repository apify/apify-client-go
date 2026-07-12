package apify

import "context"

// WebhookCollectionClient is a client for a webhook collection: the account-wide collection
// (GET/POST /v2/webhooks) or webhooks nested under an Actor or task (read-only there).
type WebhookCollectionClient struct {
	ctx *resourceContext
}

func newWebhookCollectionClient(hc *httpClient, baseURL string) *WebhookCollectionClient {
	return &WebhookCollectionClient{ctx: newCollectionContext(hc, baseURL, "webhooks")}
}

// newWebhookCollectionClientWithBase builds a webhook collection nested under another
// resource (e.g. an Actor or task), using base as the parent URL.
func newWebhookCollectionClientWithBase(hc *httpClient, base string) *WebhookCollectionClient {
	return &WebhookCollectionClient{ctx: newCollectionContext(hc, base, "webhooks")}
}

// List lists webhooks.
func (c *WebhookCollectionClient) List(ctx context.Context, options ListOptions) (PaginationList[Webhook], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[Webhook](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over the webhooks matching the options, fetching pages on
// demand. The options' Limit caps the total number of webhooks yielded (unset means all); the
// per-page size is chunkSize (nil for the server default). Mirrors the reference client's
// iterable list().
func (c *WebhookCollectionClient) Iterate(options ListOptions, chunkSize *int64) *ListIterator[Webhook] {
	return newListIterator(options.Limit, chunkSize, offsetVal(options.Offset), func(ctx context.Context, offset, limit int64) (PaginationList[Webhook], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts)
	})
}

// Create creates a new webhook. webhook is any JSON-serializable webhook definition.
func (c *WebhookCollectionClient) Create(ctx context.Context, webhook any) (Webhook, error) {
	return createResource[Webhook](ctx, c.ctx, NewQueryParams(), webhook)
}

// WebhookClient is a client for a specific webhook (/v2/webhooks/{webhookId}).
type WebhookClient struct {
	ctx *resourceContext
}

func newWebhookClient(hc *httpClient, baseURL, id string) *WebhookClient {
	return &WebhookClient{ctx: newSingleContext(hc, baseURL, "webhooks", id)}
}

// Get fetches the webhook. The bool reports whether it exists.
func (c *WebhookClient) Get(ctx context.Context) (Webhook, bool, error) {
	return getResource[Webhook](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the webhook with the given fields and returns the updated object.
func (c *WebhookClient) Update(ctx context.Context, newFields any) (Webhook, error) {
	return updateResource[Webhook](ctx, c.ctx, "", newFields)
}

// Delete deletes the webhook.
func (c *WebhookClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// Test dispatches the webhook immediately and returns the resulting dispatch.
func (c *WebhookClient) Test(ctx context.Context) (WebhookDispatch, error) {
	return postWithBody[WebhookDispatch](ctx, c.ctx, "test", NewQueryParams(), nil, "")
}

// Dispatches returns a client for this webhook's dispatch collection.
func (c *WebhookClient) Dispatches() *WebhookDispatchCollectionClient {
	return newWebhookDispatchCollectionClientWithBase(c.ctx.http, c.ctx.subURL(""))
}
