package apify

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

// mockBackend is a deterministic HTTPBackend for offline unit tests. It serves a queue of
// scripted responses/errors and records how many times it was called.
type mockBackend struct {
	mu          sync.Mutex
	responses   []mockResponse
	calls       int
	lastHeaders http.Header
	lastURL     string
	lastBody    string
	bodies      []string
}

type mockResponse struct {
	status int
	body   string
	err    error
}

func (m *mockBackend) Do(req *http.Request) (*http.Response, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	idx := m.calls
	m.calls++
	m.lastHeaders = req.Header.Clone()
	m.lastURL = req.URL.String()
	if req.Body != nil {
		data, _ := io.ReadAll(req.Body)
		m.lastBody = string(data)
		m.bodies = append(m.bodies, string(data))
	}
	// If we run past the script, repeat the last entry (so "constant" behaviour is easy).
	if idx >= len(m.responses) {
		idx = len(m.responses) - 1
	}
	r := m.responses[idx]
	if r.err != nil {
		return nil, r.err
	}
	return &http.Response{
		StatusCode: r.status,
		Header:     http.Header{},
		Body:       io.NopCloser(bytes.NewReader([]byte(r.body))),
	}, nil
}

func constant(status int, body string) []mockResponse {
	return []mockResponse{{status: status, body: body}}
}

// testClient builds a client wired to the given backend with a tiny retry delay so tests
// are fast.
func testClient(backend HTTPBackend, maxRetries int) *ApifyClient {
	return NewClient(
		"test-token",
		WithHTTPBackend(backend),
		WithMaxRetries(maxRetries),
		WithMinDelayBetweenRetries(time.Millisecond),
	)
}

func TestSuccessSingleCall(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"u1","username":"bob"}}`)}
	client := testClient(backend, 8)

	user, ok, err := client.Me().Get(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok || user.ID != "u1" || user.Username != "bob" {
		t.Fatalf("unexpected user: %+v ok=%v", user, ok)
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 call, got %d", backend.calls)
	}
}

// TestBearerTokenHeaderPresentAndAbsentOnEmptyToken locks in the constructor's token
// handling: a non-empty token is sent as a Bearer Authorization header, and NewClient("")
// yields an unauthenticated client that sends no Authorization header (valid for public
// endpoints, matching the JS reference).
func TestBearerTokenHeaderPresentAndAbsentOnEmptyToken(t *testing.T) {
	t.Run("token set", func(t *testing.T) {
		backend := &mockBackend{responses: constant(200, `{"data":{"id":"u1"}}`)}
		client := NewClient("secret-token", WithHTTPBackend(backend), WithMinDelayBetweenRetries(time.Millisecond))
		if _, _, err := client.Me().Get(context.Background()); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := backend.lastHeaders.Get("Authorization"); got != "Bearer secret-token" {
			t.Fatalf("expected Authorization header %q, got %q", "Bearer secret-token", got)
		}
	})

	t.Run("empty token", func(t *testing.T) {
		backend := &mockBackend{responses: constant(200, `{"data":{"id":"u1"}}`)}
		client := NewClient("", WithHTTPBackend(backend), WithMinDelayBetweenRetries(time.Millisecond))
		if _, _, err := client.Me().Get(context.Background()); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := backend.lastHeaders.Get("Authorization"); got != "" {
			t.Fatalf("expected no Authorization header for empty token, got %q", got)
		}
	})
}

func TestRateLimitIsRetried(t *testing.T) {
	backend := &mockBackend{responses: constant(429, `{"error":{"type":"rate-limit-exceeded","message":"slow down"}}`)}
	client := testClient(backend, 2)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	apiErr, ok := AsAPIError(err)
	if !ok || apiErr.StatusCode != 429 {
		t.Fatalf("expected 429 APIError, got %v", err)
	}
	// 1 initial + 2 retries = 3 attempts.
	if backend.calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", backend.calls)
	}
	if apiErr.Attempt != 3 {
		t.Fatalf("expected attempt 3, got %d", apiErr.Attempt)
	}
}

func TestServerErrorIsRetried(t *testing.T) {
	backend := &mockBackend{responses: constant(503, `{"error":{"type":"internal","message":"boom"}}`)}
	client := testClient(backend, 1)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 2 {
		t.Fatalf("expected 2 attempts, got %d", backend.calls)
	}
}

func TestClientErrorNotRetried(t *testing.T) {
	backend := &mockBackend{responses: constant(400, `{"error":{"type":"bad-request","message":"nope"}}`)}
	client := testClient(backend, 5)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt for 4xx, got %d", backend.calls)
	}
}

func TestNetworkErrorIsRetried(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{{err: errors.New("connection refused")}}}
	client := testClient(backend, 3)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 4 {
		t.Fatalf("expected 4 attempts, got %d", backend.calls)
	}
}

func TestRetryThenSuccess(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{
		{status: 500, body: `{"error":{"type":"internal","message":"x"}}`},
		{status: 500, body: `{"error":{"type":"internal","message":"x"}}`},
		{status: 200, body: `{"data":{"id":"ok"}}`},
	}}
	client := testClient(backend, 5)

	user, ok, err := client.Me().Get(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok || user.ID != "ok" {
		t.Fatalf("unexpected user: %+v", user)
	}
	if backend.calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", backend.calls)
	}
}

func TestNotFoundMapsToNone(t *testing.T) {
	backend := &mockBackend{responses: constant(404, `{"error":{"type":"record-not-found","message":"missing"}}`)}
	client := testClient(backend, 5)

	_, ok, err := client.Actor("nope").Get(context.Background())
	if err != nil {
		t.Fatalf("expected nil error for not-found, got %v", err)
	}
	if ok {
		t.Fatal("expected ok=false for missing resource")
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt (no retry on 404), got %d", backend.calls)
	}
}

func TestErrorBodyIsParsed(t *testing.T) {
	backend := &mockBackend{responses: constant(400, `{"error":{"type":"bad-request","message":"invalid input","data":{"field":"name"}}}`)}
	client := testClient(backend, 0)

	_, _, err := client.Me().Get(context.Background())
	apiErr, ok := AsAPIError(err)
	if !ok {
		t.Fatalf("expected APIError, got %v", err)
	}
	if apiErr.StatusCode != 400 || apiErr.Type != "bad-request" || apiErr.Message != "invalid input" {
		t.Fatalf("unexpected fields: %+v", apiErr)
	}
	if apiErr.HTTPMethod != http.MethodGet {
		t.Fatalf("expected GET, got %q", apiErr.HTTPMethod)
	}
	if apiErr.Path == "" {
		t.Fatal("expected a non-empty path")
	}
	if apiErr.Data["field"] == nil {
		t.Fatalf("expected error data to be parsed, got %+v", apiErr.Data)
	}
}

func TestZeroRetriesSingleAttempt(t *testing.T) {
	backend := &mockBackend{responses: constant(500, `{"error":{"type":"internal","message":"x"}}`)}
	client := testClient(backend, 0)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt with max_retries=0, got %d", backend.calls)
	}
}
