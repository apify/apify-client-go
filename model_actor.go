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

// checks if the Actor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Actor{}

// Actor struct for Actor
type Actor struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Name string `json:"name"`
	Username string `json:"username"`
	Description NullableString `json:"description,omitempty"`
	// Deprecated
	RestartOnError *bool `json:"restartOnError,omitempty"`
	IsPublic bool `json:"isPublic"`
	ActorPermissionLevel *ActorPermissionLevel `json:"actorPermissionLevel,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Stats ActorStats `json:"stats"`
	// 
	Versions []Version `json:"versions"`
	PricingInfos []ActorRunPricingInfo `json:"pricingInfos,omitempty"`
	DefaultRunOptions DefaultRunOptions `json:"defaultRunOptions"`
	ExampleRunInput NullableExampleRunInput `json:"exampleRunInput,omitempty"`
	IsDeprecated NullableBool `json:"isDeprecated,omitempty"`
	DeploymentKey *string `json:"deploymentKey,omitempty"`
	Title NullableString `json:"title,omitempty"`
	// A dictionary mapping build tag names (e.g., \"latest\", \"beta\") to their build information.
	TaggedBuilds map[string]TaggedBuildsValue `json:"taggedBuilds,omitempty"`
	ActorStandby NullableActorStandby `json:"actorStandby,omitempty"`
	// A brief, LLM-generated readme summary
	ReadmeSummary *string `json:"readmeSummary,omitempty"`
	SeoTitle NullableString `json:"seoTitle,omitempty"`
	SeoDescription NullableString `json:"seoDescription,omitempty"`
	PictureUrl NullableString `json:"pictureUrl,omitempty"`
	StandbyUrl NullableString `json:"standbyUrl,omitempty"`
	Notice *string `json:"notice,omitempty"`
	Categories []string `json:"categories,omitempty"`
	IsCritical *bool `json:"isCritical,omitempty"`
	IsGeneric *bool `json:"isGeneric,omitempty"`
	IsSourceCodeHidden *bool `json:"isSourceCodeHidden,omitempty"`
	HasNoDataset *bool `json:"hasNoDataset,omitempty"`
}

type _Actor Actor

// NewActor instantiates a new Actor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActor(id string, userId string, name string, username string, isPublic bool, createdAt time.Time, modifiedAt time.Time, stats ActorStats, versions []Version, defaultRunOptions DefaultRunOptions) *Actor {
	this := Actor{}
	this.Id = id
	this.UserId = userId
	this.Name = name
	this.Username = username
	this.IsPublic = isPublic
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Stats = stats
	this.Versions = versions
	this.DefaultRunOptions = defaultRunOptions
	return &this
}

// NewActorWithDefaults instantiates a new Actor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorWithDefaults() *Actor {
	this := Actor{}
	return &this
}

// GetId returns the Id field value
func (o *Actor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Actor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Actor) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *Actor) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Actor) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Actor) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *Actor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Actor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Actor) SetName(v string) {
	o.Name = v
}

// GetUsername returns the Username field value
func (o *Actor) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Actor) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Actor) SetUsername(v string) {
	o.Username = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Actor) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Actor) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Actor) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Actor) UnsetDescription() {
	o.Description.Unset()
}

// GetRestartOnError returns the RestartOnError field value if set, zero value otherwise.
// Deprecated
func (o *Actor) GetRestartOnError() bool {
	if o == nil || IsNil(o.RestartOnError) {
		var ret bool
		return ret
	}
	return *o.RestartOnError
}

// GetRestartOnErrorOk returns a tuple with the RestartOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Actor) GetRestartOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.RestartOnError) {
		return nil, false
	}
	return o.RestartOnError, true
}

// HasRestartOnError returns a boolean if a field has been set.
func (o *Actor) HasRestartOnError() bool {
	if o != nil && !IsNil(o.RestartOnError) {
		return true
	}

	return false
}

// SetRestartOnError gets a reference to the given bool and assigns it to the RestartOnError field.
// Deprecated
func (o *Actor) SetRestartOnError(v bool) {
	o.RestartOnError = &v
}

// GetIsPublic returns the IsPublic field value
func (o *Actor) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *Actor) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *Actor) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetActorPermissionLevel returns the ActorPermissionLevel field value if set, zero value otherwise.
func (o *Actor) GetActorPermissionLevel() ActorPermissionLevel {
	if o == nil || IsNil(o.ActorPermissionLevel) {
		var ret ActorPermissionLevel
		return ret
	}
	return *o.ActorPermissionLevel
}

// GetActorPermissionLevelOk returns a tuple with the ActorPermissionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetActorPermissionLevelOk() (*ActorPermissionLevel, bool) {
	if o == nil || IsNil(o.ActorPermissionLevel) {
		return nil, false
	}
	return o.ActorPermissionLevel, true
}

// HasActorPermissionLevel returns a boolean if a field has been set.
func (o *Actor) HasActorPermissionLevel() bool {
	if o != nil && !IsNil(o.ActorPermissionLevel) {
		return true
	}

	return false
}

// SetActorPermissionLevel gets a reference to the given ActorPermissionLevel and assigns it to the ActorPermissionLevel field.
func (o *Actor) SetActorPermissionLevel(v ActorPermissionLevel) {
	o.ActorPermissionLevel = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Actor) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Actor) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Actor) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *Actor) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *Actor) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *Actor) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetStats returns the Stats field value
func (o *Actor) GetStats() ActorStats {
	if o == nil {
		var ret ActorStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *Actor) GetStatsOk() (*ActorStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *Actor) SetStats(v ActorStats) {
	o.Stats = v
}

// GetVersions returns the Versions field value
func (o *Actor) GetVersions() []Version {
	if o == nil {
		var ret []Version
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *Actor) GetVersionsOk() ([]Version, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *Actor) SetVersions(v []Version) {
	o.Versions = v
}

// GetPricingInfos returns the PricingInfos field value if set, zero value otherwise.
func (o *Actor) GetPricingInfos() []ActorRunPricingInfo {
	if o == nil || IsNil(o.PricingInfos) {
		var ret []ActorRunPricingInfo
		return ret
	}
	return o.PricingInfos
}

// GetPricingInfosOk returns a tuple with the PricingInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetPricingInfosOk() ([]ActorRunPricingInfo, bool) {
	if o == nil || IsNil(o.PricingInfos) {
		return nil, false
	}
	return o.PricingInfos, true
}

// HasPricingInfos returns a boolean if a field has been set.
func (o *Actor) HasPricingInfos() bool {
	if o != nil && !IsNil(o.PricingInfos) {
		return true
	}

	return false
}

// SetPricingInfos gets a reference to the given []ActorRunPricingInfo and assigns it to the PricingInfos field.
func (o *Actor) SetPricingInfos(v []ActorRunPricingInfo) {
	o.PricingInfos = v
}

// GetDefaultRunOptions returns the DefaultRunOptions field value
func (o *Actor) GetDefaultRunOptions() DefaultRunOptions {
	if o == nil {
		var ret DefaultRunOptions
		return ret
	}

	return o.DefaultRunOptions
}

// GetDefaultRunOptionsOk returns a tuple with the DefaultRunOptions field value
// and a boolean to check if the value has been set.
func (o *Actor) GetDefaultRunOptionsOk() (*DefaultRunOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRunOptions, true
}

// SetDefaultRunOptions sets field value
func (o *Actor) SetDefaultRunOptions(v DefaultRunOptions) {
	o.DefaultRunOptions = v
}

// GetExampleRunInput returns the ExampleRunInput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetExampleRunInput() ExampleRunInput {
	if o == nil || IsNil(o.ExampleRunInput.Get()) {
		var ret ExampleRunInput
		return ret
	}
	return *o.ExampleRunInput.Get()
}

// GetExampleRunInputOk returns a tuple with the ExampleRunInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetExampleRunInputOk() (*ExampleRunInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExampleRunInput.Get(), o.ExampleRunInput.IsSet()
}

// HasExampleRunInput returns a boolean if a field has been set.
func (o *Actor) HasExampleRunInput() bool {
	if o != nil && o.ExampleRunInput.IsSet() {
		return true
	}

	return false
}

// SetExampleRunInput gets a reference to the given NullableExampleRunInput and assigns it to the ExampleRunInput field.
func (o *Actor) SetExampleRunInput(v ExampleRunInput) {
	o.ExampleRunInput.Set(&v)
}
// SetExampleRunInputNil sets the value for ExampleRunInput to be an explicit nil
func (o *Actor) SetExampleRunInputNil() {
	o.ExampleRunInput.Set(nil)
}

// UnsetExampleRunInput ensures that no value is present for ExampleRunInput, not even an explicit nil
func (o *Actor) UnsetExampleRunInput() {
	o.ExampleRunInput.Unset()
}

// GetIsDeprecated returns the IsDeprecated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetIsDeprecated() bool {
	if o == nil || IsNil(o.IsDeprecated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeprecated.Get()
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeprecated.Get(), o.IsDeprecated.IsSet()
}

// HasIsDeprecated returns a boolean if a field has been set.
func (o *Actor) HasIsDeprecated() bool {
	if o != nil && o.IsDeprecated.IsSet() {
		return true
	}

	return false
}

// SetIsDeprecated gets a reference to the given NullableBool and assigns it to the IsDeprecated field.
func (o *Actor) SetIsDeprecated(v bool) {
	o.IsDeprecated.Set(&v)
}
// SetIsDeprecatedNil sets the value for IsDeprecated to be an explicit nil
func (o *Actor) SetIsDeprecatedNil() {
	o.IsDeprecated.Set(nil)
}

// UnsetIsDeprecated ensures that no value is present for IsDeprecated, not even an explicit nil
func (o *Actor) UnsetIsDeprecated() {
	o.IsDeprecated.Unset()
}

// GetDeploymentKey returns the DeploymentKey field value if set, zero value otherwise.
func (o *Actor) GetDeploymentKey() string {
	if o == nil || IsNil(o.DeploymentKey) {
		var ret string
		return ret
	}
	return *o.DeploymentKey
}

// GetDeploymentKeyOk returns a tuple with the DeploymentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetDeploymentKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentKey) {
		return nil, false
	}
	return o.DeploymentKey, true
}

// HasDeploymentKey returns a boolean if a field has been set.
func (o *Actor) HasDeploymentKey() bool {
	if o != nil && !IsNil(o.DeploymentKey) {
		return true
	}

	return false
}

// SetDeploymentKey gets a reference to the given string and assigns it to the DeploymentKey field.
func (o *Actor) SetDeploymentKey(v string) {
	o.DeploymentKey = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Actor) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Actor) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Actor) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Actor) UnsetTitle() {
	o.Title.Unset()
}

// GetTaggedBuilds returns the TaggedBuilds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetTaggedBuilds() map[string]TaggedBuildsValue {
	if o == nil {
		var ret map[string]TaggedBuildsValue
		return ret
	}
	return o.TaggedBuilds
}

// GetTaggedBuildsOk returns a tuple with the TaggedBuilds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetTaggedBuildsOk() (map[string]TaggedBuildsValue, bool) {
	if o == nil || IsNil(o.TaggedBuilds) {
		return map[string]TaggedBuildsValue{}, false
	}
	return o.TaggedBuilds, true
}

// HasTaggedBuilds returns a boolean if a field has been set.
func (o *Actor) HasTaggedBuilds() bool {
	if o != nil && !IsNil(o.TaggedBuilds) {
		return true
	}

	return false
}

// SetTaggedBuilds gets a reference to the given map[string]TaggedBuildsValue and assigns it to the TaggedBuilds field.
func (o *Actor) SetTaggedBuilds(v map[string]TaggedBuildsValue) {
	o.TaggedBuilds = v
}

// GetActorStandby returns the ActorStandby field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetActorStandby() ActorStandby {
	if o == nil || IsNil(o.ActorStandby.Get()) {
		var ret ActorStandby
		return ret
	}
	return *o.ActorStandby.Get()
}

// GetActorStandbyOk returns a tuple with the ActorStandby field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetActorStandbyOk() (*ActorStandby, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorStandby.Get(), o.ActorStandby.IsSet()
}

// HasActorStandby returns a boolean if a field has been set.
func (o *Actor) HasActorStandby() bool {
	if o != nil && o.ActorStandby.IsSet() {
		return true
	}

	return false
}

// SetActorStandby gets a reference to the given NullableActorStandby and assigns it to the ActorStandby field.
func (o *Actor) SetActorStandby(v ActorStandby) {
	o.ActorStandby.Set(&v)
}
// SetActorStandbyNil sets the value for ActorStandby to be an explicit nil
func (o *Actor) SetActorStandbyNil() {
	o.ActorStandby.Set(nil)
}

// UnsetActorStandby ensures that no value is present for ActorStandby, not even an explicit nil
func (o *Actor) UnsetActorStandby() {
	o.ActorStandby.Unset()
}

// GetReadmeSummary returns the ReadmeSummary field value if set, zero value otherwise.
func (o *Actor) GetReadmeSummary() string {
	if o == nil || IsNil(o.ReadmeSummary) {
		var ret string
		return ret
	}
	return *o.ReadmeSummary
}

// GetReadmeSummaryOk returns a tuple with the ReadmeSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetReadmeSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ReadmeSummary) {
		return nil, false
	}
	return o.ReadmeSummary, true
}

// HasReadmeSummary returns a boolean if a field has been set.
func (o *Actor) HasReadmeSummary() bool {
	if o != nil && !IsNil(o.ReadmeSummary) {
		return true
	}

	return false
}

// SetReadmeSummary gets a reference to the given string and assigns it to the ReadmeSummary field.
func (o *Actor) SetReadmeSummary(v string) {
	o.ReadmeSummary = &v
}

// GetSeoTitle returns the SeoTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetSeoTitle() string {
	if o == nil || IsNil(o.SeoTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SeoTitle.Get()
}

// GetSeoTitleOk returns a tuple with the SeoTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetSeoTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeoTitle.Get(), o.SeoTitle.IsSet()
}

// HasSeoTitle returns a boolean if a field has been set.
func (o *Actor) HasSeoTitle() bool {
	if o != nil && o.SeoTitle.IsSet() {
		return true
	}

	return false
}

// SetSeoTitle gets a reference to the given NullableString and assigns it to the SeoTitle field.
func (o *Actor) SetSeoTitle(v string) {
	o.SeoTitle.Set(&v)
}
// SetSeoTitleNil sets the value for SeoTitle to be an explicit nil
func (o *Actor) SetSeoTitleNil() {
	o.SeoTitle.Set(nil)
}

// UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
func (o *Actor) UnsetSeoTitle() {
	o.SeoTitle.Unset()
}

// GetSeoDescription returns the SeoDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetSeoDescription() string {
	if o == nil || IsNil(o.SeoDescription.Get()) {
		var ret string
		return ret
	}
	return *o.SeoDescription.Get()
}

// GetSeoDescriptionOk returns a tuple with the SeoDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetSeoDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeoDescription.Get(), o.SeoDescription.IsSet()
}

// HasSeoDescription returns a boolean if a field has been set.
func (o *Actor) HasSeoDescription() bool {
	if o != nil && o.SeoDescription.IsSet() {
		return true
	}

	return false
}

// SetSeoDescription gets a reference to the given NullableString and assigns it to the SeoDescription field.
func (o *Actor) SetSeoDescription(v string) {
	o.SeoDescription.Set(&v)
}
// SetSeoDescriptionNil sets the value for SeoDescription to be an explicit nil
func (o *Actor) SetSeoDescriptionNil() {
	o.SeoDescription.Set(nil)
}

// UnsetSeoDescription ensures that no value is present for SeoDescription, not even an explicit nil
func (o *Actor) UnsetSeoDescription() {
	o.SeoDescription.Unset()
}

// GetPictureUrl returns the PictureUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetPictureUrl() string {
	if o == nil || IsNil(o.PictureUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PictureUrl.Get()
}

// GetPictureUrlOk returns a tuple with the PictureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetPictureUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PictureUrl.Get(), o.PictureUrl.IsSet()
}

// HasPictureUrl returns a boolean if a field has been set.
func (o *Actor) HasPictureUrl() bool {
	if o != nil && o.PictureUrl.IsSet() {
		return true
	}

	return false
}

// SetPictureUrl gets a reference to the given NullableString and assigns it to the PictureUrl field.
func (o *Actor) SetPictureUrl(v string) {
	o.PictureUrl.Set(&v)
}
// SetPictureUrlNil sets the value for PictureUrl to be an explicit nil
func (o *Actor) SetPictureUrlNil() {
	o.PictureUrl.Set(nil)
}

// UnsetPictureUrl ensures that no value is present for PictureUrl, not even an explicit nil
func (o *Actor) UnsetPictureUrl() {
	o.PictureUrl.Unset()
}

// GetStandbyUrl returns the StandbyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Actor) GetStandbyUrl() string {
	if o == nil || IsNil(o.StandbyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.StandbyUrl.Get()
}

// GetStandbyUrlOk returns a tuple with the StandbyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetStandbyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StandbyUrl.Get(), o.StandbyUrl.IsSet()
}

// HasStandbyUrl returns a boolean if a field has been set.
func (o *Actor) HasStandbyUrl() bool {
	if o != nil && o.StandbyUrl.IsSet() {
		return true
	}

	return false
}

// SetStandbyUrl gets a reference to the given NullableString and assigns it to the StandbyUrl field.
func (o *Actor) SetStandbyUrl(v string) {
	o.StandbyUrl.Set(&v)
}
// SetStandbyUrlNil sets the value for StandbyUrl to be an explicit nil
func (o *Actor) SetStandbyUrlNil() {
	o.StandbyUrl.Set(nil)
}

// UnsetStandbyUrl ensures that no value is present for StandbyUrl, not even an explicit nil
func (o *Actor) UnsetStandbyUrl() {
	o.StandbyUrl.Unset()
}

// GetNotice returns the Notice field value if set, zero value otherwise.
func (o *Actor) GetNotice() string {
	if o == nil || IsNil(o.Notice) {
		var ret string
		return ret
	}
	return *o.Notice
}

// GetNoticeOk returns a tuple with the Notice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetNoticeOk() (*string, bool) {
	if o == nil || IsNil(o.Notice) {
		return nil, false
	}
	return o.Notice, true
}

// HasNotice returns a boolean if a field has been set.
func (o *Actor) HasNotice() bool {
	if o != nil && !IsNil(o.Notice) {
		return true
	}

	return false
}

// SetNotice gets a reference to the given string and assigns it to the Notice field.
func (o *Actor) SetNotice(v string) {
	o.Notice = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Actor) GetCategories() []string {
	if o == nil || IsNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Actor) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *Actor) SetCategories(v []string) {
	o.Categories = v
}

// GetIsCritical returns the IsCritical field value if set, zero value otherwise.
func (o *Actor) GetIsCritical() bool {
	if o == nil || IsNil(o.IsCritical) {
		var ret bool
		return ret
	}
	return *o.IsCritical
}

// GetIsCriticalOk returns a tuple with the IsCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetIsCriticalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCritical) {
		return nil, false
	}
	return o.IsCritical, true
}

// HasIsCritical returns a boolean if a field has been set.
func (o *Actor) HasIsCritical() bool {
	if o != nil && !IsNil(o.IsCritical) {
		return true
	}

	return false
}

// SetIsCritical gets a reference to the given bool and assigns it to the IsCritical field.
func (o *Actor) SetIsCritical(v bool) {
	o.IsCritical = &v
}

// GetIsGeneric returns the IsGeneric field value if set, zero value otherwise.
func (o *Actor) GetIsGeneric() bool {
	if o == nil || IsNil(o.IsGeneric) {
		var ret bool
		return ret
	}
	return *o.IsGeneric
}

// GetIsGenericOk returns a tuple with the IsGeneric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetIsGenericOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGeneric) {
		return nil, false
	}
	return o.IsGeneric, true
}

// HasIsGeneric returns a boolean if a field has been set.
func (o *Actor) HasIsGeneric() bool {
	if o != nil && !IsNil(o.IsGeneric) {
		return true
	}

	return false
}

// SetIsGeneric gets a reference to the given bool and assigns it to the IsGeneric field.
func (o *Actor) SetIsGeneric(v bool) {
	o.IsGeneric = &v
}

// GetIsSourceCodeHidden returns the IsSourceCodeHidden field value if set, zero value otherwise.
func (o *Actor) GetIsSourceCodeHidden() bool {
	if o == nil || IsNil(o.IsSourceCodeHidden) {
		var ret bool
		return ret
	}
	return *o.IsSourceCodeHidden
}

// GetIsSourceCodeHiddenOk returns a tuple with the IsSourceCodeHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetIsSourceCodeHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSourceCodeHidden) {
		return nil, false
	}
	return o.IsSourceCodeHidden, true
}

// HasIsSourceCodeHidden returns a boolean if a field has been set.
func (o *Actor) HasIsSourceCodeHidden() bool {
	if o != nil && !IsNil(o.IsSourceCodeHidden) {
		return true
	}

	return false
}

// SetIsSourceCodeHidden gets a reference to the given bool and assigns it to the IsSourceCodeHidden field.
func (o *Actor) SetIsSourceCodeHidden(v bool) {
	o.IsSourceCodeHidden = &v
}

// GetHasNoDataset returns the HasNoDataset field value if set, zero value otherwise.
func (o *Actor) GetHasNoDataset() bool {
	if o == nil || IsNil(o.HasNoDataset) {
		var ret bool
		return ret
	}
	return *o.HasNoDataset
}

// GetHasNoDatasetOk returns a tuple with the HasNoDataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetHasNoDatasetOk() (*bool, bool) {
	if o == nil || IsNil(o.HasNoDataset) {
		return nil, false
	}
	return o.HasNoDataset, true
}

// HasHasNoDataset returns a boolean if a field has been set.
func (o *Actor) HasHasNoDataset() bool {
	if o != nil && !IsNil(o.HasNoDataset) {
		return true
	}

	return false
}

// SetHasNoDataset gets a reference to the given bool and assigns it to the HasNoDataset field.
func (o *Actor) SetHasNoDataset(v bool) {
	o.HasNoDataset = &v
}

func (o Actor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Actor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["userId"] = o.UserId
	toSerialize["name"] = o.Name
	toSerialize["username"] = o.Username
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.RestartOnError) {
		toSerialize["restartOnError"] = o.RestartOnError
	}
	toSerialize["isPublic"] = o.IsPublic
	if !IsNil(o.ActorPermissionLevel) {
		toSerialize["actorPermissionLevel"] = o.ActorPermissionLevel
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["stats"] = o.Stats
	toSerialize["versions"] = o.Versions
	if !IsNil(o.PricingInfos) {
		toSerialize["pricingInfos"] = o.PricingInfos
	}
	toSerialize["defaultRunOptions"] = o.DefaultRunOptions
	if o.ExampleRunInput.IsSet() {
		toSerialize["exampleRunInput"] = o.ExampleRunInput.Get()
	}
	if o.IsDeprecated.IsSet() {
		toSerialize["isDeprecated"] = o.IsDeprecated.Get()
	}
	if !IsNil(o.DeploymentKey) {
		toSerialize["deploymentKey"] = o.DeploymentKey
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TaggedBuilds != nil {
		toSerialize["taggedBuilds"] = o.TaggedBuilds
	}
	if o.ActorStandby.IsSet() {
		toSerialize["actorStandby"] = o.ActorStandby.Get()
	}
	if !IsNil(o.ReadmeSummary) {
		toSerialize["readmeSummary"] = o.ReadmeSummary
	}
	if o.SeoTitle.IsSet() {
		toSerialize["seoTitle"] = o.SeoTitle.Get()
	}
	if o.SeoDescription.IsSet() {
		toSerialize["seoDescription"] = o.SeoDescription.Get()
	}
	if o.PictureUrl.IsSet() {
		toSerialize["pictureUrl"] = o.PictureUrl.Get()
	}
	if o.StandbyUrl.IsSet() {
		toSerialize["standbyUrl"] = o.StandbyUrl.Get()
	}
	if !IsNil(o.Notice) {
		toSerialize["notice"] = o.Notice
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.IsCritical) {
		toSerialize["isCritical"] = o.IsCritical
	}
	if !IsNil(o.IsGeneric) {
		toSerialize["isGeneric"] = o.IsGeneric
	}
	if !IsNil(o.IsSourceCodeHidden) {
		toSerialize["isSourceCodeHidden"] = o.IsSourceCodeHidden
	}
	if !IsNil(o.HasNoDataset) {
		toSerialize["hasNoDataset"] = o.HasNoDataset
	}
	return toSerialize, nil
}

func (o *Actor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"userId",
		"name",
		"username",
		"isPublic",
		"createdAt",
		"modifiedAt",
		"stats",
		"versions",
		"defaultRunOptions",
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

	varActor := _Actor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActor)

	if err != nil {
		return err
	}

	*o = Actor(varActor)

	return err
}

type NullableActor struct {
	value *Actor
	isSet bool
}

func (v NullableActor) Get() *Actor {
	return v.value
}

func (v *NullableActor) Set(val *Actor) {
	v.value = val
	v.isSet = true
}

func (v NullableActor) IsSet() bool {
	return v.isSet
}

func (v *NullableActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActor(val *Actor) *NullableActor {
	return &NullableActor{value: val, isSet: true}
}

func (v NullableActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


