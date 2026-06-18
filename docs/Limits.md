# Limits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMonthlyUsageUsd** | **float32** |  | 
**MaxMonthlyActorComputeUnits** | **float32** |  | 
**MaxMonthlyExternalDataTransferGbytes** | **float32** |  | 
**MaxMonthlyProxySerps** | **int32** |  | 
**MaxMonthlyResidentialProxyGbytes** | **float32** |  | 
**MaxActorMemoryGbytes** | **float32** |  | 
**MaxActorCount** | **int32** |  | 
**MaxActorTaskCount** | **int32** |  | 
**MaxConcurrentActorJobs** | **int32** |  | 
**MaxTeamAccountSeatCount** | **int32** |  | 
**DataRetentionDays** | **int32** |  | 
**MaxScheduleCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewLimits

`func NewLimits(maxMonthlyUsageUsd float32, maxMonthlyActorComputeUnits float32, maxMonthlyExternalDataTransferGbytes float32, maxMonthlyProxySerps int32, maxMonthlyResidentialProxyGbytes float32, maxActorMemoryGbytes float32, maxActorCount int32, maxActorTaskCount int32, maxConcurrentActorJobs int32, maxTeamAccountSeatCount int32, dataRetentionDays int32, ) *Limits`

NewLimits instantiates a new Limits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitsWithDefaults

`func NewLimitsWithDefaults() *Limits`

NewLimitsWithDefaults instantiates a new Limits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMonthlyUsageUsd

`func (o *Limits) GetMaxMonthlyUsageUsd() float32`

GetMaxMonthlyUsageUsd returns the MaxMonthlyUsageUsd field if non-nil, zero value otherwise.

### GetMaxMonthlyUsageUsdOk

`func (o *Limits) GetMaxMonthlyUsageUsdOk() (*float32, bool)`

GetMaxMonthlyUsageUsdOk returns a tuple with the MaxMonthlyUsageUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyUsageUsd

`func (o *Limits) SetMaxMonthlyUsageUsd(v float32)`

SetMaxMonthlyUsageUsd sets MaxMonthlyUsageUsd field to given value.


### GetMaxMonthlyActorComputeUnits

`func (o *Limits) GetMaxMonthlyActorComputeUnits() float32`

GetMaxMonthlyActorComputeUnits returns the MaxMonthlyActorComputeUnits field if non-nil, zero value otherwise.

### GetMaxMonthlyActorComputeUnitsOk

`func (o *Limits) GetMaxMonthlyActorComputeUnitsOk() (*float32, bool)`

GetMaxMonthlyActorComputeUnitsOk returns a tuple with the MaxMonthlyActorComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyActorComputeUnits

`func (o *Limits) SetMaxMonthlyActorComputeUnits(v float32)`

SetMaxMonthlyActorComputeUnits sets MaxMonthlyActorComputeUnits field to given value.


### GetMaxMonthlyExternalDataTransferGbytes

`func (o *Limits) GetMaxMonthlyExternalDataTransferGbytes() float32`

GetMaxMonthlyExternalDataTransferGbytes returns the MaxMonthlyExternalDataTransferGbytes field if non-nil, zero value otherwise.

### GetMaxMonthlyExternalDataTransferGbytesOk

`func (o *Limits) GetMaxMonthlyExternalDataTransferGbytesOk() (*float32, bool)`

GetMaxMonthlyExternalDataTransferGbytesOk returns a tuple with the MaxMonthlyExternalDataTransferGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyExternalDataTransferGbytes

`func (o *Limits) SetMaxMonthlyExternalDataTransferGbytes(v float32)`

SetMaxMonthlyExternalDataTransferGbytes sets MaxMonthlyExternalDataTransferGbytes field to given value.


### GetMaxMonthlyProxySerps

`func (o *Limits) GetMaxMonthlyProxySerps() int32`

GetMaxMonthlyProxySerps returns the MaxMonthlyProxySerps field if non-nil, zero value otherwise.

### GetMaxMonthlyProxySerpsOk

`func (o *Limits) GetMaxMonthlyProxySerpsOk() (*int32, bool)`

GetMaxMonthlyProxySerpsOk returns a tuple with the MaxMonthlyProxySerps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyProxySerps

`func (o *Limits) SetMaxMonthlyProxySerps(v int32)`

SetMaxMonthlyProxySerps sets MaxMonthlyProxySerps field to given value.


### GetMaxMonthlyResidentialProxyGbytes

`func (o *Limits) GetMaxMonthlyResidentialProxyGbytes() float32`

GetMaxMonthlyResidentialProxyGbytes returns the MaxMonthlyResidentialProxyGbytes field if non-nil, zero value otherwise.

### GetMaxMonthlyResidentialProxyGbytesOk

