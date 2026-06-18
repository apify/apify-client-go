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

// checks if the ActorDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActorDefinition{}

// ActorDefinition The definition of the Actor, the full specification of this field can be found in [Apify docs](https://docs.apify.com/platform/actors/development/actor-definition/actor-json)
type ActorDefinition struct {
	// The Actor specification version that this Actor follows. This property must be set to 1.
	ActorSpecification *int32 `json:"actorSpecification,omitempty"`
	// The name of the Actor.
	Name *string `json:"name,omitempty"`
	// The version of the Actor, typically a dot-separated sequence of numbers (e.g., `0.1`, `1.0`, or `0.0.1`).
	Version *string `json:"version,omitempty" validate:"regexp=^[0-9]+(\\\\.[0-9]+)+$"`
	// The tag name to be applied to a successful build of the Actor. Defaults to 'latest' if not specified.
	BuildTag *string `json:"buildTag,omitempty"`
	// A map of environment variables to be used during local development and deployment.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	// The path to the Dockerfile used for building the Actor on the platform.
	Dockerfile *string `json:"dockerfile,omitempty"`
	// The path to the directory used as the Docker context when building the Actor.
	DockerContextDir *string `json:"dockerContextDir,omitempty"`
	// The path to the README file for the Actor.
	Readme *string `json:"readme,omitempty"`
	// The input schema object, the full specification can be found in [Apify docs](https://docs.apify.com/platform/actors/development/actor-definition/input-schema)
	Input map[string]interface{} `json:"input,omitempty"`
	// The path to the CHANGELOG file displayed in the Actor's information tab.
	Changelog *string `json:"changelog,omitempty"`
	Storages *ActorDefinitionStorages `json:"storages,omitempty"`
	DefaultMemoryMbytes *ActorDefinitionDefaultMemoryMbytes `json:"defaultMemoryMbytes,omitempty"`
	// Specifies the minimum amount of memory in megabytes required by the Actor.
	MinMemoryMbytes *int32 `json:"minMemoryMbytes,omitempty"`
	// Specifies the maximum amount of memory in megabytes required by the Actor.
	MaxMemoryMbytes *int32 `json:"maxMemoryMbytes,omitempty"`
	// Specifies whether Standby mode is enabled for the Actor.
	UsesStandbyMode *bool `json:"usesStandbyMode,omitempty"`
}

// NewActorDefinition instantiates a new ActorDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActorDefinition() *ActorDefinition {
	this := ActorDefinition{}
	return &this
}

// NewActorDefinitionWithDefaults instantiates a new ActorDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorDefinitionWithDefaults() *ActorDefinition {
	this := ActorDefinition{}
	return &this
}

// GetActorSpecification returns the ActorSpecification field value if set, zero value otherwise.
func (o *ActorDefinition) GetActorSpecification() int32 {
	if o == nil || IsNil(o.ActorSpecification) {
		var ret int32
		return ret
	}
	return *o.ActorSpecification
}

// GetActorSpecificationOk returns a tuple with the ActorSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetActorSpecificationOk() (*int32, bool) {
	if o == nil || IsNil(o.ActorSpecification) {
		return nil, false
	}
	return o.ActorSpecification, true
}

// HasActorSpecification returns a boolean if a field has been set.
func (o *ActorDefinition) HasActorSpecification() bool {
	if o != nil && !IsNil(o.ActorSpecification) {
		return true
	}

	return false
}

// SetActorSpecification gets a reference to the given int32 and assigns it to the ActorSpecification field.
func (o *ActorDefinition) SetActorSpecification(v int32) {
	o.ActorSpecification = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActorDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActorDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActorDefinition) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ActorDefinition) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ActorDefinition) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ActorDefinition) SetVersion(v string) {
	o.Version = &v
}

// GetBuildTag returns the BuildTag field value if set, zero value otherwise.
func (o *ActorDefinition) GetBuildTag() string {
	if o == nil || IsNil(o.BuildTag) {
		var ret string
		return ret
	}
	return *o.BuildTag
}

