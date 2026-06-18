# RequestQueueStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteCount** | Pointer to **int32** | The number of delete operations performed on the request queue. | [optional] 
**HeadItemReadCount** | Pointer to **int32** | The number of times requests from the head were read. | [optional] 
**ReadCount** | Pointer to **int32** | The total number of read operations performed on the request queue. | [optional] 
**StorageBytes** | Pointer to **int32** | The total storage size in bytes used by the request queue. | [optional] 
**WriteCount** | Pointer to **int32** | The total number of write operations performed on the request queue. | [optional] 

## Methods

### NewRequestQueueStats

`func NewRequestQueueStats() *RequestQueueStats`

NewRequestQueueStats instantiates a new RequestQueueStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestQueueStatsWithDefaults

`func NewRequestQueueStatsWithDefaults() *RequestQueueStats`

NewRequestQueueStatsWithDefaults instantiates a new RequestQueueStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteCount

`func (o *RequestQueueStats) GetDeleteCount() int32`

GetDeleteCount returns the DeleteCount field if non-nil, zero value otherwise.

### GetDeleteCountOk

`func (o *RequestQueueStats) GetDeleteCountOk() (*int32, bool)`

GetDeleteCountOk returns a tuple with the DeleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCount

`func (o *RequestQueueStats) SetDeleteCount(v int32)`

SetDeleteCount sets DeleteCount field to given value.

### HasDeleteCount

`func (o *RequestQueueStats) HasDeleteCount() bool`

HasDeleteCount returns a boolean if a field has been set.

### GetHeadItemReadCount

`func (o *RequestQueueStats) GetHeadItemReadCount() int32`

GetHeadItemReadCount returns the HeadItemReadCount field if non-nil, zero value otherwise.

### GetHeadItemReadCountOk

`func (o *RequestQueueStats) GetHeadItemReadCountOk() (*int32, bool)`

GetHeadItemReadCountOk returns a tuple with the HeadItemReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadItemReadCount

`func (o *RequestQueueStats) SetHeadItemReadCount(v int32)`

SetHeadItemReadCount sets HeadItemReadCount field to given value.

### HasHeadItemReadCount

`func (o *RequestQueueStats) HasHeadItemReadCount() bool`

HasHeadItemReadCount returns a boolean if a field has been set.

### GetReadCount

`func (o *RequestQueueStats) GetReadCount() int32`

GetReadCount returns the ReadCount field if non-nil, zero value otherwise.

### GetReadCountOk

`func (o *RequestQueueStats) GetReadCountOk() (*int32, bool)`

GetReadCountOk returns a tuple with the ReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCount

`func (o *RequestQueueStats) SetReadCount(v int32)`

SetReadCount sets ReadCount field to given value.

### HasReadCount

`func (o *RequestQueueStats) HasReadCount() bool`

HasReadCount returns a boolean if a field has been set.

### GetStorageBytes

`func (o *RequestQueueStats) GetStorageBytes() int32`

GetStorageBytes returns the StorageBytes field if non-nil, zero value otherwise.

### GetStorageBytesOk

`func (o *RequestQueueStats) GetStorageBytesOk() (*int32, bool)`

GetStorageBytesOk returns a tuple with the StorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBytes

`func (o *RequestQueueStats) SetStorageBytes(v int32)`

SetStorageBytes sets StorageBytes field to given value.

### HasStorageBytes

`func (o *RequestQueueStats) HasStorageBytes() bool`

HasStorageBytes returns a boolean if a field has been set.

### GetWriteCount

`func (o *RequestQueueStats) GetWriteCount() int32`

GetWriteCount returns the WriteCount field if non-nil, zero value otherwise.

### GetWriteCountOk

`func (o *RequestQueueStats) GetWriteCountOk() (*int32, bool)`

GetWriteCountOk returns a tuple with the WriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCount

`func (o *RequestQueueStats) SetWriteCount(v int32)`

SetWriteCount sets WriteCount field to given value.

### HasWriteCount

`func (o *RequestQueueStats) HasWriteCount() bool`

HasWriteCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


