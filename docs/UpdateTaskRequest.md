# UpdateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**NullableTaskOptions**](TaskOptions.md) |  | [optional] 
**Input** | Pointer to **map[string]interface{}** | The input configuration for the Actor task. This is a user-defined JSON object that will be passed to the Actor when the task is run.  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ActorStandby** | Pointer to [**NullableActorStandby**](ActorStandby.md) |  | [optional] 

## Methods

### NewUpdateTaskRequest

`func NewUpdateTaskRequest() *UpdateTaskRequest`

NewUpdateTaskRequest instantiates a new UpdateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskRequestWithDefaults

`func NewUpdateTaskRequestWithDefaults() *UpdateTaskRequest`

NewUpdateTaskRequestWithDefaults instantiates a new UpdateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTaskRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTaskRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateTaskRequest) GetOptions() TaskOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateTaskRequest) GetOptionsOk() (*TaskOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateTaskRequest) SetOptions(v TaskOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateTaskRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *UpdateTaskRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateTaskRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetInput

`func (o *UpdateTaskRequest) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *UpdateTaskRequest) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *UpdateTaskRequest) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *UpdateTaskRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *UpdateTaskRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *UpdateTaskRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetTitle

`func (o *UpdateTaskRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateTaskRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateTaskRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateTaskRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateTaskRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateTaskRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetActorStandby

`func (o *UpdateTaskRequest) GetActorStandby() ActorStandby`

GetActorStandby returns the ActorStandby field if non-nil, zero value otherwise.

### GetActorStandbyOk

`func (o *UpdateTaskRequest) GetActorStandbyOk() (*ActorStandby, bool)`

GetActorStandbyOk returns a tuple with the ActorStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorStandby

`func (o *UpdateTaskRequest) SetActorStandby(v ActorStandby)`

SetActorStandby sets ActorStandby field to given value.

### HasActorStandby

`func (o *UpdateTaskRequest) HasActorStandby() bool`

HasActorStandby returns a boolean if a field has been set.

### SetActorStandbyNil

`func (o *UpdateTaskRequest) SetActorStandbyNil(b bool)`

 SetActorStandbyNil sets the value for ActorStandby to be an explicit nil

### UnsetActorStandby
`func (o *UpdateTaskRequest) UnsetActorStandby()`

UnsetActorStandby ensures that no value is present for ActorStandby, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


