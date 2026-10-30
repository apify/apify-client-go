package apify

import (
	"context"
	"net/url"
	"strings"
	"testing"
)

func TestRunListsSendTheDeploymentFilter(t *testing.T) {
	ctx := context.Background()
	filter := RunListOptions{Status: []string{"RUNNING"}, Deployment: []string{"SINGLE_TENANT", "MULTI_TENANT_FREE_USERS"}}

	for name, list := range map[string]func(client *ApifyClient) error{
		"account runs": func(c *ApifyClient) error { _, err := c.Runs().List(ctx, ListOptions{}, filter); return err },
		"actor runs":   func(c *ApifyClient) error { _, err := c.Actor("a").Runs().List(ctx, ListOptions{}, filter); return err },
	} {
		backend := &mockBackend{responses: okData()}
		if err := list(testClient(backend, 0)); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		query, _ := url.ParseQuery(strings.SplitN(backend.lastURL, "?", 2)[1])
		if got := query.Get("deployment"); got != "SINGLE_TENANT,MULTI_TENANT_FREE_USERS" {
			t.Errorf("%s: deployment = %q", name, got)
		}
	}

	backend := &mockBackend{responses: okData()}
	if _, err := testClient(backend, 0).Runs().List(ctx, ListOptions{}, RunListOptions{}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(backend.lastURL, "deployment") {
		t.Errorf("deployment must be omitted when unset, got %q", backend.lastURL)
	}
}

func TestStoreListSendsIncludeInputSchema(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)
	if _, err := client.Store().List(context.Background(), StoreListOptions{IncludeInputSchema: Ptr(true)}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(backend.lastURL, "includeInputSchema=1") {
		t.Errorf("expected includeInputSchema=1, got %q", backend.lastURL)
	}
	if _, err := client.Store().List(context.Background(), StoreListOptions{}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(backend.lastURL, "includeInputSchema") {
		t.Errorf("includeInputSchema must be omitted when unset, got %q", backend.lastURL)
	}
}

// The storages and the log of a run reached through LastRun carry its status and origin filters,
// so they address the storages and the log of the same run, not those of the newest run of any
// status.
func TestStoragesAndLogOfAFilteredLastRunCarryItsFilters(t *testing.T) {
	ctx := context.Background()
	lastRun := func(client *ApifyClient) *RunClient {
		return client.Actor("a").LastRunWithOptions(LastRunOptions{Status: "SUCCEEDED", Origin: "API"})
	}
	taskLastRun := func(client *ApifyClient) *RunClient {
		return client.Task("t").LastRunWithOptions(LastRunOptions{Status: "SUCCEEDED", Origin: "API"})
	}

	cases := map[string]func(client *ApifyClient) error{
		"dataset get": func(c *ApifyClient) error { _, _, err := lastRun(c).Dataset().Get(ctx); return err },
		"dataset items": func(c *ApifyClient) error {
			_, err := lastRun(c).Dataset().ListItems(ctx, DatasetListItemsOptions{})
			return err
		},
		"dataset download": func(c *ApifyClient) error {
			_, err := lastRun(c).Dataset().DownloadItems(ctx, "json", DatasetDownloadOptions{})
			return err
		},
		"dataset push": func(c *ApifyClient) error {
			return lastRun(c).Dataset().PushItems(ctx, map[string]any{"a": 1})
		},
		"dataset public URL": func(c *ApifyClient) error {
			_, err := lastRun(c).Dataset().CreateItemsPublicURL(ctx, DatasetListItemsOptions{}, nil)
			return err
		},
		"store get": func(c *ApifyClient) error { _, _, err := lastRun(c).KeyValueStore().Get(ctx); return err },
		"store keys": func(c *ApifyClient) error {
			_, err := lastRun(c).KeyValueStore().ListKeys(ctx, ListKeysOptions{})
			return err
		},
		"store record": func(c *ApifyClient) error {
			_, _, err := lastRun(c).KeyValueStore().GetRecord(ctx, "OUTPUT")
			return err
		},
		"store record public URL": func(c *ApifyClient) error {
			_, err := lastRun(c).KeyValueStore().GetRecordPublicURL(ctx, "OUTPUT")
			return err
		},
		"store keys public URL": func(c *ApifyClient) error {
			_, err := lastRun(c).KeyValueStore().CreateKeysPublicURL(ctx, nil)
			return err
		},
		"queue get":  func(c *ApifyClient) error { _, _, err := lastRun(c).RequestQueue().Get(ctx); return err },
		"queue head": func(c *ApifyClient) error { _, err := lastRun(c).RequestQueue().ListHead(ctx, nil); return err },
		"log get":    func(c *ApifyClient) error { _, _, err := lastRun(c).Log().Get(ctx); return err },
		"log stream": func(c *ApifyClient) error {
			stream, err := lastRun(c).GetStreamedLog(ctx)
			if err == nil {
				_ = stream.Close()
			}
			return err
		},
		"task last run dataset": func(c *ApifyClient) error { _, _, err := taskLastRun(c).Dataset().Get(ctx); return err },
		"task last run log":     func(c *ApifyClient) error { _, _, err := taskLastRun(c).Log().Get(ctx); return err },
	}
	for name, call := range cases {
		t.Run(name, func(t *testing.T) {
			body := `{"data":{"id":"x","urlSigningSecretKey":"s"}}`
			if name == "dataset items" {
				body = `[]` // the items endpoint answers with a bare array
			}
			backend := &mockBackend{responses: []mockResponse{{status: 200, body: body}}}
			if err := call(testClient(backend, 0)); err != nil {
				t.Fatalf("call failed: %v", err)
			}
			if !strings.Contains(backend.lastURL, "/runs/last/") {
				t.Fatalf("expected a runs/last URL, got %q", backend.lastURL)
			}
			for _, want := range []string{"status=SUCCEEDED", "origin=API"} {
				if !strings.Contains(backend.lastURL, want) {
					t.Errorf("%s is missing from %q", want, backend.lastURL)
				}
			}
		})
	}
}

// A storage reached through a run that has no filters does not gain any.
func TestStoragesOfAnUnfilteredRunHaveNoFilterParams(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	if _, _, err := testClient(backend, 0).Run("r").Dataset().Get(context.Background()); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(backend.lastURL, "?") {
		t.Errorf("unexpected query string in %q", backend.lastURL)
	}
}

// The API returns a pending request with a null handledAt, and a schedule whose Console editor
// state was dropped with a null uiState: both parse, the null kept in Extra.
func TestNullFieldsOfRequestsAndSchedulesParse(t *testing.T) {
	var request RequestQueueRequest
	mustUnmarshal(t, `{"id":"r1","url":"https://example.com/","uniqueKey":"k","handledAt":null}`, &request)
	if request.ID != "r1" || request.UniqueKey != "k" || string(request.Extra["handledAt"]) != "null" {
		t.Errorf("unexpected request %+v", request)
	}
	if request.Method != "" || request.UserData != nil {
		t.Errorf("the optional fields of a minimal request stay empty: %+v", request)
	}

	var schedule Schedule
	mustUnmarshal(t, `{"id":"s1","userId":"u","name":"n","cronExpression":"* * * * *","isEnabled":true,"uiState":null}`, &schedule)
	if schedule.ID != "s1" || string(schedule.Extra["uiState"]) != "null" {
		t.Errorf("unexpected schedule %+v", schedule)
	}

	backend := &mockBackend{responses: constant(200, `{"data":{"id":"r1","url":"https://example.com/","uniqueKey":"k","handledAt":null}}`)}
	got, present, err := testClient(backend, 0).RequestQueue("q").GetRequest(context.Background(), "r1")
	if err != nil || !present || got.ID != "r1" {
		t.Fatalf("GetRequest of a pending request: %+v, %v, %v", got, present, err)
	}
}
