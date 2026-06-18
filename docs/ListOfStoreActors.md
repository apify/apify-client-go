# ListOfStoreActors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total number of items available across all pages. | 
**Offset** | **int32** | The starting position for this page of results. | 
**Limit** | **int32** | The maximum number of items returned per page. | 
**Desc** | **bool** | Whether the results are sorted in descending order. | 
**Count** | **int32** | The number of items returned in this response. | 
**Items** | [**[]StoreListActor**](StoreListActor.md) |  | 

## Methods

### NewListOfStoreActors

`func NewListOfStoreActors(total int32, offset int32, limit int32, desc bool, count int32, items []StoreListActor, ) *ListOfStoreActors`

NewListOfStoreActors instantiates a new ListOfStoreActors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfStoreActorsWithDefaults

`func NewListOfStoreActorsWithDefaults() *ListOfStoreActors`

NewListOfStoreActorsWithDefaults instantiates a new ListOfStoreActors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListOfStoreActors) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListOfStoreActors) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListOfStoreActors) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ListOfStoreActors) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListOfStoreActors) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListOfStoreActors) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ListOfStoreActors) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOfStoreActors) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOfStoreActors) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetDesc

`func (o *ListOfStoreActors) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ListOfStoreActors) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ListOfStoreActors) SetDesc(v bool)`

SetDesc sets Desc field to given value.


### GetCount

`func (o *ListOfStoreActors) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOfStoreActors) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOfStoreActors) SetCount(v int32)`

SetCount sets Count field to given value.


### GetItems

`func (o *ListOfStoreActors) GetItems() []StoreListActor`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfStoreActors) GetItemsOk() (*[]StoreListActor, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfStoreActors) SetItems(v []StoreListActor)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


