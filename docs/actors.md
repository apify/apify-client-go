# Actors

Actors are the programs that run on the Apify platform. Access the Actor collection with
`client.Actors()` and a single Actor with `client.Actor(id)`, where `id` is the Actor ID or
`username~name`.

## Actor collection

| Method | Description |
| --- | --- |
| `List(ctx, ActorListOptions) (PaginationList[Actor], error)` | List the account's Actors. |
| `Create(ctx, definition any) (Actor, error)` | Create a new Actor. |

`ActorListOptions`: `Offset`, `Limit`, `Desc`, `My`, `SortBy` (all optional pointers).

```go
page, err := client.Actors().List(ctx, apify.ActorListOptions{My: ptr(true), Limit: ptr(int64(10))})
```

## Single Actor

| Method | Description |
| --- | --- |
| `Get(ctx) (Actor, bool, error)` | Fetch the Actor. |
| `Update(ctx, newFields any) (Actor, error)` | Update the Actor. |
| `Delete(ctx) error` | Delete the Actor. |
| `Start(ctx, input any, ActorStartOptions) (ActorRun, error)` | Start a run, return immediately. |
| `Call(ctx, input any, ActorStartOptions, waitSecs *int64) (ActorRun, error)` | Start and wait for the run to finish. |
| `Build(ctx, versionNumber string, ActorBuildOptions) (Build, error)` | Build a version. |
| `DefaultBuild(ctx, waitForFinish *int64) (*BuildClient, error)` | Resolve the default build. |
| `ValidateInput(ctx, input any) (json.RawMessage, error)` | Validate input against the schema. |
| `LastRun(status string) *RunClient` | Client for the last run (optional status filter). |
| `Builds() *BuildCollectionClient` | This Actor's builds. |
| `Runs() *RunCollectionClient` | This Actor's runs. |
| `Version(n) *ActorVersionClient` / `Versions() *ActorVersionCollectionClient` | Versions. |
| `Webhooks() *WebhookCollectionClient` | This Actor's webhooks. |

`ActorStartOptions`: `Build`, `MemoryMbytes`, `TimeoutSecs`, `WaitForFinish`, `MaxItems`,
`MaxTotalChargeUsd`, `ContentType`, `RestartOnError`, `ForcePermissionLevel`, `Webhooks`.

`ActorBuildOptions`: `BetaPackages`, `Tag`, `UseCache`, `WaitForFinish`.

```go
// Start an Actor with input and wait for it to finish.
waitSecs := int64(120)
run, err := client.Actor("apify/hello-world").Call(ctx,
	map[string]any{"message": "hi"},
	apify.ActorStartOptions{MemoryMbytes: ptr(int64(512))},
	&waitSecs,
)
```

## Versions and environment variables

`client.Actor(id).Versions()` and `.Version(n)`:

| Method | Description |
| --- | --- |
| `Versions().List(ctx, ListOptions)` / `Versions().Create(ctx, def)` | List/create versions. |
| `Version(n).Get/Update/Delete(ctx)` | Manage a single version. |
| `Version(n).EnvVars().List(ctx)` / `.Create(ctx, ActorEnvVar)` | List/create env vars. |
| `Version(n).EnvVar(name).Get/Update/Delete(ctx)` | Manage a single env var. |

```go
v := client.Actor(actorID).Version("0.0")
_, err := v.EnvVars().Create(ctx, apify.ActorEnvVar{Name: "API_KEY", Value: "secret", IsSecret: ptr(true)})
```

> `ptr` is a helper returning a pointer to its argument: `func ptr[T any](v T) *T { return &v }`.
