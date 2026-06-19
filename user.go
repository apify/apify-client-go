package apify

import (
	"context"
	"encoding/json"
	"errors"
)

// UserClient is a client for accessing user data (/v2/users/{userId} or /v2/users/me).
//
// For the current user ("me"), it also exposes account usage and limits. Those endpoints
// only exist for "me" and return an error if called on another user's client.
type UserClient struct {
	ctx  *resourceContext
	isMe bool
}

func newUserClient(hc *httpClient, baseURL, id string) *UserClient {
	return &UserClient{
		ctx:  newSingleContext(hc, baseURL, "users", id),
		isMe: id == meUserPlaceholder,
	}
}

// errNotMe is returned by the me-only methods when called on another user's client.
var errNotMe = errors.New("this operation is only available for the current user (use Me())")

// Get fetches the user. For "me" it returns private account details; for other users it
// returns the public profile. The bool reports whether the user exists.
func (c *UserClient) Get(ctx context.Context) (User, bool, error) {
	return getResource[User](ctx, c.ctx, "", NewQueryParams())
}

// MonthlyUsage fetches the current account's monthly usage for the current month. Only
// available for "me". To select a different month, use [UserClient.MonthlyUsageForDate].
//
// It returns the raw JSON usage report from the API (a JSON object with the account's usage
// breakdown and totals for the period).
func (c *UserClient) MonthlyUsage(ctx context.Context) (json.RawMessage, error) {
	return c.MonthlyUsageForDate(ctx, "")
}

// MonthlyUsageForDate fetches the current account's monthly usage for the month containing
// the given date (formatted as YYYY-MM-DD). An empty date omits the parameter and returns
// usage for the current month. Only available for "me".
//
// It returns the raw JSON usage report from the API (a JSON object with the account's usage
// breakdown and totals for the period).
func (c *UserClient) MonthlyUsageForDate(ctx context.Context, date string) (json.RawMessage, error) {
	if !c.isMe {
		return nil, errNotMe
	}
	params := NewQueryParams()
	if date != "" {
		params.AddString("date", &date)
	}
	return getResourceRequired[json.RawMessage](ctx, c.ctx, "usage/monthly", params)
}

// Limits fetches the current account's resource limits. Only available for "me".
func (c *UserClient) Limits(ctx context.Context) (json.RawMessage, error) {
	if !c.isMe {
		return nil, errNotMe
	}
	return getResourceRequired[json.RawMessage](ctx, c.ctx, "limits", NewQueryParams())
}

// UpdateLimits updates the current account's resource limits. Only available for "me".
func (c *UserClient) UpdateLimits(ctx context.Context, newLimits any) error {
	if !c.isMe {
		return errNotMe
	}
	data, err := json.Marshal(newLimits)
	if err != nil {
		return err
	}
	url := c.ctx.subURL("limits")
	_, err = c.ctx.http.call(ctx, "PUT", url, data, contentTypeJSON, defaultRequestTimeout)
	return err
}
