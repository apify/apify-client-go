# WebhookDispatchWebhookSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to [**WebhookCondition**](WebhookCondition.md) |  | [optional] 
**RequestUrl** | Pointer to **string** |  | [optional] 
**IsAdHoc** | Pointer to **bool** |  | [optional] 

## Methods

### NewWebhookDispatchWebhookSummary

`func NewWebhookDispatchWebhookSummary() *WebhookDispatchWebhookSummary`

NewWebhookDispatchWebhookSummary instantiates a new WebhookDispatchWebhookSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDispatchWebhookSummaryWithDefaults

`func NewWebhookDispatchWebhookSummaryWithDefaults() *WebhookDispatchWebhookSummary`

NewWebhookDispatchWebhookSummaryWithDefaults instantiates a new WebhookDispatchWebhookSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *WebhookDispatchWebhookSummary) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WebhookDispatchWebhookSummary) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WebhookDispatchWebhookSummary) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WebhookDispatchWebhookSummary) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetCondition

`func (o *WebhookDispatchWebhookSummary) GetCondition() WebhookCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WebhookDispatchWebhookSummary) GetConditionOk() (*WebhookCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WebhookDispatchWebhookSummary) SetCondition(v WebhookCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WebhookDispatchWebhookSummary) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetRequestUrl

`func (o *WebhookDispatchWebhookSummary) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *WebhookDispatchWebhookSummary) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *WebhookDispatchWebhookSummary) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *WebhookDispatchWebhookSummary) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### GetIsAdHoc

`func (o *WebhookDispatchWebhookSummary) GetIsAdHoc() bool`

GetIsAdHoc returns the IsAdHoc field if non-nil, zero value otherwise.

### GetIsAdHocOk

`func (o *WebhookDispatchWebhookSummary) GetIsAdHocOk() (*bool, bool)`

GetIsAdHocOk returns a tuple with the IsAdHoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdHoc

`func (o *WebhookDispatchWebhookSummary) SetIsAdHoc(v bool)`

SetIsAdHoc sets IsAdHoc field to given value.

### HasIsAdHoc

`func (o *WebhookDispatchWebhookSummary) HasIsAdHoc() bool`

HasIsAdHoc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


