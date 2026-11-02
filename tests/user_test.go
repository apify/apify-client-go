package apify_test

import (
	"strings"
	"testing"
)

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

// TestGetUserByID covers the public-user getter client.User(id).Get() — GET /v2/users/{userId}
// — which is distinct from the client.Me() ("users/me") path. It resolves the test account's
// own userId via Me() first, then fetches it through the by-ID client. Both calls are
// read-only, so the test stays safe to run concurrently on the shared test account.
func TestGetUserByID(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	me, ok, err := client.Me().Get(ctx)
	if err != nil || !ok || me.ID == "" {
		t.Fatalf("Me().Get: ok=%v err=%v id=%q", ok, err, me.ID)
	}

	user, ok, err := client.User(me.ID).Get(ctx)
	if err != nil {
		t.Fatalf("User(%q).Get: %v", me.ID, err)
	}
	if !ok || user.ID != me.ID {
		t.Fatalf("expected the public profile of %q, got ok=%v id=%q", me.ID, ok, user.ID)
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

// TestGetLimitsAndUsageByUsername reads the limits and the monthly usage through the by-username
// client: the API serves them for your own account by ID or username as well as through "me".
func TestGetLimitsAndUsageByUsername(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	me, ok, err := client.Me().Get(ctx)
	if err != nil || !ok || me.Username == "" {
		t.Fatalf("Me().Get: ok=%v err=%v username=%q", ok, err, me.Username)
	}
	user := client.User(me.Username)
	limits, err := user.Limits(ctx)
	if err != nil {
		t.Fatalf("User(%q).Limits: %v", me.Username, err)
	}
	if len(limits) == 0 {
		t.Fatal("expected a non-empty limits object")
	}
	usage, err := user.MonthlyUsage(ctx)
	if err != nil {
		t.Fatalf("User(%q).MonthlyUsage: %v", me.Username, err)
	}
	if len(usage) == 0 {
		t.Fatal("expected a non-empty monthly usage object")
	}
}

// TestUserProfileURLsAreNormalized checks that the profile of the own account is parsed and its
// URL fields, when set, are absolute URLs.
func TestUserProfileURLsAreNormalized(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	me, ok, err := client.Me().Get(ctx)
	if err != nil || !ok {
		t.Fatalf("Me().Get: ok=%v err=%v", ok, err)
	}
	if me.Profile == nil {
		t.Fatal("expected the profile of the own account")
	}
	for name, value := range map[string]*string{"pictureUrl": me.Profile.PictureURL, "websiteUrl": me.Profile.WebsiteURL} {
		if value != nil && !strings.Contains(*value, "://") {
			t.Errorf("%s is not an absolute URL: %q", name, *value)
		}
	}
}
