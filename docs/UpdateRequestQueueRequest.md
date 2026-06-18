# UpdateRequestQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The new name for the request queue. | [optional] 
**GeneralAccess** | Pointer to [**GeneralAccess**](GeneralAccess.md) |  | [optional] 

## Methods

### NewUpdateRequestQueueRequest

`func NewUpdateRequestQueueRequest() *UpdateRequestQueueRequest`

NewUpdateRequestQueueRequest instantiates a new UpdateRequestQueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequestQueueRequestWithDefaults

`func NewUpdateRequestQueueRequestWithDefaults() *UpdateRequestQueueRequest`

NewUpdateRequestQueueRequestWithDefaults instantiates a new UpdateRequestQueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRequestQueueRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRequestQueueRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRequestQueueRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRequestQueueRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateRequestQueueRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateRequestQueueRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGeneralAccess

`func (o *UpdateRequestQueueRequest) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *UpdateRequestQueueRequest) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *UpdateRequestQueueRequest) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.

### HasGeneralAccess

`func (o *UpdateRequestQueueRequest) HasGeneralAccess() bool`

HasGeneralAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


