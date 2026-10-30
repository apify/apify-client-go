package apify

import (
	"context"
	"strings"
	"testing"
)

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
