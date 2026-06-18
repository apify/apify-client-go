# InvalidItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemPosition** | Pointer to **int32** | The position of the invalid item in the array. | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | A complete list of AJV validation error objects for the invalid item. | [optional] 

## Methods

### NewInvalidItem

`func NewInvalidItem() *InvalidItem`

NewInvalidItem instantiates a new InvalidItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidItemWithDefaults

`func NewInvalidItemWithDefaults() *InvalidItem`

NewInvalidItemWithDefaults instantiates a new InvalidItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemPosition

`func (o *InvalidItem) GetItemPosition() int32`

GetItemPosition returns the ItemPosition field if non-nil, zero value otherwise.

### GetItemPositionOk

`func (o *InvalidItem) GetItemPositionOk() (*int32, bool)`

GetItemPositionOk returns a tuple with the ItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPosition

`func (o *InvalidItem) SetItemPosition(v int32)`

SetItemPosition sets ItemPosition field to given value.

### HasItemPosition

`func (o *InvalidItem) HasItemPosition() bool`

HasItemPosition returns a boolean if a field has been set.

### GetValidationErrors

`func (o *InvalidItem) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *InvalidItem) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *InvalidItem) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *InvalidItem) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


