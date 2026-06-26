package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

func TestListRuns(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Runs().List(ctx, apify.ListOptions{Limit: ptr(int64(5))}, apify.RunListOptions{})
	if err != nil {
		t.Fatalf("Runs().List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

// TestRunActorAndReadOutputs runs the public apify/hello-world Actor, waits for it to
// finish, then reads its log, default dataset, and default key-value store.
func TestRunActorAndReadOutputs(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{}, ptr(int64(120)))
	if err != nil {
		t.Fatalf("call hello-world: %v", err)
	}
	if run.Status != "SUCCEEDED" {
		t.Fatalf("expected SUCCEEDED, got %q", run.Status)
	}

	if _, ok, err := client.Run(run.ID).Get(ctx); err != nil || !ok {
		t.Fatalf("get run: ok=%v err=%v", ok, err)
	}

	if logText, ok, err := client.Run(run.ID).Log().Get(ctx); err != nil || !ok || logText == "" {
		t.Fatalf("get run log: ok=%v err=%v len=%d", ok, err, len(logText))
	}

	if _, err := client.Run(run.ID).Dataset().ListItems(ctx, apify.DatasetListItemsOptions{}); err != nil {
		t.Fatalf("run dataset list: %v", err)
	}

	if _, _, err := client.Run(run.ID).KeyValueStore().GetRecord(ctx, "OUTPUT"); err != nil {
		t.Fatalf("run kvs get record: %v", err)
	}
}

// TestLastRunAccess runs hello-world then fetches its last succeeded run via the convenience
// accessor.
func TestLastRunAccess(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	if _, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{}, ptr(int64(120))); err != nil {
		t.Fatalf("call hello-world: %v", err)
	}

	lastRun, ok, err := client.Actor("apify/hello-world").LastRun("SUCCEEDED").Get(ctx)
	if err != nil || !ok {
		t.Fatalf("last run: ok=%v err=%v", ok, err)
	}
	if lastRun.Status != "SUCCEEDED" {
		t.Fatalf("expected last succeeded run, got %q", lastRun.Status)
	}

	// The run was started via the API, so filtering the last run by both status and origin
	// must still resolve it. This exercises the origin filter on LastRunWithOptions.
	lastRunByOrigin, ok, err := client.Actor("apify/hello-world").
		LastRunWithOptions(apify.LastRunOptions{Status: "SUCCEEDED", Origin: "API"}).Get(ctx)
	if err != nil || !ok {
		t.Fatalf("last run by origin: ok=%v err=%v", ok, err)
	}
	if lastRunByOrigin.Status != "SUCCEEDED" {
		t.Fatalf("expected last succeeded run filtered by origin, got %q", lastRunByOrigin.Status)
	}
}
