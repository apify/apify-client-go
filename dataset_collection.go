package apify

import "context"

// DatasetCollectionClient is a client for the dataset collection (GET/POST /v2/datasets).
type DatasetCollectionClient struct {
	ctx *resourceContext
}

func newDatasetCollectionClient(hc *httpClient, baseURL string) *DatasetCollectionClient {
	return &DatasetCollectionClient{ctx: newCollectionContext(hc, baseURL, "datasets")}
}

// List lists datasets.
func (c *DatasetCollectionClient) List(ctx context.Context, options StorageListOptions) (PaginationList[Dataset], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[Dataset](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over the datasets matching the options, fetching pages on
// demand. The options' Limit caps the total number of datasets yielded (unset means all); the
// per-page size is chunkSize (nil for the server default). Mirrors the reference client's
// iterable list().
func (c *DatasetCollectionClient) Iterate(options StorageListOptions, chunkSize *int64) *ListIterator[Dataset] {
	return newListIterator(options.Limit, chunkSize, func(ctx context.Context, offset, limit int64) (PaginationList[Dataset], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return c.List(ctx, opts)
	})
}

// GetOrCreate gets the dataset with the given name, creating it if it does not exist. An
// empty name creates a new unnamed dataset.
func (c *DatasetCollectionClient) GetOrCreate(ctx context.Context, name string) (Dataset, error) {
	return getOrCreateNamed[Dataset](ctx, c.ctx, name)
}
