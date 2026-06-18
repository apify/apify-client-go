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

// checks if the PricePerDatasetItemActorPricingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricePerDatasetItemActorPricingInfo{}

// PricePerDatasetItemActorPricingInfo struct for PricePerDatasetItemActorPricingInfo
type PricePerDatasetItemActorPricingInfo struct {
	// In [0, 1], fraction of pricePerUnitUsd that goes to Apify
	ApifyMarginPercentage float32 `json:"apifyMarginPercentage"`
	// When this pricing info record has been created
	CreatedAt time.Time `json:"createdAt"`
	// Since when is this pricing info record effective for a given Actor
	StartedAt time.Time `json:"startedAt"`
	NotifiedAboutFutureChangeAt *time.Time `json:"notifiedAboutFutureChangeAt,omitempty"`
	NotifiedAboutChangeAt *time.Time `json:"notifiedAboutChangeAt,omitempty"`
	ReasonForChange *string `json:"reasonForChange,omitempty"`
	IsPriceChangeNotificationSuppressed *bool `json:"isPriceChangeNotificationSuppressed,omitempty"`
	ForceContainsSignificantPriceChange *bool `json:"forceContainsSignificantPriceChange,omitempty"`
	PricingModel string `json:"pricingModel"`
	// Name of the unit that is being charged
	UnitName string `json:"unitName"`
	// Price per unit in USD. Mutually exclusive with `tieredPricing` - exactly one of the two is present on a pricing record. 
	PricePerUnitUsd *float32 `json:"pricePerUnitUsd,omitempty"`
	// Tiered price-per-dataset-item pricing, keyed by subscription tier (e.g. `FREE`, `BRONZE`, `SILVER`, `GOLD`, `PLATINUM`, `DIAMOND`). The actual price applied to a run is resolved from the user's tier. 
	TieredPricing map[string]TieredPricingPerDatasetItemEntry `json:"tieredPricing,omitempty"`
}

type _PricePerDatasetItemActorPricingInfo PricePerDatasetItemActorPricingInfo

// NewPricePerDatasetItemActorPricingInfo instantiates a new PricePerDatasetItemActorPricingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricePerDatasetItemActorPricingInfo(apifyMarginPercentage float32, createdAt time.Time, startedAt time.Time, pricingModel string, unitName string) *PricePerDatasetItemActorPricingInfo {
	this := PricePerDatasetItemActorPricingInfo{}
	this.ApifyMarginPercentage = apifyMarginPercentage
	this.CreatedAt = createdAt
	this.StartedAt = startedAt
	this.PricingModel = pricingModel
	this.UnitName = unitName
	return &this
}

// NewPricePerDatasetItemActorPricingInfoWithDefaults instantiates a new PricePerDatasetItemActorPricingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricePerDatasetItemActorPricingInfoWithDefaults() *PricePerDatasetItemActorPricingInfo {
	this := PricePerDatasetItemActorPricingInfo{}
	return &this
}

// GetApifyMarginPercentage returns the ApifyMarginPercentage field value
func (o *PricePerDatasetItemActorPricingInfo) GetApifyMarginPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ApifyMarginPercentage
}

// GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field value
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetApifyMarginPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApifyMarginPercentage, true
}

// SetApifyMarginPercentage sets field value
func (o *PricePerDatasetItemActorPricingInfo) SetApifyMarginPercentage(v float32) {
	o.ApifyMarginPercentage = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PricePerDatasetItemActorPricingInfo) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PricePerDatasetItemActorPricingInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStartedAt returns the StartedAt field value
func (o *PricePerDatasetItemActorPricingInfo) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *PricePerDatasetItemActorPricingInfo) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field value if set, zero value otherwise.
func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time {
	if o == nil || IsNil(o.NotifiedAboutFutureChangeAt) {
		var ret time.Time
		return ret
	}
	return *o.NotifiedAboutFutureChangeAt
}

// GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotifiedAboutFutureChangeAt) {
		return nil, false
	}
	return o.NotifiedAboutFutureChangeAt, true
}

// HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.
func (o *PricePerDatasetItemActorPricingInfo) HasNotifiedAboutFutureChangeAt() bool {
	if o != nil && !IsNil(o.NotifiedAboutFutureChangeAt) {
		return true
	}

	return false
}

// SetNotifiedAboutFutureChangeAt gets a reference to the given time.Time and assigns it to the NotifiedAboutFutureChangeAt field.
func (o *PricePerDatasetItemActorPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time) {
	o.NotifiedAboutFutureChangeAt = &v
}

// GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field value if set, zero value otherwise.
func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutChangeAt() time.Time {
	if o == nil || IsNil(o.NotifiedAboutChangeAt) {
		var ret time.Time
		return ret
	}
	return *o.NotifiedAboutChangeAt
}

// GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotifiedAboutChangeAt) {
		return nil, false
	}
	return o.NotifiedAboutChangeAt, true
}

// HasNotifiedAboutChangeAt returns a boolean if a field has been set.
func (o *PricePerDatasetItemActorPricingInfo) HasNotifiedAboutChangeAt() bool {
	if o != nil && !IsNil(o.NotifiedAboutChangeAt) {
		return true
	}

	return false
}

// SetNotifiedAboutChangeAt gets a reference to the given time.Time and assigns it to the NotifiedAboutChangeAt field.
func (o *PricePerDatasetItemActorPricingInfo) SetNotifiedAboutChangeAt(v time.Time) {
	o.NotifiedAboutChangeAt = &v
}

// GetReasonForChange returns the ReasonForChange field value if set, zero value otherwise.
func (o *PricePerDatasetItemActorPricingInfo) GetReasonForChange() string {
	if o == nil || IsNil(o.ReasonForChange) {
		var ret string
		return ret
	}
	return *o.ReasonForChange
}

// GetReasonForChangeOk returns a tuple with the ReasonForChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetReasonForChangeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonForChange) {
		return nil, false
	}
	return o.ReasonForChange, true
}

// HasReasonForChange returns a boolean if a field has been set.
func (o *PricePerDatasetItemActorPricingInfo) HasReasonForChange() bool {
	if o != nil && !IsNil(o.ReasonForChange) {
		return true
	}

	return false
}

// SetReasonForChange gets a reference to the given string and assigns it to the ReasonForChange field.
func (o *PricePerDatasetItemActorPricingInfo) SetReasonForChange(v string) {
	o.ReasonForChange = &v
}

// GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field value if set, zero value otherwise.
func (o *PricePerDatasetItemActorPricingInfo) GetIsPriceChangeNotificationSuppressed() bool {
	if o == nil || IsNil(o.IsPriceChangeNotificationSuppressed) {
		var ret bool
		return ret
	}
	return *o.IsPriceChangeNotificationSuppressed
}

// GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPriceChangeNotificationSuppressed) {
		return nil, false
	}
	return o.IsPriceChangeNotificationSuppressed, true
}

// HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.
func (o *PricePerDatasetItemActorPricingInfo) HasIsPriceChangeNotificationSuppressed() bool {
	if o != nil && !IsNil(o.IsPriceChangeNotificationSuppressed) {
		return true
	}

	return false
}

// SetIsPriceChangeNotificationSuppressed gets a reference to the given bool and assigns it to the IsPriceChangeNotificationSuppressed field.
func (o *PricePerDatasetItemActorPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool) {
	o.IsPriceChangeNotificationSuppressed = &v
}

// GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field value if set, zero value otherwise.
func (o *PricePerDatasetItemActorPricingInfo) GetForceContainsSignificantPriceChange() bool {
	if o == nil || IsNil(o.ForceContainsSignificantPriceChange) {
		var ret bool
		return ret
	}
	return *o.ForceContainsSignificantPriceChange
}

// GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceContainsSignificantPriceChange) {
		return nil, false
	}
	return o.ForceContainsSignificantPriceChange, true
}

// HasForceContainsSignificantPriceChange returns a boolean if a field has been set.
func (o *PricePerDatasetItemActorPricingInfo) HasForceContainsSignificantPriceChange() bool {
	if o != nil && !IsNil(o.ForceContainsSignificantPriceChange) {
		return true
	}

	return false
}

// SetForceContainsSignificantPriceChange gets a reference to the given bool and assigns it to the ForceContainsSignificantPriceChange field.
func (o *PricePerDatasetItemActorPricingInfo) SetForceContainsSignificantPriceChange(v bool) {
	o.ForceContainsSignificantPriceChange = &v
}

// GetPricingModel returns the PricingModel field value
func (o *PricePerDatasetItemActorPricingInfo) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetPricingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *PricePerDatasetItemActorPricingInfo) SetPricingModel(v string) {
	o.PricingModel = v
}

// GetUnitName returns the UnitName field value
func (o *PricePerDatasetItemActorPricingInfo) GetUnitName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitName, true
}

// SetUnitName sets field value
func (o *PricePerDatasetItemActorPricingInfo) SetUnitName(v string) {
	o.UnitName = v
}

// GetPricePerUnitUsd returns the PricePerUnitUsd field value if set, zero value otherwise.
func (o *PricePerDatasetItemActorPricingInfo) GetPricePerUnitUsd() float32 {
	if o == nil || IsNil(o.PricePerUnitUsd) {
		var ret float32
		return ret
	}
	return *o.PricePerUnitUsd
}

