/*
Apify API

 The Apify API (version 2) provides programmatic access to the [Apify platform](https://docs.apify.com). The API is organized around [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) HTTP endpoints.  You can download the complete OpenAPI schema of Apify API in the [YAML](http://docs.apify.com/api/openapi.yaml) or [JSON](http://docs.apify.com/api/openapi.json) formats. The source code is also available on [GitHub](https://github.com/apify/apify-docs/tree/master/apify-api/openapi).  All requests and responses (including errors) are encoded in [JSON](http://www.json.org/) format with UTF-8 encoding, with a few exceptions that are explicitly described in the reference.  - To access the API using [Node.js](https://nodejs.org/en/), we recommend the [`apify-client`](https://docs.apify.com/api/client/js) [NPM package](https://www.npmjs.com/package/apify-client). - To access the API using [Python](https://www.python.org/), we recommend the [`apify-client`](https://docs.apify.com/api/client/python) [PyPI package](https://pypi.org/project/apify-client/).  The clients' functions correspond to the API endpoints and have the same parameters. This simplifies development of apps that depend on the Apify platform.  :::note Important Request Details  - `Content-Type` header: For requests with a JSON body, you must include the `Content-Type: application/json` header.  - Method override: You can override the HTTP method using the `method` query parameter. This is useful for clients that can only send `GET` requests. For example, to call a `POST` endpoint, append `?method=POST` to the URL of your `GET` request.  :::  ## Authentication <span id=\"/introduction/authentication\"></span>  **You can find your API token on the [Integrations](https://console.apify.com/settings/integrations) page in the Apify Console.**  To use your token in a request, either:  - Add the token to your request's `Authorization` header as `Bearer <token>`. E.g., `Authorization: Bearer xxxxxxx`. [More info](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization). (Recommended). - Add it as the `token` parameter to your request URL. (Less secure).  Using your token in the request header is more secure than using it as a URL parameter because URLs are often stored in browser history and server logs. This creates a chance for someone unauthorized to access your API token.  **Never share your API token or password with untrusted parties!**  For more information, see our [integrations](https://docs.apify.com/platform/integrations) documentation.  ### Agentic payments  AI agents can authenticate and pay for Actor runs without an Apify account using agentic payments. Instead of an API token, the request carries a payment credential that both authorizes and pays for the call. Apify supports the [x402 protocol](https://docs.apify.com/platform/integrations/x402) (`PAYMENT-SIGNATURE` header) and [Skyfire](https://docs.apify.com/platform/integrations/skyfire) (`skyfire-pay-id` header).  ## Basic usage <span id=\"/introduction/basic-usage\"></span>  To run an Actor, send a POST request to the [Run Actor](#/reference/actors/run-collection/run-actor) endpoint using either the Actor ID code (e.g. `vKg4IjxZbEYTYeW8T`) or its name (e.g. `janedoe~my-actor`):  `https://api.apify.com/v2/actors/[actor_id]/runs`  If the Actor is not runnable anonymously, you will receive a 401 or 403 [response code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status). This means you need to add your [secret API token](https://console.apify.com/account#/integrations) to the request's `Authorization` header ([recommended](#/introduction/authentication)) or as a URL query parameter `?token=[your_token]` (less secure).  Optionally, you can include the query parameters described in the [Run Actor](#/reference/actors/run-collection/run-actor) section to customize your run.  If you're using Node.js, the best way to run an Actor is using the `Apify.call()` method from the [Apify SDK](https://sdk.apify.com/docs/api/apify#apifycallactid-input-options). It runs the Actor using the account you are currently logged into (determined by the [secret API token](https://console.apify.com/account#/integrations)). The result is an [Actor run object](https://sdk.apify.com/docs/typedefs/actor-run) and its output (if any).  A typical workflow is as follows:  1. Run an Actor or task using the [Run Actor](#/reference/actors/run-collection/run-actor) or [Run task](#/reference/actor-tasks/run-collection/run-task) API endpoints. 2. Monitor the Actor run by periodically polling its progress using the [Get run](#/reference/actor-runs/run-object-and-its-storages/get-run) API endpoint. 3. Fetch the results from the [Get items](#/reference/datasets/item-collection/get-items) API endpoint using the `defaultDatasetId`, which you receive in the Run request response. Additional data may be stored in a key-value store. You can fetch them from the [Get record](#/reference/key-value-stores/record/get-record) API endpoint using the `defaultKeyValueStoreId` and the store's `key`.  **Note**: Instead of periodic polling, you can also run your [Actor](#/reference/actors/run-actor-synchronously) or [task](#/reference/actor-tasks/runs-collection/run-task-synchronously) synchronously. This will ensure that the request waits for 300 seconds (5 minutes) for the run to finish and returns its output. If the run takes longer, the request will time out and throw an error.  ## Legacy `/v2/acts/` URL prefix <span id=\"/introduction/legacy-acts-prefix\"></span>  The `/v2/acts/` prefix is deprecated but still fully functional, and  such endpoint routes to the same handler as its `/v2/actors/...` counterpart.  New integrations should use the canonical /v2/actors/ prefix,  but existing clients keep working without changes.  ## Response structure <span id=\"/introduction/response-structure\"></span>  Most API endpoints return a JSON object with the `data` property:  ``` {     \"data\": {         ...     } } ```  However, there are a few explicitly described exceptions, such as [Get dataset items](#/reference/datasets/item-collection/get-items) or Key-value store [Get record](#/reference/key-value-stores/record/get-record) API endpoints, which return data in other formats. In case of an error, the response has the HTTP status code in the range of 4xx or 5xx and the `data` property is replaced with `error`. For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  See [Errors](#/introduction/errors) for more details.  ## Pagination <span id=\"/introduction/pagination\"></span>  All API endpoints that return a list of records (e.g. [Get list of Actors](#/reference/actors/actor-collection/get-list-of-actors)) enforce pagination in order to limit the size of their responses.  Most of these API endpoints are paginated using the `offset` and `limit` query parameters. The only exception is [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys), which is paginated using the `exclusiveStartKey` query parameter.  **IMPORTANT**: Each API endpoint that supports pagination enforces a certain maximum value for the `limit` parameter, in order to reduce the load on Apify servers. The maximum limit could change in future so you should never rely on a specific value and check the responses of these API endpoints.  ### Using offset <span id=\"/introduction/pagination/using-offset\"></span>  Most API endpoints that return a list of records enable pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number of items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>Skips a number of items from the beginning of the list, e.g. <code>offset=100</code>.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td>     By default, items are sorted in the order in which they were created or added to the list.     This feature is useful when fetching all the items, because it ensures that items     created after the client started the pagination will not be skipped.     If you specify the <code>desc=1</code> parameter, the items will be returned in the reverse order,     i.e. from the newest to the oldest items.     </td>   </tr> </table>  The response of these API endpoints is always a JSON object with the following structure:  ``` {     \"data\": {         \"total\": 2560,         \"offset\": 250,         \"limit\": 1000,         \"count\": 1000,         \"desc\": false,         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>total</code></td>     <td>The total number of items available in the list.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>The number of items that were skipped at the start.     This is equal to the <code>offset</code> query parameter if it was provided, otherwise it is <code>0</code>.</td>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular API endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>count</code></td>     <td>The actual number of items returned in the HTTP response.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td><code>true</code> if data were requested in descending order and <code>false</code> otherwise.</td>   </tr>   <tr>     <td><code>items</code></td>     <td>An array of requested items.</td>   </tr> </table>  ### Using key <span id=\"/introduction/pagination/using-key\"></span>  The records in the [key-value store](https://docs.apify.com/platform/storage/key-value-store) are not ordered based on numerical indexes, but rather by their keys in the UTF-8 binary order. Therefore the [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys) API endpoint only supports pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>Skips all records with keys up to the given key including the given key,     in the UTF-8 binary order.</td>   </tr> </table>  The response of the API endpoint is always a JSON object with following structure:  ``` {     \"data\": {         \"limit\": 1000,         \"isTruncated\": true,         \"exclusiveStartKey\": \"my-key\",         \"nextExclusiveStartKey\": \"some-other-key\",         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>isTruncated</code></td>     <td><code>true</code> if there are more items left to be queried. Otherwise <code>false</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>The last key that was skipped at the start. Is `null` for the first page.</td>   </tr>   <tr>     <td><code>nextExclusiveStartKey</code></td>     <td>The value for the <code>exclusiveStartKey</code> parameter to query the next page of items.</td>   </tr> </table>  ## Errors <span id=\"/introduction/errors\"></span>  The Apify API uses common HTTP status codes: `2xx` range for success, `4xx` range for errors caused by the caller (invalid requests) and `5xx` range for server errors (these are rare). Each error response contains a JSON object defining the `error` property, which is an object with the `type` and `message` properties that contain the error code and a human-readable error description, respectively.  For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  Here is the table of the most common errors that can occur for many API endpoints:  <table>   <tr>     <th>status</th>     <th>type</th>     <th>message</th>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-request</code></td>     <td>POST data must be a JSON object</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-value</code></td>     <td>Invalid value provided: Comments required</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-record-key</code></td>     <td>Record key contains invalid character</td>   </tr>   <tr>     <td><code>401</code></td>     <td><code>token-not-provided</code></td>     <td>Authentication token was not provided</td>   </tr>   <tr>     <td><code>404</code></td>     <td><code>record-not-found</code></td>     <td>Store was not found</td>   </tr>   <tr>     <td><code>429</code></td>     <td><code>rate-limit-exceeded</code></td>     <td>You have exceeded the rate limit of ... requests per second</td>   </tr>   <tr>     <td><code>405</code></td>     <td><code>method-not-allowed</code></td>     <td>This API endpoint can only be accessed using the following HTTP methods: OPTIONS, POST</td>   </tr> </table>  ## Rate limiting <span id=\"/introduction/rate-limiting\"></span>  All API endpoints limit the rate of requests in order to prevent overloading of Apify servers by misbehaving clients.  There are two kinds of rate limits - a global rate limit and a per-resource rate limit.  ### Global rate limit <span id=\"/introduction/rate-limiting/global-rate-limit\"></span>  The global rate limit is set to _250 000 requests per minute_. For [authenticated](#/introduction/authentication) requests, it is counted per user, and for unauthenticated requests, it is counted per IP address.  ### Per-resource rate limit <span id=\"/introduction/rate-limiting/per-resource-rate-limit\"></span>  The default per-resource rate limit is _60 requests per second per resource_, which in this context means a single Actor, a single Actor run, a single dataset, single key-value store etc. The default rate limit is applied to every API endpoint except a few select ones, which have higher rate limits. Each API endpoint returns its rate limit in `X-RateLimit-Limit` header.  These endpoints have a rate limit of _200 requests per second per resource_:  * CRUD ([get](#/reference/key-value-stores/record/get-record),   [put](#/reference/key-value-stores/record/put-record),   [delete](#/reference/key-value-stores/record/delete-record))   operations on key-value store records  These endpoints have a rate limit of _400 requests per second per resource_: * [Run Actor](#/reference/actors/run-collection/run-actor) * [Run Actor task asynchronously](#/reference/actor-tasks/runs-collection/run-task-asynchronously) * [Run Actor task synchronously](#/reference/actor-tasks/runs-collection/run-task-synchronously) * [Metamorph Actor run](#/reference/actors/metamorph-run/metamorph-run) * [Push items](#/reference/datasets/item-collection/put-items) to dataset * CRUD   ([add](#/reference/request-queues/request-collection/add-request),   [get](#/reference/request-queues/request-collection/get-request),   [update](#/reference/request-queues/request-collection/update-request),   [delete](#/reference/request-queues/request-collection/delete-request))   operations on requests in request queues  ### Rate limit exceeded errors <span id=\"/introduction/rate-limiting/rate-limit-exceeded-errors\"></span>  If the client is sending too many requests, the API endpoints respond with the HTTP status code `429 Too Many Requests` and the following body:  ``` {     \"error\": {         \"type\": \"rate-limit-exceeded\",         \"message\": \"You have exceeded the rate limit of ... requests per second\"     } } ```  ### Retrying rate-limited requests with exponential backoff <span id=\"/introduction/rate-limiting/retrying-rate-limited-requests-with-exponential-backoff\"></span>  If the client receives the rate limit error, it should wait a certain period of time and then retry the request. If the error happens again, the client should double the wait period and retry the request, and so on. This algorithm is known as _exponential backoff_ and it can be described using the following pseudo-code:  1. Define a variable `DELAY=500` 2. Send the HTTP request to the API endpoint 3. If the response has status code not equal to `429` then you are done. Otherwise:    * Wait for a period of time chosen randomly from the interval `DELAY` to `2*DELAY` milliseconds    * Double the future wait period by setting `DELAY = 2*DELAY`    * Continue with step 2  If all requests sent by the client implement the above steps, the client will automatically use the maximum available bandwidth for its requests.  Note that the Apify API clients [for JavaScript](https://docs.apify.com/api/client/js) and [for Python](https://docs.apify.com/api/client/python) use the exponential backoff algorithm transparently, so that you do not need to worry about it.  ## Referring to resources <span id=\"/introduction/referring-to-resources\"></span>  There are three main ways to refer to a resource you're accessing via API.  - the resource ID (e.g. `iKkPcIgVvwmztduf8`) - `username~resourcename` - when using this access method, you will need to use your API token, and access will only work if you have the correct permissions. - `~resourcename` - for this, you need to use an API token, and the `resourcename` refers to a resource in the API token owner's account. 

API version: v2-2026-06-16T064758Z
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apifyclient

import (
	"encoding/json"
)

// checks if the ActorStandby type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActorStandby{}

// ActorStandby struct for ActorStandby
type ActorStandby struct {
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
	DesiredRequestsPerActorRun NullableInt32 `json:"desiredRequestsPerActorRun,omitempty"`
	MaxRequestsPerActorRun NullableInt32 `json:"maxRequestsPerActorRun,omitempty"`
	IdleTimeoutSecs NullableInt32 `json:"idleTimeoutSecs,omitempty"`
	Build NullableString `json:"build,omitempty"`
	MemoryMbytes NullableInt32 `json:"memoryMbytes,omitempty"`
	DisableStandbyFieldsOverride NullableBool `json:"disableStandbyFieldsOverride,omitempty"`
	ShouldPassActorInput NullableBool `json:"shouldPassActorInput,omitempty"`
}

// NewActorStandby instantiates a new ActorStandby object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActorStandby() *ActorStandby {
	this := ActorStandby{}
	return &this
}

// NewActorStandbyWithDefaults instantiates a new ActorStandby object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorStandbyWithDefaults() *ActorStandby {
	this := ActorStandby{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ActorStandby) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *ActorStandby) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *ActorStandby) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *ActorStandby) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetDesiredRequestsPerActorRun returns the DesiredRequestsPerActorRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetDesiredRequestsPerActorRun() int32 {
	if o == nil || IsNil(o.DesiredRequestsPerActorRun.Get()) {
		var ret int32
		return ret
	}
	return *o.DesiredRequestsPerActorRun.Get()
}

// GetDesiredRequestsPerActorRunOk returns a tuple with the DesiredRequestsPerActorRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetDesiredRequestsPerActorRunOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DesiredRequestsPerActorRun.Get(), o.DesiredRequestsPerActorRun.IsSet()
}

// HasDesiredRequestsPerActorRun returns a boolean if a field has been set.
func (o *ActorStandby) HasDesiredRequestsPerActorRun() bool {
	if o != nil && o.DesiredRequestsPerActorRun.IsSet() {
		return true
	}

	return false
}

// SetDesiredRequestsPerActorRun gets a reference to the given NullableInt32 and assigns it to the DesiredRequestsPerActorRun field.
func (o *ActorStandby) SetDesiredRequestsPerActorRun(v int32) {
	o.DesiredRequestsPerActorRun.Set(&v)
}
// SetDesiredRequestsPerActorRunNil sets the value for DesiredRequestsPerActorRun to be an explicit nil
func (o *ActorStandby) SetDesiredRequestsPerActorRunNil() {
	o.DesiredRequestsPerActorRun.Set(nil)
}

// UnsetDesiredRequestsPerActorRun ensures that no value is present for DesiredRequestsPerActorRun, not even an explicit nil
func (o *ActorStandby) UnsetDesiredRequestsPerActorRun() {
	o.DesiredRequestsPerActorRun.Unset()
}

// GetMaxRequestsPerActorRun returns the MaxRequestsPerActorRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetMaxRequestsPerActorRun() int32 {
	if o == nil || IsNil(o.MaxRequestsPerActorRun.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRequestsPerActorRun.Get()
}

// GetMaxRequestsPerActorRunOk returns a tuple with the MaxRequestsPerActorRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetMaxRequestsPerActorRunOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRequestsPerActorRun.Get(), o.MaxRequestsPerActorRun.IsSet()
}

// HasMaxRequestsPerActorRun returns a boolean if a field has been set.
func (o *ActorStandby) HasMaxRequestsPerActorRun() bool {
	if o != nil && o.MaxRequestsPerActorRun.IsSet() {
		return true
	}

	return false
}

// SetMaxRequestsPerActorRun gets a reference to the given NullableInt32 and assigns it to the MaxRequestsPerActorRun field.
func (o *ActorStandby) SetMaxRequestsPerActorRun(v int32) {
	o.MaxRequestsPerActorRun.Set(&v)
}
// SetMaxRequestsPerActorRunNil sets the value for MaxRequestsPerActorRun to be an explicit nil
func (o *ActorStandby) SetMaxRequestsPerActorRunNil() {
	o.MaxRequestsPerActorRun.Set(nil)
}

// UnsetMaxRequestsPerActorRun ensures that no value is present for MaxRequestsPerActorRun, not even an explicit nil
func (o *ActorStandby) UnsetMaxRequestsPerActorRun() {
	o.MaxRequestsPerActorRun.Unset()
}

// GetIdleTimeoutSecs returns the IdleTimeoutSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetIdleTimeoutSecs() int32 {
	if o == nil || IsNil(o.IdleTimeoutSecs.Get()) {
		var ret int32
		return ret
	}
	return *o.IdleTimeoutSecs.Get()
}

// GetIdleTimeoutSecsOk returns a tuple with the IdleTimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetIdleTimeoutSecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdleTimeoutSecs.Get(), o.IdleTimeoutSecs.IsSet()
}

// HasIdleTimeoutSecs returns a boolean if a field has been set.
func (o *ActorStandby) HasIdleTimeoutSecs() bool {
	if o != nil && o.IdleTimeoutSecs.IsSet() {
		return true
	}

	return false
}

// SetIdleTimeoutSecs gets a reference to the given NullableInt32 and assigns it to the IdleTimeoutSecs field.
func (o *ActorStandby) SetIdleTimeoutSecs(v int32) {
	o.IdleTimeoutSecs.Set(&v)
}
// SetIdleTimeoutSecsNil sets the value for IdleTimeoutSecs to be an explicit nil
func (o *ActorStandby) SetIdleTimeoutSecsNil() {
	o.IdleTimeoutSecs.Set(nil)
}

// UnsetIdleTimeoutSecs ensures that no value is present for IdleTimeoutSecs, not even an explicit nil
func (o *ActorStandby) UnsetIdleTimeoutSecs() {
	o.IdleTimeoutSecs.Unset()
}

// GetBuild returns the Build field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetBuild() string {
	if o == nil || IsNil(o.Build.Get()) {
		var ret string
		return ret
	}
	return *o.Build.Get()
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Build.Get(), o.Build.IsSet()
}

// HasBuild returns a boolean if a field has been set.
func (o *ActorStandby) HasBuild() bool {
	if o != nil && o.Build.IsSet() {
		return true
	}

	return false
}

// SetBuild gets a reference to the given NullableString and assigns it to the Build field.
func (o *ActorStandby) SetBuild(v string) {
	o.Build.Set(&v)
}
// SetBuildNil sets the value for Build to be an explicit nil
func (o *ActorStandby) SetBuildNil() {
	o.Build.Set(nil)
}

// UnsetBuild ensures that no value is present for Build, not even an explicit nil
func (o *ActorStandby) UnsetBuild() {
	o.Build.Unset()
}

// GetMemoryMbytes returns the MemoryMbytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetMemoryMbytes() int32 {
	if o == nil || IsNil(o.MemoryMbytes.Get()) {
		var ret int32
		return ret
	}
	return *o.MemoryMbytes.Get()
}

// GetMemoryMbytesOk returns a tuple with the MemoryMbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetMemoryMbytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryMbytes.Get(), o.MemoryMbytes.IsSet()
}

// HasMemoryMbytes returns a boolean if a field has been set.
func (o *ActorStandby) HasMemoryMbytes() bool {
	if o != nil && o.MemoryMbytes.IsSet() {
		return true
	}

	return false
}

// SetMemoryMbytes gets a reference to the given NullableInt32 and assigns it to the MemoryMbytes field.
func (o *ActorStandby) SetMemoryMbytes(v int32) {
	o.MemoryMbytes.Set(&v)
}
// SetMemoryMbytesNil sets the value for MemoryMbytes to be an explicit nil
func (o *ActorStandby) SetMemoryMbytesNil() {
	o.MemoryMbytes.Set(nil)
}

// UnsetMemoryMbytes ensures that no value is present for MemoryMbytes, not even an explicit nil
func (o *ActorStandby) UnsetMemoryMbytes() {
	o.MemoryMbytes.Unset()
}

// GetDisableStandbyFieldsOverride returns the DisableStandbyFieldsOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetDisableStandbyFieldsOverride() bool {
	if o == nil || IsNil(o.DisableStandbyFieldsOverride.Get()) {
		var ret bool
		return ret
	}
	return *o.DisableStandbyFieldsOverride.Get()
}

// GetDisableStandbyFieldsOverrideOk returns a tuple with the DisableStandbyFieldsOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetDisableStandbyFieldsOverrideOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisableStandbyFieldsOverride.Get(), o.DisableStandbyFieldsOverride.IsSet()
}

// HasDisableStandbyFieldsOverride returns a boolean if a field has been set.
func (o *ActorStandby) HasDisableStandbyFieldsOverride() bool {
	if o != nil && o.DisableStandbyFieldsOverride.IsSet() {
		return true
	}

	return false
}

// SetDisableStandbyFieldsOverride gets a reference to the given NullableBool and assigns it to the DisableStandbyFieldsOverride field.
func (o *ActorStandby) SetDisableStandbyFieldsOverride(v bool) {
	o.DisableStandbyFieldsOverride.Set(&v)
}
// SetDisableStandbyFieldsOverrideNil sets the value for DisableStandbyFieldsOverride to be an explicit nil
func (o *ActorStandby) SetDisableStandbyFieldsOverrideNil() {
	o.DisableStandbyFieldsOverride.Set(nil)
}

// UnsetDisableStandbyFieldsOverride ensures that no value is present for DisableStandbyFieldsOverride, not even an explicit nil
func (o *ActorStandby) UnsetDisableStandbyFieldsOverride() {
	o.DisableStandbyFieldsOverride.Unset()
}

// GetShouldPassActorInput returns the ShouldPassActorInput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActorStandby) GetShouldPassActorInput() bool {
	if o == nil || IsNil(o.ShouldPassActorInput.Get()) {
		var ret bool
		return ret
	}
	return *o.ShouldPassActorInput.Get()
}

// GetShouldPassActorInputOk returns a tuple with the ShouldPassActorInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActorStandby) GetShouldPassActorInputOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldPassActorInput.Get(), o.ShouldPassActorInput.IsSet()
}

// HasShouldPassActorInput returns a boolean if a field has been set.
func (o *ActorStandby) HasShouldPassActorInput() bool {
	if o != nil && o.ShouldPassActorInput.IsSet() {
		return true
	}

	return false
}

// SetShouldPassActorInput gets a reference to the given NullableBool and assigns it to the ShouldPassActorInput field.
func (o *ActorStandby) SetShouldPassActorInput(v bool) {
	o.ShouldPassActorInput.Set(&v)
}
// SetShouldPassActorInputNil sets the value for ShouldPassActorInput to be an explicit nil
func (o *ActorStandby) SetShouldPassActorInputNil() {
	o.ShouldPassActorInput.Set(nil)
}

// UnsetShouldPassActorInput ensures that no value is present for ShouldPassActorInput, not even an explicit nil
func (o *ActorStandby) UnsetShouldPassActorInput() {
	o.ShouldPassActorInput.Unset()
}

func (o ActorStandby) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActorStandby) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	if o.DesiredRequestsPerActorRun.IsSet() {
		toSerialize["desiredRequestsPerActorRun"] = o.DesiredRequestsPerActorRun.Get()
	}
	if o.MaxRequestsPerActorRun.IsSet() {
		toSerialize["maxRequestsPerActorRun"] = o.MaxRequestsPerActorRun.Get()
	}
	if o.IdleTimeoutSecs.IsSet() {
		toSerialize["idleTimeoutSecs"] = o.IdleTimeoutSecs.Get()
	}
	if o.Build.IsSet() {
		toSerialize["build"] = o.Build.Get()
	}
	if o.MemoryMbytes.IsSet() {
		toSerialize["memoryMbytes"] = o.MemoryMbytes.Get()
	}
	if o.DisableStandbyFieldsOverride.IsSet() {
		toSerialize["disableStandbyFieldsOverride"] = o.DisableStandbyFieldsOverride.Get()
	}
	if o.ShouldPassActorInput.IsSet() {
		toSerialize["shouldPassActorInput"] = o.ShouldPassActorInput.Get()
	}
	return toSerialize, nil
}

type NullableActorStandby struct {
	value *ActorStandby
	isSet bool
}

func (v NullableActorStandby) Get() *ActorStandby {
	return v.value
}

func (v *NullableActorStandby) Set(val *ActorStandby) {
	v.value = val
	v.isSet = true
}

func (v NullableActorStandby) IsSet() bool {
	return v.isSet
}

func (v *NullableActorStandby) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActorStandby(val *ActorStandby) *NullableActorStandby {
	return &NullableActorStandby{value: val, isSet: true}
}

func (v NullableActorStandby) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActorStandby) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


