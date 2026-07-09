# Builds

A build is a compiled version of an Actor. Access all builds with `client.Builds()`, a
single build with `client.Build(id)`, and an Actor's builds with `client.Actor(id).Builds()`.

## Build collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions) (PaginationList[Build], error)` | List builds. |

## Single build

| Method | Description |
| --- | --- |
| `Get(ctx) (Build, bool, error)` | Fetch the build. |
| `Abort(ctx) (Build, error)` | Abort the build. |
| `Delete(ctx) error` | Delete the build. |
| `WaitForFinish(ctx, waitSecs *int64) (Build, error)` | Poll until the build is terminal. |
| `GetOpenAPIDefinition(ctx) (json.RawMessage, bool, error)` | The build's generated OpenAPI document. |
| `Log() *LogClient` | The build's log. |

```go
build, err := client.Actor(actorID).Build(ctx, "0.0", apify.ActorBuildOptions{})
if err != nil {
	log.Fatal(err)
}
finished, err := client.Build(build.ID).WaitForFinish(ctx, apify.Ptr(int64(300)))
if err != nil {
	log.Fatal(err)
}
fmt.Println("build status:", finished.Status)

logText, ok, err := client.Build(build.ID).Log().Get(ctx)
```

### `Build` fields

The `Build` value returned by the build methods carries the build's metadata:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique build ID. |
| `ActID` | `string` | ID of the Actor this build belongs to. |
| `Status` | `ActorJobStatus` | Build status. One of the `ActorJobStatus` constants (shared with runs; raw wire values `READY`, `RUNNING`, `SUCCEEDED`, `FAILED`, `TIMING-OUT`, `TIMED-OUT`, `ABORTING`, `ABORTED`). |
| `StartedAt` | `*time.Time` | When the build started. |
| `FinishedAt` | `*time.Time` | When the build finished (`nil` while still building). |
| `BuildNumber` | `string` | Human-readable build number (e.g. `"0.1.2"`). |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API (forward compatibility). |

`Build.IsTerminal()` reports whether a build has finished.
