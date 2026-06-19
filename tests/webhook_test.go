package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

func webhookDef(url string) map[string]any {
	return map[string]any{
		"isAdHoc":    true,
		"eventTypes": []string{"ACTOR.RUN.SUCCEEDED"},
		"condition":  map[string]any{"actorRunId": "ZZZZZZZZZZZZZZZZZ"},
		"requestUrl": url,
	}
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

	wh, err := client.Webhooks().Create(ctx, webhookDef("https://example.com/webhook"))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Webhook(wh.ID).Delete(ctx) }()

	got, ok, err := client.Webhook(wh.ID).Get(ctx)
	if err != nil || !ok || got.ID != wh.ID {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
}

func TestGetWebhookDispatch(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	wh, err := client.Webhooks().Create(ctx, webhookDef("https://example.com/webhook"))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Webhook(wh.ID).Delete(ctx) }()

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

	wh, err := client.Webhooks().Create(ctx, webhookDef("https://example.com/webhook"))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Webhook(wh.ID).Delete(ctx) }()
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
