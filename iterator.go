package apify

import "context"

// ListIterator lazily iterates over an offset/limit-paginated collection, fetching one page
// at a time on demand. Obtain one from a collection client's Iterate method and drain it by
// calling Next until it returns (nil, nil).
//
// Its end-user semantics match the reference JS client's iterable list(): the list options'
// Limit is a cap on the total number of items yielded across all pages (unset means "all
// matching items"), the page size is the separate chunkSize argument passed to Iterate (unset
// means the server default), and a caller-set Offset on the options is honored as the starting
// point (iteration begins there and yields at most Limit items from that offset onward). This
// keeps the two clients consistent for callers reasoning about offset/limit/chunk behaviour.
//
// The Limit uses the "0 == unset" convention: a nil Limit, or Limit set to ptr(0), yields the
// whole collection (no cap), rather than zero items.
type ListIterator[T any] struct {
	// fetch fetches one page starting at offset. limit is the per-page limit to request
	// (0 means "unset", i.e. let the server choose). The collection's Iterate method bakes the
	// filters into this closure and overrides offset/limit per page.
	fetch func(ctx context.Context, offset, limit int64) (PaginationList[T], error)
	// limit caps the total number of items yielded (nil or <=0 means no cap).
	limit *int64
	// chunkSize is the per-page size (nil or <=0 means the server default).
	chunkSize *int64
	// unpaginated marks an endpoint that ignores offset/limit and always returns every item in
	// a single response (the Actor version and environment-variable collections). loadPage then
	// fetches exactly once and never asks for a second page, instead of paging to an empty one.
	unpaginated bool

	buffer      []T
	pos         int
	startOffset int64 // offset the caller asked iteration to start from (0 when unset)
	offset      int64
	remaining   int64 // items still allowed to be yielded after the current buffer; valid once started and hasLimit
	hasLimit    bool  // whether limit was set, i.e. whether remaining is meaningful (see loadPage)
	started     bool
	exhausted   bool
}

// newListIterator builds a ListIterator from a page-fetch closure, the total-item cap (limit),
// the per-page size (chunkSize) and the starting offset (startOffset, the caller's Offset, 0
// when unset). It is the constructor behind every offset-paginated collection's Iterate helper,
// keeping the paging logic in one place (DRY).
func newListIterator[T any](limit, chunkSize *int64, startOffset int64, fetch func(ctx context.Context, offset, limit int64) (PaginationList[T], error)) *ListIterator[T] {
	return &ListIterator[T]{fetch: fetch, limit: limit, chunkSize: chunkSize, startOffset: startOffset, offset: startOffset}
}

// newUnpaginatedIterator builds a ListIterator for an endpoint that ignores offset/limit and
// always returns every item in one response (e.g. the Actor version and environment-variable
// collections — see ActorVersionCollectionClient.Iterate and ActorEnvVarCollectionClient.Iterate).
// It fetches once and never asks for a second page. Matches the reference client's
// listUnpaginatedResources: under the general paging algorithm, an endpoint like this would
// return the same non-empty page forever, since it never "scans zero rows".
func newUnpaginatedIterator[T any](fetch func(ctx context.Context) (PaginationList[T], error)) *ListIterator[T] {
	return &ListIterator[T]{
		unpaginated: true,
		fetch:       func(ctx context.Context, _, _ int64) (PaginationList[T], error) { return fetch(ctx) },
	}
}

// Next returns the next item, or (nil, nil) once the collection (or the total-item cap) is
// exhausted. It calls the API for another page only when the current in-memory page is used up.
func (it *ListIterator[T]) Next(ctx context.Context) (*T, error) {
	for it.pos >= len(it.buffer) {
		if it.exhausted {
			return nil, nil
		}
		if err := it.loadPage(ctx); err != nil {
			return nil, err
		}
	}
	item := it.buffer[it.pos]
	it.pos++
	return &item, nil
}

