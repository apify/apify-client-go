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

// Iterate returns a lazy iterator over the version's environment variables. The env-vars
// endpoint is not offset-paginated (it always returns the full set in one response), so there
// is no Limit/chunkSize control; the iterator drains that one page. Mirrors the reference
// client's envVars().list(), which likewise exposes no pagination options.
func (c *ActorEnvVarCollectionClient) Iterate() *ListIterator[ActorEnvVar] {
	return newUnpaginatedIterator(c.List)
}

// Create creates a new environment variable.
func (c *ActorEnvVarCollectionClient) Create(ctx context.Context, envVar ActorEnvVar) (ActorEnvVar, error) {
	return createResource[ActorEnvVar](ctx, c.ctx, NewQueryParams(), envVar, c.ctx.shortTimeout())
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

// Update updates the environment variable and returns the updated object. A secret's Value is
// empty in responses (Get, List, Iterate), so set it before passing a fetched object back, or the
// secret is overwritten with an empty value.
func (c *ActorEnvVarClient) Update(ctx context.Context, envVar ActorEnvVar) (ActorEnvVar, error) {
	return updateResource[ActorEnvVar](ctx, c.ctx, "", envVar)
}

// Delete deletes the environment variable.
func (c *ActorEnvVarClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}
