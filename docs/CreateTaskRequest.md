# CreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**NullableTaskOptions**](TaskOptions.md) |  | [optional] 
**Input** | Pointer to **map[string]interface{}** | The input configuration for the Actor task. This is a user-defined JSON object that will be passed to the Actor when the task is run.  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ActorStandby** | Pointer to [**NullableActorStandby**](ActorStandby.md) |  | [optional] 

## Methods

### NewCreateTaskRequest

`func NewCreateTaskRequest(actId string, ) *CreateTaskRequest`

NewCreateTaskRequest instantiates a new CreateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskRequestWithDefaults

`func NewCreateTaskRequestWithDefaults() *CreateTaskRequest`

NewCreateTaskRequestWithDefaults instantiates a new CreateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActId

`func (o *CreateTaskRequest) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *CreateTaskRequest) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *CreateTaskRequest) SetActId(v string)`

SetActId sets ActId field to given value.


### GetName

`func (o *CreateTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTaskRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTaskRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *CreateTaskRequest) GetOptions() TaskOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateTaskRequest) GetOptionsOk() (*TaskOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateTaskRequest) SetOptions(v TaskOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateTaskRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CreateTaskRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CreateTaskRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetInput

`func (o *CreateTaskRequest) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CreateTaskRequest) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CreateTaskRequest) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *CreateTaskRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *CreateTaskRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *CreateTaskRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetTitle

`func (o *CreateTaskRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateTaskRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateTaskRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateTaskRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateTaskRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateTaskRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetActorStandby

`func (o *CreateTaskRequest) GetActorStandby() ActorStandby`

GetActorStandby returns the ActorStandby field if non-nil, zero value otherwise.

### GetActorStandbyOk

`func (o *CreateTaskRequest) GetActorStandbyOk() (*ActorStandby, bool)`

GetActorStandbyOk returns a tuple with the ActorStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorStandby

`func (o *CreateTaskRequest) SetActorStandby(v ActorStandby)`

SetActorStandby sets ActorStandby field to given value.

### HasActorStandby

`func (o *CreateTaskRequest) HasActorStandby() bool`

HasActorStandby returns a boolean if a field has been set.

### SetActorStandbyNil

`func (o *CreateTaskRequest) SetActorStandbyNil(b bool)`

 SetActorStandbyNil sets the value for ActorStandby to be an explicit nil

### UnsetActorStandby
`func (o *CreateTaskRequest) UnsetActorStandby()`

UnsetActorStandby ensures that no value is present for ActorStandby, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


