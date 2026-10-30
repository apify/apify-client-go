package apify_test

import (
	"bytes"
	"strings"
	"testing"

	apify "github.com/apify/apify-client-go"
)

// TestCallRawSendsRawBytesInput exercises ActorClient.CallRaw end to end against the real API:
// a raw-bytes input (here, pre-serialized JSON) still runs the Actor to completion.
func TestCallRawSendsRawBytesInput(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	run, err := client.Actor("apify/hello-world").CallRaw(
		ctx,
		[]byte("{}"),
		apify.ActorStartOptions{ContentType: ptr("application/json")},
		ptr(int64(120)),
	)
	if err != nil {
		t.Fatalf("call raw hello-world: %v", err)
	}
	if run.Status != "SUCCEEDED" {
		t.Fatalf("expected SUCCEEDED, got %q", run.Status)
	}
}

// TestNestedRunStorageClientsPropagate404 verifies that a run's default dataset/key-value-store/
// request-queue/log clients (which have no id of their own) surface a 404 as an error instead of
// silently resolving to "absent": the response cannot tell the run apart from the sub-resource as
// what is missing. A record/request lookup by key or id is unaffected, since that is unambiguous.
func TestNestedRunStorageClientsPropagate404(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	const missingRunID = "nonexistent-run-id-xyz"

	if _, _, err := client.Run(missingRunID).Dataset().Get(ctx); err == nil {
		t.Fatal("expected an error fetching the default dataset of a nonexistent run")
	} else if apiErr, ok := apify.AsAPIError(err); !ok || !apiErr.IsNotFound() {
		t.Fatalf("expected a 404 *apify.APIError, got %v", err)
	}

	if _, _, err := client.Run(missingRunID).Log().Get(ctx); err == nil {
		t.Fatal("expected an error fetching the log of a nonexistent run")
	} else if apiErr, ok := apify.AsAPIError(err); !ok || !apiErr.IsNotFound() {
		t.Fatalf("expected a 404 *apify.APIError, got %v", err)
	}

	// Unambiguous lookups on the same nonexistent run still resolve to "absent", not an error.
	if _, present, err := client.Run(missingRunID).KeyValueStore().GetRecord(ctx, "OUTPUT"); err != nil {
		t.Fatalf("expected a missing record to resolve to absent, got error: %v", err)
	} else if present {
		t.Fatal("expected present=false for a record on a nonexistent run's key-value store")
	}
}

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

	logText, ok, err := client.Run(run.ID).Log().Get(ctx)
	if err != nil || !ok || logText == "" {
		t.Fatalf("get run log: ok=%v err=%v len=%d", ok, err, len(logText))
	}

	// GetBuffer is the same request as Get, decoded differently: the raw bytes must match the
	// string content byte-for-byte.
	logBuffer, ok, err := client.Run(run.ID).Log().GetBuffer(ctx)
	if err != nil || !ok {
		t.Fatalf("get run log buffer: ok=%v err=%v", ok, err)
	}
	if string(logBuffer) != logText {
		t.Fatalf("Log().Get and Log().GetBuffer disagree: %q vs %q", logText, string(logBuffer))
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

// TestCallRedirectsTheRunLog runs hello-world through Call and checks that its log was redirected
// to the configured output, prefixed with the Actor name and the run ID.
func TestCallRedirectsTheRunLog(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var out bytes.Buffer
	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{
		Log: &apify.StreamedLogOptions{Output: &out},
	}, ptr(int64(120)))
	if err != nil {
		t.Fatalf("call hello-world: %v", err)
	}
	if want := "hello-world runId:" + run.ID + " -> "; !strings.Contains(out.String(), want) {
		t.Fatalf("expected the redirected log to carry the prefix %q, got %q", want, out.String())
	}
}

