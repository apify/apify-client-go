# Tasks

A task is a pre-configured Actor run with stored input. Access the task collection with
`client.Tasks()` and a single task with `client.Task(id)`.

## Task collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions) (PaginationList[Task], error)` | List the account's tasks. |
| `Iterate(ListOptions, chunkSize *int64) *ListIterator[Task]` | Lazy iterator over matching tasks. `Limit` caps the total yielded; `chunkSize` is the page size. |
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
| `IsPublic` | `*bool` | Whether the task is published on its public landing page, derived from `PublicConfig.PublishedAt`; use `Publish`/`Unpublish` to change it. |
| `PublicConfig` | `*TaskPublicConfig` | Public-facing display configuration of the landing page, set once the task has been configured for publishing. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

### `TaskPublicConfig` fields

| Field | Type | Meaning |
|---|---|---|
| `PublishedAt` | `*time.Time` | When the task was published, or `nil` if it is not published. |
| `SEOTitle` | `*string` | Title shown in search-engine results for the landing page. |
| `SEODescription` | `*string` | Description shown in search-engine results for the landing page. |
| `Categorization` | `*string` | Free-form category label for the landing page. |
| `InputSchemaFields` | `[]string` | Input schema field names highlighted on the landing page. |
| `DatasetName` | `*string` | Display name for the task's default dataset on the landing page. |
| `DatasetView` | `*string` | Name of the dataset view shown on the landing page. |

## Single task

| Method | Description |
| --- | --- |
| `Get(ctx) (Task, bool, error)` | Fetch the task. |
| `Update(ctx, newFields any) (Task, error)` | Update the task. |
| `Delete(ctx) error` | Delete the task. |
| `Publish(ctx) (Task, error)` | Publish the task on its public landing page. |
| `Unpublish(ctx) (Task, error)` | Unpublish the task from its public landing page. |
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

`Publish`/`Unpublish` toggle the task's public landing page by updating `IsPublic`; both reuse
the same `PUT /actor-tasks/{id}` endpoint as `Update`. `IsPublic` is a `*bool` (nil-checked
before use below) so a missing field can be distinguished from `false`.

```go
published, err := client.Task("my-task-id").Publish(ctx)
if err != nil {
	log.Fatal(err)
}
if published.IsPublic != nil {
	fmt.Println(*published.IsPublic) // true
}

unpublished, err := client.Task("my-task-id").Unpublish(ctx)
if err != nil {
	log.Fatal(err)
}
if unpublished.IsPublic != nil {
	fmt.Println(*unpublished.IsPublic) // false
}
```
