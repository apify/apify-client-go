# \LastActorTaskRunsAbortAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorTaskRunsLastAbortPost**](LastActorTaskRunsAbortAPI.md#ActorTaskRunsLastAbortPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/abort | Abort Actor task&#39;s last run



## ActorTaskRunsLastAbortPost

> RunResponse ActorTaskRunsLastAbortPost(ctx, actorTaskId).Status(status).Gracefully(gracefully).Execute()

Abort Actor task's last run



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
	gracefully := true // bool | If true passed, the Actor run will abort gracefully. It will send `aborting` and `persistState` event into run and force-stop the run after 30 seconds. It is helpful in cases where you plan to resurrect the run later.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsAbortAPI.ActorTaskRunsLastAbortPost(context.Background(), actorTaskId).Status(status).Gracefully(gracefully).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsAbortAPI.ActorTaskRunsLastAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastAbortPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsAbortAPI.ActorTaskRunsLastAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
 **gracefully** | **bool** | If true passed, the Actor run will abort gracefully. It will send &#x60;aborting&#x60; and &#x60;persistState&#x60; event into run and force-stop the run after 30 seconds. It is helpful in cases where you plan to resurrect the run later.  | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

