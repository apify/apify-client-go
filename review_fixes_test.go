package apify

import (
	"context"
	"strings"
	"testing"
	"time"
)

// okData is a constant 200 response with an empty data envelope, reused by the offline
// tests below.
func okData() []mockResponse { return constant(200, `{"data":{}}`) }

// Charge must always send an idempotency-key header so a transport retry cannot double-charge.
func TestChargeSendsIdempotencyKey(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "my-event"})
	if err != nil {
		t.Fatalf("charge: %v", err)
	}
	key := backend.lastHeaders.Get("idempotency-key")
	if key == "" {
		t.Fatal("expected an idempotency-key header to be sent")
	}
	if !strings.HasPrefix(key, "run123-my-event-") {
		t.Fatalf("idempotency key should embed run id and event name, got %q", key)
	}
	if !strings.Contains(backend.lastBody, `"eventName":"my-event"`) {
		t.Fatalf("unexpected charge body: %s", backend.lastBody)
	}
}

// A caller-supplied idempotency key must be used verbatim.
func TestChargeHonorsExplicitIdempotencyKey(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "e", IdempotencyKey: "fixed-key"})
	if err != nil {
		t.Fatalf("charge: %v", err)
	}
	if got := backend.lastHeaders.Get("idempotency-key"); got != "fixed-key" {
		t.Fatalf("expected explicit idempotency key, got %q", got)
	}
}

// GetRecord must default attachment=true (matching the reference client).
func TestGetRecordDefaultsAttachment(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "raw-bytes")}
	client := testClient(backend, 0)

	_, _, err := client.KeyValueStore("store1").GetRecord(context.Background(), "OUTPUT")
	if err != nil {
		t.Fatalf("get record: %v", err)
	}
	if !strings.Contains(backend.lastURL, "attachment=1") {
		t.Fatalf("expected attachment=1 in URL, got %q", backend.lastURL)
	}
}

// BatchAddRequests must split inputs larger than the API's 25-per-batch limit into chunks.
func TestBatchAddRequestsChunks(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"processedRequests":[],"unprocessedRequests":[]}}`)}
	client := testClient(backend, 0)

	requests := make([]RequestQueueRequest, 60) // 60 -> 25 + 25 + 10 = 3 chunks
	for i := range requests {
		requests[i] = RequestQueueRequest{URL: "https://example.com"}
	}
	if _, err := client.RequestQueue("q1").BatchAddRequests(context.Background(), requests, false); err != nil {
		t.Fatalf("batch add: %v", err)
	}
	if backend.calls != 3 {
		t.Fatalf("expected 3 batch calls for 60 requests, got %d", backend.calls)
	}
}

// Run.GetWithWait must forward the waitForFinish query parameter.
func TestRunGetWithWaitForwardsParam(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"r1"}}`)}
	client := testClient(backend, 0)

	secs := int64(30)
	if _, _, err := client.Run("r1").GetWithWait(context.Background(), &secs); err != nil {
		t.Fatalf("get with wait: %v", err)
	}
	if !strings.Contains(backend.lastURL, "waitForFinish=30") {
		t.Fatalf("expected waitForFinish=30 in URL, got %q", backend.lastURL)
	}
}

// GetStreamedLog must open a raw, streaming log request (stream=1 & raw=1).
func TestGetStreamedLogParams(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "live log bytes")}
	client := testClient(backend, 0)

	stream, err := client.Run("r1").GetStreamedLog(context.Background())
	if err != nil {
		t.Fatalf("get streamed log: %v", err)
	}
	defer func() { _ = stream.Close() }()
	if !strings.Contains(backend.lastURL, "stream=1") {
		t.Fatalf("expected stream=1 in URL, got %q", backend.lastURL)
	}
	if !strings.Contains(backend.lastURL, "raw=1") {
		t.Fatalf("expected raw=1 in URL, got %q", backend.lastURL)
	}
}

// notFound is a constant 404 response, used to simulate a resource that never becomes
// available (e.g. database-replica lag on a just-started run that never resolves).
func notFound() []mockResponse {
	return constant(404, `{"error":{"type":"record-not-found","message":"missing"}}`)
}

// A bounded WaitForFinish against a resource that stays 404 must terminate when the budget
// is exhausted and return a descriptive error (not the raw 404), matching the reference
// client's "Waiting for run to finish failed" behaviour.
func TestWaitForFinishBoundedReturnsDescriptiveErrorOn404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	zeroSecs := int64(0)
	_, err := client.Run("r1").WaitForFinish(context.Background(), &zeroSecs)
	if err == nil {
		t.Fatal("expected a descriptive error when the run never becomes available")
	}
	if !strings.Contains(err.Error(), "waiting for run to finish failed") {
		t.Fatalf("expected descriptive wait error, got %q", err.Error())
	}
	// 404 must not be retried by the transport, and the budget=0 path does exactly one fetch.
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 fetch for a zero-budget wait, got %d", backend.calls)
	}
}

// An INDEFINITE WaitForFinish (nil waitSecs) against a permanent 404 must NOT hang forever:
// it polls through 404s on a pure time bound, so a cancelled context unblocks it promptly.
// This is the regression test for the waitForFinish hang bug.
func TestWaitForFinishIndefiniteDoesNotHangOn404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		_, err := client.Run("r1").WaitForFinish(ctx, nil)
		done <- err
	}()

	select {
	case err := <-done:
		// It must return because the context was cancelled (or, in principle, the budget),
		// never spin indefinitely.
		if err == nil {
			t.Fatal("expected a non-nil error (context cancelled), got nil")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("WaitForFinish hung on a permanent 404 during an indefinite wait")
	}
	// Must have polled at least once through the 404 (proving 404 is not treated as fatal).
	if backend.calls < 1 {
		t.Fatalf("expected at least one poll through the 404, got %d", backend.calls)
	}
}

// ListRequests must reject mutually-exclusive options and an invalid Filter before any API call.
func TestListRequestsValidation(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	queue := client.RequestQueue("q1")

	start, cursor := "a", "b"
	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{ExclusiveStartID: &start, Cursor: &cursor}); err == nil {
		t.Fatal("expected error for mutually-exclusive ExclusiveStartID and Cursor")
	}
	bad := "bogus"
	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{Filter: &bad}); err == nil {
		t.Fatal("expected error for invalid Filter value")
	}
	if backend.calls != 0 {
		t.Fatalf("validation must short-circuit before any API call, got %d calls", backend.calls)
	}
}
