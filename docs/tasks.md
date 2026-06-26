# Tasks

A task is a pre-configured Actor run with stored input. Access the task collection with
`client.Tasks()` and a single task with `client.Task(id)`.

## Task collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions) (PaginationList[Task], error)` | List the account's tasks. |
| `Create(ctx, definition any) (Task, error)` | Create a new task. |

### `Task` fields

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique task ID. |
| `ActID` | `string` | ID of the Actor this task runs. |
| `UserID` | `string` | ID of the user who owns the task. |
| `Name` | `string` | Technical name of the task. |
| `Title` | `string` | Human-readable title shown in the UI. |
| `CreatedAt` | `*time.Time` | When the task was created. |
| `ModifiedAt` | `*time.Time` | When the task was last modified. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

## Single task

| Method | Description |
| --- | --- |
| `Get(ctx) (Task, bool, error)` | Fetch the task. |
| `Update(ctx, newFields any) (Task, error)` | Update the task. |
| `Delete(ctx) error` | Delete the task. |
| `Start(ctx, input any, TaskStartOptions) (ActorRun, error)` | Start a run (input overrides stored input). |
| `Call(ctx, input any, TaskStartOptions, waitSecs *int64) (ActorRun, error)` | Start and wait. |
| `GetInput(ctx) (json.RawMessage, bool, error)` | Fetch the stored input. |
| `UpdateInput(ctx, input any) (json.RawMessage, error)` | Replace the stored input. |
| `LastRun(status string) *RunClient` | Client for the last run (optional status filter). |
| `LastRunWithOptions(options LastRunOptions) *RunClient` | Client for the last run, filtered by status and/or origin. |
| `Runs() *RunCollectionClient` | This task's runs. |
| `Webhooks() *WebhookCollectionClient` | This task's webhooks. |

```go
task, err := client.Tasks().Create(ctx, map[string]any{
	"actId":   "apify/hello-world",
	"name":    "my-task",
	"options": map[string]any{"build": "latest", "memoryMbytes": 256, "timeoutSecs": 60},
	"input":   map[string]any{"message": "hi"},
})
if err != nil {
	log.Fatal(err)
}
run, err := client.Task(task.ID).Call(ctx, nil, apify.TaskStartOptions{}, apify.Ptr(int64(120)))
```

`TaskStartOptions` mirrors `ActorStartOptions` (see [actors.md](actors.md)) but omits the
Actor-only `ContentType` and `ForcePermissionLevel`, which the task run endpoint does not
accept.

`LastRun(status)` / `LastRunWithOptions(LastRunOptions{Status, Origin})` resolve the task's most
recent run, optionally narrowed by status and/or origin. `LastRunOptions` is the same type used
by the Actor client — see [actors.md](actors.md#single-actor) for its field reference.

```go
// Most recent task run that SUCCEEDED.
lastRun, ok, err := client.Task("my-task-id").
	LastRunWithOptions(apify.LastRunOptions{Status: "SUCCEEDED"}).
	Get(ctx)
if err != nil {
	log.Fatal(err)
}
if ok {
	fmt.Printf("last run: %s (%s)\n", lastRun.ID, lastRun.Status)
}
```
