package apify

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// slowBackend answers every request with a 200 and body after delay, unless the request's
// context ends first (then it fails with the context's error, as a timed-out transport does). It
// counts the requests it received and keeps the last URL.
type slowBackend struct {
	delay time.Duration
	body  string
	calls atomic.Int64
	url   atomic.Value
}

func (b *slowBackend) Do(req *http.Request) (*http.Response, error) {
	b.calls.Add(1)
	b.url.Store(req.URL.String())
	select {
	case <-time.After(b.delay):
	case <-req.Context().Done():
		return nil, req.Context().Err()
	}
	return &http.Response{StatusCode: http.StatusOK, Header: http.Header{}, Body: io.NopCloser(bytes.NewReader([]byte(b.body)))}, nil
}

const (
	// slowResponseDelay is how long slowBackend takes to answer: longer than the tiny tier used in
	// the tests below, so a request bound by that tier times out, but short enough to keep them fast.
	slowResponseDelay = 400 * time.Millisecond
	// tinyTier is the timeout tier of the test clients below.
	tinyTier = 40 * time.Millisecond
	// retriesAllowed is the number of retries of the test clients below, so a retried request
	// shows up as extra calls.
	retriesAllowed = 3
)

// noLogRedirection turns the log redirection of Call off, so a test counts only the requests of
// the call itself.
var noLogRedirection = &StreamedLogOptions{Disabled: true}

func slowClient(backend HTTPBackend, options ...Option) *ApifyClient {
	base := []Option{
		WithToken("t"),
		WithHTTPBackend(backend),
		WithMaxRetries(retriesAllowed),
		WithMinDelayBetweenRetries(time.Millisecond),
		WithTimeoutShort(tinyTier),
		WithTimeoutMedium(tinyTier),
	}
	return NewClient(append(base, options...)...)
}

// A request that is not idempotent must not be retried after it timed out: it may already have
// taken effect on the server.
func TestNonIdempotentRequestsAreNotRetriedAfterATimeout(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		run  func(client *ApifyClient) error
	}{
		{"Actor Start", func(c *ApifyClient) error {
			_, err := c.Actor("a").Start(ctx, map[string]any{}, ActorStartOptions{})
			return err
		}},
		{"Actor StartRaw", func(c *ApifyClient) error {
			_, err := c.Actor("a").StartRaw(ctx, []byte("x"), ActorStartOptions{})
			return err
		}},
		{"Actor Build", func(c *ApifyClient) error {
			_, err := c.Actor("a").Build(ctx, "0.1", ActorBuildOptions{})
			return err
		}},
		{"Task Start", func(c *ApifyClient) error {
			_, err := c.Task("t").Start(ctx, nil, TaskStartOptions{})
			return err
		}},
		{"Dataset PushItems", func(c *ApifyClient) error {
			return c.Dataset("d").PushItems(ctx, map[string]any{"a": 1})
		}},
		{"KeyValueStore SetRecordRawWithOptions", func(c *ApifyClient) error {
			return c.KeyValueStore("s").SetRecordRawWithOptions(ctx, "k", []byte("v"), "text/plain", SetRecordOptions{DoNotRetryTimeouts: true})
		}},
		{"KeyValueStore SetRecordJSONWithOptions", func(c *ApifyClient) error {
			return c.KeyValueStore("s").SetRecordJSONWithOptions(ctx, "k", map[string]any{"a": 1}, SetRecordOptions{DoNotRetryTimeouts: true})
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			backend := &slowBackend{delay: slowResponseDelay, body: `{"data":{}}`}
			// The record upload and the dataset push run on tiers the tiny medium one covers; the
			// upload uses the long tier, so shrink it too.
			if err := c.run(slowClient(backend, WithTimeoutLong(tinyTier))); err == nil {
				t.Fatal("expected a timeout error")
			}
			if got := backend.calls.Load(); got != 1 {
				t.Fatalf("a timed-out request must not be retried, got %d attempts", got)
			}
		})
	}
}

