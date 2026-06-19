package apify

import (
	"strings"
	"testing"
)

func TestVersionConstants(t *testing.T) {
	if CLIENT_VERSION == "" || CLIENT_VERSION[0] < '0' || CLIENT_VERSION[0] > '9' {
		t.Fatalf("CLIENT_VERSION should start with a digit, got %q", CLIENT_VERSION)
	}
	if !strings.HasPrefix(API_SPEC_VERSION, "v2-") || !strings.HasSuffix(API_SPEC_VERSION, "Z") {
		t.Fatalf("API_SPEC_VERSION should match v2-...Z, got %q", API_SPEC_VERSION)
	}
}

func TestUserAgentFormat(t *testing.T) {
	client := NewClient("token")
	ua := client.UserAgent()
	if !strings.HasPrefix(ua, "ApifyClient/"+CLIENT_VERSION+" (") {
		t.Fatalf("unexpected user-agent prefix: %q", ua)
	}
	if !strings.Contains(ua, "Go/") {
		t.Fatalf("user-agent must report the Go version: %q", ua)
	}
	if !strings.Contains(ua, "isAtHome/") {
		t.Fatalf("user-agent must report isAtHome: %q", ua)
	}
	// The Go/ token must be followed by a real version digit, not empty.
	after := strings.SplitN(ua, "Go/", 2)[1]
	if len(after) == 0 || after[0] < '0' || after[0] > '9' {
		t.Fatalf("Go version must be a real version, got %q", ua)
	}
}

func TestUserAgentIsAtHomeFlag(t *testing.T) {
	// isAtHome flag must be driven by the platform env var, not flipped arbitrarily.
	offClient := NewClientWithOptions(WithToken("t"), withIsAtHomeFn(func() bool { return false }))
	if !strings.Contains(offClient.UserAgent(), "isAtHome/false") {
		t.Fatalf("expected isAtHome/false, got %q", offClient.UserAgent())
	}
	onClient := NewClientWithOptions(WithToken("t"), withIsAtHomeFn(func() bool { return true }))
	if !strings.Contains(onClient.UserAgent(), "isAtHome/true") {
		t.Fatalf("expected isAtHome/true, got %q", onClient.UserAgent())
	}
}

// defaultIsAtHome must honour BOTH the JS-consistent APIFY_IS_AT_HOME env var and the bare
// isAtHome name mandated by the requirements doc, resolving the two conflicting requirements.
func TestDefaultIsAtHomeReadsBothEnvVars(t *testing.T) {
	t.Setenv(envIsAtHomePrimary, "")
	t.Setenv(envIsAtHomeLegacy, "")
	if defaultIsAtHome() {
		t.Fatal("expected isAtHome=false when neither env var is set")
	}

	t.Setenv(envIsAtHomePrimary, "1")
	if !defaultIsAtHome() {
		t.Fatalf("expected isAtHome=true when %s is set", envIsAtHomePrimary)
	}

	t.Setenv(envIsAtHomePrimary, "")
	t.Setenv(envIsAtHomeLegacy, "1")
	if !defaultIsAtHome() {
		t.Fatalf("expected isAtHome=true when %s is set", envIsAtHomeLegacy)
	}
}

func TestUserAgentSuffix(t *testing.T) {
	client := NewClientWithOptions(WithToken("t"), WithUserAgentSuffix("MyTool/1.0"))
	if !strings.HasSuffix(client.UserAgent(), "; MyTool/1.0") {
		t.Fatalf("expected user-agent suffix, got %q", client.UserAgent())
	}
}

func TestBaseURLDefaultAndV2Suffix(t *testing.T) {
	client := NewClient("t")
	if client.APIBaseURL() != "https://api.apify.com/v2" {
		t.Fatalf("unexpected default base URL: %q", client.APIBaseURL())
	}
}

func TestBaseURLOverrideAppendsV2(t *testing.T) {
	client := NewClientWithOptions(WithToken("t"), WithBaseURL("https://api.example.com/"))
	if client.APIBaseURL() != "https://api.example.com/v2" {
		t.Fatalf("unexpected base URL: %q", client.APIBaseURL())
	}
}
