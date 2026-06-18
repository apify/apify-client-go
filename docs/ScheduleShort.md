# ScheduleShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**CronExpression** | **string** |  | 
**Timezone** | **string** |  | 
**IsEnabled** | **bool** |  | 
**IsExclusive** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**NextRunAt** | Pointer to **time.Time** |  | [optional] 
**LastRunAt** | Pointer to **time.Time** |  | [optional] 
**Actions** | [**[]ScheduleActionShort**](ScheduleActionShort.md) |  | 

## Methods

### NewScheduleShort

`func NewScheduleShort(id string, userId string, name string, cronExpression string, timezone string, isEnabled bool, isExclusive bool, createdAt time.Time, modifiedAt time.Time, actions []ScheduleActionShort, ) *ScheduleShort`

NewScheduleShort instantiates a new ScheduleShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleShortWithDefaults

`func NewScheduleShortWithDefaults() *ScheduleShort`

NewScheduleShortWithDefaults instantiates a new ScheduleShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleShort) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ScheduleShort) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ScheduleShort) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ScheduleShort) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *ScheduleShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleShort) SetName(v string)`

SetName sets Name field to given value.


### GetCronExpression

`func (o *ScheduleShort) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ScheduleShort) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ScheduleShort) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetTimezone

`func (o *ScheduleShort) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScheduleShort) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScheduleShort) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetIsEnabled

`func (o *ScheduleShort) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScheduleShort) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScheduleShort) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsExclusive

`func (o *ScheduleShort) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *ScheduleShort) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *ScheduleShort) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.


### GetCreatedAt

`func (o *ScheduleShort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduleShort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduleShort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ScheduleShort) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ScheduleShort) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ScheduleShort) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetNextRunAt

`func (o *ScheduleShort) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *ScheduleShort) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *ScheduleShort) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *ScheduleShort) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetLastRunAt

`func (o *ScheduleShort) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *ScheduleShort) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *ScheduleShort) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *ScheduleShort) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### GetActions

`func (o *ScheduleShort) GetActions() []ScheduleActionShort`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ScheduleShort) GetActionsOk() (*[]ScheduleActionShort, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ScheduleShort) SetActions(v []ScheduleActionShort)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


