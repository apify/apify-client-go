# RunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputBodyLen** | Pointer to **NullableInt32** |  | [optional] 
**MigrationCount** | Pointer to **int32** |  | [optional] 
**RebootCount** | Pointer to **int32** |  | [optional] 
**RestartCount** | **int32** |  | 
**ResurrectCount** | **int32** |  | 
**MemAvgBytes** | Pointer to **float32** |  | [optional] 
**MemMaxBytes** | Pointer to **int32** |  | [optional] 
**MemCurrentBytes** | Pointer to **int32** |  | [optional] 
**CpuAvgUsage** | Pointer to **float32** |  | [optional] 
**CpuMaxUsage** | Pointer to **float32** |  | [optional] 
**CpuCurrentUsage** | Pointer to **float32** |  | [optional] 
**NetRxBytes** | Pointer to **int32** |  | [optional] 
**NetTxBytes** | Pointer to **int32** |  | [optional] 
**DurationMillis** | Pointer to **int32** |  | [optional] 
**RunTimeSecs** | Pointer to **float32** |  | [optional] 
**Metamorph** | Pointer to **int32** |  | [optional] 
**ComputeUnits** | **float32** |  | 

## Methods

### NewRunStats

`func NewRunStats(restartCount int32, resurrectCount int32, computeUnits float32, ) *RunStats`

NewRunStats instantiates a new RunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunStatsWithDefaults

`func NewRunStatsWithDefaults() *RunStats`

NewRunStatsWithDefaults instantiates a new RunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputBodyLen

`func (o *RunStats) GetInputBodyLen() int32`

GetInputBodyLen returns the InputBodyLen field if non-nil, zero value otherwise.

### GetInputBodyLenOk

`func (o *RunStats) GetInputBodyLenOk() (*int32, bool)`

GetInputBodyLenOk returns a tuple with the InputBodyLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBodyLen

`func (o *RunStats) SetInputBodyLen(v int32)`

SetInputBodyLen sets InputBodyLen field to given value.

### HasInputBodyLen

`func (o *RunStats) HasInputBodyLen() bool`

HasInputBodyLen returns a boolean if a field has been set.

### SetInputBodyLenNil

`func (o *RunStats) SetInputBodyLenNil(b bool)`

 SetInputBodyLenNil sets the value for InputBodyLen to be an explicit nil

### UnsetInputBodyLen
`func (o *RunStats) UnsetInputBodyLen()`

UnsetInputBodyLen ensures that no value is present for InputBodyLen, not even an explicit nil
### GetMigrationCount

`func (o *RunStats) GetMigrationCount() int32`

GetMigrationCount returns the MigrationCount field if non-nil, zero value otherwise.

### GetMigrationCountOk

`func (o *RunStats) GetMigrationCountOk() (*int32, bool)`

GetMigrationCountOk returns a tuple with the MigrationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationCount

`func (o *RunStats) SetMigrationCount(v int32)`

SetMigrationCount sets MigrationCount field to given value.

### HasMigrationCount

`func (o *RunStats) HasMigrationCount() bool`

HasMigrationCount returns a boolean if a field has been set.

### GetRebootCount

`func (o *RunStats) GetRebootCount() int32`

GetRebootCount returns the RebootCount field if non-nil, zero value otherwise.

### GetRebootCountOk

`func (o *RunStats) GetRebootCountOk() (*int32, bool)`

GetRebootCountOk returns a tuple with the RebootCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootCount

`func (o *RunStats) SetRebootCount(v int32)`

SetRebootCount sets RebootCount field to given value.

### HasRebootCount

`func (o *RunStats) HasRebootCount() bool`

HasRebootCount returns a boolean if a field has been set.

### GetRestartCount

