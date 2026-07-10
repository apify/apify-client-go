# Changelog

All notable changes to the Apify Go client are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.6.0] - 2026-07-10

### Added

- Lazy `Iterate` helpers on every list collection (`Actors`, `Runs`, `Builds`, `Tasks`,
  `Datasets`, `KeyValueStores`, `RequestQueues`, `Schedules`, `Webhooks`, `WebhookDispatches`,
  actor versions and env vars) plus dataset-item iteration (`DatasetClient.IterateItems` and the
  generic `IterateDatasetItems[T]`), backed by a new exported generic iterator type
  `ListIterator[T]`. As in the reference client's iterable `list()`, the options' `Limit` caps the
  total number of items yielded across all pages (unset means all), and the per-page size is a
  separate `chunkSize` argument (nil for the server default).

### Changed

- Bumped `APISpecVersion` to `v2-2026-07-10T105921Z`.
- Bumped `ClientVersion` to `0.6.0`.
- **Breaking:** `StoreCollectionClient.Iterate` now takes a second `chunkSize *int64` argument and
  treats the options' `Limit` as a total-item cap rather than the per-page size, to match the
  reference client's iterator semantics. `StoreActorIterator` is now an alias of
  `ListIterator[ActorStoreListItem]`.
- Synced the `APISpecVersion` reference in the `README.md` "Versioning" section to match `version.go`.

## [0.5.0] - 2026-07-09

### Changed

- **Breaking:** Renamed the exported version constants `CLIENT_VERSION` -> `ClientVersion` and
  `API_SPEC_VERSION` -> `APISpecVersion` to follow idiomatic Go naming (MixedCaps, not ALL_CAPS).
- Bumped `APISpecVersion` to `v2-2026-07-08T143931Z`.
- Bumped `ClientVersion` to `0.5.0`.

- Mapped the `User-Agent` OS token to match Node's `os.platform()` exactly: `windows` -> `win32`,
  `solaris`/`illumos` -> `sunos`, `ios` -> `darwin`; all other platforms are unchanged.

### Added

- Request bodies of 1024 bytes or larger are now compressed before being sent, preferring Brotli
  (`Content-Encoding: br`) and falling back to gzip (`Content-Encoding: gzip`).

