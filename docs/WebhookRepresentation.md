# WebhookRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypes** | [**[]WebhookEventType**](WebhookEventType.md) |  | 
**RequestUrl** | **string** | The URL to which the webhook sends its payload. | 
**PayloadTemplate** | Pointer to **NullableString** | Optional template for the JSON payload sent by the webhook. | [optional] 
**HeadersTemplate** | Pointer to **NullableString** | Optional template for the HTTP headers sent by the webhook. | [optional] 
**ShouldInterpolateStrings** | Pointer to **NullableBool** | Flag to also interpolate &#x60;{{...}}&#x60; variables inside string values of the payload and headers templates. | [optional] 
**IdempotencyKey** | Pointer to **NullableString** | Key that prevents creating duplicate webhooks, e.g. when the run-starting request is retried. | [optional] 
**IgnoreSslErrors** | Pointer to **NullableBool** | Flag to ignore SSL errors when the webhook sends the request. | [optional] 
**DoNotRetry** | Pointer to **NullableBool** | Flag to skip retrying the webhook request on failure. | [optional] 

## Methods

### NewWebhookRepresentation

`func NewWebhookRepresentation(eventTypes []WebhookEventType, requestUrl string, ) *WebhookRepresentation`

NewWebhookRepresentation instantiates a new WebhookRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRepresentationWithDefaults

`func NewWebhookRepresentationWithDefaults() *WebhookRepresentation`

NewWebhookRepresentationWithDefaults instantiates a new WebhookRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypes

`func (o *WebhookRepresentation) GetEventTypes() []WebhookEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookRepresentation) GetEventTypesOk() (*[]WebhookEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookRepresentation) SetEventTypes(v []WebhookEventType)`

SetEventTypes sets EventTypes field to given value.


### GetRequestUrl

`func (o *WebhookRepresentation) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *WebhookRepresentation) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *WebhookRepresentation) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetPayloadTemplate

`func (o *WebhookRepresentation) GetPayloadTemplate() string`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *WebhookRepresentation) GetPayloadTemplateOk() (*string, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *WebhookRepresentation) SetPayloadTemplate(v string)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *WebhookRepresentation) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### SetPayloadTemplateNil

`func (o *WebhookRepresentation) SetPayloadTemplateNil(b bool)`

 SetPayloadTemplateNil sets the value for PayloadTemplate to be an explicit nil

### UnsetPayloadTemplate
`func (o *WebhookRepresentation) UnsetPayloadTemplate()`

UnsetPayloadTemplate ensures that no value is present for PayloadTemplate, not even an explicit nil
### GetHeadersTemplate

`func (o *WebhookRepresentation) GetHeadersTemplate() string`

GetHeadersTemplate returns the HeadersTemplate field if non-nil, zero value otherwise.

### GetHeadersTemplateOk

`func (o *WebhookRepresentation) GetHeadersTemplateOk() (*string, bool)`

GetHeadersTemplateOk returns a tuple with the HeadersTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersTemplate

`func (o *WebhookRepresentation) SetHeadersTemplate(v string)`

SetHeadersTemplate sets HeadersTemplate field to given value.

### HasHeadersTemplate

`func (o *WebhookRepresentation) HasHeadersTemplate() bool`

HasHeadersTemplate returns a boolean if a field has been set.

### SetHeadersTemplateNil

`func (o *WebhookRepresentation) SetHeadersTemplateNil(b bool)`

 SetHeadersTemplateNil sets the value for HeadersTemplate to be an explicit nil

### UnsetHeadersTemplate
`func (o *WebhookRepresentation) UnsetHeadersTemplate()`

UnsetHeadersTemplate ensures that no value is present for HeadersTemplate, not even an explicit nil
### GetShouldInterpolateStrings

`func (o *WebhookRepresentation) GetShouldInterpolateStrings() bool`

GetShouldInterpolateStrings returns the ShouldInterpolateStrings field if non-nil, zero value otherwise.

### GetShouldInterpolateStringsOk

`func (o *WebhookRepresentation) GetShouldInterpolateStringsOk() (*bool, bool)`

GetShouldInterpolateStringsOk returns a tuple with the ShouldInterpolateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInterpolateStrings

`func (o *WebhookRepresentation) SetShouldInterpolateStrings(v bool)`

SetShouldInterpolateStrings sets ShouldInterpolateStrings field to given value.

### HasShouldInterpolateStrings

`func (o *WebhookRepresentation) HasShouldInterpolateStrings() bool`

HasShouldInterpolateStrings returns a boolean if a field has been set.

### SetShouldInterpolateStringsNil

`func (o *WebhookRepresentation) SetShouldInterpolateStringsNil(b bool)`

 SetShouldInterpolateStringsNil sets the value for ShouldInterpolateStrings to be an explicit nil

### UnsetShouldInterpolateStrings
`func (o *WebhookRepresentation) UnsetShouldInterpolateStrings()`

UnsetShouldInterpolateStrings ensures that no value is present for ShouldInterpolateStrings, not even an explicit nil
### GetIdempotencyKey

`func (o *WebhookRepresentation) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *WebhookRepresentation) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *WebhookRepresentation) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *WebhookRepresentation) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *WebhookRepresentation) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *WebhookRepresentation) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetIgnoreSslErrors

`func (o *WebhookRepresentation) GetIgnoreSslErrors() bool`

GetIgnoreSslErrors returns the IgnoreSslErrors field if non-nil, zero value otherwise.

### GetIgnoreSslErrorsOk

`func (o *WebhookRepresentation) GetIgnoreSslErrorsOk() (*bool, bool)`

GetIgnoreSslErrorsOk returns a tuple with the IgnoreSslErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSslErrors

`func (o *WebhookRepresentation) SetIgnoreSslErrors(v bool)`

SetIgnoreSslErrors sets IgnoreSslErrors field to given value.

### HasIgnoreSslErrors

`func (o *WebhookRepresentation) HasIgnoreSslErrors() bool`

HasIgnoreSslErrors returns a boolean if a field has been set.

### SetIgnoreSslErrorsNil

`func (o *WebhookRepresentation) SetIgnoreSslErrorsNil(b bool)`

 SetIgnoreSslErrorsNil sets the value for IgnoreSslErrors to be an explicit nil

### UnsetIgnoreSslErrors
`func (o *WebhookRepresentation) UnsetIgnoreSslErrors()`

UnsetIgnoreSslErrors ensures that no value is present for IgnoreSslErrors, not even an explicit nil
### GetDoNotRetry

`func (o *WebhookRepresentation) GetDoNotRetry() bool`

GetDoNotRetry returns the DoNotRetry field if non-nil, zero value otherwise.

### GetDoNotRetryOk

`func (o *WebhookRepresentation) GetDoNotRetryOk() (*bool, bool)`

GetDoNotRetryOk returns a tuple with the DoNotRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotRetry

`func (o *WebhookRepresentation) SetDoNotRetry(v bool)`

SetDoNotRetry sets DoNotRetry field to given value.

### HasDoNotRetry

`func (o *WebhookRepresentation) HasDoNotRetry() bool`

HasDoNotRetry returns a boolean if a field has been set.

### SetDoNotRetryNil

`func (o *WebhookRepresentation) SetDoNotRetryNil(b bool)`

 SetDoNotRetryNil sets the value for DoNotRetry to be an explicit nil

### UnsetDoNotRetry
`func (o *WebhookRepresentation) UnsetDoNotRetry()`

UnsetDoNotRetry ensures that no value is present for DoNotRetry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


