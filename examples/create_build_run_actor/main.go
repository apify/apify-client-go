// Command create_build_run_actor demonstrates the full Actor lifecycle: create a new Actor
// from source files, build it, run it, wait for it to finish, then fetch and print the run
// log.
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/create_build_run_actor
package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	apify "github.com/apify/apify-client-go"
	"github.com/apify/apify-client-go/examples/internal/exampleclient"
)

func main() {
	client := exampleclient.New()
	ctx := context.Background()

	// Actor names are unique per account, so use a random suffix to stay collision-free when
	// this example runs concurrently (e.g. in CI, or alongside the sibling clients' examples on
	// the same test account).
	var suffix [6]byte
	_, _ = rand.Read(suffix[:])
	name := "go-example-" + hex.EncodeToString(suffix[:])
	actorDef := map[string]any{
		"name":     name,
		"isPublic": false,
		"versions": []any{
			map[string]any{
				"versionNumber": "0.0",
				"sourceType":    "SOURCE_FILES",
				"buildTag":      "latest",
				"sourceFiles": []any{
					map[string]any{"name": "Dockerfile", "format": "TEXT", "content": "FROM apify/actor-node:20\nCOPY . ./\nCMD node main.js"},
					map[string]any{"name": "main.js", "format": "TEXT", "content": "console.log('hello from the go client example');"},
				},
			},
		},
	}

	created, err := client.Actors().Create(ctx, actorDef)
	if err != nil {
		log.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()
	fmt.Printf("Created actor %s\n", created.ID)

	actor := client.Actor(created.ID)

	// Build version 0.0 and wait for the build to finish.
	build, err := actor.Build(ctx, "0.0", apify.ActorBuildOptions{})
	if err != nil {
		log.Fatalf("start build: %v", err)
	}
	if _, err := client.Build(build.ID).WaitForFinish(ctx, apify.Ptr(int64(300))); err != nil {
		log.Fatalf("wait for build: %v", err)
	}
	fmt.Printf("Built actor (build %s)\n", build.ID)

	// Run the actor and wait for it to finish.
	run, err := actor.Call(ctx, nil, apify.ActorStartOptions{}, apify.Ptr(int64(120)))
	if err != nil {
		log.Fatalf("run actor: %v", err)
	}
	fmt.Printf("Run %s finished with status %s\n", run.ID, run.Status)

	// Fetch and print the run log.
	logText, ok, err := client.Run(run.ID).Log().Get(ctx)
	if err != nil {
		log.Fatalf("fetch log: %v", err)
	}
	if ok {
		fmt.Println("--- run log ---")
		fmt.Println(logText)
	}
}
