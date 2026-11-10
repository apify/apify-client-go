package apify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// parentRequestIDHeader carries the ID of the request that a Server Actor is handling.
const parentRequestIDHeader = "X-Actor-Request-ID"

// ErrNotServerActor is returned, wrapped, when an [ActorServerClient] method is used on an Actor
// that is not a Server Actor, which is an Actor without a server URL (`standbyUrl`).
var ErrNotServerActor = errors.New("the Actor is not a Server Actor")

// ActorServerClient calls the server of a Server Actor (formerly Actor Standby), a web server
// whose run the platform keeps ready, so a call skips the run startup. Get one with
// [ActorClient.Server]. It is safe for concurrent use.
type ActorServerClient struct {
	ctx     *resourceContext
	baseURL string
	id      string

	mu sync.Mutex
	// url is the server URL once it is fetched. Fetching it takes an API call for the Actor, so
	// it is kept for [ActorServerClient.GetURL] and [ActorServerClient.Request].
	url string
	// lookup is the API call in flight, which concurrent first calls wait for instead of
	// starting their own.
	lookup *serverURLLookup
}

// serverURLLookup is a fetch of the Actor that concurrent callers share.
type serverURLLookup struct {
	done  chan struct{}
	actor serverActor
	err   error
}

func newActorServerClient(hc *httpClient, baseURL, id string) *ActorServerClient {
	return &ActorServerClient{ctx: newSingleContext(hc, baseURL, "actors", id), baseURL: baseURL, id: id}
}

// serverActor is the part of the Actor that the server client reads.
type serverActor struct {
	StandbyURL   *string `json:"standbyUrl"`
	ActorStandby *struct {
		Build *string `json:"build"`
	} `json:"actorStandby"`
	TaggedBuilds map[string]struct {
		BuildID string `json:"buildId"`
	} `json:"taggedBuilds"`
}

// ActorServerRequestOptions are the options of [ActorServerClient.Request].
type ActorServerRequestOptions struct {
	// Method is the HTTP method of the request. Defaults to GET.
	Method string
	// Params are the query parameters of the request.
	Params url.Values
	// Headers are additional headers of the request. They reach the Actor's server as they are,
	// Authorization included, and win over the headers the client sets. A header with no values
	// removes the client's header of that name.
	Headers http.Header
	// Body is the request body. A string or []byte is sent as it is, without a Content-Type
	// unless Headers sets one. A [json.RawMessage] and any other value are sent as JSON, with
	// Content-Type application/json unless Headers sets one.
	Body any
	// ParentRequestID is the X-Actor-Request-ID header of the request that the calling Server
	// Actor is handling. It is sent on as the same header, so that the called Actor charges at the
	// pricing tier of that request, and it wins over the same header in Headers. Use a client
	// created with the run's APIFY_TOKEN.
	ParentRequestID string
	// Compression compresses a body of 1 KiB or more with the chosen codec. It is nil by
	// default, which sends the body uncompressed, since the Actor's server may not decode a
	// compressed body. Set it only if the server decodes the resulting Content-Encoding. A
	// Content-Encoding header in Headers sends the body untouched instead.
	Compression *RequestCompression
}

// ActorServerResponse is a response of an Actor's server, with its body read in full.
type ActorServerResponse struct {
	// StatusCode is the HTTP status of the response. Error statuses are returned, not raised.
	StatusCode int
	// Header holds the response headers.
	Header http.Header
	// Body is the response body.
	Body []byte
}

// JSON decodes the body of the response as JSON into v.
func (r *ActorServerResponse) JSON(v any) error {
	return json.Unmarshal(r.Body, v)
}

// ActorServerStream is a response of an Actor's server with its body unread.
type ActorServerStream struct {
	// StatusCode is the HTTP status of the response. Error statuses are returned, not raised.
	StatusCode int
	// Header holds the response headers.
	Header http.Header
	// Body is the unread response body. The caller must close it, which releases the
	// connection; until then the connection stays open.
	Body io.ReadCloser
}

