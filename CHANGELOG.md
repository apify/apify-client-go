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
- `KeyValueStoreClient.GetRecord` now defaults `attachment=true`, matching the reference
  client's record-fetch behaviour.

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

### Notes

- Intentionally unimplemented endpoints (matching the JavaScript reference scope): the
  synchronous run endpoints (`run-sync`, `run-sync-get-dataset-items`), and the keyed-POST
  record aliases. Only documented endpoints are implemented.
- `Dataset`/`KeyValueStore` `GetOrCreate` take only a name: this spec version's create
  endpoints declare no request body or `schema` parameter, so a schema argument would be an
  undocumented extension. The `runs/last` endpoints accept only `status` (no `origin`) in
  this spec version. These choices favour strict OpenAPI compliance over the JS superset.
