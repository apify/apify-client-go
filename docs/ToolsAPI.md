# \ToolsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ToolsBrowserInfoDelete**](ToolsAPI.md#ToolsBrowserInfoDelete) | **Delete** /v2/browser-info | Get browser info
[**ToolsBrowserInfoGet**](ToolsAPI.md#ToolsBrowserInfoGet) | **Get** /v2/browser-info | Get browser info
[**ToolsBrowserInfoPost**](ToolsAPI.md#ToolsBrowserInfoPost) | **Post** /v2/browser-info | Get browser info
[**ToolsBrowserInfoPut**](ToolsAPI.md#ToolsBrowserInfoPut) | **Put** /v2/browser-info | Get browser info
[**ToolsDecodeAndVerifyPost**](ToolsAPI.md#ToolsDecodeAndVerifyPost) | **Post** /v2/tools/decode-and-verify | Decode and verify object
[**ToolsEncodeAndSignPost**](ToolsAPI.md#ToolsEncodeAndSignPost) | **Post** /v2/tools/encode-and-sign | Encode and sign object



## ToolsBrowserInfoDelete

> BrowserInfoResponse ToolsBrowserInfoDelete(ctx).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()

Get browser info



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
	skipHeaders := true // bool | If `true` or `1`, the response omits the `headers` field. (optional)
	rawHeaders := true // bool | If `true` or `1`, the response includes the `rawHeaders` field with the raw request headers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsBrowserInfoDelete(context.Background()).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsBrowserInfoDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsBrowserInfoDelete`: BrowserInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsBrowserInfoDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsBrowserInfoDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response omits the &#x60;headers&#x60; field. | 
 **rawHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response includes the &#x60;rawHeaders&#x60; field with the raw request headers. | 

### Return type

[**BrowserInfoResponse**](BrowserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsBrowserInfoGet

> BrowserInfoResponse ToolsBrowserInfoGet(ctx).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()

Get browser info



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
	skipHeaders := true // bool | If `true` or `1`, the response omits the `headers` field. (optional)
	rawHeaders := true // bool | If `true` or `1`, the response includes the `rawHeaders` field with the raw request headers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsBrowserInfoGet(context.Background()).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsBrowserInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsBrowserInfoGet`: BrowserInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsBrowserInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsBrowserInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response omits the &#x60;headers&#x60; field. | 
 **rawHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response includes the &#x60;rawHeaders&#x60; field with the raw request headers. | 

### Return type

[**BrowserInfoResponse**](BrowserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsBrowserInfoPost

> BrowserInfoResponse ToolsBrowserInfoPost(ctx).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()

Get browser info



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
	skipHeaders := true // bool | If `true` or `1`, the response omits the `headers` field. (optional)
	rawHeaders := true // bool | If `true` or `1`, the response includes the `rawHeaders` field with the raw request headers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsBrowserInfoPost(context.Background()).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsBrowserInfoPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsBrowserInfoPost`: BrowserInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsBrowserInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsBrowserInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response omits the &#x60;headers&#x60; field. | 
 **rawHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response includes the &#x60;rawHeaders&#x60; field with the raw request headers. | 

### Return type

[**BrowserInfoResponse**](BrowserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsBrowserInfoPut

> BrowserInfoResponse ToolsBrowserInfoPut(ctx).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()

Get browser info



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
	skipHeaders := true // bool | If `true` or `1`, the response omits the `headers` field. (optional)
	rawHeaders := true // bool | If `true` or `1`, the response includes the `rawHeaders` field with the raw request headers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsBrowserInfoPut(context.Background()).SkipHeaders(skipHeaders).RawHeaders(rawHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsBrowserInfoPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsBrowserInfoPut`: BrowserInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsBrowserInfoPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsBrowserInfoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response omits the &#x60;headers&#x60; field. | 
 **rawHeaders** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response includes the &#x60;rawHeaders&#x60; field with the raw request headers. | 

### Return type

[**BrowserInfoResponse**](BrowserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsDecodeAndVerifyPost

> DecodeAndVerifyResponse ToolsDecodeAndVerifyPost(ctx).DecodeAndVerifyRequest(decodeAndVerifyRequest).Execute()

Decode and verify object



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
	decodeAndVerifyRequest := *openapiclient.NewDecodeAndVerifyRequest("Encoded_example") // DecodeAndVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsDecodeAndVerifyPost(context.Background()).DecodeAndVerifyRequest(decodeAndVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsDecodeAndVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsDecodeAndVerifyPost`: DecodeAndVerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsDecodeAndVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsDecodeAndVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decodeAndVerifyRequest** | [**DecodeAndVerifyRequest**](DecodeAndVerifyRequest.md) |  | 

### Return type

[**DecodeAndVerifyResponse**](DecodeAndVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolsEncodeAndSignPost

> EncodeAndSignResponse ToolsEncodeAndSignPost(ctx).Body(body).Execute()

Encode and sign object



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.ToolsEncodeAndSignPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ToolsEncodeAndSignPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolsEncodeAndSignPost`: EncodeAndSignResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ToolsEncodeAndSignPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolsEncodeAndSignPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**EncodeAndSignResponse**](EncodeAndSignResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

