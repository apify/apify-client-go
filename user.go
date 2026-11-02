package apify

import (
	"context"
	"encoding/json"
	"net/http"
)

// UserClient is a client for accessing user data (/v2/users/{userId} or /v2/users/me).
//
// It also exposes account usage and limits. For another user than the current one, the API only
// serves them for your own account (by ID or username) or for the organization your API token
// belongs to, and rejects tokens with limited permissions.
type UserClient struct {
	ctx *resourceContext
}

func newUserClient(hc *httpClient, baseURL, id string) *UserClient {
	return &UserClient{ctx: newSingleContext(hc, baseURL, "users", id)}
}

// Get fetches the user. For "me" it returns private account details; for other users it
// returns the public profile. The bool reports whether the user exists.
func (c *UserClient) Get(ctx context.Context) (User, bool, error) {
	return getResource[User](ctx, c.ctx, "", NewQueryParams())
}

// MonthlyUsage fetches the account's monthly usage for the current month.
//
// It returns the raw JSON usage report from the API (a JSON object with the account's usage
// breakdown and totals for the period).
func (c *UserClient) MonthlyUsage(ctx context.Context) (json.RawMessage, error) {
	return c.MonthlyUsageWithOptions(ctx, MonthlyUsageOptions{})
}

// MonthlyUsageOptions configures [UserClient.MonthlyUsageWithOptions].
type MonthlyUsageOptions struct {
	// Date selects the billing period that contains this date, in YYYY-MM-DD format. If empty,
	// the current period is used.
	Date string
}

// MonthlyUsageWithOptions is like [UserClient.MonthlyUsage] but can select another period
// with [MonthlyUsageOptions.Date].
func (c *UserClient) MonthlyUsageWithOptions(ctx context.Context, options MonthlyUsageOptions) (json.RawMessage, error) {
	params := NewQueryParams()
	if options.Date != "" {
		params.AddString("date", &options.Date)
	}
	return getResourceRequired[json.RawMessage](ctx, c.ctx, "usage/monthly", params, c.ctx.shortTimeout())
}

// Limits fetches the account's resource limits.
func (c *UserClient) Limits(ctx context.Context) (json.RawMessage, error) {
	return getResourceRequired[json.RawMessage](ctx, c.ctx, "limits", NewQueryParams(), c.ctx.shortTimeout())
}

// UpdateLimits updates the account's resource limits.
func (c *UserClient) UpdateLimits(ctx context.Context, newLimits any) error {
	if c.ctx.idErr != nil {
		return c.ctx.idErr
	}
	data, err := json.Marshal(newLimits)
	if err != nil {
		return err
	}
	url := c.ctx.subURL("limits")
	_, err = c.ctx.http.call(ctx, http.MethodPut, url, data, contentTypeJSON, c.ctx.shortTimeout())
	return err
}
