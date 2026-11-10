package apify

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

// JSON content types used by the client.
const (
	contentTypeJSON        = "application/json"
	contentTypeJSONCharset = "application/json; charset=utf-8"
)

// Polling constants of [waitForFinish].
const (
	// waitForFinishPollInterval is the pause between polls.
	waitForFinishPollInterval = 250 * time.Millisecond
	// waitForFinishRequestSecs is the longest the API holds one waitForFinish request.
	waitForFinishRequestSecs = int64(60)
	// maxWaitForFinishSecs stands in for "wait indefinitely", which the API cannot express.
	maxWaitForFinishSecs = int64(999999)
	// maxWaitForMissingJob is how long a run or build may keep answering 404 before the wait
	// gives up. It covers database replica lag; a successful poll resets it.
	maxWaitForMissingJob = 3 * time.Second
)

// resourceContext is the URL and shared HTTP client of a resource client. The functions in this
// file are the request primitives that every resource client builds on.
type resourceContext struct {
	http *httpClient
	// url is the full URL of the resource, for example https://api.apify.com/v2/actors/ID.
	url string
	// baseParams are query parameters inherited from a parent resource, such as status on last_run.
	baseParams *QueryParams
	// apiOrigin is the scheme and host the API is reached through.
	apiOrigin string
	// publicOrigin is the scheme and host of public URLs. It defaults to apiOrigin.
	publicOrigin string
	// hasOwnID is true when the resource is addressed by its own ID. A resource reached through a
	// fixed sub-path, such as a run's default dataset, has none: its 404 cannot be told apart
	// from a missing parent, so it is not reported as "absent".
	hasOwnID bool
	// idErr is the validation error of the ID (empty, or a dot segment). The first request returns
	// it instead of sending a malformed request.
	idErr error
}

// newCollectionContext creates a context for a collection endpoint: {base}/{resourcePath}.
func newCollectionContext(hc *httpClient, baseURL, resourcePath string) *resourceContext {
	return newResourceContext(hc, baseURL+"/"+resourcePath, baseURL, false, nil)
}

// newSingleContext creates a context for a single resource: {base}/{resourcePath}/{safeID}.
//
// The getters of resource clients return no error, so an invalid id is kept in idErr and
// returned by the first request.
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

// withoutTimeoutRetry returns a copy of this context whose requests do not retry an attempt that
// timed out (see [httpClient.withoutTimeoutRetry]). Use it for a request that is not idempotent.
func (c *resourceContext) withoutTimeoutRetry() *resourceContext {
	clone := *c
	clone.http = c.http.withoutTimeoutRetry()
	return &clone
}

// withPublicOrigin sets the origin of public URLs.
func (c *resourceContext) withPublicOrigin(publicBaseURL string) *resourceContext {
	c.publicOrigin = originOf(publicBaseURL)
	return c
}

// subURL returns the resource URL with subPath appended.
func (c *resourceContext) subURL(subPath string) string {
	if subPath == "" {
		return c.url
	}
	return c.url + "/" + subPath
}

// publicURL returns subURL with the API origin replaced by the public origin.
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

// shortTimeout, mediumTimeout and longTimeout return the configured timeout tiers (see
// [timeoutTiers]).
func (c *resourceContext) shortTimeout() time.Duration  { return c.http.tiers.short }
func (c *resourceContext) mediumTimeout() time.Duration { return c.http.tiers.medium }
func (c *resourceContext) longTimeout() time.Duration   { return c.http.tiers.long }

// mergedParams returns the inherited params extended by params.
func (c *resourceContext) mergedParams(params *QueryParams) *QueryParams {
	merged := c.baseParams.clone()
	merged.extend(params)
	return merged
}

// originOf returns scheme://host[:port] of a URL.
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

