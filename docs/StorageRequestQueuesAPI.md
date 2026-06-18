# \StorageRequestQueuesAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestQueueDelete**](StorageRequestQueuesAPI.md#RequestQueueDelete) | **Delete** /v2/request-queues/{queueId} | Delete request queue
[**RequestQueueGet**](StorageRequestQueuesAPI.md#RequestQueueGet) | **Get** /v2/request-queues/{queueId} | Get request queue
[**RequestQueuePut**](StorageRequestQueuesAPI.md#RequestQueuePut) | **Put** /v2/request-queues/{queueId} | Update request queue
[**RequestQueueRequestsBatchDelete**](StorageRequestQueuesAPI.md#RequestQueueRequestsBatchDelete) | **Delete** /v2/request-queues/{queueId}/requests/batch | Delete requests
[**RequestQueueRequestsBatchPost**](StorageRequestQueuesAPI.md#RequestQueueRequestsBatchPost) | **Post** /v2/request-queues/{queueId}/requests/batch | Add requests
[**RequestQueuesGet**](StorageRequestQueuesAPI.md#RequestQueuesGet) | **Get** /v2/request-queues | Get list of request queues
[**RequestQueuesPost**](StorageRequestQueuesAPI.md#RequestQueuesPost) | **Post** /v2/request-queues | Create request queue



## RequestQueueDelete

> RequestQueueDelete(ctx, queueId).Execute()

Delete request queue



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageRequestQueuesAPI.RequestQueueDelete(context.Background(), queueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesAPI.RequestQueueDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueDeleteRequest struct via the builder pattern


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


## RequestQueueGet

> RequestQueueResponse RequestQueueGet(ctx, queueId).Execute()

Get request queue



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesAPI.RequestQueueGet(context.Background(), queueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesAPI.RequestQueueGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueGet`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesAPI.RequestQueueGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueGetRequest struct via the builder pattern


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


## RequestQueuePut

> RequestQueueResponse RequestQueuePut(ctx, queueId).UpdateRequestQueueRequest(updateRequestQueueRequest).Execute()

Update request queue



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
	updateRequestQueueRequest := *openapiclient.NewUpdateRequestQueueRequest() // UpdateRequestQueueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesAPI.RequestQueuePut(context.Background(), queueId).UpdateRequestQueueRequest(updateRequestQueueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesAPI.RequestQueuePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueuePut`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesAPI.RequestQueuePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueuePutRequest struct via the builder pattern


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


## RequestQueueRequestsBatchDelete

> BatchDeleteResponse RequestQueueRequestsBatchDelete(ctx, queueId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).ClientKey(clientKey).Execute()

Delete requests



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
	contentType := "contentType_example" // string | 
	requestDraftDelete := []openapiclient.RequestDraftDelete{*openapiclient.NewRequestDraftDelete("Id_example", "UniqueKey_example")} // []RequestDraftDelete | 
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesAPI.RequestQueueRequestsBatchDelete(context.Background(), queueId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesAPI.RequestQueueRequestsBatchDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestsBatchDelete`: BatchDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesAPI.RequestQueueRequestsBatchDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestsBatchDeleteRequest struct via the builder pattern


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


## RequestQueueRequestsBatchPost

> BatchAddResponse RequestQueueRequestsBatchPost(ctx, queueId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()

Add requests



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
	requestBase := []openapiclient.RequestBase{*openapiclient.NewRequestBase()} // []RequestBase | 
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesAPI.RequestQueueRequestsBatchPost(context.Background(), queueId).RequestBase(requestBase).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesAPI.RequestQueueRequestsBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueueRequestsBatchPost`: BatchAddResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesAPI.RequestQueueRequestsBatchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **string** | Queue ID or &#x60;username~queue-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueueRequestsBatchPostRequest struct via the builder pattern


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


## RequestQueuesGet

> ListOfRequestQueuesResponse RequestQueuesGet(ctx).Offset(offset).Limit(limit).Desc(desc).Unnamed(unnamed).Ownership(ownership).Execute()

Get list of request queues



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
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	desc := true // bool | If `true` or `1` then the objects are sorted by the `createdAt` field in descending order. By default, they are sorted in ascending order.  (optional)
	unnamed := true // bool | If `true` or `1` then all the storages are returned. By default, only named storages are returned.  (optional)
	ownership := openapiclient.StorageOwnership("ownedByMe") // StorageOwnership | Filter by ownership. If this parameter is omitted, all accessible request queues are returned.  - `ownedByMe`: Return only request queues owned by the user. - `sharedWithMe`: Return only request queues shared with the user by other users.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesAPI.RequestQueuesGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Unnamed(unnamed).Ownership(ownership).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesAPI.RequestQueuesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueuesGet`: ListOfRequestQueuesResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesAPI.RequestQueuesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 
 **unnamed** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then all the storages are returned. By default, only named storages are returned.  | 
 **ownership** | [**StorageOwnership**](StorageOwnership.md) | Filter by ownership. If this parameter is omitted, all accessible request queues are returned.  - &#x60;ownedByMe&#x60;: Return only request queues owned by the user. - &#x60;sharedWithMe&#x60;: Return only request queues shared with the user by other users.  | 

### Return type

[**ListOfRequestQueuesResponse**](ListOfRequestQueuesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestQueuesPost

> RequestQueueResponse RequestQueuesPost(ctx).Name(name).Execute()

Create request queue



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
	name := "example-com" // string | Custom unique name to easily identify the queue in the future. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageRequestQueuesAPI.RequestQueuesPost(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageRequestQueuesAPI.RequestQueuesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestQueuesPost`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageRequestQueuesAPI.RequestQueuesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestQueuesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Custom unique name to easily identify the queue in the future. | 

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

