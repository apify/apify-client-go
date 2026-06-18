// Command get_account fetches and prints the current user's account details.
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/get_account
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

	user, ok, err := client.Me().Get(context.Background())
	if err != nil {
		log.Fatalf("failed to fetch account: %v", err)
	}
	if !ok {
		log.Fatal("current user not found")
	}

	fmt.Printf("Account ID: %s\n", user.ID)
	fmt.Printf("Username:   %s\n", user.Username)
}
