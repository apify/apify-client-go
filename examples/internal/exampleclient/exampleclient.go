// Package exampleclient builds the Apify client used by the runnable examples.
//
// It centralizes reading APIFY_TOKEN and APIFY_API_URL so every example targets the same API
// as the integration test suite: production (https://api.apify.com/v2) by default, or the
// server named by APIFY_API_URL when set (for example a staging server in CI). Keeping this in
// one place avoids repeating the base-URL handling in every example program.
package exampleclient

import (
	"os"
	"strings"

	apify "github.com/apify/apify-client-go"
)

// New builds an authenticated client from the environment: the token is read from APIFY_TOKEN
// and the base URL from APIFY_API_URL (see [baseURL]).
func New() *apify.ApifyClient {
	return apify.NewClient(
		apify.WithToken(os.Getenv("APIFY_TOKEN")),
		apify.WithBaseURL(baseURL()),
	)
}

// NewAnonymous builds a client with no API token, for the token-free examples. It still honors
// APIFY_API_URL so those examples run against the configured API too.
func NewAnonymous() *apify.ApifyClient {
	return apify.NewClient(apify.WithBaseURL(baseURL()))
}

// baseURL resolves the client base URL from APIFY_API_URL, falling back to
// https://api.apify.com/v2. APIFY_API_URL includes the /v2 suffix (per the integration-test
// contract), which is stripped here because the client appends /v2 itself.
func baseURL() string {
	u := os.Getenv("APIFY_API_URL")
	if u == "" {
		u = "https://api.apify.com/v2"
	}
	return strings.TrimSuffix(strings.TrimRight(u, "/"), "/v2")
}
