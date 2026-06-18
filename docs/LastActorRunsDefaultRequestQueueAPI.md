# \LastActorRunsDefaultRequestQueueAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActRunsLastRequestQueueDelete**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueDelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue | Delete last run&#39;s default request queue
[**ActRunsLastRequestQueueGet**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueGet) | **Get** /v2/actors/{actorId}/runs/last/request-queue | Get last run&#39;s default request queue
[**ActRunsLastRequestQueueHeadGet**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueHeadGet) | **Get** /v2/actors/{actorId}/runs/last/request-queue/head | Get last run&#39;s default request queue head
[**ActRunsLastRequestQueueHeadLockPost**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueHeadLockPost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/head/lock | Get and lock last run&#39;s default request queue head
[**ActRunsLastRequestQueuePut**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueuePut) | **Put** /v2/actors/{actorId}/runs/last/request-queue | Update last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestDelete**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestDelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId} | Delete request from last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestGet**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestGet) | **Get** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId} | Get request from last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestLockDelete**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestLockDelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}/lock | Delete lock on request in last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestLockPut**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestLockPut) | **Put** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId}/lock | Prolong lock on request in last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestPut**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestPut) | **Put** /v2/actors/{actorId}/runs/last/request-queue/requests/{requestId} | Update request in last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestsBatchDelete**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestsBatchDelete) | **Delete** /v2/actors/{actorId}/runs/last/request-queue/requests/batch | Batch delete requests from last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestsBatchPost**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestsBatchPost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/requests/batch | Batch add requests to last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestsGet**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestsGet) | **Get** /v2/actors/{actorId}/runs/last/request-queue/requests | List last run&#39;s default request queue&#39;s requests
[**ActRunsLastRequestQueueRequestsPost**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestsPost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/requests | Add request to last run&#39;s default request queue
[**ActRunsLastRequestQueueRequestsUnlockPost**](LastActorRunsDefaultRequestQueueAPI.md#ActRunsLastRequestQueueRequestsUnlockPost) | **Post** /v2/actors/{actorId}/runs/last/request-queue/requests/unlock | Unlock requests in last run&#39;s default request queue



## ActRunsLastRequestQueueDelete

> ActRunsLastRequestQueueDelete(ctx, actorId).Status(status).Execute()

Delete last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueDelete(context.Background(), actorId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastRequestQueueGet

> RequestQueueResponse ActRunsLastRequestQueueGet(ctx, actorId).Status(status).Execute()

Get last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueGet(context.Background(), actorId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueGet`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastRequestQueueHeadGet

> HeadResponse ActRunsLastRequestQueueHeadGet(ctx, actorId).Status(status).Limit(limit).ClientKey(clientKey).Execute()

Get last run's default request queue head



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	limit := float64(100) // float64 | How many items from queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueHeadGet(context.Background(), actorId).Status(status).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueHeadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueHeadGet`: HeadResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueHeadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueHeadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueHeadLockPost

> HeadAndLockResponse ActRunsLastRequestQueueHeadLockPost(ctx, actorId).LockSecs(lockSecs).Status(status).Limit(limit).ClientKey(clientKey).Execute()

Get and lock last run's default request queue head



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	limit := float64(25) // float64 | How many items from the queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueHeadLockPost(context.Background(), actorId).LockSecs(lockSecs).Status(status).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueHeadLockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueHeadLockPost`: HeadAndLockResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueHeadLockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueHeadLockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockSecs** | **float64** | How long the requests will be locked for (in seconds). | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueuePut

> RequestQueueResponse ActRunsLastRequestQueuePut(ctx, actorId).UpdateRequestQueueRequest(updateRequestQueueRequest).Status(status).Execute()

Update last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	updateRequestQueueRequest := *openapiclient.NewUpdateRequestQueueRequest() // UpdateRequestQueueRequest | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueuePut(context.Background(), actorId).UpdateRequestQueueRequest(updateRequestQueueRequest).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueuePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueuePut`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueuePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueuePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequestQueueRequest** | [**UpdateRequestQueueRequest**](UpdateRequestQueueRequest.md) |  | 
 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastRequestQueueRequestDelete

> ActRunsLastRequestQueueRequestDelete(ctx, actorId, requestId).Status(status).ClientKey(clientKey).Execute()

Delete request from last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestDelete(context.Background(), actorId, requestId).Status(status).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestGet

> RequestResponse ActRunsLastRequestQueueRequestGet(ctx, actorId, requestId).Status(status).Execute()

Get request from last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestGet(context.Background(), actorId, requestId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestGet`: RequestResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastRequestQueueRequestLockDelete

> ActRunsLastRequestQueueRequestLockDelete(ctx, actorId, requestId).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Delete lock on request in last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock was removed.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestLockDelete(context.Background(), actorId, requestId).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestLockDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestLockDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestLockPut

> ProlongRequestLockResponse ActRunsLastRequestQueueRequestLockPut(ctx, actorId, requestId).LockSecs(lockSecs).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Prolong lock on request in last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock expires.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestLockPut(context.Background(), actorId, requestId).LockSecs(lockSecs).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestLockPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestLockPut`: ProlongRequestLockResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestLockPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestLockPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lockSecs** | **float64** | How long the requests will be locked for (in seconds). | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestPut

> UpdateRequestResponse ActRunsLastRequestQueueRequestPut(ctx, actorId, requestId).Request(request).Status(status).Forefront(forefront).ClientKey(clientKey).Execute()

Update request in last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	request := *openapiclient.NewRequest() // Request | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestPut(context.Background(), actorId, requestId).Request(request).Status(status).Forefront(forefront).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestPut`: UpdateRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**Request**](Request.md) |  | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestsBatchDelete

> BatchDeleteResponse ActRunsLastRequestQueueRequestsBatchDelete(ctx, actorId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).Status(status).ClientKey(clientKey).Execute()

Batch delete requests from last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	contentType := "contentType_example" // string | 
	requestDraftDelete := []openapiclient.RequestDraftDelete{*openapiclient.NewRequestDraftDelete("Id_example", "UniqueKey_example")} // []RequestDraftDelete | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsBatchDelete(context.Background(), actorId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).Status(status).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsBatchDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestsBatchDelete`: BatchDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsBatchDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestsBatchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **requestDraftDelete** | [**[]RequestDraftDelete**](RequestDraftDelete.md) |  | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestsBatchPost

> BatchAddResponse ActRunsLastRequestQueueRequestsBatchPost(ctx, actorId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Batch add requests to last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	requestBase := []openapiclient.RequestBase{*openapiclient.NewRequestBase()} // []RequestBase | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsBatchPost(context.Background(), actorId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestsBatchPost`: BatchAddResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsBatchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestsBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBase** | [**[]RequestBase**](RequestBase.md) |  | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestsGet

> ListOfRequestsResponse ActRunsLastRequestQueueRequestsGet(ctx, actorId).Status(status).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()

List last run's default request queue's requests



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	exclusiveStartId := "Ihnsp8YrvJ8102Kj" // string | All requests up to this one (including) are skipped from the result. (Deprecated, use `cursor` instead.) (optional)
	limit := float64(100) // float64 | Number of keys to be returned. Maximum value is `10000`. (optional)
	cursor := "eyJyZXF1ZXN0SWQiOiI2OFRqQ2RaTDNvM2hiUU0ifQ" // string | A cursor string for pagination, returned in the previous response as `nextCursor`. Use this to retrieve the next page of requests. (optional)
	filter := []string{"Filter_example"} // []string | Filter requests by their state. Possible values are `locked` and `pending`. You can combine multiple values separated by commas, which will mean the union of these filters – requests matching any of the specified states will be returned. (Not compatible with deprecated `exclusiveStartId` parameter.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsGet(context.Background(), actorId).Status(status).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestsGet`: ListOfRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestsPost

> AddRequestResponse ActRunsLastRequestQueueRequestsPost(ctx, actorId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Add request to last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	requestBase := *openapiclient.NewRequestBase() // RequestBase | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsPost(context.Background(), actorId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestsPost`: AddRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBase** | [**RequestBase**](RequestBase.md) |  | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastRequestQueueRequestsUnlockPost

> UnlockRequestsResponse ActRunsLastRequestQueueRequestsUnlockPost(ctx, actorId).Status(status).ClientKey(clientKey).Execute()

Unlock requests in last run's default request queue



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsUnlockPost(context.Background(), actorId).Status(status).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsUnlockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastRequestQueueRequestsUnlockPost`: UnlockRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultRequestQueueAPI.ActRunsLastRequestQueueRequestsUnlockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastRequestQueueRequestsUnlockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
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

