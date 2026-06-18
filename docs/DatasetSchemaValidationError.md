# DatasetSchemaValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the error. | [optional] 
**Message** | Pointer to **string** | A human-readable message describing the error. | [optional] 
**Data** | Pointer to [**SchemaValidationErrorData**](SchemaValidationErrorData.md) |  | [optional] 

## Methods

### NewDatasetSchemaValidationError

`func NewDatasetSchemaValidationError() *DatasetSchemaValidationError`

NewDatasetSchemaValidationError instantiates a new DatasetSchemaValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetSchemaValidationErrorWithDefaults

`func NewDatasetSchemaValidationErrorWithDefaults() *DatasetSchemaValidationError`

NewDatasetSchemaValidationErrorWithDefaults instantiates a new DatasetSchemaValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DatasetSchemaValidationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasetSchemaValidationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasetSchemaValidationError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatasetSchemaValidationError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *DatasetSchemaValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DatasetSchemaValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DatasetSchemaValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DatasetSchemaValidationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *DatasetSchemaValidationError) GetData() SchemaValidationErrorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatasetSchemaValidationError) GetDataOk() (*SchemaValidationErrorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatasetSchemaValidationError) SetData(v SchemaValidationErrorData)`

SetData sets Data field to given value.

### HasData

`func (o *DatasetSchemaValidationError) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


