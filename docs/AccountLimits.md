# AccountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthlyUsageCycle** | [**UsageCycle**](UsageCycle.md) |  | 
**Limits** | [**Limits**](Limits.md) |  | 
**Current** | [**Current**](Current.md) |  | 

## Methods

### NewAccountLimits

`func NewAccountLimits(monthlyUsageCycle UsageCycle, limits Limits, current Current, ) *AccountLimits`

NewAccountLimits instantiates a new AccountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLimitsWithDefaults

`func NewAccountLimitsWithDefaults() *AccountLimits`

NewAccountLimitsWithDefaults instantiates a new AccountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthlyUsageCycle

`func (o *AccountLimits) GetMonthlyUsageCycle() UsageCycle`

GetMonthlyUsageCycle returns the MonthlyUsageCycle field if non-nil, zero value otherwise.

### GetMonthlyUsageCycleOk

`func (o *AccountLimits) GetMonthlyUsageCycleOk() (*UsageCycle, bool)`

GetMonthlyUsageCycleOk returns a tuple with the MonthlyUsageCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyUsageCycle

`func (o *AccountLimits) SetMonthlyUsageCycle(v UsageCycle)`

SetMonthlyUsageCycle sets MonthlyUsageCycle field to given value.


### GetLimits

`func (o *AccountLimits) GetLimits() Limits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *AccountLimits) GetLimitsOk() (*Limits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *AccountLimits) SetLimits(v Limits)`

SetLimits sets Limits field to given value.


### GetCurrent

`func (o *AccountLimits) GetCurrent() Current`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AccountLimits) GetCurrentOk() (*Current, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AccountLimits) SetCurrent(v Current)`

SetCurrent sets Current field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


