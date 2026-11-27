package apify

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

// searchIgnoringOffsetBackend plays the Store API, which ignores the offset when a search is set:
// every request for such a search gets the same page of actors.
type searchIgnoringOffsetBackend struct {
	mu   sync.Mutex
	urls []string
}

func (b *searchIgnoringOffsetBackend) Do(req *http.Request) (*http.Response, error) {
	b.mu.Lock()
	b.urls = append(b.urls, req.URL.String())
	b.mu.Unlock()
	const page = `{"data":{"total":2,"offset":0,"count":2,"limit":100,"items":[` +
		`{"id":"a1","name":"one","username":"u","title":"One","stats":{"totalUsers":1}},` +
		`{"id":"a2","name":"two","username":"u","title":"Two","stats":{"totalUsers":2}}]}}`
	return &http.Response{StatusCode: 200, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(page))}, nil
}

// With Search set the API returns one page and ignores the offset, so the iterator makes one
// request and yields each actor once.
func TestStoreIterateWithSearchFetchesOnePage(t *testing.T) {
	backend := &searchIgnoringOffsetBackend{}
	it := testClient(backend, 0).Store().Iterate(StoreListOptions{Search: Ptr("scraper"), Offset: Ptr(int64(7))}, Ptr(int64(1)))
	var ids []string
	for {
		item, err := it.Next(context.Background())
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if item == nil {
			break
		}
		ids = append(ids, item.ID)
	}
	if strings.Join(ids, ",") != "a1,a2" {
		t.Errorf("expected each actor once, got %v", ids)
	}
	if len(backend.urls) != 1 || !strings.Contains(backend.urls[0], "search=scraper") || strings.Contains(backend.urls[0], "offset=") {
		t.Errorf("expected one request for the search without an offset, got %v", backend.urls)
	}
}

// Without Search the iterator still pages through the Store with the offset.
func TestStoreIterateWithoutSearchAdvancesTheOffset(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{
		{status: 200, body: `{"data":{"total":3,"offset":0,"count":2,"limit":2,"items":[{"id":"a1"},{"id":"a2"}]}}`},
		{status: 200, body: `{"data":{"total":3,"offset":2,"count":1,"limit":2,"items":[{"id":"a3"}]}}`},
		{status: 200, body: `{"data":{"total":3,"offset":3,"count":0,"limit":2,"items":[]}}`},
	}}
	it := testClient(backend, 0).Store().Iterate(StoreListOptions{}, Ptr(int64(2)))
	count := 0
	for {
		item, err := it.Next(context.Background())
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if item == nil {
			break
		}
		count++
	}
	if count != 3 || !strings.Contains(backend.lastURL, "offset=3") {
		t.Errorf("expected 3 actors paged by offset, got %d, last URL %s", count, backend.lastURL)
	}
}

func TestStoreListSendsIncludeInputSchema(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	if _, err := client.Store().List(context.Background(), StoreListOptions{IncludeInputSchema: Ptr(true)}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(backend.lastURL, "includeInputSchema=1") {
		t.Errorf("expected includeInputSchema=1, got %q", backend.lastURL)
	}
	if _, err := client.Store().List(context.Background(), StoreListOptions{}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(backend.lastURL, "includeInputSchema") {
		t.Errorf("includeInputSchema must be omitted when unset, got %q", backend.lastURL)
	}
}
