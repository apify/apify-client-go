package apify

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/klauspost/compress/zstd"
)

const (
	serverTestActorID = "test-actor-id"
	serverTestPrefix  = "/server"
)

// receivedRequest is a request that the test server got.
type receivedRequest struct {
	method string
	path   string
	query  url.Values
	header http.Header
	body   []byte
}

// serverFixture is an HTTP server that plays both the Apify API (the Actor, its builds) and the
// Actor's own server (everything under /server), and records every request.
type serverFixture struct {
	t      *testing.T
	server *httptest.Server

	mu       sync.Mutex
	received []receivedRequest
	// actor is the Actor the API returns; nil answers 404.
	actor map[string]any
	// builds maps a build ID to its build number, in the order the builds list returns them.
	buildIDs      []string
	buildNumbers  map[string]string
	buildsPerPage int
	// definition is the actorDefinition of every build the API returns.
	definition any
	// flaky is the statuses /server/flaky answers with, one per request, before it answers 200.
	flaky []int
}

func newServerFixture(t *testing.T) *serverFixture {
	f := &serverFixture{t: t, buildNumbers: map[string]string{}, definition: map[string]any{"webServerSchema": map[string]any{"openapi": "3.0.0"}}}
	f.server = httptest.NewServer(http.HandlerFunc(f.handle))
	t.Cleanup(f.server.Close)
	f.setActor(nil)
	return f
}

func (f *serverFixture) serverURL() string { return f.server.URL + serverTestPrefix }

// setActor makes the API return a Server Actor with the given extra fields.
func (f *serverFixture) setActor(fields map[string]any) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.actor = map[string]any{
		"id": serverTestActorID, "name": "test-actor", "standbyUrl": f.serverURL(),
		"actorStandby": map[string]any{"isEnabled": true},
		"taggedBuilds": map[string]any{
			"latest": map[string]any{"buildId": "latest-build-id", "buildNumber": "0.0.2"},
			"beta":   map[string]any{"buildId": "beta-build-id", "buildNumber": "0.0.3"},
		},
	}
	for k, v := range fields {
		f.actor[k] = v
	}
}

func (f *serverFixture) noActor() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.actor = nil
}

func (f *serverFixture) requests(prefix string) []receivedRequest {
	f.mu.Lock()
	defer f.mu.Unlock()
	var out []receivedRequest
	for _, r := range f.received {
		if strings.HasPrefix(r.path, prefix) {
			out = append(out, r)
		}
	}
	return out
}

func (f *serverFixture) handle(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	f.mu.Lock()
	f.received = append(f.received, receivedRequest{r.Method, r.URL.Path, r.URL.Query(), r.Header.Clone(), body})
	actor, definition := f.actor, f.definition
	var flakyStatus int
	if r.URL.Path == serverTestPrefix+"/flaky" && len(f.flaky) > 0 {
		flakyStatus, f.flaky = f.flaky[0], f.flaky[1:]
	}
	f.mu.Unlock()

	writeJSON := func(status int, payload any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(payload)
	}
	actorPath := "/v2/actors/" + serverTestActorID
	switch {
	case r.URL.Path == actorPath:
		if actor == nil {
			writeJSON(404, map[string]any{"error": map[string]any{"type": "record-not-found", "message": "Actor was not found"}})
			return
		}
		writeJSON(200, map[string]any{"data": actor})
	case r.URL.Path == actorPath+"/builds":
		f.writeBuilds(w, r, writeJSON)
	case strings.HasPrefix(r.URL.Path, "/v2/actor-builds/"):
		id := strings.TrimPrefix(r.URL.Path, "/v2/actor-builds/")
		writeJSON(200, map[string]any{"data": map[string]any{"id": id, "actId": serverTestActorID, "status": "SUCCEEDED", "actorDefinition": definition}})
	case flakyStatus != 0:
		writeJSON(flakyStatus, map[string]any{"error": "flaky"})
	case r.URL.Path == serverTestPrefix+"/hang-up":
		conn, _, _ := w.(http.Hijacker).Hijack()
		_ = conn.Close()
	case r.URL.Path == serverTestPrefix+"/missing":
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(404)
		_, _ = w.Write([]byte("Not here"))
	case r.URL.Path == serverTestPrefix+"/slow":
		time.Sleep(300 * time.Millisecond)
		writeJSON(200, map[string]any{"ok": true})
	default:
		writeJSON(200, map[string]any{"ok": true})
	}
}

