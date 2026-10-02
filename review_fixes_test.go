package apify

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"
)

// okData is a constant 200 response with an empty data envelope, reused by the offline
// tests below.
func okData() []mockResponse { return constant(200, `{"data":{}}`) }

// Charge must always send an idempotency-key header so a transport retry cannot double-charge.
func TestChargeSendsIdempotencyKey(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "my-event"})
	if err != nil {
		t.Fatalf("charge: %v", err)
	}
	key := backend.lastHeaders.Get("idempotency-key")
	if key == "" {
		t.Fatal("expected an idempotency-key header to be sent")
	}
	if !strings.HasPrefix(key, "run123-my-event-") {
		t.Fatalf("idempotency key should embed run id and event name, got %q", key)
	}
	if !strings.Contains(backend.lastBody, `"eventName":"my-event"`) {
		t.Fatalf("unexpected charge body: %s", backend.lastBody)
	}
	if !strings.Contains(backend.lastURL, "/actor-runs/run123/charge") {
		t.Fatalf("expected charge request against /actor-runs/run123/charge, got %q", backend.lastURL)
	}
}

// A caller-supplied idempotency key must be used verbatim.
func TestChargeHonorsExplicitIdempotencyKey(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "e", IdempotencyKey: "fixed-key"})
	if err != nil {
		t.Fatalf("charge: %v", err)
	}
	if got := backend.lastHeaders.Get("idempotency-key"); got != "fixed-key" {
		t.Fatalf("expected explicit idempotency key, got %q", got)
	}
}

// GetRecord must default attachment=true (matching the reference client).
func TestGetRecordDefaultsAttachment(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "raw-bytes")}
	client := testClient(backend, 0)

	_, _, err := client.KeyValueStore("store1").GetRecord(context.Background(), "OUTPUT")
	if err != nil {
		t.Fatalf("get record: %v", err)
	}
	if !strings.Contains(backend.lastURL, "attachment=1") {
		t.Fatalf("expected attachment=1 in URL, got %q", backend.lastURL)
	}
}

// BatchAddRequests must split inputs larger than the API's 25-per-batch limit into chunks.
func TestBatchAddRequestsChunks(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"processedRequests":[],"unprocessedRequests":[]}}`)}
	client := testClient(backend, 0)

	requests := make([]RequestQueueRequest, 60) // 60 -> 25 + 25 + 10 = 3 chunks
	for i := range requests {
		requests[i] = RequestQueueRequest{URL: "https://example.com"}
	}
	if _, err := client.RequestQueue("q1").BatchAddRequests(context.Background(), requests, false); err != nil {
		t.Fatalf("batch add: %v", err)
	}
	if backend.calls != 3 {
		t.Fatalf("expected 3 batch calls for 60 requests, got %d", backend.calls)
	}
}

// Run.GetWithWait must forward the waitForFinish query parameter.
func TestRunGetWithWaitForwardsParam(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"r1"}}`)}
	client := testClient(backend, 0)

	secs := int64(30)
	if _, _, err := client.Run("r1").GetWithWait(context.Background(), &secs); err != nil {
		t.Fatalf("get with wait: %v", err)
	}
	if !strings.Contains(backend.lastURL, "waitForFinish=30") {
		t.Fatalf("expected waitForFinish=30 in URL, got %q", backend.lastURL)
	}
}

// GetStreamedLog must open a raw, streaming log request (stream=1 & raw=1).
func TestGetStreamedLogParams(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "live log bytes")}
	client := testClient(backend, 0)

	stream, err := client.Run("r1").GetStreamedLog(context.Background())
	if err != nil {
		t.Fatalf("get streamed log: %v", err)
	}
	defer func() { _ = stream.Close() }()
	if !strings.Contains(backend.lastURL, "stream=1") {
		t.Fatalf("expected stream=1 in URL, got %q", backend.lastURL)
	}
	if !strings.Contains(backend.lastURL, "raw=1") {
		t.Fatalf("expected raw=1 in URL, got %q", backend.lastURL)
	}
}