## [0.4.7] - 2026-07-07

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-07-07T132551Z`.
- Bumped `CLIENT_VERSION` to `0.4.7`.
- Synced the `API_SPEC_VERSION` reference in the `README.md` "Versioning" section to match `version.go`.

### Documentation

- Corrected the `LastRunOptions.Origin` doc comment: `origin` is now a documented query parameter of the
  `runs/last` endpoints in the OpenAPI spec (behavior unchanged; the client already sent it).

## [0.4.6] - 2026-07-07

### Changed

- Rewrote earlier `CHANGELOG.md` entries to satisfy the changelog requirements: condensed
  narrative prose into short change bullets and removed cross-client references to sibling
  implementations, references to requirement-tracking issues, and out-of-scope / not-implemented
  notes.
- Bumped `CLIENT_VERSION` to `0.4.6`.

## [0.4.5] - 2026-07-03

### Changed
- Bumped `API_SPEC_VERSION` to `v2-2026-07-02T131926Z`.
- Bumped `CLIENT_VERSION` to `0.4.5`.
- Synced the `API_SPEC_VERSION` reference in the `README.md` "Versioning" section to match `version.go`.

## [0.4.4] - 2026-07-01

### Changed
- Bumped `API_SPEC_VERSION` to `v2-2026-07-01T115402Z`.
- Bumped `CLIENT_VERSION` to `0.4.4`.
- Synced the `API_SPEC_VERSION` reference in the `README.md` "Versioning" section to match `version.go`.

## [0.4.3] - 2026-06-30

### Changed
- Bumped `API_SPEC_VERSION` to `v2-2026-06-30T091455Z`.
- Bumped `CLIENT_VERSION` to `0.4.3`.
- Synced the `API_SPEC_VERSION` reference in the `README.md` "Versioning" section to match `version.go`.

### Documentation
- `docs/actors.md`: document the nil-`waitSecs` behaviour of `Call`.

## [0.4.2] - 2026-06-29

### Changed
- Bumped `API_SPEC_VERSION` to `v2-2026-06-29T142258Z` (no in-scope API surface change).
- Bumped `CLIENT_VERSION` to `0.4.2`.
- Synced the `API_SPEC_VERSION` reference in the `README.md` "Versioning" section to match `version.go`.

## [0.4.1] - 2026-06-29

### Changed
- Documentation now states the client is **official, but experimental** and AI-generated/AI-maintained (`README.md`, `docs/README.md`, and the package-level doc comment in `client.go`).
- `README.md` "Releasing" section now documents that Go's publishing process has no central registry and therefore no "Trusted Publisher" mechanism — releases are cut by pushing a git tag read by the public Go module proxy, authenticated only by the built-in `GITHUB_TOKEN`.
- Documentation completeness pass (no API change): added a field table for `*APIError` and documented `AsAPIError` in `README.md`; added a `RunChargeOptions` field table in `docs/runs.md`; enumerated the webhook `Create` input keys in `docs/webhooks.md`; documented the request-queue locking method signatures in `docs/storages.md`.

### Fixed
- Corrected a stale `API_SPEC_VERSION` reference in the `README.md` "Versioning" section (`v2-2026-06-24T105326Z` → `v2-2026-06-25T142310Z`).
- `docs/storages.md`: `DownloadItems` now lists `JSONL` in its format list.
- Bumped `CLIENT_VERSION` to `0.4.1`.

## [0.4.0] - 2026-06-26

### Added

- `ActorClient.LastRunWithOptions` and `TaskClient.LastRunWithOptions` accept a `LastRunOptions` with `Status` and `Origin` filters, matching the reference client's `lastRun({ status, origin })`. The existing `LastRun(status string)` accessors delegate to the new methods (additive, non-breaking).

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-06-25T142310Z`.
- Bumped `CLIENT_VERSION` to `0.4.0` (minor; additive `LastRunWithOptions` API).

### Fixed

- Cleaned up stale in-code comments around the `isAtHome` User-Agent flag that quoted the older, capitalized requirement wording (no behaviour change).

## [0.3.0] - 2026-06-25

### Fixed

- `ListRequestsOptions.Filter` (for `GET /v2/request-queues/{queueId}/requests`) is now `[]string` serialized comma-joined, matching the spec (an array of the enum values `locked`/`pending`, `style=form` `explode=false`). Breaking change to that field's type. The allowed values are exported as the `RequestFilterLocked` / `RequestFilterPending` constants.
- The `User-Agent` `isAtHome` flag is now based solely on the `APIFY_IS_AT_HOME` environment variable.

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-06-24T105326Z`.
- Bumped `CLIENT_VERSION` to `0.3.0` (minor; breaking `Filter` type change).

## [0.2.3] - 2026-06-23

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-06-23T113219Z`.
- Bumped `CLIENT_VERSION` to `0.2.3`.

## [0.2.2] - 2026-06-22

### Added

- CI: a manually triggered (`workflow_dispatch`) `Publish Go client` workflow (`.github/workflows/go-publish.yml`) that releases the module. Go is distributed by pushing a semver git tag, so the workflow runs the same gofmt/vet/lint/build/unit-test quality gate as CI, derives the tag from the `CLIENT_VERSION` constant in `version.go`, refuses to proceed if that tag already exists, then creates and pushes the `vX.Y.Z` tag, opens a GitHub release, and pings the public Go module proxy. It uses only the built-in `GITHUB_TOKEN` repository secret and supports a `dry_run` input.

### Changed

