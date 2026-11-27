package apify

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// compressibleBody is larger than the compression threshold and compresses well, so the client's
// own request compression would pick it up if a stream were not left alone.
var compressibleBody = strings.Repeat(`{"key":"value"}`, 400)

// streamClient is a client with the request compression on, so that a test sees that a stream
// is sent untouched, and with retries on, so that a test sees that it is sent once.
func streamClient(backend HTTPBackend) *ApifyClient {
	return NewClient(
		WithToken("test-token"),
		WithHTTPBackend(backend),
		WithMaxRetries(retriesAllowed),
		WithMinDelayBetweenRetries(1),
		WithRequestCompression(RequestCompressionGzip),
	)
}

func assertSentOnceAsIs(t *testing.T, backend *mockBackend, wantContentType string) {
	t.Helper()
	if backend.calls != 1 {
		t.Fatalf("a stream is sent once, got %d requests", backend.calls)
	}
	if backend.lastBody != compressibleBody {
		t.Errorf("the stream was altered: got %d bytes, want %d", len(backend.lastBody), len(compressibleBody))
	}
	if got := backend.lastHeaders.Get("Content-Encoding"); got != "" {
		t.Errorf("a stream is not compressed by the client, got Content-Encoding %q", got)
	}
	if got := backend.lastHeaders.Get("Content-Type"); got != wantContentType {
		t.Errorf("Content-Type %q, want %q", got, wantContentType)
	}
}

func TestSetRecordStreamSendsTheStreamOnceAsIs(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "")}
	err := streamClient(backend).KeyValueStore("s").SetRecordStream(context.Background(), "k", strings.NewReader(compressibleBody), "application/json")
	if err != nil {
		t.Fatalf("SetRecordStream: %v", err)
	}
	assertSentOnceAsIs(t, backend, "application/json")
	if !strings.HasSuffix(backend.lastURL, "/key-value-stores/s/records/k") {
		t.Errorf("unexpected URL %s", backend.lastURL)
	}
}

func TestSetRecordStreamWithEncodingSendsTheEncodingHeader(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "")}
	err := streamClient(backend).KeyValueStore("s").SetRecordStreamWithEncoding(context.Background(), "k", strings.NewReader(compressibleBody), "text/plain", "gzip")
	if err != nil {
		t.Fatalf("SetRecordStreamWithEncoding: %v", err)
	}
	if backend.calls != 1 || backend.lastBody != compressibleBody || backend.lastHeaders.Get("Content-Encoding") != "gzip" {
		t.Errorf("unexpected request: %d calls, headers %v", backend.calls, backend.lastHeaders)
	}
}

func TestSetRecordStreamIsNotRetried(t *testing.T) {
	backend := &mockBackend{responses: constant(500, `{"error":{"type":"x","message":"boom"}}`)}
	err := streamClient(backend).KeyValueStore("s").SetRecordStream(context.Background(), "k", strings.NewReader("data"), "text/plain")
	if apiErr, ok := AsAPIError(err); !ok || apiErr.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected the 500 as an APIError, got %v", err)
	}
	if backend.calls != 1 {
		t.Fatalf("a stream is not retried, got %d requests", backend.calls)
	}
}

func TestSetRecordStreamRejectsADotSegmentKey(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "")}
	if err := streamClient(backend).KeyValueStore("s").SetRecordStream(context.Background(), "..", strings.NewReader("x"), "text/plain"); err == nil {
		t.Fatal("expected an error for a dot-segment key")
	}
	if backend.calls != 0 {
		t.Fatalf("nothing is sent for a rejected key, got %d requests", backend.calls)
	}
}

const startedRun = `{"data":{"id":"run1","status":"READY"}}`

