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

// checks if the Build type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Build{}

// Build struct for Build
type Build struct {
	Id string `json:"id"`
	ActId string `json:"actId"`
	UserId string `json:"userId"`
	StartedAt time.Time `json:"startedAt"`
	FinishedAt NullableTime `json:"finishedAt,omitempty"`
	Status ActorJobStatus `json:"status"`
	Meta BuildsMeta `json:"meta"`
	Stats NullableBuildStats `json:"stats,omitempty"`
	Options NullableBuildOptions `json:"options,omitempty"`
	Usage NullableBuildUsage `json:"usage,omitempty"`
	// Total cost in USD for this build. Requires authentication token to access.
	UsageTotalUsd NullableFloat32 `json:"usageTotalUsd,omitempty"`
	// Platform usage costs breakdown in USD for this build. Requires authentication token to access.
	UsageUsd NullableBuildUsage `json:"usageUsd,omitempty"`
	// Deprecated
	InputSchema NullableString `json:"inputSchema,omitempty"`
	// Deprecated
	Readme NullableString `json:"readme,omitempty"`
	BuildNumber string `json:"buildNumber" validate:"regexp=^([0-9]|[1-9][0-9])\\\\.([0-9]|[1-9][0-9])(\\\\.[1-9][0-9]{0,4})$"`
	ActVersion *BuildActVersion `json:"actVersion,omitempty"`
	ActorDefinition NullableActorDefinition `json:"actorDefinition,omitempty"`
}

type _Build Build

// NewBuild instantiates a new Build object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuild(id string, actId string, userId string, startedAt time.Time, status ActorJobStatus, meta BuildsMeta, buildNumber string) *Build {
	this := Build{}
	this.Id = id
	this.ActId = actId
	this.UserId = userId
	this.StartedAt = startedAt
	this.Status = status
	this.Meta = meta
	this.BuildNumber = buildNumber
	return &this
}

// NewBuildWithDefaults instantiates a new Build object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildWithDefaults() *Build {
	this := Build{}
	return &this
}

// GetId returns the Id field value
func (o *Build) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Build) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Build) SetId(v string) {
	o.Id = v
}

// GetActId returns the ActId field value
func (o *Build) GetActId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActId
}

// GetActIdOk returns a tuple with the ActId field value
// and a boolean to check if the value has been set.
func (o *Build) GetActIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActId, true
}

// SetActId sets field value
func (o *Build) SetActId(v string) {
	o.ActId = v
}

// GetUserId returns the UserId field value
func (o *Build) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Build) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Build) SetUserId(v string) {
	o.UserId = v
}

// GetStartedAt returns the StartedAt field value
func (o *Build) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *Build) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *Build) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Build) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Build) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *Build) HasFinishedAt() bool {
	if o != nil && o.FinishedAt.IsSet() {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given NullableTime and assigns it to the FinishedAt field.
func (o *Build) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}
// SetFinishedAtNil sets the value for FinishedAt to be an explicit nil
func (o *Build) SetFinishedAtNil() {
	o.FinishedAt.Set(nil)
}

// UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
func (o *Build) UnsetFinishedAt() {
	o.FinishedAt.Unset()
}

// GetStatus returns the Status field value
func (o *Build) GetStatus() ActorJobStatus {
	if o == nil {
		var ret ActorJobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Build) GetStatusOk() (*ActorJobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Build) SetStatus(v ActorJobStatus) {
	o.Status = v
}

// GetMeta returns the Meta field value
func (o *Build) GetMeta() BuildsMeta {
	if o == nil {
		var ret BuildsMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *Build) GetMetaOk() (*BuildsMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *Build) SetMeta(v BuildsMeta) {
	o.Meta = v
}

// GetStats returns the Stats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Build) GetStats() BuildStats {
	if o == nil || IsNil(o.Stats.Get()) {
		var ret BuildStats
		return ret
	}
	return *o.Stats.Get()
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Build) GetStatsOk() (*BuildStats, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stats.Get(), o.Stats.IsSet()
}

// HasStats returns a boolean if a field has been set.
func (o *Build) HasStats() bool {
	if o != nil && o.Stats.IsSet() {
		return true
	}

	return false
}

// SetStats gets a reference to the given NullableBuildStats and assigns it to the Stats field.
func (o *Build) SetStats(v BuildStats) {
	o.Stats.Set(&v)
}
// SetStatsNil sets the value for Stats to be an explicit nil
func (o *Build) SetStatsNil() {
	o.Stats.Set(nil)
}

// UnsetStats ensures that no value is present for Stats, not even an explicit nil
func (o *Build) UnsetStats() {
	o.Stats.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Build) GetOptions() BuildOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret BuildOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Build) GetOptionsOk() (*BuildOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *Build) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableBuildOptions and assigns it to the Options field.
func (o *Build) SetOptions(v BuildOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *Build) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *Build) UnsetOptions() {
	o.Options.Unset()
}

