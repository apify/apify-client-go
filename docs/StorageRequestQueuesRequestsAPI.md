# \StorageRequestQueuesRequestsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestQueueRequestDelete**](StorageRequestQueuesRequestsAPI.md#RequestQueueRequestDelete) | **Delete** /v2/request-queues/{queueId}/requests/{requestId} | Delete request
[**RequestQueueRequestGet**](StorageRequestQueuesRequestsAPI.md#RequestQueueRequestGet) | **Get** /v2/request-queues/{queueId}/requests/{requestId} | Get request
[**RequestQueueRequestPut**](StorageRequestQueuesRequestsAPI.md#RequestQueueRequestPut) | **Put** /v2/request-queues/{queueId}/requests/{requestId} | Update request
[**RequestQueueRequestsGet**](StorageRequestQueuesRequestsAPI.md#RequestQueueRequestsGet) | **Get** /v2/request-queues/{queueId}/requests | List requests
[**RequestQueueRequestsPost**](StorageRequestQueuesRequestsAPI.md#RequestQueueRequestsPost) | **Post** /v2/request-queues/{queueId}/requests | Add request



## RequestQueueRequestDelete

> RequestQueueRequestDelete(ctx, queueId, requestId).ClientKey(clientKey).Execute()

Delete request



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageRequestQueuesRequestsAPI.RequestQueueRequestDelete(context.Background(), queueId, requestId).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsAPI.RequestQueueRequestDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRequestQueueRequestDeleteRequest struct via the builder pattern


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


## RequestQueueRequestGet

> RequestResponse RequestQueueRequestGet(ctx, queueId, requestId).Execute()

Get request



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsAPI.RequestQueueRequestGet(context.Background(), queueId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsAPI.RequestQueueRequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestGet`: RequestResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsAPI.RequestQueueRequestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestGetRequest struct via the builder pattern


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


## RequestQueueRequestPut

> UpdateRequestResponse RequestQueueRequestPut(ctx, queueId, requestId).Request(request).Forefront(forefront).ClientKey(clientKey).Execute()

Update request



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
	request := *openapiclient.NewRequest() // Request | 
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsAPI.RequestQueueRequestPut(context.Background(), queueId, requestId).Request(request).Forefront(forefront).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsAPI.RequestQueueRequestPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestPut`: UpdateRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsAPI.RequestQueueRequestPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestPutRequest struct via the builder pattern


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


## RequestQueueRequestsGet

> ListOfRequestsResponse RequestQueueRequestsGet(ctx, queueId).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()

List requests



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
	exclusiveStartId := "Ihnsp8YrvJ8102Kj" // string | All requests up to this one (including) are skipped from the result. (Deprecated, use `cursor` instead.) (optional)
	limit := float64(100) // float64 | Number of keys to be returned. Maximum value is `10000`. (optional)
	cursor := "eyJyZXF1ZXN0SWQiOiI2OFRqQ2RaTDNvM2hiUU0ifQ" // string | A cursor string for pagination, returned in the previous response as `nextCursor`. Use this to retrieve the next page of requests. (optional)
	filter := []string{"Filter_example"} // []string | Filter requests by their state. Possible values are `locked` and `pending`. You can combine multiple values separated by commas, which will mean the union of these filters – requests matching any of the specified states will be returned. (Not compatible with deprecated `exclusiveStartId` parameter.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsAPI.RequestQueueRequestsGet(context.Background(), queueId).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsAPI.RequestQueueRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestsGet`: ListOfRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsAPI.RequestQueueRequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestsGetRequest struct via the builder pattern


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


## RequestQueueRequestsPost

> AddRequestResponse RequestQueueRequestsPost(ctx, queueId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()

Add request



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
	requestBase := *openapiclient.NewRequestBase() // RequestBase | 
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesRequestsAPI.RequestQueueRequestsPost(context.Background(), queueId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesRequestsAPI.RequestQueueRequestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestsPost`: AddRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesRequestsAPI.RequestQueueRequestsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestsPostRequest struct via the builder pattern


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