// loadPage loads the next page into the buffer, following the same offset/limit/chunkSize
// arithmetic as the reference client's _listPaginated: the first page requests
// min(limit, chunkSize) items, and each subsequent page requests min(remaining, chunkSize)
// items, where remaining only exists when the caller set a Limit. Iteration ends when a page
// scans no rows, or (if a Limit was set) once that many items have been yielded — never by
// consulting the reported Total, which goes stale the moment the listed resource grows or
// shrinks while the loop runs (e.g. a dataset a running Actor keeps pushing to). Draining an
// uncapped iterator therefore always ends with one extra request that finds an empty page.
//
// Advancing the offset uses the number of rows the page's fetch scanned (page.scanned, falling
// back to the item count when a page does not report one — the two are the same for every
// endpoint except dataset items, see ListDatasetItems), capped at the limit requested for that
// page. The cap matters because the scanned count can lag a fresh push down to zero even though
// the page has rows (see ListDatasetItems), and because Unwind can make items.length exceed what
// was requested; either way advancing past the requested limit would skip rows the next request
// should have read.
func (it *ListIterator[T]) loadPage(ctx context.Context) error {
	if it.unpaginated {
		return it.loadUnpaginatedPage(ctx)
	}

	var limitParam int64
	if !it.started {
		limitParam = minForLimitParam(it.limitVal(), it.chunkVal())
	} else {
		limitParam = minForLimitParam(it.remainingLimitParam(), it.chunkVal())
	}

	page, err := it.fetch(ctx, it.offset, limitParam)
	if err != nil {
		return err
	}
	scanned := scannedRows(page, limitParam)
	it.buffer = page.Items
	it.pos = 0
	it.offset += scanned

	if !it.started {
		it.started = true
		if l := it.limitVal(); l > 0 {
			it.hasLimit = true
			it.remaining = l - scanned
		}
	} else if it.hasLimit {
		it.remaining -= scanned
	}

	if scanned == 0 || (it.hasLimit && it.remaining <= 0) {
		it.exhausted = true
	}
	return nil
}

// loadUnpaginatedPage fetches the single response of an unpaginated iterator (see
// newUnpaginatedIterator) and marks the iterator exhausted unconditionally: there is never a
// second page to ask for.
func (it *ListIterator[T]) loadUnpaginatedPage(ctx context.Context) error {
	page, err := it.fetch(ctx, 0, 0)
	if err != nil {
		return err
	}
	it.buffer = page.Items
	it.pos = 0
	it.exhausted = true
	return nil
}

// scannedRows returns the number of rows page's fetch scanned (falling back to the returned
// item count when the page does not report one — see loadPage), capped at requestedLimit
// (0 meaning "unset", i.e. no cap). Matches the reference client's scannedRows.
func scannedRows[T any](page PaginationList[T], requestedLimit int64) int64 {
	scanned := maxInt64(page.scanned, int64(len(page.Items)))
	if requestedLimit > 0 && scanned > requestedLimit {
		return requestedLimit
	}
	return scanned
}

// remainingLimitParam returns the number of items still allowed by the caller's Limit, for use
// as a minForLimitParam input (0 meaning "unset"). Only meaningful once hasLimit is set; loadPage
// never calls it otherwise.
func (it *ListIterator[T]) remainingLimitParam() int64 {
	if !it.hasLimit {
		return 0
	}
	return it.remaining
}

// limitVal returns the total-item cap as a plain int64 (0 when unset).
func (it *ListIterator[T]) limitVal() int64 {
	if it.limit == nil {
		return 0
	}
	return *it.limit
}

// chunkVal returns the per-page size as a plain int64 (0 when unset, i.e. server default). A
// nil or negative chunkSize is treated as unset, matching KeyValueStoreKeysIterator.chunkVal.
func (it *ListIterator[T]) chunkVal() int64 {
	if it.chunkSize == nil || *it.chunkSize < 0 {
		return 0
	}
	return *it.chunkSize
}

// minForLimitParam mirrors the reference client's minForLimitParam: it treats 0 as "unset" and
// returns the smaller of the two defined values, or 0 when both are unset.
func minForLimitParam(a, b int64) int64 {
	if a <= 0 {
		return maxInt64(b, 0)
	}
	if b <= 0 {
		return a
	}
	return minInt64(a, b)
}

// pageLimitPtr converts a per-page limit into the pointer the list options expect: a positive
// value is sent as the page's Limit; 0 leaves it unset so the server default applies.
func pageLimitPtr(limit int64) *int64 {
	if limit > 0 {
		return &limit
	}
	return nil
}

// offsetVal reads a caller-set starting offset (from a list options struct) into a plain int64,
// treating nil or a negative value as 0 ("start from the beginning").
func offsetVal(offset *int64) int64 {
	if offset == nil || *offset < 0 {
		return 0
	}
	return *offset
}
