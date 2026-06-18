// Command log_redirection starts an Actor without waiting, then streams its log output to
// stdout in real time (log redirection).
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/log_redirection
package main

import (
	"context"
	"io"
	"log"
	"os"

	apify "github.com/apify/apify-client-go"
)

func main() {
	client := apify.NewClient(os.Getenv("APIFY_TOKEN"))
	ctx := context.Background()

	// Start the Actor and return immediately (do not wait for it to finish).
	run, err := client.Actor("apify/hello-world").Start(ctx, nil, apify.ActorStartOptions{})
	if err != nil {
		log.Fatalf("start actor: %v", err)
	}

	// Open a live (raw) log stream and copy it to stdout as the run produces output.
	stream, err := client.Run(run.ID).GetStreamedLog(ctx)
	if err != nil {
		log.Fatalf("open log stream: %v", err)
	}
	defer stream.Close()

	if _, err := io.Copy(os.Stdout, stream); err != nil {
		log.Fatalf("copy log: %v", err)
	}
}
