# Runs

A run is a single execution of an Actor. Access all runs with `client.Runs()`, a single run
with `client.Run(id)`, and an Actor's or task's runs with `client.Actor(id).Runs()` /
`client.Task(id).Runs()`.

## Run collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions, RunListOptions) (PaginationList[ActorRun], error)` | List runs. |
| `Iterate(ListOptions, RunListOptions, chunkSize *int64) *ListIterator[ActorRun]` | Lazy iterator over matching runs. `Limit` caps the total yielded; `chunkSize` is the page size. |

`RunListOptions`:

| Field | Type | Meaning |
|---|---|---|
| `Status` | `[]string` | Filter by one or more run statuses; sent as a comma-separated list. Values are the eight `ActorJobStatus` values: `READY`, `RUNNING`, `SUCCEEDED`, `FAILED`, `TIMING-OUT`, `TIMED-OUT`, `ABORTING`, `ABORTED`. |
| `StartedAfter` | `*string` | Only runs started after this ISO-8601 timestamp. |
| `StartedBefore` | `*string` | Only runs started before this ISO-8601 timestamp. |
| `Deployment` | `[]string` | Only Actor Standby runs served by these pools (`SINGLE_TENANT`, `MULTI_TENANT_FREE_USERS`, `MULTI_TENANT_PAYING_USERS`, `MULTI_TENANT_OWNER_PAID`); sent as a comma-separated list. The API requires `Status` to hold only non-terminal statuses (`READY`, `RUNNING`, `TIMING-OUT`, `ABORTING`) with it. |

`StartedAfter`, `StartedBefore` and `Deployment` apply only to the account-wide (`client.Runs()`) and Actor-scoped (`client.Actor(id).Runs()`) collections, not to a task's.

## Single run

| Method | Description |
| --- | --- |
| `Get(ctx) (ActorRun, bool, error)` | Fetch the run. |
| `GetWithWait(ctx, waitForFinishSecs *int64) (ActorRun, bool, error)` | Fetch, optionally waiting server-side. |
| `Update(ctx, newFields any) (ActorRun, error)` | Update the run (e.g. status message). |
| `Delete(ctx) error` | Delete the run. |
| `Abort(ctx, gracefully *bool) (ActorRun, error)` | Abort the run. Pass `nil` to use the server default (immediate), or a pointer to `true`/`false` to abort gracefully/immediately. |
| `Metamorph(ctx, targetActorID string, input any, MetamorphOptions) (ActorRun, error)` | Metamorph into another Actor. |
| `MetamorphRaw(ctx, targetActorID string, input []byte, MetamorphOptions) (ActorRun, error)` | Like `Metamorph`, but sends `input` as a raw request body (no JSON serialization), paired with `MetamorphOptions.ContentType`. |
| `Reboot(ctx) (ActorRun, error)` | Reboot the run. |
| `Resurrect(ctx, RunResurrectOptions) (ActorRun, error)` | Resurrect a finished run (status `SUCCEEDED`, `FAILED`, `ABORTED` or `TIMED-OUT`); it restarts with the same storages. |
| `Charge(ctx, RunChargeOptions) error` | Charge a pay-per-event run (sends an idempotency key). |
| `WaitForFinish(ctx, waitSecs *int64) (ActorRun, error)` | Poll until the run is terminal. |
| `Dataset() *DatasetClient` | The run's default dataset. |
| `KeyValueStore() *KeyValueStoreClient` | The run's default key-value store. |
| `RequestQueue() *RequestQueueClient` | The run's default request queue. |
| `Log() *LogClient` | The run's log. |
| `GetStreamedLog(ctx) (io.ReadCloser, error)` | Live stream of the run's raw log, as bytes. |
| `NewStreamedLog(ctx, StreamedLogOptions) (*StreamedLog, bool, error)` | Redirects the run's log message by message to a writer; see [Log redirection](#log-redirection). The bool is `false` for a run that does not exist (with the default prefix). |
| `IterateDatasetItems(RunIterateDatasetItemsOptions) *RunDatasetItemsIterator[json.RawMessage]` | Iterates over the items of the run's default dataset while the run is still producing them; see [Live iteration of dataset items](#live-iteration-of-dataset-items). |