func (f *serverFixture) writeBuilds(w http.ResponseWriter, r *http.Request, writeJSON func(int, any)) {
	f.mu.Lock()
	ids, numbers, perPage := f.buildIDs, f.buildNumbers, f.buildsPerPage
	f.mu.Unlock()
	offset := 0
	if o := r.URL.Query().Get("offset"); o != "" {
		for _, c := range o {
			offset = offset*10 + int(c-'0')
		}
	}
	var items []map[string]any
	for _, id := range ids {
		items = append(items, map[string]any{"id": id, "buildNumber": numbers[id], "status": "SUCCEEDED"})
	}
	end := len(items)
	if perPage > 0 && offset+perPage < end {
		end = offset + perPage
	}
	page := []map[string]any{}
	if offset < end {
		page = items[offset:end]
	}
	writeJSON(200, map[string]any{"data": map[string]any{"total": len(items), "offset": offset, "count": len(page), "limit": 1000, "items": page}})
}

func (f *serverFixture) client(options ...Option) *ApifyClient {
	base := []Option{WithBaseURL(f.server.URL), WithToken("test-token"), WithMaxRetries(0)}
	return NewClient(append(base, options...)...)
}

func (f *serverFixture) setBuilds(numbersByID map[string]string, order ...string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.buildIDs, f.buildNumbers = order, numbersByID
}

func TestActorServerRequestSendsMethodPathQueryBodyAndToken(t *testing.T) {
	f := newServerFixture(t)
	resp, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/search", ActorServerRequestOptions{
		Method: "POST", Params: url.Values{"q": {"apify"}}, Body: map[string]any{"limit": 5},
	})
	if err != nil {
		t.Fatalf("Request: %v", err)
	}
	var payload map[string]bool
	if err := resp.JSON(&payload); err != nil || resp.StatusCode != 200 || !payload["ok"] {
		t.Fatalf("response: %d %s %v", resp.StatusCode, resp.Body, err)
	}
	reqs := f.requests(serverTestPrefix + "/")
	if len(reqs) != 1 {
		t.Fatalf("expected one server request, got %d", len(reqs))
	}
	r := reqs[0]
	if r.method != "POST" || r.path != serverTestPrefix+"/search" || r.query.Get("q") != "apify" {
		t.Fatalf("request: %s %s %v", r.method, r.path, r.query)
	}
	if string(r.body) != `{"limit":5}` || r.header.Get("Content-Type") != "application/json" {
		t.Fatalf("body %q, content type %q", r.body, r.header.Get("Content-Type"))
	}
	if r.header.Get("X-Apify-Authorization") != "Bearer test-token" || r.header.Get("Authorization") != "" {
		t.Fatalf("token headers: %v", r.header)
	}
	if !strings.HasPrefix(r.header.Get("User-Agent"), "ApifyClient/") {
		t.Fatalf("User-Agent: %q", r.header.Get("User-Agent"))
	}
}

func TestActorServerRequestSendsTheParentRequestIDOverTheSameHeader(t *testing.T) {
	f := newServerFixture(t)
	_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{
		Headers:         http.Header{"x-actor-request-id": {"stale"}, "X-Other": {"kept"}},
		ParentRequestID: "parent-id",
	})
	if err != nil {
		t.Fatal(err)
	}
	r := f.requests(serverTestPrefix + "/")[0]
	if got := r.header.Values("X-Actor-Request-ID"); len(got) != 1 || got[0] != "parent-id" {
		t.Fatalf("X-Actor-Request-ID: %v", got)
	}
	if r.header.Get("X-Other") != "kept" {
		t.Fatalf("other headers must stay: %v", r.header)
	}
}

func TestActorServerRequestSendsNoRequestIDWithoutAParentRequestID(t *testing.T) {
	f := newServerFixture(t)
	if _, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{}); err != nil {
		t.Fatal(err)
	}
	if got := f.requests(serverTestPrefix + "/")[0].header.Get("X-Actor-Request-ID"); got != "" {
		t.Fatalf("X-Actor-Request-ID: %q", got)
	}
}

