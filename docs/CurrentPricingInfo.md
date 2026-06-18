# CurrentPricingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricingModel** | **string** |  | 
**ApifyMarginPercentage** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**NotifiedAboutChangeAt** | Pointer to **NullableTime** |  | [optional] 
**NotifiedAboutFutureChangeAt** | Pointer to **NullableTime** |  | [optional] 
**IsPriceChangeNotificationSuppressed** | Pointer to **bool** |  | [optional] 
**ForceContainsSignificantPriceChange** | Pointer to **bool** |  | [optional] 
**IsPPEPlatformUsagePaidByUser** | Pointer to **bool** |  | [optional] 
**ReasonForChange** | Pointer to **NullableString** |  | [optional] 
**TrialMinutes** | Pointer to **NullableInt32** |  | [optional] 
**UnitName** | Pointer to **NullableString** |  | [optional] 
**PricePerUnitUsd** | Pointer to **NullableFloat32** |  | [optional] 
**MinimalMaxTotalChargeUsd** | Pointer to **NullableFloat32** |  | [optional] 
**PricingPerEvent** | Pointer to **map[string]interface{}** | Per-event pricing configuration for pay-per-event Actors. | [optional] 

## Methods

### NewCurrentPricingInfo

`func NewCurrentPricingInfo(pricingModel string, ) *CurrentPricingInfo`

NewCurrentPricingInfo instantiates a new CurrentPricingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentPricingInfoWithDefaults

`func NewCurrentPricingInfoWithDefaults() *CurrentPricingInfo`

NewCurrentPricingInfoWithDefaults instantiates a new CurrentPricingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricingModel

