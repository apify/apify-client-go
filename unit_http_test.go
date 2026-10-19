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

	"github.com/andybalholm/brotli"
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
	status  int
	body    string
	err     error
	headers http.Header // response headers; nil means none
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
	header := r.headers
	if header == nil {
		header = http.Header{}
	}
	return &http.Response{
		StatusCode: r.status,
		Header:     header,
		Body:       io.NopCloser(bytes.NewReader([]byte(r.body))),
	}, nil
}

func constant(status int, body string) []mockResponse {
	return []mockResponse{{status: status, body: body}}
}

// testClient builds a client wired to the given backend with a tiny retry delay so tests
// are fast.
func testClient(backend HTTPBackend, maxRetries int) *ApifyClient {
	return NewClient(
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

// unbrotli decompresses brotli-encoded bytes, failing the test on error.
func unbrotli(t *testing.T, data []byte) []byte {
	t.Helper()
	out, err := io.ReadAll(brotli.NewReader(bytes.NewReader(data)))
	if err != nil {
		t.Fatalf("body is not valid brotli: %v", err)
	}
	return out
}

func TestMaybeCompressRequestBody(t *testing.T) {
	// Below the threshold: sent verbatim, no encoding.
	small := []byte("small payload")
	out, enc := maybeCompressRequestBody(small, "", compressorsFor(RequestCompressionAuto))
	if enc != "" || !bytes.Equal(out, small) {
		t.Fatalf("small body must not be compressed, got enc=%q", enc)
	}

	// A nil body stays nil (GET requests have no body).
	out, enc = maybeCompressRequestBody(nil, "", compressorsFor(RequestCompressionAuto))
	if enc != "" || out != nil {
		t.Fatalf("nil body must stay nil, got enc=%q out=%v", enc, out)
	}

	// Above the threshold and compressible: brotli is preferred, and it round-trips to the original.
	large := []byte(strings.Repeat("compress me ", 200)) // ~2400 bytes, highly repetitive
	out, enc = maybeCompressRequestBody(large, "", compressorsFor(RequestCompressionAuto))
	if enc != contentEncodingBrotli {
		t.Fatalf("large compressible body must prefer brotli, got enc=%q", enc)
	}
	if len(out) >= len(large) {
		t.Fatalf("compressed body should be smaller: %d >= %d", len(out), len(large))
	}
	if !bytes.Equal(unbrotli(t, out), large) {
		t.Fatal("decompressed body does not match original")
	}

	// Above the threshold but incompressible: neither codec can shrink random bytes, so the
	// shrink guard sends the original body uncompressed (empty encoding).
	rng := rand.New(rand.NewSource(1))
	incompressible := make([]byte, 2048)
	if _, err := rng.Read(incompressible); err != nil {
		t.Fatalf("failed to build random payload: %v", err)
	}
	out, enc = maybeCompressRequestBody(incompressible, "", compressorsFor(RequestCompressionAuto))
	if enc != "" {
		t.Fatalf("incompressible body must be sent uncompressed, got enc=%q", enc)
	}
	if !bytes.Equal(out, incompressible) {
		t.Fatal("incompressible body must be returned unchanged")
	}
}

// A large body whose Content-Type already carries its own compression (e.g. a ZIP archive) is
// sent uncompressed: compressing it again would burn CPU for little to no size reduction.
func TestMaybeCompressRequestBody_AlreadyCompressedContentType(t *testing.T) {
	large := []byte(strings.Repeat("compress me ", 200))

	out, enc := maybeCompressRequestBody(large, "application/zip", compressorsFor(RequestCompressionAuto))
	if enc != "" {
		t.Fatalf("an already-compressed content type must not be compressed, got enc=%q", enc)
	}
	if !bytes.Equal(out, large) {
		t.Fatal("an already-compressed body must be sent verbatim")
	}
}

// image/bmp sits under the already-compressed "image/" prefix but is itself raw, uncompressed
// data, so it is still compressed. Exercises the exact-media-type exception list.
func TestMaybeCompressRequestBody_CompressibleExceptionUnderPrefix(t *testing.T) {
	large := []byte(strings.Repeat("compress me ", 200))

	out, enc := maybeCompressRequestBody(large, "image/bmp; charset=binary", compressorsFor(RequestCompressionAuto))
	if enc != contentEncodingBrotli {
		t.Fatalf("image/bmp is raw despite the image/ prefix and should still be compressed, got enc=%q", enc)
	}
	if len(out) >= len(large) {
		t.Fatalf("compressed body should be smaller: %d >= %d", len(out), len(large))
	}
}

// WithRequestCompression(RequestCompressionGzip) must use gzip only, even for a body brotli
// would otherwise have compressed first.
func TestWithRequestCompressionGzipOnly(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{}}`)}
	client := NewClient(
		WithToken("t"),
		WithHTTPBackend(backend),
		WithMaxRetries(0),
		WithRequestCompression(RequestCompressionGzip),
	)

	large := []byte(strings.Repeat("compress me ", 200))
	if _, err := client.Actor("some-actor").Update(context.Background(), map[string]any{"v": string(large)}); err != nil {
		t.Fatalf("update: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got != contentEncodingGzip {
		t.Fatalf("expected gzip-only compression, got Content-Encoding %q", got)
	}
}

// The default RequestCompressionAuto still prefers brotli, matching existing behavior.
func TestDefaultRequestCompressionPrefersBrotli(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{}}`)}
	client := testClient(backend, 0)

	large := []byte(strings.Repeat("compress me ", 200))
	if _, err := client.Actor("some-actor").Update(context.Background(), map[string]any{"v": string(large)}); err != nil {
		t.Fatalf("update: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got != contentEncodingBrotli {
		t.Fatalf("expected brotli by default, got Content-Encoding %q", got)
	}
}

