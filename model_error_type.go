/*
Apify API

 The Apify API (version 2) provides programmatic access to the [Apify platform](https://docs.apify.com). The API is organized around [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) HTTP endpoints.  You can download the complete OpenAPI schema of Apify API in the [YAML](http://docs.apify.com/api/openapi.yaml) or [JSON](http://docs.apify.com/api/openapi.json) formats. The source code is also available on [GitHub](https://github.com/apify/apify-docs/tree/master/apify-api/openapi).  All requests and responses (including errors) are encoded in [JSON](http://www.json.org/) format with UTF-8 encoding, with a few exceptions that are explicitly described in the reference.  - To access the API using [Node.js](https://nodejs.org/en/), we recommend the [`apify-client`](https://docs.apify.com/api/client/js) [NPM package](https://www.npmjs.com/package/apify-client). - To access the API using [Python](https://www.python.org/), we recommend the [`apify-client`](https://docs.apify.com/api/client/python) [PyPI package](https://pypi.org/project/apify-client/).  The clients' functions correspond to the API endpoints and have the same parameters. This simplifies development of apps that depend on the Apify platform.  :::note Important Request Details  - `Content-Type` header: For requests with a JSON body, you must include the `Content-Type: application/json` header.  - Method override: You can override the HTTP method using the `method` query parameter. This is useful for clients that can only send `GET` requests. For example, to call a `POST` endpoint, append `?method=POST` to the URL of your `GET` request.  :::  ## Authentication <span id=\"/introduction/authentication\"></span>  **You can find your API token on the [Integrations](https://console.apify.com/settings/integrations) page in the Apify Console.**  To use your token in a request, either:  - Add the token to your request's `Authorization` header as `Bearer <token>`. E.g., `Authorization: Bearer xxxxxxx`. [More info](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization). (Recommended). - Add it as the `token` parameter to your request URL. (Less secure).  Using your token in the request header is more secure than using it as a URL parameter because URLs are often stored in browser history and server logs. This creates a chance for someone unauthorized to access your API token.  **Never share your API token or password with untrusted parties!**  For more information, see our [integrations](https://docs.apify.com/platform/integrations) documentation.  ### Agentic payments  AI agents can authenticate and pay for Actor runs without an Apify account using agentic payments. Instead of an API token, the request carries a payment credential that both authorizes and pays for the call. Apify supports the [x402 protocol](https://docs.apify.com/platform/integrations/x402) (`PAYMENT-SIGNATURE` header) and [Skyfire](https://docs.apify.com/platform/integrations/skyfire) (`skyfire-pay-id` header).  ## Basic usage <span id=\"/introduction/basic-usage\"></span>  To run an Actor, send a POST request to the [Run Actor](#/reference/actors/run-collection/run-actor) endpoint using either the Actor ID code (e.g. `vKg4IjxZbEYTYeW8T`) or its name (e.g. `janedoe~my-actor`):  `https://api.apify.com/v2/actors/[actor_id]/runs`  If the Actor is not runnable anonymously, you will receive a 401 or 403 [response code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status). This means you need to add your [secret API token](https://console.apify.com/account#/integrations) to the request's `Authorization` header ([recommended](#/introduction/authentication)) or as a URL query parameter `?token=[your_token]` (less secure).  Optionally, you can include the query parameters described in the [Run Actor](#/reference/actors/run-collection/run-actor) section to customize your run.  If you're using Node.js, the best way to run an Actor is using the `Apify.call()` method from the [Apify SDK](https://sdk.apify.com/docs/api/apify#apifycallactid-input-options). It runs the Actor using the account you are currently logged into (determined by the [secret API token](https://console.apify.com/account#/integrations)). The result is an [Actor run object](https://sdk.apify.com/docs/typedefs/actor-run) and its output (if any).  A typical workflow is as follows:  1. Run an Actor or task using the [Run Actor](#/reference/actors/run-collection/run-actor) or [Run task](#/reference/actor-tasks/run-collection/run-task) API endpoints. 2. Monitor the Actor run by periodically polling its progress using the [Get run](#/reference/actor-runs/run-object-and-its-storages/get-run) API endpoint. 3. Fetch the results from the [Get items](#/reference/datasets/item-collection/get-items) API endpoint using the `defaultDatasetId`, which you receive in the Run request response. Additional data may be stored in a key-value store. You can fetch them from the [Get record](#/reference/key-value-stores/record/get-record) API endpoint using the `defaultKeyValueStoreId` and the store's `key`.  **Note**: Instead of periodic polling, you can also run your [Actor](#/reference/actors/run-actor-synchronously) or [task](#/reference/actor-tasks/runs-collection/run-task-synchronously) synchronously. This will ensure that the request waits for 300 seconds (5 minutes) for the run to finish and returns its output. If the run takes longer, the request will time out and throw an error.  ## Legacy `/v2/acts/` URL prefix <span id=\"/introduction/legacy-acts-prefix\"></span>  The `/v2/acts/` prefix is deprecated but still fully functional, and  such endpoint routes to the same handler as its `/v2/actors/...` counterpart.  New integrations should use the canonical /v2/actors/ prefix,  but existing clients keep working without changes.  ## Response structure <span id=\"/introduction/response-structure\"></span>  Most API endpoints return a JSON object with the `data` property:  ``` {     \"data\": {         ...     } } ```  However, there are a few explicitly described exceptions, such as [Get dataset items](#/reference/datasets/item-collection/get-items) or Key-value store [Get record](#/reference/key-value-stores/record/get-record) API endpoints, which return data in other formats. In case of an error, the response has the HTTP status code in the range of 4xx or 5xx and the `data` property is replaced with `error`. For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  See [Errors](#/introduction/errors) for more details.  ## Pagination <span id=\"/introduction/pagination\"></span>  All API endpoints that return a list of records (e.g. [Get list of Actors](#/reference/actors/actor-collection/get-list-of-actors)) enforce pagination in order to limit the size of their responses.  Most of these API endpoints are paginated using the `offset` and `limit` query parameters. The only exception is [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys), which is paginated using the `exclusiveStartKey` query parameter.  **IMPORTANT**: Each API endpoint that supports pagination enforces a certain maximum value for the `limit` parameter, in order to reduce the load on Apify servers. The maximum limit could change in future so you should never rely on a specific value and check the responses of these API endpoints.  ### Using offset <span id=\"/introduction/pagination/using-offset\"></span>  Most API endpoints that return a list of records enable pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number of items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>Skips a number of items from the beginning of the list, e.g. <code>offset=100</code>.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td>     By default, items are sorted in the order in which they were created or added to the list.     This feature is useful when fetching all the items, because it ensures that items     created after the client started the pagination will not be skipped.     If you specify the <code>desc=1</code> parameter, the items will be returned in the reverse order,     i.e. from the newest to the oldest items.     </td>   </tr> </table>  The response of these API endpoints is always a JSON object with the following structure:  ``` {     \"data\": {         \"total\": 2560,         \"offset\": 250,         \"limit\": 1000,         \"count\": 1000,         \"desc\": false,         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>total</code></td>     <td>The total number of items available in the list.</td>   </tr>   <tr>     <td><code>offset</code></td>     <td>The number of items that were skipped at the start.     This is equal to the <code>offset</code> query parameter if it was provided, otherwise it is <code>0</code>.</td>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular API endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>count</code></td>     <td>The actual number of items returned in the HTTP response.</td>   </tr>   <tr>     <td><code>desc</code></td>     <td><code>true</code> if data were requested in descending order and <code>false</code> otherwise.</td>   </tr>   <tr>     <td><code>items</code></td>     <td>An array of requested items.</td>   </tr> </table>  ### Using key <span id=\"/introduction/pagination/using-key\"></span>  The records in the [key-value store](https://docs.apify.com/platform/storage/key-value-store) are not ordered based on numerical indexes, but rather by their keys in the UTF-8 binary order. Therefore the [Get list of keys](#/reference/key-value-stores/key-collection/get-list-of-keys) API endpoint only supports pagination using the following query parameters:  <table>   <tr>     <td><code>limit</code></td>     <td>Limits the response to contain a specific maximum number items, e.g. <code>limit=20</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>Skips all records with keys up to the given key including the given key,     in the UTF-8 binary order.</td>   </tr> </table>  The response of the API endpoint is always a JSON object with following structure:  ``` {     \"data\": {         \"limit\": 1000,         \"isTruncated\": true,         \"exclusiveStartKey\": \"my-key\",         \"nextExclusiveStartKey\": \"some-other-key\",         \"items\": [             { 1st object },             { 2nd object },             ...             { 1000th object }         ]     } } ```  The following table describes the meaning of the response properties:  <table>   <tr>     <th>Property</th>     <th>Description</th>   </tr>   <tr>     <td><code>limit</code></td>     <td>The maximum number of items that can be returned in the HTTP response.     It equals to the <code>limit</code> query parameter if it was provided or     the maximum limit enforced for the particular endpoint, whichever is smaller.</td>   </tr>   <tr>     <td><code>isTruncated</code></td>     <td><code>true</code> if there are more items left to be queried. Otherwise <code>false</code>.</td>   </tr>   <tr>     <td><code>exclusiveStartKey</code></td>     <td>The last key that was skipped at the start. Is `null` for the first page.</td>   </tr>   <tr>     <td><code>nextExclusiveStartKey</code></td>     <td>The value for the <code>exclusiveStartKey</code> parameter to query the next page of items.</td>   </tr> </table>  ## Errors <span id=\"/introduction/errors\"></span>  The Apify API uses common HTTP status codes: `2xx` range for success, `4xx` range for errors caused by the caller (invalid requests) and `5xx` range for server errors (these are rare). Each error response contains a JSON object defining the `error` property, which is an object with the `type` and `message` properties that contain the error code and a human-readable error description, respectively.  For example:  ``` {     \"error\": {         \"type\": \"record-not-found\",         \"message\": \"Store was not found.\"     } } ```  Here is the table of the most common errors that can occur for many API endpoints:  <table>   <tr>     <th>status</th>     <th>type</th>     <th>message</th>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-request</code></td>     <td>POST data must be a JSON object</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-value</code></td>     <td>Invalid value provided: Comments required</td>   </tr>   <tr>     <td><code>400</code></td>     <td><code>invalid-record-key</code></td>     <td>Record key contains invalid character</td>   </tr>   <tr>     <td><code>401</code></td>     <td><code>token-not-provided</code></td>     <td>Authentication token was not provided</td>   </tr>   <tr>     <td><code>404</code></td>     <td><code>record-not-found</code></td>     <td>Store was not found</td>   </tr>   <tr>     <td><code>429</code></td>     <td><code>rate-limit-exceeded</code></td>     <td>You have exceeded the rate limit of ... requests per second</td>   </tr>   <tr>     <td><code>405</code></td>     <td><code>method-not-allowed</code></td>     <td>This API endpoint can only be accessed using the following HTTP methods: OPTIONS, POST</td>   </tr> </table>  ## Rate limiting <span id=\"/introduction/rate-limiting\"></span>  All API endpoints limit the rate of requests in order to prevent overloading of Apify servers by misbehaving clients.  There are two kinds of rate limits - a global rate limit and a per-resource rate limit.  ### Global rate limit <span id=\"/introduction/rate-limiting/global-rate-limit\"></span>  The global rate limit is set to _250 000 requests per minute_. For [authenticated](#/introduction/authentication) requests, it is counted per user, and for unauthenticated requests, it is counted per IP address.  ### Per-resource rate limit <span id=\"/introduction/rate-limiting/per-resource-rate-limit\"></span>  The default per-resource rate limit is _60 requests per second per resource_, which in this context means a single Actor, a single Actor run, a single dataset, single key-value store etc. The default rate limit is applied to every API endpoint except a few select ones, which have higher rate limits. Each API endpoint returns its rate limit in `X-RateLimit-Limit` header.  These endpoints have a rate limit of _200 requests per second per resource_:  * CRUD ([get](#/reference/key-value-stores/record/get-record),   [put](#/reference/key-value-stores/record/put-record),   [delete](#/reference/key-value-stores/record/delete-record))   operations on key-value store records  These endpoints have a rate limit of _400 requests per second per resource_: * [Run Actor](#/reference/actors/run-collection/run-actor) * [Run Actor task asynchronously](#/reference/actor-tasks/runs-collection/run-task-asynchronously) * [Run Actor task synchronously](#/reference/actor-tasks/runs-collection/run-task-synchronously) * [Metamorph Actor run](#/reference/actors/metamorph-run/metamorph-run) * [Push items](#/reference/datasets/item-collection/put-items) to dataset * CRUD   ([add](#/reference/request-queues/request-collection/add-request),   [get](#/reference/request-queues/request-collection/get-request),   [update](#/reference/request-queues/request-collection/update-request),   [delete](#/reference/request-queues/request-collection/delete-request))   operations on requests in request queues  ### Rate limit exceeded errors <span id=\"/introduction/rate-limiting/rate-limit-exceeded-errors\"></span>  If the client is sending too many requests, the API endpoints respond with the HTTP status code `429 Too Many Requests` and the following body:  ``` {     \"error\": {         \"type\": \"rate-limit-exceeded\",         \"message\": \"You have exceeded the rate limit of ... requests per second\"     } } ```  ### Retrying rate-limited requests with exponential backoff <span id=\"/introduction/rate-limiting/retrying-rate-limited-requests-with-exponential-backoff\"></span>  If the client receives the rate limit error, it should wait a certain period of time and then retry the request. If the error happens again, the client should double the wait period and retry the request, and so on. This algorithm is known as _exponential backoff_ and it can be described using the following pseudo-code:  1. Define a variable `DELAY=500` 2. Send the HTTP request to the API endpoint 3. If the response has status code not equal to `429` then you are done. Otherwise:    * Wait for a period of time chosen randomly from the interval `DELAY` to `2*DELAY` milliseconds    * Double the future wait period by setting `DELAY = 2*DELAY`    * Continue with step 2  If all requests sent by the client implement the above steps, the client will automatically use the maximum available bandwidth for its requests.  Note that the Apify API clients [for JavaScript](https://docs.apify.com/api/client/js) and [for Python](https://docs.apify.com/api/client/python) use the exponential backoff algorithm transparently, so that you do not need to worry about it.  ## Referring to resources <span id=\"/introduction/referring-to-resources\"></span>  There are three main ways to refer to a resource you're accessing via API.  - the resource ID (e.g. `iKkPcIgVvwmztduf8`) - `username~resourcename` - when using this access method, you will need to use your API token, and access will only work if you have the correct permissions. - `~resourcename` - for this, you need to use an API token, and the `resourcename` refers to a resource in the API token owner's account. 

API version: v2-2026-06-16T064758Z
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apifyclient

import (
	"encoding/json"
	"fmt"
)

// ErrorType Machine-processable error type identifier.
type ErrorType string

// List of ErrorType
const (
	_3D_SECURE_AUTH_FAILED ErrorType = "3d-secure-auth-failed"
	ACCESS_RIGHT_ALREADY_EXISTS ErrorType = "access-right-already-exists"
	ACTION_NOT_FOUND ErrorType = "action-not-found"
	ACTOR_ALREADY_RENTED ErrorType = "actor-already-rented"
	ACTOR_CAN_NOT_BE_RENTED ErrorType = "actor-can-not-be-rented"
	ACTOR_DISABLED ErrorType = "actor-disabled"
	ACTOR_IS_NOT_RENTED ErrorType = "actor-is-not-rented"
	ACTOR_MEMORY_LIMIT_EXCEEDED ErrorType = "actor-memory-limit-exceeded"
	ACTOR_NAME_EXISTS_NEW_OWNER ErrorType = "actor-name-exists-new-owner"
	ACTOR_NAME_NOT_UNIQUE ErrorType = "actor-name-not-unique"
	ACTOR_NOT_FOUND ErrorType = "actor-not-found"
	ACTOR_NOT_GITHUB_ACTOR ErrorType = "actor-not-github-actor"
	ACTOR_NOT_PUBLIC ErrorType = "actor-not-public"
	ACTOR_PERMISSION_LEVEL_NOT_SUPPORTED_FOR_AGENTIC_PAYMENTS ErrorType = "actor-permission-level-not-supported-for-agentic-payments"
	ACTOR_REVIEW_ALREADY_EXISTS ErrorType = "actor-review-already-exists"
	ACTOR_RUN_FAILED ErrorType = "actor-run-failed"
	ACTOR_STANDBY_NOT_SUPPORTED_FOR_AGENTIC_PAYMENTS ErrorType = "actor-standby-not-supported-for-agentic-payments"
	ACTOR_TASK_NAME_NOT_UNIQUE ErrorType = "actor-task-name-not-unique"
	AGENTIC_PAYMENT_INFO_RETRIEVAL_ERROR ErrorType = "agentic-payment-info-retrieval-error"
	AGENTIC_PAYMENT_INFORMATION_MISSING ErrorType = "agentic-payment-information-missing"
	AGENTIC_PAYMENT_INSUFFICIENT_AMOUNT ErrorType = "agentic-payment-insufficient-amount"
	AGENTIC_PAYMENT_PROVIDER_INTERNAL_ERROR ErrorType = "agentic-payment-provider-internal-error"
	AGENTIC_PAYMENT_PROVIDER_UNAUTHORIZED ErrorType = "agentic-payment-provider-unauthorized"
	AIRTABLE_WEBHOOK_DEPRECATED ErrorType = "airtable-webhook-deprecated"
	ALREADY_SUBSCRIBED_TO_PAID_ACTOR ErrorType = "already-subscribed-to-paid-actor"
	APIFY_PLAN_REQUIRED_TO_USE_PAID_ACTOR ErrorType = "apify-plan-required-to-use-paid-actor"
	APIFY_SIGNUP_NOT_ALLOWED ErrorType = "apify-signup-not-allowed"
	AUTH_METHOD_NOT_SUPPORTED ErrorType = "auth-method-not-supported"
	AUTHORIZATION_SERVER_NOT_FOUND ErrorType = "authorization-server-not-found"
	AUTO_ISSUE_DATE_INVALID ErrorType = "auto-issue-date-invalid"
	BACKGROUND_CHECK_REQUIRED ErrorType = "background-check-required"
	BILLING_SYSTEM_ERROR ErrorType = "billing-system-error"
	BLACK_FRIDAY_PLAN_EXPIRED ErrorType = "black-friday-plan-expired"
	BRAINTREE_ERROR ErrorType = "braintree-error"
	BRAINTREE_NOT_LINKED ErrorType = "braintree-not-linked"
	BRAINTREE_OPERATION_TIMED_OUT ErrorType = "braintree-operation-timed-out"
	BRAINTREE_UNSUPPORTED_CURRENCY ErrorType = "braintree-unsupported-currency"
	BUILD_NOT_FOUND ErrorType = "build-not-found"
	BUILD_OUTDATED ErrorType = "build-outdated"
	CANNOT_ADD_APIFY_EVENTS_TO_PPE_ACTOR ErrorType = "cannot-add-apify-events-to-ppe-actor"
	CANNOT_ADD_MULTIPLE_PRICING_INFOS ErrorType = "cannot-add-multiple-pricing-infos"
	CANNOT_ADD_PRICING_INFO_THAT_ALTERS_PAST ErrorType = "cannot-add-pricing-info-that-alters-past"
	CANNOT_ADD_SECOND_FUTURE_PRICING_INFO ErrorType = "cannot-add-second-future-pricing-info"
	CANNOT_BUILD_ACTOR_FROM_WEBHOOK ErrorType = "cannot-build-actor-from-webhook"
	CANNOT_CHANGE_BILLING_INTERVAL ErrorType = "cannot-change-billing-interval"
	CANNOT_CHANGE_OWNER ErrorType = "cannot-change-owner"
	CANNOT_CHARGE_APIFY_EVENT ErrorType = "cannot-charge-apify-event"
	CANNOT_CHARGE_NON_PAY_PER_EVENT_ACTOR ErrorType = "cannot-charge-non-pay-per-event-actor"
	CANNOT_COMMENT_AS_OTHER_USER ErrorType = "cannot-comment-as-other-user"
	CANNOT_COPY_ACTOR_TASK ErrorType = "cannot-copy-actor-task"
	CANNOT_CREATE_PAYOUT ErrorType = "cannot-create-payout"
	CANNOT_CREATE_PUBLIC_ACTOR ErrorType = "cannot-create-public-actor"
	CANNOT_CREATE_TAX_TRANSACTION ErrorType = "cannot-create-tax-transaction"
	CANNOT_DELETE_CRITICAL_ACTOR ErrorType = "cannot-delete-critical-actor"
	CANNOT_DELETE_INVOICE ErrorType = "cannot-delete-invoice"
	CANNOT_DELETE_PAID_ACTOR ErrorType = "cannot-delete-paid-actor"
	CANNOT_DISABLE_ONE_TIME_EVENT_FOR_APIFY_START_EVENT ErrorType = "cannot-disable-one-time-event-for-apify-start-event"
	CANNOT_DISABLE_ORGANIZATION_WITH_ENABLED_MEMBERS ErrorType = "cannot-disable-organization-with-enabled-members"
	CANNOT_DISABLE_USER_WITH_SUBSCRIPTION ErrorType = "cannot-disable-user-with-subscription"
	CANNOT_LINK_OAUTH_TO_UNVERIFIED_EMAIL ErrorType = "cannot-link-oauth-to-unverified-email"
	CANNOT_METAMORPH_TO_PAY_PER_RESULT_ACTOR ErrorType = "cannot-metamorph-to-pay-per-result-actor"
	CANNOT_MODIFY_ACTOR_PRICING_TOO_FREQUENTLY ErrorType = "cannot-modify-actor-pricing-too-frequently"
	CANNOT_MODIFY_ACTOR_PRICING_WITH_IMMEDIATE_EFFECT ErrorType = "cannot-modify-actor-pricing-with-immediate-effect"
	CANNOT_OVERRIDE_PAID_ACTOR_TRIAL ErrorType = "cannot-override-paid-actor-trial"
	CANNOT_PERMANENTLY_DELETE_SUBSCRIPTION ErrorType = "cannot-permanently-delete-subscription"
	CANNOT_PUBLISH_ACTOR ErrorType = "cannot-publish-actor"
	CANNOT_REDUCE_LAST_FULL_TOKEN ErrorType = "cannot-reduce-last-full-token"
	CANNOT_REIMBURSE_MORE_THAN_ORIGINAL_CHARGE ErrorType = "cannot-reimburse-more-than-original-charge"
	CANNOT_REIMBURSE_NON_RENTAL_CHARGE ErrorType = "cannot-reimburse-non-rental-charge"
	CANNOT_REMOVE_OWN_ACTOR_FROM_RECENTLY_USED ErrorType = "cannot-remove-own-actor-from-recently-used"
	CANNOT_REMOVE_PAYMENT_METHOD ErrorType = "cannot-remove-payment-method"
	CANNOT_REMOVE_PRICING_INFO ErrorType = "cannot-remove-pricing-info"
	CANNOT_REMOVE_RUNNING_RUN ErrorType = "cannot-remove-running-run"
	CANNOT_REMOVE_USER_WITH_PUBLIC_ACTORS ErrorType = "cannot-remove-user-with-public-actors"
	CANNOT_REMOVE_USER_WITH_SUBSCRIPTION ErrorType = "cannot-remove-user-with-subscription"
	CANNOT_REMOVE_USER_WITH_UNPAID_INVOICE ErrorType = "cannot-remove-user-with-unpaid-invoice"
	CANNOT_RENAME_ENV_VAR ErrorType = "cannot-rename-env-var"
	CANNOT_RENT_PAID_ACTOR ErrorType = "cannot-rent-paid-actor"
	CANNOT_REVIEW_OWN_ACTOR ErrorType = "cannot-review-own-actor"
	CANNOT_SET_ACCESS_RIGHTS_FOR_OWNER ErrorType = "cannot-set-access-rights-for-owner"
	CANNOT_SET_IS_STATUS_MESSAGE_TERMINAL ErrorType = "cannot-set-is-status-message-terminal"
	CANNOT_UNPUBLISH_CRITICAL_ACTOR ErrorType = "cannot-unpublish-critical-actor"
	CANNOT_UNPUBLISH_PAID_ACTOR ErrorType = "cannot-unpublish-paid-actor"
	CANNOT_UNPUBLISH_PROFILE ErrorType = "cannot-unpublish-profile"
	CANNOT_UPDATE_INVOICE_FIELD ErrorType = "cannot-update-invoice-field"
	CONCURRENT_RUNS_LIMIT_EXCEEDED ErrorType = "concurrent-runs-limit-exceeded"
	CONCURRENT_UPDATE_DETECTED ErrorType = "concurrent-update-detected"
	CONFERENCE_TOKEN_NOT_FOUND ErrorType = "conference-token-not-found"
	CONTENT_ENCODING_FORBIDDEN_FOR_HTML ErrorType = "content-encoding-forbidden-for-html"
	COUPON_ALREADY_REDEEMED ErrorType = "coupon-already-redeemed"
	COUPON_EXPIRED ErrorType = "coupon-expired"
	COUPON_FOR_NEW_CUSTOMERS ErrorType = "coupon-for-new-customers"
	COUPON_FOR_SUBSCRIBED_USERS ErrorType = "coupon-for-subscribed-users"
	COUPON_LIMITS_ARE_IN_CONFLICT_WITH_CURRENT_LIMITS ErrorType = "coupon-limits-are-in-conflict-with-current-limits"
	COUPON_MAX_NUMBER_OF_REDEMPTIONS_REACHED ErrorType = "coupon-max-number-of-redemptions-reached"
	COUPON_NOT_FOUND ErrorType = "coupon-not-found"
	COUPON_NOT_UNIQUE ErrorType = "coupon-not-unique"
	COUPONS_DISABLED ErrorType = "coupons-disabled"
	CREATE_GITHUB_ISSUE_NOT_ALLOWED ErrorType = "create-github-issue-not-allowed"
	CREATOR_PLAN_NOT_AVAILABLE ErrorType = "creator-plan-not-available"
	CRON_EXPRESSION_INVALID ErrorType = "cron-expression-invalid"
	DAILY_AI_TOKEN_LIMIT_EXCEEDED ErrorType = "daily-ai-token-limit-exceeded"
	DAILY_PUBLICATION_LIMIT_EXCEEDED ErrorType = "daily-publication-limit-exceeded"
	DATASET_DOES_NOT_HAVE_FIELDS_SCHEMA ErrorType = "dataset-does-not-have-fields-schema"
	DATASET_DOES_NOT_HAVE_SCHEMA ErrorType = "dataset-does-not-have-schema"
	DATASET_LOCKED ErrorType = "dataset-locked"
	DATASET_SCHEMA_INVALID ErrorType = "dataset-schema-invalid"
	DCR_NOT_SUPPORTED ErrorType = "dcr-not-supported"
	DEFAULT_DATASET_NOT_FOUND ErrorType = "default-dataset-not-found"
	DELETING_DEFAULT_BUILD ErrorType = "deleting-default-build"
	DELETING_UNFINISHED_BUILD ErrorType = "deleting-unfinished-build"
	EMAIL_ALREADY_TAKEN ErrorType = "email-already-taken"
	EMAIL_ALREADY_TAKEN_REMOVED_USER ErrorType = "email-already-taken-removed-user"
	EMAIL_DOMAIN_NOT_ALLOWED_FOR_COUPON ErrorType = "email-domain-not-allowed-for-coupon"
	EMAIL_INVALID ErrorType = "email-invalid"
	EMAIL_NOT_ALLOWED ErrorType = "email-not-allowed"
	EMAIL_NOT_VALID ErrorType = "email-not-valid"
	EMAIL_UPDATE_TOO_SOON ErrorType = "email-update-too-soon"
	ELEVATED_PERMISSIONS_NEEDED ErrorType = "elevated-permissions-needed"
	ENV_VAR_ALREADY_EXISTS ErrorType = "env-var-already-exists"
	EXCHANGE_RATE_FETCH_FAILED ErrorType = "exchange-rate-fetch-failed"
	EXPIRED_CONFERENCE_TOKEN ErrorType = "expired-conference-token"
	FAILED_TO_CHARGE_USER ErrorType = "failed-to-charge-user"
	FINAL_INVOICE_NEGATIVE ErrorType = "final-invoice-negative"
	FULL_PERMISSION_ACTOR_NOT_APPROVED ErrorType = "full-permission-actor-not-approved"
	GITHUB_BRANCH_EMPTY ErrorType = "github-branch-empty"
	GITHUB_ISSUE_ALREADY_EXISTS ErrorType = "github-issue-already-exists"
	GITHUB_PUBLIC_KEY_NOT_FOUND ErrorType = "github-public-key-not-found"
	GITHUB_REPOSITORY_NOT_FOUND ErrorType = "github-repository-not-found"
	GITHUB_SIGNATURE_DOES_NOT_MATCH_PAYLOAD ErrorType = "github-signature-does-not-match-payload"
	GITHUB_USER_NOT_AUTHORIZED_FOR_ISSUES ErrorType = "github-user-not-authorized-for-issues"
	GMAIL_NOT_ALLOWED ErrorType = "gmail-not-allowed"
	ID_DOES_NOT_MATCH ErrorType = "id-does-not-match"
	INCOMPATIBLE_BILLING_INTERVAL ErrorType = "incompatible-billing-interval"
	INCOMPLETE_PAYOUT_BILLING_INFO ErrorType = "incomplete-payout-billing-info"
	INCONSISTENT_CURRENCIES ErrorType = "inconsistent-currencies"
	INCORRECT_PRICING_MODIFIER_PREFIX ErrorType = "incorrect-pricing-modifier-prefix"
	INPUT_JSON_INVALID_CHARACTERS ErrorType = "input-json-invalid-characters"
	INPUT_JSON_NOT_OBJECT ErrorType = "input-json-not-object"
	INPUT_JSON_TOO_LONG ErrorType = "input-json-too-long"
	INPUT_UPDATE_COLLISION ErrorType = "input-update-collision"
	INSUFFICIENT_PERMISSIONS ErrorType = "insufficient-permissions"
	INSUFFICIENT_PERMISSIONS_TO_CHANGE_FIELD ErrorType = "insufficient-permissions-to-change-field"
	INSUFFICIENT_SECURITY_MEASURES ErrorType = "insufficient-security-measures"
	INSUFFICIENT_TAX_COUNTRY_EVIDENCE ErrorType = "insufficient-tax-country-evidence"
	INTEGRATION_AUTH_ERROR ErrorType = "integration-auth-error"
	INTERNAL_SERVER_ERROR ErrorType = "internal-server-error"
	INVALID_BILLING_INFO ErrorType = "invalid-billing-info"
	INVALID_BILLING_PERIOD_FOR_PAYOUT ErrorType = "invalid-billing-period-for-payout"
	INVALID_BUILD ErrorType = "invalid-build"
	INVALID_CLIENT_KEY ErrorType = "invalid-client-key"
	INVALID_COLLECTION ErrorType = "invalid-collection"
	INVALID_CONFERENCE_LOGIN_PASSWORD ErrorType = "invalid-conference-login-password"
	INVALID_CONTENT_TYPE_HEADER ErrorType = "invalid-content-type-header"
	INVALID_CREDENTIALS ErrorType = "invalid-credentials"
	INVALID_GIT_AUTH_TOKEN ErrorType = "invalid-git-auth-token"
	INVALID_GITHUB_ISSUE_URL ErrorType = "invalid-github-issue-url"
	INVALID_HEADER ErrorType = "invalid-header"
	INVALID_ID ErrorType = "invalid-id"
	INVALID_IDEMPOTENCY_KEY ErrorType = "invalid-idempotency-key"
	INVALID_INPUT ErrorType = "invalid-input"
	INVALID_INPUT_SCHEMA ErrorType = "invalid-input-schema"
	INVALID_INVOICE ErrorType = "invalid-invoice"
	INVALID_INVOICE_TYPE ErrorType = "invalid-invoice-type"
	INVALID_ISSUE_DATE ErrorType = "invalid-issue-date"
	INVALID_LABEL_PARAMS ErrorType = "invalid-label-params"
	INVALID_MAIN_ACCOUNT_USER_ID ErrorType = "invalid-main-account-user-id"
	INVALID_OAUTH_APP ErrorType = "invalid-oauth-app"
	INVALID_OAUTH_SCOPE ErrorType = "invalid-oauth-scope"
	INVALID_ONE_TIME_INVOICE ErrorType = "invalid-one-time-invoice"
	INVALID_PARAMETER ErrorType = "invalid-parameter"
	INVALID_PAYOUT_STATUS ErrorType = "invalid-payout-status"
	INVALID_PICTURE_URL ErrorType = "invalid-picture-url"
	INVALID_RECORD_KEY ErrorType = "invalid-record-key"
	INVALID_REQUEST ErrorType = "invalid-request"
	INVALID_RESOURCE_TYPE ErrorType = "invalid-resource-type"
	INVALID_SIGNATURE ErrorType = "invalid-signature"
	INVALID_SUBSCRIPTION_PLAN ErrorType = "invalid-subscription-plan"
	INVALID_TAX_NUMBER ErrorType = "invalid-tax-number"
	INVALID_TAX_NUMBER_FORMAT ErrorType = "invalid-tax-number-format"
	INVALID_TOKEN ErrorType = "invalid-token"
	INVALID_TOKEN_TYPE ErrorType = "invalid-token-type"
	INVALID_TWO_FACTOR_CODE ErrorType = "invalid-two-factor-code"
	INVALID_TWO_FACTOR_CODE_OR_RECOVERY_CODE ErrorType = "invalid-two-factor-code-or-recovery-code"
	INVALID_TWO_FACTOR_RECOVERY_CODE ErrorType = "invalid-two-factor-recovery-code"
	INVALID_USERNAME ErrorType = "invalid-username"
	INVALID_VALUE ErrorType = "invalid-value"
	INVITATION_INVALID_RESOURCE_TYPE ErrorType = "invitation-invalid-resource-type"
	INVITATION_NO_LONGER_VALID ErrorType = "invitation-no-longer-valid"
	INVOICE_CANCELED ErrorType = "invoice-canceled"
	INVOICE_CANNOT_BE_REFUNDED_DUE_TO_TOO_HIGH_AMOUNT ErrorType = "invoice-cannot-be-refunded-due-to-too-high-amount"
	INVOICE_INCOMPLETE ErrorType = "invoice-incomplete"
	INVOICE_IS_DRAFT ErrorType = "invoice-is-draft"
	INVOICE_LOCKED ErrorType = "invoice-locked"
	INVOICE_MUST_BE_BUFFER ErrorType = "invoice-must-be-buffer"
	INVOICE_NOT_CANCELED ErrorType = "invoice-not-canceled"
	INVOICE_NOT_DRAFT ErrorType = "invoice-not-draft"
	INVOICE_NOT_FOUND ErrorType = "invoice-not-found"
	INVOICE_OUTDATED ErrorType = "invoice-outdated"
	INVOICE_PAID_ALREADY ErrorType = "invoice-paid-already"
	ISSUE_ALREADY_CONNECTED_TO_GITHUB ErrorType = "issue-already-connected-to-github"
	ISSUE_NOT_FOUND ErrorType = "issue-not-found"
	ISSUES_BAD_REQUEST ErrorType = "issues-bad-request"
	ISSUER_NOT_REGISTERED ErrorType = "issuer-not-registered"
	JOB_FINISHED ErrorType = "job-finished"
	LABEL_ALREADY_LINKED ErrorType = "label-already-linked"
	LAST_API_TOKEN ErrorType = "last-api-token"
	LIMIT_REACHED ErrorType = "limit-reached"
	MAX_ITEMS_MUST_BE_GREATER_THAN_ZERO ErrorType = "max-items-must-be-greater-than-zero"
	MAX_METAMORPHS_EXCEEDED ErrorType = "max-metamorphs-exceeded"
	MAX_TOTAL_CHARGE_USD_BELOW_MINIMUM ErrorType = "max-total-charge-usd-below-minimum"
	MAX_TOTAL_CHARGE_USD_MUST_BE_GREATER_THAN_ZERO ErrorType = "max-total-charge-usd-must-be-greater-than-zero"
	METHOD_NOT_ALLOWED ErrorType = "method-not-allowed"
	MIGRATION_DISABLED ErrorType = "migration-disabled"
	MISSING_ACTOR_RIGHTS ErrorType = "missing-actor-rights"
	MISSING_API_TOKEN ErrorType = "missing-api-token"
	MISSING_BILLING_INFO ErrorType = "missing-billing-info"
	MISSING_LINE_ITEMS ErrorType = "missing-line-items"
	MISSING_PAYMENT_DATE ErrorType = "missing-payment-date"
	MISSING_PAYOUT_BILLING_INFO ErrorType = "missing-payout-billing-info"
	MISSING_PROXY_PASSWORD ErrorType = "missing-proxy-password"
	MISSING_REPORTING_FIELDS ErrorType = "missing-reporting-fields"
	MISSING_RESOURCE_NAME ErrorType = "missing-resource-name"
	MISSING_SETTINGS ErrorType = "missing-settings"
	MISSING_USERNAME ErrorType = "missing-username"
	MONTHLY_USAGE_LIMIT_TOO_LOW ErrorType = "monthly-usage-limit-too-low"
	MORE_THAN_ONE_UPDATE_NOT_ALLOWED ErrorType = "more-than-one-update-not-allowed"
	MULTIPLE_RECORDS_FOUND ErrorType = "multiple-records-found"
	MUST_BE_ADMIN ErrorType = "must-be-admin"
	NAME_NOT_UNIQUE ErrorType = "name-not-unique"
	NEXT_RUNTIME_COMPUTATION_FAILED ErrorType = "next-runtime-computation-failed"
	NO_COLUMNS_IN_EXPORTED_DATASET ErrorType = "no-columns-in-exported-dataset"
	NO_PAYMENT_ATTEMPT_FOR_REFUND_FOUND ErrorType = "no-payment-attempt-for-refund-found"
	NO_PAYMENT_METHOD_AVAILABLE ErrorType = "no-payment-method-available"
	NO_TEAM_ACCOUNT_SEATS_AVAILABLE ErrorType = "no-team-account-seats-available"
	NON_TEMPORARY_EMAIL ErrorType = "non-temporary-email"
	NOT_ENOUGH_USAGE_TO_RUN_PAID_ACTOR ErrorType = "not-enough-usage-to-run-paid-actor"
	NOT_IMPLEMENTED ErrorType = "not-implemented"
	NOT_SUPPORTED_CURRENCIES ErrorType = "not-supported-currencies"
	O_AUTH_SERVICE_ALREADY_CONNECTED ErrorType = "o-auth-service-already-connected"
	O_AUTH_SERVICE_NOT_CONNECTED ErrorType = "o-auth-service-not-connected"
	OAUTH_RESOURCE_ACCESS_FAILED ErrorType = "oauth-resource-access-failed"
	ONE_TIME_INVOICE_ALREADY_MARKED_PAID ErrorType = "one-time-invoice-already-marked-paid"
	ONLY_DRAFTS_CAN_BE_DELETED ErrorType = "only-drafts-can-be-deleted"
	OPERATION_CANCELED ErrorType = "operation-canceled"
	OPERATION_NOT_ALLOWED ErrorType = "operation-not-allowed"
	OPERATION_TIMED_OUT ErrorType = "operation-timed-out"
	ORGANIZATION_CANNOT_OWN_ITSELF ErrorType = "organization-cannot-own-itself"
	ORGANIZATION_ROLE_NOT_FOUND ErrorType = "organization-role-not-found"
	OVERLAPPING_PAYOUT_BILLING_PERIODS ErrorType = "overlapping-payout-billing-periods"
	OWN_TOKEN_REQUIRED ErrorType = "own-token-required"
	PAGE_NOT_FOUND ErrorType = "page-not-found"
	PARAM_NOT_ONE_OF ErrorType = "param-not-one-of"
	PARAMETER_REQUIRED ErrorType = "parameter-required"
	PARAMETERS_MISMATCHED ErrorType = "parameters-mismatched"
	PASSWORD_RESET_EMAIL_ALREADY_SENT ErrorType = "password-reset-email-already-sent"
	PASSWORD_RESET_TOKEN_EXPIRED ErrorType = "password-reset-token-expired"
	PAY_AS_YOU_GO_WITHOUT_MONTHLY_INTERVAL ErrorType = "pay-as-you-go-without-monthly-interval"
	PAYMENT_ATTEMPT_STATUS_MESSAGE_REQUIRED ErrorType = "payment-attempt-status-message-required"
	PAYOUT_ALREADY_PAID ErrorType = "payout-already-paid"
	PAYOUT_CANCELED ErrorType = "payout-canceled"
	PAYOUT_INVALID_STATE ErrorType = "payout-invalid-state"
	PAYOUT_MUST_BE_APPROVED_TO_BE_MARKED_PAID ErrorType = "payout-must-be-approved-to-be-marked-paid"
	PAYOUT_NOT_FOUND ErrorType = "payout-not-found"
	PAYOUT_NUMBER_ALREADY_EXISTS ErrorType = "payout-number-already-exists"
	PHONE_NUMBER_INVALID ErrorType = "phone-number-invalid"
	PHONE_NUMBER_LANDLINE ErrorType = "phone-number-landline"
	PHONE_NUMBER_OPTED_OUT ErrorType = "phone-number-opted-out"
	PHONE_VERIFICATION_DISABLED ErrorType = "phone-verification-disabled"
	PLATFORM_FEATURE_DISABLED ErrorType = "platform-feature-disabled"
	PRICE_OVERRIDES_VALIDATION_FAILED ErrorType = "price-overrides-validation-failed"
	PRICING_MODEL_NOT_SUPPORTED ErrorType = "pricing-model-not-supported"
	PROMOTIONAL_PLAN_NOT_AVAILABLE ErrorType = "promotional-plan-not-available"
	PROXY_AUTH_IP_NOT_UNIQUE ErrorType = "proxy-auth-ip-not-unique"
	PUBLIC_ACTOR_DISABLED ErrorType = "public-actor-disabled"
	QUERY_TIMEOUT ErrorType = "query-timeout"
	QUOTED_PRICE_OUTDATED ErrorType = "quoted-price-outdated"
	RATE_LIMIT_EXCEEDED ErrorType = "rate-limit-exceeded"
	RECAPTCHA_INVALID ErrorType = "recaptcha-invalid"
	RECAPTCHA_REQUIRED ErrorType = "recaptcha-required"
	RECORD_NOT_FOUND ErrorType = "record-not-found"
	RECORD_NOT_PUBLIC ErrorType = "record-not-public"
	RECORD_OR_TOKEN_NOT_FOUND ErrorType = "record-or-token-not-found"
	RECORD_TOO_LARGE ErrorType = "record-too-large"
	REDIRECT_URI_MISMATCH ErrorType = "redirect-uri-mismatch"
	REDUCED_PLAN_NOT_AVAILABLE ErrorType = "reduced-plan-not-available"
	RENTAL_CHARGE_ALREADY_REIMBURSED ErrorType = "rental-charge-already-reimbursed"
	RENTAL_NOT_ALLOWED ErrorType = "rental-not-allowed"
	REQUEST_ABORTED_PREMATURELY ErrorType = "request-aborted-prematurely"
	REQUEST_HANDLED_OR_LOCKED ErrorType = "request-handled-or-locked"
	REQUEST_ID_INVALID ErrorType = "request-id-invalid"
	REQUEST_QUEUE_DUPLICATE_REQUESTS ErrorType = "request-queue-duplicate-requests"
	REQUEST_TOO_LARGE ErrorType = "request-too-large"
	REQUESTED_DATASET_VIEW_DOES_NOT_EXIST ErrorType = "requested-dataset-view-does-not-exist"
	RESUME_TOKEN_EXPIRED ErrorType = "resume-token-expired"
	RUN_FAILED ErrorType = "run-failed"
	RUN_INPUT_BODY_NOT_VALID_JSON ErrorType = "run-input-body-not-valid-json"
	RUN_TIMEOUT_EXCEEDED ErrorType = "run-timeout-exceeded"
	RUSSIA_IS_EVIL ErrorType = "russia-is-evil"
	SAME_USER ErrorType = "same-user"
	SCHEDULE_ACTOR_NOT_FOUND ErrorType = "schedule-actor-not-found"
	SCHEDULE_ACTOR_TASK_NOT_FOUND ErrorType = "schedule-actor-task-not-found"
	SCHEDULE_NAME_NOT_UNIQUE ErrorType = "schedule-name-not-unique"
	SCHEMA_VALIDATION ErrorType = "schema-validation"
	SCHEMA_VALIDATION_ERROR ErrorType = "schema-validation-error"
	SCHEMA_VALIDATION_FAILED ErrorType = "schema-validation-failed"
	SIGN_UP_METHOD_NOT_ALLOWED ErrorType = "sign-up-method-not-allowed"
	SLACK_INTEGRATION_NOT_CUSTOM ErrorType = "slack-integration-not-custom"
	SOCKET_CLOSED ErrorType = "socket-closed"
	SOCKET_DESTROYED ErrorType = "socket-destroyed"
	STORE_SCHEMA_INVALID ErrorType = "store-schema-invalid"
	STORE_TERMS_NOT_ACCEPTED ErrorType = "store-terms-not-accepted"
	STRIPE_ENABLED ErrorType = "stripe-enabled"
	STRIPE_GENERIC_DECLINE ErrorType = "stripe-generic-decline"
	STRIPE_NOT_ENABLED ErrorType = "stripe-not-enabled"
	STRIPE_NOT_ENABLED_FOR_USER ErrorType = "stripe-not-enabled-for-user"
	TAGGED_BUILD_REQUIRED ErrorType = "tagged-build-required"
	TAX_COUNTRY_INVALID ErrorType = "tax-country-invalid"
	TAX_NUMBER_INVALID ErrorType = "tax-number-invalid"
	TAX_NUMBER_VALIDATION_FAILED ErrorType = "tax-number-validation-failed"
	TAXAMO_CALL_FAILED ErrorType = "taxamo-call-failed"
	TAXAMO_REQUEST_FAILED ErrorType = "taxamo-request-failed"
	TESTING_ERROR ErrorType = "testing-error"
	TOKEN_NOT_PROVIDED ErrorType = "token-not-provided"
	TOO_FEW_VERSIONS ErrorType = "too-few-versions"
	TOO_MANY_ACTOR_TASKS ErrorType = "too-many-actor-tasks"
	TOO_MANY_ACTORS ErrorType = "too-many-actors"
	TOO_MANY_LABELS_ON_RESOURCE ErrorType = "too-many-labels-on-resource"
	TOO_MANY_MCP_CONNECTORS ErrorType = "too-many-mcp-connectors"
	TOO_MANY_O_AUTH_APPS ErrorType = "too-many-o-auth-apps"
	TOO_MANY_ORGANIZATIONS ErrorType = "too-many-organizations"
	TOO_MANY_REQUESTS ErrorType = "too-many-requests"
	TOO_MANY_SCHEDULES ErrorType = "too-many-schedules"
	TOO_MANY_UI_ACCESS_KEYS ErrorType = "too-many-ui-access-keys"
	TOO_MANY_USER_LABELS ErrorType = "too-many-user-labels"
	TOO_MANY_VALUES ErrorType = "too-many-values"
	TOO_MANY_VERSIONS ErrorType = "too-many-versions"
	TOO_MANY_WEBHOOKS ErrorType = "too-many-webhooks"
	UNEXPECTED_ROUTE ErrorType = "unexpected-route"
	UNKNOWN_BUILD_TAG ErrorType = "unknown-build-tag"
	UNKNOWN_PAYMENT_PROVIDER ErrorType = "unknown-payment-provider"
	UNSUBSCRIBE_TOKEN_INVALID ErrorType = "unsubscribe-token-invalid"
	UNSUPPORTED_ACTOR_PRICING_MODEL_FOR_AGENTIC_PAYMENTS ErrorType = "unsupported-actor-pricing-model-for-agentic-payments"
	UNSUPPORTED_CONTENT_ENCODING ErrorType = "unsupported-content-encoding"
	UNSUPPORTED_FILE_TYPE_FOR_ISSUE ErrorType = "unsupported-file-type-for-issue"
	UNSUPPORTED_FILE_TYPE_IMAGE_EXPECTED ErrorType = "unsupported-file-type-image-expected"
	UNSUPPORTED_FILE_TYPE_TEXT_OR_JSON_EXPECTED ErrorType = "unsupported-file-type-text-or-json-expected"
	UNSUPPORTED_PERMISSION ErrorType = "unsupported-permission"
	UPCOMING_SUBSCRIPTION_BILL_NOT_UP_TO_DATE ErrorType = "upcoming-subscription-bill-not-up-to-date"
	USER_ALREADY_EXISTS ErrorType = "user-already-exists"
	USER_ALREADY_VERIFIED ErrorType = "user-already-verified"
	USER_CREATES_ORGANIZATIONS_TOO_FAST ErrorType = "user-creates-organizations-too-fast"
	USER_DISABLED ErrorType = "user-disabled"
	USER_EMAIL_IS_DISPOSABLE ErrorType = "user-email-is-disposable"
	USER_EMAIL_NOT_SET ErrorType = "user-email-not-set"
	USER_EMAIL_NOT_VERIFIED ErrorType = "user-email-not-verified"
	USER_HAS_NO_SUBSCRIPTION ErrorType = "user-has-no-subscription"
	USER_INTEGRATION_NOT_FOUND ErrorType = "user-integration-not-found"
	USER_IS_ALREADY_INVITED ErrorType = "user-is-already-invited"
	USER_IS_ALREADY_ORGANIZATION_MEMBER ErrorType = "user-is-already-organization-member"
	USER_IS_NOT_MEMBER_OF_ORGANIZATION ErrorType = "user-is-not-member-of-organization"
	USER_IS_NOT_ORGANIZATION ErrorType = "user-is-not-organization"
	USER_IS_ORGANIZATION ErrorType = "user-is-organization"
	USER_IS_ORGANIZATION_OWNER ErrorType = "user-is-organization-owner"
	USER_IS_REMOVED ErrorType = "user-is-removed"
	USER_NOT_FOUND ErrorType = "user-not-found"
	USER_NOT_LOGGED_IN ErrorType = "user-not-logged-in"
	USER_NOT_VERIFIED ErrorType = "user-not-verified"
	USER_OR_TOKEN_NOT_FOUND ErrorType = "user-or-token-not-found"
	USER_PLAN_NOT_ALLOWED_FOR_COUPON ErrorType = "user-plan-not-allowed-for-coupon"
	USER_PROBLEM_WITH_CARD ErrorType = "user-problem-with-card"
	USER_RECORD_NOT_FOUND ErrorType = "user-record-not-found"
	USERNAME_ALREADY_TAKEN ErrorType = "username-already-taken"
	USERNAME_MISSING ErrorType = "username-missing"
	USERNAME_NOT_ALLOWED ErrorType = "username-not-allowed"
	USERNAME_REMOVAL_FORBIDDEN ErrorType = "username-removal-forbidden"
	USERNAME_REQUIRED ErrorType = "username-required"
	VERIFICATION_EMAIL_ALREADY_SENT ErrorType = "verification-email-already-sent"
	VERIFICATION_TOKEN_EXPIRED ErrorType = "verification-token-expired"
	VERSION_ALREADY_EXISTS ErrorType = "version-already-exists"
	VERSIONS_SIZE_EXCEEDED ErrorType = "versions-size-exceeded"
	WEAK_PASSWORD ErrorType = "weak-password"
	X402_AGENTIC_PAYMENT_ALREADY_FINALIZED ErrorType = "x402-agentic-payment-already-finalized"
	X402_AGENTIC_PAYMENT_INSUFFICIENT_AMOUNT ErrorType = "x402-agentic-payment-insufficient-amount"
	X402_AGENTIC_PAYMENT_MALFORMED_TOKEN ErrorType = "x402-agentic-payment-malformed-token"
	X402_AGENTIC_PAYMENT_SETTLEMENT_FAILED ErrorType = "x402-agentic-payment-settlement-failed"
	X402_AGENTIC_PAYMENT_SETTLEMENT_IN_PROGRESS ErrorType = "x402-agentic-payment-settlement-in-progress"
	X402_AGENTIC_PAYMENT_SETTLEMENT_STUCK ErrorType = "x402-agentic-payment-settlement-stuck"
	X402_AGENTIC_PAYMENT_UNAUTHORIZED ErrorType = "x402-agentic-payment-unauthorized"
	X402_PAYMENT_REQUIRED ErrorType = "x402-payment-required"
	ZERO_INVOICE ErrorType = "zero-invoice"
)

// All allowed values of ErrorType enum
var AllowedErrorTypeEnumValues = []ErrorType{
	"3d-secure-auth-failed",
	"access-right-already-exists",
	"action-not-found",
	"actor-already-rented",
	"actor-can-not-be-rented",
	"actor-disabled",
	"actor-is-not-rented",
	"actor-memory-limit-exceeded",
	"actor-name-exists-new-owner",
	"actor-name-not-unique",
	"actor-not-found",
	"actor-not-github-actor",
	"actor-not-public",
	"actor-permission-level-not-supported-for-agentic-payments",
	"actor-review-already-exists",
	"actor-run-failed",
	"actor-standby-not-supported-for-agentic-payments",
	"actor-task-name-not-unique",
	"agentic-payment-info-retrieval-error",
	"agentic-payment-information-missing",
	"agentic-payment-insufficient-amount",
	"agentic-payment-provider-internal-error",
	"agentic-payment-provider-unauthorized",
	"airtable-webhook-deprecated",
	"already-subscribed-to-paid-actor",
	"apify-plan-required-to-use-paid-actor",
	"apify-signup-not-allowed",
	"auth-method-not-supported",
	"authorization-server-not-found",
	"auto-issue-date-invalid",
	"background-check-required",
	"billing-system-error",
	"black-friday-plan-expired",
	"braintree-error",
	"braintree-not-linked",
	"braintree-operation-timed-out",
	"braintree-unsupported-currency",
	"build-not-found",
	"build-outdated",
	"cannot-add-apify-events-to-ppe-actor",
	"cannot-add-multiple-pricing-infos",
	"cannot-add-pricing-info-that-alters-past",
	"cannot-add-second-future-pricing-info",
	"cannot-build-actor-from-webhook",
	"cannot-change-billing-interval",
	"cannot-change-owner",
	"cannot-charge-apify-event",
	"cannot-charge-non-pay-per-event-actor",
	"cannot-comment-as-other-user",
	"cannot-copy-actor-task",
	"cannot-create-payout",
	"cannot-create-public-actor",
	"cannot-create-tax-transaction",
	"cannot-delete-critical-actor",
	"cannot-delete-invoice",
	"cannot-delete-paid-actor",
	"cannot-disable-one-time-event-for-apify-start-event",
	"cannot-disable-organization-with-enabled-members",
	"cannot-disable-user-with-subscription",
	"cannot-link-oauth-to-unverified-email",
	"cannot-metamorph-to-pay-per-result-actor",
	"cannot-modify-actor-pricing-too-frequently",
	"cannot-modify-actor-pricing-with-immediate-effect",
	"cannot-override-paid-actor-trial",
	"cannot-permanently-delete-subscription",
	"cannot-publish-actor",
	"cannot-reduce-last-full-token",
	"cannot-reimburse-more-than-original-charge",
	"cannot-reimburse-non-rental-charge",
	"cannot-remove-own-actor-from-recently-used",
	"cannot-remove-payment-method",
	"cannot-remove-pricing-info",
	"cannot-remove-running-run",
	"cannot-remove-user-with-public-actors",
	"cannot-remove-user-with-subscription",
	"cannot-remove-user-with-unpaid-invoice",
	"cannot-rename-env-var",
	"cannot-rent-paid-actor",
	"cannot-review-own-actor",
	"cannot-set-access-rights-for-owner",
	"cannot-set-is-status-message-terminal",
	"cannot-unpublish-critical-actor",
	"cannot-unpublish-paid-actor",
	"cannot-unpublish-profile",
	"cannot-update-invoice-field",
	"concurrent-runs-limit-exceeded",
	"concurrent-update-detected",
	"conference-token-not-found",
	"content-encoding-forbidden-for-html",
	"coupon-already-redeemed",
	"coupon-expired",
	"coupon-for-new-customers",
	"coupon-for-subscribed-users",
	"coupon-limits-are-in-conflict-with-current-limits",
	"coupon-max-number-of-redemptions-reached",
	"coupon-not-found",
	"coupon-not-unique",
	"coupons-disabled",
	"create-github-issue-not-allowed",
	"creator-plan-not-available",
	"cron-expression-invalid",
	"daily-ai-token-limit-exceeded",
	"daily-publication-limit-exceeded",
	"dataset-does-not-have-fields-schema",
	"dataset-does-not-have-schema",
	"dataset-locked",
	"dataset-schema-invalid",
	"dcr-not-supported",
	"default-dataset-not-found",
	"deleting-default-build",
	"deleting-unfinished-build",
	"email-already-taken",
	"email-already-taken-removed-user",
	"email-domain-not-allowed-for-coupon",
	"email-invalid",
	"email-not-allowed",
	"email-not-valid",
	"email-update-too-soon",
	"elevated-permissions-needed",
	"env-var-already-exists",
	"exchange-rate-fetch-failed",
	"expired-conference-token",
	"failed-to-charge-user",
	"final-invoice-negative",
	"full-permission-actor-not-approved",
	"github-branch-empty",
	"github-issue-already-exists",
	"github-public-key-not-found",
	"github-repository-not-found",
	"github-signature-does-not-match-payload",
	"github-user-not-authorized-for-issues",
	"gmail-not-allowed",
	"id-does-not-match",
	"incompatible-billing-interval",
	"incomplete-payout-billing-info",
	"inconsistent-currencies",
	"incorrect-pricing-modifier-prefix",
	"input-json-invalid-characters",
	"input-json-not-object",
	"input-json-too-long",
	"input-update-collision",
	"insufficient-permissions",
	"insufficient-permissions-to-change-field",
	"insufficient-security-measures",
	"insufficient-tax-country-evidence",
	"integration-auth-error",
	"internal-server-error",
	"invalid-billing-info",
	"invalid-billing-period-for-payout",
	"invalid-build",
	"invalid-client-key",
	"invalid-collection",
	"invalid-conference-login-password",
	"invalid-content-type-header",
	"invalid-credentials",
	"invalid-git-auth-token",
	"invalid-github-issue-url",
	"invalid-header",
	"invalid-id",
	"invalid-idempotency-key",
	"invalid-input",
	"invalid-input-schema",
	"invalid-invoice",
	"invalid-invoice-type",
	"invalid-issue-date",
	"invalid-label-params",
	"invalid-main-account-user-id",
	"invalid-oauth-app",
	"invalid-oauth-scope",
	"invalid-one-time-invoice",
	"invalid-parameter",
	"invalid-payout-status",
	"invalid-picture-url",
	"invalid-record-key",
	"invalid-request",
	"invalid-resource-type",
	"invalid-signature",
	"invalid-subscription-plan",
	"invalid-tax-number",
	"invalid-tax-number-format",
	"invalid-token",
	"invalid-token-type",
	"invalid-two-factor-code",
	"invalid-two-factor-code-or-recovery-code",
	"invalid-two-factor-recovery-code",
	"invalid-username",
	"invalid-value",
	"invitation-invalid-resource-type",
	"invitation-no-longer-valid",
	"invoice-canceled",
	"invoice-cannot-be-refunded-due-to-too-high-amount",
	"invoice-incomplete",
	"invoice-is-draft",
	"invoice-locked",
	"invoice-must-be-buffer",
	"invoice-not-canceled",
	"invoice-not-draft",
	"invoice-not-found",
	"invoice-outdated",
	"invoice-paid-already",
	"issue-already-connected-to-github",
	"issue-not-found",
	"issues-bad-request",
	"issuer-not-registered",
	"job-finished",
	"label-already-linked",
	"last-api-token",
	"limit-reached",
	"max-items-must-be-greater-than-zero",
	"max-metamorphs-exceeded",
	"max-total-charge-usd-below-minimum",
	"max-total-charge-usd-must-be-greater-than-zero",
	"method-not-allowed",
	"migration-disabled",
	"missing-actor-rights",
	"missing-api-token",
	"missing-billing-info",
	"missing-line-items",
	"missing-payment-date",
	"missing-payout-billing-info",
	"missing-proxy-password",
	"missing-reporting-fields",
	"missing-resource-name",
	"missing-settings",
	"missing-username",
	"monthly-usage-limit-too-low",
	"more-than-one-update-not-allowed",
	"multiple-records-found",
	"must-be-admin",
	"name-not-unique",
	"next-runtime-computation-failed",
	"no-columns-in-exported-dataset",
	"no-payment-attempt-for-refund-found",
	"no-payment-method-available",
	"no-team-account-seats-available",
	"non-temporary-email",
	"not-enough-usage-to-run-paid-actor",
	"not-implemented",
	"not-supported-currencies",
	"o-auth-service-already-connected",
	"o-auth-service-not-connected",
	"oauth-resource-access-failed",
	"one-time-invoice-already-marked-paid",
	"only-drafts-can-be-deleted",
	"operation-canceled",
	"operation-not-allowed",
	"operation-timed-out",
	"organization-cannot-own-itself",
	"organization-role-not-found",
	"overlapping-payout-billing-periods",
	"own-token-required",
	"page-not-found",
	"param-not-one-of",
	"parameter-required",
	"parameters-mismatched",
	"password-reset-email-already-sent",
	"password-reset-token-expired",
	"pay-as-you-go-without-monthly-interval",
	"payment-attempt-status-message-required",
	"payout-already-paid",
	"payout-canceled",
	"payout-invalid-state",
	"payout-must-be-approved-to-be-marked-paid",
	"payout-not-found",
	"payout-number-already-exists",
	"phone-number-invalid",
	"phone-number-landline",
	"phone-number-opted-out",
	"phone-verification-disabled",
	"platform-feature-disabled",
	"price-overrides-validation-failed",
	"pricing-model-not-supported",
	"promotional-plan-not-available",
	"proxy-auth-ip-not-unique",
	"public-actor-disabled",
	"query-timeout",
	"quoted-price-outdated",
	"rate-limit-exceeded",
	"recaptcha-invalid",
	"recaptcha-required",
	"record-not-found",
	"record-not-public",
	"record-or-token-not-found",
	"record-too-large",
	"redirect-uri-mismatch",
	"reduced-plan-not-available",
	"rental-charge-already-reimbursed",
	"rental-not-allowed",
	"request-aborted-prematurely",
	"request-handled-or-locked",
	"request-id-invalid",
	"request-queue-duplicate-requests",
	"request-too-large",
	"requested-dataset-view-does-not-exist",
	"resume-token-expired",
	"run-failed",
	"run-input-body-not-valid-json",
	"run-timeout-exceeded",
	"russia-is-evil",
	"same-user",
	"schedule-actor-not-found",
	"schedule-actor-task-not-found",
	"schedule-name-not-unique",
	"schema-validation",
	"schema-validation-error",
	"schema-validation-failed",
	"sign-up-method-not-allowed",
	"slack-integration-not-custom",
	"socket-closed",
	"socket-destroyed",
	"store-schema-invalid",
	"store-terms-not-accepted",
	"stripe-enabled",
	"stripe-generic-decline",
	"stripe-not-enabled",
	"stripe-not-enabled-for-user",
	"tagged-build-required",
	"tax-country-invalid",
	"tax-number-invalid",
	"tax-number-validation-failed",
	"taxamo-call-failed",
	"taxamo-request-failed",
	"testing-error",
	"token-not-provided",
	"too-few-versions",
	"too-many-actor-tasks",
	"too-many-actors",
	"too-many-labels-on-resource",
	"too-many-mcp-connectors",
	"too-many-o-auth-apps",
	"too-many-organizations",
	"too-many-requests",
	"too-many-schedules",
	"too-many-ui-access-keys",
	"too-many-user-labels",
	"too-many-values",
	"too-many-versions",
	"too-many-webhooks",
	"unexpected-route",
	"unknown-build-tag",
	"unknown-payment-provider",
	"unsubscribe-token-invalid",
	"unsupported-actor-pricing-model-for-agentic-payments",
	"unsupported-content-encoding",
	"unsupported-file-type-for-issue",
	"unsupported-file-type-image-expected",
	"unsupported-file-type-text-or-json-expected",
	"unsupported-permission",
	"upcoming-subscription-bill-not-up-to-date",
	"user-already-exists",
	"user-already-verified",
	"user-creates-organizations-too-fast",
	"user-disabled",
	"user-email-is-disposable",
	"user-email-not-set",
	"user-email-not-verified",
	"user-has-no-subscription",
	"user-integration-not-found",
	"user-is-already-invited",
	"user-is-already-organization-member",
	"user-is-not-member-of-organization",
	"user-is-not-organization",
	"user-is-organization",
	"user-is-organization-owner",
	"user-is-removed",
	"user-not-found",
	"user-not-logged-in",
	"user-not-verified",
	"user-or-token-not-found",
	"user-plan-not-allowed-for-coupon",
	"user-problem-with-card",
	"user-record-not-found",
	"username-already-taken",
	"username-missing",
	"username-not-allowed",
	"username-removal-forbidden",
	"username-required",
	"verification-email-already-sent",
	"verification-token-expired",
	"version-already-exists",
	"versions-size-exceeded",
	"weak-password",
	"x402-agentic-payment-already-finalized",
	"x402-agentic-payment-insufficient-amount",
	"x402-agentic-payment-malformed-token",
	"x402-agentic-payment-settlement-failed",
	"x402-agentic-payment-settlement-in-progress",
	"x402-agentic-payment-settlement-stuck",
	"x402-agentic-payment-unauthorized",
	"x402-payment-required",
	"zero-invoice",
}

func (v *ErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorType(value)
	for _, existing := range AllowedErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorType", value)
}

// NewErrorTypeFromValue returns a pointer to a valid ErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorTypeFromValue(v string) (*ErrorType, error) {
	ev := ErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorType: valid values are %v", v, AllowedErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorType) IsValid() bool {
	for _, existing := range AllowedErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorType value
func (v ErrorType) Ptr() *ErrorType {
	return &v
}

type NullableErrorType struct {
	value *ErrorType
	isSet bool
}

func (v NullableErrorType) Get() *ErrorType {
	return v.value
}

func (v *NullableErrorType) Set(val *ErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorType(val *ErrorType) *NullableErrorType {
	return &NullableErrorType{value: val, isSet: true}
}

func (v NullableErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

