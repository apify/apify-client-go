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

// The users endpoints can answer with the reduced proxy view of a user, which has no ID or
// username: it parses, and its fields are kept in Extra.
func TestUserProxyInfoParses(t *testing.T) {
	var user User
	mustUnmarshal(t, `{"plan":{"availableProxyGroups":["RESIDENTIAL"]},"effectivePlatformFeatures":{},"isPaying":true}`, &user)
	if user.ID != "" || string(user.Extra["isPaying"]) != "true" || user.Extra["plan"] == nil {
		t.Errorf("unexpected user %+v", user)
	}
}

// A Store Actor carries its categories, the Store statistics and the verified flag as typed
// fields: the statistics that older Store responses lacked stay nil.
func TestActorStoreListItemParsesStatsAndCategories(t *testing.T) {
	var item ActorStoreListItem
	mustUnmarshal(t, `{"id":"a1","name":"n","username":"u","title":"T","categories":["AI_AND_ML","ECOMMERCE"],
		"stats":{"totalUsers":6,"totalRuns":16,"actorReviewRating":4.7,"actorReviewCount":69},
		"isVerified":true,"badge":null}`, &item)
	if len(item.Categories) != 2 || item.Categories[0] != "AI_AND_ML" {
		t.Errorf("unexpected categories %v", item.Categories)
	}
	stats := item.Stats
	if stats.TotalUsers != 6 || stats.ActorReviewRating != 4.7 || stats.ActorReviewCount != 69 {
		t.Errorf("unexpected required stats %+v", stats)
	}
	if stats.TotalRuns == nil || *stats.TotalRuns != 16 || stats.TotalUsers7Days != nil {
		t.Errorf("unexpected optional stats %+v", stats)
	}
	if item.IsVerified == nil || !*item.IsVerified || string(item.Extra["badge"]) != "null" {
		t.Errorf("unexpected item %+v", item)
	}

	var minimal ActorStoreListItem
	mustUnmarshal(t, `{"id":"a2","name":"n","username":"u","title":"T","categories":null,"stats":{"totalUsers":0,"actorReviewRating":0,"actorReviewCount":0}}`, &minimal)
	if minimal.Categories != nil || minimal.IsVerified != nil {
		t.Errorf("unexpected minimal item %+v", minimal)
	}
}
