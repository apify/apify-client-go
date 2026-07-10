package apify

import "context"

// ListIterator lazily iterates over an offset/limit-paginated collection, fetching one page
// at a time on demand. Obtain one from a collection client's Iterate method and drain it by
// calling Next until it returns (nil, nil).
//
// Its end-user semantics match the reference JS client's iterable list(): the list options'
// Limit is a cap on the total number of items yielded across all pages (unset means "all
// matching items"), and the page size is the separate chunkSize argument passed to Iterate
// (unset means the server default). This keeps the two clients consistent for callers reasoning
// about limit/chunk behaviour.
type ListIterator[T any] struct {
	// fetch fetches one page starting at offset. limit is the per-page limit to request
	// (0 means "unset", i.e. let the server choose). The collection's Iterate method bakes the
	// filters into this closure and overrides offset/limit per page.
	fetch func(ctx context.Context, offset, limit int64) (PaginationList[T], error)
	// limit caps the total number of items yielded (nil or <=0 means no cap).
	limit *int64
	// chunkSize is the per-page size (nil or <=0 means the server default).
	chunkSize *int64

	buffer    []T
	pos       int
	offset    int64
	remaining int64 // items still allowed to be yielded after the current buffer; valid once started
	started   bool
	exhausted bool
}

// newListIterator builds a ListIterator from a page-fetch closure, the total-item cap (limit)
// and the per-page size (chunkSize). It is the single constructor behind every collection's
// Iterate helper, keeping the paging logic in one place (DRY).
func newListIterator[T any](limit, chunkSize *int64, fetch func(ctx context.Context, offset, limit int64) (PaginationList[T], error)) *ListIterator[T] {
	return &ListIterator[T]{fetch: fetch, limit: limit, chunkSize: chunkSize}
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
// min(limit, chunkSize) items; the total cap is then bounded by the reported total, and each
// subsequent page requests min(remaining, chunkSize) items. Iteration ends when a page comes
// back empty or the remaining cap reaches zero.
func (it *ListIterator[T]) loadPage(ctx context.Context) error {
	var limitParam int64
	if !it.started {
		limitParam = minForLimitParam(it.limitVal(), it.chunkVal())
	} else {
		limitParam = minForLimitParam(it.remaining, it.chunkVal())
	}

	page, err := it.fetch(ctx, it.offset, limitParam)
	if err != nil {
		return err
	}
	n := int64(len(page.Items))
	it.buffer = page.Items
	it.pos = 0
	it.offset += n

	if !it.started {
		it.started = true
		// Cap the total number of yielded items by the reported total (a Limit larger than
		// the collection yields the whole collection, matching the reference).
		cap := page.Total
		if l := it.limitVal(); l > 0 && l < cap {
			cap = l
		}
		it.remaining = cap - n
	} else {
		it.remaining -= n
	}

	if n == 0 || it.remaining <= 0 {
		it.exhausted = true
	}
	return nil
}

// limitVal returns the total-item cap as a plain int64 (0 when unset).
func (it *ListIterator[T]) limitVal() int64 {
	if it.limit == nil {
		return 0
	}
	return *it.limit
}

// chunkVal returns the per-page size as a plain int64 (0 when unset).
func (it *ListIterator[T]) chunkVal() int64 {
	if it.chunkSize == nil {
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