func TestActorStartStreamSendsTheStreamOnce(t *testing.T) {
	backend := &mockBackend{responses: constant(201, startedRun)}
	run, err := streamClient(backend).Actor("a").StartStream(context.Background(), strings.NewReader(compressibleBody), ActorStartOptions{WaitForResources: Ptr(int64(60))})
	if err != nil || run.ID != "run1" {
		t.Fatalf("StartStream: %+v, %v", run, err)
	}
	assertSentOnceAsIs(t, backend, "application/octet-stream")

	backend = &mockBackend{responses: constant(201, startedRun)}
	options := ActorStartOptions{ContentType: Ptr("application/zip")}
	if _, err := streamClient(backend).Actor("a").StartStream(context.Background(), strings.NewReader(compressibleBody), options); err != nil {
		t.Fatalf("StartStream with a content type: %v", err)
	}
	assertSentOnceAsIs(t, backend, "application/zip")
}

func TestActorStartStreamIsNotRetriedOnResourceShortage(t *testing.T) {
	backend := &mockBackend{responses: constant(402, `{"error":{"type":"actor-memory-limit-exceeded","message":"no memory"}}`)}
	_, err := streamClient(backend).Actor("a").StartStream(context.Background(), strings.NewReader("input"), ActorStartOptions{WaitForResources: Ptr(int64(60))})
	if err == nil || backend.calls != 1 {
		t.Fatalf("expected the first rejection after one request, got %v after %d", err, backend.calls)
	}
}

func TestActorCallStreamWaitsForTheRun(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{
		{status: 201, body: startedRun},
		{status: 200, body: `{"data":{"id":"run1","status":"SUCCEEDED"}}`},
	}}
	run, err := streamClient(backend).Actor("a").CallStream(context.Background(), strings.NewReader(compressibleBody), ActorStartOptions{Log: noLogRedirection}, nil)
	if err != nil || run.Status != "SUCCEEDED" {
		t.Fatalf("CallStream: %+v, %v", run, err)
	}
	if len(backend.bodies) == 0 || backend.bodies[0] != compressibleBody {
		t.Errorf("the start request did not carry the stream")
	}
}

func TestTaskStartStreamSendsTheStreamOnceAsJSON(t *testing.T) {
	backend := &mockBackend{responses: constant(201, startedRun)}
	run, err := streamClient(backend).Task("t").StartStream(context.Background(), strings.NewReader(compressibleBody), TaskStartOptions{})
	if err != nil || run.ID != "run1" {
		t.Fatalf("StartStream: %+v, %v", run, err)
	}
	// The API accepts only a JSON object as the input of a task.
	assertSentOnceAsIs(t, backend, "application/json")
}

func TestTaskCallStreamWaitsForTheRun(t *testing.T) {
	backend := &mockBackend{responses: []mockResponse{
		{status: 201, body: startedRun},
		{status: 200, body: `{"data":{"id":"run1","status":"SUCCEEDED"}}`},
	}}
	run, err := streamClient(backend).Task("t").CallStream(context.Background(), strings.NewReader(`{"a":1}`), TaskStartOptions{}, nil)
	if err != nil || run.Status != "SUCCEEDED" {
		t.Fatalf("CallStream: %+v, %v", run, err)
	}
}

// A redirect is not followed, since a stream cannot be replayed: it comes back as an error.
func TestStreamCallsDoNotFollowRedirects(t *testing.T) {
	var targetHits atomic.Int64
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { targetHits.Add(1) }))
	defer target.Close()
	redirecting := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusTemporaryRedirect)
	}))
	defer redirecting.Close()

	client := NewClient(WithToken("t"), WithBaseURL(redirecting.URL), WithMaxRetries(0))
	err := client.KeyValueStore("s").SetRecordStream(context.Background(), "k", strings.NewReader("data"), "text/plain")
	if err == nil {
		t.Fatal("expected the redirect to be returned as an error")
	}
	if targetHits.Load() != 0 {
		t.Fatalf("the redirect target was called %d times", targetHits.Load())
	}
}

// closeRecorder is a stream that counts its closes.
type closeRecorder struct {
	io.Reader
	closes atomic.Int32
}

func (r *closeRecorder) Close() error {
	r.closes.Add(1)
	return nil
}