// notFound is a constant 404 response, used to simulate a resource that never becomes
// available (e.g. database-replica lag on a just-started run that never resolves).
func notFound() []mockResponse {
	return constant(404, `{"error":{"type":"record-not-found","message":"missing"}}`)
}

// A bounded WaitForFinish against a resource that stays 404 must terminate when the budget
// is exhausted and return a descriptive error (not the raw 404), matching the reference
// client's "Waiting for run to finish failed" behaviour.
func TestWaitForFinishBoundedReturnsDescriptiveErrorOn404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	zeroSecs := int64(0)
	_, err := client.Run("r1").WaitForFinish(context.Background(), &zeroSecs)
	if err == nil {
		t.Fatal("expected a descriptive error when the run never becomes available")
	}
	if !strings.Contains(err.Error(), "waiting for run to finish failed") {
		t.Fatalf("expected descriptive wait error, got %q", err.Error())
	}
	// 404 must not be retried by the transport, and the budget=0 path does exactly one fetch.
	if backend.calls != 1 {
		t.Fatalf("expected exactly 1 fetch for a zero-budget wait, got %d", backend.calls)
	}
}

// An INDEFINITE WaitForFinish (nil waitSecs) against a permanent 404 must NOT hang forever:
// it polls through 404s on a pure time bound, so a cancelled context unblocks it promptly.
// This is the regression test for the waitForFinish hang bug.
func TestWaitForFinishIndefiniteDoesNotHangOn404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		_, err := client.Run("r1").WaitForFinish(ctx, nil)
		done <- err
	}()

	select {
	case err := <-done:
		// It must return because the context was cancelled (or, in principle, the budget),
		// never spin indefinitely.
		if err == nil {
			t.Fatal("expected a non-nil error (context cancelled), got nil")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("WaitForFinish hung on a permanent 404 during an indefinite wait")
	}
	// Must have polled at least once through the 404 (proving 404 is not treated as fatal).
	if backend.calls < 1 {
		t.Fatalf("expected at least one poll through the 404, got %d", backend.calls)
	}
}

// ListRequests must reject mutually-exclusive options and an invalid Filter before any API call.
func TestListRequestsValidation(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	queue := client.RequestQueue("q1")

	start, cursor := "a", "b"
	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{ExclusiveStartID: &start, Cursor: &cursor}); err == nil {
		t.Fatal("expected error for mutually-exclusive ExclusiveStartID and Cursor")
	}
	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{Filter: []string{"bogus"}}); err == nil {
		t.Fatal("expected error for invalid Filter value")
	}
	if backend.calls != 0 {
		t.Fatalf("validation must short-circuit before any API call, got %d calls", backend.calls)
	}
}

// A resource id, record key or request id that is a dot segment (".." or ".") must be rejected
// before any request is sent: a URL parser or proxy can resolve it after the path is built,
// letting it escape to a different endpoint (path traversal).
func TestPathTraversalRejected(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	cases := []struct {
		name string
		call func() error
	}{
		{"resource id", func() error { _, _, err := client.Actor("..").Get(context.Background()); return err }},
		{"record key", func() error {
			_, _, err := client.KeyValueStore("store1").GetRecord(context.Background(), "..")
			return err
		}},
		{"request id", func() error {
			_, _, err := client.RequestQueue("q1").GetRequest(context.Background(), "..")
			return err
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := c.call(); err == nil {
				t.Fatal("expected an error for a dot-segment id, got nil")
			}
		})
	}
	if backend.calls != 0 {
		t.Fatalf("a dot-segment id must be rejected before any API call, got %d calls", backend.calls)
	}
}