func TestActorServerRequestSendsTheCallersAuthorizationAsItIs(t *testing.T) {
	f := newServerFixture(t)
	headers := http.Header{"Authorization": {"Bearer actor-token"}}
	if _, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{Headers: headers}); err != nil {
		t.Fatal(err)
	}
	r := f.requests(serverTestPrefix + "/")[0]
	if r.header.Get("Authorization") != "Bearer actor-token" || r.header.Get("X-Apify-Authorization") != "Bearer test-token" {
		t.Fatalf("headers: %v", r.header)
	}
}

// A redirect of the server is returned as it is: following it would send the token to a host the
// Actor's code picked.
func TestActorServerRequestDoesNotFollowRedirects(t *testing.T) {
	var reached atomic.Bool
	other := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { reached.Store(true) }))
	defer other.Close()
	redirecting := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, other.URL+"/", http.StatusFound)
	}))
	defer redirecting.Close()

	f := newServerFixture(t)
	f.setActor(map[string]any{"standbyUrl": redirecting.URL})
	resp, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusFound || resp.Header.Get("Location") != other.URL+"/" {
		t.Fatalf("expected the redirect as it is, got %d %v", resp.StatusCode, resp.Header)
	}
	if reached.Load() {
		t.Fatal("the redirect target was reached")
	}
}

// redirectCheckingBackend is the custom backend of the README: a plain *http.Client that honors
// FollowsRedirects.
type redirectCheckingBackend struct{ inner *http.Client }

func (b *redirectCheckingBackend) Do(req *http.Request) (*http.Response, error) {
	if FollowsRedirects(req) {
		return b.inner.Do(req)
	}
	client := *b.inner
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	return client.Do(req)
}

// A custom backend that checks FollowsRedirects, as the README's example does, keeps the token
// away from the target of a redirect too.
func TestActorServerRequestDoesNotFollowRedirectsThroughACustomBackend(t *testing.T) {
	var tokenReceived atomic.Bool
	other := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Apify-Authorization") != "" {
			tokenReceived.Store(true)
		}
	}))
	defer other.Close()
	redirecting := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, other.URL+"/", http.StatusFound)
	}))
	defer redirecting.Close()

	f := newServerFixture(t)
	f.setActor(map[string]any{"standbyUrl": redirecting.URL})
	client := f.client(WithHTTPBackend(&redirectCheckingBackend{inner: http.DefaultClient}))
	resp, err := client.Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{})
	if err != nil || resp.StatusCode != http.StatusFound {
		t.Fatalf("expected the redirect as it is, got %+v, %v", resp, err)
	}
	if tokenReceived.Load() {
		t.Fatal("the token reached the redirect target")
	}
}

func TestActorServerClientFetchesTheServerURLOnce(t *testing.T) {
	f := newServerFixture(t)
	actor := f.client().Actor(serverTestActorID)
	ctx := context.Background()

	for i := 0; i < 2; i++ {
		got, err := actor.Server().GetURL(ctx)
		if err != nil || got != f.serverURL() {
			t.Fatalf("GetURL = %q, %v", got, err)
		}
	}
	if _, err := actor.Server().Request(ctx, "", ActorServerRequestOptions{}); err != nil {
		t.Fatal(err)
	}
	if n := len(f.requests("/v2/actors/"+serverTestActorID)) - len(f.requests("/v2/actors/"+serverTestActorID+"/")); n != 1 {
		t.Fatalf("expected the Actor to be fetched once, got %d", n)
	}
	first, second := actor.Server(), actor.Server()
	if first != second {
		t.Fatal("Server() must return the same client every time")
	}
}

func TestActorServerClientSharesOneLookupBetweenConcurrentCalls(t *testing.T) {
	f := newServerFixture(t)
	server := f.client().Actor(serverTestActorID).Server()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got, err := server.GetURL(context.Background()); err != nil || got != f.serverURL() {
				t.Errorf("GetURL = %q, %v", got, err)
			}
		}()
	}
	wg.Wait()
	if n := len(f.requests("/v2/actors/" + serverTestActorID)); n != 1 {
		t.Fatalf("expected one Actor fetch, got %d", n)
	}
}

