package apify

import "context"

// ListIterator lazily iterates over a paginated collection, fetching one page at a time. Get one
// from the Iterate method of a collection client, and call Next until it returns (nil, nil).
//
// The Limit of the list options caps the total number of items yielded, and nil or 0 means all of
// them. The chunkSize argument of Iterate is the page size. If it is unset, the page size is the
// default of the API, capped by the remaining Limit. Dataset items have no default limit, so
// iterating them without a chunkSize or Limit loads the whole dataset in one response. The Offset
// of the options is where the iteration starts.
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

	buffer    []T
	pos       int
	offset    int64 // current page offset, seeded from the caller's starting offset (0 when unset)
	remaining int64 // items still allowed to be yielded after the current buffer; valid once started and hasLimit
	hasLimit  bool  // whether limit was set, i.e. whether remaining is meaningful (see loadPage)
	started   bool
	exhausted bool
}

// newListIterator builds a ListIterator from a page-fetch closure, the total-item cap (limit),
// the per-page size (chunkSize) and the starting offset (startOffset, the caller's Offset, 0
// when unset). It is the constructor behind every offset-paginated collection's Iterate helper,
// keeping the paging logic in one place (DRY).
func newListIterator[T any](limit, chunkSize *int64, startOffset int64, fetch func(ctx context.Context, offset, limit int64) (PaginationList[T], error)) *ListIterator[T] {
	return &ListIterator[T]{fetch: fetch, limit: limit, chunkSize: chunkSize, offset: startOffset}
}

// newUnpaginatedIterator iterates an endpoint that ignores offset and limit and returns every item
// in one response, such as the Actor versions. It fetches once.
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

// loadPage loads the next page. The first page requests min(limit, chunkSize) items, and later
// pages min(remaining, chunkSize), where remaining exists only if a Limit is set. Iteration ends
// when a page scans no rows or the Limit is reached, never by Total, which goes stale while the
// resource changes. An uncapped iterator therefore ends with one request for an empty page.
//
// The offset advances by the larger of the page's Count and its item count (see scannedRows),
// capped at the limit requested for the page, so that Unwind cannot make it skip rows.
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

// scannedRows returns the number of rows page covers: its Count, or the returned item count when
// that is larger, capped at requestedLimit (0 meaning "unset", i.e. no cap).
func scannedRows[T any](page PaginationList[T], requestedLimit int64) int64 {
	scanned := maxInt64(page.Count, int64(len(page.Items)))
	if requestedLimit > 0 && scanned > requestedLimit {
		return requestedLimit
	}
	return scanned
}

// remainingLimitParam returns the number of items still allowed by the caller's Limit, for use
// as a minForLimitParam input, where 0 means unset. It is meaningful only once hasLimit is set.
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

// minForLimitParam returns the smaller of a and b, ignoring a value of 0 or below, which means
// unset. It returns 0 when neither is set.
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
