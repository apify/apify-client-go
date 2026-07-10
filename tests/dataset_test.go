package apify_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	apify "github.com/apify/apify-client-go"
)

func TestListDatasets(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Datasets().List(ctx, apify.StorageListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("Datasets().List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestGetDataset(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	ds, err := client.Datasets().GetOrCreate(ctx, uniqueName("ds-get"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.Dataset(ds.ID).Delete(ctx) }()

	got, ok, err := client.Dataset(ds.ID).Get(ctx)
	if err != nil || !ok || got.ID != ds.ID {
		t.Fatalf("get dataset: ok=%v err=%v", ok, err)
	}
}

func TestDatasetCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	ds, err := client.Datasets().GetOrCreate(ctx, uniqueName("ds-crud"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.Dataset(ds.ID).Delete(ctx) }()
	dataset := client.Dataset(ds.ID)

	if _, ok, err := dataset.Get(ctx); err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}

	items := []map[string]any{
		{"url": "https://a.com", "n": 1},
		{"url": "https://b.com", "n": 2},
		{"url": "https://c.com", "n": 3},
	}
	if err := dataset.PushItems(ctx, items); err != nil {
		t.Fatalf("push items: %v", err)
	}

	page, err := dataset.ListItems(ctx, apify.DatasetListItemsOptions{})
	if err != nil {
		t.Fatalf("list items: %v", err)
	}
	// The pushed items should be readable immediately. The X-Apify-Pagination-Total header
	// can lag behind under eventual consistency, so assert on the returned item count.
	if page.Count != 3 || len(page.Items) != 3 {
		t.Fatalf("expected 3 items, got count=%d len=%d total=%d", page.Count, len(page.Items), page.Total)
	}

	csv, err := dataset.DownloadItems(ctx, apify.FormatCSV, apify.DatasetDownloadOptions{Bom: ptr(true)})
	if err != nil {
		t.Fatalf("download items: %v", err)
	}
	if !bytes.Contains(csv, []byte("url")) {
		t.Fatalf("expected CSV to contain a header, got %q", string(csv))
	}

	url, err := dataset.CreateItemsPublicURL(ctx, apify.DatasetListItemsOptions{}, nil)
	if err != nil || url == "" {
		t.Fatalf("public url: %q err=%v", url, err)
	}

	if _, _, err := dataset.GetStatistics(ctx); err != nil {
		t.Fatalf("statistics: %v", err)
	}

	updated, err := dataset.Update(ctx, map[string]any{"name": uniqueName("ds-renamed")})
	if err != nil {
		t.Fatalf("update: %v", err)
	}
	if updated.Name == "" {
		t.Fatalf("expected renamed dataset, got %+v", updated)
	}

	// Sanity-check that items decode into typed values.
	var first struct {
		N int `json:"n"`
	}
	if err := json.Unmarshal(page.Items[0], &first); err != nil {
		t.Fatalf("decode item: %v", err)
	}
}

// TestDatasetPushCompressedBody verifies that the live API accepts a gzip-compressed request
// body (Content-Encoding: gzip) and stores it intact. The client automatically gzips request
// bodies of 1024 bytes or larger, so pushing an item well above that threshold exercises the
// compression path end-to-end and confirms the round-trip is lossless.
func TestDatasetPushCompressedBody(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	ds, err := client.Datasets().GetOrCreate(ctx, uniqueName("ds-gzip"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.Dataset(ds.ID).Delete(ctx) }()
	dataset := client.Dataset(ds.ID)

	// A payload comfortably above the 1024-byte compression threshold. A distinctive marker
	// plus a long repeated body lets us assert the stored value survived compression intact.
	marker := "gzip-marker-" + uniqueName("v")
	blob := marker + strings.Repeat("x", 4096)
	if err := dataset.PushItems(ctx, map[string]any{"blob": blob}); err != nil {
		t.Fatalf("push large (gzipped) item: %v", err)
	}

	page, err := dataset.ListItems(ctx, apify.DatasetListItemsOptions{})
	if err != nil {
		t.Fatalf("list items: %v", err)
	}
	if page.Count != 1 || len(page.Items) != 1 {
		t.Fatalf("expected 1 item, got count=%d len=%d", page.Count, len(page.Items))
	}
	var got struct {
		Blob string `json:"blob"`
	}
	if err := json.Unmarshal(page.Items[0], &got); err != nil {
		t.Fatalf("decode item: %v", err)
	}
	if got.Blob != blob {
		t.Fatalf("stored blob does not round-trip: got %d bytes (prefix %q), want %d bytes",
			len(got.Blob), got.Blob[:min(len(got.Blob), len(marker))], len(blob))
	}
}
