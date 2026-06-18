# \StorageKeyValueStoresAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyValueStoreDelete**](StorageKeyValueStoresAPI.md#KeyValueStoreDelete) | **Delete** /v2/key-value-stores/{storeId} | Delete store
[**KeyValueStoreGet**](StorageKeyValueStoresAPI.md#KeyValueStoreGet) | **Get** /v2/key-value-stores/{storeId} | Get store
[**KeyValueStoreKeysGet**](StorageKeyValueStoresAPI.md#KeyValueStoreKeysGet) | **Get** /v2/key-value-stores/{storeId}/keys | Get list of keys
[**KeyValueStorePut**](StorageKeyValueStoresAPI.md#KeyValueStorePut) | **Put** /v2/key-value-stores/{storeId} | Update store
[**KeyValueStoreRecordDelete**](StorageKeyValueStoresAPI.md#KeyValueStoreRecordDelete) | **Delete** /v2/key-value-stores/{storeId}/records/{recordKey} | Delete record
[**KeyValueStoreRecordGet**](StorageKeyValueStoresAPI.md#KeyValueStoreRecordGet) | **Get** /v2/key-value-stores/{storeId}/records/{recordKey} | Get record
[**KeyValueStoreRecordHead**](StorageKeyValueStoresAPI.md#KeyValueStoreRecordHead) | **Head** /v2/key-value-stores/{storeId}/records/{recordKey} | Check if a record exists
[**KeyValueStoreRecordPost**](StorageKeyValueStoresAPI.md#KeyValueStoreRecordPost) | **Post** /v2/key-value-stores/{storeId}/records/{recordKey} | Store record (POST)
[**KeyValueStoreRecordPut**](StorageKeyValueStoresAPI.md#KeyValueStoreRecordPut) | **Put** /v2/key-value-stores/{storeId}/records/{recordKey} | Store record
[**KeyValueStoreRecordsGet**](StorageKeyValueStoresAPI.md#KeyValueStoreRecordsGet) | **Get** /v2/key-value-stores/{storeId}/records | Download records
[**KeyValueStoresGet**](StorageKeyValueStoresAPI.md#KeyValueStoresGet) | **Get** /v2/key-value-stores | Get list of key-value stores
[**KeyValueStoresPost**](StorageKeyValueStoresAPI.md#KeyValueStoresPost) | **Post** /v2/key-value-stores | Create key-value store



## KeyValueStoreDelete

> KeyValueStoreDelete(ctx, storeId).Execute()

Delete store



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreDelete(context.Background(), storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreDeleteRequest struct via the builder pattern


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


## KeyValueStoreGet

> KeyValueStoreResponse KeyValueStoreGet(ctx, storeId).Execute()

Get store



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreGet(context.Background(), storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoreGet`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoreGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreGetRequest struct via the builder pattern


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


## KeyValueStoreKeysGet

> ListOfKeysResponse KeyValueStoreKeysGet(ctx, storeId).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Get list of keys



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	exclusiveStartKey := "Ihnsp8YrvJ8102Kj" // string | All keys up to this one (including) are skipped from the result. (optional)
	limit := float32(100) // float32 | Number of keys to be returned. (optional) (default to 1000)
	collection := "postImages" // string | Limit the results to keys that belong to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work. (optional)
	prefix := "post-images-" // string | Limit the results to keys that start with a specific prefix. (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreKeysGet(context.Background(), storeId).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoreKeysGet`: ListOfKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoreKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreKeysGetRequest struct via the builder pattern


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


## KeyValueStorePut

> KeyValueStoreResponse KeyValueStorePut(ctx, storeId).UpdateStoreRequest(updateStoreRequest).Execute()

Update store



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	updateStoreRequest := *openapiclient.NewUpdateStoreRequest() // UpdateStoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStorePut(context.Background(), storeId).UpdateStoreRequest(updateStoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStorePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStorePut`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStorePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStorePutRequest struct via the builder pattern


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


## KeyValueStoreRecordDelete

> KeyValueStoreRecordDelete(ctx, storeId, recordKey).Execute()

Delete record



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	recordKey := "someKey" // string | Key of the record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreRecordDelete(context.Background(), storeId, recordKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreRecordDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreRecordDeleteRequest struct via the builder pattern


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


## KeyValueStoreRecordGet

> map[string]interface{} KeyValueStoreRecordGet(ctx, storeId, recordKey).Attachment(attachment).Signature(signature).Execute()

Get record



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	recordKey := "someKey" // string | Key of the record.
	attachment := true // bool | If `true` or `1`, the response will be served with `Content-Disposition: attachment` header, causing web browsers to offer downloading HTML records instead of displaying them.  (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreRecordGet(context.Background(), storeId, recordKey).Attachment(attachment).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreRecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoreRecordGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoreRecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachment** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the response will be served with &#x60;Content-Disposition: attachment&#x60; header, causing web browsers to offer downloading HTML records instead of displaying them.  | 
 **signature** | **string** | Signature used for the access. | 

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


## KeyValueStoreRecordHead

> KeyValueStoreRecordHead(ctx, storeId, recordKey).Execute()

Check if a record exists



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	recordKey := "someKey" // string | Key of the record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreRecordHead(context.Background(), storeId, recordKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreRecordHead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreRecordHeadRequest struct via the builder pattern


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


## KeyValueStoreRecordPost

> map[string]interface{} KeyValueStoreRecordPost(ctx, storeId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()

Store record (POST)



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreRecordPost(context.Background(), storeId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreRecordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoreRecordPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoreRecordPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreRecordPostRequest struct via the builder pattern


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


## KeyValueStoreRecordPut

> map[string]interface{} KeyValueStoreRecordPut(ctx, storeId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()

Store record



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreRecordPut(context.Background(), storeId, recordKey).RequestBody(requestBody).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreRecordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoreRecordPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoreRecordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreRecordPutRequest struct via the builder pattern


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


## KeyValueStoreRecordsGet

> *os.File KeyValueStoreRecordsGet(ctx, storeId).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Download records



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
	storeId := "WkzbQMuFYuamGv3YF" // string | Key-value store ID or `username~store-name`.
	collection := "my-collection" // string | If specified, only records belonging to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work.  (optional)
	prefix := "my-prefix/" // string | If specified, only records whose key starts with the given prefix are included in the archive.  (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoreRecordsGet(context.Background(), storeId).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoreRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoreRecordsGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoreRecordsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Key-value store ID or &#x60;username~store-name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoreRecordsGetRequest struct via the builder pattern


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


## KeyValueStoresGet

> ListOfKeyValueStoresResponse KeyValueStoresGet(ctx).Offset(offset).Limit(limit).Desc(desc).Unnamed(unnamed).Ownership(ownership).Execute()

Get list of key-value stores



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
	unnamed := true // bool | If `true` or `1` then all the storages are returned. By default, only named storages are returned.  (optional)
	ownership := openapiclient.StorageOwnership("ownedByMe") // StorageOwnership | Filter by ownership. If this parameter is omitted, all accessible key-value stores are returned.  - `ownedByMe`: Return only key-value stores owned by the user. - `sharedWithMe`: Return only key-value stores shared with the user by other users.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoresGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Unnamed(unnamed).Ownership(ownership).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoresGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoresGet`: ListOfKeyValueStoresResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoresGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoresGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 
 **unnamed** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then all the storages are returned. By default, only named storages are returned.  | 
 **ownership** | [**StorageOwnership**](StorageOwnership.md) | Filter by ownership. If this parameter is omitted, all accessible key-value stores are returned.  - &#x60;ownedByMe&#x60;: Return only key-value stores owned by the user. - &#x60;sharedWithMe&#x60;: Return only key-value stores shared with the user by other users.  | 

### Return type

[**ListOfKeyValueStoresResponse**](ListOfKeyValueStoresResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyValueStoresPost

> KeyValueStoreResponse KeyValueStoresPost(ctx).Name(name).Execute()

Create key-value store



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
	name := "eshop-values" // string | Custom unique name to easily identify the store in the future. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageKeyValueStoresAPI.KeyValueStoresPost(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageKeyValueStoresAPI.KeyValueStoresPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyValueStoresPost`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageKeyValueStoresAPI.KeyValueStoresPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyValueStoresPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Custom unique name to easily identify the store in the future. | 

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

