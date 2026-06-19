# Schedules

Schedules start Actor or task runs at specified times. Access the schedule collection with
`client.Schedules()` and a single schedule with `client.Schedule(id)`.

## Schedule collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions) (PaginationList[Schedule], error)` | List the account's schedules. |
| `Create(ctx, definition any) (Schedule, error)` | Create a new schedule. |

### `Schedule` fields

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique schedule ID. |
| `UserID` | `string` | ID of the user who owns the schedule. |
| `Name` | `string` | The schedule name. |
| `CronExpression` | `string` | Cron expression governing when the schedule fires. |
| `IsEnabled` | `bool` | Whether the schedule is currently active. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

## Single schedule

| Method | Description |
| --- | --- |
| `Get(ctx) (Schedule, bool, error)` | Fetch the schedule. |
| `Update(ctx, newFields any) (Schedule, error)` | Update the schedule. |
| `Delete(ctx) error` | Delete the schedule. |
| `GetLog(ctx) (string, bool, error)` | The schedule's invocation log. |

`Create`/`Update` take a free-form definition (`any`) that is serialized to JSON, so the
schedule's fields are passed as a map. The key fields are `name`, `cronExpression`, `isEnabled`,
`isExclusive` (whether overlapping invocations are skipped), and `actions` — the list of things
the schedule does when it fires.

Each entry in `actions` describes one action. The common type is `RUN_ACTOR` (or
`RUN_ACTOR_TASK`), which starts an Actor (or task) run:

| Action field | Type | Meaning |
|---|---|---|
| `type` | `string` | `"RUN_ACTOR"` or `"RUN_ACTOR_TASK"`. |
| `actorId` | `string` | ID of the Actor to run (for `RUN_ACTOR`). |
| `actorTaskId` | `string` | ID of the task to run (for `RUN_ACTOR_TASK`). |
| `runInput` | `object` | Optional run input (`body`, `contentType`). |
| `runOptions` | `object` | Optional run options (`build`, `memoryMbytes`, `timeoutSecs`). |

```go
sch, err := client.Schedules().Create(ctx, map[string]any{
	"name":           "nightly",
	"cronExpression": "0 0 * * *",
	"isEnabled":      true,
	"isExclusive":    true,
	"actions": []any{
		map[string]any{
			"type":    "RUN_ACTOR",
			"actorId": "apify/hello-world",
			"runInput": map[string]any{
				"body":        `{"message":"hi"}`,
				"contentType": "application/json",
			},
			"runOptions": map[string]any{
				"build":        "latest",
				"memoryMbytes": 256,
				"timeoutSecs":  60,
			},
		},
	},
})
if err != nil {
	log.Fatal(err)
}
_, err = client.Schedule(sch.ID).Update(ctx, map[string]any{"cronExpression": "0 12 * * *"})
```
