package apify

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Per-endpoint base timeouts. These mirror the reference clients' SMALL/MEDIUM/DEFAULT
// timeouts and are the single source of truth (no magic durations sprinkled in callers).
const (
	// defaultRequestTimeout is used when an endpoint does not specify its own (6 minutes).
	defaultRequestTimeout = 360 * time.Second
	// smallRequestTimeout is for fast, idempotent metadata operations (mirrors JS SMALL_TIMEOUT).
	smallRequestTimeout = 5 * time.Second
	// mediumRequestTimeout is for operations that may take a little longer (JS MEDIUM_TIMEOUT).
	mediumRequestTimeout = 30 * time.Second
)

// JSON content types used by the client.
const (
	contentTypeJSON        = "application/json"
	contentTypeJSONCharset = "application/json; charset=utf-8"
)

// How long to wait between polls while waiting for a run/build to finish, and the
// server-side waitForFinish chunk size (the API caps server waiting at 60 seconds).
const (
	waitForFinishPollInterval = 250 * time.Millisecond
	waitForFinishRequestSecs  = int64(60)
)

// resourceContext is the resolved context for a resource client: its base URL and the
// shared HTTP client. The free functions in this file implement the CRUD primitives once,
// so each resource client stays small and consistent (DRY).
type resourceContext struct {
	http *httpClient
	// url is the fully-qualified base URL of the resource,
	// e.g. https://api.apify.com/v2/actors/ID.
	url string
	// baseParams are parameters inherited from a parent resource (e.g. status on last_run).
	baseParams *QueryParams
	// apiOrigin is the origin (scheme + host) the API is reached through.
	apiOrigin string
	// publicOrigin is the origin used to build public, shareable URLs (defaults to apiOrigin).
	publicOrigin string
}

// newCollectionContext creates a context for a collection endpoint: {base}/{resourcePath}.
func newCollectionContext(hc *httpClient, baseURL, resourcePath string) *resourceContext {
	return newResourceContext(hc, baseURL+"/"+resourcePath, baseURL)
}

// newSingleContext creates a context for a single resource: {base}/{resourcePath}/{safeID}.
func newSingleContext(hc *httpClient, baseURL, resourcePath, id string) *resourceContext {
	return newResourceContext(hc, baseURL+"/"+resourcePath+"/"+toSafeID(id), baseURL)
}

func newResourceContext(hc *httpClient, url, baseURL string) *resourceContext {
	origin := originOf(baseURL)
	return &resourceContext{
		http:         hc,
		url:          url,
		baseParams:   NewQueryParams(),
		apiOrigin:    origin,
		publicOrigin: origin,
	}
}

// withPublicOrigin overrides the origin used when building public URLs.
func (c *resourceContext) withPublicOrigin(publicBaseURL string) *resourceContext {
	c.publicOrigin = originOf(publicBaseURL)
	return c
}

// subURL returns this resource's URL with an optional extra path segment appended.
func (c *resourceContext) subURL(subPath string) string {
	if subPath == "" {
		return c.url
	}
	return c.url + "/" + subPath
}

// publicURL builds the public (shareable) form of this resource's URL with an optional
// extra path segment, swapping the API origin for the configured public origin.
func (c *resourceContext) publicURL(subPath string) string {
	apiURL := c.subURL(subPath)
	if c.publicOrigin == c.apiOrigin {
		return apiURL
	}
	if rest, ok := strings.CutPrefix(apiURL, c.apiOrigin); ok {
		return c.publicOrigin + rest
	}
	return apiURL
}

// mergedParams merges the inherited base params with per-call params.
func (c *resourceContext) mergedParams(params *QueryParams) *QueryParams {
	merged := c.baseParams.clone()
	merged.extend(params)
	return merged
}

// originOf extracts the origin (scheme://host[:port]) from a URL, dropping any path.
func originOf(rawURL string) string {
	rest := rawURL
	scheme := ""
	if i := strings.Index(rest, "://"); i >= 0 {
		scheme = rest[:i+3]
		rest = rest[i+3:]
	}
	if i := strings.IndexByte(rest, '/'); i >= 0 {
		rest = rest[:i]
	}
	return scheme + rest
}

