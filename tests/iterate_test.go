package apify_test

import (
	"context"
	"testing"
	"time"

	apify "github.com/apify/apify-client-go"
)

// iterFindAttempts / iterFindBackoff bound how long iterateFindIDs waits for a freshly
// created resource to surface in a collection listing. Collection list indexes are eventually
// consistent, so a resource can be missing from the first listing right after creation; the
// helper re-iterates a fresh iterator until it converges.
const (
	iterFindAttempts = 8
	iterFindBackoff  = time.Second
)

// iterateFindIDs repeatedly drains a fresh iterator (from makeIter) and checks that every id
// in want is yielded. Within one pass it stops early once all wanted ids are seen and gives up
// after safetyCap items so a shared account with many pre-existing resources cannot make the
// test iterate unboundedly. Across passes it retries with backoff to absorb list-index
// eventual consistency. A small per-page Limit combined with several wanted ids exercises
// paging across more than one page.
func iterateFindIDs[T any](t *testing.T, ctx context.Context, makeIter func() *apify.ListIterator[T], idOf func(*T) string, want []string, safetyCap int) {
	t.Helper()
	missing := make(map[string]bool, len(want))
	for _, id := range want {
		missing[id] = true
	}
	for attempt := 0; attempt < iterFindAttempts; attempt++ {
		it := makeIter()
		seen := 0
		for {
			item, err := it.Next(ctx)
			if err != nil {
				t.Fatalf("iterate: %v", err)
			}
			if item == nil {
				break
			}
			seen++
			if id := idOf(item); missing[id] {
				delete(missing, id)
				if len(missing) == 0 {
					return
				}
			}
			if safetyCap > 0 && seen >= safetyCap {
				break
			}
		}
		if len(missing) == 0 {
			return
		}
		time.Sleep(iterFindBackoff)
	}
	t.Fatalf("iteration did not yield all created resources; missing %v", missing)
}

// smallPage returns list options that page one item at a time, newest first, so the just-created
// resources appear near the front and multiple pages are traversed.
func smallPageDesc() apify.ListOptions {
	return apify.ListOptions{Desc: ptr(true)}
}

func TestIterateActors(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var want []string
	for i := 0; i < 3; i++ {
		created, err := client.Actors().Create(ctx, minimalActor(uniqueName("iter-actor")))
		if err != nil {
			t.Fatalf("create actor %d: %v", i, err)
		}
		defer func(id string) { _ = client.Actor(id).Delete(ctx) }(created.ID)
		want = append(want, created.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.Actor] {
		return client.Actors().Iterate(apify.ActorListOptions{My: ptr(true), Desc: ptr(true)}, ptr(int64(1)))
	}, func(a *apify.Actor) string { return a.ID }, want, 500)
}

func TestIterateTasks(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var want []string
	for i := 0; i < 3; i++ {
		task, err := client.Tasks().Create(ctx, taskDef(uniqueName("iter-task")))
		if err != nil {
			t.Fatalf("create task %d: %v", i, err)
		}
		defer func(id string) { _ = client.Task(id).Delete(ctx) }(task.ID)
		want = append(want, task.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.Task] {
		return client.Tasks().Iterate(smallPageDesc(), ptr(int64(1)))
	}, func(x *apify.Task) string { return x.ID }, want, 500)
}

func TestIterateSchedules(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var want []string
	for i := 0; i < 3; i++ {
		sch, err := client.Schedules().Create(ctx, scheduleDef(uniqueName("iter-sch")))
		if err != nil {
			t.Fatalf("create schedule %d: %v", i, err)
		}
		defer func(id string) { _ = client.Schedule(id).Delete(ctx) }(sch.ID)
		want = append(want, sch.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.Schedule] {
		return client.Schedules().Iterate(smallPageDesc(), ptr(int64(1)))
	}, func(s *apify.Schedule) string { return s.ID }, want, 500)
}

