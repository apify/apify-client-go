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
	"time"
)


// ActorsActorRunsAPIService ActorsActorRunsAPI service
type ActorsActorRunsAPIService service

type ApiActRunAbortPostRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	runId string
	gracefully *bool
}

// If true passed, the Actor run will abort gracefully. It will send &#x60;aborting&#x60; and &#x60;persistState&#x60; event into run and force-stop the run after 30 seconds. It is helpful in cases where you plan to resurrect the run later. 
func (r ApiActRunAbortPostRequest) Gracefully(gracefully bool) ApiActRunAbortPostRequest {
	r.gracefully = &gracefully
	return r
}

func (r ApiActRunAbortPostRequest) Execute() (*RunResponse, *http.Response, error) {
	return r.ApiService.ActRunAbortPostExecute(r)
}

/*
ActRunAbortPost Abort run

**[DEPRECATED]** API endpoints related to run of the Actor were moved under
new namespace [`actor-runs`](#/reference/actor-runs). Aborts an Actor run and
returns an object that contains all the details about the run.

Only runs that are starting or running are aborted. For runs with status
`FINISHED`, `FAILED`, `ABORTING` and `TIMED-OUT` this call does nothing.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param runId Actor run ID.
 @return ApiActRunAbortPostRequest

Deprecated
*/
func (a *ActorsActorRunsAPIService) ActRunAbortPost(ctx context.Context, actorId string, runId string) ApiActRunAbortPostRequest {
	return ApiActRunAbortPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		runId: runId,
	}
}

// Execute executes the request
//  @return RunResponse
// Deprecated
func (a *ActorsActorRunsAPIService) ActRunAbortPostExecute(r ApiActRunAbortPostRequest) (*RunResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RunResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunAbortPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/{runId}/abort"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"runId"+"}", url.PathEscape(parameterValueToString(r.runId, "runId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.gracefully != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gracefully", r.gracefully, "form", "")
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

type ApiActRunGetRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	runId string
	waitForFinish *float64
}

// The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;). 
func (r ApiActRunGetRequest) WaitForFinish(waitForFinish float64) ApiActRunGetRequest {
	r.waitForFinish = &waitForFinish
	return r
}

func (r ApiActRunGetRequest) Execute() (*RunResponse, *http.Response, error) {
	return r.ApiService.ActRunGetExecute(r)
}

/*
ActRunGet Get run

**[DEPRECATED]** API endpoints related to run of the Actor were moved under
new namespace [`actor-runs`](#/reference/actor-runs).

Gets an object that contains all the details about a specific run of an Actor.

By passing the optional `waitForFinish` parameter the API endpoint will
synchronously wait for the run to finish.
This is useful to avoid periodic polling when waiting for Actor run to
complete.

This endpoint does not require the authentication token. Instead, calls are authenticated using a hard-to-guess ID of the run. However,
if you access the endpoint without the token, certain attributes, such as `usageUsd` and `usageTotalUsd`, will be hidden.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param runId Actor run ID.
 @return ApiActRunGetRequest

Deprecated
*/
func (a *ActorsActorRunsAPIService) ActRunGet(ctx context.Context, actorId string, runId string) ApiActRunGetRequest {
	return ApiActRunGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		runId: runId,
	}
}

