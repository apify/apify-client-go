package apify

import (
	"context"
	"net/http"
	"testing"
)

func countHeaders(total, offset, limit int64, scanned *int64) http.Header {
	h := http.Header{}
	h.Set("X-Apify-Pagination-Total", itoa(total))
	h.Set("X-Apify-Pagination-Offset", itoa(offset))
	h.Set("X-Apify-Pagination-Limit", itoa(limit))
	if scanned != nil {
		h.Set("X-Apify-Pagination-Count", itoa(*scanned))
	}
	return h
}

// On a page of dataset items, Count is the number of rows the API scanned, or the number of items
// returned when that is larger.
func TestListItemsCountIsTheScannedRowsOrTheReturnedItems(t *testing.T) {
	cases := []struct {
		name    string
		body    string
		scanned *int64
		want    int64
	}{
		{"a filter dropped items", `[1,2]`, Ptr(int64(5)), 5},
		{"unwind returned more items than rows", `[1,2,3,4]`, Ptr(int64(2)), 4},
		{"a lagging scanned count falls below the items", `[1,2,3]`, Ptr(int64(0)), 3},
		{"no header", `[1,2,3]`, nil, 3},
		{"an empty page that scanned nothing", `[]`, Ptr(int64(0)), 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			backend := &mockBackend{responses: []mockResponse{{status: 200, body: c.body, headers: countHeaders(10, 0, 5, c.scanned)}}}
			page, err := testClient(backend, 0).Dataset("ds").ListItems(context.Background(), DatasetListItemsOptions{})
			if err != nil {
				t.Fatalf("ListItems: %v", err)
			}
			if page.Count != c.want {
				t.Fatalf("Count = %d, want %d", page.Count, c.want)
			}
		})
	}
}
