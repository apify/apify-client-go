package apify

import "context"

// KeyValueStoreCollectionClient is a client for the key-value store collection
// (GET/POST /v2/key-value-stores).
type KeyValueStoreCollectionClient struct {
	ctx *resourceContext
}

func newKeyValueStoreCollectionClient(hc *httpClient, baseURL string) *KeyValueStoreCollectionClient {
	return &KeyValueStoreCollectionClient{ctx: newCollectionContext(hc, baseURL, "key-value-stores")}
}

// List lists key-value stores.
func (c *KeyValueStoreCollectionClient) List(ctx context.Context, options StorageListOptions) (PaginationList[KeyValueStore], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[KeyValueStore](ctx, c.ctx, "", params)
}

// Iterate returns a lazy iterator over all key-value stores matching the options, fetching
// pages on demand. The options' Limit (if set) is used as the per-page size. Mirrors the
// reference client's iterable list().
func (c *KeyValueStoreCollectionClient) Iterate(options StorageListOptions) *ListIterator[KeyValueStore] {
	return newListIterator(func(ctx context.Context, offset int64) (PaginationList[KeyValueStore], error) {
		opts := options
		opts.Offset = &offset
		return c.List(ctx, opts)
	})
}

// GetOrCreate gets the store with the given name, creating it if it does not exist. An
// empty name creates a new unnamed store.
func (c *KeyValueStoreCollectionClient) GetOrCreate(ctx context.Context, name string) (KeyValueStore, error) {
	return getOrCreateNamed[KeyValueStore](ctx, c.ctx, name)
}
