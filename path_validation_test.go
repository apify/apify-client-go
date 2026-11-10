package apify

import (
	"context"
	"testing"
)

// A resource id, record key or request id that is a dot segment (".." or ".") must be rejected
// before any request is sent: a URL parser or proxy can resolve it after the path is built,
// letting it escape to a different endpoint (path traversal).
func TestPathTraversalRejected(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	cases := []struct {
		name string
		call func() error
	}{
		{"resource id", func() error { _, _, err := client.Actor("..").Get(context.Background()); return err }},
		{"record key", func() error {
			_, _, err := client.KeyValueStore("store1").GetRecord(context.Background(), "..")
			return err
		}},
		{"request id", func() error {
			_, _, err := client.RequestQueue("q1").GetRequest(context.Background(), "..")
			return err
		}},
		{"prolong request lock", func() error {
			_, err := client.RequestQueue("q1").ProlongRequestLock(context.Background(), "..", 30, false)
			return err
		}},
		{"delete request lock", func() error {
			return client.RequestQueue("q1").DeleteRequestLock(context.Background(), "..", false)
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := c.call(); err == nil {
				t.Fatal("expected an error for a dot-segment id, got nil")
			}
		})
	}
	if backend.calls != 0 {
		t.Fatalf("a dot-segment id must be rejected before any API call, got %d calls", backend.calls)
	}
}

// An empty resource id must be rejected the same way as a dot segment, not silently read as
// "the whole collection" (e.g. an empty Actor version number).
func TestEmptyIDRejected(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	if _, _, err := client.Actor("some-actor").Version("").Get(context.Background()); err == nil {
		t.Fatal("expected an error for an empty version number, got nil")
	}
	if backend.calls != 0 {
		t.Fatalf("an empty id must be rejected before any API call, got %d calls", backend.calls)
	}
}
