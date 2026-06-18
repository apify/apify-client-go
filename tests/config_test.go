package apify_test

import "testing"

func TestResolveBaseURLDefault(t *testing.T) {
	if got := resolveBaseURL(""); got != "https://api.apify.com" {
		t.Fatalf("expected default origin, got %q", got)
	}
}

func TestResolveBaseURLStripsV2(t *testing.T) {
	if got := resolveBaseURL("https://api.example.com/v2"); got != "https://api.example.com" {
		t.Fatalf("got %q", got)
	}
	if got := resolveBaseURL("https://api.example.com/v2/"); got != "https://api.example.com" {
		t.Fatalf("trailing-slash variant: got %q", got)
	}
}
