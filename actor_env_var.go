package apify

import "context"

// ActorEnvVarCollectionClient is a client for an Actor version's environment variable
// collection (GET/POST /v2/actors/{actorId}/versions/{versionNumber}/env-vars).
type ActorEnvVarCollectionClient struct {
	ctx *resourceContext
}

func newActorEnvVarCollectionClient(hc *httpClient, versionURL string) *ActorEnvVarCollectionClient {
	return &ActorEnvVarCollectionClient{ctx: newCollectionContext(hc, versionURL, "env-vars")}
}

// List lists the version's environment variables.
func (c *ActorEnvVarCollectionClient) List(ctx context.Context) (PaginationList[ActorEnvVar], error) {
	return listResource[ActorEnvVar](ctx, c.ctx, "", NewQueryParams())
}

// Iterate returns a lazy iterator over the version's environment variables. Mirrors the
// reference client's iterable list(). The env-vars endpoint is not offset-paginated (it
// returns the full set in a single page), so there is no Limit/chunkSize control and the
// closure ignores the offset/limit arguments; the iterator drains that one page.
func (c *ActorEnvVarCollectionClient) Iterate() *ListIterator[ActorEnvVar] {
	return newListIterator(nil, nil, 0, func(ctx context.Context, _, _ int64) (PaginationList[ActorEnvVar], error) {
		return c.List(ctx)
	})
}

// Create creates a new environment variable.
func (c *ActorEnvVarCollectionClient) Create(ctx context.Context, envVar ActorEnvVar) (ActorEnvVar, error) {
	return createResource[ActorEnvVar](ctx, c.ctx, NewQueryParams(), envVar)
}

// ActorEnvVarClient is a client for a single environment variable
// (GET/PUT/DELETE /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{name}).
type ActorEnvVarClient struct {
	ctx *resourceContext
}

func newActorEnvVarClient(hc *httpClient, versionURL, name string) *ActorEnvVarClient {
	return &ActorEnvVarClient{ctx: newSingleContext(hc, versionURL, "env-vars", name)}
}

// Get fetches the environment variable. The bool reports whether it exists.
func (c *ActorEnvVarClient) Get(ctx context.Context) (ActorEnvVar, bool, error) {
	return getResource[ActorEnvVar](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the environment variable and returns the updated object.
func (c *ActorEnvVarClient) Update(ctx context.Context, envVar ActorEnvVar) (ActorEnvVar, error) {
	return updateResource[ActorEnvVar](ctx, c.ctx, "", envVar)
}

// Delete deletes the environment variable.
func (c *ActorEnvVarClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}
