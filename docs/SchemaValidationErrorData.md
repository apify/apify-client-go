# SchemaValidationErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidItems** | [**[]InvalidItem**](InvalidItem.md) | A list of invalid items in the received array of items. | 

## Methods

### NewSchemaValidationErrorData

`func NewSchemaValidationErrorData(invalidItems []InvalidItem, ) *SchemaValidationErrorData`

NewSchemaValidationErrorData instantiates a new SchemaValidationErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaValidationErrorDataWithDefaults

`func NewSchemaValidationErrorDataWithDefaults() *SchemaValidationErrorData`

NewSchemaValidationErrorDataWithDefaults instantiates a new SchemaValidationErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidItems

`func (o *SchemaValidationErrorData) GetInvalidItems() []InvalidItem`

GetInvalidItems returns the InvalidItems field if non-nil, zero value otherwise.

### GetInvalidItemsOk

`func (o *SchemaValidationErrorData) GetInvalidItemsOk() (*[]InvalidItem, bool)`

GetInvalidItemsOk returns a tuple with the InvalidItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidItems

`func (o *SchemaValidationErrorData) SetInvalidItems(v []InvalidItem)`

SetInvalidItems sets InvalidItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


