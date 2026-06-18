# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**IsAdHoc** | Pointer to **NullableBool** |  | [optional] 
**ShouldInterpolateStrings** | Pointer to **NullableBool** |  | [optional] 
**EventTypes** | [**[]WebhookEventType**](WebhookEventType.md) |  | 
**Condition** | [**WebhookCondition**](WebhookCondition.md) |  | 
**IgnoreSslErrors** | **bool** |  | 
**DoNotRetry** | Pointer to **NullableBool** |  | [optional] 
**RequestUrl** | **string** |  | 
**PayloadTemplate** | Pointer to **NullableString** |  | [optional] 
**HeadersTemplate** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**LastDispatch** | Pointer to [**NullableExampleWebhookDispatch**](ExampleWebhookDispatch.md) |  | [optional] 
**Stats** | Pointer to [**NullableWebhookStats**](WebhookStats.md) |  | [optional] 

## Methods

### NewWebhook

`func NewWebhook(id string, createdAt time.Time, modifiedAt time.Time, userId string, eventTypes []WebhookEventType, condition WebhookCondition, ignoreSslErrors bool, requestUrl string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Webhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Webhook) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Webhook) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Webhook) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetUserId

`func (o *Webhook) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Webhook) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Webhook) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsAdHoc

`func (o *Webhook) GetIsAdHoc() bool`

GetIsAdHoc returns the IsAdHoc field if non-nil, zero value otherwise.

### GetIsAdHocOk

`func (o *Webhook) GetIsAdHocOk() (*bool, bool)`

GetIsAdHocOk returns a tuple with the IsAdHoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdHoc

`func (o *Webhook) SetIsAdHoc(v bool)`

SetIsAdHoc sets IsAdHoc field to given value.

### HasIsAdHoc

`func (o *Webhook) HasIsAdHoc() bool`

HasIsAdHoc returns a boolean if a field has been set.

### SetIsAdHocNil

`func (o *Webhook) SetIsAdHocNil(b bool)`

 SetIsAdHocNil sets the value for IsAdHoc to be an explicit nil

### UnsetIsAdHoc
`func (o *Webhook) UnsetIsAdHoc()`

UnsetIsAdHoc ensures that no value is present for IsAdHoc, not even an explicit nil
### GetShouldInterpolateStrings

`func (o *Webhook) GetShouldInterpolateStrings() bool`

GetShouldInterpolateStrings returns the ShouldInterpolateStrings field if non-nil, zero value otherwise.

### GetShouldInterpolateStringsOk

`func (o *Webhook) GetShouldInterpolateStringsOk() (*bool, bool)`

GetShouldInterpolateStringsOk returns a tuple with the ShouldInterpolateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInterpolateStrings

`func (o *Webhook) SetShouldInterpolateStrings(v bool)`

SetShouldInterpolateStrings sets ShouldInterpolateStrings field to given value.

### HasShouldInterpolateStrings

`func (o *Webhook) HasShouldInterpolateStrings() bool`

HasShouldInterpolateStrings returns a boolean if a field has been set.

### SetShouldInterpolateStringsNil

`func (o *Webhook) SetShouldInterpolateStringsNil(b bool)`

 SetShouldInterpolateStringsNil sets the value for ShouldInterpolateStrings to be an explicit nil

### UnsetShouldInterpolateStrings
`func (o *Webhook) UnsetShouldInterpolateStrings()`

UnsetShouldInterpolateStrings ensures that no value is present for ShouldInterpolateStrings, not even an explicit nil
### GetEventTypes

