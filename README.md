# Apify API client for Go

> **Official, but experimental — AI-generated and AI-maintained.** This is an official Apify
> client, but it is experimental: it is generated and maintained by AI. Review the code before
> relying on it in production and report issues on the repository.

An idiomatic Go client for the [Apify API](https://docs.apify.com/api/v2).

It provides a resource-oriented interface that mirrors the official
[JavaScript](https://github.com/apify/apify-client-js) and Rust clients: start from an
`ApifyClient`, then drill down into resources (Actors, runs, datasets, key-value stores,
request queues, tasks, schedules, webhooks, the store, users and logs).

## Features

- Resource-oriented API surface consistent with the reference clients.
- Transparent authentication, `User-Agent` header, retries with exponential backoff, and
  per-request timeouts applied to every call.
- Replaceable HTTP transport (the `HTTPBackend` interface) with a default implementation.
- Convenience helpers: start-and-wait, build/run polling, lazy store and request-queue
  iterators, dataset export, signed public URLs, request-queue locking, and more.
- Forward-compatible models that keep unknown API fields in an `Extra` map.
- A single third-party dependency (`github.com/andybalholm/brotli`, used for Brotli
  request-body compression); everything else is the Go standard library.

## Installation

```sh
go get github.com/apify/apify-client-go
```

Requires Go 1.23 or newer.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	apify "github.com/apify/apify-client-go"
)

func main() {
	client := apify.NewClient(apify.WithToken(os.Getenv("APIFY_TOKEN")))
	ctx := context.Background()

	// Start an Actor and wait for it to finish.
	waitSecs := int64(120)
	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{}, &waitSecs)
	if err != nil {
		log.Fatal(err)
	}

	// Read items from the run's default dataset.
	page, err := client.Dataset(run.DefaultDatasetID).ListItems(ctx, apify.DatasetListItemsOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got %d items\n", page.Total)
}
```

Get your API token from the [Apify Console](https://console.apify.com/) under **Settings →
Integrations** (the **Personal API tokens** section). The client never reads it from the
environment itself: pass the token via `apify.WithToken` explicitly (the examples above
read `APIFY_TOKEN` from the environment only as a convenience in `main`).

## Configuration

`NewClient` takes functional options. Pass `WithToken` for authentication, plus any other
options for full control. `WithToken` is optional: omit it to create an unauthenticated
client, which can still call the few endpoints that require no token (for example, resolving a
public Actor's default build with `client.Actor("apify/hello-world").DefaultBuild(ctx, nil)` —
see [`examples/public_build_no_token`](examples/public_build_no_token)). Most endpoints —
anything account-scoped or that reads or writes your resources — require a token.

```go
client := apify.NewClient(
	apify.WithToken("my-api-token"),
	apify.WithBaseURL("https://api.apify.com"),       // /v2 is appended automatically
	apify.WithPublicBaseURL("https://api.apify.com"), // base for signed, shareable URLs
	apify.WithMaxRetries(8),                          // default 8
	apify.WithMinDelayBetweenRetries(500*time.Millisecond),
	apify.WithTimeout(360*time.Second), // default 6 minutes
	apify.WithUserAgentSuffix("MyTool/1.0"),
	apify.WithHTTPBackend(apify.NewDefaultHTTPBackend()),
)
```

| Option | Default | Description |
| --- | --- | --- |
| `WithToken` | — | API token, sent as a `Bearer` token. Optional; omit for an unauthenticated client limited to endpoints that need no token. |
| `WithBaseURL` | `https://api.apify.com` | API base URL; `/v2` is appended automatically. |
| `WithPublicBaseURL` | API base URL | Base URL used for building public, shareable URLs. |
| `WithMaxRetries` | `8` | Maximum retries for failed requests. |
| `WithMinDelayBetweenRetries` | `500ms` | Minimum delay between retries (doubled each retry). |
| `WithTimeout` | `360s` | Overall per-request timeout. |
| `WithUserAgentSuffix` | — | Suffix appended to the `User-Agent` header. |
| `WithHTTPBackend` | `DefaultHTTPBackend` | Replaceable HTTP transport. |

The `User-Agent` header reports an `isAtHome` flag indicating whether the client runs on the
Apify platform. It is driven solely by the `APIFY_IS_AT_HOME` environment variable (the same
variable the JavaScript reference client reads); if it is set to a non-empty value, the flag is
`true`, otherwise `false`.

## Resource clients

| Accessor | Returns | Purpose |
| --- | --- | --- |
| `client.Actors()` / `client.Actor(id)` | `*ActorCollectionClient` / `*ActorClient` | List/create Actors; manage a single Actor and its runs, builds, versions, webhooks. |
| `client.Builds()` / `client.Build(id)` | `*BuildCollectionClient` / `*BuildClient` | List builds; inspect, abort, wait for, and read a build's log/OpenAPI. |
| `client.Runs()` / `client.Run(id)` | `*RunCollectionClient` / `*RunClient` | List runs; manage a run, its default storages, and its log. |
| `client.Datasets()` / `client.Dataset(id)` | `*DatasetCollectionClient` / `*DatasetClient` | List/get-or-create datasets; read/write/export items. |
| `client.KeyValueStores()` / `client.KeyValueStore(id)` | `*KeyValueStoreCollectionClient` / `*KeyValueStoreClient` | List/get-or-create stores; read/write records and keys. |
| `client.RequestQueues()` / `client.RequestQueue(id)` | `*RequestQueueCollectionClient` / `*RequestQueueClient` | List/get-or-create queues; add/list/lock requests. |
| `client.Tasks()` / `client.Task(id)` | `*TaskCollectionClient` / `*TaskClient` | List/create tasks; manage a task, its input and runs. |
| `client.Schedules()` / `client.Schedule(id)` | `*ScheduleCollectionClient` / `*ScheduleClient` | List/create schedules; manage a single schedule. |
| `client.Webhooks()` / `client.Webhook(id)` | `*WebhookCollectionClient` / `*WebhookClient` | List/create webhooks; manage and test a single webhook. |
| `client.WebhookDispatches()` / `client.WebhookDispatch(id)` | `*WebhookDispatchCollectionClient` / `*WebhookDispatchClient` | List/inspect webhook dispatches. |
| `client.Store()` | `*StoreCollectionClient` | Browse and iterate the Apify Store. |
| `client.Me()` / `client.User(id)` | `*UserClient` | Account details, usage and limits (account scope for `Me()`). |
| `client.Log(buildOrRunID)` | `*LogClient` | Read or stream a build/run log. |

## Error handling

API errors are returned as `*APIError`. Recover it from any returned `error` with
`apify.AsAPIError(err) (*APIError, bool)` — the boolean is `false` when the error is not an API
error (e.g. a network or context error). `get`/`delete` on a missing resource is *not* an error:
the methods report absence via a boolean (`ok`) instead.

`*APIError` exposes:

| Field | Type | Meaning |
|---|---|---|
| `StatusCode` | `int` | HTTP status code of the error response. |
| `Type` | `string` | Machine-readable error type returned by the API (e.g. `"record-not-found"`). |
| `Message` | `string` | Human-readable error description returned by the API. |
| `Attempt` | `int` | 1-based number of the API call attempt that produced this error. |
| `HTTPMethod` | `string` | HTTP method of the failing call (e.g. `"GET"`, `"POST"`). |
| `Path` | `string` | Request path of the endpoint (URL excluding origin). |

```go
user, ok, err := client.Me().Get(ctx)
if err != nil {
	if apiErr, isAPI := apify.AsAPIError(err); isAPI {
		fmt.Printf("API error %d (%s): %s\n", apiErr.StatusCode, apiErr.Type, apiErr.Message)
	}
	log.Fatal(err)
}
if !ok {
	log.Fatal("user not found")
}
```

## Custom HTTP transport

The transport is replaceable. Implement `HTTPBackend` (a single `Do` method) to integrate a
custom client, proxy, or test double, and pass it with `WithHTTPBackend`:

```go
package main

import (
	"net/http"

	apify "github.com/apify/apify-client-go"
)

// myBackend is a custom HTTPBackend wrapping a standard *http.Client.
type myBackend struct{ inner *http.Client }

func (b *myBackend) Do(req *http.Request) (*http.Response, error) {
	return b.inner.Do(req)
}

func main() {
	client := apify.NewClient(
		apify.WithToken("my-api-token"),
		apify.WithHTTPBackend(&myBackend{inner: http.DefaultClient}),
	)
	_ = client
}
```

## Versioning

- `apify.ClientVersion` — the semantic version of this library.
- `apify.APISpecVersion` — the Apify OpenAPI spec version this client was built against
  (`v2-2026-07-13T092445Z`).

### Releasing

Go modules are distributed by pushing a semver git tag — there is no separate package registry to
upload to. The [`Publish Go client`](.github/workflows/go-publish.yml) workflow is the release
mechanism: a maintainer triggers it manually (`workflow_dispatch`), it runs the same quality gate
as CI, then creates and pushes the `v<ClientVersion>` tag, opens a GitHub release, and asks the Go
module proxy to index the new version so it appears on
[pkg.go.dev](https://pkg.go.dev/github.com/apify/apify-client-go). The tag is derived from
`ClientVersion` in [`version.go`](version.go), so bump that constant before releasing. The workflow
uses only the built-in `GITHUB_TOKEN`; no extra credentials are required.

There is no "Trusted Publisher" step: that mechanism applies to registries that authenticate
uploads (e.g. PyPI, npm, crates.io). Go has no central registry and no upload step — a module is
published purely by pushing a git tag that the public module proxy reads — so there is no token or
trusted-publisher relationship to configure. The workflow therefore relies only on the repository's
built-in `GITHUB_TOKEN` to push the tag and open the release.

## Examples

Runnable examples live in [`examples/`](examples) and are exercised in CI. Most need a token;
run them like:

```sh
APIFY_TOKEN=<your-token> go run ./examples/run_store_actor
```

`public_build_no_token` needs no token — run it with `go run ./examples/public_build_no_token`.

| Example | Description |
| --- | --- |
| `get_account` | Fetch and print the current account details. |
| `storages` | Create, write to, and read from a dataset, key-value store, and request queue. |
| `run_store_actor` | Run a Store Actor, wait for it, and read its default dataset. |
| `run_and_last_run_storages` | Start a run, then fetch the Actor's last run and its storages. |
| `iterate_store` | Lazily iterate Actors in the Apify Store. |
| `log_redirection` | Start an Actor and stream its log in real time. |
| `create_build_run_actor` | Create an Actor, build it, run it, and print the run log. |
| `public_build_no_token` | Fetch a public Actor's default build with an unauthenticated client (no token). |

## Documentation

Per-resource guides are in [`docs/`](docs):
[actors](docs/actors.md), [builds](docs/builds.md), [runs](docs/runs.md),
[tasks](docs/tasks.md), [storages](docs/storages.md), [schedules](docs/schedules.md),
[webhooks](docs/webhooks.md), [misc (store, users, logs)](docs/misc.md).

## Scope

The client implements only documented API endpoints. Matching the JavaScript reference (and
the Rust sibling) for cross-client consistency, the following documented endpoints are
**intentionally not implemented**:

- Synchronous run endpoints (`run-sync`, `run-sync-get-dataset-items`).
- The keyed-`POST` record aliases.
- Cryptographic tools: `POST /v2/tools/encode-and-sign` and `POST /v2/tools/decode-and-verify`.
  These are server-side conveniences for the same HMAC signing this client already performs
  locally in `signature.go`; the reference clients do not expose them, so the Go client omits
  them too for parity. (Should a future requirement need them, they can be added alongside
  `signature.go`.)
- `/v2/browser-info`. Not exposed by the reference clients; omitted for parity.

This is a deliberate, parity-driven decision, not an accidental gap. See the CHANGELOG for the
same note.

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE).
