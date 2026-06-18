# \ActorBuildsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorBuildAbortPost**](ActorBuildsAPI.md#ActorBuildAbortPost) | **Post** /v2/actor-builds/{buildId}/abort | Abort build
[**ActorBuildDelete**](ActorBuildsAPI.md#ActorBuildDelete) | **Delete** /v2/actor-builds/{buildId} | Delete build
[**ActorBuildGet**](ActorBuildsAPI.md#ActorBuildGet) | **Get** /v2/actor-builds/{buildId} | Get build
[**ActorBuildLogGet**](ActorBuildsAPI.md#ActorBuildLogGet) | **Get** /v2/actor-builds/{buildId}/log | Get build&#39;s Log
[**ActorBuildOpenapiJsonGet**](ActorBuildsAPI.md#ActorBuildOpenapiJsonGet) | **Get** /v2/actor-builds/{buildId}/openapi.json | Get OpenAPI definition
[**ActorBuildsGet**](ActorBuildsAPI.md#ActorBuildsGet) | **Get** /v2/actor-builds | Get user builds list



## ActorBuildAbortPost

> BuildResponse ActorBuildAbortPost(ctx, buildId).Execute()

Abort build



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
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorBuildsAPI.ActorBuildAbortPost(context.Background(), buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorBuildsAPI.ActorBuildAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorBuildAbortPost`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorBuildsAPI.ActorBuildAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorBuildAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorBuildDelete

> ActorBuildDelete(ctx, buildId).Execute()

Delete build



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
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActorBuildsAPI.ActorBuildDelete(context.Background(), buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorBuildsAPI.ActorBuildDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorBuildDeleteRequest struct via the builder pattern


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


## ActorBuildGet

> BuildResponse ActorBuildGet(ctx, buildId).WaitForFinish(waitForFinish).Execute()

Get build



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
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab.
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the build to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the build finishes in time then the returned build object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorBuildsAPI.ActorBuildGet(context.Background(), buildId).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorBuildsAPI.ActorBuildGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorBuildGet`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorBuildsAPI.ActorBuildGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorBuildGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **waitForFinish** | **float64** | The maximum number of seconds the server waits for the build to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the build finishes in time then the returned build object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;).  | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorBuildLogGet

> string ActorBuildLogGet(ctx, buildId).Stream(stream).Download(download).Execute()

Get build's Log



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
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab.
	stream := false // bool | If `true` or `1` then the logs will be streamed as long as the run or build is running.  (optional)
	download := false // bool | If `true` or `1` then the web browser will download the log file rather than open it in a tab.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorBuildsAPI.ActorBuildLogGet(context.Background(), buildId).Stream(stream).Download(download).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorBuildsAPI.ActorBuildLogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorBuildLogGet`: string
	fmt.Fprintf(os.Stdout, "Response from `ActorBuildsAPI.ActorBuildLogGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorBuildLogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stream** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the logs will be streamed as long as the run or build is running.  | 
 **download** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the web browser will download the log file rather than open it in a tab.  | 

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


## ActorBuildOpenapiJsonGet

> map[string]interface{} ActorBuildOpenapiJsonGet(ctx, buildId).Execute()

Get OpenAPI definition



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
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab. Use the special value `default` to get the OpenAPI schema for the Actor's default build. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorBuildsAPI.ActorBuildOpenapiJsonGet(context.Background(), buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorBuildsAPI.ActorBuildOpenapiJsonGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorBuildOpenapiJsonGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorBuildsAPI.ActorBuildOpenapiJsonGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. Use the special value &#x60;default&#x60; to get the OpenAPI schema for the Actor&#39;s default build.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorBuildOpenapiJsonGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorBuildsGet

> ListOfBuildsResponse ActorBuildsGet(ctx).Offset(offset).Limit(limit).Desc(desc).Execute()

Get user builds list



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
	desc := true // bool | If `true` or `1` then the objects are sorted by the `startedAt` field in descending order. By default, they are sorted in ascending order.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorBuildsAPI.ActorBuildsGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorBuildsAPI.ActorBuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorBuildsGet`: ListOfBuildsResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorBuildsAPI.ActorBuildsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActorBuildsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;startedAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 

### Return type

[**ListOfBuildsResponse**](ListOfBuildsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

