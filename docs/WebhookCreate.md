# WebhookCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAdHoc** | Pointer to **NullableBool** |  | [optional] 
**EventTypes** | [**[]WebhookEventType**](WebhookEventType.md) |  | 
**Condition** | [**WebhookCondition**](WebhookCondition.md) |  | 
**IdempotencyKey** | Pointer to **NullableString** |  | [optional] 
**IgnoreSslErrors** | Pointer to **NullableBool** |  | [optional] 
**DoNotRetry** | Pointer to **NullableBool** |  | [optional] 
**RequestUrl** | **string** |  | 
**PayloadTemplate** | Pointer to **NullableString** |  | [optional] 
**HeadersTemplate** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShouldInterpolateStrings** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWebhookCreate

`func NewWebhookCreate(eventTypes []WebhookEventType, condition WebhookCondition, requestUrl string, ) *WebhookCreate`

NewWebhookCreate instantiates a new WebhookCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateWithDefaults

`func NewWebhookCreateWithDefaults() *WebhookCreate`

NewWebhookCreateWithDefaults instantiates a new WebhookCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdHoc

`func (o *WebhookCreate) GetIsAdHoc() bool`

GetIsAdHoc returns the IsAdHoc field if non-nil, zero value otherwise.

### GetIsAdHocOk

`func (o *WebhookCreate) GetIsAdHocOk() (*bool, bool)`

GetIsAdHocOk returns a tuple with the IsAdHoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdHoc

`func (o *WebhookCreate) SetIsAdHoc(v bool)`

SetIsAdHoc sets IsAdHoc field to given value.

### HasIsAdHoc

`func (o *WebhookCreate) HasIsAdHoc() bool`

HasIsAdHoc returns a boolean if a field has been set.

### SetIsAdHocNil

`func (o *WebhookCreate) SetIsAdHocNil(b bool)`

 SetIsAdHocNil sets the value for IsAdHoc to be an explicit nil

### UnsetIsAdHoc
`func (o *WebhookCreate) UnsetIsAdHoc()`

UnsetIsAdHoc ensures that no value is present for IsAdHoc, not even an explicit nil
### GetEventTypes

`func (o *WebhookCreate) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookCreate) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookCreate) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.


### GetCondition

`func (o *WebhookCreate) GetCondition() WebhookCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WebhookCreate) GetConditionOk() (*WebhookCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WebhookCreate) SetCondition(v WebhookCondition)`

SetCondition sets Condition field to given value.


### GetIdempotencyKey

`func (o *WebhookCreate) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *WebhookCreate) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *WebhookCreate) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *WebhookCreate) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *WebhookCreate) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *WebhookCreate) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetIgnoreSslErrors

`func (o *WebhookCreate) GetIgnoreSslErrors() bool`

GetIgnoreSslErrors returns the IgnoreSslErrors field if non-nil, zero value otherwise.

### GetIgnoreSslErrorsOk

`func (o *WebhookCreate) GetIgnoreSslErrorsOk() (*bool, bool)`

GetIgnoreSslErrorsOk returns a tuple with the IgnoreSslErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSslErrors

`func (o *WebhookCreate) SetIgnoreSslErrors(v bool)`

SetIgnoreSslErrors sets IgnoreSslErrors field to given value.

### HasIgnoreSslErrors

`func (o *WebhookCreate) HasIgnoreSslErrors() bool`

HasIgnoreSslErrors returns a boolean if a field has been set.

### SetIgnoreSslErrorsNil

`func (o *WebhookCreate) SetIgnoreSslErrorsNil(b bool)`

 SetIgnoreSslErrorsNil sets the value for IgnoreSslErrors to be an explicit nil

### UnsetIgnoreSslErrors
`func (o *WebhookCreate) UnsetIgnoreSslErrors()`

UnsetIgnoreSslErrors ensures that no value is present for IgnoreSslErrors, not even an explicit nil
### GetDoNotRetry

`func (o *WebhookCreate) GetDoNotRetry() bool`

GetDoNotRetry returns the DoNotRetry field if non-nil, zero value otherwise.