func TestActorServerClientLooksTheActorUpAgainAfterAFailure(t *testing.T) {
	f := newServerFixture(t)
	server := f.client().Actor(serverTestActorID).Server()
	f.noActor()
	_, err := server.GetURL(context.Background())
	if apiErr, ok := AsAPIError(err); !ok || !apiErr.IsNotFound() {
		t.Fatalf("expected a not-found API error, got %v", err)
	}
	f.setActor(nil)
	if got, err := server.GetURL(context.Background()); err != nil || got != f.serverURL() {
		t.Fatalf("GetURL after the failure = %q, %v", got, err)
	}
	if n := len(f.requests("/v2/actors/" + serverTestActorID)); n != 2 {
		t.Fatalf("expected two Actor fetches, got %d", n)
	}
}

func TestActorServerClientRejectsAnActorThatIsNotAServerActor(t *testing.T) {
	f := newServerFixture(t)
	f.setActor(map[string]any{"standbyUrl": nil})
	server := f.client().Actor(serverTestActorID).Server()
	ctx := context.Background()
	_, urlErr := server.GetURL(ctx)
	_, specErr := server.GetOpenAPISpec(ctx)
	_, reqErr := server.Request(ctx, "/", ActorServerRequestOptions{})
	for name, err := range map[string]error{"GetURL": urlErr, "GetOpenAPISpec": specErr, "Request": reqErr} {
		if !errors.Is(err, ErrNotServerActor) {
			t.Errorf("%s: expected ErrNotServerActor, got %v", name, err)
		}
	}
	if n := len(f.requests(serverTestPrefix + "/")); n != 0 {
		t.Fatalf("no request may reach the server, got %d", n)
	}
}

// Whatever the path looks like, it is a path on the server, never another host.
func TestActorServerRequestKeepsEveryPathOnTheServer(t *testing.T) {
	for _, path := range []string{"//evil.example/steal", "@evil.example/steal", "https://evil.example/steal"} {
		f := newServerFixture(t)
		resp, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), path, ActorServerRequestOptions{})
		if err != nil || resp.StatusCode != 200 {
			t.Fatalf("%s: %v", path, err)
		}
		if n := len(f.requests(serverTestPrefix + "/")); n != 1 {
			t.Fatalf("%s: expected one request on the server, got %d", path, n)
		}
	}
}

func TestActorServerRequestDropsAFragmentOfThePath(t *testing.T) {
	f := newServerFixture(t)
	_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/search#results", ActorServerRequestOptions{Params: url.Values{"q": {"apify"}}})
	if err != nil {
		t.Fatal(err)
	}
	r := f.requests(serverTestPrefix + "/")[0]
	if r.path != serverTestPrefix+"/search" || r.query.Get("q") != "apify" {
		t.Fatalf("request: %s %v", r.path, r.query)
	}
}

