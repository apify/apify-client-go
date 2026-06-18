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

// checks if the Run type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Run{}

// Run Represents an Actor run and its associated data.
type Run struct {
	// Unique identifier of the Actor run.
	Id string `json:"id"`
	// ID of the Actor that was run.
	ActId string `json:"actId"`
	// ID of the user who started the run.
	UserId string `json:"userId"`
	// ID of the Actor task, if the run was started from a task.
	ActorTaskId NullableString `json:"actorTaskId,omitempty"`
	// Time when the Actor run started.
	StartedAt time.Time `json:"startedAt"`
	// Time when the Actor run finished.
	FinishedAt NullableTime `json:"finishedAt,omitempty"`
	// Current status of the Actor run.
	Status ActorJobStatus `json:"status"`
	// Detailed message about the run status.
	StatusMessage NullableString `json:"statusMessage,omitempty"`
	// Whether the status message is terminal (final).
	IsStatusMessageTerminal NullableBool `json:"isStatusMessageTerminal,omitempty"`
	// Metadata about the Actor run.
	Meta RunMeta `json:"meta"`
	// Pricing information for the Actor.
	PricingInfo *ActorRunPricingInfo `json:"pricingInfo,omitempty"`
	// Statistics of the Actor run.
	Stats RunStats `json:"stats"`
	// A map of charged event types to their counts. The keys are event type identifiers defined by the Actor's pricing model (pay-per-event), and the values are the number of times each event was charged during this run.
	ChargedEventCounts map[string]int32 `json:"chargedEventCounts,omitempty"`
	// Configuration options for the Actor run.
	Options RunOptions `json:"options"`
	// ID of the Actor build used for this run.
	BuildId string `json:"buildId"`
	// Exit code of the Actor run process.
	ExitCode NullableInt32 `json:"exitCode,omitempty"`
	// General access level for the Actor run.
	GeneralAccess GeneralAccess `json:"generalAccess"`
	// ID of the default key-value store for this run.
	DefaultKeyValueStoreId string `json:"defaultKeyValueStoreId"`
	// ID of the default dataset for this run.
	DefaultDatasetId string `json:"defaultDatasetId"`
	// ID of the default request queue for this run.
	DefaultRequestQueueId string `json:"defaultRequestQueueId"`
	StorageIds *RunStorageIds `json:"storageIds,omitempty"`
	// Build number of the Actor build used for this run.
	BuildNumber NullableString `json:"buildNumber,omitempty" validate:"regexp=^([0-9]|[1-9][0-9])\\\\.([0-9]|[1-9][0-9])(\\\\.[1-9][0-9]{0,4})$"`
	// URL of the container running the Actor.
	ContainerUrl *string `json:"containerUrl,omitempty"`
	// Whether the container's HTTP server is ready to accept requests.
	IsContainerServerReady NullableBool `json:"isContainerServerReady,omitempty"`
	// Name of the git branch used for the Actor build.
	GitBranchName NullableString `json:"gitBranchName,omitempty"`
	// Resource usage statistics for the run.
	Usage NullableRunUsage `json:"usage,omitempty"`
	// Total cost in USD for this run. Represents what you actually pay. For run owners: includes platform usage (compute units) and/or event costs depending on the Actor's pricing model. For run non-owners: only available for Pay-Per-Event Actors (event costs only). Requires authentication token to access.
	UsageTotalUsd NullableFloat32 `json:"usageTotalUsd,omitempty"`
	// Platform usage costs breakdown in USD. Only present if you own the run AND are paying for platform usage (Pay-Per-Usage, Rental, or Pay-Per-Event with usage costs like standby Actors). Not available for standard Pay-Per-Event Actors. Requires authentication token to access.
	UsageUsd NullableRunUsageUsd `json:"usageUsd,omitempty"`
	// List of metamorph events that occurred during the run.
	Metamorphs []Metamorph `json:"metamorphs,omitempty"`
	// Indicates which party covers platform usage costs for this run.
	PlatformUsageBillingModel *string `json:"platformUsageBillingModel,omitempty"`
}

type _Run Run

