# DailyServiceUsages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**ServiceUsage** | [**map[string]UsageItem**](UsageItem.md) | A map of service usage item names to their usage details. | 
**TotalUsageCreditsUsd** | **float32** |  | 

## Methods

### NewDailyServiceUsages

`func NewDailyServiceUsages(date string, serviceUsage map[string]UsageItem, totalUsageCreditsUsd float32, ) *DailyServiceUsages`

NewDailyServiceUsages instantiates a new DailyServiceUsages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyServiceUsagesWithDefaults

`func NewDailyServiceUsagesWithDefaults() *DailyServiceUsages`

NewDailyServiceUsagesWithDefaults instantiates a new DailyServiceUsages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DailyServiceUsages) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DailyServiceUsages) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DailyServiceUsages) SetDate(v string)`

SetDate sets Date field to given value.


### GetServiceUsage

`func (o *DailyServiceUsages) GetServiceUsage() map[string]UsageItem`

GetServiceUsage returns the ServiceUsage field if non-nil, zero value otherwise.

### GetServiceUsageOk

`func (o *DailyServiceUsages) GetServiceUsageOk() (*map[string]UsageItem, bool)`

GetServiceUsageOk returns a tuple with the ServiceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsage

`func (o *DailyServiceUsages) SetServiceUsage(v map[string]UsageItem)`

SetServiceUsage sets ServiceUsage field to given value.


### GetTotalUsageCreditsUsd

`func (o *DailyServiceUsages) GetTotalUsageCreditsUsd() float32`

GetTotalUsageCreditsUsd returns the TotalUsageCreditsUsd field if non-nil, zero value otherwise.

### GetTotalUsageCreditsUsdOk

`func (o *DailyServiceUsages) GetTotalUsageCreditsUsdOk() (*float32, bool)`

GetTotalUsageCreditsUsdOk returns a tuple with the TotalUsageCreditsUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCreditsUsd

`func (o *DailyServiceUsages) SetTotalUsageCreditsUsd(v float32)`

SetTotalUsageCreditsUsd sets TotalUsageCreditsUsd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


