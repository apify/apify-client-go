# PayPerEventActorPricingInfo

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
**MinimalMaxTotalChargeUsd** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewPayPerEventActorPricingInfo

`func NewPayPerEventActorPricingInfo(apifyMarginPercentage float32, createdAt time.Time, startedAt time.Time, pricingModel string, pricingPerEvent PayPerEventActorPricingInfoAllOfPricingPerEvent, ) *PayPerEventActorPricingInfo`

NewPayPerEventActorPricingInfo instantiates a new PayPerEventActorPricingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayPerEventActorPricingInfoWithDefaults

`func NewPayPerEventActorPricingInfoWithDefaults() *PayPerEventActorPricingInfo`

NewPayPerEventActorPricingInfoWithDefaults instantiates a new PayPerEventActorPricingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApifyMarginPercentage

`func (o *PayPerEventActorPricingInfo) GetApifyMarginPercentage() float32`

GetApifyMarginPercentage returns the ApifyMarginPercentage field if non-nil, zero value otherwise.

### GetApifyMarginPercentageOk

`func (o *PayPerEventActorPricingInfo) GetApifyMarginPercentageOk() (*float32, bool)`

GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApifyMarginPercentage

`func (o *PayPerEventActorPricingInfo) SetApifyMarginPercentage(v float32)`

SetApifyMarginPercentage sets ApifyMarginPercentage field to given value.


### GetCreatedAt

`func (o *PayPerEventActorPricingInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PayPerEventActorPricingInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PayPerEventActorPricingInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *PayPerEventActorPricingInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PayPerEventActorPricingInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PayPerEventActorPricingInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetNotifiedAboutFutureChangeAt

`func (o *PayPerEventActorPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time`

GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutFutureChangeAtOk

`func (o *PayPerEventActorPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutFutureChangeAt

`func (o *PayPerEventActorPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time)`

SetNotifiedAboutFutureChangeAt sets NotifiedAboutFutureChangeAt field to given value.

### HasNotifiedAboutFutureChangeAt

`func (o *PayPerEventActorPricingInfo) HasNotifiedAboutFutureChangeAt() bool`

HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.

### GetNotifiedAboutChangeAt

`func (o *PayPerEventActorPricingInfo) GetNotifiedAboutChangeAt() time.Time`

GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutChangeAtOk

`func (o *PayPerEventActorPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutChangeAt

`func (o *PayPerEventActorPricingInfo) SetNotifiedAboutChangeAt(v time.Time)`

SetNotifiedAboutChangeAt sets NotifiedAboutChangeAt field to given value.

### HasNotifiedAboutChangeAt

`func (o *PayPerEventActorPricingInfo) HasNotifiedAboutChangeAt() bool`

HasNotifiedAboutChangeAt returns a boolean if a field has been set.

### GetReasonForChange

`func (o *PayPerEventActorPricingInfo) GetReasonForChange() string`

GetReasonForChange returns the ReasonForChange field if non-nil, zero value otherwise.

### GetReasonForChangeOk

`func (o *PayPerEventActorPricingInfo) GetReasonForChangeOk() (*string, bool)`

GetReasonForChangeOk returns a tuple with the ReasonForChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForChange

`func (o *PayPerEventActorPricingInfo) SetReasonForChange(v string)`

SetReasonForChange sets ReasonForChange field to given value.

### HasReasonForChange

`func (o *PayPerEventActorPricingInfo) HasReasonForChange() bool`

HasReasonForChange returns a boolean if a field has been set.

### GetIsPriceChangeNotificationSuppressed

`func (o *PayPerEventActorPricingInfo) GetIsPriceChangeNotificationSuppressed() bool`

GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field if non-nil, zero value otherwise.

### GetIsPriceChangeNotificationSuppressedOk

`func (o *PayPerEventActorPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool)`

GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceChangeNotificationSuppressed

`func (o *PayPerEventActorPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool)`

SetIsPriceChangeNotificationSuppressed sets IsPriceChangeNotificationSuppressed field to given value.

### HasIsPriceChangeNotificationSuppressed

`func (o *PayPerEventActorPricingInfo) HasIsPriceChangeNotificationSuppressed() bool`

HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.

### GetForceContainsSignificantPriceChange

`func (o *PayPerEventActorPricingInfo) GetForceContainsSignificantPriceChange() bool`

GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field if non-nil, zero value otherwise.

### GetForceContainsSignificantPriceChangeOk

`func (o *PayPerEventActorPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool)`

GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceContainsSignificantPriceChange

`func (o *PayPerEventActorPricingInfo) SetForceContainsSignificantPriceChange(v bool)`

SetForceContainsSignificantPriceChange sets ForceContainsSignificantPriceChange field to given value.

### HasForceContainsSignificantPriceChange

`func (o *PayPerEventActorPricingInfo) HasForceContainsSignificantPriceChange() bool`

HasForceContainsSignificantPriceChange returns a boolean if a field has been set.

### GetPricingModel

`func (o *PayPerEventActorPricingInfo) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *PayPerEventActorPricingInfo) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *PayPerEventActorPricingInfo) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetPricingPerEvent

`func (o *PayPerEventActorPricingInfo) GetPricingPerEvent() PayPerEventActorPricingInfoAllOfPricingPerEvent`

GetPricingPerEvent returns the PricingPerEvent field if non-nil, zero value otherwise.

### GetPricingPerEventOk

`func (o *PayPerEventActorPricingInfo) GetPricingPerEventOk() (*PayPerEventActorPricingInfoAllOfPricingPerEvent, bool)`

GetPricingPerEventOk returns a tuple with the PricingPerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPerEvent

`func (o *PayPerEventActorPricingInfo) SetPricingPerEvent(v PayPerEventActorPricingInfoAllOfPricingPerEvent)`

SetPricingPerEvent sets PricingPerEvent field to given value.


### GetMinimalMaxTotalChargeUsd

`func (o *PayPerEventActorPricingInfo) GetMinimalMaxTotalChargeUsd() float32`

GetMinimalMaxTotalChargeUsd returns the MinimalMaxTotalChargeUsd field if non-nil, zero value otherwise.

### GetMinimalMaxTotalChargeUsdOk

`func (o *PayPerEventActorPricingInfo) GetMinimalMaxTotalChargeUsdOk() (*float32, bool)`

GetMinimalMaxTotalChargeUsdOk returns a tuple with the MinimalMaxTotalChargeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalMaxTotalChargeUsd

`func (o *PayPerEventActorPricingInfo) SetMinimalMaxTotalChargeUsd(v float32)`

SetMinimalMaxTotalChargeUsd sets MinimalMaxTotalChargeUsd field to given value.

### HasMinimalMaxTotalChargeUsd

`func (o *PayPerEventActorPricingInfo) HasMinimalMaxTotalChargeUsd() bool`

HasMinimalMaxTotalChargeUsd returns a boolean if a field has been set.

### SetMinimalMaxTotalChargeUsdNil

`func (o *PayPerEventActorPricingInfo) SetMinimalMaxTotalChargeUsdNil(b bool)`

 SetMinimalMaxTotalChargeUsdNil sets the value for MinimalMaxTotalChargeUsd to be an explicit nil

### UnsetMinimalMaxTotalChargeUsd
`func (o *PayPerEventActorPricingInfo) UnsetMinimalMaxTotalChargeUsd()`

UnsetMinimalMaxTotalChargeUsd ensures that no value is present for MinimalMaxTotalChargeUsd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


