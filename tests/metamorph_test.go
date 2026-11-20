package apify_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	apify "github.com/apify/apify-client-go"
)

// metamorphDockerfile and metamorphSource build an Actor that stores its run token under `TOKEN`
// in its key-value store, then waits for 240 seconds.
const metamorphDockerfile = "FROM apify/actor-node:26\nCOPY . ./\nCMD [\"node\", \"main.js\"]\n"

const metamorphSource = `
const { Actor } = require('apify');
Actor.main(async () => {
    await Actor.setValue('TOKEN', process.env.APIFY_TOKEN, { contentType: 'text/plain' });
    await new Promise((resolve) => setTimeout(resolve, 240_000));
});
`

// TestMetamorphRun starts an Actor that publishes its own run token, then metamorphs the run into
// hello-world with a client that uses that token, since only the run itself can metamorph.
func TestMetamorphRun(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, map[string]any{
		"name": uniqueName("metamorph"),
		"versions": []any{map[string]any{
			"versionNumber": "0.0", "sourceType": "SOURCE_FILES", "buildTag": "latest",
			"sourceFiles": []any{
				map[string]any{"name": "Dockerfile", "format": "TEXT", "content": metamorphDockerfile},
				map[string]any{"name": "main.js", "format": "TEXT", "content": metamorphSource},
			},
		}},
	})
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	// Cleanups run in reverse order: the run goes first, then the Actor.
	t.Cleanup(func() { _ = client.Actor(created.ID).Delete(context.Background()) })

	build, err := client.Actor(created.ID).Build(ctx, "0.0", apify.ActorBuildOptions{})
	if err != nil {
		t.Fatalf("start build: %v", err)
	}
	finished, err := client.Build(build.ID).WaitForFinish(ctx, ptr(int64(300)))
	if err != nil || finished.Status != "SUCCEEDED" {
		t.Fatalf("build: status=%q err=%v", finished.Status, err)
	}

	run, err := client.Actor(created.ID).Start(ctx, nil, apify.ActorStartOptions{
		MemoryMbytes: ptr(int64(256)), TimeoutSecs: ptr(int64(300)),
	})
	if err != nil {
		t.Fatalf("start run: %v", err)
	}
	runClient := client.Run(run.ID)
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		// Aborting a finished run returns it unchanged.
		_, _ = runClient.Abort(cleanupCtx, nil)
		_, _ = runClient.WaitForFinish(cleanupCtx, ptr(int64(120)))
		_ = runClient.Delete(cleanupCtx)
	})

	// A metamorph needs the run's own token, which the Actor stores in its key-value store.
	var runToken string
	deadline := time.Now().Add(2 * time.Minute)
	for runToken == "" {
		record, ok, err := client.KeyValueStore(run.DefaultKeyValueStoreID).GetRecord(ctx, "TOKEN")
		if err != nil {
			t.Fatalf("read the run token: %v", err)
		}
		if ok {
			runToken = string(record.Value)
			break
		}
		current, _, err := runClient.Get(ctx)
		if err != nil || (current.Status != "READY" && current.Status != "RUNNING") {
			t.Fatalf("the run ended as %q without a token: %v", current.Status, err)
		}
		if time.Now().After(deadline) {
			t.Fatal("the run did not store its token in time")
		}
		time.Sleep(2 * time.Second)
	}

	runOwn := apify.NewClient(
		apify.WithToken(runToken),
		apify.WithBaseURL(resolveBaseURL(os.Getenv("APIFY_API_URL"))),
	)
	metamorphed, err := runOwn.Run(run.ID).Metamorph(ctx, "apify/hello-world", map[string]any{"message": "Hello from metamorph!"}, apify.MetamorphOptions{})
	if err != nil {
		t.Fatalf("metamorph: %v", err)
	}
	if metamorphed.ID != run.ID {
		t.Fatalf("metamorph changed the run ID: %q != %q", metamorphed.ID, run.ID)
	}

	// Hello-world replaces the waiting Actor, so the run finishes on its own.
	final, err := runClient.WaitForFinish(ctx, ptr(int64(240)))
	if err != nil || final.Status != "SUCCEEDED" {
		t.Fatalf("the metamorphed run: status=%q err=%v", final.Status, err)
	}
	helloWorld, ok, err := client.Actor("apify/hello-world").Get(ctx)
	if err != nil || !ok {
		t.Fatalf("get hello-world: ok=%v err=%v", ok, err)
	}
	// The run lists the Actors it metamorphed into, in a field the model keeps in Extra.
	var metamorphs []struct {
		ActorID string `json:"actorId"`
	}
	if err := json.Unmarshal(final.Extra["metamorphs"], &metamorphs); err != nil || len(metamorphs) != 1 || metamorphs[0].ActorID != helloWorld.ID {
		t.Fatalf("metamorphs = %s (err=%v), want one entry for %q", final.Extra["metamorphs"], err, helloWorld.ID)
	}
}
