# ActorRunPricingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApifyMarginPercentage** | **float32** | In [0, 1], fraction of pricePerUnitUsd that goes to Apify | 
**CreatedAt** | **time.Time** | When this pricing info record has been created | 
**StartedAt** | **time.Time** | Since when is this pricing info record effective for a given Actor | 
**NotifiedAboutFutureChangeAt** | Pointer to **time.Time** |  | [optional] 
**NotifiedAboutChangeAt** | Pointer to **time.Time** |  | [optional] 
**ReasonForChange** | Pointer to **string** |  | [optional] 
**IsPriceChangeNotificationSuppressed** | Pointer to **bool** |  | [optional] 
**ForceContainsSignificantPriceChange** | Pointer to **bool** |  | [optional] 
**PricingModel** | **string** |  | 
**PricingPerEvent** | [**PayPerEventActorPricingInfoAllOfPricingPerEvent**](PayPerEventActorPricingInfoAllOfPricingPerEvent.md) |  | 
**MinimalMaxTotalChargeUsd** | Pointer to **float32** |  | [optional] 
**UnitName** | **string** | Name of the unit that is being charged | 
**PricePerUnitUsd** | **float32** | Monthly flat price in USD | 
**TieredPricing** | Pointer to [**map[string]TieredPricingPerDatasetItemEntry**](TieredPricingPerDatasetItemEntry.md) | Tiered price-per-dataset-item pricing, keyed by subscription tier (e.g. &#x60;FREE&#x60;, &#x60;BRONZE&#x60;, &#x60;SILVER&#x60;, &#x60;GOLD&#x60;, &#x60;PLATINUM&#x60;, &#x60;DIAMOND&#x60;). The actual price applied to a run is resolved from the user&#39;s tier.  | [optional] 
**TrialMinutes** | **int32** | For how long this Actor can be used for free in trial period | 

## Methods

### NewActorRunPricingInfo

`func NewActorRunPricingInfo(apifyMarginPercentage float32, createdAt time.Time, startedAt time.Time, pricingModel string, pricingPerEvent PayPerEventActorPricingInfoAllOfPricingPerEvent, unitName string, pricePerUnitUsd float32, trialMinutes int32, ) *ActorRunPricingInfo`

NewActorRunPricingInfo instantiates a new ActorRunPricingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorRunPricingInfoWithDefaults

`func NewActorRunPricingInfoWithDefaults() *ActorRunPricingInfo`

NewActorRunPricingInfoWithDefaults instantiates a new ActorRunPricingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApifyMarginPercentage

`func (o *ActorRunPricingInfo) GetApifyMarginPercentage() float32`

GetApifyMarginPercentage returns the ApifyMarginPercentage field if non-nil, zero value otherwise.

### GetApifyMarginPercentageOk

`func (o *ActorRunPricingInfo) GetApifyMarginPercentageOk() (*float32, bool)`

GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApifyMarginPercentage

`func (o *ActorRunPricingInfo) SetApifyMarginPercentage(v float32)`

SetApifyMarginPercentage sets ApifyMarginPercentage field to given value.


### GetCreatedAt

`func (o *ActorRunPricingInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActorRunPricingInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActorRunPricingInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *ActorRunPricingInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ActorRunPricingInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ActorRunPricingInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetNotifiedAboutFutureChangeAt

`func (o *ActorRunPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time`

GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutFutureChangeAtOk

`func (o *ActorRunPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutFutureChangeAt

`func (o *ActorRunPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time)`

SetNotifiedAboutFutureChangeAt sets NotifiedAboutFutureChangeAt field to given value.

### HasNotifiedAboutFutureChangeAt

`func (o *ActorRunPricingInfo) HasNotifiedAboutFutureChangeAt() bool`

HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.

### GetNotifiedAboutChangeAt

`func (o *ActorRunPricingInfo) GetNotifiedAboutChangeAt() time.Time`

GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutChangeAtOk

`func (o *ActorRunPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutChangeAt

`func (o *ActorRunPricingInfo) SetNotifiedAboutChangeAt(v time.Time)`

SetNotifiedAboutChangeAt sets NotifiedAboutChangeAt field to given value.

### HasNotifiedAboutChangeAt

`func (o *ActorRunPricingInfo) HasNotifiedAboutChangeAt() bool`

HasNotifiedAboutChangeAt returns a boolean if a field has been set.

### GetReasonForChange

`func (o *ActorRunPricingInfo) GetReasonForChange() string`

GetReasonForChange returns the ReasonForChange field if non-nil, zero value otherwise.

### GetReasonForChangeOk

`func (o *ActorRunPricingInfo) GetReasonForChangeOk() (*string, bool)`

GetReasonForChangeOk returns a tuple with the ReasonForChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForChange

`func (o *ActorRunPricingInfo) SetReasonForChange(v string)`

SetReasonForChange sets ReasonForChange field to given value.

### HasReasonForChange

`func (o *ActorRunPricingInfo) HasReasonForChange() bool`

HasReasonForChange returns a boolean if a field has been set.

### GetIsPriceChangeNotificationSuppressed

`func (o *ActorRunPricingInfo) GetIsPriceChangeNotificationSuppressed() bool`

GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field if non-nil, zero value otherwise.

### GetIsPriceChangeNotificationSuppressedOk

`func (o *ActorRunPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool)`

GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceChangeNotificationSuppressed

`func (o *ActorRunPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool)`

SetIsPriceChangeNotificationSuppressed sets IsPriceChangeNotificationSuppressed field to given value.

### HasIsPriceChangeNotificationSuppressed

`func (o *ActorRunPricingInfo) HasIsPriceChangeNotificationSuppressed() bool`

HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.

### GetForceContainsSignificantPriceChange

`func (o *ActorRunPricingInfo) GetForceContainsSignificantPriceChange() bool`

GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field if non-nil, zero value otherwise.

### GetForceContainsSignificantPriceChangeOk

`func (o *ActorRunPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool)`

GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceContainsSignificantPriceChange

`func (o *ActorRunPricingInfo) SetForceContainsSignificantPriceChange(v bool)`

SetForceContainsSignificantPriceChange sets ForceContainsSignificantPriceChange field to given value.

### HasForceContainsSignificantPriceChange

`func (o *ActorRunPricingInfo) HasForceContainsSignificantPriceChange() bool`

HasForceContainsSignificantPriceChange returns a boolean if a field has been set.

### GetPricingModel

`func (o *ActorRunPricingInfo) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *ActorRunPricingInfo) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *ActorRunPricingInfo) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetPricingPerEvent

`func (o *ActorRunPricingInfo) GetPricingPerEvent() PayPerEventActorPricingInfoAllOfPricingPerEvent`

GetPricingPerEvent returns the PricingPerEvent field if non-nil, zero value otherwise.

### GetPricingPerEventOk

`func (o *ActorRunPricingInfo) GetPricingPerEventOk() (*PayPerEventActorPricingInfoAllOfPricingPerEvent, bool)`

GetPricingPerEventOk returns a tuple with the PricingPerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPerEvent

`func (o *ActorRunPricingInfo) SetPricingPerEvent(v PayPerEventActorPricingInfoAllOfPricingPerEvent)`

SetPricingPerEvent sets PricingPerEvent field to given value.


### GetMinimalMaxTotalChargeUsd

`func (o *ActorRunPricingInfo) GetMinimalMaxTotalChargeUsd() float32`

GetMinimalMaxTotalChargeUsd returns the MinimalMaxTotalChargeUsd field if non-nil, zero value otherwise.

### GetMinimalMaxTotalChargeUsdOk

`func (o *ActorRunPricingInfo) GetMinimalMaxTotalChargeUsdOk() (*float32, bool)`

GetMinimalMaxTotalChargeUsdOk returns a tuple with the MinimalMaxTotalChargeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalMaxTotalChargeUsd

`func (o *ActorRunPricingInfo) SetMinimalMaxTotalChargeUsd(v float32)`

SetMinimalMaxTotalChargeUsd sets MinimalMaxTotalChargeUsd field to given value.

### HasMinimalMaxTotalChargeUsd

`func (o *ActorRunPricingInfo) HasMinimalMaxTotalChargeUsd() bool`

HasMinimalMaxTotalChargeUsd returns a boolean if a field has been set.

### GetUnitName

`func (o *ActorRunPricingInfo) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *ActorRunPricingInfo) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *ActorRunPricingInfo) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.


### GetPricePerUnitUsd

`func (o *ActorRunPricingInfo) GetPricePerUnitUsd() float32`

GetPricePerUnitUsd returns the PricePerUnitUsd field if non-nil, zero value otherwise.

### GetPricePerUnitUsdOk

`func (o *ActorRunPricingInfo) GetPricePerUnitUsdOk() (*float32, bool)`

GetPricePerUnitUsdOk returns a tuple with the PricePerUnitUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnitUsd

`func (o *ActorRunPricingInfo) SetPricePerUnitUsd(v float32)`

SetPricePerUnitUsd sets PricePerUnitUsd field to given value.


### GetTieredPricing

`func (o *ActorRunPricingInfo) GetTieredPricing() map[string]TieredPricingPerDatasetItemEntry`

GetTieredPricing returns the TieredPricing field if non-nil, zero value otherwise.

### GetTieredPricingOk

`func (o *ActorRunPricingInfo) GetTieredPricingOk() (*map[string]TieredPricingPerDatasetItemEntry, bool)`

GetTieredPricingOk returns a tuple with the TieredPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredPricing

`func (o *ActorRunPricingInfo) SetTieredPricing(v map[string]TieredPricingPerDatasetItemEntry)`

SetTieredPricing sets TieredPricing field to given value.

### HasTieredPricing

`func (o *ActorRunPricingInfo) HasTieredPricing() bool`

HasTieredPricing returns a boolean if a field has been set.

### GetTrialMinutes

`func (o *ActorRunPricingInfo) GetTrialMinutes() int32`

GetTrialMinutes returns the TrialMinutes field if non-nil, zero value otherwise.

### GetTrialMinutesOk

`func (o *ActorRunPricingInfo) GetTrialMinutesOk() (*int32, bool)`

GetTrialMinutesOk returns a tuple with the TrialMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialMinutes

`func (o *ActorRunPricingInfo) SetTrialMinutes(v int32)`

SetTrialMinutes sets TrialMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


