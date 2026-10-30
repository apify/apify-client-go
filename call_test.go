package apify

import (
	"context"
	"testing"
	"time"
)

// noLogRedirection turns the log redirection of Call off, so a test counts only the requests of
// the call itself.
var noLogRedirection = &StreamedLogOptions{Disabled: true}

// Call has no client-side timeout of its own: a start slower than the medium tier must neither be
// aborted nor repeated (which would start a second, billed run). Start itself is bound by it.
func TestCallStartsWithoutATimeout(t *testing.T) {
	run := `{"data":{"id":"run1","status":"SUCCEEDED"}}`
	t.Run("Actor Call", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		got, err := slowClient(backend).Actor("a").Call(context.Background(), map[string]any{}, ActorStartOptions{Log: noLogRedirection}, Ptr(int64(0)))
		if err != nil || got.ID != "run1" {
			t.Fatalf("Call: %+v, %v", got, err)
		}
		if n := backend.calls.Load(); n != 2 { // the start and one poll
			t.Fatalf("expected the start once and one poll, got %d requests", n)
		}
	})
	t.Run("Actor CallRaw", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		if _, err := slowClient(backend).Actor("a").CallRaw(context.Background(), []byte("x"), ActorStartOptions{Log: noLogRedirection}, Ptr(int64(0))); err != nil {
			t.Fatalf("CallRaw: %v", err)
		}
		if n := backend.calls.Load(); n != 2 {
			t.Fatalf("expected the start once and one poll, got %d requests", n)
		}
	})
	t.Run("Task Call", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		if _, err := slowClient(backend).Task("t").Call(context.Background(), nil, TaskStartOptions{}, Ptr(int64(0))); err != nil {
			t.Fatalf("Task Call: %v", err)
		}
		if n := backend.calls.Load(); n != 2 {
			t.Fatalf("expected the start once and one poll, got %d requests", n)
		}
	})
	t.Run("Start stays bound by the medium tier", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		if _, err := slowClient(backend).Actor("a").Start(context.Background(), nil, ActorStartOptions{}); err == nil {
			t.Fatal("expected Start to time out")
		}
	})
	t.Run("ctx still bounds Call", func(t *testing.T) {
		backend := &slowBackend{delay: time.Minute, body: run}
		ctx, cancel := context.WithTimeout(context.Background(), 3*tinyTier)
		defer cancel()
		started := time.Now()
		if _, err := slowClient(backend).Actor("a").Call(ctx, nil, ActorStartOptions{Log: noLogRedirection}, nil); err == nil {
			t.Fatal("expected an error once ctx is done")
		}
		if time.Since(started) > 5*time.Second {
			t.Fatal("Call did not stop at the ctx deadline")
		}
	})
}