// GetUsage returns the Usage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Build) GetUsage() BuildUsage {
	if o == nil || IsNil(o.Usage.Get()) {
		var ret BuildUsage
		return ret
	}
	return *o.Usage.Get()
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Build) GetUsageOk() (*BuildUsage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage.Get(), o.Usage.IsSet()
}

// HasUsage returns a boolean if a field has been set.
func (o *Build) HasUsage() bool {
	if o != nil && o.Usage.IsSet() {
		return true
	}

	return false
}

// SetUsage gets a reference to the given NullableBuildUsage and assigns it to the Usage field.
func (o *Build) SetUsage(v BuildUsage) {
	o.Usage.Set(&v)
}
// SetUsageNil sets the value for Usage to be an explicit nil
func (o *Build) SetUsageNil() {
	o.Usage.Set(nil)
}

// UnsetUsage ensures that no value is present for Usage, not even an explicit nil
func (o *Build) UnsetUsage() {
	o.Usage.Unset()
}

// GetUsageTotalUsd returns the UsageTotalUsd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Build) GetUsageTotalUsd() float32 {
	if o == nil || IsNil(o.UsageTotalUsd.Get()) {
		var ret float32
		return ret
	}
	return *o.UsageTotalUsd.Get()
}

// GetUsageTotalUsdOk returns a tuple with the UsageTotalUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Build) GetUsageTotalUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageTotalUsd.Get(), o.UsageTotalUsd.IsSet()
}

// HasUsageTotalUsd returns a boolean if a field has been set.
func (o *Build) HasUsageTotalUsd() bool {
	if o != nil && o.UsageTotalUsd.IsSet() {
		return true
	}

	return false
}

// SetUsageTotalUsd gets a reference to the given NullableFloat32 and assigns it to the UsageTotalUsd field.
func (o *Build) SetUsageTotalUsd(v float32) {
	o.UsageTotalUsd.Set(&v)
}
// SetUsageTotalUsdNil sets the value for UsageTotalUsd to be an explicit nil
func (o *Build) SetUsageTotalUsdNil() {
	o.UsageTotalUsd.Set(nil)
}

// UnsetUsageTotalUsd ensures that no value is present for UsageTotalUsd, not even an explicit nil
func (o *Build) UnsetUsageTotalUsd() {
	o.UsageTotalUsd.Unset()
}

// GetUsageUsd returns the UsageUsd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Build) GetUsageUsd() BuildUsage {
	if o == nil || IsNil(o.UsageUsd.Get()) {
		var ret BuildUsage
		return ret
	}
	return *o.UsageUsd.Get()
}

// GetUsageUsdOk returns a tuple with the UsageUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Build) GetUsageUsdOk() (*BuildUsage, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageUsd.Get(), o.UsageUsd.IsSet()
}

// HasUsageUsd returns a boolean if a field has been set.
func (o *Build) HasUsageUsd() bool {
	if o != nil && o.UsageUsd.IsSet() {
		return true
	}

	return false
}

// SetUsageUsd gets a reference to the given NullableBuildUsage and assigns it to the UsageUsd field.
func (o *Build) SetUsageUsd(v BuildUsage) {
	o.UsageUsd.Set(&v)
}
// SetUsageUsdNil sets the value for UsageUsd to be an explicit nil
func (o *Build) SetUsageUsdNil() {
	o.UsageUsd.Set(nil)
}

// UnsetUsageUsd ensures that no value is present for UsageUsd, not even an explicit nil
func (o *Build) UnsetUsageUsd() {
	o.UsageUsd.Unset()
}

// GetInputSchema returns the InputSchema field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *Build) GetInputSchema() string {
	if o == nil || IsNil(o.InputSchema.Get()) {
		var ret string
		return ret
	}
	return *o.InputSchema.Get()
}

// GetInputSchemaOk returns a tuple with the InputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *Build) GetInputSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputSchema.Get(), o.InputSchema.IsSet()
}

// HasInputSchema returns a boolean if a field has been set.
func (o *Build) HasInputSchema() bool {
	if o != nil && o.InputSchema.IsSet() {
		return true
	}

	return false
}

// SetInputSchema gets a reference to the given NullableString and assigns it to the InputSchema field.
// Deprecated
func (o *Build) SetInputSchema(v string) {
	o.InputSchema.Set(&v)
}
// SetInputSchemaNil sets the value for InputSchema to be an explicit nil
func (o *Build) SetInputSchemaNil() {
	o.InputSchema.Set(nil)
}

// UnsetInputSchema ensures that no value is present for InputSchema, not even an explicit nil
func (o *Build) UnsetInputSchema() {
	o.InputSchema.Unset()
}

// GetReadme returns the Readme field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *Build) GetReadme() string {
	if o == nil || IsNil(o.Readme.Get()) {
		var ret string
		return ret
	}
	return *o.Readme.Get()
}

// GetReadmeOk returns a tuple with the Readme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *Build) GetReadmeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Readme.Get(), o.Readme.IsSet()
}

