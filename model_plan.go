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

// checks if the Plan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plan{}

// Plan struct for Plan
type Plan struct {
	Id string `json:"id"`
	Description string `json:"description"`
	IsEnabled bool `json:"isEnabled"`
	MonthlyBasePriceUsd float32 `json:"monthlyBasePriceUsd"`
	MonthlyUsageCreditsUsd float32 `json:"monthlyUsageCreditsUsd"`
	UsageDiscountPercent *float32 `json:"usageDiscountPercent,omitempty"`
	EnabledPlatformFeatures []string `json:"enabledPlatformFeatures"`
	MaxMonthlyUsageUsd float32 `json:"maxMonthlyUsageUsd"`
	MaxActorMemoryGbytes float32 `json:"maxActorMemoryGbytes"`
	MaxMonthlyActorComputeUnits float32 `json:"maxMonthlyActorComputeUnits"`
	MaxMonthlyResidentialProxyGbytes float32 `json:"maxMonthlyResidentialProxyGbytes"`
	MaxMonthlyProxySerps int32 `json:"maxMonthlyProxySerps"`
	MaxMonthlyExternalDataTransferGbytes float32 `json:"maxMonthlyExternalDataTransferGbytes"`
	MaxActorCount int32 `json:"maxActorCount"`
	MaxActorTaskCount int32 `json:"maxActorTaskCount"`
	DataRetentionDays int32 `json:"dataRetentionDays"`
	// A dictionary mapping proxy group names to the number of available proxies in each group. The keys are proxy group names (e.g., \"RESIDENTIAL\", \"DATACENTER\") and values are the count of available proxies. 
	AvailableProxyGroups map[string]int32 `json:"availableProxyGroups"`
	TeamAccountSeatCount int32 `json:"teamAccountSeatCount"`
	SupportLevel string `json:"supportLevel"`
	AvailableAddOns []string `json:"availableAddOns"`
	Tier *string `json:"tier,omitempty"`
	ApiRateLimitBoosts *int32 `json:"apiRateLimitBoosts,omitempty"`
	MaxScheduleCount *int32 `json:"maxScheduleCount,omitempty"`
	MaxConcurrentActorRuns *int32 `json:"maxConcurrentActorRuns,omitempty"`
	// Pricing details for this plan.
	PlanPricing map[string]interface{} `json:"planPricing,omitempty"`
}

type _Plan Plan

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan(id string, description string, isEnabled bool, monthlyBasePriceUsd float32, monthlyUsageCreditsUsd float32, enabledPlatformFeatures []string, maxMonthlyUsageUsd float32, maxActorMemoryGbytes float32, maxMonthlyActorComputeUnits float32, maxMonthlyResidentialProxyGbytes float32, maxMonthlyProxySerps int32, maxMonthlyExternalDataTransferGbytes float32, maxActorCount int32, maxActorTaskCount int32, dataRetentionDays int32, availableProxyGroups map[string]int32, teamAccountSeatCount int32, supportLevel string, availableAddOns []string) *Plan {
	this := Plan{}
	this.Id = id
	this.Description = description
	this.IsEnabled = isEnabled
	this.MonthlyBasePriceUsd = monthlyBasePriceUsd
	this.MonthlyUsageCreditsUsd = monthlyUsageCreditsUsd
	this.EnabledPlatformFeatures = enabledPlatformFeatures
	this.MaxMonthlyUsageUsd = maxMonthlyUsageUsd
	this.MaxActorMemoryGbytes = maxActorMemoryGbytes
	this.MaxMonthlyActorComputeUnits = maxMonthlyActorComputeUnits
	this.MaxMonthlyResidentialProxyGbytes = maxMonthlyResidentialProxyGbytes
	this.MaxMonthlyProxySerps = maxMonthlyProxySerps
	this.MaxMonthlyExternalDataTransferGbytes = maxMonthlyExternalDataTransferGbytes
	this.MaxActorCount = maxActorCount
	this.MaxActorTaskCount = maxActorTaskCount
	this.DataRetentionDays = dataRetentionDays
	this.AvailableProxyGroups = availableProxyGroups
	this.TeamAccountSeatCount = teamAccountSeatCount
	this.SupportLevel = supportLevel
	this.AvailableAddOns = availableAddOns
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	return &this
}

