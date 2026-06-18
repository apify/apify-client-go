# ListOfDatasets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total number of items available across all pages. | 
**Offset** | **int32** | The starting position for this page of results. | 
**Limit** | **int32** | The maximum number of items returned per page. | 
**Desc** | **bool** | Whether the results are sorted in descending order. | 
**Count** | **int32** | The number of items returned in this response. | 
**Unnamed** | Pointer to **bool** | Whether the listing was filtered to only unnamed datasets. | [optional] 
**Items** | [**[]DatasetListItem**](DatasetListItem.md) |  | 

## Methods

### NewListOfDatasets

`func NewListOfDatasets(total int32, offset int32, limit int32, desc bool, count int32, items []DatasetListItem, ) *ListOfDatasets`

NewListOfDatasets instantiates a new ListOfDatasets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfDatasetsWithDefaults

`func NewListOfDatasetsWithDefaults() *ListOfDatasets`

NewListOfDatasetsWithDefaults instantiates a new ListOfDatasets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListOfDatasets) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListOfDatasets) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListOfDatasets) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ListOfDatasets) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListOfDatasets) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListOfDatasets) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ListOfDatasets) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOfDatasets) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOfDatasets) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetDesc

`func (o *ListOfDatasets) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ListOfDatasets) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ListOfDatasets) SetDesc(v bool)`

SetDesc sets Desc field to given value.


### GetCount

`func (o *ListOfDatasets) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOfDatasets) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOfDatasets) SetCount(v int32)`

SetCount sets Count field to given value.


### GetUnnamed

`func (o *ListOfDatasets) GetUnnamed() bool`

GetUnnamed returns the Unnamed field if non-nil, zero value otherwise.

### GetUnnamedOk

`func (o *ListOfDatasets) GetUnnamedOk() (*bool, bool)`

GetUnnamedOk returns a tuple with the Unnamed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnnamed

`func (o *ListOfDatasets) SetUnnamed(v bool)`

SetUnnamed sets Unnamed field to given value.

### HasUnnamed

`func (o *ListOfDatasets) HasUnnamed() bool`

HasUnnamed returns a boolean if a field has been set.

### GetItems

`func (o *ListOfDatasets) GetItems() []DatasetListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfDatasets) GetItemsOk() (*[]DatasetListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfDatasets) SetItems(v []DatasetListItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


