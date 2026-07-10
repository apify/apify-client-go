package apify

import "context"

// StoreListOptions configures listing/iterating the Apify Store.
type StoreListOptions struct {
	// Offset is the number of Actors to skip.
	Offset *int64
	// Limit is the maximum number of Actors to return.
	Limit *int64
	// Search filters Actors by a full-text search query.
	Search *string
	// SortBy sets the sort field (e.g. "popularity", "newest").
	SortBy *string
	// Category filters Actors by category.
	Category *string
	// Username filters Actors by owner username.
	Username *string
	// PricingModel filters Actors by pricing model (e.g. "FREE", "FLAT_PRICE_PER_MONTH").
	PricingModel *string
	// IncludeUnrunnableActors includes Actors the current user cannot run.
	IncludeUnrunnableActors *bool
	// AllowsAgenticUsers filters to Actors that allow agentic users.
	AllowsAgenticUsers *bool
	// ResponseFormat selects the response format.
	ResponseFormat *string
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
		AddString("responseFormat", o.ResponseFormat)
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

// Iterate returns a lazy iterator over all Store Actors matching the options, fetching pages
// on demand. The options' Limit (if set) is used as the per-page size.
func (c *StoreCollectionClient) Iterate(options StoreListOptions) *StoreActorIterator {
	return newListIterator(func(ctx context.Context, offset int64) (PaginationList[ActorStoreListItem], error) {
		opts := options
		opts.Offset = &offset
		return c.List(ctx, opts)
	})
}
