# UpdateLimitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMonthlyUsageUsd** | Pointer to **float32** | If your platform usage in the billing period exceeds the prepaid usage, you will be charged extra. Setting this property you can update your hard limit on monthly platform usage to prevent accidental overage or to limit the extra charges.  | [optional] 
**DataRetentionDays** | Pointer to **int32** | Apify securely stores your ten most recent Actor runs indefinitely, ensuring they are always accessible. Unnamed storages and other Actor runs are automatically deleted after the retention period. If you&#39;re subscribed, you can change it to keep data for longer or to limit your usage. [Lear more](https://docs.apify.com/platform/storage/usage#data-retention).  | [optional] 

## Methods

### NewUpdateLimitsRequest

`func NewUpdateLimitsRequest() *UpdateLimitsRequest`

NewUpdateLimitsRequest instantiates a new UpdateLimitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLimitsRequestWithDefaults

`func NewUpdateLimitsRequestWithDefaults() *UpdateLimitsRequest`

NewUpdateLimitsRequestWithDefaults instantiates a new UpdateLimitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMonthlyUsageUsd

`func (o *UpdateLimitsRequest) GetMaxMonthlyUsageUsd() float32`

GetMaxMonthlyUsageUsd returns the MaxMonthlyUsageUsd field if non-nil, zero value otherwise.

### GetMaxMonthlyUsageUsdOk

`func (o *UpdateLimitsRequest) GetMaxMonthlyUsageUsdOk() (*float32, bool)`

GetMaxMonthlyUsageUsdOk returns a tuple with the MaxMonthlyUsageUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyUsageUsd

`func (o *UpdateLimitsRequest) SetMaxMonthlyUsageUsd(v float32)`

SetMaxMonthlyUsageUsd sets MaxMonthlyUsageUsd field to given value.

### HasMaxMonthlyUsageUsd

`func (o *UpdateLimitsRequest) HasMaxMonthlyUsageUsd() bool`

HasMaxMonthlyUsageUsd returns a boolean if a field has been set.

### GetDataRetentionDays

`func (o *UpdateLimitsRequest) GetDataRetentionDays() int32`

GetDataRetentionDays returns the DataRetentionDays field if non-nil, zero value otherwise.

### GetDataRetentionDaysOk

`func (o *UpdateLimitsRequest) GetDataRetentionDaysOk() (*int32, bool)`

GetDataRetentionDaysOk returns a tuple with the DataRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionDays

`func (o *UpdateLimitsRequest) SetDataRetentionDays(v int32)`

SetDataRetentionDays sets DataRetentionDays field to given value.

### HasDataRetentionDays

`func (o *UpdateLimitsRequest) HasDataRetentionDays() bool`

HasDataRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


