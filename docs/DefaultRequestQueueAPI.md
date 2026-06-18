# \DefaultRequestQueueAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorRunRequestQueueDelete**](DefaultRequestQueueAPI.md#ActorRunRequestQueueDelete) | **Delete** /v2/actor-runs/{runId}/request-queue | Delete default request queue
[**ActorRunRequestQueueGet**](DefaultRequestQueueAPI.md#ActorRunRequestQueueGet) | **Get** /v2/actor-runs/{runId}/request-queue | Get default request queue
[**ActorRunRequestQueueHeadGet**](DefaultRequestQueueAPI.md#ActorRunRequestQueueHeadGet) | **Get** /v2/actor-runs/{runId}/request-queue/head | Get default request queue head
[**ActorRunRequestQueueHeadLockPost**](DefaultRequestQueueAPI.md#ActorRunRequestQueueHeadLockPost) | **Post** /v2/actor-runs/{runId}/request-queue/head/lock | Get and lock default request queue head
[**ActorRunRequestQueuePut**](DefaultRequestQueueAPI.md#ActorRunRequestQueuePut) | **Put** /v2/actor-runs/{runId}/request-queue | Update default request queue
[**ActorRunRequestQueueRequestDelete**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestDelete) | **Delete** /v2/actor-runs/{runId}/request-queue/requests/{requestId} | Delete request from default request queue
[**ActorRunRequestQueueRequestGet**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestGet) | **Get** /v2/actor-runs/{runId}/request-queue/requests/{requestId} | Get request from default request queue
[**ActorRunRequestQueueRequestLockDelete**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestLockDelete) | **Delete** /v2/actor-runs/{runId}/request-queue/requests/{requestId}/lock | Delete lock on request in default request queue
[**ActorRunRequestQueueRequestLockPut**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestLockPut) | **Put** /v2/actor-runs/{runId}/request-queue/requests/{requestId}/lock | Prolong lock on request in default request queue
[**ActorRunRequestQueueRequestPut**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestPut) | **Put** /v2/actor-runs/{runId}/request-queue/requests/{requestId} | Update request in default request queue
[**ActorRunRequestQueueRequestsBatchDelete**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestsBatchDelete) | **Delete** /v2/actor-runs/{runId}/request-queue/requests/batch | Batch delete requests from default request queue
[**ActorRunRequestQueueRequestsBatchPost**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestsBatchPost) | **Post** /v2/actor-runs/{runId}/request-queue/requests/batch | Batch add requests to default request queue
[**ActorRunRequestQueueRequestsGet**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestsGet) | **Get** /v2/actor-runs/{runId}/request-queue/requests | List default request queue&#39;s requests
[**ActorRunRequestQueueRequestsPost**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestsPost) | **Post** /v2/actor-runs/{runId}/request-queue/requests | Add request to default request queue
[**ActorRunRequestQueueRequestsUnlockPost**](DefaultRequestQueueAPI.md#ActorRunRequestQueueRequestsUnlockPost) | **Post** /v2/actor-runs/{runId}/request-queue/requests/unlock | Unlock requests in default request queue



## ActorRunRequestQueueDelete

> ActorRunRequestQueueDelete(ctx, runId).Execute()

Delete default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueDelete(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueGet

> RequestQueueResponse ActorRunRequestQueueGet(ctx, runId).Execute()

Get default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueGet`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestQueueResponse**](RequestQueueResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueHeadGet

> HeadResponse ActorRunRequestQueueHeadGet(ctx, runId).Limit(limit).ClientKey(clientKey).Execute()

Get default request queue head



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	limit := float64(100) // float64 | How many items from queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueHeadGet(context.Background(), runId).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueHeadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueHeadGet`: HeadResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueHeadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueHeadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float64** | How many items from queue should be returned. | 
 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 

### Return type

[**HeadResponse**](HeadResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueHeadLockPost

> HeadAndLockResponse ActorRunRequestQueueHeadLockPost(ctx, runId).LockSecs(lockSecs).Limit(limit).ClientKey(clientKey).Execute()

Get and lock default request queue head



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	limit := float64(25) // float64 | How many items from the queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueHeadLockPost(context.Background(), runId).LockSecs(lockSecs).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueHeadLockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueHeadLockPost`: HeadAndLockResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueHeadLockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueHeadLockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockSecs** | **float64** | How long the requests will be locked for (in seconds). | 
 **limit** | **float64** | How many items from the queue should be returned. | 
 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 

### Return type

[**HeadAndLockResponse**](HeadAndLockResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueuePut

> RequestQueueResponse ActorRunRequestQueuePut(ctx, runId).UpdateRequestQueueRequest(updateRequestQueueRequest).Execute()

Update default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	updateRequestQueueRequest := *openapiclient.NewUpdateRequestQueueRequest() // UpdateRequestQueueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueuePut(context.Background(), runId).UpdateRequestQueueRequest(updateRequestQueueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueuePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueuePut`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueuePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueuePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequestQueueRequest** | [**UpdateRequestQueueRequest**](UpdateRequestQueueRequest.md) |  | 

### Return type

[**RequestQueueResponse**](RequestQueueResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestDelete

> ActorRunRequestQueueRequestDelete(ctx, runId, requestId).ClientKey(clientKey).Execute()

Delete request from default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestDelete(context.Background(), runId, requestId).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestGet

> RequestResponse ActorRunRequestQueueRequestGet(ctx, runId, requestId).Execute()

Get request from default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestGet(context.Background(), runId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestGet`: RequestResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequestResponse**](RequestResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestLockDelete

> ActorRunRequestQueueRequestLockDelete(ctx, runId, requestId).ClientKey(clientKey).Forefront(forefront).Execute()

Delete lock on request in default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock was removed.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestLockDelete(context.Background(), runId, requestId).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestLockDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestLockDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 
 **forefront** | **string** | Determines if request should be added to the head of the queue or to the end after lock was removed.  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestLockPut

> ProlongRequestLockResponse ActorRunRequestQueueRequestLockPut(ctx, runId, requestId).LockSecs(lockSecs).ClientKey(clientKey).Forefront(forefront).Execute()

Prolong lock on request in default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock expires.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestLockPut(context.Background(), runId, requestId).LockSecs(lockSecs).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestLockPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestLockPut`: ProlongRequestLockResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestLockPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestLockPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lockSecs** | **float64** | How long the requests will be locked for (in seconds). | 
 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 
 **forefront** | **string** | Determines if request should be added to the head of the queue or to the end after lock expires.  | 

### Return type

[**ProlongRequestLockResponse**](ProlongRequestLockResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestPut

> UpdateRequestResponse ActorRunRequestQueueRequestPut(ctx, runId, requestId).Request(request).Forefront(forefront).ClientKey(clientKey).Execute()

Update request in default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	request := *openapiclient.NewRequest() // Request | 
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestPut(context.Background(), runId, requestId).Request(request).Forefront(forefront).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestPut`: UpdateRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**Request**](Request.md) |  | 
 **forefront** | **string** | Determines if request should be added to the head of the queue or to the end. Default value is &#x60;false&#x60; (end of queue).  | 
 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 

### Return type

[**UpdateRequestResponse**](UpdateRequestResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestsBatchDelete

> BatchDeleteResponse ActorRunRequestQueueRequestsBatchDelete(ctx, runId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).ClientKey(clientKey).Execute()

Batch delete requests from default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	contentType := "contentType_example" // string | 
	requestDraftDelete := []openapiclient.RequestDraftDelete{*openapiclient.NewRequestDraftDelete("Id_example", "UniqueKey_example")} // []RequestDraftDelete | 
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestsBatchDelete(context.Background(), runId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsBatchDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestsBatchDelete`: BatchDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsBatchDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestsBatchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **requestDraftDelete** | [**[]RequestDraftDelete**](RequestDraftDelete.md) |  | 
 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 

### Return type

[**BatchDeleteResponse**](BatchDeleteResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestsBatchPost

> BatchAddResponse ActorRunRequestQueueRequestsBatchPost(ctx, runId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()

Batch add requests to default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	requestBase := []openapiclient.RequestBase{*openapiclient.NewRequestBase()} // []RequestBase | 
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestsBatchPost(context.Background(), runId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestsBatchPost`: BatchAddResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsBatchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestsBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBase** | [**[]RequestBase**](RequestBase.md) |  | 
 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 
 **forefront** | **string** | Determines if request should be added to the head of the queue or to the end. Default value is &#x60;false&#x60; (end of queue).  | 

### Return type

[**BatchAddResponse**](BatchAddResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestsGet

> ListOfRequestsResponse ActorRunRequestQueueRequestsGet(ctx, runId).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()

List default request queue's requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	exclusiveStartId := "Ihnsp8YrvJ8102Kj" // string | All requests up to this one (including) are skipped from the result. (Deprecated, use `cursor` instead.) (optional)
	limit := float64(100) // float64 | Number of keys to be returned. Maximum value is `10000`. (optional)
	cursor := "eyJyZXF1ZXN0SWQiOiI2OFRqQ2RaTDNvM2hiUU0ifQ" // string | A cursor string for pagination, returned in the previous response as `nextCursor`. Use this to retrieve the next page of requests. (optional)
	filter := []string{"Filter_example"} // []string | Filter requests by their state. Possible values are `locked` and `pending`. You can combine multiple values separated by commas, which will mean the union of these filters – requests matching any of the specified states will be returned. (Not compatible with deprecated `exclusiveStartId` parameter.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestsGet(context.Background(), runId).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestsGet`: ListOfRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 
 **exclusiveStartId** | **string** | All requests up to this one (including) are skipped from the result. (Deprecated, use &#x60;cursor&#x60; instead.) | 
 **limit** | **float64** | Number of keys to be returned. Maximum value is &#x60;10000&#x60;. | 
 **cursor** | **string** | A cursor string for pagination, returned in the previous response as &#x60;nextCursor&#x60;. Use this to retrieve the next page of requests. | 
 **filter** | **[]string** | Filter requests by their state. Possible values are &#x60;locked&#x60; and &#x60;pending&#x60;. You can combine multiple values separated by commas, which will mean the union of these filters – requests matching any of the specified states will be returned. (Not compatible with deprecated &#x60;exclusiveStartId&#x60; parameter.) | 

### Return type

[**ListOfRequestsResponse**](ListOfRequestsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestsPost

> AddRequestResponse ActorRunRequestQueueRequestsPost(ctx, runId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()

Add request to default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	requestBase := *openapiclient.NewRequestBase() // RequestBase | 
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestsPost(context.Background(), runId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestsPost`: AddRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBase** | [**RequestBase**](RequestBase.md) |  | 
 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 
 **forefront** | **string** | Determines if request should be added to the head of the queue or to the end. Default value is &#x60;false&#x60; (end of queue).  | 

### Return type

[**AddRequestResponse**](AddRequestResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRequestQueueRequestsUnlockPost

> UnlockRequestsResponse ActorRunRequestQueueRequestsUnlockPost(ctx, runId).ClientKey(clientKey).Execute()

Unlock requests in default request queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultRequestQueueAPI.ActorRunRequestQueueRequestsUnlockPost(context.Background(), runId).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsUnlockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRequestQueueRequestsUnlockPost`: UnlockRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultRequestQueueAPI.ActorRunRequestQueueRequestsUnlockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRequestQueueRequestsUnlockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientKey** | **string** | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If &#x60;clientKey&#x60; is not provided, the system considers this API call to come from a new client. For details, see the &#x60;hadMultipleClients&#x60; field returned by the [Get head](#/reference/request-queues/queue-head) operation.  | 

### Return type

[**UnlockRequestsResponse**](UnlockRequestsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

