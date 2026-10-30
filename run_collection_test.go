package apify

import (
	"context"
	"net/url"
	"strings"
	"testing"
)

func TestRunListsSendTheDeploymentFilter(t *testing.T) {
	ctx := context.Background()
	filter := RunListOptions{Status: []string{"RUNNING"}, Deployment: []string{"SINGLE_TENANT", "MULTI_TENANT_FREE_USERS"}}

	for name, list := range map[string]func(client *ApifyClient) error{
		"account runs": func(c *ApifyClient) error { _, err := c.Runs().List(ctx, ListOptions{}, filter); return err },
		"actor runs":   func(c *ApifyClient) error { _, err := c.Actor("a").Runs().List(ctx, ListOptions{}, filter); return err },
	} {
		backend := &mockBackend{responses: okData()}
		if err := list(testClient(backend, 0)); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		query, _ := url.ParseQuery(strings.SplitN(backend.lastURL, "?", 2)[1])
		if got := query.Get("deployment"); got != "SINGLE_TENANT,MULTI_TENANT_FREE_USERS" {
			t.Errorf("%s: deployment = %q", name, got)
		}
	}

	backend := &mockBackend{responses: okData()}
	if _, err := testClient(backend, 0).Runs().List(ctx, ListOptions{}, RunListOptions{}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(backend.lastURL, "deployment") {
		t.Errorf("deployment must be omitted when unset, got %q", backend.lastURL)
	}
}
