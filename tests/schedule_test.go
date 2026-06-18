package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

func scheduleDef(name string) map[string]any {
	return map[string]any{
		"name":           name,
		"cronExpression": "0 0 * * *",
		"isEnabled":      false,
		"isExclusive":    true,
		"actions":        []any{},
	}
}

func TestListSchedules(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Schedules().List(ctx, apify.ListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestGetSchedule(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	sch, err := client.Schedules().Create(ctx, scheduleDef(uniqueName("sch-get")))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Schedule(sch.ID).Delete(ctx) }()

	got, ok, err := client.Schedule(sch.ID).Get(ctx)
	if err != nil || !ok || got.ID != sch.ID {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
}

func TestScheduleCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	sch, err := client.Schedules().Create(ctx, scheduleDef(uniqueName("sch-crud")))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Schedule(sch.ID).Delete(ctx) }()
	schedule := client.Schedule(sch.ID)

	if _, ok, err := schedule.Get(ctx); err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	updated, err := schedule.Update(ctx, map[string]any{"cronExpression": "0 12 * * *"})
	if err != nil {
		t.Fatalf("update: %v", err)
	}
	if updated.CronExpression != "0 12 * * *" {
		t.Fatalf("cron not updated: %+v", updated)
	}
	if _, _, err := schedule.GetLog(ctx); err != nil {
		t.Fatalf("get log: %v", err)
	}
}
