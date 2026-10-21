package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

// createTestWebhook creates a persistent webhook conditioned on a task that only this test owns,
// so it can fire only for runs of that task, which nothing else starts. The API rejects an ad hoc
// webhook whose run does not exist, which is why a persistent one is used. Both resources are
// deleted on t.Cleanup with a fresh context, so a test that fails or times out leaves nothing
// behind.
func createTestWebhook(t *testing.T, client *apify.ApifyClient, url string) apify.Webhook {
	t.Helper()
	ctx, cancel := testContext(t)
	defer cancel()

	task, err := client.Tasks().Create(ctx, taskDef(uniqueName("webhook-task")))
	if err != nil {
		t.Fatalf("create webhook condition task: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := testContext(t)
		defer cleanupCancel()
		_ = client.Task(task.ID).Delete(cleanupCtx)
	})

	wh, err := client.Webhooks().Create(ctx, map[string]any{
		"eventTypes": []string{"ACTOR.RUN.SUCCEEDED"},
		"condition":  map[string]any{"actorTaskId": task.ID},
		"requestUrl": url,
	})
	if err != nil {
		t.Fatalf("create webhook: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := testContext(t)
		defer cleanupCancel()
		_ = client.Webhook(wh.ID).Delete(cleanupCtx)
	})
	return wh
}

func TestListWebhooks(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Webhooks().List(ctx, apify.ListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestListWebhookDispatches(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.WebhookDispatches().List(ctx, apify.ListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestGetWebhook(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	wh := createTestWebhook(t, client, "https://example.com/webhook")

	got, ok, err := client.Webhook(wh.ID).Get(ctx)
	if err != nil || !ok || got.ID != wh.ID {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
}

func TestGetWebhookDispatch(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	wh := createTestWebhook(t, client, "https://example.com/webhook")

	dispatch, err := client.Webhook(wh.ID).Test(ctx)
	if err != nil {
		t.Fatalf("test webhook: %v", err)
	}
	got, ok, err := client.WebhookDispatch(dispatch.ID).Get(ctx)
	if err != nil || !ok || got.ID != dispatch.ID {
		t.Fatalf("get dispatch: ok=%v err=%v", ok, err)
	}
}

func TestWebhookCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	wh := createTestWebhook(t, client, "https://example.com/webhook")
	webhook := client.Webhook(wh.ID)

	if _, ok, err := webhook.Get(ctx); err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	updated, err := webhook.Update(ctx, map[string]any{"requestUrl": "https://example.com/updated"})
	if err != nil {
		t.Fatalf("update: %v", err)
	}
	if updated.RequestURL != "https://example.com/updated" {
		t.Fatalf("request url not updated: %+v", updated)
	}
	if _, err := webhook.Dispatches().List(ctx, apify.ListOptions{}); err != nil {
		t.Fatalf("dispatches list: %v", err)
	}
	if _, err := webhook.Test(ctx); err != nil {
		t.Fatalf("test: %v", err)
	}
}
