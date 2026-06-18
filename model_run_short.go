/*
Apify API

 The Apify API (version 2) provides programmatic access to the [Apify platform](https://docs.apify.com). The API is organized around [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) HTTP endpoints.  You can download the complete OpenAPI schema of Apify API in the [YAML](http://docs.apify.com/api/openapi.yaml) or [JSON](http://docs.apify.com/api/openapi.json) formats. The source code is also available on [GitHub](https://github.com/apify/apify-docs/tree/master/apify-api/openapi).  All requests and responses (including errors) are encoded in [JSON](http://www.json.org/) format with UTF-8 encoding, with a few exceptions that are explicitly described in the reference.  - To access the API using [Node.js](https://nodejs.org/en/), we recommend the [`apify-client`](https://docs.apify.com/api/client/js) [NPM package](https://www.npmjs.com/package/apify-client). - To access the API using [Python](https://www.python.org/), we recommend the [`apify-client`](https://docs.apify.com/api/client/python) [PyPI package](https://pypi.org/project/apify-client/).  The clients' functions correspond to the API endpoints and have the same parameters. This simplifies development of apps that depend on the Apify platform.  :::note Important Request Details  - `Content-Type` header: For requests with a JSON body, you must include the `Content-Type: application/json` header.  - Method override: You can override the HTTP method using the `method` query parameter. This is useful for clients that can only send `GET` requests. For example, to call a `POST` endpoint, append `?method=POST` to the URL of your `GET` request.  :::  ## Authentication <span id=\"/introduction/authentication\"></span>  **You can find your API token on the [Integrations](https://console.apify.com/settings/integrations) page in the Apify Console.**  To use your token in a request, either:  - Add the token to your request's `Authorization` header as `Bearer <token>`. E.g., `Authorization: Bearer xxxxxxx`. [More info](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization). (Recommended). - Add it as the `token` parameter to your request URL. (Less secure).  Using your token in the request header is more secure than using it as a URL parameter because URLs are often stored in browser history and server logs. This creates a chance for someone unauthorized to access your API token.  **Never share your API token or password with untrusted parties!**  For more information, see our [integrations](https://docs.apify.com/platform/integrations) documentation.  ### Agentic payments  AI agents can authenticate and pay for Actor runs without an Apify account using agentic payments. Instead of an API token, the request carries a payment credential that both authorizes and pays for the call. Apify supports the [x402 protocol](https://docs.apify.com/platform/integrations/x402) (`PAYMENT-SIGNATURE` header) and [Skyfire](https://docs.apify.com/platform/integrations/skyfire) (`skyfire-pay-id` header).  ## Basic usage <span id=\"/introduction/basic-usage\"></span>  To run an Actor, send a POST request to the [Run Actor](#/reference/actors/run-collection/run-actor) endpoint using either the Actor ID code (e.g. `vKg4IjxZbEYTYeW8T`) or its name (e.g. `janedoe~my-actor`):  `https://api.apify.com/v2/actors/[actor_id]/runs`  If the Actor is not runnable anonymously, you will receive a 401 or 403 [response code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status). This means you need to add your [secret API token](https://console.apify.com/account#/integrations) to the request's `Authorization` header ([recommended](#/introduction/authentication)) or as a URL query parameter `?token=[your_token]` (less secure).  Optionally, you can include the query parameters described in the [Run Actor](#/reference/actors/run-collection/run-actor) section to customize your run.  If you're using Node.js, the best way to run an Actor is using the `Apify.call()` method from the [Apify SDK](https://sdk.apify.com/docs/api/apify#apifycallactid-input-options). It runs the Actor using the account you are currently logged into (determined by the [secret API token](https://console.apify.com/account#/integrations)). The result is an [Actor run object](https://sdk.apify.com/docs/typedefs/actor-run) and its output (if any).  A typical workflow is as follows:  1. Run an Actor or task using the [Run Actor](#/reference/actors/run-collection/run-actor) or [Run task](#/reference/actor-tasks/run-collection/run-task) API endpoints. 2. Monitor the Actor run by periodically polling its progress using the [Get run](#/reference/actor-runs/run-object-and-its-storages/get-run) API endpoint. 3. Fetch the results from the [Get items](#/reference/datasets/item-collection/get-items) API endpoint using the `defaultDatasetId`, which you receive in the Run request response. Additional data may be stored in a key-value store. You can fetch them from the [Get record](#/reference/key-value-stores/record/get-record) API endpoint using the `defaultKeyValueStoreId` and the store's `key`.  **Note**: Instead of periodic polling, you can also run your [Actor](#/reference/actors/run-actor-synchronously) or [task](#/reference/actor-tasks/runs-collection/run-task-synchronously) synchronously. This will ensure that the request waits for 300 seconds (5 minutes) for the run to finish and returns its output. If the run takes longer, the request will time out and throw an error.  ## Legacy `/v2/acts/` URL prefix <span id=\"/introduction/legacy-acts-prefix\"></span>  The `/v2/acts/` prefix is deprecated but still fully functional, and  such endpoint routes to the same handler as its `/v2/actors/...` counterpart.  New integrations should use the canonical /v2/actors/ prefix,  but existing clients keep working without changes.  ## Response structure <span id=\"/introduction/response-structure\"></span>  Most API endpoints return a JSON object with the `data` property:  ``` {     \"data\": {         ...     } } ```  However, there are a few explicitly described exceptions, such as [Get dataset items](#/reference/datasets/item-collection/get-items) or Key-value store [Get record](#/reference/key-value-stores/record/get-record) API endpoints, which return data in other formats. In case of an error, the response has the HTTP status code in the range of 4xx or 5xx and the `data` property is replaced with `error`. For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  See [Errors](#/introduction/errors) for more details.  ## Pagination <span id=\"/introduction/pagination\"></span>  All API endpoints that return a list of records (e.g. [Get list of Actors](#/reference/actors/actor-collection/get-list-of-actors)) enforce pagination in order to limit the size of their responses.  Most of these API endpoints are paginated using the `offset` and `limit` query parameters. The only exception is [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys), which is paginated using the `exclusiveStartKey` query parameter.  **IMPORTANT**: Each API endpoint that supports pagination enforces a certain maximum value for the `limit` parameter, in order to reduce the load on Apify servers. The maximum limit could change in future so you should never rely on a specific value and check the responses of these API endpoints.  ### Using offset <span id=\"/introduction/pagination/using-offset\"></span>  Most API endpoints that return a list of records enable pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number of items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>Skips a number of items from the beginning of the list, e.g. <code>offset=100</code>.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td>     By default, items are sorted in the order in which they were created or added to the list.     This feature is useful when fetching all the items, because it ensures that items     created after the client started the pagination will not be skipped.     If you specify the <code>desc=1</code> parameter, the items will be returned in the reverse order,     i.e. from the newest to the oldest items.     </td>   </tr> </table>  The response of these API endpoints is always a JSON object with the following structure:  ``` {     \"data\": {         \"total\": 2560,         \"offset\": 250,         \"limit\": 1000,         \"count\": 1000,         \"desc\": false,         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>total</code></td>     <td>The total number of items available in the list.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>The number of items that were skipped at the start.     This is equal to the <code>offset</code> query parameter if it was provided, otherwise it is <code>0</code>.</td>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular API endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>count</code></td>     <td>The actual number of items returned in the HTTP response.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td><code>true</code> if data were requested in descending order and <code>false</code> otherwise.</td>   </tr>   <tr>     <td><code>items</code></td>     <td>An array of requested items.</td>   </tr> </table>  ### Using key <span id=\"/introduction/pagination/using-key\"></span>  The records in the [key-value store](https://docs.apify.com/platform/storage/key-value-store) are not ordered based on numerical indexes, but rather by their keys in the UTF-8 binary order. Therefore the [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys) API endpoint only supports pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>Skips all records with keys up to the given key including the given key,     in the UTF-8 binary order.</td>   </tr> </table>  The response of the API endpoint is always a JSON object with following structure:  ``` {     \"data\": {         \"limit\": 1000,         \"isTruncated\": true,         \"exclusiveStartKey\": \"my-key\",         \"nextExclusiveStartKey\": \"some-other-key\",         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>isTruncated</code></td>     <td><code>true</code> if there are more items left to be queried. Otherwise <code>false</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>The last key that was skipped at the start. Is `null` for the first page.</td>   </tr>   <tr>     <td><code>nextExclusiveStartKey</code></td>     <td>The value for the <code>exclusiveStartKey</code> parameter to query the next page of items.</td>   </tr> </table>  ## Errors <span id=\"/introduction/errors\"></span>  The Apify API uses common HTTP status codes: `2xx` range for success, `4xx` range for errors caused by the caller (invalid requests) and `5xx` range for server errors (these are rare). Each error response contains a JSON object defining the `error` property, which is an object with the `type` and `message` properties that contain the error code and a human-readable error description, respectively.  For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  Here is the table of the most common errors that can occur for many API endpoints:  <table>   <tr>     <th>status</th>     <th>type</th>     <th>message</th>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-request</code></td>     <td>POST data must be a JSON object</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-value</code></td>     <td>Invalid value provided: Comments required</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-record-key</code></td>     <td>Record key contains invalid character</td>   </tr>   <tr>     <td><code>401</code></td>     <td><code>token-not-provided</code></td>     <td>Authentication token was not provided</td>   </tr>   <tr>     <td><code>404</code></td>     <td><code>record-not-found</code></td>     <td>Store was not found</td>   </tr>   <tr>     <td><code>429</code></td>     <td><code>rate-limit-exceeded</code></td>     <td>You have exceeded the rate limit of ... requests per second</td>   </tr>   <tr>     <td><code>405</code></td>     <td><code>method-not-allowed</code></td>     <td>This API endpoint can only be accessed using the following HTTP methods: OPTIONS, POST</td>   </tr> </table>  ## Rate limiting <span id=\"/introduction/rate-limiting\"></span>  All API endpoints limit the rate of requests in order to prevent overloading of Apify servers by misbehaving clients.  There are two kinds of rate limits - a global rate limit and a per-resource rate limit.  ### Global rate limit <span id=\"/introduction/rate-limiting/global-rate-limit\"></span>  The global rate limit is set to _250 000 requests per minute_. For [authenticated](#/introduction/authentication) requests, it is counted per user, and for unauthenticated requests, it is counted per IP address.  ### Per-resource rate limit <span id=\"/introduction/rate-limiting/per-resource-rate-limit\"></span>  The default per-resource rate limit is _60 requests per second per resource_, which in this context means a single Actor, a single Actor run, a single dataset, single key-value store etc. The default rate limit is applied to every API endpoint except a few select ones, which have higher rate limits. Each API endpoint returns its rate limit in `X-RateLimit-Limit` header.  These endpoints have a rate limit of _200 requests per second per resource_:  * CRUD ([get](#/reference/key-value-stores/record/get-record),   [put](#/reference/key-value-stores/record/put-record),   [delete](#/reference/key-value-stores/record/delete-record))   operations on key-value store records  These endpoints have a rate limit of _400 requests per second per resource_: * [Run Actor](#/reference/actors/run-collection/run-actor) * [Run Actor task asynchronously](#/reference/actor-tasks/runs-collection/run-task-asynchronously) * [Run Actor task synchronously](#/reference/actor-tasks/runs-collection/run-task-synchronously) * [Metamorph Actor run](#/reference/actors/metamorph-run/metamorph-run) * [Push items](#/reference/datasets/item-collection/put-items) to dataset * CRUD   ([add](#/reference/request-queues/request-collection/add-request),   [get](#/reference/request-queues/request-collection/get-request),   [update](#/reference/request-queues/request-collection/update-request),   [delete](#/reference/request-queues/request-collection/delete-request))   operations on requests in request queues  ### Rate limit exceeded errors <span id=\"/introduction/rate-limiting/rate-limit-exceeded-errors\"></span>  If the client is sending too many requests, the API endpoints respond with the HTTP status code `429 Too Many Requests` and the following body:  ``` {     \"error\": {         \"type\": \"rate-limit-exceeded\",         \"message\": \"You have exceeded the rate limit of ... requests per second\"     } } ```  ### Retrying rate-limited requests with exponential backoff <span id=\"/introduction/rate-limiting/retrying-rate-limited-requests-with-exponential-backoff\"></span>  If the client receives the rate limit error, it should wait a certain period of time and then retry the request. If the error happens again, the client should double the wait period and retry the request, and so on. This algorithm is known as _exponential backoff_ and it can be described using the following pseudo-code:  1. Define a variable `DELAY=500` 2. Send the HTTP request to the API endpoint 3. If the response has status code not equal to `429` then you are done. Otherwise:    * Wait for a period of time chosen randomly from the interval `DELAY` to `2*DELAY` milliseconds    * Double the future wait period by setting `DELAY = 2*DELAY`    * Continue with step 2  If all requests sent by the client implement the above steps, the client will automatically use the maximum available bandwidth for its requests.  Note that the Apify API clients [for JavaScript](https://docs.apify.com/api/client/js) and [for Python](https://docs.apify.com/api/client/python) use the exponential backoff algorithm transparently, so that you do not need to worry about it.  ## Referring to resources <span id=\"/introduction/referring-to-resources\"></span>  There are three main ways to refer to a resource you're accessing via API.  - the resource ID (e.g. `iKkPcIgVvwmztduf8`) - `username~resourcename` - when using this access method, you will need to use your API token, and access will only work if you have the correct permissions. - `~resourcename` - for this, you need to use an API token, and the `resourcename` refers to a resource in the API token owner's account. 

API version: v2-2026-06-16T064758Z
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apifyclient

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RunShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunShort{}

// RunShort struct for RunShort
type RunShort struct {
	Id string `json:"id"`
	ActId string `json:"actId"`
	UserId *string `json:"userId,omitempty"`
	ActorTaskId NullableString `json:"actorTaskId,omitempty"`
	Status ActorJobStatus `json:"status"`
	StartedAt time.Time `json:"startedAt"`
	FinishedAt NullableTime `json:"finishedAt,omitempty"`
	BuildId string `json:"buildId"`
	BuildNumber *string `json:"buildNumber,omitempty" validate:"regexp=^([0-9]|[1-9][0-9])\\\\.([0-9]|[1-9][0-9])(\\\\.[1-9][0-9]{0,4})$"`
	BuildNumberInt *int32 `json:"buildNumberInt,omitempty"`
	Meta RunMeta `json:"meta"`
	UsageTotalUsd float32 `json:"usageTotalUsd"`
	DefaultKeyValueStoreId string `json:"defaultKeyValueStoreId"`
	DefaultDatasetId string `json:"defaultDatasetId"`
	DefaultRequestQueueId string `json:"defaultRequestQueueId"`
}

type _RunShort RunShort

// NewRunShort instantiates a new RunShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunShort(id string, actId string, status ActorJobStatus, startedAt time.Time, buildId string, meta RunMeta, usageTotalUsd float32, defaultKeyValueStoreId string, defaultDatasetId string, defaultRequestQueueId string) *RunShort {
	this := RunShort{}
	this.Id = id
	this.ActId = actId
	this.Status = status
	this.StartedAt = startedAt
	this.BuildId = buildId
	this.Meta = meta
	this.UsageTotalUsd = usageTotalUsd
	this.DefaultKeyValueStoreId = defaultKeyValueStoreId
	this.DefaultDatasetId = defaultDatasetId
	this.DefaultRequestQueueId = defaultRequestQueueId
	return &this
}

// NewRunShortWithDefaults instantiates a new RunShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunShortWithDefaults() *RunShort {
	this := RunShort{}
	return &this
}

// GetId returns the Id field value
func (o *RunShort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunShort) SetId(v string) {
	o.Id = v
}

// GetActId returns the ActId field value
func (o *RunShort) GetActId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActId
}

// GetActIdOk returns a tuple with the ActId field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetActIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActId, true
}

// SetActId sets field value
func (o *RunShort) SetActId(v string) {
	o.ActId = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RunShort) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunShort) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RunShort) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *RunShort) SetUserId(v string) {
	o.UserId = &v
}

// GetActorTaskId returns the ActorTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunShort) GetActorTaskId() string {
	if o == nil || IsNil(o.ActorTaskId.Get()) {
		var ret string
		return ret
	}
	return *o.ActorTaskId.Get()
}

// GetActorTaskIdOk returns a tuple with the ActorTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunShort) GetActorTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorTaskId.Get(), o.ActorTaskId.IsSet()
}

// HasActorTaskId returns a boolean if a field has been set.
func (o *RunShort) HasActorTaskId() bool {
	if o != nil && o.ActorTaskId.IsSet() {
		return true
	}

	return false
}

// SetActorTaskId gets a reference to the given NullableString and assigns it to the ActorTaskId field.
func (o *RunShort) SetActorTaskId(v string) {
	o.ActorTaskId.Set(&v)
}
// SetActorTaskIdNil sets the value for ActorTaskId to be an explicit nil
func (o *RunShort) SetActorTaskIdNil() {
	o.ActorTaskId.Set(nil)
}

// UnsetActorTaskId ensures that no value is present for ActorTaskId, not even an explicit nil
func (o *RunShort) UnsetActorTaskId() {
	o.ActorTaskId.Unset()
}

// GetStatus returns the Status field value
func (o *RunShort) GetStatus() ActorJobStatus {
	if o == nil {
		var ret ActorJobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetStatusOk() (*ActorJobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RunShort) SetStatus(v ActorJobStatus) {
	o.Status = v
}

// GetStartedAt returns the StartedAt field value
func (o *RunShort) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *RunShort) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunShort) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunShort) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *RunShort) HasFinishedAt() bool {
	if o != nil && o.FinishedAt.IsSet() {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given NullableTime and assigns it to the FinishedAt field.
func (o *RunShort) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}
// SetFinishedAtNil sets the value for FinishedAt to be an explicit nil
func (o *RunShort) SetFinishedAtNil() {
	o.FinishedAt.Set(nil)
}

// UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
func (o *RunShort) UnsetFinishedAt() {
	o.FinishedAt.Unset()
}

// GetBuildId returns the BuildId field value
func (o *RunShort) GetBuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetBuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *RunShort) SetBuildId(v string) {
	o.BuildId = v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *RunShort) GetBuildNumber() string {
	if o == nil || IsNil(o.BuildNumber) {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunShort) GetBuildNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BuildNumber) {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *RunShort) HasBuildNumber() bool {
	if o != nil && !IsNil(o.BuildNumber) {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *RunShort) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// GetBuildNumberInt returns the BuildNumberInt field value if set, zero value otherwise.
func (o *RunShort) GetBuildNumberInt() int32 {
	if o == nil || IsNil(o.BuildNumberInt) {
		var ret int32
		return ret
	}
	return *o.BuildNumberInt
}

// GetBuildNumberIntOk returns a tuple with the BuildNumberInt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunShort) GetBuildNumberIntOk() (*int32, bool) {
	if o == nil || IsNil(o.BuildNumberInt) {
		return nil, false
	}
	return o.BuildNumberInt, true
}

// HasBuildNumberInt returns a boolean if a field has been set.
func (o *RunShort) HasBuildNumberInt() bool {
	if o != nil && !IsNil(o.BuildNumberInt) {
		return true
	}

	return false
}

// SetBuildNumberInt gets a reference to the given int32 and assigns it to the BuildNumberInt field.
func (o *RunShort) SetBuildNumberInt(v int32) {
	o.BuildNumberInt = &v
}

// GetMeta returns the Meta field value
func (o *RunShort) GetMeta() RunMeta {
	if o == nil {
		var ret RunMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetMetaOk() (*RunMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *RunShort) SetMeta(v RunMeta) {
	o.Meta = v
}

// GetUsageTotalUsd returns the UsageTotalUsd field value
func (o *RunShort) GetUsageTotalUsd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UsageTotalUsd
}

// GetUsageTotalUsdOk returns a tuple with the UsageTotalUsd field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetUsageTotalUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageTotalUsd, true
}

// SetUsageTotalUsd sets field value
func (o *RunShort) SetUsageTotalUsd(v float32) {
	o.UsageTotalUsd = v
}

// GetDefaultKeyValueStoreId returns the DefaultKeyValueStoreId field value
func (o *RunShort) GetDefaultKeyValueStoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultKeyValueStoreId
}

// GetDefaultKeyValueStoreIdOk returns a tuple with the DefaultKeyValueStoreId field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetDefaultKeyValueStoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultKeyValueStoreId, true
}

// SetDefaultKeyValueStoreId sets field value
func (o *RunShort) SetDefaultKeyValueStoreId(v string) {
	o.DefaultKeyValueStoreId = v
}

// GetDefaultDatasetId returns the DefaultDatasetId field value
func (o *RunShort) GetDefaultDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultDatasetId
}

// GetDefaultDatasetIdOk returns a tuple with the DefaultDatasetId field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetDefaultDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDatasetId, true
}

// SetDefaultDatasetId sets field value
func (o *RunShort) SetDefaultDatasetId(v string) {
	o.DefaultDatasetId = v
}

// GetDefaultRequestQueueId returns the DefaultRequestQueueId field value
func (o *RunShort) GetDefaultRequestQueueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultRequestQueueId
}

// GetDefaultRequestQueueIdOk returns a tuple with the DefaultRequestQueueId field value
// and a boolean to check if the value has been set.
func (o *RunShort) GetDefaultRequestQueueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRequestQueueId, true
}

// SetDefaultRequestQueueId sets field value
func (o *RunShort) SetDefaultRequestQueueId(v string) {
	o.DefaultRequestQueueId = v
}

func (o RunShort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["actId"] = o.ActId
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if o.ActorTaskId.IsSet() {
		toSerialize["actorTaskId"] = o.ActorTaskId.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["startedAt"] = o.StartedAt
	if o.FinishedAt.IsSet() {
		toSerialize["finishedAt"] = o.FinishedAt.Get()
	}
	toSerialize["buildId"] = o.BuildId
	if !IsNil(o.BuildNumber) {
		toSerialize["buildNumber"] = o.BuildNumber
	}
	if !IsNil(o.BuildNumberInt) {
		toSerialize["buildNumberInt"] = o.BuildNumberInt
	}
	toSerialize["meta"] = o.Meta
	toSerialize["usageTotalUsd"] = o.UsageTotalUsd
	toSerialize["defaultKeyValueStoreId"] = o.DefaultKeyValueStoreId
	toSerialize["defaultDatasetId"] = o.DefaultDatasetId
	toSerialize["defaultRequestQueueId"] = o.DefaultRequestQueueId
	return toSerialize, nil
}

func (o *RunShort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"actId",
		"status",
		"startedAt",
		"buildId",
		"meta",
		"usageTotalUsd",
		"defaultKeyValueStoreId",
		"defaultDatasetId",
		"defaultRequestQueueId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRunShort := _RunShort{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunShort)

	if err != nil {
		return err
	}

	*o = RunShort(varRunShort)

	return err
}

type NullableRunShort struct {
	value *RunShort
	isSet bool
}

func (v NullableRunShort) Get() *RunShort {
	return v.value
}

func (v *NullableRunShort) Set(val *RunShort) {
	v.value = val
	v.isSet = true
}

func (v NullableRunShort) IsSet() bool {
	return v.isSet
}

func (v *NullableRunShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunShort(val *RunShort) *NullableRunShort {
	return &NullableRunShort{value: val, isSet: true}
}

func (v NullableRunShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


