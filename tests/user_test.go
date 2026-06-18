package apify_test

import "testing"

func TestGetOwnAccount(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	user, ok, err := client.Me().Get(ctx)
	if err != nil {
		t.Fatalf("Me().Get: %v", err)
	}
	if !ok || user.ID == "" {
		t.Fatalf("expected a user with a non-empty ID, got ok=%v %+v", ok, user)
	}
}

func TestGetMonthlyUsage(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	usage, err := client.Me().MonthlyUsage(ctx)
	if err != nil {
		t.Fatalf("MonthlyUsage: %v", err)
	}
	if len(usage) == 0 {
		t.Fatal("expected a non-empty monthly usage object")
	}
}

func TestGetLimits(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	limits, err := client.Me().Limits(ctx)
	if err != nil {
		t.Fatalf("Limits: %v", err)
	}
	if len(limits) == 0 {
		t.Fatal("expected a non-empty limits object")
	}
}
