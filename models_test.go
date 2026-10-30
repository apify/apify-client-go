package apify

import (
	"context"
	"testing"
)

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
