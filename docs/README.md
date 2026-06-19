# Apify Go client documentation

This directory documents the public API of the Apify Go client, organized by resource. Each
page lists the available methods with their parameters and short, runnable snippets. For an
overview, configuration, error handling, and the full resource table, see the
[top-level README](../README.md).

All snippets assume a configured client and a context:

```go
client := apify.NewClient("my-api-token")
ctx := context.Background()
```

`NewClient` takes the token as an explicit argument — it does **not** read `APIFY_TOKEN` (or
any other environment variable) automatically. Read it yourself if you want that, e.g.
`apify.NewClient(os.Getenv("APIFY_TOKEN"))`. Use `apify.NewClientWithOptions(...)` for
non-default settings (base URL, retries, timeout, user-agent suffix, custom HTTP backend).

Methods that fetch a single resource return a `(value, ok, error)` triple: a missing
resource is reported by `ok == false` rather than an error. API failures are returned as
`*apify.APIError` (see [error handling](../README.md#error-handling)).

## Setting optional fields — `apify.Ptr`

Optional option-struct fields are pointer-typed (so "unset" is distinguishable from a zero
value). The exported generic helper `apify.Ptr[T](v T) *T` sets them inline without a named
local:

```go
page, err := client.Actors().List(ctx, apify.ActorListOptions{
    My:    apify.Ptr(true),
    Limit: apify.Ptr(int64(10)),
})
```

## Pagination

List/iterate methods return `apify.PaginationList[T]`, one page plus the API's pagination
metadata:

| Field | Type | Meaning |
|---|---|---|
| `Items` | `[]T` | The items in this page. |
| `Total` | `int64` | Total items available across all pages. |
| `Count` | `int64` | Number of items actually returned in this page. |
| `Offset` | `int64` | Items skipped before this page. |
| `Limit` | `int64` | Max items the API would return for this request. |
| `Desc` | `bool` | Whether items are in descending order. |

> **`Total` can lag right after a write.** The API computes the count asynchronously, so
> immediately after a `PushItems` (or other write) `Total` may not yet include the new items.
> Re-read after a short delay if you need an exact post-write total.

## Pages

- [Actors](actors.md) — Actors, versions, environment variables.
- [Builds](builds.md) — Actor builds.
- [Runs](runs.md) — Actor runs and their default storages.
- [Tasks](tasks.md) — Actor tasks.
- [Storages](storages.md) — Datasets, key-value stores, request queues.
- [Schedules](schedules.md) — Schedules.
- [Webhooks](webhooks.md) — Webhooks and webhook dispatches.
- [Misc](misc.md) — Apify Store, users, logs.

## Examples

Full runnable programs are in [`../examples`](../examples) and are tested in CI.
