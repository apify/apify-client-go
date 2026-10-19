package apify_test

import (
	"bytes"
	"compress/gzip"
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

// TestSetRecordRawWithEncodingRoundTrips uploads a pre-gzipped value via
// SetRecordRawWithEncoding (a Content-Encoding: gzip PUT, which the client sends as-is rather
// than compressing a second time) and verifies the original content survives the round trip.
//
// The API accepts the compressed upload (saving upload bandwidth, per its documented contract)
// but may decompress it on ingest rather than echoing the same Content-Encoding back on GET; this
// asserts only the content client code can rely on either way — the value is never corrupted —
// without assuming which of the two the API currently does.
func TestSetRecordRawWithEncodingRoundTrips(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	store, err := client.KeyValueStores().GetOrCreate(ctx, uniqueName("kvs-encoding"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.KeyValueStore(store.ID).Delete(ctx) }()
	kvs := client.KeyValueStore(store.ID)

	original := []byte("the quick brown fox jumps over the lazy dog, repeated for good measure. " +
		"the quick brown fox jumps over the lazy dog, repeated for good measure.")
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(original); err != nil {
		t.Fatalf("gzip write: %v", err)
	}
	if err := gz.Close(); err != nil {
		t.Fatalf("gzip close: %v", err)
	}
	compressed := buf.Bytes()

	if err := kvs.SetRecordRawWithEncoding(ctx, "REPORT", compressed, "text/plain", "gzip"); err != nil {
		t.Fatalf("set record with encoding: %v", err)
	}

	rec, ok, err := kvs.GetRecord(ctx, "REPORT")
	if err != nil || !ok {
		t.Fatalf("get record: ok=%v err=%v", ok, err)
	}

	got := rec.Value
	// Accept either form: the record stored (and served back) still gzip-compressed, or
	// decompressed by the API on ingest.
	if bytes.HasPrefix(got, []byte{0x1f, 0x8b}) {
		reader, err := gzip.NewReader(bytes.NewReader(got))
		if err != nil {
			t.Fatalf("record value looks gzip-compressed but does not decode: %v", err)
		}
		got, err = io.ReadAll(reader)
		if err != nil {
			t.Fatalf("gunzip record value: %v", err)
		}
	}
	if !bytes.Equal(got, original) {
		t.Fatalf("round-tripped content mismatch: got %q, want %q", got, original)
	}
}

// TestGetRecordStream verifies that GetRecordStream serves the same record GetRecord would, as
// a stream the caller drains incrementally.
func TestGetRecordStream(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	store, err := client.KeyValueStores().GetOrCreate(ctx, uniqueName("kvs-stream"))
	if err != nil {
		t.Fatalf("get-or-create: %v", err)
	}
	defer func() { _ = client.KeyValueStore(store.ID).Delete(ctx) }()
	kvs := client.KeyValueStore(store.ID)

	if err := kvs.SetRecordRaw(ctx, "OUTPUT", []byte("hello from a stream"), "text/plain"); err != nil {
		t.Fatalf("set record: %v", err)
	}

	rec, ok, err := kvs.GetRecordStream(ctx, "OUTPUT")
	if err != nil || !ok {
		t.Fatalf("get record stream: ok=%v err=%v", ok, err)
	}
	defer func() { _ = rec.Value.Close() }()

	streamed, err := io.ReadAll(rec.Value)
	if err != nil {
		t.Fatalf("read stream: %v", err)
	}
	if string(streamed) != "hello from a stream" {
		t.Fatalf("unexpected streamed content: %q", string(streamed))
	}

	buffered, ok, err := kvs.GetRecord(ctx, "OUTPUT")
	if err != nil || !ok {
		t.Fatalf("get record: ok=%v err=%v", ok, err)
	}
	if string(buffered.Value) != string(streamed) || buffered.ContentType != rec.ContentType {
		t.Fatalf("streamed and buffered records differ: %+v vs stream content %q (content type %q)",
			buffered, string(streamed), rec.ContentType)
	}

	// A missing record is unambiguous: GetRecordStream resolves to absent, not an error.
	missing, ok, err := kvs.GetRecordStream(ctx, "NO-SUCH-KEY")
	if err != nil || ok || missing != nil {
		t.Fatalf("expected absent for a missing key, got missing=%v ok=%v err=%v", missing, ok, err)
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
