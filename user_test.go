package apify

import (
	"context"
	"strings"
	"testing"
)

// MonthlyUsageWithOptions sends the date query parameter only when it is set.
func TestMonthlyUsageWithOptionsDate(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	if _, err := client.Me().MonthlyUsageWithOptions(context.Background(), MonthlyUsageOptions{Date: "2026-09-14"}); err != nil {
		t.Fatalf("monthly usage: %v", err)
	}
	if !strings.Contains(backend.lastURL, "/users/me/usage/monthly") || !strings.Contains(backend.lastURL, "date=2026-09-14") {
		t.Fatalf("expected date query parameter, got %q", backend.lastURL)
	}
	if _, err := client.Me().MonthlyUsage(context.Background()); err != nil {
		t.Fatalf("monthly usage: %v", err)
	}
	if strings.Contains(backend.lastURL, "date=") {
		t.Fatalf("date must be omitted when empty, got %q", backend.lastURL)
	}
}

// The usage and limits endpoints are documented for a user by ID or username as well, so a
// client from User(id) calls them instead of failing.
func TestUserLimitsAndUsageByID(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	ctx := context.Background()

	if _, err := client.User("someone").MonthlyUsageWithOptions(ctx, MonthlyUsageOptions{Date: "2026-09-14"}); err != nil {
		t.Fatalf("monthly usage by ID: %v", err)
	}
	if !strings.Contains(backend.lastURL, "/users/someone/usage/monthly") || !strings.Contains(backend.lastURL, "date=2026-09-14") {
		t.Fatalf("unexpected monthly usage URL %q", backend.lastURL)
	}
	if _, err := client.User("someone").Limits(ctx); err != nil || !strings.HasSuffix(backend.lastURL, "/users/someone/limits") {
		t.Fatalf("limits by ID: %v, URL %q", err, backend.lastURL)
	}
	if err := client.User("someone").UpdateLimits(ctx, map[string]any{"maxMonthlyUsageUsd": 5}); err != nil || !strings.HasSuffix(backend.lastURL, "/users/someone/limits") {
		t.Fatalf("update limits by ID: %v, URL %q", err, backend.lastURL)
	}
}
