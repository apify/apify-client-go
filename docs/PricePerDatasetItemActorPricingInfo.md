# PricePerDatasetItemActorPricingInfo

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
**UnitName** | **string** | Name of the unit that is being charged | 
**PricePerUnitUsd** | Pointer to **float32** | Price per unit in USD. Mutually exclusive with &#x60;tieredPricing&#x60; - exactly one of the two is present on a pricing record.  | [optional] 
**TieredPricing** | Pointer to [**map[string]TieredPricingPerDatasetItemEntry**](TieredPricingPerDatasetItemEntry.md) | Tiered price-per-dataset-item pricing, keyed by subscription tier (e.g. &#x60;FREE&#x60;, &#x60;BRONZE&#x60;, &#x60;SILVER&#x60;, &#x60;GOLD&#x60;, &#x60;PLATINUM&#x60;, &#x60;DIAMOND&#x60;). The actual price applied to a run is resolved from the user&#39;s tier.  | [optional] 

## Methods

### NewPricePerDatasetItemActorPricingInfo

`func NewPricePerDatasetItemActorPricingInfo(apifyMarginPercentage float32, createdAt time.Time, startedAt time.Time, pricingModel string, unitName string, ) *PricePerDatasetItemActorPricingInfo`

NewPricePerDatasetItemActorPricingInfo instantiates a new PricePerDatasetItemActorPricingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricePerDatasetItemActorPricingInfoWithDefaults

`func NewPricePerDatasetItemActorPricingInfoWithDefaults() *PricePerDatasetItemActorPricingInfo`

NewPricePerDatasetItemActorPricingInfoWithDefaults instantiates a new PricePerDatasetItemActorPricingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApifyMarginPercentage

`func (o *PricePerDatasetItemActorPricingInfo) GetApifyMarginPercentage() float32`

GetApifyMarginPercentage returns the ApifyMarginPercentage field if non-nil, zero value otherwise.

### GetApifyMarginPercentageOk

`func (o *PricePerDatasetItemActorPricingInfo) GetApifyMarginPercentageOk() (*float32, bool)`

GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApifyMarginPercentage

`func (o *PricePerDatasetItemActorPricingInfo) SetApifyMarginPercentage(v float32)`

SetApifyMarginPercentage sets ApifyMarginPercentage field to given value.


### GetCreatedAt

`func (o *PricePerDatasetItemActorPricingInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PricePerDatasetItemActorPricingInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PricePerDatasetItemActorPricingInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *PricePerDatasetItemActorPricingInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PricePerDatasetItemActorPricingInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PricePerDatasetItemActorPricingInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetNotifiedAboutFutureChangeAt

`func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time`

GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutFutureChangeAtOk

`func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutFutureChangeAt

`func (o *PricePerDatasetItemActorPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time)`

SetNotifiedAboutFutureChangeAt sets NotifiedAboutFutureChangeAt field to given value.

### HasNotifiedAboutFutureChangeAt

`func (o *PricePerDatasetItemActorPricingInfo) HasNotifiedAboutFutureChangeAt() bool`

HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.

### GetNotifiedAboutChangeAt

`func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutChangeAt() time.Time`

GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutChangeAtOk

`func (o *PricePerDatasetItemActorPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutChangeAt

`func (o *PricePerDatasetItemActorPricingInfo) SetNotifiedAboutChangeAt(v time.Time)`

SetNotifiedAboutChangeAt sets NotifiedAboutChangeAt field to given value.

### HasNotifiedAboutChangeAt

`func (o *PricePerDatasetItemActorPricingInfo) HasNotifiedAboutChangeAt() bool`

HasNotifiedAboutChangeAt returns a boolean if a field has been set.

### GetReasonForChange

`func (o *PricePerDatasetItemActorPricingInfo) GetReasonForChange() string`

GetReasonForChange returns the ReasonForChange field if non-nil, zero value otherwise.

### GetReasonForChangeOk

`func (o *PricePerDatasetItemActorPricingInfo) GetReasonForChangeOk() (*string, bool)`

GetReasonForChangeOk returns a tuple with the ReasonForChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForChange

`func (o *PricePerDatasetItemActorPricingInfo) SetReasonForChange(v string)`

SetReasonForChange sets ReasonForChange field to given value.

### HasReasonForChange

`func (o *PricePerDatasetItemActorPricingInfo) HasReasonForChange() bool`

HasReasonForChange returns a boolean if a field has been set.

### GetIsPriceChangeNotificationSuppressed

`func (o *PricePerDatasetItemActorPricingInfo) GetIsPriceChangeNotificationSuppressed() bool`

GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field if non-nil, zero value otherwise.

### GetIsPriceChangeNotificationSuppressedOk

`func (o *PricePerDatasetItemActorPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool)`

GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceChangeNotificationSuppressed

`func (o *PricePerDatasetItemActorPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool)`

SetIsPriceChangeNotificationSuppressed sets IsPriceChangeNotificationSuppressed field to given value.

### HasIsPriceChangeNotificationSuppressed

`func (o *PricePerDatasetItemActorPricingInfo) HasIsPriceChangeNotificationSuppressed() bool`

HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.

### GetForceContainsSignificantPriceChange

`func (o *PricePerDatasetItemActorPricingInfo) GetForceContainsSignificantPriceChange() bool`

GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field if non-nil, zero value otherwise.

### GetForceContainsSignificantPriceChangeOk

`func (o *PricePerDatasetItemActorPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool)`

GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceContainsSignificantPriceChange

`func (o *PricePerDatasetItemActorPricingInfo) SetForceContainsSignificantPriceChange(v bool)`

SetForceContainsSignificantPriceChange sets ForceContainsSignificantPriceChange field to given value.

### HasForceContainsSignificantPriceChange

`func (o *PricePerDatasetItemActorPricingInfo) HasForceContainsSignificantPriceChange() bool`

HasForceContainsSignificantPriceChange returns a boolean if a field has been set.

### GetPricingModel

`func (o *PricePerDatasetItemActorPricingInfo) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *PricePerDatasetItemActorPricingInfo) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *PricePerDatasetItemActorPricingInfo) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetUnitName

