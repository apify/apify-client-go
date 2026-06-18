# \UsersAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGet**](UsersAPI.md#UserGet) | **Get** /v2/users/{userId} | Get public user data
[**UsersMeGet**](UsersAPI.md#UsersMeGet) | **Get** /v2/users/me | Get private user data
[**UsersMeLimitsGet**](UsersAPI.md#UsersMeLimitsGet) | **Get** /v2/users/me/limits | Get limits
[**UsersMeLimitsPut**](UsersAPI.md#UsersMeLimitsPut) | **Put** /v2/users/me/limits | Update limits
[**UsersMeUsageMonthlyGet**](UsersAPI.md#UsersMeUsageMonthlyGet) | **Get** /v2/users/me/usage/monthly | Get monthly usage



## UserGet

> PublicUserDataResponse UserGet(ctx, userId).Execute()

Get public user data



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
	userId := "HGzIk8z78YcAPEB" // string | User ID or username.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserGet`: PublicUserDataResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID or username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicUserDataResponse**](PublicUserDataResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeGet

> PrivateUserDataResponse UsersMeGet(ctx).Execute()

Get private user data



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersMeGet`: PrivateUserDataResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMeGetRequest struct via the builder pattern


### Return type

[**PrivateUserDataResponse**](PrivateUserDataResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeLimitsGet

> LimitsResponse UsersMeLimitsGet(ctx).Execute()

Get limits



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersMeLimitsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersMeLimitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersMeLimitsGet`: LimitsResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersMeLimitsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMeLimitsGetRequest struct via the builder pattern


### Return type

[**LimitsResponse**](LimitsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMeLimitsPut

> map[string]interface{} UsersMeLimitsPut(ctx).UpdateLimitsRequest(updateLimitsRequest).Execute()

Update limits



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
	updateLimitsRequest := *openapiclient.NewUpdateLimitsRequest() // UpdateLimitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersMeLimitsPut(context.Background()).UpdateLimitsRequest(updateLimitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersMeLimitsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersMeLimitsPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersMeLimitsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersMeLimitsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLimitsRequest** | [**UpdateLimitsRequest**](UpdateLimitsRequest.md) |  | 

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


## UsersMeUsageMonthlyGet

> MonthlyUsageResponse UsersMeUsageMonthlyGet(ctx).Date(date).Execute()

Get monthly usage



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
	date := "2020-06-14" // string | Date in the YYYY-MM-DD format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersMeUsageMonthlyGet(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersMeUsageMonthlyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersMeUsageMonthlyGet`: MonthlyUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersMeUsageMonthlyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersMeUsageMonthlyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | Date in the YYYY-MM-DD format. | 

### Return type

[**MonthlyUsageResponse**](MonthlyUsageResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

