# \LastActorTaskRunsMetamorphAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorTaskRunsLastMetamorphPost**](LastActorTaskRunsMetamorphAPI.md#ActorTaskRunsLastMetamorphPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/metamorph | Metamorph Actor task&#39;s last run



## ActorTaskRunsLastMetamorphPost

> RunResponse ActorTaskRunsLastMetamorphPost(ctx, actorTaskId).TargetActorId(targetActorId).Status(status).Build(build).Execute()

Metamorph Actor task's last run



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
	targetActorId := "HDSasDasz78YcAPEB" // string | ID of a target Actor that the run should be transformed into.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	build := "beta" // string | Optional build of the target Actor.  It can be either a build tag or build number. By default, the run uses the build specified in the default run configuration for the target Actor (typically `latest`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsMetamorphAPI.ActorTaskRunsLastMetamorphPost(context.Background(), actorTaskId).TargetActorId(targetActorId).Status(status).Build(build).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsMetamorphAPI.ActorTaskRunsLastMetamorphPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastMetamorphPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsMetamorphAPI.ActorTaskRunsLastMetamorphPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastMetamorphPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetActorId** | **string** | ID of a target Actor that the run should be transformed into. | 
 **status** | **string** | Filter for the run status. | 
 **build** | **string** | Optional build of the target Actor.  It can be either a build tag or build number. By default, the run uses the build specified in the default run configuration for the target Actor (typically &#x60;latest&#x60;).  | 

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

