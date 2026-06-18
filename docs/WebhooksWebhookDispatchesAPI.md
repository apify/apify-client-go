# \WebhooksWebhookDispatchesAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookDispatchGet**](WebhooksWebhookDispatchesAPI.md#WebhookDispatchGet) | **Get** /v2/webhook-dispatches/{dispatchId} | Get webhook dispatch
[**WebhookDispatchesGet**](WebhooksWebhookDispatchesAPI.md#WebhookDispatchesGet) | **Get** /v2/webhook-dispatches | Get list of webhook dispatches



## WebhookDispatchGet

> WebhookDispatchResponse WebhookDispatchGet(ctx, dispatchId).Execute()

Get webhook dispatch



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
	dispatchId := "Zib4xbZsmvZeK55ua" // string | Webhook dispatch ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhookDispatchesAPI.WebhookDispatchGet(context.Background(), dispatchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhookDispatchesAPI.WebhookDispatchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDispatchGet`: WebhookDispatchResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhookDispatchesAPI.WebhookDispatchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dispatchId** | **string** | Webhook dispatch ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDispatchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookDispatchResponse**](WebhookDispatchResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDispatchesGet

> ListOfWebhookDispatchesResponse WebhookDispatchesGet(ctx).Offset(offset).Limit(limit).Desc(desc).Execute()

Get list of webhook dispatches



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
	desc := true // bool | If `true` or `1` then the objects are sorted by the `createdAt` field in descending order. By default, they are sorted in ascending order.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhookDispatchesAPI.WebhookDispatchesGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhookDispatchesAPI.WebhookDispatchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDispatchesGet`: ListOfWebhookDispatchesResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhookDispatchesAPI.WebhookDispatchesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDispatchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 

### Return type

[**ListOfWebhookDispatchesResponse**](ListOfWebhookDispatchesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

