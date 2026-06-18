# ListOfRequestQueues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total number of items available across all pages. | 
**Offset** | **int32** | The starting position for this page of results. | 
**Limit** | **int32** | The maximum number of items returned per page. | 
**Desc** | **bool** | Whether the results are sorted in descending order. | 
**Count** | **int32** | The number of items returned in this response. | 
**Unnamed** | Pointer to **bool** | Whether the listing was filtered to only unnamed request queues. | [optional] 
**Items** | [**[]RequestQueueShort**](RequestQueueShort.md) | The array of request queues. | 

## Methods

### NewListOfRequestQueues

`func NewListOfRequestQueues(total int32, offset int32, limit int32, desc bool, count int32, items []RequestQueueShort, ) *ListOfRequestQueues`

NewListOfRequestQueues instantiates a new ListOfRequestQueues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfRequestQueuesWithDefaults

`func NewListOfRequestQueuesWithDefaults() *ListOfRequestQueues`

NewListOfRequestQueuesWithDefaults instantiates a new ListOfRequestQueues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListOfRequestQueues) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListOfRequestQueues) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListOfRequestQueues) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ListOfRequestQueues) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListOfRequestQueues) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListOfRequestQueues) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ListOfRequestQueues) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOfRequestQueues) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOfRequestQueues) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetDesc

`func (o *ListOfRequestQueues) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ListOfRequestQueues) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ListOfRequestQueues) SetDesc(v bool)`

SetDesc sets Desc field to given value.


### GetCount

`func (o *ListOfRequestQueues) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOfRequestQueues) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOfRequestQueues) SetCount(v int32)`

SetCount sets Count field to given value.


### GetUnnamed

`func (o *ListOfRequestQueues) GetUnnamed() bool`

GetUnnamed returns the Unnamed field if non-nil, zero value otherwise.

### GetUnnamedOk

`func (o *ListOfRequestQueues) GetUnnamedOk() (*bool, bool)`

GetUnnamedOk returns a tuple with the Unnamed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnamed

`func (o *ListOfRequestQueues) SetUnnamed(v bool)`

SetUnnamed sets Unnamed field to given value.

### HasUnnamed

`func (o *ListOfRequestQueues) HasUnnamed() bool`

HasUnnamed returns a boolean if a field has been set.

### GetItems

`func (o *ListOfRequestQueues) GetItems() []RequestQueueShort`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfRequestQueues) GetItemsOk() (*[]RequestQueueShort, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfRequestQueues) SetItems(v []RequestQueueShort)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


