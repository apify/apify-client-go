package apify

import (
	"context"
	"strings"
	"testing"
)

// TestSetStatusMessageRequiresRunID verifies the convenience method errors when it is not
// running inside an Actor run (ACTOR_RUN_ID unset).
func TestSetStatusMessageRequiresRunID(t *testing.T) {
	t.Setenv("ACTOR_RUN_ID", "")
	client := testClient(&mockBackend{responses: constant(200, `{"data":{}}`)}, 0)

	if _, err := client.SetStatusMessage(context.Background(), "hi", SetStatusMessageOptions{}); err == nil {
		t.Fatal("expected an error when ACTOR_RUN_ID is unset")
	}
}

// TestSetStatusMessageBuildsBody verifies the request targets the current run and carries the
// message and the Terminal flag (from the options struct) in the body.
func TestSetStatusMessageBuildsBody(t *testing.T) {
	t.Setenv("ACTOR_RUN_ID", "run123")
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"run123","status":"RUNNING"}}`)}
	client := testClient(backend, 0)

	run, err := client.SetStatusMessage(context.Background(), "almost done", SetStatusMessageOptions{Terminal: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if run.ID != "run123" {
		t.Fatalf("unexpected run: %+v", run)
	}
	if !strings.Contains(backend.lastURL, "/actor-runs/run123") {
		t.Fatalf("expected request to target the current run, got %q", backend.lastURL)
	}
	if !strings.Contains(backend.lastBody, `"statusMessage":"almost done"`) {
		t.Fatalf("body missing status message: %q", backend.lastBody)
	}
	if !strings.Contains(backend.lastBody, `"isStatusMessageTerminal":true`) {
		t.Fatalf("body missing terminal flag: %q", backend.lastBody)
	}
}
