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

	it := client.Store().Iterate(apify.StoreListOptions{Limit: ptr(int64(5))})
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
