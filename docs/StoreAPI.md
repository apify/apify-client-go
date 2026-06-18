# \StoreAPI

All URIs are relative to *https://api.apify.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StoreGet**](StoreAPI.md#StoreGet) | **Get** /v2/store | Get list of Actors in Store



## StoreGet

> ListOfActorsInStoreResponse StoreGet(ctx).Limit(limit).Offset(offset).Search(search).SortBy(sortBy).Category(category).Username(username).PricingModel(pricingModel).AllowsAgenticUsers(allowsAgenticUsers).ResponseFormat(responseFormat).IncludeUnrunnableActors(includeUnrunnableActors).Execute()

Get list of Actors in Store



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
	limit := float64(1000) // float64 | Maximum number of items to return. The default value as well as the maximum is `1000`.  (optional)
	offset := float64(0) // float64 | Number of items that should be skipped at the start. The default value is `0`.  (optional)
	search := "web scraper" // string | String to search by. The search runs on the following fields: `title`, `name`, `description`, `username`, `readme`.  (optional)
	sortBy := "'popularity'" // string | Specifies the field by which to sort the results. The supported values are `relevance` (default), `popularity`, `newest` and `lastUpdate`.  (optional)
	category := "'AI'" // string | Filters the results by the specified category. (optional)
	username := "'apify'" // string | Filters the results by the specified username. (optional)
	pricingModel := "FREE" // string | Only return Actors with the specified pricing model.  (optional)
	allowsAgenticUsers := true // bool | If true, only return Actors that allow agentic users. If false, only return Actors that do not allow agentic users.  (optional)
	responseFormat := "agent" // string | Controls the shape of the response. Use `full` (default) for the complete response including image URLs and all fields. Use `agent` for a reduced field set optimized for LLM consumers, which only includes `id`, `title`, `name`, `username`, `description`, `notice`, `badge`, `categories`, and minimal `stats`.  (optional) (default to "full")
	includeUnrunnableActors := true // bool | By default, search results exclude Actors that are not safe to run automatically (e.g. Actors from developers who haven't passed KYC, or full-permission Actors without a large user base). Set to `true` to bypass this safety filtering and include all Actors in the results.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.StoreGet(context.Background()).Limit(limit).Offset(offset).Search(search).SortBy(sortBy).Category(category).Username(username).PricingModel(pricingModel).AllowsAgenticUsers(allowsAgenticUsers).ResponseFormat(responseFormat).IncludeUnrunnableActors(includeUnrunnableActors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.StoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreGet`: ListOfActorsInStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.StoreGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float64** | Maximum number of items to return. The default value as well as the maximum is &#x60;1000&#x60;.  | 
 **offset** | **float64** | Number of items that should be skipped at the start. The default value is &#x60;0&#x60;.  | 
 **search** | **string** | String to search by. The search runs on the following fields: &#x60;title&#x60;, &#x60;name&#x60;, &#x60;description&#x60;, &#x60;username&#x60;, &#x60;readme&#x60;.  | 
 **sortBy** | **string** | Specifies the field by which to sort the results. The supported values are &#x60;relevance&#x60; (default), &#x60;popularity&#x60;, &#x60;newest&#x60; and &#x60;lastUpdate&#x60;.  | 
 **category** | **string** | Filters the results by the specified category. | 
 **username** | **string** | Filters the results by the specified username. | 
 **pricingModel** | **string** | Only return Actors with the specified pricing model.  | 
 **allowsAgenticUsers** | **bool** | If true, only return Actors that allow agentic users. If false, only return Actors that do not allow agentic users.  | 
 **responseFormat** | **string** | Controls the shape of the response. Use &#x60;full&#x60; (default) for the complete response including image URLs and all fields. Use &#x60;agent&#x60; for a reduced field set optimized for LLM consumers, which only includes &#x60;id&#x60;, &#x60;title&#x60;, &#x60;name&#x60;, &#x60;username&#x60;, &#x60;description&#x60;, &#x60;notice&#x60;, &#x60;badge&#x60;, &#x60;categories&#x60;, and minimal &#x60;stats&#x60;.  | [default to &quot;full&quot;]
 **includeUnrunnableActors** | **bool** | By default, search results exclude Actors that are not safe to run automatically (e.g. Actors from developers who haven&#39;t passed KYC, or full-permission Actors without a large user base). Set to &#x60;true&#x60; to bypass this safety filtering and include all Actors in the results.  | [default to false]

### Return type

[**ListOfActorsInStoreResponse**](ListOfActorsInStoreResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [httpBearer](../README.md#httpBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

