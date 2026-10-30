package apify

import (
	"context"
	"strings"
	"testing"
)

func TestStoreListSendsIncludeInputSchema(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	if _, err := client.Store().List(context.Background(), StoreListOptions{IncludeInputSchema: Ptr(true)}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(backend.lastURL, "includeInputSchema=1") {
		t.Errorf("expected includeInputSchema=1, got %q", backend.lastURL)
	}
	if _, err := client.Store().List(context.Background(), StoreListOptions{}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(backend.lastURL, "includeInputSchema") {
		t.Errorf("includeInputSchema must be omitted when unset, got %q", backend.lastURL)
	}
}
