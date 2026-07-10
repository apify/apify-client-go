# Actors

Actors are the programs that run on the Apify platform. Access the Actor collection with
`client.Actors()` and a single Actor with `client.Actor(id)`, where `id` is the Actor ID or
`username~name`.

## Actor collection

| Method | Description |
| --- | --- |
| `List(ctx, ActorListOptions) (PaginationList[Actor], error)` | List the account's Actors. |
| `Iterate(ActorListOptions, chunkSize *int64) *ListIterator[Actor]` | Lazy iterator over matching Actors. `Limit` caps the total yielded; `chunkSize` is the page size. |
| `Create(ctx, definition any) (Actor, error)` | Create a new Actor. |

`Create` takes a free-form definition (`any`) serialized to JSON, so the Actor's fields are
passed as a map (`name`, `title`, `versions`, etc.). A version's `sourceType` selects how its
source is supplied, and is one of the closed `VersionSourceType` values: `SOURCE_FILES`,
`GIT_REPO`, `TARBALL`, `GITHUB_GIST`, `SOURCE_CODE`. For `SOURCE_FILES`, each entry in
`sourceFiles` has a `format` of `TEXT` or `BASE64`. A minimal `SOURCE_FILES` Actor:

```go
actor, err := client.Actors().Create(ctx, map[string]any{
	"name":  "my-actor",
	"title": "My Actor",
	"versions": []any{
		map[string]any{
			"versionNumber": "0.0",
			"sourceType":    "SOURCE_FILES", // VersionSourceType: SOURCE_FILES|GIT_REPO|TARBALL|GITHUB_GIST|SOURCE_CODE
			"sourceFiles": []any{
				map[string]any{
					"name":    "Dockerfile",
					"format":  "TEXT", // SourceCodeFileFormat: TEXT|BASE64
					"content": "FROM apify/actor-node:20\n",
				},
			},
		},
	},
})
if err != nil {
	log.Fatal(err)
}
_ = actor
```

`ActorListOptions` (all fields optional pointers):

| Field | Type | Meaning |
|---|---|---|
| `Offset` | `*int64` | Number of Actors to skip. |
| `Limit` | `*int64` | Maximum number of Actors to return. |
| `Desc` | `*bool` | Return Actors newest-first. |
| `My` | `*bool` | Return only Actors owned by the current user. |
| `SortBy` | `*string` | Sort field. Accepted values: `createdAt`, `stats.lastRunStartedAt`. |

```go
page, err := client.Actors().List(ctx, apify.ActorListOptions{My: apify.Ptr(true), Limit: apify.Ptr(int64(10))})
```

### `Actor` fields

The `Actor` value returned by `Get`/`Create`/`Update` and listed by `List`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique Actor ID. |
| `UserID` | `string` | ID of the user who owns the Actor. |
| `Name` | `string` | Technical name of the Actor (used in API paths). |
| `Username` | `string` | Username of the Actor's owner. |
| `Title` | `string` | Human-readable title shown in the UI. |
| `Description` | `string` | What the Actor does. |
| `IsPublic` | `bool` | Whether the Actor is public in Apify Store. |
| `CreatedAt` | `*time.Time` | When the Actor was created. |
| `ModifiedAt` | `*time.Time` | When the Actor was last modified. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

## Single Actor

| Method | Description |
| --- | --- |
| `Get(ctx) (Actor, bool, error)` | Fetch the Actor. |
| `Update(ctx, newFields any) (Actor, error)` | Update the Actor. |
| `Delete(ctx) error` | Delete the Actor. |
| `Start(ctx, input any, ActorStartOptions) (ActorRun, error)` | Start a run, return immediately. |
| `Call(ctx, input any, ActorStartOptions, waitSecs *int64) (ActorRun, error)` | Start and wait for the run to finish. `waitSecs` bounds the client-side wait; pass `nil` to wait indefinitely (until the run reaches a terminal state). |
| `Build(ctx, versionNumber string, ActorBuildOptions) (Build, error)` | Build a version. |
| `DefaultBuild(ctx, waitForFinish *int64) (*BuildClient, error)` | Resolve the default build. |
| `ValidateInput(ctx, input any) (json.RawMessage, error)` | Validate input against the `latest` build's schema. |
| `ValidateInputForBuild(ctx, input any, build string) (json.RawMessage, error)` | Validate input against a specific build's schema (tag or number). |
| `LastRun(status string) *RunClient` | Client for the last run (optional status filter). |
| `LastRunWithOptions(options LastRunOptions) *RunClient` | Client for the last run, filtered by status and/or origin. |
| `Builds() *BuildCollectionClient` | This Actor's builds. |
| `Runs() *RunCollectionClient` | This Actor's runs. |
| `Version(n) *ActorVersionClient` / `Versions() *ActorVersionCollectionClient` | Versions. |
| `Webhooks() *WebhookCollectionClient` | This Actor's webhooks. |

> **Note — two different `Build`s.** `Actor.Build(ctx, versionNumber, ActorBuildOptions)` here
> *starts* a build of a version and returns the resulting `Build`. It is unrelated to the
> top-level accessor `client.Build(id)`, which returns a `*BuildClient` for inspecting an
> existing build by ID (see [builds.md](builds.md)). Same name, different jobs.

`ValidateInput` is equivalent to `ValidateInputForBuild(ctx, input, "")`: an empty `build`
omits the parameter, so the API validates against the build tagged `latest` (per the API
specification). Both return the raw JSON validation result from the API — a JSON object
reporting whether the input is valid and, if not, the schema violations.

