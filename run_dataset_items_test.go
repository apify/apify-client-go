package apify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// fakeRunBackend simulates a run that keeps pushing items to its default dataset. Each answer to
// a run poll (a GET of the run or a waitForFinish) moves the simulation one step: runSteps[i] is
// the run status and itemCounts[i] the dataset's itemCount after it, and every item row exists
// from the start of the test, so a page past itemCount still returns rows, like the lagging
// itemCount of the real API.
type fakeRunBackend struct {
	mu         sync.Mutex
	runID      string
	runSteps   []string
	itemCounts []int64
	totalRows  int
	step       int
	requests   []string
}

func (b *fakeRunBackend) Do(req *http.Request) (*http.Response, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.requests = append(b.requests, req.URL.RequestURI())
	path := req.URL.Path
	query := req.URL.Query()
	respond := func(body string, header http.Header) (*http.Response, error) {
		if header == nil {
			header = http.Header{}
		}
		return &http.Response{StatusCode: http.StatusOK, Header: header, Body: io.NopCloser(bytes.NewReader([]byte(body)))}, nil
	}
	current := min(b.step, len(b.runSteps)-1)
	switch {
	case strings.HasSuffix(path, "/dataset/items"):
		offset, _ := strconv.Atoi(query.Get("offset"))
		limit, _ := strconv.Atoi(query.Get("limit"))
		var rows []map[string]int
		for i := offset; i < offset+limit && i < b.totalRows; i++ {
			rows = append(rows, map[string]int{"n": i})
		}
		body, _ := json.Marshal(rows)
		if rows == nil {
			body = []byte("[]")
		}
		header := http.Header{}
		header.Set("X-Apify-Pagination-Count", strconv.Itoa(len(rows)))
		return respond(string(body), header)
	case strings.HasSuffix(path, "/dataset"):
		return respond(fmt.Sprintf(`{"data":{"id":"d","itemCount":%d}}`, b.itemCounts[current]), nil)
	case strings.HasSuffix(path, "/actor-runs/"+b.runID), strings.HasSuffix(path, "/runs/last"):
		if query.Get("waitForFinish") != "" {
			b.step++
			current = min(b.step, len(b.runSteps)-1)
		}
		return respond(fmt.Sprintf(`{"data":{"id":%q,"status":%q}}`, b.runID, b.runSteps[current]), nil)
	}
	return nil, fmt.Errorf("unexpected request %s", req.URL)
}

func drainRunItems(t *testing.T, it *RunDatasetItemsIterator[map[string]int]) []int {
	t.Helper()
	var got []int
	for {
		item, err := it.Next(context.Background())
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		if item == nil {
			return got
		}
		got = append(got, (*item)["n"])
	}
}

