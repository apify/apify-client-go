package apify

import (
	"context"
	"strings"
	"testing"
)

// The storages and the log of a run reached through LastRun carry its status and origin filters,
// so they address the storages and the log of the same run, not those of the newest run of any
// status.
func TestStoragesAndLogOfAFilteredLastRunCarryItsFilters(t *testing.T) {
	ctx := context.Background()
	lastRun := func(client *ApifyClient) *RunClient {
		return client.Actor("a").LastRunWithOptions(LastRunOptions{Status: "SUCCEEDED", Origin: "API"})
	}
	taskLastRun := func(client *ApifyClient) *RunClient {
		return client.Task("t").LastRunWithOptions(LastRunOptions{Status: "SUCCEEDED", Origin: "API"})
	}

	cases := map[string]func(client *ApifyClient) error{
		"dataset get": func(c *ApifyClient) error { _, _, err := lastRun(c).Dataset().Get(ctx); return err },
		"dataset items": func(c *ApifyClient) error {
			_, err := lastRun(c).Dataset().ListItems(ctx, DatasetListItemsOptions{})
			return err
		},
		"dataset download": func(c *ApifyClient) error {
			_, err := lastRun(c).Dataset().DownloadItems(ctx, "json", DatasetDownloadOptions{})
			return err
		},
		"dataset push": func(c *ApifyClient) error {
			return lastRun(c).Dataset().PushItems(ctx, map[string]any{"a": 1})
		},
		"dataset public URL": func(c *ApifyClient) error {
			_, err := lastRun(c).Dataset().CreateItemsPublicURL(ctx, DatasetListItemsOptions{}, nil)
			return err
		},
		"store get": func(c *ApifyClient) error { _, _, err := lastRun(c).KeyValueStore().Get(ctx); return err },
		"store keys": func(c *ApifyClient) error {
			_, err := lastRun(c).KeyValueStore().ListKeys(ctx, ListKeysOptions{})
			return err
		},
		"store record": func(c *ApifyClient) error {
			_, _, err := lastRun(c).KeyValueStore().GetRecord(ctx, "OUTPUT")
			return err
		},
		"store record public URL": func(c *ApifyClient) error {
			_, err := lastRun(c).KeyValueStore().GetRecordPublicURL(ctx, "OUTPUT")
			return err
		},
		"store keys public URL": func(c *ApifyClient) error {
			_, err := lastRun(c).KeyValueStore().CreateKeysPublicURL(ctx, nil)
			return err
		},
		"queue get":  func(c *ApifyClient) error { _, _, err := lastRun(c).RequestQueue().Get(ctx); return err },
		"queue head": func(c *ApifyClient) error { _, err := lastRun(c).RequestQueue().ListHead(ctx, nil); return err },
		"log get":    func(c *ApifyClient) error { _, _, err := lastRun(c).Log().Get(ctx); return err },
		"log stream": func(c *ApifyClient) error {
			stream, err := lastRun(c).GetStreamedLog(ctx)
			if err == nil {
				_ = stream.Close()
			}
			return err
		},
		"task last run dataset": func(c *ApifyClient) error { _, _, err := taskLastRun(c).Dataset().Get(ctx); return err },
		"task last run log":     func(c *ApifyClient) error { _, _, err := taskLastRun(c).Log().Get(ctx); return err },
	}
	for name, call := range cases {
		t.Run(name, func(t *testing.T) {
			body := `{"data":{"id":"x","urlSigningSecretKey":"s"}}`
			if name == "dataset items" {
				body = `[]` // the items endpoint answers with a bare array
			}
			backend := &mockBackend{responses: []mockResponse{{status: 200, body: body}}}
			if err := call(testClient(backend, 0)); err != nil {
				t.Fatalf("call failed: %v", err)
			}
			if !strings.Contains(backend.lastURL, "/runs/last/") {
				t.Fatalf("expected a runs/last URL, got %q", backend.lastURL)
			}
			for _, want := range []string{"status=SUCCEEDED", "origin=API"} {
				if !strings.Contains(backend.lastURL, want) {
					t.Errorf("%s is missing from %q", want, backend.lastURL)
				}
			}
		})
	}
}

// A storage reached through a run that has no filters does not gain any.
func TestStoragesOfAnUnfilteredRunHaveNoFilterParams(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	if _, _, err := testClient(backend, 0).Run("r").Dataset().Get(context.Background()); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(backend.lastURL, "?") {
		t.Errorf("unexpected query string in %q", backend.lastURL)
	}
}