### GetDoNotRetryOk

`func (o *WebhookCreate) GetDoNotRetryOk() (*bool, bool)`

GetDoNotRetryOk returns a tuple with the DoNotRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotRetry

`func (o *WebhookCreate) SetDoNotRetry(v bool)`

SetDoNotRetry sets DoNotRetry field to given value.

### HasDoNotRetry

`func (o *WebhookCreate) HasDoNotRetry() bool`

HasDoNotRetry returns a boolean if a field has been set.

### SetDoNotRetryNil

`func (o *WebhookCreate) SetDoNotRetryNil(b bool)`

 SetDoNotRetryNil sets the value for DoNotRetry to be an explicit nil

### UnsetDoNotRetry
`func (o *WebhookCreate) UnsetDoNotRetry()`

UnsetDoNotRetry ensures that no value is present for DoNotRetry, not even an explicit nil
### GetRequestUrl

`func (o *WebhookCreate) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *WebhookCreate) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *WebhookCreate) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetPayloadTemplate

`func (o *WebhookCreate) GetPayloadTemplate() string`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *WebhookCreate) GetPayloadTemplateOk() (*string, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *WebhookCreate) SetPayloadTemplate(v string)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *WebhookCreate) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### SetPayloadTemplateNil

`func (o *WebhookCreate) SetPayloadTemplateNil(b bool)`

 SetPayloadTemplateNil sets the value for PayloadTemplate to be an explicit nil

### UnsetPayloadTemplate
`func (o *WebhookCreate) UnsetPayloadTemplate()`

UnsetPayloadTemplate ensures that no value is present for PayloadTemplate, not even an explicit nil
### GetHeadersTemplate

`func (o *WebhookCreate) GetHeadersTemplate() string`

GetHeadersTemplate returns the HeadersTemplate field if non-nil, zero value otherwise.

### GetHeadersTemplateOk

`func (o *WebhookCreate) GetHeadersTemplateOk() (*string, bool)`

GetHeadersTemplateOk returns a tuple with the HeadersTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersTemplate

`func (o *WebhookCreate) SetHeadersTemplate(v string)`

SetHeadersTemplate sets HeadersTemplate field to given value.

### HasHeadersTemplate

`func (o *WebhookCreate) HasHeadersTemplate() bool`

HasHeadersTemplate returns a boolean if a field has been set.

### SetHeadersTemplateNil

`func (o *WebhookCreate) SetHeadersTemplateNil(b bool)`

 SetHeadersTemplateNil sets the value for HeadersTemplate to be an explicit nil

### UnsetHeadersTemplate
`func (o *WebhookCreate) UnsetHeadersTemplate()`

UnsetHeadersTemplate ensures that no value is present for HeadersTemplate, not even an explicit nil
### GetDescription

`func (o *WebhookCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebhookCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShouldInterpolateStrings

`func (o *WebhookCreate) GetShouldInterpolateStrings() bool`

GetShouldInterpolateStrings returns the ShouldInterpolateStrings field if non-nil, zero value otherwise.

### GetShouldInterpolateStringsOk

`func (o *WebhookCreate) GetShouldInterpolateStringsOk() (*bool, bool)`

GetShouldInterpolateStringsOk returns a tuple with the ShouldInterpolateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInterpolateStrings

`func (o *WebhookCreate) SetShouldInterpolateStrings(v bool)`

SetShouldInterpolateStrings sets ShouldInterpolateStrings field to given value.

### HasShouldInterpolateStrings

`func (o *WebhookCreate) HasShouldInterpolateStrings() bool`

HasShouldInterpolateStrings returns a boolean if a field has been set.

### SetShouldInterpolateStringsNil

`func (o *WebhookCreate) SetShouldInterpolateStringsNil(b bool)`

 SetShouldInterpolateStringsNil sets the value for ShouldInterpolateStrings to be an explicit nil

### UnsetShouldInterpolateStrings
`func (o *WebhookCreate) UnsetShouldInterpolateStrings()`

UnsetShouldInterpolateStrings ensures that no value is present for ShouldInterpolateStrings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


