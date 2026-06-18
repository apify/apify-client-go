# Store, users, and logs

## Apify Store

Browse public Actors with `client.Store()`:

| Method | Description |
| --- | --- |
| `List(ctx, StoreListOptions) (PaginationList[ActorStoreListItem], error)` | One page of Store Actors. |
| `Iterate(StoreListOptions) *StoreActorIterator` | Lazy iterator over all matching Actors. |

`StoreListOptions`: `Offset`, `Limit`, `Search`, `SortBy`, `Category`, `Username`,
`PricingModel`, `IncludeUnrunnableActors`, `AllowsAgenticUsers`, `ResponseFormat`.

```go
it := client.Store().Iterate(apify.StoreListOptions{Limit: ptr(int64(20)), Search: ptr("scraper")})
for {
	actor, err := it.Next(ctx)
	if err != nil { log.Fatal(err) }
	if actor == nil { break }
	fmt.Println(actor.Title, actor.ID)
}
```

## Users

`client.Me()` returns a client for the current account; `client.User(id)` for another user.

| Method | Description |
| --- | --- |
| `Get(ctx) (User, bool, error)` | Account details (private for `Me()`, public otherwise). |
| `MonthlyUsage(ctx) (json.RawMessage, error)` | Current account's monthly usage (`Me()` only). |
| `Limits(ctx) (json.RawMessage, error)` | Current account's limits (`Me()` only). |
| `UpdateLimits(ctx, newLimits any) error` | Update the account's limits (`Me()` only). |

The usage/limits methods return an error if called on a non-`Me()` client.

```go
user, ok, err := client.Me().Get(ctx)
usage, err := client.Me().MonthlyUsage(ctx)
```

## Logs

`client.Log(buildOrRunID)` accesses a build or run log (also via `client.Run(id).Log()` /
`client.Build(id).Log()`):

| Method | Description |
| --- | --- |
| `Get(ctx) (string, bool, error)` | The entire log as text. |
| `GetWithOptions(ctx, LogOptions) (string, bool, error)` | Read with options (`Raw`, `Download`). |
| `Stream(ctx) (io.ReadCloser, error)` | A live stream of the log; close when done. |
| `StreamWithOptions(ctx, LogOptions) (io.ReadCloser, error)` | Stream with options (`Raw`, `Download`). |

For convenient live redirection of a run's log, `client.Run(id).GetStreamedLog(ctx)` returns
a raw live stream directly.

```go
// Read the whole log.
logText, ok, err := client.Run(runID).Log().Get(ctx)

// Or stream it live (log redirection).
stream, err := client.Run(runID).Log().Stream(ctx)
if err != nil { log.Fatal(err) }
defer stream.Close()
_, _ = io.Copy(os.Stdout, stream)
```