- Bumped `CLIENT_VERSION` to `0.2.2`.
- Documentation: added a "Releasing" subsection to the README "Versioning" section describing the tag-based Go distribution mechanism and the publish workflow.
- Documentation: added the "experimental, AI-generated and AI-maintained" disclaimer to `README.md` and `docs/README.md`, and softened the "official" wording accordingly.

## [0.2.1] - 2026-06-19

### Added

- CI: a standalone `Test examples` step in the Go integration workflow that runs the documentation example code end-to-end (the `TestExample*` smoke tests, each running `go run ./examples/<name>`) and validates that every in-documentation `go` snippet is valid, runnable, and gofmt-formatted (the new `TestDocSnippets*` tests). The `Integration tests` step now skips these so the two concerns stay separate.
- Tests: `tests/docs_snippets_test.go`, an offline doc-snippet harness that enforces that each in-documentation code snippet is valid, runnable, and properly formatted.

### Changed

- Workflow now also triggers on `docs/**` and `README.md` changes.

### Fixed

- Reformatted all `docs/` and README code snippets to canonical gofmt output and made the custom-HTTP-transport snippet a complete, compilable program.
- Corrected the README versioning note (`v2-2026-06-16T064758Z` → `v2-2026-06-18T095846Z`).

### Documentation

- Documented the shared `ListOptions` type and the `apify.ListDatasetItems[T]` generic helper (with a runnable typed-decoding example) in `docs/README.md` / `docs/storages.md`.
- Added field listings for the response models (`ActorRun`, `Build`, `User`, `ActorStoreListItem`, `Actor`, `Task`, `Schedule`, `Webhook`, `WebhookDispatch`) across `docs/runs.md`, `docs/builds.md`, `docs/misc.md`, and their resource pages.
- Added field tables in `docs/storages.md` for the storage option/parameter structs (`DatasetListItemsOptions`, `DatasetDownloadOptions`, `ListKeysOptions`, `GetRecordOptions`, `GetRecordsOptions`, `ListRequestsOptions`, `RequestQueueRequest`) and the storage return types (`KeyValueStoreRecord`, `KeyValueStoreKeysPage`, `RequestQueueHead`, `RequestQueueOperationInfo`, `BatchAddResult`).
- Expanded the run/Actor/store input option structs (`ActorStartOptions`, `ActorBuildOptions`, `ActorListOptions`, `StoreListOptions`, `RunResurrectOptions`, `MetamorphOptions`, `LogOptions`) into full field/type/meaning tables, and stated the closed enum sets definitively (`RunListOptions.Status`, `ListRequestsOptions.Filter`, `DatasetListItemsOptions.View`, `StorageListOptions`, `StoreListOptions.PricingModel`/`ResponseFormat`, `ActorStartOptions.ForcePermissionLevel`, `ActorListOptions.SortBy`, the 8-value `ActorJobStatus`, the version `sourceType`/`format` enums, and the 12-value `WebhookEventType`).
- Documented the schedule `actions` payload shape with a runnable `RUN_ACTOR` action example in `docs/schedules.md`.

## [0.2.0] - 2026-06-19

### Added

