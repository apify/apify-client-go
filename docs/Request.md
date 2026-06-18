# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueKey** | Pointer to **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | [optional] 
**Url** | Pointer to **string** | The URL of the request. | [optional] 
**Method** | Pointer to [**HttpMethod**](HttpMethod.md) |  | [optional] 
**RetryCount** | Pointer to **int32** | The number of times this request has been retried. | [optional] 
**LoadedUrl** | Pointer to **string** | The final URL that was loaded, after redirects (if any). | [optional] 
**Payload** | Pointer to [**NullableRequestBasePayload**](RequestBasePayload.md) |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** | HTTP headers sent with the request. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | Custom user data attached to the request. Can contain arbitrary fields. | [optional] 
**NoRetry** | Pointer to **bool** | Indicates whether the request should not be retried if processing fails. | [optional] 
**ErrorMessages** | Pointer to **[]string** | Error messages recorded from failed processing attempts. | [optional] 
**HandledAt** | Pointer to **time.Time** | The timestamp when the request was marked as handled, if applicable. | [optional] 
**Id** | Pointer to **string** | A unique identifier assigned to the request. | [optional] 

## Methods

### NewRequest

`func NewRequest() *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueKey

`func (o *Request) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *Request) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *Request) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.

### HasUniqueKey

`func (o *Request) HasUniqueKey() bool`

HasUniqueKey returns a boolean if a field has been set.

### GetUrl

`func (o *Request) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Request) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Request) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Request) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *Request) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Request) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Request) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Request) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRetryCount

`func (o *Request) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *Request) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *Request) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *Request) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetLoadedUrl

`func (o *Request) GetLoadedUrl() string`

GetLoadedUrl returns the LoadedUrl field if non-nil, zero value otherwise.

### GetLoadedUrlOk

`func (o *Request) GetLoadedUrlOk() (*string, bool)`

GetLoadedUrlOk returns a tuple with the LoadedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedUrl

`func (o *Request) SetLoadedUrl(v string)`

SetLoadedUrl sets LoadedUrl field to given value.

### HasLoadedUrl

`func (o *Request) HasLoadedUrl() bool`

HasLoadedUrl returns a boolean if a field has been set.

### GetPayload

`func (o *Request) GetPayload() RequestBasePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Request) GetPayloadOk() (*RequestBasePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Request) SetPayload(v RequestBasePayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Request) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *Request) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Request) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetHeaders

`func (o *Request) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Request) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Request) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Request) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUserData

`func (o *Request) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *Request) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *Request) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *Request) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetNoRetry

`func (o *Request) GetNoRetry() bool`

GetNoRetry returns the NoRetry field if non-nil, zero value otherwise.

### GetNoRetryOk

`func (o *Request) GetNoRetryOk() (*bool, bool)`

GetNoRetryOk returns a tuple with the NoRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRetry

`func (o *Request) SetNoRetry(v bool)`

SetNoRetry sets NoRetry field to given value.

### HasNoRetry

`func (o *Request) HasNoRetry() bool`

HasNoRetry returns a boolean if a field has been set.

### GetErrorMessages

`func (o *Request) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *Request) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *Request) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *Request) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetHandledAt

`func (o *Request) GetHandledAt() time.Time`

GetHandledAt returns the HandledAt field if non-nil, zero value otherwise.

### GetHandledAtOk

`func (o *Request) GetHandledAtOk() (*time.Time, bool)`

GetHandledAtOk returns a tuple with the HandledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAt

`func (o *Request) SetHandledAt(v time.Time)`

SetHandledAt sets HandledAt field to given value.

### HasHandledAt

`func (o *Request) HasHandledAt() bool`

HasHandledAt returns a boolean if a field has been set.

### GetId

`func (o *Request) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Request) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Request) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Request) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


