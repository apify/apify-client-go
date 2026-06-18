# Current

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthlyUsageUsd** | **float32** |  | 
**MonthlyActorComputeUnits** | **float32** |  | 
**MonthlyExternalDataTransferGbytes** | **float32** |  | 
**MonthlyProxySerps** | **int32** |  | 
**MonthlyResidentialProxyGbytes** | **float32** |  | 
**ActorMemoryGbytes** | **float32** |  | 
**ActorCount** | **int32** |  | 
**ActorTaskCount** | **int32** |  | 
**ActiveActorJobCount** | **int32** |  | 
**TeamAccountSeatCount** | **int32** |  | 
**ScheduleCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCurrent

`func NewCurrent(monthlyUsageUsd float32, monthlyActorComputeUnits float32, monthlyExternalDataTransferGbytes float32, monthlyProxySerps int32, monthlyResidentialProxyGbytes float32, actorMemoryGbytes float32, actorCount int32, actorTaskCount int32, activeActorJobCount int32, teamAccountSeatCount int32, ) *Current`

NewCurrent instantiates a new Current object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentWithDefaults

`func NewCurrentWithDefaults() *Current`

NewCurrentWithDefaults instantiates a new Current object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthlyUsageUsd

`func (o *Current) GetMonthlyUsageUsd() float32`

GetMonthlyUsageUsd returns the MonthlyUsageUsd field if non-nil, zero value otherwise.

### GetMonthlyUsageUsdOk

`func (o *Current) GetMonthlyUsageUsdOk() (*float32, bool)`

GetMonthlyUsageUsdOk returns a tuple with the MonthlyUsageUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyUsageUsd

`func (o *Current) SetMonthlyUsageUsd(v float32)`

SetMonthlyUsageUsd sets MonthlyUsageUsd field to given value.


### GetMonthlyActorComputeUnits

`func (o *Current) GetMonthlyActorComputeUnits() float32`

GetMonthlyActorComputeUnits returns the MonthlyActorComputeUnits field if non-nil, zero value otherwise.

### GetMonthlyActorComputeUnitsOk

`func (o *Current) GetMonthlyActorComputeUnitsOk() (*float32, bool)`

GetMonthlyActorComputeUnitsOk returns a tuple with the MonthlyActorComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyActorComputeUnits

`func (o *Current) SetMonthlyActorComputeUnits(v float32)`

SetMonthlyActorComputeUnits sets MonthlyActorComputeUnits field to given value.


### GetMonthlyExternalDataTransferGbytes

`func (o *Current) GetMonthlyExternalDataTransferGbytes() float32`

GetMonthlyExternalDataTransferGbytes returns the MonthlyExternalDataTransferGbytes field if non-nil, zero value otherwise.

### GetMonthlyExternalDataTransferGbytesOk

`func (o *Current) GetMonthlyExternalDataTransferGbytesOk() (*float32, bool)`

GetMonthlyExternalDataTransferGbytesOk returns a tuple with the MonthlyExternalDataTransferGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyExternalDataTransferGbytes

`func (o *Current) SetMonthlyExternalDataTransferGbytes(v float32)`

SetMonthlyExternalDataTransferGbytes sets MonthlyExternalDataTransferGbytes field to given value.


### GetMonthlyProxySerps

`func (o *Current) GetMonthlyProxySerps() int32`

GetMonthlyProxySerps returns the MonthlyProxySerps field if non-nil, zero value otherwise.

### GetMonthlyProxySerpsOk

`func (o *Current) GetMonthlyProxySerpsOk() (*int32, bool)`

GetMonthlyProxySerpsOk returns a tuple with the MonthlyProxySerps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyProxySerps

`func (o *Current) SetMonthlyProxySerps(v int32)`

SetMonthlyProxySerps sets MonthlyProxySerps field to given value.


### GetMonthlyResidentialProxyGbytes

`func (o *Current) GetMonthlyResidentialProxyGbytes() float32`

GetMonthlyResidentialProxyGbytes returns the MonthlyResidentialProxyGbytes field if non-nil, zero value otherwise.

### GetMonthlyResidentialProxyGbytesOk

`func (o *Current) GetMonthlyResidentialProxyGbytesOk() (*float32, bool)`

GetMonthlyResidentialProxyGbytesOk returns a tuple with the MonthlyResidentialProxyGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyResidentialProxyGbytes

`func (o *Current) SetMonthlyResidentialProxyGbytes(v float32)`

SetMonthlyResidentialProxyGbytes sets MonthlyResidentialProxyGbytes field to given value.


### GetActorMemoryGbytes

`func (o *Current) GetActorMemoryGbytes() float32`

GetActorMemoryGbytes returns the ActorMemoryGbytes field if non-nil, zero value otherwise.

### GetActorMemoryGbytesOk

`func (o *Current) GetActorMemoryGbytesOk() (*float32, bool)`

GetActorMemoryGbytesOk returns a tuple with the ActorMemoryGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorMemoryGbytes

`func (o *Current) SetActorMemoryGbytes(v float32)`

SetActorMemoryGbytes sets ActorMemoryGbytes field to given value.


### GetActorCount

`func (o *Current) GetActorCount() int32`

GetActorCount returns the ActorCount field if non-nil, zero value otherwise.

### GetActorCountOk

`func (o *Current) GetActorCountOk() (*int32, bool)`

GetActorCountOk returns a tuple with the ActorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorCount

`func (o *Current) SetActorCount(v int32)`

SetActorCount sets ActorCount field to given value.


### GetActorTaskCount

`func (o *Current) GetActorTaskCount() int32`

GetActorTaskCount returns the ActorTaskCount field if non-nil, zero value otherwise.

### GetActorTaskCountOk

`func (o *Current) GetActorTaskCountOk() (*int32, bool)`

GetActorTaskCountOk returns a tuple with the ActorTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTaskCount

`func (o *Current) SetActorTaskCount(v int32)`

SetActorTaskCount sets ActorTaskCount field to given value.


### GetActiveActorJobCount

`func (o *Current) GetActiveActorJobCount() int32`

GetActiveActorJobCount returns the ActiveActorJobCount field if non-nil, zero value otherwise.

### GetActiveActorJobCountOk

`func (o *Current) GetActiveActorJobCountOk() (*int32, bool)`

GetActiveActorJobCountOk returns a tuple with the ActiveActorJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActorJobCount

`func (o *Current) SetActiveActorJobCount(v int32)`

SetActiveActorJobCount sets ActiveActorJobCount field to given value.


### GetTeamAccountSeatCount

`func (o *Current) GetTeamAccountSeatCount() int32`

GetTeamAccountSeatCount returns the TeamAccountSeatCount field if non-nil, zero value otherwise.

### GetTeamAccountSeatCountOk

`func (o *Current) GetTeamAccountSeatCountOk() (*int32, bool)`

GetTeamAccountSeatCountOk returns a tuple with the TeamAccountSeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAccountSeatCount

`func (o *Current) SetTeamAccountSeatCount(v int32)`

SetTeamAccountSeatCount sets TeamAccountSeatCount field to given value.


### GetScheduleCount

`func (o *Current) GetScheduleCount() int32`

GetScheduleCount returns the ScheduleCount field if non-nil, zero value otherwise.

### GetScheduleCountOk

`func (o *Current) GetScheduleCountOk() (*int32, bool)`

GetScheduleCountOk returns a tuple with the ScheduleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCount

`func (o *Current) SetScheduleCount(v int32)`

SetScheduleCount sets ScheduleCount field to given value.

### HasScheduleCount

`func (o *Current) HasScheduleCount() bool`

HasScheduleCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


