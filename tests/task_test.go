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
