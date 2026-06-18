# RequestQueueShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier assigned to the request queue. | 
**Name** | **string** | The name of the request queue. | 
**UserId** | **string** | The ID of the user who owns the request queue. | 
**Username** | **string** | The username of the user who owns the request queue. | 
**CreatedAt** | **time.Time** | The timestamp when the request queue was created. | 
**ModifiedAt** | **time.Time** | The timestamp when the request queue was last modified. Modifications include adding, updating, or removing requests, as well as locking or unlocking requests in the request queue. | 
**AccessedAt** | **time.Time** | The timestamp when the request queue was last accessed. | 
**ExpireAt** | Pointer to **time.Time** | The timestamp when the request queue will expire and be deleted. | [optional] 
**TotalRequestCount** | **int32** | The total number of requests in the request queue. | 
**HandledRequestCount** | **int32** | The number of requests that have been handled. | 
**PendingRequestCount** | **int32** | The number of requests that are pending and have not been handled yet. | 
**ActId** | Pointer to **NullableString** | The ID of the Actor that created this request queue. | [optional] 
**ActRunId** | Pointer to **NullableString** | The ID of the Actor run that created this request queue. | [optional] 
**HadMultipleClients** | **bool** | Whether the request queue has been accessed by multiple different clients. | 
**GeneralAccess** | Pointer to [**GeneralAccess**](GeneralAccess.md) |  | [optional] 
**Stats** | Pointer to [**RequestQueueStats**](RequestQueueStats.md) |  | [optional] 

## Methods

### NewRequestQueueShort

`func NewRequestQueueShort(id string, name string, userId string, username string, createdAt time.Time, modifiedAt time.Time, accessedAt time.Time, totalRequestCount int32, handledRequestCount int32, pendingRequestCount int32, hadMultipleClients bool, ) *RequestQueueShort`

NewRequestQueueShort instantiates a new RequestQueueShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestQueueShortWithDefaults

`func NewRequestQueueShortWithDefaults() *RequestQueueShort`

NewRequestQueueShortWithDefaults instantiates a new RequestQueueShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestQueueShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestQueueShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestQueueShort) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RequestQueueShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestQueueShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestQueueShort) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *RequestQueueShort) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RequestQueueShort) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RequestQueueShort) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUsername

`func (o *RequestQueueShort) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RequestQueueShort) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RequestQueueShort) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCreatedAt

`func (o *RequestQueueShort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestQueueShort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestQueueShort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *RequestQueueShort) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RequestQueueShort) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RequestQueueShort) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetAccessedAt

`func (o *RequestQueueShort) GetAccessedAt() time.Time`

GetAccessedAt returns the AccessedAt field if non-nil, zero value otherwise.

### GetAccessedAtOk

`func (o *RequestQueueShort) GetAccessedAtOk() (*time.Time, bool)`

GetAccessedAtOk returns a tuple with the AccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedAt

`func (o *RequestQueueShort) SetAccessedAt(v time.Time)`

SetAccessedAt sets AccessedAt field to given value.


### GetExpireAt

`func (o *RequestQueueShort) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *RequestQueueShort) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *RequestQueueShort) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *RequestQueueShort) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetTotalRequestCount

`func (o *RequestQueueShort) GetTotalRequestCount() int32`

GetTotalRequestCount returns the TotalRequestCount field if non-nil, zero value otherwise.

### GetTotalRequestCountOk

`func (o *RequestQueueShort) GetTotalRequestCountOk() (*int32, bool)`

GetTotalRequestCountOk returns a tuple with the TotalRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequestCount

`func (o *RequestQueueShort) SetTotalRequestCount(v int32)`

SetTotalRequestCount sets TotalRequestCount field to given value.


### GetHandledRequestCount

`func (o *RequestQueueShort) GetHandledRequestCount() int32`

GetHandledRequestCount returns the HandledRequestCount field if non-nil, zero value otherwise.

### GetHandledRequestCountOk

`func (o *RequestQueueShort) GetHandledRequestCountOk() (*int32, bool)`

GetHandledRequestCountOk returns a tuple with the HandledRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledRequestCount

`func (o *RequestQueueShort) SetHandledRequestCount(v int32)`

SetHandledRequestCount sets HandledRequestCount field to given value.


### GetPendingRequestCount

`func (o *RequestQueueShort) GetPendingRequestCount() int32`

GetPendingRequestCount returns the PendingRequestCount field if non-nil, zero value otherwise.

### GetPendingRequestCountOk

`func (o *RequestQueueShort) GetPendingRequestCountOk() (*int32, bool)`

GetPendingRequestCountOk returns a tuple with the PendingRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRequestCount

`func (o *RequestQueueShort) SetPendingRequestCount(v int32)`

SetPendingRequestCount sets PendingRequestCount field to given value.


### GetActId

`func (o *RequestQueueShort) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *RequestQueueShort) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *RequestQueueShort) SetActId(v string)`

SetActId sets ActId field to given value.

### HasActId

`func (o *RequestQueueShort) HasActId() bool`

HasActId returns a boolean if a field has been set.

### SetActIdNil

`func (o *RequestQueueShort) SetActIdNil(b bool)`

 SetActIdNil sets the value for ActId to be an explicit nil

### UnsetActId
`func (o *RequestQueueShort) UnsetActId()`

UnsetActId ensures that no value is present for ActId, not even an explicit nil
### GetActRunId

`func (o *RequestQueueShort) GetActRunId() string`

GetActRunId returns the ActRunId field if non-nil, zero value otherwise.

### GetActRunIdOk

`func (o *RequestQueueShort) GetActRunIdOk() (*string, bool)`

GetActRunIdOk returns a tuple with the ActRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActRunId

`func (o *RequestQueueShort) SetActRunId(v string)`

SetActRunId sets ActRunId field to given value.

### HasActRunId

`func (o *RequestQueueShort) HasActRunId() bool`

HasActRunId returns a boolean if a field has been set.

### SetActRunIdNil

`func (o *RequestQueueShort) SetActRunIdNil(b bool)`

 SetActRunIdNil sets the value for ActRunId to be an explicit nil

### UnsetActRunId
`func (o *RequestQueueShort) UnsetActRunId()`

UnsetActRunId ensures that no value is present for ActRunId, not even an explicit nil
### GetHadMultipleClients

`func (o *RequestQueueShort) GetHadMultipleClients() bool`

GetHadMultipleClients returns the HadMultipleClients field if non-nil, zero value otherwise.

### GetHadMultipleClientsOk

`func (o *RequestQueueShort) GetHadMultipleClientsOk() (*bool, bool)`

GetHadMultipleClientsOk returns a tuple with the HadMultipleClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadMultipleClients

`func (o *RequestQueueShort) SetHadMultipleClients(v bool)`

SetHadMultipleClients sets HadMultipleClients field to given value.


### GetGeneralAccess

`func (o *RequestQueueShort) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *RequestQueueShort) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *RequestQueueShort) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.

### HasGeneralAccess

`func (o *RequestQueueShort) HasGeneralAccess() bool`

HasGeneralAccess returns a boolean if a field has been set.

### GetStats

`func (o *RequestQueueShort) GetStats() RequestQueueStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *RequestQueueShort) GetStatsOk() (*RequestQueueStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *RequestQueueShort) SetStats(v RequestQueueStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *RequestQueueShort) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