`func (o *Limits) GetMaxMonthlyResidentialProxyGbytesOk() (*float32, bool)`

GetMaxMonthlyResidentialProxyGbytesOk returns a tuple with the MaxMonthlyResidentialProxyGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyResidentialProxyGbytes

`func (o *Limits) SetMaxMonthlyResidentialProxyGbytes(v float32)`

SetMaxMonthlyResidentialProxyGbytes sets MaxMonthlyResidentialProxyGbytes field to given value.


### GetMaxActorMemoryGbytes

`func (o *Limits) GetMaxActorMemoryGbytes() float32`

GetMaxActorMemoryGbytes returns the MaxActorMemoryGbytes field if non-nil, zero value otherwise.

### GetMaxActorMemoryGbytesOk

`func (o *Limits) GetMaxActorMemoryGbytesOk() (*float32, bool)`

GetMaxActorMemoryGbytesOk returns a tuple with the MaxActorMemoryGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActorMemoryGbytes

`func (o *Limits) SetMaxActorMemoryGbytes(v float32)`

SetMaxActorMemoryGbytes sets MaxActorMemoryGbytes field to given value.


### GetMaxActorCount

`func (o *Limits) GetMaxActorCount() int32`

GetMaxActorCount returns the MaxActorCount field if non-nil, zero value otherwise.

### GetMaxActorCountOk

`func (o *Limits) GetMaxActorCountOk() (*int32, bool)`

GetMaxActorCountOk returns a tuple with the MaxActorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActorCount

`func (o *Limits) SetMaxActorCount(v int32)`

SetMaxActorCount sets MaxActorCount field to given value.


### GetMaxActorTaskCount

`func (o *Limits) GetMaxActorTaskCount() int32`

GetMaxActorTaskCount returns the MaxActorTaskCount field if non-nil, zero value otherwise.

### GetMaxActorTaskCountOk

`func (o *Limits) GetMaxActorTaskCountOk() (*int32, bool)`

GetMaxActorTaskCountOk returns a tuple with the MaxActorTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActorTaskCount

`func (o *Limits) SetMaxActorTaskCount(v int32)`

SetMaxActorTaskCount sets MaxActorTaskCount field to given value.


### GetMaxConcurrentActorJobs

`func (o *Limits) GetMaxConcurrentActorJobs() int32`

GetMaxConcurrentActorJobs returns the MaxConcurrentActorJobs field if non-nil, zero value otherwise.

### GetMaxConcurrentActorJobsOk

`func (o *Limits) GetMaxConcurrentActorJobsOk() (*int32, bool)`

GetMaxConcurrentActorJobsOk returns a tuple with the MaxConcurrentActorJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentActorJobs

`func (o *Limits) SetMaxConcurrentActorJobs(v int32)`

SetMaxConcurrentActorJobs sets MaxConcurrentActorJobs field to given value.


### GetMaxTeamAccountSeatCount

`func (o *Limits) GetMaxTeamAccountSeatCount() int32`

GetMaxTeamAccountSeatCount returns the MaxTeamAccountSeatCount field if non-nil, zero value otherwise.

### GetMaxTeamAccountSeatCountOk

`func (o *Limits) GetMaxTeamAccountSeatCountOk() (*int32, bool)`

GetMaxTeamAccountSeatCountOk returns a tuple with the MaxTeamAccountSeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTeamAccountSeatCount

`func (o *Limits) SetMaxTeamAccountSeatCount(v int32)`

SetMaxTeamAccountSeatCount sets MaxTeamAccountSeatCount field to given value.


### GetDataRetentionDays

`func (o *Limits) GetDataRetentionDays() int32`

GetDataRetentionDays returns the DataRetentionDays field if non-nil, zero value otherwise.

### GetDataRetentionDaysOk

`func (o *Limits) GetDataRetentionDaysOk() (*int32, bool)`

GetDataRetentionDaysOk returns a tuple with the DataRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionDays

`func (o *Limits) SetDataRetentionDays(v int32)`

SetDataRetentionDays sets DataRetentionDays field to given value.


### GetMaxScheduleCount

`func (o *Limits) GetMaxScheduleCount() int32`

GetMaxScheduleCount returns the MaxScheduleCount field if non-nil, zero value otherwise.

### GetMaxScheduleCountOk

`func (o *Limits) GetMaxScheduleCountOk() (*int32, bool)`

GetMaxScheduleCountOk returns a tuple with the MaxScheduleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScheduleCount

`func (o *Limits) SetMaxScheduleCount(v int32)`

SetMaxScheduleCount sets MaxScheduleCount field to given value.

### HasMaxScheduleCount

`func (o *Limits) HasMaxScheduleCount() bool`

HasMaxScheduleCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