// NewRun instantiates a new Run object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRun(id string, actId string, userId string, startedAt time.Time, status ActorJobStatus, meta RunMeta, stats RunStats, options RunOptions, buildId string, generalAccess GeneralAccess, defaultKeyValueStoreId string, defaultDatasetId string, defaultRequestQueueId string) *Run {
	this := Run{}
	this.Id = id
	this.ActId = actId
	this.UserId = userId
	this.StartedAt = startedAt
	this.Status = status
	this.Meta = meta
	this.Stats = stats
	this.Options = options
	this.BuildId = buildId
	this.GeneralAccess = generalAccess
	this.DefaultKeyValueStoreId = defaultKeyValueStoreId
	this.DefaultDatasetId = defaultDatasetId
	this.DefaultRequestQueueId = defaultRequestQueueId
	return &this
}

// NewRunWithDefaults instantiates a new Run object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunWithDefaults() *Run {
	this := Run{}
	return &this
}

// GetId returns the Id field value
func (o *Run) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Run) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Run) SetId(v string) {
	o.Id = v
}

// GetActId returns the ActId field value
func (o *Run) GetActId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActId
}

// GetActIdOk returns a tuple with the ActId field value
// and a boolean to check if the value has been set.
func (o *Run) GetActIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActId, true
}

// SetActId sets field value
func (o *Run) SetActId(v string) {
	o.ActId = v
}

// GetUserId returns the UserId field value
func (o *Run) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Run) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Run) SetUserId(v string) {
	o.UserId = v
}

// GetActorTaskId returns the ActorTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetActorTaskId() string {
	if o == nil || IsNil(o.ActorTaskId.Get()) {
		var ret string
		return ret
	}
	return *o.ActorTaskId.Get()
}

// GetActorTaskIdOk returns a tuple with the ActorTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetActorTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorTaskId.Get(), o.ActorTaskId.IsSet()
}

// HasActorTaskId returns a boolean if a field has been set.
func (o *Run) HasActorTaskId() bool {
	if o != nil && o.ActorTaskId.IsSet() {
		return true
	}

	return false
}

// SetActorTaskId gets a reference to the given NullableString and assigns it to the ActorTaskId field.
func (o *Run) SetActorTaskId(v string) {
	o.ActorTaskId.Set(&v)
}
// SetActorTaskIdNil sets the value for ActorTaskId to be an explicit nil
func (o *Run) SetActorTaskIdNil() {
	o.ActorTaskId.Set(nil)
}

// UnsetActorTaskId ensures that no value is present for ActorTaskId, not even an explicit nil
func (o *Run) UnsetActorTaskId() {
	o.ActorTaskId.Unset()
}

// GetStartedAt returns the StartedAt field value
func (o *Run) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *Run) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *Run) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *Run) HasFinishedAt() bool {
	if o != nil && o.FinishedAt.IsSet() {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given NullableTime and assigns it to the FinishedAt field.
func (o *Run) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}
// SetFinishedAtNil sets the value for FinishedAt to be an explicit nil
func (o *Run) SetFinishedAtNil() {
	o.FinishedAt.Set(nil)
}

// UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
func (o *Run) UnsetFinishedAt() {
	o.FinishedAt.Unset()
}

// GetStatus returns the Status field value
func (o *Run) GetStatus() ActorJobStatus {
	if o == nil {
		var ret ActorJobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Run) GetStatusOk() (*ActorJobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Run) SetStatus(v ActorJobStatus) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage.Get()) {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Run) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *Run) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *Run) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *Run) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetIsStatusMessageTerminal returns the IsStatusMessageTerminal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetIsStatusMessageTerminal() bool {
	if o == nil || IsNil(o.IsStatusMessageTerminal.Get()) {
		var ret bool
		return ret
	}
	return *o.IsStatusMessageTerminal.Get()
}

// GetIsStatusMessageTerminalOk returns a tuple with the IsStatusMessageTerminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetIsStatusMessageTerminalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsStatusMessageTerminal.Get(), o.IsStatusMessageTerminal.IsSet()
}

// HasIsStatusMessageTerminal returns a boolean if a field has been set.
func (o *Run) HasIsStatusMessageTerminal() bool {
	if o != nil && o.IsStatusMessageTerminal.IsSet() {
		return true
	}

	return false
}

