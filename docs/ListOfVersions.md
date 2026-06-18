# ListOfVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Items** | [**[]Version**](Version.md) |  | 

## Methods

### NewListOfVersions

`func NewListOfVersions(total int32, items []Version, ) *ListOfVersions`

NewListOfVersions instantiates a new ListOfVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfVersionsWithDefaults

`func NewListOfVersionsWithDefaults() *ListOfVersions`

NewListOfVersionsWithDefaults instantiates a new ListOfVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListOfVersions) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListOfVersions) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListOfVersions) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListOfVersions) GetItems() []Version`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfVersions) GetItemsOk() (*[]Version, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfVersions) SetItems(v []Version)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


