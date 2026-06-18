# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Description** | **string** |  | 
**IsEnabled** | **bool** |  | 
**MonthlyBasePriceUsd** | **float32** |  | 
**MonthlyUsageCreditsUsd** | **float32** |  | 
**UsageDiscountPercent** | Pointer to **float32** |  | [optional] 
**EnabledPlatformFeatures** | **[]string** |  | 
**MaxMonthlyUsageUsd** | **float32** |  | 
**MaxActorMemoryGbytes** | **float32** |  | 
**MaxMonthlyActorComputeUnits** | **float32** |  | 
**MaxMonthlyResidentialProxyGbytes** | **float32** |  | 
**MaxMonthlyProxySerps** | **int32** |  | 
**MaxMonthlyExternalDataTransferGbytes** | **float32** |  | 
**MaxActorCount** | **int32** |  | 
**MaxActorTaskCount** | **int32** |  | 
**DataRetentionDays** | **int32** |  | 
**AvailableProxyGroups** | **map[string]int32** | A dictionary mapping proxy group names to the number of available proxies in each group. The keys are proxy group names (e.g., \&quot;RESIDENTIAL\&quot;, \&quot;DATACENTER\&quot;) and values are the count of available proxies.  | 
**TeamAccountSeatCount** | **int32** |  | 
**SupportLevel** | **string** |  | 
**AvailableAddOns** | **[]string** |  | 
**Tier** | Pointer to **string** |  | [optional] 
**ApiRateLimitBoosts** | Pointer to **int32** |  | [optional] 
**MaxScheduleCount** | Pointer to **int32** |  | [optional] 
**MaxConcurrentActorRuns** | Pointer to **int32** |  | [optional] 
**PlanPricing** | Pointer to **map[string]interface{}** | Pricing details for this plan. | [optional] 

## Methods

### NewPlan

`func NewPlan(id string, description string, isEnabled bool, monthlyBasePriceUsd float32, monthlyUsageCreditsUsd float32, enabledPlatformFeatures []string, maxMonthlyUsageUsd float32, maxActorMemoryGbytes float32, maxMonthlyActorComputeUnits float32, maxMonthlyResidentialProxyGbytes float32, maxMonthlyProxySerps int32, maxMonthlyExternalDataTransferGbytes float32, maxActorCount int32, maxActorTaskCount int32, dataRetentionDays int32, availableProxyGroups map[string]int32, teamAccountSeatCount int32, supportLevel string, availableAddOns []string, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsEnabled

`func (o *Plan) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Plan) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Plan) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMonthlyBasePriceUsd

`func (o *Plan) GetMonthlyBasePriceUsd() float32`

GetMonthlyBasePriceUsd returns the MonthlyBasePriceUsd field if non-nil, zero value otherwise.

### GetMonthlyBasePriceUsdOk

`func (o *Plan) GetMonthlyBasePriceUsdOk() (*float32, bool)`

GetMonthlyBasePriceUsdOk returns a tuple with the MonthlyBasePriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyBasePriceUsd

`func (o *Plan) SetMonthlyBasePriceUsd(v float32)`

SetMonthlyBasePriceUsd sets MonthlyBasePriceUsd field to given value.


### GetMonthlyUsageCreditsUsd

`func (o *Plan) GetMonthlyUsageCreditsUsd() float32`

GetMonthlyUsageCreditsUsd returns the MonthlyUsageCreditsUsd field if non-nil, zero value otherwise.

### GetMonthlyUsageCreditsUsdOk

`func (o *Plan) GetMonthlyUsageCreditsUsdOk() (*float32, bool)`

GetMonthlyUsageCreditsUsdOk returns a tuple with the MonthlyUsageCreditsUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyUsageCreditsUsd

`func (o *Plan) SetMonthlyUsageCreditsUsd(v float32)`

SetMonthlyUsageCreditsUsd sets MonthlyUsageCreditsUsd field to given value.


### GetUsageDiscountPercent

`func (o *Plan) GetUsageDiscountPercent() float32`