func TestActorServerRequestCompression(t *testing.T) {
	payload := map[string]any{"items": strings.Repeat("x", 5000)}
	t.Run("uncompressed by default", func(t *testing.T) {
		f := newServerFixture(t)
		if _, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/items", ActorServerRequestOptions{Method: "POST", Body: payload}); err != nil {
			t.Fatal(err)
		}
		r := f.requests(serverTestPrefix + "/")[0]
		if r.header.Get("Content-Encoding") != "" || !bytes.Contains(r.body, []byte(`"items"`)) {
			t.Fatalf("expected the body as it is, got encoding %q", r.header.Get("Content-Encoding"))
		}
	})
	t.Run("gzip on request", func(t *testing.T) {
		f := newServerFixture(t)
		_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/items", ActorServerRequestOptions{
			Method: "POST", Body: payload, Compression: Ptr(RequestCompressionGzip),
		})
		if err != nil {
			t.Fatal(err)
		}
		r := f.requests(serverTestPrefix + "/")[0]
		zr, err := gzip.NewReader(bytes.NewReader(r.body))
		if err != nil || r.header.Get("Content-Encoding") != "gzip" {
			t.Fatalf("expected a gzip body, got encoding %q: %v", r.header.Get("Content-Encoding"), err)
		}
		plain, _ := io.ReadAll(zr)
		if !bytes.Contains(plain, []byte(`"items"`)) {
			t.Fatalf("decompressed body: %q", plain[:20])
		}
	})
	t.Run("zstd on request", func(t *testing.T) {
		f := newServerFixture(t)
		_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/items", ActorServerRequestOptions{
			Method: "POST", Body: payload, Compression: Ptr(RequestCompressionZstd),
		})
		if err != nil {
			t.Fatal(err)
		}
		r := f.requests(serverTestPrefix + "/")[0]
		zr, err := zstd.NewReader(bytes.NewReader(r.body))
		if err != nil || r.header.Get("Content-Encoding") != "zstd" {
			t.Fatalf("expected a zstd body, got encoding %q: %v", r.header.Get("Content-Encoding"), err)
		}
		defer zr.Close()
		plain, _ := io.ReadAll(zr)
		if !bytes.Contains(plain, []byte(`"items"`)) {
			t.Fatalf("decompressed body: %q", plain)
		}
	})
	t.Run("a Content-Type of the caller that is already compressed is not compressed again", func(t *testing.T) {
		f := newServerFixture(t)
		body := bytes.Repeat([]byte("x"), 5000)
		_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/items", ActorServerRequestOptions{
			Method: "POST", Body: body, Headers: http.Header{"Content-Type": {"image/png"}}, Compression: Ptr(RequestCompressionGzip),
		})
		if err != nil {
			t.Fatal(err)
		}
		r := f.requests(serverTestPrefix + "/")[0]
		if r.header.Get("Content-Encoding") != "" || !bytes.Equal(r.body, body) {
			t.Fatalf("expected the body as it is, got encoding %q and %d bytes", r.header.Get("Content-Encoding"), len(r.body))
		}
	})
	t.Run("lower-case header names of the caller count too", func(t *testing.T) {
		f := newServerFixture(t)
		body := bytes.Repeat([]byte("x"), 5000)
		server := f.client().Actor(serverTestActorID).Server()
		ctx := context.Background()
		_, err := server.Request(ctx, "/items", ActorServerRequestOptions{
			Method: "POST", Body: body, Headers: http.Header{"content-type": {"image/png"}}, Compression: Ptr(RequestCompressionGzip),
		})
		if err != nil {
			t.Fatal(err)
		}
		_, err = server.Request(ctx, "/items", ActorServerRequestOptions{
			Method: "POST", Body: body, Headers: http.Header{"content-encoding": {"br"}}, Compression: Ptr(RequestCompressionGzip),
		})
		if err != nil {
			t.Fatal(err)
		}
		requests := f.requests(serverTestPrefix + "/")
		if len(requests) != 2 {
			t.Fatalf("expected 2 requests, got %d", len(requests))
		}
		for i, r := range requests {
			wantEncoding := []string{"", "br"}[i]
			if r.header.Get("Content-Encoding") != wantEncoding || !bytes.Equal(r.body, body) {
				t.Errorf("request %d: expected the body as it is under encoding %q, got encoding %q and %d bytes", i, wantEncoding, r.header.Get("Content-Encoding"), len(r.body))
			}
		}
	})
	t.Run("a Content-Encoding header of the caller without values leaves the body untouched too", func(t *testing.T) {
		f := newServerFixture(t)
		body := bytes.Repeat([]byte("x"), 5000)
		_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/items", ActorServerRequestOptions{
			Method: "POST", Body: body, Headers: http.Header{"Content-Encoding": {}}, Compression: Ptr(RequestCompressionGzip),
		})
		if err != nil {
			t.Fatal(err)
		}
		requests := f.requests(serverTestPrefix + "/")
		if len(requests) != 1 {
			t.Fatalf("expected 1 request, got %d", len(requests))
		}
		if r := requests[0]; r.header.Get("Content-Encoding") != "" || !bytes.Equal(r.body, body) {
			t.Fatalf("expected the body as it is, got encoding %q and %d bytes", r.header.Get("Content-Encoding"), len(r.body))
		}
	})
	t.Run("a Content-Encoding header of the caller leaves the body untouched", func(t *testing.T) {
		f := newServerFixture(t)
		body := bytes.Repeat([]byte("x"), 5000)
		_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/items", ActorServerRequestOptions{
			Method: "POST", Body: body, Headers: http.Header{"Content-Encoding": {"gzip"}}, Compression: Ptr(RequestCompressionGzip),
		})
		if err != nil {
			t.Fatal(err)
		}
		r := f.requests(serverTestPrefix + "/")[0]
		if r.header.Get("Content-Encoding") != "gzip" || !bytes.Equal(r.body, body) {
			t.Fatalf("expected the body untouched under gzip, got %d bytes", len(r.body))
		}
	})
}