func TestIsCompressibleContentType(t *testing.T) {
	cases := []struct {
		contentType string
		want        bool
	}{
		{"", true},
		{"application/json", true},
		{"application/json; charset=utf-8", true},
		{"image/svg+xml", true},
		{"image/png", false},
		{"IMAGE/PNG", false},
		{"audio/mpeg", false},
		{"video/mp4", false},
		{"application/zip", false},
		{"application/octet-stream", true},
		{"image/bmp", true},
		{"audio/wav", true},
		{"font/woff2", false},
	}
	for _, c := range cases {
		if got := isCompressibleContentType(c.contentType); got != c.want {
			t.Errorf("isCompressibleContentType(%q) = %v, want %v", c.contentType, got, c.want)
		}
	}
}

// Both codecs must round-trip a payload and actually shrink it, so each compression path is
// exercised directly rather than only through the preference selector.
func TestRequestCompressorsRoundTrip(t *testing.T) {
	payload := []byte(strings.Repeat("compress me ", 200))

	brOut, err := brotliCompress(payload)
	if err != nil {
		t.Fatalf("brotliCompress error: %v", err)
	}
	if len(brOut) >= len(payload) {
		t.Fatalf("brotli should shrink payload: %d >= %d", len(brOut), len(payload))
	}
	if !bytes.Equal(unbrotli(t, brOut), payload) {
		t.Fatal("brotli round-trip mismatch")
	}

	gzOut, err := gzipCompress(payload)
	if err != nil {
		t.Fatalf("gzipCompress error: %v", err)
	}
	if len(gzOut) >= len(payload) {
		t.Fatalf("gzip should shrink payload: %d >= %d", len(gzOut), len(payload))
	}
	if !bytes.Equal(gunzip(t, gzOut), payload) {
		t.Fatal("gzip round-trip mismatch")
	}
}

// The gzip fallback must be genuinely reachable: when the preferred (brotli) codec errors, the
// selector falls through to gzip and emits Content-Encoding: gzip.
func TestCompressionFallsBackToGzipWhenBrotliFails(t *testing.T) {
	payload := []byte(strings.Repeat("compress me ", 200))
	failingBrotli := requestCompressor{
		encoding: contentEncodingBrotli,
		compress: func([]byte) ([]byte, error) { return nil, errors.New("brotli unavailable") },
	}
	compressors := []requestCompressor{failingBrotli, {encoding: contentEncodingGzip, compress: gzipCompress}}

	out, enc := compressWith(payload, compressors)
	if enc != contentEncodingGzip {
		t.Fatalf("expected gzip fallback, got enc=%q", enc)
	}
	if !bytes.Equal(gunzip(t, out), payload) {
		t.Fatal("gzip fallback body does not round-trip")
	}
}

