# LockedHeadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier assigned to the request. | 
**UniqueKey** | **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | 
**Url** | **string** | The URL of the request. | 
**Method** | Pointer to [**HttpMethod**](HttpMethod.md) |  | [optional] 
**RetryCount** | Pointer to **int32** | The number of times this request has been retried. | [optional] 
**LockExpiresAt** | **time.Time** | The timestamp when the lock on this request expires. | 

## Methods

### NewLockedHeadRequest

`func NewLockedHeadRequest(id string, uniqueKey string, url string, lockExpiresAt time.Time, ) *LockedHeadRequest`

NewLockedHeadRequest instantiates a new LockedHeadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockedHeadRequestWithDefaults

`func NewLockedHeadRequestWithDefaults() *LockedHeadRequest`

NewLockedHeadRequestWithDefaults instantiates a new LockedHeadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LockedHeadRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LockedHeadRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LockedHeadRequest) SetId(v string)`

SetId sets Id field to given value.


### GetUniqueKey

`func (o *LockedHeadRequest) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *LockedHeadRequest) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *LockedHeadRequest) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.


### GetUrl

`func (o *LockedHeadRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LockedHeadRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LockedHeadRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *LockedHeadRequest) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LockedHeadRequest) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LockedHeadRequest) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LockedHeadRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRetryCount

`func (o *LockedHeadRequest) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *LockedHeadRequest) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *LockedHeadRequest) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *LockedHeadRequest) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetLockExpiresAt

`func (o *LockedHeadRequest) GetLockExpiresAt() time.Time`

GetLockExpiresAt returns the LockExpiresAt field if non-nil, zero value otherwise.

### GetLockExpiresAtOk

`func (o *LockedHeadRequest) GetLockExpiresAtOk() (*time.Time, bool)`

GetLockExpiresAtOk returns a tuple with the LockExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpiresAt

`func (o *LockedHeadRequest) SetLockExpiresAt(v time.Time)`

SetLockExpiresAt sets LockExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