func sequence(from, to int) []int {
	var out []int
	for i := from; i < to; i++ {
		out = append(out, i)
	}
	return out
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestIterateDatasetItemsReadsAFinishedRun(t *testing.T) {
	backend := &fakeRunBackend{runID: "run1", runSteps: []string{"SUCCEEDED"}, itemCounts: []int64{5}, totalRows: 5}
	client := testClient(backend, 0)
	got := drainRunItems(t, IterateRunDatasetItems[map[string]int](client.Run("run1"), RunIterateDatasetItemsOptions{}))
	if !equalInts(got, sequence(0, 5)) {
		t.Fatalf("got %v", got)
	}
}

// While the run is running, the iterator reads up to itemCount, waits for the run to finish, and
// then reads the rows past a stale itemCount until a page is empty.
func TestIterateDatasetItemsFollowsARunningRun(t *testing.T) {
	backend := &fakeRunBackend{
		runID:      "run1",
		runSteps:   []string{"RUNNING", "RUNNING", "SUCCEEDED"},
		itemCounts: []int64{2, 4, 4},
		totalRows:  7, // itemCount stays at 4 although 7 rows exist: the tail catches up
	}
	client := testClient(backend, 0)
	it := IterateRunDatasetItems[map[string]int](client.Run("run1"), RunIterateDatasetItemsOptions{
		ChunkSize:        Ptr(int64(3)),
		PollIntervalSecs: Ptr(int64(1)),
	})
	got := drainRunItems(t, it)
	if !equalInts(got, sequence(0, 7)) {
		t.Fatalf("got %v", got)
	}
	polls := 0
	for _, r := range backend.requests {
		if strings.Contains(r, "/actor-runs/run1?waitForFinish=") {
			polls++
		}
	}
	if polls == 0 {
		t.Fatalf("expected the iterator to wait for the run to finish, got %v", backend.requests)
	}
	// The first round stops at the itemCount of 2 although 3 rows fit in a chunk.
	if !strings.Contains(backend.requests[2], "offset=0&limit=2") {
		t.Fatalf("the first page must end at itemCount: %v", backend.requests)
	}
	for _, r := range backend.requests {
		if strings.Contains(r, "/dataset/items") && strings.Contains(r, "limit=") {
			query, _ := url.ParseQuery(strings.SplitN(r, "?", 2)[1])
			if limit, _ := strconv.Atoi(query.Get("limit")); limit > 3 {
				t.Fatalf("a page asked for %d rows, more than the chunk size 3: %s", limit, r)
			}
		}
	}
}

func TestIterateDatasetItemsHonorsOffsetAndLimit(t *testing.T) {
	backend := &fakeRunBackend{runID: "run1", runSteps: []string{"SUCCEEDED"}, itemCounts: []int64{10}, totalRows: 10}
	client := testClient(backend, 0)
	it := IterateRunDatasetItems[map[string]int](client.Run("run1"), RunIterateDatasetItemsOptions{
		Items:     DatasetListItemsOptions{Offset: Ptr(int64(2)), Limit: Ptr(int64(5))},
		ChunkSize: Ptr(int64(2)),
	})
	if got := drainRunItems(t, it); !equalInts(got, sequence(2, 7)) {
		t.Fatalf("got %v, want rows 2..6", got)
	}
}

// The tail reads past itemCount even when the run is finished before the first request, up to the
// Limit.
func TestIterateDatasetItemsLimitEndsTheTail(t *testing.T) {
	backend := &fakeRunBackend{runID: "run1", runSteps: []string{"SUCCEEDED"}, itemCounts: []int64{0}, totalRows: 10}
	client := testClient(backend, 0)
	it := IterateRunDatasetItems[map[string]int](client.Run("run1"), RunIterateDatasetItemsOptions{
		Items:     DatasetListItemsOptions{Limit: Ptr(int64(4))},
		ChunkSize: Ptr(int64(3)),
	})
	if got := drainRunItems(t, it); !equalInts(got, sequence(0, 4)) {
		t.Fatalf("got %v, want rows 0..3", got)
	}
}

// A LastRun client resolves runs/last on each request: the iteration pins itself to the run its
// first request resolved to.
func TestIterateDatasetItemsPinsALastRunClient(t *testing.T) {
	backend := &fakeRunBackend{runID: "run1", runSteps: []string{"SUCCEEDED"}, itemCounts: []int64{3}, totalRows: 3}
	client := testClient(backend, 0)
	got := drainRunItems(t, IterateRunDatasetItems[map[string]int](client.Actor("a").LastRun("SUCCEEDED"), RunIterateDatasetItemsOptions{}))
	if !equalInts(got, sequence(0, 3)) {
		t.Fatalf("got %v", got)
	}
	for _, r := range backend.requests[1:] {
		if strings.Contains(r, "/runs/last") {
			t.Fatalf("after resolving the run, the requests must name it, not runs/last: %s", r)
		}
	}
	if !strings.Contains(backend.requests[0], "/runs/last") || !strings.Contains(backend.requests[0], "status=SUCCEEDED") {
		t.Fatalf("the first request resolves the last run: %s", backend.requests[0])
	}
}

func TestIterateDatasetItemsRejectsDescendingOrder(t *testing.T) {
	backend := &fakeRunBackend{runID: "run1", runSteps: []string{"SUCCEEDED"}, itemCounts: []int64{3}, totalRows: 3}
	client := testClient(backend, 0)
	it := client.Run("run1").IterateDatasetItems(RunIterateDatasetItemsOptions{Items: DatasetListItemsOptions{Desc: Ptr(true)}})
	for range 2 {
		if _, err := it.Next(context.Background()); err == nil {
			t.Fatal("expected an error for a descending order")
		}
	}
	if len(backend.requests) != 0 {
		t.Fatalf("no request may be sent, got %v", backend.requests)
	}
}

func TestIterateDatasetItemsDecodesIntoRawMessages(t *testing.T) {
	backend := &fakeRunBackend{runID: "run1", runSteps: []string{"SUCCEEDED"}, itemCounts: []int64{2}, totalRows: 2}
	it := testClient(backend, 0).Run("run1").IterateDatasetItems(RunIterateDatasetItemsOptions{})
	var raws []string
	for {
		item, err := it.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if item == nil {
			break
		}
		raws = append(raws, string(*item))
	}
	if strings.Join(raws, ",") != `{"n":0},{"n":1}` {
		t.Fatalf("got %v", raws)
	}
}
