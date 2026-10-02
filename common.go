package apify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"strconv"
	"strings"
)

// dataEnvelope unwraps the top-level `{ "data": ... }` wrapper used by most Apify endpoints.
type dataEnvelope[T any] struct {
	Data T `json:"data"`
}

// parseDataEnvelope parses a JSON response body wrapped in a `data` envelope.
func parseDataEnvelope[T any](body []byte) (T, error) {
	var env dataEnvelope[T]
	err := json.Unmarshal(body, &env)
	return env.Data, err
}

const notFoundStatusCode = 404

// isNotFound reports whether err is an API error with a 404 status, mirroring the reference
// clients' NotFoundError: a get/delete/head on a missing resource should resolve to "absent"
// rather than raise. Any 404 counts, regardless of its machine-readable `type` — the reference
// clients used to swallow only the "record-not-found"/"record-or-token-not-found" types, but
// that let an unrelated 404 (e.g. a typo'd sub-path) pass silently for "not found" too, so v3
// dropped the type check.
func isNotFound(err error) bool {
	apiErr, ok := AsAPIError(err)
	if !ok {
		return false
	}
	return apiErr.StatusCode == notFoundStatusCode
}

// QueryParams is an ordered collection of query parameters that omits absent values and
// encodes booleans as 1/0, matching the Apify API conventions.
type QueryParams struct {
	pairs [][2]string
}

// NewQueryParams returns an empty QueryParams.
func NewQueryParams() *QueryParams { return &QueryParams{} }

// AddString adds a string parameter if value is non-nil.
func (q *QueryParams) AddString(key string, value *string) *QueryParams {
	if value != nil {
		q.pairs = append(q.pairs, [2]string{key, *value})
	}
	return q
}

// AddInt adds an integer parameter if value is non-nil.
func (q *QueryParams) AddInt(key string, value *int64) *QueryParams {
	if value != nil {
		q.pairs = append(q.pairs, [2]string{key, strconv.FormatInt(*value, 10)})
	}
	return q
}

// AddFloat adds a floating-point parameter if value is non-nil.
func (q *QueryParams) AddFloat(key string, value *float64) *QueryParams {
	if value != nil {
		q.pairs = append(q.pairs, [2]string{key, strconv.FormatFloat(*value, 'f', -1, 64)})
	}
	return q
}

// AddBool adds a boolean parameter, encoded as 1/0, if value is non-nil.
func (q *QueryParams) AddBool(key string, value *bool) *QueryParams {
	if value != nil {
		v := "0"
		if *value {
			v = "1"
		}
		q.pairs = append(q.pairs, [2]string{key, v})
	}
	return q
}

// AddCSV adds a comma-joined list parameter if value is non-empty.
func (q *QueryParams) AddCSV(key string, value []string) *QueryParams {
	if len(value) > 0 {
		q.pairs = append(q.pairs, [2]string{key, strings.Join(value, ",")})
	}
	return q
}

// addRaw appends an already-stringified key/value pair.
func (q *QueryParams) addRaw(key, value string) *QueryParams {
	q.pairs = append(q.pairs, [2]string{key, value})
	return q
}

// IsEmpty reports whether no parameters were added.
func (q *QueryParams) IsEmpty() bool { return len(q.pairs) == 0 }

// clone returns a copy of q.
func (q *QueryParams) clone() *QueryParams {
	out := &QueryParams{pairs: make([][2]string, len(q.pairs))}
	copy(out.pairs, q.pairs)
	return out
}

// extend appends all pairs from other to q.
func (q *QueryParams) extend(other *QueryParams) {
	if other == nil {
		return
	}
	q.pairs = append(q.pairs, other.pairs...)
}

// applyToURL appends the parameters to rawURL as a query string.
func (q *QueryParams) applyToURL(rawURL string) string {
	if len(q.pairs) == 0 {
		return rawURL
	}
	var b strings.Builder
	for i, p := range q.pairs {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(url.QueryEscape(p[0]))
		b.WriteByte('=')
		b.WriteString(url.QueryEscape(p[1]))
	}
	sep := "?"
	if strings.Contains(rawURL, "?") {
		sep = "&"
	}
	return rawURL + sep + b.String()
}

// ListOptions holds the standard offset/limit pagination shared by most list endpoints.
type ListOptions struct {
	// Offset is the number of items to skip from the beginning of the list.
	Offset *int64
	// Limit is the maximum number of items to return.
	Limit *int64
	// Desc, if true, returns items newest-first.
	Desc *bool
}

func (o ListOptions) apply(q *QueryParams) {
	q.AddInt("offset", o.Offset).AddInt("limit", o.Limit).AddBool("desc", o.Desc)
}

// StorageListOptions holds the options shared by the storage collection list endpoints
// (GET /v2/datasets, /v2/key-value-stores, /v2/request-queues), which add `unnamed` and
// `ownership` filters on top of the standard pagination.
type StorageListOptions struct {
	// Offset is the number of items to skip from the beginning of the list.
	Offset *int64
	// Limit is the maximum number of items to return.
	Limit *int64
	// Desc, if true, returns items newest-first.
	Desc *bool
	// Unnamed, if true, includes unnamed storages in the result.
	Unnamed *bool
	// Ownership filters by ownership (e.g. "OWNED" / "ACCESSIBLE").
	Ownership *string
}

func (o StorageListOptions) apply(q *QueryParams) {
	q.AddInt("offset", o.Offset).
		AddInt("limit", o.Limit).
		AddBool("desc", o.Desc).
		AddBool("unnamed", o.Unnamed).
		AddString("ownership", o.Ownership)
}