`func (o *Webhook) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *Webhook) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *Webhook) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.


### GetCondition

`func (o *Webhook) GetCondition() WebhookCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Webhook) GetConditionOk() (*WebhookCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Webhook) SetCondition(v WebhookCondition)`

SetCondition sets Condition field to given value.


### GetIgnoreSslErrors

`func (o *Webhook) GetIgnoreSslErrors() bool`

GetIgnoreSslErrors returns the IgnoreSslErrors field if non-nil, zero value otherwise.

### GetIgnoreSslErrorsOk

`func (o *Webhook) GetIgnoreSslErrorsOk() (*bool, bool)`

GetIgnoreSslErrorsOk returns a tuple with the IgnoreSslErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSslErrors

`func (o *Webhook) SetIgnoreSslErrors(v bool)`

SetIgnoreSslErrors sets IgnoreSslErrors field to given value.


### GetDoNotRetry

`func (o *Webhook) GetDoNotRetry() bool`

GetDoNotRetry returns the DoNotRetry field if non-nil, zero value otherwise.

### GetDoNotRetryOk

`func (o *Webhook) GetDoNotRetryOk() (*bool, bool)`

GetDoNotRetryOk returns a tuple with the DoNotRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotRetry

`func (o *Webhook) SetDoNotRetry(v bool)`

SetDoNotRetry sets DoNotRetry field to given value.

### HasDoNotRetry

`func (o *Webhook) HasDoNotRetry() bool`

HasDoNotRetry returns a boolean if a field has been set.

### SetDoNotRetryNil

`func (o *Webhook) SetDoNotRetryNil(b bool)`

 SetDoNotRetryNil sets the value for DoNotRetry to be an explicit nil

### UnsetDoNotRetry
`func (o *Webhook) UnsetDoNotRetry()`

UnsetDoNotRetry ensures that no value is present for DoNotRetry, not even an explicit nil
### GetRequestUrl

`func (o *Webhook) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *Webhook) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *Webhook) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetPayloadTemplate

`func (o *Webhook) GetPayloadTemplate() string`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *Webhook) GetPayloadTemplateOk() (*string, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *Webhook) SetPayloadTemplate(v string)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *Webhook) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### SetPayloadTemplateNil

`func (o *Webhook) SetPayloadTemplateNil(b bool)`

 SetPayloadTemplateNil sets the value for PayloadTemplate to be an explicit nil

### UnsetPayloadTemplate
`func (o *Webhook) UnsetPayloadTemplate()`

UnsetPayloadTemplate ensures that no value is present for PayloadTemplate, not even an explicit nil
### GetHeadersTemplate

`func (o *Webhook) GetHeadersTemplate() string`

GetHeadersTemplate returns the HeadersTemplate field if non-nil, zero value otherwise.

### GetHeadersTemplateOk

`func (o *Webhook) GetHeadersTemplateOk() (*string, bool)`

GetHeadersTemplateOk returns a tuple with the HeadersTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersTemplate

`func (o *Webhook) SetHeadersTemplate(v string)`

SetHeadersTemplate sets HeadersTemplate field to given value.

### HasHeadersTemplate

`func (o *Webhook) HasHeadersTemplate() bool`

HasHeadersTemplate returns a boolean if a field has been set.

### SetHeadersTemplateNil

`func (o *Webhook) SetHeadersTemplateNil(b bool)`

 SetHeadersTemplateNil sets the value for HeadersTemplate to be an explicit nil

### UnsetHeadersTemplate
`func (o *Webhook) UnsetHeadersTemplate()`

UnsetHeadersTemplate ensures that no value is present for HeadersTemplate, not even an explicit nil
### GetDescription

`func (o *Webhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Webhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Webhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Webhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Webhook) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Webhook) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLastDispatch

`func (o *Webhook) GetLastDispatch() ExampleWebhookDispatch`

GetLastDispatch returns the LastDispatch field if non-nil, zero value otherwise.

### GetLastDispatchOk

`func (o *Webhook) GetLastDispatchOk() (*ExampleWebhookDispatch, bool)`

GetLastDispatchOk returns a tuple with the LastDispatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDispatch

`func (o *Webhook) SetLastDispatch(v ExampleWebhookDispatch)`

SetLastDispatch sets LastDispatch field to given value.

### HasLastDispatch

`func (o *Webhook) HasLastDispatch() bool`

HasLastDispatch returns a boolean if a field has been set.

### SetLastDispatchNil

`func (o *Webhook) SetLastDispatchNil(b bool)`

 SetLastDispatchNil sets the value for LastDispatch to be an explicit nil

### UnsetLastDispatch
`func (o *Webhook) UnsetLastDispatch()`

UnsetLastDispatch ensures that no value is present for LastDispatch, not even an explicit nil
### GetStats

`func (o *Webhook) GetStats() WebhookStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Webhook) GetStatsOk() (*WebhookStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Webhook) SetStats(v WebhookStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Webhook) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *Webhook) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *Webhook) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


