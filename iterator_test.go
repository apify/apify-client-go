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
	it := newListIterator(ptrInt64(3), ptrInt64(2), stubFetcher(10, 1000, &calls))
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
	it := newListIterator(nil, ptrInt64(2), stubFetcher(5, 1000, &calls))
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
	it := newListIterator(nil, nil, stubFetcher(3, 1000, &calls))
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
	it := newListIterator(ptrInt64(100), nil, stubFetcher(4, 1000, nil))
	if got := drainInts(t, it); len(got) != 4 {
		t.Fatalf("expected 4 items (whole collection), got %d", len(got))
	}
}

func TestListIteratorEmptyCollection(t *testing.T) {
	it := newListIterator(nil, ptrInt64(2), stubFetcher(0, 1000, nil))
	if got := drainInts(t, it); len(got) != 0 {
		t.Fatalf("expected 0 items, got %d", len(got))
	}
}

// ptrInt64 is a local helper for the pointer-typed limit/chunkSize arguments.
func ptrInt64(v int64) *int64 { return &v }
