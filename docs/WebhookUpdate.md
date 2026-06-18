# WebhookUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAdHoc** | Pointer to **NullableBool** |  | [optional] 
**EventTypes** | Pointer to [**[]WebhookEventType**](WebhookEventType.md) |  | [optional] 
**Condition** | Pointer to [**NullableWebhookCondition**](WebhookCondition.md) |  | [optional] 
**IgnoreSslErrors** | Pointer to **NullableBool** |  | [optional] 
**DoNotRetry** | Pointer to **NullableBool** |  | [optional] 
**RequestUrl** | Pointer to **NullableString** |  | [optional] 
**PayloadTemplate** | Pointer to **NullableString** |  | [optional] 
**HeadersTemplate** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShouldInterpolateStrings** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWebhookUpdate

`func NewWebhookUpdate() *WebhookUpdate`

NewWebhookUpdate instantiates a new WebhookUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateWithDefaults

`func NewWebhookUpdateWithDefaults() *WebhookUpdate`

NewWebhookUpdateWithDefaults instantiates a new WebhookUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdHoc

`func (o *WebhookUpdate) GetIsAdHoc() bool`

GetIsAdHoc returns the IsAdHoc field if non-nil, zero value otherwise.

### GetIsAdHocOk

`func (o *WebhookUpdate) GetIsAdHocOk() (*bool, bool)`

GetIsAdHocOk returns a tuple with the IsAdHoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdHoc

`func (o *WebhookUpdate) SetIsAdHoc(v bool)`

SetIsAdHoc sets IsAdHoc field to given value.

### HasIsAdHoc

`func (o *WebhookUpdate) HasIsAdHoc() bool`

HasIsAdHoc returns a boolean if a field has been set.

### SetIsAdHocNil

`func (o *WebhookUpdate) SetIsAdHocNil(b bool)`

 SetIsAdHocNil sets the value for IsAdHoc to be an explicit nil

### UnsetIsAdHoc
`func (o *WebhookUpdate) UnsetIsAdHoc()`

UnsetIsAdHoc ensures that no value is present for IsAdHoc, not even an explicit nil
### GetEventTypes

`func (o *WebhookUpdate) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookUpdate) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookUpdate) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *WebhookUpdate) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### SetEventTypesNil

`func (o *WebhookUpdate) SetEventTypesNil(b bool)`

 SetEventTypesNil sets the value for EventTypes to be an explicit nil

### UnsetEventTypes
`func (o *WebhookUpdate) UnsetEventTypes()`

UnsetEventTypes ensures that no value is present for EventTypes, not even an explicit nil
### GetCondition

`func (o *WebhookUpdate) GetCondition() WebhookCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WebhookUpdate) GetConditionOk() (*WebhookCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WebhookUpdate) SetCondition(v WebhookCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WebhookUpdate) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *WebhookUpdate) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *WebhookUpdate) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetIgnoreSslErrors

`func (o *WebhookUpdate) GetIgnoreSslErrors() bool`

GetIgnoreSslErrors returns the IgnoreSslErrors field if non-nil, zero value otherwise.

### GetIgnoreSslErrorsOk

`func (o *WebhookUpdate) GetIgnoreSslErrorsOk() (*bool, bool)`

GetIgnoreSslErrorsOk returns a tuple with the IgnoreSslErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSslErrors

`func (o *WebhookUpdate) SetIgnoreSslErrors(v bool)`

SetIgnoreSslErrors sets IgnoreSslErrors field to given value.

### HasIgnoreSslErrors

`func (o *WebhookUpdate) HasIgnoreSslErrors() bool`

HasIgnoreSslErrors returns a boolean if a field has been set.

### SetIgnoreSslErrorsNil

`func (o *WebhookUpdate) SetIgnoreSslErrorsNil(b bool)`

 SetIgnoreSslErrorsNil sets the value for IgnoreSslErrors to be an explicit nil

### UnsetIgnoreSslErrors
`func (o *WebhookUpdate) UnsetIgnoreSslErrors()`

UnsetIgnoreSslErrors ensures that no value is present for IgnoreSslErrors, not even an explicit nil
### GetDoNotRetry

`func (o *WebhookUpdate) GetDoNotRetry() bool`

GetDoNotRetry returns the DoNotRetry field if non-nil, zero value otherwise.

### GetDoNotRetryOk

`func (o *WebhookUpdate) GetDoNotRetryOk() (*bool, bool)`

