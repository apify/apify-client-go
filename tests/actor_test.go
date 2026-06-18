package apify_test

import (
	"testing"

	apify "github.com/apify/apify-client-go"
)

// minimalActor returns a minimal Actor definition with the given name, suitable for create.
// The API requires at least one version, so a trivial source-files version "0.0" is embedded.
func minimalActor(name string) map[string]any {
	return map[string]any{
		"name":        name,
		"isPublic":    false,
		"description": "Integration test actor",
		"versions": []any{
			map[string]any{
				"versionNumber": "0.0",
				"sourceType":    "SOURCE_FILES",
				"buildTag":      "latest",
				"sourceFiles": []any{
					map[string]any{
						"name":    "Dockerfile",
						"format":  "TEXT",
						"content": "FROM apify/actor-node:20\nCOPY . ./\nCMD node main.js",
					},
					map[string]any{
						"name":    "main.js",
						"format":  "TEXT",
						"content": "console.log('hello from go client test');",
					},
				},
			},
		},
	}
}

func TestListActors(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	page, err := client.Actors().List(ctx, apify.ActorListOptions{My: ptr(true), Limit: ptr(int64(5))})
	if err != nil {
		t.Fatalf("Actors().List: %v", err)
	}
	if page.Total < 0 {
		t.Fatalf("unexpected total: %d", page.Total)
	}
}

func TestGetActor(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("get")))
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()

	got, ok, err := client.Actor(created.ID).Get(ctx)
	if err != nil {
		t.Fatalf("get actor: %v", err)
	}
	if !ok || got.ID != created.ID {
		t.Fatalf("expected to fetch created actor, got ok=%v %+v", ok, got)
	}
}

func TestActorCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("crud")))
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()

	actor := client.Actor(created.ID)

	if _, ok, err := actor.Get(ctx); err != nil || !ok {
		t.Fatalf("get: ok=%v err=%v", ok, err)
	}

	updated, err := actor.Update(ctx, map[string]any{"title": "Updated Title"})
	if err != nil {
		t.Fatalf("update: %v", err)
	}
	if updated.Title != "Updated Title" {
		t.Fatalf("title not updated: %+v", updated)
	}

	if _, err := actor.Builds().List(ctx, apify.ListOptions{}); err != nil {
		t.Fatalf("builds list: %v", err)
	}
	if _, err := actor.Versions().List(ctx, apify.ListOptions{}); err != nil {
		t.Fatalf("versions list: %v", err)
	}
}

func TestActorVersionCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("ver")))
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()
	actor := client.Actor(created.ID)

	versionDef := map[string]any{
		"versionNumber": "0.1",
		"sourceType":    "SOURCE_FILES",
		"buildTag":      "latest",
		"sourceFiles":   []any{},
	}
	version, err := actor.Versions().Create(ctx, versionDef)
	if err != nil {
		t.Fatalf("create version: %v", err)
	}
	if version.VersionNumber != "0.1" {
		t.Fatalf("unexpected version: %+v", version)
	}

	if _, ok, err := actor.Version("0.1").Get(ctx); err != nil || !ok {
		t.Fatalf("get version: ok=%v err=%v", ok, err)
	}
	if _, err := actor.Versions().List(ctx, apify.ListOptions{}); err != nil {
		t.Fatalf("list versions: %v", err)
	}
	if _, err := actor.Version("0.1").Update(ctx, map[string]any{"buildTag": "beta", "sourceType": "SOURCE_FILES", "sourceFiles": []any{}}); err != nil {
		t.Fatalf("update version: %v", err)
	}
	if err := actor.Version("0.1").Delete(ctx); err != nil {
		t.Fatalf("delete version: %v", err)
	}
}

func TestActorEnvVarCRUDFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("env")))
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()
	actor := client.Actor(created.ID)

	// The actor was created with version "0.0"; attach env vars to it.
	envVars := actor.Version("0.0").EnvVars()

	if _, err := envVars.Create(ctx, apify.ActorEnvVar{Name: "MY_VAR", Value: "value1"}); err != nil {
		t.Fatalf("create env var: %v", err)
	}
	if _, ok, err := actor.Version("0.0").EnvVar("MY_VAR").Get(ctx); err != nil || !ok {
		t.Fatalf("get env var: ok=%v err=%v", ok, err)
	}
	if _, err := envVars.List(ctx); err != nil {
		t.Fatalf("list env vars: %v", err)
	}
	if _, err := actor.Version("0.0").EnvVar("MY_VAR").Update(ctx, apify.ActorEnvVar{Name: "MY_VAR", Value: "value2"}); err != nil {
		t.Fatalf("update env var: %v", err)
	}
	if err := actor.Version("0.0").EnvVar("MY_VAR").Delete(ctx); err != nil {
		t.Fatalf("delete env var: %v", err)
	}
}
