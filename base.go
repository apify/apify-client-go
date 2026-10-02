package apify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
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
	// hasOwnID is true when this resource is addressed by its own id (constructed via
	// newSingleContext). It is false for a resource reached through a fixed sub-path with no id
	// of its own (e.g. a run's default dataset, via newCollectionContext): there, a 404 cannot
	// be told apart from the parent being missing, so it must not be swallowed as "absent".
	hasOwnID bool
	// idErr, if non-nil, is a validation error discovered while building this resource's URL
	// (an empty or path-traversing id). It is returned by the first request this context
	// attempts, instead of silently sending a malformed request.
	idErr error
}

// newCollectionContext creates a context for a collection endpoint: {base}/{resourcePath}.
func newCollectionContext(hc *httpClient, baseURL, resourcePath string) *resourceContext {
	return newResourceContext(hc, baseURL+"/"+resourcePath, baseURL, false, nil)
}

// newSingleContext creates a context for a single resource: {base}/{resourcePath}/{safeID}.
//
// id is validated and percent-encoded (see safePathSegment); an invalid id (empty, or a dot
// segment) does not fail here; synchronously since every resource getter in this package
// returns a client struct directly rather than an error. Instead, the resource context
// remembers the error and every CRUD call this client later attempts fails fast with it.
func newSingleContext(hc *httpClient, baseURL, resourcePath, id string) *resourceContext {
	safeID, err := toSafeID(id)
	return newResourceContext(hc, baseURL+"/"+resourcePath+"/"+safeID, baseURL, true, err)
}

