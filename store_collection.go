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

// Iterate returns a lazy iterator over all Store Actors matching the options, fetching pages
// on demand. The options' Limit (if set) is used as the per-page size.
func (c *StoreCollectionClient) Iterate(options StoreListOptions) *StoreActorIterator {
	return &StoreActorIterator{client: c, options: options}
}

// StoreActorIterator lazily iterates over Apify Store Actors, fetching one page at a time.
type StoreActorIterator struct {
	client  *StoreCollectionClient
	options StoreListOptions

	buffer    []ActorStoreListItem
	pos       int
	offset    int64
	total     int64
	exhausted bool
}

// Next returns the next Store Actor, or (nil, nil) when the iterator is exhausted.
func (it *StoreActorIterator) Next(ctx context.Context) (*ActorStoreListItem, error) {
	for it.pos >= len(it.buffer) {
		if it.exhausted {
			return nil, nil
		}
		if err := it.fetchPage(ctx); err != nil {
			return nil, err
		}
	}
	item := it.buffer[it.pos]
	it.pos++
	return &item, nil
}

// fetchPage loads the next page of Store Actors into the buffer.
func (it *StoreActorIterator) fetchPage(ctx context.Context) error {
	opts := it.options
	opts.Offset = &it.offset
	page, err := it.client.List(ctx, opts)
	if err != nil {
		return err
	}
	it.buffer = page.Items
	it.pos = 0
	it.total = page.Total
	it.offset += int64(len(page.Items))

	// Exhausted when the page is empty or we have reached the reported total.
	if len(page.Items) == 0 || it.offset >= it.total {
		it.exhausted = true
	}
	return nil
}
