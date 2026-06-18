# \DefaultKeyValueStoreAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorRunKeyValueStoreDelete**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreDelete) | **Delete** /v2/actor-runs/{runId}/key-value-store | Delete default store
[**ActorRunKeyValueStoreGet**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreGet) | **Get** /v2/actor-runs/{runId}/key-value-store | Get default store
[**ActorRunKeyValueStoreKeysGet**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreKeysGet) | **Get** /v2/actor-runs/{runId}/key-value-store/keys | Get default store&#39;s list of keys
[**ActorRunKeyValueStorePut**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStorePut) | **Put** /v2/actor-runs/{runId}/key-value-store | Update default store
[**ActorRunKeyValueStoreRecordDelete**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreRecordDelete) | **Delete** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Delete default store&#39;s record
[**ActorRunKeyValueStoreRecordGet**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreRecordGet) | **Get** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Get default store&#39;s record
[**ActorRunKeyValueStoreRecordPost**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreRecordPost) | **Post** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Store record in default store (POST)
[**ActorRunKeyValueStoreRecordPut**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreRecordPut) | **Put** /v2/actor-runs/{runId}/key-value-store/records/{recordKey} | Store record in default store
[**ActorRunKeyValueStoreRecordsGet**](DefaultKeyValueStoreAPI.md#ActorRunKeyValueStoreRecordsGet) | **Get** /v2/actor-runs/{runId}/key-value-store/records | Download default store&#39;s records



## ActorRunKeyValueStoreDelete

> ActorRunKeyValueStoreDelete(ctx, runId).Execute()

Delete default store



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreDelete(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreDeleteRequest struct via the builder pattern


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


## ActorRunKeyValueStoreGet

> KeyValueStoreResponse ActorRunKeyValueStoreGet(ctx, runId).Execute()

Get default store



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunKeyValueStoreGet`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyValueStoreResponse**](KeyValueStoreResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunKeyValueStoreKeysGet

> ListOfKeysResponse ActorRunKeyValueStoreKeysGet(ctx, runId).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Get default store's list of keys



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	exclusiveStartKey := "Ihnsp8YrvJ8102Kj" // string | All keys up to this one (including) are skipped from the result. (optional)
	limit := float32(100) // float32 | Number of keys to be returned. (optional) (default to 1000)
	collection := "postImages" // string | Limit the results to keys that belong to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work. (optional)
	prefix := "post-images-" // string | Limit the results to keys that start with a specific prefix. (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreKeysGet(context.Background(), runId).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunKeyValueStoreKeysGet`: ListOfKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclusiveStartKey** | **string** | All keys up to this one (including) are skipped from the result. | 
 **limit** | **float32** | Number of keys to be returned. | [default to 1000]
 **collection** | **string** | Limit the results to keys that belong to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work. | 
 **prefix** | **string** | Limit the results to keys that start with a specific prefix. | 
 **signature** | **string** | Signature used for the access. | 

### Return type

[**ListOfKeysResponse**](ListOfKeysResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunKeyValueStorePut

> KeyValueStoreResponse ActorRunKeyValueStorePut(ctx, runId).UpdateStoreRequest(updateStoreRequest).Execute()

Update default store



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	updateStoreRequest := *openapiclient.NewUpdateStoreRequest() // UpdateStoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStorePut(context.Background(), runId).UpdateStoreRequest(updateStoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStorePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunKeyValueStorePut`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultKeyValueStoreAPI.ActorRunKeyValueStorePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStorePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStoreRequest** | [**UpdateStoreRequest**](UpdateStoreRequest.md) |  | 

### Return type

[**KeyValueStoreResponse**](KeyValueStoreResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunKeyValueStoreRecordDelete

> ActorRunKeyValueStoreRecordDelete(ctx, runId, recordKey).Execute()

Delete default store's record



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	recordKey := "someKey" // string | Key of the record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordDelete(context.Background(), runId, recordKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreRecordDeleteRequest struct via the builder pattern


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


## ActorRunKeyValueStoreRecordGet

> map[string]interface{} ActorRunKeyValueStoreRecordGet(ctx, runId, recordKey).Signature(signature).Attachment(attachment).Execute()

Get default store's record



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	recordKey := "someKey" // string | Key of the record.
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)
	attachment := true // bool | If `true` or `1`, the response will be served with `Content-Disposition: attachment` header, causing web browsers to offer downloading HTML records instead of displaying them.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordGet(context.Background(), runId, recordKey).Signature(signature).Attachment(attachment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunKeyValueStoreRecordGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signature** | **string** | Signature used for the access. | 
 **attachment** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response will be served with &#x60;Content-Disposition: attachment&#x60; header, causing web browsers to offer downloading HTML records instead of displaying them.  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunKeyValueStoreRecordPost

> map[string]interface{} ActorRunKeyValueStoreRecordPost(ctx, runId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()

Store record in default store (POST)



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordPost(context.Background(), runId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunKeyValueStoreRecordPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreRecordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **contentEncoding** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunKeyValueStoreRecordPut

> map[string]interface{} ActorRunKeyValueStoreRecordPut(ctx, runId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()

Store record in default store



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordPut(context.Background(), runId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunKeyValueStoreRecordPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreRecordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **contentEncoding** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunKeyValueStoreRecordsGet

> *os.File ActorRunKeyValueStoreRecordsGet(ctx, runId).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Download default store's records



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
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	collection := "my-collection" // string | If specified, only records belonging to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work.  (optional)
	prefix := "my-prefix/" // string | If specified, only records whose key starts with the given prefix are included in the archive.  (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordsGet(context.Background(), runId).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunKeyValueStoreRecordsGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultKeyValueStoreAPI.ActorRunKeyValueStoreRecordsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunKeyValueStoreRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collection** | **string** | If specified, only records belonging to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work.  | 
 **prefix** | **string** | If specified, only records whose key starts with the given prefix are included in the archive.  | 
 **signature** | **string** | Signature used for the access. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

