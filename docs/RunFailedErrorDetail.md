# RunFailedErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** |  | [optional] 
**Message** | Pointer to **string** | Human-readable error message describing what went wrong. | [optional] 

## Methods

### NewRunFailedErrorDetail

`func NewRunFailedErrorDetail() *RunFailedErrorDetail`

NewRunFailedErrorDetail instantiates a new RunFailedErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFailedErrorDetailWithDefaults

`func NewRunFailedErrorDetailWithDefaults() *RunFailedErrorDetail`

NewRunFailedErrorDetailWithDefaults instantiates a new RunFailedErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RunFailedErrorDetail) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunFailedErrorDetail) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunFailedErrorDetail) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *RunFailedErrorDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RunFailedErrorDetail) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RunFailedErrorDetail) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *RunFailedErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RunFailedErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RunFailedErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RunFailedErrorDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


