# Webhooks

Webhooks notify external services when events occur. Access the webhook collection with
`client.Webhooks()`, a single webhook with `client.Webhook(id)`, dispatches with
`client.WebhookDispatches()` / `client.WebhookDispatch(id)`.

## Webhook collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions) (PaginationList[Webhook], error)` | List webhooks. |
| `Create(ctx, definition any) (Webhook, error)` | Create a webhook. |

An Actor's or task's webhooks are also listable via `client.Actor(id).Webhooks()` /
`client.Task(id).Webhooks()`.

## Single webhook

| Method | Description |
| --- | --- |
| `Get(ctx) (Webhook, bool, error)` | Fetch the webhook. |
| `Update(ctx, newFields any) (Webhook, error)` | Update the webhook. |
| `Delete(ctx) error` | Delete the webhook. |
| `Test(ctx) (WebhookDispatch, error)` | Dispatch the webhook immediately. |
| `Dispatches() *WebhookDispatchCollectionClient` | The webhook's dispatches. |

## Webhook dispatches

| Method | Description |
| --- | --- |
| `WebhookDispatches().List(ctx, ListOptions)` | List dispatches. |
| `WebhookDispatch(id).Get(ctx) (WebhookDispatch, bool, error)` | Fetch a dispatch. |

```go
wh, err := client.Webhooks().Create(ctx, map[string]any{
	"isAdHoc":    true,
	"eventTypes": []string{"ACTOR.RUN.SUCCEEDED"},
	"condition":  map[string]any{"actorRunId": "RUN_ID"},
	"requestUrl": "https://example.com/webhook",
})
if err != nil {
	log.Fatal(err)
}
dispatch, err := client.Webhook(wh.ID).Test(ctx)
```
