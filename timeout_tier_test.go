package apify

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// The calls on the long tier are not cut off by a shorter short tier: here the short tier
// (tinyTier) is shorter than the response time, and the calls still succeed.
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

// extendForServerWait adds the server-side wait to the base timeout, capped at
// maxWaitForFinishHoldSecs.
func TestExtendForServerWait(t *testing.T) {
	cases := []struct {
		name             string
		base             time.Duration
		waitForFinishSec *int64
		want             time.Duration
	}{
		{"nil wait leaves base unchanged", 5 * time.Second, nil, 5 * time.Second},
		{"zero wait leaves base unchanged", 5 * time.Second, Ptr(int64(0)), 5 * time.Second},
		{"negative wait leaves base unchanged", 5 * time.Second, Ptr(int64(-1)), 5 * time.Second},
		{"wait under the cap adds its full duration", 5 * time.Second, Ptr(int64(30)), 35 * time.Second},
		{"wait over the cap is capped at 60s", 5 * time.Second, Ptr(int64(999)), 65 * time.Second},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := extendForServerWait(c.base, c.waitForFinishSec); got != c.want {
				t.Errorf("extendForServerWait(%v, %v) = %v, want %v", c.base, c.waitForFinishSec, got, c.want)
			}
		})
	}
}

// blockingBackend blocks until the request's context is done and returns its error, which
// simulates an unresponsive server.
type blockingBackend struct{}

func (blockingBackend) Do(req *http.Request) (*http.Response, error) {
	<-req.Context().Done()
	return nil, req.Context().Err()
}

// A metadata call (Get, which uses the "short" tier) must be aborted by the client once the
// short tier elapses, even though the caller's own context has no deadline of its own.
func TestShortTierTimesOutMetadataCall(t *testing.T) {
	client := NewClient(
		WithToken("t"),
		WithHTTPBackend(blockingBackend{}),
		WithMaxRetries(0),
		WithTimeoutShort(30*time.Millisecond),
	)

	start := time.Now()
	_, _, err := client.Actor("some-actor").Get(context.Background())
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected a timeout error from a backend that never responds")
	}
	if elapsed > 2*time.Second {
		t.Fatalf("expected the short tier (30ms) to abort the request quickly, took %v", elapsed)
	}
}
