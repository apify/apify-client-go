# \LastActorTaskRunsDefaultRequestQueueAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorTaskRunsLastRequestQueueDelete**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueDelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue | Delete last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueGet**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue | Get last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueHeadGet**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueHeadGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/head | Get last task run&#39;s default request queue head
[**ActorTaskRunsLastRequestQueueHeadLockPost**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueHeadLockPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/head/lock | Get and lock last task run&#39;s default request queue head
[**ActorTaskRunsLastRequestQueuePut**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueuePut) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue | Update last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestDelete**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestDelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId} | Delete request from last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestGet**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId} | Get request from last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestLockDelete**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestLockDelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId}/lock | Delete lock on request in last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestLockPut**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestLockPut) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId}/lock | Prolong lock on request in last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestPut**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestPut) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/{requestId} | Update request in last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestsBatchDelete**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestsBatchDelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/batch | Batch delete requests from last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestsBatchPost**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestsBatchPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/batch | Batch add requests to last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestsGet**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestsGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests | List last task run&#39;s default request queue&#39;s requests
[**ActorTaskRunsLastRequestQueueRequestsPost**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestsPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests | Add request to last task run&#39;s default request queue
[**ActorTaskRunsLastRequestQueueRequestsUnlockPost**](LastActorTaskRunsDefaultRequestQueueAPI.md#ActorTaskRunsLastRequestQueueRequestsUnlockPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/request-queue/requests/unlock | Unlock requests in last task run&#39;s default request queue



## ActorTaskRunsLastRequestQueueDelete

> ActorTaskRunsLastRequestQueueDelete(ctx, actorTaskId).Status(status).Execute()

Delete last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueDelete(context.Background(), actorTaskId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueDeleteRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueGet

> RequestQueueResponse ActorTaskRunsLastRequestQueueGet(ctx, actorTaskId).Status(status).Execute()

Get last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueGet(context.Background(), actorTaskId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueGet`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueGetRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueHeadGet

> HeadResponse ActorTaskRunsLastRequestQueueHeadGet(ctx, actorTaskId).Status(status).Limit(limit).ClientKey(clientKey).Execute()

Get last task run's default request queue head



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	limit := float64(100) // float64 | How many items from queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueHeadGet(context.Background(), actorTaskId).Status(status).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueHeadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueHeadGet`: HeadResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueHeadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueHeadGetRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueHeadLockPost

> HeadAndLockResponse ActorTaskRunsLastRequestQueueHeadLockPost(ctx, actorTaskId).LockSecs(lockSecs).Status(status).Limit(limit).ClientKey(clientKey).Execute()

Get and lock last task run's default request queue head



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	limit := float64(25) // float64 | How many items from the queue should be returned. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueHeadLockPost(context.Background(), actorTaskId).LockSecs(lockSecs).Status(status).Limit(limit).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueHeadLockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueHeadLockPost`: HeadAndLockResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueHeadLockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueHeadLockPostRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueuePut

> RequestQueueResponse ActorTaskRunsLastRequestQueuePut(ctx, actorTaskId).UpdateRequestQueueRequest(updateRequestQueueRequest).Status(status).Execute()

Update last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	updateRequestQueueRequest := *openapiclient.NewUpdateRequestQueueRequest() // UpdateRequestQueueRequest | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueuePut(context.Background(), actorTaskId).UpdateRequestQueueRequest(updateRequestQueueRequest).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueuePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueuePut`: RequestQueueResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueuePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueuePutRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestDelete

> ActorTaskRunsLastRequestQueueRequestDelete(ctx, actorTaskId, requestId).Status(status).ClientKey(clientKey).Execute()

Delete request from last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestDelete(context.Background(), actorTaskId, requestId).Status(status).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestDeleteRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestGet

> RequestResponse ActorTaskRunsLastRequestQueueRequestGet(ctx, actorTaskId, requestId).Status(status).Execute()

Get request from last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestGet(context.Background(), actorTaskId, requestId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestGet`: RequestResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestGetRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestLockDelete

> ActorTaskRunsLastRequestQueueRequestLockDelete(ctx, actorTaskId, requestId).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Delete lock on request in last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock was removed.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestLockDelete(context.Background(), actorTaskId, requestId).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestLockDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestLockDeleteRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestLockPut

> ProlongRequestLockResponse ActorTaskRunsLastRequestQueueRequestLockPut(ctx, actorTaskId, requestId).LockSecs(lockSecs).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Prolong lock on request in last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	lockSecs := float64(60) // float64 | How long the requests will be locked for (in seconds).
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end after lock expires.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestLockPut(context.Background(), actorTaskId, requestId).LockSecs(lockSecs).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestLockPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestLockPut`: ProlongRequestLockResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestLockPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestLockPutRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestPut

> UpdateRequestResponse ActorTaskRunsLastRequestQueueRequestPut(ctx, actorTaskId, requestId).Request(request).Status(status).Forefront(forefront).ClientKey(clientKey).Execute()

Update request in last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	requestId := "xpsmkDlspokDSmklS" // string | Request ID.
	request := *openapiclient.NewRequest() // Request | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestPut(context.Background(), actorTaskId, requestId).Request(request).Status(status).Forefront(forefront).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestPut`: UpdateRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**requestId** | **string** | Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestPutRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestsBatchDelete

> BatchDeleteResponse ActorTaskRunsLastRequestQueueRequestsBatchDelete(ctx, actorTaskId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).Status(status).ClientKey(clientKey).Execute()

Batch delete requests from last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	contentType := "contentType_example" // string | 
	requestDraftDelete := []openapiclient.RequestDraftDelete{*openapiclient.NewRequestDraftDelete("Id_example", "UniqueKey_example")} // []RequestDraftDelete | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsBatchDelete(context.Background(), actorTaskId).ContentType(contentType).RequestDraftDelete(requestDraftDelete).Status(status).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsBatchDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestsBatchDelete`: BatchDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsBatchDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestsBatchDeleteRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestsBatchPost

> BatchAddResponse ActorTaskRunsLastRequestQueueRequestsBatchPost(ctx, actorTaskId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Batch add requests to last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	requestBase := []openapiclient.RequestBase{*openapiclient.NewRequestBase()} // []RequestBase | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsBatchPost(context.Background(), actorTaskId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestsBatchPost`: BatchAddResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsBatchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestsBatchPostRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestsGet

> ListOfRequestsResponse ActorTaskRunsLastRequestQueueRequestsGet(ctx, actorTaskId).Status(status).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()

List last task run's default request queue's requests



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	exclusiveStartId := "Ihnsp8YrvJ8102Kj" // string | All requests up to this one (including) are skipped from the result. (Deprecated, use `cursor` instead.) (optional)
	limit := float64(100) // float64 | Number of keys to be returned. Maximum value is `10000`. (optional)
	cursor := "eyJyZXF1ZXN0SWQiOiI2OFRqQ2RaTDNvM2hiUU0ifQ" // string | A cursor string for pagination, returned in the previous response as `nextCursor`. Use this to retrieve the next page of requests. (optional)
	filter := []string{"Filter_example"} // []string | Filter requests by their state. Possible values are `locked` and `pending`. You can combine multiple values separated by commas, which will mean the union of these filters – requests matching any of the specified states will be returned. (Not compatible with deprecated `exclusiveStartId` parameter.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsGet(context.Background(), actorTaskId).Status(status).ClientKey(clientKey).ExclusiveStartId(exclusiveStartId).Limit(limit).Cursor(cursor).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestsGet`: ListOfRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestsGetRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestsPost

> AddRequestResponse ActorTaskRunsLastRequestQueueRequestsPost(ctx, actorTaskId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()

Add request to last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	requestBase := *openapiclient.NewRequestBase() // RequestBase | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)
	forefront := "false" // string | Determines if request should be added to the head of the queue or to the end. Default value is `false` (end of queue).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsPost(context.Background(), actorTaskId).RequestBase(requestBase).Status(status).ClientKey(clientKey).Forefront(forefront).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestsPost`: AddRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestsPostRequest struct via the builder pattern


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


## ActorTaskRunsLastRequestQueueRequestsUnlockPost

> UnlockRequestsResponse ActorTaskRunsLastRequestQueueRequestsUnlockPost(ctx, actorTaskId).Status(status).ClientKey(clientKey).Execute()

Unlock requests in last task run's default request queue



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	clientKey := "client-abc" // string | A unique identifier of the client accessing the request queue. It must be a string between 1 and 32 characters long. This identifier is used to determine whether the queue was accessed by multiple clients. If `clientKey` is not provided, the system considers this API call to come from a new client. For details, see the `hadMultipleClients` field returned by the [Get head](#/reference/request-queues/queue-head) operation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsUnlockPost(context.Background(), actorTaskId).Status(status).ClientKey(clientKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsUnlockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastRequestQueueRequestsUnlockPost`: UnlockRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultRequestQueueAPI.ActorTaskRunsLastRequestQueueRequestsUnlockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastRequestQueueRequestsUnlockPostRequest struct via the builder pattern


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

