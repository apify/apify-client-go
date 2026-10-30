package apify

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

// syncBuffer is a bytes.Buffer that is safe to read while the redirection writes to it.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

// logBackend serves a run's log: streamBodies[i] answers the i-th request for the live stream
// (the last one repeats), and fullLog answers the read of the whole log. Run and Actor lookups
// answer the fixed documents below.
type logBackend struct {
	mu            sync.Mutex
	streamBodies  []string
	streamErr     error
	fullLog       string
	streamCalls   int
	fullLogCalls  int
	rawOnStream   bool
	streamTailErr error // makes the stream fail with this error after its body
	// blockFullLog makes the read of the whole log wait until its request is canceled.
	blockFullLog bool
	// holdStream keeps the live stream open after its body, until the request is canceled.
	holdStream bool
	// fullLogFailure, when set, makes the read of the whole log fail with this status after
	// fullLogDelay.
	fullLogFailure int
	fullLogDelay   time.Duration
}

type failingReader struct {
	data io.Reader
	err  error
}

func (r *failingReader) Read(p []byte) (int, error) {
	n, err := r.data.Read(p)
	if errors.Is(err, io.EOF) {
		return n, r.err
	}
	return n, err
}

// holdingReader serves data, then blocks until ctx is done, like an open log stream of a run that
// is still running but logs nothing.
type holdingReader struct {
	data io.Reader
	ctx  context.Context
}

func (r *holdingReader) Read(p []byte) (int, error) {
	n, err := r.data.Read(p)
	if errors.Is(err, io.EOF) && n == 0 {
		<-r.ctx.Done()
		return 0, r.ctx.Err()
	}
	if errors.Is(err, io.EOF) {
		return n, nil
	}
	return n, err
}

func (b *logBackend) fullLogReads() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.fullLogCalls
}

func (b *logBackend) Do(req *http.Request) (*http.Response, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	ok := func(body io.Reader) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Header: http.Header{}, Body: io.NopCloser(body)}, nil
	}
	path := req.URL.Path
	switch {
	case strings.HasSuffix(path, "/log"):
		if req.URL.Query().Get("stream") != "" {
			b.rawOnStream = req.URL.Query().Get("raw") == "1"
			if b.streamErr != nil {
				return nil, b.streamErr
			}
			body := b.streamBodies[min(b.streamCalls, len(b.streamBodies)-1)]
			b.streamCalls++
			var reader io.Reader = strings.NewReader(body)
			if b.holdStream {
				reader = &holdingReader{data: reader, ctx: req.Context()}
			}
			if b.streamTailErr != nil {
				reader = &failingReader{data: reader, err: b.streamTailErr}
			}
			return ok(reader)
		}
		b.fullLogCalls++
		if b.blockFullLog {
			b.mu.Unlock()
			<-req.Context().Done()
			b.mu.Lock()
			return nil, req.Context().Err()
		}
		if b.fullLogFailure != 0 {
			b.mu.Unlock()
			time.Sleep(b.fullLogDelay)
			b.mu.Lock()
			return &http.Response{StatusCode: b.fullLogFailure, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(`{"error":{"type":"forbidden","message":"no"}}`))}, nil
		}
		return ok(strings.NewReader(b.fullLog))
	case strings.HasSuffix(path, "/actors/actor1"):
		return ok(strings.NewReader(`{"data":{"id":"actor1","name":"hello-world"}}`))
	case strings.Contains(path, "/actor-runs/"):
		return ok(strings.NewReader(`{"data":{"id":"run1","actId":"actor1","status":"RUNNING"}}`))
	}
	return nil, errors.New("unexpected request " + req.URL.String())
}

const sampleLog = "2026-10-30T10:00:00.000Z ACTOR: Starting\n" +
	"2026-10-30T10:00:01.000Z first line\nsecond line of the same message\n" +
	"2026-10-30T10:00:02.000Z last message\n"