// An empty resource id must be rejected the same way as a dot segment, not silently read as
// "the whole collection" (e.g. an empty Actor version number).
func TestEmptyIDRejected(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	if _, _, err := client.Actor("some-actor").Version("").Get(context.Background()); err == nil {
		t.Fatal("expected an error for an empty version number, got nil")
	}
	if backend.calls != 0 {
		t.Fatalf("an empty id must be rejected before any API call, got %d calls", backend.calls)
	}
}

// Get/Delete on a resource addressed by its own id must swallow a 404 of any `type`, not just
// "record-not-found"/"record-or-token-not-found" (matching the reference client's NotFoundError,
// which classifies solely by status).
func TestGetSwallowsAnyNotFoundType(t *testing.T) {
	backend := &mockBackend{responses: constant(404, `{"error":{"type":"some-other-error","message":"gone"}}`)}
	client := testClient(backend, 0)

	_, present, err := client.Actor("missing").Get(context.Background())
	if err != nil {
		t.Fatalf("expected a 404 of any type to resolve to absent, got error: %v", err)
	}
	if present {
		t.Fatal("expected present=false for a 404 response")
	}
}

// Get/Delete on a client reached through a fixed sub-path with no id of its own (e.g. a run's
// default dataset) must NOT swallow a 404: the response cannot tell the run apart from the
// dataset as what is missing.
func TestNestedResourceGetPropagates404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)

	_, _, err := client.Run("run1").Dataset().Get(context.Background())
	if err == nil {
		t.Fatal("expected a 404 on a nested dataset client to propagate as an error")
	}
	apiErr, ok := AsAPIError(err)
	if !ok || apiErr.StatusCode != 404 {
		t.Fatalf("expected a 404 *APIError, got %v", err)
	}

	err = client.Run("run1").Dataset().Delete(context.Background())
	if err == nil {
		t.Fatal("expected a 404 on a nested dataset client's Delete to propagate as an error")
	}
}

// GetRecord/RecordExists on a nested key-value store (no id of its own) must still swallow a
// 404: unlike the store's own Get/Delete, a record lookup by key is never ambiguous.
func TestNestedKeyValueStoreRecordLookupStillSwallows404(t *testing.T) {
	backend := &mockBackend{responses: notFound()}
	client := testClient(backend, 0)
	store := client.Run("run1").KeyValueStore()

	record, present, err := store.GetRecord(context.Background(), "OUTPUT")
	if err != nil {
		t.Fatalf("expected a missing record to resolve to absent, got error: %v", err)
	}
	if present || record != nil {
		t.Fatalf("expected present=false and a nil record, got present=%v record=%v", present, record)
	}

	exists, err := store.RecordExists(context.Background(), "OUTPUT")
	if err != nil {
		t.Fatalf("expected RecordExists to resolve false, got error: %v", err)
	}
	if exists {
		t.Fatal("expected exists=false for a 404 response")
	}
}

// DatasetClient.GetStatistics, TaskClient.GetInput and ScheduleClient.GetLog must propagate a
// 404 as an error rather than an "absent" result: it always means the parent resource is gone.
func TestSingletonSubResourcesPropagate404(t *testing.T) {
	cases := []struct {
		name string
		call func(c *ApifyClient) error
	}{
		{"dataset statistics", func(c *ApifyClient) error { _, err := c.Dataset("ds1").GetStatistics(context.Background()); return err }},
		{"task input", func(c *ApifyClient) error { _, err := c.Task("task1").GetInput(context.Background()); return err }},
		{"schedule log", func(c *ApifyClient) error { _, err := c.Schedule("sch1").GetLog(context.Background()); return err }},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			backend := &mockBackend{responses: notFound()}
			client := testClient(backend, 0)
			if err := c.call(client); err == nil {
				t.Fatal("expected a 404 to propagate as an error")
			}
		})
	}
}

