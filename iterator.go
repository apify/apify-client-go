package apify

import "context"

// ListIterator lazily iterates over an offset/limit-paginated collection, fetching one page
// at a time on demand. Obtain one from a collection client's Iterate method and drain it by
// calling Next until it returns (nil, nil).
//
// It is the shared engine behind every collection's Iterate helper, mirroring the reference
// JS client whose collection list() returns an async iterator over all matching items. The
// options' Limit (if set) is used as the per-page size; when it is unset the server default
// applies. The iterator walks every matching item across pages — the Limit bounds the page
// size, not the total number of items yielded.
type ListIterator[T any] struct {
	// fetchPage fetches one page starting at the given offset. The per-call filters and
	// per-page limit are baked into the closure by the collection's Iterate method; only the
	// offset varies between pages.
	fetchPage func(ctx context.Context, offset int64) (PaginationList[T], error)

	buffer    []T
	pos       int
	offset    int64
	total     int64
	exhausted bool
}

// newListIterator builds a ListIterator from a page-fetch closure. It is the single
// constructor used by every collection's Iterate helper, keeping the paging logic in one
// place (DRY).
func newListIterator[T any](fetchPage func(ctx context.Context, offset int64) (PaginationList[T], error)) *ListIterator[T] {
	return &ListIterator[T]{fetchPage: fetchPage}
}

// Next returns the next item, or (nil, nil) once the collection is exhausted. It calls the
// API for another page only when the current in-memory page is used up.
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

// loadPage loads the next page into the buffer and advances the offset. The collection is
// exhausted once a page comes back empty or the offset reaches the reported total.
func (it *ListIterator[T]) loadPage(ctx context.Context) error {
	page, err := it.fetchPage(ctx, it.offset)
	if err != nil {
		return err
	}
	it.buffer = page.Items
	it.pos = 0
	it.total = page.Total
	it.offset += int64(len(page.Items))
	if len(page.Items) == 0 || it.offset >= it.total {
		it.exhausted = true
	}
	return nil
}
