// Command public_build_no_token demonstrates the unauthenticated client. A client built with
// NewClient and no WithToken option can still call the endpoints that require no API token —
// here it resolves and fetches the default build of a public Actor, which needs only the
// public Actor ID (GET /v2/actors/{actorId}/builds/default is a token-free endpoint).
//
// Run with (no token needed):
//
//	go run ./examples/public_build_no_token
package main

import (
	"context"
	"fmt"
	"log"

	apify "github.com/apify/apify-client-go"
)

// publicActor is a well-known public Actor whose default build is readable without a token.
const publicActor = "apify/hello-world"

func main() {
	// No WithToken: an unauthenticated client, usable for token-free endpoints only.
	client := apify.NewClient()
	ctx := context.Background()

	// DefaultBuild resolves the Actor's default build by public Actor ID, so no build ID has
	// to be known in advance and no token is required.
	buildClient, err := client.Actor(publicActor).DefaultBuild(ctx, nil)
	if err != nil {
		log.Fatalf("failed to resolve default build of %s: %v", publicActor, err)
	}

	build, ok, err := buildClient.Get(ctx)
	if err != nil {
		log.Fatalf("failed to fetch build: %v", err)
	}
	if !ok {
		log.Fatalf("default build of %s not found", publicActor)
	}

	fmt.Printf("Fetched the default build of %s without a token (build ID: %s)\n", publicActor, build.ID)
}
