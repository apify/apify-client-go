# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstancePath** | Pointer to **string** | The path to the instance being validated. | [optional] 
**SchemaPath** | Pointer to **string** | The path to the schema that failed the validation. | [optional] 
**Keyword** | Pointer to **string** | The validation keyword that caused the error. | [optional] 
**Message** | Pointer to **string** | A message describing the validation error. | [optional] 
**Params** | Pointer to **map[string]interface{}** | Additional parameters specific to the validation error. | [optional] 

## Methods

### NewValidationError

`func NewValidationError() *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstancePath

`func (o *ValidationError) GetInstancePath() string`

GetInstancePath returns the InstancePath field if non-nil, zero value otherwise.

### GetInstancePathOk

`func (o *ValidationError) GetInstancePathOk() (*string, bool)`

GetInstancePathOk returns a tuple with the InstancePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePath

`func (o *ValidationError) SetInstancePath(v string)`

SetInstancePath sets InstancePath field to given value.

### HasInstancePath

`func (o *ValidationError) HasInstancePath() bool`

HasInstancePath returns a boolean if a field has been set.

### GetSchemaPath

`func (o *ValidationError) GetSchemaPath() string`

GetSchemaPath returns the SchemaPath field if non-nil, zero value otherwise.

### GetSchemaPathOk

`func (o *ValidationError) GetSchemaPathOk() (*string, bool)`

GetSchemaPathOk returns a tuple with the SchemaPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaPath

`func (o *ValidationError) SetSchemaPath(v string)`

SetSchemaPath sets SchemaPath field to given value.

### HasSchemaPath

`func (o *ValidationError) HasSchemaPath() bool`

HasSchemaPath returns a boolean if a field has been set.

### GetKeyword

`func (o *ValidationError) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *ValidationError) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *ValidationError) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *ValidationError) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetMessage

`func (o *ValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParams

`func (o *ValidationError) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ValidationError) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ValidationError) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ValidationError) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


