# \ActorTasksAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorTaskDelete**](ActorTasksAPI.md#ActorTaskDelete) | **Delete** /v2/actor-tasks/{actorTaskId} | Delete task
[**ActorTaskGet**](ActorTasksAPI.md#ActorTaskGet) | **Get** /v2/actor-tasks/{actorTaskId} | Get task
[**ActorTaskInputGet**](ActorTasksAPI.md#ActorTaskInputGet) | **Get** /v2/actor-tasks/{actorTaskId}/input | Get task input
[**ActorTaskInputPut**](ActorTasksAPI.md#ActorTaskInputPut) | **Put** /v2/actor-tasks/{actorTaskId}/input | Update task input
[**ActorTaskPut**](ActorTasksAPI.md#ActorTaskPut) | **Put** /v2/actor-tasks/{actorTaskId} | Update task
[**ActorTaskRunSyncGet**](ActorTasksAPI.md#ActorTaskRunSyncGet) | **Get** /v2/actor-tasks/{actorTaskId}/run-sync | Run task synchronously
[**ActorTaskRunSyncGetDatasetItemsGet**](ActorTasksAPI.md#ActorTaskRunSyncGetDatasetItemsGet) | **Get** /v2/actor-tasks/{actorTaskId}/run-sync-get-dataset-items | Run task synchronously and get dataset items
[**ActorTaskRunSyncGetDatasetItemsPost**](ActorTasksAPI.md#ActorTaskRunSyncGetDatasetItemsPost) | **Post** /v2/actor-tasks/{actorTaskId}/run-sync-get-dataset-items | Run task synchronously and get dataset items
[**ActorTaskRunSyncPost**](ActorTasksAPI.md#ActorTaskRunSyncPost) | **Post** /v2/actor-tasks/{actorTaskId}/run-sync | Run task synchronously
[**ActorTaskRunsGet**](ActorTasksAPI.md#ActorTaskRunsGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs | Get list of task runs
[**ActorTaskRunsLastGet**](ActorTasksAPI.md#ActorTaskRunsLastGet) | **Get** /v2/actor-tasks/{actorTaskId}/runs/last | Get last run
[**ActorTaskRunsPost**](ActorTasksAPI.md#ActorTaskRunsPost) | **Post** /v2/actor-tasks/{actorTaskId}/runs | Run task
[**ActorTaskWebhooksGet**](ActorTasksAPI.md#ActorTaskWebhooksGet) | **Get** /v2/actor-tasks/{actorTaskId}/webhooks | Get list of webhooks
[**ActorTasksGet**](ActorTasksAPI.md#ActorTasksGet) | **Get** /v2/actor-tasks | Get list of tasks
[**ActorTasksPost**](ActorTasksAPI.md#ActorTasksPost) | **Post** /v2/actor-tasks | Create task



## ActorTaskDelete

> map[string]interface{} ActorTaskDelete(ctx, actorTaskId).Execute()

Delete task



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskDelete(context.Background(), actorTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskDeleteRequest struct via the builder pattern


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


## ActorTaskGet

> ActorTaskGet200Response ActorTaskGet(ctx, actorTaskId).Execute()

Get task



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskGet(context.Background(), actorTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskGet`: ActorTaskGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActorTaskGet200Response**](ActorTaskGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTaskInputGet

> map[string]interface{} ActorTaskInputGet(ctx, actorTaskId).Execute()

Get task input



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskInputGet(context.Background(), actorTaskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskInputGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskInputGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskInputGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskInputGetRequest struct via the builder pattern


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


## ActorTaskInputPut

> map[string]interface{} ActorTaskInputPut(ctx, actorTaskId).Body(body).Execute()

Update task input



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskInputPut(context.Background(), actorTaskId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskInputPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskInputPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskInputPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskInputPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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


## ActorTaskPut

> ActorTaskGet200Response ActorTaskPut(ctx, actorTaskId).UpdateTaskRequest(updateTaskRequest).Execute()

Update task



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
	updateTaskRequest := *openapiclient.NewUpdateTaskRequest() // UpdateTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskPut(context.Background(), actorTaskId).UpdateTaskRequest(updateTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskPut`: ActorTaskGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskRequest** | [**UpdateTaskRequest**](UpdateTaskRequest.md) |  | 

### Return type

[**ActorTaskGet200Response**](ActorTaskGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTaskRunSyncGet

> map[string]interface{} ActorTaskRunSyncGet(ctx, actorTaskId).Timeout(timeout).Memory(memory).MaxItems(maxItems).Build(build).OutputRecordKey(outputRecordKey).Webhooks(webhooks).Execute()

Run task synchronously



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
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	outputRecordKey := "OUTPUT" // string | Key of the record from run's default key-value store to be returned in the response. By default, it is `OUTPUT`.  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskRunSyncGet(context.Background(), actorTaskId).Timeout(timeout).Memory(memory).MaxItems(maxItems).Build(build).OutputRecordKey(outputRecordKey).Webhooks(webhooks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskRunSyncGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunSyncGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskRunSyncGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunSyncGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;).  | 
 **outputRecordKey** | **string** | Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;.  | 
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


## ActorTaskRunSyncGetDatasetItemsGet

> map[string]interface{} ActorTaskRunSyncGetDatasetItemsGet(ctx, actorTaskId).Timeout(timeout).Memory(memory).MaxItems(maxItems).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()

Run task synchronously and get dataset items



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
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
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
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskRunSyncGetDatasetItemsGet(context.Background(), actorTaskId).Timeout(timeout).Memory(memory).MaxItems(maxItems).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskRunSyncGetDatasetItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunSyncGetDatasetItemsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskRunSyncGetDatasetItemsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunSyncGetDatasetItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
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

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTaskRunSyncGetDatasetItemsPost

> map[string]interface{} ActorTaskRunSyncGetDatasetItemsPost(ctx, actorTaskId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()

Run task synchronously and get dataset items



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
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskRunSyncGetDatasetItemsPost(context.Background(), actorTaskId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).Webhooks(webhooks).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskRunSyncGetDatasetItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunSyncGetDatasetItemsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskRunSyncGetDatasetItemsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunSyncGetDatasetItemsPostRequest struct via the builder pattern


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

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTaskRunSyncPost

> map[string]interface{} ActorTaskRunSyncPost(ctx, actorTaskId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).OutputRecordKey(outputRecordKey).Webhooks(webhooks).Execute()

Run task synchronously



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the run will be restarted if it fails.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	outputRecordKey := "OUTPUT" // string | Key of the record from run's default key-value store to be returned in the response. By default, it is `OUTPUT`.  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskRunSyncPost(context.Background(), actorTaskId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).OutputRecordKey(outputRecordKey).Webhooks(webhooks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskRunSyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunSyncPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskRunSyncPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunSyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **timeout** | **float64** | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  | 
 **memory** | **float64** | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  | 
 **maxItems** | **float64** | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won&#39;t be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using &#x60;ACTOR_MAX_PAID_DATASET_ITEMS&#x60; environment variable.  | 
 **maxTotalChargeUsd** | **float64** | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the &#x60;ACTOR_MAX_TOTAL_CHARGE_USD&#x60; environment variable.  | 
 **restartOnError** | **bool** | Determines whether the run will be restarted if it fails.  | 
 **build** | **string** | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically &#x60;latest&#x60;).  | 
 **outputRecordKey** | **string** | Key of the record from run&#39;s default key-value store to be returned in the response. By default, it is &#x60;OUTPUT&#x60;.  | 
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


## ActorTaskRunsGet

> ActorTaskRunsGet200Response ActorTaskRunsGet(ctx, actorTaskId).Offset(offset).Limit(limit).Desc(desc).Status(status).Execute()

Get list of task runs



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
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	desc := true // bool | If `true` or `1` then the objects are sorted by the `startedAt` field in descending order. By default, they are sorted in ascending order.  (optional)
	status := []string{"Inner_example"} // []string | Single status or comma-separated list of statuses, see ([available statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)). Used to filter runs by the specified statuses only.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskRunsGet(context.Background(), actorTaskId).Offset(offset).Limit(limit).Desc(desc).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskRunsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsGet`: ActorTaskRunsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskRunsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;startedAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 
 **status** | **[]string** | Single status or comma-separated list of statuses, see ([available statuses](https://docs.apify.com/platform/actors/running/runs-and-builds#lifecycle)). Used to filter runs by the specified statuses only.  | 

### Return type

[**ActorTaskRunsGet200Response**](ActorTaskRunsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTaskRunsLastGet

> ActorTaskRunsPost201Response ActorTaskRunsLastGet(ctx, actorTaskId).Status(status).WaitForFinish(waitForFinish).Execute()

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
	actorTaskId := "janedoe~my-task" // string | Task ID or a tilde-separated owner's username and task's name.
	status := "SUCCEEDED" // string | Filter for the run status. (optional)
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the run to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the run finishes in time then the returned run object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskRunsLastGet(context.Background(), actorTaskId).Status(status).WaitForFinish(waitForFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskRunsLastGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsLastGet`: ActorTaskRunsPost201Response
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskRunsLastGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsLastGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter for the run status. | 
 **waitForFinish** | **float64** | The maximum number of seconds the server waits for the run to finish. By default it is &#x60;0&#x60;, the maximum value is &#x60;60&#x60;. &lt;!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --&gt; If the run finishes in time then the returned run object will have a terminal status (e.g. &#x60;SUCCEEDED&#x60;), otherwise it will have a transitional status (e.g. &#x60;RUNNING&#x60;).  | 

### Return type

[**ActorTaskRunsPost201Response**](ActorTaskRunsPost201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTaskRunsPost

> ActorTaskRunsPost201Response ActorTaskRunsPost(ctx, actorTaskId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).WaitForFinish(waitForFinish).Webhooks(webhooks).Execute()

Run task



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	timeout := float64(60) // float64 | Optional timeout for the run, in seconds. By default, the run uses the timeout from its configuration.  (optional)
	memory := float64(256) // float64 | Memory limit for the run, in megabytes. The amount of memory can be set to a power of 2 with a minimum of 128. By default, the run uses the memory limit from its configuration.  (optional)
	maxItems := float64(1000) // float64 | Specifies the maximum number of dataset items that will be charged for pay-per-result Actors. This does NOT guarantee that the Actor will return only this many items. It only ensures you won't be charged for more than this number of items. Only works for pay-per-result Actors. Value can be accessed in the actor run using `ACTOR_MAX_PAID_DATASET_ITEMS` environment variable.  (optional)
	maxTotalChargeUsd := float64(5) // float64 | Specifies the maximum cost of the run. This parameter is useful for pay-per-event Actors, as it allows you to limit the amount charged to your subscription. You can access the maximum cost in your Actor by using the `ACTOR_MAX_TOTAL_CHARGE_USD` environment variable.  (optional)
	restartOnError := false // bool | Determines whether the run will be restarted if it fails.  (optional)
	build := "0.1.234" // string | Specifies the Actor build to run. It can be either a build tag or build number. By default, the run uses the build from its configuration (typically `latest`).  (optional)
	waitForFinish := float64(60) // float64 | The maximum number of seconds the server waits for the run to finish. By default it is `0`, the maximum value is `60`. <!-- MAX_ACTOR_JOB_ASYNC_WAIT_SECS --> If the run finishes in time then the returned run object will have a terminal status (e.g. `SUCCEEDED`), otherwise it will have a transitional status (e.g. `RUNNING`).  (optional)
	webhooks := "dGhpcyBpcyBqdXN0IGV4YW1wbGUK..." // string | Specifies optional webhooks associated with the Actor run, which can be used to receive a notification e.g. when the Actor finished or failed. The value is a Base64-encoded JSON array whose items follow the WebhookRepresentation schema. For more information, see [Webhooks documentation](https://docs.apify.com/platform/integrations/webhooks).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskRunsPost(context.Background(), actorTaskId).Body(body).Timeout(timeout).Memory(memory).MaxItems(maxItems).MaxTotalChargeUsd(maxTotalChargeUsd).RestartOnError(restartOnError).Build(build).WaitForFinish(waitForFinish).Webhooks(webhooks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskRunsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskRunsPost`: ActorTaskRunsPost201Response
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskRunsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskRunsPostRequest struct via the builder pattern


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

### Return type

[**ActorTaskRunsPost201Response**](ActorTaskRunsPost201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTaskWebhooksGet

> ActorTaskWebhooksGet200Response ActorTaskWebhooksGet(ctx, actorTaskId).Offset(offset).Limit(limit).Desc(desc).Execute()

Get list of webhooks



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
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	desc := true // bool | If `true` or `1` then the objects are sorted by the `createdAt` field in descending order. By default, they are sorted in ascending order.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTaskWebhooksGet(context.Background(), actorTaskId).Offset(offset).Limit(limit).Desc(desc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTaskWebhooksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTaskWebhooksGet`: ActorTaskWebhooksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTaskWebhooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actorTaskId** | **string** | Task ID or a tilde-separated owner&#39;s username and task&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorTaskWebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 

### Return type

[**ActorTaskWebhooksGet200Response**](ActorTaskWebhooksGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTasksGet

> ListOfTasksResponse ActorTasksGet(ctx).Offset(offset).Limit(limit).Desc(desc).Execute()

Get list of tasks



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
	resp, r, err := apiClient.ActorTasksAPI.ActorTasksGet(context.Background()).Offset(offset).Limit(limit).Desc(desc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTasksGet`: ListOfTasksResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActorTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **desc** | **bool** | If &#x60;true&#x60; or &#x60;1&#x60; then the objects are sorted by the &#x60;createdAt&#x60; field in descending order. By default, they are sorted in ascending order.  | 

### Return type

[**ListOfTasksResponse**](ListOfTasksResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorTasksPost

> TaskResponse ActorTasksPost(ctx).CreateTaskRequest(createTaskRequest).Execute()

Create task



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
	createTaskRequest := *openapiclient.NewCreateTaskRequest("ActId_example") // CreateTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActorTasksAPI.ActorTasksPost(context.Background()).CreateTaskRequest(createTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActorTasksAPI.ActorTasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorTasksPost`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ActorTasksAPI.ActorTasksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActorTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTaskRequest** | [**CreateTaskRequest**](CreateTaskRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