GetUsageDiscountPercent returns the UsageDiscountPercent field if non-nil, zero value otherwise.

### GetUsageDiscountPercentOk

`func (o *Plan) GetUsageDiscountPercentOk() (*float32, bool)`

GetUsageDiscountPercentOk returns a tuple with the UsageDiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDiscountPercent

`func (o *Plan) SetUsageDiscountPercent(v float32)`

SetUsageDiscountPercent sets UsageDiscountPercent field to given value.

### HasUsageDiscountPercent

`func (o *Plan) HasUsageDiscountPercent() bool`

HasUsageDiscountPercent returns a boolean if a field has been set.

### GetEnabledPlatformFeatures

`func (o *Plan) GetEnabledPlatformFeatures() []string`

GetEnabledPlatformFeatures returns the EnabledPlatformFeatures field if non-nil, zero value otherwise.

### GetEnabledPlatformFeaturesOk

`func (o *Plan) GetEnabledPlatformFeaturesOk() (*[]string, bool)`

GetEnabledPlatformFeaturesOk returns a tuple with the EnabledPlatformFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledPlatformFeatures

`func (o *Plan) SetEnabledPlatformFeatures(v []string)`

SetEnabledPlatformFeatures sets EnabledPlatformFeatures field to given value.


### GetMaxMonthlyUsageUsd

`func (o *Plan) GetMaxMonthlyUsageUsd() float32`

GetMaxMonthlyUsageUsd returns the MaxMonthlyUsageUsd field if non-nil, zero value otherwise.

### GetMaxMonthlyUsageUsdOk

`func (o *Plan) GetMaxMonthlyUsageUsdOk() (*float32, bool)`

GetMaxMonthlyUsageUsdOk returns a tuple with the MaxMonthlyUsageUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyUsageUsd

`func (o *Plan) SetMaxMonthlyUsageUsd(v float32)`

SetMaxMonthlyUsageUsd sets MaxMonthlyUsageUsd field to given value.


### GetMaxActorMemoryGbytes

`func (o *Plan) GetMaxActorMemoryGbytes() float32`

GetMaxActorMemoryGbytes returns the MaxActorMemoryGbytes field if non-nil, zero value otherwise.

### GetMaxActorMemoryGbytesOk

`func (o *Plan) GetMaxActorMemoryGbytesOk() (*float32, bool)`

GetMaxActorMemoryGbytesOk returns a tuple with the MaxActorMemoryGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActorMemoryGbytes

`func (o *Plan) SetMaxActorMemoryGbytes(v float32)`

SetMaxActorMemoryGbytes sets MaxActorMemoryGbytes field to given value.


### GetMaxMonthlyActorComputeUnits

`func (o *Plan) GetMaxMonthlyActorComputeUnits() float32`

GetMaxMonthlyActorComputeUnits returns the MaxMonthlyActorComputeUnits field if non-nil, zero value otherwise.

### GetMaxMonthlyActorComputeUnitsOk

`func (o *Plan) GetMaxMonthlyActorComputeUnitsOk() (*float32, bool)`

GetMaxMonthlyActorComputeUnitsOk returns a tuple with the MaxMonthlyActorComputeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyActorComputeUnits

`func (o *Plan) SetMaxMonthlyActorComputeUnits(v float32)`

SetMaxMonthlyActorComputeUnits sets MaxMonthlyActorComputeUnits field to given value.


### GetMaxMonthlyResidentialProxyGbytes

`func (o *Plan) GetMaxMonthlyResidentialProxyGbytes() float32`

GetMaxMonthlyResidentialProxyGbytes returns the MaxMonthlyResidentialProxyGbytes field if non-nil, zero value otherwise.

### GetMaxMonthlyResidentialProxyGbytesOk

`func (o *Plan) GetMaxMonthlyResidentialProxyGbytesOk() (*float32, bool)`

GetMaxMonthlyResidentialProxyGbytesOk returns a tuple with the MaxMonthlyResidentialProxyGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyResidentialProxyGbytes

`func (o *Plan) SetMaxMonthlyResidentialProxyGbytes(v float32)`

