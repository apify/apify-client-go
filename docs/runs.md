# Runs

A run is a single execution of an Actor. Access all runs with `client.Runs()`, a single run
with `client.Run(id)`, and an Actor's or task's runs with `client.Actor(id).Runs()` /
`client.Task(id).Runs()`.

## Run collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions, RunListOptions) (PaginationList[ActorRun], error)` | List runs. |

`RunListOptions`:

| Field | Type | Meaning |
|---|---|---|
| `Status` | `[]string` | Filter by one or more run statuses; sent as a comma-separated list. Values are the eight `ActorJobStatus` values: `READY`, `RUNNING`, `SUCCEEDED`, `FAILED`, `TIMING-OUT`, `TIMED-OUT`, `ABORTING`, `ABORTED`. |
| `StartedAfter` | `*string` | Only runs started after this ISO-8601 timestamp. |
| `StartedBefore` | `*string` | Only runs started before this ISO-8601 timestamp. |

The time filters apply only to Actor- and task-scoped collections.

## Single run

| Method | Description |
| --- | --- |
| `Get(ctx) (ActorRun, bool, error)` | Fetch the run. |
| `GetWithWait(ctx, waitForFinishSecs *int64) (ActorRun, bool, error)` | Fetch, optionally waiting server-side. |
| `Update(ctx, newFields any) (ActorRun, error)` | Update the run (e.g. status message). |
| `Delete(ctx) error` | Delete the run. |
| `Abort(ctx, gracefully *bool) (ActorRun, error)` | Abort the run. Pass `nil` to use the server default (immediate), or a pointer to `true`/`false` to abort gracefully/immediately. |
| `Metamorph(ctx, targetActorID string, input any, MetamorphOptions) (ActorRun, error)` | Metamorph into another Actor. |
| `Reboot(ctx) (ActorRun, error)` | Reboot the run. |
| `Resurrect(ctx, RunResurrectOptions) (ActorRun, error)` | Resurrect a finished run. |
| `Charge(ctx, RunChargeOptions) error` | Charge a pay-per-event run (sends an idempotency key). |
| `WaitForFinish(ctx, waitSecs *int64) (ActorRun, error)` | Poll until the run is terminal. |
| `Dataset() *DatasetClient` | The run's default dataset. |
| `KeyValueStore() *KeyValueStoreClient` | The run's default key-value store. |
| `RequestQueue() *RequestQueueClient` | The run's default request queue. |
| `Log() *LogClient` | The run's log. |
| `GetStreamedLog(ctx) (io.ReadCloser, error)` | Live stream of the run's raw log (log redirection). |

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
| `ContentType` | `string` | Content type of the input body (default `application/json`). |

`RunChargeOptions`: `EventName` (required), `Count`, `IdempotencyKey` (auto-generated if
empty, so a retried charge is applied at most once).

```go
run, err := client.Run(runID).WaitForFinish(ctx, nil) // nil waits indefinitely
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
| `ContainerURL` | `string` | URL of the run's container (for live access). |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API (forward compatibility). |

`ActorRun.IsTerminal()` reports whether a run has finished. Status message convenience:
`client.SetStatusMessage(ctx, message, isTerminal)` updates the current run identified by the
`ACTOR_RUN_ID` environment variable.
