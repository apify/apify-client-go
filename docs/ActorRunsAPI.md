# \ActorRunsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorRunAbortPost**](ActorRunsAPI.md#ActorRunAbortPost) | **Post** /v2/actor-runs/{runId}/abort | Abort run
[**ActorRunDelete**](ActorRunsAPI.md#ActorRunDelete) | **Delete** /v2/actor-runs/{runId} | Delete run
[**ActorRunGet**](ActorRunsAPI.md#ActorRunGet) | **Get** /v2/actor-runs/{runId} | Get run
[**ActorRunLogGet**](ActorRunsAPI.md#ActorRunLogGet) | **Get** /v2/actor-runs/{runId}/log | Get run&#39;s log
[**ActorRunMetamorphPost**](ActorRunsAPI.md#ActorRunMetamorphPost) | **Post** /v2/actor-runs/{runId}/metamorph | Metamorph run
[**ActorRunPut**](ActorRunsAPI.md#ActorRunPut) | **Put** /v2/actor-runs/{runId} | Update run
[**ActorRunRebootPost**](ActorRunsAPI.md#ActorRunRebootPost) | **Post** /v2/actor-runs/{runId}/reboot | Reboot run
[**ActorRunsGet**](ActorRunsAPI.md#ActorRunsGet) | **Get** /v2/actor-runs | Get user runs list
[**PostChargeRun**](ActorRunsAPI.md#PostChargeRun) | **Post** /v2/actor-runs/{runId}/charge | Charge events in run
[**PostResurrectRun**](ActorRunsAPI.md#PostResurrectRun) | **Post** /v2/actor-runs/{runId}/resurrect | Resurrect run



## ActorRunAbortPost

> RunResponse ActorRunAbortPost(ctx, runId).Gracefully(gracefully).Execute()

Abort run



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
	gracefully := true // bool | If true passed, the Actor run will abort gracefully. It will send `aborting` and `persistState` event into run and force-stop the run after 30 seconds. It is helpful in cases where you plan to resurrect the run later.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.ActorRunAbortPost(context.Background(), runId).Gracefully(gracefully).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunAbortPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.ActorRunAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gracefully** | **bool** | If true passed, the Actor run will abort gracefully. It will send &#x60;aborting&#x60; and &#x60;persistState&#x60; event into run and force-stop the run after 30 seconds. It is helpful in cases where you plan to resurrect the run later.  | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunDelete

> ActorRunDelete(ctx, runId).Execute()

Delete run



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
	r, err := apiClient.ActorRunsAPI.ActorRunDelete(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActorRunDeleteRequest struct via the builder pattern


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


## ActorRunGet

> RunResponse ActorRunGet(ctx, runId).WaitForFinish(waitForFinish).Execute()

Get run



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
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the run to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the run finishes in time then the returned run object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.ActorRunGet(context.Background(), runId).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunGet`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.ActorRunGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **waitForFinish** | **float64** | The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;).  | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunLogGet

> string ActorRunLogGet(ctx, runId).Stream(stream).Download(download).Raw(raw).Execute()

Get run's log



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
	stream := false // bool | If `true` or `1` then the logs will be streamed as long as the run or build is running.  (optional)
	download := false // bool | If `true` or `1` then the web browser will download the log file rather than open it in a tab.  (optional)
	raw := false // bool | If `true` or `1`, the logs will be kept verbatim. By default, the API removes ANSI escape codes from the logs, keeping only printable characters.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.ActorRunLogGet(context.Background(), runId).Stream(stream).Download(download).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunLogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunLogGet`: string
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.ActorRunLogGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunLogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stream** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the logs will be streamed as long as the run or build is running.  | 
 **download** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the web browser will download the log file rather than open it in a tab.  | 
 **raw** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60;, the logs will be kept verbatim. By default, the API removes ANSI escape codes from the logs, keeping only printable characters.  | 

### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunMetamorphPost

> RunResponse ActorRunMetamorphPost(ctx, runId).TargetActorId(targetActorId).Build(build).Execute()

Metamorph run



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
	targetActorId := "HDSasDasz78YcAPEB" // string | ID of a target Actor that the run should be transformed into.
	build := "beta" // string | Optional build of the target Actor.  It can be either a build tag or build number. By default, the run uses the build specified in the default run configuration for the target Actor (typically `latest`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.ActorRunMetamorphPost(context.Background(), runId).TargetActorId(targetActorId).Build(build).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunMetamorphPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunMetamorphPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.ActorRunMetamorphPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunMetamorphPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetActorId** | **string** | ID of a target Actor that the run should be transformed into. | 
 **build** | **string** | Optional build of the target Actor.  It can be either a build tag or build number. By default, the run uses the build specified in the default run configuration for the target Actor (typically &#x60;latest&#x60;).  | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunPut

> RunResponse ActorRunPut(ctx, runId).UpdateRunRequest(updateRunRequest).Execute()

Update run



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
	updateRunRequest := *openapiclient.NewUpdateRunRequest() // UpdateRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.ActorRunPut(context.Background(), runId).UpdateRunRequest(updateRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunPut`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.ActorRunPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRunRequest** | [**UpdateRunRequest**](UpdateRunRequest.md) |  | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunRebootPost

> RunResponse ActorRunRebootPost(ctx, runId).Execute()

Reboot run



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
	resp, r, err := apiClient.ActorRunsAPI.ActorRunRebootPost(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunRebootPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunRebootPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.ActorRunRebootPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunRebootPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunsGet

> ListOfRunsResponse ActorRunsGet(ctx).Offset(offset).Limit(limit).Desc(desc).Status(status).StartedAfter(startedAfter).StartedBefore(startedBefore).Execute()

Get user runs list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/apify/apify-client-go"
)

func main() {
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	desc := true // bool | If `true` or `1` then the objects are sorted by the `startedAt` field in descending order. By default, they are sorted in ascending order.  (optional)
	status := []string{"Inner_example"} // []string | Single status or comma-separated list of statuses, see ([available statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)). Used to filter runs by the specified statuses only.  (optional)
	startedAfter := time.Now() // time.Time | Filter runs that started after the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC).  (optional)
	startedBefore := time.Now() // time.Time | Filter runs that started before the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.ActorRunsGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Status(status).StartedAfter(startedAfter).StartedBefore(startedBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.ActorRunsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunsGet`: ListOfRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.ActorRunsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActorRunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;startedAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 
 **status** | **[]string** | Single status or comma-separated list of statuses, see ([available statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)). Used to filter runs by the specified statuses only.  | 
 **startedAfter** | **time.Time** | Filter runs that started after the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC).  | 
 **startedBefore** | **time.Time** | Filter runs that started before the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC).  | 

### Return type

[**ListOfRunsResponse**](ListOfRunsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostChargeRun

> map[string]interface{} PostChargeRun(ctx, runId).ChargeRunRequest(chargeRunRequest).IdempotencyKey(idempotencyKey).Execute()

Charge events in run



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
	chargeRunRequest := *openapiclient.NewChargeRunRequest("EventName_example", int32(123)) // ChargeRunRequest | Define which event, and how many times, you want to charge for.
	idempotencyKey := "2024-12-09T01:23:45.000Z-random-uuid" // string | Always pass a unique idempotency key (any unique string) for each charge to avoid double charging in case of retries or network errors. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.PostChargeRun(context.Background(), runId).ChargeRunRequest(chargeRunRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.PostChargeRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostChargeRun`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.PostChargeRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostChargeRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeRunRequest** | [**ChargeRunRequest**](ChargeRunRequest.md) | Define which event, and how many times, you want to charge for. | 
 **idempotencyKey** | **string** | Always pass a unique idempotency key (any unique string) for each charge to avoid double charging in case of retries or network errors. | 

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


## PostResurrectRun

> RunResponse PostResurrectRun(ctx, runId).Build(build).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Execute()

Resurrect run



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
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run is resurrected with the same build it originally used. Specifically, if a run was first started with the `latest` tag, which resolves to version `0.0.3` at the time, a run resurrected without this parameter will continue running with `0.0.3`, even if `latest` already points to a newer build.  (optional)
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout specified in the run that is being resurrected.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit specified in the run that is being resurrected.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the resurrected run will be restarted if it fails. By default, the resurrected run uses the same setting as before.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorRunsAPI.PostResurrectRun(context.Background(), runId).Build(build).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorRunsAPI.PostResurrectRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResurrectRun`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorRunsAPI.PostResurrectRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostResurrectRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run is resurrected with the same build it originally used. Specifically, if a run was first started with the &#x60;latest&#x60; tag, which resolves to version &#x60;0.0.3&#x60; at the time, a run resurrected without this parameter will continue running with &#x60;0.0.3&#x60;, even if &#x60;latest&#x60; already points to a newer build.  | 
 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout specified in the run that is being resurrected.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit specified in the run that is being resurrected.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **maxTotalChargeUsd** | **float64** | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable.  | 
 **restartOnError** | **bool** | Determines whether the resurrected run will be restarted if it fails. By default, the resurrected run uses the same setting as before.  | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

