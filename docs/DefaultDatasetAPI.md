# \DefaultDatasetAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorRunDatasetDelete**](DefaultDatasetAPI.md#ActorRunDatasetDelete) | **Delete** /v2/actor-runs/{runId}/dataset | Delete default dataset
[**ActorRunDatasetGet**](DefaultDatasetAPI.md#ActorRunDatasetGet) | **Get** /v2/actor-runs/{runId}/dataset | Get default dataset
[**ActorRunDatasetItemsGet**](DefaultDatasetAPI.md#ActorRunDatasetItemsGet) | **Get** /v2/actor-runs/{runId}/dataset/items | Get default dataset items
[**ActorRunDatasetItemsPost**](DefaultDatasetAPI.md#ActorRunDatasetItemsPost) | **Post** /v2/actor-runs/{runId}/dataset/items | Store items
[**ActorRunDatasetPut**](DefaultDatasetAPI.md#ActorRunDatasetPut) | **Put** /v2/actor-runs/{runId}/dataset | Update default dataset
[**ActorRunDatasetStatisticsGet**](DefaultDatasetAPI.md#ActorRunDatasetStatisticsGet) | **Get** /v2/actor-runs/{runId}/dataset/statistics | Get default dataset statistics



## ActorRunDatasetDelete

> ActorRunDatasetDelete(ctx, runId).Execute()

Delete default dataset



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
	r, err := apiClient.DefaultDatasetAPI.ActorRunDatasetDelete(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultDatasetAPI.ActorRunDatasetDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActorRunDatasetDeleteRequest struct via the builder pattern


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


## ActorRunDatasetGet

> DatasetResponse ActorRunDatasetGet(ctx, runId).Execute()

Get default dataset



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
	resp, r, err := apiClient.DefaultDatasetAPI.ActorRunDatasetGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultDatasetAPI.ActorRunDatasetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunDatasetGet`: DatasetResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultDatasetAPI.ActorRunDatasetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunDatasetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasetResponse**](DatasetResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunDatasetItemsGet

> []map[string]interface{} ActorRunDatasetItemsGet(ctx, runId).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Signature(signature).Execute()

Get default dataset items



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
	signature := "2wTI46Bg8qWQrV7tavlPI" // string | Signature used for the access. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultDatasetAPI.ActorRunDatasetItemsGet(context.Background(), runId).Format(format).Clean(clean).Offset(offset).Limit(limit).Fields(fields).OutputFields(outputFields).Omit(omit).Unwind(unwind).Flatten(flatten).Desc(desc).Attachment(attachment).Delimiter(delimiter).Bom(bom).XmlRoot(xmlRoot).XmlRow(xmlRow).SkipHeaderRow(skipHeaderRow).SkipHidden(skipHidden).SkipEmpty(skipEmpty).Simplified(simplified).View(view).SkipFailedPages(skipFailedPages).FeedTitle(feedTitle).FeedDescription(feedDescription).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultDatasetAPI.ActorRunDatasetItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunDatasetItemsGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultDatasetAPI.ActorRunDatasetItemsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunDatasetItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 **signature** | **string** | Signature used for the access. | 

### Return type

**[]map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/jsonl, text/csv, text/html, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/rss+xml, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunDatasetItemsPost

> map[string]interface{} ActorRunDatasetItemsPost(ctx, runId).ActRunsLastDatasetItemsPostRequest(actRunsLastDatasetItemsPostRequest).Execute()

Store items



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
	actRunsLastDatasetItemsPostRequest := openapiclient.act_runs_last_dataset_items_post_request{ArrayOfMapmapOfStringAny: new([]map[string]interface{})} // ActRunsLastDatasetItemsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultDatasetAPI.ActorRunDatasetItemsPost(context.Background(), runId).ActRunsLastDatasetItemsPostRequest(actRunsLastDatasetItemsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultDatasetAPI.ActorRunDatasetItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunDatasetItemsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultDatasetAPI.ActorRunDatasetItemsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunDatasetItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actRunsLastDatasetItemsPostRequest** | [**ActRunsLastDatasetItemsPostRequest**](ActRunsLastDatasetItemsPostRequest.md) |  | 

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


## ActorRunDatasetPut

> DatasetResponse ActorRunDatasetPut(ctx, runId).UpdateDatasetRequest(updateDatasetRequest).Execute()

Update default dataset



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
	updateDatasetRequest := *openapiclient.NewUpdateDatasetRequest() // UpdateDatasetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultDatasetAPI.ActorRunDatasetPut(context.Background(), runId).UpdateDatasetRequest(updateDatasetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultDatasetAPI.ActorRunDatasetPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunDatasetPut`: DatasetResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultDatasetAPI.ActorRunDatasetPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunDatasetPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatasetRequest** | [**UpdateDatasetRequest**](UpdateDatasetRequest.md) |  | 

### Return type

[**DatasetResponse**](DatasetResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorRunDatasetStatisticsGet

> DatasetStatisticsResponse ActorRunDatasetStatisticsGet(ctx, runId).Execute()

Get default dataset statistics



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
	resp, r, err := apiClient.DefaultDatasetAPI.ActorRunDatasetStatisticsGet(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultDatasetAPI.ActorRunDatasetStatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActorRunDatasetStatisticsGet`: DatasetStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultDatasetAPI.ActorRunDatasetStatisticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Actor run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorRunDatasetStatisticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasetStatisticsResponse**](DatasetStatisticsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

