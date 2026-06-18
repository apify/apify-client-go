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
	"os"

	apify "github.com/apify/apify-client-go"
)

func main() {
	client := apify.NewClient(os.Getenv("APIFY_TOKEN"))
	ctx := context.Background()

	// Iterate the store lazily, fetching pages of 5 on demand.
	limit := int64(5)
	it := client.Store().Iterate(apify.StoreListOptions{Limit: &limit})

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
