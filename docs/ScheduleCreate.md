# ScheduleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**IsExclusive** | Pointer to **NullableBool** |  | [optional] 
**CronExpression** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Actions** | Pointer to [**[]ScheduleCreateAction**](ScheduleCreateAction.md) |  | [optional] 

## Methods

### NewScheduleCreate

`func NewScheduleCreate() *ScheduleCreate`

NewScheduleCreate instantiates a new ScheduleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCreateWithDefaults

`func NewScheduleCreateWithDefaults() *ScheduleCreate`

NewScheduleCreateWithDefaults instantiates a new ScheduleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScheduleCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ScheduleCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ScheduleCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsEnabled

`func (o *ScheduleCreate) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScheduleCreate) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScheduleCreate) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ScheduleCreate) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *ScheduleCreate) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ScheduleCreate) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIsExclusive

`func (o *ScheduleCreate) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *ScheduleCreate) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *ScheduleCreate) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *ScheduleCreate) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### SetIsExclusiveNil

`func (o *ScheduleCreate) SetIsExclusiveNil(b bool)`

 SetIsExclusiveNil sets the value for IsExclusive to be an explicit nil

### UnsetIsExclusive
`func (o *ScheduleCreate) UnsetIsExclusive()`

UnsetIsExclusive ensures that no value is present for IsExclusive, not even an explicit nil
### GetCronExpression

`func (o *ScheduleCreate) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ScheduleCreate) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ScheduleCreate) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *ScheduleCreate) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *ScheduleCreate) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *ScheduleCreate) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetTimezone

`func (o *ScheduleCreate) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScheduleCreate) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScheduleCreate) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ScheduleCreate) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *ScheduleCreate) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *ScheduleCreate) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetDescription

`func (o *ScheduleCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScheduleCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScheduleCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTitle

`func (o *ScheduleCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ScheduleCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ScheduleCreate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ScheduleCreate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ScheduleCreate) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ScheduleCreate) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetActions

`func (o *ScheduleCreate) GetActions() []ScheduleCreateAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ScheduleCreate) GetActionsOk() (*[]ScheduleCreateAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ScheduleCreate) SetActions(v []ScheduleCreateAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ScheduleCreate) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *ScheduleCreate) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *ScheduleCreate) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