// getResource performs a GET that unwraps the data envelope; a not-found maps to (zero,
// false, nil). The bool reports presence.
func getResource[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (T, bool, error) {
	var zero T
	result, err := getResourceRequired[T](ctx, c, subPath, params)
	if err != nil {
		if isNotFound(err) {
			return zero, false, nil
		}
		return zero, false, err
	}
	return result, true, nil
}

// getResourceRequired performs a GET that unwraps the data envelope and propagates errors.
func getResourceRequired[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (T, error) {
	var zero T
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodGet, url, nil, "", defaultRequestTimeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// updateResource performs a PUT with a JSON body, unwrapping the data envelope.
func updateResource[T any](ctx context.Context, c *resourceContext, subPath string, body any) (T, error) {
	var zero T
	data, err := json.Marshal(body)
	if err != nil {
		return zero, err
	}
	url := c.mergedParams(NewQueryParams()).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodPut, url, data, contentTypeJSON, defaultRequestTimeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// deleteResource performs a DELETE; a not-found is treated as a successful no-op.
func deleteResource(ctx context.Context, c *resourceContext, subPath string) error {
	url := c.mergedParams(NewQueryParams()).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodDelete, url, nil, "", defaultRequestTimeout)
	if err != nil && !isNotFound(err) {
		return err
	}
	return nil
}

// listResource performs a GET returning a paginated list (data envelope wrapping items).
func listResource[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (PaginationList[T], error) {
	return getResourceRequired[PaginationList[T]](ctx, c, subPath, params)
}

// createResource performs a POST with a JSON body, unwrapping the data envelope.
func createResource[T any](ctx context.Context, c *resourceContext, params *QueryParams, body any) (T, error) {
	var zero T
	data, err := json.Marshal(body)
	if err != nil {
		return zero, err
	}
	url := c.mergedParams(params).applyToURL(c.subURL(""))
	resp, err := c.http.call(ctx, http.MethodPost, url, data, contentTypeJSON, defaultRequestTimeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// getOrCreateNamed performs a POST that gets-or-creates a named resource
// (POST {collection}?name=...), unwrapping the data envelope.
func getOrCreateNamed[T any](ctx context.Context, c *resourceContext, name string) (T, error) {
	var zero T
	params := NewQueryParams()
	if name != "" {
		params.AddString("name", &name)
	}
	url := params.applyToURL(c.subURL(""))
	resp, err := c.http.call(ctx, http.MethodPost, url, nil, "", defaultRequestTimeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// postWithBody performs a POST with a raw body (optional) and content type, unwrapping the
// data envelope. Used where the input is arbitrary user JSON (actor.start, run.metamorph).
func postWithBody[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body []byte, contentType string) (T, error) {
	var zero T
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodPost, url, body, contentType, defaultRequestTimeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// deleteWithBody performs a DELETE with a JSON body (used for batch request deletion),
// unwrapping the data envelope.
func deleteWithBody[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body any) (T, error) {
	var zero T
	data, err := json.Marshal(body)
	if err != nil {
		return zero, err
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodDelete, url, data, contentTypeJSON, defaultRequestTimeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// getRaw performs a GET returning the raw response (no data envelope). A not-found maps to
// (nil, nil). Used for logs and key-value-store record values.
func getRaw(ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (*apiResponse, error) {
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodGet, url, nil, "", defaultRequestTimeout)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}

// headExists performs a HEAD request, returning whether the resource exists.
func headExists(ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (bool, error) {
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodHead, url, nil, "", defaultRequestTimeout)
	if err != nil {
		if isNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// putRaw performs a PUT with raw bytes and a content type (used for KVS record uploads).
func putRaw(ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body []byte, contentType string) error {
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodPut, url, body, contentType, defaultRequestTimeout)
	return err
}

// terminalChecker reports whether a fetched resource has reached a terminal state.
type terminalChecker[T any] func(*T) bool

// waitForFinish polls a GET endpoint with waitForFinish until the resource reaches a
// terminal state or waitSecs elapses. waitSecs nil waits indefinitely. Returns the latest
// resource (finished, or still running if the wait budget was exhausted).
func waitForFinish[T any](ctx context.Context, c *resourceContext, waitSecs *int64, isTerminal terminalChecker[T]) (T, error) {
	var zero T
	start := time.Now()
	var budget *time.Duration
	if waitSecs != nil {
		d := time.Duration(maxInt64(*waitSecs, 0)) * time.Second
		budget = &d
	}

	for {
		var requestSecs int64
		if budget != nil {
			elapsed := time.Since(start)
			if elapsed >= *budget {
				// Budget exhausted: one final immediate fetch.
				params := NewQueryParams()
				zeroSecs := int64(0)
				params.AddInt("waitForFinish", &zeroSecs)
				return getResourceRequired[T](ctx, c, "", params)
			}
			remaining := int64((*budget - elapsed).Seconds())
			requestSecs = minInt64(remaining, waitForFinishRequestSecs)
		} else {
			requestSecs = waitForFinishRequestSecs
		}

		params := NewQueryParams()
		params.AddInt("waitForFinish", &requestSecs)

		resource, present, err := getResource[T](ctx, c, "", params)
		if err != nil {
			return zero, err
		}
		if present {
			if isTerminal(&resource) {
				return resource, nil
			}
			if budget != nil && time.Since(start) >= *budget {
				return resource, nil
			}
		}

		if !sleepWithContext(ctx, waitForFinishPollInterval) {
			return zero, ctx.Err()
		}
	}
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// mustMarshal serializes v to JSON, returning nil on error. Used for small, internally
// constructed bodies (maps of primitives) that cannot fail to marshal in practice.
func mustMarshal(v any) []byte {
	data, _ := json.Marshal(v)
	return data
}
