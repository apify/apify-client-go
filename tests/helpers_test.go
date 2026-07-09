// Package apify_test holds the integration test suite for the Apify Go client.
//
// All integration tests require a valid APIFY_TOKEN for the test account. The API base URL
// is taken from APIFY_API_URL (which includes the /v2 suffix) and falls back to
// https://api.apify.com/v2.
//
// Tests are designed to run concurrently — including against the same test account from
// several language clients at once — so every test creates uniquely-named resources and
// cleans them up afterwards.
package apify_test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"strings"
	"testing"
	"time"

	apify "github.com/apify/apify-client-go"
)

// defaultAPIURL mirrors the integration-test contract fallback.
const defaultAPIURL = "https://api.apify.com/v2"

// resolveBaseURL derives the client base_url from an optional APIFY_API_URL.
//
// APIFY_API_URL includes the /v2 suffix (per the integration-test contract) and falls back
// to https://api.apify.com/v2. Since the client appends /v2 itself, the suffix is stripped.
func resolveBaseURL(apiURL string) string {
	if apiURL == "" {
		apiURL = defaultAPIURL
	}
	trimmed := strings.TrimRight(apiURL, "/")
	trimmed = strings.TrimSuffix(trimmed, "/v2")
	return trimmed
}

// requireClient returns a configured client, or skips the test if APIFY_TOKEN is unset.
func requireClient(t *testing.T) *apify.ApifyClient {
	t.Helper()
	token := os.Getenv("APIFY_TOKEN")
	if token == "" {
		t.Skip("skipping: APIFY_TOKEN is not set")
	}
	return apify.NewClient(
		token,
		apify.WithBaseURL(resolveBaseURL(os.Getenv("APIFY_API_URL"))),
	)
}

// testContext returns a context with a generous timeout for a single test.
func testContext(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()
	return context.WithTimeout(context.Background(), 5*time.Minute)
}

// uniqueName generates a collision-resistant resource name for test isolation.
//
// The name embeds the test-specific prefix and a random hex fragment, kept short enough for
// Apify's naming limits. The random component lets the same test run in parallel (across
// processes and languages) without clobbering shared state.
func uniqueName(prefix string) string {
	var buf [6]byte
	_, _ = rand.Read(buf[:])
	return "go-test-" + prefix + "-" + hex.EncodeToString(buf[:])
}

// ptr returns a pointer to v, for the *T option fields.
func ptr[T any](v T) *T { return &v }
