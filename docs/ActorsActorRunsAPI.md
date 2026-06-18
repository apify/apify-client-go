# \ActorsActorRunsAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActRunAbortPost**](ActorsActorRunsAPI.md#ActRunAbortPost) | **Post** /v2/actors/{actorId}/runs/{runId}/abort | Abort run
[**ActRunGet**](ActorsActorRunsAPI.md#ActRunGet) | **Get** /v2/actors/{actorId}/runs/{runId} | Get run
[**ActRunMetamorphPost**](ActorsActorRunsAPI.md#ActRunMetamorphPost) | **Post** /v2/actors/{actorId}/runs/{runId}/metamorph | Metamorph run
[**ActRunResurrectPost**](ActorsActorRunsAPI.md#ActRunResurrectPost) | **Post** /v2/actors/{actorId}/runs/{runId}/resurrect | Resurrect run
[**ActRunSyncGet**](ActorsActorRunsAPI.md#ActRunSyncGet) | **Get** /v2/actors/{actorId}/run-sync | Run Actor synchronously without input
[**ActRunSyncGetDatasetItemsGet**](ActorsActorRunsAPI.md#ActRunSyncGetDatasetItemsGet) | **Get** /v2/actors/{actorId}/run-sync-get-dataset-items | Run Actor synchronously without input and get dataset items
[**ActRunSyncGetDatasetItemsPost**](ActorsActorRunsAPI.md#ActRunSyncGetDatasetItemsPost) | **Post** /v2/actors/{actorId}/run-sync-get-dataset-items | Run Actor synchronously and get dataset items
[**ActRunSyncPost**](ActorsActorRunsAPI.md#ActRunSyncPost) | **Post** /v2/actors/{actorId}/run-sync | Run Actor synchronously and return output
[**ActRunsGet**](ActorsActorRunsAPI.md#ActRunsGet) | **Get** /v2/actors/{actorId}/runs | Get list of runs
[**ActRunsLastGet**](ActorsActorRunsAPI.md#ActRunsLastGet) | **Get** /v2/actors/{actorId}/runs/last | Get last run
[**ActRunsPost**](ActorsActorRunsAPI.md#ActRunsPost) | **Post** /v2/actors/{actorId}/runs | Run Actor



## ActRunAbortPost

> RunResponse ActRunAbortPost(ctx, actorId, runId).Gracefully(gracefully).Execute()

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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	gracefully := true // bool | If true passed, the Actor run will abort gracefully. It will send `aborting` and `persistState` event into run and force-stop the run after 30 seconds. It is helpful in cases where you plan to resurrect the run later.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunAbortPost(context.Background(), actorId, runId).Gracefully(gracefully).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunAbortPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunAbortPostRequest struct via the builder pattern


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


## ActRunGet

> RunResponse ActRunGet(ctx, actorId, runId).WaitForFinish(waitForFinish).Execute()

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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the run to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the run finishes in time then the returned run object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunGet(context.Background(), actorId, runId).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunGet`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunGetRequest struct via the builder pattern


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


## ActRunMetamorphPost

> RunResponse ActRunMetamorphPost(ctx, actorId, runId).TargetActorId(targetActorId).Build(build).Execute()

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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	targetActorId := "HDSasDasz78YcAPEB" // string | ID of a target Actor that the run should be transformed into.
	build := "beta" // string | Optional build of the target Actor.  It can be either a build tag or build number. By default, the run uses the build specified in the default run configuration for the target Actor (typically `latest`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunMetamorphPost(context.Background(), actorId, runId).TargetActorId(targetActorId).Build(build).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunMetamorphPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunMetamorphPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunMetamorphPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunMetamorphPostRequest struct via the builder pattern


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


## ActRunResurrectPost

> RunResponse ActRunResurrectPost(ctx, actorId, runId).Build(build).Timeout(timeout).Memory(memory).RestartOnError(restartOnError).Execute()

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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	runId := "3KH8gEpp4d8uQSe8T" // string | Actor run ID.
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run is resurrected with the same build it originally used. Specifically, if a run was first started with the `latest` tag, which resolves to version `0.0.3` at the time, a run resurrected without this parameter will continue running with `0.0.3`, even if `latest` already points to a newer build.  (optional)
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout specified in the run that is being resurrected.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit specified in the run that is being resurrected.  (optional)
	restartOnError := false // bool | Determines whether the resurrected run will be restarted if it fails. By default, the resurrected run uses the same setting as before.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunResurrectPost(context.Background(), actorId, runId).Build(build).Timeout(timeout).Memory(memory).RestartOnError(restartOnError).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunResurrectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunResurrectPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunResurrectPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunResurrectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run is resurrected with the same build it originally used. Specifically, if a run was first started with the &#x60;latest&#x60; tag, which resolves to version &#x60;0.0.3&#x60; at the time, a run resurrected without this parameter will continue running with &#x60;0.0.3&#x60;, even if &#x60;latest&#x60; already points to a newer build.  | 
 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout specified in the run that is being resurrected.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit specified in the run that is being resurrected.  | 
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


## ActRunSyncGet

> map[string]interface{} ActRunSyncGet(ctx, actorId).OutputRecordKey(outputRecordKey).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Execute()

Run Actor synchronously without input



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
	outputRecordKey := "OUTPUT" // string | Key of the record from run's default key-value store to be returned in the response. By default, it is `OUTPUT`.  (optional)
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the run will be restarted if it fails.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunSyncGet(context.Background(), actorId).OutputRecordKey(outputRecordKey).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunSyncGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunSyncGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunSyncGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunSyncGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **outputRecordKey** | **string** | Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;.  | 
 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **maxTotalChargeUsd** | **float64** | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable.  | 
 **restartOnError** | **bool** | Determines whether the run will be restarted if it fails.  | 
 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;).  | 
 **webhooks** | **string** | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  | 

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


## ActRunSyncGetDatasetItemsGet

> []map[string]interface{} ActRunSyncGetDatasetItemsGet(ctx, actorId).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()

Run Actor synchronously without input and get dataset items



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
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the run will be restarted if it fails.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)
	format := "json" // string | Format of the results, possible values are: `json`, `jsonl`, `csv`, `html`, `xlsx`, `xml` and `rss`. The default value is `json`.  (optional)
	clean := false // bool | If `true` or `1` then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The `clean` parameter is just a shortcut for `skipHidden=true` and `skipEmpty=true` parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the `limit` value.  (optional)
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1.2) // float64 | Maximum number of items to return. By default there is no limit. (optional)
	fields := "myValue,myOtherValue" // string | A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the `fields` query parameter. You can use this feature to effectively fix the output format.  (optional)
	outputFields := "title,link" // string | A comma-separated list of output field names that positionally rename the fields specified in the `fields` parameter. For example, `?fields=headline,url&outputFields=title,link` renames `headline` to `title` and `url` to `link` in the output. The number of names in `outputFields` must match the number of names in `fields`. Requires the `fields` parameter to be specified as well.  (optional)
	omit := "myValue,myOtherValue" // string | A comma-separated list of fields which should be omitted from the items. (optional)
	unwind := "myValue,myOtherValue" // string | A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the `desc` parameter.  (optional)
	flatten := "myValue" // string | A comma-separated list of fields which should transform nested objects into flat structures.  For example, with `flatten=\"foo\"` the object `{\"foo\":{\"bar\": \"hello\"}}` is turned into `{\"foo.bar\": \"hello\"}`.  The original object with properties is replaced with the flattened object.  (optional)
	desc := true // bool | By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to `true` or `1`.  (optional)
	attachment := true // bool | If `true` or `1` then the response will define the `Content-Disposition: attachment` header, forcing a web browser to download the file rather than to display it. By default this header is not present.  (optional)
	delimiter := ";" // string | A delimiter character for CSV files, only used if `format=csv`. You might need to URL-encode the character (e.g. use `%09` for tab or `%3B` for semicolon). The default delimiter is a simple comma (`,`).  (optional)
	bom := false // bool | All text responses are encoded in UTF-8 encoding. By default, the `format=csv` files are prefixed with the UTF-8 Byte Order Mark (BOM), while `json`, `jsonl`, `xml`, `html` and `rss` files are not.  If you want to override this default behavior, specify `bom=1` query parameter to include the BOM or `bom=0` to skip it.  (optional)
	xmlRoot := "items" // string | Overrides default root element name of `xml` output. By default the root element is `items`.  (optional)
	xmlRow := "item" // string | Overrides default element name that wraps each page or page function result object in `xml` output. By default the element name is `item`.  (optional)
	skipHeaderRow := true // bool | If `true` or `1` then header row in the `csv` format is skipped. (optional)
	skipHidden := false // bool | If `true` or `1` then hidden fields are skipped from the output, i.e. fields starting with the `#` character.  (optional)
	skipEmpty := false // bool | If `true` or `1` then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value.  (optional)
	simplified := false // bool | If `true` or `1` then, the endpoint applies the `fields=url,pageFunctionResult,errorInfo` and `unwind=pageFunctionResult` query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it's not recommended to use it in new integrations.  (optional)
	view := "overview" // string | Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema).  (optional)
	skipFailedPages := false // bool | If `true` or `1` then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it's not recommended to use it in new integrations.  (optional)
	feedTitle := "Latest posts from r/pasta" // string | Overrides the auto-generated RSS channel `<title>` element. Only used when `format=rss`. If not provided, the title defaults to `Dataset <label>`.  (optional)
	feedDescription := "Scraped forum posts" // string | Overrides the auto-generated RSS channel `<description>` element. Only used when `format=rss`. If not provided, the description defaults to `Items in dataset with id \"<datasetId>\".`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunSyncGetDatasetItemsGet(context.Background(), actorId).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunSyncGetDatasetItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunSyncGetDatasetItemsGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunSyncGetDatasetItemsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunSyncGetDatasetItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **maxTotalChargeUsd** | **float64** | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable.  | 
 **restartOnError** | **bool** | Determines whether the run will be restarted if it fails.  | 
 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;).  | 
 **webhooks** | **string** | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  | 
 **format** | **string** | Format of the results, possible values are: &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;csv&#x60;, &#x60;html&#x60;, &#x60;xlsx&#x60;, &#x60;xml&#x60; and &#x60;rss&#x60;. The default value is &#x60;json&#x60;.  | 
 **clean** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The &#x60;clean&#x60; parameter is just a shortcut for &#x60;skipHidden&#x3D;true&#x60; and &#x60;skipEmpty&#x3D;true&#x60; parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the &#x60;limit&#x60; value.  | 
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. By default there is no limit. | 
 **fields** | **string** | A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the &#x60;fields&#x60; query parameter. You can use this feature to effectively fix the output format.  | 
 **outputFields** | **string** | A comma-separated list of output field names that positionally rename the fields specified in the &#x60;fields&#x60; parameter. For example, &#x60;?fields&#x3D;headline,url&amp;outputFields&#x3D;title,link&#x60; renames &#x60;headline&#x60; to &#x60;title&#x60; and &#x60;url&#x60; to &#x60;link&#x60; in the output. The number of names in &#x60;outputFields&#x60; must match the number of names in &#x60;fields&#x60;. Requires the &#x60;fields&#x60; parameter to be specified as well.  | 
 **omit** | **string** | A comma-separated list of fields which should be omitted from the items. | 
 **unwind** | **string** | A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the &#x60;desc&#x60; parameter.  | 
 **flatten** | **string** | A comma-separated list of fields which should transform nested objects into flat structures.  For example, with &#x60;flatten&#x3D;\&quot;foo\&quot;&#x60; the object &#x60;{\&quot;foo\&quot;:{\&quot;bar\&quot;: \&quot;hello\&quot;}}&#x60; is turned into &#x60;{\&quot;foo.bar\&quot;: \&quot;hello\&quot;}&#x60;.  The original object with properties is replaced with the flattened object.  | 
 **desc** | **bool** | By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to &#x60;true&#x60; or &#x60;1&#x60;.  | 
 **attachment** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the response will define the &#x60;Content-Disposition: attachment&#x60; header, forcing a web browser to download the file rather than to display it. By default this header is not present.  | 
 **delimiter** | **string** | A delimiter character for CSV files, only used if &#x60;format&#x3D;csv&#x60;. You might need to URL-encode the character (e.g. use &#x60;%09&#x60; for tab or &#x60;%3B&#x60; for semicolon). The default delimiter is a simple comma (&#x60;,&#x60;).  | 
 **bom** | **bool** | All text responses are encoded in UTF-8 encoding. By default, the &#x60;format&#x3D;csv&#x60; files are prefixed with the UTF-8 Byte Order Mark (BOM), while &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;xml&#x60;, &#x60;html&#x60; and &#x60;rss&#x60; files are not.  If you want to override this default behavior, specify &#x60;bom&#x3D;1&#x60; query parameter to include the BOM or &#x60;bom&#x3D;0&#x60; to skip it.  | 
 **xmlRoot** | **string** | Overrides default root element name of &#x60;xml&#x60; output. By default the root element is &#x60;items&#x60;.  | 
 **xmlRow** | **string** | Overrides default element name that wraps each page or page function result object in &#x60;xml&#x60; output. By default the element name is &#x60;item&#x60;.  | 
 **skipHeaderRow** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then header row in the &#x60;csv&#x60; format is skipped. | 
 **skipHidden** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then hidden fields are skipped from the output, i.e. fields starting with the &#x60;#&#x60; character.  | 
 **skipEmpty** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value.  | 
 **simplified** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then, the endpoint applies the &#x60;fields&#x3D;url,pageFunctionResult,errorInfo&#x60; and &#x60;unwind&#x3D;pageFunctionResult&#x60; query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations.  | 
 **view** | **string** | Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema).  | 
 **skipFailedPages** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations.  | 
 **feedTitle** | **string** | Overrides the auto-generated RSS channel &#x60;&lt;title&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the title defaults to &#x60;Dataset &lt;label&gt;&#x60;.  | 
 **feedDescription** | **string** | Overrides the auto-generated RSS channel &#x60;&lt;description&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the description defaults to &#x60;Items in dataset with id \&quot;&lt;datasetId&gt;\&quot;.&#x60;  | 

### Return type

**[]map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActRunSyncGetDatasetItemsPost

> []map[string]interface{} ActRunSyncGetDatasetItemsPost(ctx, actorId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()

Run Actor synchronously and get dataset items



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the run will be restarted if it fails.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)
	format := "json" // string | Format of the results, possible values are: `json`, `jsonl`, `csv`, `html`, `xlsx`, `xml` and `rss`. The default value is `json`.  (optional)
	clean := false // bool | If `true` or `1` then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The `clean` parameter is just a shortcut for `skipHidden=true` and `skipEmpty=true` parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the `limit` value.  (optional)
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1.2) // float64 | Maximum number of items to return. By default there is no limit. (optional)
	fields := "myValue,myOtherValue" // string | A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the `fields` query parameter. You can use this feature to effectively fix the output format.  (optional)
	outputFields := "title,link" // string | A comma-separated list of output field names that positionally rename the fields specified in the `fields` parameter. For example, `?fields=headline,url&outputFields=title,link` renames `headline` to `title` and `url` to `link` in the output. The number of names in `outputFields` must match the number of names in `fields`. Requires the `fields` parameter to be specified as well.  (optional)
	omit := "myValue,myOtherValue" // string | A comma-separated list of fields which should be omitted from the items. (optional)
	unwind := "myValue,myOtherValue" // string | A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the `desc` parameter.  (optional)
	flatten := "myValue" // string | A comma-separated list of fields which should transform nested objects into flat structures.  For example, with `flatten=\"foo\"` the object `{\"foo\":{\"bar\": \"hello\"}}` is turned into `{\"foo.bar\": \"hello\"}`.  The original object with properties is replaced with the flattened object.  (optional)
	desc := true // bool | By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to `true` or `1`.  (optional)
	attachment := true // bool | If `true` or `1` then the response will define the `Content-Disposition: attachment` header, forcing a web browser to download the file rather than to display it. By default this header is not present.  (optional)
	delimiter := ";" // string | A delimiter character for CSV files, only used if `format=csv`. You might need to URL-encode the character (e.g. use `%09` for tab or `%3B` for semicolon). The default delimiter is a simple comma (`,`).  (optional)
	bom := false // bool | All text responses are encoded in UTF-8 encoding. By default, the `format=csv` files are prefixed with the UTF-8 Byte Order Mark (BOM), while `json`, `jsonl`, `xml`, `html` and `rss` files are not.  If you want to override this default behavior, specify `bom=1` query parameter to include the BOM or `bom=0` to skip it.  (optional)
	xmlRoot := "items" // string | Overrides default root element name of `xml` output. By default the root element is `items`.  (optional)
	xmlRow := "item" // string | Overrides default element name that wraps each page or page function result object in `xml` output. By default the element name is `item`.  (optional)
	skipHeaderRow := true // bool | If `true` or `1` then header row in the `csv` format is skipped. (optional)
	skipHidden := false // bool | If `true` or `1` then hidden fields are skipped from the output, i.e. fields starting with the `#` character.  (optional)
	skipEmpty := false // bool | If `true` or `1` then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value.  (optional)
	simplified := false // bool | If `true` or `1` then, the endpoint applies the `fields=url,pageFunctionResult,errorInfo` and `unwind=pageFunctionResult` query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it's not recommended to use it in new integrations.  (optional)
	view := "overview" // string | Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema).  (optional)
	skipFailedPages := false // bool | If `true` or `1` then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it's not recommended to use it in new integrations.  (optional)
	feedTitle := "Latest posts from r/pasta" // string | Overrides the auto-generated RSS channel `<title>` element. Only used when `format=rss`. If not provided, the title defaults to `Dataset <label>`.  (optional)
	feedDescription := "Scraped forum posts" // string | Overrides the auto-generated RSS channel `<description>` element. Only used when `format=rss`. If not provided, the description defaults to `Items in dataset with id \"<datasetId>\".`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunSyncGetDatasetItemsPost(context.Background(), actorId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunSyncGetDatasetItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunSyncGetDatasetItemsPost`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunSyncGetDatasetItemsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunSyncGetDatasetItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **maxTotalChargeUsd** | **float64** | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable.  | 
 **restartOnError** | **bool** | Determines whether the run will be restarted if it fails.  | 
 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;).  | 
 **webhooks** | **string** | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  | 
 **format** | **string** | Format of the results, possible values are: &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;csv&#x60;, &#x60;html&#x60;, &#x60;xlsx&#x60;, &#x60;xml&#x60; and &#x60;rss&#x60;. The default value is &#x60;json&#x60;.  | 
 **clean** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the API endpoint returns only non-empty items and skips hidden fields (i.e. fields starting with the # character). The &#x60;clean&#x60; parameter is just a shortcut for &#x60;skipHidden&#x3D;true&#x60; and &#x60;skipEmpty&#x3D;true&#x60; parameters. Note that since some objects might be skipped from the output, that the result might contain less items than the &#x60;limit&#x60; value.  | 
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. By default there is no limit. | 
 **fields** | **string** | A comma-separated list of fields which should be picked from the items, only these fields will remain in the resulting record objects. Note that the fields in the outputted items are sorted the same way as they are specified in the &#x60;fields&#x60; query parameter. You can use this feature to effectively fix the output format.  | 
 **outputFields** | **string** | A comma-separated list of output field names that positionally rename the fields specified in the &#x60;fields&#x60; parameter. For example, &#x60;?fields&#x3D;headline,url&amp;outputFields&#x3D;title,link&#x60; renames &#x60;headline&#x60; to &#x60;title&#x60; and &#x60;url&#x60; to &#x60;link&#x60; in the output. The number of names in &#x60;outputFields&#x60; must match the number of names in &#x60;fields&#x60;. Requires the &#x60;fields&#x60; parameter to be specified as well.  | 
 **omit** | **string** | A comma-separated list of fields which should be omitted from the items. | 
 **unwind** | **string** | A comma-separated list of fields which should be unwound, in order which they should be processed. Each field should be either an array or an object. If the field is an array then every element of the array will become a separate record and merged with parent object. If the unwound field is an object then it is merged with the parent object. If the unwound field is missing or its value is neither an array nor an object and therefore cannot be merged with a parent object then the item gets preserved as it is. Note that the unwound items ignore the &#x60;desc&#x60; parameter.  | 
 **flatten** | **string** | A comma-separated list of fields which should transform nested objects into flat structures.  For example, with &#x60;flatten&#x3D;\&quot;foo\&quot;&#x60; the object &#x60;{\&quot;foo\&quot;:{\&quot;bar\&quot;: \&quot;hello\&quot;}}&#x60; is turned into &#x60;{\&quot;foo.bar\&quot;: \&quot;hello\&quot;}&#x60;.  The original object with properties is replaced with the flattened object.  | 
 **desc** | **bool** | By default, results are returned in the same order as they were stored. To reverse the order, set this parameter to &#x60;true&#x60; or &#x60;1&#x60;.  | 
 **attachment** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the response will define the &#x60;Content-Disposition: attachment&#x60; header, forcing a web browser to download the file rather than to display it. By default this header is not present.  | 
 **delimiter** | **string** | A delimiter character for CSV files, only used if &#x60;format&#x3D;csv&#x60;. You might need to URL-encode the character (e.g. use &#x60;%09&#x60; for tab or &#x60;%3B&#x60; for semicolon). The default delimiter is a simple comma (&#x60;,&#x60;).  | 
 **bom** | **bool** | All text responses are encoded in UTF-8 encoding. By default, the &#x60;format&#x3D;csv&#x60; files are prefixed with the UTF-8 Byte Order Mark (BOM), while &#x60;json&#x60;, &#x60;jsonl&#x60;, &#x60;xml&#x60;, &#x60;html&#x60; and &#x60;rss&#x60; files are not.  If you want to override this default behavior, specify &#x60;bom&#x3D;1&#x60; query parameter to include the BOM or &#x60;bom&#x3D;0&#x60; to skip it.  | 
 **xmlRoot** | **string** | Overrides default root element name of &#x60;xml&#x60; output. By default the root element is &#x60;items&#x60;.  | 
 **xmlRow** | **string** | Overrides default element name that wraps each page or page function result object in &#x60;xml&#x60; output. By default the element name is &#x60;item&#x60;.  | 
 **skipHeaderRow** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then header row in the &#x60;csv&#x60; format is skipped. | 
 **skipHidden** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then hidden fields are skipped from the output, i.e. fields starting with the &#x60;#&#x60; character.  | 
 **skipEmpty** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then empty items are skipped from the output.  Note that if used, the results might contain less items than the limit value.  | 
 **simplified** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then, the endpoint applies the &#x60;fields&#x3D;url,pageFunctionResult,errorInfo&#x60; and &#x60;unwind&#x3D;pageFunctionResult&#x60; query parameters. This feature is used to emulate simplified results provided by the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations.  | 
 **view** | **string** | Defines the view configuration for dataset items based on the schema definition. This parameter determines how the data will be filtered and presented. For complete specification details, see the [dataset schema documentation](/platform/actors/development/actor-definition/dataset-schema).  | 
 **skipFailedPages** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then, the all the items with errorInfo property will be skipped from the output.  This feature is here to emulate functionality of API version 1 used for the legacy Apify Crawler product and it&#39;s not recommended to use it in new integrations.  | 
 **feedTitle** | **string** | Overrides the auto-generated RSS channel &#x60;&lt;title&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the title defaults to &#x60;Dataset &lt;label&gt;&#x60;.  | 
 **feedDescription** | **string** | Overrides the auto-generated RSS channel &#x60;&lt;description&gt;&#x60; element. Only used when &#x60;format&#x3D;rss&#x60;. If not provided, the description defaults to &#x60;Items in dataset with id \&quot;&lt;datasetId&gt;\&quot;.&#x60;  | 

### Return type

**[]map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActRunSyncPost

> map[string]interface{} ActRunSyncPost(ctx, actorId).Body(body).OutputRecordKey(outputRecordKey).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Execute()

Run Actor synchronously and return output



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	outputRecordKey := "OUTPUT" // string | Key of the record from run's default key-value store to be returned in the response. By default, it is `OUTPUT`.  (optional)
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the run will be restarted if it fails.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunSyncPost(context.Background(), actorId).Body(body).OutputRecordKey(outputRecordKey).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunSyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunSyncPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunSyncPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunSyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **outputRecordKey** | **string** | Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;.  | 
 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **maxTotalChargeUsd** | **float64** | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable.  | 
 **restartOnError** | **bool** | Determines whether the run will be restarted if it fails.  | 
 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;).  | 
 **webhooks** | **string** | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  | 

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


## ActRunsGet

> ListOfRunsResponse ActRunsGet(ctx, actorId).Offset(offset).Limit(limit).Desc(desc).Status(status).StartedAfter(startedAfter).StartedBefore(startedBefore).Execute()

Get list of runs



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
	actorId := "janedoe~my-actor" // string | Actor ID or a tilde-separated owner's username and Actor name.
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	desc := true // bool | If `true` or `1` then the objects are sorted by the `startedAt` field in descending order. By default, they are sorted in ascending order.  (optional)
	status := []string{"Inner_example"} // []string | Single status or comma-separated list of statuses, see ([available statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)). Used to filter runs by the specified statuses only.  (optional)
	startedAfter := time.Now() // time.Time | Filter runs that started after the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC).  (optional)
	startedBefore := time.Now() // time.Time | Filter runs that started before the specified date and time (inclusive). The value must be a valid ISO 8601 datetime string (UTC).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunsGet(context.Background(), actorId).Offset(offset).Limit(limit).Desc(desc).Status(status).StartedAfter(startedAfter).StartedBefore(startedBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsGet`: ListOfRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsGetRequest struct via the builder pattern


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


## ActRunsLastGet

> RunResponse ActRunsLastGet(ctx, actorId).Status(status).WaitForFinish(waitForFinish).Execute()

Get last run



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
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the run to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the run finishes in time then the returned run object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunsLastGet(context.Background(), actorId).Status(status).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunsLastGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsLastGet`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunsLastGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsLastGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
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


## ActRunsPost

> RunResponse ActRunsPost(ctx, actorId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).WaitForFinish(waitForFinish).Webhooks(webhooks).ForcePermissionLevel(forcePermissionLevel).Execute()

Run Actor



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the run will be restarted if it fails.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the run to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the run finishes in time then the returned run object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)
	forcePermissionLevel := "LIMITED_PERMISSIONS" // string | Overrides the Actor's permission level for this specific run. Use to test restricted permissions before deploying changes to your Actor or to temporarily elevate or restrict access. If you don't specify this parameter, the Actor uses its configured default permission level. For more information on permissions, see the [documentation](https://docs.apify.com/platform/actors/development/permissions).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorsActorRunsAPI.ActRunsPost(context.Background(), actorId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).WaitForFinish(waitForFinish).Webhooks(webhooks).ForcePermissionLevel(forcePermissionLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorsActorRunsAPI.ActRunsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActRunsPost`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorsActorRunsAPI.ActRunsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorId** | **string** | Actor ID or a tilde-separated owner&#39;s username and Actor name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActRunsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **maxTotalChargeUsd** | **float64** | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable.  | 
 **restartOnError** | **bool** | Determines whether the run will be restarted if it fails.  | 
 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;).  | 
 **waitForFinish** | **float64** | The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;).  | 
 **webhooks** | **string** | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  | 
 **forcePermissionLevel** | **string** | Overrides the Actor&#39;s permission level for this specific run. Use to test restricted permissions before deploying changes to your Actor or to temporarily elevate or restrict access. If you don&#39;t specify this parameter, the Actor uses its configured default permission level. For more information on permissions, see the [documentation](https://docs.apify.com/platform/actors/development/permissions).  | 

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

