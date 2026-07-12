# Webhooks

Webhooks notify external services when events occur. Access the webhook collection with
`client.Webhooks()`, a single webhook with `client.Webhook(id)`, dispatches with
`client.WebhookDispatches()` / `client.WebhookDispatch(id)`.

## Webhook collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions) (PaginationList[Webhook], error)` | List webhooks. |
| `Iterate(ListOptions, chunkSize *int64) *ListIterator[Webhook]` | Lazy iterator over matching webhooks. `Limit` caps the total yielded; `chunkSize` is the page size. |
| `Create(ctx, definition any) (Webhook, error)` | Create a webhook. |

An Actor's or task's webhooks are also listable via `client.Actor(id).Webhooks()` /
`client.Task(id).Webhooks()`.

`Create` (and `Update`) take a free-form JSON object (`any`) that mirrors the webhook shape in
the Apify API, so it can carry any field the API accepts. The commonly used input keys are:

| Key | Type | Meaning |
|---|---|---|
| `eventTypes` | `[]string` | Events that trigger the webhook (see [Event types](#event-types)). Required. |
| `requestUrl` | `string` | URL the webhook posts to. Required. |
| `condition` | `object` | Scope the webhook, e.g. `{"actorId": "..."}`, `{"actorTaskId": "..."}`, or `{"actorRunId": "..."}`. |
| `isAdHoc` | `bool` | `true` for a one-off webhook bound to a single run/build via `condition`. |
| `payloadTemplate` | `string` | Optional template string for the dispatched payload. |
| `description` | `string` | Optional human-readable description. |

### Event types

A webhook's `eventTypes` is a list drawn from the closed `WebhookEventType` enum (12 values):

| Build events | Run events | Other |
|---|---|---|
| `ACTOR.BUILD.CREATED` | `ACTOR.RUN.CREATED` | `TEST` |
| `ACTOR.BUILD.SUCCEEDED` | `ACTOR.RUN.SUCCEEDED` | |
| `ACTOR.BUILD.FAILED` | `ACTOR.RUN.FAILED` | |
| `ACTOR.BUILD.ABORTED` | `ACTOR.RUN.ABORTED` | |
| `ACTOR.BUILD.TIMED_OUT` | `ACTOR.RUN.TIMED_OUT` | |
| | `ACTOR.RUN.RESURRECTED` | |

The same values apply to the ad-hoc `ActorStartOptions.Webhooks` element (see
[actors.md](actors.md)).

### `Webhook` and `WebhookDispatch` fields

`Webhook`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique webhook ID. |
| `UserID` | `string` | ID of the user who owns the webhook. |
| `RequestURL` | `string` | URL the webhook posts to. |
| `EventTypes` | `[]string` | Events that trigger the webhook. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

`WebhookDispatch`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique dispatch ID. |
| `WebhookID` | `string` | ID of the webhook that produced this dispatch. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

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
| `WebhookDispatches().List(ctx, ListOptions)` / `WebhookDispatches().Iterate(ListOptions, chunkSize *int64)` | List/iterate dispatches. |
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
