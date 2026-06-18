# RequestLockInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockExpiresAt** | **time.Time** | The timestamp when the lock on this request expires. | 

## Methods

### NewRequestLockInfo

`func NewRequestLockInfo(lockExpiresAt time.Time, ) *RequestLockInfo`

NewRequestLockInfo instantiates a new RequestLockInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestLockInfoWithDefaults

`func NewRequestLockInfoWithDefaults() *RequestLockInfo`

NewRequestLockInfoWithDefaults instantiates a new RequestLockInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockExpiresAt

`func (o *RequestLockInfo) GetLockExpiresAt() time.Time`

GetLockExpiresAt returns the LockExpiresAt field if non-nil, zero value otherwise.

### GetLockExpiresAtOk

`func (o *RequestLockInfo) GetLockExpiresAtOk() (*time.Time, bool)`

GetLockExpiresAtOk returns a tuple with the LockExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpiresAt

`func (o *RequestLockInfo) SetLockExpiresAt(v time.Time)`

SetLockExpiresAt sets LockExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


