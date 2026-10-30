package apify

import (
	"bytes"
	"context"
	"io"
	"net/http"
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