// GetBuildTagOk returns a tuple with the BuildTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetBuildTagOk() (*string, bool) {
	if o == nil || IsNil(o.BuildTag) {
		return nil, false
	}
	return o.BuildTag, true
}

// HasBuildTag returns a boolean if a field has been set.
func (o *ActorDefinition) HasBuildTag() bool {
	if o != nil && !IsNil(o.BuildTag) {
		return true
	}

	return false
}

// SetBuildTag gets a reference to the given string and assigns it to the BuildTag field.
func (o *ActorDefinition) SetBuildTag(v string) {
	o.BuildTag = &v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *ActorDefinition) GetEnvironmentVariables() map[string]string {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret map[string]string
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetEnvironmentVariablesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return map[string]string{}, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *ActorDefinition) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given map[string]string and assigns it to the EnvironmentVariables field.
func (o *ActorDefinition) SetEnvironmentVariables(v map[string]string) {
	o.EnvironmentVariables = v
}

// GetDockerfile returns the Dockerfile field value if set, zero value otherwise.
func (o *ActorDefinition) GetDockerfile() string {
	if o == nil || IsNil(o.Dockerfile) {
		var ret string
		return ret
	}
	return *o.Dockerfile
}

// GetDockerfileOk returns a tuple with the Dockerfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetDockerfileOk() (*string, bool) {
	if o == nil || IsNil(o.Dockerfile) {
		return nil, false
	}
	return o.Dockerfile, true
}

// HasDockerfile returns a boolean if a field has been set.
func (o *ActorDefinition) HasDockerfile() bool {
	if o != nil && !IsNil(o.Dockerfile) {
		return true
	}

	return false
}

// SetDockerfile gets a reference to the given string and assigns it to the Dockerfile field.
func (o *ActorDefinition) SetDockerfile(v string) {
	o.Dockerfile = &v
}

// GetDockerContextDir returns the DockerContextDir field value if set, zero value otherwise.
func (o *ActorDefinition) GetDockerContextDir() string {
	if o == nil || IsNil(o.DockerContextDir) {
		var ret string
		return ret
	}
	return *o.DockerContextDir
}

// GetDockerContextDirOk returns a tuple with the DockerContextDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetDockerContextDirOk() (*string, bool) {
	if o == nil || IsNil(o.DockerContextDir) {
		return nil, false
	}
	return o.DockerContextDir, true
}

// HasDockerContextDir returns a boolean if a field has been set.
func (o *ActorDefinition) HasDockerContextDir() bool {
	if o != nil && !IsNil(o.DockerContextDir) {
		return true
	}

	return false
}

// SetDockerContextDir gets a reference to the given string and assigns it to the DockerContextDir field.
func (o *ActorDefinition) SetDockerContextDir(v string) {
	o.DockerContextDir = &v
}

// GetReadme returns the Readme field value if set, zero value otherwise.
func (o *ActorDefinition) GetReadme() string {
	if o == nil || IsNil(o.Readme) {
		var ret string
		return ret
	}
	return *o.Readme
}

// GetReadmeOk returns a tuple with the Readme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetReadmeOk() (*string, bool) {
	if o == nil || IsNil(o.Readme) {
		return nil, false
	}
	return o.Readme, true
}

// HasReadme returns a boolean if a field has been set.
func (o *ActorDefinition) HasReadme() bool {
	if o != nil && !IsNil(o.Readme) {
		return true
	}

	return false
}

// SetReadme gets a reference to the given string and assigns it to the Readme field.
func (o *ActorDefinition) SetReadme(v string) {
	o.Readme = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ActorDefinition) GetInput() map[string]interface{} {
	if o == nil || IsNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ActorDefinition) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *ActorDefinition) SetInput(v map[string]interface{}) {
	o.Input = v
}

