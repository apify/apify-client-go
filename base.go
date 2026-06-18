package apify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// defaultRequestTimeout is the per-request timeout applied to all API calls (6 minutes),
// matching the reference client's DEFAULT_TIMEOUT_MILLIS. It is the single source of truth
// for the request timeout (no magic durations sprinkled in callers).
const defaultRequestTimeout = 360 * time.Second

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
	// maxWaitForFinishSecs is the finite upper bound used when the caller asks to wait
	// "indefinitely" (waitSecs == nil). The API will not accept "Infinity", and an
	// unbounded client loop can spin forever if the resource keeps returning 404 (e.g. a
	// just-started run whose database replica lags). 999999 seconds is more than 11 days,
	// effectively indefinite while still guaranteeing termination. Mirrors the reference
	// client's MAX_WAIT_FOR_FINISH.
	maxWaitForFinishSecs = int64(999999)
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
// terminal state or the wait budget elapses. waitSecs nil means "wait indefinitely", which
// is implemented as a finite but very large bound (maxWaitForFinishSecs) so the loop always
// terminates.
//
// The budget is a pure time bound, evaluated independently of whether the resource is
// currently present. A just-started run/build can transiently return 404 (database-replica
// lag); like the reference client, we treat that as "not yet available", keep polling on the
// time bound, and only after the budget is exhausted do we decide the outcome. If the
// resource never became available within the budget, a descriptive error is returned naming
// the resource (resourceName, e.g. "run" or "build").
func waitForFinish[T any](ctx context.Context, c *resourceContext, waitSecs *int64, resourceName string, isTerminal terminalChecker[T]) (T, error) {
	var zero T

	effectiveWaitSecs := maxWaitForFinishSecs
	if waitSecs != nil {
		effectiveWaitSecs = maxInt64(*waitSecs, 0)
	}
	budget := time.Duration(effectiveWaitSecs) * time.Second

	start := time.Now()
	var (
		resource T
		present  bool
	)

	// do-while: always perform at least one fetch (matching the reference client), then
	// repeat only while the pure time bound has not been reached.
	for {
		elapsed := time.Since(start)
		remaining := int64((budget - elapsed).Seconds())
		requestSecs := minInt64(maxInt64(remaining, 0), waitForFinishRequestSecs)

		params := NewQueryParams()
		params.AddInt("waitForFinish", &requestSecs)

		res, ok, err := getResource[T](ctx, c, "", params)
		if err != nil {
			return zero, err
		}
		// A transient 404 (ok == false) is not fatal: keep the last known state and poll
		// again until the time budget runs out.
		if ok {
			resource, present = res, true
			if isTerminal(&resource) {
				return resource, nil
			}
		}

		// Pure time bound: stop once the budget is exhausted, regardless of presence.
		if time.Since(start) >= budget {
			break
		}

		if !sleepWithContext(ctx, waitForFinishPollInterval) {
			return zero, ctx.Err()
		}
	}

	if present {
		// Budget exhausted but the resource was seen at least once: return its latest state
		// (e.g. still running). Mirrors the reference client returning the last fetched job.
		return resource, nil
	}

	// The resource never became available within the wait budget.
	return zero, fmt.Errorf(
		"waiting for %s to finish failed: cannot fetch %s details from the server",
		resourceName, resourceName,
	)
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
