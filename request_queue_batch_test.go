package apify

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
)

// batchBackend answers batch-add requests like the API: it adds every request except the ones
// that reject says to leave unprocessed.
type batchBackend struct {
	mu       sync.Mutex
	calls    int
	bodies   [][]RequestQueueRequest
	inflight atomic.Int64
	peak     atomic.Int64
	// reject returns the unique keys that the call with the given index leaves unprocessed.
	reject func(call int, requests []RequestQueueRequest) map[string]bool
	// fail returns the error of the call with the given index, or nil.
	fail func(call int) error
	// delay is how long each call takes.
	delay time.Duration
}

func (b *batchBackend) Do(req *http.Request) (*http.Response, error) {
	data := decodeBody(req)
	var requests []RequestQueueRequest
	if err := json.Unmarshal(data, &requests); err != nil {
		return nil, err
	}
	current := b.inflight.Add(1)
	defer b.inflight.Add(-1)
	for peak := b.peak.Load(); current > peak && !b.peak.CompareAndSwap(peak, current); peak = b.peak.Load() {
	}
	if b.delay > 0 {
		time.Sleep(b.delay)
	}

	b.mu.Lock()
	call := b.calls
	b.calls++
	b.bodies = append(b.bodies, requests)
	b.mu.Unlock()

	if b.fail != nil {
		if err := b.fail(call); err != nil {
			return nil, err
		}
	}
	var rejected map[string]bool
	if b.reject != nil {
		rejected = b.reject(call, requests)
	}
	processed, unprocessed := []map[string]any{}, []map[string]any{}
	for _, r := range requests {
		if rejected[r.UniqueKey] {
			unprocessed = append(unprocessed, map[string]any{"uniqueKey": r.UniqueKey, "url": r.URL, "method": r.Method})
		} else {
			processed = append(processed, map[string]any{"requestId": "id-" + r.UniqueKey, "uniqueKey": r.UniqueKey})
		}
	}
	body, _ := json.Marshal(map[string]any{"data": map[string]any{"processedRequests": processed, "unprocessedRequests": unprocessed}})
	return &http.Response{StatusCode: 200, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(string(body)))}, nil
}

// decodeBody reads the body of req, which the client may have compressed.
func decodeBody(req *http.Request) []byte {
	var reader io.Reader = req.Body
	switch req.Header.Get("Content-Encoding") {
	case "br":
		reader = brotli.NewReader(req.Body)
	case "gzip":
		gz, err := gzip.NewReader(req.Body)
		if err != nil {
			return nil
		}
		reader = gz
	case "zstd":
		zr, err := zstd.NewReader(req.Body)
		if err != nil {
			return nil
		}
		defer zr.Close()
		reader = zr
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil
	}
	return data
}

func uniqueRequests(n int) []RequestQueueRequest {
	requests := make([]RequestQueueRequest, n)
	for i := range requests {
		key := "key-" + string(rune('a'+i/26)) + string(rune('a'+i%26))
		requests[i] = RequestQueueRequest{URL: "https://example.com/" + key, UniqueKey: key, Method: "GET"}
	}
	return requests
}

// fastBatchOptions retries without a noticeable delay.
func fastBatchOptions() BatchAddRequestsOptions {
	return BatchAddRequestsOptions{MinDelayBetweenUnprocessedRequestsRetries: Ptr(time.Millisecond)}
}

func TestBatchAddRequestsSplitsInputIntoBatchesOfTheAPILimit(t *testing.T) {
	backend := &batchBackend{}
	result, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequests(context.Background(), uniqueRequests(60), false)
	if err != nil {
		t.Fatal(err)
	}
	if backend.calls != 3 || len(result.ProcessedRequests) != 60 || len(result.UnprocessedRequests) != 0 {
		t.Fatalf("calls %d, processed %d, unprocessed %d", backend.calls, len(result.ProcessedRequests), len(result.UnprocessedRequests))
	}
}

func TestBatchAddRequestsRetriesTheUnprocessedRequestsOnly(t *testing.T) {
	backend := &batchBackend{reject: func(call int, requests []RequestQueueRequest) map[string]bool {
		if call == 0 {
			return map[string]bool{"key-aa": true, "key-ab": true}
		}
		return nil
	}}
	options := fastBatchOptions()
	result, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(5), options)
	if err != nil {
		t.Fatal(err)
	}
	if backend.calls != 2 {
		t.Fatalf("expected one retry, got %d calls", backend.calls)
	}
	if retried := backend.bodies[1]; len(retried) != 2 || retried[0].UniqueKey != "key-aa" || retried[1].UniqueKey != "key-ab" {
		t.Fatalf("the retry must send the unprocessed requests only, got %+v", retried)
	}
	if len(result.ProcessedRequests) != 5 || len(result.UnprocessedRequests) != 0 {
		t.Fatalf("processed %d, unprocessed %d", len(result.ProcessedRequests), len(result.UnprocessedRequests))
	}
}

