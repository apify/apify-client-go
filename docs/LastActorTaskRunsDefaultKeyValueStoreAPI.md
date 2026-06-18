# \LastActorTaskRunsDefaultKeyValueStoreAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorTaskRunsLastKeyValueStoreDelete**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreDelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store | Delete last task run&#39;s default store
[**ActorTaskRunsLastKeyValueStoreGet**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store | Get last task run&#39;s default store
[**ActorTaskRunsLastKeyValueStoreKeysGet**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreKeysGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/keys | Get last task run&#39;s default store&#39;s list of keys
[**ActorTaskRunsLastKeyValueStorePut**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStorePut) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store | Update last task run&#39;s default store
[**ActorTaskRunsLastKeyValueStoreRecordDelete**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreRecordDelete) | **Delete** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Delete last task run&#39;s default store&#39;s record
[**ActorTaskRunsLastKeyValueStoreRecordGet**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreRecordGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Get last task run&#39;s default store&#39;s record
[**ActorTaskRunsLastKeyValueStoreRecordPost**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreRecordPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Store record in last task run&#39;s default store (POST)
[**ActorTaskRunsLastKeyValueStoreRecordPut**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreRecordPut) | **Put** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records/{recordKey} | Store record in last task run&#39;s default store
[**ActorTaskRunsLastKeyValueStoreRecordsGet**](LastActorTaskRunsDefaultKeyValueStoreAPI.md#ActorTaskRunsLastKeyValueStoreRecordsGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last/key-value-store/records | Download last task run&#39;s default store&#39;s records



## ActorTaskRunsLastKeyValueStoreDelete

> ActorTaskRunsLastKeyValueStoreDelete(ctx, actorTaskId).Status(status).Execute()

Delete last task run's default store



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreDelete(context.Background(), actorTaskId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreDeleteRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStoreGet

> KeyValueStoreResponse ActorTaskRunsLastKeyValueStoreGet(ctx, actorTaskId).Status(status).Execute()

Get last task run's default store



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreGet(context.Background(), actorTaskId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastKeyValueStoreGet`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreGetRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStoreKeysGet

> ListOfKeysResponse ActorTaskRunsLastKeyValueStoreKeysGet(ctx, actorTaskId).Status(status).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Get last task run's default store's list of keys



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	exclusiveStartKey := "Ihnsp8YrvJ8102Kj" // string | All keys up to this one (including) are skipped from the result. (optional)
	limit := float32(100) // float32 | Number of keys to be returned. (optional) (default to 1000)
	collection := "postImages" // string | Limit the results to keys that belong to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work. (optional)
	prefix := "post-images-" // string | Limit the results to keys that start with a specific prefix. (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreKeysGet(context.Background(), actorTaskId).Status(status).ExclusiveStartKey(exclusiveStartKey).Limit(limit).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastKeyValueStoreKeysGet`: ListOfKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreKeysGetRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStorePut

> KeyValueStoreResponse ActorTaskRunsLastKeyValueStorePut(ctx, actorTaskId).UpdateStoreRequest(updateStoreRequest).Status(status).Execute()

Update last task run's default store



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	updateStoreRequest := *openapiclient.NewUpdateStoreRequest() // UpdateStoreRequest | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStorePut(context.Background(), actorTaskId).UpdateStoreRequest(updateStoreRequest).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStorePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastKeyValueStorePut`: KeyValueStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStorePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStorePutRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStoreRecordDelete

> ActorTaskRunsLastKeyValueStoreRecordDelete(ctx, actorTaskId, recordKey).Status(status).Execute()

Delete last task run's default store's record



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	recordKey := "someKey" // string | Key of the record.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordDelete(context.Background(), actorTaskId, recordKey).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreRecordDeleteRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStoreRecordGet

> map[string]interface{} ActorTaskRunsLastKeyValueStoreRecordGet(ctx, actorTaskId, recordKey).Status(status).Signature(signature).Attachment(attachment).Execute()

Get last task run's default store's record



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	recordKey := "someKey" // string | Key of the record.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)
	attachment := true // bool | If `true` or `1`, the response will be served with `Content-Disposition: attachment` header, causing web browsers to offer downloading HTML records instead of displaying them.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordGet(context.Background(), actorTaskId, recordKey).Status(status).Signature(signature).Attachment(attachment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastKeyValueStoreRecordGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreRecordGetRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStoreRecordPost

> map[string]interface{} ActorTaskRunsLastKeyValueStoreRecordPost(ctx, actorTaskId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()

Store record in last task run's default store (POST)



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordPost(context.Background(), actorTaskId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastKeyValueStoreRecordPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreRecordPostRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStoreRecordPut

> map[string]interface{} ActorTaskRunsLastKeyValueStoreRecordPut(ctx, actorTaskId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()

Store record in last task run's default store



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	recordKey := "someKey" // string | Key of the record.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	contentEncoding := "contentEncoding_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordPut(context.Background(), actorTaskId, recordKey).RequestBody(requestBody).Status(status).ContentEncoding(contentEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastKeyValueStoreRecordPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 
**recordKey** | **string** | Key of the record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreRecordPutRequest struct via the builder pattern


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


## ActorTaskRunsLastKeyValueStoreRecordsGet

> *os.File ActorTaskRunsLastKeyValueStoreRecordsGet(ctx, actorTaskId).Status(status).Collection(collection).Prefix(prefix).Signature(signature).Execute()

Download last task run's default store's records



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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	collection := "my-collection" // string | If specified, only records belonging to a specific collection from the key-value store schema. The key-value store need to have a schema defined for this parameter to work.  (optional)
	prefix := "my-prefix/" // string | If specified, only records whose key starts with the given prefix are included in the archive.  (optional)
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordsGet(context.Background(), actorTaskId).Status(status).Collection(collection).Prefix(prefix).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastKeyValueStoreRecordsGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LastActorTaskRunsDefaultKeyValueStoreAPI.ActorTaskRunsLastKeyValueStoreRecordsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastKeyValueStoreRecordsGetRequest struct via the builder pattern


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