// Execute executes the request
//  @return RunResponse
// Deprecated
func (a *ActorsActorRunsAPIService) ActRunGetExecute(r ApiActRunGetRequest) (*RunResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RunResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/{runId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"runId"+"}", url.PathEscape(parameterValueToString(r.runId, "runId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiActRunMetamorphPostRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	runId string
	targetActorId *string
	build *string
}

// ID of a target Actor that the run should be transformed into.
func (r ApiActRunMetamorphPostRequest) TargetActorId(targetActorId string) ApiActRunMetamorphPostRequest {
	r.targetActorId = &targetActorId
	return r
}

// Optional build of the target Actor.  It can be either a build tag or build number. By default, the run uses the build specified in the default run configuration for the target Actor (typically &#x60;latest&#x60;). 
func (r ApiActRunMetamorphPostRequest) Build(build string) ApiActRunMetamorphPostRequest {
	r.build = &build
	return r
}

func (r ApiActRunMetamorphPostRequest) Execute() (*RunResponse, *http.Response, error) {
	return r.ApiService.ActRunMetamorphPostExecute(r)
}

/*
ActRunMetamorphPost Metamorph run

**[DEPRECATED]** API endpoints related to run of the Actor were moved under
new namespace [`actor-runs`](#/reference/actor-runs). Transforms an Actor run
into a run of another Actor with a new input.

This is useful if you want to use another Actor to finish the work
of your current Actor run, without the need to create a completely new run
and waiting for its finish.
For the users of your Actors, the metamorph operation is transparent, they
will just see your Actor got the work done.

There is a limit on how many times you can metamorph a single run. You can
check the limit in [the Actor runtime limits](https://docs.apify.com/platform/limits#actor-limits).

Internally, the system stops the Docker container corresponding to the Actor
run and starts a new container using a different Docker image.
All the default storages are preserved and the new input is stored under the
`INPUT-METAMORPH-1` key in the same default key-value store.

For more information, see the [Actor docs](https://docs.apify.com/platform/actors/development/programming-interface/metamorph).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param runId Actor run ID.
 @return ApiActRunMetamorphPostRequest

Deprecated
*/
func (a *ActorsActorRunsAPIService) ActRunMetamorphPost(ctx context.Context, actorId string, runId string) ApiActRunMetamorphPostRequest {
	return ApiActRunMetamorphPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		runId: runId,
	}
}

// Execute executes the request
//  @return RunResponse
// Deprecated
func (a *ActorsActorRunsAPIService) ActRunMetamorphPostExecute(r ApiActRunMetamorphPostRequest) (*RunResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RunResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunMetamorphPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/{runId}/metamorph"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"runId"+"}", url.PathEscape(parameterValueToString(r.runId, "runId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targetActorId == nil {
		return localVarReturnValue, nil, reportError("targetActorId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "targetActorId", r.targetActorId, "form", "")
	if r.build != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build", r.build, "form", "")
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

type ApiActRunResurrectPostRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	runId string
	build *string
	timeout *float64
	memory *float64
	restartOnError *bool
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run is resurrected with the same build it originally used. Specifically, if a run was first started with the &#x60;latest&#x60; tag, which resolves to version &#x60;0.0.3&#x60; at the time, a run resurrected without this parameter will continue running with &#x60;0.0.3&#x60;, even if &#x60;latest&#x60; already points to a newer build. 
func (r ApiActRunResurrectPostRequest) Build(build string) ApiActRunResurrectPostRequest {
	r.build = &build
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout specified in the run that is being resurrected. 
func (r ApiActRunResurrectPostRequest) Timeout(timeout float64) ApiActRunResurrectPostRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit specified in the run that is being resurrected. 
func (r ApiActRunResurrectPostRequest) Memory(memory float64) ApiActRunResurrectPostRequest {
	r.memory = &memory
	return r
}

// Determines whether the resurrected run will be restarted if it fails. By default, the resurrected run uses the same setting as before. 
func (r ApiActRunResurrectPostRequest) RestartOnError(restartOnError bool) ApiActRunResurrectPostRequest {
	r.restartOnError = &restartOnError
	return r
}

func (r ApiActRunResurrectPostRequest) Execute() (*RunResponse, *http.Response, error) {
	return r.ApiService.ActRunResurrectPostExecute(r)
}

/*
ActRunResurrectPost Resurrect run

**[DEPRECATED]** API endpoints related to run of the Actor were moved under
new namespace [`actor-runs`](#/reference/actor-runs).Resurrects a finished
Actor run and returns an object that contains all the details about the
resurrected run.

Only finished runs, i.e. runs with status `FINISHED`, `FAILED`, `ABORTED`
and `TIMED-OUT` can be resurrected.
Run status will be updated to RUNNING and its container will be restarted
with the same storages
(the same behaviour as when the run gets migrated to the new server).

For more information, see the [Actor
docs](https://docs.apify.com/platform/actors/running/runs-and-builds#resurrection-of-finished-run).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @param runId Actor run ID.
 @return ApiActRunResurrectPostRequest
*/
func (a *ActorsActorRunsAPIService) ActRunResurrectPost(ctx context.Context, actorId string, runId string) ApiActRunResurrectPostRequest {
	return ApiActRunResurrectPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
		runId: runId,
	}
}

// Execute executes the request
//  @return RunResponse
func (a *ActorsActorRunsAPIService) ActRunResurrectPostExecute(r ApiActRunResurrectPostRequest) (*RunResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RunResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunResurrectPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/{runId}/resurrect"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"runId"+"}", url.PathEscape(parameterValueToString(r.runId, "runId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.build != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build", r.build, "form", "")
	}
	if r.timeout != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeout", r.timeout, "form", "")
	}
	if r.memory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "memory", r.memory, "form", "")
	}
	if r.restartOnError != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "restartOnError", r.restartOnError, "form", "")
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

type ApiActRunSyncGetRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	outputRecordKey *string
	timeout *float64
	memory *float64
	maxItems *float64
	maxTotalChargeUsd *float64
	restartOnError *bool
	build *string
	webhooks *string
}

// Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;. 
func (r ApiActRunSyncGetRequest) OutputRecordKey(outputRecordKey string) ApiActRunSyncGetRequest {
	r.outputRecordKey = &outputRecordKey
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActRunSyncGetRequest) Timeout(timeout float64) ApiActRunSyncGetRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActRunSyncGetRequest) Memory(memory float64) ApiActRunSyncGetRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActRunSyncGetRequest) MaxItems(maxItems float64) ApiActRunSyncGetRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActRunSyncGetRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActRunSyncGetRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActRunSyncGetRequest) RestartOnError(restartOnError bool) ApiActRunSyncGetRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActRunSyncGetRequest) Build(build string) ApiActRunSyncGetRequest {
	r.build = &build
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActRunSyncGetRequest) Webhooks(webhooks string) ApiActRunSyncGetRequest {
	r.webhooks = &webhooks
	return r
}

func (r ApiActRunSyncGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActRunSyncGetExecute(r)
}

/*
ActRunSyncGet Run Actor synchronously without input

Runs a specific Actor and returns its output.
The run must finish in 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds
otherwise the API endpoint returns a timeout error.
The Actor is not passed any input.

Beware that it might be impossible to maintain an idle HTTP connection for a
long period of time,
due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.
If the connection breaks, you will not receive any information about the run
and its status.

To run the Actor asynchronously, use the [Run
Actor](#/reference/actors/run-collection/run-actor) API endpoint instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunSyncGetRequest
*/
func (a *ActorsActorRunsAPIService) ActRunSyncGet(ctx context.Context, actorId string) ApiActRunSyncGetRequest {
	return ApiActRunSyncGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorsActorRunsAPIService) ActRunSyncGetExecute(r ApiActRunSyncGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunSyncGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/run-sync"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.outputRecordKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputRecordKey", r.outputRecordKey, "form", "")
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
			var v ActorRunFailedError
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
		if localVarHTTPResponse.StatusCode == 408 {
			var v ActorRunTimeoutExceededError
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

type ApiActRunSyncGetDatasetItemsGetRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
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

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Timeout(timeout float64) ApiActRunSyncGetDatasetItemsGetRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Memory(memory float64) ApiActRunSyncGetDatasetItemsGetRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) MaxItems(maxItems float64) ApiActRunSyncGetDatasetItemsGetRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActRunSyncGetDatasetItemsGetRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) RestartOnError(restartOnError bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Build(build string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.build = &build
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Webhooks(webhooks string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.webhooks = &webhooks
	return r
}

// Format of the results, possible values are: &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;csv&#x60;, &#x60;html&#x60;, &#x60;xlsx&#x60;, &#x60;xml&#x60; and &#x60;rss&#x60;. The default value is &#x60;json&#x60;. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Format(format string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.format = &format
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The &#x60;clean&#x60; parameter is just a shortcut for &#x60;skipHidden&#x3D;true&#x60; and &#x60;skipEmpty&#x3D;true&#x60; parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the &#x60;limit&#x60; value. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Clean(clean bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.clean = &clean
	return r
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Offset(offset float64) ApiActRunSyncGetDatasetItemsGetRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. By default there is no limit.
func (r ApiActRunSyncGetDatasetItemsGetRequest) Limit(limit float64) ApiActRunSyncGetDatasetItemsGetRequest {
	r.limit = &limit
	return r
}

// A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the &#x60;fields&#x60; query parameter. You can use this feature to effectively fix the output format. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Fields(fields string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of output field names that positionally rename the fields specified in the &#x60;fields&#x60; parameter. For example, &#x60;?fields&#x3D;headline,url&amp;outputFields&#x3D;title,link&#x60; renames &#x60;headline&#x60; to &#x60;title&#x60; and &#x60;url&#x60; to &#x60;link&#x60; in the output. The number of names in &#x60;outputFields&#x60; must match the number of names in &#x60;fields&#x60;. Requires the &#x60;fields&#x60; parameter to be specified as well. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) OutputFields(outputFields string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.outputFields = &outputFields
	return r
}

// A comma-separated list of fields which should be omitted from the items.
func (r ApiActRunSyncGetDatasetItemsGetRequest) Omit(omit string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.omit = &omit
	return r
}

// A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the &#x60;desc&#x60; parameter. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Unwind(unwind string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.unwind = &unwind
	return r
}

// A comma-separated list of fields which should transform nested objects into flat structures.  For example, with &#x60;flatten&#x3D;\&quot;foo\&quot;&#x60; the object &#x60;{\&quot;foo\&quot;:{\&quot;bar\&quot;: \&quot;hello\&quot;}}&#x60; is turned into &#x60;{\&quot;foo.bar\&quot;: \&quot;hello\&quot;}&#x60;.  The original object with properties is replaced with the flattened object. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Flatten(flatten string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.flatten = &flatten
	return r
}

// By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to &#x60;true&#x60; or &#x60;1&#x60;. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Desc(desc bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.desc = &desc
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the response will define the &#x60;Content-Disposition: attachment&#x60; header, forcing a web browser to download the file rather than to display it. By default this header is not present. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Attachment(attachment bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.attachment = &attachment
	return r
}

