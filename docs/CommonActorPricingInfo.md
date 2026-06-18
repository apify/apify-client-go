# CommonActorPricingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApifyMarginPercentage** | **float32** | In [0, 1], fraction of pricePerUnitUsd that goes to Apify | 
**CreatedAt** | **time.Time** | When this pricing info record has been created | 
**StartedAt** | **time.Time** | Since when is this pricing info record effective for a given Actor | 
**NotifiedAboutFutureChangeAt** | Pointer to **NullableTime** |  | [optional] 
**NotifiedAboutChangeAt** | Pointer to **NullableTime** |  | [optional] 
**ReasonForChange** | Pointer to **NullableString** |  | [optional] 
**IsPriceChangeNotificationSuppressed** | Pointer to **bool** |  | [optional] 
**ForceContainsSignificantPriceChange** | Pointer to **bool** |  | [optional] 

## Methods

### NewCommonActorPricingInfo

`func NewCommonActorPricingInfo(apifyMarginPercentage float32, createdAt time.Time, startedAt time.Time, ) *CommonActorPricingInfo`

NewCommonActorPricingInfo instantiates a new CommonActorPricingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonActorPricingInfoWithDefaults

`func NewCommonActorPricingInfoWithDefaults() *CommonActorPricingInfo`

NewCommonActorPricingInfoWithDefaults instantiates a new CommonActorPricingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApifyMarginPercentage

`func (o *CommonActorPricingInfo) GetApifyMarginPercentage() float32`

GetApifyMarginPercentage returns the ApifyMarginPercentage field if non-nil, zero value otherwise.

### GetApifyMarginPercentageOk

`func (o *CommonActorPricingInfo) GetApifyMarginPercentageOk() (*float32, bool)`

GetApifyMarginPercentageOk returns a tuple with the ApifyMarginPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApifyMarginPercentage

`func (o *CommonActorPricingInfo) SetApifyMarginPercentage(v float32)`

SetApifyMarginPercentage sets ApifyMarginPercentage field to given value.


### GetCreatedAt

`func (o *CommonActorPricingInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommonActorPricingInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommonActorPricingInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *CommonActorPricingInfo) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CommonActorPricingInfo) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CommonActorPricingInfo) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetNotifiedAboutFutureChangeAt

`func (o *CommonActorPricingInfo) GetNotifiedAboutFutureChangeAt() time.Time`

GetNotifiedAboutFutureChangeAt returns the NotifiedAboutFutureChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutFutureChangeAtOk

`func (o *CommonActorPricingInfo) GetNotifiedAboutFutureChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutFutureChangeAtOk returns a tuple with the NotifiedAboutFutureChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutFutureChangeAt

`func (o *CommonActorPricingInfo) SetNotifiedAboutFutureChangeAt(v time.Time)`

SetNotifiedAboutFutureChangeAt sets NotifiedAboutFutureChangeAt field to given value.

### HasNotifiedAboutFutureChangeAt

`func (o *CommonActorPricingInfo) HasNotifiedAboutFutureChangeAt() bool`

HasNotifiedAboutFutureChangeAt returns a boolean if a field has been set.

### SetNotifiedAboutFutureChangeAtNil

`func (o *CommonActorPricingInfo) SetNotifiedAboutFutureChangeAtNil(b bool)`

 SetNotifiedAboutFutureChangeAtNil sets the value for NotifiedAboutFutureChangeAt to be an explicit nil

### UnsetNotifiedAboutFutureChangeAt
`func (o *CommonActorPricingInfo) UnsetNotifiedAboutFutureChangeAt()`

UnsetNotifiedAboutFutureChangeAt ensures that no value is present for NotifiedAboutFutureChangeAt, not even an explicit nil
### GetNotifiedAboutChangeAt

`func (o *CommonActorPricingInfo) GetNotifiedAboutChangeAt() time.Time`

GetNotifiedAboutChangeAt returns the NotifiedAboutChangeAt field if non-nil, zero value otherwise.

### GetNotifiedAboutChangeAtOk

`func (o *CommonActorPricingInfo) GetNotifiedAboutChangeAtOk() (*time.Time, bool)`