On a client from `LastRun(...)`, `Dataset()`, `KeyValueStore()`, `RequestQueue()` and `Log()` carry
the same `status` and `origin` filters, so they address the storages and the log of the run the
filters resolve to, not those of the newest run of any status.

> **Missing resources.** `Dataset()`/`KeyValueStore()`/`RequestQueue()`/`Log()` address a
> sub-resource with no ID of its own, so a 404 from `Get`/`Delete`/`Log().Get`/`Log().Stream`
> on the client they return can mean either the run or the default sub-resource is missing —
> the response does not say which, so the error propagates instead of resolving to "absent".
> `KeyValueStore().GetRecord`/`RecordExists` and `RequestQueue().GetRequest` are unaffected:
> there a 404 unambiguously means the record/request itself is missing.

`RunResurrectOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `Build` | `*string` | Tag or number of the build to resurrect with. |
| `MemoryMbytes` | `*int64` | Memory in megabytes to allocate. |
| `TimeoutSecs` | `*int64` | Run timeout in seconds. |
| `MaxItems` | `*int64` | Maximum dataset items to charge (pay-per-result Actors). |
| `MaxTotalChargeUsd` | `*float64` | Maximum total charge in USD (pay-per-event Actors). |
| `RestartOnError` | `*bool` | Restart the run if it fails. |

`MetamorphOptions`:

| Field | Type | Meaning |
|---|---|---|
| `Build` | `string` | Pin the target Actor's build (empty for default). |
| `ContentType` | `string` | Content type of the input body. Defaults to `application/json` for `Metamorph`, or `application/octet-stream` for `MetamorphRaw`, when empty. |

`RunChargeOptions`:

| Field | Type | Meaning |
|---|---|---|
| `EventName` | `string` | Name of the event to charge for. **Required.** |
| `Count` | `*int64` | Number of times to charge the event (defaults to 1 when nil). |
| `IdempotencyKey` | `string` | Deduplicates the charge across retries; auto-generated when empty, so a retried charge is applied at most once. The API requires this header on every charge and expires the key 3 minutes after the charge. |
| `RequestID` | `string` | ID of the Server Actor request being charged, from the `X-Actor-Request-ID` header of the incoming request. Required for multi-tenant Server Actor runs, ignored elsewhere; omitted from the request when empty. |

> Note: unlike most option structs (whose optional fields are pointer-typed — see
> [docs/README.md](README.md)), `MetamorphOptions` and the string fields of `RunChargeOptions`
> use plain values. `Build`/`ContentType` and `EventName`/`IdempotencyKey`/`RequestID` treat
> the zero value (empty string) as "unset"; only `RunChargeOptions.Count` is a pointer.

```go
// nil polls until the run is terminal and is not cut off by WithTimeout (which bounds each
// poll request, not the total wait) — see docs/actors.md for the full explanation.
run, err := client.Run(runID).WaitForFinish(ctx, nil)
if err != nil {
	log.Fatal(err)
}