`func (o *CurrentPricingInfo) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *CurrentPricingInfo) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *CurrentPricingInfo) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetApifyMarginPercentage

`func (o *CurrentPricingInfo) GetApifyMarginPercentage() float32`

GetApifyMarginPercentage returns the ApifyMarginPercentage field if non-nil, zero value otherwise.

### GetApifyMarginPercentageOk

`func (o *CurrentPricingInfo) GetApifyMarginPercentageOk() (*float32, bool)`

GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApifyMarginPercentage

`func (o *CurrentPricingInfo) SetApifyMarginPercentage(v float32)`

SetApifyMarginPercentage sets ApifyMarginPercentage field to given value.

### HasApifyMarginPercentage

`func (o *CurrentPricingInfo) HasApifyMarginPercentage() bool`

HasApifyMarginPercentage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CurrentPricingInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CurrentPricingInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CurrentPricingInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CurrentPricingInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *CurrentPricingInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CurrentPricingInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CurrentPricingInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CurrentPricingInfo) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetNotifiedAboutChangeAt

`func (o *CurrentPricingInfo) GetNotifiedAboutChangeAt() time.Time`

GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutChangeAtOk

`func (o *CurrentPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutChangeAt

`func (o *CurrentPricingInfo) SetNotifiedAboutChangeAt(v time.Time)`

SetNotifiedAboutChangeAt sets NotifiedAboutChangeAt field to given value.

### HasNotifiedAboutChangeAt

`func (o *CurrentPricingInfo) HasNotifiedAboutChangeAt() bool`

HasNotifiedAboutChangeAt returns a boolean if a field has been set.

### SetNotifiedAboutChangeAtNil

`func (o *CurrentPricingInfo) SetNotifiedAboutChangeAtNil(b bool)`

 SetNotifiedAboutChangeAtNil sets the value for NotifiedAboutChangeAt to be an explicit nil

### UnsetNotifiedAboutChangeAt
`func (o *CurrentPricingInfo) UnsetNotifiedAboutChangeAt()`

UnsetNotifiedAboutChangeAt ensures that no value is present for NotifiedAboutChangeAt, not even an explicit nil
### GetNotifiedAboutFutureChangeAt

`func (o *CurrentPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time`

GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutFutureChangeAtOk

`func (o *CurrentPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutFutureChangeAt

`func (o *CurrentPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time)`

SetNotifiedAboutFutureChangeAt sets NotifiedAboutFutureChangeAt field to given value.

### HasNotifiedAboutFutureChangeAt

`func (o *CurrentPricingInfo) HasNotifiedAboutFutureChangeAt() bool`

HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.

### SetNotifiedAboutFutureChangeAtNil

`func (o *CurrentPricingInfo) SetNotifiedAboutFutureChangeAtNil(b bool)`

 SetNotifiedAboutFutureChangeAtNil sets the value for NotifiedAboutFutureChangeAt to be an explicit nil

### UnsetNotifiedAboutFutureChangeAt
`func (o *CurrentPricingInfo) UnsetNotifiedAboutFutureChangeAt()`

UnsetNotifiedAboutFutureChangeAt ensures that no value is present for NotifiedAboutFutureChangeAt, not even an explicit nil
### GetIsPriceChangeNotificationSuppressed

`func (o *CurrentPricingInfo) GetIsPriceChangeNotificationSuppressed() bool`

GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field if non-nil, zero value otherwise.

### GetIsPriceChangeNotificationSuppressedOk

`func (o *CurrentPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool)`

GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceChangeNotificationSuppressed

`func (o *CurrentPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool)`

SetIsPriceChangeNotificationSuppressed sets IsPriceChangeNotificationSuppressed field to given value.

### HasIsPriceChangeNotificationSuppressed

`func (o *CurrentPricingInfo) HasIsPriceChangeNotificationSuppressed() bool`

HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.

### GetForceContainsSignificantPriceChange

`func (o *CurrentPricingInfo) GetForceContainsSignificantPriceChange() bool`

GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field if non-nil, zero value otherwise.

### GetForceContainsSignificantPriceChangeOk

`func (o *CurrentPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool)`

GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceContainsSignificantPriceChange

`func (o *CurrentPricingInfo) SetForceContainsSignificantPriceChange(v bool)`

SetForceContainsSignificantPriceChange sets ForceContainsSignificantPriceChange field to given value.

### HasForceContainsSignificantPriceChange

`func (o *CurrentPricingInfo) HasForceContainsSignificantPriceChange() bool`

HasForceContainsSignificantPriceChange returns a boolean if a field has been set.

### GetIsPPEPlatformUsagePaidByUser

`func (o *CurrentPricingInfo) GetIsPPEPlatformUsagePaidByUser() bool`

GetIsPPEPlatformUsagePaidByUser returns the IsPPEPlatformUsagePaidByUser field if non-nil, zero value otherwise.

### GetIsPPEPlatformUsagePaidByUserOk

`func (o *CurrentPricingInfo) GetIsPPEPlatformUsagePaidByUserOk() (*bool, bool)`

GetIsPPEPlatformUsagePaidByUserOk returns a tuple with the IsPPEPlatformUsagePaidByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPPEPlatformUsagePaidByUser

`func (o *CurrentPricingInfo) SetIsPPEPlatformUsagePaidByUser(v bool)`

SetIsPPEPlatformUsagePaidByUser sets IsPPEPlatformUsagePaidByUser field to given value.

### HasIsPPEPlatformUsagePaidByUser

`func (o *CurrentPricingInfo) HasIsPPEPlatformUsagePaidByUser() bool`

HasIsPPEPlatformUsagePaidByUser returns a boolean if a field has been set.

### GetReasonForChange

`func (o *CurrentPricingInfo) GetReasonForChange() string`

GetReasonForChange returns the ReasonForChange field if non-nil, zero value otherwise.

### GetReasonForChangeOk

`func (o *CurrentPricingInfo) GetReasonForChangeOk() (*string, bool)`

GetReasonForChangeOk returns a tuple with the ReasonForChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForChange

`func (o *CurrentPricingInfo) SetReasonForChange(v string)`

SetReasonForChange sets ReasonForChange field to given value.

### HasReasonForChange

`func (o *CurrentPricingInfo) HasReasonForChange() bool`

HasReasonForChange returns a boolean if a field has been set.

### SetReasonForChangeNil

`func (o *CurrentPricingInfo) SetReasonForChangeNil(b bool)`

 SetReasonForChangeNil sets the value for ReasonForChange to be an explicit nil

### UnsetReasonForChange
`func (o *CurrentPricingInfo) UnsetReasonForChange()`

UnsetReasonForChange ensures that no value is present for ReasonForChange, not even an explicit nil
### GetTrialMinutes

`func (o *CurrentPricingInfo) GetTrialMinutes() int32`

GetTrialMinutes returns the TrialMinutes field if non-nil, zero value otherwise.

### GetTrialMinutesOk

`func (o *CurrentPricingInfo) GetTrialMinutesOk() (*int32, bool)`

GetTrialMinutesOk returns a tuple with the TrialMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialMinutes

`func (o *CurrentPricingInfo) SetTrialMinutes(v int32)`

SetTrialMinutes sets TrialMinutes field to given value.

### HasTrialMinutes

`func (o *CurrentPricingInfo) HasTrialMinutes() bool`

HasTrialMinutes returns a boolean if a field has been set.

### SetTrialMinutesNil

`func (o *CurrentPricingInfo) SetTrialMinutesNil(b bool)`

 SetTrialMinutesNil sets the value for TrialMinutes to be an explicit nil

### UnsetTrialMinutes
`func (o *CurrentPricingInfo) UnsetTrialMinutes()`

UnsetTrialMinutes ensures that no value is present for TrialMinutes, not even an explicit nil
### GetUnitName

`func (o *CurrentPricingInfo) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *CurrentPricingInfo) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *CurrentPricingInfo) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *CurrentPricingInfo) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### SetUnitNameNil

