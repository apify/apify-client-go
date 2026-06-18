# WebhookShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**IsAdHoc** | Pointer to **NullableBool** |  | [optional] 
**IsApifyIntegration** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**ActionType** | Pointer to **string** |  | [optional] 
**ShouldInterpolateStrings** | Pointer to **NullableBool** |  | [optional] 
**EventTypes** | [**[]WebhookEventType**](WebhookEventType.md) |  | 
**Condition** | [**WebhookCondition**](WebhookCondition.md) |  | 
**IgnoreSslErrors** | **bool** |  | 
**DoNotRetry** | **bool** |  | 
**RequestUrl** | **string** |  | 
**LastDispatch** | Pointer to [**NullableExampleWebhookDispatch**](ExampleWebhookDispatch.md) |  | [optional] 
**Stats** | Pointer to [**NullableWebhookStats**](WebhookStats.md) |  | [optional] 

## Methods

### NewWebhookShort

`func NewWebhookShort(id string, createdAt time.Time, modifiedAt time.Time, userId string, eventTypes []WebhookEventType, condition WebhookCondition, ignoreSslErrors bool, doNotRetry bool, requestUrl string, ) *WebhookShort`

NewWebhookShort instantiates a new WebhookShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookShortWithDefaults

`func NewWebhookShortWithDefaults() *WebhookShort`

NewWebhookShortWithDefaults instantiates a new WebhookShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookShort) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *WebhookShort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookShort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookShort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *WebhookShort) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *WebhookShort) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *WebhookShort) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetUserId

`func (o *WebhookShort) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookShort) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookShort) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsAdHoc

`func (o *WebhookShort) GetIsAdHoc() bool`

GetIsAdHoc returns the IsAdHoc field if non-nil, zero value otherwise.

### GetIsAdHocOk

`func (o *WebhookShort) GetIsAdHocOk() (*bool, bool)`

GetIsAdHocOk returns a tuple with the IsAdHoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdHoc

`func (o *WebhookShort) SetIsAdHoc(v bool)`

SetIsAdHoc sets IsAdHoc field to given value.

### HasIsAdHoc

`func (o *WebhookShort) HasIsAdHoc() bool`

HasIsAdHoc returns a boolean if a field has been set.

### SetIsAdHocNil

`func (o *WebhookShort) SetIsAdHocNil(b bool)`

 SetIsAdHocNil sets the value for IsAdHoc to be an explicit nil

### UnsetIsAdHoc
`func (o *WebhookShort) UnsetIsAdHoc()`

UnsetIsAdHoc ensures that no value is present for IsAdHoc, not even an explicit nil
### GetIsApifyIntegration

`func (o *WebhookShort) GetIsApifyIntegration() bool`

GetIsApifyIntegration returns the IsApifyIntegration field if non-nil, zero value otherwise.

### GetIsApifyIntegrationOk

`func (o *WebhookShort) GetIsApifyIntegrationOk() (*bool, bool)`

GetIsApifyIntegrationOk returns a tuple with the IsApifyIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApifyIntegration

`func (o *WebhookShort) SetIsApifyIntegration(v bool)`

SetIsApifyIntegration sets IsApifyIntegration field to given value.

### HasIsApifyIntegration

`func (o *WebhookShort) HasIsApifyIntegration() bool`

HasIsApifyIntegration returns a boolean if a field has been set.

### GetIsEnabled

`func (o *WebhookShort) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WebhookShort) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WebhookShort) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *WebhookShort) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetActionType

`func (o *WebhookShort) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WebhookShort) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WebhookShort) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WebhookShort) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetShouldInterpolateStrings

`func (o *WebhookShort) GetShouldInterpolateStrings() bool`

GetShouldInterpolateStrings returns the ShouldInterpolateStrings field if non-nil, zero value otherwise.

### GetShouldInterpolateStringsOk

`func (o *WebhookShort) GetShouldInterpolateStringsOk() (*bool, bool)`

GetShouldInterpolateStringsOk returns a tuple with the ShouldInterpolateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInterpolateStrings

`func (o *WebhookShort) SetShouldInterpolateStrings(v bool)`

SetShouldInterpolateStrings sets ShouldInterpolateStrings field to given value.

### HasShouldInterpolateStrings

`func (o *WebhookShort) HasShouldInterpolateStrings() bool`

HasShouldInterpolateStrings returns a boolean if a field has been set.

### SetShouldInterpolateStringsNil

`func (o *WebhookShort) SetShouldInterpolateStringsNil(b bool)`

 SetShouldInterpolateStringsNil sets the value for ShouldInterpolateStrings to be an explicit nil

### UnsetShouldInterpolateStrings
`func (o *WebhookShort) UnsetShouldInterpolateStrings()`

UnsetShouldInterpolateStrings ensures that no value is present for ShouldInterpolateStrings, not even an explicit nil
### GetEventTypes

`func (o *WebhookShort) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookShort) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookShort) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.


### GetCondition

`func (o *WebhookShort) GetCondition() WebhookCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WebhookShort) GetConditionOk() (*WebhookCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WebhookShort) SetCondition(v WebhookCondition)`

SetCondition sets Condition field to given value.


### GetIgnoreSslErrors

`func (o *WebhookShort) GetIgnoreSslErrors() bool`

GetIgnoreSslErrors returns the IgnoreSslErrors field if non-nil, zero value otherwise.

### GetIgnoreSslErrorsOk

`func (o *WebhookShort) GetIgnoreSslErrorsOk() (*bool, bool)`

GetIgnoreSslErrorsOk returns a tuple with the IgnoreSslErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSslErrors

`func (o *WebhookShort) SetIgnoreSslErrors(v bool)`

SetIgnoreSslErrors sets IgnoreSslErrors field to given value.


### GetDoNotRetry

`func (o *WebhookShort) GetDoNotRetry() bool`

GetDoNotRetry returns the DoNotRetry field if non-nil, zero value otherwise.

### GetDoNotRetryOk

`func (o *WebhookShort) GetDoNotRetryOk() (*bool, bool)`

GetDoNotRetryOk returns a tuple with the DoNotRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotRetry

`func (o *WebhookShort) SetDoNotRetry(v bool)`

SetDoNotRetry sets DoNotRetry field to given value.


### GetRequestUrl

`func (o *WebhookShort) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *WebhookShort) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *WebhookShort) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetLastDispatch

`func (o *WebhookShort) GetLastDispatch() ExampleWebhookDispatch`

GetLastDispatch returns the LastDispatch field if non-nil, zero value otherwise.

### GetLastDispatchOk

`func (o *WebhookShort) GetLastDispatchOk() (*ExampleWebhookDispatch, bool)`

GetLastDispatchOk returns a tuple with the LastDispatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDispatch

`func (o *WebhookShort) SetLastDispatch(v ExampleWebhookDispatch)`

SetLastDispatch sets LastDispatch field to given value.

### HasLastDispatch

`func (o *WebhookShort) HasLastDispatch() bool`

HasLastDispatch returns a boolean if a field has been set.

### SetLastDispatchNil

`func (o *WebhookShort) SetLastDispatchNil(b bool)`

 SetLastDispatchNil sets the value for LastDispatch to be an explicit nil

### UnsetLastDispatch
`func (o *WebhookShort) UnsetLastDispatch()`

UnsetLastDispatch ensures that no value is present for LastDispatch, not even an explicit nil
### GetStats

`func (o *WebhookShort) GetStats() WebhookStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *WebhookShort) GetStatsOk() (*WebhookStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *WebhookShort) SetStats(v WebhookStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *WebhookShort) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *WebhookShort) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *WebhookShort) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


