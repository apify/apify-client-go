package apify

import (
	"context"
	"testing"
)

// stubFetcher returns a page-fetch closure over a synthetic collection of `total` sequential
// ints. It honors offset and the requested page limit (limit <= 0 means "server default",
// modeled by defaultPage), and reports the true Total. Requested (offset, limit) pairs are
// appended to calls so tests can assert the paging arithmetic. This keeps the ListIterator
// unit tests hermetic — no network or APIFY_TOKEN required.
func stubFetcher(total int, defaultPage int64, calls *[][2]int64) func(context.Context, int64, int64) (PaginationList[int], error) {
	return func(_ context.Context, offset, limit int64) (PaginationList[int], error) {
		if calls != nil {
			*calls = append(*calls, [2]int64{offset, limit})
		}
		pageLen := limit
		if pageLen <= 0 {
			pageLen = defaultPage
		}
		var items []int
		for i := offset; i < offset+pageLen && i < int64(total); i++ {
			items = append(items, int(i))
		}
		return PaginationList[int]{
			Total:  int64(total),
			Offset: offset,
			Limit:  pageLen,
			Count:  int64(len(items)),
			Items:  items,
		}, nil
	}
}

// boundedFetcher models two server behaviors the plain stubFetcher cannot: (a) a server-side
// page cap (serverPage), so a page can come back non-empty but shorter than requested, and
// (b) an over-reported / lagging Total (reportedTotal), which may exceed the real backing size.
// It backs a collection of `backing` sequential ints and appends each requested (offset,limit)
// to calls.
func boundedFetcher(backing int, reportedTotal, serverPage int64, calls *[][2]int64) func(context.Context, int64, int64) (PaginationList[int], error) {
	return func(_ context.Context, offset, limit int64) (PaginationList[int], error) {
		if calls != nil {
			*calls = append(*calls, [2]int64{offset, limit})
		}
		pageLen := limit
		if pageLen <= 0 {
			pageLen = serverPage
		}
		if serverPage > 0 && pageLen > serverPage {
			pageLen = serverPage
		}
		var items []int
		for i := offset; i < offset+pageLen && i < int64(backing); i++ {
			items = append(items, int(i))
		}
		return PaginationList[int]{
			Total:  reportedTotal,
			Offset: offset,
			Limit:  pageLen,
			Count:  int64(len(items)),
			Items:  items,
		}, nil
	}
}

// drainInts fully drains an int iterator, with a hard safety bound so a broken termination
// condition fails fast instead of hanging the test.
func drainInts(t *testing.T, it *ListIterator[int]) []int {
	t.Helper()
	var out []int
	for len(out) <= 100000 {
		item, err := it.Next(context.Background())
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if item == nil {
			return out
		}
		out = append(out, *item)
	}
	t.Fatal("iterator did not terminate")
	return nil
}

func TestListIteratorTotalCap(t *testing.T) {
	// Limit=3 (total cap), page size 2, backing collection of 10 → exactly 3 items across 2 pages.
	var calls [][2]int64
	it := newListIterator(ptrInt64(3), ptrInt64(2), 0, stubFetcher(10, 1000, &calls))
	got := drainInts(t, it)
	if len(got) != 3 {
		t.Fatalf("expected 3 items (Limit cap), got %d: %v", len(got), got)
	}
	// First page requests min(limit,chunk)=2 at offset 0; second requests min(remaining=1,chunk)=1 at offset 2.
	want := [][2]int64{{0, 2}, {2, 1}}
	if len(calls) != len(want) {
		t.Fatalf("expected %d page fetches, got %d: %v", len(want), len(calls), calls)
	}
	for i, w := range want {
		if calls[i] != w {
			t.Fatalf("page %d: requested (offset,limit)=%v, want %v (all calls: %v)", i, calls[i], w, calls)
		}
	}
}

func TestListIteratorNoCapPagesAll(t *testing.T) {
	// No total cap, page size 2, backing 5 → all 5 items across 3 pages (2+2+1).
	var calls [][2]int64
	it := newListIterator(nil, ptrInt64(2), 0, stubFetcher(5, 1000, &calls))
	got := drainInts(t, it)
	if len(got) != 5 {
		t.Fatalf("expected all 5 items, got %d: %v", len(got), got)
	}
	for i, v := range got {
		if v != i {
			t.Fatalf("item %d = %d, want %d (items: %v)", i, v, i, got)
		}
	}
	if len(calls) != 3 {
		t.Fatalf("expected 3 page fetches (2+2+1), got %d: %v", len(calls), calls)
	}
}

