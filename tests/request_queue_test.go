package apify_test

import (
	"fmt"
	"testing"

	apify "github.com/apify/apify-client-go"
)

func TestListRequestQueues(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.RequestQueues().List(ctx, apify.StorageListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestGetRequestQueue(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	rq, err := client.RequestQueues().GetOrCreate(ctx, uniqueName("rq-get"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.RequestQueue(rq.ID).Delete(ctx) }()

	got, ok, err := client.RequestQueue(rq.ID).Get(ctx)
	if err != nil || !ok || got.ID != rq.ID {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
}

func TestRequestQueueCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	rq, err := client.RequestQueues().GetOrCreate(ctx, uniqueName("rq-crud"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.RequestQueue(rq.ID).Delete(ctx) }()
	queue := client.RequestQueue(rq.ID)

	if _, ok, err := queue.Get(ctx); err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}

	info, err := queue.AddRequest(ctx, apify.RequestQueueRequest{
		URL:       "https://example.com",
		UniqueKey: "example",
		Method:    "GET",
		UserData:  []byte(`{"label":"DETAIL"}`),
	}, false)
	if err != nil {
		t.Fatalf("add request: %v", err)
	}
	if info.RequestID == "" {
		t.Fatal("expected a request ID")
	}

	got, ok, err := queue.GetRequest(ctx, info.RequestID)
	if err != nil || !ok || got.URL != "https://example.com" {
		t.Fatalf("get request: ok=%v err=%v %+v", ok, err, got)
	}

	head, err := queue.ListHead(ctx, ptr(int64(10)))
	if err != nil {
		t.Fatalf("list head: %v", err)
	}
	if len(head.Items) == 0 {
		t.Fatal("expected at least one request at the head")
	}

	if _, err := queue.Update(ctx, map[string]any{"name": uniqueName("rq-renamed")}); err != nil {
		t.Fatalf("update: %v", err)
	}
	if err := queue.DeleteRequest(ctx, info.RequestID); err != nil {
		t.Fatalf("delete request: %v", err)
	}
}

// TestRequestQueuePaginateMultiplePages validates cursor-based pagination across pages.
func TestRequestQueuePaginateMultiplePages(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	rq, err := client.RequestQueues().GetOrCreate(ctx, uniqueName("rq-page"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.RequestQueue(rq.ID).Delete(ctx) }()
	queue := client.RequestQueue(rq.ID)

	const total = 5
	for i := 0; i < total; i++ {
		url := fmt.Sprintf("https://example.com/%d", i)
		if _, err := queue.AddRequest(ctx, apify.RequestQueueRequest{URL: url, UniqueKey: url}, false); err != nil {
			t.Fatalf("add request %d: %v", i, err)
		}
	}

	seen := map[string]int{}
	it := queue.PaginateRequests(ptr(int64(2))) // small page size forces multiple pages
	for {
		req, err := it.Next(ctx)
		if err != nil {
			t.Fatalf("iterate: %v", err)
		}
		if req == nil {
			break
		}
		seen[req.URL]++
	}
	if len(seen) != total {
		t.Fatalf("expected %d distinct requests, saw %d (%v)", total, len(seen), seen)
	}
	for url, count := range seen {
		if count != 1 {
			t.Fatalf("request %s returned %d times (expected once)", url, count)
		}
	}
}

// TestRequestQueueBatchAddRequests exercises the chunked batch-add path with more than the
// 25-per-call API limit, verifying all requests are accepted across chunks.
func TestRequestQueueBatchAddRequests(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	rq, err := client.RequestQueues().GetOrCreate(ctx, uniqueName("rq-batch"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.RequestQueue(rq.ID).Delete(ctx) }()
	queue := client.RequestQueue(rq.ID)

	const total = 30 // > 25, so the client must split into multiple chunks
	requests := make([]apify.RequestQueueRequest, total)
	for i := 0; i < total; i++ {
		url := fmt.Sprintf("https://batch.example.com/%d", i)
		requests[i] = apify.RequestQueueRequest{URL: url, UniqueKey: url}
	}
	result, err := queue.BatchAddRequests(ctx, requests, false)
	if err != nil {
		t.Fatalf("batch add: %v", err)
	}
	if len(result.ProcessedRequests) != total {
		t.Fatalf("expected %d processed requests, got %d (unprocessed=%d)",
			total, len(result.ProcessedRequests), len(result.UnprocessedRequests))
	}
}

// TestRequestQueueLockLifecycle exercises the lock operations end-to-end.
func TestRequestQueueLockLifecycle(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	rq, err := client.RequestQueues().GetOrCreate(ctx, uniqueName("rq-lock"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.RequestQueue(rq.ID).Delete(ctx) }()

	// A stable client key is required to unlock the client's own locks.
	queue := client.RequestQueue(rq.ID).WithClientKey("go-test-client-key")

	info, err := queue.AddRequest(ctx, apify.RequestQueueRequest{URL: "https://lock.example.com", UniqueKey: "lock"}, false)
	if err != nil {
		t.Fatalf("add request: %v", err)
	}
	if _, err := queue.ListRequests(ctx, apify.ListRequestsOptions{}); err != nil {
		t.Fatalf("list requests: %v", err)
	}
	// Exercise the multi-value `filter` query param (array of enum locked|pending, sent
	// comma-joined) end-to-end so the spec-faithful serialization is covered against the API.
	if _, err := queue.ListRequests(ctx, apify.ListRequestsOptions{
		Filter: []string{apify.RequestFilterLocked, apify.RequestFilterPending},
	}); err != nil {
		t.Fatalf("list requests with filter: %v", err)
	}
	if _, err := queue.ListAndLockHead(ctx, 60, ptr(int64(10))); err != nil {
		t.Fatalf("list and lock head: %v", err)
	}
	if _, err := queue.ProlongRequestLock(ctx, info.RequestID, 30, false); err != nil {
		t.Fatalf("prolong lock: %v", err)
	}
	if err := queue.DeleteRequestLock(ctx, info.RequestID, false); err != nil {
		t.Fatalf("delete lock: %v", err)
	}
	if _, err := queue.UnlockRequests(ctx); err != nil {
		t.Fatalf("unlock requests: %v", err)
	}
}
