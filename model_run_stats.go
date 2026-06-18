/*
Apify API

 The Apify API (version 2) provides programmatic access to the [Apify platform](https://docs.apify.com). The API is organized around [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) HTTP endpoints.  You can download the complete OpenAPI schema of Apify API in the [YAML](http://docs.apify.com/api/openapi.yaml) or [JSON](http://docs.apify.com/api/openapi.json) formats. The source code is also available on [GitHub](https://github.com/apify/apify-docs/tree/master/apify-api/openapi).  All requests and responses (including errors) are encoded in [JSON](http://www.json.org/) format with UTF-8 encoding, with a few exceptions that are explicitly described in the reference.  - To access the API using [Node.js](https://nodejs.org/en/), we recommend the [`apify-client`](https://docs.apify.com/api/client/js) [NPM package](https://www.npmjs.com/package/apify-client). - To access the API using [Python](https://www.python.org/), we recommend the [`apify-client`](https://docs.apify.com/api/client/python) [PyPI package](https://pypi.org/project/apify-client/).  The clients' functions correspond to the API endpoints and have the same parameters. This simplifies development of apps that depend on the Apify platform.  :::note Important Request Details  - `Content-Type` header: For requests with a JSON body, you must include the `Content-Type: application/json` header.  - Method override: You can override the HTTP method using the `method` query parameter. This is useful for clients that can only send `GET` requests. For example, to call a `POST` endpoint, append `?method=POST` to the URL of your `GET` request.  :::  ## Authentication <span id=\"/introduction/authentication\"></span>  **You can find your API token on the [Integrations](https://console.apify.com/settings/integrations) page in the Apify Console.**  To use your token in a request, either:  - Add the token to your request's `Authorization` header as `Bearer <token>`. E.g., `Authorization: Bearer xxxxxxx`. [More info](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization). (Recommended). - Add it as the `token` parameter to your request URL. (Less secure).  Using your token in the request header is more secure than using it as a URL parameter because URLs are often stored in browser history and server logs. This creates a chance for someone unauthorized to access your API token.  **Never share your API token or password with untrusted parties!**  For more information, see our [integrations](https://docs.apify.com/platform/integrations) documentation.  ### Agentic payments  AI agents can authenticate and pay for Actor runs without an Apify account using agentic payments. Instead of an API token, the request carries a payment credential that both authorizes and pays for the call. Apify supports the [x402 protocol](https://docs.apify.com/platform/integrations/x402) (`PAYMENT-SIGNATURE` header) and [Skyfire](https://docs.apify.com/platform/integrations/skyfire) (`skyfire-pay-id` header).  ## Basic usage <span id=\"/introduction/basic-usage\"></span>  To run an Actor, send a POST request to the [Run Actor](#/reference/actors/run-collection/run-actor) endpoint using either the Actor ID code (e.g. `vKg4IjxZbEYTYeW8T`) or its name (e.g. `janedoe~my-actor`):  `https://api.apify.com/v2/actors/[actor_id]/runs`  If the Actor is not runnable anonymously, you will receive a 401 or 403 [response code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status). This means you need to add your [secret API token](https://console.apify.com/account#/integrations) to the request's `Authorization` header ([recommended](#/introduction/authentication)) or as a URL query parameter `?token=[your_token]` (less secure).  Optionally, you can include the query parameters described in the [Run Actor](#/reference/actors/run-collection/run-actor) section to customize your run.  If you're using Node.js, the best way to run an Actor is using the `Apify.call()` method from the [Apify SDK](https://sdk.apify.com/docs/api/apify#apifycallactid-input-options). It runs the Actor using the account you are currently logged into (determined by the [secret API token](https://console.apify.com/account#/integrations)). The result is an [Actor run object](https://sdk.apify.com/docs/typedefs/actor-run) and its output (if any).  A typical workflow is as follows:  1. Run an Actor or task using the [Run Actor](#/reference/actors/run-collection/run-actor) or [Run task](#/reference/actor-tasks/run-collection/run-task) API endpoints. 2. Monitor the Actor run by periodically polling its progress using the [Get run](#/reference/actor-runs/run-object-and-its-storages/get-run) API endpoint. 3. Fetch the results from the [Get items](#/reference/datasets/item-collection/get-items) API endpoint using the `defaultDatasetId`, which you receive in the Run request response. Additional data may be stored in a key-value store. You can fetch them from the [Get record](#/reference/key-value-stores/record/get-record) API endpoint using the `defaultKeyValueStoreId` and the store's `key`.  **Note**: Instead of periodic polling, you can also run your [Actor](#/reference/actors/run-actor-synchronously) or [task](#/reference/actor-tasks/runs-collection/run-task-synchronously) synchronously. This will ensure that the request waits for 300 seconds (5 minutes) for the run to finish and returns its output. If the run takes longer, the request will time out and throw an error.  ## Legacy `/v2/acts/` URL prefix <span id=\"/introduction/legacy-acts-prefix\"></span>  The `/v2/acts/` prefix is deprecated but still fully functional, and  such endpoint routes to the same handler as its `/v2/actors/...` counterpart.  New integrations should use the canonical /v2/actors/ prefix,  but existing clients keep working without changes.  ## Response structure <span id=\"/introduction/response-structure\"></span>  Most API endpoints return a JSON object with the `data` property:  ``` {     \"data\": {         ...     } } ```  However, there are a few explicitly described exceptions, such as [Get dataset items](#/reference/datasets/item-collection/get-items) or Key-value store [Get record](#/reference/key-value-stores/record/get-record) API endpoints, which return data in other formats. In case of an error, the response has the HTTP status code in the range of 4xx or 5xx and the `data` property is replaced with `error`. For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  See [Errors](#/introduction/errors) for more details.  ## Pagination <span id=\"/introduction/pagination\"></span>  All API endpoints that return a list of records (e.g. [Get list of Actors](#/reference/actors/actor-collection/get-list-of-actors)) enforce pagination in order to limit the size of their responses.  Most of these API endpoints are paginated using the `offset` and `limit` query parameters. The only exception is [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys), which is paginated using the `exclusiveStartKey` query parameter.  **IMPORTANT**: Each API endpoint that supports pagination enforces a certain maximum value for the `limit` parameter, in order to reduce the load on Apify servers. The maximum limit could change in future so you should never rely on a specific value and check the responses of these API endpoints.  ### Using offset <span id=\"/introduction/pagination/using-offset\"></span>  Most API endpoints that return a list of records enable pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number of items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>Skips a number of items from the beginning of the list, e.g. <code>offset=100</code>.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td>     By default, items are sorted in the order in which they were created or added to the list.     This feature is useful when fetching all the items, because it ensures that items     created after the client started the pagination will not be skipped.     If you specify the <code>desc=1</code> parameter, the items will be returned in the reverse order,     i.e. from the newest to the oldest items.     </td>   </tr> </table>  The response of these API endpoints is always a JSON object with the following structure:  ``` {     \"data\": {         \"total\": 2560,         \"offset\": 250,         \"limit\": 1000,         \"count\": 1000,         \"desc\": false,         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>total</code></td>     <td>The total number of items available in the list.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>The number of items that were skipped at the start.     This is equal to the <code>offset</code> query parameter if it was provided, otherwise it is <code>0</code>.</td>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular API endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>count</code></td>     <td>The actual number of items returned in the HTTP response.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td><code>true</code> if data were requested in descending order and <code>false</code> otherwise.</td>   </tr>   <tr>     <td><code>items</code></td>     <td>An array of requested items.</td>   </tr> </table>  ### Using key <span id=\"/introduction/pagination/using-key\"></span>  The records in the [key-value store](https://docs.apify.com/platform/storage/key-value-store) are not ordered based on numerical indexes, but rather by their keys in the UTF-8 binary order. Therefore the [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys) API endpoint only supports pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>Skips all records with keys up to the given key including the given key,     in the UTF-8 binary order.</td>   </tr> </table>  The response of the API endpoint is always a JSON object with following structure:  ``` {     \"data\": {         \"limit\": 1000,         \"isTruncated\": true,         \"exclusiveStartKey\": \"my-key\",         \"nextExclusiveStartKey\": \"some-other-key\",         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>isTruncated</code></td>     <td><code>true</code> if there are more items left to be queried. Otherwise <code>false</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>The last key that was skipped at the start. Is `null` for the first page.</td>   </tr>   <tr>     <td><code>nextExclusiveStartKey</code></td>     <td>The value for the <code>exclusiveStartKey</code> parameter to query the next page of items.</td>   </tr> </table>  ## Errors <span id=\"/introduction/errors\"></span>  The Apify API uses common HTTP status codes: `2xx` range for success, `4xx` range for errors caused by the caller (invalid requests) and `5xx` range for server errors (these are rare). Each error response contains a JSON object defining the `error` property, which is an object with the `type` and `message` properties that contain the error code and a human-readable error description, respectively.  For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  Here is the table of the most common errors that can occur for many API endpoints:  <table>   <tr>     <th>status</th>     <th>type</th>     <th>message</th>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-request</code></td>     <td>POST data must be a JSON object</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-value</code></td>     <td>Invalid value provided: Comments required</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-record-key</code></td>     <td>Record key contains invalid character</td>   </tr>   <tr>     <td><code>401</code></td>     <td><code>token-not-provided</code></td>     <td>Authentication token was not provided</td>   </tr>   <tr>     <td><code>404</code></td>     <td><code>record-not-found</code></td>     <td>Store was not found</td>   </tr>   <tr>     <td><code>429</code></td>     <td><code>rate-limit-exceeded</code></td>     <td>You have exceeded the rate limit of ... requests per second</td>   </tr>   <tr>     <td><code>405</code></td>     <td><code>method-not-allowed</code></td>     <td>This API endpoint can only be accessed using the following HTTP methods: OPTIONS, POST</td>   </tr> </table>  ## Rate limiting <span id=\"/introduction/rate-limiting\"></span>  All API endpoints limit the rate of requests in order to prevent overloading of Apify servers by misbehaving clients.  There are two kinds of rate limits - a global rate limit and a per-resource rate limit.  ### Global rate limit <span id=\"/introduction/rate-limiting/global-rate-limit\"></span>  The global rate limit is set to _250 000 requests per minute_. For [authenticated](#/introduction/authentication) requests, it is counted per user, and for unauthenticated requests, it is counted per IP address.  ### Per-resource rate limit <span id=\"/introduction/rate-limiting/per-resource-rate-limit\"></span>  The default per-resource rate limit is _60 requests per second per resource_, which in this context means a single Actor, a single Actor run, a single dataset, single key-value store etc. The default rate limit is applied to every API endpoint except a few select ones, which have higher rate limits. Each API endpoint returns its rate limit in `X-RateLimit-Limit` header.  These endpoints have a rate limit of _200 requests per second per resource_:  * CRUD ([get](#/reference/key-value-stores/record/get-record),   [put](#/reference/key-value-stores/record/put-record),   [delete](#/reference/key-value-stores/record/delete-record))   operations on key-value store records  These endpoints have a rate limit of _400 requests per second per resource_: * [Run Actor](#/reference/actors/run-collection/run-actor) * [Run Actor task asynchronously](#/reference/actor-tasks/runs-collection/run-task-asynchronously) * [Run Actor task synchronously](#/reference/actor-tasks/runs-collection/run-task-synchronously) * [Metamorph Actor run](#/reference/actors/metamorph-run/metamorph-run) * [Push items](#/reference/datasets/item-collection/put-items) to dataset * CRUD   ([add](#/reference/request-queues/request-collection/add-request),   [get](#/reference/request-queues/request-collection/get-request),   [update](#/reference/request-queues/request-collection/update-request),   [delete](#/reference/request-queues/request-collection/delete-request))   operations on requests in request queues  ### Rate limit exceeded errors <span id=\"/introduction/rate-limiting/rate-limit-exceeded-errors\"></span>  If the client is sending too many requests, the API endpoints respond with the HTTP status code `429 Too Many Requests` and the following body:  ``` {     \"error\": {         \"type\": \"rate-limit-exceeded\",         \"message\": \"You have exceeded the rate limit of ... requests per second\"     } } ```  ### Retrying rate-limited requests with exponential backoff <span id=\"/introduction/rate-limiting/retrying-rate-limited-requests-with-exponential-backoff\"></span>  If the client receives the rate limit error, it should wait a certain period of time and then retry the request. If the error happens again, the client should double the wait period and retry the request, and so on. This algorithm is known as _exponential backoff_ and it can be described using the following pseudo-code:  1. Define a variable `DELAY=500` 2. Send the HTTP request to the API endpoint 3. If the response has status code not equal to `429` then you are done. Otherwise:    * Wait for a period of time chosen randomly from the interval `DELAY` to `2*DELAY` milliseconds    * Double the future wait period by setting `DELAY = 2*DELAY`    * Continue with step 2  If all requests sent by the client implement the above steps, the client will automatically use the maximum available bandwidth for its requests.  Note that the Apify API clients [for JavaScript](https://docs.apify.com/api/client/js) and [for Python](https://docs.apify.com/api/client/python) use the exponential backoff algorithm transparently, so that you do not need to worry about it.  ## Referring to resources <span id=\"/introduction/referring-to-resources\"></span>  There are three main ways to refer to a resource you're accessing via API.  - the resource ID (e.g. `iKkPcIgVvwmztduf8`) - `username~resourcename` - when using this access method, you will need to use your API token, and access will only work if you have the correct permissions. - `~resourcename` - for this, you need to use an API token, and the `resourcename` refers to a resource in the API token owner's account. 

API version: v2-2026-06-16T064758Z
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apifyclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RunStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStats{}

// RunStats struct for RunStats
type RunStats struct {
	InputBodyLen NullableInt32 `json:"inputBodyLen,omitempty"`
	MigrationCount *int32 `json:"migrationCount,omitempty"`
	RebootCount *int32 `json:"rebootCount,omitempty"`
	RestartCount int32 `json:"restartCount"`
	ResurrectCount int32 `json:"resurrectCount"`
	MemAvgBytes *float32 `json:"memAvgBytes,omitempty"`
	MemMaxBytes *int32 `json:"memMaxBytes,omitempty"`
	MemCurrentBytes *int32 `json:"memCurrentBytes,omitempty"`
	CpuAvgUsage *float32 `json:"cpuAvgUsage,omitempty"`
	CpuMaxUsage *float32 `json:"cpuMaxUsage,omitempty"`
	CpuCurrentUsage *float32 `json:"cpuCurrentUsage,omitempty"`
	NetRxBytes *int32 `json:"netRxBytes,omitempty"`
	NetTxBytes *int32 `json:"netTxBytes,omitempty"`
	DurationMillis *int32 `json:"durationMillis,omitempty"`
	RunTimeSecs *float32 `json:"runTimeSecs,omitempty"`
	Metamorph *int32 `json:"metamorph,omitempty"`
	ComputeUnits float32 `json:"computeUnits"`
}

type _RunStats RunStats

// NewRunStats instantiates a new RunStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStats(restartCount int32, resurrectCount int32, computeUnits float32) *RunStats {
	this := RunStats{}
	this.RestartCount = restartCount
	this.ResurrectCount = resurrectCount
	this.ComputeUnits = computeUnits
	return &this
}

// NewRunStatsWithDefaults instantiates a new RunStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStatsWithDefaults() *RunStats {
	this := RunStats{}
	return &this
}

// GetInputBodyLen returns the InputBodyLen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunStats) GetInputBodyLen() int32 {
	if o == nil || IsNil(o.InputBodyLen.Get()) {
		var ret int32
		return ret
	}
	return *o.InputBodyLen.Get()
}

// GetInputBodyLenOk returns a tuple with the InputBodyLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunStats) GetInputBodyLenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputBodyLen.Get(), o.InputBodyLen.IsSet()
}

// HasInputBodyLen returns a boolean if a field has been set.
func (o *RunStats) HasInputBodyLen() bool {
	if o != nil && o.InputBodyLen.IsSet() {
		return true
	}

	return false
}

// SetInputBodyLen gets a reference to the given NullableInt32 and assigns it to the InputBodyLen field.
func (o *RunStats) SetInputBodyLen(v int32) {
	o.InputBodyLen.Set(&v)
}
// SetInputBodyLenNil sets the value for InputBodyLen to be an explicit nil
func (o *RunStats) SetInputBodyLenNil() {
	o.InputBodyLen.Set(nil)
}

// UnsetInputBodyLen ensures that no value is present for InputBodyLen, not even an explicit nil
func (o *RunStats) UnsetInputBodyLen() {
	o.InputBodyLen.Unset()
}

// GetMigrationCount returns the MigrationCount field value if set, zero value otherwise.
func (o *RunStats) GetMigrationCount() int32 {
	if o == nil || IsNil(o.MigrationCount) {
		var ret int32
		return ret
	}
	return *o.MigrationCount
}

// GetMigrationCountOk returns a tuple with the MigrationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetMigrationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MigrationCount) {
		return nil, false
	}
	return o.MigrationCount, true
}

// HasMigrationCount returns a boolean if a field has been set.
func (o *RunStats) HasMigrationCount() bool {
	if o != nil && !IsNil(o.MigrationCount) {
		return true
	}

	return false
}

// SetMigrationCount gets a reference to the given int32 and assigns it to the MigrationCount field.
func (o *RunStats) SetMigrationCount(v int32) {
	o.MigrationCount = &v
}

// GetRebootCount returns the RebootCount field value if set, zero value otherwise.
func (o *RunStats) GetRebootCount() int32 {
	if o == nil || IsNil(o.RebootCount) {
		var ret int32
		return ret
	}
	return *o.RebootCount
}

// GetRebootCountOk returns a tuple with the RebootCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetRebootCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RebootCount) {
		return nil, false
	}
	return o.RebootCount, true
}

// HasRebootCount returns a boolean if a field has been set.
func (o *RunStats) HasRebootCount() bool {
	if o != nil && !IsNil(o.RebootCount) {
		return true
	}

	return false
}

// SetRebootCount gets a reference to the given int32 and assigns it to the RebootCount field.
func (o *RunStats) SetRebootCount(v int32) {
	o.RebootCount = &v
}

// GetRestartCount returns the RestartCount field value
func (o *RunStats) GetRestartCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value
// and a boolean to check if the value has been set.
func (o *RunStats) GetRestartCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartCount, true
}

// SetRestartCount sets field value
func (o *RunStats) SetRestartCount(v int32) {
	o.RestartCount = v
}

// GetResurrectCount returns the ResurrectCount field value
func (o *RunStats) GetResurrectCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResurrectCount
}

// GetResurrectCountOk returns a tuple with the ResurrectCount field value
// and a boolean to check if the value has been set.
func (o *RunStats) GetResurrectCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResurrectCount, true
}

// SetResurrectCount sets field value
func (o *RunStats) SetResurrectCount(v int32) {
	o.ResurrectCount = v
}

// GetMemAvgBytes returns the MemAvgBytes field value if set, zero value otherwise.
func (o *RunStats) GetMemAvgBytes() float32 {
	if o == nil || IsNil(o.MemAvgBytes) {
		var ret float32
		return ret
	}
	return *o.MemAvgBytes
}

// GetMemAvgBytesOk returns a tuple with the MemAvgBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetMemAvgBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.MemAvgBytes) {
		return nil, false
	}
	return o.MemAvgBytes, true
}

// HasMemAvgBytes returns a boolean if a field has been set.
func (o *RunStats) HasMemAvgBytes() bool {
	if o != nil && !IsNil(o.MemAvgBytes) {
		return true
	}

	return false
}

// SetMemAvgBytes gets a reference to the given float32 and assigns it to the MemAvgBytes field.
func (o *RunStats) SetMemAvgBytes(v float32) {
	o.MemAvgBytes = &v
}

// GetMemMaxBytes returns the MemMaxBytes field value if set, zero value otherwise.
func (o *RunStats) GetMemMaxBytes() int32 {
	if o == nil || IsNil(o.MemMaxBytes) {
		var ret int32
		return ret
	}
	return *o.MemMaxBytes
}

// GetMemMaxBytesOk returns a tuple with the MemMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetMemMaxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.MemMaxBytes) {
		return nil, false
	}
	return o.MemMaxBytes, true
}

// HasMemMaxBytes returns a boolean if a field has been set.
func (o *RunStats) HasMemMaxBytes() bool {
	if o != nil && !IsNil(o.MemMaxBytes) {
		return true
	}

	return false
}

// SetMemMaxBytes gets a reference to the given int32 and assigns it to the MemMaxBytes field.
func (o *RunStats) SetMemMaxBytes(v int32) {
	o.MemMaxBytes = &v
}

// GetMemCurrentBytes returns the MemCurrentBytes field value if set, zero value otherwise.
func (o *RunStats) GetMemCurrentBytes() int32 {
	if o == nil || IsNil(o.MemCurrentBytes) {
		var ret int32
		return ret
	}
	return *o.MemCurrentBytes
}

// GetMemCurrentBytesOk returns a tuple with the MemCurrentBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetMemCurrentBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.MemCurrentBytes) {
		return nil, false
	}
	return o.MemCurrentBytes, true
}

// HasMemCurrentBytes returns a boolean if a field has been set.
func (o *RunStats) HasMemCurrentBytes() bool {
	if o != nil && !IsNil(o.MemCurrentBytes) {
		return true
	}

	return false
}

// SetMemCurrentBytes gets a reference to the given int32 and assigns it to the MemCurrentBytes field.
func (o *RunStats) SetMemCurrentBytes(v int32) {
	o.MemCurrentBytes = &v
}

// GetCpuAvgUsage returns the CpuAvgUsage field value if set, zero value otherwise.
func (o *RunStats) GetCpuAvgUsage() float32 {
	if o == nil || IsNil(o.CpuAvgUsage) {
		var ret float32
		return ret
	}
	return *o.CpuAvgUsage
}

// GetCpuAvgUsageOk returns a tuple with the CpuAvgUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetCpuAvgUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.CpuAvgUsage) {
		return nil, false
	}
	return o.CpuAvgUsage, true
}

// HasCpuAvgUsage returns a boolean if a field has been set.
func (o *RunStats) HasCpuAvgUsage() bool {
	if o != nil && !IsNil(o.CpuAvgUsage) {
		return true
	}

	return false
}

// SetCpuAvgUsage gets a reference to the given float32 and assigns it to the CpuAvgUsage field.
func (o *RunStats) SetCpuAvgUsage(v float32) {
	o.CpuAvgUsage = &v
}

// GetCpuMaxUsage returns the CpuMaxUsage field value if set, zero value otherwise.
func (o *RunStats) GetCpuMaxUsage() float32 {
	if o == nil || IsNil(o.CpuMaxUsage) {
		var ret float32
		return ret
	}
	return *o.CpuMaxUsage
}

// GetCpuMaxUsageOk returns a tuple with the CpuMaxUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetCpuMaxUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.CpuMaxUsage) {
		return nil, false
	}
	return o.CpuMaxUsage, true
}

// HasCpuMaxUsage returns a boolean if a field has been set.
func (o *RunStats) HasCpuMaxUsage() bool {
	if o != nil && !IsNil(o.CpuMaxUsage) {
		return true
	}

	return false
}

// SetCpuMaxUsage gets a reference to the given float32 and assigns it to the CpuMaxUsage field.
func (o *RunStats) SetCpuMaxUsage(v float32) {
	o.CpuMaxUsage = &v
}

// GetCpuCurrentUsage returns the CpuCurrentUsage field value if set, zero value otherwise.
func (o *RunStats) GetCpuCurrentUsage() float32 {
	if o == nil || IsNil(o.CpuCurrentUsage) {
		var ret float32
		return ret
	}
	return *o.CpuCurrentUsage
}

// GetCpuCurrentUsageOk returns a tuple with the CpuCurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetCpuCurrentUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.CpuCurrentUsage) {
		return nil, false
	}
	return o.CpuCurrentUsage, true
}

// HasCpuCurrentUsage returns a boolean if a field has been set.
func (o *RunStats) HasCpuCurrentUsage() bool {
	if o != nil && !IsNil(o.CpuCurrentUsage) {
		return true
	}

	return false
}

// SetCpuCurrentUsage gets a reference to the given float32 and assigns it to the CpuCurrentUsage field.
func (o *RunStats) SetCpuCurrentUsage(v float32) {
	o.CpuCurrentUsage = &v
}

// GetNetRxBytes returns the NetRxBytes field value if set, zero value otherwise.
func (o *RunStats) GetNetRxBytes() int32 {
	if o == nil || IsNil(o.NetRxBytes) {
		var ret int32
		return ret
	}
	return *o.NetRxBytes
}

// GetNetRxBytesOk returns a tuple with the NetRxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetNetRxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.NetRxBytes) {
		return nil, false
	}
	return o.NetRxBytes, true
}

// HasNetRxBytes returns a boolean if a field has been set.
func (o *RunStats) HasNetRxBytes() bool {
	if o != nil && !IsNil(o.NetRxBytes) {
		return true
	}

	return false
}

// SetNetRxBytes gets a reference to the given int32 and assigns it to the NetRxBytes field.
func (o *RunStats) SetNetRxBytes(v int32) {
	o.NetRxBytes = &v
}

// GetNetTxBytes returns the NetTxBytes field value if set, zero value otherwise.
func (o *RunStats) GetNetTxBytes() int32 {
	if o == nil || IsNil(o.NetTxBytes) {
		var ret int32
		return ret
	}
	return *o.NetTxBytes
}

// GetNetTxBytesOk returns a tuple with the NetTxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetNetTxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.NetTxBytes) {
		return nil, false
	}
	return o.NetTxBytes, true
}

// HasNetTxBytes returns a boolean if a field has been set.
func (o *RunStats) HasNetTxBytes() bool {
	if o != nil && !IsNil(o.NetTxBytes) {
		return true
	}

	return false
}

// SetNetTxBytes gets a reference to the given int32 and assigns it to the NetTxBytes field.
func (o *RunStats) SetNetTxBytes(v int32) {
	o.NetTxBytes = &v
}

// GetDurationMillis returns the DurationMillis field value if set, zero value otherwise.
func (o *RunStats) GetDurationMillis() int32 {
	if o == nil || IsNil(o.DurationMillis) {
		var ret int32
		return ret
	}
	return *o.DurationMillis
}

// GetDurationMillisOk returns a tuple with the DurationMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetDurationMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMillis) {
		return nil, false
	}
	return o.DurationMillis, true
}

// HasDurationMillis returns a boolean if a field has been set.
func (o *RunStats) HasDurationMillis() bool {
	if o != nil && !IsNil(o.DurationMillis) {
		return true
	}

	return false
}

// SetDurationMillis gets a reference to the given int32 and assigns it to the DurationMillis field.
func (o *RunStats) SetDurationMillis(v int32) {
	o.DurationMillis = &v
}

// GetRunTimeSecs returns the RunTimeSecs field value if set, zero value otherwise.
func (o *RunStats) GetRunTimeSecs() float32 {
	if o == nil || IsNil(o.RunTimeSecs) {
		var ret float32
		return ret
	}
	return *o.RunTimeSecs
}

// GetRunTimeSecsOk returns a tuple with the RunTimeSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetRunTimeSecsOk() (*float32, bool) {
	if o == nil || IsNil(o.RunTimeSecs) {
		return nil, false
	}
	return o.RunTimeSecs, true
}

// HasRunTimeSecs returns a boolean if a field has been set.
func (o *RunStats) HasRunTimeSecs() bool {
	if o != nil && !IsNil(o.RunTimeSecs) {
		return true
	}

	return false
}

// SetRunTimeSecs gets a reference to the given float32 and assigns it to the RunTimeSecs field.
func (o *RunStats) SetRunTimeSecs(v float32) {
	o.RunTimeSecs = &v
}

// GetMetamorph returns the Metamorph field value if set, zero value otherwise.
func (o *RunStats) GetMetamorph() int32 {
	if o == nil || IsNil(o.Metamorph) {
		var ret int32
		return ret
	}
	return *o.Metamorph
}

// GetMetamorphOk returns a tuple with the Metamorph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStats) GetMetamorphOk() (*int32, bool) {
	if o == nil || IsNil(o.Metamorph) {
		return nil, false
	}
	return o.Metamorph, true
}

// HasMetamorph returns a boolean if a field has been set.
func (o *RunStats) HasMetamorph() bool {
	if o != nil && !IsNil(o.Metamorph) {
		return true
	}

	return false
}

// SetMetamorph gets a reference to the given int32 and assigns it to the Metamorph field.
func (o *RunStats) SetMetamorph(v int32) {
	o.Metamorph = &v
}

// GetComputeUnits returns the ComputeUnits field value
func (o *RunStats) GetComputeUnits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ComputeUnits
}

// GetComputeUnitsOk returns a tuple with the ComputeUnits field value
// and a boolean to check if the value has been set.
func (o *RunStats) GetComputeUnitsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputeUnits, true
}

// SetComputeUnits sets field value
func (o *RunStats) SetComputeUnits(v float32) {
	o.ComputeUnits = v
}

func (o RunStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InputBodyLen.IsSet() {
		toSerialize["inputBodyLen"] = o.InputBodyLen.Get()
	}
	if !IsNil(o.MigrationCount) {
		toSerialize["migrationCount"] = o.MigrationCount
	}
	if !IsNil(o.RebootCount) {
		toSerialize["rebootCount"] = o.RebootCount
	}
	toSerialize["restartCount"] = o.RestartCount
	toSerialize["resurrectCount"] = o.ResurrectCount
	if !IsNil(o.MemAvgBytes) {
		toSerialize["memAvgBytes"] = o.MemAvgBytes
	}
	if !IsNil(o.MemMaxBytes) {
		toSerialize["memMaxBytes"] = o.MemMaxBytes
	}
	if !IsNil(o.MemCurrentBytes) {
		toSerialize["memCurrentBytes"] = o.MemCurrentBytes
	}
	if !IsNil(o.CpuAvgUsage) {
		toSerialize["cpuAvgUsage"] = o.CpuAvgUsage
	}
	if !IsNil(o.CpuMaxUsage) {
		toSerialize["cpuMaxUsage"] = o.CpuMaxUsage
	}
	if !IsNil(o.CpuCurrentUsage) {
		toSerialize["cpuCurrentUsage"] = o.CpuCurrentUsage
	}
	if !IsNil(o.NetRxBytes) {
		toSerialize["netRxBytes"] = o.NetRxBytes
	}
	if !IsNil(o.NetTxBytes) {
		toSerialize["netTxBytes"] = o.NetTxBytes
	}
	if !IsNil(o.DurationMillis) {
		toSerialize["durationMillis"] = o.DurationMillis
	}
	if !IsNil(o.RunTimeSecs) {
		toSerialize["runTimeSecs"] = o.RunTimeSecs
	}
	if !IsNil(o.Metamorph) {
		toSerialize["metamorph"] = o.Metamorph
	}
	toSerialize["computeUnits"] = o.ComputeUnits
	return toSerialize, nil
}

func (o *RunStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"restartCount",
		"resurrectCount",
		"computeUnits",
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

	varRunStats := _RunStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunStats)

	if err != nil {
		return err
	}

	*o = RunStats(varRunStats)

	return err
}

type NullableRunStats struct {
	value *RunStats
	isSet bool
}

func (v NullableRunStats) Get() *RunStats {
	return v.value
}

func (v *NullableRunStats) Set(val *RunStats) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStats) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStats(val *RunStats) *NullableRunStats {
	return &NullableRunStats{value: val, isSet: true}
}

func (v NullableRunStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