// Read the run's default dataset.
items, err := client.Run(run.ID).Dataset().ListItems(ctx, apify.DatasetListItemsOptions{})
```

### `ActorRun` fields

The `ActorRun` value returned by the run methods (and by `Actor(id).Call`/`Start`) carries the
run's metadata. The commonly used fields:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique run ID. |
| `ActID` | `string` | ID of the Actor that produced the run. |
| `ActorTaskID` | `string` | ID of the task that started the run, if any. |
| `UserID` | `string` | ID of the user who owns the run. |
| `Status` | `string` | Run status. One of the eight `ActorJobStatus` values: `READY`, `RUNNING`, `SUCCEEDED`, `FAILED`, `TIMING-OUT`, `TIMED-OUT`, `ABORTING`, `ABORTED`. |
| `StatusMessage` | `string` | Optional human-readable status message. |
| `StartedAt` | `*time.Time` | When the run started. |
| `FinishedAt` | `*time.Time` | When the run finished (`nil` while still running). |
| `BuildID` | `string` | ID of the build used for the run. |
| `DefaultDatasetID` | `string` | ID of the run's default dataset. |
| `DefaultKeyValueStoreID` | `string` | ID of the run's default key-value store. |
| `DefaultRequestQueueID` | `string` | ID of the run's default request queue. |
| `ContainerURL` | `string` | URL of the run's container (for live access), normalized. |
| `ConsoleURL` | `string` | URL of the run in the Apify Console, normalized. Empty on items from `List`/`Iterate`. |
| `Output` | `map[string]string` | Output links of the run (name to normalized URL), generated from the output schema of the Actor build. `nil` if the build defines none, and on items from `List`/`Iterate`. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API (forward compatibility). |

`ActorRun.IsTerminal()` reports whether a run has finished. Status message convenience:
`client.SetStatusMessage(ctx, message, isTerminal)` updates the current run identified by the
`ACTOR_RUN_ID` environment variable.

## Log redirection

`ActorClient.Call` redirects the log of the run while it waits (set `ActorStartOptions.Log` to
change where it goes or turn it off). To redirect the log of any run, create a `StreamedLog`:

```go
streamed, found, err := client.Run(runID).NewStreamedLog(ctx, apify.StreamedLogOptions{Output: os.Stdout})
if err != nil {
	log.Fatal(err)
}
if !found {
	log.Fatal("run not found")
}
if err := streamed.Start(ctx); err != nil {
	log.Fatal(err)
}
// ... wait for the run to finish ...
if err := streamed.Stop(ctx); err != nil {
	log.Fatal(err)
}
```

`StreamedLogOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `Output` | `io.Writer` | Receives every message, one per line, with the prefix. Defaults to `os.Stderr`. Writes are serialized. |
| `Prefix` | `*string` | Written in front of every message. The default is `"{Actor name} runId:{run ID} -> "`, which costs two requests (the run and its Actor); a non-nil `Prefix`, an empty one included, skips them. |
| `FromStart` | `*bool` | `false` drops the messages that were logged before the redirection was created. Default `true`. |
| `Disabled` | `bool` | Turns the redirection off in `ActorStartOptions.Log`. |

The log is split into messages at the timestamps the platform puts in front of each of them, so a
message that spans several lines is written as one. `Start` begins the redirection in the
background, and it also ends when the `ctx` it got is done. `Stop` gives the messages in flight up
to one second, then aborts the stream; if the stream delivered nothing yet, it reads the whole log
in one request, so a short run that finished before the stream connected is not lost. The `ctx`
given to `Stop` bounds that wait and that read: when it is done first, `Stop` returns its error
and the rest of the log is not redirected, unless the `ctx` given to `Start` ended first, which
ends the redirection on purpose and is not an error. When both end at about the same time,
either result is possible. A second `Stop` while one is running returns an error. A log that has
no bytes yet (the run has not logged anything) is reopened, pausing from 500ms up to 5s, until its
first bytes arrive. `Stop` returns the error that ended the redirection early, if any; `Call`
reports it on the redirection's output and goes on.

## Live iteration of dataset items

`RunClient.IterateDatasetItems` (and the generic `apify.IterateRunDatasetItems[T]`, which decodes
every item into `T`) iterates over the items of the run's default dataset *while the run is still
producing them*. While the run has not finished, each round yields the rows below the dataset's
`itemCount` and then waits up to `PollIntervalSecs` for the run to finish, so the last rows are
read as soon as it does. `itemCount` lags a few seconds behind the pushed items, so once the run
has finished the rows past it are read a page at a time until none are left. On a client from
`LastRun(...)`, the iterator sticks to the run its first request resolves to.

`RunIterateDatasetItemsOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `Items` | `DatasetListItemsOptions` | Filters and projections applied to every page (see [storages.md](storages.md)). `Items.Offset` is the number of rows to skip (default 0) and `Items.Limit` the maximum number of dataset rows to scan (default no limit): fewer items are yielded when filters drop some, more when `Unwind` splits a row into several. `Items.Desc` is not supported, because the rows are read by offset from the oldest; it makes `Next` fail. |
| `ChunkSize` | `*int64` | Maximum number of dataset rows requested per API call. Default 1000. |
| `PollIntervalSecs` | `*int64` | How long to wait for the run to finish between polls, in seconds. Default 5. |

```go
it := client.Run(runID).IterateDatasetItems(apify.RunIterateDatasetItemsOptions{})
for {
	item, err := it.Next(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if item == nil {
		break // every item has been read
	}
	fmt.Println(string(*item))
}
```
