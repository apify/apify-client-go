# BuildStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMillis** | Pointer to **int32** |  | [optional] 
**RunTimeSecs** | Pointer to **float32** |  | [optional] 
**ComputeUnits** | **float32** |  | 
**ImageSizeBytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewBuildStats

`func NewBuildStats(computeUnits float32, ) *BuildStats`

NewBuildStats instantiates a new BuildStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildStatsWithDefaults

`func NewBuildStatsWithDefaults() *BuildStats`

NewBuildStatsWithDefaults instantiates a new BuildStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMillis

`func (o *BuildStats) GetDurationMillis() int32`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *BuildStats) GetDurationMillisOk() (*int32, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *BuildStats) SetDurationMillis(v int32)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *BuildStats) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetRunTimeSecs

`func (o *BuildStats) GetRunTimeSecs() float32`

GetRunTimeSecs returns the RunTimeSecs field if non-nil, zero value otherwise.

### GetRunTimeSecsOk

`func (o *BuildStats) GetRunTimeSecsOk() (*float32, bool)`

GetRunTimeSecsOk returns a tuple with the RunTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeSecs

`func (o *BuildStats) SetRunTimeSecs(v float32)`

SetRunTimeSecs sets RunTimeSecs field to given value.

### HasRunTimeSecs

`func (o *BuildStats) HasRunTimeSecs() bool`

HasRunTimeSecs returns a boolean if a field has been set.

### GetComputeUnits

`func (o *BuildStats) GetComputeUnits() float32`

GetComputeUnits returns the ComputeUnits field if non-nil, zero value otherwise.

### GetComputeUnitsOk

`func (o *BuildStats) GetComputeUnitsOk() (*float32, bool)`

GetComputeUnitsOk returns a tuple with the ComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnits

`func (o *BuildStats) SetComputeUnits(v float32)`

SetComputeUnits sets ComputeUnits field to given value.


### GetImageSizeBytes

`func (o *BuildStats) GetImageSizeBytes() int32`

GetImageSizeBytes returns the ImageSizeBytes field if non-nil, zero value otherwise.

### GetImageSizeBytesOk

`func (o *BuildStats) GetImageSizeBytesOk() (*int32, bool)`

GetImageSizeBytesOk returns a tuple with the ImageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSizeBytes

`func (o *BuildStats) SetImageSizeBytes(v int32)`

SetImageSizeBytes sets ImageSizeBytes field to given value.

### HasImageSizeBytes

`func (o *BuildStats) HasImageSizeBytes() bool`

HasImageSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