func TestBatchAddRequestsReportsWhatIsStillUnprocessedWhenTheRetriesRunOut(t *testing.T) {
	backend := &batchBackend{reject: func(call int, requests []RequestQueueRequest) map[string]bool {
		return map[string]bool{"key-aa": true}
	}}
	options := fastBatchOptions()
	options.MaxUnprocessedRequestsRetries = Ptr(int64(2))
	result, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(3), options)
	if err != nil {
		t.Fatal(err)
	}
	if backend.calls != 3 {
		t.Fatalf("expected the first attempt and 2 retries, got %d calls", backend.calls)
	}
	if len(result.ProcessedRequests) != 2 || len(result.UnprocessedRequests) != 1 || result.UnprocessedRequests[0].UniqueKey != "key-aa" {
		t.Fatalf("result: %+v", result)
	}
}

func TestBatchAddRequestsDoesNotRetryWithZeroRetries(t *testing.T) {
	backend := &batchBackend{reject: func(int, []RequestQueueRequest) map[string]bool { return map[string]bool{"key-aa": true} }}
	options := fastBatchOptions()
	options.MaxUnprocessedRequestsRetries = Ptr(int64(0))
	result, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(3), options)
	if err != nil || backend.calls != 1 || len(result.UnprocessedRequests) != 1 {
		t.Fatalf("calls %d, unprocessed %d, err %v", backend.calls, len(result.UnprocessedRequests), err)
	}
}

func TestBatchAddRequestsSendsAtMostMaxParallelBatchesAtOnce(t *testing.T) {
	backend := &batchBackend{delay: 20 * time.Millisecond}
	options := BatchAddRequestsOptions{MaxParallel: Ptr(int64(2))}
	if _, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(125), options); err != nil {
		t.Fatal(err)
	}
	if backend.calls != 5 || backend.peak.Load() != 2 {
		t.Fatalf("calls %d, peak parallelism %d", backend.calls, backend.peak.Load())
	}
}

func TestBatchAddRequestsLeavesTheRequestsOfAFailedBatchUnprocessed(t *testing.T) {
	backend := &batchBackend{fail: func(call int) error {
		if call == 0 {
			return errors.New("connection refused")
		}
		return nil
	}}
	options := BatchAddRequestsOptions{MaxParallel: Ptr(int64(1))}
	result, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(30), options)
	if err != nil {
		t.Fatalf("a failed batch is not an error of the call: %v", err)
	}
	if len(result.UnprocessedRequests) != 25 || len(result.ProcessedRequests) != 5 {
		t.Fatalf("processed %d, unprocessed %d", len(result.ProcessedRequests), len(result.UnprocessedRequests))
	}
	if first := result.UnprocessedRequests[0]; first.UniqueKey != "key-aa" || first.URL == "" || first.Method != "GET" {
		t.Fatalf("an unprocessed request keeps its method, unique key and URL: %+v", first)
	}
}

func TestBatchAddRequestsKeepsWhatWasAddedWhenARetryFails(t *testing.T) {
	backend := &batchBackend{
		reject: func(call int, requests []RequestQueueRequest) map[string]bool {
			if call == 0 {
				return map[string]bool{"key-ab": true, "key-ac": true}
			}
			return nil
		},
		fail: func(call int) error {
			if call == 1 {
				return errors.New("connection reset")
			}
			return nil
		},
	}
	result, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(4), fastBatchOptions())
	if err != nil {
		t.Fatal(err)
	}
	if len(result.ProcessedRequests) != 2 || len(result.UnprocessedRequests) != 2 {
		t.Fatalf("processed %d, unprocessed %d", len(result.ProcessedRequests), len(result.UnprocessedRequests))
	}
	for _, r := range result.UnprocessedRequests {
		if r.UniqueKey != "key-ab" && r.UniqueKey != "key-ac" {
			t.Fatalf("an added request is reported as unprocessed: %+v", r)
		}
	}
}

func TestBatchAddRequestsReturnsTheErrorWhenTheContextEndsDuringARequest(t *testing.T) {
	backend := &batchBackend{delay: 100 * time.Millisecond, fail: func(int) error { return context.DeadlineExceeded }}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequests(ctx, uniqueRequests(2), false)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v", err)
	}
}