// TestRunAndTaskStoragesAndLogCanBeRead covers the simple GETs of a finished run: the run-scoped
// default storages (dataset, key-value store, request queue), the log by run id, and the default
// build, webhook listings and last run of the Actor and of a task.
func TestRunAndTaskStoragesAndLogCanBeRead(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	actor := client.Actor("apify/hello-world")
	run, err := actor.Call(ctx, nil, apify.ActorStartOptions{Log: &apify.StreamedLogOptions{Disabled: true}}, ptr(int64(120)))
	if err != nil {
		t.Fatalf("call hello-world: %v", err)
	}
	runClient := client.Run(run.ID)

	if logText, ok, err := client.Log(run.ID).Get(ctx); err != nil || !ok || logText == "" {
		t.Fatalf("Log(%s).Get: ok=%v err=%v len=%d", run.ID, ok, err, len(logText))
	}

	if ds, ok, err := runClient.Dataset().Get(ctx); err != nil || !ok || ds.ID != run.DefaultDatasetID {
		t.Fatalf("run dataset Get: %+v ok=%v err=%v", ds, ok, err)
	}
	if _, err := runClient.Dataset().GetStatistics(ctx); err != nil {
		t.Fatalf("run dataset GetStatistics: %v", err)
	}
	if store, ok, err := runClient.KeyValueStore().Get(ctx); err != nil || !ok || store.ID != run.DefaultKeyValueStoreID {
		t.Fatalf("run key-value store Get: %+v ok=%v err=%v", store, ok, err)
	}
	if _, err := runClient.KeyValueStore().ListKeys(ctx, apify.ListKeysOptions{}); err != nil {
		t.Fatalf("run key-value store ListKeys: %v", err)
	}
	if queue, ok, err := runClient.RequestQueue().Get(ctx); err != nil || !ok || queue.ID != run.DefaultRequestQueueID {
		t.Fatalf("run request queue Get: %+v ok=%v err=%v", queue, ok, err)
	}
	if _, err := runClient.RequestQueue().ListHead(ctx, nil); err != nil {
		t.Fatalf("run request queue ListHead: %v", err)
	}
	if _, err := runClient.RequestQueue().ListRequests(ctx, apify.ListRequestsOptions{}); err != nil {
		t.Fatalf("run request queue ListRequests: %v", err)
	}

	if _, err := actor.DefaultBuild(ctx, nil); err != nil {
		t.Fatalf("DefaultBuild: %v", err)
	}
	if _, err := actor.Webhooks().List(ctx, apify.ListOptions{Limit: ptr(int64(5))}); err != nil {
		t.Fatalf("actor webhooks List: %v", err)
	}

	task, err := client.Tasks().Create(ctx, taskDef(uniqueName("task-last-run")))
	if err != nil {
		t.Fatalf("create task: %v", err)
	}
	defer func() { _ = client.Task(task.ID).Delete(ctx) }()
	taskRun, err := client.Task(task.ID).Call(ctx, nil, apify.TaskStartOptions{}, ptr(int64(120)))
	if err != nil {
		t.Fatalf("call task: %v", err)
	}
	if _, err := client.Task(task.ID).Webhooks().List(ctx, apify.ListOptions{Limit: ptr(int64(5))}); err != nil {
		t.Fatalf("task webhooks List: %v", err)
	}
	last, ok, err := client.Task(task.ID).LastRun("SUCCEEDED").Get(ctx)
	if err != nil || !ok || last.ID != taskRun.ID {
		t.Fatalf("task LastRun: %+v ok=%v err=%v", last, ok, err)
	}
	if logText, ok, err := client.Task(task.ID).LastRun("SUCCEEDED").Log().Get(ctx); err != nil || !ok || logText == "" {
		t.Fatalf("task last run log: ok=%v err=%v len=%d", ok, err, len(logText))
	}
	// The storages of a filtered last run are those of the run the filter resolves to.
	if ds, ok, err := client.Task(task.ID).LastRun("SUCCEEDED").Dataset().Get(ctx); err != nil || !ok || ds.ID != taskRun.DefaultDatasetID {
		t.Fatalf("task last run dataset: %+v ok=%v err=%v", ds, ok, err)
	}
}

