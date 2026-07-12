// Command run_store_actor runs an existing Apify Store Actor, waits for it to finish, and
// reads items from its default dataset.
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/run_store_actor
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

	// Start the public hello-world Actor and wait up to 120 seconds for it to finish.
	waitSecs := int64(120)
	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{}, &waitSecs)
	if err != nil {
		log.Fatalf("failed to run actor: %v", err)
	}
	fmt.Printf("Run %s finished with status %s\n", run.ID, run.Status)

	// Read items from the run's default dataset.
	page, err := client.Dataset(run.DefaultDatasetID).ListItems(ctx, apify.DatasetListItemsOptions{})
	if err != nil {
		log.Fatalf("failed to read dataset: %v", err)
	}
	fmt.Printf("Default dataset contains %d items\n", page.Total)
}