// ApiError's Is* status-classification predicates match the response status, and only that
// status's predicate is true.
func TestAPIErrorStatusPredicates(t *testing.T) {
	errorFor := func(status int) *APIError {
		backend := &mockBackend{responses: constant(status, `{"error":{"type":"x","message":"m"}}`)}
		client := testClient(backend, 0)
		_, err := client.Me().MonthlyUsage(context.Background())
		apiErr, ok := AsAPIError(err)
		if !ok {
			t.Fatalf("expected an *APIError for status %d, got %v", status, err)
		}
		return apiErr
	}

	notFoundErr := errorFor(404)
	if !notFoundErr.IsNotFound() {
		t.Error("expected IsNotFound() for a 404")
	}
	for _, predicate := range []struct {
		name string
		ok   bool
	}{
		{"IsInvalidRequest", notFoundErr.IsInvalidRequest()},
		{"IsUnauthorized", notFoundErr.IsUnauthorized()},
		{"IsForbidden", notFoundErr.IsForbidden()},
		{"IsConflict", notFoundErr.IsConflict()},
		{"IsRateLimited", notFoundErr.IsRateLimited()},
		{"IsServerError", notFoundErr.IsServerError()},
	} {
		if predicate.ok {
			t.Errorf("expected %s() to be false for a 404", predicate.name)
		}
	}

	if !errorFor(400).IsInvalidRequest() {
		t.Error("expected IsInvalidRequest() for a 400")
	}
	if !errorFor(401).IsUnauthorized() {
		t.Error("expected IsUnauthorized() for a 401")
	}
	if !errorFor(403).IsForbidden() {
		t.Error("expected IsForbidden() for a 403")
	}
	if !errorFor(409).IsConflict() {
		t.Error("expected IsConflict() for a 409")
	}
	if !errorFor(429).IsRateLimited() {
		t.Error("expected IsRateLimited() for a 429")
	}
	if !errorFor(500).IsServerError() {
		t.Error("expected IsServerError() for a 500")
	}
	if !errorFor(503).IsServerError() {
		t.Error("expected IsServerError() for a 503")
	}
}

// ActorClient.StartRaw sends the given bytes verbatim (no JSON serialization), defaulting to
// application/octet-stream when the caller does not set options.ContentType.
func TestStartRawSendsBytesVerbatimWithDefaultContentType(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"run1"}}`)}
	client := testClient(backend, 0)

	input := []byte("not json, just bytes: \x00\x01\x02")
	if _, err := client.Actor("me/actor").StartRaw(context.Background(), input, ActorStartOptions{}); err != nil {
		t.Fatalf("start raw: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Type"); got != "application/octet-stream" {
		t.Fatalf("expected default content type application/octet-stream, got %q", got)
	}
	if backend.lastBody != string(input) {
		t.Fatalf("expected the raw bytes to be sent verbatim, got %q", backend.lastBody)
	}
}

// RunClient.MetamorphRaw sends the given bytes verbatim and still forwards targetActorId/build
// as query parameters like Metamorph.
func TestMetamorphRawSendsBytesVerbatim(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"run1"}}`)}
	client := testClient(backend, 0)

	input := []byte("raw metamorph input")
	_, err := client.Run("run1").MetamorphRaw(context.Background(), "other/actor", input, MetamorphOptions{Build: "1.0"})
	if err != nil {
		t.Fatalf("metamorph raw: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Type"); got != "application/octet-stream" {
		t.Fatalf("expected default content type application/octet-stream, got %q", got)
	}
	if backend.lastBody != string(input) {
		t.Fatalf("expected the raw bytes to be sent verbatim, got %q", backend.lastBody)
	}
	if !strings.Contains(backend.lastURL, "targetActorId=other~actor") {
		t.Fatalf("expected targetActorId in URL, got %q", backend.lastURL)
	}
	if !strings.Contains(backend.lastURL, "build=1.0") {
		t.Fatalf("expected build in URL, got %q", backend.lastURL)
	}
}

