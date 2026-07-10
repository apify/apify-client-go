package apify

import (
	"runtime"
	"strings"
	"testing"
)

func TestVersionConstants(t *testing.T) {
	if ClientVersion == "" || ClientVersion[0] < '0' || ClientVersion[0] > '9' {
		t.Fatalf("ClientVersion should start with a digit, got %q", ClientVersion)
	}
	if !strings.HasPrefix(APISpecVersion, "v2-") || !strings.HasSuffix(APISpecVersion, "Z") {
		t.Fatalf("APISpecVersion should match v2-...Z, got %q", APISpecVersion)
	}
}

func TestUserAgentFormat(t *testing.T) {
	client := NewClient("token")
	ua := client.UserAgent()
	if !strings.HasPrefix(ua, "ApifyClient/"+ClientVersion+" (") {
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

// The User-Agent OS token must match the platform names the other Apify clients emit
// (Node os.platform() / Python sys.platform): Windows is "win32" and Solaris is "sunos",
// while every other GOOS value (linux, darwin, android, ...) passes through unchanged.
func TestPlatformToken(t *testing.T) {
	cases := map[string]string{
		"windows": "win32",
		"solaris": "sunos",
		"linux":   "linux",
		"darwin":  "darwin",
		"android": "android",
		"freebsd": "freebsd",
	}
	for goos, want := range cases {
		if got := platformToken(goos); got != want {
			t.Errorf("platformToken(%q) = %q, want %q", goos, got, want)
		}
	}
}

// The OS token emitted for the platform the tests run on must be the aligned token for that
// platform, i.e. the mapping applied to runtime.GOOS.
func TestUserAgentOSTokenMatchesPlatform(t *testing.T) {
	ua := NewClient("token").UserAgent()
	want := "(" + platformToken(runtime.GOOS) + "; "
	if !strings.Contains(ua, want) {
		t.Fatalf("user-agent OS token = %q, want it to contain %q", ua, want)
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

// defaultIsAtHome must be driven solely by APIFY_IS_AT_HOME (per the requirements and the JS
// reference); a bare isAtHome env var must NOT affect it.
func TestDefaultIsAtHomeReadsOnlyApifyIsAtHome(t *testing.T) {
	t.Setenv(envIsAtHome, "")
	t.Setenv("isAtHome", "")
	if defaultIsAtHome() {
		t.Fatal("expected isAtHome=false when APIFY_IS_AT_HOME is not set")
	}

	// A bare isAtHome env var must NOT flip the flag (it is not the mandated variable).
	t.Setenv("isAtHome", "1")
	if defaultIsAtHome() {
		t.Fatal("expected isAtHome=false: a bare isAtHome env var must not affect the flag")
	}

	// Only APIFY_IS_AT_HOME drives the flag.
	t.Setenv("isAtHome", "")
	t.Setenv(envIsAtHome, "1")
	if !defaultIsAtHome() {
		t.Fatalf("expected isAtHome=true when %s is set", envIsAtHome)
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
