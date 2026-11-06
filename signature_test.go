package apify

import (
	"context"
	"encoding/base64"
	"math/big"
	"testing"
)

// Known-answer test pinning createHmacSignature against upstream @apify/utilities:
// HMAC-SHA256 -> hex -> first 30 hex chars -> big integer -> base62 (alphabet 0-9a-zA-Z).
// For key="secret", msg="message" the upstream output is "11GYWmGxviysIBMtnQHBk".
func TestHmacSignatureMatchesUpstream(t *testing.T) {
	sig := createHmacSignature("secret", "message")
	if sig != "11GYWmGxviysIBMtnQHBk" {
		t.Fatalf("unexpected signature %q", sig)
	}
}

func TestToBase62(t *testing.T) {
	cases := map[int64]string{0: "0", 61: "Z", 62: "10"}
	for in, want := range cases {
		got := toBase62(big.NewInt(in))
		if got != want {
			t.Fatalf("toBase62(%d) = %q, want %q", in, got, want)
		}
	}
}

// A non-expiring storage-content signature uses version "0" and expiresAt 0, and is the
// base64url (no-pad) encoding of "0.0.<hmac>" where the hmac is over "0.0.RESID".
func TestStorageContentSignatureNonExpiring(t *testing.T) {
	sig := signStorageContent("secret", "RESID", nil)
	decoded, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		t.Fatalf("not valid base64url: %v", err)
	}
	expected := "0.0." + createHmacSignature("secret", "0.0.RESID")
	if string(decoded) != expected {
		t.Fatalf("got %q, want %q", string(decoded), expected)
	}
}

// A signed public URL needs a positive expiry: zero and a negative value give a URL that is
// already expired, so both are rejected before any request is sent.
func TestPublicURLsRejectANonPositiveExpiry(t *testing.T) {
	ctx := context.Background()
	for _, expiry := range []int64{0, -1} {
		backend := &mockBackend{responses: okData()}
		client := testClient(backend, 0)
		calls := map[string]func() (string, error){
			"CreateItemsPublicURL": func() (string, error) {
				return client.Dataset("d").CreateItemsPublicURL(ctx, DatasetListItemsOptions{}, Ptr(expiry))
			},
			"CreateItemsPublicURLWithFormat": func() (string, error) {
				return client.Dataset("d").CreateItemsPublicURLWithFormat(ctx, DatasetListItemsOptions{}, Ptr(expiry), FormatCSV)
			},
			"CreateKeysPublicURL": func() (string, error) {
				return client.KeyValueStore("s").CreateKeysPublicURL(ctx, Ptr(expiry))
			},
		}
		for name, call := range calls {
			if url, err := call(); err == nil {
				t.Errorf("%s with expiresInSecs %d: expected an error, got %q", name, expiry, url)
			}
		}
		if backend.calls != 0 {
			t.Errorf("expiresInSecs %d: expected no request, got %d", expiry, backend.calls)
		}
	}
}

func TestPublicURLsAcceptAPositiveExpiry(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	if _, err := testClient(backend, 0).Dataset("d").CreateItemsPublicURL(context.Background(), DatasetListItemsOptions{}, Ptr(int64(60))); err != nil {
		t.Fatalf("a positive expiry must be accepted: %v", err)
	}
}