func newResourceContext(hc *httpClient, url, baseURL string, hasOwnID bool, idErr error) *resourceContext {
	origin := originOf(baseURL)
	return &resourceContext{
		http:         hc,
		url:          url,
		baseParams:   NewQueryParams(),
		apiOrigin:    origin,
		publicOrigin: origin,
		hasOwnID:     hasOwnID,
		idErr:        idErr,
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

// shortTimeout, mediumTimeout and longTimeout return this resource's configured timeout tiers
// (see [timeoutTiers]), read through to the shared httpClient so every resource client picks up
// a [WithTimeoutShort]/[WithTimeoutMedium]/[WithTimeoutLong] override.
func (c *resourceContext) shortTimeout() time.Duration  { return c.http.tiers.short }
func (c *resourceContext) mediumTimeout() time.Duration { return c.http.tiers.medium }
func (c *resourceContext) longTimeout() time.Duration   { return c.http.tiers.long }

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

// getResource performs a GET addressing this client's own resource (subPath "" for the
// resource itself, or a fixed sub-path such as a singleton the resource owns) and unwraps the
// data envelope. A not-found maps to (zero, false, nil) only when the client has its own id
// (c.hasOwnID): there, the 404 unambiguously means this resource is gone. A client reached
// through a fixed sub-path with no id of its own (e.g. a run's default dataset) cannot tell its
// own 404 apart from the parent's, so the error always propagates instead. Mirrors the
// reference client's catchNotFoundForResourceOrThrow. For a 404 that is unambiguous regardless
// of how the client was reached (e.g. a lookup by key or request id), use [getResourceAlways].
func getResource[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (T, bool, error) {
	var zero T
	result, err := getResourceRequired[T](ctx, c, subPath, params, c.shortTimeout())
	if err != nil {
		if c.hasOwnID && isNotFound(err) {
			return zero, false, nil
		}
		return zero, false, err
	}
	return result, true, nil
}

// maxWaitForFinishHoldSecs caps how much extra time a waitForFinish request parameter can add
// to a request's client-side timeout, matching the reference client's
// MAX_WAIT_FOR_FINISH_HOLD_SECS (the server itself never holds a single request open longer).
const maxWaitForFinishHoldSecs = 60

// extendForServerWait extends base by the time the server may legitimately hold the connection
// open to honor a waitForFinish request parameter (capped at maxWaitForFinishHoldSecs), so the
// client does not abort a request the server is still correctly working on. waitForFinishSecs
// nil or <= 0 returns base unchanged. Mirrors the reference client's timeoutForWaitForFinish.
func extendForServerWait(base time.Duration, waitForFinishSecs *int64) time.Duration {
	if waitForFinishSecs == nil || *waitForFinishSecs <= 0 {
		return base
	}
	hold := minInt64(*waitForFinishSecs, maxWaitForFinishHoldSecs)
	return base + time.Duration(hold)*time.Second
}

// getResourceWithServerWait is [getResource] with its timeout extended by
// [extendForServerWait], for a get(waitForFinish) call that asks the server to hold the
// response (RunClient.GetWithWait, BuildClient.GetWithWait).
func getResourceWithServerWait[T any](ctx context.Context, c *resourceContext, params *QueryParams, waitForFinishSecs *int64) (T, bool, error) {
	var zero T
	timeout := extendForServerWait(c.shortTimeout(), waitForFinishSecs)
	result, err := getResourceRequired[T](ctx, c, "", params, timeout)
	if err != nil {
		if c.hasOwnID && isNotFound(err) {
			return zero, false, nil
		}
		return zero, false, err
	}
	return result, true, nil
}

// getResourceAlways performs a GET that unwraps the data envelope; a not-found always maps to
// (zero, false, nil), regardless of c.hasOwnID. Used for a lookup that is unambiguous no matter
// how the parent client was reached, such as a request-queue request by id.
func getResourceAlways[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (T, bool, error) {
	var zero T
	result, err := getResourceRequired[T](ctx, c, subPath, params, c.shortTimeout())
	if err != nil {
		if isNotFound(err) {
			return zero, false, nil
		}
		return zero, false, err
	}
	return result, true, nil
}

// getResourceRequired performs a GET that unwraps the data envelope and propagates every error,
// including a 404. Used both where a 404 is never ambiguous (list, batch and lookup endpoints)
// and, since [DatasetClient.GetStatistics]/[ScheduleClient.GetLog]/[TaskClient.GetInput], where
// it always means the parent resource is gone.
//
// timeout is the caller's chosen timeout tier (see [timeoutTiers]), or [noRequestTimeout] for a
// request that must not be bounded by one (the polling behind [waitForFinish]).
func getResourceRequired[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams, timeout time.Duration) (T, error) {
	var zero T
	if c.idErr != nil {
		return zero, c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodGet, url, nil, "", timeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// updateResource performs a PUT with a JSON body, unwrapping the data envelope. Metadata writes
// use the "short" timeout tier.
func updateResource[T any](ctx context.Context, c *resourceContext, subPath string, body any) (T, error) {
	var zero T
	if c.idErr != nil {
		return zero, c.idErr
	}
	data, err := json.Marshal(body)
	if err != nil {
		return zero, err
	}
	url := c.mergedParams(NewQueryParams()).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodPut, url, data, contentTypeJSON, c.shortTimeout())
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// deleteResource performs a DELETE addressing this client's own resource. A not-found is
// treated as a successful no-op only when the client has its own id (c.hasOwnID); otherwise the
// error propagates. See [getResource] for why. For a delete that is unambiguous regardless of
// how the parent client was reached (e.g. deleting a key-value-store record by key), use
// [deleteResourceAlways].
func deleteResource(ctx context.Context, c *resourceContext, subPath string) error {
	if c.idErr != nil {
		return c.idErr
	}
	url := c.mergedParams(NewQueryParams()).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodDelete, url, nil, "", c.shortTimeout())
	if err != nil && (!c.hasOwnID || !isNotFound(err)) {
		return err
	}
	return nil
}

// deleteResourceAlways performs a DELETE; a not-found is always treated as a successful no-op,
// regardless of c.hasOwnID.
func deleteResourceAlways(ctx context.Context, c *resourceContext, subPath string) error {
	if c.idErr != nil {
		return c.idErr
	}
	url := c.mergedParams(NewQueryParams()).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodDelete, url, nil, "", c.shortTimeout())
	if err != nil && !isNotFound(err) {
		return err
	}
	return nil
}

// listResource performs a GET returning a paginated list (data envelope wrapping items). Every
// collection's List uses the "medium" timeout tier.
func listResource[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (PaginationList[T], error) {
	return getResourceRequired[PaginationList[T]](ctx, c, subPath, params, c.mediumTimeout())
}

// createResource performs a POST with a JSON body, unwrapping the data envelope. timeout is the
// caller's chosen tier: most Create endpoints are "short", a few (e.g. ActorCollectionClient's)
// are "medium", matching the reference client per resource.
func createResource[T any](ctx context.Context, c *resourceContext, params *QueryParams, body any, timeout time.Duration) (T, error) {
	var zero T
	if c.idErr != nil {
		return zero, c.idErr
	}
	data, err := json.Marshal(body)
	if err != nil {
		return zero, err
	}
	url := c.mergedParams(params).applyToURL(c.subURL(""))
	resp, err := c.http.call(ctx, http.MethodPost, url, data, contentTypeJSON, timeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// getOrCreateNamed performs a POST that gets-or-creates a named resource
// (POST {collection}?name=...), unwrapping the data envelope. Uses the "short" timeout tier.
func getOrCreateNamed[T any](ctx context.Context, c *resourceContext, name string) (T, error) {
	var zero T
	if c.idErr != nil {
		return zero, c.idErr
	}
	params := NewQueryParams()
	if name != "" {
		params.AddString("name", &name)
	}
	url := params.applyToURL(c.subURL(""))
	resp, err := c.http.call(ctx, http.MethodPost, url, nil, "", c.shortTimeout())
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// postWithBody performs a POST with a raw body (optional) and content type, unwrapping the
// data envelope. Used where the input is arbitrary user JSON (actor.start, run.metamorph).
// timeout is the caller's chosen tier (trigger/batch operations are typically "medium").
func postWithBody[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body []byte, contentType string, timeout time.Duration) (T, error) {
	var zero T
	if c.idErr != nil {
		return zero, c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodPost, url, body, contentType, timeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// deleteWithBody performs a DELETE with a JSON body (used for batch request deletion),
// unwrapping the data envelope. Uses the "short" timeout tier (matching the reference client's
// batchDeleteRequests).
func deleteWithBody[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body any) (T, error) {
	var zero T
	if c.idErr != nil {
		return zero, c.idErr
	}
	data, err := json.Marshal(body)
	if err != nil {
		return zero, err
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodDelete, url, data, contentTypeJSON, c.shortTimeout())
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// getRaw performs a GET returning the raw response (no data envelope), using the "long" timeout
// tier (a data download). A not-found maps to (nil, nil), regardless of c.hasOwnID: used only
// where that is unambiguous, such as a key-value-store record lookup by key. For a fixed
// sub-path where a 404 always means the parent resource is gone (e.g. dataset statistics), use
// [getRawRequired] instead.
func getRaw(ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (*apiResponse, error) {
	resp, err := getRawRequired(ctx, c, subPath, params, c.longTimeout())
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp, nil
}

// getRawRequired performs a GET returning the raw response (no data envelope) and propagates
// every error, including a 404. timeout is the caller's chosen tier.
func getRawRequired(ctx context.Context, c *resourceContext, subPath string, params *QueryParams, timeout time.Duration) (*apiResponse, error) {
	if c.idErr != nil {
		return nil, c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	return c.http.call(ctx, http.MethodGet, url, nil, "", timeout)
}

// headExists performs a HEAD request, returning whether the resource exists, using the "short"
// timeout tier. A not-found always maps to (false, nil): used only for an unambiguous lookup (a
// key-value-store record by key).
func headExists(ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (bool, error) {
	if c.idErr != nil {
		return false, c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodHead, url, nil, "", c.shortTimeout())
	if err != nil {
		if isNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// putRaw performs a PUT with raw bytes and a content type (used for KVS record uploads), using
// the "long" timeout tier (a data upload).
func putRaw(ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body []byte, contentType string) error {
	if c.idErr != nil {
		return c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodPut, url, body, contentType, c.longTimeout())
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

		// The poll itself runs with no client-imposed timeout (noRequestTimeout): the server
		// may legitimately hold the connection open for up to waitForFinishRequestSecs while
		// waiting for the job to finish, so a fixed per-request budget would abort it for doing
		// exactly what was asked. The overall wait is still bounded by the pure time budget
		// above.
		res, err := getResourceRequired[T](ctx, c, "", params, noRequestTimeout)
		ok := err == nil
		if err != nil && (!c.hasOwnID || !isNotFound(err)) {
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