// A header of the caller with no values removes the header the client would set.
func TestActorServerRequestRemovesAHeaderOfTheClient(t *testing.T) {
	f := newServerFixture(t)
	headers := http.Header{"x-apify-authorization": nil, "User-Agent": {}}
	if _, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{Headers: headers}); err != nil {
		t.Fatal(err)
	}
	requests := f.requests(serverTestPrefix + "/")
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
	if token := requests[0].header.Get("X-Apify-Authorization"); token != "" {
		t.Fatalf("the token header must be removed, got %q", token)
	}
	if strings.HasPrefix(requests[0].header.Get("User-Agent"), "ApifyClient/") {
		t.Fatalf("the User-Agent must be removed, got %q", requests[0].header.Get("User-Agent"))
	}
}

func TestActorServerRequestBodyForms(t *testing.T) {
	f := newServerFixture(t)
	server := f.client().Actor(serverTestActorID).Server()
	ctx := context.Background()
	for _, body := range []any{"text", []byte("bytes"), json.RawMessage(`{"a":1}`)} {
		if _, err := server.Request(ctx, "/", ActorServerRequestOptions{Method: "POST", Body: body}); err != nil {
			t.Fatal(err)
		}
	}
	reqs := f.requests(serverTestPrefix + "/")
	wantBodies := []string{"text", "bytes", `{"a":1}`}
	wantTypes := []string{"", "", "application/json"}
	for i, r := range reqs {
		if string(r.body) != wantBodies[i] || r.header.Get("Content-Type") != wantTypes[i] {
			t.Errorf("request %d: body %q, content type %q", i, r.body, r.header.Get("Content-Type"))
		}
	}
}

func TestActorServerRequestRejectsAnUnknownMethodWithoutSendingARequest(t *testing.T) {
	f := newServerFixture(t)
	for _, method := range []string{"get", "FETCH"} {
		_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{Method: method})
		if err == nil {
			t.Fatalf("%q: expected an error", method)
		}
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.received) != 0 {
		t.Fatalf("expected no request, got %d", len(f.received))
	}
}

func TestActorServerRequestReturnsAnErrorResponseAsItIs(t *testing.T) {
	f := newServerFixture(t)
	resp, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/missing", ActorServerRequestOptions{})
	if err != nil || resp.StatusCode != 404 || string(resp.Body) != "Not here" {
		t.Fatalf("response: %+v, %v", resp, err)
	}
}

func TestActorServerRequestReturnsNotFoundForAMissingActor(t *testing.T) {
	f := newServerFixture(t)
	f.noActor()
	_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{})
	if apiErr, ok := AsAPIError(err); !ok || !apiErr.IsNotFound() {
		t.Fatalf("expected a not-found API error, got %v", err)
	}
}

func TestActorServerRequestRetriesOnlyRateLimitsAndServiceUnavailable(t *testing.T) {
	retrying := func(f *serverFixture, retries int) *ActorServerClient {
		return f.client(WithMaxRetries(retries), WithMinDelayBetweenRetries(time.Millisecond)).Actor(serverTestActorID).Server()
	}
	for _, status := range []int{429, 503} {
		f := newServerFixture(t)
		f.flaky = []int{status, status}
		resp, err := retrying(f, 2).Request(context.Background(), "/flaky", ActorServerRequestOptions{})
		if err != nil || resp.StatusCode != 200 {
			t.Fatalf("%d: %+v, %v", status, resp, err)
		}
		if n := len(f.requests(serverTestPrefix + "/")); n != 3 {
			t.Fatalf("%d: expected 3 requests, got %d", status, n)
		}
	}
	t.Run("the last response is returned once the retries run out", func(t *testing.T) {
		f := newServerFixture(t)
		f.flaky = []int{503, 503, 503}
		resp, err := retrying(f, 1).Request(context.Background(), "/flaky", ActorServerRequestOptions{Method: "POST", Body: map[string]int{"a": 1}})
		if err != nil || resp.StatusCode != 503 {
			t.Fatalf("%+v, %v", resp, err)
		}
		reqs := f.requests(serverTestPrefix + "/")
		if len(reqs) != 2 || string(reqs[0].body) != string(reqs[1].body) {
			t.Fatalf("expected 2 requests with the same body, got %d", len(reqs))
		}
	})
	t.Run("another server error is not retried", func(t *testing.T) {
		f := newServerFixture(t)
		f.flaky = []int{500}
		resp, err := retrying(f, 2).Request(context.Background(), "/flaky", ActorServerRequestOptions{})
		if err != nil || resp.StatusCode != 500 {
			t.Fatalf("%+v, %v", resp, err)
		}
		if n := len(f.requests(serverTestPrefix + "/")); n != 1 {
			t.Fatalf("expected 1 request, got %d", n)
		}
	})
}

