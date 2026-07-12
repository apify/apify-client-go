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

	buffer      []T
	pos         int
	startOffset int64 // offset the caller asked iteration to start from (0 when unset)
	offset      int64
	remaining   int64 // items still allowed to be yielded after the current buffer; valid once started
	started     bool
	exhausted   bool
}

// newListIterator builds a ListIterator from a page-fetch closure, the total-item cap (limit),
// the per-page size (chunkSize) and the starting offset (startOffset, the caller's Offset, 0
// when unset). It is the single constructor behind every collection's Iterate helper, keeping
// the paging logic in one place (DRY).
func newListIterator[T any](limit, chunkSize *int64, startOffset int64, fetch func(ctx context.Context, offset, limit int64) (PaginationList[T], error)) *ListIterator[T] {
	return &ListIterator[T]{fetch: fetch, limit: limit, chunkSize: chunkSize, startOffset: startOffset, offset: startOffset}
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
		// Cap the number of items to yield: from the start offset onward at most
		// (Total - startOffset) items remain, and no more than the caller's Limit. This mirrors
		// the reference client's remainingItems = min(total - offset, limit) - firstPageCount,
		// so a caller-set Offset is honored as the starting point and a Limit larger than the
		// collection still yields everything from the offset onward.
		capItems := page.Total - it.startOffset
		if capItems < 0 {
			capItems = 0
		}
		if l := it.limitVal(); l > 0 && l < capItems {
			capItems = l
		}
		it.remaining = capItems - n
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
