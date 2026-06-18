# ScheduleActionRunActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **interface{}** |  | 
**ActorId** | **string** |  | 
**RunInput** | Pointer to [**NullableScheduleActionRunInput**](ScheduleActionRunInput.md) |  | [optional] 
**RunOptions** | Pointer to [**NullableTaskOptions**](TaskOptions.md) |  | [optional] 

## Methods

### NewScheduleActionRunActor

`func NewScheduleActionRunActor(id string, type_ interface{}, actorId string, ) *ScheduleActionRunActor`

NewScheduleActionRunActor instantiates a new ScheduleActionRunActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleActionRunActorWithDefaults

`func NewScheduleActionRunActorWithDefaults() *ScheduleActionRunActor`

NewScheduleActionRunActorWithDefaults instantiates a new ScheduleActionRunActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleActionRunActor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleActionRunActor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleActionRunActor) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ScheduleActionRunActor) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleActionRunActor) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleActionRunActor) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ScheduleActionRunActor) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ScheduleActionRunActor) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetActorId

`func (o *ScheduleActionRunActor) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *ScheduleActionRunActor) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *ScheduleActionRunActor) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetRunInput

`func (o *ScheduleActionRunActor) GetRunInput() ScheduleActionRunInput`

GetRunInput returns the RunInput field if non-nil, zero value otherwise.

### GetRunInputOk

`func (o *ScheduleActionRunActor) GetRunInputOk() (*ScheduleActionRunInput, bool)`

GetRunInputOk returns a tuple with the RunInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunInput

`func (o *ScheduleActionRunActor) SetRunInput(v ScheduleActionRunInput)`

SetRunInput sets RunInput field to given value.

### HasRunInput

`func (o *ScheduleActionRunActor) HasRunInput() bool`

HasRunInput returns a boolean if a field has been set.

### SetRunInputNil

`func (o *ScheduleActionRunActor) SetRunInputNil(b bool)`

 SetRunInputNil sets the value for RunInput to be an explicit nil

### UnsetRunInput
`func (o *ScheduleActionRunActor) UnsetRunInput()`

UnsetRunInput ensures that no value is present for RunInput, not even an explicit nil
### GetRunOptions

`func (o *ScheduleActionRunActor) GetRunOptions() TaskOptions`

GetRunOptions returns the RunOptions field if non-nil, zero value otherwise.

### GetRunOptionsOk

`func (o *ScheduleActionRunActor) GetRunOptionsOk() (*TaskOptions, bool)`

GetRunOptionsOk returns a tuple with the RunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOptions

`func (o *ScheduleActionRunActor) SetRunOptions(v TaskOptions)`

SetRunOptions sets RunOptions field to given value.

### HasRunOptions

`func (o *ScheduleActionRunActor) HasRunOptions() bool`

HasRunOptions returns a boolean if a field has been set.

### SetRunOptionsNil

`func (o *ScheduleActionRunActor) SetRunOptionsNil(b bool)`

 SetRunOptionsNil sets the value for RunOptions to be an explicit nil

### UnsetRunOptions
`func (o *ScheduleActionRunActor) UnsetRunOptions()`

UnsetRunOptions ensures that no value is present for RunOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