// A delimiter character for CSV files, only used if &#x60;format&#x3D;csv&#x60;. You might need to URL-encode the character (e.g. use &#x60;%09&#x60; for tab or &#x60;%3B&#x60; for semicolon). The default delimiter is a simple comma (&#x60;,&#x60;). 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Delimiter(delimiter string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.delimiter = &delimiter
	return r
}

// All text responses are encoded in UTF-8 encoding. By default, the &#x60;format&#x3D;csv&#x60; files are prefixed with the UTF-8 Byte Order Mark (BOM), while &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;xml&#x60;, &#x60;html&#x60; and &#x60;rss&#x60; files are not.  If you want to override this default behavior, specify &#x60;bom&#x3D;1&#x60; query parameter to include the BOM or &#x60;bom&#x3D;0&#x60; to skip it. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Bom(bom bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.bom = &bom
	return r
}

// Overrides default root element name of &#x60;xml&#x60; output. By default the root element is &#x60;items&#x60;. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) XmlRoot(xmlRoot string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.xmlRoot = &xmlRoot
	return r
}

// Overrides default element name that wraps each page or page function result object in &#x60;xml&#x60; output. By default the element name is &#x60;item&#x60;. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) XmlRow(xmlRow string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.xmlRow = &xmlRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then header row in the &#x60;csv&#x60; format is skipped.
func (r ApiActRunSyncGetDatasetItemsGetRequest) SkipHeaderRow(skipHeaderRow bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.skipHeaderRow = &skipHeaderRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then hidden fields are skipped from the output, i.e. fields starting with the &#x60;#&#x60; character. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) SkipHidden(skipHidden bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.skipHidden = &skipHidden
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) SkipEmpty(skipEmpty bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.skipEmpty = &skipEmpty
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the endpoint applies the &#x60;fields&#x3D;url,pageFunctionResult,errorInfo&#x60; and &#x60;unwind&#x3D;pageFunctionResult&#x60; query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) Simplified(simplified bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.simplified = &simplified
	return r
}

// Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema). 
func (r ApiActRunSyncGetDatasetItemsGetRequest) View(view string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.view = &view
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) SkipFailedPages(skipFailedPages bool) ApiActRunSyncGetDatasetItemsGetRequest {
	r.skipFailedPages = &skipFailedPages
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;title&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the title defaults to &#x60;Dataset &lt;label&gt;&#x60;. 
func (r ApiActRunSyncGetDatasetItemsGetRequest) FeedTitle(feedTitle string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.feedTitle = &feedTitle
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;description&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the description defaults to &#x60;Items in dataset with id \&quot;&lt;datasetId&gt;\&quot;.&#x60; 
func (r ApiActRunSyncGetDatasetItemsGetRequest) FeedDescription(feedDescription string) ApiActRunSyncGetDatasetItemsGetRequest {
	r.feedDescription = &feedDescription
	return r
}

func (r ApiActRunSyncGetDatasetItemsGetRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActRunSyncGetDatasetItemsGetExecute(r)
}

/*
ActRunSyncGetDatasetItemsGet Run Actor synchronously without input and get dataset items

Runs a specific Actor and returns its dataset items.
The run must finish in 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds
otherwise the API endpoint returns a timeout error.
The Actor is not passed any input.

It allows to send all possible options in parameters from [Get Dataset
Items](#/reference/datasets/item-collection/get-items) API endpoint.

Beware that it might be impossible to maintain an idle HTTP connection for a
long period of time,
due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.
If the connection breaks, you will not receive any information about the run
and its status.

To run the Actor asynchronously, use the [Run
Actor](#/reference/actors/run-collection/run-actor) API endpoint instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunSyncGetDatasetItemsGetRequest
*/
func (a *ActorsActorRunsAPIService) ActRunSyncGetDatasetItemsGet(ctx context.Context, actorId string) ApiActRunSyncGetDatasetItemsGetRequest {
	return ApiActRunSyncGetDatasetItemsGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ActorsActorRunsAPIService) ActRunSyncGetDatasetItemsGetExecute(r ApiActRunSyncGetDatasetItemsGetRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunSyncGetDatasetItemsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/run-sync-get-dataset-items"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

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
			var v ActorRunFailedError
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
		if localVarHTTPResponse.StatusCode == 408 {
			var v ActorRunTimeoutExceededError
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

type ApiActRunSyncGetDatasetItemsPostRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
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
func (r ApiActRunSyncGetDatasetItemsPostRequest) Body(body map[string]interface{}) ApiActRunSyncGetDatasetItemsPostRequest {
	r.body = &body
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Timeout(timeout float64) ApiActRunSyncGetDatasetItemsPostRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Memory(memory float64) ApiActRunSyncGetDatasetItemsPostRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) MaxItems(maxItems float64) ApiActRunSyncGetDatasetItemsPostRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActRunSyncGetDatasetItemsPostRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) RestartOnError(restartOnError bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Build(build string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.build = &build
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Webhooks(webhooks string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.webhooks = &webhooks
	return r
}

// Format of the results, possible values are: &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;csv&#x60;, &#x60;html&#x60;, &#x60;xlsx&#x60;, &#x60;xml&#x60; and &#x60;rss&#x60;. The default value is &#x60;json&#x60;. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Format(format string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.format = &format
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The &#x60;clean&#x60; parameter is just a shortcut for &#x60;skipHidden&#x3D;true&#x60; and &#x60;skipEmpty&#x3D;true&#x60; parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the &#x60;limit&#x60; value. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Clean(clean bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.clean = &clean
	return r
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Offset(offset float64) ApiActRunSyncGetDatasetItemsPostRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. By default there is no limit.
func (r ApiActRunSyncGetDatasetItemsPostRequest) Limit(limit float64) ApiActRunSyncGetDatasetItemsPostRequest {
	r.limit = &limit
	return r
}

// A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the &#x60;fields&#x60; query parameter. You can use this feature to effectively fix the output format. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Fields(fields string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of output field names that positionally rename the fields specified in the &#x60;fields&#x60; parameter. For example, &#x60;?fields&#x3D;headline,url&amp;outputFields&#x3D;title,link&#x60; renames &#x60;headline&#x60; to &#x60;title&#x60; and &#x60;url&#x60; to &#x60;link&#x60; in the output. The number of names in &#x60;outputFields&#x60; must match the number of names in &#x60;fields&#x60;. Requires the &#x60;fields&#x60; parameter to be specified as well. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) OutputFields(outputFields string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.outputFields = &outputFields
	return r
}

// A comma-separated list of fields which should be omitted from the items.
func (r ApiActRunSyncGetDatasetItemsPostRequest) Omit(omit string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.omit = &omit
	return r
}

// A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the &#x60;desc&#x60; parameter. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Unwind(unwind string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.unwind = &unwind
	return r
}

// A comma-separated list of fields which should transform nested objects into flat structures.  For example, with &#x60;flatten&#x3D;\&quot;foo\&quot;&#x60; the object &#x60;{\&quot;foo\&quot;:{\&quot;bar\&quot;: \&quot;hello\&quot;}}&#x60; is turned into &#x60;{\&quot;foo.bar\&quot;: \&quot;hello\&quot;}&#x60;.  The original object with properties is replaced with the flattened object. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Flatten(flatten string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.flatten = &flatten
	return r
}

// By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to &#x60;true&#x60; or &#x60;1&#x60;. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Desc(desc bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.desc = &desc
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the response will define the &#x60;Content-Disposition: attachment&#x60; header, forcing a web browser to download the file rather than to display it. By default this header is not present. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Attachment(attachment bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.attachment = &attachment
	return r
}

// A delimiter character for CSV files, only used if &#x60;format&#x3D;csv&#x60;. You might need to URL-encode the character (e.g. use &#x60;%09&#x60; for tab or &#x60;%3B&#x60; for semicolon). The default delimiter is a simple comma (&#x60;,&#x60;). 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Delimiter(delimiter string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.delimiter = &delimiter
	return r
}

// All text responses are encoded in UTF-8 encoding. By default, the &#x60;format&#x3D;csv&#x60; files are prefixed with the UTF-8 Byte Order Mark (BOM), while &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;xml&#x60;, &#x60;html&#x60; and &#x60;rss&#x60; files are not.  If you want to override this default behavior, specify &#x60;bom&#x3D;1&#x60; query parameter to include the BOM or &#x60;bom&#x3D;0&#x60; to skip it. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Bom(bom bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.bom = &bom
	return r
}

// Overrides default root element name of &#x60;xml&#x60; output. By default the root element is &#x60;items&#x60;. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) XmlRoot(xmlRoot string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.xmlRoot = &xmlRoot
	return r
}

// Overrides default element name that wraps each page or page function result object in &#x60;xml&#x60; output. By default the element name is &#x60;item&#x60;. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) XmlRow(xmlRow string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.xmlRow = &xmlRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then header row in the &#x60;csv&#x60; format is skipped.
func (r ApiActRunSyncGetDatasetItemsPostRequest) SkipHeaderRow(skipHeaderRow bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.skipHeaderRow = &skipHeaderRow
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then hidden fields are skipped from the output, i.e. fields starting with the &#x60;#&#x60; character. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) SkipHidden(skipHidden bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.skipHidden = &skipHidden
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) SkipEmpty(skipEmpty bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.skipEmpty = &skipEmpty
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the endpoint applies the &#x60;fields&#x3D;url,pageFunctionResult,errorInfo&#x60; and &#x60;unwind&#x3D;pageFunctionResult&#x60; query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) Simplified(simplified bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.simplified = &simplified
	return r
}

// Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema). 
func (r ApiActRunSyncGetDatasetItemsPostRequest) View(view string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.view = &view
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) SkipFailedPages(skipFailedPages bool) ApiActRunSyncGetDatasetItemsPostRequest {
	r.skipFailedPages = &skipFailedPages
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;title&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the title defaults to &#x60;Dataset &lt;label&gt;&#x60;. 
func (r ApiActRunSyncGetDatasetItemsPostRequest) FeedTitle(feedTitle string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.feedTitle = &feedTitle
	return r
}

// Overrides the auto-generated RSS channel &#x60;&lt;description&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the description defaults to &#x60;Items in dataset with id \&quot;&lt;datasetId&gt;\&quot;.&#x60; 
func (r ApiActRunSyncGetDatasetItemsPostRequest) FeedDescription(feedDescription string) ApiActRunSyncGetDatasetItemsPostRequest {
	r.feedDescription = &feedDescription
	return r
}

func (r ApiActRunSyncGetDatasetItemsPostRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActRunSyncGetDatasetItemsPostExecute(r)
}

/*
ActRunSyncGetDatasetItemsPost Run Actor synchronously and get dataset items

Runs a specific Actor and returns its dataset items.

The POST payload including its `Content-Type` header is passed as `INPUT` to
the Actor (usually `application/json`).
The HTTP response contains the Actors dataset items, while the format of
items depends on specifying dataset items' `format` parameter.

You can send all the same options in parameters as the [Get Dataset
Items](#/reference/datasets/item-collection/get-items) API endpoint.

The Actor is started with the default options; you can override them using
URL query parameters.
If the Actor run exceeds 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds,
the HTTP response will return the 408 status code (Request Timeout).

Beware that it might be impossible to maintain an idle HTTP connection for a
long period of time,
due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.
If the connection breaks, you will not receive any information about the run
and its status.

To run the Actor asynchronously, use the [Run
Actor](#/reference/actors/run-collection/run-actor) API endpoint instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunSyncGetDatasetItemsPostRequest
*/
func (a *ActorsActorRunsAPIService) ActRunSyncGetDatasetItemsPost(ctx context.Context, actorId string) ApiActRunSyncGetDatasetItemsPostRequest {
	return ApiActRunSyncGetDatasetItemsPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ActorsActorRunsAPIService) ActRunSyncGetDatasetItemsPostExecute(r ApiActRunSyncGetDatasetItemsPostRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunSyncGetDatasetItemsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/run-sync-get-dataset-items"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

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
			var v ActorRunFailedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
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
		if localVarHTTPResponse.StatusCode == 408 {
			var v ActorRunTimeoutExceededError
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

type ApiActRunSyncPostRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	body *map[string]interface{}
	outputRecordKey *string
	timeout *float64
	memory *float64
	maxItems *float64
	maxTotalChargeUsd *float64
	restartOnError *bool
	build *string
	webhooks *string
}

// 
func (r ApiActRunSyncPostRequest) Body(body map[string]interface{}) ApiActRunSyncPostRequest {
	r.body = &body
	return r
}

// Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;. 
func (r ApiActRunSyncPostRequest) OutputRecordKey(outputRecordKey string) ApiActRunSyncPostRequest {
	r.outputRecordKey = &outputRecordKey
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActRunSyncPostRequest) Timeout(timeout float64) ApiActRunSyncPostRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActRunSyncPostRequest) Memory(memory float64) ApiActRunSyncPostRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActRunSyncPostRequest) MaxItems(maxItems float64) ApiActRunSyncPostRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActRunSyncPostRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActRunSyncPostRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActRunSyncPostRequest) RestartOnError(restartOnError bool) ApiActRunSyncPostRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActRunSyncPostRequest) Build(build string) ApiActRunSyncPostRequest {
	r.build = &build
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActRunSyncPostRequest) Webhooks(webhooks string) ApiActRunSyncPostRequest {
	r.webhooks = &webhooks
	return r
}

func (r ApiActRunSyncPostRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ActRunSyncPostExecute(r)
}

