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
finished, err := client.Build(build.ID).WaitForFinish(ctx, ptr(int64(300)))
if err != nil {
	log.Fatal(err)
}
fmt.Println("build status:", finished.Status)

logText, ok, err := client.Build(build.ID).Log().Get(ctx)
```

`Build.IsTerminal()` reports whether a build has finished.