func TestBatchAddRequestsReturnsAnErrorForAnUndecodableResponse(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":`)}
	_, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequests(context.Background(), uniqueRequests(1), false)
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestUnprocessedRetryDelayStaysWithinTheCap(t *testing.T) {
	client := testClient(&batchBackend{}, 0).RequestQueue("q1")
	timeoutMax := client.ctx.http.retry.timeoutMax
	for _, minDelay := range []time.Duration{500 * time.Millisecond, 2 * timeoutMax} {
		for _, attempt := range []int{0, 1, 5, 40, 200} {
			delay := client.unprocessedRetryDelay(minDelay, attempt)
			if delay <= 0 || delay >= 2*timeoutMax {
				t.Fatalf("minDelay %v, attempt %d: delay %v", minDelay, attempt, delay)
			}
		}
	}
}

func TestBatchAddRequestsReturnsTheErrorOfAnInvalidQueueID(t *testing.T) {
	backend := &batchBackend{}
	_, err := testClient(backend, 0).RequestQueue("").BatchAddRequests(context.Background(), uniqueRequests(2), false)
	if err == nil || backend.calls != 0 {
		t.Fatalf("err = %v, calls = %d", err, backend.calls)
	}
}

func TestBatchAddRequestsStopsRetryingWhenTheContextEnds(t *testing.T) {
	backend := &batchBackend{reject: func(int, []RequestQueueRequest) map[string]bool { return map[string]bool{"key-aa": true} }}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	options := BatchAddRequestsOptions{MinDelayBetweenUnprocessedRequestsRetries: Ptr(time.Hour)}
	result, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(ctx, uniqueRequests(2), options)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v", err)
	}
	if len(result.ProcessedRequests) != 1 || len(result.UnprocessedRequests) != 1 {
		t.Fatalf("processed %d, unprocessed %d", len(result.ProcessedRequests), len(result.UnprocessedRequests))
	}
}

func TestBatchAddRequestsRejectsInvalidOptionsBeforeSendingAnything(t *testing.T) {
	cases := map[string]BatchAddRequestsOptions{
		"negative retries": {MaxUnprocessedRequestsRetries: Ptr(int64(-1))},
		"zero parallelism": {MaxParallel: Ptr(int64(0))},
		"negative delay":   {MinDelayBetweenUnprocessedRequestsRetries: Ptr(-time.Second)},
	}
	for name, options := range cases {
		t.Run(name, func(t *testing.T) {
			backend := &batchBackend{}
			if _, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(1), options); err == nil || backend.calls != 0 {
				t.Fatalf("err = %v, calls = %d", err, backend.calls)
			}
		})
	}
}

func TestBatchAddRequestsSendsForefront(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"processedRequests":[],"unprocessedRequests":[]}}`)}
	options := BatchAddRequestsOptions{Forefront: true, MaxUnprocessedRequestsRetries: Ptr(int64(0))}
	if _, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequestsWithOptions(context.Background(), uniqueRequests(1), options); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(backend.lastURL, "forefront=1") {
		t.Fatalf("URL: %s", backend.lastURL)
	}
}

func TestBatchAddRequestsRejectsAnOversizedRequestBeforeSendingAnything(t *testing.T) {
	backend := &batchBackend{}
	requests := uniqueRequests(30)
	// A URL this long exceeds payloadSizeLimitBytes, so the request cannot fit a batch of its own.
	requests[29].URL = "https://example.com/" + strings.Repeat("a", payloadSizeLimitBytes)

	_, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequests(context.Background(), requests, false)
	if err == nil || !strings.Contains(err.Error(), "index 29") {
		t.Fatalf("expected an error naming index 29, got %v", err)
	}
	if backend.calls != 0 {
		t.Fatalf("expected no batch to be sent, got %d calls", backend.calls)
	}
}

func TestBatchAddRequestsSplitsABatchThatExceedsTheByteLimit(t *testing.T) {
	backend := &batchBackend{}
	big := strings.Repeat("a", payloadSizeLimitBytes/2)
	requests := []RequestQueueRequest{
		{URL: "https://example.com/" + big, UniqueKey: "one"},
		{URL: "https://example.com/" + big, UniqueKey: "two"},
	}
	if _, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequests(context.Background(), requests, false); err != nil {
		t.Fatal(err)
	}
	if backend.calls != 2 {
		t.Fatalf("expected 2 batches, got %d calls", backend.calls)
	}
}

func TestBatchAddRequestsSendsEachBatchAsOneJSONArray(t *testing.T) {
	backend := &batchBackend{}
	requests := uniqueRequests(2)
	if _, err := testClient(backend, 0).RequestQueue("q1").BatchAddRequests(context.Background(), requests, false); err != nil {
		t.Fatal(err)
	}
	if got := backend.bodies[0]; len(got) != 2 || got[0].URL != requests[0].URL || got[1].URL != requests[1].URL {
		t.Fatalf("batch body: %+v", got)
	}
}