// The reverse of the test above: an idempotent request is retried after a timeout, and so is a
// record upload that did not ask for DoNotRetryTimeouts.
func TestIdempotentRequestsAreRetriedAfterATimeout(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		run  func(client *ApifyClient) error
	}{
		{"Actor Get", func(c *ApifyClient) error {
			_, _, err := c.Actor("a").Get(ctx)
			return err
		}},
		{"KeyValueStore SetRecordRaw", func(c *ApifyClient) error {
			return c.KeyValueStore("s").SetRecordRaw(ctx, "k", []byte("v"), "text/plain")
		}},
		{"KeyValueStore SetRecordRawWithOptions", func(c *ApifyClient) error {
			return c.KeyValueStore("s").SetRecordRawWithOptions(ctx, "k", []byte("v"), "text/plain", SetRecordOptions{})
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			backend := &slowBackend{delay: slowResponseDelay, body: `{"data":{}}`}
			if err := c.run(slowClient(backend, WithTimeoutLong(tinyTier))); err == nil {
				t.Fatal("expected a timeout error")
			}
			if got := backend.calls.Load(); got != retriesAllowed+1 {
				t.Fatalf("expected %d attempts, got %d", retriesAllowed+1, got)
			}
		})
	}
}

// A non-idempotent request still retries failures that are not timeouts (here a 500).
func TestNonIdempotentRequestsStillRetryServerErrors(t *testing.T) {
	backend := &mockBackend{responses: constant(500, `{"error":{"type":"x","message":"m"}}`)}
	client := testClient(backend, 2)
	if err := client.Dataset("d").PushItems(context.Background(), map[string]any{"a": 1}); err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", backend.calls)
	}
}

// Call has no client-side timeout of its own: a start slower than the medium tier must neither be
// aborted nor repeated (which would start a second, billed run). Start itself is bound by it.
func TestCallStartsWithoutATimeout(t *testing.T) {
	run := `{"data":{"id":"run1","status":"SUCCEEDED"}}`
	t.Run("Actor Call", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		got, err := slowClient(backend).Actor("a").Call(context.Background(), map[string]any{}, ActorStartOptions{Log: noLogRedirection}, Ptr(int64(0)))
		if err != nil || got.ID != "run1" {
			t.Fatalf("Call: %+v, %v", got, err)
		}
		if n := backend.calls.Load(); n != 2 { // the start and one poll
			t.Fatalf("expected the start once and one poll, got %d requests", n)
		}
	})
	t.Run("Actor CallRaw", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		if _, err := slowClient(backend).Actor("a").CallRaw(context.Background(), []byte("x"), ActorStartOptions{Log: noLogRedirection}, Ptr(int64(0))); err != nil {
			t.Fatalf("CallRaw: %v", err)
		}
		if n := backend.calls.Load(); n != 2 {
			t.Fatalf("expected the start once and one poll, got %d requests", n)
		}
	})
	t.Run("Task Call", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		if _, err := slowClient(backend).Task("t").Call(context.Background(), nil, TaskStartOptions{}, Ptr(int64(0))); err != nil {
			t.Fatalf("Task Call: %v", err)
		}
		if n := backend.calls.Load(); n != 2 {
			t.Fatalf("expected the start once and one poll, got %d requests", n)
		}
	})
	t.Run("Start stays bound by the medium tier", func(t *testing.T) {
		backend := &slowBackend{delay: slowResponseDelay, body: run}
		if _, err := slowClient(backend).Actor("a").Start(context.Background(), nil, ActorStartOptions{}); err == nil {
			t.Fatal("expected Start to time out")
		}
	})
	t.Run("ctx still bounds Call", func(t *testing.T) {
		backend := &slowBackend{delay: time.Minute, body: run}
		ctx, cancel := context.WithTimeout(context.Background(), 3*tinyTier)
		defer cancel()
		started := time.Now()
		if _, err := slowClient(backend).Actor("a").Call(ctx, nil, ActorStartOptions{Log: noLogRedirection}, nil); err == nil {
			t.Fatal("expected an error once ctx is done")
		}
		if time.Since(started) > 5*time.Second {
			t.Fatal("Call did not stop at the ctx deadline")
		}
	})
}

// The calls the reference client puts on the long tier must not be cut off by a short tier: here
// the short tier (tinyTier) is shorter than the response time, and the calls still succeed.
func TestKeyValueStoreAndPublicURLCallsUseTheLongTier(t *testing.T) {
	ctx := context.Background()
	store := `{"data":{"id":"s","urlSigningSecretKey":"secret"}}`
	cases := []struct {
		name string
		run  func(client *ApifyClient) error
	}{
		{"KeyValueStore Update", func(c *ApifyClient) error {
			_, err := c.KeyValueStore("s").Update(ctx, map[string]any{"name": "n"})
			return err
		}},
		{"KeyValueStore RecordExists", func(c *ApifyClient) error {
			_, err := c.KeyValueStore("s").RecordExists(ctx, "k")
			return err
		}},
		{"KeyValueStore GetRecordPublicURL", func(c *ApifyClient) error {
			_, err := c.KeyValueStore("s").GetRecordPublicURL(ctx, "k")
			return err
		}},
		{"KeyValueStore CreateKeysPublicURL", func(c *ApifyClient) error {
			_, err := c.KeyValueStore("s").CreateKeysPublicURL(ctx, nil)
			return err
		}},
		{"Dataset CreateItemsPublicURL", func(c *ApifyClient) error {
			_, err := c.Dataset("d").CreateItemsPublicURL(ctx, DatasetListItemsOptions{}, nil)
			return err
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			backend := &slowBackend{delay: slowResponseDelay, body: store}
			if err := c.run(slowClient(backend)); err != nil {
				t.Fatalf("a long-tier call must outlive the %v short and medium tiers: %v", tinyTier, err)
			}
			if got := backend.calls.Load(); got != 1 {
				t.Fatalf("expected a single request, got %d", got)
			}
		})
	}
}