// GetChangelog returns the Changelog field value if set, zero value otherwise.
func (o *ActorDefinition) GetChangelog() string {
	if o == nil || IsNil(o.Changelog) {
		var ret string
		return ret
	}
	return *o.Changelog
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetChangelogOk() (*string, bool) {
	if o == nil || IsNil(o.Changelog) {
		return nil, false
	}
	return o.Changelog, true
}

// HasChangelog returns a boolean if a field has been set.
func (o *ActorDefinition) HasChangelog() bool {
	if o != nil && !IsNil(o.Changelog) {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given string and assigns it to the Changelog field.
func (o *ActorDefinition) SetChangelog(v string) {
	o.Changelog = &v
}

// GetStorages returns the Storages field value if set, zero value otherwise.
func (o *ActorDefinition) GetStorages() ActorDefinitionStorages {
	if o == nil || IsNil(o.Storages) {
		var ret ActorDefinitionStorages
		return ret
	}
	return *o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetStoragesOk() (*ActorDefinitionStorages, bool) {
	if o == nil || IsNil(o.Storages) {
		return nil, false
	}
	return o.Storages, true
}

// HasStorages returns a boolean if a field has been set.
func (o *ActorDefinition) HasStorages() bool {
	if o != nil && !IsNil(o.Storages) {
		return true
	}

	return false
}

// SetStorages gets a reference to the given ActorDefinitionStorages and assigns it to the Storages field.
func (o *ActorDefinition) SetStorages(v ActorDefinitionStorages) {
	o.Storages = &v
}

// GetDefaultMemoryMbytes returns the DefaultMemoryMbytes field value if set, zero value otherwise.
func (o *ActorDefinition) GetDefaultMemoryMbytes() ActorDefinitionDefaultMemoryMbytes {
	if o == nil || IsNil(o.DefaultMemoryMbytes) {
		var ret ActorDefinitionDefaultMemoryMbytes
		return ret
	}
	return *o.DefaultMemoryMbytes
}

// GetDefaultMemoryMbytesOk returns a tuple with the DefaultMemoryMbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetDefaultMemoryMbytesOk() (*ActorDefinitionDefaultMemoryMbytes, bool) {
	if o == nil || IsNil(o.DefaultMemoryMbytes) {
		return nil, false
	}
	return o.DefaultMemoryMbytes, true
}

// HasDefaultMemoryMbytes returns a boolean if a field has been set.
func (o *ActorDefinition) HasDefaultMemoryMbytes() bool {
	if o != nil && !IsNil(o.DefaultMemoryMbytes) {
		return true
	}

	return false
}

// SetDefaultMemoryMbytes gets a reference to the given ActorDefinitionDefaultMemoryMbytes and assigns it to the DefaultMemoryMbytes field.
func (o *ActorDefinition) SetDefaultMemoryMbytes(v ActorDefinitionDefaultMemoryMbytes) {
	o.DefaultMemoryMbytes = &v
}

// GetMinMemoryMbytes returns the MinMemoryMbytes field value if set, zero value otherwise.
func (o *ActorDefinition) GetMinMemoryMbytes() int32 {
	if o == nil || IsNil(o.MinMemoryMbytes) {
		var ret int32
		return ret
	}
	return *o.MinMemoryMbytes
}

// GetMinMemoryMbytesOk returns a tuple with the MinMemoryMbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetMinMemoryMbytesOk() (*int32, bool) {
	if o == nil || IsNil(o.MinMemoryMbytes) {
		return nil, false
	}
	return o.MinMemoryMbytes, true
}

// HasMinMemoryMbytes returns a boolean if a field has been set.
func (o *ActorDefinition) HasMinMemoryMbytes() bool {
	if o != nil && !IsNil(o.MinMemoryMbytes) {
		return true
	}

	return false
}

// SetMinMemoryMbytes gets a reference to the given int32 and assigns it to the MinMemoryMbytes field.
func (o *ActorDefinition) SetMinMemoryMbytes(v int32) {
	o.MinMemoryMbytes = &v
}

// GetMaxMemoryMbytes returns the MaxMemoryMbytes field value if set, zero value otherwise.
func (o *ActorDefinition) GetMaxMemoryMbytes() int32 {
	if o == nil || IsNil(o.MaxMemoryMbytes) {
		var ret int32
		return ret
	}
	return *o.MaxMemoryMbytes
}

// GetMaxMemoryMbytesOk returns a tuple with the MaxMemoryMbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetMaxMemoryMbytesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMemoryMbytes) {
		return nil, false
	}
	return o.MaxMemoryMbytes, true
}

