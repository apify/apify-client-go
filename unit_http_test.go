package apify

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

// mockBackend is a deterministic HTTPBackend for offline unit tests. It serves a queue of
// scripted responses/errors and records how many times it was called.
type mockBackend struct {
	mu          sync.Mutex
	responses   []mockResponse
	calls       int
	lastHeaders http.Header
	lastURL     string
	lastBody    string
	bodies      []string
}

type mockResponse struct {
	status int
	body   string
	err    error
}

func (m *mockBackend) Do(req *http.Request) (*http.Response, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	idx := m.calls
	m.calls++
	m.lastHeaders = req.Header.Clone()
	m.lastURL = req.URL.String()
	if req.Body != nil {
		data, _ := io.ReadAll(req.Body)
		m.lastBody = string(data)
		m.bodies = append(m.bodies, string(data))
	}
	// If we run past the script, repeat the last entry (so "constant" behaviour is easy).
	if idx >= len(m.responses) {
		idx = len(m.responses) - 1
	}
	r := m.responses[idx]
	if r.err != nil {
		return nil, r.err
	}
	return &http.Response{
		StatusCode: r.status,
		Header:     http.Header{},
		Body:       io.NopCloser(bytes.NewReader([]byte(r.body))),
	}, nil
}

func constant(status int, body string) []mockResponse {
	return []mockResponse{{status: status, body: body}}
}

// testClient builds a client wired to the given backend with a tiny retry delay so tests
// are fast.
func testClient(backend HTTPBackend, maxRetries int) *ApifyClient {
	return NewClientWithOptions(
		WithToken("test-token"),
		WithHTTPBackend(backend),
		WithMaxRetries(maxRetries),
		WithMinDelayBetweenRetries(time.Millisecond),
	)
}

