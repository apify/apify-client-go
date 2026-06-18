# KeyValueStoreStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadCount** | **int32** |  | 
**WriteCount** | **int32** |  | 
**DeleteCount** | **int32** |  | 
**ListCount** | **int32** |  | 
**S3StorageBytes** | Pointer to **int32** |  | [optional] 
**StorageBytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewKeyValueStoreStats

`func NewKeyValueStoreStats(readCount int32, writeCount int32, deleteCount int32, listCount int32, ) *KeyValueStoreStats`

NewKeyValueStoreStats instantiates a new KeyValueStoreStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyValueStoreStatsWithDefaults

`func NewKeyValueStoreStatsWithDefaults() *KeyValueStoreStats`

NewKeyValueStoreStatsWithDefaults instantiates a new KeyValueStoreStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadCount

`func (o *KeyValueStoreStats) GetReadCount() int32`

GetReadCount returns the ReadCount field if non-nil, zero value otherwise.

### GetReadCountOk

`func (o *KeyValueStoreStats) GetReadCountOk() (*int32, bool)`

GetReadCountOk returns a tuple with the ReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCount

`func (o *KeyValueStoreStats) SetReadCount(v int32)`

SetReadCount sets ReadCount field to given value.


### GetWriteCount

`func (o *KeyValueStoreStats) GetWriteCount() int32`

GetWriteCount returns the WriteCount field if non-nil, zero value otherwise.

### GetWriteCountOk

`func (o *KeyValueStoreStats) GetWriteCountOk() (*int32, bool)`

GetWriteCountOk returns a tuple with the WriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCount

`func (o *KeyValueStoreStats) SetWriteCount(v int32)`

SetWriteCount sets WriteCount field to given value.


### GetDeleteCount

`func (o *KeyValueStoreStats) GetDeleteCount() int32`

GetDeleteCount returns the DeleteCount field if non-nil, zero value otherwise.

### GetDeleteCountOk

`func (o *KeyValueStoreStats) GetDeleteCountOk() (*int32, bool)`

GetDeleteCountOk returns a tuple with the DeleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCount

`func (o *KeyValueStoreStats) SetDeleteCount(v int32)`

SetDeleteCount sets DeleteCount field to given value.


### GetListCount

`func (o *KeyValueStoreStats) GetListCount() int32`

GetListCount returns the ListCount field if non-nil, zero value otherwise.

### GetListCountOk

`func (o *KeyValueStoreStats) GetListCountOk() (*int32, bool)`

GetListCountOk returns a tuple with the ListCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListCount

`func (o *KeyValueStoreStats) SetListCount(v int32)`

SetListCount sets ListCount field to given value.


### GetS3StorageBytes

`func (o *KeyValueStoreStats) GetS3StorageBytes() int32`

GetS3StorageBytes returns the S3StorageBytes field if non-nil, zero value otherwise.

### GetS3StorageBytesOk

`func (o *KeyValueStoreStats) GetS3StorageBytesOk() (*int32, bool)`

GetS3StorageBytesOk returns a tuple with the S3StorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3StorageBytes

`func (o *KeyValueStoreStats) SetS3StorageBytes(v int32)`

SetS3StorageBytes sets S3StorageBytes field to given value.

### HasS3StorageBytes

`func (o *KeyValueStoreStats) HasS3StorageBytes() bool`

HasS3StorageBytes returns a boolean if a field has been set.

### GetStorageBytes

`func (o *KeyValueStoreStats) GetStorageBytes() int32`

GetStorageBytes returns the StorageBytes field if non-nil, zero value otherwise.

### GetStorageBytesOk

`func (o *KeyValueStoreStats) GetStorageBytesOk() (*int32, bool)`

GetStorageBytesOk returns a tuple with the StorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBytes

`func (o *KeyValueStoreStats) SetStorageBytes(v int32)`

SetStorageBytes sets StorageBytes field to given value.

### HasStorageBytes

`func (o *KeyValueStoreStats) HasStorageBytes() bool`

HasStorageBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


