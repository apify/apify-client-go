package apify

import "context"

// ActorVersionCollectionClient is a client for an Actor's version collection
// (GET/POST /v2/actors/{actorId}/versions).
type ActorVersionCollectionClient struct {
	ctx *resourceContext
}

func newActorVersionCollectionClient(hc *httpClient, actorURL string) *ActorVersionCollectionClient {
	return &ActorVersionCollectionClient{ctx: newCollectionContext(hc, actorURL, "versions")}
}

// List lists the Actor's versions.
func (c *ActorVersionCollectionClient) List(ctx context.Context, options ListOptions) (PaginationList[ActorVersion], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[ActorVersion](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over the Actor's versions matching the options, fetching
// pages on demand. The options' Limit caps the total number of versions yielded (unset means
// all); the per-page size is chunkSize (nil for the server default). Mirrors the reference
// client's iterable list().
func (c *ActorVersionCollectionClient) Iterate(options ListOptions, chunkSize *int64) *ListIterator[ActorVersion] {
	return newListIterator(options.Limit, chunkSize, offsetVal(options.Offset), func(ctx context.Context, offset, limit int64) (PaginationList[ActorVersion], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts)
	})
}

// Create creates a new Actor version. version is any JSON-serializable version definition.
func (c *ActorVersionCollectionClient) Create(ctx context.Context, version any) (ActorVersion, error) {
	return createResource[ActorVersion](ctx, c.ctx, NewQueryParams(), version)
}

// ActorVersionClient is a client for a specific Actor version
// (GET/PUT/DELETE /v2/actors/{actorId}/versions/{versionNumber}).
type ActorVersionClient struct {
	ctx        *resourceContext
	versionURL string
}

func newActorVersionClient(hc *httpClient, actorURL, versionNumber string) *ActorVersionClient {
	ctx := newSingleContext(hc, actorURL, "versions", versionNumber)
	return &ActorVersionClient{ctx: ctx, versionURL: ctx.subURL("")}
}

// Get fetches the version. The bool reports whether it exists.
func (c *ActorVersionClient) Get(ctx context.Context) (ActorVersion, bool, error) {
	return getResource[ActorVersion](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the version with the given fields and returns the updated object.
func (c *ActorVersionClient) Update(ctx context.Context, newFields any) (ActorVersion, error) {
	return updateResource[ActorVersion](ctx, c.ctx, "", newFields)
}

// Delete deletes the version.
func (c *ActorVersionClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// EnvVar returns a client for a specific environment variable of this version.
func (c *ActorVersionClient) EnvVar(name string) *ActorEnvVarClient {
	return newActorEnvVarClient(c.ctx.http, c.versionURL, name)
}

// EnvVars returns a client for this version's environment variable collection.
func (c *ActorVersionClient) EnvVars() *ActorEnvVarCollectionClient {
	return newActorEnvVarCollectionClient(c.ctx.http, c.versionURL)
}