// A request body large enough to compress is sent brotli-encoded (the preferred codec), with the
// Content-Encoding header set, and the bytes on the wire decompress back to the original payload.
func TestLargeRequestBodyIsBrotliCompressed(t *testing.T) {
	backend := &mockBackend{responses: constant(201, `{"data":{}}`)}
	client := testClient(backend, 0)

	bigValue := strings.Repeat("x", 4096)
	if err := client.Dataset("ds1").PushItems(context.Background(), map[string]string{"blob": bigValue}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := backend.lastHeaders.Get("Content-Encoding"); got != contentEncodingBrotli {
		t.Fatalf("expected Content-Encoding %q, got %q", contentEncodingBrotli, got)
	}
	decoded := unbrotli(t, []byte(backend.lastBody))
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

// SetRecordRawWithEncoding must send the caller's Content-Encoding verbatim and must never run
// the value through the client's own request compression — even though the value is well above
// the auto-compression threshold, it must reach the backend byte-for-byte, or a pre-compressed
// upload would be corrupted by being compressed a second time.
func TestSetRecordRawWithEncodingSendsHeaderVerbatimAndSkipsCompression(t *testing.T) {
	backend := &mockBackend{responses: constant(201, "")}
	client := testClient(backend, 0)

	// Larger than minCompressRequestBytes, and already "pre-compressed" opaque bytes (not
	// valid gzip/brotli, which would make a false-positive auto-compression attempt obvious
	// by corrupting the body).
	preCompressed := []byte(strings.Repeat("not-real-gzip-bytes-", 100))

	err := client.KeyValueStore("kvs1").SetRecordRawWithEncoding(
		context.Background(), "OUTPUT", preCompressed, "application/octet-stream", "gzip",
	)
	if err != nil {
		t.Fatalf("set record with encoding: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got != "gzip" {
		t.Fatalf("expected Content-Encoding: gzip, got %q", got)
	}
	if backend.lastBody != string(preCompressed) {
		t.Fatalf("body must be sent as-is, not re-compressed: got %d bytes, want %d bytes", len(backend.lastBody), len(preCompressed))
	}
}

// An empty contentEncoding must behave exactly like SetRecordRaw: no Content-Encoding header.
func TestSetRecordRawWithEncodingEmptyMeansNoHeader(t *testing.T) {
	backend := &mockBackend{responses: constant(201, "")}
	client := testClient(backend, 0)

	err := client.KeyValueStore("kvs1").SetRecordRawWithEncoding(
		context.Background(), "OUTPUT", []byte("hello"), "text/plain", "",
	)
	if err != nil {
		t.Fatalf("set record with encoding: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got != "" {
		t.Fatalf("expected no Content-Encoding header, got %q", got)
	}
}

// Even with contentEncoding == "", SetRecordRawWithEncoding must still skip the client's own
// request compression: a body at or above minCompressRequestBytes that SetRecordRaw would
// brotli/gzip-compress must be sent uncompressed here instead (the whole point of this call is
// "send value exactly as given").
func TestSetRecordRawWithEncodingEmptyStillSkipsAutoCompression(t *testing.T) {
	backend := &mockBackend{responses: constant(201, "")}
	client := testClient(backend, 0)

	large := strings.Repeat("x", minCompressRequestBytes+1)
	err := client.KeyValueStore("kvs1").SetRecordRawWithEncoding(
		context.Background(), "OUTPUT", []byte(large), "text/plain", "",
	)
	if err != nil {
		t.Fatalf("set record with encoding: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got != "" {
		t.Fatalf("expected no Content-Encoding header, got %q", got)
	}
	if backend.lastBody != large {
		t.Fatalf("expected the body to be sent uncompressed and unmodified, got %d bytes (want %d)",
			len(backend.lastBody), len(large))
	}

	// SetRecordRaw, by contrast, does auto-compress a body this size: this is the behaviour
	// SetRecordRawWithEncoding is documented to not have.
	backend2 := &mockBackend{responses: constant(201, "")}
	client2 := testClient(backend2, 0)
	if err := client2.KeyValueStore("kvs1").SetRecordRaw(context.Background(), "OUTPUT", []byte(large), "text/plain"); err != nil {
		t.Fatalf("set record raw: %v", err)
	}
	if got := backend2.lastHeaders.Get("Content-Encoding"); got == "" {
		t.Fatal("expected SetRecordRaw to auto-compress a body this size (sanity check on the contrast above)")
	}
}

// KeyValueStoreClient.GetRecordStream must open a streaming GET (attachment=1 by default,
// mirroring GetRecord) and report the same ContentType a buffered GetRecord would.
func TestGetRecordStreamOpensStreamingRequest(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{{
		status:  200,
		body:    "record bytes",
		headers: http.Header{"Content-Type": []string{"application/octet-stream"}},
	}}}
	client := testClient(backend, 0)

	rec, ok, err := client.KeyValueStore("kvs1").GetRecordStream(context.Background(), "OUTPUT")
	if err != nil || !ok {
		t.Fatalf("get record stream: ok=%v err=%v", ok, err)
	}
	defer func() { _ = rec.Value.Close() }()
	if rec.ContentType != "application/octet-stream" {
		t.Fatalf("unexpected content type: %q", rec.ContentType)
	}
	data, err := io.ReadAll(rec.Value)
	if err != nil {
		t.Fatalf("read stream: %v", err)
	}
	if string(data) != "record bytes" {
		t.Fatalf("unexpected stream contents: %q", string(data))
	}
	if !strings.Contains(backend.lastURL, "attachment=1") {
		t.Fatalf("expected attachment=1 in URL, got %q", backend.lastURL)
	}
}

// A missing record must resolve to (nil, false, nil), matching GetRecord's behaviour, since a
// key-value-store record lookup by key is unambiguous.
func TestGetRecordStreamNotFoundMapsToAbsent(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	rec, ok, err := client.KeyValueStore("kvs1").GetRecordStream(context.Background(), "MISSING")
	if err != nil || ok || rec != nil {
		t.Fatalf("expected absent, got rec=%v ok=%v err=%v", rec, ok, err)
	}
}

// LogClient.GetBuffer must return the exact response bytes, and GetWithOptions/Get must still
// decode them to the equivalent string — the two are the same request, differing only in return
// type.
func TestLogGetBufferReturnsRawBytes(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "log line one\nlog line two\n")}
	client := testClient(backend, 0)

	buf, ok, err := client.Run("r1").Log().GetBuffer(context.Background())
	if err != nil || !ok {
		t.Fatalf("get buffer: ok=%v err=%v", ok, err)
	}
	if string(buf) != "log line one\nlog line two\n" {
		t.Fatalf("unexpected buffer contents: %q", string(buf))
	}

	text, ok, err := client.Run("r1").Log().Get(context.Background())
	if err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}
	if text != string(buf) {
		t.Fatalf("Get and GetBuffer must read the same content: %q vs %q", text, string(buf))
	}
}

// A 404 for LogClient.GetBuffer must resolve to "absent" under the same rule as Get: only when
// the client has its own id.
func TestLogGetBufferNotFoundMapsToAbsent(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	buf, ok, err := client.Log("log1").GetBuffer(context.Background())
	if err != nil || ok || buf != nil {
		t.Fatalf("expected absent, got buf=%v ok=%v err=%v", buf, ok, err)
	}
}

// DatasetClient.DownloadItemsStream must open a streaming GET against the items endpoint (format
// in the query) and hand back the raw response body unbuffered.
func TestDatasetDownloadItemsStreamOpensStreamingRequest(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "a,b,c\n1,2,3\n")}
	client := testClient(backend, 0)

	stream, err := client.Dataset("ds1").DownloadItemsStream(context.Background(), FormatCSV, DatasetDownloadOptions{})
	if err != nil {
		t.Fatalf("download items stream: %v", err)
	}
	defer func() { _ = stream.Close() }()

	data, err := io.ReadAll(stream)
	if err != nil {
		t.Fatalf("read stream: %v", err)
	}
	if string(data) != "a,b,c\n1,2,3\n" {
		t.Fatalf("unexpected stream contents: %q", string(data))
	}
	if !strings.Contains(backend.lastURL, "/datasets/ds1/items") {
		t.Fatalf("expected request against the items endpoint, got %q", backend.lastURL)
	}
	if !strings.Contains(backend.lastURL, "format=csv") {
		t.Fatalf("expected format=csv in URL, got %q", backend.lastURL)
	}
}

// A 404 from DownloadItemsStream must always surface as an error, matching DownloadItems (the
// buffered counterpart), since a missing dataset is unambiguous here.
func TestDatasetDownloadItemsStreamNotFoundIsError(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	_, err := client.Dataset("ds1").DownloadItemsStream(context.Background(), FormatCSV, DatasetDownloadOptions{})
	if err == nil {
		t.Fatal("expected an error for a 404 dataset")
	}
	apiErr, ok := AsAPIError(err)
	if !ok || !apiErr.IsNotFound() {
		t.Fatalf("expected a 404 *APIError, got %v", err)
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
