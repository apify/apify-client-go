// Command storages demonstrates creating, writing to, and reading from all three Apify
// storage types: datasets, key-value stores, and request queues.
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/storages
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	apify "github.com/apify/apify-client-go"
	"github.com/apify/apify-client-go/examples/internal/exampleclient"
)

func main() {
	client := exampleclient.New()
	ctx := context.Background()

	datasetExample(ctx, client)
	keyValueStoreExample(ctx, client)
	requestQueueExample(ctx, client)
}

func datasetExample(ctx context.Context, client *apify.ApifyClient) {
	ds, err := client.Datasets().GetOrCreate(ctx, "")
	if err != nil {
		log.Fatalf("create dataset: %v", err)
	}
	defer func() { _ = client.Dataset(ds.ID).Delete(ctx) }()

	dataset := client.Dataset(ds.ID)
	if err := dataset.PushItems(ctx, []map[string]any{{"name": "Alice"}, {"name": "Bob"}}); err != nil {
		log.Fatalf("push items: %v", err)
	}
	page, err := dataset.ListItems(ctx, apify.DatasetListItemsOptions{})
	if err != nil {
		log.Fatalf("list items: %v", err)
	}
	// page.Count is the number of items on this page (not the dataset total, which the API
	// updates asynchronously and can briefly lag right after a push).
	fmt.Printf("Dataset %s has %d items on this page\n", ds.ID, page.Count)
}

func keyValueStoreExample(ctx context.Context, client *apify.ApifyClient) {
	store, err := client.KeyValueStores().GetOrCreate(ctx, "")
	if err != nil {
		log.Fatalf("create store: %v", err)
	}
	defer func() { _ = client.KeyValueStore(store.ID).Delete(ctx) }()

	kvs := client.KeyValueStore(store.ID)
	if err := kvs.SetRecordJSON(ctx, "OUTPUT", map[string]any{"answer": 42}); err != nil {
		log.Fatalf("set record: %v", err)
	}
	rec, ok, err := kvs.GetRecord(ctx, "OUTPUT")
	if err != nil || !ok {
		log.Fatalf("get record: ok=%v err=%v", ok, err)
	}
	fmt.Printf("Key-value store %s OUTPUT = %s\n", store.ID, string(rec.Value))
}

func requestQueueExample(ctx context.Context, client *apify.ApifyClient) {
	rq, err := client.RequestQueues().GetOrCreate(ctx, "")
	if err != nil {
		log.Fatalf("create queue: %v", err)
	}
	defer func() { _ = client.RequestQueue(rq.ID).Delete(ctx) }()

	queue := client.RequestQueue(rq.ID)
	if _, err := queue.AddRequest(ctx, apify.RequestQueueRequest{URL: "https://example.com", UniqueKey: "example"}, false); err != nil {
		log.Fatalf("add request: %v", err)
	}
	head, err := queue.ListHead(ctx, nil)
	if err != nil {
		log.Fatalf("list head: %v", err)
	}
	fmt.Printf("Request queue %s head has %d requests\n", rq.ID, len(head.Items))

	// ListRequests accepts a multi-value Filter (an array of the enum values "locked"/"pending",
	// sent comma-joined, with union semantics). Use the exported constants for type-safety. It
	// returns the raw API response, which you unmarshal into your own type.
	raw, err := queue.ListRequests(ctx, apify.ListRequestsOptions{
		Filter: []string{apify.RequestFilterPending, apify.RequestFilterLocked},
	})
	if err != nil {
		log.Fatalf("list requests: %v", err)
	}
	var listed struct {
		Items []apify.RequestQueueRequest `json:"items"`
		// NextCursor is the opaque pagination cursor; feed it back via
		// ListRequestsOptions.Cursor to fetch the next page manually. It is empty on the last
		// page. (For most use cases prefer the typed PaginateRequests iterator, which handles
		// this loop for you.)
		NextCursor string `json:"nextCursor"`
	}
	if err := json.Unmarshal(raw, &listed); err != nil {
		log.Fatalf("decode listed requests: %v", err)
	}
	fmt.Printf("Request queue %s listed %d pending/locked requests\n", rq.ID, len(listed.Items))

	// Manual pagination: feed nextCursor back through ListRequestsOptions.Cursor to get the
	// next page. (Cursor and ExclusiveStartID are mutually exclusive.)
	if listed.NextCursor != "" {
		if _, err := queue.ListRequests(ctx, apify.ListRequestsOptions{Cursor: &listed.NextCursor}); err != nil {
			log.Fatalf("list requests next page: %v", err)
		}
	}
}