/*
ActRunSyncPost Run Actor synchronously and return output

Runs a specific Actor and returns its output.

The POST payload including its `Content-Type` header is passed as `INPUT` to
the Actor (usually <code>application/json</code>).
The HTTP response contains Actors `OUTPUT` record from its default
key-value store.

The Actor is started with the default options; you can override them using
various URL query parameters.
If the Actor run exceeds 300<!-- MAX_ACTOR_JOB_SYNC_WAIT_SECS --> seconds,
the HTTP response will have status 408 (Request Timeout).

Beware that it might be impossible to maintain an idle HTTP connection for a
long period of time, due to client timeout or network conditions. Make sure your HTTP client is
configured to have a long enough connection timeout.
If the connection breaks, you will not receive any information about the run
and its status.

To run the Actor asynchronously, use the [Run
Actor](#/reference/actors/run-collection/run-actor) API endpoint instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunSyncPostRequest
*/
func (a *ActorsActorRunsAPIService) ActRunSyncPost(ctx context.Context, actorId string) ApiActRunSyncPostRequest {
	return ApiActRunSyncPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActorsActorRunsAPIService) ActRunSyncPostExecute(r ApiActRunSyncPostRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunSyncPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/run-sync"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.outputRecordKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputRecordKey", r.outputRecordKey, "form", "")
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
			var v ActorRunFailedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
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
		if localVarHTTPResponse.StatusCode == 408 {
			var v ActorRunTimeoutExceededError
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

type ApiActRunsGetRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	offset *float64
	limit *float64
	desc *bool
	status *[]string
	startedAfter *time.Time
	startedBefore *time.Time
}

// Number of items that should be skipped at the start. The default value is &#x60;0&#x60;. 
func (r ApiActRunsGetRequest) Offset(offset float64) ApiActRunsGetRequest {
	r.offset = &offset
	return r
}

// Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;. 
func (r ApiActRunsGetRequest) Limit(limit float64) ApiActRunsGetRequest {
	r.limit = &limit
	return r
}

// If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;startedAt&#x60; field in descending order. By default, they are sorted in ascending order. 
func (r ApiActRunsGetRequest) Desc(desc bool) ApiActRunsGetRequest {
	r.desc = &desc
	return r
}

// Single status or comma-separated list of statuses, see ([available statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)). Used to filter runs by the specified statuses only. 
func (r ApiActRunsGetRequest) Status(status []string) ApiActRunsGetRequest {
	r.status = &status
	return r
}

// Filter runs that started after the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC). 
func (r ApiActRunsGetRequest) StartedAfter(startedAfter time.Time) ApiActRunsGetRequest {
	r.startedAfter = &startedAfter
	return r
}

// Filter runs that started before the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC). 
func (r ApiActRunsGetRequest) StartedBefore(startedBefore time.Time) ApiActRunsGetRequest {
	r.startedBefore = &startedBefore
	return r
}

func (r ApiActRunsGetRequest) Execute() (*ListOfRunsResponse, *http.Response, error) {
	return r.ApiService.ActRunsGetExecute(r)
}

/*
ActRunsGet Get list of runs

Gets the list of runs of a specific Actor. The response is a list of
objects, where each object contains basic information about a single Actor run.

The endpoint supports pagination using the `limit` and `offset` parameters
and it will not return more than 1000 array elements.

By default, the records are sorted by the `startedAt` field in ascending
order, therefore you can use pagination to incrementally fetch all records while
new ones are still being created. To sort the records in descending order, use
`desc=1` parameter. You can also filter runs by status ([available
statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsGetRequest
*/
func (a *ActorsActorRunsAPIService) ActRunsGet(ctx context.Context, actorId string) ApiActRunsGetRequest {
	return ApiActRunsGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return ListOfRunsResponse
func (a *ActorsActorRunsAPIService) ActRunsGetExecute(r ApiActRunsGetRequest) (*ListOfRunsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOfRunsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

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
	if r.startedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startedAfter", r.startedAfter, "form", "")
	}
	if r.startedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startedBefore", r.startedBefore, "form", "")
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

type ApiActRunsLastGetRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	status *string
	waitForFinish *float64
}

// Filter for the run status.
func (r ApiActRunsLastGetRequest) Status(status string) ApiActRunsLastGetRequest {
	r.status = &status
	return r
}

// The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;). 
func (r ApiActRunsLastGetRequest) WaitForFinish(waitForFinish float64) ApiActRunsLastGetRequest {
	r.waitForFinish = &waitForFinish
	return r
}

func (r ApiActRunsLastGetRequest) Execute() (*RunResponse, *http.Response, error) {
	return r.ApiService.ActRunsLastGetExecute(r)
}

