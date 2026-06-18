# ListOfKeyValueStores

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total number of items available across all pages. | 
**Offset** | **int32** | The starting position for this page of results. | 
**Limit** | **int32** | The maximum number of items returned per page. | 
**Desc** | **bool** | Whether the results are sorted in descending order. | 
**Count** | **int32** | The number of items returned in this response. | 
**Unnamed** | Pointer to **bool** | Whether the listing was filtered to only unnamed key-value stores. | [optional] 
**Items** | [**[]KeyValueStore**](KeyValueStore.md) |  | 

## Methods

### NewListOfKeyValueStores

`func NewListOfKeyValueStores(total int32, offset int32, limit int32, desc bool, count int32, items []KeyValueStore, ) *ListOfKeyValueStores`

NewListOfKeyValueStores instantiates a new ListOfKeyValueStores object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfKeyValueStoresWithDefaults

`func NewListOfKeyValueStoresWithDefaults() *ListOfKeyValueStores`

NewListOfKeyValueStoresWithDefaults instantiates a new ListOfKeyValueStores object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListOfKeyValueStores) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListOfKeyValueStores) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListOfKeyValueStores) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ListOfKeyValueStores) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListOfKeyValueStores) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListOfKeyValueStores) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ListOfKeyValueStores) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOfKeyValueStores) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOfKeyValueStores) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetDesc

`func (o *ListOfKeyValueStores) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ListOfKeyValueStores) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ListOfKeyValueStores) SetDesc(v bool)`

SetDesc sets Desc field to given value.


### GetCount

`func (o *ListOfKeyValueStores) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOfKeyValueStores) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOfKeyValueStores) SetCount(v int32)`

SetCount sets Count field to given value.


### GetUnnamed

`func (o *ListOfKeyValueStores) GetUnnamed() bool`

GetUnnamed returns the Unnamed field if non-nil, zero value otherwise.

### GetUnnamedOk

`func (o *ListOfKeyValueStores) GetUnnamedOk() (*bool, bool)`

GetUnnamedOk returns a tuple with the Unnamed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnamed

`func (o *ListOfKeyValueStores) SetUnnamed(v bool)`

SetUnnamed sets Unnamed field to given value.

### HasUnnamed

`func (o *ListOfKeyValueStores) HasUnnamed() bool`

HasUnnamed returns a boolean if a field has been set.

### GetItems

`func (o *ListOfKeyValueStores) GetItems() []KeyValueStore`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfKeyValueStores) GetItemsOk() (*[]KeyValueStore, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfKeyValueStores) SetItems(v []KeyValueStore)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


