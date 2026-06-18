# \ActorsActorBuildsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActBuildAbortPost**](ActorsActorBuildsAPI.md#ActBuildAbortPost) | **Post** /v2/actors/{actorId}/builds/{buildId}/abort | Abort build
[**ActBuildDefaultGet**](ActorsActorBuildsAPI.md#ActBuildDefaultGet) | **Get** /v2/actors/{actorId}/builds/default | Get default build
[**ActBuildGet**](ActorsActorBuildsAPI.md#ActBuildGet) | **Get** /v2/actors/{actorId}/builds/{buildId} | Get build
[**ActBuildsGet**](ActorsActorBuildsAPI.md#ActBuildsGet) | **Get** /v2/actors/{actorId}/builds | Get list of builds
[**ActBuildsPost**](ActorsActorBuildsAPI.md#ActBuildsPost) | **Post** /v2/actors/{actorId}/builds | Build Actor
[**ActOpenapiJsonGet**](ActorsActorBuildsAPI.md#ActOpenapiJsonGet) | **Get** /v2/actors/{actorId}/builds/{buildId}/openapi.json | Get OpenAPI definition



## ActBuildAbortPost

> BuildResponse ActBuildAbortPost(ctx, actorId, buildId).Execute()

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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorBuildsAPI.ActBuildAbortPost(context.Background(), actorId, buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorBuildsAPI.ActBuildAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActBuildAbortPost`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorBuildsAPI.ActBuildAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActBuildAbortPostRequest struct via the builder pattern


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


## ActBuildDefaultGet

> BuildResponse ActBuildDefaultGet(ctx, actorId).WaitForFinish(waitForFinish).Execute()

Get default build



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
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the build to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the build finishes in time then the returned build object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorBuildsAPI.ActBuildDefaultGet(context.Background(), actorId).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorBuildsAPI.ActBuildDefaultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActBuildDefaultGet`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorBuildsAPI.ActBuildDefaultGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActBuildDefaultGetRequest struct via the builder pattern


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


## ActBuildGet

> BuildResponse ActBuildGet(ctx, actorId, buildId).WaitForFinish(waitForFinish).Execute()

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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab.
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the build to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the build finishes in time then the returned build object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorBuildsAPI.ActBuildGet(context.Background(), actorId, buildId).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorBuildsAPI.ActBuildGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActBuildGet`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorBuildsAPI.ActBuildGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActBuildGetRequest struct via the builder pattern


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


## ActBuildsGet

> ListOfBuildsResponse ActBuildsGet(ctx, actorId).Offset(offset).Limit(limit).Desc(desc).Execute()

Get list of builds



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
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	desc := true // bool | If `true` or `1` then the objects are sorted by the `startedAt` field in descending order. By default, they are sorted in ascending order.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorBuildsAPI.ActBuildsGet(context.Background(), actorId).Offset(offset).Limit(limit).Desc(desc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorBuildsAPI.ActBuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActBuildsGet`: ListOfBuildsResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorBuildsAPI.ActBuildsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActBuildsGetRequest struct via the builder pattern


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


## ActBuildsPost

> BuildResponse ActBuildsPost(ctx, actorId).Version(version).UseCache(useCache).BetaPackages(betaPackages).Tag(tag).WaitForFinish(waitForFinish).Execute()

Build Actor



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
	version := "0.0" // string | Actor version number to be built.
	useCache := true // bool | If `true` or `1`, the system will use a cache to speed up the build process. By default, cache is not used.  (optional)
	betaPackages := true // bool | If `true` or `1` then the Actor is built with beta versions of Apify NPM packages. By default, the build uses `latest` packages.  (optional)
	tag := "latest" // string | Tag to be applied to the build on success. By default, the tag is taken from Actor version's `buildTag` property.  (optional)
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the build to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the build finishes in time then the returned build object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorBuildsAPI.ActBuildsPost(context.Background(), actorId).Version(version).UseCache(useCache).BetaPackages(betaPackages).Tag(tag).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorBuildsAPI.ActBuildsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActBuildsPost`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorBuildsAPI.ActBuildsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActBuildsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | Actor version number to be built. | 
 **useCache** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the system will use a cache to speed up the build process. By default, cache is not used.  | 
 **betaPackages** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the Actor is built with beta versions of Apify NPM packages. By default, the build uses &#x60;latest&#x60; packages.  | 
 **tag** | **string** | Tag to be applied to the build on success. By default, the tag is taken from Actor version&#39;s &#x60;buildTag&#x60; property.  | 
 **waitForFinish** | **float64** | The maximum number of seconds the server waits for the build to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the build finishes in time then the returned build object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;).  | 

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


## ActOpenapiJsonGet

> map[string]interface{} ActOpenapiJsonGet(ctx, actorId, buildId).Execute()

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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	buildId := "soSkq9ekdmfOslopH" // string | ID of the build, found in the build's Info tab. Use the special value `default` to get the OpenAPI schema for the Actor's default build. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorBuildsAPI.ActOpenapiJsonGet(context.Background(), actorId, buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorBuildsAPI.ActOpenapiJsonGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActOpenapiJsonGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorBuildsAPI.ActOpenapiJsonGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**buildId** | **string** | ID of the build, found in the build&#39;s Info tab. Use the special value &#x60;default&#x60; to get the OpenAPI schema for the Actor&#39;s default build.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActOpenapiJsonGetRequest struct via the builder pattern


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