func TestStreamedLogRedirectsMessagesWithThePrefixOfTheActorAndTheRun(t *testing.T) {
	backend := &logBackend{streamBodies: []string{sampleLog}}
	out := &syncBuffer{}
	ctx := context.Background()

	streamed, found, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: out})
	if err != nil || !found {
		t.Fatalf("NewStreamedLog: %v, %v, %v", streamed, found, err)
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "last message") })
	if err := streamed.Stop(ctx); err != nil {
		t.Fatalf("Stop: %v", err)
	}

	const prefix = "hello-world runId:run1 -> "
	want := prefix + "2026-10-30T10:00:00.000Z ACTOR: Starting\n" +
		prefix + "2026-10-30T10:00:01.000Z first line\nsecond line of the same message\n" +
		prefix + "2026-10-30T10:00:02.000Z last message\n"
	if got := out.String(); got != want {
		t.Fatalf("redirected log:\n%q\nwant:\n%q", got, want)
	}
	if !backend.rawOnStream {
		t.Fatal("the raw log must be streamed")
	}
}

func TestStreamedLogPrefixOverrideSkipsTheLookups(t *testing.T) {
	backend := &logBackend{streamBodies: []string{sampleLog}}
	out := &syncBuffer{}
	ctx := context.Background()
	streamed, found, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: out, Prefix: Ptr("")})
	if err != nil || !found {
		t.Fatalf("NewStreamedLog: %v, %v, %v", streamed, found, err)
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "last message") })
	if err := streamed.Stop(ctx); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out.String(), "2026-10-30T10:00:00.000Z ACTOR: Starting\n") {
		t.Fatalf("an empty prefix is used as it is: %q", out.String())
	}
}

func TestStreamedLogDropsMessagesBeforeTheStartWhenNotFromStart(t *testing.T) {
	future := time.Now().Add(time.Hour).UTC().Format("2006-01-02T15:04:05.000Z")
	log := "2026-10-30T10:00:00.000Z old\n" + future + " new\n"
	backend := &logBackend{streamBodies: []string{log}}
	out := &syncBuffer{}
	ctx := context.Background()
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: out, Prefix: Ptr(""), FromStart: Ptr(false)})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "new") })
	if err := streamed.Stop(ctx); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(out.String(), "old") {
		t.Fatalf("a message logged before the redirection was created must be dropped: %q", out.String())
	}
}

// The API serves the log of a run that has not logged anything as an empty stream that ends at
// once: the redirection reopens it until the first bytes arrive.
func TestStreamedLogReopensAnEmptyStream(t *testing.T) {
	backend := &logBackend{streamBodies: []string{"", "", sampleLog}}
	out := &syncBuffer{}
	ctx := context.Background()
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: out, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "last message") })
	if err := streamed.Stop(ctx); err != nil {
		t.Fatal(err)
	}
	if backend.streamCalls != 3 {
		t.Fatalf("expected the empty stream to be reopened twice, got %d stream requests", backend.streamCalls)
	}
}

// A run that finished before the stream delivered anything is not lost: stopping reads the whole
// log in one request.
func TestStreamedLogStoppedBeforeTheFirstBytesReadsTheWholeLog(t *testing.T) {
	backend := &logBackend{streamBodies: []string{""}, fullLog: sampleLog}
	out := &syncBuffer{}
	ctx := context.Background()
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: out, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	if err := streamed.Stop(ctx); err != nil {
		t.Fatalf("Stop: %v", err)
	}
	if backend.fullLogCalls != 1 {
		t.Fatalf("expected one read of the whole log, got %d", backend.fullLogCalls)
	}
	if !strings.Contains(out.String(), "last message") || strings.Count(out.String(), "\n") < 4 {
		t.Fatalf("the whole log must be redirected: %q", out.String())
	}
}

func TestStreamedLogFlushesALastLineWithoutAnEndOfLine(t *testing.T) {
	backend := &logBackend{streamBodies: []string{"2026-10-30T10:00:00.000Z one\n2026-10-30T10:00:01.000Z unterminated"}}
	out := &syncBuffer{}
	ctx := context.Background()
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: out, Prefix: Ptr(">")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "unterminated") })
	if err := streamed.Stop(ctx); err != nil {
		t.Fatal(err)
	}
	if got := out.String(); got != ">2026-10-30T10:00:00.000Z one\n>2026-10-30T10:00:01.000Z unterminated\n" {
		t.Fatalf("got %q", got)
	}
}

