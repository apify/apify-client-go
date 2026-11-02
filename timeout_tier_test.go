package apify

import (
	"context"
	"testing"
)

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
