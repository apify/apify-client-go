package apify

import (
	"context"
	"encoding/json"
	"testing"
)

func TestPushItemsAcceptsOnlyObjectsAndArraysOfObjects(t *testing.T) {
	type item struct {
		A int `json:"a"`
	}
	valid := map[string]any{
		"map":                 map[string]any{"a": 1},
		"struct":              item{A: 1},
		"struct pointer":      &item{A: 1},
		"slice of structs":    []item{{A: 1}, {A: 2}},
		"slice of maps":       []map[string]any{{"a": 1}},
		"empty slice":         []item{},
		"slice of any":        []any{map[string]any{"a": 1}, item{A: 2}},
		"raw JSON object":     rawJSON(`{"a":1}`),
		"raw JSON array":      rawJSON(`[{"a":1},{"a":2}]`),
		"padded raw JSON obj": rawJSON(" \n{\"a\":1} "),
	}
	for name, items := range valid {
		t.Run("accepts "+name, func(t *testing.T) {
			backend := &mockBackend{responses: okData()}
			if err := testClient(backend, 0).Dataset("d").PushItems(context.Background(), items); err != nil {
				t.Fatalf("PushItems rejected a valid shape: %v", err)
			}
			if backend.calls != 1 {
				t.Fatalf("expected 1 request, got %d", backend.calls)
			}
		})
	}

	invalid := map[string]any{
		"string":                "[1,2]",
		"[]byte":                []byte(`{"a":1}`),
		"number":                1,
		"bool":                  true,
		"nil":                   nil,
		"slice of strings":      []string{"a", "b"},
		"slice of numbers":      []int{1, 2},
		"mixed slice":           []any{map[string]any{"a": 1}, "b"},
		"slice of slices":       [][]int{{1}},
		"raw JSON string":       rawJSON(`"a"`),
		"raw JSON null element": rawJSON(`[{"a":1},null]`),
	}
	for name, items := range invalid {
		t.Run("rejects "+name, func(t *testing.T) {
			backend := &mockBackend{responses: okData()}
			if err := testClient(backend, 0).Dataset("d").PushItems(context.Background(), items); err == nil {
				t.Fatal("PushItems accepted a shape the API rejects")
			}
			if backend.calls != 0 {
				t.Fatalf("nothing may be sent for an invalid shape, got %d requests", backend.calls)
			}
		})
	}
}

func TestPushItemsRejectsAnInvalidDatasetIDBeforeSending(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	if err := testClient(backend, 0).Dataset("").PushItems(context.Background(), map[string]any{"a": 1}); err == nil {
		t.Fatal("expected an error for an empty dataset ID")
	}
	if backend.calls != 0 {
		t.Fatalf("nothing may be sent for an invalid ID, got %d requests", backend.calls)
	}
}

func rawJSON(s string) json.RawMessage { return json.RawMessage(s) }
