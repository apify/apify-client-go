// Command log_redirection runs an Actor and redirects its log to stdout in real time while it
// runs. Call does this by default (to os.Stderr); here the output is set to stdout.
//
// Run with:
//
//	APIFY_TOKEN=<your-token> go run ./examples/log_redirection
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	apify "github.com/apify/apify-client-go"
	"github.com/apify/apify-client-go/examples/internal/exampleclient"
)

func main() {
	client := exampleclient.New()
	ctx := context.Background()

	// Every message of the run's log is printed as the run produces it, prefixed with the Actor
	// name and the run ID, until the run finishes.
	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{
		Log: &apify.StreamedLogOptions{Output: os.Stdout},
	}, nil)
	if err != nil {
		log.Fatalf("call actor: %v", err)
	}
	fmt.Println("run finished with status", run.Status)

	// To redirect the log of a run that was started elsewhere, create the redirection for it and
	// stop it once the run is done. Starting it for a finished run reads the whole log.
	streamed, err := client.Run(run.ID).NewStreamedLog(ctx, apify.StreamedLogOptions{Output: os.Stdout})
	if err != nil {
		log.Fatalf("create log redirection: %v", err)
	}
	if err := streamed.Start(ctx); err != nil {
		log.Fatalf("start log redirection: %v", err)
	}
	if err := streamed.Stop(ctx); err != nil {
		log.Fatalf("stop log redirection: %v", err)
	}
}