GetNotifiedAboutChangeAtOk returns a tuple with the NotifiedAboutChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAboutChangeAt

`func (o *CommonActorPricingInfo) SetNotifiedAboutChangeAt(v time.Time)`

SetNotifiedAboutChangeAt sets NotifiedAboutChangeAt field to given value.

### HasNotifiedAboutChangeAt

`func (o *CommonActorPricingInfo) HasNotifiedAboutChangeAt() bool`

HasNotifiedAboutChangeAt returns a boolean if a field has been set.

### SetNotifiedAboutChangeAtNil

`func (o *CommonActorPricingInfo) SetNotifiedAboutChangeAtNil(b bool)`

 SetNotifiedAboutChangeAtNil sets the value for NotifiedAboutChangeAt to be an explicit nil

### UnsetNotifiedAboutChangeAt
`func (o *CommonActorPricingInfo) UnsetNotifiedAboutChangeAt()`

UnsetNotifiedAboutChangeAt ensures that no value is present for NotifiedAboutChangeAt, not even an explicit nil
### GetReasonForChange

`func (o *CommonActorPricingInfo) GetReasonForChange() string`

GetReasonForChange returns the ReasonForChange field if non-nil, zero value otherwise.

### GetReasonForChangeOk

`func (o *CommonActorPricingInfo) GetReasonForChangeOk() (*string, bool)`

GetReasonForChangeOk returns a tuple with the ReasonForChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForChange

`func (o *CommonActorPricingInfo) SetReasonForChange(v string)`

SetReasonForChange sets ReasonForChange field to given value.

### HasReasonForChange

`func (o *CommonActorPricingInfo) HasReasonForChange() bool`

HasReasonForChange returns a boolean if a field has been set.

### SetReasonForChangeNil

`func (o *CommonActorPricingInfo) SetReasonForChangeNil(b bool)`

 SetReasonForChangeNil sets the value for ReasonForChange to be an explicit nil

### UnsetReasonForChange
`func (o *CommonActorPricingInfo) UnsetReasonForChange()`

UnsetReasonForChange ensures that no value is present for ReasonForChange, not even an explicit nil
### GetIsPriceChangeNotificationSuppressed

`func (o *CommonActorPricingInfo) GetIsPriceChangeNotificationSuppressed() bool`

GetIsPriceChangeNotificationSuppressed returns the IsPriceChangeNotificationSuppressed field if non-nil, zero value otherwise.

### GetIsPriceChangeNotificationSuppressedOk

`func (o *CommonActorPricingInfo) GetIsPriceChangeNotificationSuppressedOk() (*bool, bool)`

GetIsPriceChangeNotificationSuppressedOk returns a tuple with the IsPriceChangeNotificationSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceChangeNotificationSuppressed

`func (o *CommonActorPricingInfo) SetIsPriceChangeNotificationSuppressed(v bool)`

SetIsPriceChangeNotificationSuppressed sets IsPriceChangeNotificationSuppressed field to given value.

### HasIsPriceChangeNotificationSuppressed

`func (o *CommonActorPricingInfo) HasIsPriceChangeNotificationSuppressed() bool`

HasIsPriceChangeNotificationSuppressed returns a boolean if a field has been set.

### GetForceContainsSignificantPriceChange

`func (o *CommonActorPricingInfo) GetForceContainsSignificantPriceChange() bool`

GetForceContainsSignificantPriceChange returns the ForceContainsSignificantPriceChange field if non-nil, zero value otherwise.

### GetForceContainsSignificantPriceChangeOk

`func (o *CommonActorPricingInfo) GetForceContainsSignificantPriceChangeOk() (*bool, bool)`

GetForceContainsSignificantPriceChangeOk returns a tuple with the ForceContainsSignificantPriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceContainsSignificantPriceChange

`func (o *CommonActorPricingInfo) SetForceContainsSignificantPriceChange(v bool)`

SetForceContainsSignificantPriceChange sets ForceContainsSignificantPriceChange field to given value.

### HasForceContainsSignificantPriceChange

`func (o *CommonActorPricingInfo) HasForceContainsSignificantPriceChange() bool`

HasForceContainsSignificantPriceChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