// DatasetClient.CreateItemsPublicURLWithFormat adds a `format` query parameter, and
// CreateItemsPublicURL (unchanged) omits it, defaulting to the endpoint's own json default.
func TestCreateItemsPublicURLWithFormat(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"ds1"}}`)}
	client := testClient(backend, 0)
	dataset := client.Dataset("ds1")

	url, err := dataset.CreateItemsPublicURLWithFormat(context.Background(), DatasetListItemsOptions{}, nil, FormatCSV)
	if err != nil {
		t.Fatalf("create items public url with format: %v", err)
	}
	if !strings.Contains(url, "format=csv") {
		t.Fatalf("expected format=csv in url, got %q", url)
	}

	url, err = dataset.CreateItemsPublicURL(context.Background(), DatasetListItemsOptions{}, nil)
	if err != nil {
		t.Fatalf("create items public url: %v", err)
	}
	if strings.Contains(url, "format=") {
		t.Fatalf("CreateItemsPublicURL must not set a format param, url was: %q", url)
	}
}

// IterateDatasetItems must advance and stop by the number of rows scanned
// (X-Apify-Pagination-Count), not the number of items a page returned: a filter can make a page
// return fewer items than the rows it covered (even zero, on a page that still has more rows
// behind it), and the iterator must not stop early in that case.
func TestIterateDatasetItemsUsesScannedCount(t *testing.T) {
	// pageHeaders builds the X-Apify-Pagination-* headers for a page that scanned `scanned` rows
	// of a dataset of `total` rows, starting at `offset`, having requested up to `limit`.
	pageHeaders := func(total, offset, limit, scanned int64) http.Header {
		h := http.Header{}
		h.Set("X-Apify-Pagination-Total", itoa(total))
		h.Set("X-Apify-Pagination-Offset", itoa(offset))
		h.Set("X-Apify-Pagination-Limit", itoa(limit))
		h.Set("X-Apify-Pagination-Count", itoa(scanned))
		return h
	}
	backend := &mockBackend{responses: []mockResponse{
		// Page 1: the API scanned 5 rows (offset 0-4) but a filter dropped all of them.
		{status: 200, body: `[]`, headers: pageHeaders(10, 0, 5, 5)},
		// Page 2: scanned the remaining 5 rows (offset 5-9), all kept.
		{status: 200, body: `[1,2,3,4,5]`, headers: pageHeaders(10, 5, 5, 5)},
	}}
	client := testClient(backend, 0)

	it := client.Dataset("ds1").IterateItems(DatasetListItemsOptions{}, Ptr(int64(5)))
	var items []json.RawMessage
	for {
		item, err := it.Next(context.Background())
		if err != nil {
			t.Fatalf("iterate: %v", err)
		}
		if item == nil {
			break
		}
		items = append(items, *item)
	}
	if len(items) != 5 {
		t.Fatalf("expected 5 items across both pages, got %d", len(items))
	}
	if backend.calls != 2 {
		t.Fatalf("expected exactly 2 page fetches, got %d", backend.calls)
	}
}

// ListRequests must serialize a multi-value Filter as a single comma-joined query parameter,
// matching the spec (array of enum locked|pending, style=form explode=false) and the JS reference.
func TestListRequestsFilterCommaJoined(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	queue := client.RequestQueue("q1")

	if _, err := queue.ListRequests(context.Background(), ListRequestsOptions{
		Filter: []string{RequestFilterLocked, RequestFilterPending},
	}); err != nil {
		t.Fatalf("ListRequests with multi-value Filter: %v", err)
	}
	if backend.calls != 1 {
		t.Fatalf("expected exactly one API call, got %d", backend.calls)
	}
	gotURL := backend.lastURL
	if !strings.Contains(gotURL, "filter=locked%2Cpending") && !strings.Contains(gotURL, "filter=locked,pending") {
		t.Fatalf("expected filter sent comma-joined as a single param, got URL: %s", gotURL)
	}
}