// GetId returns the Id field value
func (o *Plan) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Plan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Plan) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *Plan) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Plan) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Plan) SetDescription(v string) {
	o.Description = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *Plan) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *Plan) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *Plan) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetMonthlyBasePriceUsd returns the MonthlyBasePriceUsd field value
func (o *Plan) GetMonthlyBasePriceUsd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MonthlyBasePriceUsd
}

// GetMonthlyBasePriceUsdOk returns a tuple with the MonthlyBasePriceUsd field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMonthlyBasePriceUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyBasePriceUsd, true
}

// SetMonthlyBasePriceUsd sets field value
func (o *Plan) SetMonthlyBasePriceUsd(v float32) {
	o.MonthlyBasePriceUsd = v
}

// GetMonthlyUsageCreditsUsd returns the MonthlyUsageCreditsUsd field value
func (o *Plan) GetMonthlyUsageCreditsUsd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MonthlyUsageCreditsUsd
}

// GetMonthlyUsageCreditsUsdOk returns a tuple with the MonthlyUsageCreditsUsd field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMonthlyUsageCreditsUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyUsageCreditsUsd, true
}

// SetMonthlyUsageCreditsUsd sets field value
func (o *Plan) SetMonthlyUsageCreditsUsd(v float32) {
	o.MonthlyUsageCreditsUsd = v
}

// GetUsageDiscountPercent returns the UsageDiscountPercent field value if set, zero value otherwise.
func (o *Plan) GetUsageDiscountPercent() float32 {
	if o == nil || IsNil(o.UsageDiscountPercent) {
		var ret float32
		return ret
	}
	return *o.UsageDiscountPercent
}

// GetUsageDiscountPercentOk returns a tuple with the UsageDiscountPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetUsageDiscountPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.UsageDiscountPercent) {
		return nil, false
	}
	return o.UsageDiscountPercent, true
}

// HasUsageDiscountPercent returns a boolean if a field has been set.
func (o *Plan) HasUsageDiscountPercent() bool {
	if o != nil && !IsNil(o.UsageDiscountPercent) {
		return true
	}

	return false
}

// SetUsageDiscountPercent gets a reference to the given float32 and assigns it to the UsageDiscountPercent field.
func (o *Plan) SetUsageDiscountPercent(v float32) {
	o.UsageDiscountPercent = &v
}

// GetEnabledPlatformFeatures returns the EnabledPlatformFeatures field value
func (o *Plan) GetEnabledPlatformFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EnabledPlatformFeatures
}

// GetEnabledPlatformFeaturesOk returns a tuple with the EnabledPlatformFeatures field value
// and a boolean to check if the value has been set.
func (o *Plan) GetEnabledPlatformFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledPlatformFeatures, true
}

// SetEnabledPlatformFeatures sets field value
func (o *Plan) SetEnabledPlatformFeatures(v []string) {
	o.EnabledPlatformFeatures = v
}

// GetMaxMonthlyUsageUsd returns the MaxMonthlyUsageUsd field value
func (o *Plan) GetMaxMonthlyUsageUsd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxMonthlyUsageUsd
}

// GetMaxMonthlyUsageUsdOk returns a tuple with the MaxMonthlyUsageUsd field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxMonthlyUsageUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMonthlyUsageUsd, true
}

// SetMaxMonthlyUsageUsd sets field value
func (o *Plan) SetMaxMonthlyUsageUsd(v float32) {
	o.MaxMonthlyUsageUsd = v
}

// GetMaxActorMemoryGbytes returns the MaxActorMemoryGbytes field value
func (o *Plan) GetMaxActorMemoryGbytes() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxActorMemoryGbytes
}

// GetMaxActorMemoryGbytesOk returns a tuple with the MaxActorMemoryGbytes field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxActorMemoryGbytesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxActorMemoryGbytes, true
}

// SetMaxActorMemoryGbytes sets field value
func (o *Plan) SetMaxActorMemoryGbytes(v float32) {
	o.MaxActorMemoryGbytes = v
}

// GetMaxMonthlyActorComputeUnits returns the MaxMonthlyActorComputeUnits field value
func (o *Plan) GetMaxMonthlyActorComputeUnits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxMonthlyActorComputeUnits
}

