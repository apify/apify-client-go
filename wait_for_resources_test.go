package apify

import (
	"context"
	"errors"
	"testing"
	"time"
)

// memoryLimitExceeded is a constant 402 response of type actor-memory-limit-exceeded, the
// error the API rejects a start with while the account has no free memory for the run.
func memoryLimitExceeded() []mockResponse {
	return constant(402, `{"error":{"type":"actor-memory-limit-exceeded","message":"not enough memory"}}`)
}

// Without WaitForResources set, the first resource-limit rejection is returned immediately.
func TestWaitForResourcesDisabledReturnsFirstError(t *testing.T) {
	backend := &mockBackend{responses: memoryLimitExceeded()}
	client := testClient(backend, 0)

	_, err := client.Actor("me/actor").Start(context.Background(), nil, ActorStartOptions{})
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Type != "actor-memory-limit-exceeded" {
		t.Fatalf("expected the raw resource-limit error, got %v", err)
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt, got %d", backend.calls)
	}
}

// WaitForResources: 0 makes a single attempt (the deadline is already reached after it), like
// disabling retrying.
func TestWaitForResourcesZeroSecondsMakesSingleAttempt(t *testing.T) {
	backend := &mockBackend{responses: memoryLimitExceeded()}
	client := testClient(backend, 0)

	zero := int64(0)
	_, err := client.Actor("me/actor").Start(context.Background(), nil, ActorStartOptions{WaitForResources: &zero})
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Type != "actor-memory-limit-exceeded" {
		t.Fatalf("expected the resource-limit error, got %v", err)
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt, got %d", backend.calls)
	}
}

// A non-resource-limit error is never retried, however WaitForResources is set.
func TestWaitForResourcesNeverRetriesOtherErrors(t *testing.T) {
	backend := &mockBackend{responses: constant(400, `{"error":{"type":"invalid-input","message":"bad input"}}`)}
	client := testClient(backend, 0)

	_, err := client.Actor("me/actor").Start(context.Background(), nil, ActorStartOptions{WaitForResources: Ptr(WaitForResourcesIndefinitely)})
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Type != "invalid-input" {
		t.Fatalf("expected the invalid-input error returned immediately, got %v", err)
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt (no retry for a non-resource-limit error), got %d", backend.calls)
	}
}

// A negative WaitForResources (other than WaitForResourcesIndefinitely) is rejected before any
// attempt is made.
func TestWaitForResourcesRejectsNegativeValue(t *testing.T) {
	backend := &mockBackend{responses: memoryLimitExceeded()}
	client := testClient(backend, 0)

	negative := int64(-5)
	_, err := client.Actor("me/actor").Start(context.Background(), nil, ActorStartOptions{WaitForResources: &negative})
	if err == nil {
		t.Fatal("expected a validation error for a negative WaitForResources")
	}
	if backend.calls != 0 {
		t.Fatalf("expected validation to short-circuit before any attempt, got %d calls", backend.calls)
	}
}

// With a bounded WaitForResources, a start rejected for lack of resources is retried and
// succeeds once the rejection stops.
func TestWaitForResourcesRetriesThenSucceeds(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{
		{status: 402, body: `{"error":{"type":"concurrent-runs-limit-exceeded","message":"too many runs"}}`},
		{status: 200, body: `{"data":{"id":"run1","status":"READY"}}`},
	}}
	client := testClient(backend, 0)

	start := time.Now()
	run, err := client.Actor("me/actor").Start(context.Background(), nil, ActorStartOptions{WaitForResources: Ptr(int64(1))})
	if err != nil {
		t.Fatalf("expected the retried start to succeed, got %v", err)
	}
	if run.ID != "run1" {
		t.Fatalf("expected the started run, got %+v", run)
	}
	if backend.calls != 2 {
		t.Fatalf("expected exactly 2 attempts (one rejection, one success), got %d", backend.calls)
	}
	if elapsed := time.Since(start); elapsed < 900*time.Millisecond {
		t.Fatalf("expected the retry to wait out roughly the 1s budget, took %s", elapsed)
	}
}

// A bounded WaitForResources that never stops being rejected throws the last rejection once the
// budget elapses, rather than retrying indefinitely.
func TestWaitForResourcesBoundedThrowsLastErrorAfterDeadline(t *testing.T) {
	backend := &mockBackend{responses: memoryLimitExceeded()}
	client := testClient(backend, 0)

	_, err := client.Actor("me/actor").Start(context.Background(), nil, ActorStartOptions{WaitForResources: Ptr(int64(1))})
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Type != "actor-memory-limit-exceeded" {
		t.Fatalf("expected the last resource-limit error once the budget elapses, got %v", err)
	}
	if backend.calls < 2 {
		t.Fatalf("expected at least 2 attempts (initial + after the cooldown), got %d", backend.calls)
	}
}

// TaskClient.Start retries exactly like ActorClient.Start; this is a thin smoke test for the
// shared wiring rather than a repeat of the full behavior matrix above.
func TestTaskWaitForResourcesRetriesThenSucceeds(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{
		{status: 402, body: `{"error":{"type":"actor-memory-limit-exceeded","message":"not enough memory"}}`},
		{status: 200, body: `{"data":{"id":"run1","status":"READY"}}`},
	}}
	client := testClient(backend, 0)

	run, err := client.Task("my-task").Start(context.Background(), nil, TaskStartOptions{WaitForResources: Ptr(int64(1))})
	if err != nil {
		t.Fatalf("expected the retried task start to succeed, got %v", err)
	}
	if run.ID != "run1" {
		t.Fatalf("expected the started run, got %+v", run)
	}
	if backend.calls != 2 {
		t.Fatalf("expected exactly 2 attempts, got %d", backend.calls)
	}
}