GetDoNotRetryOk returns a tuple with the DoNotRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotRetry

`func (o *WebhookUpdate) SetDoNotRetry(v bool)`

SetDoNotRetry sets DoNotRetry field to given value.

### HasDoNotRetry

`func (o *WebhookUpdate) HasDoNotRetry() bool`

HasDoNotRetry returns a boolean if a field has been set.

### SetDoNotRetryNil

`func (o *WebhookUpdate) SetDoNotRetryNil(b bool)`

 SetDoNotRetryNil sets the value for DoNotRetry to be an explicit nil

### UnsetDoNotRetry
`func (o *WebhookUpdate) UnsetDoNotRetry()`

UnsetDoNotRetry ensures that no value is present for DoNotRetry, not even an explicit nil
### GetRequestUrl

`func (o *WebhookUpdate) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *WebhookUpdate) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *WebhookUpdate) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *WebhookUpdate) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### SetRequestUrlNil

`func (o *WebhookUpdate) SetRequestUrlNil(b bool)`

 SetRequestUrlNil sets the value for RequestUrl to be an explicit nil

### UnsetRequestUrl
`func (o *WebhookUpdate) UnsetRequestUrl()`

UnsetRequestUrl ensures that no value is present for RequestUrl, not even an explicit nil
### GetPayloadTemplate

`func (o *WebhookUpdate) GetPayloadTemplate() string`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *WebhookUpdate) GetPayloadTemplateOk() (*string, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *WebhookUpdate) SetPayloadTemplate(v string)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *WebhookUpdate) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### SetPayloadTemplateNil

`func (o *WebhookUpdate) SetPayloadTemplateNil(b bool)`

 SetPayloadTemplateNil sets the value for PayloadTemplate to be an explicit nil

### UnsetPayloadTemplate
`func (o *WebhookUpdate) UnsetPayloadTemplate()`

UnsetPayloadTemplate ensures that no value is present for PayloadTemplate, not even an explicit nil
### GetHeadersTemplate

`func (o *WebhookUpdate) GetHeadersTemplate() string`

GetHeadersTemplate returns the HeadersTemplate field if non-nil, zero value otherwise.

### GetHeadersTemplateOk

`func (o *WebhookUpdate) GetHeadersTemplateOk() (*string, bool)`

GetHeadersTemplateOk returns a tuple with the HeadersTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersTemplate

`func (o *WebhookUpdate) SetHeadersTemplate(v string)`

SetHeadersTemplate sets HeadersTemplate field to given value.

### HasHeadersTemplate

`func (o *WebhookUpdate) HasHeadersTemplate() bool`

HasHeadersTemplate returns a boolean if a field has been set.

### SetHeadersTemplateNil

`func (o *WebhookUpdate) SetHeadersTemplateNil(b bool)`

 SetHeadersTemplateNil sets the value for HeadersTemplate to be an explicit nil

### UnsetHeadersTemplate
`func (o *WebhookUpdate) UnsetHeadersTemplate()`

UnsetHeadersTemplate ensures that no value is present for HeadersTemplate, not even an explicit nil
### GetDescription

`func (o *WebhookUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebhookUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShouldInterpolateStrings

`func (o *WebhookUpdate) GetShouldInterpolateStrings() bool`

GetShouldInterpolateStrings returns the ShouldInterpolateStrings field if non-nil, zero value otherwise.

### GetShouldInterpolateStringsOk

`func (o *WebhookUpdate) GetShouldInterpolateStringsOk() (*bool, bool)`

GetShouldInterpolateStringsOk returns a tuple with the ShouldInterpolateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInterpolateStrings

`func (o *WebhookUpdate) SetShouldInterpolateStrings(v bool)`

SetShouldInterpolateStrings sets ShouldInterpolateStrings field to given value.

### HasShouldInterpolateStrings

`func (o *WebhookUpdate) HasShouldInterpolateStrings() bool`

HasShouldInterpolateStrings returns a boolean if a field has been set.

### SetShouldInterpolateStringsNil

`func (o *WebhookUpdate) SetShouldInterpolateStringsNil(b bool)`

 SetShouldInterpolateStringsNil sets the value for ShouldInterpolateStrings to be an explicit nil

### UnsetShouldInterpolateStrings
`func (o *WebhookUpdate) UnsetShouldInterpolateStrings()`

UnsetShouldInterpolateStrings ensures that no value is present for ShouldInterpolateStrings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