// SetIsStatusMessageTerminal gets a reference to the given NullableBool and assigns it to the IsStatusMessageTerminal field.
func (o *Run) SetIsStatusMessageTerminal(v bool) {
	o.IsStatusMessageTerminal.Set(&v)
}
// SetIsStatusMessageTerminalNil sets the value for IsStatusMessageTerminal to be an explicit nil
func (o *Run) SetIsStatusMessageTerminalNil() {
	o.IsStatusMessageTerminal.Set(nil)
}

// UnsetIsStatusMessageTerminal ensures that no value is present for IsStatusMessageTerminal, not even an explicit nil
func (o *Run) UnsetIsStatusMessageTerminal() {
	o.IsStatusMessageTerminal.Unset()
}

// GetMeta returns the Meta field value
func (o *Run) GetMeta() RunMeta {
	if o == nil {
		var ret RunMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *Run) GetMetaOk() (*RunMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *Run) SetMeta(v RunMeta) {
	o.Meta = v
}

// GetPricingInfo returns the PricingInfo field value if set, zero value otherwise.
func (o *Run) GetPricingInfo() ActorRunPricingInfo {
	if o == nil || IsNil(o.PricingInfo) {
		var ret ActorRunPricingInfo
		return ret
	}
	return *o.PricingInfo
}

// GetPricingInfoOk returns a tuple with the PricingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetPricingInfoOk() (*ActorRunPricingInfo, bool) {
	if o == nil || IsNil(o.PricingInfo) {
		return nil, false
	}
	return o.PricingInfo, true
}

// HasPricingInfo returns a boolean if a field has been set.
func (o *Run) HasPricingInfo() bool {
	if o != nil && !IsNil(o.PricingInfo) {
		return true
	}

	return false
}

// SetPricingInfo gets a reference to the given ActorRunPricingInfo and assigns it to the PricingInfo field.
func (o *Run) SetPricingInfo(v ActorRunPricingInfo) {
	o.PricingInfo = &v
}

// GetStats returns the Stats field value
func (o *Run) GetStats() RunStats {
	if o == nil {
		var ret RunStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *Run) GetStatsOk() (*RunStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *Run) SetStats(v RunStats) {
	o.Stats = v
}

// GetChargedEventCounts returns the ChargedEventCounts field value if set, zero value otherwise.
func (o *Run) GetChargedEventCounts() map[string]int32 {
	if o == nil || IsNil(o.ChargedEventCounts) {
		var ret map[string]int32
		return ret
	}
	return o.ChargedEventCounts
}

// GetChargedEventCountsOk returns a tuple with the ChargedEventCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetChargedEventCountsOk() (map[string]int32, bool) {
	if o == nil || IsNil(o.ChargedEventCounts) {
		return map[string]int32{}, false
	}
	return o.ChargedEventCounts, true
}

// HasChargedEventCounts returns a boolean if a field has been set.
func (o *Run) HasChargedEventCounts() bool {
	if o != nil && !IsNil(o.ChargedEventCounts) {
		return true
	}

	return false
}

// SetChargedEventCounts gets a reference to the given map[string]int32 and assigns it to the ChargedEventCounts field.
func (o *Run) SetChargedEventCounts(v map[string]int32) {
	o.ChargedEventCounts = v
}

