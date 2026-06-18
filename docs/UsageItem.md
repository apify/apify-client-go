# UsageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **float32** |  | 
**BaseAmountUsd** | **float32** |  | 
**BaseUnitPriceUsd** | Pointer to **float32** |  | [optional] 
**AmountAfterVolumeDiscountUsd** | Pointer to **float32** |  | [optional] 
**PriceTiers** | Pointer to [**[]PriceTiers**](PriceTiers.md) |  | [optional] 

## Methods

### NewUsageItem

`func NewUsageItem(quantity float32, baseAmountUsd float32, ) *UsageItem`

NewUsageItem instantiates a new UsageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageItemWithDefaults

`func NewUsageItemWithDefaults() *UsageItem`

NewUsageItemWithDefaults instantiates a new UsageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *UsageItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetBaseAmountUsd

`func (o *UsageItem) GetBaseAmountUsd() float32`

GetBaseAmountUsd returns the BaseAmountUsd field if non-nil, zero value otherwise.

### GetBaseAmountUsdOk

`func (o *UsageItem) GetBaseAmountUsdOk() (*float32, bool)`

GetBaseAmountUsdOk returns a tuple with the BaseAmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmountUsd

`func (o *UsageItem) SetBaseAmountUsd(v float32)`

SetBaseAmountUsd sets BaseAmountUsd field to given value.


### GetBaseUnitPriceUsd

`func (o *UsageItem) GetBaseUnitPriceUsd() float32`

GetBaseUnitPriceUsd returns the BaseUnitPriceUsd field if non-nil, zero value otherwise.

### GetBaseUnitPriceUsdOk

`func (o *UsageItem) GetBaseUnitPriceUsdOk() (*float32, bool)`

GetBaseUnitPriceUsdOk returns a tuple with the BaseUnitPriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnitPriceUsd

`func (o *UsageItem) SetBaseUnitPriceUsd(v float32)`

SetBaseUnitPriceUsd sets BaseUnitPriceUsd field to given value.

### HasBaseUnitPriceUsd

`func (o *UsageItem) HasBaseUnitPriceUsd() bool`

HasBaseUnitPriceUsd returns a boolean if a field has been set.

### GetAmountAfterVolumeDiscountUsd

`func (o *UsageItem) GetAmountAfterVolumeDiscountUsd() float32`

GetAmountAfterVolumeDiscountUsd returns the AmountAfterVolumeDiscountUsd field if non-nil, zero value otherwise.

### GetAmountAfterVolumeDiscountUsdOk

`func (o *UsageItem) GetAmountAfterVolumeDiscountUsdOk() (*float32, bool)`

GetAmountAfterVolumeDiscountUsdOk returns a tuple with the AmountAfterVolumeDiscountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAfterVolumeDiscountUsd

`func (o *UsageItem) SetAmountAfterVolumeDiscountUsd(v float32)`

SetAmountAfterVolumeDiscountUsd sets AmountAfterVolumeDiscountUsd field to given value.

### HasAmountAfterVolumeDiscountUsd

`func (o *UsageItem) HasAmountAfterVolumeDiscountUsd() bool`

HasAmountAfterVolumeDiscountUsd returns a boolean if a field has been set.

### GetPriceTiers

`func (o *UsageItem) GetPriceTiers() []PriceTiers`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *UsageItem) GetPriceTiersOk() (*[]PriceTiers, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *UsageItem) SetPriceTiers(v []PriceTiers)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *UsageItem) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