`func (o *RunStats) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *RunStats) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *RunStats) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.


### GetResurrectCount

`func (o *RunStats) GetResurrectCount() int32`

GetResurrectCount returns the ResurrectCount field if non-nil, zero value otherwise.

### GetResurrectCountOk

`func (o *RunStats) GetResurrectCountOk() (*int32, bool)`

GetResurrectCountOk returns a tuple with the ResurrectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResurrectCount

`func (o *RunStats) SetResurrectCount(v int32)`

SetResurrectCount sets ResurrectCount field to given value.


### GetMemAvgBytes

`func (o *RunStats) GetMemAvgBytes() float32`

GetMemAvgBytes returns the MemAvgBytes field if non-nil, zero value otherwise.

### GetMemAvgBytesOk

`func (o *RunStats) GetMemAvgBytesOk() (*float32, bool)`

GetMemAvgBytesOk returns a tuple with the MemAvgBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemAvgBytes

`func (o *RunStats) SetMemAvgBytes(v float32)`

SetMemAvgBytes sets MemAvgBytes field to given value.

### HasMemAvgBytes

`func (o *RunStats) HasMemAvgBytes() bool`

HasMemAvgBytes returns a boolean if a field has been set.

### GetMemMaxBytes

`func (o *RunStats) GetMemMaxBytes() int32`

GetMemMaxBytes returns the MemMaxBytes field if non-nil, zero value otherwise.

### GetMemMaxBytesOk

`func (o *RunStats) GetMemMaxBytesOk() (*int32, bool)`

GetMemMaxBytesOk returns a tuple with the MemMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemMaxBytes

`func (o *RunStats) SetMemMaxBytes(v int32)`

SetMemMaxBytes sets MemMaxBytes field to given value.

### HasMemMaxBytes

`func (o *RunStats) HasMemMaxBytes() bool`

HasMemMaxBytes returns a boolean if a field has been set.

### GetMemCurrentBytes

`func (o *RunStats) GetMemCurrentBytes() int32`

GetMemCurrentBytes returns the MemCurrentBytes field if non-nil, zero value otherwise.

### GetMemCurrentBytesOk

`func (o *RunStats) GetMemCurrentBytesOk() (*int32, bool)`

GetMemCurrentBytesOk returns a tuple with the MemCurrentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemCurrentBytes

`func (o *RunStats) SetMemCurrentBytes(v int32)`

SetMemCurrentBytes sets MemCurrentBytes field to given value.

### HasMemCurrentBytes

`func (o *RunStats) HasMemCurrentBytes() bool`

HasMemCurrentBytes returns a boolean if a field has been set.

### GetCpuAvgUsage

`func (o *RunStats) GetCpuAvgUsage() float32`

GetCpuAvgUsage returns the CpuAvgUsage field if non-nil, zero value otherwise.

### GetCpuAvgUsageOk

`func (o *RunStats) GetCpuAvgUsageOk() (*float32, bool)`

GetCpuAvgUsageOk returns a tuple with the CpuAvgUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAvgUsage

`func (o *RunStats) SetCpuAvgUsage(v float32)`

SetCpuAvgUsage sets CpuAvgUsage field to given value.

### HasCpuAvgUsage

`func (o *RunStats) HasCpuAvgUsage() bool`

HasCpuAvgUsage returns a boolean if a field has been set.

### GetCpuMaxUsage

`func (o *RunStats) GetCpuMaxUsage() float32`

GetCpuMaxUsage returns the CpuMaxUsage field if non-nil, zero value otherwise.

### GetCpuMaxUsageOk

`func (o *RunStats) GetCpuMaxUsageOk() (*float32, bool)`

GetCpuMaxUsageOk returns a tuple with the CpuMaxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMaxUsage

`func (o *RunStats) SetCpuMaxUsage(v float32)`

SetCpuMaxUsage sets CpuMaxUsage field to given value.

### HasCpuMaxUsage

`func (o *RunStats) HasCpuMaxUsage() bool`

HasCpuMaxUsage returns a boolean if a field has been set.

### GetCpuCurrentUsage

`func (o *RunStats) GetCpuCurrentUsage() float32`

GetCpuCurrentUsage returns the CpuCurrentUsage field if non-nil, zero value otherwise.

### GetCpuCurrentUsageOk

`func (o *RunStats) GetCpuCurrentUsageOk() (*float32, bool)`

GetCpuCurrentUsageOk returns a tuple with the CpuCurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCurrentUsage

`func (o *RunStats) SetCpuCurrentUsage(v float32)`

SetCpuCurrentUsage sets CpuCurrentUsage field to given value.

### HasCpuCurrentUsage

`func (o *RunStats) HasCpuCurrentUsage() bool`

HasCpuCurrentUsage returns a boolean if a field has been set.

### GetNetRxBytes

`func (o *RunStats) GetNetRxBytes() int32`

GetNetRxBytes returns the NetRxBytes field if non-nil, zero value otherwise.

### GetNetRxBytesOk

`func (o *RunStats) GetNetRxBytesOk() (*int32, bool)`

GetNetRxBytesOk returns a tuple with the NetRxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxBytes

`func (o *RunStats) SetNetRxBytes(v int32)`

SetNetRxBytes sets NetRxBytes field to given value.

### HasNetRxBytes

`func (o *RunStats) HasNetRxBytes() bool`

HasNetRxBytes returns a boolean if a field has been set.

### GetNetTxBytes

`func (o *RunStats) GetNetTxBytes() int32`

GetNetTxBytes returns the NetTxBytes field if non-nil, zero value otherwise.

### GetNetTxBytesOk

`func (o *RunStats) GetNetTxBytesOk() (*int32, bool)`

GetNetTxBytesOk returns a tuple with the NetTxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxBytes

`func (o *RunStats) SetNetTxBytes(v int32)`

SetNetTxBytes sets NetTxBytes field to given value.

### HasNetTxBytes

`func (o *RunStats) HasNetTxBytes() bool`

HasNetTxBytes returns a boolean if a field has been set.

### GetDurationMillis

`func (o *RunStats) GetDurationMillis() int32`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *RunStats) GetDurationMillisOk() (*int32, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *RunStats) SetDurationMillis(v int32)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *RunStats) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetRunTimeSecs

`func (o *RunStats) GetRunTimeSecs() float32`

GetRunTimeSecs returns the RunTimeSecs field if non-nil, zero value otherwise.

### GetRunTimeSecsOk

`func (o *RunStats) GetRunTimeSecsOk() (*float32, bool)`

GetRunTimeSecsOk returns a tuple with the RunTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeSecs

`func (o *RunStats) SetRunTimeSecs(v float32)`

SetRunTimeSecs sets RunTimeSecs field to given value.

### HasRunTimeSecs

`func (o *RunStats) HasRunTimeSecs() bool`

HasRunTimeSecs returns a boolean if a field has been set.

### GetMetamorph

`func (o *RunStats) GetMetamorph() int32`

GetMetamorph returns the Metamorph field if non-nil, zero value otherwise.

### GetMetamorphOk

`func (o *RunStats) GetMetamorphOk() (*int32, bool)`

GetMetamorphOk returns a tuple with the Metamorph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetamorph

`func (o *RunStats) SetMetamorph(v int32)`

SetMetamorph sets Metamorph field to given value.

### HasMetamorph

`func (o *RunStats) HasMetamorph() bool`

HasMetamorph returns a boolean if a field has been set.

### GetComputeUnits

`func (o *RunStats) GetComputeUnits() float32`

GetComputeUnits returns the ComputeUnits field if non-nil, zero value otherwise.

### GetComputeUnitsOk

`func (o *RunStats) GetComputeUnitsOk() (*float32, bool)`

GetComputeUnitsOk returns a tuple with the ComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnits

`func (o *RunStats) SetComputeUnits(v float32)`

SetComputeUnits sets ComputeUnits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


