# \LastActorRunsDefaultKeyValueStoreAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActRunsLastKeyValueStoreDelete**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreDelete) | **Delete** /v2/actors/{actorId}/runs/last/key-value-store | Delete last run&#39;s default store
[**ActRunsLastKeyValueStoreGet**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreGet) | **Get** /v2/actors/{actorId}/runs/last/key-value-store | Get last run&#39;s default store
[**ActRunsLastKeyValueStoreKeysGet**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreKeysGet) | **Get** /v2/actors/{actorId}/runs/last/key-value-store/keys | Get last run&#39;s default store&#39;s list of keys
[**ActRunsLastKeyValueStorePut**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStorePut) | **Put** /v2/actors/{actorId}/runs/last/key-value-store | Update last run&#39;s default store
[**ActRunsLastKeyValueStoreRecordDelete**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreRecordDelete) | **Delete** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Delete last run&#39;s default store&#39;s record
[**ActRunsLastKeyValueStoreRecordGet**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreRecordGet) | **Get** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Get last run&#39;s default store&#39;s record
[**ActRunsLastKeyValueStoreRecordPost**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreRecordPost) | **Post** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Store record in last run&#39;s default store (POST)
[**ActRunsLastKeyValueStoreRecordPut**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreRecordPut) | **Put** /v2/actors/{actorId}/runs/last/key-value-store/records/{recordKey} | Store record in last run&#39;s default store
[**ActRunsLastKeyValueStoreRecordsGet**](LastActorRunsDefaultKeyValueStoreAPI.md#ActRunsLastKeyValueStoreRecordsGet) | **Get** /v2/actors/{actorId}/runs/last/key-value-store/records | Download last run&#39;s default store&#39;s records



## ActRunsLastKeyValueStoreDelete

> ActRunsLastKeyValueStoreDelete(ctx, actorId).Status(status).Execute()

Delete last run's default store



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
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreDelete(context.Background(), actorId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastKeyValueStoreGet

> KeyValueStoreResponse ActRunsLastKeyValueStoreGet(ctx, actorId).Status(status).Execute()

Get last run's default store



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
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreGet(context.Background(), actorId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastKeyValueStoreGet`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastKeyValueStoreKeysGet

> ListOfKeysResponse ActRunsLastKeyValueStoreKeysGet(ctx, actorId).Status(status).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Get last run's default store's list of keys



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
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	exclusiveStartKey := "Ihnsp8YrvJ8102Kj" // string | All keys up to this one (including) are skipped from the result. (optional)
	limit := float32(100) // float32 | Number of keys to be returned. (optional) (default to 1000)
	collection := "postImages" // string | Limit the results to keys that belong to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work. (optional)
	prefix := "post-images-" // string | Limit the results to keys that start with a specific prefix. (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreKeysGet(context.Background(), actorId).Status(status).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastKeyValueStoreKeysGet`: ListOfKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastKeyValueStorePut

> KeyValueStoreResponse ActRunsLastKeyValueStorePut(ctx, actorId).UpdateStoreRequest(updateStoreRequest).Status(status).Execute()

Update last run's default store



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
	updateStoreRequest := *openapiclient.NewUpdateStoreRequest() // UpdateStoreRequest | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStorePut(context.Background(), actorId).UpdateStoreRequest(updateStoreRequest).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStorePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastKeyValueStorePut`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStorePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStorePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStoreRequest** | [**UpdateStoreRequest**](UpdateStoreRequest.md) |  | 
 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastKeyValueStoreRecordDelete

> ActRunsLastKeyValueStoreRecordDelete(ctx, actorId, recordKey).Status(status).Execute()

Delete last run's default store's record



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
	recordKey := "someKey" // string | Key of the record.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordDelete(context.Background(), actorId, recordKey).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreRecordDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter for the run status. | 

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


## ActRunsLastKeyValueStoreRecordGet

> map[string]interface{} ActRunsLastKeyValueStoreRecordGet(ctx, actorId, recordKey).Status(status).Signature(signature).Attachment(attachment).Execute()

Get last run's default store's record



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
	recordKey := "someKey" // string | Key of the record.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)
	attachment := true // bool | If `true` or `1`, the response will be served with `Content-Disposition: attachment` header, causing web browsers to offer downloading HTML records instead of displaying them.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordGet(context.Background(), actorId, recordKey).Status(status).Signature(signature).Attachment(attachment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastKeyValueStoreRecordGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastKeyValueStoreRecordPost

> map[string]interface{} ActRunsLastKeyValueStoreRecordPost(ctx, actorId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()

Store record in last run's default store (POST)



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
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordPost(context.Background(), actorId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastKeyValueStoreRecordPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreRecordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastKeyValueStoreRecordPut

> map[string]interface{} ActRunsLastKeyValueStoreRecordPut(ctx, actorId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()

Store record in last run's default store



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
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordPut(context.Background(), actorId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastKeyValueStoreRecordPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreRecordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **status** | **string** | Filter for the run status. | 
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


## ActRunsLastKeyValueStoreRecordsGet

> *os.File ActRunsLastKeyValueStoreRecordsGet(ctx, actorId).Status(status).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Download last run's default store's records



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
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	collection := "my-collection" // string | If specified, only records belonging to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work.  (optional)
	prefix := "my-prefix/" // string | If specified, only records whose key starts with the given prefix are included in the archive.  (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordsGet(context.Background(), actorId).Status(status).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastKeyValueStoreRecordsGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LastActorRunsDefaultKeyValueStoreAPI.ActRunsLastKeyValueStoreRecordsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastKeyValueStoreRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
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

