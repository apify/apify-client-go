# Schedules

Schedules start Actor or task runs at specified times. Access the schedule collection with
`client.Schedules()` and a single schedule with `client.Schedule(id)`.

## Schedule collection

| Method | Description |
| --- | --- |
| `List(ctx, ListOptions) (PaginationList[Schedule], error)` | List the account's schedules. |
| `Create(ctx, definition any) (Schedule, error)` | Create a new schedule. |

## Single schedule

| Method | Description |
| --- | --- |
| `Get(ctx) (Schedule, bool, error)` | Fetch the schedule. |
| `Update(ctx, newFields any) (Schedule, error)` | Update the schedule. |
| `Delete(ctx) error` | Delete the schedule. |
| `GetLog(ctx) (string, bool, error)` | The schedule's invocation log. |

```go
sch, err := client.Schedules().Create(ctx, map[string]any{
	"name":           "nightly",
	"cronExpression": "0 0 * * *",
	"isEnabled":      true,
	"isExclusive":    true,
	"actions":        []any{},
})
if err != nil {
	log.Fatal(err)
}
_, err = client.Schedule(sch.ID).Update(ctx, map[string]any{"cronExpression": "0 12 * * *"})
```
