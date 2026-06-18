# \LogsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogGet**](LogsAPI.md#LogGet) | **Get** /v2/logs/{buildOrRunId} | Get log



## LogGet

> string LogGet(ctx, buildOrRunId).Stream(stream).Download(download).Raw(raw).Execute()

Get log



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
	buildOrRunId := "HG7ML7M8z78YcAPE" // string | ID of the Actor build or run.
	stream := false // bool | If `true` or `1` then the logs will be streamed as long as the run or build is running.  (optional)
	download := false // bool | If `true` or `1` then the web browser will download the log file rather than open it in a tab.  (optional)
	raw := false // bool | If `true` or `1`, the logs will be kept verbatim. By default, the API removes ANSI escape codes from the logs, keeping only printable characters.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.LogGet(context.Background(), buildOrRunId).Stream(stream).Download(download).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogGet`: string
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.LogGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildOrRunId** | **string** | ID of the Actor build or run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogGetRequest struct via the builder pattern


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

