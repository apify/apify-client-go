# ChargeRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** |  | 
**Count** | **int32** |  | 

## Methods

### NewChargeRunRequest

`func NewChargeRunRequest(eventName string, count int32, ) *ChargeRunRequest`

NewChargeRunRequest instantiates a new ChargeRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeRunRequestWithDefaults

`func NewChargeRunRequestWithDefaults() *ChargeRunRequest`

NewChargeRunRequestWithDefaults instantiates a new ChargeRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *ChargeRunRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ChargeRunRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ChargeRunRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetCount

`func (o *ChargeRunRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ChargeRunRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ChargeRunRequest) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


