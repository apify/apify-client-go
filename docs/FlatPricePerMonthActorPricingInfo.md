# FlatPricePerMonthActorPricingInfo

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
**TrialMinutes** | **int32** | For how long this Actor can be used for free in trial period | 
**PricePerUnitUsd** | **float32** | Monthly flat price in USD | 

## Methods

### NewFlatPricePerMonthActorPricingInfo

`func NewFlatPricePerMonthActorPricingInfo(apifyMarginPercentage float32, createdAt time.Time, startedAt time.Time, pricingModel string, trialMinutes int32, pricePerUnitUsd float32, ) *FlatPricePerMonthActorPricingInfo`

NewFlatPricePerMonthActorPricingInfo instantiates a new FlatPricePerMonthActorPricingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatPricePerMonthActorPricingInfoWithDefaults

`func NewFlatPricePerMonthActorPricingInfoWithDefaults() *FlatPricePerMonthActorPricingInfo`

NewFlatPricePerMonthActorPricingInfoWithDefaults instantiates a new FlatPricePerMonthActorPricingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApifyMarginPercentage

`func (o *FlatPricePerMonthActorPricingInfo) GetApifyMarginPercentage() float32`

GetApifyMarginPercentage returns the ApifyMarginPercentage field if non-nil, zero value otherwise.

### GetApifyMarginPercentageOk

`func (o *FlatPricePerMonthActorPricingInfo) GetApifyMarginPercentageOk() (*float32, bool)`

GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApifyMarginPercentage

`func (o *FlatPricePerMonthActorPricingInfo) SetApifyMarginPercentage(v float32)`

SetApifyMarginPercentage sets ApifyMarginPercentage field to given value.


### GetCreatedAt

`func (o *FlatPricePerMonthActorPricingInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlatPricePerMonthActorPricingInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlatPricePerMonthActorPricingInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *FlatPricePerMonthActorPricingInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FlatPricePerMonthActorPricingInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FlatPricePerMonthActorPricingInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetNotifiedAboutFutureChangeAt

`func (o *FlatPricePerMonthActorPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time`

GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutFutureChangeAtOk

`func (o *FlatPricePerMonthActorPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutFutureChangeAt

`func (o *FlatPricePerMonthActorPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time)`

SetNotifiedAboutFutureChangeAt sets NotifiedAboutFutureChangeAt field to given value.

### HasNotifiedAboutFutureChangeAt

`func (o *FlatPricePerMonthActorPricingInfo) HasNotifiedAboutFutureChangeAt() bool`

HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.

### GetNotifiedAboutChangeAt

`func (o *FlatPricePerMonthActorPricingInfo) GetNotifiedAboutChangeAt() time.Time`

GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutChangeAtOk

`func (o *FlatPricePerMonthActorPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutChangeAt

`func (o *FlatPricePerMonthActorPricingInfo) SetNotifiedAboutChangeAt(v time.Time)`

SetNotifiedAboutChangeAt sets NotifiedAboutChangeAt field to given value.

### HasNotifiedAboutChangeAt

`func (o *FlatPricePerMonthActorPricingInfo) HasNotifiedAboutChangeAt() bool`

HasNotifiedAboutChangeAt returns a boolean if a field has been set.

### GetReasonForChange

`func (o *FlatPricePerMonthActorPricingInfo) GetReasonForChange() string`

GetReasonForChange returns the ReasonForChange field if non-nil, zero value otherwise.

### GetReasonForChangeOk

`func (o *FlatPricePerMonthActorPricingInfo) GetReasonForChangeOk() (*string, bool)`

GetReasonForChangeOk returns a tuple with the ReasonForChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForChange

`func (o *FlatPricePerMonthActorPricingInfo) SetReasonForChange(v string)`

SetReasonForChange sets ReasonForChange field to given value.

### HasReasonForChange

`func (o *FlatPricePerMonthActorPricingInfo) HasReasonForChange() bool`

HasReasonForChange returns a boolean if a field has been set.

### GetIsPriceChangeNotificationSuppressed

`func (o *FlatPricePerMonthActorPricingInfo) GetIsPriceChangeNotificationSuppressed() bool`

GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field if non-nil, zero value otherwise.

### GetIsPriceChangeNotificationSuppressedOk

`func (o *FlatPricePerMonthActorPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool)`

GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceChangeNotificationSuppressed

`func (o *FlatPricePerMonthActorPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool)`

SetIsPriceChangeNotificationSuppressed sets IsPriceChangeNotificationSuppressed field to given value.

### HasIsPriceChangeNotificationSuppressed

`func (o *FlatPricePerMonthActorPricingInfo) HasIsPriceChangeNotificationSuppressed() bool`

HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.

### GetForceContainsSignificantPriceChange

`func (o *FlatPricePerMonthActorPricingInfo) GetForceContainsSignificantPriceChange() bool`

GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field if non-nil, zero value otherwise.

### GetForceContainsSignificantPriceChangeOk

`func (o *FlatPricePerMonthActorPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool)`

GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceContainsSignificantPriceChange

`func (o *FlatPricePerMonthActorPricingInfo) SetForceContainsSignificantPriceChange(v bool)`

SetForceContainsSignificantPriceChange sets ForceContainsSignificantPriceChange field to given value.

### HasForceContainsSignificantPriceChange

`func (o *FlatPricePerMonthActorPricingInfo) HasForceContainsSignificantPriceChange() bool`

HasForceContainsSignificantPriceChange returns a boolean if a field has been set.

### GetPricingModel

`func (o *FlatPricePerMonthActorPricingInfo) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *FlatPricePerMonthActorPricingInfo) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *FlatPricePerMonthActorPricingInfo) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetTrialMinutes

`func (o *FlatPricePerMonthActorPricingInfo) GetTrialMinutes() int32`

GetTrialMinutes returns the TrialMinutes field if non-nil, zero value otherwise.

### GetTrialMinutesOk

`func (o *FlatPricePerMonthActorPricingInfo) GetTrialMinutesOk() (*int32, bool)`

GetTrialMinutesOk returns a tuple with the TrialMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialMinutes

`func (o *FlatPricePerMonthActorPricingInfo) SetTrialMinutes(v int32)`

SetTrialMinutes sets TrialMinutes field to given value.


### GetPricePerUnitUsd

`func (o *FlatPricePerMonthActorPricingInfo) GetPricePerUnitUsd() float32`

GetPricePerUnitUsd returns the PricePerUnitUsd field if non-nil, zero value otherwise.

### GetPricePerUnitUsdOk

`func (o *FlatPricePerMonthActorPricingInfo) GetPricePerUnitUsdOk() (*float32, bool)`

GetPricePerUnitUsdOk returns a tuple with the PricePerUnitUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnitUsd

`func (o *FlatPricePerMonthActorPricingInfo) SetPricePerUnitUsd(v float32)`

SetPricePerUnitUsd sets PricePerUnitUsd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


