# RequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueKey** | Pointer to **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | [optional] 
**Url** | Pointer to **string** | The URL of the request. | [optional] 
**Method** | Pointer to [**HttpMethod**](HttpMethod.md) |  | [optional] 
**RetryCount** | Pointer to **int32** | The number of times this request has been retried. | [optional] 
**LoadedUrl** | Pointer to **NullableString** | The final URL that was loaded, after redirects (if any). | [optional] 
**Payload** | Pointer to [**NullableRequestBasePayload**](RequestBasePayload.md) |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** | HTTP headers sent with the request. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | Custom user data attached to the request. Can contain arbitrary fields. | [optional] 
**NoRetry** | Pointer to **NullableBool** | Indicates whether the request should not be retried if processing fails. | [optional] 
**ErrorMessages** | Pointer to **[]string** | Error messages recorded from failed processing attempts. | [optional] 
**HandledAt** | Pointer to **NullableTime** | The timestamp when the request was marked as handled, if applicable. | [optional] 

## Methods

### NewRequestBase

`func NewRequestBase() *RequestBase`

NewRequestBase instantiates a new RequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBaseWithDefaults

`func NewRequestBaseWithDefaults() *RequestBase`

NewRequestBaseWithDefaults instantiates a new RequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueKey

`func (o *RequestBase) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *RequestBase) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *RequestBase) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.

### HasUniqueKey

`func (o *RequestBase) HasUniqueKey() bool`

HasUniqueKey returns a boolean if a field has been set.

### GetUrl

`func (o *RequestBase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestBase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestBase) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RequestBase) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *RequestBase) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RequestBase) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RequestBase) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RequestBase) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRetryCount

`func (o *RequestBase) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *RequestBase) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *RequestBase) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *RequestBase) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetLoadedUrl

`func (o *RequestBase) GetLoadedUrl() string`

GetLoadedUrl returns the LoadedUrl field if non-nil, zero value otherwise.

### GetLoadedUrlOk

`func (o *RequestBase) GetLoadedUrlOk() (*string, bool)`

GetLoadedUrlOk returns a tuple with the LoadedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedUrl

`func (o *RequestBase) SetLoadedUrl(v string)`

SetLoadedUrl sets LoadedUrl field to given value.

### HasLoadedUrl

`func (o *RequestBase) HasLoadedUrl() bool`

HasLoadedUrl returns a boolean if a field has been set.

### SetLoadedUrlNil

`func (o *RequestBase) SetLoadedUrlNil(b bool)`

 SetLoadedUrlNil sets the value for LoadedUrl to be an explicit nil

### UnsetLoadedUrl
`func (o *RequestBase) UnsetLoadedUrl()`

UnsetLoadedUrl ensures that no value is present for LoadedUrl, not even an explicit nil
### GetPayload

`func (o *RequestBase) GetPayload() RequestBasePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *RequestBase) GetPayloadOk() (*RequestBasePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *RequestBase) SetPayload(v RequestBasePayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *RequestBase) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *RequestBase) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *RequestBase) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetHeaders

`func (o *RequestBase) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RequestBase) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RequestBase) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RequestBase) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *RequestBase) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *RequestBase) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetUserData

`func (o *RequestBase) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *RequestBase) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *RequestBase) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *RequestBase) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetNoRetry

`func (o *RequestBase) GetNoRetry() bool`

GetNoRetry returns the NoRetry field if non-nil, zero value otherwise.

### GetNoRetryOk

`func (o *RequestBase) GetNoRetryOk() (*bool, bool)`

GetNoRetryOk returns a tuple with the NoRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRetry

`func (o *RequestBase) SetNoRetry(v bool)`

SetNoRetry sets NoRetry field to given value.

### HasNoRetry

`func (o *RequestBase) HasNoRetry() bool`

HasNoRetry returns a boolean if a field has been set.

### SetNoRetryNil

`func (o *RequestBase) SetNoRetryNil(b bool)`

 SetNoRetryNil sets the value for NoRetry to be an explicit nil

### UnsetNoRetry
`func (o *RequestBase) UnsetNoRetry()`

UnsetNoRetry ensures that no value is present for NoRetry, not even an explicit nil
### GetErrorMessages

`func (o *RequestBase) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *RequestBase) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *RequestBase) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *RequestBase) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### SetErrorMessagesNil

`func (o *RequestBase) SetErrorMessagesNil(b bool)`

 SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil

### UnsetErrorMessages
`func (o *RequestBase) UnsetErrorMessages()`

UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil
### GetHandledAt

`func (o *RequestBase) GetHandledAt() time.Time`

GetHandledAt returns the HandledAt field if non-nil, zero value otherwise.

### GetHandledAtOk

`func (o *RequestBase) GetHandledAtOk() (*time.Time, bool)`

GetHandledAtOk returns a tuple with the HandledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAt

`func (o *RequestBase) SetHandledAt(v time.Time)`

SetHandledAt sets HandledAt field to given value.

### HasHandledAt

`func (o *RequestBase) HasHandledAt() bool`

HasHandledAt returns a boolean if a field has been set.

### SetHandledAtNil

`func (o *RequestBase) SetHandledAtNil(b bool)`

 SetHandledAtNil sets the value for HandledAt to be an explicit nil

### UnsetHandledAt
`func (o *RequestBase) UnsetHandledAt()`

UnsetHandledAt ensures that no value is present for HandledAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


