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


// ActorTasksAPIService ActorTasksAPI service
type ActorTasksAPIService service

type ApiActorTaskDeleteRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
}

func (r ApiActorTaskDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActorTaskDeleteExecute(r)
}

/*
ActorTaskDelete Delete task

Delete the task specified through the `actorTaskId` parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskDeleteRequest
*/
func (a *ActorTasksAPIService) ActorTaskDelete(ctx context.Context, actorTaskId string) ApiActorTaskDeleteRequest {
	return ApiActorTaskDeleteRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorTasksAPIService) ActorTaskDeleteExecute(r ApiActorTaskDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiActorTaskGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
}

func (r ApiActorTaskGetRequest) Execute() (*ActorTaskGet200Response, *http.Response, error) {
	return r.ApiService.ActorTaskGetExecute(r)
}

/*
ActorTaskGet Get task

Get an object that contains all the details about a task.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskGetRequest
*/
func (a *ActorTasksAPIService) ActorTaskGet(ctx context.Context, actorTaskId string) ApiActorTaskGetRequest {
	return ApiActorTaskGetRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return ActorTaskGet200Response
func (a *ActorTasksAPIService) ActorTaskGetExecute(r ApiActorTaskGetRequest) (*ActorTaskGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActorTaskGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiActorTaskInputGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
}

func (r ApiActorTaskInputGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActorTaskInputGetExecute(r)
}

/*
ActorTaskInputGet Get task input

Returns the input of a given task.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskInputGetRequest
*/
func (a *ActorTasksAPIService) ActorTaskInputGet(ctx context.Context, actorTaskId string) ApiActorTaskInputGetRequest {
	return ApiActorTaskInputGetRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorTasksAPIService) ActorTaskInputGetExecute(r ApiActorTaskInputGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskInputGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/input"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiActorTaskInputPutRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	body *map[string]interface{}
}

// 
func (r ApiActorTaskInputPutRequest) Body(body map[string]interface{}) ApiActorTaskInputPutRequest {
	r.body = &body
	return r
}

func (r ApiActorTaskInputPutRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActorTaskInputPutExecute(r)
}

/*
ActorTaskInputPut Update task input

Updates the input of a task using values specified by an object passed as
JSON in the PUT payload.

If the object does not define a specific property, its value is not updated.

The response is the full task input as returned by the
[Get task input](#/reference/tasks/task-input-object/get-task-input) endpoint.

The request needs to specify the `Content-Type: application/json` HTTP
header!

When providing your API authentication token, we recommend using the
request's `Authorization` header, rather than the URL. ([More
info](#/introduction/authentication)).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskInputPutRequest
*/
func (a *ActorTasksAPIService) ActorTaskInputPut(ctx context.Context, actorTaskId string) ApiActorTaskInputPutRequest {
	return ApiActorTaskInputPutRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorTasksAPIService) ActorTaskInputPutExecute(r ApiActorTaskInputPutRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskInputPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/input"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type ApiActorTaskPutRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	updateTaskRequest *UpdateTaskRequest
}

// 
func (r ApiActorTaskPutRequest) UpdateTaskRequest(updateTaskRequest UpdateTaskRequest) ApiActorTaskPutRequest {
	r.updateTaskRequest = &updateTaskRequest
	return r
}

func (r ApiActorTaskPutRequest) Execute() (*ActorTaskGet200Response, *http.Response, error) {
	return r.ApiService.ActorTaskPutExecute(r)
}

/*
ActorTaskPut Update task

Update settings of a task using values specified by an object passed as JSON
in the POST payload.

If the object does not define a specific property, its value is not updated.

The response is the full task object as returned by the
[Get task](/api/v2/actor-task-get) endpoint.

The request needs to specify the `Content-Type: application/json` HTTP
header!

When providing your API authentication token, we recommend using the
request's `Authorization` header, rather than the URL.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskPutRequest
*/
func (a *ActorTasksAPIService) ActorTaskPut(ctx context.Context, actorTaskId string) ApiActorTaskPutRequest {
	return ApiActorTaskPutRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return ActorTaskGet200Response
func (a *ActorTasksAPIService) ActorTaskPutExecute(r ApiActorTaskPutRequest) (*ActorTaskGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActorTaskGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateTaskRequest == nil {
		return localVarReturnValue, nil, reportError("updateTaskRequest is required and must be specified")
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
	localVarPostBody = r.updateTaskRequest
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

type ApiActorTaskRunSyncGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	timeout *float64
	memory *float64
	maxItems *float64
	build *string
	outputRecordKey *string
	webhooks *string
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActorTaskRunSyncGetRequest) Timeout(timeout float64) ApiActorTaskRunSyncGetRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActorTaskRunSyncGetRequest) Memory(memory float64) ApiActorTaskRunSyncGetRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActorTaskRunSyncGetRequest) MaxItems(maxItems float64) ApiActorTaskRunSyncGetRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActorTaskRunSyncGetRequest) Build(build string) ApiActorTaskRunSyncGetRequest {
	r.build = &build
	return r
}

// Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;. 
func (r ApiActorTaskRunSyncGetRequest) OutputRecordKey(outputRecordKey string) ApiActorTaskRunSyncGetRequest {
	r.outputRecordKey = &outputRecordKey
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActorTaskRunSyncGetRequest) Webhooks(webhooks string) ApiActorTaskRunSyncGetRequest {
	r.webhooks = &webhooks
	return r
}

func (r ApiActorTaskRunSyncGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActorTaskRunSyncGetExecute(r)
}

/*
ActorTaskRunSyncGet Run task synchronously

Run a specific task and return its output.

The run must finish in 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds
otherwise the HTTP request fails with a timeout error (this won't abort
the run itself).

Beware that it might be impossible to maintain an idle HTTP connection for
an extended period, due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.

If the connection breaks, you will not receive any information about the run
and its status.

To run the Task asynchronously, use the
[Run task asynchronously](#/reference/actor-tasks/run-collection/run-task)
endpoint instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskRunSyncGetRequest
*/
func (a *ActorTasksAPIService) ActorTaskRunSyncGet(ctx context.Context, actorTaskId string) ApiActorTaskRunSyncGetRequest {
	return ApiActorTaskRunSyncGetRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorTasksAPIService) ActorTaskRunSyncGetExecute(r ApiActorTaskRunSyncGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskRunSyncGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/run-sync"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.timeout != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeout", r.timeout, "form", "")
	}
	if r.memory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "memory", r.memory, "form", "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxItems", r.maxItems, "form", "")
	}
	if r.build != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build", r.build, "form", "")
	}
	if r.outputRecordKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputRecordKey", r.outputRecordKey, "form", "")
	}
	if r.webhooks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "webhooks", r.webhooks, "form", "")
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
		if localVarHTTPResponse.StatusCode == 408 {
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

type ApiActorTaskRunSyncGetDatasetItemsGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	timeout *float64
	memory *float64
	maxItems *float64
	build *string
	webhooks *string
	format *string
	clean *bool
	offset *float64
	limit *float64
	fields *string
	outputFields *string
	omit *string
	unwind *string
	flatten *string
	desc *bool
	attachment *bool
	delimiter *string
	bom *bool
	xmlRoot *string
	xmlRow *string
	skipHeaderRow *bool
	skipHidden *bool
	skipEmpty *bool
	simplified *bool
	view *string
	skipFailedPages *bool
	feedTitle *string
	feedDescription *string
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Timeout(timeout float64) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Memory(memory float64) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) MaxItems(maxItems float64) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Build(build string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.build = &build
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Webhooks(webhooks string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.webhooks = &webhooks
	return r
}

// Format of the results, possible values are: &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;csv&#x60;, &#x60;html&#x60;, &#x60;xlsx&#x60;, &#x60;xml&#x60; and &#x60;rss&#x60;. The default value is &#x60;json&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Format(format string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.format = &format
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The &#x60;clean&#x60; parameter is just a shortcut for &#x60;skipHidden&#x3D;true&#x60; and &#x60;skipEmpty&#x3D;true&#x60; parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the &#x60;limit&#x60; value. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Clean(clean bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.clean = &clean
	return r
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Offset(offset float64) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. By default there is no limit.
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Limit(limit float64) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.limit = &limit
	return r
}

// A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the &#x60;fields&#x60; query parameter. You can use this feature to effectively fix the output format. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Fields(fields string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of output field names that positionally rename the fields specified in the &#x60;fields&#x60; parameter. For example, &#x60;?fields&#x3D;headline,url&amp;outputFields&#x3D;title,link&#x60; renames &#x60;headline&#x60; to &#x60;title&#x60; and &#x60;url&#x60; to &#x60;link&#x60; in the output. The number of names in &#x60;outputFields&#x60; must match the number of names in &#x60;fields&#x60;. Requires the &#x60;fields&#x60; parameter to be specified as well. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) OutputFields(outputFields string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.outputFields = &outputFields
	return r
}

// A comma-separated list of fields which should be omitted from the items.
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Omit(omit string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.omit = &omit
	return r
}

// A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the &#x60;desc&#x60; parameter. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Unwind(unwind string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.unwind = &unwind
	return r
}

// A comma-separated list of fields which should transform nested objects into flat structures.  For example, with &#x60;flatten&#x3D;\&quot;foo\&quot;&#x60; the object &#x60;{\&quot;foo\&quot;:{\&quot;bar\&quot;: \&quot;hello\&quot;}}&#x60; is turned into &#x60;{\&quot;foo.bar\&quot;: \&quot;hello\&quot;}&#x60;.  The original object with properties is replaced with the flattened object. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Flatten(flatten string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.flatten = &flatten
	return r
}

// By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to &#x60;true&#x60; or &#x60;1&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Desc(desc bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.desc = &desc
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the response will define the &#x60;Content-Disposition: attachment&#x60; header, forcing a web browser to download the file rather than to display it. By default this header is not present. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Attachment(attachment bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.attachment = &attachment
	return r
}

// A delimiter character for CSV files, only used if &#x60;format&#x3D;csv&#x60;. You might need to URL-encode the character (e.g. use &#x60;%09&#x60; for tab or &#x60;%3B&#x60; for semicolon). The default delimiter is a simple comma (&#x60;,&#x60;). 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Delimiter(delimiter string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.delimiter = &delimiter
	return r
}

