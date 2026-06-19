package apify

import "context"

// TaskCollectionClient is a client for the Actor task collection
// (GET/POST /v2/actor-tasks).
type TaskCollectionClient struct {
	ctx *resourceContext
}

func newTaskCollectionClient(hc *httpClient, baseURL string) *TaskCollectionClient {
	return &TaskCollectionClient{ctx: newCollectionContext(hc, baseURL, "actor-tasks")}
}

// List lists the account's tasks.
func (c *TaskCollectionClient) List(ctx context.Context, options ListOptions) (PaginationList[Task], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[Task](ctx, c.ctx, "", params)
}

// Create creates a new task. task is any JSON-serializable task definition.
func (c *TaskCollectionClient) Create(ctx context.Context, task any) (Task, error) {
	return createResource[Task](ctx, c.ctx, NewQueryParams(), task)
}