// GetOptions returns the Options field value
func (o *Run) GetOptions() RunOptions {
	if o == nil {
		var ret RunOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Run) GetOptionsOk() (*RunOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *Run) SetOptions(v RunOptions) {
	o.Options = v
}

// GetBuildId returns the BuildId field value
func (o *Run) GetBuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *Run) GetBuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *Run) SetBuildId(v string) {
	o.BuildId = v
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetExitCode() int32 {
	if o == nil || IsNil(o.ExitCode.Get()) {
		var ret int32
		return ret
	}
	return *o.ExitCode.Get()
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetExitCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExitCode.Get(), o.ExitCode.IsSet()
}

// HasExitCode returns a boolean if a field has been set.
func (o *Run) HasExitCode() bool {
	if o != nil && o.ExitCode.IsSet() {
		return true
	}

	return false
}

// SetExitCode gets a reference to the given NullableInt32 and assigns it to the ExitCode field.
func (o *Run) SetExitCode(v int32) {
	o.ExitCode.Set(&v)
}
// SetExitCodeNil sets the value for ExitCode to be an explicit nil
func (o *Run) SetExitCodeNil() {
	o.ExitCode.Set(nil)
}

// UnsetExitCode ensures that no value is present for ExitCode, not even an explicit nil
func (o *Run) UnsetExitCode() {
	o.ExitCode.Unset()
}

// GetGeneralAccess returns the GeneralAccess field value
func (o *Run) GetGeneralAccess() GeneralAccess {
	if o == nil {
		var ret GeneralAccess
		return ret
	}

	return o.GeneralAccess
}

// GetGeneralAccessOk returns a tuple with the GeneralAccess field value
// and a boolean to check if the value has been set.
func (o *Run) GetGeneralAccessOk() (*GeneralAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneralAccess, true
}

// SetGeneralAccess sets field value
func (o *Run) SetGeneralAccess(v GeneralAccess) {
	o.GeneralAccess = v
}

// GetDefaultKeyValueStoreId returns the DefaultKeyValueStoreId field value
func (o *Run) GetDefaultKeyValueStoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultKeyValueStoreId
}

// GetDefaultKeyValueStoreIdOk returns a tuple with the DefaultKeyValueStoreId field value
// and a boolean to check if the value has been set.
func (o *Run) GetDefaultKeyValueStoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultKeyValueStoreId, true
}

// SetDefaultKeyValueStoreId sets field value
func (o *Run) SetDefaultKeyValueStoreId(v string) {
	o.DefaultKeyValueStoreId = v
}

// GetDefaultDatasetId returns the DefaultDatasetId field value
func (o *Run) GetDefaultDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultDatasetId
}

// GetDefaultDatasetIdOk returns a tuple with the DefaultDatasetId field value
// and a boolean to check if the value has been set.
func (o *Run) GetDefaultDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDatasetId, true
}

// SetDefaultDatasetId sets field value
func (o *Run) SetDefaultDatasetId(v string) {
	o.DefaultDatasetId = v
}

// GetDefaultRequestQueueId returns the DefaultRequestQueueId field value
func (o *Run) GetDefaultRequestQueueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultRequestQueueId
}

// GetDefaultRequestQueueIdOk returns a tuple with the DefaultRequestQueueId field value
// and a boolean to check if the value has been set.
func (o *Run) GetDefaultRequestQueueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRequestQueueId, true
}

// SetDefaultRequestQueueId sets field value
func (o *Run) SetDefaultRequestQueueId(v string) {
	o.DefaultRequestQueueId = v
}

// GetStorageIds returns the StorageIds field value if set, zero value otherwise.
func (o *Run) GetStorageIds() RunStorageIds {
	if o == nil || IsNil(o.StorageIds) {
		var ret RunStorageIds
		return ret
	}
	return *o.StorageIds
}

// GetStorageIdsOk returns a tuple with the StorageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetStorageIdsOk() (*RunStorageIds, bool) {
	if o == nil || IsNil(o.StorageIds) {
		return nil, false
	}
	return o.StorageIds, true
}

// HasStorageIds returns a boolean if a field has been set.
func (o *Run) HasStorageIds() bool {
	if o != nil && !IsNil(o.StorageIds) {
		return true
	}

	return false
}

