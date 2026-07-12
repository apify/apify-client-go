// Command run_and_last_run_storages starts a Store Actor run, waits for it, then fetches the
// Actor's last run and accesses all three of its default storages.
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/run_and_last_run_storages
package main

import (
	"context"
	"fmt"
	"log"

	apify "github.com/apify/apify-client-go"
	"github.com/apify/apify-client-go/examples/internal/exampleclient"
)

func main() {
	client := exampleclient.New()
	ctx := context.Background()

	waitSecs := int64(120)
	if _, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{}, &waitSecs); err != nil {
		log.Fatalf("run actor: %v", err)
	}

	// Fetch the Actor's last successful run.
	lastRun, ok, err := client.Actor("apify/hello-world").LastRun("SUCCEEDED").Get(ctx)
	if err != nil || !ok {
		log.Fatalf("fetch last run: ok=%v err=%v", ok, err)
	}
	fmt.Printf("Last run %s status %s\n", lastRun.ID, lastRun.Status)

	run := client.Run(lastRun.ID)

	items, err := run.Dataset().ListItems(ctx, apify.DatasetListItemsOptions{})
	if err != nil {
		log.Fatalf("read dataset: %v", err)
	}
	fmt.Printf("Default dataset items: %d\n", items.Count)

	keys, err := run.KeyValueStore().ListKeys(ctx, apify.ListKeysOptions{})
	if err != nil {
		log.Fatalf("read kvs keys: %v", err)
	}
	fmt.Printf("Default key-value store keys: %d\n", len(keys.Items))

	head, err := run.RequestQueue().ListHead(ctx, nil)
	if err != nil {
		log.Fatalf("read queue head: %v", err)
	}
	fmt.Printf("Default request queue head: %d\n", len(head.Items))
}
