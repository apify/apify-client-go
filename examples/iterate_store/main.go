// Command iterate_store lazily iterates over Actors in the Apify Store using the convenience
// iterator, printing the first few.
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/iterate_store
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

	// Iterate the store lazily, fetching pages of 5 on demand (chunkSize). The options' Limit
	// would cap the total number of Actors yielded; here it is left unset so iteration would
	// cover the whole store, and the loop below stops after the first few.
	chunkSize := int64(5)
	it := client.Store().Iterate(apify.StoreListOptions{}, &chunkSize)

	const want = 10
	for i := 0; i < want; i++ {
		item, err := it.Next(ctx)
		if err != nil {
			log.Fatalf("iterate store: %v", err)
		}
		if item == nil {
			break // store exhausted
		}
		title := item.Title
		if title == "" {
			title = item.Name
		}
		fmt.Printf("%2d. %s (%s)\n", i+1, title, item.ID)
	}
}