// A call that fails before it sends anything closes a stream that is an io.Closer, once: the
// HTTP transport, which closes the body of a request it sent, never saw it.
func TestStreamCallsThatFailBeforeSendingCloseTheStreamOnce(t *testing.T) {
	ctx := context.Background()
	client := streamClient(&mockBackend{responses: constant(200, startedRun)})
	badURLClient := NewClient(WithToken("t"), WithBaseURL("http://[::1"), WithMaxRetries(0))
	cases := []struct {
		name string
		run  func(t *testing.T, stream io.Reader) error
	}{
		{"invalid key", func(_ *testing.T, s io.Reader) error {
			return client.KeyValueStore("s").SetRecordStream(ctx, "..", s, "text/plain")
		}},
		{"empty store ID", func(_ *testing.T, s io.Reader) error {
			return client.KeyValueStore("").SetRecordStream(ctx, "k", s, "text/plain")
		}},
		{"empty task ID", func(_ *testing.T, s io.Reader) error {
			_, err := client.Task("").StartStream(ctx, s, TaskStartOptions{})
			return err
		}},
		{"invalid URL", func(_ *testing.T, s io.Reader) error {
			_, err := badURLClient.Actor("a").StartStream(ctx, s, ActorStartOptions{})
			return err
		}},
		{"server request with an unknown method", func(_ *testing.T, s io.Reader) error {
			_, err := client.Actor("a").Server().Request(ctx, "/", ActorServerRequestOptions{Method: "NOPE", Body: s})
			return err
		}},
		{"server request of an Actor that does not exist", func(t *testing.T, s io.Reader) error {
			f := newServerFixture(t)
			f.noActor()
			_, err := f.client().Actor(serverTestActorID).Server().Request(ctx, "/", ActorServerRequestOptions{Method: "POST", Body: s})
			return err
		}},
		{"server request to a malformed server URL", func(t *testing.T, s io.Reader) error {
			f := newServerFixture(t)
			f.setActor(map[string]any{"standbyUrl": "http://[::1"})
			_, err := f.client().Actor(serverTestActorID).Server().Request(ctx, "/", ActorServerRequestOptions{Method: "POST", Body: s})
			return err
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			stream := &closeRecorder{Reader: strings.NewReader("data")}
			if err := c.run(t, stream); err == nil {
				t.Fatal("expected the call to fail")
			}
			if got := stream.closes.Load(); got != 1 {
				t.Errorf("the stream was closed %d times, want once", got)
			}
		})
	}
}

// The body of a request that was sent is closed by the HTTP transport, and the client does not
// close the stream a second time, whether the request succeeds or fails.
func TestServerRequestStreamIsClosedOnceByTheTransport(t *testing.T) {
	f := newServerFixture(t)
	unreachable := httptest.NewServer(http.NotFoundHandler())
	unreachableURL := unreachable.URL
	unreachable.Close()
	f.setActor(map[string]any{"standbyUrl": unreachableURL})

	stream := &closeRecorder{Reader: strings.NewReader("data")}
	_, err := f.client().Actor(serverTestActorID).Server().Request(context.Background(), "/", ActorServerRequestOptions{Method: "POST", Body: stream})
	if err == nil {
		t.Fatal("expected the connection to be refused")
	}
	if got := stream.closes.Load(); got != 1 {
		t.Errorf("the stream was closed %d times, want once", got)
	}
}

// A stream that was sent successfully is closed once, by the HTTP transport of the default
// backend.
func TestSentStreamIsClosedOnce(t *testing.T) {
	f := newServerFixture(t)
	ctx := context.Background()
	cases := map[string]func(stream io.Reader) error{
		"record": func(s io.Reader) error {
			return f.client().KeyValueStore("s").SetRecordStream(ctx, "k", s, "text/plain")
		},
		"server request": func(s io.Reader) error {
			_, err := f.client().Actor(serverTestActorID).Server().Request(ctx, "/search", ActorServerRequestOptions{Method: "POST", Body: s})
			return err
		},
	}
	for name, run := range cases {
		t.Run(name, func(t *testing.T) {
			stream := &closeRecorder{Reader: strings.NewReader("data")}
			if err := run(stream); err != nil {
				t.Fatalf("the call failed: %v", err)
			}
			if got := stream.closes.Load(); got != 1 {
				t.Errorf("the stream was closed %d times, want once", got)
			}
		})
	}
}