// GetURL returns the URL of the Actor's server, which is its `standbyUrl`. The first call of this
// method, [ActorServerClient.GetOpenAPISpec] or [ActorServerClient.Request] fetches it, and later
// calls reuse it. Use it to call the server with another HTTP client.
//
// It returns an *[APIError] with IsNotFound true if the Actor does not exist, and an error
// wrapping [ErrNotServerActor] if it is not a Server Actor.
func (c *ActorServerClient) GetURL(ctx context.Context) (string, error) {
	actor, err := c.getServerActor(ctx)
	if err != nil {
		return "", err
	}
	return *actor.StandbyURL, nil
}

// getServerActor returns the Actor's server fields, fetching the Actor unless the server URL is
// already known. Concurrent first calls share one API request, and a failed lookup is not kept.
func (c *ActorServerClient) getServerActor(ctx context.Context) (serverActor, error) {
	for {
		c.mu.Lock()
		if c.url != "" {
			url := c.url
			c.mu.Unlock()
			return serverActor{StandbyURL: &url}, nil
		}
		if lookup := c.lookup; lookup != nil {
			c.mu.Unlock()
			select {
			case <-lookup.done:
			case <-ctx.Done():
				return serverActor{}, ctx.Err()
			}
			// The caller whose context ended did not fail for this caller's reasons.
			if lookup.err != nil && ctx.Err() == nil && isContextError(lookup.err) {
				continue
			}
			return lookup.actor, lookup.err
		}
		lookup := &serverURLLookup{done: make(chan struct{})}
		c.lookup = lookup
		c.mu.Unlock()

		lookup.actor, lookup.err = c.fetchServerActor(ctx)

		c.mu.Lock()
		if lookup.err == nil {
			c.url = *lookup.actor.StandbyURL
		}
		c.lookup = nil
		c.mu.Unlock()
		close(lookup.done)
		return lookup.actor, lookup.err
	}
}

// fetchServerActor fetches the Actor and checks that it is a Server Actor.
func (c *ActorServerClient) fetchServerActor(ctx context.Context) (serverActor, error) {
	actor, err := getResourceRequired[serverActor](ctx, c.ctx, "", NewQueryParams(), c.ctx.shortTimeout())
	if err != nil {
		return serverActor{}, err
	}
	if actor.StandbyURL == nil || *actor.StandbyURL == "" {
		return serverActor{}, fmt.Errorf("%w: %s", ErrNotServerActor, c.id)
	}
	if err := normalizeOptionalURLField("standbyUrl", &actor.StandbyURL); err != nil {
		return serverActor{}, err
	}
	return actor, nil
}

