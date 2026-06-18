# RequestQueueHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The maximum number of requests returned. | 
**QueueModifiedAt** | **time.Time** | The timestamp when the request queue was last modified. Modifications include adding, updating, or removing requests, as well as locking or unlocking requests in the request queue. | 
**HadMultipleClients** | **bool** | Whether the request queue has been accessed by multiple different clients. | 
**Items** | [**[]HeadRequest**](HeadRequest.md) | The array of requests from the request queue head. | 

## Methods

### NewRequestQueueHead

`func NewRequestQueueHead(limit int32, queueModifiedAt time.Time, hadMultipleClients bool, items []HeadRequest, ) *RequestQueueHead`

NewRequestQueueHead instantiates a new RequestQueueHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestQueueHeadWithDefaults

`func NewRequestQueueHeadWithDefaults() *RequestQueueHead`

NewRequestQueueHeadWithDefaults instantiates a new RequestQueueHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *RequestQueueHead) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RequestQueueHead) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RequestQueueHead) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetQueueModifiedAt

`func (o *RequestQueueHead) GetQueueModifiedAt() time.Time`

GetQueueModifiedAt returns the QueueModifiedAt field if non-nil, zero value otherwise.

### GetQueueModifiedAtOk

`func (o *RequestQueueHead) GetQueueModifiedAtOk() (*time.Time, bool)`

GetQueueModifiedAtOk returns a tuple with the QueueModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueModifiedAt

`func (o *RequestQueueHead) SetQueueModifiedAt(v time.Time)`

SetQueueModifiedAt sets QueueModifiedAt field to given value.


### GetHadMultipleClients

`func (o *RequestQueueHead) GetHadMultipleClients() bool`

GetHadMultipleClients returns the HadMultipleClients field if non-nil, zero value otherwise.

### GetHadMultipleClientsOk

`func (o *RequestQueueHead) GetHadMultipleClientsOk() (*bool, bool)`

GetHadMultipleClientsOk returns a tuple with the HadMultipleClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadMultipleClients

`func (o *RequestQueueHead) SetHadMultipleClients(v bool)`

SetHadMultipleClients sets HadMultipleClients field to given value.


### GetItems

`func (o *RequestQueueHead) GetItems() []HeadRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RequestQueueHead) GetItemsOk() (*[]HeadRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RequestQueueHead) SetItems(v []HeadRequest)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