func TestStreamedLogStartAndStopMustBePaired(t *testing.T) {
	backend := &logBackend{streamBodies: []string{sampleLog}}
	ctx := context.Background()
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: io.Discard, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Stop(ctx); err == nil {
		t.Fatal("Stop before Start must fail")
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(ctx); err == nil {
		t.Fatal("a second Start must fail")
	}
	if err := streamed.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}

// Two goroutines may call Stop at once: one stops the redirection, the other gets an error
// instead of closing a channel twice.
func TestStreamedLogConcurrentStopDoesNotPanic(t *testing.T) {
	backend := &logBackend{streamBodies: []string{""}, fullLog: sampleLog}
	ctx := context.Background()
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(ctx, StreamedLogOptions{Output: io.Discard, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(ctx); err != nil {
		t.Fatal(err)
	}
	const callers = 8
	results := make(chan error, callers)
	for range callers {
		go func() { results <- streamed.Stop(ctx) }()
	}
	succeeded := 0
	for range callers {
		if <-results == nil {
			succeeded++
		}
	}
	if succeeded != 1 {
		t.Fatalf("exactly one Stop must succeed, %d did", succeeded)
	}
}

// The context of Stop bounds the final read of the whole log, not only the wait for the stream.
func TestStreamedLogStopContextBoundsTheReadOfTheWholeLog(t *testing.T) {
	backend := &logBackend{streamBodies: []string{""}, blockFullLog: true}
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(context.Background(), StreamedLogOptions{Output: io.Discard, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	stopCtx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	started := time.Now()
	err = streamed.Stop(stopCtx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected the deadline error of the Stop context, got %v", err)
	}
	if elapsed := time.Since(started); elapsed > 5*time.Second {
		t.Fatalf("Stop must return once its context is done, took %v", elapsed)
	}
}

// Stop with a context that is already done does not turn a log that was redirected completely
// into an error: nothing was cut short.
func TestStreamedLogStopWithADoneContextAfterACompleteLogIsNotAnError(t *testing.T) {
	backend := &logBackend{streamBodies: []string{sampleLog}}
	out := &syncBuffer{}
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(context.Background(), StreamedLogOptions{Output: out, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "last message") })
	done, cancel := context.WithCancel(context.Background())
	cancel()
	if err := streamed.Stop(done); err != nil {
		t.Fatalf("a complete log is not an error, got %v", err)
	}
}

// A Call whose context ends while it waits does not report a failed redirection after a log that
// was redirected completely.
func TestCallWhoseContextEndsDoesNotReportAFalseRedirectionError(t *testing.T) {
	out := &syncBuffer{}
	backend := &callLogBackend{logBackend: logBackend{streamBodies: []string{sampleLog}, fullLog: sampleLog}, pollDelay: 300 * time.Millisecond}
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	_, err := testClient(backend, 0).Actor("actor1").Call(ctx, nil, ActorStartOptions{Log: &StreamedLogOptions{Output: out, Prefix: Ptr("")}}, Ptr(int64(0)))
	if err == nil {
		t.Fatal("expected the call to end with the error of its context")
	}
	if strings.Contains(out.String(), "Log redirection stopped due to error") {
		t.Fatalf("a complete log is not a redirection error: %q", out.String())
	}
}

// A failed final read of the whole log is reported even when it fails after the grace period of
// Stop: the log is lost, and the caller must hear of it.
func TestStreamedLogStopReportsAFailedReadOfTheWholeLog(t *testing.T) {
	backend := &logBackend{streamBodies: []string{""}, fullLogFailure: http.StatusForbidden, fullLogDelay: stopGrace + 500*time.Millisecond}
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(context.Background(), StreamedLogOptions{Output: io.Discard, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	err = streamed.Stop(context.Background())
	apiErr, ok := AsAPIError(err)
	if !ok || apiErr.StatusCode != http.StatusForbidden {
		t.Fatalf("expected the 403 of the final read, got %v", err)
	}
}

// Stop with a context that ends while the stream is still open and has delivered lines cuts the
// log off: it returns the error of its context, so a truncated log is not mistaken for a complete
// one.
func TestStreamedLogStopReportsALogCutOffByItsContext(t *testing.T) {
	backend := &logBackend{streamBodies: []string{sampleLog}, holdStream: true}
	out := &syncBuffer{}
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(context.Background(), StreamedLogOptions{Output: out, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "last message") })
	stopCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := streamed.Stop(stopCtx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected the error of the Stop context for a log that was cut off, got %v", err)
	}
}

// Without a Stop context to cut it off, a live stream ends after the grace period: that is how a
// redirection ends normally, and it is not an error.
func TestStreamedLogStopOfALiveStreamAfterTheGracePeriodIsNotAnError(t *testing.T) {
	backend := &logBackend{streamBodies: []string{sampleLog}, holdStream: true}
	out := &syncBuffer{}
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(context.Background(), StreamedLogOptions{Output: out, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "last message") })
	if err := streamed.Stop(context.Background()); err != nil {
		t.Fatalf("a normal stop is not an error, got %v", err)
	}
}

// The context given to Start ends the redirection on purpose, so its end is not an error of Stop,
// whether it ends before Stop or during the final read of the whole log.
func TestStreamedLogStartContextEndingIsNotAnErrorOfStop(t *testing.T) {
	t.Run("before Stop", func(t *testing.T) {
		backend := &logBackend{streamBodies: []string{""}, blockFullLog: true}
		parent, cancelParent := context.WithCancel(context.Background())
		streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(parent, StreamedLogOptions{Output: io.Discard, Prefix: Ptr("")})
		if err != nil {
			t.Fatal(err)
		}
		if err := streamed.Start(parent); err != nil {
			t.Fatal(err)
		}
		cancelParent()
		time.Sleep(50 * time.Millisecond)
		if err := streamed.Stop(context.Background()); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})
	t.Run("during the final read", func(t *testing.T) {
		backend := &logBackend{streamBodies: []string{""}, blockFullLog: true}
		parent, cancelParent := context.WithCancel(context.Background())
		streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(parent, StreamedLogOptions{Output: io.Discard, Prefix: Ptr("")})
		if err != nil {
			t.Fatal(err)
		}
		if err := streamed.Start(parent); err != nil {
			t.Fatal(err)
		}
		stopped := make(chan error, 1)
		go func() { stopped <- streamed.Stop(context.Background()) }()
		waitFor(t, func() bool { return backend.fullLogReads() == 1 })
		cancelParent()
		if err := <-stopped; err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})
}

// The end of the context given to Start ends the redirection on purpose: Stop with the same
// context, as Call does, reports no error for an open stream that is cut off by it.
func TestStreamedLogStopAfterTheStartContextEndedIsNotAnError(t *testing.T) {
	backend := &logBackend{streamBodies: []string{sampleLog}, holdStream: true}
	out := &syncBuffer{}
	parent, cancel := context.WithCancel(context.Background())
	streamed, _, err := testClient(backend, 0).Run("run1").NewStreamedLog(parent, StreamedLogOptions{Output: out, Prefix: Ptr("")})
	if err != nil {
		t.Fatal(err)
	}
	if err := streamed.Start(parent); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return strings.Contains(out.String(), "last message") })
	cancel()
	if err := streamed.Stop(parent); err != nil {
		t.Fatalf("the end of the Start context is not an error, got %v", err)
	}
}

// A Call whose context ends while the log stream is still open reports no redirection error.
func TestCallWhoseContextEndsWithAnOpenLogStreamReportsNoError(t *testing.T) {
	for range 10 {
		out := &syncBuffer{}
		backend := &callLogBackend{logBackend: logBackend{streamBodies: []string{sampleLog}, holdStream: true}, pollDelay: 300 * time.Millisecond}
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
		_, err := testClient(backend, 0).Actor("actor1").Call(ctx, nil, ActorStartOptions{Log: &StreamedLogOptions{Output: out, Prefix: Ptr("")}}, Ptr(int64(0)))
		cancel()
		if err == nil {
			t.Fatal("expected the call to end with the error of its context")
		}
		if strings.Contains(out.String(), "Log redirection stopped due to error") {
			t.Fatalf("the end of the context of the call is not a redirection error: %q", out.String())
		}
	}
}

func TestNewStreamedLogOfAMissingRunIsNotFound(t *testing.T) {
	backend := &mockBackend{responses: constant(404, `{"error":{"type":"record-not-found","message":"m"}}`)}
	streamed, found, err := testClient(backend, 0).Run("gone").NewStreamedLog(context.Background(), StreamedLogOptions{})
	if err != nil || found || streamed != nil {
		t.Fatalf("expected (nil, false, nil), got %v, %v, %v", streamed, found, err)
	}
}

// Call redirects the log of the run by default, and the call goes on when the redirection fails.
func TestCallRedirectsTheLogOfTheRun(t *testing.T) {
	out := &syncBuffer{}
	backend := &callLogBackend{logBackend: logBackend{streamBodies: []string{sampleLog}, fullLog: sampleLog}}
	got, err := testClient(backend, 0).Actor("actor1").Call(context.Background(), nil, ActorStartOptions{Log: &StreamedLogOptions{Output: out}}, Ptr(int64(0)))
	if err != nil || got.Status != "SUCCEEDED" {
		t.Fatalf("Call: %+v, %v", got, err)
	}
	if !strings.Contains(out.String(), "hello-world runId:run1 -> 2026-10-30T10:00:02.000Z last message") {
		t.Fatalf("the log must be redirected with the prefix of the Actor and the run: %q", out.String())
	}
}

func TestCallWithDisabledLogRedirectionRequestsNoLog(t *testing.T) {
	backend := &callLogBackend{logBackend: logBackend{streamBodies: []string{sampleLog}, fullLog: sampleLog}}
	_, err := testClient(backend, 0).Actor("actor1").Call(context.Background(), nil, ActorStartOptions{Log: &StreamedLogOptions{Disabled: true}}, Ptr(int64(0)))
	if err != nil {
		t.Fatal(err)
	}
	if backend.streamCalls != 0 || backend.fullLogCalls != 0 {
		t.Fatalf("no log may be requested, got %d stream and %d full reads", backend.streamCalls, backend.fullLogCalls)
	}
}

func TestACallGoesOnWhenTheLogRedirectionFails(t *testing.T) {
	out := &syncBuffer{}
	backend := &callLogBackend{logBackend: logBackend{streamBodies: []string{sampleLog}, streamErr: errors.New("boom")}, pollDelay: 100 * time.Millisecond}
	got, err := testClient(backend, 0).Actor("actor1").Call(context.Background(), nil, ActorStartOptions{Log: &StreamedLogOptions{Output: out, Prefix: Ptr("")}}, Ptr(int64(0)))
	if err != nil || got.Status != "SUCCEEDED" {
		t.Fatalf("a failing log redirection must not fail the call: %+v, %v", got, err)
	}
	if !strings.Contains(out.String(), "Log redirection stopped due to error") || !strings.Contains(out.String(), "boom") {
		t.Fatalf("the failure must be reported: %q", out.String())
	}
}

// callLogBackend also serves the start of a run, and answers every poll of it as finished after
// pollDelay, which leaves the log redirection time to run before the call stops it.
type callLogBackend struct {
	logBackend
	pollDelay time.Duration
}

func (b *callLogBackend) Do(req *http.Request) (*http.Response, error) {
	if req.Method == http.MethodPost && strings.HasSuffix(req.URL.Path, "/runs") {
		return &http.Response{StatusCode: 200, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(`{"data":{"id":"run1","actId":"actor1","status":"RUNNING"}}`))}, nil
	}
	if strings.Contains(req.URL.Path, "/actor-runs/") && !strings.HasSuffix(req.URL.Path, "/log") && req.URL.Query().Get("waitForFinish") != "" {
		select {
		case <-time.After(b.pollDelay):
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
		return &http.Response{StatusCode: 200, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(`{"data":{"id":"run1","actId":"actor1","status":"SUCCEEDED"}}`))}, nil
	}
	return b.logBackend.Do(req)
}

func waitFor(t *testing.T, condition func() bool) {
	t.Helper()
	deadline := time.Now().Add(10 * time.Second)
	for !condition() {
		if time.Now().After(deadline) {
			t.Fatal("timed out waiting for the condition")
		}
		time.Sleep(5 * time.Millisecond)
	}
}