SetMaxMonthlyResidentialProxyGbytes sets MaxMonthlyResidentialProxyGbytes field to given value.


### GetMaxMonthlyProxySerps

`func (o *Plan) GetMaxMonthlyProxySerps() int32`

GetMaxMonthlyProxySerps returns the MaxMonthlyProxySerps field if non-nil, zero value otherwise.

### GetMaxMonthlyProxySerpsOk

`func (o *Plan) GetMaxMonthlyProxySerpsOk() (*int32, bool)`

GetMaxMonthlyProxySerpsOk returns a tuple with the MaxMonthlyProxySerps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyProxySerps

`func (o *Plan) SetMaxMonthlyProxySerps(v int32)`

SetMaxMonthlyProxySerps sets MaxMonthlyProxySerps field to given value.


### GetMaxMonthlyExternalDataTransferGbytes

`func (o *Plan) GetMaxMonthlyExternalDataTransferGbytes() float32`

GetMaxMonthlyExternalDataTransferGbytes returns the MaxMonthlyExternalDataTransferGbytes field if non-nil, zero value otherwise.

### GetMaxMonthlyExternalDataTransferGbytesOk

`func (o *Plan) GetMaxMonthlyExternalDataTransferGbytesOk() (*float32, bool)`

GetMaxMonthlyExternalDataTransferGbytesOk returns a tuple with the MaxMonthlyExternalDataTransferGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyExternalDataTransferGbytes

`func (o *Plan) SetMaxMonthlyExternalDataTransferGbytes(v float32)`

SetMaxMonthlyExternalDataTransferGbytes sets MaxMonthlyExternalDataTransferGbytes field to given value.


### GetMaxActorCount

`func (o *Plan) GetMaxActorCount() int32`

GetMaxActorCount returns the MaxActorCount field if non-nil, zero value otherwise.

### GetMaxActorCountOk

`func (o *Plan) GetMaxActorCountOk() (*int32, bool)`

GetMaxActorCountOk returns a tuple with the MaxActorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActorCount

`func (o *Plan) SetMaxActorCount(v int32)`

SetMaxActorCount sets MaxActorCount field to given value.


### GetMaxActorTaskCount

`func (o *Plan) GetMaxActorTaskCount() int32`

GetMaxActorTaskCount returns the MaxActorTaskCount field if non-nil, zero value otherwise.

### GetMaxActorTaskCountOk

`func (o *Plan) GetMaxActorTaskCountOk() (*int32, bool)`

GetMaxActorTaskCountOk returns a tuple with the MaxActorTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActorTaskCount

`func (o *Plan) SetMaxActorTaskCount(v int32)`

SetMaxActorTaskCount sets MaxActorTaskCount field to given value.


### GetDataRetentionDays

`func (o *Plan) GetDataRetentionDays() int32`

GetDataRetentionDays returns the DataRetentionDays field if non-nil, zero value otherwise.

### GetDataRetentionDaysOk

`func (o *Plan) GetDataRetentionDaysOk() (*int32, bool)`

GetDataRetentionDaysOk returns a tuple with the DataRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionDays

`func (o *Plan) SetDataRetentionDays(v int32)`

SetDataRetentionDays sets DataRetentionDays field to given value.


### GetAvailableProxyGroups

`func (o *Plan) GetAvailableProxyGroups() map[string]int32`

GetAvailableProxyGroups returns the AvailableProxyGroups field if non-nil, zero value otherwise.

### GetAvailableProxyGroupsOk

`func (o *Plan) GetAvailableProxyGroupsOk() (*map[string]int32, bool)`

GetAvailableProxyGroupsOk returns a tuple with the AvailableProxyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableProxyGroups

`func (o *Plan) SetAvailableProxyGroups(v map[string]int32)`

SetAvailableProxyGroups sets AvailableProxyGroups field to given value.


### GetTeamAccountSeatCount

`func (o *Plan) GetTeamAccountSeatCount() int32`

GetTeamAccountSeatCount returns the TeamAccountSeatCount field if non-nil, zero value otherwise.