func TestIterateWebhooks(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var want []string
	for i := 0; i < 3; i++ {
		wh, err := client.Webhooks().Create(ctx, webhookDef("https://example.com/iter-webhook"))
		if err != nil {
			t.Fatalf("create webhook %d: %v", i, err)
		}
		defer func(id string) { _ = client.Webhook(id).Delete(ctx) }(wh.ID)
		want = append(want, wh.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.Webhook] {
		return client.Webhooks().Iterate(smallPageDesc(), ptr(int64(1)))
	}, func(w *apify.Webhook) string { return w.ID }, want, 500)
}

func TestIterateDatasets(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var want []string
	for i := 0; i < 3; i++ {
		ds, err := client.Datasets().GetOrCreate(ctx, uniqueName("iter-ds"))
		if err != nil {
			t.Fatalf("create dataset %d: %v", i, err)
		}
		defer func(id string) { _ = client.Dataset(id).Delete(ctx) }(ds.ID)
		want = append(want, ds.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.Dataset] {
		return client.Datasets().Iterate(apify.StorageListOptions{Desc: ptr(true)}, ptr(int64(1)))
	}, func(d *apify.Dataset) string { return d.ID }, want, 500)
}

func TestIterateKeyValueStores(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var want []string
	for i := 0; i < 3; i++ {
		store, err := client.KeyValueStores().GetOrCreate(ctx, uniqueName("iter-kvs"))
		if err != nil {
			t.Fatalf("create store %d: %v", i, err)
		}
		defer func(id string) { _ = client.KeyValueStore(id).Delete(ctx) }(store.ID)
		want = append(want, store.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.KeyValueStore] {
		return client.KeyValueStores().Iterate(apify.StorageListOptions{Desc: ptr(true)}, ptr(int64(1)))
	}, func(s *apify.KeyValueStore) string { return s.ID }, want, 500)
}

func TestIterateRequestQueues(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	var want []string
	for i := 0; i < 3; i++ {
		rq, err := client.RequestQueues().GetOrCreate(ctx, uniqueName("iter-rq"))
		if err != nil {
			t.Fatalf("create queue %d: %v", i, err)
		}
		defer func(id string) { _ = client.RequestQueue(id).Delete(ctx) }(rq.ID)
		want = append(want, rq.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.RequestQueue] {
		return client.RequestQueues().Iterate(apify.StorageListOptions{Desc: ptr(true)}, ptr(int64(1)))
	}, func(q *apify.RequestQueue) string { return q.ID }, want, 500)
}

func TestIterateActorVersions(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("iter-ver")))
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()
	actor := client.Actor(created.ID)

	// The actor is created with version "0.0"; add a second version so paging spans >1 page.
	if _, err := actor.Versions().Create(ctx, map[string]any{
		"versionNumber": "0.1",
		"sourceType":    "SOURCE_FILES",
		"buildTag":      "latest",
		"sourceFiles":   []any{},
	}); err != nil {
		t.Fatalf("create version: %v", err)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.ActorVersion] {
		return actor.Versions().Iterate(smallPageDesc(), ptr(int64(1)))
	}, func(v *apify.ActorVersion) string { return v.VersionNumber }, []string{"0.0", "0.1"}, 100)
}

func TestIterateActorEnvVars(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("iter-env")))
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()
	envVars := client.Actor(created.ID).Version("0.0").EnvVars()

	for _, name := range []string{"VAR_A", "VAR_B", "VAR_C"} {
		if _, err := envVars.Create(ctx, apify.ActorEnvVar{Name: name, Value: "v"}); err != nil {
			t.Fatalf("create env var %s: %v", name, err)
		}
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.ActorEnvVar] {
		return envVars.Iterate()
	}, func(e *apify.ActorEnvVar) string { return e.Name }, []string{"VAR_A", "VAR_B", "VAR_C"}, 100)
}

func TestIterateRuns(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	// Start two quick runs of the public hello-world Actor (Start does not wait for finish),
	// so the Actor-scoped run collection has more than one item to page through.
	actor := client.Actor("apify/hello-world")
	var want []string
	for i := 0; i < 2; i++ {
		run, err := actor.Start(ctx, nil, apify.ActorStartOptions{})
		if err != nil {
			t.Fatalf("start run %d: %v", i, err)
		}
		want = append(want, run.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.ActorRun] {
		return actor.Runs().Iterate(smallPageDesc(), apify.RunListOptions{}, ptr(int64(1)))
	}, func(r *apify.ActorRun) string { return r.ID }, want, 200)
}