// All text responses are encoded in UTF-8 encoding. By default, the &#x60;format&#x3D;csv&#x60; files are prefixed with the UTF-8 Byte Order Mark (BOM), while &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;xml&#x60;, &#x60;html&#x60; and &#x60;rss&#x60; files are not.  If you want to override this default behavior, specify &#x60;bom&#x3D;1&#x60; query parameter to include the BOM or &#x60;bom&#x3D;0&#x60; to skip it. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Bom(bom bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.bom = &bom
	return r
}

// Overrides default root element name of &#x60;xml&#x60; output. By default the root element is &#x60;items&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) XmlRoot(xmlRoot string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.xmlRoot = &xmlRoot
	return r
}

// Overrides default element name that wraps each page or page function result object in &#x60;xml&#x60; output. By default the element name is &#x60;item&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) XmlRow(xmlRow string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.xmlRow = &xmlRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then header row in the &#x60;csv&#x60; format is skipped.
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) SkipHeaderRow(skipHeaderRow bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.skipHeaderRow = &skipHeaderRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then hidden fields are skipped from the output, i.e. fields starting with the &#x60;#&#x60; character. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) SkipHidden(skipHidden bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.skipHidden = &skipHidden
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) SkipEmpty(skipEmpty bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.skipEmpty = &skipEmpty
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the endpoint applies the &#x60;fields&#x3D;url,pageFunctionResult,errorInfo&#x60; and &#x60;unwind&#x3D;pageFunctionResult&#x60; query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Simplified(simplified bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.simplified = &simplified
	return r
}

// Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema). 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) View(view string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.view = &view
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) SkipFailedPages(skipFailedPages bool) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.skipFailedPages = &skipFailedPages
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;title&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the title defaults to &#x60;Dataset &lt;label&gt;&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) FeedTitle(feedTitle string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.feedTitle = &feedTitle
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;description&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the description defaults to &#x60;Items in dataset with id \&quot;&lt;datasetId&gt;\&quot;.&#x60; 
func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) FeedDescription(feedDescription string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	r.feedDescription = &feedDescription
	return r
}

func (r ApiActorTaskRunSyncGetDatasetItemsGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActorTaskRunSyncGetDatasetItemsGetExecute(r)
}

/*
ActorTaskRunSyncGetDatasetItemsGet Run task synchronously and get dataset items

Run a specific task and return its dataset items.

The run must finish in 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds
otherwise the HTTP request fails with a timeout error (this won't abort
the run itself).

You can send all the same options in parameters as the [Get Dataset
Items](#/reference/datasets/item-collection/get-items) API endpoint.

Beware that it might be impossible to maintain an idle HTTP connection for
an extended period, due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.

If the connection breaks, you will not receive any information about the run
and its status.

To run the Task asynchronously, use the [Run task
asynchronously](#/reference/actor-tasks/run-collection/run-task) endpoint
instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskRunSyncGetDatasetItemsGetRequest
*/
func (a *ActorTasksAPIService) ActorTaskRunSyncGetDatasetItemsGet(ctx context.Context, actorTaskId string) ApiActorTaskRunSyncGetDatasetItemsGetRequest {
	return ApiActorTaskRunSyncGetDatasetItemsGetRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorTasksAPIService) ActorTaskRunSyncGetDatasetItemsGetExecute(r ApiActorTaskRunSyncGetDatasetItemsGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskRunSyncGetDatasetItemsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/run-sync-get-dataset-items"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.timeout != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeout", r.timeout, "form", "")
	}
	if r.memory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "memory", r.memory, "form", "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxItems", r.maxItems, "form", "")
	}
	if r.build != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build", r.build, "form", "")
	}
	if r.webhooks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "webhooks", r.webhooks, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.clean != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clean", r.clean, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.outputFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputFields", r.outputFields, "form", "")
	}
	if r.omit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "omit", r.omit, "form", "")
	}
	if r.unwind != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unwind", r.unwind, "form", "")
	}
	if r.flatten != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flatten", r.flatten, "form", "")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "")
	}
	if r.attachment != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "attachment", r.attachment, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.bom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bom", r.bom, "form", "")
	}
	if r.xmlRoot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "xmlRoot", r.xmlRoot, "form", "")
	}
	if r.xmlRow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "xmlRow", r.xmlRow, "form", "")
	}
	if r.skipHeaderRow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipHeaderRow", r.skipHeaderRow, "form", "")
	}
	if r.skipHidden != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipHidden", r.skipHidden, "form", "")
	}
	if r.skipEmpty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipEmpty", r.skipEmpty, "form", "")
	}
	if r.simplified != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "simplified", r.simplified, "form", "")
	}
	if r.view != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "view", r.view, "form", "")
	}
	if r.skipFailedPages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipFailedPages", r.skipFailedPages, "form", "")
	}
	if r.feedTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "feedTitle", r.feedTitle, "form", "")
	}
	if r.feedDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "feedDescription", r.feedDescription, "form", "")
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
		if localVarHTTPResponse.StatusCode == 408 {
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

type ApiActorTaskRunSyncGetDatasetItemsPostRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	body *map[string]interface{}
	timeout *float64
	memory *float64
	maxItems *float64
	maxTotalChargeUsd *float64
	restartOnError *bool
	build *string
	webhooks *string
	format *string
	clean *bool
	offset *float64
	limit *float64
	fields *string
	outputFields *string
	omit *string
	unwind *string
	flatten *string
	desc *bool
	attachment *bool
	delimiter *string
	bom *bool
	xmlRoot *string
	xmlRow *string
	skipHeaderRow *bool
	skipHidden *bool
	skipEmpty *bool
	simplified *bool
	view *string
	skipFailedPages *bool
	feedTitle *string
	feedDescription *string
}

