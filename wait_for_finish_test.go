package apify

import (
	"context"
	"errors"
	"testing"
	"time"
)

// notFound is a constant 404 response, used to simulate a resource that never becomes
// available (e.g. database-replica lag on a just-started run that never resolves).
func notFound() []mockResponse {
	return constant(404, `{"error":{"type":"record-not-found","message":"missing"}}`)
}

// A bounded WaitForFinish on a resource that stays 404 ends when the budget runs out and returns
// the 404 as an *APIError.
func TestWaitForFinishBoundedReturnsDescriptiveErrorOn404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	zeroSecs := int64(0)
	_, err := client.Run("r1").WaitForFinish(context.Background(), &zeroSecs)
	if err == nil {
		t.Fatal("expected an error when the run never becomes available")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) || !apiErr.IsNotFound() {
		t.Fatalf("expected a not-found *APIError, got %v (%T)", err, err)
	}
	// 404 must not be retried by the transport, and the budget=0 path does exactly one fetch.
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 fetch for a zero-budget wait, got %d", backend.calls)
	}
}

// WaitForFinish gives up on a resource that keeps answering 404 for the grace period (3s), long
// before the overall wait budget.
func TestWaitForFinishGivesUpAfterGracePeriodOnPersistent404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	generousSecs := int64(30)
	start := time.Now()
	_, err := client.Run("r1").WaitForFinish(context.Background(), &generousSecs)
	elapsed := time.Since(start)

	var apiErr *APIError
	if !errors.As(err, &apiErr) || !apiErr.IsNotFound() {
		t.Fatalf("expected a not-found *APIError, got %v (%T)", err, err)
	}
	// Must stop around the 3s grace period, nowhere near the 30s budget.
	if elapsed < maxWaitForMissingJob || elapsed > maxWaitForMissingJob+2*time.Second {
		t.Fatalf("expected to give up around the %s grace period, took %s", maxWaitForMissingJob, elapsed)
	}
}

// A successful poll in between resets the 404 grace period, so a run/build that goes
// 404 -> briefly visible -> 404 again is not penalized for the earlier 404 streak.
func TestWaitForFinishResetsGracePeriodOnSuccessfulPoll(t *testing.T) {
	notFoundResp := mockResponse{status: 404, body: `{"error":{"type":"record-not-found","message":"missing"}}`}
	runningResp := mockResponse{status: 200, body: `{"data":{"id":"r1","status":"RUNNING"}}`}
	succeededResp := mockResponse{status: 200, body: `{"data":{"id":"r1","status":"SUCCEEDED"}}`}
	backend := &mockBackend{responses: []mockResponse{
		notFoundResp, notFoundResp, notFoundResp, runningResp,
		notFoundResp, notFoundResp, notFoundResp, succeededResp,
	}}
	client := testClient(backend, 0)

	generousSecs := int64(30)
	run, err := client.Run("r1").WaitForFinish(context.Background(), &generousSecs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if run.Status != "SUCCEEDED" {
		t.Fatalf("expected the final SUCCEEDED run, got status %q", run.Status)
	}
	if backend.calls != 8 {
		t.Fatalf("expected exactly 8 polls (both 404 streaks fully consumed), got %d", backend.calls)
	}
}

// A run that was seen once and then answers 404 until the wait budget runs out, which is within
// the grace period, is reported as not found and not as its stale last state.
func TestWaitForFinishReturnsNotFoundWhenBudgetExhaustedWhileCurrently404(t *testing.T) {
	runningResp := mockResponse{status: 200, body: `{"data":{"id":"r1","status":"RUNNING"}}`}
	notFoundResp := mockResponse{status: 404, body: `{"error":{"type":"record-not-found","message":"missing"}}`}
	// Only the first poll succeeds; every poll after that 404s, clamped by mockBackend for as
	// long as the (sub-3s) budget keeps the loop going.
	backend := &mockBackend{responses: []mockResponse{runningResp, notFoundResp}}
	client := testClient(backend, 0)

	oneSec := int64(1)
	_, err := client.Run("r1").WaitForFinish(context.Background(), &oneSec)
	var apiErr *APIError
	if !errors.As(err, &apiErr) || !apiErr.IsNotFound() {
		t.Fatalf("expected a not-found *APIError (not the stale RUNNING run), got %v", err)
	}
	if backend.calls < 2 {
		t.Fatalf("expected at least 2 polls (the success, then at least one 404), got %d", backend.calls)
	}
}

// An INDEFINITE WaitForFinish (nil waitSecs) against a permanent 404 must NOT hang forever:
// it polls through 404s on a pure time bound, so a cancelled context unblocks it promptly.
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

// WaitForFinish's polling request must NOT be bounded by the short tier (it uses
// noRequestTimeout): with an unresponsive backend, it is the caller's own context deadline that
// ends the call, not a tiny configured short tier.
func TestWaitForFinishPollIgnoresShortTier(t *testing.T) {
	client := NewClient(
		WithToken("t"),
		WithHTTPBackend(blockingBackend{}),
		WithMaxRetries(0),
		WithTimeoutShort(10*time.Millisecond),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := client.Run("run1").WaitForFinish(ctx, Ptr(int64(1)))
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected an error once the caller's context deadline is reached")
	}
	// If the 10ms short tier were (wrongly) applied to the poll, this would return almost
	// instantly instead of waiting for the ~150ms caller deadline.
	if elapsed < 100*time.Millisecond {
		t.Fatalf("expected the poll to run until the caller's context deadline (~150ms), returned after %v", elapsed)
	}
}