// GetMaxMonthlyActorComputeUnitsOk returns a tuple with the MaxMonthlyActorComputeUnits field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxMonthlyActorComputeUnitsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMonthlyActorComputeUnits, true
}

// SetMaxMonthlyActorComputeUnits sets field value
func (o *Plan) SetMaxMonthlyActorComputeUnits(v float32) {
	o.MaxMonthlyActorComputeUnits = v
}

// GetMaxMonthlyResidentialProxyGbytes returns the MaxMonthlyResidentialProxyGbytes field value
func (o *Plan) GetMaxMonthlyResidentialProxyGbytes() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxMonthlyResidentialProxyGbytes
}

// GetMaxMonthlyResidentialProxyGbytesOk returns a tuple with the MaxMonthlyResidentialProxyGbytes field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxMonthlyResidentialProxyGbytesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMonthlyResidentialProxyGbytes, true
}

// SetMaxMonthlyResidentialProxyGbytes sets field value
func (o *Plan) SetMaxMonthlyResidentialProxyGbytes(v float32) {
	o.MaxMonthlyResidentialProxyGbytes = v
}

// GetMaxMonthlyProxySerps returns the MaxMonthlyProxySerps field value
func (o *Plan) GetMaxMonthlyProxySerps() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxMonthlyProxySerps
}

// GetMaxMonthlyProxySerpsOk returns a tuple with the MaxMonthlyProxySerps field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxMonthlyProxySerpsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMonthlyProxySerps, true
}

// SetMaxMonthlyProxySerps sets field value
func (o *Plan) SetMaxMonthlyProxySerps(v int32) {
	o.MaxMonthlyProxySerps = v
}

// GetMaxMonthlyExternalDataTransferGbytes returns the MaxMonthlyExternalDataTransferGbytes field value
func (o *Plan) GetMaxMonthlyExternalDataTransferGbytes() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxMonthlyExternalDataTransferGbytes
}

// GetMaxMonthlyExternalDataTransferGbytesOk returns a tuple with the MaxMonthlyExternalDataTransferGbytes field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxMonthlyExternalDataTransferGbytesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMonthlyExternalDataTransferGbytes, true
}

// SetMaxMonthlyExternalDataTransferGbytes sets field value
func (o *Plan) SetMaxMonthlyExternalDataTransferGbytes(v float32) {
	o.MaxMonthlyExternalDataTransferGbytes = v
}

// GetMaxActorCount returns the MaxActorCount field value
func (o *Plan) GetMaxActorCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxActorCount
}

// GetMaxActorCountOk returns a tuple with the MaxActorCount field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxActorCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxActorCount, true
}

// SetMaxActorCount sets field value
func (o *Plan) SetMaxActorCount(v int32) {
	o.MaxActorCount = v
}

// GetMaxActorTaskCount returns the MaxActorTaskCount field value
func (o *Plan) GetMaxActorTaskCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxActorTaskCount
}

// GetMaxActorTaskCountOk returns a tuple with the MaxActorTaskCount field value
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxActorTaskCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxActorTaskCount, true
}

// SetMaxActorTaskCount sets field value
func (o *Plan) SetMaxActorTaskCount(v int32) {
	o.MaxActorTaskCount = v
}

// GetDataRetentionDays returns the DataRetentionDays field value
func (o *Plan) GetDataRetentionDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DataRetentionDays
}

// GetDataRetentionDaysOk returns a tuple with the DataRetentionDays field value
// and a boolean to check if the value has been set.
func (o *Plan) GetDataRetentionDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataRetentionDays, true
}

// SetDataRetentionDays sets field value
func (o *Plan) SetDataRetentionDays(v int32) {
	o.DataRetentionDays = v
}

// GetAvailableProxyGroups returns the AvailableProxyGroups field value
func (o *Plan) GetAvailableProxyGroups() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.AvailableProxyGroups
}

// GetAvailableProxyGroupsOk returns a tuple with the AvailableProxyGroups field value
// and a boolean to check if the value has been set.
func (o *Plan) GetAvailableProxyGroupsOk() (map[string]int32, bool) {
	if o == nil {
		return map[string]int32{}, false
	}
	return o.AvailableProxyGroups, true
}

// SetAvailableProxyGroups sets field value
func (o *Plan) SetAvailableProxyGroups(v map[string]int32) {
	o.AvailableProxyGroups = v
}