### GetTeamAccountSeatCountOk

`func (o *Plan) GetTeamAccountSeatCountOk() (*int32, bool)`

GetTeamAccountSeatCountOk returns a tuple with the TeamAccountSeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAccountSeatCount

`func (o *Plan) SetTeamAccountSeatCount(v int32)`

SetTeamAccountSeatCount sets TeamAccountSeatCount field to given value.


### GetSupportLevel

`func (o *Plan) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *Plan) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *Plan) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.


### GetAvailableAddOns

`func (o *Plan) GetAvailableAddOns() []string`

GetAvailableAddOns returns the AvailableAddOns field if non-nil, zero value otherwise.

### GetAvailableAddOnsOk

`func (o *Plan) GetAvailableAddOnsOk() (*[]string, bool)`

GetAvailableAddOnsOk returns a tuple with the AvailableAddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAddOns

`func (o *Plan) SetAvailableAddOns(v []string)`

SetAvailableAddOns sets AvailableAddOns field to given value.


### GetTier

`func (o *Plan) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Plan) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Plan) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *Plan) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetApiRateLimitBoosts

`func (o *Plan) GetApiRateLimitBoosts() int32`

GetApiRateLimitBoosts returns the ApiRateLimitBoosts field if non-nil, zero value otherwise.

### GetApiRateLimitBoostsOk

`func (o *Plan) GetApiRateLimitBoostsOk() (*int32, bool)`

GetApiRateLimitBoostsOk returns a tuple with the ApiRateLimitBoosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRateLimitBoosts

`func (o *Plan) SetApiRateLimitBoosts(v int32)`

SetApiRateLimitBoosts sets ApiRateLimitBoosts field to given value.

### HasApiRateLimitBoosts

`func (o *Plan) HasApiRateLimitBoosts() bool`

HasApiRateLimitBoosts returns a boolean if a field has been set.

### GetMaxScheduleCount

`func (o *Plan) GetMaxScheduleCount() int32`

GetMaxScheduleCount returns the MaxScheduleCount field if non-nil, zero value otherwise.

### GetMaxScheduleCountOk

`func (o *Plan) GetMaxScheduleCountOk() (*int32, bool)`

GetMaxScheduleCountOk returns a tuple with the MaxScheduleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScheduleCount

`func (o *Plan) SetMaxScheduleCount(v int32)`

SetMaxScheduleCount sets MaxScheduleCount field to given value.

### HasMaxScheduleCount

`func (o *Plan) HasMaxScheduleCount() bool`

HasMaxScheduleCount returns a boolean if a field has been set.

### GetMaxConcurrentActorRuns

`func (o *Plan) GetMaxConcurrentActorRuns() int32`

GetMaxConcurrentActorRuns returns the MaxConcurrentActorRuns field if non-nil, zero value otherwise.

### GetMaxConcurrentActorRunsOk

`func (o *Plan) GetMaxConcurrentActorRunsOk() (*int32, bool)`

GetMaxConcurrentActorRunsOk returns a tuple with the MaxConcurrentActorRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentActorRuns

`func (o *Plan) SetMaxConcurrentActorRuns(v int32)`

SetMaxConcurrentActorRuns sets MaxConcurrentActorRuns field to given value.

### HasMaxConcurrentActorRuns

`func (o *Plan) HasMaxConcurrentActorRuns() bool`

HasMaxConcurrentActorRuns returns a boolean if a field has been set.

### GetPlanPricing

`func (o *Plan) GetPlanPricing() map[string]interface{}`

GetPlanPricing returns the PlanPricing field if non-nil, zero value otherwise.

### GetPlanPricingOk

`func (o *Plan) GetPlanPricingOk() (*map[string]interface{}, bool)`

GetPlanPricingOk returns a tuple with the PlanPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPricing

`func (o *Plan) SetPlanPricing(v map[string]interface{})`

SetPlanPricing sets PlanPricing field to given value.

### HasPlanPricing

`func (o *Plan) HasPlanPricing() bool`

HasPlanPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


