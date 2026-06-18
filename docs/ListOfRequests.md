# ListOfRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Request**](Request.md) | The array of requests. | 
**Limit** | **int32** | The maximum number of requests returned in this response. | 
**ExclusiveStartId** | Pointer to **string** | The ID of the last request from the previous page, used for pagination. | [optional] 
**Cursor** | Pointer to **string** | A cursor string used for current page of results. | [optional] 
**NextCursor** | Pointer to **string** | A cursor string to be used to continue pagination. | [optional] 

## Methods

### NewListOfRequests

`func NewListOfRequests(items []Request, limit int32, ) *ListOfRequests`

NewListOfRequests instantiates a new ListOfRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfRequestsWithDefaults

`func NewListOfRequestsWithDefaults() *ListOfRequests`

NewListOfRequestsWithDefaults instantiates a new ListOfRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListOfRequests) GetItems() []Request`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfRequests) GetItemsOk() (*[]Request, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfRequests) SetItems(v []Request)`

SetItems sets Items field to given value.


### GetLimit

`func (o *ListOfRequests) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOfRequests) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOfRequests) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetExclusiveStartId

`func (o *ListOfRequests) GetExclusiveStartId() string`

GetExclusiveStartId returns the ExclusiveStartId field if non-nil, zero value otherwise.

### GetExclusiveStartIdOk

`func (o *ListOfRequests) GetExclusiveStartIdOk() (*string, bool)`

GetExclusiveStartIdOk returns a tuple with the ExclusiveStartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveStartId

`func (o *ListOfRequests) SetExclusiveStartId(v string)`

SetExclusiveStartId sets ExclusiveStartId field to given value.

### HasExclusiveStartId

`func (o *ListOfRequests) HasExclusiveStartId() bool`

HasExclusiveStartId returns a boolean if a field has been set.

### GetCursor

`func (o *ListOfRequests) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListOfRequests) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListOfRequests) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListOfRequests) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListOfRequests) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListOfRequests) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListOfRequests) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListOfRequests) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


