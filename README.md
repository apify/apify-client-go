# Go API client for apifyclient


The Apify API (version 2) provides programmatic access to the [Apify
platform](https://docs.apify.com). The API is organized
around [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer)
HTTP endpoints.

You can download the complete OpenAPI schema of Apify API in the [YAML](http://docs.apify.com/api/openapi.yaml) or [JSON](http://docs.apify.com/api/openapi.json) formats. The source code is also available on [GitHub](https://github.com/apify/apify-docs/tree/master/apify-api/openapi).

All requests and responses (including errors) are encoded in
[JSON](http://www.json.org/) format with UTF-8 encoding,
with a few exceptions that are explicitly described in the reference.

- To access the API using [Node.js](https://nodejs.org/en/), we recommend the [`apify-client`](https://docs.apify.com/api/client/js) [NPM
package](https://www.npmjs.com/package/apify-client).
- To access the API using [Python](https://www.python.org/), we recommend the [`apify-client`](https://docs.apify.com/api/client/python) [PyPI
package](https://pypi.org/project/apify-client/).

The clients' functions correspond to the API endpoints and have the same
parameters. This simplifies development of apps that depend on the Apify
platform.

:::note Important Request Details

- `Content-Type` header: For requests with a JSON body, you must include the `Content-Type: application/json` header.

- Method override: You can override the HTTP method using the `method` query parameter. This is useful for clients that can only send `GET` requests. For example, to call a `POST` endpoint, append `?method=POST` to the URL of your `GET` request.

:::

## Authentication
<span id=\"/introduction/authentication\"></span>

**You can find your API token on the
[Integrations](https://console.apify.com/settings/integrations) page in the
Apify Console.**

To use your token in a request, either:

- Add the token to your request's `Authorization` header as `Bearer <token>`.
E.g., `Authorization: Bearer xxxxxxx`.
[More info](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization).
(Recommended).
- Add it as the `token` parameter to your request URL. (Less secure).

Using your token in the request header is more secure than using it as a URL
parameter because URLs are often stored
in browser history and server logs. This creates a chance for someone
unauthorized to access your API token.

**Never share your API token or password with untrusted parties!**

For more information, see our
[integrations](https://docs.apify.com/platform/integrations) documentation.

### Agentic payments

AI agents can authenticate and pay for Actor runs without an Apify account
using agentic payments. Instead of an API token, the request carries a
payment credential that both authorizes and pays for the call. Apify supports
the [x402 protocol](https://docs.apify.com/platform/integrations/x402)
(`PAYMENT-SIGNATURE` header) and
[Skyfire](https://docs.apify.com/platform/integrations/skyfire)
(`skyfire-pay-id` header).

## Basic usage
<span id=\"/introduction/basic-usage\"></span>

To run an Actor, send a POST request to the [Run
Actor](#/reference/actors/run-collection/run-actor) endpoint using either the
Actor ID code (e.g. `vKg4IjxZbEYTYeW8T`) or its name (e.g.
`janedoe~my-actor`):

`https://api.apify.com/v2/actors/[actor_id]/runs`

If the Actor is not runnable anonymously, you will receive a 401 or 403
[response code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status).
This means you need to add your [secret API
token](https://console.apify.com/account#/integrations) to the request's
`Authorization` header ([recommended](#/introduction/authentication)) or as a
URL query parameter `?token=[your_token]` (less secure).

Optionally, you can include the query parameters described in the [Run
Actor](#/reference/actors/run-collection/run-actor) section to customize your
run.

If you're using Node.js, the best way to run an Actor is using the
`Apify.call()` method from the [Apify
SDK](https://sdk.apify.com/docs/api/apify#apifycallactid-input-options). It
runs the Actor using the account you are currently logged into (determined
by the [secret API token](https://console.apify.com/account#/integrations)).
The result is an [Actor run
object](https://sdk.apify.com/docs/typedefs/actor-run) and its output (if
any).

A typical workflow is as follows:

1. Run an Actor or task using the [Run
Actor](#/reference/actors/run-collection/run-actor) or [Run
task](#/reference/actor-tasks/run-collection/run-task) API endpoints.
2. Monitor the Actor run by periodically polling its progress using the [Get
run](#/reference/actor-runs/run-object-and-its-storages/get-run) API
endpoint.
3. Fetch the results from the [Get
items](#/reference/datasets/item-collection/get-items) API endpoint using the
`defaultDatasetId`, which you receive in the Run request response.
Additional data may be stored in a key-value store. You can fetch them from
the [Get record](#/reference/key-value-stores/record/get-record) API endpoint
using the `defaultKeyValueStoreId` and the store's `key`.

**Note**: Instead of periodic polling, you can also run your
[Actor](#/reference/actors/run-actor-synchronously) or
[task](#/reference/actor-tasks/runs-collection/run-task-synchronously)
synchronously. This will ensure that the request waits for 300 seconds (5
minutes) for the run to finish and returns its output. If the run takes
longer, the request will time out and throw an error.

## Legacy `/v2/acts/` URL prefix
<span id=\"/introduction/legacy-acts-prefix\"></span>

The `/v2/acts/` prefix is deprecated but still fully functional, and 
such endpoint routes to the same handler as its `/v2/actors/...` counterpart. 
New integrations should use the canonical /v2/actors/ prefix, 
but existing clients keep working without changes.

## Response structure
<span id=\"/introduction/response-structure\"></span>

Most API endpoints return a JSON object with the `data` property:

```
{
    \"data\": {
        ...
    }
}
```

However, there are a few explicitly described exceptions, such as
[Get dataset items](#/reference/datasets/item-collection/get-items) or
Key-value store [Get record](#/reference/key-value-stores/record/get-record)
API endpoints, which return data in other formats.
In case of an error, the response has the HTTP status code in the range of
4xx or 5xx and the `data` property is replaced with `error`. For example:

```
{
    \"error\": {
        \"type\": \"record-not-found\",
        \"message\": \"Store was not found.\"
    }
}
```

See [Errors](#/introduction/errors) for more details.

## Pagination
<span id=\"/introduction/pagination\"></span>

All API endpoints that return a list of records
(e.g. [Get list of
Actors](#/reference/actors/actor-collection/get-list-of-actors))
enforce pagination in order to limit the size of their responses.

Most of these API endpoints are paginated using the `offset` and `limit`
query parameters.
The only exception is [Get list of
keys](#/reference/key-value-stores/key-collection/get-list-of-keys),
which is paginated using the `exclusiveStartKey` query parameter.

**IMPORTANT**: Each API endpoint that supports pagination enforces a certain
maximum value for the `limit` parameter,
in order to reduce the load on Apify servers.
The maximum limit could change in future so you should never
rely on a specific value and check the responses of these API endpoints.

### Using offset
<span id=\"/introduction/pagination/using-offset\"></span>

Most API endpoints that return a list of records enable pagination using the
following query parameters:

<table>
  <tr>
    <td><code>limit</code></td>
    <td>Limits the response to contain a specific maximum number of items, e.g. <code>limit=20</code>.</td>
  </tr>
  <tr>
    <td><code>offset</code></td>
    <td>Skips a number of items from the beginning of the list, e.g. <code>offset=100</code>.</td>
  </tr>
  <tr>
    <td><code>desc</code></td>
    <td>
    By default, items are sorted in the order in which they were created or added to the list.
    This feature is useful when fetching all the items, because it ensures that items
    created after the client started the pagination will not be skipped.
    If you specify the <code>desc=1</code> parameter, the items will be returned in the reverse order,
    i.e. from the newest to the oldest items.
    </td>
  </tr>
</table>

The response of these API endpoints is always a JSON object with the
following structure:

```
{
    \"data\": {
        \"total\": 2560,
        \"offset\": 250,
        \"limit\": 1000,
        \"count\": 1000,
        \"desc\": false,
        \"items\": [
            { 1st object },
            { 2nd object },
            ...
            { 1000th object }
        ]
    }
}
```

The following table describes the meaning of the response properties:

<table>
  <tr>
    <th>Property</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><code>total</code></td>
    <td>The total number of items available in the list.</td>
  </tr>
  <tr>
    <td><code>offset</code></td>
    <td>The number of items that were skipped at the start.
    This is equal to the <code>offset</code> query parameter if it was provided, otherwise it is <code>0</code>.</td>
  </tr>
  <tr>
    <td><code>limit</code></td>
    <td>The maximum number of items that can be returned in the HTTP response.
    It equals to the <code>limit</code> query parameter if it was provided or
    the maximum limit enforced for the particular API endpoint, whichever is smaller.</td>
  </tr>
  <tr>
    <td><code>count</code></td>
    <td>The actual number of items returned in the HTTP response.</td>
  </tr>
  <tr>
    <td><code>desc</code></td>
    <td><code>true</code> if data were requested in descending order and <code>false</code> otherwise.</td>
  </tr>
  <tr>
    <td><code>items</code></td>
    <td>An array of requested items.</td>
  </tr>
</table>

### Using key
<span id=\"/introduction/pagination/using-key\"></span>

The records in the [key-value
store](https://docs.apify.com/platform/storage/key-value-store)
are not ordered based on numerical indexes,
but rather by their keys in the UTF-8 binary order.
Therefore the [Get list of
keys](#/reference/key-value-stores/key-collection/get-list-of-keys)
API endpoint only supports pagination using the following query parameters:

<table>
  <tr>
    <td><code>limit</code></td>
    <td>Limits the response to contain a specific maximum number items, e.g. <code>limit=20</code>.</td>
  </tr>
  <tr>
    <td><code>exclusiveStartKey</code></td>
    <td>Skips all records with keys up to the given key including the given key,
    in the UTF-8 binary order.</td>
  </tr>
</table>

The response of the API endpoint is always a JSON object with following
structure:

```
{
    \"data\": {
        \"limit\": 1000,
        \"isTruncated\": true,
        \"exclusiveStartKey\": \"my-key\",
        \"nextExclusiveStartKey\": \"some-other-key\",
        \"items\": [
            { 1st object },
            { 2nd object },
            ...
            { 1000th object }
        ]
    }
}
```

The following table describes the meaning of the response properties:

<table>
  <tr>
    <th>Property</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><code>limit</code></td>
    <td>The maximum number of items that can be returned in the HTTP response.
    It equals to the <code>limit</code> query parameter if it was provided or
    the maximum limit enforced for the particular endpoint, whichever is smaller.</td>
  </tr>
  <tr>
    <td><code>isTruncated</code></td>
    <td><code>true</code> if there are more items left to be queried. Otherwise <code>false</code>.</td>
  </tr>
  <tr>
    <td><code>exclusiveStartKey</code></td>
    <td>The last key that was skipped at the start. Is `null` for the first page.</td>
  </tr>
  <tr>
    <td><code>nextExclusiveStartKey</code></td>
    <td>The value for the <code>exclusiveStartKey</code> parameter to query the next page of items.</td>
  </tr>
</table>

## Errors
<span id=\"/introduction/errors\"></span>

The Apify API uses common HTTP status codes: `2xx` range for success, `4xx`
range for errors caused by the caller
(invalid requests) and `5xx` range for server errors (these are rare).
Each error response contains a JSON object defining the `error` property,
which is an object with
the `type` and `message` properties that contain the error code and a
human-readable error description, respectively.

For example:

```
{
    \"error\": {
        \"type\": \"record-not-found\",
        \"message\": \"Store was not found.\"
    }
}
```

Here is the table of the most common errors that can occur for many API
endpoints:

<table>
  <tr>
    <th>status</th>
    <th>type</th>
    <th>message</th>
  </tr>
  <tr>
    <td><code>400</code></td>
    <td><code>invalid-request</code></td>
    <td>POST data must be a JSON object</td>
  </tr>
  <tr>
    <td><code>400</code></td>
    <td><code>invalid-value</code></td>
    <td>Invalid value provided: Comments required</td>
  </tr>
  <tr>
    <td><code>400</code></td>
    <td><code>invalid-record-key</code></td>
    <td>Record key contains invalid character</td>
  </tr>
  <tr>
    <td><code>401</code></td>
    <td><code>token-not-provided</code></td>
    <td>Authentication token was not provided</td>
  </tr>
  <tr>
    <td><code>404</code></td>
    <td><code>record-not-found</code></td>
    <td>Store was not found</td>
  </tr>
  <tr>
    <td><code>429</code></td>
    <td><code>rate-limit-exceeded</code></td>
    <td>You have exceeded the rate limit of ... requests per second</td>
  </tr>
  <tr>
    <td><code>405</code></td>
    <td><code>method-not-allowed</code></td>
    <td>This API endpoint can only be accessed using the following HTTP methods: OPTIONS, POST</td>
  </tr>
</table>

## Rate limiting
<span id=\"/introduction/rate-limiting\"></span>

All API endpoints limit the rate of requests in order to prevent overloading of Apify servers by misbehaving clients.

There are two kinds of rate limits - a global rate limit and a per-resource rate limit.

### Global rate limit
<span id=\"/introduction/rate-limiting/global-rate-limit\"></span>

The global rate limit is set to _250 000 requests per minute_.
For [authenticated](#/introduction/authentication) requests, it is counted per user,
and for unauthenticated requests, it is counted per IP address.

### Per-resource rate limit
<span id=\"/introduction/rate-limiting/per-resource-rate-limit\"></span>

The default per-resource rate limit is _60 requests per second per resource_, which in this context means a single Actor, a single Actor run, a single dataset, single key-value store etc.
The default rate limit is applied to every API endpoint except a few select ones, which have higher rate limits.
Each API endpoint returns its rate limit in `X-RateLimit-Limit` header.

These endpoints have a rate limit of _200 requests per second per resource_:

* CRUD ([get](#/reference/key-value-stores/record/get-record),
  [put](#/reference/key-value-stores/record/put-record),
  [delete](#/reference/key-value-stores/record/delete-record))
  operations on key-value store records

These endpoints have a rate limit of _400 requests per second per resource_:
* [Run Actor](#/reference/actors/run-collection/run-actor)
* [Run Actor task asynchronously](#/reference/actor-tasks/runs-collection/run-task-asynchronously)
* [Run Actor task synchronously](#/reference/actor-tasks/runs-collection/run-task-synchronously)
* [Metamorph Actor run](#/reference/actors/metamorph-run/metamorph-run)
* [Push items](#/reference/datasets/item-collection/put-items) to dataset
* CRUD
  ([add](#/reference/request-queues/request-collection/add-request),
  [get](#/reference/request-queues/request-collection/get-request),
  [update](#/reference/request-queues/request-collection/update-request),
  [delete](#/reference/request-queues/request-collection/delete-request))
  operations on requests in request queues

### Rate limit exceeded errors
<span id=\"/introduction/rate-limiting/rate-limit-exceeded-errors\"></span>

If the client is sending too many requests, the API endpoints respond with the HTTP status code `429 Too Many Requests`
and the following body:

```
{
    \"error\": {
        \"type\": \"rate-limit-exceeded\",
        \"message\": \"You have exceeded the rate limit of ... requests per second\"
    }
}
```

### Retrying rate-limited requests with exponential backoff
<span id=\"/introduction/rate-limiting/retrying-rate-limited-requests-with-exponential-backoff\"></span>

If the client receives the rate limit error, it should wait a certain period of time and then retry the request.
If the error happens again, the client should double the wait period and retry the request,
and so on. This algorithm is known as _exponential backoff_
and it can be described using the following pseudo-code:

1. Define a variable `DELAY=500`
2. Send the HTTP request to the API endpoint
3. If the response has status code not equal to `429` then you are done. Otherwise:
   * Wait for a period of time chosen randomly from the interval `DELAY` to `2*DELAY` milliseconds
   * Double the future wait period by setting `DELAY = 2*DELAY`
   * Continue with step 2

If all requests sent by the client implement the above steps,
the client will automatically use the maximum available bandwidth for its requests.

Note that the Apify API clients [for JavaScript](https://docs.apify.com/api/client/js)
and [for Python](https://docs.apify.com/api/client/python)
use the exponential backoff algorithm transparently, so that you do not need to worry about it.

## Referring to resources
<span id=\"/introduction/referring-to-resources\"></span>

There are three main ways to refer to a resource you're accessing via API.

- the resource ID (e.g. `iKkPcIgVvwmztduf8`)
- `username~resourcename` - when using this access method, you will need to
use your API token, and access will only work if you have the correct
permissions.
- `~resourcename` - for this, you need to use an API token, and the
`resourcename` refers to a resource in the API token owner's account.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2-2026-06-16T064758Z
- Package version: 1.0.0
- Generator version: 7.23.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Import the package in a go file in your project and run `go mod tidy`:

```go
import apifyclient "github.com/apify/apify-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `apifyclient.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), apifyclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `apifyclient.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), apifyclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `apifyclient.ContextOperationServerIndices` and `apifyclient.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), apifyclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), apifyclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.apify.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActorBuildsAPI* | [**ActorBuildAbortPost**](docs/ActorBuildsAPI.md#actorbuildabortpost) | **Post** /v2/actor-builds/{buildId}/abort | Abort build
*ActorBuildsAPI* | [**ActorBuildDelete**](docs/ActorBuildsAPI.md#actorbuilddelete) | **Delete** /v2/actor-builds/{buildId} | Delete build
*ActorBuildsAPI* | [**ActorBuildGet**](docs/ActorBuildsAPI.md#actorbuildget) | **Get** /v2/actor-builds/{buildId} | Get build
*ActorBuildsAPI* | [**ActorBuildLogGet**](docs/ActorBuildsAPI.md#actorbuildlogget) | **Get** /v2/actor-builds/{buildId}/log | Get build&#39;s Log
*ActorBuildsAPI* | [**ActorBuildOpenapiJsonGet**](docs/ActorBuildsAPI.md#actorbuildopenapijsonget) | **Get** /v2/actor-builds/{buildId}/openapi.json | Get OpenAPI definition
*ActorBuildsAPI* | [**ActorBuildsGet**](docs/ActorBuildsAPI.md#actorbuildsget) | **Get** /v2/actor-builds | Get user builds list
*ActorRunsAPI* | [**ActorRunAbortPost**](docs/ActorRunsAPI.md#actorrunabortpost) | **Post** /v2/actor-runs/{runId}/abort | Abort run
*ActorRunsAPI* | [**ActorRunDelete**](docs/ActorRunsAPI.md#actorrundelete) | **Delete** /v2/actor-runs/{runId} | Delete run
*ActorRunsAPI* | [**ActorRunGet**](docs/ActorRunsAPI.md#actorrunget) | **Get** /v2/actor-runs/{runId} | Get run
*ActorRunsAPI* | [**ActorRunLogGet**](docs/ActorRunsAPI.md#actorrunlogget) | **Get** /v2/actor-runs/{runId}/log | Get run&#39;s log
*ActorRunsAPI* | [**ActorRunMetamorphPost**](docs/ActorRunsAPI.md#actorrunmetamorphpost) | **Post** /v2/actor-runs/{runId}/metamorph | Metamorph run
*ActorRunsAPI* | [**ActorRunPut**](docs/ActorRunsAPI.md#actorrunput) | **Put** /v2/actor-runs/{runId} | Update run
*ActorRunsAPI* | [**ActorRunRebootPost**](docs/ActorRunsAPI.md#actorrunrebootpost) | **Post** /v2/actor-runs/{runId}/reboot | Reboot run
*ActorRunsAPI* | [**ActorRunsGet**](docs/ActorRunsAPI.md#actorrunsget) | **Get** /v2/actor-runs | Get user runs list
*ActorRunsAPI* | [**PostChargeRun**](docs/ActorRunsAPI.md#postchargerun) | **Post** /v2/actor-runs/{runId}/charge | Charge events in run
*ActorRunsAPI* | [**PostResurrectRun**](docs/ActorRunsAPI.md#postresurrectrun) | **Post** /v2/actor-runs/{runId}/resurrect | Resurrect run
*ActorTasksAPI* | [**ActorTaskDelete**](docs/ActorTasksAPI.md#actortaskdelete) | **Delete** /v2/actor-tasks/{actorTaskId} | Delete task
*ActorTasksAPI* | [**ActorTaskGet**](docs/ActorTasksAPI.md#actortaskget) | **Get** /v2/actor-tasks/{actorTaskId} | Get task
*ActorTasksAPI* | [**ActorTaskInputGet**](docs/ActorTasksAPI.md#actortaskinputget) | **Get** /v2/actor-tasks/{actorTaskId}/input | Get task input
*ActorTasksAPI* | [**ActorTaskInputPut**](docs/ActorTasksAPI.md#actortaskinputput) | **Put** /v2/actor-tasks/{actorTaskId}/input | Update task input
*ActorTasksAPI* | [**ActorTaskPut**](docs/ActorTasksAPI.md#actortaskput) | **Put** /v2/actor-tasks/{actorTaskId} | Update task
*ActorTasksAPI* | [**ActorTaskRunSyncGet**](docs/ActorTasksAPI.md#actortaskrunsyncget) | **Get** /v2/actor-tasks/{actorTaskId}/run-sync | Run task synchronously
*ActorTasksAPI* | [**ActorTaskRunSyncGetDatasetItemsGet**](docs/ActorTasksAPI.md#actortaskrunsyncgetdatasetitemsget) | **Get** /v2/actor-tasks/{actorTaskId}/run-sync-get-dataset-items | Run task synchronously and get dataset items
*ActorTasksAPI* | [**ActorTaskRunSyncGetDatasetItemsPost**](docs/ActorTasksAPI.md#actortaskrunsyncgetdatasetitemspost) | **Post** /v2/actor-tasks/{actorTaskId}/run-sync-get-dataset-items | Run task synchronously and get dataset items
*ActorTasksAPI* | [**ActorTaskRunSyncPost**](docs/ActorTasksAPI.md#actortaskrunsyncpost) | **Post** /v2/actor-tasks/{actorTaskId}/run-sync | Run task synchronously
*ActorTasksAPI* | [**ActorTaskRunsGet**](docs/ActorTasksAPI.md#actortaskrunsget) | **Get** /v2/actor-tasks/{actorTaskId}/runs | Get list of task runs
*ActorTasksAPI* | [**ActorTaskRunsLastGet**](docs/ActorTasksAPI.md#actortaskrunslastget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last | Get last run
*ActorTasksAPI* | [**ActorTaskRunsPost**](docs/ActorTasksAPI.md#actortaskrunspost) | **Post** /v2/actor-tasks/{actorTaskId}/runs | Run task
*ActorTasksAPI* | [**ActorTaskWebhooksGet**](docs/ActorTasksAPI.md#actortaskwebhooksget) | **Get** /v2/actor-tasks/{actorTaskId}/webhooks | Get list of webhooks
*ActorTasksAPI* | [**ActorTasksGet**](docs/ActorTasksAPI.md#actortasksget) | **Get** /v2/actor-tasks | Get list of tasks
*ActorTasksAPI* | [**ActorTasksPost**](docs/ActorTasksAPI.md#actortaskspost) | **Post** /v2/actor-tasks | Create task
*ActorsAPI* | [**ActDelete**](docs/ActorsAPI.md#actdelete) | **Delete** /v2/actors/{actorId} | Delete Actor
*ActorsAPI* | [**ActGet**](docs/ActorsAPI.md#actget) | **Get** /v2/actors/{actorId} | Get Actor
*ActorsAPI* | [**ActPut**](docs/ActorsAPI.md#actput) | **Put** /v2/actors/{actorId} | Update Actor
*ActorsAPI* | [**ActValidateInputPost**](docs/ActorsAPI.md#actvalidateinputpost) | **Post** /v2/actors/{actorId}/validate-input | Validate Actor input
*ActorsAPI* | [**ActsGet**](docs/ActorsAPI.md#actsget) | **Get** /v2/actors | Get list of Actors
*ActorsAPI* | [**ActsPost**](docs/ActorsAPI.md#actspost) | **Post** /v2/actors | Create Actor
*ActorsActorBuildsAPI* | [**ActBuildAbortPost**](docs/ActorsActorBuildsAPI.md#actbuildabortpost) | **Post** /v2/actors/{actorId}/builds/{buildId}/abort | Abort build
*ActorsActorBuildsAPI* | [**ActBuildDefaultGet**](docs/ActorsActorBuildsAPI.md#actbuilddefaultget) | **Get** /v2/actors/{actorId}/builds/default | Get default build
*ActorsActorBuildsAPI* | [**ActBuildGet**](docs/ActorsActorBuildsAPI.md#actbuildget) | **Get** /v2/actors/{actorId}/builds/{buildId} | Get build
*ActorsActorBuildsAPI* | [**ActBuildsGet**](docs/ActorsActorBuildsAPI.md#actbuildsget) | **Get** /v2/actors/{actorId}/builds | Get list of builds
*ActorsActorBuildsAPI* | [**ActBuildsPost**](docs/ActorsActorBuildsAPI.md#actbuildspost) | **Post** /v2/actors/{actorId}/builds | Build Actor
*ActorsActorBuildsAPI* | [**ActOpenapiJsonGet**](docs/ActorsActorBuildsAPI.md#actopenapijsonget) | **Get** /v2/actors/{actorId}/builds/{buildId}/openapi.json | Get OpenAPI definition
*ActorsActorRunsAPI* | [**ActRunAbortPost**](docs/ActorsActorRunsAPI.md#actrunabortpost) | **Post** /v2/actors/{actorId}/runs/{runId}/abort | Abort run
*ActorsActorRunsAPI* | [**ActRunGet**](docs/ActorsActorRunsAPI.md#actrunget) | **Get** /v2/actors/{actorId}/runs/{runId} | Get run
*ActorsActorRunsAPI* | [**ActRunMetamorphPost**](docs/ActorsActorRunsAPI.md#actrunmetamorphpost) | **Post** /v2/actors/{actorId}/runs/{runId}/metamorph | Metamorph run
*ActorsActorRunsAPI* | [**ActRunResurrectPost**](docs/ActorsActorRunsAPI.md#actrunresurrectpost) | **Post** /v2/actors/{actorId}/runs/{runId}/resurrect | Resurrect run
*ActorsActorRunsAPI* | [**ActRunSyncGet**](docs/ActorsActorRunsAPI.md#actrunsyncget) | **Get** /v2/actors/{actorId}/run-sync | Run Actor synchronously without input
*ActorsActorRunsAPI* | [**ActRunSyncGetDatasetItemsGet**](docs/ActorsActorRunsAPI.md#actrunsyncgetdatasetitemsget) | **Get** /v2/actors/{actorId}/run-sync-get-dataset-items | Run Actor synchronously without input and get dataset items
*ActorsActorRunsAPI* | [**ActRunSyncGetDatasetItemsPost**](docs/ActorsActorRunsAPI.md#actrunsyncgetdatasetitemspost) | **Post** /v2/actors/{actorId}/run-sync-get-dataset-items | Run Actor synchronously and get dataset items
*ActorsActorRunsAPI* | [**ActRunSyncPost**](docs/ActorsActorRunsAPI.md#actrunsyncpost) | **Post** /v2/actors/{actorId}/run-sync | Run Actor synchronously and return output
*ActorsActorRunsAPI* | [**ActRunsGet**](docs/ActorsActorRunsAPI.md#actrunsget) | **Get** /v2/actors/{actorId}/runs | Get list of runs
*ActorsActorRunsAPI* | [**ActRunsLastGet**](docs/ActorsActorRunsAPI.md#actrunslastget) | **Get** /v2/actors/{actorId}/runs/last | Get last run
*ActorsActorRunsAPI* | [**ActRunsPost**](docs/ActorsActorRunsAPI.md#actrunspost) | **Post** /v2/actors/{actorId}/runs | Run Actor
*ActorsActorVersionsAPI* | [**ActVersionDelete**](docs/ActorsActorVersionsAPI.md#actversiondelete) | **Delete** /v2/actors/{actorId}/versions/{versionNumber} | Delete version
*ActorsActorVersionsAPI* | [**ActVersionEnvVarDelete**](docs/ActorsActorVersionsAPI.md#actversionenvvardelete) | **Delete** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Delete environment variable
*ActorsActorVersionsAPI* | [**ActVersionEnvVarGet**](docs/ActorsActorVersionsAPI.md#actversionenvvarget) | **Get** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Get environment variable
*ActorsActorVersionsAPI* | [**ActVersionEnvVarPost**](docs/ActorsActorVersionsAPI.md#actversionenvvarpost) | **Post** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Update environment variable (POST)
*ActorsActorVersionsAPI* | [**ActVersionEnvVarPut**](docs/ActorsActorVersionsAPI.md#actversionenvvarput) | **Put** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Update environment variable
*ActorsActorVersionsAPI* | [**ActVersionEnvVarsGet**](docs/ActorsActorVersionsAPI.md#actversionenvvarsget) | **Get** /v2/actors/{actorId}/versions/{versionNumber}/env-vars | Get list of environment variables
*ActorsActorVersionsAPI* | [**ActVersionEnvVarsPost**](docs/ActorsActorVersionsAPI.md#actversionenvvarspost) | **Post** /v2/actors/{actorId}/versions/{versionNumber}/env-vars | Create environment variable
*ActorsActorVersionsAPI* | [**ActVersionGet**](docs/ActorsActorVersionsAPI.md#actversionget) | **Get** /v2/actors/{actorId}/versions/{versionNumber} | Get version
*ActorsActorVersionsAPI* | [**ActVersionPost**](docs/ActorsActorVersionsAPI.md#actversionpost) | **Post** /v2/actors/{actorId}/versions/{versionNumber} | Update version (POST)
*ActorsActorVersionsAPI* | [**ActVersionPut**](docs/ActorsActorVersionsAPI.md#actversionput) | **Put** /v2/actors/{actorId}/versions/{versionNumber} | Update version
*ActorsActorVersionsAPI* | [**ActVersionsGet**](docs/ActorsActorVersionsAPI.md#actversionsget) | **Get** /v2/actors/{actorId}/versions | Get list of versions
*ActorsActorVersionsAPI* | [**ActVersionsPost**](docs/ActorsActorVersionsAPI.md#actversionspost) | **Post** /v2/actors/{actorId}/versions | Create version
*ActorsWebhookCollectionAPI* | [**ActWebhooksGet**](docs/ActorsWebhookCollectionAPI.md#actwebhooksget) | **Get** /v2/actors/{actorId}/webhooks | Get list of webhooks
*DefaultDatasetAPI* | [**ActorRunDatasetDelete**](docs/DefaultDatasetAPI.md#actorrundatasetdelete) | **Delete** /v2/actor-runs/{runId}/dataset | Delete default dataset
*DefaultDatasetAPI* | [**ActorRunDatasetGet**](docs/DefaultDatasetAPI.md#actorrundatasetget) | **Get** /v2/actor-runs/{runId}/dataset | Get default dataset
*DefaultDatasetAPI* | [**ActorRunDatasetItemsGet**](docs/DefaultDatasetAPI.md#actorrundatasetitemsget) | **Get** /v2/actor-runs/{runId}/dataset/items | Get default dataset items
*DefaultDatasetAPI* | [**ActorRunDatasetItemsPost**](docs/DefaultDatasetAPI.md#actorrundatasetitemspost) | **Post** /v2/actor-runs/{runId}/dataset/items | Store items
*DefaultDatasetAPI* | [**ActorRunDatasetPut**](docs/DefaultDatasetAPI.md#actorrundatasetput) | **Put** /v2/actor-runs/{runId}/dataset | Update default dataset
*DefaultDatasetAPI* | [**ActorRunDatasetStatisticsGet**](docs/DefaultDatasetAPI.md#actorrundatasetstatisticsget) | **Get** /v2/actor-runs/{runId}/dataset/statistics | Get default dataset statistics
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreDelete**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestoredelete) | **Delete** /v2/actor-runs/{runId}/key-value-store | Delete default store
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreGet**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestoreget) | **Get** /v2/actor-runs/{runId}/key-value-store | Get default store
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreKeysGet**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestorekeysget) | **Get** /v2/actor-runs/{runId}/key-value-store/keys | Get default store&#39;s list of keys
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStorePut**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestoreput) | **Put** /v2/actor-runs/{runId}/key-value-store | Update default store
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreRecordDelete**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestorerecorddelete) | **Delete** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Delete default store&#39;s record
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreRecordGet**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestorerecordget) | **Get** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Get default store&#39;s record
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreRecordPost**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestorerecordpost) | **Post** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Store record in default store (POST)
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreRecordPut**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestorerecordput) | **Put** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Store record in default store
*DefaultKeyValueStoreAPI* | [**ActorRunKeyValueStoreRecordsGet**](docs/DefaultKeyValueStoreAPI.md#actorrunkeyvaluestorerecordsget) | **Get** /v2/actor-runs/{runId}/key-value-store/records | Download default store&#39;s records
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueDelete**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuedelete) | **Delete** /v2/actor-runs/{runId}/request-queue | Delete default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueGet**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueueget) | **Get** /v2/actor-runs/{runId}/request-queue | Get default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueHeadGet**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueueheadget) | **Get** /v2/actor-runs/{runId}/request-queue/head | Get default request queue head
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueHeadLockPost**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueueheadlockpost) | **Post** /v2/actor-runs/{runId}/request-queue/head/lock | Get and lock default request queue head
*DefaultRequestQueueAPI* | [**ActorRunRequestQueuePut**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueueput) | **Put** /v2/actor-runs/{runId}/request-queue | Update default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestDelete**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestdelete) | **Delete** /v2/actor-runs/{runId}/request-queue/requests/{requestId} | Delete request from default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestGet**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestget) | **Get** /v2/actor-runs/{runId}/request-queue/requests/{requestId} | Get request from default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestLockDelete**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestlockdelete) | **Delete** /v2/actor-runs/{runId}/request-queue/requests/{requestId}/lock | Delete lock on request in default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestLockPut**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestlockput) | **Put** /v2/actor-runs/{runId}/request-queue/requests/{requestId}/lock | Prolong lock on request in default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestPut**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestput) | **Put** /v2/actor-runs/{runId}/request-queue/requests/{requestId} | Update request in default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestsBatchDelete**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestsbatchdelete) | **Delete** /v2/actor-runs/{runId}/request-queue/requests/batch | Batch delete requests from default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestsBatchPost**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestsbatchpost) | **Post** /v2/actor-runs/{runId}/request-queue/requests/batch | Batch add requests to default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestsGet**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestsget) | **Get** /v2/actor-runs/{runId}/request-queue/requests | List default request queue&#39;s requests
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestsPost**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestspost) | **Post** /v2/actor-runs/{runId}/request-queue/requests | Add request to default request queue
*DefaultRequestQueueAPI* | [**ActorRunRequestQueueRequestsUnlockPost**](docs/DefaultRequestQueueAPI.md#actorrunrequestqueuerequestsunlockpost) | **Post** /v2/actor-runs/{runId}/request-queue/requests/unlock | Unlock requests in default request queue
*LastActorRunsAbortAPI* | [**ActRunsLastAbortPost**](docs/LastActorRunsAbortAPI.md#actrunslastabortpost) | **Post** /v2/actors/{actorId}/runs/last/abort | Abort Actor&#39;s last run
*LastActorRunsDefaultDatasetAPI* | [**ActRunsLastDatasetDelete**](docs/LastActorRunsDefaultDatasetAPI.md#actrunslastdatasetdelete) | **Delete** /v2/actors/{actorId}/runs/last/dataset | Delete last run&#39;s default dataset
*LastActorRunsDefaultDatasetAPI* | [**ActRunsLastDatasetGet**](docs/LastActorRunsDefaultDatasetAPI.md#actrunslastdatasetget) | **Get** /v2/actors/{actorId}/runs/last/dataset | Get last run&#39;s default dataset
*LastActorRunsDefaultDatasetAPI* | [**ActRunsLastDatasetItemsGet**](docs/LastActorRunsDefaultDatasetAPI.md#actrunslastdatasetitemsget) | **Get** /v2/actors/{actorId}/runs/last/dataset/items | Get last run&#39;s dataset items
*LastActorRunsDefaultDatasetAPI* | [**ActRunsLastDatasetItemsPost**](docs/LastActorRunsDefaultDatasetAPI.md#actrunslastdatasetitemspost) | **Post** /v2/actors/{actorId}/runs/last/dataset/items | Store items in last run&#39;s dataset
*LastActorRunsDefaultDatasetAPI* | [**ActRunsLastDatasetPut**](docs/LastActorRunsDefaultDatasetAPI.md#actrunslastdatasetput) | **Put** /v2/actors/{actorId}/runs/last/dataset | Update last run&#39;s default dataset
*LastActorRunsDefaultDatasetAPI* | [**ActRunsLastDatasetStatisticsGet**](docs/LastActorRunsDefaultDatasetAPI.md#actrunslastdatasetstatisticsget) | **Get** /v2/actors/{actorId}/runs/last/dataset/statistics | Get last run&#39;s dataset statistics
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreDelete**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestoredelete) | **Delete** /v2/actors/{actorId}/runs/last/key-value-store | Delete last run&#39;s default store
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreGet**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestoreget) | **Get** /v2/actors/{actorId}/runs/last/key-value-store | Get last run&#39;s default store
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreKeysGet**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestorekeysget) | **Get** /v2/actors/{actorId}/runs/last/key-value-store/keys | Get last run&#39;s default store&#39;s list of keys
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStorePut**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestoreput) | **Put** /v2/actors/{actorId}/runs/last/key-value-store | Update last run&#39;s default store
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreRecordDelete**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestorerecorddelete) | **Delete** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Delete last run&#39;s default store&#39;s record
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreRecordGet**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestorerecordget) | **Get** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Get last run&#39;s default store&#39;s record
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreRecordPost**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestorerecordpost) | **Post** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Store record in last run&#39;s default store (POST)
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreRecordPut**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestorerecordput) | **Put** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Store record in last run&#39;s default store
*LastActorRunsDefaultKeyValueStoreAPI* | [**ActRunsLastKeyValueStoreRecordsGet**](docs/LastActorRunsDefaultKeyValueStoreAPI.md#actrunslastkeyvaluestorerecordsget) | **Get** /v2/actors/{actorId}/runs/last/key-value-store/records | Download last run&#39;s default store&#39;s records
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueDelete**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuedelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue | Delete last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueGet**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueueget) | **Get** /v2/actors/{actorId}/runs/last/request-queue | Get last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueHeadGet**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueueheadget) | **Get** /v2/actors/{actorId}/runs/last/request-queue/head | Get last run&#39;s default request queue head
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueHeadLockPost**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueueheadlockpost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/head/lock | Get and lock last run&#39;s default request queue head
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueuePut**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueueput) | **Put** /v2/actors/{actorId}/runs/last/request-queue | Update last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestDelete**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestdelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId} | Delete request from last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestGet**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestget) | **Get** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId} | Get request from last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestLockDelete**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestlockdelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}/lock | Delete lock on request in last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestLockPut**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestlockput) | **Put** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}/lock | Prolong lock on request in last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestPut**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestput) | **Put** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId} | Update request in last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestsBatchDelete**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestsbatchdelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue/requests/batch | Batch delete requests from last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestsBatchPost**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestsbatchpost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/requests/batch | Batch add requests to last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestsGet**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestsget) | **Get** /v2/actors/{actorId}/runs/last/request-queue/requests | List last run&#39;s default request queue&#39;s requests
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestsPost**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestspost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/requests | Add request to last run&#39;s default request queue
*LastActorRunsDefaultRequestQueueAPI* | [**ActRunsLastRequestQueueRequestsUnlockPost**](docs/LastActorRunsDefaultRequestQueueAPI.md#actrunslastrequestqueuerequestsunlockpost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/requests/unlock | Unlock requests in last run&#39;s default request queue
*LastActorRunsLogAPI* | [**ActRunsLastLogGet**](docs/LastActorRunsLogAPI.md#actrunslastlogget) | **Get** /v2/actors/{actorId}/runs/last/log | Get last Actor run&#39;s log
*LastActorRunsMetamorphAPI* | [**ActRunsLastMetamorphPost**](docs/LastActorRunsMetamorphAPI.md#actrunslastmetamorphpost) | **Post** /v2/actors/{actorId}/runs/last/metamorph | Metamorph Actor&#39;s last run
*LastActorRunsRebootAPI* | [**ActRunsLastRebootPost**](docs/LastActorRunsRebootAPI.md#actrunslastrebootpost) | **Post** /v2/actors/{actorId}/runs/last/reboot | Reboot Actor&#39;s last run
*LastActorTaskRunsAbortAPI* | [**ActorTaskRunsLastAbortPost**](docs/LastActorTaskRunsAbortAPI.md#actortaskrunslastabortpost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/abort | Abort Actor task&#39;s last run
*LastActorTaskRunsDefaultDatasetAPI* | [**ActorTaskRunsLastDatasetDelete**](docs/LastActorTaskRunsDefaultDatasetAPI.md#actortaskrunslastdatasetdelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/dataset | Delete last task run&#39;s default dataset
*LastActorTaskRunsDefaultDatasetAPI* | [**ActorTaskRunsLastDatasetGet**](docs/LastActorTaskRunsDefaultDatasetAPI.md#actortaskrunslastdatasetget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/dataset | Get last task run&#39;s default dataset
*LastActorTaskRunsDefaultDatasetAPI* | [**ActorTaskRunsLastDatasetItemsGet**](docs/LastActorTaskRunsDefaultDatasetAPI.md#actortaskrunslastdatasetitemsget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/dataset/items | Get last task run&#39;s dataset items
*LastActorTaskRunsDefaultDatasetAPI* | [**ActorTaskRunsLastDatasetItemsPost**](docs/LastActorTaskRunsDefaultDatasetAPI.md#actortaskrunslastdatasetitemspost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/dataset/items | Store items in last task run&#39;s dataset
*LastActorTaskRunsDefaultDatasetAPI* | [**ActorTaskRunsLastDatasetPut**](docs/LastActorTaskRunsDefaultDatasetAPI.md#actortaskrunslastdatasetput) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/dataset | Update last task run&#39;s default dataset
*LastActorTaskRunsDefaultDatasetAPI* | [**ActorTaskRunsLastDatasetStatisticsGet**](docs/LastActorTaskRunsDefaultDatasetAPI.md#actortaskrunslastdatasetstatisticsget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/dataset/statistics | Get last task run&#39;s dataset statistics
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreDelete**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestoredelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store | Delete last task run&#39;s default store
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreGet**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestoreget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store | Get last task run&#39;s default store
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreKeysGet**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestorekeysget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/keys | Get last task run&#39;s default store&#39;s list of keys
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStorePut**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestoreput) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store | Update last task run&#39;s default store
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreRecordDelete**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestorerecorddelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Delete last task run&#39;s default store&#39;s record
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreRecordGet**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestorerecordget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Get last task run&#39;s default store&#39;s record
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreRecordPost**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestorerecordpost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Store record in last task run&#39;s default store (POST)
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreRecordPut**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestorerecordput) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Store record in last task run&#39;s default store
*LastActorTaskRunsDefaultKeyValueStoreAPI* | [**ActorTaskRunsLastKeyValueStoreRecordsGet**](docs/LastActorTaskRunsDefaultKeyValueStoreAPI.md#actortaskrunslastkeyvaluestorerecordsget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records | Download last task run&#39;s default store&#39;s records
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueDelete**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuedelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue | Delete last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueGet**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueueget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue | Get last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueHeadGet**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueueheadget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/head | Get last task run&#39;s default request queue head
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueHeadLockPost**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueueheadlockpost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/head/lock | Get and lock last task run&#39;s default request queue head
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueuePut**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueueput) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue | Update last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestDelete**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestdelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId} | Delete request from last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestGet**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId} | Get request from last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestLockDelete**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestlockdelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId}/lock | Delete lock on request in last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestLockPut**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestlockput) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId}/lock | Prolong lock on request in last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestPut**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestput) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId} | Update request in last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestsBatchDelete**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestsbatchdelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/batch | Batch delete requests from last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestsBatchPost**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestsbatchpost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/batch | Batch add requests to last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestsGet**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestsget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests | List last task run&#39;s default request queue&#39;s requests
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestsPost**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestspost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests | Add request to last task run&#39;s default request queue
*LastActorTaskRunsDefaultRequestQueueAPI* | [**ActorTaskRunsLastRequestQueueRequestsUnlockPost**](docs/LastActorTaskRunsDefaultRequestQueueAPI.md#actortaskrunslastrequestqueuerequestsunlockpost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/unlock | Unlock requests in last task run&#39;s default request queue
*LastActorTaskRunsLogAPI* | [**ActorTaskLastLogGet**](docs/LastActorTaskRunsLogAPI.md#actortasklastlogget) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/log | Get last Actor task run&#39;s log
*LastActorTaskRunsMetamorphAPI* | [**ActorTaskRunsLastMetamorphPost**](docs/LastActorTaskRunsMetamorphAPI.md#actortaskrunslastmetamorphpost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/metamorph | Metamorph Actor task&#39;s last run
*LastActorTaskRunsRebootAPI* | [**ActorTaskRunsLastRebootPost**](docs/LastActorTaskRunsRebootAPI.md#actortaskrunslastrebootpost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/reboot | Reboot Actor task&#39;s last run
*LogsAPI* | [**LogGet**](docs/LogsAPI.md#logget) | **Get** /v2/logs/{buildOrRunId} | Get log
*SchedulesAPI* | [**ScheduleDelete**](docs/SchedulesAPI.md#scheduledelete) | **Delete** /v2/schedules/{scheduleId} | Delete schedule
*SchedulesAPI* | [**ScheduleGet**](docs/SchedulesAPI.md#scheduleget) | **Get** /v2/schedules/{scheduleId} | Get schedule
*SchedulesAPI* | [**ScheduleLogGet**](docs/SchedulesAPI.md#schedulelogget) | **Get** /v2/schedules/{scheduleId}/log | Get schedule log
*SchedulesAPI* | [**SchedulePut**](docs/SchedulesAPI.md#scheduleput) | **Put** /v2/schedules/{scheduleId} | Update schedule
*SchedulesAPI* | [**SchedulesGet**](docs/SchedulesAPI.md#schedulesget) | **Get** /v2/schedules | Get list of schedules
*SchedulesAPI* | [**SchedulesPost**](docs/SchedulesAPI.md#schedulespost) | **Post** /v2/schedules | Create schedule
*StorageDatasetsAPI* | [**DatasetDelete**](docs/StorageDatasetsAPI.md#datasetdelete) | **Delete** /v2/datasets/{datasetId} | Delete dataset
*StorageDatasetsAPI* | [**DatasetGet**](docs/StorageDatasetsAPI.md#datasetget) | **Get** /v2/datasets/{datasetId} | Get dataset
*StorageDatasetsAPI* | [**DatasetItemsGet**](docs/StorageDatasetsAPI.md#datasetitemsget) | **Get** /v2/datasets/{datasetId}/items | Get dataset items
*StorageDatasetsAPI* | [**DatasetItemsHead**](docs/StorageDatasetsAPI.md#datasetitemshead) | **Head** /v2/datasets/{datasetId}/items | Get dataset items headers
*StorageDatasetsAPI* | [**DatasetItemsPost**](docs/StorageDatasetsAPI.md#datasetitemspost) | **Post** /v2/datasets/{datasetId}/items | Store items
*StorageDatasetsAPI* | [**DatasetPut**](docs/StorageDatasetsAPI.md#datasetput) | **Put** /v2/datasets/{datasetId} | Update dataset
*StorageDatasetsAPI* | [**DatasetStatisticsGet**](docs/StorageDatasetsAPI.md#datasetstatisticsget) | **Get** /v2/datasets/{datasetId}/statistics | Get dataset statistics
*StorageDatasetsAPI* | [**DatasetsGet**](docs/StorageDatasetsAPI.md#datasetsget) | **Get** /v2/datasets | Get list of datasets
*StorageDatasetsAPI* | [**DatasetsPost**](docs/StorageDatasetsAPI.md#datasetspost) | **Post** /v2/datasets | Create dataset
*StorageKeyValueStoresAPI* | [**KeyValueStoreDelete**](docs/StorageKeyValueStoresAPI.md#keyvaluestoredelete) | **Delete** /v2/key-value-stores/{storeId} | Delete store
*StorageKeyValueStoresAPI* | [**KeyValueStoreGet**](docs/StorageKeyValueStoresAPI.md#keyvaluestoreget) | **Get** /v2/key-value-stores/{storeId} | Get store
*StorageKeyValueStoresAPI* | [**KeyValueStoreKeysGet**](docs/StorageKeyValueStoresAPI.md#keyvaluestorekeysget) | **Get** /v2/key-value-stores/{storeId}/keys | Get list of keys
*StorageKeyValueStoresAPI* | [**KeyValueStorePut**](docs/StorageKeyValueStoresAPI.md#keyvaluestoreput) | **Put** /v2/key-value-stores/{storeId} | Update store
*StorageKeyValueStoresAPI* | [**KeyValueStoreRecordDelete**](docs/StorageKeyValueStoresAPI.md#keyvaluestorerecorddelete) | **Delete** /v2/key-value-stores/{storeId}/records/{recordKey} | Delete record
*StorageKeyValueStoresAPI* | [**KeyValueStoreRecordGet**](docs/StorageKeyValueStoresAPI.md#keyvaluestorerecordget) | **Get** /v2/key-value-stores/{storeId}/records/{recordKey} | Get record
*StorageKeyValueStoresAPI* | [**KeyValueStoreRecordHead**](docs/StorageKeyValueStoresAPI.md#keyvaluestorerecordhead) | **Head** /v2/key-value-stores/{storeId}/records/{recordKey} | Check if a record exists
*StorageKeyValueStoresAPI* | [**KeyValueStoreRecordPost**](docs/StorageKeyValueStoresAPI.md#keyvaluestorerecordpost) | **Post** /v2/key-value-stores/{storeId}/records/{recordKey} | Store record (POST)
*StorageKeyValueStoresAPI* | [**KeyValueStoreRecordPut**](docs/StorageKeyValueStoresAPI.md#keyvaluestorerecordput) | **Put** /v2/key-value-stores/{storeId}/records/{recordKey} | Store record
*StorageKeyValueStoresAPI* | [**KeyValueStoreRecordsGet**](docs/StorageKeyValueStoresAPI.md#keyvaluestorerecordsget) | **Get** /v2/key-value-stores/{storeId}/records | Download records
*StorageKeyValueStoresAPI* | [**KeyValueStoresGet**](docs/StorageKeyValueStoresAPI.md#keyvaluestoresget) | **Get** /v2/key-value-stores | Get list of key-value stores
*StorageKeyValueStoresAPI* | [**KeyValueStoresPost**](docs/StorageKeyValueStoresAPI.md#keyvaluestorespost) | **Post** /v2/key-value-stores | Create key-value store
*StorageRequestQueuesAPI* | [**RequestQueueDelete**](docs/StorageRequestQueuesAPI.md#requestqueuedelete) | **Delete** /v2/request-queues/{queueId} | Delete request queue
*StorageRequestQueuesAPI* | [**RequestQueueGet**](docs/StorageRequestQueuesAPI.md#requestqueueget) | **Get** /v2/request-queues/{queueId} | Get request queue
*StorageRequestQueuesAPI* | [**RequestQueuePut**](docs/StorageRequestQueuesAPI.md#requestqueueput) | **Put** /v2/request-queues/{queueId} | Update request queue
*StorageRequestQueuesAPI* | [**RequestQueueRequestsBatchDelete**](docs/StorageRequestQueuesAPI.md#requestqueuerequestsbatchdelete) | **Delete** /v2/request-queues/{queueId}/requests/batch | Delete requests
*StorageRequestQueuesAPI* | [**RequestQueueRequestsBatchPost**](docs/StorageRequestQueuesAPI.md#requestqueuerequestsbatchpost) | **Post** /v2/request-queues/{queueId}/requests/batch | Add requests
*StorageRequestQueuesAPI* | [**RequestQueuesGet**](docs/StorageRequestQueuesAPI.md#requestqueuesget) | **Get** /v2/request-queues | Get list of request queues
*StorageRequestQueuesAPI* | [**RequestQueuesPost**](docs/StorageRequestQueuesAPI.md#requestqueuespost) | **Post** /v2/request-queues | Create request queue
*StorageRequestQueuesRequestsAPI* | [**RequestQueueRequestDelete**](docs/StorageRequestQueuesRequestsAPI.md#requestqueuerequestdelete) | **Delete** /v2/request-queues/{queueId}/requests/{requestId} | Delete request
*StorageRequestQueuesRequestsAPI* | [**RequestQueueRequestGet**](docs/StorageRequestQueuesRequestsAPI.md#requestqueuerequestget) | **Get** /v2/request-queues/{queueId}/requests/{requestId} | Get request
*StorageRequestQueuesRequestsAPI* | [**RequestQueueRequestPut**](docs/StorageRequestQueuesRequestsAPI.md#requestqueuerequestput) | **Put** /v2/request-queues/{queueId}/requests/{requestId} | Update request
*StorageRequestQueuesRequestsAPI* | [**RequestQueueRequestsGet**](docs/StorageRequestQueuesRequestsAPI.md#requestqueuerequestsget) | **Get** /v2/request-queues/{queueId}/requests | List requests
*StorageRequestQueuesRequestsAPI* | [**RequestQueueRequestsPost**](docs/StorageRequestQueuesRequestsAPI.md#requestqueuerequestspost) | **Post** /v2/request-queues/{queueId}/requests | Add request
*StorageRequestQueuesRequestsLocksAPI* | [**RequestQueueHeadGet**](docs/StorageRequestQueuesRequestsLocksAPI.md#requestqueueheadget) | **Get** /v2/request-queues/{queueId}/head | Get head
*StorageRequestQueuesRequestsLocksAPI* | [**RequestQueueHeadLockPost**](docs/StorageRequestQueuesRequestsLocksAPI.md#requestqueueheadlockpost) | **Post** /v2/request-queues/{queueId}/head/lock | Get head and lock
*StorageRequestQueuesRequestsLocksAPI* | [**RequestQueueRequestLockDelete**](docs/StorageRequestQueuesRequestsLocksAPI.md#requestqueuerequestlockdelete) | **Delete** /v2/request-queues/{queueId}/requests/{requestId}/lock | Delete request lock
*StorageRequestQueuesRequestsLocksAPI* | [**RequestQueueRequestLockPut**](docs/StorageRequestQueuesRequestsLocksAPI.md#requestqueuerequestlockput) | **Put** /v2/request-queues/{queueId}/requests/{requestId}/lock | Prolong request lock
*StorageRequestQueuesRequestsLocksAPI* | [**RequestQueueRequestsUnlockPost**](docs/StorageRequestQueuesRequestsLocksAPI.md#requestqueuerequestsunlockpost) | **Post** /v2/request-queues/{queueId}/requests/unlock | Unlock requests
*StoreAPI* | [**StoreGet**](docs/StoreAPI.md#storeget) | **Get** /v2/store | Get list of Actors in Store
*ToolsAPI* | [**ToolsBrowserInfoDelete**](docs/ToolsAPI.md#toolsbrowserinfodelete) | **Delete** /v2/browser-info | Get browser info
*ToolsAPI* | [**ToolsBrowserInfoGet**](docs/ToolsAPI.md#toolsbrowserinfoget) | **Get** /v2/browser-info | Get browser info
*ToolsAPI* | [**ToolsBrowserInfoPost**](docs/ToolsAPI.md#toolsbrowserinfopost) | **Post** /v2/browser-info | Get browser info
*ToolsAPI* | [**ToolsBrowserInfoPut**](docs/ToolsAPI.md#toolsbrowserinfoput) | **Put** /v2/browser-info | Get browser info
*ToolsAPI* | [**ToolsDecodeAndVerifyPost**](docs/ToolsAPI.md#toolsdecodeandverifypost) | **Post** /v2/tools/decode-and-verify | Decode and verify object
*ToolsAPI* | [**ToolsEncodeAndSignPost**](docs/ToolsAPI.md#toolsencodeandsignpost) | **Post** /v2/tools/encode-and-sign | Encode and sign object
*UsersAPI* | [**UserGet**](docs/UsersAPI.md#userget) | **Get** /v2/users/{userId} | Get public user data
*UsersAPI* | [**UsersMeGet**](docs/UsersAPI.md#usersmeget) | **Get** /v2/users/me | Get private user data
*UsersAPI* | [**UsersMeLimitsGet**](docs/UsersAPI.md#usersmelimitsget) | **Get** /v2/users/me/limits | Get limits
*UsersAPI* | [**UsersMeLimitsPut**](docs/UsersAPI.md#usersmelimitsput) | **Put** /v2/users/me/limits | Update limits
*UsersAPI* | [**UsersMeUsageMonthlyGet**](docs/UsersAPI.md#usersmeusagemonthlyget) | **Get** /v2/users/me/usage/monthly | Get monthly usage
*WebhooksWebhookDispatchesAPI* | [**WebhookDispatchGet**](docs/WebhooksWebhookDispatchesAPI.md#webhookdispatchget) | **Get** /v2/webhook-dispatches/{dispatchId} | Get webhook dispatch
*WebhooksWebhookDispatchesAPI* | [**WebhookDispatchesGet**](docs/WebhooksWebhookDispatchesAPI.md#webhookdispatchesget) | **Get** /v2/webhook-dispatches | Get list of webhook dispatches
*WebhooksWebhooksAPI* | [**WebhookDelete**](docs/WebhooksWebhooksAPI.md#webhookdelete) | **Delete** /v2/webhooks/{webhookId} | Delete webhook
*WebhooksWebhooksAPI* | [**WebhookGet**](docs/WebhooksWebhooksAPI.md#webhookget) | **Get** /v2/webhooks/{webhookId} | Get webhook
*WebhooksWebhooksAPI* | [**WebhookPut**](docs/WebhooksWebhooksAPI.md#webhookput) | **Put** /v2/webhooks/{webhookId} | Update webhook
*WebhooksWebhooksAPI* | [**WebhookTestPost**](docs/WebhooksWebhooksAPI.md#webhooktestpost) | **Post** /v2/webhooks/{webhookId}/test | Test webhook
*WebhooksWebhooksAPI* | [**WebhookWebhookDispatchesGet**](docs/WebhooksWebhooksAPI.md#webhookwebhookdispatchesget) | **Get** /v2/webhooks/{webhookId}/dispatches | Get collection
*WebhooksWebhooksAPI* | [**WebhooksGet**](docs/WebhooksWebhooksAPI.md#webhooksget) | **Get** /v2/webhooks | Get list of webhooks
*WebhooksWebhooksAPI* | [**WebhooksPost**](docs/WebhooksWebhooksAPI.md#webhookspost) | **Post** /v2/webhooks | Create webhook


## Documentation For Models

 - [AccountLimits](docs/AccountLimits.md)
 - [ActRunsLastDatasetItemsPost400Response](docs/ActRunsLastDatasetItemsPost400Response.md)
 - [ActRunsLastDatasetItemsPostRequest](docs/ActRunsLastDatasetItemsPostRequest.md)
 - [ActValidateInputPost200Response](docs/ActValidateInputPost200Response.md)
 - [Actor](docs/Actor.md)
 - [ActorChargeEvent](docs/ActorChargeEvent.md)
 - [ActorDefinition](docs/ActorDefinition.md)
 - [ActorDefinitionDefaultMemoryMbytes](docs/ActorDefinitionDefaultMemoryMbytes.md)
 - [ActorDefinitionStorages](docs/ActorDefinitionStorages.md)
 - [ActorJobStatus](docs/ActorJobStatus.md)
 - [ActorPermissionLevel](docs/ActorPermissionLevel.md)
 - [ActorResponse](docs/ActorResponse.md)
 - [ActorRunFailedError](docs/ActorRunFailedError.md)
 - [ActorRunPricingInfo](docs/ActorRunPricingInfo.md)
 - [ActorRunTimeoutExceededError](docs/ActorRunTimeoutExceededError.md)
 - [ActorShort](docs/ActorShort.md)
 - [ActorStandby](docs/ActorStandby.md)
 - [ActorStats](docs/ActorStats.md)
 - [ActorStatsPublicActorRunStats30Days](docs/ActorStatsPublicActorRunStats30Days.md)
 - [ActorTaskGet200Response](docs/ActorTaskGet200Response.md)
 - [ActorTaskRunsGet200Response](docs/ActorTaskRunsGet200Response.md)
 - [ActorTaskRunsGet200ResponseData](docs/ActorTaskRunsGet200ResponseData.md)
 - [ActorTaskRunsPost201Response](docs/ActorTaskRunsPost201Response.md)
 - [ActorTaskWebhooksGet200Response](docs/ActorTaskWebhooksGet200Response.md)
 - [ActorTaskWebhooksGet200ResponseData](docs/ActorTaskWebhooksGet200ResponseData.md)
 - [AddRequestResponse](docs/AddRequestResponse.md)
 - [AddedRequest](docs/AddedRequest.md)
 - [BatchAddResponse](docs/BatchAddResponse.md)
 - [BatchAddResult](docs/BatchAddResult.md)
 - [BatchDeleteResponse](docs/BatchDeleteResponse.md)
 - [BatchDeleteResult](docs/BatchDeleteResult.md)
 - [BrowserInfoResponse](docs/BrowserInfoResponse.md)
 - [BrowserInfoResponseHeadersValue](docs/BrowserInfoResponseHeadersValue.md)
 - [Build](docs/Build.md)
 - [BuildActVersion](docs/BuildActVersion.md)
 - [BuildOptions](docs/BuildOptions.md)
 - [BuildResponse](docs/BuildResponse.md)
 - [BuildShort](docs/BuildShort.md)
 - [BuildStats](docs/BuildStats.md)
 - [BuildTag](docs/BuildTag.md)
 - [BuildUsage](docs/BuildUsage.md)
 - [BuildsMeta](docs/BuildsMeta.md)
 - [CallsInner](docs/CallsInner.md)
 - [ChargeRunRequest](docs/ChargeRunRequest.md)
 - [CommonActorPricingInfo](docs/CommonActorPricingInfo.md)
 - [CreateActorRequest](docs/CreateActorRequest.md)
 - [CreateOrUpdateVersionRequest](docs/CreateOrUpdateVersionRequest.md)
 - [CreateTaskRequest](docs/CreateTaskRequest.md)
 - [Current](docs/Current.md)
 - [CurrentPricingInfo](docs/CurrentPricingInfo.md)
 - [DailyServiceUsages](docs/DailyServiceUsages.md)
 - [Dataset](docs/Dataset.md)
 - [DatasetFieldStatistics](docs/DatasetFieldStatistics.md)
 - [DatasetListItem](docs/DatasetListItem.md)
 - [DatasetResponse](docs/DatasetResponse.md)
 - [DatasetSchemaValidationError](docs/DatasetSchemaValidationError.md)
 - [DatasetStatistics](docs/DatasetStatistics.md)
 - [DatasetStatisticsResponse](docs/DatasetStatisticsResponse.md)
 - [DatasetStats](docs/DatasetStats.md)
 - [DecodeAndVerifyData](docs/DecodeAndVerifyData.md)
 - [DecodeAndVerifyRequest](docs/DecodeAndVerifyRequest.md)
 - [DecodeAndVerifyResponse](docs/DecodeAndVerifyResponse.md)
 - [DefaultRunOptions](docs/DefaultRunOptions.md)
 - [DeletedRequest](docs/DeletedRequest.md)
 - [DeletedRequestById](docs/DeletedRequestById.md)
 - [DeletedRequestByUniqueKey](docs/DeletedRequestByUniqueKey.md)
 - [EffectivePlatformFeature](docs/EffectivePlatformFeature.md)
 - [EffectivePlatformFeatures](docs/EffectivePlatformFeatures.md)
 - [EncodeAndSignData](docs/EncodeAndSignData.md)
 - [EncodeAndSignResponse](docs/EncodeAndSignResponse.md)
 - [EnvVar](docs/EnvVar.md)
 - [EnvVarRequest](docs/EnvVarRequest.md)
 - [EnvVarResponse](docs/EnvVarResponse.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorType](docs/ErrorType.md)
 - [EventData](docs/EventData.md)
 - [ExampleRunInput](docs/ExampleRunInput.md)
 - [ExampleWebhookDispatch](docs/ExampleWebhookDispatch.md)
 - [FlatPricePerMonthActorPricingInfo](docs/FlatPricePerMonthActorPricingInfo.md)
 - [FreeActorPricingInfo](docs/FreeActorPricingInfo.md)
 - [GeneralAccess](docs/GeneralAccess.md)
 - [HeadAndLockResponse](docs/HeadAndLockResponse.md)
 - [HeadRequest](docs/HeadRequest.md)
 - [HeadResponse](docs/HeadResponse.md)
 - [HttpMethod](docs/HttpMethod.md)
 - [InvalidItem](docs/InvalidItem.md)
 - [KeyValueStore](docs/KeyValueStore.md)
 - [KeyValueStoreKey](docs/KeyValueStoreKey.md)
 - [KeyValueStoreResponse](docs/KeyValueStoreResponse.md)
 - [KeyValueStoreStats](docs/KeyValueStoreStats.md)
 - [Limits](docs/Limits.md)
 - [LimitsResponse](docs/LimitsResponse.md)
 - [ListOfActors](docs/ListOfActors.md)
 - [ListOfActorsInStoreResponse](docs/ListOfActorsInStoreResponse.md)
 - [ListOfActorsResponse](docs/ListOfActorsResponse.md)
 - [ListOfBuilds](docs/ListOfBuilds.md)
 - [ListOfBuildsResponse](docs/ListOfBuildsResponse.md)
 - [ListOfDatasets](docs/ListOfDatasets.md)
 - [ListOfDatasetsResponse](docs/ListOfDatasetsResponse.md)
 - [ListOfEnvVars](docs/ListOfEnvVars.md)
 - [ListOfEnvVarsResponse](docs/ListOfEnvVarsResponse.md)
 - [ListOfKeyValueStores](docs/ListOfKeyValueStores.md)
 - [ListOfKeyValueStoresResponse](docs/ListOfKeyValueStoresResponse.md)
 - [ListOfKeys](docs/ListOfKeys.md)
 - [ListOfKeysResponse](docs/ListOfKeysResponse.md)
 - [ListOfRequestQueues](docs/ListOfRequestQueues.md)
 - [ListOfRequestQueuesResponse](docs/ListOfRequestQueuesResponse.md)
 - [ListOfRequests](docs/ListOfRequests.md)
 - [ListOfRequestsResponse](docs/ListOfRequestsResponse.md)
 - [ListOfRuns](docs/ListOfRuns.md)
 - [ListOfRunsResponse](docs/ListOfRunsResponse.md)
 - [ListOfSchedules](docs/ListOfSchedules.md)
 - [ListOfSchedulesResponse](docs/ListOfSchedulesResponse.md)
 - [ListOfStoreActors](docs/ListOfStoreActors.md)
 - [ListOfTasks](docs/ListOfTasks.md)
 - [ListOfTasksResponse](docs/ListOfTasksResponse.md)
 - [ListOfVersions](docs/ListOfVersions.md)
 - [ListOfVersionsResponse](docs/ListOfVersionsResponse.md)
 - [ListOfWebhookDispatches](docs/ListOfWebhookDispatches.md)
 - [ListOfWebhookDispatchesResponse](docs/ListOfWebhookDispatchesResponse.md)
 - [ListOfWebhooks](docs/ListOfWebhooks.md)
 - [ListOfWebhooksResponse](docs/ListOfWebhooksResponse.md)
 - [LockedHeadRequest](docs/LockedHeadRequest.md)
 - [LockedRequestQueueHead](docs/LockedRequestQueueHead.md)
 - [Metamorph](docs/Metamorph.md)
 - [MonthlyUsage](docs/MonthlyUsage.md)
 - [MonthlyUsageResponse](docs/MonthlyUsageResponse.md)
 - [PaginationResponse](docs/PaginationResponse.md)
 - [PayPerEventActorPricingInfo](docs/PayPerEventActorPricingInfo.md)
 - [PayPerEventActorPricingInfoAllOfPricingPerEvent](docs/PayPerEventActorPricingInfoAllOfPricingPerEvent.md)
 - [Plan](docs/Plan.md)
 - [PricePerDatasetItemActorPricingInfo](docs/PricePerDatasetItemActorPricingInfo.md)
 - [PriceTiers](docs/PriceTiers.md)
 - [PrivateUserDataResponse](docs/PrivateUserDataResponse.md)
 - [Profile](docs/Profile.md)
 - [ProlongRequestLockResponse](docs/ProlongRequestLockResponse.md)
 - [Proxy](docs/Proxy.md)
 - [ProxyGroup](docs/ProxyGroup.md)
 - [PublicUserDataResponse](docs/PublicUserDataResponse.md)
 - [PutItemResponseError](docs/PutItemResponseError.md)
 - [Request](docs/Request.md)
 - [RequestBase](docs/RequestBase.md)
 - [RequestBasePayload](docs/RequestBasePayload.md)
 - [RequestDraft](docs/RequestDraft.md)
 - [RequestDraftDelete](docs/RequestDraftDelete.md)
 - [RequestDraftDeleteById](docs/RequestDraftDeleteById.md)
 - [RequestDraftDeleteByUniqueKey](docs/RequestDraftDeleteByUniqueKey.md)
 - [RequestLockInfo](docs/RequestLockInfo.md)
 - [RequestQueue](docs/RequestQueue.md)
 - [RequestQueueHead](docs/RequestQueueHead.md)
 - [RequestQueueResponse](docs/RequestQueueResponse.md)
 - [RequestQueueShort](docs/RequestQueueShort.md)
 - [RequestQueueStats](docs/RequestQueueStats.md)
 - [RequestRegistration](docs/RequestRegistration.md)
 - [RequestResponse](docs/RequestResponse.md)
 - [Run](docs/Run.md)
 - [RunFailedErrorDetail](docs/RunFailedErrorDetail.md)
 - [RunMeta](docs/RunMeta.md)
 - [RunOptions](docs/RunOptions.md)
 - [RunOrigin](docs/RunOrigin.md)
 - [RunResponse](docs/RunResponse.md)
 - [RunShort](docs/RunShort.md)
 - [RunStats](docs/RunStats.md)
 - [RunStorageIds](docs/RunStorageIds.md)
 - [RunStorageIdsDatasets](docs/RunStorageIdsDatasets.md)
 - [RunTimeoutExceededErrorDetail](docs/RunTimeoutExceededErrorDetail.md)
 - [RunUsage](docs/RunUsage.md)
 - [RunUsageUsd](docs/RunUsageUsd.md)
 - [Schedule](docs/Schedule.md)
 - [ScheduleAction](docs/ScheduleAction.md)
 - [ScheduleActionRunActor](docs/ScheduleActionRunActor.md)
 - [ScheduleActionRunActorTask](docs/ScheduleActionRunActorTask.md)
 - [ScheduleActionRunInput](docs/ScheduleActionRunInput.md)
 - [ScheduleActionShort](docs/ScheduleActionShort.md)
 - [ScheduleActionShortRunActor](docs/ScheduleActionShortRunActor.md)
 - [ScheduleActionShortRunActorTask](docs/ScheduleActionShortRunActorTask.md)
 - [ScheduleBase](docs/ScheduleBase.md)
 - [ScheduleCreate](docs/ScheduleCreate.md)
 - [ScheduleCreateAction](docs/ScheduleCreateAction.md)
 - [ScheduleCreateActionRunActor](docs/ScheduleCreateActionRunActor.md)
 - [ScheduleCreateActionRunActorTask](docs/ScheduleCreateActionRunActorTask.md)
 - [ScheduleInvoked](docs/ScheduleInvoked.md)
 - [ScheduleLogResponse](docs/ScheduleLogResponse.md)
 - [ScheduleNotifications](docs/ScheduleNotifications.md)
 - [ScheduleResponse](docs/ScheduleResponse.md)
 - [ScheduleShort](docs/ScheduleShort.md)
 - [SchemaValidationErrorData](docs/SchemaValidationErrorData.md)
 - [SourceCodeFile](docs/SourceCodeFile.md)
 - [SourceCodeFileFormat](docs/SourceCodeFileFormat.md)
 - [SourceCodeFolder](docs/SourceCodeFolder.md)
 - [StorageOwnership](docs/StorageOwnership.md)
 - [StoreListActor](docs/StoreListActor.md)
 - [TaggedBuildInfo](docs/TaggedBuildInfo.md)
 - [TaggedBuildsValue](docs/TaggedBuildsValue.md)
 - [Task](docs/Task.md)
 - [TaskOptions](docs/TaskOptions.md)
 - [TaskResponse](docs/TaskResponse.md)
 - [TaskShort](docs/TaskShort.md)
 - [TaskStats](docs/TaskStats.md)
 - [TestWebhookResponse](docs/TestWebhookResponse.md)
 - [TieredPricingPerDatasetItemEntry](docs/TieredPricingPerDatasetItemEntry.md)
 - [TieredPricingPerEventEntry](docs/TieredPricingPerEventEntry.md)
 - [UnlockRequestsResponse](docs/UnlockRequestsResponse.md)
 - [UnlockRequestsResult](docs/UnlockRequestsResult.md)
 - [UpdateActorRequest](docs/UpdateActorRequest.md)
 - [UpdateDatasetRequest](docs/UpdateDatasetRequest.md)
 - [UpdateLimitsRequest](docs/UpdateLimitsRequest.md)
 - [UpdateRequestQueueRequest](docs/UpdateRequestQueueRequest.md)
 - [UpdateRequestResponse](docs/UpdateRequestResponse.md)
 - [UpdateRunRequest](docs/UpdateRunRequest.md)
 - [UpdateStoreRequest](docs/UpdateStoreRequest.md)
 - [UpdateTaskRequest](docs/UpdateTaskRequest.md)
 - [UsageCycle](docs/UsageCycle.md)
 - [UsageItem](docs/UsageItem.md)
 - [UserPrivateInfo](docs/UserPrivateInfo.md)
 - [UserPublicInfo](docs/UserPublicInfo.md)
 - [ValidationError](docs/ValidationError.md)
 - [Version](docs/Version.md)
 - [VersionResponse](docs/VersionResponse.md)
 - [VersionSourceFilesInner](docs/VersionSourceFilesInner.md)
 - [VersionSourceType](docs/VersionSourceType.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookCondition](docs/WebhookCondition.md)
 - [WebhookCreate](docs/WebhookCreate.md)
 - [WebhookDispatch](docs/WebhookDispatch.md)
 - [WebhookDispatchResponse](docs/WebhookDispatchResponse.md)
 - [WebhookDispatchStatus](docs/WebhookDispatchStatus.md)
 - [WebhookDispatchWebhookSummary](docs/WebhookDispatchWebhookSummary.md)
 - [WebhookEventType](docs/WebhookEventType.md)
 - [WebhookRepresentation](docs/WebhookRepresentation.md)
 - [WebhookResponse](docs/WebhookResponse.md)
 - [WebhookShort](docs/WebhookShort.md)
 - [WebhookStats](docs/WebhookStats.md)
 - [WebhookUpdate](docs/WebhookUpdate.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### httpBearer

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), apifyclient.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### apiKey

- **Type**: API key
- **API key parameter name**: token
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apiKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		apifyclient.ContextAPIKeys,
		map[string]apifyclient.APIKey{
			"apiKey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



