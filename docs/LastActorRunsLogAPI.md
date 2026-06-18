# \LastActorRunsLogAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActRunsLastLogGet**](LastActorRunsLogAPI.md#ActRunsLastLogGet) | **Get** /v2/actors/{actorId}/runs/last/log | Get last Actor run&#39;s log



## ActRunsLastLogGet

> string ActRunsLastLogGet(ctx, actorId).Stream(stream).Download(download).Raw(raw).Execute()

Get last Actor run's log



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
	stream := false // bool | If `true` or `1` then the logs will be streamed as long as the run or build is running.  (optional)
	download := false // bool | If `true` or `1` then the web browser will download the log file rather than open it in a tab.  (optional)
	raw := false // bool | If `true` or `1`, the logs will be kept verbatim. By default, the API removes ANSI escape codes from the logs, keeping only printable characters.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsLogAPI.ActRunsLastLogGet(context.Background(), actorId).Stream(stream).Download(download).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsLogAPI.ActRunsLastLogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastLogGet`: string
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsLogAPI.ActRunsLastLogGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastLogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stream** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the logs will be streamed as long as the run or build is running.  | 
 **download** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the web browser will download the log file rather than open it in a tab.  | 
 **raw** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the logs will be kept verbatim. By default, the API removes ANSI escape codes from the logs, keeping only printable characters.  | 

### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

