package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

func TestListBuilds(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Builds().List(ctx, apify.ListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("Builds().List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

// TestBuildActorFlow creates an Actor, builds it, waits for the build to finish, then reads
// the build, its log, and its OpenAPI definition.
func TestBuildActorFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("build")))
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()

	build, err := client.Actor(created.ID).Build(ctx, "0.0", apify.ActorBuildOptions{})
	if err != nil {
		t.Fatalf("start build: %v", err)
	}

	finished, err := client.Build(build.ID).WaitForFinish(ctx, ptr(int64(300)))
	if err != nil {
		t.Fatalf("wait for finish: %v", err)
	}
	if !finished.IsTerminal() {
		t.Fatalf("build did not finish: status=%q", finished.Status)
	}

	if _, ok, err := client.Build(build.ID).Get(ctx); err != nil || !ok {
		t.Fatalf("get build: ok=%v err=%v", ok, err)
	}
	if _, _, err := client.Build(build.ID).Log().Get(ctx); err != nil {
		t.Fatalf("build log: %v", err)
	}
	// The OpenAPI definition may or may not be present depending on the build; only assert
	// that the call itself succeeds.
	if _, _, err := client.Build(build.ID).GetOpenAPIDefinition(ctx); err != nil {
		t.Fatalf("openapi definition: %v", err)
	}
}