// The Actor may have handled a request that failed in transit, so repeating it is not safe.
func TestActorServerRequestDoesNotRetryATransportError(t *testing.T) {
	f := newServerFixture(t)
	server := f.client(WithMaxRetries(2), WithMinDelayBetweenRetries(time.Millisecond)).Actor(serverTestActorID).Server()
	_, err := server.Request(context.Background(), "/hang-up", ActorServerRequestOptions{Method: "POST", Body: map[string]int{"a": 1}})
	if err == nil {
		t.Fatal("expected an error")
	}
	if n := len(f.requests(serverTestPrefix + "/")); n != 1 {
		t.Fatalf("expected 1 request, got %d", n)
	}
}

func TestActorServerRequestTimesOutWithoutRetrying(t *testing.T) {
	f := newServerFixture(t)
	server := f.client(WithTimeoutLong(50*time.Millisecond), WithMaxRetries(3), WithMinDelayBetweenRetries(time.Millisecond)).Actor(serverTestActorID).Server()
	_, err := server.Request(context.Background(), "/slow", ActorServerRequestOptions{})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected a deadline error, got %v", err)
	}
	if n := len(f.requests(serverTestPrefix + "/")); n != 1 {
		t.Fatalf("expected 1 request, got %d", n)
	}
}

func TestActorServerRequestStreamReturnsTheBodyUnread(t *testing.T) {
	f := newServerFixture(t)
	stream, err := f.client().Actor(serverTestActorID).Server().RequestStream(context.Background(), "/", ActorServerRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = stream.Body.Close() }()
	body, err := io.ReadAll(stream.Body)
	if err != nil || stream.StatusCode != 200 || !strings.Contains(string(body), `"ok":true`) {
		t.Fatalf("stream: %d %q %v", stream.StatusCode, body, err)
	}
}