// GetPricePerUnitUsdOk returns a tuple with the PricePerUnitUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetPricePerUnitUsdOk() (*float32, bool) {
	if o == nil || IsNil(o.PricePerUnitUsd) {
		return nil, false
	}
	return o.PricePerUnitUsd, true
}

// HasPricePerUnitUsd returns a boolean if a field has been set.
func (o *PricePerDatasetItemActorPricingInfo) HasPricePerUnitUsd() bool {
	if o != nil && !IsNil(o.PricePerUnitUsd) {
		return true
	}

	return false
}

// SetPricePerUnitUsd gets a reference to the given float32 and assigns it to the PricePerUnitUsd field.
func (o *PricePerDatasetItemActorPricingInfo) SetPricePerUnitUsd(v float32) {
	o.PricePerUnitUsd = &v
}

// GetTieredPricing returns the TieredPricing field value if set, zero value otherwise.
func (o *PricePerDatasetItemActorPricingInfo) GetTieredPricing() map[string]TieredPricingPerDatasetItemEntry {
	if o == nil || IsNil(o.TieredPricing) {
		var ret map[string]TieredPricingPerDatasetItemEntry
		return ret
	}
	return o.TieredPricing
}

// GetTieredPricingOk returns a tuple with the TieredPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricePerDatasetItemActorPricingInfo) GetTieredPricingOk() (map[string]TieredPricingPerDatasetItemEntry, bool) {
	if o == nil || IsNil(o.TieredPricing) {
		return map[string]TieredPricingPerDatasetItemEntry{}, false
	}
	return o.TieredPricing, true
}

// HasTieredPricing returns a boolean if a field has been set.
func (o *PricePerDatasetItemActorPricingInfo) HasTieredPricing() bool {
	if o != nil && !IsNil(o.TieredPricing) {
		return true
	}

	return false
}

// SetTieredPricing gets a reference to the given map[string]TieredPricingPerDatasetItemEntry and assigns it to the TieredPricing field.
func (o *PricePerDatasetItemActorPricingInfo) SetTieredPricing(v map[string]TieredPricingPerDatasetItemEntry) {
	o.TieredPricing = v
}

func (o PricePerDatasetItemActorPricingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricePerDatasetItemActorPricingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apifyMarginPercentage"] = o.ApifyMarginPercentage
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.NotifiedAboutFutureChangeAt) {
		toSerialize["notifiedAboutFutureChangeAt"] = o.NotifiedAboutFutureChangeAt
	}
	if !IsNil(o.NotifiedAboutChangeAt) {
		toSerialize["notifiedAboutChangeAt"] = o.NotifiedAboutChangeAt
	}
	if !IsNil(o.ReasonForChange) {
		toSerialize["reasonForChange"] = o.ReasonForChange
	}
	if !IsNil(o.IsPriceChangeNotificationSuppressed) {
		toSerialize["isPriceChangeNotificationSuppressed"] = o.IsPriceChangeNotificationSuppressed
	}
	if !IsNil(o.ForceContainsSignificantPriceChange) {
		toSerialize["forceContainsSignificantPriceChange"] = o.ForceContainsSignificantPriceChange
	}
	toSerialize["pricingModel"] = o.PricingModel
	toSerialize["unitName"] = o.UnitName
	if !IsNil(o.PricePerUnitUsd) {
		toSerialize["pricePerUnitUsd"] = o.PricePerUnitUsd
	}
	if !IsNil(o.TieredPricing) {
		toSerialize["tieredPricing"] = o.TieredPricing
	}
	return toSerialize, nil
}

func (o *PricePerDatasetItemActorPricingInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apifyMarginPercentage",
		"createdAt",
		"startedAt",
		"pricingModel",
		"unitName",
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

	varPricePerDatasetItemActorPricingInfo := _PricePerDatasetItemActorPricingInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPricePerDatasetItemActorPricingInfo)

	if err != nil {
		return err
	}

	*o = PricePerDatasetItemActorPricingInfo(varPricePerDatasetItemActorPricingInfo)

	return err
}

type NullablePricePerDatasetItemActorPricingInfo struct {
	value *PricePerDatasetItemActorPricingInfo
	isSet bool
}

func (v NullablePricePerDatasetItemActorPricingInfo) Get() *PricePerDatasetItemActorPricingInfo {
	return v.value
}

func (v *NullablePricePerDatasetItemActorPricingInfo) Set(val *PricePerDatasetItemActorPricingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePricePerDatasetItemActorPricingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePricePerDatasetItemActorPricingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricePerDatasetItemActorPricingInfo(val *PricePerDatasetItemActorPricingInfo) *NullablePricePerDatasetItemActorPricingInfo {
	return &NullablePricePerDatasetItemActorPricingInfo{value: val, isSet: true}
}

func (v NullablePricePerDatasetItemActorPricingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricePerDatasetItemActorPricingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


