# Store, users, and logs

## Apify Store

Browse public Actors with `client.Store()`:

| Method | Description |
| --- | --- |
| `List(ctx, StoreListOptions) (PaginationList[ActorStoreListItem], error)` | One page of Store Actors. |
| `Iterate(StoreListOptions, chunkSize *int64) *StoreActorIterator` | Lazy iterator over matching Actors. `Limit` caps the total yielded; `chunkSize` is the page size. |

`StoreListOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `Offset` | `*int64` | Number of Actors to skip. |
| `Limit` | `*int64` | Maximum number of Actors to return. |
| `Search` | `*string` | Full-text search query. |
| `SortBy` | `*string` | Sort field (e.g. `"popularity"`, `"newest"`). |
| `Category` | `*string` | Filter Actors by category. |
| `Username` | `*string` | Filter Actors by owner username. |
| `PricingModel` | `*string` | Filter by pricing model. Accepted values: `FREE`, `FLAT_PRICE_PER_MONTH`, `PRICE_PER_DATASET_ITEM`, `PAY_PER_EVENT`. |
| `IncludeUnrunnableActors` | `*bool` | Include Actors the current user cannot run. |
| `AllowsAgenticUsers` | `*bool` | Filter to Actors that allow agentic users. |
| `ResponseFormat` | `*string` | Select the response format. Accepted values: `full`, `agent`. |

Each item is an `ActorStoreListItem`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique Actor ID. |
| `Name` | `string` | Technical name of the Actor. |
| `Username` | `string` | Username of the Actor's owner. |
| `Title` | `string` | Human-readable title (may be empty; fall back to `Name`). |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API (forward compatibility). |

```go
// Limit caps the total number of Actors yielded; the second argument is the per-page size.
it := client.Store().Iterate(apify.StoreListOptions{Limit: apify.Ptr(int64(20)), Search: apify.Ptr("scraper")}, apify.Ptr(int64(10)))
for {
	actor, err := it.Next(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if actor == nil {
		break
	}
	fmt.Println(actor.Title, actor.ID)
}
```

## Users

`client.Me()` returns a client for the current account; `client.User(id)` for another user.

| Method | Description |
| --- | --- |
| `Get(ctx) (User, bool, error)` | Account details (private for `Me()`, public otherwise). |
| `MonthlyUsage(ctx) (json.RawMessage, error)` | Current month's usage (`Me()` only). |
| `MonthlyUsageForDate(ctx, date string) (json.RawMessage, error)` | Usage for the month containing `date` (YYYY-MM-DD, `Me()` only). |
| `Limits(ctx) (json.RawMessage, error)` | Current account's limits (`Me()` only). |
| `UpdateLimits(ctx, newLimits any) error` | Update the account's limits (`Me()` only). |

The `User` value returned by `Get` carries the account's public fields; for `Me()` the API
also returns private account details, which are preserved in `Extra`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique user ID. |
| `Username` | `string` | The user's username. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API (private details for `Me()`). |

The usage/limits methods return an error if called on a non-`Me()` client.

`MonthlyUsage` is equivalent to `MonthlyUsageForDate(ctx, "")`: an empty `date` omits the
parameter, so the API reports usage for the current month. Pass a `YYYY-MM-DD` date to select
the month containing that date. Both return the raw JSON usage report from the API — a JSON
object with the account's usage breakdown and totals for the period.

```go
user, ok, err := client.Me().Get(ctx)
if err != nil {
	log.Fatal(err)
}
if !ok {
	log.Fatal("account not found")
}
fmt.Println(user.Username)

usage, err := client.Me().MonthlyUsage(ctx)
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(usage)) // raw JSON usage report for the current month

// Usage for the month containing a specific date.
mayUsage, err := client.Me().MonthlyUsageForDate(ctx, "2026-05-15")
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(mayUsage)) // raw JSON usage report
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

`LogOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `Raw` | `*bool` | Return the unprocessed log content (no platform post-processing). |
| `Download` | `*bool` | Set `Content-Disposition` so the log is served as a download. |

For convenient live redirection of a run's log, `client.Run(id).GetStreamedLog(ctx)` returns
a raw live stream directly.

```go
// Read the whole log.
logText, ok, err := client.Run(runID).Log().Get(ctx)

// Or stream it live (log redirection).
stream, err := client.Run(runID).Log().Stream(ctx)
if err != nil {
	log.Fatal(err)
}
defer stream.Close()
_, _ = io.Copy(os.Stdout, stream)
```