func TestSetRecordRawWithOptionsSendsTheContentEncoding(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	value := []byte(strings.Repeat("x", 2*minCompressRequestBytes))

	err := client.KeyValueStore("s").SetRecordRawWithOptions(context.Background(), "k", value, "text/plain", SetRecordOptions{ContentEncoding: "gzip"})
	if err != nil {
		t.Fatalf("SetRecordRawWithOptions: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got != "gzip" {
		t.Fatalf("expected Content-Encoding gzip, got %q", got)
	}
	if backend.lastBody != string(value) {
		t.Fatal("a pre-compressed value must be uploaded as given, never compressed again")
	}

	if err := client.KeyValueStore("s").SetRecordRawWithOptions(context.Background(), "k", value, "text/plain", SetRecordOptions{}); err != nil {
		t.Fatalf("SetRecordRawWithOptions: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got == "" {
		t.Fatal("without a ContentEncoding the client compresses a large value of its own accord")
	}
}

func TestPushItemsAcceptsOnlyObjectsAndArraysOfObjects(t *testing.T) {
	type item struct {
		A int `json:"a"`
	}
	valid := map[string]any{
		"map":                 map[string]any{"a": 1},
		"struct":              item{A: 1},
		"struct pointer":      &item{A: 1},
		"slice of structs":    []item{{A: 1}, {A: 2}},
		"slice of maps":       []map[string]any{{"a": 1}},
		"empty slice":         []item{},
		"slice of any":        []any{map[string]any{"a": 1}, item{A: 2}},
		"raw JSON object":     rawJSON(`{"a":1}`),
		"raw JSON array":      rawJSON(`[{"a":1},{"a":2}]`),
		"padded raw JSON obj": rawJSON(" \n{\"a\":1} "),
	}
	for name, items := range valid {
		t.Run("accepts "+name, func(t *testing.T) {
			backend := &mockBackend{responses: okData()}
			if err := testClient(backend, 0).Dataset("d").PushItems(context.Background(), items); err != nil {
				t.Fatalf("PushItems rejected a valid shape: %v", err)
			}
			if backend.calls != 1 {
				t.Fatalf("expected 1 request, got %d", backend.calls)
			}
		})
	}

	invalid := map[string]any{
		"string":                "[1,2]",
		"[]byte":                []byte(`{"a":1}`),
		"number":                1,
		"bool":                  true,
		"nil":                   nil,
		"slice of strings":      []string{"a", "b"},
		"slice of numbers":      []int{1, 2},
		"mixed slice":           []any{map[string]any{"a": 1}, "b"},
		"slice of slices":       [][]int{{1}},
		"raw JSON string":       rawJSON(`"a"`),
		"raw JSON null element": rawJSON(`[{"a":1},null]`),
	}
	for name, items := range invalid {
		t.Run("rejects "+name, func(t *testing.T) {
			backend := &mockBackend{responses: okData()}
			if err := testClient(backend, 0).Dataset("d").PushItems(context.Background(), items); err == nil {
				t.Fatal("PushItems accepted a shape the API rejects")
			}
			if backend.calls != 0 {
				t.Fatalf("nothing may be sent for an invalid shape, got %d requests", backend.calls)
			}
		})
	}
}

func TestPushItemsRejectsAnInvalidDatasetIDBeforeSending(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	if err := testClient(backend, 0).Dataset("").PushItems(context.Background(), map[string]any{"a": 1}); err == nil {
		t.Fatal("expected an error for an empty dataset ID")
	}
	if backend.calls != 0 {
		t.Fatalf("nothing may be sent for an invalid ID, got %d requests", backend.calls)
	}
}

func rawJSON(s string) json.RawMessage { return json.RawMessage(s) }
