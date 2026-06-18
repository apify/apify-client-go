# \LastActorRunsMetamorphAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActRunsLastMetamorphPost**](LastActorRunsMetamorphAPI.md#ActRunsLastMetamorphPost) | **Post** /v2/actors/{actorId}/runs/last/metamorph | Metamorph Actor&#39;s last run



## ActRunsLastMetamorphPost

> RunResponse ActRunsLastMetamorphPost(ctx, actorId).TargetActorId(targetActorId).Status(status).Build(build).Execute()

Metamorph Actor's last run



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
	targetActorId := "HDSasDasz78YcAPEB" // string | ID of a target Actor that the run should be transformed into.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	build := "beta" // string | Optional build of the target Actor.  It can be either a build tag or build number. By default, the run uses the build specified in the default run configuration for the target Actor (typically `latest`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsMetamorphAPI.ActRunsLastMetamorphPost(context.Background(), actorId).TargetActorId(targetActorId).Status(status).Build(build).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsMetamorphAPI.ActRunsLastMetamorphPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastMetamorphPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsMetamorphAPI.ActRunsLastMetamorphPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastMetamorphPostRequest struct via the builder pattern


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