func TestListIteratorServerDefaultPage(t *testing.T) {
	// No cap and no chunk size → the page limit is left unset (0) so the server default applies.
	var calls [][2]int64
	it := newListIterator(nil, nil, 0, stubFetcher(3, 1000, &calls))
	got := drainInts(t, it)
	if len(got) != 3 {
		t.Fatalf("expected 3 items, got %d", len(got))
	}
	if calls[0][1] != 0 {
		t.Fatalf("first page should request limit 0 (server default), got %d", calls[0][1])
	}
}

func TestListIteratorLimitLargerThanTotal(t *testing.T) {
	// A Limit larger than the collection yields the whole collection (cap bounded by Total).
	it := newListIterator(ptrInt64(100), nil, 0, stubFetcher(4, 1000, nil))
	if got := drainInts(t, it); len(got) != 4 {
		t.Fatalf("expected 4 items (whole collection), got %d", len(got))
	}
}

func TestListIteratorEmptyCollection(t *testing.T) {
	it := newListIterator(nil, ptrInt64(2), 0, stubFetcher(0, 1000, nil))
	if got := drainInts(t, it); len(got) != 0 {
		t.Fatalf("expected 0 items, got %d", len(got))
	}
}

func TestListIteratorHonorsStartOffset(t *testing.T) {
	// Caller-set Offset=4 on a backing collection of 10, page size 3: iteration must start at
	// item 4 and yield 4..9 (6 items), mirroring the reference's options.offset start point.
	var calls [][2]int64
	it := newListIterator(nil, ptrInt64(3), 4, stubFetcher(10, 1000, &calls))
	got := drainInts(t, it)
	want := []int{4, 5, 6, 7, 8, 9}
	if len(got) != len(want) {
		t.Fatalf("expected %v (from offset 4), got %v", want, got)
	}
	for i, v := range want {
		if got[i] != v {
			t.Fatalf("item %d = %d, want %d (items: %v)", i, got[i], v, got)
		}
	}
	// The first page must be requested at the caller's offset, not 0.
	if calls[0][0] != 4 {
		t.Fatalf("first page requested at offset %d, want 4 (calls: %v)", calls[0][0], calls)
	}
}

func TestListIteratorStartOffsetWithLimit(t *testing.T) {
	// Offset=4 with a total cap of 3 must yield exactly items 4,5,6 (cap counts from the offset).
	it := newListIterator(ptrInt64(3), ptrInt64(2), 4, stubFetcher(10, 1000, nil))
	got := drainInts(t, it)
	want := []int{4, 5, 6}
	if len(got) != len(want) {
		t.Fatalf("expected %v (offset 4, cap 3), got %v", want, got)
	}
	for i, v := range want {
		if got[i] != v {
			t.Fatalf("item %d = %d, want %d (items: %v)", i, got[i], v, got)
		}
	}
}

func TestListIteratorNonFinalShortPage(t *testing.T) {
	// The server caps every page at 2 items even when the iterator requests a larger chunk (10),
	// so pages 1 and 2 are non-final yet shorter than requested. Iteration must keep paging (via
	// the reported Total) and still yield the whole collection rather than stopping early on the
	// first short page.
	var calls [][2]int64
	it := newListIterator(nil, ptrInt64(10), 0, boundedFetcher(6, 6, 2, &calls))
	got := drainInts(t, it)
	if len(got) != 6 {
		t.Fatalf("expected all 6 items despite short pages, got %d: %v", len(got), got)
	}
	for i, v := range got {
		if v != i {
			t.Fatalf("item %d = %d, want %d (items: %v)", i, v, i, got)
		}
	}
	// Pages at offsets 0,2,4 return 2 items each (remaining reaches 0 after the third), so the
	// iterator must not issue a needless fourth (empty) fetch.
	if len(calls) != 3 {
		t.Fatalf("expected 3 page fetches (2+2+2), got %d: %v", len(calls), calls)
	}
}

func TestListIteratorOverReportedTotalTerminates(t *testing.T) {
	// The endpoint over-reports Total=100 but the collection really holds 3 items (e.g. a lagging
	// pagination-total header). Once a page comes back empty the iterator must stop instead of
	// looping toward the phantom cap.
	var calls [][2]int64
	it := newListIterator(nil, ptrInt64(2), 0, boundedFetcher(3, 100, 10, &calls))
	got := drainInts(t, it)
	if len(got) != 3 {
		t.Fatalf("expected exactly 3 items (real backing size), got %d: %v", len(got), got)
	}
	// offset 0 -> 2 items, offset 2 -> 1 item, offset 3 -> empty (terminates).
	if len(calls) != 3 {
		t.Fatalf("expected 3 page fetches, got %d: %v", len(calls), calls)
	}
}

// ptrInt64 is a local helper for the pointer-typed limit/chunkSize arguments.
func ptrInt64(v int64) *int64 { return &v }
