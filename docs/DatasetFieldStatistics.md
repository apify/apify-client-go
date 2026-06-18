# DatasetFieldStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Min** | Pointer to **NullableFloat32** | Minimum value of the field. For numbers, this is calculated directly. For strings, this is the length of the shortest string. For arrays, this is the length of the shortest array. For objects, this is the number of keys in the smallest object. | [optional] 
**Max** | Pointer to **NullableFloat32** | Maximum value of the field. For numbers, this is calculated directly. For strings, this is the length of the longest string. For arrays, this is the length of the longest array. For objects, this is the number of keys in the largest object. | [optional] 
**NullCount** | Pointer to **NullableInt32** | How many items in the dataset have a null value for this field. | [optional] 
**EmptyCount** | Pointer to **NullableInt32** | How many items in the dataset are &#x60;undefined&#x60;, meaning that for example empty string is not considered empty. | [optional] 

## Methods

### NewDatasetFieldStatistics

`func NewDatasetFieldStatistics() *DatasetFieldStatistics`

NewDatasetFieldStatistics instantiates a new DatasetFieldStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetFieldStatisticsWithDefaults

`func NewDatasetFieldStatisticsWithDefaults() *DatasetFieldStatistics`

NewDatasetFieldStatisticsWithDefaults instantiates a new DatasetFieldStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMin

`func (o *DatasetFieldStatistics) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DatasetFieldStatistics) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DatasetFieldStatistics) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *DatasetFieldStatistics) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *DatasetFieldStatistics) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *DatasetFieldStatistics) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *DatasetFieldStatistics) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DatasetFieldStatistics) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DatasetFieldStatistics) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *DatasetFieldStatistics) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *DatasetFieldStatistics) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *DatasetFieldStatistics) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetNullCount

`func (o *DatasetFieldStatistics) GetNullCount() int32`

GetNullCount returns the NullCount field if non-nil, zero value otherwise.

### GetNullCountOk

`func (o *DatasetFieldStatistics) GetNullCountOk() (*int32, bool)`

GetNullCountOk returns a tuple with the NullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullCount

`func (o *DatasetFieldStatistics) SetNullCount(v int32)`

SetNullCount sets NullCount field to given value.

### HasNullCount

`func (o *DatasetFieldStatistics) HasNullCount() bool`

HasNullCount returns a boolean if a field has been set.

### SetNullCountNil

`func (o *DatasetFieldStatistics) SetNullCountNil(b bool)`

 SetNullCountNil sets the value for NullCount to be an explicit nil

### UnsetNullCount
`func (o *DatasetFieldStatistics) UnsetNullCount()`

UnsetNullCount ensures that no value is present for NullCount, not even an explicit nil
### GetEmptyCount

`func (o *DatasetFieldStatistics) GetEmptyCount() int32`

GetEmptyCount returns the EmptyCount field if non-nil, zero value otherwise.

### GetEmptyCountOk

`func (o *DatasetFieldStatistics) GetEmptyCountOk() (*int32, bool)`

GetEmptyCountOk returns a tuple with the EmptyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyCount

`func (o *DatasetFieldStatistics) SetEmptyCount(v int32)`

SetEmptyCount sets EmptyCount field to given value.

### HasEmptyCount

`func (o *DatasetFieldStatistics) HasEmptyCount() bool`

HasEmptyCount returns a boolean if a field has been set.

### SetEmptyCountNil

`func (o *DatasetFieldStatistics) SetEmptyCountNil(b bool)`

 SetEmptyCountNil sets the value for EmptyCount to be an explicit nil

### UnsetEmptyCount
`func (o *DatasetFieldStatistics) UnsetEmptyCount()`

UnsetEmptyCount ensures that no value is present for EmptyCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


