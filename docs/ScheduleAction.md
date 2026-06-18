# ScheduleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **interface{}** |  | 
**ActorId** | **string** |  | 
**RunInput** | Pointer to [**ScheduleActionRunInput**](ScheduleActionRunInput.md) |  | [optional] 
**RunOptions** | Pointer to [**TaskOptions**](TaskOptions.md) |  | [optional] 
**ActorTaskId** | **string** |  | 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewScheduleAction

`func NewScheduleAction(id string, type_ interface{}, actorId string, actorTaskId string, ) *ScheduleAction`

NewScheduleAction instantiates a new ScheduleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleActionWithDefaults

`func NewScheduleActionWithDefaults() *ScheduleAction`

NewScheduleActionWithDefaults instantiates a new ScheduleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleAction) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ScheduleAction) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleAction) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleAction) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ScheduleAction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ScheduleAction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetActorId

`func (o *ScheduleAction) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *ScheduleAction) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *ScheduleAction) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetRunInput

`func (o *ScheduleAction) GetRunInput() ScheduleActionRunInput`

GetRunInput returns the RunInput field if non-nil, zero value otherwise.

### GetRunInputOk

`func (o *ScheduleAction) GetRunInputOk() (*ScheduleActionRunInput, bool)`

GetRunInputOk returns a tuple with the RunInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunInput

`func (o *ScheduleAction) SetRunInput(v ScheduleActionRunInput)`

SetRunInput sets RunInput field to given value.

### HasRunInput

`func (o *ScheduleAction) HasRunInput() bool`

HasRunInput returns a boolean if a field has been set.

### GetRunOptions

`func (o *ScheduleAction) GetRunOptions() TaskOptions`

GetRunOptions returns the RunOptions field if non-nil, zero value otherwise.

### GetRunOptionsOk

`func (o *ScheduleAction) GetRunOptionsOk() (*TaskOptions, bool)`

GetRunOptionsOk returns a tuple with the RunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOptions

`func (o *ScheduleAction) SetRunOptions(v TaskOptions)`

SetRunOptions sets RunOptions field to given value.

### HasRunOptions

`func (o *ScheduleAction) HasRunOptions() bool`

HasRunOptions returns a boolean if a field has been set.

### GetActorTaskId

`func (o *ScheduleAction) GetActorTaskId() string`

GetActorTaskId returns the ActorTaskId field if non-nil, zero value otherwise.

### GetActorTaskIdOk

`func (o *ScheduleAction) GetActorTaskIdOk() (*string, bool)`

GetActorTaskIdOk returns a tuple with the ActorTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTaskId

`func (o *ScheduleAction) SetActorTaskId(v string)`

SetActorTaskId sets ActorTaskId field to given value.


### GetInput

`func (o *ScheduleAction) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ScheduleAction) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ScheduleAction) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ScheduleAction) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


