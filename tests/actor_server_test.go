package apify_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	apify "github.com/apify/apify-client-go"
)

const echoOpenAPISpec = `{"openapi":"3.0.0","info":{"title":"Echo server","version":"1.0.0"},"paths":{"/echo":{"post":{"responses":{"200":{"description":"The echoed request."}}}}}}`

// echoServerSource is a web server that answers every request, including the platform's
// readiness probe, with the request it received, under the status in the `status` query
// parameter, or 200.
const echoServerSource = `
const http = require('http');

http.createServer((req, res) => {
    let body = '';
    req.on('data', (chunk) => { body += chunk; });
    req.on('end', () => {
        const status = Number(new URL(req.url, 'http://localhost').searchParams.get('status')) || 200;
        res.writeHead(status, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ method: req.method, url: req.url, headers: req.headers, body }));
    });
}).listen(process.env.ACTOR_WEB_SERVER_PORT);
`

const echoDockerfile = `
FROM node:22-alpine
COPY main.js ./
CMD ["node", "main.js"]
`

// echoResponse is the body of the echo server's answers.
type echoResponse struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// TestActorServerFlow builds a Server Actor that echoes its requests, then gets its URL and
// OpenAPI specification and sends requests to it.
func TestActorServerFlow(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	name := uniqueName("server")
	actorJSON, err := json.Marshal(map[string]any{
		"actorSpecification": 1, "name": name, "version": "0.0", "usesStandbyMode": true,
		"webServerSchema": json.RawMessage(echoOpenAPISpec),
	})
	if err != nil {
		t.Fatal(err)
	}
	created, err := client.Actors().Create(ctx, map[string]any{
		"name": name,
		"versions": []any{map[string]any{
			"versionNumber": "0.0", "sourceType": "SOURCE_FILES", "buildTag": "latest",
			"sourceFiles": []any{
				map[string]any{"name": "Dockerfile", "format": "TEXT", "content": echoDockerfile},
				map[string]any{"name": "main.js", "format": "TEXT", "content": echoServerSource},
				map[string]any{"name": ".actor/actor.json", "format": "TEXT", "content": string(actorJSON)},
			},
		}},
		"actorStandby": map[string]any{"isEnabled": true, "memoryMbytes": 256, "idleTimeoutSecs": 60},
	})
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()

	build, err := client.Actor(created.ID).Build(ctx, "0.0", apify.ActorBuildOptions{})
	if err != nil {
		t.Fatalf("start build: %v", err)
	}
	finished, err := client.Build(build.ID).WaitForFinish(ctx, ptr(int64(300)))
	if err != nil || finished.Status != "SUCCEEDED" {
		t.Fatalf("build: status=%q err=%v", finished.Status, err)
	}

	actor, ok, err := client.Actor(created.ID).Get(ctx)
	if err != nil || !ok || actor.StandbyURL == nil {
		t.Fatalf("get actor: ok=%v standbyUrl=%v err=%v", ok, actor.StandbyURL, err)
	}
	server := client.Actor(created.ID).Server()

	serverURL, err := server.GetURL(ctx)
	if err != nil || serverURL != *actor.StandbyURL {
		t.Fatalf("GetURL = %q, %v; want %q", serverURL, err, *actor.StandbyURL)
	}
	spec, err := server.GetOpenAPISpec(ctx)
	if err != nil {
		t.Fatalf("GetOpenAPISpec: %v", err)
	}
	var gotSpec, wantSpec any
	if err := json.Unmarshal(spec, &gotSpec); err != nil {
		t.Fatalf("spec is not JSON: %v", err)
	}
	_ = json.Unmarshal([]byte(echoOpenAPISpec), &wantSpec)
	gotJSON, _ := json.Marshal(gotSpec)
	wantJSON, _ := json.Marshal(wantSpec)
	if string(gotJSON) != string(wantJSON) {
		t.Fatalf("spec = %s, want %s", gotJSON, wantJSON)
	}

	resp, err := server.Request(ctx, "/echo", apify.ActorServerRequestOptions{
		Method: http.MethodPost, Params: url.Values{"q": {"apify"}}, Body: map[string]int{"limit": 5},
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("request: %+v, %v", resp, err)
	}
	var echoed echoResponse
	if err := resp.JSON(&echoed); err != nil {
		t.Fatal(err)
	}
	if echoed.Method != "POST" || echoed.URL != "/echo?q=apify" || echoed.Body != `{"limit":5}` {
		t.Fatalf("echoed request: %+v", echoed)
	}
	// The platform consumes the client's token, so it never reaches the Actor's code.
	for _, header := range []string{"x-apify-authorization", "authorization"} {
		if _, present := echoed.Headers[header]; present {
			t.Fatalf("the %s header reached the Actor: %v", header, echoed.Headers)
		}
	}

	own, err := server.Request(ctx, "/echo", apify.ActorServerRequestOptions{Headers: http.Header{"Authorization": {"Bearer actor-token"}}})
	if err != nil {
		t.Fatal(err)
	}
	if err := own.JSON(&echoed); err != nil || echoed.Headers["authorization"] != "Bearer actor-token" {
		t.Fatalf("the caller's Authorization header did not reach the Actor: %+v, %v", echoed, err)
	}

	errorResponse, err := server.Request(ctx, "/echo", apify.ActorServerRequestOptions{Params: url.Values{"status": {"404"}}})
	if err != nil || errorResponse.StatusCode != http.StatusNotFound {
		t.Fatalf("an error status must be returned, got %+v, %v", errorResponse, err)
	}

	stream, err := server.RequestStream(ctx, "/echo", apify.ActorServerRequestOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = stream.Body.Close() }()
	body, err := io.ReadAll(stream.Body)
	if err != nil || stream.StatusCode != http.StatusOK || !strings.Contains(string(body), `"method":"GET"`) {
		t.Fatalf("stream: %d %s %v", stream.StatusCode, body, err)
	}
}

// TestActorServerErrors checks the errors for an Actor without a server
// and for a missing one.
func TestActorServerErrors(t *testing.T) {
	client := requireClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	created, err := client.Actors().Create(ctx, minimalActor(uniqueName("noserver")))
	if err != nil {
		t.Fatalf("create actor: %v", err)
	}
	defer func() { _ = client.Actor(created.ID).Delete(ctx) }()

	if _, err := client.Actor(created.ID).Server().GetURL(ctx); !errors.Is(err, apify.ErrNotServerActor) {
		t.Fatalf("expected ErrNotServerActor, got %v", err)
	}
	_, err = client.Actor(uniqueName("missing")).Server().Request(ctx, "/", apify.ActorServerRequestOptions{})
	if apiErr, ok := apify.AsAPIError(err); !ok || !apiErr.IsNotFound() {
		t.Fatalf("expected a not-found API error, got %v", err)
	}
}
