# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**ActId** | **string** |  | 
**Name** | **string** |  | 
**Username** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**RemovedAt** | Pointer to **NullableTime** |  | [optional] 
**Stats** | Pointer to [**NullableTaskStats**](TaskStats.md) |  | [optional] 
**Options** | Pointer to [**NullableTaskOptions**](TaskOptions.md) |  | [optional] 
**Input** | Pointer to **map[string]interface{}** | The input configuration for the Actor task. This is a user-defined JSON object that will be passed to the Actor when the task is run.  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ActorStandby** | Pointer to [**NullableActorStandby**](ActorStandby.md) |  | [optional] 
**StandbyUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTask

`func NewTask(id string, userId string, actId string, name string, createdAt time.Time, modifiedAt time.Time, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *Task) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Task) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Task) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetActId

`func (o *Task) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *Task) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *Task) SetActId(v string)`

SetActId sets ActId field to given value.


### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *Task) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Task) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Task) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Task) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *Task) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *Task) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Task) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Task) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Task) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetRemovedAt

`func (o *Task) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *Task) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *Task) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.

### HasRemovedAt

`func (o *Task) HasRemovedAt() bool`

HasRemovedAt returns a boolean if a field has been set.

### SetRemovedAtNil

`func (o *Task) SetRemovedAtNil(b bool)`

 SetRemovedAtNil sets the value for RemovedAt to be an explicit nil

### UnsetRemovedAt
`func (o *Task) UnsetRemovedAt()`

UnsetRemovedAt ensures that no value is present for RemovedAt, not even an explicit nil
### GetStats

`func (o *Task) GetStats() TaskStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Task) GetStatsOk() (*TaskStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Task) SetStats(v TaskStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Task) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *Task) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *Task) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetOptions

`func (o *Task) GetOptions() TaskOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Task) GetOptionsOk() (*TaskOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Task) SetOptions(v TaskOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Task) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *Task) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *Task) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetInput

`func (o *Task) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Task) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Task) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Task) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *Task) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *Task) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetTitle

`func (o *Task) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Task) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Task) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Task) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Task) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Task) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetActorStandby

`func (o *Task) GetActorStandby() ActorStandby`

GetActorStandby returns the ActorStandby field if non-nil, zero value otherwise.

### GetActorStandbyOk

`func (o *Task) GetActorStandbyOk() (*ActorStandby, bool)`

GetActorStandbyOk returns a tuple with the ActorStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorStandby

`func (o *Task) SetActorStandby(v ActorStandby)`

SetActorStandby sets ActorStandby field to given value.

### HasActorStandby

`func (o *Task) HasActorStandby() bool`

HasActorStandby returns a boolean if a field has been set.

### SetActorStandbyNil

`func (o *Task) SetActorStandbyNil(b bool)`

 SetActorStandbyNil sets the value for ActorStandby to be an explicit nil

### UnsetActorStandby
`func (o *Task) UnsetActorStandby()`

UnsetActorStandby ensures that no value is present for ActorStandby, not even an explicit nil
### GetStandbyUrl

`func (o *Task) GetStandbyUrl() string`

GetStandbyUrl returns the StandbyUrl field if non-nil, zero value otherwise.

### GetStandbyUrlOk

`func (o *Task) GetStandbyUrlOk() (*string, bool)`

GetStandbyUrlOk returns a tuple with the StandbyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyUrl

`func (o *Task) SetStandbyUrl(v string)`

SetStandbyUrl sets StandbyUrl field to given value.

### HasStandbyUrl

`func (o *Task) HasStandbyUrl() bool`

HasStandbyUrl returns a boolean if a field has been set.

### SetStandbyUrlNil

`func (o *Task) SetStandbyUrlNil(b bool)`

 SetStandbyUrlNil sets the value for StandbyUrl to be an explicit nil

### UnsetStandbyUrl
`func (o *Task) UnsetStandbyUrl()`

UnsetStandbyUrl ensures that no value is present for StandbyUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


