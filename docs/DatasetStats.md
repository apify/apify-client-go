# DatasetStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadCount** | Pointer to **int32** |  | [optional] 
**WriteCount** | Pointer to **int32** |  | [optional] 
**StorageBytes** | Pointer to **int32** | Total storage size in bytes. Only returned by the single-dataset endpoint. | [optional] 
**InflatedBytes** | Pointer to **int32** | Uncompressed size in bytes. Only returned by the dataset list endpoint. | [optional] 

## Methods

### NewDatasetStats

`func NewDatasetStats() *DatasetStats`

NewDatasetStats instantiates a new DatasetStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetStatsWithDefaults

`func NewDatasetStatsWithDefaults() *DatasetStats`

NewDatasetStatsWithDefaults instantiates a new DatasetStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadCount

`func (o *DatasetStats) GetReadCount() int32`

GetReadCount returns the ReadCount field if non-nil, zero value otherwise.

### GetReadCountOk

`func (o *DatasetStats) GetReadCountOk() (*int32, bool)`

GetReadCountOk returns a tuple with the ReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCount

`func (o *DatasetStats) SetReadCount(v int32)`

SetReadCount sets ReadCount field to given value.

### HasReadCount

`func (o *DatasetStats) HasReadCount() bool`

HasReadCount returns a boolean if a field has been set.

### GetWriteCount

`func (o *DatasetStats) GetWriteCount() int32`

GetWriteCount returns the WriteCount field if non-nil, zero value otherwise.

### GetWriteCountOk

`func (o *DatasetStats) GetWriteCountOk() (*int32, bool)`

GetWriteCountOk returns a tuple with the WriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCount

`func (o *DatasetStats) SetWriteCount(v int32)`

SetWriteCount sets WriteCount field to given value.

### HasWriteCount

`func (o *DatasetStats) HasWriteCount() bool`

HasWriteCount returns a boolean if a field has been set.

### GetStorageBytes

`func (o *DatasetStats) GetStorageBytes() int32`

GetStorageBytes returns the StorageBytes field if non-nil, zero value otherwise.

### GetStorageBytesOk

`func (o *DatasetStats) GetStorageBytesOk() (*int32, bool)`

GetStorageBytesOk returns a tuple with the StorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBytes

`func (o *DatasetStats) SetStorageBytes(v int32)`

SetStorageBytes sets StorageBytes field to given value.

### HasStorageBytes

`func (o *DatasetStats) HasStorageBytes() bool`

HasStorageBytes returns a boolean if a field has been set.

### GetInflatedBytes

`func (o *DatasetStats) GetInflatedBytes() int32`

GetInflatedBytes returns the InflatedBytes field if non-nil, zero value otherwise.

### GetInflatedBytesOk

`func (o *DatasetStats) GetInflatedBytesOk() (*int32, bool)`

GetInflatedBytesOk returns a tuple with the InflatedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflatedBytes

`func (o *DatasetStats) SetInflatedBytes(v int32)`

SetInflatedBytes sets InflatedBytes field to given value.

### HasInflatedBytes

`func (o *DatasetStats) HasInflatedBytes() bool`

HasInflatedBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


