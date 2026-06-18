# \LastActorRunsAbortAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActRunsLastAbortPost**](LastActorRunsAbortAPI.md#ActRunsLastAbortPost) | **Post** /v2/actors/{actorId}/runs/last/abort | Abort Actor&#39;s last run



## ActRunsLastAbortPost

> RunResponse ActRunsLastAbortPost(ctx, actorId).Status(status).Gracefully(gracefully).Execute()

Abort Actor's last run



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
	gracefully := true // bool | If true passed, the Actor run will abort gracefully. It will send `aborting` and `persistState` event into run and force-stop the run after 30 seconds. It is helpful in cases where you plan to resurrect the run later.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsAbortAPI.ActRunsLastAbortPost(context.Background(), actorId).Status(status).Gracefully(gracefully).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsAbortAPI.ActRunsLastAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastAbortPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsAbortAPI.ActRunsLastAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastAbortPostRequest struct via the builder pattern


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

