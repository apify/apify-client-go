# FreeActorPricingInfo

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

## Methods

### NewFreeActorPricingInfo

`func NewFreeActorPricingInfo(apifyMarginPercentage float32, createdAt time.Time, startedAt time.Time, pricingModel string, ) *FreeActorPricingInfo`

NewFreeActorPricingInfo instantiates a new FreeActorPricingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeActorPricingInfoWithDefaults

`func NewFreeActorPricingInfoWithDefaults() *FreeActorPricingInfo`

NewFreeActorPricingInfoWithDefaults instantiates a new FreeActorPricingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApifyMarginPercentage

`func (o *FreeActorPricingInfo) GetApifyMarginPercentage() float32`

GetApifyMarginPercentage returns the ApifyMarginPercentage field if non-nil, zero value otherwise.

### GetApifyMarginPercentageOk

`func (o *FreeActorPricingInfo) GetApifyMarginPercentageOk() (*float32, bool)`

GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApifyMarginPercentage

`func (o *FreeActorPricingInfo) SetApifyMarginPercentage(v float32)`

SetApifyMarginPercentage sets ApifyMarginPercentage field to given value.


### GetCreatedAt

`func (o *FreeActorPricingInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FreeActorPricingInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FreeActorPricingInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *FreeActorPricingInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FreeActorPricingInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FreeActorPricingInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetNotifiedAboutFutureChangeAt

`func (o *FreeActorPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time`

GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutFutureChangeAtOk

`func (o *FreeActorPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutFutureChangeAt

`func (o *FreeActorPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time)`

SetNotifiedAboutFutureChangeAt sets NotifiedAboutFutureChangeAt field to given value.

### HasNotifiedAboutFutureChangeAt

`func (o *FreeActorPricingInfo) HasNotifiedAboutFutureChangeAt() bool`

HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.

### GetNotifiedAboutChangeAt

`func (o *FreeActorPricingInfo) GetNotifiedAboutChangeAt() time.Time`

GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutChangeAtOk

`func (o *FreeActorPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutChangeAt

`func (o *FreeActorPricingInfo) SetNotifiedAboutChangeAt(v time.Time)`

SetNotifiedAboutChangeAt sets NotifiedAboutChangeAt field to given value.

### HasNotifiedAboutChangeAt

`func (o *FreeActorPricingInfo) HasNotifiedAboutChangeAt() bool`

HasNotifiedAboutChangeAt returns a boolean if a field has been set.

### GetReasonForChange

`func (o *FreeActorPricingInfo) GetReasonForChange() string`

GetReasonForChange returns the ReasonForChange field if non-nil, zero value otherwise.

### GetReasonForChangeOk

`func (o *FreeActorPricingInfo) GetReasonForChangeOk() (*string, bool)`

GetReasonForChangeOk returns a tuple with the ReasonForChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForChange

`func (o *FreeActorPricingInfo) SetReasonForChange(v string)`

SetReasonForChange sets ReasonForChange field to given value.

### HasReasonForChange

`func (o *FreeActorPricingInfo) HasReasonForChange() bool`

HasReasonForChange returns a boolean if a field has been set.

### GetIsPriceChangeNotificationSuppressed

`func (o *FreeActorPricingInfo) GetIsPriceChangeNotificationSuppressed() bool`

GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field if non-nil, zero value otherwise.

### GetIsPriceChangeNotificationSuppressedOk

`func (o *FreeActorPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool)`

GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceChangeNotificationSuppressed

`func (o *FreeActorPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool)`

SetIsPriceChangeNotificationSuppressed sets IsPriceChangeNotificationSuppressed field to given value.

### HasIsPriceChangeNotificationSuppressed

`func (o *FreeActorPricingInfo) HasIsPriceChangeNotificationSuppressed() bool`

HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.

### GetForceContainsSignificantPriceChange

`func (o *FreeActorPricingInfo) GetForceContainsSignificantPriceChange() bool`

GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field if non-nil, zero value otherwise.

### GetForceContainsSignificantPriceChangeOk

`func (o *FreeActorPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool)`

GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceContainsSignificantPriceChange

`func (o *FreeActorPricingInfo) SetForceContainsSignificantPriceChange(v bool)`

SetForceContainsSignificantPriceChange sets ForceContainsSignificantPriceChange field to given value.

### HasForceContainsSignificantPriceChange

`func (o *FreeActorPricingInfo) HasForceContainsSignificantPriceChange() bool`

HasForceContainsSignificantPriceChange returns a boolean if a field has been set.

### GetPricingModel

`func (o *FreeActorPricingInfo) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *FreeActorPricingInfo) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *FreeActorPricingInfo) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


