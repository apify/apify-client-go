package apify

import (
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
