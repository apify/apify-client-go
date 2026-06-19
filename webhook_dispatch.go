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