// GetTeamAccountSeatCount returns the TeamAccountSeatCount field value
func (o *Plan) GetTeamAccountSeatCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TeamAccountSeatCount
}

// GetTeamAccountSeatCountOk returns a tuple with the TeamAccountSeatCount field value
// and a boolean to check if the value has been set.
func (o *Plan) GetTeamAccountSeatCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamAccountSeatCount, true
}

// SetTeamAccountSeatCount sets field value
func (o *Plan) SetTeamAccountSeatCount(v int32) {
	o.TeamAccountSeatCount = v
}

// GetSupportLevel returns the SupportLevel field value
func (o *Plan) GetSupportLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportLevel
}

// GetSupportLevelOk returns a tuple with the SupportLevel field value
// and a boolean to check if the value has been set.
func (o *Plan) GetSupportLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportLevel, true
}

// SetSupportLevel sets field value
func (o *Plan) SetSupportLevel(v string) {
	o.SupportLevel = v
}

// GetAvailableAddOns returns the AvailableAddOns field value
func (o *Plan) GetAvailableAddOns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableAddOns
}

// GetAvailableAddOnsOk returns a tuple with the AvailableAddOns field value
// and a boolean to check if the value has been set.
func (o *Plan) GetAvailableAddOnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableAddOns, true
}

// SetAvailableAddOns sets field value
func (o *Plan) SetAvailableAddOns(v []string) {
	o.AvailableAddOns = v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *Plan) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *Plan) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *Plan) SetTier(v string) {
	o.Tier = &v
}

// GetApiRateLimitBoosts returns the ApiRateLimitBoosts field value if set, zero value otherwise.
func (o *Plan) GetApiRateLimitBoosts() int32 {
	if o == nil || IsNil(o.ApiRateLimitBoosts) {
		var ret int32
		return ret
	}
	return *o.ApiRateLimitBoosts
}

// GetApiRateLimitBoostsOk returns a tuple with the ApiRateLimitBoosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetApiRateLimitBoostsOk() (*int32, bool) {
	if o == nil || IsNil(o.ApiRateLimitBoosts) {
		return nil, false
	}
	return o.ApiRateLimitBoosts, true
}

// HasApiRateLimitBoosts returns a boolean if a field has been set.
func (o *Plan) HasApiRateLimitBoosts() bool {
	if o != nil && !IsNil(o.ApiRateLimitBoosts) {
		return true
	}

	return false
}

// SetApiRateLimitBoosts gets a reference to the given int32 and assigns it to the ApiRateLimitBoosts field.
func (o *Plan) SetApiRateLimitBoosts(v int32) {
	o.ApiRateLimitBoosts = &v
}

// GetMaxScheduleCount returns the MaxScheduleCount field value if set, zero value otherwise.
func (o *Plan) GetMaxScheduleCount() int32 {
	if o == nil || IsNil(o.MaxScheduleCount) {
		var ret int32
		return ret
	}
	return *o.MaxScheduleCount
}

// GetMaxScheduleCountOk returns a tuple with the MaxScheduleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxScheduleCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxScheduleCount) {
		return nil, false
	}
	return o.MaxScheduleCount, true
}

// HasMaxScheduleCount returns a boolean if a field has been set.
func (o *Plan) HasMaxScheduleCount() bool {
	if o != nil && !IsNil(o.MaxScheduleCount) {
		return true
	}

	return false
}

// SetMaxScheduleCount gets a reference to the given int32 and assigns it to the MaxScheduleCount field.
func (o *Plan) SetMaxScheduleCount(v int32) {
	o.MaxScheduleCount = &v
}

// GetMaxConcurrentActorRuns returns the MaxConcurrentActorRuns field value if set, zero value otherwise.
func (o *Plan) GetMaxConcurrentActorRuns() int32 {
	if o == nil || IsNil(o.MaxConcurrentActorRuns) {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentActorRuns
}

// GetMaxConcurrentActorRunsOk returns a tuple with the MaxConcurrentActorRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetMaxConcurrentActorRunsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConcurrentActorRuns) {
		return nil, false
	}
	return o.MaxConcurrentActorRuns, true
}

// HasMaxConcurrentActorRuns returns a boolean if a field has been set.
func (o *Plan) HasMaxConcurrentActorRuns() bool {
	if o != nil && !IsNil(o.MaxConcurrentActorRuns) {
		return true
	}

	return false
}

