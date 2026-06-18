# ExampleWebhookDispatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**WebhookDispatchStatus**](WebhookDispatchStatus.md) |  | 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 
**RemovedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewExampleWebhookDispatch

`func NewExampleWebhookDispatch(status WebhookDispatchStatus, ) *ExampleWebhookDispatch`

NewExampleWebhookDispatch instantiates a new ExampleWebhookDispatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExampleWebhookDispatchWithDefaults

`func NewExampleWebhookDispatchWithDefaults() *ExampleWebhookDispatch`

NewExampleWebhookDispatchWithDefaults instantiates a new ExampleWebhookDispatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ExampleWebhookDispatch) GetStatus() WebhookDispatchStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExampleWebhookDispatch) GetStatusOk() (*WebhookDispatchStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExampleWebhookDispatch) SetStatus(v WebhookDispatchStatus)`

SetStatus sets Status field to given value.


### GetFinishedAt

`func (o *ExampleWebhookDispatch) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ExampleWebhookDispatch) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ExampleWebhookDispatch) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ExampleWebhookDispatch) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *ExampleWebhookDispatch) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ExampleWebhookDispatch) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetRemovedAt

`func (o *ExampleWebhookDispatch) GetRemovedAt() time.Time`

GetRemovedAt returns the RemovedAt field if non-nil, zero value otherwise.

### GetRemovedAtOk

`func (o *ExampleWebhookDispatch) GetRemovedAtOk() (*time.Time, bool)`

GetRemovedAtOk returns a tuple with the RemovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAt

`func (o *ExampleWebhookDispatch) SetRemovedAt(v time.Time)`

SetRemovedAt sets RemovedAt field to given value.

### HasRemovedAt

`func (o *ExampleWebhookDispatch) HasRemovedAt() bool`

HasRemovedAt returns a boolean if a field has been set.

### SetRemovedAtNil

`func (o *ExampleWebhookDispatch) SetRemovedAtNil(b bool)`

 SetRemovedAtNil sets the value for RemovedAt to be an explicit nil

### UnsetRemovedAt
`func (o *ExampleWebhookDispatch) UnsetRemovedAt()`

UnsetRemovedAt ensures that no value is present for RemovedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


