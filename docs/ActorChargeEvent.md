# ActorChargeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTitle** | **string** | Human-readable title shown to users in the billing UI. | 
**EventDescription** | **string** | Human-readable description of what triggers this event. | 
**EventPriceUsd** | Pointer to **float32** | Flat price per event in USD. Present only for non-tiered events. Mutually exclusive with &#x60;eventTieredPricingUsd&#x60;.  | [optional] 
**EventTieredPricingUsd** | Pointer to [**map[string]TieredPricingPerEventEntry**](TieredPricingPerEventEntry.md) | Tiered price-per-event pricing for a single charge event, keyed by subscription tier (e.g. &#x60;FREE&#x60;, &#x60;BRONZE&#x60;, &#x60;SILVER&#x60;, &#x60;GOLD&#x60;, &#x60;PLATINUM&#x60;, &#x60;DIAMOND&#x60;). The actual price applied is resolved from the user&#39;s tier.  | [optional] 
**IsPrimaryEvent** | Pointer to **bool** | Whether this event is the Actor&#39;s primary chargeable event. | [optional] 
**IsOneTimeEvent** | Pointer to **bool** | Whether this event can only be charged once per Actor run. | [optional] 

## Methods

### NewActorChargeEvent

`func NewActorChargeEvent(eventTitle string, eventDescription string, ) *ActorChargeEvent`

NewActorChargeEvent instantiates a new ActorChargeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorChargeEventWithDefaults

`func NewActorChargeEventWithDefaults() *ActorChargeEvent`

NewActorChargeEventWithDefaults instantiates a new ActorChargeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTitle

`func (o *ActorChargeEvent) GetEventTitle() string`

GetEventTitle returns the EventTitle field if non-nil, zero value otherwise.

### GetEventTitleOk

`func (o *ActorChargeEvent) GetEventTitleOk() (*string, bool)`

GetEventTitleOk returns a tuple with the EventTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTitle

`func (o *ActorChargeEvent) SetEventTitle(v string)`

SetEventTitle sets EventTitle field to given value.


### GetEventDescription

`func (o *ActorChargeEvent) GetEventDescription() string`

GetEventDescription returns the EventDescription field if non-nil, zero value otherwise.

### GetEventDescriptionOk

`func (o *ActorChargeEvent) GetEventDescriptionOk() (*string, bool)`

GetEventDescriptionOk returns a tuple with the EventDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDescription

`func (o *ActorChargeEvent) SetEventDescription(v string)`

SetEventDescription sets EventDescription field to given value.


### GetEventPriceUsd

`func (o *ActorChargeEvent) GetEventPriceUsd() float32`

GetEventPriceUsd returns the EventPriceUsd field if non-nil, zero value otherwise.

### GetEventPriceUsdOk

`func (o *ActorChargeEvent) GetEventPriceUsdOk() (*float32, bool)`

GetEventPriceUsdOk returns a tuple with the EventPriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPriceUsd

`func (o *ActorChargeEvent) SetEventPriceUsd(v float32)`

SetEventPriceUsd sets EventPriceUsd field to given value.

### HasEventPriceUsd

`func (o *ActorChargeEvent) HasEventPriceUsd() bool`

HasEventPriceUsd returns a boolean if a field has been set.

### GetEventTieredPricingUsd

`func (o *ActorChargeEvent) GetEventTieredPricingUsd() map[string]TieredPricingPerEventEntry`

GetEventTieredPricingUsd returns the EventTieredPricingUsd field if non-nil, zero value otherwise.

### GetEventTieredPricingUsdOk

`func (o *ActorChargeEvent) GetEventTieredPricingUsdOk() (*map[string]TieredPricingPerEventEntry, bool)`

GetEventTieredPricingUsdOk returns a tuple with the EventTieredPricingUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTieredPricingUsd

`func (o *ActorChargeEvent) SetEventTieredPricingUsd(v map[string]TieredPricingPerEventEntry)`

SetEventTieredPricingUsd sets EventTieredPricingUsd field to given value.

### HasEventTieredPricingUsd

`func (o *ActorChargeEvent) HasEventTieredPricingUsd() bool`

HasEventTieredPricingUsd returns a boolean if a field has been set.

### GetIsPrimaryEvent

`func (o *ActorChargeEvent) GetIsPrimaryEvent() bool`

GetIsPrimaryEvent returns the IsPrimaryEvent field if non-nil, zero value otherwise.

### GetIsPrimaryEventOk

`func (o *ActorChargeEvent) GetIsPrimaryEventOk() (*bool, bool)`

GetIsPrimaryEventOk returns a tuple with the IsPrimaryEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryEvent

`func (o *ActorChargeEvent) SetIsPrimaryEvent(v bool)`

SetIsPrimaryEvent sets IsPrimaryEvent field to given value.

### HasIsPrimaryEvent

`func (o *ActorChargeEvent) HasIsPrimaryEvent() bool`

HasIsPrimaryEvent returns a boolean if a field has been set.

### GetIsOneTimeEvent

`func (o *ActorChargeEvent) GetIsOneTimeEvent() bool`

GetIsOneTimeEvent returns the IsOneTimeEvent field if non-nil, zero value otherwise.

### GetIsOneTimeEventOk

`func (o *ActorChargeEvent) GetIsOneTimeEventOk() (*bool, bool)`

GetIsOneTimeEventOk returns a tuple with the IsOneTimeEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOneTimeEvent

`func (o *ActorChargeEvent) SetIsOneTimeEvent(v bool)`

SetIsOneTimeEvent sets IsOneTimeEvent field to given value.

### HasIsOneTimeEvent

`func (o *ActorChargeEvent) HasIsOneTimeEvent() bool`

HasIsOneTimeEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


