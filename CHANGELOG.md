# Changelog

All notable changes to the Apify Go client are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-06-18

Initial release of the official Go client for the Apify API, verified against OpenAPI
specification version `v2-2026-06-16T064758Z`.

### Added

- Resource-oriented `ApifyClient` mirroring the official JavaScript and Rust clients, with
  accessors for Actors, Actor versions and environment variables, builds, runs, datasets,
  key-value stores, request queues, tasks, schedules, webhooks, webhook dispatches, the
  Apify Store, users, and logs.
- Replaceable HTTP transport via the `HTTPBackend` interface (default `DefaultHTTPBackend`),
  configurable through `WithHTTPBackend`.
- Cross-cutting request behaviour applied to every call: bearer-token authentication, the
  mandated `User-Agent` header, exponential-backoff-with-jitter retries (429 and 5xx and
  network errors), and a growing-but-capped per-attempt timeout.
- Convenience helpers matching the reference clients: `Actor.Call`/`Task.Call` (start and
  wait), `Build.WaitForFinish`/`Run.WaitForFinish`, `Actor.DefaultBuild`,
  `Actor.ValidateInput`, `Run.Metamorph`/`Reboot`/`Resurrect`/`Charge`,
  run-nested default storages (`Run.Dataset`/`KeyValueStore`/`RequestQueue`/`Log`),
  lazy `Store.Iterate` and `RequestQueue.PaginateRequests` iterators,
  dataset `DownloadItems`/`GetStatistics`/`CreateItemsPublicURL`,
  key-value-store `GetRecords` (ZIP), record public URLs and key-list public URLs with
  HMAC-SHA256 signing, request-queue lock lifecycle, and `ApifyClient.SetStatusMessage`.
- Public version constants `CLIENT_VERSION` and `API_SPEC_VERSION`.
- Forward-compatible models that capture unmodelled API fields in an `Extra` map.
- Offline unit tests (mock HTTP backend) covering retries, error parsing, 404→absent
  mapping, the User-Agent format, base-URL resolution, and the storage-signature scheme.
- Integration test suite (one simple GET plus one CRUD flow per resource) and runnable,
  CI-tested documentation examples.
- A language-specific GitHub Actions workflow that runs gofmt, go vet, build, unit tests and
  integration tests, triggered by PRs to master touching Go code or manual dispatch.

### Fixed

- `RunClient.Charge` now always sends an `idempotency-key` header (auto-generated when not
  supplied), so a transport-retried charge is applied at most once — matching the reference
  client and preventing double-charging.
- `KeyValueStoreClient.GetRecord` now defaults `attachment` to the truthy form, matching the
  reference client's record-fetch behaviour. (On the wire this client serialises booleans as
  `1`/`0`, so the request carries `attachment=1`; functionally equivalent to the reference
  client's `attachment=true`.)
- `WaitForFinish` (used by `Run.WaitForFinish`, `Build.WaitForFinish`, `Actor.Call`,
  `Task.Call`) no longer hangs forever on a transient `404` during an indefinite wait
  (`waitSecs == nil`). A just-started run/build can briefly return `404` because of
  database-replica lag; the wait now polls through `404`s on a pure time bound (defaulting an
  indefinite wait to a finite upper bound, mirroring the reference client's
  `MAX_WAIT_FOR_FINISH`) and, if the resource never becomes available within the budget,
  returns a descriptive error instead of spinning until the context is cancelled.

### Changed

- `RunClient.Charge` now takes a `RunChargeOptions{EventName, Count, IdempotencyKey}` struct.
- `RunClient.Metamorph` now takes a `MetamorphOptions{Build, ContentType}` struct.
- `RunResurrectOptions` gained `MaxItems`, `MaxTotalChargeUsd`, `RestartOnError` (all in the
  spec and the reference client).
- `RunClient.GetWithWait` and `BuildClient.GetWithWait` expose the spec's `waitForFinish`
  query parameter for a server-side synchronous fetch.
- `RunCollectionClient.List` `Status` filter accepts multiple statuses (`[]string`,
  comma-separated), as the API allows.
- `RequestQueueClient.BatchAddRequests` auto-chunks inputs at the API's 25-per-call limit and
  returns a typed `BatchAddResult{ProcessedRequests, UnprocessedRequests}`.
- `RequestQueueClient.ListRequests` validates its options (mutually-exclusive
  `ExclusiveStartID`/`Cursor`; `Filter` restricted to `locked`/`pending`).
- Tasks use a dedicated `TaskStartOptions` that omits the Actor-only `contentType` and
  `forcePermissionLevel` fields (which the task run endpoint does not accept).
- `LogClient` exposes the spec's `raw`/`download` options via `GetWithOptions`/
  `StreamWithOptions`, and `RunClient.GetStreamedLog` provides a raw live-stream convenience.
- `RunClient.Abort` now takes `gracefully *bool` instead of `bool`. Passing `nil` omits the
  `gracefully` query parameter entirely (letting the server apply its default, immediate
  abort), matching the reference client's optional `gracefully` option; pass a pointer to
  `true`/`false` to request a graceful/immediate abort explicitly.

### Notes

- Intentionally unimplemented endpoints (a deliberate, parity-driven decision matching the
  JavaScript reference and the Rust sibling — not an accidental gap):
  - The synchronous run endpoints (`run-sync`, `run-sync-get-dataset-items`).
  - The keyed-`POST` record aliases.
  - The cryptographic tools `POST /v2/tools/encode-and-sign` and
    `POST /v2/tools/decode-and-verify`. These are server-side conveniences for the same
    HMAC-SHA256 signing this client already performs locally in `signature.go`; the reference
    clients do not expose them, so the Go client omits them for cross-client parity. They can
    be added alongside `signature.go` if a future requirement needs them.
  - `/v2/browser-info`. Not exposed by the reference clients; omitted for parity.
  Only documented endpoints are implemented; the omissions above are the conscious exceptions.
- User-Agent `isAtHome` flag: the canonical reference (JS) reads the `APIFY_IS_AT_HOME`
  environment variable, while `client_requirements.md`'s worked example uses the bare name
  `isAtHome`. These two same-priority requirements conflict, so the client honours **both**
  variable names (either being set marks the client "at home"). The flag is rendered
  lowercase (`true`/`false`) to stay byte-consistent with the JS reference; the requirements'
  `True`/`False` capitalisation is treated as a cosmetic example, with JS consistency winning.
- `Dataset`/`KeyValueStore` `GetOrCreate` take only a name: this spec version's create
  endpoints declare no request body or `schema` parameter, so a schema argument would be an
  undocumented extension. The `runs/last` endpoints accept only `status` (no `origin`) in
  this spec version. These choices favour strict OpenAPI compliance over the JS superset.
- CI runs `golangci-lint` (default high-signal analyzers — errcheck, govet, ineffassign,
  staticcheck, unused — plus `misspell`) in addition to `gofmt`, `go vet`, and `go build`,
  satisfying the coding-rule mandate that linting run in CI. Config in `.golangci.yml`.
