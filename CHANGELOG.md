# Changelog

All notable changes to the Apify Go client are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.1] - 2026-06-29

Documentation-only compliance pass against the current orchestration requirements (no spec bump;
the client remains on `v2-2026-06-25T142310Z`). No change to the public API surface, so this is a
SemVer patch.

### Changed

- Documentation now states the client is **official, but experimental** and **AI-generated and
  AI-maintained**, replacing the previous "not (yet) an officially supported Apify product"
  wording that contradicted the requirement. Updated consistently in `README.md`, `docs/README.md`,
  and the package-level doc comment in `client.go`.
- `README.md` "Releasing" section now explicitly documents that Go's publishing process has no
  central registry and therefore no "Trusted Publisher" mechanism — releases are cut purely by
  pushing a git tag read by the public Go module proxy, authenticated only by the built-in
  `GITHUB_TOKEN`. The `go-publish.yml` workflow already reflected this; the README now matches.
- Documentation completeness pass (no API change): added a field table for `*APIError` and
  documented `AsAPIError` in `README.md`; added a `RunChargeOptions` field table and a note that
  `MetamorphOptions`/`RunChargeOptions` use plain (non-pointer) values in `docs/runs.md`;
  enumerated the webhook `Create` input keys in `docs/webhooks.md`; documented the request-queue
  locking method signatures (`ListAndLockHead`, `ProlongRequestLock`, `DeleteRequestLock`,
  `UnlockRequests`) in `docs/storages.md`.

### Fixed

- Corrected a stale `API_SPEC_VERSION` reference in the `README.md` "Versioning" section
  (`v2-2026-06-24T105326Z` → `v2-2026-06-25T142310Z`) so it matches `version.go`.
- `docs/storages.md`: `DownloadItems` now lists `JSONL` in its format list, matching the format
  constants table (previously the prose omitted it).
- Bumped `CLIENT_VERSION` to `0.4.1` (patch; documentation-only, no public API change).

## [0.4.0] - 2026-06-26

Updated to Apify OpenAPI specification `v2-2026-06-25T142310Z` (previously
`v2-2026-06-24T105326Z`). An operation- and parameter-level audit against the new specification
found no changes to the in-scope API surface; the spec update itself is a version bump only. This
release adds an `origin` filter to the last-run convenience accessors (additively, with no
breaking change) for parity with the `apify-client-js` reference, and cleans up stale in-code
comments.

### Added

