# \WebhooksWebhooksAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookDelete**](WebhooksWebhooksAPI.md#WebhookDelete) | **Delete** /v2/webhooks/{webhookId} | Delete webhook
[**WebhookGet**](WebhooksWebhooksAPI.md#WebhookGet) | **Get** /v2/webhooks/{webhookId} | Get webhook
[**WebhookPut**](WebhooksWebhooksAPI.md#WebhookPut) | **Put** /v2/webhooks/{webhookId} | Update webhook
[**WebhookTestPost**](WebhooksWebhooksAPI.md#WebhookTestPost) | **Post** /v2/webhooks/{webhookId}/test | Test webhook
[**WebhookWebhookDispatchesGet**](WebhooksWebhooksAPI.md#WebhookWebhookDispatchesGet) | **Get** /v2/webhooks/{webhookId}/dispatches | Get collection
[**WebhooksGet**](WebhooksWebhooksAPI.md#WebhooksGet) | **Get** /v2/webhooks | Get list of webhooks
[**WebhooksPost**](WebhooksWebhooksAPI.md#WebhooksPost) | **Post** /v2/webhooks | Create webhook



## WebhookDelete

> map[string]interface{} WebhookDelete(ctx, webhookId).Execute()

Delete webhook



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
	webhookId := "pVJtoTelgYUq4qJOt" // string | Webhook ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhooksAPI.WebhookDelete(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhooksAPI.WebhookDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhooksAPI.WebhookDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeleteRequest struct via the builder pattern


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


## WebhookGet

> WebhookResponse WebhookGet(ctx, webhookId).Execute()

Get webhook



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
	webhookId := "pVJtoTelgYUq4qJOt" // string | Webhook ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhooksAPI.WebhookGet(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhooksAPI.WebhookGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookGet`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhooksAPI.WebhookGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookPut

> WebhookResponse WebhookPut(ctx, webhookId).WebhookUpdate(webhookUpdate).Execute()

Update webhook



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
	webhookId := "pVJtoTelgYUq4qJOt" // string | Webhook ID.
	webhookUpdate := *openapiclient.NewWebhookUpdate() // WebhookUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhooksAPI.WebhookPut(context.Background(), webhookId).WebhookUpdate(webhookUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhooksAPI.WebhookPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookPut`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhooksAPI.WebhookPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookTestPost

> TestWebhookResponse WebhookTestPost(ctx, webhookId).Execute()

Test webhook



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
	webhookId := "pVJtoTelgYUq4qJOt" // string | Webhook ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhooksAPI.WebhookTestPost(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhooksAPI.WebhookTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookTestPost`: TestWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhooksAPI.WebhookTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestWebhookResponse**](TestWebhookResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookWebhookDispatchesGet

> ListOfWebhookDispatchesResponse WebhookWebhookDispatchesGet(ctx, webhookId).Execute()

Get collection



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
	webhookId := "pVJtoTelgYUq4qJOt" // string | Webhook ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhooksAPI.WebhookWebhookDispatchesGet(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhooksAPI.WebhookWebhookDispatchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookWebhookDispatchesGet`: ListOfWebhookDispatchesResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhooksAPI.WebhookWebhookDispatchesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookWebhookDispatchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## WebhooksGet

> ListOfWebhooksResponse WebhooksGet(ctx).Offset(offset).Limit(limit).Desc(desc).Execute()

Get list of webhooks



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
	resp, r, err := apiClient.WebhooksWebhooksAPI.WebhooksGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhooksAPI.WebhooksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksGet`: ListOfWebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhooksAPI.WebhooksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 

### Return type

[**ListOfWebhooksResponse**](ListOfWebhooksResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPost

> WebhookResponse WebhooksPost(ctx).WebhookCreate(webhookCreate).Execute()

Create webhook



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
	webhookCreate := *openapiclient.NewWebhookCreate([]openapiclient.WebhookEventType{openapiclient.WebhookEventType("ACTOR.BUILD.ABORTED")}, *openapiclient.NewWebhookCondition(), "RequestUrl_example") // WebhookCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksWebhooksAPI.WebhooksPost(context.Background()).WebhookCreate(webhookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksWebhooksAPI.WebhooksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksPost`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksWebhooksAPI.WebhooksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreate** | [**WebhookCreate**](WebhookCreate.md) |  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