func TestIterateBuilds(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	// The public hello-world Actor has at least one build; iterate its build collection one
	// build per page and confirm the iterator yields real builds with non-empty IDs.
	it := client.Actor("apify/hello-world").Builds().Iterate(smallPageDesc(), ptr(int64(1)))
	count := 0
	for count < 50 {
		build, err := it.Next(ctx)
		if err != nil {
			t.Fatalf("iterate: %v", err)
		}
		if build == nil {
			break
		}
		if build.ID == "" {
			t.Fatal("expected a non-empty build ID")
		}
		count++
	}
	if count == 0 {
		t.Fatal("expected at least one build for apify/hello-world")
	}
}

func TestIterateWebhookDispatches(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	wh, err := client.Webhooks().Create(ctx, webhookDef("https://example.com/iter-dispatch"))
	if err != nil {
		t.Fatalf("create webhook: %v", err)
	}
	defer func() { _ = client.Webhook(wh.ID).Delete(ctx) }()
	webhook := client.Webhook(wh.ID)

	// Each Test() dispatches the webhook, creating a dispatch record to page through.
	var want []string
	for i := 0; i < 2; i++ {
		dispatch, err := webhook.Test(ctx)
		if err != nil {
			t.Fatalf("test webhook %d: %v", i, err)
		}
		want = append(want, dispatch.ID)
	}

	iterateFindIDs(t, ctx, func() *apify.ListIterator[apify.WebhookDispatch] {
		return webhook.Dispatches().Iterate(smallPageDesc(), ptr(int64(1)))
	}, func(d *apify.WebhookDispatch) string { return d.ID }, want, 200)
}

func TestIterateDatasetItems(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	ds, err := client.Datasets().GetOrCreate(ctx, uniqueName("iter-items"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.Dataset(ds.ID).Delete(ctx) }()
	dataset := client.Dataset(ds.ID)

	const total = 5
	items := make([]map[string]any, total)
	for i := 0; i < total; i++ {
		items[i] = map[string]any{"n": i}
	}
	if err := dataset.PushItems(ctx, items); err != nil {
		t.Fatalf("push items: %v", err)
	}

	// The X-Apify-Pagination-Total header can lag right after a push, and offset-based
	// iteration relies on it to page. Wait for the total to converge before iterating.
	settled := false
	for attempt := 0; attempt < 20; attempt++ {
		page, err := dataset.ListItems(ctx, apify.DatasetListItemsOptions{Limit: ptr(int64(1))})
		if err != nil {
			t.Fatalf("list items: %v", err)
		}
		if page.Total >= total {
			settled = true
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if !settled {
		t.Fatalf("dataset item total did not converge to %d", total)
	}

	// Page two items at a time so iteration crosses multiple pages.
	it := dataset.IterateItems(apify.DatasetListItemsOptions{}, ptr(int64(2)))
	count := 0
	for {
		item, err := it.Next(ctx)
		if err != nil {
			t.Fatalf("iterate: %v", err)
		}
		if item == nil {
			break
		}
		count++
	}
	if count != total {
		t.Fatalf("expected to iterate %d items, got %d", total, count)
	}

	// Limit is a cap on the total number of items yielded (not the page size), matching the
	// reference client. With Limit=3 and a page size of 2, iteration must stop at exactly 3.
	const cap = 3
	capped := dataset.IterateItems(apify.DatasetListItemsOptions{Limit: ptr(int64(cap))}, ptr(int64(2)))
	capCount := 0
	for {
		item, err := capped.Next(ctx)
		if err != nil {
			t.Fatalf("iterate (capped): %v", err)
		}
		if item == nil {
			break
		}
		capCount++
	}
	if capCount != cap {
		t.Fatalf("expected Limit to cap iteration at %d items, got %d", cap, capCount)
	}
}