- `ActorClient.LastRunWithOptions` and `TaskClient.LastRunWithOptions` accept a `LastRunOptions`
  with `Status` and `Origin` filters, matching the reference client's `lastRun({ status, origin })`.
  The existing `LastRun(status string)` accessors are unchanged (they delegate to the new methods
  with only `Status` set), so this is a purely additive, non-breaking change. `Origin` is a
  reference-client convenience threaded to the same `runs/last` endpoint; the OpenAPI spec does not
  document it as a query parameter.

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-06-25T142310Z`.
- Bumped `CLIENT_VERSION` to `0.4.0` (minor bump per SemVer for the additive `LastRunWithOptions`
  API).

### Fixed

- Cleaned up stale in-code comments around the `isAtHome` User-Agent flag that quoted the older,
  capitalized requirement wording; the comments now match the current lowercase requirement text.
  No behavior change — the flag already rendered lowercase (`true`/`false`).

## [0.3.0] - 2026-06-25

Updated to Apify OpenAPI specification `v2-2026-06-24T105326Z` (previously
`v2-2026-06-23T113219Z`). A full operation- and parameter-level audit of every in-scope endpoint
against the new specification found no changes to the in-scope API surface (same 131 paths and
231 operations, identical parameters and request/response schemas); the spec update itself is a
version bump only. The new `/v2/browser-info` and `/v2/tools/*` operations in the spec are
out-of-scope (absent from the `apify-client-js` reference) and are intentionally not implemented.
This release also includes a spec-compliance fix surfaced by the review pass. The minor-version
bump (rather than patch) reflects the breaking change to `ListRequestsOptions.Filter` below.

### Fixed

- `ListRequestsOptions.Filter` (for `GET /v2/request-queues/{queueId}/requests`) is now
  `[]string` and serialized comma-joined, matching the spec (an array of the enum values
  `locked`/`pending`, `style=form` `explode=false`) and the JS reference (`filter.join(',')`).
  Previously it was a single `*string`, which could not express the multi-value union. This is a
  spec-compliance bugfix to a type that did not match the specification; it is a breaking change to
  that field's type. The allowed values are now exported as the `RequestFilterLocked` /
  `RequestFilterPending` constants.
- The `User-Agent` `isAtHome` flag is now based **solely** on the `APIFY_IS_AT_HOME` environment
  variable, as mandated by the requirements and matching the JS reference (which reads only that
  variable). A previously-honored bare `isAtHome` environment variable is no longer consulted; it
  was a non-standard accommodation not present in the requirements or the reference.

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-06-24T105326Z`.
- Bumped `CLIENT_VERSION` to `0.3.0` (minor bump per SemVer because of the breaking `Filter`
  type change above).

## [0.2.3] - 2026-06-23

Updated to Apify OpenAPI specification `v2-2026-06-23T113219Z` (previously
`v2-2026-06-18T095846Z`). The spec delta over the in-scope API surface is fully additive — no
new in-scope endpoints and no removed or renamed parameters — so there are no breaking changes
to the public interface. Response models already capture unmodelled API fields in their `Extra`
maps, so additive schema fields decode cleanly.

### Changed

- Bumped `API_SPEC_VERSION` to `v2-2026-06-23T113219Z`.
- Bumped `CLIENT_VERSION` to `0.2.3`.

## [0.2.2] - 2026-06-22

Publishing compliance for the updated client requirements (apify-client-orchestration PR #7).
No changes to the public interface or to the OpenAPI spec version (`v2-2026-06-18T095846Z`);
packaging metadata, a release workflow, and documentation only.

### Added

- CI: a manually triggered (`workflow_dispatch`) `Publish Go client` workflow
  (`.github/workflows/go-publish.yml`) that releases the module. Go is distributed by pushing a
  semver git tag (there is no upload registry), so the workflow runs the same gofmt/vet/lint/build/
  unit-test quality gate as CI, derives the tag from the `CLIENT_VERSION` constant in `version.go`
  (the single source of truth), refuses to proceed if that tag already exists, then creates and
  pushes the `vX.Y.Z` tag, opens a GitHub release, and pings the public Go module proxy to trigger
  pkg.go.dev indexing. It uses only the built-in `GITHUB_TOKEN` repository secret — no registry
  credential is needed for Go — and supports a `dry_run` input that runs every check without
  creating the release. Mirrors the Rust sibling's `Publish Rust client to crates.io` workflow.

### Changed

- Bumped `CLIENT_VERSION` to `0.2.2`.
- Documentation: added a "Releasing" subsection to the README "Versioning" section describing the
  tag-based Go distribution mechanism and the publish workflow.
- Documentation: added the required "experimental, AI-generated and AI-maintained" disclaimer to
  the top of `README.md` and `docs/README.md`, and softened the "official" wording accordingly, to
  satisfy the documentation requirement and match the Rust sibling.

## [0.2.1] - 2026-06-19

Documentation and CI compliance with the updated client requirements. No changes to the
public API or to the OpenAPI spec version (`v2-2026-06-18T095846Z`), so there are no breaking
changes.

### Added

- CI: a standalone `Test examples` step in the Go integration workflow that actually runs the
  documentation example code end-to-end. It executes the example programs in `examples/` against
  the live API (the `TestExample*` smoke tests, each running `go run ./examples/<name>`) and
  validates that every in-documentation `go` snippet is valid, runnable, and gofmt-formatted
  (the new `TestDocSnippets*` tests, which extract every ```go block from the README and `docs/`
  and compile each one). The `Integration tests` step now skips these so the two concerns stay
  separate, mirroring the Rust sibling client.
- Tests: `tests/docs_snippets_test.go`, an offline doc-snippet harness (Go has no Markdown
  doctest equivalent) that enforces the requirement that each in-documentation code snippet is
  valid, runnable, and properly formatted.

### Changed

- Workflow now also triggers on `docs/**` and `README.md` changes, so documentation edits
  re-run the snippet validation.

### Fixed

- Reformatted all `docs/` and README code snippets to canonical gofmt output (one-line
  `if err != nil { ... }` blocks expanded, tab indentation, aligned trailing comments) and made
  the custom-HTTP-transport snippet a complete, compilable program.
- Corrected the README versioning note, which referenced the older spec version
  `v2-2026-06-16T064758Z` instead of the current `v2-2026-06-18T095846Z`.

### Documentation

- Documented the shared `ListOptions` type (fields + example) in `docs/README.md`, which several
  resource pages reference as a method argument but which was previously undefined in the docs.
- Documented the `apify.ListDatasetItems[T]` generic helper's argument types and added a runnable
  typed-decoding example in `docs/storages.md`.
- Added explicit field listings for the `ActorRun`, `Build`, `User`, and `ActorStoreListItem`
  response models to `docs/runs.md`, `docs/builds.md`, and `docs/misc.md`.
- Added full field listings in `docs/storages.md` for the storage option/parameter structs that
  were previously named in method tables but not enumerated: `DatasetListItemsOptions`,
  `DatasetDownloadOptions`, `ListKeysOptions`, `GetRecordOptions`, `GetRecordsOptions`,
  `ListRequestsOptions`, and the `RequestQueueRequest` payload.
- Documented the storage *return* types that examples dereference: `KeyValueStoreRecord`,
  `KeyValueStoreKeysPage` (and `KeyValueStoreKey`), `RequestQueueHead`,
  `RequestQueueOperationInfo`, and `BatchAddResult` in `docs/storages.md`.
- Added field tables for the remaining response models (`Actor`, `Task`, `Schedule`, `Webhook`,
  `WebhookDispatch`) to their resource pages, matching the treatment of `ActorRun`/`Build`.
- Documented the accepted values / details of the enum-like parameters `RunListOptions.Status`,
  `ListRequestsOptions.Filter` (`"locked"`/`"pending"`), `DatasetListItemsOptions.View`, and the
  full `StorageListOptions` field table; and clarified in `docs/README.md` that the
  within-storage listers (`ListKeys`, `ListHead`) return their own page/head containers rather
  than `PaginationList[T]`.
- Expanded the run/Actor/store *input* option structs from bare name lists into full
  field/type/meaning tables: `ActorStartOptions` (including the nested ad-hoc `Webhooks` element
  shape and `ForcePermissionLevel`), `ActorBuildOptions`, `ActorListOptions`, `StoreListOptions`
  (with enum values for `SortBy`/`PricingModel`), `RunResurrectOptions`, `MetamorphOptions`, and
  `LogOptions`; and made `StorageListOptions.Ownership` state its accepted values definitively.
- Documented the schedule `actions` payload shape with a runnable `RUN_ACTOR` action example in
  `docs/schedules.md`, replacing the empty `[]any{}` placeholder.
- Stated the closed enum sets definitively (verified against the OpenAPI spec) instead of hedging
  with "e.g.": `StoreListOptions.PricingModel` (`FREE`, `FLAT_PRICE_PER_MONTH`,
  `PRICE_PER_DATASET_ITEM`, `PAY_PER_EVENT` — previously omitted `PAY_PER_EVENT`),
  `StoreListOptions.ResponseFormat` (`full`, `agent` — previously unspecified),
  `ActorStartOptions.ForcePermissionLevel` (`LIMITED_PERMISSIONS`, `FULL_PERMISSIONS`), and
  `ActorListOptions.SortBy` (`createdAt`, `stats.lastRunStartedAt`).
- Corrected and completed the run/build status enum documentation: the canonical `ActorJobStatus`
  enum has eight values, but the docs (and the `ActorRun.Status` in-code comment) listed only six
  for runs and four for builds. Now `ActorRun.Status`, `RunListOptions.Status`, and `Build.Status`
  all document the full set `READY`, `RUNNING`, `SUCCEEDED`, `FAILED`, `TIMING-OUT`, `TIMED-OUT`,
  `ABORTING`, `ABORTED`. (Behavior unchanged: `IsTerminal` still treats only the four terminal
  states as finished, which is correct.)
- Documented the closed enums in the actor/version `Create` definition (`sourceType` =
  `VersionSourceType`: `SOURCE_FILES`/`GIT_REPO`/`TARBALL`/`GITHUB_GIST`/`SOURCE_CODE`; source-file
  `format` = `TEXT`/`BASE64`) with a runnable `SOURCE_FILES` example in `docs/actors.md`.
- Enumerated the closed 12-value `WebhookEventType` set in `docs/webhooks.md` and cross-referenced
  it from the `ActorStartOptions.Webhooks` note.

## [0.2.0] - 2026-06-19

Verified against OpenAPI specification version `v2-2026-06-18T095846Z` (bumped from
`v2-2026-06-16T064758Z`). The spec delta is purely additive — two optional query
parameters — so there are no breaking changes.

### Added

- `ActorClient.ValidateInputForBuild(ctx, input, build)` exposes the new optional `build`
  query parameter on `POST /v2/actors/{actorId}/validate-input`, validating input against the
  input schema of a specific Actor build (tag or number). `ValidateInput` is unchanged and
  now delegates with an empty build (omitting the parameter, so the API validates against the
  build tagged `latest`, per the OpenAPI specification).
- `UserClient.MonthlyUsageForDate(ctx, date)` exposes the new optional `date` query parameter
  (YYYY-MM-DD) on `GET /v2/users/me/usage/monthly`, selecting the month to report usage for.
  `MonthlyUsage` is unchanged and now delegates with an empty date (omitting the parameter, so
  the current month is reported).

### Changed

- `API_SPEC_VERSION` bumped to `v2-2026-06-18T095846Z`.

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
  _(Superseded in [0.3.0]: the flag is now based solely on `APIFY_IS_AT_HOME`; the bare `isAtHome`
  variable is no longer consulted.)_
- `Dataset`/`KeyValueStore` `GetOrCreate` take only a name: this spec version's create
  endpoints declare no request body or `schema` parameter, so a schema argument would be an
  undocumented extension. The `runs/last` endpoints accept only `status` (no `origin`) in
  this spec version. These choices favour strict OpenAPI compliance over the JS superset.
- CI runs `golangci-lint` (default high-signal analyzers — errcheck, govet, ineffassign,
  staticcheck, unused — plus `misspell`) in addition to `gofmt`, `go vet`, and `go build`,
  satisfying the coding-rule mandate that linting run in CI. Config in `.golangci.yml`.
