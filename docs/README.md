# Apify Go client documentation

> **Official, but experimental — AI-generated and AI-maintained.** This is an official Apify
> client, but it is experimental: it is generated and maintained by AI. Review the code before
> relying on it in production and report issues on the repository.

This directory documents the public API of the Apify Go client, organized by resource. Each
page lists the available methods with their parameters and short, runnable snippets. For an
overview, configuration, error handling, and the full resource table, see the
[top-level README](../README.md).

All snippets assume a configured client and a context:

```go
client := apify.NewClient(apify.WithToken("my-api-token"))
ctx := context.Background()
```

`NewClient` takes functional options; authentication is supplied via `apify.WithToken`. It
does **not** read `APIFY_TOKEN` (or any other environment variable) automatically. Read it
yourself if you want that, e.g. `apify.NewClient(apify.WithToken(os.Getenv("APIFY_TOKEN")))`.
Pass additional options for non-default settings (base URL, retries, timeout, user-agent
suffix, custom HTTP backend).

`WithToken` is optional. Omit it to create an unauthenticated client that can still call the
few endpoints that require no token, such as fetching a public Actor build by ID:

```go
publicClient := apify.NewClient()
build, ok, err := publicClient.Build(buildID).Get(ctx)
```

Account-scoped endpoints and anything that reads or writes your resources require a token.

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

## Common list options — `apify.ListOptions`

Most `List` methods (builds, runs, tasks, schedules, webhooks, Actor versions) take the shared
`apify.ListOptions`, which carries the standard pagination/ordering controls. All fields are
optional pointers; leave a field `nil` to use the API default. Use `apify.Ptr` to set them
inline.

| Field | Type | Meaning |
|---|---|---|
| `Offset` | `*int64` | Number of items to skip from the start of the list. |
| `Limit` | `*int64` | Maximum number of items to return. |
| `Desc` | `*bool` | If `true`, return items newest-first. |

```go
page, err := client.Builds().List(ctx, apify.ListOptions{
	Limit: apify.Ptr(int64(50)),
	Desc:  apify.Ptr(true),
})
```

Collections with extra filters use a dedicated options type instead of (or in addition to)
`ListOptions`: `ActorListOptions` (Actors), `StorageListOptions` (datasets/key-value
stores/request queues), `StoreListOptions` (the Store), and `RunListOptions` (runs, passed
alongside `ListOptions`). Each is documented on its resource page.

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

The *collection* `List` methods (Actors, builds, runs, tasks, schedules, webhooks, datasets,
key-value stores, request queues) return `PaginationList[T]`. The *within-storage* listers use
their own page/head containers instead, because the underlying API endpoints paginate
differently: `KeyValueStoreClient.ListKeys` returns `KeyValueStoreKeysPage` (key-based
pagination) and `RequestQueueClient.ListHead` returns `RequestQueueHead`. Both are documented on
the [storages](storages.md) page.

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
