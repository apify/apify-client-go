package apify

import "context"

// StoreListOptions configures listing/iterating the Apify Store.
type StoreListOptions struct {
	// Offset is the number of Actors to skip. The API ignores it when Search is set.
	Offset *int64
	// Limit is the maximum number of Actors to return. The default and the maximum are 1000, or
	// 100 when Search is set, where a larger value is reduced to 100.
	Limit *int64
	// Search ranks Actors by relevance to the query and returns them as a single page of at most
	// 100 Actors, with no second page. Narrow the results with Category, Username or
	// PricingModel.
	Search *string
	// SortBy sets the sort field.
	//
	// Deprecated: The API ignores it. Results are ordered by relevance when Search is set, and by
	// Store position otherwise. Sort the results on the client for another order.
	SortBy *string
	// Category filters Actors by category. It takes one or more category IDs, separated by commas
	// and case-insensitive, and an Actor must match all of them (for example "AI_AND_ML"). An
	// unknown category makes the API answer with a 400 error.
	Category *string
	// Username filters Actors by owner username.
	Username *string
	// PricingModel filters Actors by pricing model (e.g. "FREE", "PAY_PER_EVENT").
	PricingModel *string
	// IncludeUnrunnableActors, if true, turns off the safety filter that leaves out Actors that
	// are not safe to run automatically, such as those of developers without a verified identity
	// and Actors with full permissions that have few users.
	IncludeUnrunnableActors *bool
	// AllowsAgenticUsers, if true, returns only Actors that allow agentic users, and if false,
	// only Actors that do not.
	AllowsAgenticUsers *bool
	// ResponseFormat controls the shape of the response: "full" (the default) returns every field,
	// and "agent" a reduced field set for LLM consumers, in which URL, UserPictureURL, IsVerified
	// and the optional Stats counts are nil.
	ResponseFormat *string
	// IncludeInputSchema, if true, adds to each listed Actor an inputSchema field with a compact
	// version of its default build's input schema (property types and required properties only).
	// With it set, the API defaults Limit to 10 and does not accept a larger one.
	IncludeInputSchema *bool
}

func (o StoreListOptions) apply(q *QueryParams) {
	q.AddInt("offset", o.Offset).
		AddInt("limit", o.Limit).
		AddString("search", o.Search).
		AddString("sortBy", o.SortBy).
		AddString("category", o.Category).
		AddString("username", o.Username).
		AddString("pricingModel", o.PricingModel).
		AddBool("includeUnrunnableActors", o.IncludeUnrunnableActors).
		AddBool("allowsAgenticUsers", o.AllowsAgenticUsers).
		AddString("responseFormat", o.ResponseFormat).
		AddBool("includeInputSchema", o.IncludeInputSchema)
}

// StoreCollectionClient is a client for browsing the Apify Store (GET /v2/store).
type StoreCollectionClient struct {
	ctx *resourceContext
}

func newStoreCollectionClient(hc *httpClient, baseURL string) *StoreCollectionClient {
	return &StoreCollectionClient{ctx: newCollectionContext(hc, baseURL, "store")}
}

// List returns a single page of Store Actors matching the options.
func (c *StoreCollectionClient) List(ctx context.Context, options StoreListOptions) (PaginationList[ActorStoreListItem], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[ActorStoreListItem](ctx, c.ctx, "", params)
}

// StoreActorIterator lazily iterates over Apify Store Actors, fetching one page at a time.
// It is the generic [ListIterator] specialized to Store Actors; drain it with Next.
type StoreActorIterator = ListIterator[ActorStoreListItem]

// Iterate returns a lazy iterator over the Store Actors matching the options, fetching pages on
// demand. The options' Limit caps the total number of Actors yielded (unset means all); the page
// size is chunkSize (nil for the server default).
//
// With Search set, the API ranks the Actors by relevance and returns a single page of at most 100
// of them, ignoring the offset, so a second page would repeat the first. The iterator then makes
// one request: Limit is the page size, chunkSize is ignored, and no more than 100 Actors are
// yielded.
func (c *StoreCollectionClient) Iterate(options StoreListOptions, chunkSize *int64) *StoreActorIterator {
	if options.Search != nil && *options.Search != "" {
		options.Offset = nil
		return newUnpaginatedIterator(func(ctx context.Context) (PaginationList[ActorStoreListItem], error) {
			return c.List(ctx, options)
		})
	}
	return newListIterator(options.Limit, chunkSize, offsetVal(options.Offset), func(ctx context.Context, offset, limit int64) (PaginationList[ActorStoreListItem], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts)
	})
}