// Ptr returns a pointer to v. It is a convenience for setting the optional, pointer-typed
// fields on the option structs (e.g. Limit, Desc, My) without needing a named local variable:
//
//	client.Actors().List(ctx, apify.ActorListOptions{My: apify.Ptr(true), Limit: apify.Ptr(int64(10))})
func Ptr[T any](v T) *T { return &v }

// PaginationList is a single page of an offset/limit-paginated list.
//
// The pagination metadata (Total, Offset, Limit, Count, Desc) accompanies the Items slice.
// Note: Total reflects the API's reported total, which can briefly lag immediately after a
// write (e.g. right after PushItems) because the count is computed asynchronously — re-read
// after a short delay if you need an exact post-write total.
type PaginationList[T any] struct {
	// Total is the total number of items available across all pages.
	Total int64 `json:"total"`
	// Offset is the number of items skipped at the start.
	Offset int64 `json:"offset"`
	// Limit is the maximum number of items the API would return for this request.
	Limit int64 `json:"limit"`
	// Count is the number of items actually returned in this page.
	Count int64 `json:"count"`
	// Desc reports whether the items are in descending order.
	Desc bool `json:"desc"`
	// Items are the items of this page.
	Items []T `json:"items"`
	// scanned is the number of underlying rows the API scanned to produce this page, if the
	// endpoint reports one distinct from Count (currently only dataset items, via the
	// X-Apify-Pagination-Count header). Zero means "not reported"; [ListIterator] then falls
	// back to Count, which is correct for every endpoint that does not reshape its rows.
	//
	// Unexported and so never (un)marshaled: it is paging plumbing for ListIterator, not part
	// of this type's public, JSON-mirroring shape.
	scanned int64
}

// osTokenOverrides maps Go's runtime.GOOS values to the platform names the reference Apify JS
// client emits in the User-Agent OS token. That token comes from Node's os.platform(), which
// spells a few platforms differently from Go's runtime.GOOS:
//   - "windows" -> "win32"
//   - "solaris" -> "sunos", and Go's "illumos" (a Solaris/OpenSolaris derivative that Node/libuv
//     also reports as "sunos") -> "sunos"
//   - "ios" -> "darwin" (Node/libuv reports Apple platforms uniformly as "darwin")
//
// Every other GOOS value already matches os.platform() verbatim (e.g. "linux", "darwin",
// "android", "freebsd", "openbsd", "aix"), so it is passed through unchanged.
var osTokenOverrides = map[string]string{
	"windows": "win32",
	"solaris": "sunos",
	"illumos": "sunos",
	"ios":     "darwin",
}

// platformToken maps a Go GOOS value to the aligned User-Agent OS token (see osTokenOverrides).
// It is a pure function of its argument so the mapping can be unit-tested for every platform,
// not just the one the tests happen to run on.
func platformToken(goos string) string {
	if mapped, ok := osTokenOverrides[goos]; ok {
		return mapped
	}
	return goos
}

// osToken returns the User-Agent OS token for the current platform. runtime.GOOS is the idiomatic
// Go source of the platform identifier; platformToken aligns the two names that Go spells
// differently from the reference clients.
func osToken() string {
	return platformToken(runtime.GOOS)
}

// BuildUserAgent builds the User-Agent header value mandated by the client requirements:
// `ApifyClient/{version} ({os}; {language version}); isAtHome/{isAtHome}`.
//
// isAtHome is driven solely by the platform's APIFY_IS_AT_HOME environment variable (matching
// the requirements and the reference JS client, which reads it via @apify/consts) and is
// rendered lowercase (true/false). The {os} token uses osToken so it matches the platform names
// the reference clients emit.
func BuildUserAgent(suffix string, isAtHomeFn func() bool) string {
	atHome := "false"
	if isAtHomeFn() {
		atHome = "true"
	}
	ua := "ApifyClient/" + ClientVersion + " (" + osToken() + "; Go/" + goVersion() + "); isAtHome/" + atHome
	if suffix != "" {
		ua += "; " + suffix
	}
	return ua
}

// goVersion returns the runtime Go version (e.g. "1.24.7"), stripping the leading "go".
func goVersion() string {
	return strings.TrimPrefix(runtime.Version(), "go")
}

// toSafeID rewrites a resource id into the `username~resourcename` form Apify's URLs use
// (every `/` becomes `~`, matching the reference client's replaceAll) and validates the result
// as a URL path segment. See [safePathSegment] for what that rejects and why.
func toSafeID(id string) (string, error) {
	return safePathSegment(strings.ReplaceAll(id, "/", "~"))
}

// safePathSegment validates value as a single URL path segment and percent-encodes it. Used
// directly to validate a value interpolated into the path below the resource's own id (a
// record key, a request id), and by [toSafeID] for the id itself.
//
// An empty string or a dot segment ("." or "..") is rejected outright rather than encoded:
// percent-encoding alone is not enough to stop a path-traversal attempt, because a URL parser
// (or an intermediate proxy) can decode and resolve dot segments after this client has built
// the request path, collapsing e.g. "records/%2E%2E" to the parent endpoint exactly like a
// literal "records/..". Mirrors the reference client's toPathSegment.
func safePathSegment(value string) (string, error) {
	if value == "" {
		return "", errors.New("resource id must not be empty")
	}
	if value == "." || value == ".." {
		return "", fmt.Errorf("resource id must not be a dot segment, got %q", value)
	}
	return url.PathEscape(value), nil
}