`func (o *PricePerDatasetItemActorPricingInfo) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *PricePerDatasetItemActorPricingInfo) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *PricePerDatasetItemActorPricingInfo) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.


### GetPricePerUnitUsd

`func (o *PricePerDatasetItemActorPricingInfo) GetPricePerUnitUsd() float32`

GetPricePerUnitUsd returns the PricePerUnitUsd field if non-nil, zero value otherwise.

### GetPricePerUnitUsdOk

`func (o *PricePerDatasetItemActorPricingInfo) GetPricePerUnitUsdOk() (*float32, bool)`

GetPricePerUnitUsdOk returns a tuple with the PricePerUnitUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnitUsd

`func (o *PricePerDatasetItemActorPricingInfo) SetPricePerUnitUsd(v float32)`

SetPricePerUnitUsd sets PricePerUnitUsd field to given value.

### HasPricePerUnitUsd

`func (o *PricePerDatasetItemActorPricingInfo) HasPricePerUnitUsd() bool`

HasPricePerUnitUsd returns a boolean if a field has been set.

### GetTieredPricing

`func (o *PricePerDatasetItemActorPricingInfo) GetTieredPricing() map[string]TieredPricingPerDatasetItemEntry`

GetTieredPricing returns the TieredPricing field if non-nil, zero value otherwise.

### GetTieredPricingOk

`func (o *PricePerDatasetItemActorPricingInfo) GetTieredPricingOk() (*map[string]TieredPricingPerDatasetItemEntry, bool)`

GetTieredPricingOk returns a tuple with the TieredPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredPricing

`func (o *PricePerDatasetItemActorPricingInfo) SetTieredPricing(v map[string]TieredPricingPerDatasetItemEntry)`

SetTieredPricing sets TieredPricing field to given value.

### HasTieredPricing

`func (o *PricePerDatasetItemActorPricingInfo) HasTieredPricing() bool`

HasTieredPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