func TestActorServerGetOpenAPISpec(t *testing.T) {
	spec := `{"openapi":"3.0.0"}`
	cases := []struct {
		name        string
		fields      map[string]any
		wantBuildID string
	}{
		{"no build set uses latest", nil, "latest-build-id"},
		{"a build tag", map[string]any{"actorStandby": map[string]any{"isEnabled": true, "build": "beta"}}, "beta-build-id"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newServerFixture(t)
			f.setActor(c.fields)
			got, err := f.client().Actor(serverTestActorID).Server().GetOpenAPISpec(context.Background())
			if err != nil || string(got) != spec {
				t.Fatalf("spec %q, %v", got, err)
			}
			reqs := f.requests("/v2/actor-builds/")
			if len(reqs) != 1 || reqs[0].path != "/v2/actor-builds/"+c.wantBuildID {
				t.Fatalf("build requests: %+v", reqs)
			}
		})
	}
	t.Run("looks a build number up in the builds of the Actor", func(t *testing.T) {
		f := newServerFixture(t)
		f.setActor(map[string]any{"actorStandby": map[string]any{"build": "0.0.1"}})
		f.setBuilds(map[string]string{"latest-build-id": "0.0.2", "old-build-id": "0.0.1"}, "latest-build-id", "old-build-id")
		got, err := f.client().Actor(serverTestActorID).Server().GetOpenAPISpec(context.Background())
		if err != nil || string(got) != spec {
			t.Fatalf("spec %q, %v", got, err)
		}
		if reqs := f.requests("/v2/actor-builds/"); reqs[len(reqs)-1].path != "/v2/actor-builds/old-build-id" {
			t.Fatalf("build requests: %+v", reqs)
		}
	})
	t.Run("a build without a web server schema gives nil", func(t *testing.T) {
		for name, definition := range map[string]any{"no schema": map[string]any{"name": "x"}, "no definition": nil, "null schema": map[string]any{"webServerSchema": nil}} {
			f := newServerFixture(t)
			f.definition = definition
			got, err := f.client().Actor(serverTestActorID).Server().GetOpenAPISpec(context.Background())
			if err != nil || got != nil {
				t.Fatalf("%s: spec %q, %v", name, got, err)
			}
		}
	})
	t.Run("a build that is neither a tag nor a listed number does not exist", func(t *testing.T) {
		f := newServerFixture(t)
		f.setActor(map[string]any{"actorStandby": map[string]any{"build": "9.9.9"}})
		f.setBuilds(map[string]string{"latest-build-id": "0.0.2"}, "latest-build-id")
		_, err := f.client().Actor(serverTestActorID).Server().GetOpenAPISpec(context.Background())
		if err == nil || !strings.Contains(err.Error(), "build 9.9.9 of Actor test-actor-id does not exist") {
			t.Fatalf("got %v", err)
		}
	})
	t.Run("the number lookup stops at an older build of the same version", func(t *testing.T) {
		f := newServerFixture(t)
		f.setActor(map[string]any{"actorStandby": map[string]any{"build": "0.0.4"}})
		f.setBuilds(map[string]string{"b5": "0.0.5", "b3": "0.0.3", "b2": "0.0.2", "b1": "0.0.1"}, "b5", "b3", "b2", "b1")
		f.buildsPerPage = 1
		_, err := f.client().Actor(serverTestActorID).Server().GetOpenAPISpec(context.Background())
		if err == nil || !strings.Contains(err.Error(), "does not exist") {
			t.Fatalf("got %v", err)
		}
		if n := len(f.requests("/v2/actors/" + serverTestActorID + "/builds")); n != 2 {
			t.Fatalf("expected 2 build list requests, got %d", n)
		}
	})
	t.Run("an unknown tag does not list the builds", func(t *testing.T) {
		f := newServerFixture(t)
		f.setActor(map[string]any{"actorStandby": map[string]any{"build": "nightly"}})
		_, err := f.client().Actor(serverTestActorID).Server().GetOpenAPISpec(context.Background())
		if err == nil || !strings.Contains(err.Error(), "build nightly of Actor test-actor-id does not exist") {
			t.Fatalf("got %v", err)
		}
		if n := len(f.requests("/v2/actors/" + serverTestActorID + "/builds")); n != 0 {
			t.Fatalf("expected no build list request, got %d", n)
		}
	})
	t.Run("the URL fetched for the spec serves Request", func(t *testing.T) {
		f := newServerFixture(t)
		server := f.client().Actor(serverTestActorID).Server()
		if _, err := server.GetOpenAPISpec(context.Background()); err != nil {
			t.Fatal(err)
		}
		if _, err := server.Request(context.Background(), "/", ActorServerRequestOptions{}); err != nil {
			t.Fatal(err)
		}
		if n := len(f.requests("/v2/actors/" + serverTestActorID)); n != 1 {
			t.Fatalf("expected one Actor fetch, got %d", n)
		}
	})
}

func TestActorStandbyIsTheServerClient(t *testing.T) {
	actor := NewClient(WithToken("t")).Actor("a")
	if actor.Standby() != actor.Server() {
		t.Fatal("Standby() must return the server client")
	}
}

func TestActorStandbyURLIsNormalized(t *testing.T) {
	var actor Actor
	if err := json.Unmarshal([]byte(`{"id":"a","standbyUrl":"HTTPS://Example.COM:443/x"}`), &actor); err != nil {
		t.Fatal(err)
	}
	if actor.StandbyURL == nil || *actor.StandbyURL != "https://example.com/x" {
		t.Fatalf("StandbyURL = %v", actor.StandbyURL)
	}
	if _, ok := actor.Extra["standbyUrl"]; ok {
		t.Fatal("standbyUrl must not be in Extra")
	}
}

func TestActorVersionCarriesTheActorPath(t *testing.T) {
	var v ActorVersion
	if err := json.Unmarshal([]byte(`{"versionNumber":"0.1","sourceType":"SOURCE_FILES","actorPath":"actors/a"}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.ActorPath == nil || *v.ActorPath != "actors/a" {
		t.Fatalf("ActorPath = %v", v.ActorPath)
	}
	if err := json.Unmarshal([]byte(`{"versionNumber":"0.1","sourceType":"SOURCE_FILES","actorPath":null}`), &v); err != nil || v.ActorPath != nil {
		t.Fatalf("a null actorPath must decode to nil: %v, %v", v.ActorPath, err)
	}
}
