package apify

import "context"

// ActorListOptions configures [ActorCollectionClient.List].
type ActorListOptions struct {
	// Offset is the number of Actors to skip.
	Offset *int64
	// Limit is the maximum number of Actors to return.
	Limit *int64
	// Desc, if true, returns Actors newest-first.
	Desc *bool
	// My, if true, returns only Actors owned by the current user.
	My *bool
	// SortBy sets the sort field (e.g. "createdAt", "stats.lastRunStartedAt").
	SortBy *string
}

func (o ActorListOptions) apply(q *QueryParams) {
	q.AddInt("offset", o.Offset).
		AddInt("limit", o.Limit).
		AddBool("desc", o.Desc).
		AddBool("my", o.My).
		AddString("sortBy", o.SortBy)
}

// ActorCollectionClient is a client for the Actor collection (GET/POST /v2/actors).
type ActorCollectionClient struct {
	ctx *resourceContext
}

func newActorCollectionClient(hc *httpClient, baseURL string) *ActorCollectionClient {
	return &ActorCollectionClient{ctx: newCollectionContext(hc, baseURL, "actors")}
}

// List lists the account's Actors.
func (c *ActorCollectionClient) List(ctx context.Context, options ActorListOptions) (PaginationList[Actor], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[Actor](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over the Actors matching the options, fetching pages on
// demand. The options' Limit caps the total number of Actors yielded (unset means all); the
// per-page size is chunkSize (nil for the server default). Mirrors the reference client's
// iterable list().
func (c *ActorCollectionClient) Iterate(options ActorListOptions, chunkSize *int64) *ListIterator[Actor] {
	return newListIterator(options.Limit, chunkSize, func(ctx context.Context, offset, limit int64) (PaginationList[Actor], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts)
	})
}

// Create creates a new Actor. actor is any JSON-serializable Actor definition.
func (c *ActorCollectionClient) Create(ctx context.Context, actor any) (Actor, error) {
	return createResource[Actor](ctx, c.ctx, NewQueryParams(), actor)
}
