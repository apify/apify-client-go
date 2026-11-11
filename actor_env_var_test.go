package apify

import (
	"encoding/json"
	"testing"
)

// The API requires `value` on env var create/update, so an empty value must still be sent.
func TestActorEnvVarSendsEmptyValue(t *testing.T) {
	body, err := json.Marshal(ActorEnvVar{Name: "EMPTY"})
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"name":"EMPTY","value":""}` {
		t.Fatalf("unexpected body: %s", body)
	}
}