func TestSuccessSingleCall(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"u1","username":"bob"}}`)}
	client := testClient(backend, 8)

	user, ok, err := client.Me().Get(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok || user.ID != "u1" || user.Username != "bob" {
		t.Fatalf("unexpected user: %+v ok=%v", user, ok)
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 call, got %d", backend.calls)
	}
}

func TestRateLimitIsRetried(t *testing.T) {
	backend := &mockBackend{responses: constant(429, `{"error":{"type":"rate-limit-exceeded","message":"slow down"}}`)}
	client := testClient(backend, 2)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	apiErr, ok := AsAPIError(err)
	if !ok || apiErr.StatusCode != 429 {
		t.Fatalf("expected 429 APIError, got %v", err)
	}
	// 1 initial + 2 retries = 3 attempts.
	if backend.calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", backend.calls)
	}
	if apiErr.Attempt != 3 {
		t.Fatalf("expected attempt 3, got %d", apiErr.Attempt)
	}
}

func TestServerErrorIsRetried(t *testing.T) {
	backend := &mockBackend{responses: constant(503, `{"error":{"type":"internal","message":"boom"}}`)}
	client := testClient(backend, 1)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 2 {
		t.Fatalf("expected 2 attempts, got %d", backend.calls)
	}
}

func TestClientErrorNotRetried(t *testing.T) {
	backend := &mockBackend{responses: constant(400, `{"error":{"type":"bad-request","message":"nope"}}`)}
	client := testClient(backend, 5)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt for 4xx, got %d", backend.calls)
	}
}

func TestNetworkErrorIsRetried(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{{err: errors.New("connection refused")}}}
	client := testClient(backend, 3)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 4 {
		t.Fatalf("expected 4 attempts, got %d", backend.calls)
	}
}

func TestRetryThenSuccess(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{
		{status: 500, body: `{"error":{"type":"internal","message":"x"}}`},
		{status: 500, body: `{"error":{"type":"internal","message":"x"}}`},
		{status: 200, body: `{"data":{"id":"ok"}}`},
	}}
	client := testClient(backend, 5)

	user, ok, err := client.Me().Get(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok || user.ID != "ok" {
		t.Fatalf("unexpected user: %+v", user)
	}
	if backend.calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", backend.calls)
	}
}

func TestNotFoundMapsToNone(t *testing.T) {
	backend := &mockBackend{responses: constant(404, `{"error":{"type":"record-not-found","message":"missing"}}`)}
	client := testClient(backend, 5)

	_, ok, err := client.Actor("nope").Get(context.Background())
	if err != nil {
		t.Fatalf("expected nil error for not-found, got %v", err)
	}
	if ok {
		t.Fatal("expected ok=false for missing resource")
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt (no retry on 404), got %d", backend.calls)
	}
}

func TestErrorBodyIsParsed(t *testing.T) {
	backend := &mockBackend{responses: constant(400, `{"error":{"type":"bad-request","message":"invalid input","data":{"field":"name"}}}`)}
	client := testClient(backend, 0)

	_, _, err := client.Me().Get(context.Background())
	apiErr, ok := AsAPIError(err)
	if !ok {
		t.Fatalf("expected APIError, got %v", err)
	}
	if apiErr.StatusCode != 400 || apiErr.Type != "bad-request" || apiErr.Message != "invalid input" {
		t.Fatalf("unexpected fields: %+v", apiErr)
	}
	if apiErr.HTTPMethod != http.MethodGet {
		t.Fatalf("expected GET, got %q", apiErr.HTTPMethod)
	}
	if apiErr.Path == "" {
		t.Fatal("expected a non-empty path")
	}
	if apiErr.Data["field"] == nil {
		t.Fatalf("expected error data to be parsed, got %+v", apiErr.Data)
	}
}

// gunzip decompresses gzip-encoded bytes, failing the test on error.
func gunzip(t *testing.T, data []byte) []byte {
	t.Helper()
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("body is not valid gzip: %v", err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read gzip body: %v", err)
	}
	return out
}

func TestMaybeCompressRequestBody(t *testing.T) {
	// Below the threshold: sent verbatim, no encoding.
	small := []byte("small payload")
	out, enc := maybeCompressRequestBody(small)
	if enc != "" || !bytes.Equal(out, small) {
		t.Fatalf("small body must not be compressed, got enc=%q", enc)
	}

	// A nil body stays nil (GET requests have no body).
	out, enc = maybeCompressRequestBody(nil)
	if enc != "" || out != nil {
		t.Fatalf("nil body must stay nil, got enc=%q out=%v", enc, out)
	}

	// Above the threshold and compressible: gzip-encoded, round-trips to the original.
	large := []byte(strings.Repeat("compress me ", 200)) // ~2400 bytes, highly repetitive
	out, enc = maybeCompressRequestBody(large)
	if enc != contentEncodingGzip {
		t.Fatalf("large compressible body must be gzip, got enc=%q", enc)
	}
	if len(out) >= len(large) {
		t.Fatalf("compressed body should be smaller: %d >= %d", len(out), len(large))
	}
	if !bytes.Equal(gunzip(t, out), large) {
		t.Fatal("decompressed body does not match original")
	}

	// Above the threshold but incompressible: gzip cannot shrink random bytes, so the shrink
	// guard sends the original body uncompressed (empty encoding).
	rng := rand.New(rand.NewSource(1))
	incompressible := make([]byte, 2048)
	if _, err := rng.Read(incompressible); err != nil {
		t.Fatalf("failed to build random payload: %v", err)
	}
	out, enc = maybeCompressRequestBody(incompressible)
	if enc != "" {
		t.Fatalf("incompressible body must be sent uncompressed, got enc=%q", enc)
	}
	if !bytes.Equal(out, incompressible) {
		t.Fatal("incompressible body must be returned unchanged")
	}
}

// A request body large enough to compress is sent gzip-encoded, with the Content-Encoding
// header set, and the bytes on the wire decompress back to the original payload.
func TestLargeRequestBodyIsGzipCompressed(t *testing.T) {
	backend := &mockBackend{responses: constant(201, `{"data":{}}`)}
	client := testClient(backend, 0)

	bigValue := strings.Repeat("x", 4096)
	if err := client.Dataset("ds1").PushItems(context.Background(), map[string]string{"blob": bigValue}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := backend.lastHeaders.Get("Content-Encoding"); got != contentEncodingGzip {
		t.Fatalf("expected Content-Encoding %q, got %q", contentEncodingGzip, got)
	}
	decoded := gunzip(t, []byte(backend.lastBody))
	if !strings.Contains(string(decoded), bigValue) {
		t.Fatal("decompressed request body does not contain the pushed payload")
	}
}

// A small request body is sent uncompressed and carries no Content-Encoding header.
func TestSmallRequestBodyIsNotCompressed(t *testing.T) {
	backend := &mockBackend{responses: constant(201, `{"data":{}}`)}
	client := testClient(backend, 0)

	if err := client.Dataset("ds1").PushItems(context.Background(), map[string]string{"k": "v"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := backend.lastHeaders.Get("Content-Encoding"); got != "" {
		t.Fatalf("small body must not set Content-Encoding, got %q", got)
	}
	if !strings.Contains(backend.lastBody, `"v"`) {
		t.Fatalf("small body should be sent as plain JSON, got %q", backend.lastBody)
	}
}

func TestZeroRetriesSingleAttempt(t *testing.T) {
	backend := &mockBackend{responses: constant(500, `{"error":{"type":"internal","message":"x"}}`)}
	client := testClient(backend, 0)

	_, _, err := client.Me().Get(context.Background())
	if err == nil {
		t.Fatal("expected an error")
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 attempt with max_retries=0, got %d", backend.calls)
	}
}
