# \ActorsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActDelete**](ActorsAPI.md#ActDelete) | **Delete** /v2/actors/{actorId} | Delete Actor
[**ActGet**](ActorsAPI.md#ActGet) | **Get** /v2/actors/{actorId} | Get Actor
[**ActPut**](ActorsAPI.md#ActPut) | **Put** /v2/actors/{actorId} | Update Actor
[**ActValidateInputPost**](ActorsAPI.md#ActValidateInputPost) | **Post** /v2/actors/{actorId}/validate-input | Validate Actor input
[**ActsGet**](ActorsAPI.md#ActsGet) | **Get** /v2/actors | Get list of Actors
[**ActsPost**](ActorsAPI.md#ActsPost) | **Post** /v2/actors | Create Actor



## ActDelete

> map[string]interface{} ActDelete(ctx, actorId).Execute()

Delete Actor



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsAPI.ActDelete(context.Background(), actorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActGet

> ActorResponse ActGet(ctx, actorId).Execute()

Get Actor



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsAPI.ActGet(context.Background(), actorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActGet`: ActorResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActorResponse**](ActorResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActPut

> ActorResponse ActPut(ctx, actorId).UpdateActorRequest(updateActorRequest).Execute()

Update Actor



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
	updateActorRequest := *openapiclient.NewUpdateActorRequest() // UpdateActorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsAPI.ActPut(context.Background(), actorId).UpdateActorRequest(updateActorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActPut`: ActorResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateActorRequest** | [**UpdateActorRequest**](UpdateActorRequest.md) |  | 

### Return type

[**ActorResponse**](ActorResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActValidateInputPost

> ActValidateInputPost200Response ActValidateInputPost(ctx, actorId).Body(body).Build(build).Execute()

Validate Actor input



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
	body := map[string]interface{}{ ... } // map[string]interface{} | JSON input to validate against the Actor's input schema.
	build := "latest" // string | Optional tag or number of the Actor build to use for input schema validation. By default, the `latest` build tag is used.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsAPI.ActValidateInputPost(context.Background(), actorId).Body(body).Build(build).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActValidateInputPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActValidateInputPost`: ActValidateInputPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActValidateInputPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActValidateInputPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | JSON input to validate against the Actor&#39;s input schema. | 
 **build** | **string** | Optional tag or number of the Actor build to use for input schema validation. By default, the &#x60;latest&#x60; build tag is used.  | 

### Return type

[**ActValidateInputPost200Response**](ActValidateInputPost200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActsGet

> ListOfActorsResponse ActsGet(ctx).My(my).Offset(offset).Limit(limit).Desc(desc).SortBy(sortBy).Execute()

Get list of Actors



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
	my := true // bool | If `true` or `1` then the returned list only contains Actors owned by the user. The default value is `false`.  (optional)
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	desc := true // bool | If `true` or `1` then the objects are sorted by the `createdAt` field in descending order. By default, they are sorted in ascending order.  (optional)
	sortBy := "createdAt" // string | Field to sort the records by. The default is `createdAt`. You can also use `stats.lastRunStartedAt` to sort by the most recently ran Actors.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsAPI.ActsGet(context.Background()).My(my).Offset(offset).Limit(limit).Desc(desc).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActsGet`: ListOfActorsResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **my** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the returned list only contains Actors owned by the user. The default value is &#x60;false&#x60;.  | 
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 
 **sortBy** | **string** | Field to sort the records by. The default is &#x60;createdAt&#x60;. You can also use &#x60;stats.lastRunStartedAt&#x60; to sort by the most recently ran Actors.  | 

### Return type

[**ListOfActorsResponse**](ListOfActorsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActsPost

> ActorResponse ActsPost(ctx).CreateActorRequest(createActorRequest).Execute()

Create Actor



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
	createActorRequest := *openapiclient.NewCreateActorRequest() // CreateActorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsAPI.ActsPost(context.Background()).CreateActorRequest(createActorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActsPost`: ActorResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createActorRequest** | [**CreateActorRequest**](CreateActorRequest.md) |  | 

### Return type

[**ActorResponse**](ActorResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

