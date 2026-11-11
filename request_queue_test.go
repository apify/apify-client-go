package apify

import (
	"context"
	"strings"
	"testing"
)

// ListRequests must reject mutually-exclusive options and an invalid Filter before any API call.
func TestListRequestsValidation(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	queue := client.RequestQueue("q1")

	start, cursor := "a", "b"
	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{ExclusiveStartID: &start, Cursor: &cursor}); err == nil {
		t.Fatal("expected error for mutually-exclusive ExclusiveStartID and Cursor")
	}
	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{Filter: []string{"bogus"}}); err == nil {
		t.Fatal("expected error for invalid Filter value")
	}
	if backend.calls != 0 {
		t.Fatalf("validation must short-circuit before any API call, got %d calls", backend.calls)
	}
}

// Every RequestQueueClient method that sends a request — including the ones that bypass the
// shared base.go CRUD helpers (ProlongRequestLock, DeleteRequestLock) — must fail fast on a
// queue built with an invalid id, instead of sending a request built from it.
func TestRequestQueueLockMethodsFailFastOnInvalidQueueID(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	queue := client.RequestQueue("..")

	if _, err := queue.ProlongRequestLock(context.Background(), "req1", 30, false); err == nil {
		t.Fatal("expected ProlongRequestLock to fail fast on an invalid queue id")
	}
	if err := queue.DeleteRequestLock(context.Background(), "req1", false); err == nil {
		t.Fatal("expected DeleteRequestLock to fail fast on an invalid queue id")
	}
	if backend.calls != 0 {
		t.Fatalf("an invalid queue id must be rejected before any API call, got %d calls", backend.calls)
	}
}

// ListRequests sends a Filter with several values as one comma-joined query parameter.
func TestListRequestsFilterCommaJoined(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	queue := client.RequestQueue("q1")

	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{
		Filter: []string{RequestFilterLocked, RequestFilterPending},
	}); err != nil {
		t.Fatalf("ListRequests with multi-value Filter: %v", err)
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly one API call, got %d", backend.calls)
	}
	gotURL := backend.lastURL
	if !strings.Contains(gotURL, "filter=locked%2Cpending") && !strings.Contains(gotURL, "filter=locked,pending") {
		t.Fatalf("expected filter sent comma-joined as a single param, got URL: %s", gotURL)
	}
}