// TestModelURLFieldsAreNormalized checks, against the real API, that the URL fields of a run and
// of its storages come back in their normalized form: an absolute URL with a path.
func TestModelURLFieldsAreNormalized(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{Log: &apify.StreamedLogOptions{Disabled: true}}, ptr(int64(120)))
	if err != nil {
		t.Fatalf("call hello-world: %v", err)
	}
	got, ok, err := client.Run(run.ID).Get(ctx)
	if err != nil || !ok {
		t.Fatalf("get run: ok=%v err=%v", ok, err)
	}
	for name, value := range map[string]string{"consoleUrl": got.ConsoleURL, "containerUrl": got.ContainerURL} {
		if value == "" {
			continue // the API does not return the field for this run
		}
		if !strings.HasPrefix(value, "https://") || strings.Count(value, "/") < 3 {
			t.Errorf("%s is not a normalized absolute URL: %q", name, value)
		}
	}
	if got.ConsoleURL == "" {
		t.Error("the API is expected to return the consoleUrl of a run")
	}

	ds, ok, err := client.Dataset(run.DefaultDatasetID).Get(ctx)
	if err != nil || !ok {
		t.Fatalf("get dataset: ok=%v err=%v", ok, err)
	}
	if !strings.HasPrefix(ds.ConsoleURL, "https://") || ds.ItemsPublicURL == "" {
		t.Errorf("unexpected dataset URLs: %q, %q", ds.ConsoleURL, ds.ItemsPublicURL)
	}
	store, ok, err := client.KeyValueStore(run.DefaultKeyValueStoreID).Get(ctx)
	if err != nil || !ok {
		t.Fatalf("get key-value store: ok=%v err=%v", ok, err)
	}
	if !strings.HasPrefix(store.ConsoleURL, "https://") || store.KeysPublicURL == "" || store.RecordsPublicURL == "" {
		t.Errorf("unexpected key-value store URLs: %+v", store)
	}
}

// TestIterateRunDatasetItems pushes items into the default dataset of a finished run, then reads
// them back through RunClient.IterateDatasetItems, which must also catch the rows past a lagging
// itemCount once the run is finished.
func TestIterateRunDatasetItems(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{Log: &apify.StreamedLogOptions{Disabled: true}}, ptr(int64(120)))
	if err != nil {
		t.Fatalf("call hello-world: %v", err)
	}
	runClient := client.Run(run.ID)
	const total = 7
	items := make([]map[string]any, total)
	for i := range items {
		items[i] = map[string]any{"n": i}
	}
	if err := runClient.Dataset().PushItems(ctx, items); err != nil {
		t.Fatalf("push items: %v", err)
	}

	// hello-world pushes an item of its own, which has no "n" and is left out below.
	type item struct {
		N *int `json:"n"`
	}
	it := apify.IterateRunDatasetItems[item](runClient, apify.RunIterateDatasetItemsOptions{
		ChunkSize:        ptr(int64(3)),
		PollIntervalSecs: ptr(int64(1)),
	})
	var got []int
	for {
		next, err := it.Next(ctx)
		if err != nil {
			t.Fatalf("iterate run dataset items: %v", err)
		}
		if next == nil {
			break
		}
		if next.N != nil {
			got = append(got, *next.N)
		}
	}
	if len(got) != total {
		t.Fatalf("expected the %d pushed items, got %v", total, got)
	}
	for i, n := range got {
		if n != i {
			t.Fatalf("items out of order: %v", got)
		}
	}

	// A Limit caps the rows scanned.
	limited := runClient.IterateDatasetItems(apify.RunIterateDatasetItemsOptions{
		Items: apify.DatasetListItemsOptions{Limit: ptr(int64(3))},
	})
	count := 0
	for {
		next, err := limited.Next(ctx)
		if err != nil {
			t.Fatalf("iterate with offset and limit: %v", err)
		}
		if next == nil {
			break
		}
		count++
	}
	if count != 3 {
		t.Fatalf("expected 3 items with limit 3, got %d", count)
	}
}

// TestRunListDeploymentFilter lists the Actor Standby runs of a pool through the account-wide
// and Actor-scoped run collections: the filter is accepted (the result may well be empty).
func TestRunListDeploymentFilter(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	filter := apify.RunListOptions{
		Status:     []string{"READY", "RUNNING"},
		Deployment: []string{"SINGLE_TENANT", "MULTI_TENANT_FREE_USERS"},
	}
	if _, err := client.Runs().List(ctx, apify.ListOptions{Limit: ptr(int64(5))}, filter); err != nil {
		t.Fatalf("Runs().List with a deployment filter: %v", err)
	}
	if _, err := client.Actor("apify/hello-world").Runs().List(ctx, apify.ListOptions{Limit: ptr(int64(5))}, filter); err != nil {
		t.Fatalf("Actor.Runs().List with a deployment filter: %v", err)
	}
}