`LastRun(status)` returns the Actor's most recent run, optionally narrowed to a status. Use
`LastRunWithOptions` to also narrow by origin (how the run was started). Both return a
`RunClient` for the resolved run, so you can chain `.Get(ctx)`, `.Dataset()`, etc.

`LastRunOptions` (all fields optional; an empty field leaves that filter unset):

| Field | Type | Meaning |
|---|---|---|
| `Status` | `string` | Filter by run status (e.g. `SUCCEEDED`, `FAILED`, `RUNNING`). |
| `Origin` | `string` | Filter by how the run was started: `DEVELOPMENT`, `WEB`, `API`, `SCHEDULER`, `TEST`, `WEBHOOK`, `ACTOR`, `CLI`, `CI`, `STANDBY`, `MCP`. |

```go
// Most recent run that both SUCCEEDED and was started via the API.
lastRun, ok, err := client.Actor("apify/hello-world").
	LastRunWithOptions(apify.LastRunOptions{Status: "SUCCEEDED", Origin: "API"}).
	Get(ctx)
if err != nil {
	log.Fatal(err)
}
if ok {
	fmt.Printf("last API run: %s (%s)\n", lastRun.ID, lastRun.Status)
}
```

```go
// Validate input against a specific build's input schema.
result, err := client.Actor("apify/hello-world").ValidateInputForBuild(ctx,
	map[string]any{"message": "hi"},
	"latest",
)
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(result)) // raw JSON validation result
```

`ActorStartOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `Build` | `*string` | Tag or number of the build to run (e.g. `"latest"`, `"0.1.2"`). |
| `MemoryMbytes` | `*int64` | Memory in megabytes allocated for the run. |
| `TimeoutSecs` | `*int64` | Run timeout in seconds (`0` means no timeout). |
| `WaitForFinish` | `*int64` | Max seconds to wait server-side for the run to finish (max 60). |
| `MaxItems` | `*int64` | Maximum dataset items to charge (pay-per-result Actors). |
| `MaxTotalChargeUsd` | `*float64` | Maximum total charge in USD (pay-per-event Actors). |
| `ContentType` | `*string` | Content type of the input body (default `application/json`). |
| `RestartOnError` | `*bool` | Restart the run if it fails. |
| `ForcePermissionLevel` | `*string` | Override the Actor's permission level for this run. Accepted values: `LIMITED_PERMISSIONS`, `FULL_PERMISSIONS`. |
| `Webhooks` | `[]any` | Ad-hoc webhooks to attach to this run. Each element is a map describing one webhook: `{"eventTypes": []string, "requestUrl": string, "payloadTemplate": string}` (same shape as a webhook definition; `eventTypes` are the `WebhookEventType` values listed in [webhooks.md](webhooks.md)). |

`ActorBuildOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `BetaPackages` | `*bool` | Use beta versions of Apify packages. |
| `Tag` | `*string` | Tag to apply to the build (e.g. `"latest"`). |
| `UseCache` | `*bool` | Whether to use the Docker build cache (default `true`). |
| `WaitForFinish` | `*int64` | Max seconds to wait server-side for the build (max 60). |

```go
// Start an Actor with input and wait for it to finish.
waitSecs := int64(120)
run, err := client.Actor("apify/hello-world").Call(ctx,
	map[string]any{"message": "hi"},
	apify.ActorStartOptions{MemoryMbytes: apify.Ptr(int64(512))},
	&waitSecs,
)
```

> **How `Call`/`WaitForFinish` relate to `WithTimeout`.** The wait is done client-side by
> *polling*: the client repeatedly re-fetches the run, each poll being a separate HTTP request
> that asks the server to block for at most 60 seconds (the API's per-request wait cap). Because
> every poll returns within that 60-second server cap — comfortably inside the client's
> `WithTimeout` budget (default 360s, which bounds each individual request, not the total wait) —
> the overall wait is **not** cut off at 360s. It continues across as many polls as needed until
> the run reaches a terminal state. Passing `waitSecs == nil` therefore genuinely waits until the
> run finishes (bounded only by a very large internal cap of ~11.5 days, or by cancelling the
> `ctx` you pass in); a non-nil `waitSecs` bounds the total client-side wait instead. The distinct
> `ActorStartOptions.WaitForFinish` field is unrelated: it only controls the single server-side
> wait on the initial `Start` request (max 60s), not the client-side polling loop.

## Versions and environment variables

`client.Actor(id).Versions()` and `.Version(n)`:

| Method | Description |
| --- | --- |
| `Versions().List(ctx, ListOptions)` / `Versions().Iterate(ListOptions, chunkSize *int64)` / `Versions().Create(ctx, def)` | List/iterate/create versions. |
| `Version(n).Get/Update/Delete(ctx)` | Manage a single version. |
| `Version(n).EnvVars().List(ctx)` / `.Iterate()` / `.Create(ctx, ActorEnvVar)` | List/iterate/create env vars. |
| `Version(n).EnvVar(name).Get/Update/Delete(ctx)` | Manage a single env var. |

`ActorEnvVar` fields:

| Field | Type | Meaning |
|---|---|---|
| `Name` | `string` | The environment variable name (required). |
| `Value` | `string` | The environment variable value. |
| `IsSecret` | `*bool` | Whether the value is stored as a secret (encrypted at rest). |

```go
v := client.Actor(actorID).Version("0.0")
_, err := v.EnvVars().Create(ctx, apify.ActorEnvVar{Name: "API_KEY", Value: "secret", IsSecret: apify.Ptr(true)})
```

> `apify.Ptr` is an exported helper returning a pointer to its argument
> (`func Ptr[T any](v T) *T`). Use it to set the optional, pointer-typed option fields inline.