// HasReadme returns a boolean if a field has been set.
func (o *Build) HasReadme() bool {
	if o != nil && o.Readme.IsSet() {
		return true
	}

	return false
}

// SetReadme gets a reference to the given NullableString and assigns it to the Readme field.
// Deprecated
func (o *Build) SetReadme(v string) {
	o.Readme.Set(&v)
}
// SetReadmeNil sets the value for Readme to be an explicit nil
func (o *Build) SetReadmeNil() {
	o.Readme.Set(nil)
}

// UnsetReadme ensures that no value is present for Readme, not even an explicit nil
func (o *Build) UnsetReadme() {
	o.Readme.Unset()
}

// GetBuildNumber returns the BuildNumber field value
func (o *Build) GetBuildNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value
// and a boolean to check if the value has been set.
func (o *Build) GetBuildNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildNumber, true
}

// SetBuildNumber sets field value
func (o *Build) SetBuildNumber(v string) {
	o.BuildNumber = v
}

// GetActVersion returns the ActVersion field value if set, zero value otherwise.
func (o *Build) GetActVersion() BuildActVersion {
	if o == nil || IsNil(o.ActVersion) {
		var ret BuildActVersion
		return ret
	}
	return *o.ActVersion
}

// GetActVersionOk returns a tuple with the ActVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetActVersionOk() (*BuildActVersion, bool) {
	if o == nil || IsNil(o.ActVersion) {
		return nil, false
	}
	return o.ActVersion, true
}

// HasActVersion returns a boolean if a field has been set.
func (o *Build) HasActVersion() bool {
	if o != nil && !IsNil(o.ActVersion) {
		return true
	}

	return false
}

// SetActVersion gets a reference to the given BuildActVersion and assigns it to the ActVersion field.
func (o *Build) SetActVersion(v BuildActVersion) {
	o.ActVersion = &v
}

// GetActorDefinition returns the ActorDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Build) GetActorDefinition() ActorDefinition {
	if o == nil || IsNil(o.ActorDefinition.Get()) {
		var ret ActorDefinition
		return ret
	}
	return *o.ActorDefinition.Get()
}

// GetActorDefinitionOk returns a tuple with the ActorDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Build) GetActorDefinitionOk() (*ActorDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorDefinition.Get(), o.ActorDefinition.IsSet()
}

// HasActorDefinition returns a boolean if a field has been set.
func (o *Build) HasActorDefinition() bool {
	if o != nil && o.ActorDefinition.IsSet() {
		return true
	}

	return false
}

// SetActorDefinition gets a reference to the given NullableActorDefinition and assigns it to the ActorDefinition field.
func (o *Build) SetActorDefinition(v ActorDefinition) {
	o.ActorDefinition.Set(&v)
}
// SetActorDefinitionNil sets the value for ActorDefinition to be an explicit nil
func (o *Build) SetActorDefinitionNil() {
	o.ActorDefinition.Set(nil)
}

// UnsetActorDefinition ensures that no value is present for ActorDefinition, not even an explicit nil
func (o *Build) UnsetActorDefinition() {
	o.ActorDefinition.Unset()
}

func (o Build) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Build) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["actId"] = o.ActId
	toSerialize["userId"] = o.UserId
	toSerialize["startedAt"] = o.StartedAt
	if o.FinishedAt.IsSet() {
		toSerialize["finishedAt"] = o.FinishedAt.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["meta"] = o.Meta
	if o.Stats.IsSet() {
		toSerialize["stats"] = o.Stats.Get()
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if o.Usage.IsSet() {
		toSerialize["usage"] = o.Usage.Get()
	}
	if o.UsageTotalUsd.IsSet() {
		toSerialize["usageTotalUsd"] = o.UsageTotalUsd.Get()
	}
	if o.UsageUsd.IsSet() {
		toSerialize["usageUsd"] = o.UsageUsd.Get()
	}
	if o.InputSchema.IsSet() {
		toSerialize["inputSchema"] = o.InputSchema.Get()
	}
	if o.Readme.IsSet() {
		toSerialize["readme"] = o.Readme.Get()
	}
	toSerialize["buildNumber"] = o.BuildNumber
	if !IsNil(o.ActVersion) {
		toSerialize["actVersion"] = o.ActVersion
	}
	if o.ActorDefinition.IsSet() {
		toSerialize["actorDefinition"] = o.ActorDefinition.Get()
	}
	return toSerialize, nil
}

func (o *Build) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"actId",
		"userId",
		"startedAt",
		"status",
		"meta",
		"buildNumber",
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

	varBuild := _Build{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuild)

	if err != nil {
		return err
	}

	*o = Build(varBuild)

	return err
}

type NullableBuild struct {
	value *Build
	isSet bool
}

func (v NullableBuild) Get() *Build {
	return v.value
}

func (v *NullableBuild) Set(val *Build) {
	v.value = val
	v.isSet = true
}

func (v NullableBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuild(val *Build) *NullableBuild {
	return &NullableBuild{value: val, isSet: true}
}

func (v NullableBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