- `ActorClient.ValidateInputForBuild(ctx, input, build)` exposes the optional `build` query parameter on `POST /v2/actors/{actorId}/validate-input`. `ValidateInput` delegates with an empty build.
- `UserClient.MonthlyUsageForDate(ctx, date)` exposes the optional `date` query parameter on `GET /v2/users/me/usage/monthly`. `MonthlyUsage` delegates with an empty date.

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-06-18T095846Z`.

## [0.1.0] - 2026-06-18

Initial release of the official Go client for the Apify API, verified against OpenAPI
specification version `v2-2026-06-16T064758Z`.

### Added

- Resource-oriented `ApifyClient` mirroring the JavaScript reference client, with accessors for Actors, Actor versions and environment variables, builds, runs, datasets, key-value stores, request queues, tasks, schedules, webhooks, webhook dispatches, the Apify Store, users, and logs.
- Replaceable HTTP transport via the `HTTPBackend` interface (default `DefaultHTTPBackend`), configurable through `WithHTTPBackend`.
- Cross-cutting request behaviour applied to every call: bearer-token authentication, the mandated `User-Agent` header, exponential-backoff-with-jitter retries (429, 5xx and network errors), and a growing-but-capped per-attempt timeout.
- Convenience helpers matching the reference client: `Actor.Call`/`Task.Call`, `Build.WaitForFinish`/`Run.WaitForFinish`, `Actor.DefaultBuild`, `Actor.ValidateInput`, `Run.Metamorph`/`Reboot`/`Resurrect`/`Charge`, run-nested default storages (`Run.Dataset`/`KeyValueStore`/`RequestQueue`/`Log`), lazy `Store.Iterate` and `RequestQueue.PaginateRequests` iterators, dataset `DownloadItems`/`GetStatistics`/`CreateItemsPublicURL`, key-value-store `GetRecords` (ZIP), record and key-list public URLs with HMAC-SHA256 signing, the request-queue lock lifecycle, and `ApifyClient.SetStatusMessage`.
- Public version constants `CLIENT_VERSION` and `API_SPEC_VERSION`.
- Forward-compatible models that capture unmodelled API fields in an `Extra` map.
- Offline unit tests (mock HTTP backend) covering retries, error parsing, 404→absent mapping, the User-Agent format, base-URL resolution, and the storage-signature scheme.
- Integration test suite (one simple GET plus one CRUD flow per resource) and runnable, CI-tested documentation examples.
- A language-specific GitHub Actions workflow that runs gofmt, go vet, build, unit tests and integration tests, triggered by PRs to master touching Go code or manual dispatch. CI also runs `golangci-lint` (errcheck, govet, ineffassign, staticcheck, unused, misspell), configured in `.golangci.yml`.

### Fixed

- `RunClient.Charge` now always sends an `idempotency-key` header (auto-generated when not supplied), so a transport-retried charge is applied at most once — matching the reference client and preventing double-charging.
- `KeyValueStoreClient.GetRecord` now defaults `attachment` to the truthy form, matching the reference client's record-fetch behaviour (serialised on the wire as `attachment=1`).
- `WaitForFinish` (used by `Run.WaitForFinish`, `Build.WaitForFinish`, `Actor.Call`, `Task.Call`) no longer hangs forever on a transient `404` during an indefinite wait (`waitSecs == nil`); it polls through `404`s on a pure time bound and returns a descriptive error if the resource never becomes available within the budget.

### Changed

- `RunClient.Charge` now takes a `RunChargeOptions{EventName, Count, IdempotencyKey}` struct.
- `RunClient.Metamorph` now takes a `MetamorphOptions{Build, ContentType}` struct.
- `RunResurrectOptions` gained `MaxItems`, `MaxTotalChargeUsd`, `RestartOnError`.
- `RunClient.GetWithWait` and `BuildClient.GetWithWait` expose the spec's `waitForFinish` query parameter for a server-side synchronous fetch.
- `RunCollectionClient.List` `Status` filter accepts multiple statuses (`[]string`, comma-separated).
- `RequestQueueClient.BatchAddRequests` auto-chunks inputs at the API's 25-per-call limit and returns a typed `BatchAddResult{ProcessedRequests, UnprocessedRequests}`.
- `RequestQueueClient.ListRequests` validates its options (mutually-exclusive `ExclusiveStartID`/`Cursor`; `Filter` restricted to `locked`/`pending`).
- Tasks use a dedicated `TaskStartOptions` that omits the Actor-only `contentType` and `forcePermissionLevel` fields (which the task run endpoint does not accept).
- `LogClient` exposes the spec's `raw`/`download` options via `GetWithOptions`/`StreamWithOptions`, and `RunClient.GetStreamedLog` provides a raw live-stream convenience.
- `RunClient.Abort` now takes `gracefully *bool` instead of `bool` (`nil` omits the `gracefully` query parameter).
