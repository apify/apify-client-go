# \ActorsActorVersionsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActVersionDelete**](ActorsActorVersionsAPI.md#ActVersionDelete) | **Delete** /v2/actors/{actorId}/versions/{versionNumber} | Delete version
[**ActVersionEnvVarDelete**](ActorsActorVersionsAPI.md#ActVersionEnvVarDelete) | **Delete** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Delete environment variable
[**ActVersionEnvVarGet**](ActorsActorVersionsAPI.md#ActVersionEnvVarGet) | **Get** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Get environment variable
[**ActVersionEnvVarPost**](ActorsActorVersionsAPI.md#ActVersionEnvVarPost) | **Post** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Update environment variable (POST)
[**ActVersionEnvVarPut**](ActorsActorVersionsAPI.md#ActVersionEnvVarPut) | **Put** /v2/actors/{actorId}/versions/{versionNumber}/env-vars/{envVarName} | Update environment variable
[**ActVersionEnvVarsGet**](ActorsActorVersionsAPI.md#ActVersionEnvVarsGet) | **Get** /v2/actors/{actorId}/versions/{versionNumber}/env-vars | Get list of environment variables
[**ActVersionEnvVarsPost**](ActorsActorVersionsAPI.md#ActVersionEnvVarsPost) | **Post** /v2/actors/{actorId}/versions/{versionNumber}/env-vars | Create environment variable
[**ActVersionGet**](ActorsActorVersionsAPI.md#ActVersionGet) | **Get** /v2/actors/{actorId}/versions/{versionNumber} | Get version
[**ActVersionPost**](ActorsActorVersionsAPI.md#ActVersionPost) | **Post** /v2/actors/{actorId}/versions/{versionNumber} | Update version (POST)
[**ActVersionPut**](ActorsActorVersionsAPI.md#ActVersionPut) | **Put** /v2/actors/{actorId}/versions/{versionNumber} | Update version
[**ActVersionsGet**](ActorsActorVersionsAPI.md#ActVersionsGet) | **Get** /v2/actors/{actorId}/versions | Get list of versions
[**ActVersionsPost**](ActorsActorVersionsAPI.md#ActVersionsPost) | **Post** /v2/actors/{actorId}/versions | Create version



## ActVersionDelete

> map[string]interface{} ActVersionDelete(ctx, actorId, versionNumber).Execute()

Delete version



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
	versionNumber := "0.1" // string | Actor version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionDelete(context.Background(), actorId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionDeleteRequest struct via the builder pattern


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


## ActVersionEnvVarDelete

> map[string]interface{} ActVersionEnvVarDelete(ctx, actorId, versionNumber, envVarName).Execute()

Delete environment variable



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
	versionNumber := "0.1" // string | Actor version.
	envVarName := "MY_ENV_VAR" // string | The name of the environment variable

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionEnvVarDelete(context.Background(), actorId, versionNumber, envVarName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionEnvVarDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionEnvVarDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionEnvVarDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 
**envVarName** | **string** | The name of the environment variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionEnvVarDeleteRequest struct via the builder pattern


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


## ActVersionEnvVarGet

> EnvVarResponse ActVersionEnvVarGet(ctx, actorId, versionNumber, envVarName).Execute()

Get environment variable



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
	versionNumber := "0.1" // string | Actor version.
	envVarName := "MY_ENV_VAR" // string | The name of the environment variable

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionEnvVarGet(context.Background(), actorId, versionNumber, envVarName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionEnvVarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionEnvVarGet`: EnvVarResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionEnvVarGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 
**envVarName** | **string** | The name of the environment variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionEnvVarGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnvVarResponse**](EnvVarResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionEnvVarPost

> EnvVarResponse ActVersionEnvVarPost(ctx, actorId, versionNumber, envVarName).EnvVarRequest(envVarRequest).Execute()

Update environment variable (POST)



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
	versionNumber := "0.1" // string | Actor version.
	envVarName := "MY_ENV_VAR" // string | The name of the environment variable
	envVarRequest := *openapiclient.NewEnvVarRequest("Name_example") // EnvVarRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionEnvVarPost(context.Background(), actorId, versionNumber, envVarName).EnvVarRequest(envVarRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionEnvVarPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionEnvVarPost`: EnvVarResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionEnvVarPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 
**envVarName** | **string** | The name of the environment variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionEnvVarPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envVarRequest** | [**EnvVarRequest**](EnvVarRequest.md) |  | 

### Return type

[**EnvVarResponse**](EnvVarResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionEnvVarPut

> EnvVarResponse ActVersionEnvVarPut(ctx, actorId, versionNumber, envVarName).EnvVarRequest(envVarRequest).Execute()

Update environment variable



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
	versionNumber := "0.1" // string | Actor version.
	envVarName := "MY_ENV_VAR" // string | The name of the environment variable
	envVarRequest := *openapiclient.NewEnvVarRequest("Name_example") // EnvVarRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionEnvVarPut(context.Background(), actorId, versionNumber, envVarName).EnvVarRequest(envVarRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionEnvVarPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionEnvVarPut`: EnvVarResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionEnvVarPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 
**envVarName** | **string** | The name of the environment variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionEnvVarPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envVarRequest** | [**EnvVarRequest**](EnvVarRequest.md) |  | 

### Return type

[**EnvVarResponse**](EnvVarResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionEnvVarsGet

> ListOfEnvVarsResponse ActVersionEnvVarsGet(ctx, actorId, versionNumber).Execute()

Get list of environment variables



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
	versionNumber := "0.1" // string | Actor version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionEnvVarsGet(context.Background(), actorId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionEnvVarsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionEnvVarsGet`: ListOfEnvVarsResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionEnvVarsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionEnvVarsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListOfEnvVarsResponse**](ListOfEnvVarsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionEnvVarsPost

> EnvVarResponse ActVersionEnvVarsPost(ctx, actorId, versionNumber).EnvVarRequest(envVarRequest).Execute()

Create environment variable



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
	versionNumber := "0.1" // string | Actor version.
	envVarRequest := *openapiclient.NewEnvVarRequest("Name_example") // EnvVarRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionEnvVarsPost(context.Background(), actorId, versionNumber).EnvVarRequest(envVarRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionEnvVarsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionEnvVarsPost`: EnvVarResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionEnvVarsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionEnvVarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envVarRequest** | [**EnvVarRequest**](EnvVarRequest.md) |  | 

### Return type

[**EnvVarResponse**](EnvVarResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionGet

> VersionResponse ActVersionGet(ctx, actorId, versionNumber).Execute()

Get version



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
	versionNumber := "0.1" // string | Actor version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionGet(context.Background(), actorId, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionGet`: VersionResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionPost

> VersionResponse ActVersionPost(ctx, actorId, versionNumber).CreateOrUpdateVersionRequest(createOrUpdateVersionRequest).Execute()

Update version (POST)



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
	versionNumber := "0.1" // string | Actor version.
	createOrUpdateVersionRequest := *openapiclient.NewCreateOrUpdateVersionRequest() // CreateOrUpdateVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionPost(context.Background(), actorId, versionNumber).CreateOrUpdateVersionRequest(createOrUpdateVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionPost`: VersionResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOrUpdateVersionRequest** | [**CreateOrUpdateVersionRequest**](CreateOrUpdateVersionRequest.md) |  | 

### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionPut

> VersionResponse ActVersionPut(ctx, actorId, versionNumber).CreateOrUpdateVersionRequest(createOrUpdateVersionRequest).Execute()

Update version



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
	versionNumber := "0.1" // string | Actor version.
	createOrUpdateVersionRequest := *openapiclient.NewCreateOrUpdateVersionRequest() // CreateOrUpdateVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionPut(context.Background(), actorId, versionNumber).CreateOrUpdateVersionRequest(createOrUpdateVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionPut`: VersionResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**versionNumber** | **string** | Actor version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOrUpdateVersionRequest** | [**CreateOrUpdateVersionRequest**](CreateOrUpdateVersionRequest.md) |  | 

### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionsGet

> ListOfVersionsResponse ActVersionsGet(ctx, actorId).Execute()

Get list of versions



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
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionsGet(context.Background(), actorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionsGet`: ListOfVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListOfVersionsResponse**](ListOfVersionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActVersionsPost

> VersionResponse ActVersionsPost(ctx, actorId).CreateOrUpdateVersionRequest(createOrUpdateVersionRequest).Execute()

Create version



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
	createOrUpdateVersionRequest := *openapiclient.NewCreateOrUpdateVersionRequest() // CreateOrUpdateVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorVersionsAPI.ActVersionsPost(context.Background(), actorId).CreateOrUpdateVersionRequest(createOrUpdateVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorVersionsAPI.ActVersionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActVersionsPost`: VersionResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorVersionsAPI.ActVersionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActVersionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrUpdateVersionRequest** | [**CreateOrUpdateVersionRequest**](CreateOrUpdateVersionRequest.md) |  | 

### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