// 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Body(body map[string]interface{}) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.body = &body
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Timeout(timeout float64) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Memory(memory float64) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) MaxItems(maxItems float64) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) RestartOnError(restartOnError bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Build(build string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.build = &build
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Webhooks(webhooks string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.webhooks = &webhooks
	return r
}

// Format of the results, possible values are: &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;csv&#x60;, &#x60;html&#x60;, &#x60;xlsx&#x60;, &#x60;xml&#x60; and &#x60;rss&#x60;. The default value is &#x60;json&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Format(format string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.format = &format
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The &#x60;clean&#x60; parameter is just a shortcut for &#x60;skipHidden&#x3D;true&#x60; and &#x60;skipEmpty&#x3D;true&#x60; parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the &#x60;limit&#x60; value. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Clean(clean bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.clean = &clean
	return r
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Offset(offset float64) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. By default there is no limit.
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Limit(limit float64) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.limit = &limit
	return r
}

// A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the &#x60;fields&#x60; query parameter. You can use this feature to effectively fix the output format. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Fields(fields string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of output field names that positionally rename the fields specified in the &#x60;fields&#x60; parameter. For example, &#x60;?fields&#x3D;headline,url&amp;outputFields&#x3D;title,link&#x60; renames &#x60;headline&#x60; to &#x60;title&#x60; and &#x60;url&#x60; to &#x60;link&#x60; in the output. The number of names in &#x60;outputFields&#x60; must match the number of names in &#x60;fields&#x60;. Requires the &#x60;fields&#x60; parameter to be specified as well. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) OutputFields(outputFields string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.outputFields = &outputFields
	return r
}

// A comma-separated list of fields which should be omitted from the items.
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Omit(omit string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.omit = &omit
	return r
}

// A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the &#x60;desc&#x60; parameter. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Unwind(unwind string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.unwind = &unwind
	return r
}

// A comma-separated list of fields which should transform nested objects into flat structures.  For example, with &#x60;flatten&#x3D;\&quot;foo\&quot;&#x60; the object &#x60;{\&quot;foo\&quot;:{\&quot;bar\&quot;: \&quot;hello\&quot;}}&#x60; is turned into &#x60;{\&quot;foo.bar\&quot;: \&quot;hello\&quot;}&#x60;.  The original object with properties is replaced with the flattened object. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Flatten(flatten string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.flatten = &flatten
	return r
}

// By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to &#x60;true&#x60; or &#x60;1&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Desc(desc bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.desc = &desc
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the response will define the &#x60;Content-Disposition: attachment&#x60; header, forcing a web browser to download the file rather than to display it. By default this header is not present. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Attachment(attachment bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.attachment = &attachment
	return r
}

// A delimiter character for CSV files, only used if &#x60;format&#x3D;csv&#x60;. You might need to URL-encode the character (e.g. use &#x60;%09&#x60; for tab or &#x60;%3B&#x60; for semicolon). The default delimiter is a simple comma (&#x60;,&#x60;). 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Delimiter(delimiter string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.delimiter = &delimiter
	return r
}

// All text responses are encoded in UTF-8 encoding. By default, the &#x60;format&#x3D;csv&#x60; files are prefixed with the UTF-8 Byte Order Mark (BOM), while &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;xml&#x60;, &#x60;html&#x60; and &#x60;rss&#x60; files are not.  If you want to override this default behavior, specify &#x60;bom&#x3D;1&#x60; query parameter to include the BOM or &#x60;bom&#x3D;0&#x60; to skip it. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Bom(bom bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.bom = &bom
	return r
}

// Overrides default root element name of &#x60;xml&#x60; output. By default the root element is &#x60;items&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) XmlRoot(xmlRoot string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.xmlRoot = &xmlRoot
	return r
}

// Overrides default element name that wraps each page or page function result object in &#x60;xml&#x60; output. By default the element name is &#x60;item&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) XmlRow(xmlRow string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.xmlRow = &xmlRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then header row in the &#x60;csv&#x60; format is skipped.
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) SkipHeaderRow(skipHeaderRow bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.skipHeaderRow = &skipHeaderRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then hidden fields are skipped from the output, i.e. fields starting with the &#x60;#&#x60; character. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) SkipHidden(skipHidden bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.skipHidden = &skipHidden
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) SkipEmpty(skipEmpty bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.skipEmpty = &skipEmpty
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the endpoint applies the &#x60;fields&#x3D;url,pageFunctionResult,errorInfo&#x60; and &#x60;unwind&#x3D;pageFunctionResult&#x60; query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Simplified(simplified bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.simplified = &simplified
	return r
}

// Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema). 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) View(view string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.view = &view
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) SkipFailedPages(skipFailedPages bool) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.skipFailedPages = &skipFailedPages
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;title&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the title defaults to &#x60;Dataset &lt;label&gt;&#x60;. 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) FeedTitle(feedTitle string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.feedTitle = &feedTitle
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;description&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the description defaults to &#x60;Items in dataset with id \&quot;&lt;datasetId&gt;\&quot;.&#x60; 
func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) FeedDescription(feedDescription string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	r.feedDescription = &feedDescription
	return r
}

func (r ApiActorTaskRunSyncGetDatasetItemsPostRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActorTaskRunSyncGetDatasetItemsPostExecute(r)
}

/*
ActorTaskRunSyncGetDatasetItemsPost Run task synchronously and get dataset items

Runs an Actor task and synchronously returns its dataset items.

The run must finish in 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds
otherwise the HTTP request fails with a timeout error (this won't abort
the run itself).

Optionally, you can override the Actor input configuration by passing a JSON
object as the POST payload and setting the `Content-Type: application/json` HTTP header.

Note that if the object in the POST payload does not define a particular
input property, the Actor run uses the default value defined by the task (or the Actor's
input schema if not defined by the task).

You can send all the same options in parameters as the [Get Dataset
Items](#/reference/datasets/item-collection/get-items) API endpoint.

Beware that it might be impossible to maintain an idle HTTP connection for
an extended period, due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.

If the connection breaks, you will not receive any information about the run
and its status.

Input fields from Actor task configuration can be overloaded with values
passed as the POST payload.

Just make sure to specify the `Content-Type` header as `application/json`
and that the input is an object.

To run the task asynchronously, use the [Run
task](#/reference/actor-tasks/run-collection/run-task) API endpoint instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskRunSyncGetDatasetItemsPostRequest
*/
func (a *ActorTasksAPIService) ActorTaskRunSyncGetDatasetItemsPost(ctx context.Context, actorTaskId string) ApiActorTaskRunSyncGetDatasetItemsPostRequest {
	return ApiActorTaskRunSyncGetDatasetItemsPostRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorTasksAPIService) ActorTaskRunSyncGetDatasetItemsPostExecute(r ApiActorTaskRunSyncGetDatasetItemsPostRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskRunSyncGetDatasetItemsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/run-sync-get-dataset-items"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.timeout != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeout", r.timeout, "form", "")
	}
	if r.memory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "memory", r.memory, "form", "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxItems", r.maxItems, "form", "")
	}
	if r.maxTotalChargeUsd != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxTotalChargeUsd", r.maxTotalChargeUsd, "form", "")
	}
	if r.restartOnError != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "restartOnError", r.restartOnError, "form", "")
	}
	if r.build != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build", r.build, "form", "")
	}
	if r.webhooks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "webhooks", r.webhooks, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.clean != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clean", r.clean, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.outputFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputFields", r.outputFields, "form", "")
	}
	if r.omit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "omit", r.omit, "form", "")
	}
	if r.unwind != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unwind", r.unwind, "form", "")
	}
	if r.flatten != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flatten", r.flatten, "form", "")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "")
	}
	if r.attachment != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "attachment", r.attachment, "form", "")
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	}
	if r.bom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bom", r.bom, "form", "")
	}
	if r.xmlRoot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "xmlRoot", r.xmlRoot, "form", "")
	}
	if r.xmlRow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "xmlRow", r.xmlRow, "form", "")
	}
	if r.skipHeaderRow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipHeaderRow", r.skipHeaderRow, "form", "")
	}
	if r.skipHidden != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipHidden", r.skipHidden, "form", "")
	}
	if r.skipEmpty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipEmpty", r.skipEmpty, "form", "")
	}
	if r.simplified != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "simplified", r.simplified, "form", "")
	}
	if r.view != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "view", r.view, "form", "")
	}
	if r.skipFailedPages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipFailedPages", r.skipFailedPages, "form", "")
	}
	if r.feedTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "feedTitle", r.feedTitle, "form", "")
	}
	if r.feedDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "feedDescription", r.feedDescription, "form", "")
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
	localVarPostBody = r.body
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

type ApiActorTaskRunSyncPostRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	body *map[string]interface{}
	timeout *float64
	memory *float64
	maxItems *float64
	maxTotalChargeUsd *float64
	restartOnError *bool
	build *string
	outputRecordKey *string
	webhooks *string
}

// 
func (r ApiActorTaskRunSyncPostRequest) Body(body map[string]interface{}) ApiActorTaskRunSyncPostRequest {
	r.body = &body
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActorTaskRunSyncPostRequest) Timeout(timeout float64) ApiActorTaskRunSyncPostRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActorTaskRunSyncPostRequest) Memory(memory float64) ApiActorTaskRunSyncPostRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActorTaskRunSyncPostRequest) MaxItems(maxItems float64) ApiActorTaskRunSyncPostRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActorTaskRunSyncPostRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActorTaskRunSyncPostRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActorTaskRunSyncPostRequest) RestartOnError(restartOnError bool) ApiActorTaskRunSyncPostRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActorTaskRunSyncPostRequest) Build(build string) ApiActorTaskRunSyncPostRequest {
	r.build = &build
	return r
}

// Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;. 
func (r ApiActorTaskRunSyncPostRequest) OutputRecordKey(outputRecordKey string) ApiActorTaskRunSyncPostRequest {
	r.outputRecordKey = &outputRecordKey
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActorTaskRunSyncPostRequest) Webhooks(webhooks string) ApiActorTaskRunSyncPostRequest {
	r.webhooks = &webhooks
	return r
}

func (r ApiActorTaskRunSyncPostRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActorTaskRunSyncPostExecute(r)
}

/*
ActorTaskRunSyncPost Run task synchronously

Runs an Actor task and synchronously returns its output.

The run must finish in 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds
otherwise the HTTP request fails with a timeout error (this won't abort
the run itself).

Optionally, you can override the Actor input configuration by passing a JSON
object as the POST payload and setting the `Content-Type: application/json` HTTP header.

Note that if the object in the POST payload does not define a particular
input property, the Actor run uses the default value defined by the task (or Actor's input
schema if not defined by the task).

Beware that it might be impossible to maintain an idle HTTP connection for
an extended period, due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.

If the connection breaks, you will not receive any information about the run
and its status.

Input fields from Actor task configuration can be overloaded with values
passed as the POST payload.

Just make sure to specify `Content-Type` header to be `application/json` and
input to be an object.

To run the task asynchronously, use the [Run
task](#/reference/actor-tasks/run-collection/run-task) API endpoint instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskRunSyncPostRequest
*/
func (a *ActorTasksAPIService) ActorTaskRunSyncPost(ctx context.Context, actorTaskId string) ApiActorTaskRunSyncPostRequest {
	return ApiActorTaskRunSyncPostRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorTasksAPIService) ActorTaskRunSyncPostExecute(r ApiActorTaskRunSyncPostRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskRunSyncPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/run-sync"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.timeout != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeout", r.timeout, "form", "")
	}
	if r.memory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "memory", r.memory, "form", "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxItems", r.maxItems, "form", "")
	}
	if r.maxTotalChargeUsd != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxTotalChargeUsd", r.maxTotalChargeUsd, "form", "")
	}
	if r.restartOnError != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "restartOnError", r.restartOnError, "form", "")
	}
	if r.build != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build", r.build, "form", "")
	}
	if r.outputRecordKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputRecordKey", r.outputRecordKey, "form", "")
	}
	if r.webhooks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "webhooks", r.webhooks, "form", "")
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
	localVarPostBody = r.body
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

type ApiActorTaskRunsGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	offset *float64
	limit *float64
	desc *bool
	status *[]string
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActorTaskRunsGetRequest) Offset(offset float64) ApiActorTaskRunsGetRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;. 
func (r ApiActorTaskRunsGetRequest) Limit(limit float64) ApiActorTaskRunsGetRequest {
	r.limit = &limit
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;startedAt&#x60; field in descending order. By default, they are sorted in ascending order. 
func (r ApiActorTaskRunsGetRequest) Desc(desc bool) ApiActorTaskRunsGetRequest {
	r.desc = &desc
	return r
}

// Single status or comma-separated list of statuses, see ([available statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)). Used to filter runs by the specified statuses only. 
func (r ApiActorTaskRunsGetRequest) Status(status []string) ApiActorTaskRunsGetRequest {
	r.status = &status
	return r
}

func (r ApiActorTaskRunsGetRequest) Execute() (*ActorTaskRunsGet200Response, *http.Response, error) {
	return r.ApiService.ActorTaskRunsGetExecute(r)
}

/*
ActorTaskRunsGet Get list of task runs

Get a list of runs of a specific task. The response is a list of objects,
where each object contains essential information about a single task run.

The endpoint supports pagination using the `limit` and `offset` parameters,
and it does not return more than a 1000 array elements.

By default, the records are sorted by the `startedAt` field in ascending
order; therefore you can use pagination to incrementally fetch all records while
new ones are still being created. To sort the records in descending order, use
the `desc=1` parameter. You can also filter runs by status ([available
statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskRunsGetRequest
*/
func (a *ActorTasksAPIService) ActorTaskRunsGet(ctx context.Context, actorTaskId string) ApiActorTaskRunsGetRequest {
	return ApiActorTaskRunsGetRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return ActorTaskRunsGet200Response
func (a *ActorTasksAPIService) ActorTaskRunsGetExecute(r ApiActorTaskRunsGetRequest) (*ActorTaskRunsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActorTaskRunsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskRunsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "csv")
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

type ApiActorTaskRunsLastGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	status *string
	waitForFinish *float64
}

// Filter for the run status.
func (r ApiActorTaskRunsLastGetRequest) Status(status string) ApiActorTaskRunsLastGetRequest {
	r.status = &status
	return r
}

// The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;). 
func (r ApiActorTaskRunsLastGetRequest) WaitForFinish(waitForFinish float64) ApiActorTaskRunsLastGetRequest {
	r.waitForFinish = &waitForFinish
	return r
}

func (r ApiActorTaskRunsLastGetRequest) Execute() (*ActorTaskRunsPost201Response, *http.Response, error) {
	return r.ApiService.ActorTaskRunsLastGetExecute(r)
}

/*
ActorTaskRunsLastGet Get last run

This is not a single endpoint, but an entire group of endpoints that lets you to
retrieve and manage the last run of given actor task or any of its default storages.
All the endpoints require an authentication token.

The base path represents the last actor task run object is:

`/v2/actor-tasks/{actorTaskId}/runs/last{?token,status}`

Using the `status` query parameter you can ensure to only get a run with a certain status
(e.g. `status=SUCCEEDED`). The output of this endpoint and other query parameters
are the same as in the [Run object](/api/v2/actor-run-get) endpoint.

##### Convenience endpoints for last Actor task run

* [Dataset](/api/v2/last-actor-task-runs-default-dataset)

* [Key-value store](/api/v2/last-actor-task-runs-default-key-value-store)

* [Request queue](/api/v2/last-actor-task-runs-default-request-queue)

* [Log](/api/v2/last-actor-task-runs-log)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskRunsLastGetRequest
*/
func (a *ActorTasksAPIService) ActorTaskRunsLastGet(ctx context.Context, actorTaskId string) ApiActorTaskRunsLastGetRequest {
	return ApiActorTaskRunsLastGetRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return ActorTaskRunsPost201Response
func (a *ActorTasksAPIService) ActorTaskRunsLastGetExecute(r ApiActorTaskRunsLastGetRequest) (*ActorTaskRunsPost201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActorTaskRunsPost201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskRunsLastGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/runs/last"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.waitForFinish != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "waitForFinish", r.waitForFinish, "form", "")
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

type ApiActorTaskRunsPostRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	body *map[string]interface{}
	timeout *float64
	memory *float64
	maxItems *float64
	maxTotalChargeUsd *float64
	restartOnError *bool
	build *string
	waitForFinish *float64
	webhooks *string
}

// 
func (r ApiActorTaskRunsPostRequest) Body(body map[string]interface{}) ApiActorTaskRunsPostRequest {
	r.body = &body
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActorTaskRunsPostRequest) Timeout(timeout float64) ApiActorTaskRunsPostRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActorTaskRunsPostRequest) Memory(memory float64) ApiActorTaskRunsPostRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActorTaskRunsPostRequest) MaxItems(maxItems float64) ApiActorTaskRunsPostRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActorTaskRunsPostRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActorTaskRunsPostRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActorTaskRunsPostRequest) RestartOnError(restartOnError bool) ApiActorTaskRunsPostRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActorTaskRunsPostRequest) Build(build string) ApiActorTaskRunsPostRequest {
	r.build = &build
	return r
}

// The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;). 
func (r ApiActorTaskRunsPostRequest) WaitForFinish(waitForFinish float64) ApiActorTaskRunsPostRequest {
	r.waitForFinish = &waitForFinish
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActorTaskRunsPostRequest) Webhooks(webhooks string) ApiActorTaskRunsPostRequest {
	r.webhooks = &webhooks
	return r
}

func (r ApiActorTaskRunsPostRequest) Execute() (*ActorTaskRunsPost201Response, *http.Response, error) {
	return r.ApiService.ActorTaskRunsPostExecute(r)
}

/*
ActorTaskRunsPost Run task

Runs an Actor task and immediately returns without waiting for the run to
finish.

Optionally, you can override the Actor input configuration by passing a JSON
object as the POST payload and setting the `Content-Type: application/json` HTTP header.

Note that if the object in the POST payload does not define a particular
input property, the Actor run uses the default value defined by the task (or Actor's input
schema if not defined by the task).

The response is the Actor Run object as returned by the [Get
run](#/reference/actor-runs/run-object-and-its-storages/get-run) endpoint.

If you want to wait for the run to finish and receive the actual output of
the Actor run as the response, use one of the [Run task
synchronously](#/reference/actor-tasks/run-task-synchronously) API endpoints
instead.

To fetch the Actor run results that are typically stored in the default
dataset, you'll need to pass the ID received in the `defaultDatasetId` field
received in the response JSON to the
[Get dataset items](#/reference/datasets/item-collection/get-items) API endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskRunsPostRequest
*/
func (a *ActorTasksAPIService) ActorTaskRunsPost(ctx context.Context, actorTaskId string) ApiActorTaskRunsPostRequest {
	return ApiActorTaskRunsPostRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return ActorTaskRunsPost201Response
func (a *ActorTasksAPIService) ActorTaskRunsPostExecute(r ApiActorTaskRunsPostRequest) (*ActorTaskRunsPost201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActorTaskRunsPost201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskRunsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.timeout != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeout", r.timeout, "form", "")
	}
	if r.memory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "memory", r.memory, "form", "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxItems", r.maxItems, "form", "")
	}
	if r.maxTotalChargeUsd != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxTotalChargeUsd", r.maxTotalChargeUsd, "form", "")
	}
	if r.restartOnError != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "restartOnError", r.restartOnError, "form", "")
	}
	if r.build != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build", r.build, "form", "")
	}
	if r.waitForFinish != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "waitForFinish", r.waitForFinish, "form", "")
	}
	if r.webhooks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "webhooks", r.webhooks, "form", "")
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
	localVarPostBody = r.body
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

type ApiActorTaskWebhooksGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	actorTaskId string
	offset *float64
	limit *float64
	desc *bool
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActorTaskWebhooksGetRequest) Offset(offset float64) ApiActorTaskWebhooksGetRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;. 
func (r ApiActorTaskWebhooksGetRequest) Limit(limit float64) ApiActorTaskWebhooksGetRequest {
	r.limit = &limit
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order. 
func (r ApiActorTaskWebhooksGetRequest) Desc(desc bool) ApiActorTaskWebhooksGetRequest {
	r.desc = &desc
	return r
}

func (r ApiActorTaskWebhooksGetRequest) Execute() (*ActorTaskWebhooksGet200Response, *http.Response, error) {
	return r.ApiService.ActorTaskWebhooksGetExecute(r)
}

/*
ActorTaskWebhooksGet Get list of webhooks

Gets the list of webhooks of a specific Actor task. The response is a JSON
with the list of objects, where each object contains basic information about a single webhook.

The endpoint supports pagination using the `limit` and `offset` parameters
and it will not return more than 1000 records.

By default, the records are sorted by the `createdAt` field in ascending
order, to sort the records in descending order, use the `desc=1` parameter.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorTaskId Task ID or a tilde-separated owner's username and task's name.
 @return ApiActorTaskWebhooksGetRequest
*/
func (a *ActorTasksAPIService) ActorTaskWebhooksGet(ctx context.Context, actorTaskId string) ApiActorTaskWebhooksGetRequest {
	return ApiActorTaskWebhooksGetRequest{
		ApiService: a,
		ctx: ctx,
		actorTaskId: actorTaskId,
	}
}

// Execute executes the request
//  @return ActorTaskWebhooksGet200Response
func (a *ActorTasksAPIService) ActorTaskWebhooksGetExecute(r ApiActorTaskWebhooksGetRequest) (*ActorTaskWebhooksGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActorTaskWebhooksGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTaskWebhooksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks/{actorTaskId}/webhooks"
	localVarPath = strings.Replace(localVarPath, "{"+"actorTaskId"+"}", url.PathEscape(parameterValueToString(r.actorTaskId, "actorTaskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "")
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

type ApiActorTasksGetRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	offset *float64
	limit *float64
	desc *bool
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActorTasksGetRequest) Offset(offset float64) ApiActorTasksGetRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;. 
func (r ApiActorTasksGetRequest) Limit(limit float64) ApiActorTasksGetRequest {
	r.limit = &limit
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order. 
func (r ApiActorTasksGetRequest) Desc(desc bool) ApiActorTasksGetRequest {
	r.desc = &desc
	return r
}

func (r ApiActorTasksGetRequest) Execute() (*ListOfTasksResponse, *http.Response, error) {
	return r.ApiService.ActorTasksGetExecute(r)
}

/*
ActorTasksGet Get list of tasks

Gets the complete list of tasks that a user has created or used.

The response is a list of objects in which each object contains essential
information about a single task.

The endpoint supports pagination using the `limit` and `offset` parameters,
and it does not return more than a 1000 records.

By default, the records are sorted by the `createdAt` field in ascending
order; therefore you can use pagination to incrementally fetch all tasks while new
ones are still being created. To sort the records in descending order, use
the `desc=1` parameter.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiActorTasksGetRequest
*/
func (a *ActorTasksAPIService) ActorTasksGet(ctx context.Context) ApiActorTasksGetRequest {
	return ApiActorTasksGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListOfTasksResponse
func (a *ActorTasksAPIService) ActorTasksGetExecute(r ApiActorTasksGetRequest) (*ListOfTasksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOfTasksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTasksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "")
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

type ApiActorTasksPostRequest struct {
	ctx context.Context
	ApiService *ActorTasksAPIService
	createTaskRequest *CreateTaskRequest
}

// 
func (r ApiActorTasksPostRequest) CreateTaskRequest(createTaskRequest CreateTaskRequest) ApiActorTasksPostRequest {
	r.createTaskRequest = &createTaskRequest
	return r
}

func (r ApiActorTasksPostRequest) Execute() (*TaskResponse, *http.Response, error) {
	return r.ApiService.ActorTasksPostExecute(r)
}

/*
ActorTasksPost Create task

Create a new task with settings specified by the object passed as JSON in
the POST payload.

The response is the full task object as returned by the
[Get task](/api/v2/actor-task-get) endpoint.

The request needs to specify the `Content-Type: application/json` HTTP header!

When providing your API authentication token, we recommend using the
request's `Authorization` header, rather than the URL.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiActorTasksPostRequest
*/
func (a *ActorTasksAPIService) ActorTasksPost(ctx context.Context) ApiActorTasksPostRequest {
	return ApiActorTasksPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TaskResponse
func (a *ActorTasksAPIService) ActorTasksPostExecute(r ApiActorTasksPostRequest) (*TaskResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorTasksAPIService.ActorTasksPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actor-tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createTaskRequest == nil {
		return localVarReturnValue, nil, reportError("createTaskRequest is required and must be specified")
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
	localVarPostBody = r.createTaskRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
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