// SetMaxConcurrentActorRuns gets a reference to the given int32 and assigns it to the MaxConcurrentActorRuns field.
func (o *Plan) SetMaxConcurrentActorRuns(v int32) {
	o.MaxConcurrentActorRuns = &v
}

// GetPlanPricing returns the PlanPricing field value if set, zero value otherwise.
func (o *Plan) GetPlanPricing() map[string]interface{} {
	if o == nil || IsNil(o.PlanPricing) {
		var ret map[string]interface{}
		return ret
	}
	return o.PlanPricing
}

// GetPlanPricingOk returns a tuple with the PlanPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetPlanPricingOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PlanPricing) {
		return map[string]interface{}{}, false
	}
	return o.PlanPricing, true
}

// HasPlanPricing returns a boolean if a field has been set.
func (o *Plan) HasPlanPricing() bool {
	if o != nil && !IsNil(o.PlanPricing) {
		return true
	}

	return false
}

// SetPlanPricing gets a reference to the given map[string]interface{} and assigns it to the PlanPricing field.
func (o *Plan) SetPlanPricing(v map[string]interface{}) {
	o.PlanPricing = v
}

func (o Plan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Plan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["description"] = o.Description
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["monthlyBasePriceUsd"] = o.MonthlyBasePriceUsd
	toSerialize["monthlyUsageCreditsUsd"] = o.MonthlyUsageCreditsUsd
	if !IsNil(o.UsageDiscountPercent) {
		toSerialize["usageDiscountPercent"] = o.UsageDiscountPercent
	}
	toSerialize["enabledPlatformFeatures"] = o.EnabledPlatformFeatures
	toSerialize["maxMonthlyUsageUsd"] = o.MaxMonthlyUsageUsd
	toSerialize["maxActorMemoryGbytes"] = o.MaxActorMemoryGbytes
	toSerialize["maxMonthlyActorComputeUnits"] = o.MaxMonthlyActorComputeUnits
	toSerialize["maxMonthlyResidentialProxyGbytes"] = o.MaxMonthlyResidentialProxyGbytes
	toSerialize["maxMonthlyProxySerps"] = o.MaxMonthlyProxySerps
	toSerialize["maxMonthlyExternalDataTransferGbytes"] = o.MaxMonthlyExternalDataTransferGbytes
	toSerialize["maxActorCount"] = o.MaxActorCount
	toSerialize["maxActorTaskCount"] = o.MaxActorTaskCount
	toSerialize["dataRetentionDays"] = o.DataRetentionDays
	toSerialize["availableProxyGroups"] = o.AvailableProxyGroups
	toSerialize["teamAccountSeatCount"] = o.TeamAccountSeatCount
	toSerialize["supportLevel"] = o.SupportLevel
	toSerialize["availableAddOns"] = o.AvailableAddOns
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.ApiRateLimitBoosts) {
		toSerialize["apiRateLimitBoosts"] = o.ApiRateLimitBoosts
	}
	if !IsNil(o.MaxScheduleCount) {
		toSerialize["maxScheduleCount"] = o.MaxScheduleCount
	}
	if !IsNil(o.MaxConcurrentActorRuns) {
		toSerialize["maxConcurrentActorRuns"] = o.MaxConcurrentActorRuns
	}
	if !IsNil(o.PlanPricing) {
		toSerialize["planPricing"] = o.PlanPricing
	}
	return toSerialize, nil
}

func (o *Plan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"description",
		"isEnabled",
		"monthlyBasePriceUsd",
		"monthlyUsageCreditsUsd",
		"enabledPlatformFeatures",
		"maxMonthlyUsageUsd",
		"maxActorMemoryGbytes",
		"maxMonthlyActorComputeUnits",
		"maxMonthlyResidentialProxyGbytes",
		"maxMonthlyProxySerps",
		"maxMonthlyExternalDataTransferGbytes",
		"maxActorCount",
		"maxActorTaskCount",
		"dataRetentionDays",
		"availableProxyGroups",
		"teamAccountSeatCount",
		"supportLevel",
		"availableAddOns",
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

	varPlan := _Plan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlan)

	if err != nil {
		return err
	}

	*o = Plan(varPlan)

	return err
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


