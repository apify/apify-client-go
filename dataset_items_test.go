package apify

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// DatasetClient.CreateItemsPublicURLWithFormat adds a `format` query parameter, and
// CreateItemsPublicURL (unchanged) omits it, defaulting to the endpoint's own json default.
func TestCreateItemsPublicURLWithFormat(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"ds1"}}`)}
	client := testClient(backend, 0)
	dataset := client.Dataset("ds1")

	url, err := dataset.CreateItemsPublicURLWithFormat(context.Background(), DatasetListItemsOptions{}, nil, FormatCSV)
	if err != nil {
		t.Fatalf("create items public url with format: %v", err)
	}
	if !strings.Contains(url, "format=csv") {
		t.Fatalf("expected format=csv in url, got %q", url)
	}

	url, err = dataset.CreateItemsPublicURL(context.Background(), DatasetListItemsOptions{}, nil)
	if err != nil {
		t.Fatalf("create items public url: %v", err)
	}
	if strings.Contains(url, "format=") {
		t.Fatalf("CreateItemsPublicURL must not set a format param, url was: %q", url)
	}
}

// IterateDatasetItems must advance and stop by the number of rows scanned
// (X-Apify-Pagination-Count), not the number of items a page returned: a filter can make a page
// return fewer items than the rows it covered (even zero, on a page that still has more rows
// behind it), and the iterator must not stop early in that case.
func TestIterateDatasetItemsUsesScannedCount(t *testing.T) {
	// pageHeaders builds the X-Apify-Pagination-* headers for a page that scanned `scanned` rows
	// of a dataset of `total` rows, starting at `offset`, having requested up to `limit`.
	pageHeaders := func(total, offset, limit, scanned int64) http.Header {
		h := http.Header{}
		h.Set("X-Apify-Pagination-Total", itoa(total))
		h.Set("X-Apify-Pagination-Offset", itoa(offset))
		h.Set("X-Apify-Pagination-Limit", itoa(limit))
		h.Set("X-Apify-Pagination-Count", itoa(scanned))
		return h
	}
	backend := &mockBackend{responses: []mockResponse{
		// Page 1: the API scanned 5 rows (offset 0-4) but a filter dropped all of them.
		{status: 200, body: `[]`, headers: pageHeaders(10, 0, 5, 5)},
		// Page 2: scanned the remaining 5 rows (offset 5-9), all kept.
		{status: 200, body: `[1,2,3,4,5]`, headers: pageHeaders(10, 5, 5, 5)},
		// Page 3: no Limit was set, so the iterator does not stop after page 2's 5 items; it
		// needs this trailing empty page (scanned 0) to confirm there is nothing left.
		{status: 200, body: `[]`, headers: pageHeaders(10, 10, 5, 0)},
	}}
	client := testClient(backend, 0)

	it := client.Dataset("ds1").IterateItems(DatasetListItemsOptions{}, Ptr(int64(5)))
	var items []json.RawMessage
	for {
		item, err := it.Next(context.Background())
		if err != nil {
			t.Fatalf("iterate: %v", err)
		}
		if item == nil {
			break
		}
		items = append(items, *item)
	}
	if len(items) != 5 {
		t.Fatalf("expected 5 items across both non-empty pages, got %d", len(items))
	}
	if backend.calls != 3 {
		t.Fatalf("expected exactly 3 page fetches (including the trailing empty one), got %d", backend.calls)
	}
}
