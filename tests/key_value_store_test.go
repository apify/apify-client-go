package apify_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	apify "github.com/apify/apify-client-go"
)

func TestListKeyValueStores(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.KeyValueStores().List(ctx, apify.StorageListOptions{Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestGetKeyValueStore(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	store, err := client.KeyValueStores().GetOrCreate(ctx, uniqueName("kvs-get"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.KeyValueStore(store.ID).Delete(ctx) }()

	got, ok, err := client.KeyValueStore(store.ID).Get(ctx)
	if err != nil || !ok || got.ID != store.ID {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
}

// TestRecordKeyWithSpecialChars validates percent-encoding of URL-reserved characters in
// record-key path segments.
func TestRecordKeyWithSpecialChars(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	store, err := client.KeyValueStores().GetOrCreate(ctx, uniqueName("kvs-special"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.KeyValueStore(store.ID).Delete(ctx) }()
	kvs := client.KeyValueStore(store.ID)

	// Only characters allowed by the API, but including URL-reserved ones that must be
	// percent-encoded in the path segment.
	key := "weird-key!'()"
	if err := kvs.SetRecordJSON(ctx, key, map[string]any{"ok": true}); err != nil {
		t.Fatalf("set record: %v", err)
	}
	exists, err := kvs.RecordExists(ctx, key)
	if err != nil || !exists {
		t.Fatalf("record exists: exists=%v err=%v", exists, err)
	}
	rec, ok, err := kvs.GetRecord(ctx, key)
	if err != nil || !ok {
		t.Fatalf("get record: ok=%v err=%v", ok, err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(rec.Value, &decoded); err != nil {
		t.Fatalf("decode record: %v", err)
	}
	if err := kvs.DeleteRecord(ctx, key); err != nil {
		t.Fatalf("delete record: %v", err)
	}
}

func TestKeyValueStoreCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	store, err := client.KeyValueStores().GetOrCreate(ctx, uniqueName("kvs-crud"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.KeyValueStore(store.ID).Delete(ctx) }()
	kvs := client.KeyValueStore(store.ID)

	if _, ok, err := kvs.Get(ctx); err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	if err := kvs.SetRecordJSON(ctx, "OUTPUT", map[string]any{"hello": "world"}); err != nil {
		t.Fatalf("set record: %v", err)
	}
	exists, err := kvs.RecordExists(ctx, "OUTPUT")
	if err != nil || !exists {
		t.Fatalf("record exists: exists=%v err=%v", exists, err)
	}
	rec, ok, err := kvs.GetRecord(ctx, "OUTPUT")
	if err != nil || !ok {
		t.Fatalf("get record: ok=%v err=%v", ok, err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(rec.Value, &decoded); err != nil || decoded["hello"] != "world" {
		t.Fatalf("unexpected record: %v err=%v", decoded, err)
	}
	if _, _, err := kvs.GetRecordWithOptions(ctx, "OUTPUT", apify.GetRecordOptions{Attachment: ptr(false)}); err != nil {
		t.Fatalf("get record with options: %v", err)
	}
	keys, err := kvs.ListKeys(ctx, apify.ListKeysOptions{})
	if err != nil {
		t.Fatalf("list keys: %v", err)
	}
	if len(keys.Items) == 0 {
		t.Fatal("expected at least one key")
	}
	if _, err := kvs.Update(ctx, map[string]any{"name": uniqueName("kvs-renamed")}); err != nil {
		t.Fatalf("update: %v", err)
	}
	if err := kvs.DeleteRecord(ctx, "OUTPUT"); err != nil {
		t.Fatalf("delete record: %v", err)
	}
}

// TestRecordPublicURLIsFetchable verifies that the public record URL is fetchable without
// authentication.
func TestRecordPublicURLIsFetchable(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	store, err := client.KeyValueStores().GetOrCreate(ctx, uniqueName("kvs-pub"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.KeyValueStore(store.ID).Delete(ctx) }()
	kvs := client.KeyValueStore(store.ID)

	if err := kvs.SetRecordJSON(ctx, "OUTPUT", map[string]any{"pub": true}); err != nil {
		t.Fatalf("set record: %v", err)
	}
	url, err := kvs.GetRecordPublicURL(ctx, "OUTPUT")
	if err != nil || url == "" {
		t.Fatalf("public url: %q err=%v", url, err)
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("fetch public url: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	_, _ = io.Copy(io.Discard, resp.Body)
	if resp.StatusCode >= 300 {
		t.Fatalf("expected success fetching public url, got %d", resp.StatusCode)
	}
}
