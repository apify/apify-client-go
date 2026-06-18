# MonthlyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageCycle** | [**UsageCycle**](UsageCycle.md) |  | 
**MonthlyServiceUsage** | [**map[string]UsageItem**](UsageItem.md) | A map of usage item names (e.g., ACTOR_COMPUTE_UNITS) to their usage details. | 
**DailyServiceUsages** | [**[]DailyServiceUsages**](DailyServiceUsages.md) |  | 
**TotalUsageCreditsUsdBeforeVolumeDiscount** | **float32** |  | 
**TotalUsageCreditsUsdAfterVolumeDiscount** | **float32** |  | 

## Methods

### NewMonthlyUsage

`func NewMonthlyUsage(usageCycle UsageCycle, monthlyServiceUsage map[string]UsageItem, dailyServiceUsages []DailyServiceUsages, totalUsageCreditsUsdBeforeVolumeDiscount float32, totalUsageCreditsUsdAfterVolumeDiscount float32, ) *MonthlyUsage`

NewMonthlyUsage instantiates a new MonthlyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyUsageWithDefaults

`func NewMonthlyUsageWithDefaults() *MonthlyUsage`

NewMonthlyUsageWithDefaults instantiates a new MonthlyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageCycle

`func (o *MonthlyUsage) GetUsageCycle() UsageCycle`

GetUsageCycle returns the UsageCycle field if non-nil, zero value otherwise.

### GetUsageCycleOk

`func (o *MonthlyUsage) GetUsageCycleOk() (*UsageCycle, bool)`

GetUsageCycleOk returns a tuple with the UsageCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCycle

`func (o *MonthlyUsage) SetUsageCycle(v UsageCycle)`

SetUsageCycle sets UsageCycle field to given value.


### GetMonthlyServiceUsage

`func (o *MonthlyUsage) GetMonthlyServiceUsage() map[string]UsageItem`

GetMonthlyServiceUsage returns the MonthlyServiceUsage field if non-nil, zero value otherwise.

### GetMonthlyServiceUsageOk

`func (o *MonthlyUsage) GetMonthlyServiceUsageOk() (*map[string]UsageItem, bool)`

GetMonthlyServiceUsageOk returns a tuple with the MonthlyServiceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyServiceUsage

`func (o *MonthlyUsage) SetMonthlyServiceUsage(v map[string]UsageItem)`

SetMonthlyServiceUsage sets MonthlyServiceUsage field to given value.


### GetDailyServiceUsages

`func (o *MonthlyUsage) GetDailyServiceUsages() []DailyServiceUsages`

GetDailyServiceUsages returns the DailyServiceUsages field if non-nil, zero value otherwise.

### GetDailyServiceUsagesOk

`func (o *MonthlyUsage) GetDailyServiceUsagesOk() (*[]DailyServiceUsages, bool)`

GetDailyServiceUsagesOk returns a tuple with the DailyServiceUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyServiceUsages

`func (o *MonthlyUsage) SetDailyServiceUsages(v []DailyServiceUsages)`

SetDailyServiceUsages sets DailyServiceUsages field to given value.


### GetTotalUsageCreditsUsdBeforeVolumeDiscount

`func (o *MonthlyUsage) GetTotalUsageCreditsUsdBeforeVolumeDiscount() float32`

GetTotalUsageCreditsUsdBeforeVolumeDiscount returns the TotalUsageCreditsUsdBeforeVolumeDiscount field if non-nil, zero value otherwise.

### GetTotalUsageCreditsUsdBeforeVolumeDiscountOk

`func (o *MonthlyUsage) GetTotalUsageCreditsUsdBeforeVolumeDiscountOk() (*float32, bool)`

GetTotalUsageCreditsUsdBeforeVolumeDiscountOk returns a tuple with the TotalUsageCreditsUsdBeforeVolumeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCreditsUsdBeforeVolumeDiscount

`func (o *MonthlyUsage) SetTotalUsageCreditsUsdBeforeVolumeDiscount(v float32)`

SetTotalUsageCreditsUsdBeforeVolumeDiscount sets TotalUsageCreditsUsdBeforeVolumeDiscount field to given value.


### GetTotalUsageCreditsUsdAfterVolumeDiscount

`func (o *MonthlyUsage) GetTotalUsageCreditsUsdAfterVolumeDiscount() float32`

GetTotalUsageCreditsUsdAfterVolumeDiscount returns the TotalUsageCreditsUsdAfterVolumeDiscount field if non-nil, zero value otherwise.

### GetTotalUsageCreditsUsdAfterVolumeDiscountOk

`func (o *MonthlyUsage) GetTotalUsageCreditsUsdAfterVolumeDiscountOk() (*float32, bool)`

GetTotalUsageCreditsUsdAfterVolumeDiscountOk returns a tuple with the TotalUsageCreditsUsdAfterVolumeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCreditsUsdAfterVolumeDiscount

`func (o *MonthlyUsage) SetTotalUsageCreditsUsdAfterVolumeDiscount(v float32)`

SetTotalUsageCreditsUsdAfterVolumeDiscount sets TotalUsageCreditsUsdAfterVolumeDiscount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


