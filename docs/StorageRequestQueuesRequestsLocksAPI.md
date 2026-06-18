# \StorageRequestQueuesRequestsLocksAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestQueueHeadGet**](StorageRequestQueuesRequestsLocksAPI.md#RequestQueueHeadGet) | **Get** /v2/request-queues/{queueId}/head | Get head
[**RequestQueueHeadLockPost**](StorageRequestQueuesRequestsLocksAPI.md#RequestQueueHeadLockPost) | **Post** /v2/request-queues/{queueId}/head/lock | Get head and lock
[**RequestQueueRequestLockDelete**](StorageRequestQueuesRequestsLocksAPI.md#RequestQueueRequestLockDelete) | **Delete** /v2/request-queues/{queueId}/requests/{requestId}/lock | Delete request lock
[**RequestQueueRequestLockPut**](StorageRequestQueuesRequestsLocksAPI.md#RequestQueueRequestLockPut) | **Put** /v2/request-queues/{queueId}/requests/{requestId}/lock | Prolong request lock
[**RequestQueueRequestsUnlockPost**](StorageRequestQueuesRequestsLocksAPI.md#RequestQueueRequestsUnlockPost) | **Post** /v2/request-queues/{queueId}/requests/unlock | Unlock requests



## RequestQueueHeadGet

> HeadResponse RequestQueueHeadGet(ctx, queueId).Limit(limit).ClientKey(clientKey).Execute()

Get head



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
	queueId := "WkzbQMuFYuamGv3YF" // string | Queue ID or `username~queue-name`.
	limit := float64(100) // float64 | How many items from queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsLocksAPI.RequestQueueHeadGet(context.Background(), queueId).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsLocksAPI.RequestQueueHeadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueHeadGet`: HeadResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsLocksAPI.RequestQueueHeadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueHeadGetRequest struct via the builder pattern


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


## RequestQueueHeadLockPost

> HeadAndLockResponse RequestQueueHeadLockPost(ctx, queueId).LockSecs(lockSecs).Limit(limit).ClientKey(clientKey).Execute()

Get head and lock



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
	queueId := "WkzbQMuFYuamGv3YF" // string | Queue ID or `username~queue-name`.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	limit := float64(25) // float64 | How many items from the queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsLocksAPI.RequestQueueHeadLockPost(context.Background(), queueId).LockSecs(lockSecs).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsLocksAPI.RequestQueueHeadLockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueHeadLockPost`: HeadAndLockResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsLocksAPI.RequestQueueHeadLockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueHeadLockPostRequest struct via the builder pattern


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


## RequestQueueRequestLockDelete

> RequestQueueRequestLockDelete(ctx, queueId, requestId).ClientKey(clientKey).Forefront(forefront).Execute()

Delete request lock



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
	queueId := "WkzbQMuFYuamGv3YF" // string | Queue ID or `username~queue-name`.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock was removed.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestLockDelete(context.Background(), queueId, requestId).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestLockDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestLockDeleteRequest struct via the builder pattern


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


## RequestQueueRequestLockPut

> ProlongRequestLockResponse RequestQueueRequestLockPut(ctx, queueId, requestId).LockSecs(lockSecs).ClientKey(clientKey).Forefront(forefront).Execute()

Prolong request lock



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
	queueId := "WkzbQMuFYuamGv3YF" // string | Queue ID or `username~queue-name`.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock expires.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestLockPut(context.Background(), queueId, requestId).LockSecs(lockSecs).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestLockPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestLockPut`: ProlongRequestLockResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestLockPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestLockPutRequest struct via the builder pattern


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


## RequestQueueRequestsUnlockPost

> UnlockRequestsResponse RequestQueueRequestsUnlockPost(ctx, queueId).ClientKey(clientKey).Execute()

Unlock requests



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
	queueId := "WkzbQMuFYuamGv3YF" // string | Queue ID or `username~queue-name`.
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestsUnlockPost(context.Background(), queueId).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestsUnlockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestsUnlockPost`: UnlockRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsLocksAPI.RequestQueueRequestsUnlockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestsUnlockPostRequest struct via the builder pattern


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

