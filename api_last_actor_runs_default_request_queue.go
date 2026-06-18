/*
Apify API

 The Apify API (version 2) provides programmatic access to the [Apify platform](https://docs.apify.com). The API is organized around [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) HTTP endpoints.  You can download the complete OpenAPI schema of Apify API in the [YAML](http://docs.apify.com/api/openapi.yaml) or [JSON](http://docs.apify.com/api/openapi.json) formats. The source code is also available on [GitHub](https://github.com/apify/apify-docs/tree/master/apify-api/openapi).  All requests and responses (including errors) are encoded in [JSON](http://www.json.org/) format with UTF-8 encoding, with a few exceptions that are explicitly described in the reference.  - To access the API using [Node.js](https://nodejs.org/en/), we recommend the [`apify-client`](https://docs.apify.com/api/client/js) [NPM package](https://www.npmjs.com/package/apify-client). - To access the API using [Python](https://www.python.org/), we recommend the [`apify-client`](https://docs.apify.com/api/client/python) [PyPI package](https://pypi.org/project/apify-client/).  The clients' functions correspond to the API endpoints and have the same parameters. This simplifies development of apps that depend on the Apify platform.  :::note Important Request Details  - `Content-Type` header: For requests with a JSON body, you must include the `Content-Type: application/json` header.  - Method override: You can override the HTTP method using the `method` query parameter. This is useful for clients that can only send `GET` requests. For example, to call a `POST` endpoint, append `?method=POST` to the URL of your `GET` request.  :::  ## Authentication <span id=\"/introduction/authentication\"></span>  **You can find your API token on the [Integrations](https://console.apify.com/settings/integrations) page in the Apify Console.**  To use your token in a request, either:  - Add the token to your request's `Authorization` header as `Bearer <token>`. E.g., `Authorization: Bearer xxxxxxx`. [More info](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization). (Recommended). - Add it as the `token` parameter to your request URL. (Less secure).  Using your token in the request header is more secure than using it as a URL parameter because URLs are often stored in browser history and server logs. This creates a chance for someone unauthorized to access your API token.  **Never share your API token or password with untrusted parties!**  For more information, see our [integrations](https://docs.apify.com/platform/integrations) documentation.  ### Agentic payments  AI agents can authenticate and pay for Actor runs without an Apify account using agentic payments. Instead of an API token, the request carries a payment credential that both authorizes and pays for the call. Apify supports the [x402 protocol](https://docs.apify.com/platform/integrations/x402) (`PAYMENT-SIGNATURE` header) and [Skyfire](https://docs.apify.com/platform/integrations/skyfire) (`skyfire-pay-id` header).  ## Basic usage <span id=\"/introduction/basic-usage\"></span>  To run an Actor, send a POST request to the [Run Actor](#/reference/actors/run-collection/run-actor) endpoint using either the Actor ID code (e.g. `vKg4IjxZbEYTYeW8T`) or its name (e.g. `janedoe~my-actor`):  `https://api.apify.com/v2/actors/[actor_id]/runs`  If the Actor is not runnable anonymously, you will receive a 401 or 403 [response code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status). This means you need to add your [secret API token](https://console.apify.com/account#/integrations) to the request's `Authorization` header ([recommended](#/introduction/authentication)) or as a URL query parameter `?token=[your_token]` (less secure).  Optionally, you can include the query parameters described in the [Run Actor](#/reference/actors/run-collection/run-actor) section to customize your run.  If you're using Node.js, the best way to run an Actor is using the `Apify.call()` method from the [Apify SDK](https://sdk.apify.com/docs/api/apify#apifycallactid-input-options). It runs the Actor using the account you are currently logged into (determined by the [secret API token](https://console.apify.com/account#/integrations)). The result is an [Actor run object](https://sdk.apify.com/docs/typedefs/actor-run) and its output (if any).  A typical workflow is as follows:  1. Run an Actor or task using the [Run Actor](#/reference/actors/run-collection/run-actor) or [Run task](#/reference/actor-tasks/run-collection/run-task) API endpoints. 2. Monitor the Actor run by periodically polling its progress using the [Get run](#/reference/actor-runs/run-object-and-its-storages/get-run) API endpoint. 3. Fetch the results from the [Get items](#/reference/datasets/item-collection/get-items) API endpoint using the `defaultDatasetId`, which you receive in the Run request response. Additional data may be stored in a key-value store. You can fetch them from the [Get record](#/reference/key-value-stores/record/get-record) API endpoint using the `defaultKeyValueStoreId` and the store's `key`.  **Note**: Instead of periodic polling, you can also run your [Actor](#/reference/actors/run-actor-synchronously) or [task](#/reference/actor-tasks/runs-collection/run-task-synchronously) synchronously. This will ensure that the request waits for 300 seconds (5 minutes) for the run to finish and returns its output. If the run takes longer, the request will time out and throw an error.  ## Legacy `/v2/acts/` URL prefix <span id=\"/introduction/legacy-acts-prefix\"></span>  The `/v2/acts/` prefix is deprecated but still fully functional, and  such endpoint routes to the same handler as its `/v2/actors/...` counterpart.  New integrations should use the canonical /v2/actors/ prefix,  but existing clients keep working without changes.  ## Response structure <span id=\"/introduction/response-structure\"></span>  Most API endpoints return a JSON object with the `data` property:  ``` {     \"data\": {         ...     } } ```  However, there are a few explicitly described exceptions, such as [Get dataset items](#/reference/datasets/item-collection/get-items) or Key-value store [Get record](#/reference/key-value-stores/record/get-record) API endpoints, which return data in other formats. In case of an error, the response has the HTTP status code in the range of 4xx or 5xx and the `data` property is replaced with `error`. For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  See [Errors](#/introduction/errors) for more details.  ## Pagination <span id=\"/introduction/pagination\"></span>  All API endpoints that return a list of records (e.g. [Get list of Actors](#/reference/actors/actor-collection/get-list-of-actors)) enforce pagination in order to limit the size of their responses.  Most of these API endpoints are paginated using the `offset` and `limit` query parameters. The only exception is [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys), which is paginated using the `exclusiveStartKey` query parameter.  **IMPORTANT**: Each API endpoint that supports pagination enforces a certain maximum value for the `limit` parameter, in order to reduce the load on Apify servers. The maximum limit could change in future so you should never rely on a specific value and check the responses of these API endpoints.  ### Using offset <span id=\"/introduction/pagination/using-offset\"></span>  Most API endpoints that return a list of records enable pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number of items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>Skips a number of items from the beginning of the list, e.g. <code>offset=100</code>.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td>     By default, items are sorted in the order in which they were created or added to the list.     This feature is useful when fetching all the items, because it ensures that items     created after the client started the pagination will not be skipped.     If you specify the <code>desc=1</code> parameter, the items will be returned in the reverse order,     i.e. from the newest to the oldest items.     </td>   </tr> </table>  The response of these API endpoints is always a JSON object with the following structure:  ``` {     \"data\": {         \"total\": 2560,         \"offset\": 250,         \"limit\": 1000,         \"count\": 1000,         \"desc\": false,         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>total</code></td>     <td>The total number of items available in the list.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>The number of items that were skipped at the start.     This is equal to the <code>offset</code> query parameter if it was provided, otherwise it is <code>0</code>.</td>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular API endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>count</code></td>     <td>The actual number of items returned in the HTTP response.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td><code>true</code> if data were requested in descending order and <code>false</code> otherwise.</td>   </tr>   <tr>     <td><code>items</code></td>     <td>An array of requested items.</td>   </tr> </table>  ### Using key <span id=\"/introduction/pagination/using-key\"></span>  The records in the [key-value store](https://docs.apify.com/platform/storage/key-value-store) are not ordered based on numerical indexes, but rather by their keys in the UTF-8 binary order. Therefore the [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys) API endpoint only supports pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>Skips all records with keys up to the given key including the given key,     in the UTF-8 binary order.</td>   </tr> </table>  The response of the API endpoint is always a JSON object with following structure:  ``` {     \"data\": {         \"limit\": 1000,         \"isTruncated\": true,         \"exclusiveStartKey\": \"my-key\",         \"nextExclusiveStartKey\": \"some-other-key\",         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>isTruncated</code></td>     <td><code>true</code> if there are more items left to be queried. Otherwise <code>false</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>The last key that was skipped at the start. Is `null` for the first page.</td>   </tr>   <tr>     <td><code>nextExclusiveStartKey</code></td>     <td>The value for the <code>exclusiveStartKey</code> parameter to query the next page of items.</td>   </tr> </table>  ## Errors <span id=\"/introduction/errors\"></span>  The Apify API uses common HTTP status codes: `2xx` range for success, `4xx` range for errors caused by the caller (invalid requests) and `5xx` range for server errors (these are rare). Each error response contains a JSON object defining the `error` property, which is an object with the `type` and `message` properties that contain the error code and a human-readable error description, respectively.  For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  Here is the table of the most common errors that can occur for many API endpoints:  <table>   <tr>     <th>status</th>     <th>type</th>     <th>message</th>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-request</code></td>     <td>POST data must be a JSON object</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-value</code></td>     <td>Invalid value provided: Comments required</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-record-key</code></td>     <td>Record key contains invalid character</td>   </tr>   <tr>     <td><code>401</code></td>     <td><code>token-not-provided</code></td>     <td>Authentication token was not provided</td>   </tr>   <tr>     <td><code>404</code></td>     <td><code>record-not-found</code></td>     <td>Store was not found</td>   </tr>   <tr>     <td><code>429</code></td>     <td><code>rate-limit-exceeded</code></td>     <td>You have exceeded the rate limit of ... requests per second</td>   </tr>   <tr>     <td><code>405</code></td>     <td><code>method-not-allowed</code></td>     <td>This API endpoint can only be accessed using the following HTTP methods: OPTIONS, POST</td>   </tr> </table>  ## Rate limiting <span id=\"/introduction/rate-limiting\"></span>  All API endpoints limit the rate of requests in order to prevent overloading of Apify servers by misbehaving clients.  There are two kinds of rate limits - a global rate limit and a per-resource rate limit.  ### Global rate limit <span id=\"/introduction/rate-limiting/global-rate-limit\"></span>  The global rate limit is set to _250 000 requests per minute_. For [authenticated](#/introduction/authentication) requests, it is counted per user, and for unauthenticated requests, it is counted per IP address.  ### Per-resource rate limit <span id=\"/introduction/rate-limiting/per-resource-rate-limit\"></span>  The default per-resource rate limit is _60 requests per second per resource_, which in this context means a single Actor, a single Actor run, a single dataset, single key-value store etc. The default rate limit is applied to every API endpoint except a few select ones, which have higher rate limits. Each API endpoint returns its rate limit in `X-RateLimit-Limit` header.  These endpoints have a rate limit of _200 requests per second per resource_:  * CRUD ([get](#/reference/key-value-stores/record/get-record),   [put](#/reference/key-value-stores/record/put-record),   [delete](#/reference/key-value-stores/record/delete-record))   operations on key-value store records  These endpoints have a rate limit of _400 requests per second per resource_: * [Run Actor](#/reference/actors/run-collection/run-actor) * [Run Actor task asynchronously](#/reference/actor-tasks/runs-collection/run-task-asynchronously) * [Run Actor task synchronously](#/reference/actor-tasks/runs-collection/run-task-synchronously) * [Metamorph Actor run](#/reference/actors/metamorph-run/metamorph-run) * [Push items](#/reference/datasets/item-collection/put-items) to dataset * CRUD   ([add](#/reference/request-queues/request-collection/add-request),   [get](#/reference/request-queues/request-collection/get-request),   [update](#/reference/request-queues/request-collection/update-request),   [delete](#/reference/request-queues/request-collection/delete-request))   operations on requests in request queues  ### Rate limit exceeded errors <span id=\"/introduction/rate-limiting/rate-limit-exceeded-errors\"></span>  If the client is sending too many requests, the API endpoints respond with the HTTP status code `429 Too Many Requests` and the following body:  ``` {     \"error\": {         \"type\": \"rate-limit-exceeded\",         \"message\": \"You have exceeded the rate limit of ... requests per second\"     } } ```  ### Retrying rate-limited requests with exponential backoff <span id=\"/introduction/rate-limiting/retrying-rate-limited-requests-with-exponential-backoff\"></span>  If the client receives the rate limit error, it should wait a certain period of time and then retry the request. If the error happens again, the client should double the wait period and retry the request, and so on. This algorithm is known as _exponential backoff_ and it can be described using the following pseudo-code:  1. Define a variable `DELAY=500` 2. Send the HTTP request to the API endpoint 3. If the response has status code not equal to `429` then you are done. Otherwise:    * Wait for a period of time chosen randomly from the interval `DELAY` to `2*DELAY` milliseconds    * Double the future wait period by setting `DELAY = 2*DELAY`    * Continue with step 2  If all requests sent by the client implement the above steps, the client will automatically use the maximum available bandwidth for its requests.  Note that the Apify API clients [for JavaScript](https://docs.apify.com/api/client/js) and [for Python](https://docs.apify.com/api/client/python) use the exponential backoff algorithm transparently, so that you do not need to worry about it.  ## Referring to resources <span id=\"/introduction/referring-to-resources\"></span>  There are three main ways to refer to a resource you're accessing via API.  - the resource ID (e.g. `iKkPcIgVvwmztduf8`) - `username~resourcename` - when using this access method, you will need to use your API token, and access will only work if you have the correct permissions. - `~resourcename` - for this, you need to use an API token, and the `resourcename` refers to a resource in the API token owner's account. 

API version: v2-2026-06-16T064758Z
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apifyclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// LastActorRunsDefaultRequestQueueAPIService LastActorRunsDefaultRequestQueueAPI service
type LastActorRunsDefaultRequestQueueAPIService service

type ApiActRunsLastRequestQueueDeleteRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	status *string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueDeleteRequest) Status(status string) ApiActRunsLastRequestQueueDeleteRequest {
	r.status = &status
	return r
}

func (r ApiActRunsLastRequestQueueDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueDeleteExecute(r)
}

/*
ActRunsLastRequestQueueDelete Delete last run's default request queue

Deletes the default request queue associated with the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Delete request queue](/api/v2/request-queue-delete) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueDeleteRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueDelete(ctx context.Context, actorId string) ApiActRunsLastRequestQueueDeleteRequest {
	return ApiActRunsLastRequestQueueDeleteRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueDeleteExecute(r ApiActRunsLastRequestQueueDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueGetRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	status *string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueGetRequest) Status(status string) ApiActRunsLastRequestQueueGetRequest {
	r.status = &status
	return r
}

func (r ApiActRunsLastRequestQueueGetRequest) Execute() (*RequestQueueResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueGetExecute(r)
}

/*
ActRunsLastRequestQueueGet Get last run's default request queue

Returns the default request queue associated with the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Get request queue](/api/v2/request-queue-get) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueGetRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueGet(ctx context.Context, actorId string) ApiActRunsLastRequestQueueGetRequest {
	return ApiActRunsLastRequestQueueGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return RequestQueueResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueGetExecute(r ApiActRunsLastRequestQueueGetRequest) (*RequestQueueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RequestQueueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueHeadGetRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	status *string
	limit *float64
	clientKey *string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueHeadGetRequest) Status(status string) ApiActRunsLastRequestQueueHeadGetRequest {
	r.status = &status
	return r
}

// How many items from queue should be returned.
func (r ApiActRunsLastRequestQueueHeadGetRequest) Limit(limit float64) ApiActRunsLastRequestQueueHeadGetRequest {
	r.limit = &limit
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueHeadGetRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueHeadGetRequest {
	r.clientKey = &clientKey
	return r
}

func (r ApiActRunsLastRequestQueueHeadGetRequest) Execute() (*HeadResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueHeadGetExecute(r)
}

/*
ActRunsLastRequestQueueHeadGet Get last run's default request queue head

Returns the given number of first requests from the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Get head](/api/v2/request-queue-head-get) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueHeadGetRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueHeadGet(ctx context.Context, actorId string) ApiActRunsLastRequestQueueHeadGetRequest {
	return ApiActRunsLastRequestQueueHeadGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return HeadResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueHeadGetExecute(r ApiActRunsLastRequestQueueHeadGetRequest) (*HeadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HeadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueHeadGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/head"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueHeadLockPostRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	lockSecs *float64
	status *string
	limit *float64
	clientKey *string
}

// How long the requests will be locked for (in seconds).
func (r ApiActRunsLastRequestQueueHeadLockPostRequest) LockSecs(lockSecs float64) ApiActRunsLastRequestQueueHeadLockPostRequest {
	r.lockSecs = &lockSecs
	return r
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueHeadLockPostRequest) Status(status string) ApiActRunsLastRequestQueueHeadLockPostRequest {
	r.status = &status
	return r
}

// How many items from the queue should be returned.
func (r ApiActRunsLastRequestQueueHeadLockPostRequest) Limit(limit float64) ApiActRunsLastRequestQueueHeadLockPostRequest {
	r.limit = &limit
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueHeadLockPostRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueHeadLockPostRequest {
	r.clientKey = &clientKey
	return r
}

func (r ApiActRunsLastRequestQueueHeadLockPostRequest) Execute() (*HeadAndLockResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueHeadLockPostExecute(r)
}

/*
ActRunsLastRequestQueueHeadLockPost Get and lock last run's default request queue head

Returns the given number of first requests from the default request queue of the last Actor run
and locks them for the given time.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Get head and lock](/api/v2/request-queue-head-lock-post) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueHeadLockPostRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueHeadLockPost(ctx context.Context, actorId string) ApiActRunsLastRequestQueueHeadLockPostRequest {
	return ApiActRunsLastRequestQueueHeadLockPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return HeadAndLockResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueHeadLockPostExecute(r ApiActRunsLastRequestQueueHeadLockPostRequest) (*HeadAndLockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HeadAndLockResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueHeadLockPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/head/lock"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lockSecs == nil {
		return localVarReturnValue, nil, reportError("lockSecs is required and must be specified")
	}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "lockSecs", r.lockSecs, "form", "")
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueuePutRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	updateRequestQueueRequest *UpdateRequestQueueRequest
	status *string
}

// 
func (r ApiActRunsLastRequestQueuePutRequest) UpdateRequestQueueRequest(updateRequestQueueRequest UpdateRequestQueueRequest) ApiActRunsLastRequestQueuePutRequest {
	r.updateRequestQueueRequest = &updateRequestQueueRequest
	return r
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueuePutRequest) Status(status string) ApiActRunsLastRequestQueuePutRequest {
	r.status = &status
	return r
}

func (r ApiActRunsLastRequestQueuePutRequest) Execute() (*RequestQueueResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueuePutExecute(r)
}

/*
ActRunsLastRequestQueuePut Update last run's default request queue

Updates the default request queue associated with the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Update request queue](/api/v2/request-queue-put) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueuePutRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueuePut(ctx context.Context, actorId string) ApiActRunsLastRequestQueuePutRequest {
	return ApiActRunsLastRequestQueuePutRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return RequestQueueResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueuePutExecute(r ApiActRunsLastRequestQueuePutRequest) (*RequestQueueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RequestQueueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueuePut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRequestQueueRequest == nil {
		return localVarReturnValue, nil, reportError("updateRequestQueueRequest is required and must be specified")
	}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateRequestQueueRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestDeleteRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	requestId string
	status *string
	clientKey *string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestDeleteRequest) Status(status string) ApiActRunsLastRequestQueueRequestDeleteRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestDeleteRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestDeleteRequest {
	r.clientKey = &clientKey
	return r
}

func (r ApiActRunsLastRequestQueueRequestDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestDeleteExecute(r)
}

/*
ActRunsLastRequestQueueRequestDelete Delete request from last run's default request queue

Deletes a request from the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Delete request](/api/v2/request-queue-request-delete) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param requestId Request ID.
 @return ApiActRunsLastRequestQueueRequestDeleteRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestDelete(ctx context.Context, actorId string, requestId string) ApiActRunsLastRequestQueueRequestDeleteRequest {
	return ApiActRunsLastRequestQueueRequestDeleteRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		requestId: requestId,
	}
}

// Execute executes the request
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestDeleteExecute(r ApiActRunsLastRequestQueueRequestDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestGetRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	requestId string
	status *string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestGetRequest) Status(status string) ApiActRunsLastRequestQueueRequestGetRequest {
	r.status = &status
	return r
}

func (r ApiActRunsLastRequestQueueRequestGetRequest) Execute() (*RequestResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestGetExecute(r)
}

/*
ActRunsLastRequestQueueRequestGet Get request from last run's default request queue

Returns a request from the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Get request](/api/v2/request-queue-request-get) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param requestId Request ID.
 @return ApiActRunsLastRequestQueueRequestGetRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestGet(ctx context.Context, actorId string, requestId string) ApiActRunsLastRequestQueueRequestGetRequest {
	return ApiActRunsLastRequestQueueRequestGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		requestId: requestId,
	}
}

// Execute executes the request
//  @return RequestResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestGetExecute(r ApiActRunsLastRequestQueueRequestGetRequest) (*RequestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RequestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestLockDeleteRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	requestId string
	status *string
	clientKey *string
	forefront *string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestLockDeleteRequest) Status(status string) ApiActRunsLastRequestQueueRequestLockDeleteRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestLockDeleteRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestLockDeleteRequest {
	r.clientKey = &clientKey
	return r
}

// Determines if request should be added to the head of the queue or to the end after lock was removed. 
func (r ApiActRunsLastRequestQueueRequestLockDeleteRequest) Forefront(forefront string) ApiActRunsLastRequestQueueRequestLockDeleteRequest {
	r.forefront = &forefront
	return r
}

func (r ApiActRunsLastRequestQueueRequestLockDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestLockDeleteExecute(r)
}

/*
ActRunsLastRequestQueueRequestLockDelete Delete lock on request in last run's default request queue

Deletes a request lock in the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Delete request lock](/api/v2/request-queue-request-lock-delete) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param requestId Request ID.
 @return ApiActRunsLastRequestQueueRequestLockDeleteRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestLockDelete(ctx context.Context, actorId string, requestId string) ApiActRunsLastRequestQueueRequestLockDeleteRequest {
	return ApiActRunsLastRequestQueueRequestLockDeleteRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		requestId: requestId,
	}
}

// Execute executes the request
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestLockDeleteExecute(r ApiActRunsLastRequestQueueRequestLockDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestLockDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}/lock"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	if r.forefront != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forefront", r.forefront, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestLockPutRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	requestId string
	lockSecs *float64
	status *string
	clientKey *string
	forefront *string
}

// How long the requests will be locked for (in seconds).
func (r ApiActRunsLastRequestQueueRequestLockPutRequest) LockSecs(lockSecs float64) ApiActRunsLastRequestQueueRequestLockPutRequest {
	r.lockSecs = &lockSecs
	return r
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestLockPutRequest) Status(status string) ApiActRunsLastRequestQueueRequestLockPutRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestLockPutRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestLockPutRequest {
	r.clientKey = &clientKey
	return r
}

// Determines if request should be added to the head of the queue or to the end after lock expires. 
func (r ApiActRunsLastRequestQueueRequestLockPutRequest) Forefront(forefront string) ApiActRunsLastRequestQueueRequestLockPutRequest {
	r.forefront = &forefront
	return r
}

func (r ApiActRunsLastRequestQueueRequestLockPutRequest) Execute() (*ProlongRequestLockResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestLockPutExecute(r)
}

/*
ActRunsLastRequestQueueRequestLockPut Prolong lock on request in last run's default request queue

Prolongs a request lock in the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Prolong request lock](/api/v2/request-queue-request-lock-put) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param requestId Request ID.
 @return ApiActRunsLastRequestQueueRequestLockPutRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestLockPut(ctx context.Context, actorId string, requestId string) ApiActRunsLastRequestQueueRequestLockPutRequest {
	return ApiActRunsLastRequestQueueRequestLockPutRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		requestId: requestId,
	}
}

// Execute executes the request
//  @return ProlongRequestLockResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestLockPutExecute(r ApiActRunsLastRequestQueueRequestLockPutRequest) (*ProlongRequestLockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProlongRequestLockResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestLockPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}/lock"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lockSecs == nil {
		return localVarReturnValue, nil, reportError("lockSecs is required and must be specified")
	}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "lockSecs", r.lockSecs, "form", "")
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	if r.forefront != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forefront", r.forefront, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestPutRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	requestId string
	request *Request
	status *string
	forefront *string
	clientKey *string
}

// 
func (r ApiActRunsLastRequestQueueRequestPutRequest) Request(request Request) ApiActRunsLastRequestQueueRequestPutRequest {
	r.request = &request
	return r
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestPutRequest) Status(status string) ApiActRunsLastRequestQueueRequestPutRequest {
	r.status = &status
	return r
}

// Determines if request should be added to the head of the queue or to the end. Default value is &#x60;false&#x60; (end of queue). 
func (r ApiActRunsLastRequestQueueRequestPutRequest) Forefront(forefront string) ApiActRunsLastRequestQueueRequestPutRequest {
	r.forefront = &forefront
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestPutRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestPutRequest {
	r.clientKey = &clientKey
	return r
}

func (r ApiActRunsLastRequestQueueRequestPutRequest) Execute() (*UpdateRequestResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestPutExecute(r)
}

/*
ActRunsLastRequestQueueRequestPut Update request in last run's default request queue

Updates a request in the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Update request](/api/v2/request-queue-request-put) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param requestId Request ID.
 @return ApiActRunsLastRequestQueueRequestPutRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestPut(ctx context.Context, actorId string, requestId string) ApiActRunsLastRequestQueueRequestPutRequest {
	return ApiActRunsLastRequestQueueRequestPutRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		requestId: requestId,
	}
}

// Execute executes the request
//  @return UpdateRequestResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestPutExecute(r ApiActRunsLastRequestQueueRequestPutRequest) (*UpdateRequestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateRequestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, reportError("request is required and must be specified")
	}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.forefront != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forefront", r.forefront, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.request
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestsBatchDeleteRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	contentType *string
	requestDraftDelete *[]RequestDraftDelete
	status *string
	clientKey *string
}

// 
func (r ApiActRunsLastRequestQueueRequestsBatchDeleteRequest) ContentType(contentType string) ApiActRunsLastRequestQueueRequestsBatchDeleteRequest {
	r.contentType = &contentType
	return r
}

// 
func (r ApiActRunsLastRequestQueueRequestsBatchDeleteRequest) RequestDraftDelete(requestDraftDelete []RequestDraftDelete) ApiActRunsLastRequestQueueRequestsBatchDeleteRequest {
	r.requestDraftDelete = &requestDraftDelete
	return r
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestsBatchDeleteRequest) Status(status string) ApiActRunsLastRequestQueueRequestsBatchDeleteRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestsBatchDeleteRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestsBatchDeleteRequest {
	r.clientKey = &clientKey
	return r
}

func (r ApiActRunsLastRequestQueueRequestsBatchDeleteRequest) Execute() (*BatchDeleteResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestsBatchDeleteExecute(r)
}

/*
ActRunsLastRequestQueueRequestsBatchDelete Batch delete requests from last run's default request queue

Batch-deletes requests from the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Delete requests](/api/v2/request-queue-requests-batch-delete) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueRequestsBatchDeleteRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsBatchDelete(ctx context.Context, actorId string) ApiActRunsLastRequestQueueRequestsBatchDeleteRequest {
	return ApiActRunsLastRequestQueueRequestsBatchDeleteRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return BatchDeleteResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsBatchDeleteExecute(r ApiActRunsLastRequestQueueRequestsBatchDeleteRequest) (*BatchDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestsBatchDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/batch"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}
	if r.requestDraftDelete == nil {
		return localVarReturnValue, nil, reportError("requestDraftDelete is required and must be specified")
	}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	// body params
	localVarPostBody = r.requestDraftDelete
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestsBatchPostRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	requestBase *[]RequestBase
	status *string
	clientKey *string
	forefront *string
}

// 
func (r ApiActRunsLastRequestQueueRequestsBatchPostRequest) RequestBase(requestBase []RequestBase) ApiActRunsLastRequestQueueRequestsBatchPostRequest {
	r.requestBase = &requestBase
	return r
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestsBatchPostRequest) Status(status string) ApiActRunsLastRequestQueueRequestsBatchPostRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestsBatchPostRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestsBatchPostRequest {
	r.clientKey = &clientKey
	return r
}

// Determines if request should be added to the head of the queue or to the end. Default value is &#x60;false&#x60; (end of queue). 
func (r ApiActRunsLastRequestQueueRequestsBatchPostRequest) Forefront(forefront string) ApiActRunsLastRequestQueueRequestsBatchPostRequest {
	r.forefront = &forefront
	return r
}

func (r ApiActRunsLastRequestQueueRequestsBatchPostRequest) Execute() (*BatchAddResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestsBatchPostExecute(r)
}

/*
ActRunsLastRequestQueueRequestsBatchPost Batch add requests to last run's default request queue

Adds requests to the default request queue of the last Actor run in batch.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Add requests](/api/v2/request-queue-requests-batch-post) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueRequestsBatchPostRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsBatchPost(ctx context.Context, actorId string) ApiActRunsLastRequestQueueRequestsBatchPostRequest {
	return ApiActRunsLastRequestQueueRequestsBatchPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return BatchAddResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsBatchPostExecute(r ApiActRunsLastRequestQueueRequestsBatchPostRequest) (*BatchAddResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchAddResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestsBatchPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/batch"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBase == nil {
		return localVarReturnValue, nil, reportError("requestBase is required and must be specified")
	}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	if r.forefront != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forefront", r.forefront, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBase
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestsGetRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	status *string
	clientKey *string
	exclusiveStartId *string
	limit *float64
	cursor *string
	filter *[]string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestsGetRequest) Status(status string) ApiActRunsLastRequestQueueRequestsGetRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestsGetRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestsGetRequest {
	r.clientKey = &clientKey
	return r
}

// All requests up to this one (including) are skipped from the result. (Deprecated, use &#x60;cursor&#x60; instead.)
// Deprecated
func (r ApiActRunsLastRequestQueueRequestsGetRequest) ExclusiveStartId(exclusiveStartId string) ApiActRunsLastRequestQueueRequestsGetRequest {
	r.exclusiveStartId = &exclusiveStartId
	return r
}

// Number of keys to be returned. Maximum value is &#x60;10000&#x60;.
func (r ApiActRunsLastRequestQueueRequestsGetRequest) Limit(limit float64) ApiActRunsLastRequestQueueRequestsGetRequest {
	r.limit = &limit
	return r
}

// A cursor string for pagination, returned in the previous response as &#x60;nextCursor&#x60;. Use this to retrieve the next page of requests.
func (r ApiActRunsLastRequestQueueRequestsGetRequest) Cursor(cursor string) ApiActRunsLastRequestQueueRequestsGetRequest {
	r.cursor = &cursor
	return r
}

// Filter requests by their state. Possible values are &#x60;locked&#x60; and &#x60;pending&#x60;. You can combine multiple values separated by commas, which will mean the union of these filters – requests matching any of the specified states will be returned. (Not compatible with deprecated &#x60;exclusiveStartId&#x60; parameter.)
func (r ApiActRunsLastRequestQueueRequestsGetRequest) Filter(filter []string) ApiActRunsLastRequestQueueRequestsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiActRunsLastRequestQueueRequestsGetRequest) Execute() (*ListOfRequestsResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestsGetExecute(r)
}

/*
ActRunsLastRequestQueueRequestsGet List last run's default request queue's requests

Returns a list of requests from the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[List requests](/api/v2/request-queue-requests-get) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueRequestsGetRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsGet(ctx context.Context, actorId string) ApiActRunsLastRequestQueueRequestsGetRequest {
	return ApiActRunsLastRequestQueueRequestsGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return ListOfRequestsResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsGetExecute(r ApiActRunsLastRequestQueueRequestsGetRequest) (*ListOfRequestsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOfRequestsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	if r.exclusiveStartId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclusiveStartId", r.exclusiveStartId, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestsPostRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	requestBase *RequestBase
	status *string
	clientKey *string
	forefront *string
}

// 
func (r ApiActRunsLastRequestQueueRequestsPostRequest) RequestBase(requestBase RequestBase) ApiActRunsLastRequestQueueRequestsPostRequest {
	r.requestBase = &requestBase
	return r
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestsPostRequest) Status(status string) ApiActRunsLastRequestQueueRequestsPostRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestsPostRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestsPostRequest {
	r.clientKey = &clientKey
	return r
}

// Determines if request should be added to the head of the queue or to the end. Default value is &#x60;false&#x60; (end of queue). 
func (r ApiActRunsLastRequestQueueRequestsPostRequest) Forefront(forefront string) ApiActRunsLastRequestQueueRequestsPostRequest {
	r.forefront = &forefront
	return r
}

func (r ApiActRunsLastRequestQueueRequestsPostRequest) Execute() (*AddRequestResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestsPostExecute(r)
}

/*
ActRunsLastRequestQueueRequestsPost Add request to last run's default request queue

Adds a request to the default request queue of the last Actor run.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Add request](/api/v2/request-queue-requests-post) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueRequestsPostRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsPost(ctx context.Context, actorId string) ApiActRunsLastRequestQueueRequestsPostRequest {
	return ApiActRunsLastRequestQueueRequestsPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return AddRequestResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsPostExecute(r ApiActRunsLastRequestQueueRequestsPostRequest) (*AddRequestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddRequestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBase == nil {
		return localVarReturnValue, nil, reportError("requestBase is required and must be specified")
	}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	if r.forefront != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forefront", r.forefront, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBase
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiActRunsLastRequestQueueRequestsUnlockPostRequest struct {
	ctx context.Context
	ApiService *LastActorRunsDefaultRequestQueueAPIService
	actorId string
	status *string
	clientKey *string
}

// Filter for the run status.
func (r ApiActRunsLastRequestQueueRequestsUnlockPostRequest) Status(status string) ApiActRunsLastRequestQueueRequestsUnlockPostRequest {
	r.status = &status
	return r
}

// A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation. 
func (r ApiActRunsLastRequestQueueRequestsUnlockPostRequest) ClientKey(clientKey string) ApiActRunsLastRequestQueueRequestsUnlockPostRequest {
	r.clientKey = &clientKey
	return r
}

func (r ApiActRunsLastRequestQueueRequestsUnlockPostRequest) Execute() (*UnlockRequestsResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastRequestQueueRequestsUnlockPostExecute(r)
}

/*
ActRunsLastRequestQueueRequestsUnlockPost Unlock requests in last run's default request queue

Unlocks requests in the default request queue of the last Actor run that are currently locked by the client.

This endpoint is a shortcut for getting the last run's `defaultRequestQueueId` and then using the
[Unlock requests](/api/v2/request-queue-requests-unlock-post) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastRequestQueueRequestsUnlockPostRequest
*/
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsUnlockPost(ctx context.Context, actorId string) ApiActRunsLastRequestQueueRequestsUnlockPostRequest {
	return ApiActRunsLastRequestQueueRequestsUnlockPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return UnlockRequestsResponse
func (a *LastActorRunsDefaultRequestQueueAPIService) ActRunsLastRequestQueueRequestsUnlockPostExecute(r ApiActRunsLastRequestQueueRequestsUnlockPostRequest) (*UnlockRequestsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UnlockRequestsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LastActorRunsDefaultRequestQueueAPIService.ActRunsLastRequestQueueRequestsUnlockPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last/request-queue/requests/unlock"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.clientKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientKey", r.clientKey, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
