# RunOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | **string** |  | 
**TimeoutSecs** | **int32** |  | 
**MemoryMbytes** | **int32** |  | 
**DiskMbytes** | **int32** |  | 
**MaxItems** | Pointer to **NullableInt32** |  | [optional] 
**MaxTotalChargeUsd** | Pointer to **float32** |  | [optional] 

## Methods

### NewRunOptions

`func NewRunOptions(build string, timeoutSecs int32, memoryMbytes int32, diskMbytes int32, ) *RunOptions`

NewRunOptions instantiates a new RunOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunOptionsWithDefaults

`func NewRunOptionsWithDefaults() *RunOptions`

NewRunOptionsWithDefaults instantiates a new RunOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *RunOptions) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *RunOptions) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *RunOptions) SetBuild(v string)`

SetBuild sets Build field to given value.


### GetTimeoutSecs

`func (o *RunOptions) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *RunOptions) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *RunOptions) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.


### GetMemoryMbytes

`func (o *RunOptions) GetMemoryMbytes() int32`

GetMemoryMbytes returns the MemoryMbytes field if non-nil, zero value otherwise.

### GetMemoryMbytesOk

`func (o *RunOptions) GetMemoryMbytesOk() (*int32, bool)`

GetMemoryMbytesOk returns a tuple with the MemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMbytes

`func (o *RunOptions) SetMemoryMbytes(v int32)`

SetMemoryMbytes sets MemoryMbytes field to given value.


### GetDiskMbytes

`func (o *RunOptions) GetDiskMbytes() int32`

GetDiskMbytes returns the DiskMbytes field if non-nil, zero value otherwise.

### GetDiskMbytesOk

`func (o *RunOptions) GetDiskMbytesOk() (*int32, bool)`

GetDiskMbytesOk returns a tuple with the DiskMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMbytes

`func (o *RunOptions) SetDiskMbytes(v int32)`

SetDiskMbytes sets DiskMbytes field to given value.


### GetMaxItems

`func (o *RunOptions) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *RunOptions) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *RunOptions) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *RunOptions) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *RunOptions) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *RunOptions) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetMaxTotalChargeUsd

`func (o *RunOptions) GetMaxTotalChargeUsd() float32`

GetMaxTotalChargeUsd returns the MaxTotalChargeUsd field if non-nil, zero value otherwise.

### GetMaxTotalChargeUsdOk

`func (o *RunOptions) GetMaxTotalChargeUsdOk() (*float32, bool)`

GetMaxTotalChargeUsdOk returns a tuple with the MaxTotalChargeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalChargeUsd

`func (o *RunOptions) SetMaxTotalChargeUsd(v float32)`

SetMaxTotalChargeUsd sets MaxTotalChargeUsd field to given value.

### HasMaxTotalChargeUsd

`func (o *RunOptions) HasMaxTotalChargeUsd() bool`

HasMaxTotalChargeUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