// HasMaxMemoryMbytes returns a boolean if a field has been set.
func (o *ActorDefinition) HasMaxMemoryMbytes() bool {
	if o != nil && !IsNil(o.MaxMemoryMbytes) {
		return true
	}

	return false
}

// SetMaxMemoryMbytes gets a reference to the given int32 and assigns it to the MaxMemoryMbytes field.
func (o *ActorDefinition) SetMaxMemoryMbytes(v int32) {
	o.MaxMemoryMbytes = &v
}

// GetUsesStandbyMode returns the UsesStandbyMode field value if set, zero value otherwise.
func (o *ActorDefinition) GetUsesStandbyMode() bool {
	if o == nil || IsNil(o.UsesStandbyMode) {
		var ret bool
		return ret
	}
	return *o.UsesStandbyMode
}

// GetUsesStandbyModeOk returns a tuple with the UsesStandbyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorDefinition) GetUsesStandbyModeOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesStandbyMode) {
		return nil, false
	}
	return o.UsesStandbyMode, true
}

// HasUsesStandbyMode returns a boolean if a field has been set.
func (o *ActorDefinition) HasUsesStandbyMode() bool {
	if o != nil && !IsNil(o.UsesStandbyMode) {
		return true
	}

	return false
}

// SetUsesStandbyMode gets a reference to the given bool and assigns it to the UsesStandbyMode field.
func (o *ActorDefinition) SetUsesStandbyMode(v bool) {
	o.UsesStandbyMode = &v
}

func (o ActorDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActorDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActorSpecification) {
		toSerialize["actorSpecification"] = o.ActorSpecification
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.BuildTag) {
		toSerialize["buildTag"] = o.BuildTag
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if !IsNil(o.Dockerfile) {
		toSerialize["dockerfile"] = o.Dockerfile
	}
	if !IsNil(o.DockerContextDir) {
		toSerialize["dockerContextDir"] = o.DockerContextDir
	}
	if !IsNil(o.Readme) {
		toSerialize["readme"] = o.Readme
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Changelog) {
		toSerialize["changelog"] = o.Changelog
	}
	if !IsNil(o.Storages) {
		toSerialize["storages"] = o.Storages
	}
	if !IsNil(o.DefaultMemoryMbytes) {
		toSerialize["defaultMemoryMbytes"] = o.DefaultMemoryMbytes
	}
	if !IsNil(o.MinMemoryMbytes) {
		toSerialize["minMemoryMbytes"] = o.MinMemoryMbytes
	}
	if !IsNil(o.MaxMemoryMbytes) {
		toSerialize["maxMemoryMbytes"] = o.MaxMemoryMbytes
	}
	if !IsNil(o.UsesStandbyMode) {
		toSerialize["usesStandbyMode"] = o.UsesStandbyMode
	}
	return toSerialize, nil
}

type NullableActorDefinition struct {
	value *ActorDefinition
	isSet bool
}

func (v NullableActorDefinition) Get() *ActorDefinition {
	return v.value
}

func (v *NullableActorDefinition) Set(val *ActorDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableActorDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableActorDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActorDefinition(val *ActorDefinition) *NullableActorDefinition {
	return &NullableActorDefinition{value: val, isSet: true}
}

func (v NullableActorDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActorDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


