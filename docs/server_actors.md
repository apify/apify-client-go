# Server Actors

A [Server Actor](https://docs.apify.com/platform/actors/running/server-actors) (formerly Actor Standby)
runs a web server. You call it over HTTP, and the platform keeps a run ready to answer, so each
call skips the run startup.

> **Formerly Actor Standby.** The API still uses the old name, for example in the Actor's
> `standbyUrl` and `actorStandby` fields. In the client, `ActorClient.Standby()` is a deprecated
> alias of `ActorClient.Server()`.

`client.Actor(id).Server()` returns the `*ActorServerClient` of the Actor. Every call returns the
same client, so the server URL it fetches is reused. The client is safe for concurrent use.

| Method | Description |
| --- | --- |
| `GetURL(ctx) (string, error)` | The URL of the Actor's server (the Actor's `standbyUrl`). |
| `GetOpenAPISpec(ctx) (json.RawMessage, error)` | The OpenAPI specification of the server, or `nil` if the Actor's definition has none. |
| `Request(ctx, path string, ActorServerRequestOptions) (*ActorServerResponse, error)` | Send a request and read the whole response. |
| `RequestStream(ctx, path string, ActorServerRequestOptions) (*ActorServerStream, error)` | Send a request and return the response body unread. Close `Body` when done. |

`Request` takes the path on the Actor's server and optionally the HTTP method, query parameters,
headers and a body:

```go
server := client.Actor("username/actor-name").Server()

response, err := server.Request(ctx, "/search", apify.ActorServerRequestOptions{
	Method: http.MethodPost,
	Params: url.Values{"q": {"apify"}},
	Body:   map[string]any{"limit": 5},
})
if err != nil {
	log.Fatal(err)
}
if response.StatusCode != http.StatusOK {
	log.Fatalf("the server answered %d: %s", response.StatusCode, response.Body)
}
var results map[string]any
if err := response.JSON(&results); err != nil {
	log.Fatal(err)
}
fmt.Println(results)
```

## Calling a Server Actor from a Server Actor

A multi-tenant Server Actor shares its runs among its users, so the platform tells the requests
apart by the `X-Actor-Request-ID` header of each request it forwards to the Actor. When your
Server Actor calls another one while handling a request, pass that header in `ParentRequestID`.
The client sends it as the same header, so the called Actor charges at the pricing tier of your
incoming request. Use a client created with the run's `APIFY_TOKEN`. For details, see
[Call another multi-tenant Server Actor](https://docs.apify.com/platform/actors/monetize/server-actors#call-another-multi-tenant-server-actor).

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	response, err := client.Actor("username/other-actor").Server().Request(r.Context(), "/search", apify.ActorServerRequestOptions{
		Params:          url.Values{"q": {"apify"}},
		ParentRequestID: r.Header.Get("X-Actor-Request-ID"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(response.StatusCode)
	_, _ = w.Write(response.Body)
})
_ = handler
```

## Request options

| Field | Type | Meaning |
|---|---|---|
| `Method` | `string` | The HTTP method, in upper case. Defaults to `GET`. An unknown method fails before any request is sent. |
| `Params` | `url.Values` | Query parameters. |
| `Headers` | `http.Header` | Additional headers. They reach the Actor's server as they are and win over the headers the client sets; a header with no values removes the client's header of that name. |
| `Body` | `any` | A `string` or `[]byte` is sent as it is. A `json.RawMessage` and any other value are sent as JSON with `Content-Type: application/json`, unless `Headers` sets a content type. |
| `ParentRequestID` | `string` | The `X-Actor-Request-ID` header of the request your Server Actor is handling. It is sent on as the same header, so the called Actor charges at the pricing tier of that request, and it wins over the same header in `Headers`. Use a client created with the run's `APIFY_TOKEN`. |
| `Compression` | `*apify.RequestCompression` | Compresses a body of 1 KiB or more. Nil by default, which sends the body uncompressed, since the Actor's server may not decode a compressed body. Set it only if the server decodes the resulting `Content-Encoding`. A `Content-Encoding` header in `Headers` sends the body untouched. |

The response has the `StatusCode`, the `Header` and the `Body` bytes, and `JSON(&v)` decodes the
body.

## How the request is sent

- The client takes the server's URL from the Actor's `standbyUrl`. It fetches the Actor on the
  first call and reuses the URL for later calls through the same client. Concurrent first calls
  share one API request, and a failed lookup is not kept. If the Actor is not a Server Actor, the
  methods return an error that wraps `apify.ErrNotServerActor`; if it does not exist, an
  `*apify.APIError` with `IsNotFound()`.
- The path is always resolved against that URL, so a path cannot send the request to another
  host. A fragment (`#...`) is dropped, since it is never sent to the server. Pass the query in
  `Params`.
- The request carries the client's API token in the `X-Apify-Authorization` header. The platform
  removes the token before it forwards the request, so the token does not reach the Actor's
  code. If the Actor reads the `Authorization` header for its own authentication, pass that
  header in `Headers`. For details, see
  [Authenticate requests](https://docs.apify.com/platform/actors/running/server-actors#authenticate-requests)
  in the platform documentation.
- Redirects are not followed. A redirect response is returned as it is, since following it would
  send the token to a host the Actor's code picked. A custom `HTTPBackend` that follows redirects
  must check `apify.FollowsRedirects(req)`, as the default backend does.

## Timeouts, retries and errors

When no server run is ready, the first request waits for the platform to start one, so requests
use the "long" timeout tier. As for every method, `ctx` can only shorten that bound. For
`RequestStream`, the tier bounds only the wait for the server to start responding.

The methods return the server's response whatever its status, so check `StatusCode` before you
use the body. Only rate limit (429) and Service Unavailable (503) responses are retried, with the
client's retry settings. Once the retries run out, the last response is returned. The platform
sends these when it cannot pass the request to a server run, but the Actor's code can return them
too, and then a retry sends the request to the Actor again. A transport error, a timeout
included, is not retried, since the Actor may have already handled the request.

The methods return an error only when they get no response from the server: the Actor does not
exist, it is not a Server Actor, the options are invalid, or the request fails or times out.

## Server URL and OpenAPI specification

`GetURL` returns the server's URL, fetching and reusing it the same way as `Request`. Use it to
call the server with another HTTP client.

`GetOpenAPISpec` returns the server's OpenAPI specification. It is the
[web server schema](https://docs.apify.com/platform/actors/development/actor-definition/web-server-schema)
from the Actor's definition, taken from the build that serves the server interface: the build set
in the Actor's `actorStandby.build` field (a build tag or a build number), or `latest` if none is
set. It returns `nil` if the Actor's definition has no web server schema, and an error if the build
does not exist.

```go
server := client.Actor("username/actor-name").Server()

serverURL, err := server.GetURL(ctx)
if err != nil {
	log.Fatal(err)
}
spec, err := server.GetOpenAPISpec(ctx)
if err != nil {
	log.Fatal(err)
}
fmt.Println("server:", serverURL)
if spec == nil {
	log.Fatal("the Actor's definition has no web server schema")
}
var document struct {
	Paths map[string]json.RawMessage `json:"paths"`
}
if err := json.Unmarshal(spec, &document); err != nil {
	log.Fatal(err)
}
for path := range document.Paths {
	fmt.Println("endpoint:", path)
}
```
