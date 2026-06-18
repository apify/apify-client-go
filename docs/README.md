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

Methods that fetch a single resource return a `(value, ok, error)` triple: a missing
resource is reported by `ok == false` rather than an error. API failures are returned as
`*apify.APIError` (see [error handling](../README.md#error-handling)).

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
