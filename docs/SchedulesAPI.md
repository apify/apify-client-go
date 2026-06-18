# \SchedulesAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduleDelete**](SchedulesAPI.md#ScheduleDelete) | **Delete** /v2/schedules/{scheduleId} | Delete schedule
[**ScheduleGet**](SchedulesAPI.md#ScheduleGet) | **Get** /v2/schedules/{scheduleId} | Get schedule
[**ScheduleLogGet**](SchedulesAPI.md#ScheduleLogGet) | **Get** /v2/schedules/{scheduleId}/log | Get schedule log
[**SchedulePut**](SchedulesAPI.md#SchedulePut) | **Put** /v2/schedules/{scheduleId} | Update schedule
[**SchedulesGet**](SchedulesAPI.md#SchedulesGet) | **Get** /v2/schedules | Get list of schedules
[**SchedulesPost**](SchedulesAPI.md#SchedulesPost) | **Post** /v2/schedules | Create schedule



## ScheduleDelete

> map[string]interface{} ScheduleDelete(ctx, scheduleId).Execute()

Delete schedule



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
	scheduleId := "asdLZtadYvn4mBZmm" // string | Schedule ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulesAPI.ScheduleDelete(context.Background(), scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.ScheduleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.ScheduleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | Schedule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleDeleteRequest struct via the builder pattern


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


## ScheduleGet

> ScheduleResponse ScheduleGet(ctx, scheduleId).Execute()

Get schedule



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
	scheduleId := "asdLZtadYvn4mBZmm" // string | Schedule ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulesAPI.ScheduleGet(context.Background(), scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.ScheduleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleGet`: ScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.ScheduleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | Schedule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduleResponse**](ScheduleResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleLogGet

> ScheduleLogResponse ScheduleLogGet(ctx, scheduleId).Execute()

Get schedule log



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
	scheduleId := "asdLZtadYvn4mBZmm" // string | Schedule ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulesAPI.ScheduleLogGet(context.Background(), scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.ScheduleLogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleLogGet`: ScheduleLogResponse
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.ScheduleLogGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | Schedule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleLogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduleLogResponse**](ScheduleLogResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulePut

> ScheduleResponse SchedulePut(ctx, scheduleId).ScheduleCreate(scheduleCreate).Execute()

Update schedule



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
	scheduleId := "asdLZtadYvn4mBZmm" // string | Schedule ID.
	scheduleCreate := *openapiclient.NewScheduleCreate() // ScheduleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulesAPI.SchedulePut(context.Background(), scheduleId).ScheduleCreate(scheduleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.SchedulePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulePut`: ScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.SchedulePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | Schedule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchedulePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleCreate** | [**ScheduleCreate**](ScheduleCreate.md) |  | 

### Return type

[**ScheduleResponse**](ScheduleResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulesGet

> ListOfSchedulesResponse SchedulesGet(ctx).Offset(offset).Limit(limit).Desc(desc).Execute()

Get list of schedules



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
	resp, r, err := apiClient.SchedulesAPI.SchedulesGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.SchedulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulesGet`: ListOfSchedulesResponse
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.SchedulesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 

### Return type

[**ListOfSchedulesResponse**](ListOfSchedulesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulesPost

> ScheduleResponse SchedulesPost(ctx).ScheduleCreate(scheduleCreate).Execute()

Create schedule



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
	scheduleCreate := *openapiclient.NewScheduleCreate() // ScheduleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulesAPI.SchedulesPost(context.Background()).ScheduleCreate(scheduleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.SchedulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulesPost`: ScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.SchedulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduleCreate** | [**ScheduleCreate**](ScheduleCreate.md) |  | 

### Return type

[**ScheduleResponse**](ScheduleResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