// buildNumberPattern matches a build number such as 0.1.23.
var buildNumberPattern = regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)$`)

// parseBuildNumber splits a build number like 0.1.23 into its version and sequence number, and
// reports false for anything else.
func parseBuildNumber(buildNumber string) ([3]int, bool) {
	match := buildNumberPattern.FindStringSubmatch(buildNumber)
	if match == nil {
		return [3]int{}, false
	}
	var parsed [3]int
	for i := range parsed {
		n, err := strconv.Atoi(match[i+1])
		if err != nil {
			return [3]int{}, false
		}
		parsed[i] = n
	}
	return parsed, true
}

// isOlderBuild reports whether buildNumber belongs to the same version as target and comes
// before it. Builds are listed newest first and the sequence numbers of a version grow over
// time, so meeting such a build means that the target does not exist.
func isOlderBuild(buildNumber string, target [3]int) bool {
	parsed, ok := parseBuildNumber(buildNumber)
	return ok && parsed[0] == target[0] && parsed[1] == target[1] && parsed[2] < target[2]
}

// GetOpenAPISpec returns the OpenAPI specification of the Actor's server, or nil if the Actor's
// definition has none. It is the web server schema of the build that serves the server: the
// `actorStandby.build` of the Actor, or `latest`.
//
// It returns the errors of [ActorServerClient.GetURL], and an error if that build does not exist.
func (c *ActorServerClient) GetOpenAPISpec(ctx context.Context) (json.RawMessage, error) {
	actor, err := c.fetchSpecActor(ctx)
	if err != nil {
		return nil, err
	}
	buildRef := "latest"
	if actor.ActorStandby != nil && actor.ActorStandby.Build != nil && *actor.ActorStandby.Build != "" {
		buildRef = *actor.ActorStandby.Build
	}

	buildID := actor.TaggedBuilds[buildRef].BuildID
	if target, isNumber := parseBuildNumber(buildRef); buildID == "" && isNumber {
		// Not a tag, so look buildRef up as a build number, which the API cannot resolve directly.
		buildID, err = c.findBuildID(ctx, buildRef, target)
		if err != nil {
			return nil, err
		}
	}

	var build Build
	var found bool
	if buildID != "" {
		build, found, err = newBuildClient(c.ctx.http, c.baseURL, buildID).Get(ctx)
		if err != nil {
			return nil, err
		}
	}
	if !found {
		return nil, fmt.Errorf("build %s of Actor %s does not exist", buildRef, c.id)
	}

	var definition struct {
		WebServerSchema json.RawMessage `json:"webServerSchema"`
	}
	if raw, ok := build.Extra["actorDefinition"]; ok {
		if err := json.Unmarshal(raw, &definition); err != nil {
			return nil, err
		}
	}
	if string(definition.WebServerSchema) == "null" {
		return nil, nil
	}
	return definition.WebServerSchema, nil
}

// fetchSpecActor fetches the Actor for the OpenAPI specification lookup, which needs more of it
// than the server URL, and remembers the server URL.
func (c *ActorServerClient) fetchSpecActor(ctx context.Context) (serverActor, error) {
	actor, err := c.fetchServerActor(ctx)
	if err != nil {
		return serverActor{}, err
	}
	c.mu.Lock()
	c.url = *actor.StandbyURL
	c.mu.Unlock()
	return actor, nil
}

// findBuildID looks buildNumber up in the builds of the Actor, newest first.
func (c *ActorServerClient) findBuildID(ctx context.Context, buildNumber string, target [3]int) (string, error) {
	builds := newBuildCollectionClientWithBase(c.ctx.http, c.ctx.subURL(""), "builds")
	it := builds.Iterate(ListOptions{Desc: Ptr(true)}, nil)
	for {
		build, err := it.Next(ctx)
		if err != nil {
			return "", err
		}
		if build == nil {
			return "", nil
		}
		if build.BuildNumber == buildNumber {
			return build.ID, nil
		}
		if isOlderBuild(build.BuildNumber, target) {
			return "", nil
		}
	}
}

// httpMethods are the methods [ActorServerClient.Request] accepts.
var httpMethods = map[string]bool{
	http.MethodGet: true, http.MethodHead: true, http.MethodPost: true, http.MethodPut: true,
	http.MethodDelete: true, http.MethodConnect: true, http.MethodOptions: true,
	http.MethodTrace: true, http.MethodPatch: true,
}

// Request sends an HTTP request to the Actor's server and reads the whole response.
//
// path is the path on the server, for example "/search", and a fragment is dropped. It is resolved
// against the server URL, so the request cannot go to another host. Pass the query in Params.
//
// The request carries the token in the X-Apify-Authorization header, which the platform removes
// before the Actor sees it. The Headers of the options reach the server as they are.
//
// The response is returned whatever its status, so check StatusCode. A 429 or 503 response is
// retried, and the last one is returned when the retries run out. A redirect is returned as it
// is, so the token never reaches another host. A transport error, a timeout included, is not
// retried, since the Actor may have handled the request. The first request to an idle Actor waits
// for a run to start, so the request uses the "long" timeout tier, which ctx can shorten.
//
// It returns an error only when there is no response: the errors of [ActorServerClient.GetURL],
// invalid options, or a failed or timed-out request. To read the body incrementally, use
// [ActorServerClient.RequestStream].
func (c *ActorServerClient) Request(ctx context.Context, path string, options ActorServerRequestOptions) (*ActorServerResponse, error) {
	resp, err := c.send(ctx, path, options, false)
	if err != nil {
		return nil, err
	}
	return &ActorServerResponse{StatusCode: resp.statusCode, Header: resp.header, Body: resp.body}, nil
}

// RequestStream is [ActorServerClient.Request], but returns the response body unread, so that a
// large or long-lived body can be processed in chunks. The caller must close the returned
// stream's Body. The long timeout tier bounds only the wait for the server to start
// responding.
func (c *ActorServerClient) RequestStream(ctx context.Context, path string, options ActorServerRequestOptions) (*ActorServerStream, error) {
	resp, err := c.send(ctx, path, options, true)
	if err != nil {
		return nil, err
	}
	return &ActorServerStream{StatusCode: resp.statusCode, Header: resp.header, Body: resp.stream}, nil
}

// send validates the options, resolves the server URL and sends the request.
func (c *ActorServerClient) send(ctx context.Context, path string, options ActorServerRequestOptions, stream bool) (*serverResponse, error) {
	method := strings.ToUpper(options.Method)
	if method == "" {
		method = http.MethodGet
	}
	if !httpMethods[method] || (options.Method != "" && options.Method != method) {
		return nil, fmt.Errorf("ActorServerRequestOptions.Method: unknown HTTP method %q", options.Method)
	}
	body, contentType, err := encodeServerBody(options.Body)
	if err != nil {
		return nil, err
	}
	var compressors []requestCompressor
	if options.Compression != nil {
		compressors = compressorsFor(*options.Compression)
	}

	serverURL, err := c.GetURL(ctx)
	if err != nil {
		return nil, err
	}
	headers := options.Headers
	if options.ParentRequestID != "" {
		headers = withHeader(headers, parentRequestIDHeader, options.ParentRequestID)
	}
	return c.ctx.http.callServer(ctx, serverCall{
		method:      method,
		url:         serverRequestURL(serverURL, path, options.Params),
		body:        body,
		contentType: contentType,
		headers:     headers,
		compressors: compressors,
		timeout:     c.ctx.longTimeout(),
		stream:      stream,
	})
}

// withHeader returns a copy of headers with name set to value, replacing the header of that
// name whatever the casing of either.
func withHeader(headers http.Header, name, value string) http.Header {
	result := http.Header{}
	for key, values := range headers {
		if !strings.EqualFold(key, name) {
			result[key] = values
		}
	}
	result.Set(name, value)
	return result
}

// serverRequestURL joins the server URL, the path and the query. Leading slashes of the path are
// dropped and a fragment is cut, so whatever the path looks like (//host, @host, an absolute URL)
// it stays a path on the server.
func serverRequestURL(serverURL, path string, params url.Values) string {
	path, _, _ = strings.Cut(path, "#")
	result := strings.TrimRight(serverURL, "/") + "/" + strings.TrimLeft(path, "/")
	if len(params) == 0 {
		return result
	}
	separator := "?"
	if strings.Contains(result, "?") {
		separator = "&"
	}
	return result + separator + params.Encode()
}

// encodeServerBody serializes a request body of [ActorServerRequestOptions.Body] and returns it
// with its Content-Type, which is empty for a body that is sent as it is.
func encodeServerBody(body any) ([]byte, string, error) {
	switch b := body.(type) {
	case nil:
		return nil, "", nil
	case json.RawMessage:
		return b, contentTypeJSON, nil
	case []byte:
		return b, "", nil
	case string:
		return []byte(b), "", nil
	default:
		encoded, err := json.Marshal(body)
		return encoded, contentTypeJSON, err
	}
}

// Server returns the [ActorServerClient] of this Actor. Every call returns the same client.
func (c *ActorClient) Server() *ActorServerClient {
	c.serverOnce.Do(func() {
		c.server = newActorServerClient(c.ctx.http, c.baseURL, c.id)
	})
	return c.server
}

// Standby returns a client for the server interface of this Actor in Standby mode.
//
// Deprecated: Actor Standby is the former name of Server Actors. Use [ActorClient.Server].
func (c *ActorClient) Standby() *ActorServerClient {
	return c.Server()
}
