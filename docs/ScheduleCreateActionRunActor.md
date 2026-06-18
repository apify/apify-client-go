# ScheduleCreateActionRunActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**ActorId** | **string** |  | 
**RunInput** | Pointer to [**NullableScheduleActionRunInput**](ScheduleActionRunInput.md) |  | [optional] 
**RunOptions** | Pointer to [**NullableTaskOptions**](TaskOptions.md) |  | [optional] 

## Methods

### NewScheduleCreateActionRunActor

`func NewScheduleCreateActionRunActor(type_ interface{}, actorId string, ) *ScheduleCreateActionRunActor`

NewScheduleCreateActionRunActor instantiates a new ScheduleCreateActionRunActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCreateActionRunActorWithDefaults

`func NewScheduleCreateActionRunActorWithDefaults() *ScheduleCreateActionRunActor`

NewScheduleCreateActionRunActorWithDefaults instantiates a new ScheduleCreateActionRunActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScheduleCreateActionRunActor) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleCreateActionRunActor) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleCreateActionRunActor) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ScheduleCreateActionRunActor) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ScheduleCreateActionRunActor) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetActorId

`func (o *ScheduleCreateActionRunActor) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *ScheduleCreateActionRunActor) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *ScheduleCreateActionRunActor) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetRunInput

`func (o *ScheduleCreateActionRunActor) GetRunInput() ScheduleActionRunInput`

GetRunInput returns the RunInput field if non-nil, zero value otherwise.

### GetRunInputOk

`func (o *ScheduleCreateActionRunActor) GetRunInputOk() (*ScheduleActionRunInput, bool)`

GetRunInputOk returns a tuple with the RunInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunInput

`func (o *ScheduleCreateActionRunActor) SetRunInput(v ScheduleActionRunInput)`

SetRunInput sets RunInput field to given value.

### HasRunInput

`func (o *ScheduleCreateActionRunActor) HasRunInput() bool`

HasRunInput returns a boolean if a field has been set.

### SetRunInputNil

`func (o *ScheduleCreateActionRunActor) SetRunInputNil(b bool)`

 SetRunInputNil sets the value for RunInput to be an explicit nil

### UnsetRunInput
`func (o *ScheduleCreateActionRunActor) UnsetRunInput()`

UnsetRunInput ensures that no value is present for RunInput, not even an explicit nil
### GetRunOptions

`func (o *ScheduleCreateActionRunActor) GetRunOptions() TaskOptions`

GetRunOptions returns the RunOptions field if non-nil, zero value otherwise.

### GetRunOptionsOk

`func (o *ScheduleCreateActionRunActor) GetRunOptionsOk() (*TaskOptions, bool)`

GetRunOptionsOk returns a tuple with the RunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOptions

`func (o *ScheduleCreateActionRunActor) SetRunOptions(v TaskOptions)`

SetRunOptions sets RunOptions field to given value.

### HasRunOptions

`func (o *ScheduleCreateActionRunActor) HasRunOptions() bool`

HasRunOptions returns a boolean if a field has been set.

### SetRunOptionsNil

`func (o *ScheduleCreateActionRunActor) SetRunOptionsNil(b bool)`

 SetRunOptionsNil sets the value for RunOptions to be an explicit nil

### UnsetRunOptions
`func (o *ScheduleCreateActionRunActor) UnsetRunOptions()`

UnsetRunOptions ensures that no value is present for RunOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