// SetStorageIds gets a reference to the given RunStorageIds and assigns it to the StorageIds field.
func (o *Run) SetStorageIds(v RunStorageIds) {
	o.StorageIds = &v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetBuildNumber() string {
	if o == nil || IsNil(o.BuildNumber.Get()) {
		var ret string
		return ret
	}
	return *o.BuildNumber.Get()
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetBuildNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildNumber.Get(), o.BuildNumber.IsSet()
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *Run) HasBuildNumber() bool {
	if o != nil && o.BuildNumber.IsSet() {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given NullableString and assigns it to the BuildNumber field.
func (o *Run) SetBuildNumber(v string) {
	o.BuildNumber.Set(&v)
}
// SetBuildNumberNil sets the value for BuildNumber to be an explicit nil
func (o *Run) SetBuildNumberNil() {
	o.BuildNumber.Set(nil)
}

// UnsetBuildNumber ensures that no value is present for BuildNumber, not even an explicit nil
func (o *Run) UnsetBuildNumber() {
	o.BuildNumber.Unset()
}

// GetContainerUrl returns the ContainerUrl field value if set, zero value otherwise.
func (o *Run) GetContainerUrl() string {
	if o == nil || IsNil(o.ContainerUrl) {
		var ret string
		return ret
	}
	return *o.ContainerUrl
}

// GetContainerUrlOk returns a tuple with the ContainerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetContainerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerUrl) {
		return nil, false
	}
	return o.ContainerUrl, true
}

// HasContainerUrl returns a boolean if a field has been set.
func (o *Run) HasContainerUrl() bool {
	if o != nil && !IsNil(o.ContainerUrl) {
		return true
	}

	return false
}

// SetContainerUrl gets a reference to the given string and assigns it to the ContainerUrl field.
func (o *Run) SetContainerUrl(v string) {
	o.ContainerUrl = &v
}

// GetIsContainerServerReady returns the IsContainerServerReady field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetIsContainerServerReady() bool {
	if o == nil || IsNil(o.IsContainerServerReady.Get()) {
		var ret bool
		return ret
	}
	return *o.IsContainerServerReady.Get()
}

// GetIsContainerServerReadyOk returns a tuple with the IsContainerServerReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetIsContainerServerReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsContainerServerReady.Get(), o.IsContainerServerReady.IsSet()
}

// HasIsContainerServerReady returns a boolean if a field has been set.
func (o *Run) HasIsContainerServerReady() bool {
	if o != nil && o.IsContainerServerReady.IsSet() {
		return true
	}

	return false
}

// SetIsContainerServerReady gets a reference to the given NullableBool and assigns it to the IsContainerServerReady field.
func (o *Run) SetIsContainerServerReady(v bool) {
	o.IsContainerServerReady.Set(&v)
}
// SetIsContainerServerReadyNil sets the value for IsContainerServerReady to be an explicit nil
func (o *Run) SetIsContainerServerReadyNil() {
	o.IsContainerServerReady.Set(nil)
}

// UnsetIsContainerServerReady ensures that no value is present for IsContainerServerReady, not even an explicit nil
func (o *Run) UnsetIsContainerServerReady() {
	o.IsContainerServerReady.Unset()
}

// GetGitBranchName returns the GitBranchName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetGitBranchName() string {
	if o == nil || IsNil(o.GitBranchName.Get()) {
		var ret string
		return ret
	}
	return *o.GitBranchName.Get()
}

// GetGitBranchNameOk returns a tuple with the GitBranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetGitBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitBranchName.Get(), o.GitBranchName.IsSet()
}

// HasGitBranchName returns a boolean if a field has been set.
func (o *Run) HasGitBranchName() bool {
	if o != nil && o.GitBranchName.IsSet() {
		return true
	}

	return false
}

// SetGitBranchName gets a reference to the given NullableString and assigns it to the GitBranchName field.
func (o *Run) SetGitBranchName(v string) {
	o.GitBranchName.Set(&v)
}
// SetGitBranchNameNil sets the value for GitBranchName to be an explicit nil
func (o *Run) SetGitBranchNameNil() {
	o.GitBranchName.Set(nil)
}

// UnsetGitBranchName ensures that no value is present for GitBranchName, not even an explicit nil
func (o *Run) UnsetGitBranchName() {
	o.GitBranchName.Unset()
}

// GetUsage returns the Usage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetUsage() RunUsage {
	if o == nil || IsNil(o.Usage.Get()) {
		var ret RunUsage
		return ret
	}
	return *o.Usage.Get()
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetUsageOk() (*RunUsage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage.Get(), o.Usage.IsSet()
}

// HasUsage returns a boolean if a field has been set.
func (o *Run) HasUsage() bool {
	if o != nil && o.Usage.IsSet() {
		return true
	}

	return false
}

