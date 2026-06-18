# WebhookDispatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**WebhookId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Status** | [**WebhookDispatchStatus**](WebhookDispatchStatus.md) |  | 
**EventType** | [**WebhookEventType**](WebhookEventType.md) |  | 
**EventData** | Pointer to [**EventData**](EventData.md) |  | [optional] 
**Webhook** | Pointer to [**NullableWebhookDispatchWebhookSummary**](WebhookDispatchWebhookSummary.md) |  | [optional] 
**Calls** | Pointer to [**[]CallsInner**](CallsInner.md) |  | [optional] 

## Methods

### NewWebhookDispatch

`func NewWebhookDispatch(id string, userId string, webhookId string, createdAt time.Time, status WebhookDispatchStatus, eventType WebhookEventType, ) *WebhookDispatch`

NewWebhookDispatch instantiates a new WebhookDispatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDispatchWithDefaults

`func NewWebhookDispatchWithDefaults() *WebhookDispatch`

NewWebhookDispatchWithDefaults instantiates a new WebhookDispatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookDispatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDispatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDispatch) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *WebhookDispatch) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookDispatch) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookDispatch) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWebhookId

`func (o *WebhookDispatch) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookDispatch) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookDispatch) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetCreatedAt

`func (o *WebhookDispatch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookDispatch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookDispatch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *WebhookDispatch) GetStatus() WebhookDispatchStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDispatch) GetStatusOk() (*WebhookDispatchStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDispatch) SetStatus(v WebhookDispatchStatus)`

SetStatus sets Status field to given value.


### GetEventType

`func (o *WebhookDispatch) GetEventType() WebhookEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookDispatch) GetEventTypeOk() (*WebhookEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookDispatch) SetEventType(v WebhookEventType)`

SetEventType sets EventType field to given value.


### GetEventData

`func (o *WebhookDispatch) GetEventData() EventData`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *WebhookDispatch) GetEventDataOk() (*EventData, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *WebhookDispatch) SetEventData(v EventData)`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *WebhookDispatch) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### GetWebhook

`func (o *WebhookDispatch) GetWebhook() WebhookDispatchWebhookSummary`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *WebhookDispatch) GetWebhookOk() (*WebhookDispatchWebhookSummary, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *WebhookDispatch) SetWebhook(v WebhookDispatchWebhookSummary)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *WebhookDispatch) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *WebhookDispatch) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *WebhookDispatch) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
### GetCalls

`func (o *WebhookDispatch) GetCalls() []CallsInner`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *WebhookDispatch) GetCallsOk() (*[]CallsInner, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *WebhookDispatch) SetCalls(v []CallsInner)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *WebhookDispatch) HasCalls() bool`

HasCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


