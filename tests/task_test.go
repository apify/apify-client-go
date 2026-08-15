package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

func taskDef(name string) map[string]any {
	return map[string]any{
		"actId":   "apify/hello-world",
		"name":    name,
		"options": map[string]any{"build": "latest", "memoryMbytes": 256, "timeoutSecs": 60},
		"input":   map[string]any{"message": "hello"},
	}
}

func TestListTasks(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Tasks().List(ctx, apify.ListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestGetTask(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	task, err := client.Tasks().Create(ctx, taskDef(uniqueName("task-get")))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Task(task.ID).Delete(ctx) }()

	got, ok, err := client.Task(task.ID).Get(ctx)
	if err != nil || !ok || got.ID != task.ID {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
}

func TestTaskCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	task, err := client.Tasks().Create(ctx, taskDef(uniqueName("task-crud")))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Task(task.ID).Delete(ctx) }()
	tc := client.Task(task.ID)

	if _, ok, err := tc.Get(ctx); err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	if _, err := tc.UpdateInput(ctx, map[string]any{"message": "updated"}); err != nil {
		t.Fatalf("update input: %v", err)
	}
	if _, ok, err := tc.GetInput(ctx); err != nil || !ok {
		t.Fatalf("get input: ok=%v err=%v", ok, err)
	}
	if _, err := tc.Update(ctx, map[string]any{"name": uniqueName("task-renamed")}); err != nil {
		t.Fatalf("update: %v", err)
	}
	if _, err := tc.Runs().List(ctx, apify.ListOptions{}, apify.RunListOptions{}); err != nil {
		t.Fatalf("runs list: %v", err)
	}
}

// TestTaskPublishUnpublish exercises Publish/Unpublish against a task whose Actor
// (apify/hello-world) is not owned by the test account.
//
// Publish (like Unpublish) requires write permission to both the task and its Actor, so it is
// expected to fail (400 for the missing PublicConfig, or 403 for the unowned Actor - the server
// may reject on either ground first). Unpublish is expected to succeed here because the task is
// already unpublished (created with IsPublic unset/false) and the API treats setting isPublic to
// its current value as a no-op, not because Unpublish has a smaller permission requirement than
// Publish - it does not create/verify an actually-published state on an unowned Actor, which is
// not achievable via this API (publishing always requires Actor write permission).
func TestTaskPublishUnpublish(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	task, err := client.Tasks().Create(ctx, taskDef(uniqueName("task-publish")))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Task(task.ID).Delete(ctx) }()
	tc := client.Task(task.ID)

	if _, err := tc.Publish(ctx); err == nil {
		t.Fatal("publish: expected an error for an unowned Actor without PublicConfig, got nil")
	} else if apiErr, ok := apify.AsAPIError(err); !ok {
		t.Fatalf("publish: expected an *apify.APIError, got %T: %v", err, err)
	} else if apiErr.StatusCode != 400 && apiErr.StatusCode != 403 {
		t.Fatalf("publish: expected status 400 or 403, got %d: %v", apiErr.StatusCode, apiErr)
	}

	unpublished, err := tc.Unpublish(ctx)
	if err != nil {
		t.Fatalf("unpublish: %v", err)
	}
	if unpublished.IsPublic != nil && *unpublished.IsPublic {
		t.Fatalf("unpublish: expected IsPublic to not be true, got %v", *unpublished.IsPublic)
	}
}