// SetUsage gets a reference to the given NullableRunUsage and assigns it to the Usage field.
func (o *Run) SetUsage(v RunUsage) {
	o.Usage.Set(&v)
}
// SetUsageNil sets the value for Usage to be an explicit nil
func (o *Run) SetUsageNil() {
	o.Usage.Set(nil)
}

// UnsetUsage ensures that no value is present for Usage, not even an explicit nil
func (o *Run) UnsetUsage() {
	o.Usage.Unset()
}

// GetUsageTotalUsd returns the UsageTotalUsd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetUsageTotalUsd() float32 {
	if o == nil || IsNil(o.UsageTotalUsd.Get()) {
		var ret float32
		return ret
	}
	return *o.UsageTotalUsd.Get()
}

// GetUsageTotalUsdOk returns a tuple with the UsageTotalUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetUsageTotalUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageTotalUsd.Get(), o.UsageTotalUsd.IsSet()
}

// HasUsageTotalUsd returns a boolean if a field has been set.
func (o *Run) HasUsageTotalUsd() bool {
	if o != nil && o.UsageTotalUsd.IsSet() {
		return true
	}

	return false
}

// SetUsageTotalUsd gets a reference to the given NullableFloat32 and assigns it to the UsageTotalUsd field.
func (o *Run) SetUsageTotalUsd(v float32) {
	o.UsageTotalUsd.Set(&v)
}
// SetUsageTotalUsdNil sets the value for UsageTotalUsd to be an explicit nil
func (o *Run) SetUsageTotalUsdNil() {
	o.UsageTotalUsd.Set(nil)
}

// UnsetUsageTotalUsd ensures that no value is present for UsageTotalUsd, not even an explicit nil
func (o *Run) UnsetUsageTotalUsd() {
	o.UsageTotalUsd.Unset()
}

// GetUsageUsd returns the UsageUsd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetUsageUsd() RunUsageUsd {
	if o == nil || IsNil(o.UsageUsd.Get()) {
		var ret RunUsageUsd
		return ret
	}
	return *o.UsageUsd.Get()
}

// GetUsageUsdOk returns a tuple with the UsageUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetUsageUsdOk() (*RunUsageUsd, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageUsd.Get(), o.UsageUsd.IsSet()
}

// HasUsageUsd returns a boolean if a field has been set.
func (o *Run) HasUsageUsd() bool {
	if o != nil && o.UsageUsd.IsSet() {
		return true
	}

	return false
}

// SetUsageUsd gets a reference to the given NullableRunUsageUsd and assigns it to the UsageUsd field.
func (o *Run) SetUsageUsd(v RunUsageUsd) {
	o.UsageUsd.Set(&v)
}
// SetUsageUsdNil sets the value for UsageUsd to be an explicit nil
func (o *Run) SetUsageUsdNil() {
	o.UsageUsd.Set(nil)
}

// UnsetUsageUsd ensures that no value is present for UsageUsd, not even an explicit nil
func (o *Run) UnsetUsageUsd() {
	o.UsageUsd.Unset()
}

// GetMetamorphs returns the Metamorphs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetMetamorphs() []Metamorph {
	if o == nil {
		var ret []Metamorph
		return ret
	}
	return o.Metamorphs
}

// GetMetamorphsOk returns a tuple with the Metamorphs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetMetamorphsOk() ([]Metamorph, bool) {
	if o == nil || IsNil(o.Metamorphs) {
		return nil, false
	}
	return o.Metamorphs, true
}

// HasMetamorphs returns a boolean if a field has been set.
func (o *Run) HasMetamorphs() bool {
	if o != nil && !IsNil(o.Metamorphs) {
		return true
	}

	return false
}

// SetMetamorphs gets a reference to the given []Metamorph and assigns it to the Metamorphs field.
func (o *Run) SetMetamorphs(v []Metamorph) {
	o.Metamorphs = v
}

// GetPlatformUsageBillingModel returns the PlatformUsageBillingModel field value if set, zero value otherwise.
func (o *Run) GetPlatformUsageBillingModel() string {
	if o == nil || IsNil(o.PlatformUsageBillingModel) {
		var ret string
		return ret
	}
	return *o.PlatformUsageBillingModel
}

