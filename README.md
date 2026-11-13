# Apify API client for Go

> **Official, but experimental — AI-generated and AI-maintained.** This is an official Apify
> client, but it is experimental: it is generated and maintained by AI. Review the code before
> relying on it in production and report issues on the repository.

An idiomatic Go client for the [Apify API](https://docs.apify.com/api/v2).

Start from an `ApifyClient` and go on to the resources: Actors, runs, datasets, key-value
stores, request queues, tasks, schedules, webhooks, the Store, users, and logs.

## Features

- One client per resource, with the same method names as the [JavaScript client](https://github.com/apify/apify-client-js).
- Transparent authentication, `User-Agent` header, retries with exponential backoff, and
  per-request timeouts applied to every call.
- Replaceable HTTP transport (the `HTTPBackend` interface) with a default implementation.
- Convenience helpers: start-and-wait with live log redirection, build/run polling, lazy store
  and request-queue iterators, live iteration of a running Actor's dataset items, dataset
  export, signed public URLs, request-queue locking, requests to Server Actors, and more.
- Forward-compatible models that keep unknown API fields in an `Extra` map.
- Two third-party dependencies for request-body compression, `github.com/andybalholm/brotli` and
  `github.com/klauspost/compress` (zstd). Everything else is the Go standard library.

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

Create an API token on the [Integrations](https://console.apify.com/settings/integrations) page
of Apify Console (the **Personal API tokens** section). The Console shows the token only once,
right after you create or rotate it, so copy it and store it somewhere safe. The client never reads it from the
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
package main

import (
	"time"

	apify "github.com/apify/apify-client-go"
)

func main() {
	client := apify.NewClient(
		apify.WithToken("my-api-token"),
		apify.WithBaseURL("https://api.apify.com"),       // /v2 is appended automatically
		apify.WithPublicBaseURL("https://api.apify.com"), // base for signed, shareable URLs
		apify.WithMaxRetries(8),                          // default 8
		apify.WithMinDelayBetweenRetries(500*time.Millisecond),
		apify.WithTimeoutShort(5*time.Second),   // metadata reads/writes; default 5s
		apify.WithTimeoutMedium(30*time.Second), // listing/batch/trigger calls; default 30s
		apify.WithTimeoutLong(360*time.Second),  // downloads/uploads/streaming; default 360s
		apify.WithTimeoutMax(360*time.Second),   // caps any single request attempt; default 360s
		apify.WithUserAgentSuffix("MyTool/1.0"),
		apify.WithHTTPBackend(apify.NewDefaultHTTPBackend()),
	)
	_ = client
}
```

| Option | Default | Description |
| --- | --- | --- |
| `WithToken` | — | API token, sent as a `Bearer` token. Optional; omit for an unauthenticated client limited to endpoints that need no token. |
| `WithBaseURL` | `https://api.apify.com` | API base URL, with or without the `/v2` version path — it is appended automatically unless already present. |
| `WithPublicBaseURL` | API base URL | Base URL used for building public, shareable URLs. Same `/v2` handling as `WithBaseURL`. |
| `WithMaxRetries` | `8` | Maximum retries for failed requests. |
| `WithMinDelayBetweenRetries` | `500ms` | Minimum delay between retries (doubled each retry). |
| `WithTimeoutShort` | `5s` | Duration of the "short" timeout tier — simple metadata reads/writes (`Get`, `Update`, `Delete`). See [Timeouts](#timeouts). |
| `WithTimeoutMedium` | `30s` | Duration of the "medium" tier — listing, batch and trigger operations (`List`, `Create`, `Start`, `BatchAddRequests`). |
| `WithTimeoutLong` | `360s` | Duration of the "long" tier — downloads, uploads and streaming (`ListItems`, `SetRecord`, `Log().Get()`). |
| `WithTimeoutMax` | `360s` | Caps the timeout of any single request attempt; a tier configured above it is capped too. `WithTimeout` is an alias for this option, kept for callers of the single-timeout API this client had before tiers. |
| `WithRequestCompression` | `RequestCompressionAuto` | Codec for request-body compression: `RequestCompressionAuto` (brotli, falling back to gzip), `RequestCompressionBrotli`, `RequestCompressionGzip`, or `RequestCompressionZstd`. See [Request compression](#request-compression). |
| `WithUserAgentSuffix` | — | Suffix appended to the `User-Agent` header. |
| `WithHTTPBackend` | `DefaultHTTPBackend` | Replaceable HTTP transport. |

### Timeouts

Every request the client sends is assigned one of three timeout tiers, matching the kind of
operation its method performs: `short` for a metadata read/write, `medium` for a listing, batch or
trigger call, `long` for a download, upload or streaming call. The duration of each tier is
configured on the client (see the table above); a timeout still grows with each retry, capped at
`WithTimeoutMax`.

Two kinds of request have no client-side timeout of their own: the polling requests behind
`WaitForFinish` and `Call`, and the start request of `Call`. `waitSecs` bounds how long `Call` and
`WaitForFinish` wait for the run, and the `context.Context` bounds the whole call.

Every method takes a `context.Context`, and its deadline or cancellation can only shorten a
request: the request ends when its tier's timeout or the context's deadline passes, whichever is
sooner. To give one call a shorter bound than its tier, pass a context created with
`context.WithTimeout`; a context cannot extend a tier, so raise the tier (`WithTimeoutShort`,
`WithTimeoutMedium`, `WithTimeoutLong`, `WithTimeoutMax`) for a call that needs longer.

A request that timed out is retried like any other transient failure, with a longer timeout,
except one that is not idempotent, where the timed-out attempt may already have taken effect:
starting a run (`Start`, `StartRaw`, a task's `Start`), building an Actor (`Build`), pushing
dataset items (`PushItems`) and, when `SetRecordOptions.DoNotRetryTimeouts` is set, storing a
record. Those fail on the first timeout. The other failures (a 429 or 5xx response, a network
error) are retried for them as well.

A method with a `waitForFinishSecs`/`WaitForFinish` parameter that asks the API to hold the
response (e.g. `RunClient.GetWithWait`, `ActorClient.Start`'s `WaitForFinish` option) extends
its tier's timeout by the requested wait (capped at 60s, the server's own hold limit), so the
request is not aborted client-side while the server is correctly honoring it.

### Request compression

Request bodies of 1 KiB or more are compressed before sending, to save upload bandwidth. By
default (`RequestCompressionAuto`) the client prefers brotli and falls back to gzip when brotli
errors or does not shrink the body. `WithRequestCompression` with `RequestCompressionBrotli`,
`RequestCompressionGzip`, or `RequestCompressionZstd` pins it to one codec, with no fallback: a
body that codec does not shrink is sent uncompressed. A body whose content type is already
compressed, such as images, audio, video, and archives, is not compressed, except for raw formats
like `image/bmp`.

The `User-Agent` header reports an `isAtHome` flag indicating whether the client runs on the
Apify platform. It is driven solely by the `APIFY_IS_AT_HOME` environment variable: if it is set
to a non-empty value, the flag is `true`, otherwise `false`.

### URL fields

The fields the API specification declares as URLs (`consoleUrl`, `containerUrl`, `itemsPublicUrl`,
`keysPublicUrl`, `recordsPublicUrl`, `recordPublicUrl`, `standbyUrl`, the webhook `requestUrl`, a
run's `output` links, a Store Actor's `url` and `userPictureUrl`, and a user profile's
`pictureUrl` and `websiteUrl`) are returned in a normalized form, close to the WHATWG URL
serialization: the scheme and host are lowercased, a non-ASCII host becomes punycode, an IPv4 or
IPv6 host is written canonically, a default port is dropped, an empty path becomes `/`, dot
segments are resolved, and unsafe characters are percent-encoded. It covers the http(s), ws(s),
ftp and file URLs the API returns; it does not apply the full Unicode (UTS #46) mapping to a
non-ASCII host. Compare a field against a stored URL after normalizing the stored one the same
way (for example by reading it back from the API). A value that is not an absolute URL makes
decoding the response fail.

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
| `client.Me()` / `client.User(id)` | `*UserClient` | Account details, usage and limits (for your own account or your organization with `User(id)`). |
| `client.Log(buildOrRunID)` | `*LogClient` | Read or stream a build/run log. |
| `client.Actor(id).Server()` | `*ActorServerClient` | Send requests to a Server Actor (formerly Actor Standby) and get its URL and OpenAPI specification. See [docs/server_actors.md](docs/server_actors.md). |

## Error handling

API errors are returned as `*APIError`. Recover it from any returned `error` with
`apify.AsAPIError(err) (*APIError, bool)` — the boolean is `false` when the error is not an API
error (e.g. a network or context error).

`get`/`delete` addressing a resource by its own ID is *not* an error on a 404: the methods
report absence via a boolean (`ok`) instead. A client reached through a fixed sub-path with no
ID of its own, such as `client.Run(id).Dataset()`, cannot tell its own 404 apart from the
parent's — there the error always propagates instead. See [docs/runs.md](docs/runs.md) for the
affected accessors.

`*APIError` exposes:

| Field | Type | Meaning |
|---|---|---|
| `StatusCode` | `int` | HTTP status code of the error response. |
| `Type` | `string` | Machine-readable error type returned by the API (e.g. `"record-not-found"`). |
| `Message` | `string` | Human-readable error description returned by the API. |
| `Attempt` | `int` | 1-based number of the API call attempt that produced this error. |
| `HTTPMethod` | `string` | HTTP method of the failing call (e.g. `"GET"`, `"POST"`). |
| `Path` | `string` | Request path of the endpoint (URL excluding origin). |

It also exposes `IsInvalidRequest()`, `IsUnauthorized()`, `IsForbidden()`, `IsNotFound()`,
`IsConflict()`, `IsRateLimited()` and `IsServerError()`, which classify the status.

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

### Cancellation

Every method takes a `context.Context`. Cancel it, or let its timeout or deadline pass, to stop
the request.

## Custom HTTP transport

The transport is replaceable. Implement `HTTPBackend` (a single `Do` method) to integrate a
custom client, proxy, or test double, and pass it with `WithHTTPBackend`. A backend that follows
redirects must not follow them for a request that `apify.FollowsRedirects(req)` reports false for,
as the requests to a [Server Actor](docs/server_actors.md) carry a token meant for one host only,
and a redirect to another host would send it there:

```go
package main

import (
	"net/http"

	apify "github.com/apify/apify-client-go"
)

// myBackend is a custom HTTPBackend wrapping a standard *http.Client.
type myBackend struct{ inner *http.Client }

func (b *myBackend) Do(req *http.Request) (*http.Response, error) {
	if apify.FollowsRedirects(req) {
		return b.inner.Do(req)
	}
	// This request carries a token meant for one host only, so a redirect is returned as it is.
	client := *b.inner
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	return client.Do(req)
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
  (`v2-2026-11-10T140017Z`).

### Releasing

A Go module is published by pushing a semver git tag. A maintainer triggers the
[`Publish Go client`](.github/workflows/go-publish.yml) workflow manually. It runs the checks of CI,
pushes the `v<ClientVersion>` tag, opens a GitHub release and asks the Go module proxy to index
the version for [pkg.go.dev](https://pkg.go.dev/github.com/apify/apify-client-go). Bump
`ClientVersion` in [`version.go`](version.go) before releasing.

Go has no registry to upload to, so there is no Trusted Publisher step, and the workflow needs only
the built-in `GITHUB_TOKEN`.

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
| `run_and_last_run_storages` | Start a run, wait for it to finish, then fetch the Actor's last run and its storages. |
| `iterate_store` | Lazily iterate Actors in the Apify Store. |
| `log_redirection` | Run an Actor and redirect its log to stdout in real time. |
| `create_build_run_actor` | Create an Actor, build it, run it, and print the run log. |
| `public_build_no_token` | Fetch a public Actor's default build with an unauthenticated client (no token). |

## Documentation

Per-resource guides are in [`docs/`](docs):
[actors](docs/actors.md), [server Actors](docs/server_actors.md), [builds](docs/builds.md), [runs](docs/runs.md),
[tasks](docs/tasks.md), [storages](docs/storages.md), [schedules](docs/schedules.md),
[webhooks](docs/webhooks.md), [misc (store, users, logs)](docs/misc.md).

## Scope

The client implements only documented API endpoints. These documented endpoints are not
implemented:

- Synchronous run endpoints (`run-sync`, `run-sync-get-dataset-items`).
- The keyed-`POST` record aliases.
- `POST /v2/tools/encode-and-sign` and `POST /v2/tools/decode-and-verify`. The client signs
  public URLs locally.
- `/v2/browser-info`.

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE).
