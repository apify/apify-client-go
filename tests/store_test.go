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
	for _, item := range page.Items {
		if item.Stats.TotalUsers <= 0 {
			t.Fatalf("expected the Store statistics of %s to count at least its owner, got %+v", item.ID, item.Stats)
		}
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

// TestIterateStoreSearch iterates a search, which the API answers with a single page: each Actor
// comes once, and the iteration ends.
func TestIterateStoreSearch(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	it := client.Store().Iterate(apify.StoreListOptions{Search: ptr("web scraper")}, ptr(int64(5)))
	seen := map[string]bool{}
	for {
		item, err := it.Next(ctx)
		if err != nil {
			t.Fatalf("iterate: %v", err)
		}
		if item == nil {
			break
		}
		if seen[item.ID] {
			t.Fatalf("Actor %s was yielded twice", item.ID)
		}
		seen[item.ID] = true
	}
	if len(seen) == 0 {
		t.Fatal("expected the search to find Actors")
	}
}
