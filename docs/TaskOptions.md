# TaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to **NullableString** |  | [optional] 
**TimeoutSecs** | Pointer to **NullableInt32** |  | [optional] 
**MemoryMbytes** | Pointer to **NullableInt32** |  | [optional] 
**MaxItems** | Pointer to **NullableInt32** |  | [optional] 
**MaxTotalChargeUsd** | Pointer to **NullableFloat32** |  | [optional] 
**RestartOnError** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewTaskOptions

`func NewTaskOptions() *TaskOptions`

NewTaskOptions instantiates a new TaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskOptionsWithDefaults

`func NewTaskOptionsWithDefaults() *TaskOptions`

NewTaskOptionsWithDefaults instantiates a new TaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *TaskOptions) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TaskOptions) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TaskOptions) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TaskOptions) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TaskOptions) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TaskOptions) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetTimeoutSecs

`func (o *TaskOptions) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *TaskOptions) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *TaskOptions) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *TaskOptions) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### SetTimeoutSecsNil

`func (o *TaskOptions) SetTimeoutSecsNil(b bool)`

 SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil

### UnsetTimeoutSecs
`func (o *TaskOptions) UnsetTimeoutSecs()`

UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil
### GetMemoryMbytes

`func (o *TaskOptions) GetMemoryMbytes() int32`

GetMemoryMbytes returns the MemoryMbytes field if non-nil, zero value otherwise.

### GetMemoryMbytesOk

`func (o *TaskOptions) GetMemoryMbytesOk() (*int32, bool)`

GetMemoryMbytesOk returns a tuple with the MemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMbytes

`func (o *TaskOptions) SetMemoryMbytes(v int32)`

SetMemoryMbytes sets MemoryMbytes field to given value.

### HasMemoryMbytes

`func (o *TaskOptions) HasMemoryMbytes() bool`

HasMemoryMbytes returns a boolean if a field has been set.

### SetMemoryMbytesNil

`func (o *TaskOptions) SetMemoryMbytesNil(b bool)`

 SetMemoryMbytesNil sets the value for MemoryMbytes to be an explicit nil

### UnsetMemoryMbytes
`func (o *TaskOptions) UnsetMemoryMbytes()`

UnsetMemoryMbytes ensures that no value is present for MemoryMbytes, not even an explicit nil
### GetMaxItems

`func (o *TaskOptions) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *TaskOptions) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *TaskOptions) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *TaskOptions) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *TaskOptions) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *TaskOptions) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetMaxTotalChargeUsd

`func (o *TaskOptions) GetMaxTotalChargeUsd() float32`

GetMaxTotalChargeUsd returns the MaxTotalChargeUsd field if non-nil, zero value otherwise.

### GetMaxTotalChargeUsdOk

`func (o *TaskOptions) GetMaxTotalChargeUsdOk() (*float32, bool)`

GetMaxTotalChargeUsdOk returns a tuple with the MaxTotalChargeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalChargeUsd

`func (o *TaskOptions) SetMaxTotalChargeUsd(v float32)`

SetMaxTotalChargeUsd sets MaxTotalChargeUsd field to given value.

### HasMaxTotalChargeUsd

`func (o *TaskOptions) HasMaxTotalChargeUsd() bool`

HasMaxTotalChargeUsd returns a boolean if a field has been set.

### SetMaxTotalChargeUsdNil

`func (o *TaskOptions) SetMaxTotalChargeUsdNil(b bool)`

 SetMaxTotalChargeUsdNil sets the value for MaxTotalChargeUsd to be an explicit nil

### UnsetMaxTotalChargeUsd
`func (o *TaskOptions) UnsetMaxTotalChargeUsd()`

UnsetMaxTotalChargeUsd ensures that no value is present for MaxTotalChargeUsd, not even an explicit nil
### GetRestartOnError

`func (o *TaskOptions) GetRestartOnError() bool`

GetRestartOnError returns the RestartOnError field if non-nil, zero value otherwise.

### GetRestartOnErrorOk

`func (o *TaskOptions) GetRestartOnErrorOk() (*bool, bool)`

GetRestartOnErrorOk returns a tuple with the RestartOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnError

`func (o *TaskOptions) SetRestartOnError(v bool)`

SetRestartOnError sets RestartOnError field to given value.

### HasRestartOnError

`func (o *TaskOptions) HasRestartOnError() bool`

HasRestartOnError returns a boolean if a field has been set.

### SetRestartOnErrorNil

`func (o *TaskOptions) SetRestartOnErrorNil(b bool)`

 SetRestartOnErrorNil sets the value for RestartOnError to be an explicit nil

### UnsetRestartOnError
`func (o *TaskOptions) UnsetRestartOnError()`

UnsetRestartOnError ensures that no value is present for RestartOnError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