`func (o *CurrentPricingInfo) SetUnitNameNil(b bool)`

 SetUnitNameNil sets the value for UnitName to be an explicit nil

### UnsetUnitName
`func (o *CurrentPricingInfo) UnsetUnitName()`

UnsetUnitName ensures that no value is present for UnitName, not even an explicit nil
### GetPricePerUnitUsd

`func (o *CurrentPricingInfo) GetPricePerUnitUsd() float32`

GetPricePerUnitUsd returns the PricePerUnitUsd field if non-nil, zero value otherwise.

### GetPricePerUnitUsdOk

`func (o *CurrentPricingInfo) GetPricePerUnitUsdOk() (*float32, bool)`

GetPricePerUnitUsdOk returns a tuple with the PricePerUnitUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnitUsd

`func (o *CurrentPricingInfo) SetPricePerUnitUsd(v float32)`

SetPricePerUnitUsd sets PricePerUnitUsd field to given value.

### HasPricePerUnitUsd

`func (o *CurrentPricingInfo) HasPricePerUnitUsd() bool`

HasPricePerUnitUsd returns a boolean if a field has been set.

### SetPricePerUnitUsdNil

`func (o *CurrentPricingInfo) SetPricePerUnitUsdNil(b bool)`

 SetPricePerUnitUsdNil sets the value for PricePerUnitUsd to be an explicit nil

### UnsetPricePerUnitUsd
`func (o *CurrentPricingInfo) UnsetPricePerUnitUsd()`

UnsetPricePerUnitUsd ensures that no value is present for PricePerUnitUsd, not even an explicit nil
### GetMinimalMaxTotalChargeUsd

`func (o *CurrentPricingInfo) GetMinimalMaxTotalChargeUsd() float32`

GetMinimalMaxTotalChargeUsd returns the MinimalMaxTotalChargeUsd field if non-nil, zero value otherwise.

### GetMinimalMaxTotalChargeUsdOk

`func (o *CurrentPricingInfo) GetMinimalMaxTotalChargeUsdOk() (*float32, bool)`

GetMinimalMaxTotalChargeUsdOk returns a tuple with the MinimalMaxTotalChargeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalMaxTotalChargeUsd

`func (o *CurrentPricingInfo) SetMinimalMaxTotalChargeUsd(v float32)`

SetMinimalMaxTotalChargeUsd sets MinimalMaxTotalChargeUsd field to given value.

### HasMinimalMaxTotalChargeUsd

`func (o *CurrentPricingInfo) HasMinimalMaxTotalChargeUsd() bool`

HasMinimalMaxTotalChargeUsd returns a boolean if a field has been set.

### SetMinimalMaxTotalChargeUsdNil

`func (o *CurrentPricingInfo) SetMinimalMaxTotalChargeUsdNil(b bool)`

 SetMinimalMaxTotalChargeUsdNil sets the value for MinimalMaxTotalChargeUsd to be an explicit nil

### UnsetMinimalMaxTotalChargeUsd
`func (o *CurrentPricingInfo) UnsetMinimalMaxTotalChargeUsd()`

UnsetMinimalMaxTotalChargeUsd ensures that no value is present for MinimalMaxTotalChargeUsd, not even an explicit nil
### GetPricingPerEvent

`func (o *CurrentPricingInfo) GetPricingPerEvent() map[string]interface{}`

GetPricingPerEvent returns the PricingPerEvent field if non-nil, zero value otherwise.

### GetPricingPerEventOk

`func (o *CurrentPricingInfo) GetPricingPerEventOk() (*map[string]interface{}, bool)`

GetPricingPerEventOk returns a tuple with the PricingPerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPerEvent

`func (o *CurrentPricingInfo) SetPricingPerEvent(v map[string]interface{})`

SetPricingPerEvent sets PricingPerEvent field to given value.

### HasPricingPerEvent

`func (o *CurrentPricingInfo) HasPricingPerEvent() bool`

HasPricingPerEvent returns a boolean if a field has been set.

### SetPricingPerEventNil

`func (o *CurrentPricingInfo) SetPricingPerEventNil(b bool)`

 SetPricingPerEventNil sets the value for PricingPerEvent to be an explicit nil

### UnsetPricingPerEvent
`func (o *CurrentPricingInfo) UnsetPricingPerEvent()`

UnsetPricingPerEvent ensures that no value is present for PricingPerEvent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


