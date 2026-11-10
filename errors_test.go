package apify

import (
	"context"
	"testing"
)

// ApiError's Is* status-classification predicates match the response status, and only that
// status's predicate is true.
func TestAPIErrorStatusPredicates(t *testing.T) {
	errorFor := func(status int) *APIError {
		backend := &mockBackend{responses: constant(status, `{"error":{"type":"x","message":"m"}}`)}
		client := testClient(backend, 0)
		_, err := client.Me().MonthlyUsage(context.Background())
		apiErr, ok := AsAPIError(err)
		if !ok {
			t.Fatalf("expected an *APIError for status %d, got %v", status, err)
		}
		return apiErr
	}

	notFoundErr := errorFor(404)
	if !notFoundErr.IsNotFound() {
		t.Error("expected IsNotFound() for a 404")
	}
	for _, predicate := range []struct {
		name string
		ok   bool
	}{
		{"IsInvalidRequest", notFoundErr.IsInvalidRequest()},
		{"IsUnauthorized", notFoundErr.IsUnauthorized()},
		{"IsForbidden", notFoundErr.IsForbidden()},
		{"IsConflict", notFoundErr.IsConflict()},
		{"IsRateLimited", notFoundErr.IsRateLimited()},
		{"IsServerError", notFoundErr.IsServerError()},
	} {
		if predicate.ok {
			t.Errorf("expected %s() to be false for a 404", predicate.name)
		}
	}

	if !errorFor(400).IsInvalidRequest() {
		t.Error("expected IsInvalidRequest() for a 400")
	}
	if !errorFor(401).IsUnauthorized() {
		t.Error("expected IsUnauthorized() for a 401")
	}
	if !errorFor(403).IsForbidden() {
		t.Error("expected IsForbidden() for a 403")
	}
	if !errorFor(409).IsConflict() {
		t.Error("expected IsConflict() for a 409")
	}
	if !errorFor(429).IsRateLimited() {
		t.Error("expected IsRateLimited() for a 429")
	}
	if !errorFor(500).IsServerError() {
		t.Error("expected IsServerError() for a 500")
	}
	if !errorFor(503).IsServerError() {
		t.Error("expected IsServerError() for a 503")
	}
}