/*
ActRunsLastGet Get last run

This is not a single endpoint, but an entire group of endpoints that lets you to
retrieve and manage the last run of given Actor or any of its default storages.
All the endpoints require an authentication token.

The base path represents the last Actor run object is:

`/v2/actors/{actorId}/runs/last{?token,status}`

Using the `status` query parameter you can ensure to only get a run with a certain status
(e.g. `status=SUCCEEDED`). The output of this endpoint and other query parameters
are the same as in the [Run object](#/reference/actors/run-object) endpoint.

##### Convenience endpoints for last Actor run

* [Dataset](/api/v2/last-actor-runs-default-dataset)

* [Key-value store](/api/v2/last-actor-runs-default-key-value-store)

* [Request queue](/api/v2/last-actor-runs-default-request-queue)

* [Log](/api/v2/last-actor-runs-log)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsLastGetRequest
*/
func (a *ActorsActorRunsAPIService) ActRunsLastGet(ctx context.Context, actorId string) ApiActRunsLastGetRequest {
	return ApiActRunsLastGetRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return RunResponse
func (a *ActorsActorRunsAPIService) ActRunsLastGetExecute(r ApiActRunsLastGetRequest) (*RunResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RunResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunsLastGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs/last"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

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

type ApiActRunsPostRequest struct {
	ctx context.Context
	ApiService *ActorsActorRunsAPIService
	actorId string
	body *map[string]interface{}
	timeout *float64
	memory *float64
	maxItems *float64
	maxTotalChargeUsd *float64
	restartOnError *bool
	build *string
	waitForFinish *float64
	webhooks *string
	forcePermissionLevel *string
}

// 
func (r ApiActRunsPostRequest) Body(body map[string]interface{}) ApiActRunsPostRequest {
	r.body = &body
	return r
}

// Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration. 
func (r ApiActRunsPostRequest) Timeout(timeout float64) ApiActRunsPostRequest {
	r.timeout = &timeout
	return r
}

// Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration. 
func (r ApiActRunsPostRequest) Memory(memory float64) ApiActRunsPostRequest {
	r.memory = &memory
	return r
}

// Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable. 
func (r ApiActRunsPostRequest) MaxItems(maxItems float64) ApiActRunsPostRequest {
	r.maxItems = &maxItems
	return r
}

// Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable. 
func (r ApiActRunsPostRequest) MaxTotalChargeUsd(maxTotalChargeUsd float64) ApiActRunsPostRequest {
	r.maxTotalChargeUsd = &maxTotalChargeUsd
	return r
}

// Determines whether the run will be restarted if it fails. 
func (r ApiActRunsPostRequest) RestartOnError(restartOnError bool) ApiActRunsPostRequest {
	r.restartOnError = &restartOnError
	return r
}

// Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;). 
func (r ApiActRunsPostRequest) Build(build string) ApiActRunsPostRequest {
	r.build = &build
	return r
}

// The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;). 
func (r ApiActRunsPostRequest) WaitForFinish(waitForFinish float64) ApiActRunsPostRequest {
	r.waitForFinish = &waitForFinish
	return r
}

// Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks). 
func (r ApiActRunsPostRequest) Webhooks(webhooks string) ApiActRunsPostRequest {
	r.webhooks = &webhooks
	return r
}

// Overrides the Actor&#39;s permission level for this specific run. Use to test restricted permissions before deploying changes to your Actor or to temporarily elevate or restrict access. If you don&#39;t specify this parameter, the Actor uses its configured default permission level. For more information on permissions, see the [documentation](https://docs.apify.com/platform/actors/development/permissions). 
func (r ApiActRunsPostRequest) ForcePermissionLevel(forcePermissionLevel string) ApiActRunsPostRequest {
	r.forcePermissionLevel = &forcePermissionLevel
	return r
}

func (r ApiActRunsPostRequest) Execute() (*RunResponse, *http.Response, error) {
	return r.ApiService.ActRunsPostExecute(r)
}

/*
ActRunsPost Run Actor

Runs an Actor and immediately returns without waiting for the run to finish.

The POST payload including its `Content-Type` header is passed as `INPUT` to
the Actor (usually `application/json`).

The Actor is started with the default options; you can override them using
various URL query parameters.

The response is the Run object as returned by the [Get
run](#/reference/actor-runs/run-object-and-its-storages/get-run) API
endpoint.

If you want to wait for the run to finish and receive the actual output of
the Actor as the response, please use one of the [Run Actor
synchronously](#/reference/actors/run-actor-synchronously) API endpoints
instead.

To fetch the Actor run results that are typically stored in the default
dataset, you'll need to pass the ID received in the `defaultDatasetId` field
received in the response JSON to the [Get dataset items](#/reference/datasets/item-collection/get-items)
API endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID or a tilde-separated owner's username and Actor name.
 @return ApiActRunsPostRequest
*/
func (a *ActorsActorRunsAPIService) ActRunsPost(ctx context.Context, actorId string) ApiActRunsPostRequest {
	return ApiActRunsPostRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return RunResponse
func (a *ActorsActorRunsAPIService) ActRunsPostExecute(r ApiActRunsPostRequest) (*RunResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RunResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActorsActorRunsAPIService.ActRunsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/actors/{actorId}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterValueToString(r.actorId, "actorId")), -1)

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
	if r.forcePermissionLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forcePermissionLevel", r.forcePermissionLevel, "form", "")
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
		if localVarHTTPResponse.StatusCode == 402 {
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
