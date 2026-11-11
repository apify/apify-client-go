package apify

import (
	"context"
	"testing"
)

// Get and Delete on a resource with its own ID treat a 404 of any `type` as absent.
func TestGetSwallowsAnyNotFoundType(t *testing.T) {
	backend := &mockBackend{responses: constant(404, `{"error":{"type":"some-other-error","message":"gone"}}`)}
	client := testClient(backend, 0)

	_, present, err := client.Actor("missing").Get(context.Background())
	if err != nil {
		t.Fatalf("expected a 404 of any type to resolve to absent, got error: %v", err)
	}
	if present {
		t.Fatal("expected present=false for a 404 response")
	}
}

// Get/Delete on a client reached through a fixed sub-path with no id of its own (e.g. a run's
// default dataset) must NOT swallow a 404: the response cannot tell the run apart from the
// dataset as what is missing.
func TestNestedResourceGetPropagates404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	_, _, err := client.Run("run1").Dataset().Get(context.Background())
	if err == nil {
		t.Fatal("expected a 404 on a nested dataset client to propagate as an error")
	}
	apiErr, ok := AsAPIError(err)
	if !ok || apiErr.StatusCode != 404 {
		t.Fatalf("expected a 404 *APIError, got %v", err)
	}

	err = client.Run("run1").Dataset().Delete(context.Background())
	if err == nil {
		t.Fatal("expected a 404 on a nested dataset client's Delete to propagate as an error")
	}
}

// GetRecord/RecordExists on a nested key-value store (no id of its own) must still swallow a
// 404: unlike the store's own Get/Delete, a record lookup by key is never ambiguous.
func TestNestedKeyValueStoreRecordLookupStillSwallows404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)
	store := client.Run("run1").KeyValueStore()

	record, present, err := store.GetRecord(context.Background(), "OUTPUT")
	if err != nil {
		t.Fatalf("expected a missing record to resolve to absent, got error: %v", err)
	}
	if present || record != nil {
		t.Fatalf("expected present=false and a nil record, got present=%v record=%v", present, record)
	}

	exists, err := store.RecordExists(context.Background(), "OUTPUT")
	if err != nil {
		t.Fatalf("expected RecordExists to resolve false, got error: %v", err)
	}
	if exists {
		t.Fatal("expected exists=false for a 404 response")
	}
}

// DatasetClient.GetStatistics, TaskClient.GetInput and ScheduleClient.GetLog must propagate a
// 404 as an error rather than an "absent" result: it always means the parent resource is gone.
func TestSingletonSubResourcesPropagate404(t *testing.T) {
	cases := []struct {
		name string
		call func(c *ApifyClient) error
	}{
		{"dataset statistics", func(c *ApifyClient) error { _, err := c.Dataset("ds1").GetStatistics(context.Background()); return err }},
		{"task input", func(c *ApifyClient) error { _, err := c.Task("task1").GetInput(context.Background()); return err }},
		{"schedule log", func(c *ApifyClient) error { _, err := c.Schedule("sch1").GetLog(context.Background()); return err }},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			backend := &mockBackend{responses: notFound()}
			client := testClient(backend, 0)
			if err := c.call(client); err == nil {
				t.Fatal("expected a 404 to propagate as an error")
			}
		})
	}
}
