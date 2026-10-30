package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

func TestListStore(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Store().List(ctx, apify.StoreListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("Store().List: %v", err)
	}
	if len(page.Items) > 5 {
		t.Fatalf("expected at most 5 items, got %d", len(page.Items))
	}
}

func TestIterateStore(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	it := client.Store().Iterate(apify.StoreListOptions{}, ptr(int64(5)))
	count := 0
	for count < 12 {
		item, err := it.Next(ctx)
		if err != nil {
			t.Fatalf("iterate: %v", err)
		}
		if item == nil {
			break
		}
		if item.ID == "" {
			t.Fatal("expected a non-empty store actor ID")
		}
		count++
	}
	if count < 12 {
		t.Fatalf("expected to iterate at least 12 store actors, got %d", count)
	}
}

// TestListStoreIncludeInputSchema asks the Store to add the compact input schema to every Actor:
// with it set, the API defaults the limit to 10 and accepts nothing above it.
func TestListStoreIncludeInputSchema(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Store().List(ctx, apify.StoreListOptions{Limit: ptr(int64(3)), IncludeInputSchema: ptr(true)})
	if err != nil {
		t.Fatalf("Store().List with includeInputSchema: %v", err)
	}
	if len(page.Items) == 0 {
		t.Fatal("expected at least one Store Actor")
	}
	if len(page.Items) > 10 {
		t.Fatalf("with includeInputSchema the API returns at most 10 Actors, got %d", len(page.Items))
	}
}
