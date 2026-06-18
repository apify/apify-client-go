# LockedRequestQueueHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The maximum number of requests returned. | 
**QueueModifiedAt** | **time.Time** | The timestamp when the request queue was last modified. Modifications include adding, updating, or removing requests, as well as locking or unlocking requests in the request queue. | 
**QueueHasLockedRequests** | Pointer to **bool** | Whether the request queue contains requests locked by any client (either the one calling the endpoint or a different one). | [optional] 
**ClientKey** | Pointer to **string** | The client key used for locking the requests. | [optional] 
**HadMultipleClients** | **bool** | Whether the request queue has been accessed by multiple different clients. | 
**LockSecs** | **int32** | The number of seconds the locks will be held. | 
**Items** | [**[]LockedHeadRequest**](LockedHeadRequest.md) | The array of locked requests from the request queue head. | 

## Methods

### NewLockedRequestQueueHead

`func NewLockedRequestQueueHead(limit int32, queueModifiedAt time.Time, hadMultipleClients bool, lockSecs int32, items []LockedHeadRequest, ) *LockedRequestQueueHead`

NewLockedRequestQueueHead instantiates a new LockedRequestQueueHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockedRequestQueueHeadWithDefaults

`func NewLockedRequestQueueHeadWithDefaults() *LockedRequestQueueHead`

NewLockedRequestQueueHeadWithDefaults instantiates a new LockedRequestQueueHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *LockedRequestQueueHead) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LockedRequestQueueHead) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LockedRequestQueueHead) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetQueueModifiedAt

`func (o *LockedRequestQueueHead) GetQueueModifiedAt() time.Time`

GetQueueModifiedAt returns the QueueModifiedAt field if non-nil, zero value otherwise.

### GetQueueModifiedAtOk

`func (o *LockedRequestQueueHead) GetQueueModifiedAtOk() (*time.Time, bool)`

GetQueueModifiedAtOk returns a tuple with the QueueModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueModifiedAt

`func (o *LockedRequestQueueHead) SetQueueModifiedAt(v time.Time)`

SetQueueModifiedAt sets QueueModifiedAt field to given value.


### GetQueueHasLockedRequests

`func (o *LockedRequestQueueHead) GetQueueHasLockedRequests() bool`

GetQueueHasLockedRequests returns the QueueHasLockedRequests field if non-nil, zero value otherwise.

### GetQueueHasLockedRequestsOk

`func (o *LockedRequestQueueHead) GetQueueHasLockedRequestsOk() (*bool, bool)`

GetQueueHasLockedRequestsOk returns a tuple with the QueueHasLockedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueHasLockedRequests

`func (o *LockedRequestQueueHead) SetQueueHasLockedRequests(v bool)`

SetQueueHasLockedRequests sets QueueHasLockedRequests field to given value.

### HasQueueHasLockedRequests

`func (o *LockedRequestQueueHead) HasQueueHasLockedRequests() bool`

HasQueueHasLockedRequests returns a boolean if a field has been set.

### GetClientKey

`func (o *LockedRequestQueueHead) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *LockedRequestQueueHead) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *LockedRequestQueueHead) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *LockedRequestQueueHead) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetHadMultipleClients

`func (o *LockedRequestQueueHead) GetHadMultipleClients() bool`

GetHadMultipleClients returns the HadMultipleClients field if non-nil, zero value otherwise.

### GetHadMultipleClientsOk

`func (o *LockedRequestQueueHead) GetHadMultipleClientsOk() (*bool, bool)`

GetHadMultipleClientsOk returns a tuple with the HadMultipleClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadMultipleClients

`func (o *LockedRequestQueueHead) SetHadMultipleClients(v bool)`

SetHadMultipleClients sets HadMultipleClients field to given value.


### GetLockSecs

`func (o *LockedRequestQueueHead) GetLockSecs() int32`

GetLockSecs returns the LockSecs field if non-nil, zero value otherwise.

### GetLockSecsOk

`func (o *LockedRequestQueueHead) GetLockSecsOk() (*int32, bool)`

GetLockSecsOk returns a tuple with the LockSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockSecs

`func (o *LockedRequestQueueHead) SetLockSecs(v int32)`

SetLockSecs sets LockSecs field to given value.


### GetItems

`func (o *LockedRequestQueueHead) GetItems() []LockedHeadRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LockedRequestQueueHead) GetItemsOk() (*[]LockedHeadRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LockedRequestQueueHead) SetItems(v []LockedHeadRequest)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