// GetPlatformUsageBillingModelOk returns a tuple with the PlatformUsageBillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Run) GetPlatformUsageBillingModelOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformUsageBillingModel) {
		return nil, false
	}
	return o.PlatformUsageBillingModel, true
}

// HasPlatformUsageBillingModel returns a boolean if a field has been set.
func (o *Run) HasPlatformUsageBillingModel() bool {
	if o != nil && !IsNil(o.PlatformUsageBillingModel) {
		return true
	}

	return false
}

// SetPlatformUsageBillingModel gets a reference to the given string and assigns it to the PlatformUsageBillingModel field.
func (o *Run) SetPlatformUsageBillingModel(v string) {
	o.PlatformUsageBillingModel = &v
}

func (o Run) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Run) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["actId"] = o.ActId
	toSerialize["userId"] = o.UserId
	if o.ActorTaskId.IsSet() {
		toSerialize["actorTaskId"] = o.ActorTaskId.Get()
	}
	toSerialize["startedAt"] = o.StartedAt
	if o.FinishedAt.IsSet() {
		toSerialize["finishedAt"] = o.FinishedAt.Get()
	}
	toSerialize["status"] = o.Status
	if o.StatusMessage.IsSet() {
		toSerialize["statusMessage"] = o.StatusMessage.Get()
	}
	if o.IsStatusMessageTerminal.IsSet() {
		toSerialize["isStatusMessageTerminal"] = o.IsStatusMessageTerminal.Get()
	}
	toSerialize["meta"] = o.Meta
	if !IsNil(o.PricingInfo) {
		toSerialize["pricingInfo"] = o.PricingInfo
	}
	toSerialize["stats"] = o.Stats
	if !IsNil(o.ChargedEventCounts) {
		toSerialize["chargedEventCounts"] = o.ChargedEventCounts
	}
	toSerialize["options"] = o.Options
	toSerialize["buildId"] = o.BuildId
	if o.ExitCode.IsSet() {
		toSerialize["exitCode"] = o.ExitCode.Get()
	}
	toSerialize["generalAccess"] = o.GeneralAccess
	toSerialize["defaultKeyValueStoreId"] = o.DefaultKeyValueStoreId
	toSerialize["defaultDatasetId"] = o.DefaultDatasetId
	toSerialize["defaultRequestQueueId"] = o.DefaultRequestQueueId
	if !IsNil(o.StorageIds) {
		toSerialize["storageIds"] = o.StorageIds
	}
	if o.BuildNumber.IsSet() {
		toSerialize["buildNumber"] = o.BuildNumber.Get()
	}
	if !IsNil(o.ContainerUrl) {
		toSerialize["containerUrl"] = o.ContainerUrl
	}
	if o.IsContainerServerReady.IsSet() {
		toSerialize["isContainerServerReady"] = o.IsContainerServerReady.Get()
	}
	if o.GitBranchName.IsSet() {
		toSerialize["gitBranchName"] = o.GitBranchName.Get()
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
	if o.Metamorphs != nil {
		toSerialize["metamorphs"] = o.Metamorphs
	}
	if !IsNil(o.PlatformUsageBillingModel) {
		toSerialize["platformUsageBillingModel"] = o.PlatformUsageBillingModel
	}
	return toSerialize, nil
}

func (o *Run) UnmarshalJSON(data []byte) (err error) {
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
		"stats",
		"options",
		"buildId",
		"generalAccess",
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

	varRun := _Run{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRun)

	if err != nil {
		return err
	}

	*o = Run(varRun)

	return err
}

type NullableRun struct {
	value *Run
	isSet bool
}

func (v NullableRun) Get() *Run {
	return v.value
}

func (v *NullableRun) Set(val *Run) {
	v.value = val
	v.isSet = true
}

func (v NullableRun) IsSet() bool {
	return v.isSet
}

func (v *NullableRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRun(val *Run) *NullableRun {
	return &NullableRun{value: val, isSet: true}
}

func (v NullableRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