// getResource performs a GET of the resource, or of a fixed sub-path of it, and unwraps the data
// envelope. A 404 gives (zero, false, nil) only if the client has its own ID, since otherwise it
// may mean that the parent is missing. For a 404 that is unambiguous, use [getResourceAlways].
func getResource[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (T, bool, error) {
	return getResourceWithTimeout[T](ctx, c, subPath, params, c.shortTimeout())
}

// getResourceWithTimeout is [getResource] with the given timeout.
func getResourceWithTimeout[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams, timeout time.Duration) (T, bool, error) {
	var zero T
	result, err := getResourceRequired[T](ctx, c, subPath, params, timeout)
	if err != nil {
		if c.hasOwnID && isNotFound(err) {
			return zero, false, nil
		}
		return zero, false, err
	}
	return result, true, nil
}

// maxWaitForFinishHoldSecs is the longest the server holds a request open for waitForFinish.
const maxWaitForFinishHoldSecs = 60

// extendForServerWait adds the time the server holds a request open for waitForFinishSecs to
// base, so the client does not abort a request that is working as asked.
func extendForServerWait(base time.Duration, waitForFinishSecs *int64) time.Duration {
	if waitForFinishSecs == nil || *waitForFinishSecs <= 0 {
		return base
	}
	hold := minInt64(*waitForFinishSecs, maxWaitForFinishHoldSecs)
	return base + time.Duration(hold)*time.Second
}

// getResourceWithServerWait is [getResource] with the timeout extended by [extendForServerWait].
func getResourceWithServerWait[T any](ctx context.Context, c *resourceContext, params *QueryParams, waitForFinishSecs *int64) (T, bool, error) {
	return getResourceWithTimeout[T](ctx, c, "", params, extendForServerWait(c.shortTimeout(), waitForFinishSecs))
}

// getResourceAlways is [getResource], but a 404 always gives (zero, false, nil). Use it for a
// lookup whose 404 is unambiguous, such as a request-queue request by ID.
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

// getResourceRequired is [getResource], but every error, a 404 included, is returned. timeout may
// be [noRequestTimeout].
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

// updateResource performs a PUT with a JSON body and unwraps the data envelope.
func updateResource[T any](ctx context.Context, c *resourceContext, subPath string, body any) (T, error) {
	return updateResourceWithTimeout[T](ctx, c, subPath, body, c.shortTimeout())
}

// updateResourceWithTimeout is [updateResource] with the given timeout.
func updateResourceWithTimeout[T any](ctx context.Context, c *resourceContext, subPath string, body any, timeout time.Duration) (T, error) {
	var zero T
	if c.idErr != nil {
		return zero, c.idErr
	}
	data, err := json.Marshal(body)
	if err != nil {
		return zero, err
	}
	url := c.mergedParams(NewQueryParams()).applyToURL(c.subURL(subPath))
	resp, err := c.http.call(ctx, http.MethodPut, url, data, contentTypeJSON, timeout)
	if err != nil {
		return zero, err
	}
	return parseDataEnvelope[T](resp.body)
}

// deleteResource performs a DELETE of the resource. A 404 is a no-op only if the client has its
// own ID (see [getResource]). For an unambiguous 404, use [deleteResourceAlways].
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

// deleteResourceAlways is [deleteResource], but a 404 is always a no-op.
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

// listResource performs a GET of a paginated list.
func listResource[T any](ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (PaginationList[T], error) {
	return getResourceRequired[PaginationList[T]](ctx, c, subPath, params, c.mediumTimeout())
}

// createResource performs a POST with a JSON body and unwraps the data envelope.
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

// getOrCreateNamed performs POST {collection}?name=..., which gets or creates a named resource.
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

// postWithBody performs a POST with a raw body and unwraps the data envelope.
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

// deleteWithBody performs a DELETE with a JSON body and unwraps the data envelope.
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

// getRaw performs a GET and returns the raw response. A 404 gives (nil, nil), so use it only for
// an unambiguous 404, such as a record lookup by key. Otherwise use [getRawRequired].
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

// getRawRequired is [getRaw], but every error, a 404 included, is returned.
func getRawRequired(ctx context.Context, c *resourceContext, subPath string, params *QueryParams, timeout time.Duration) (*apiResponse, error) {
	if c.idErr != nil {
		return nil, c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	return c.http.call(ctx, http.MethodGet, url, nil, "", timeout)
}

// headExists performs a HEAD request and reports whether the resource exists. A 404 gives false.
func headExists(ctx context.Context, c *resourceContext, subPath string, params *QueryParams) (bool, error) {
	if c.idErr != nil {
		return false, c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodHead, url, nil, "", c.longTimeout())
	if err != nil {
		if isNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// putRaw performs a PUT with raw bytes.
func putRaw(ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body []byte, contentType string) error {
	if c.idErr != nil {
		return c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	_, err := c.http.call(ctx, http.MethodPut, url, body, contentType, c.longTimeout())
	return err
}

// putRawWithEncoding is [putRaw] for a body that is already compressed under contentEncoding. The
// body is sent as it is, and the client's own request compression is skipped, even when
// contentEncoding is empty.
func putRawWithEncoding(ctx context.Context, c *resourceContext, subPath string, params *QueryParams, body []byte, contentType, contentEncoding string) error {
	if c.idErr != nil {
		return c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))
	_, err := c.http.callPreEncoded(ctx, http.MethodPut, url, body, contentType, contentEncoding, c.longTimeout())
	return err
}

// streamRaw sends a GET request and returns the response body unread, with the response headers.
// The request is not retried, and the caller closes the body.
//
// A non-2xx response is returned as an *APIError, except that a 404 gives (nil, nil, nil) when
// notFound is true.
func streamRaw(ctx context.Context, c *resourceContext, subPath string, params *QueryParams, notFound bool) (io.ReadCloser, http.Header, error) {
	if c.idErr != nil {
		return nil, nil, c.idErr
	}
	url := c.mergedParams(params).applyToURL(c.subURL(subPath))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	c.http.applyAuthHeaders(req)

	resp, err := c.http.backend.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode >= maxSuccessStatusCode {
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		apiErr := buildAPIError(resp.StatusCode, body, 1, http.MethodGet, extractPath(url))
		if notFound && isNotFound(apiErr) {
			return nil, nil, nil
		}
		return nil, nil, apiErr
	}
	return resp.Body, resp.Header, nil
}

// terminalChecker reports whether a fetched resource has reached a terminal state.
type terminalChecker[T any] func(*T) bool

// waitForFinish polls the resource until it is terminal, waitSecs elapse, or it has answered 404
// for [maxWaitForMissingJob]. A nil waitSecs waits for [maxWaitForFinishSecs]. When the wait ends
// while the resource is missing, the last 404 error is returned.
func waitForFinish[T any](ctx context.Context, c *resourceContext, waitSecs *int64, resourceName string, isTerminal terminalChecker[T]) (T, error) {
	var zero T

	effectiveWaitSecs := maxWaitForFinishSecs
	if waitSecs != nil {
		effectiveWaitSecs = maxInt64(*waitSecs, 0)
	}
	budget := time.Duration(effectiveWaitSecs) * time.Second

	start := time.Now()
	var (
		resource      T
		present       bool
		notFoundSince time.Time
		notFoundErr   error
	)

	// Fetch at least once, even with a budget of zero.
	for {
		elapsed := time.Since(start)
		remaining := int64((budget - elapsed).Seconds())
		requestSecs := minInt64(maxInt64(remaining, 0), waitForFinishRequestSecs)

		params := NewQueryParams()
		params.AddInt("waitForFinish", &requestSecs)

		// The server holds the request open on purpose, so a poll has no timeout of its own.
		res, err := getResourceRequired[T](ctx, c, "", params, noRequestTimeout)
		ok := err == nil
		if err != nil && (!c.hasOwnID || !isNotFound(err)) {
			return zero, err
		}
		if ok {
			resource, present = res, true
			notFoundSince = time.Time{}
			notFoundErr = nil
			if isTerminal(&resource) {
				return resource, nil
			}
		} else {
			// A 404 may be replica lag, so keep polling until it has lasted maxWaitForMissingJob.
			present = false
			if notFoundSince.IsZero() {
				notFoundSince = time.Now()
			}
			notFoundErr = err
			if time.Since(notFoundSince) >= maxWaitForMissingJob {
				return zero, notFoundErr
			}
		}

		if time.Since(start) >= budget {
			break
		}

		if !sleepWithContext(ctx, waitForFinishPollInterval) {
			return zero, ctx.Err()
		}
	}

	if present {
		// The wait ran out, so this is the last state, for example still running.
		return resource, nil
	}

	// The wait ran out on a 404.
	return zero, notFoundErr
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

// mustMarshal serializes v to JSON for bodies built from primitives, which cannot fail.
func mustMarshal(v any) []byte {
	data, _ := json.Marshal(v)
	return data
}
