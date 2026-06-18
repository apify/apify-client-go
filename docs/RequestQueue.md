# RequestQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier assigned to the request queue. | 
**Name** | Pointer to **NullableString** | The name of the request queue. | [optional] 
**UserId** | **string** | The ID of the user who owns the request queue. | 
**ActId** | Pointer to **NullableString** | The ID of the Actor that created this request queue. | [optional] 
**ActRunId** | Pointer to **NullableString** | The ID of the Actor run that created this request queue. | [optional] 
**CreatedAt** | **time.Time** | The timestamp when the request queue was created. | 
**ModifiedAt** | **time.Time** | The timestamp when the request queue was last modified. Modifications include adding, updating, or removing requests, as well as locking or unlocking requests in the request queue. | 
**AccessedAt** | **time.Time** | The timestamp when the request queue was last accessed. | 
**TotalRequestCount** | **int32** | The total number of requests in the request queue. | 
**HandledRequestCount** | **int32** | The number of requests that have been handled. | 
**PendingRequestCount** | **int32** | The number of requests that are pending and have not been handled yet. | 
**HadMultipleClients** | **bool** | Whether the request queue has been accessed by multiple different clients. | 
**ConsoleUrl** | **string** | The URL to view the request queue in the Apify console. | 
**Stats** | Pointer to [**RequestQueueStats**](RequestQueueStats.md) |  | [optional] 
**GeneralAccess** | Pointer to [**GeneralAccess**](GeneralAccess.md) |  | [optional] 

## Methods

### NewRequestQueue

`func NewRequestQueue(id string, userId string, createdAt time.Time, modifiedAt time.Time, accessedAt time.Time, totalRequestCount int32, handledRequestCount int32, pendingRequestCount int32, hadMultipleClients bool, consoleUrl string, ) *RequestQueue`

NewRequestQueue instantiates a new RequestQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestQueueWithDefaults

`func NewRequestQueueWithDefaults() *RequestQueue`

NewRequestQueueWithDefaults instantiates a new RequestQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestQueue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestQueue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestQueue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RequestQueue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestQueue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestQueue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestQueue) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RequestQueue) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RequestQueue) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserId

`func (o *RequestQueue) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RequestQueue) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RequestQueue) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetActId

`func (o *RequestQueue) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *RequestQueue) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *RequestQueue) SetActId(v string)`

SetActId sets ActId field to given value.

### HasActId

`func (o *RequestQueue) HasActId() bool`

HasActId returns a boolean if a field has been set.

### SetActIdNil

`func (o *RequestQueue) SetActIdNil(b bool)`

 SetActIdNil sets the value for ActId to be an explicit nil

### UnsetActId
`func (o *RequestQueue) UnsetActId()`

UnsetActId ensures that no value is present for ActId, not even an explicit nil
### GetActRunId

`func (o *RequestQueue) GetActRunId() string`

GetActRunId returns the ActRunId field if non-nil, zero value otherwise.

### GetActRunIdOk

`func (o *RequestQueue) GetActRunIdOk() (*string, bool)`

GetActRunIdOk returns a tuple with the ActRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActRunId

`func (o *RequestQueue) SetActRunId(v string)`

SetActRunId sets ActRunId field to given value.

### HasActRunId

`func (o *RequestQueue) HasActRunId() bool`

HasActRunId returns a boolean if a field has been set.

### SetActRunIdNil

`func (o *RequestQueue) SetActRunIdNil(b bool)`

 SetActRunIdNil sets the value for ActRunId to be an explicit nil

### UnsetActRunId
`func (o *RequestQueue) UnsetActRunId()`

UnsetActRunId ensures that no value is present for ActRunId, not even an explicit nil
### GetCreatedAt

`func (o *RequestQueue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestQueue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestQueue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *RequestQueue) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RequestQueue) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RequestQueue) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetAccessedAt

`func (o *RequestQueue) GetAccessedAt() time.Time`

GetAccessedAt returns the AccessedAt field if non-nil, zero value otherwise.

### GetAccessedAtOk

`func (o *RequestQueue) GetAccessedAtOk() (*time.Time, bool)`

GetAccessedAtOk returns a tuple with the AccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedAt

`func (o *RequestQueue) SetAccessedAt(v time.Time)`

SetAccessedAt sets AccessedAt field to given value.


### GetTotalRequestCount

`func (o *RequestQueue) GetTotalRequestCount() int32`

GetTotalRequestCount returns the TotalRequestCount field if non-nil, zero value otherwise.

### GetTotalRequestCountOk

`func (o *RequestQueue) GetTotalRequestCountOk() (*int32, bool)`

GetTotalRequestCountOk returns a tuple with the TotalRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequestCount

`func (o *RequestQueue) SetTotalRequestCount(v int32)`

SetTotalRequestCount sets TotalRequestCount field to given value.


### GetHandledRequestCount

`func (o *RequestQueue) GetHandledRequestCount() int32`

GetHandledRequestCount returns the HandledRequestCount field if non-nil, zero value otherwise.

### GetHandledRequestCountOk

`func (o *RequestQueue) GetHandledRequestCountOk() (*int32, bool)`

GetHandledRequestCountOk returns a tuple with the HandledRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledRequestCount

`func (o *RequestQueue) SetHandledRequestCount(v int32)`

SetHandledRequestCount sets HandledRequestCount field to given value.


### GetPendingRequestCount

`func (o *RequestQueue) GetPendingRequestCount() int32`

GetPendingRequestCount returns the PendingRequestCount field if non-nil, zero value otherwise.

### GetPendingRequestCountOk

`func (o *RequestQueue) GetPendingRequestCountOk() (*int32, bool)`

GetPendingRequestCountOk returns a tuple with the PendingRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRequestCount

`func (o *RequestQueue) SetPendingRequestCount(v int32)`

SetPendingRequestCount sets PendingRequestCount field to given value.


### GetHadMultipleClients

`func (o *RequestQueue) GetHadMultipleClients() bool`

GetHadMultipleClients returns the HadMultipleClients field if non-nil, zero value otherwise.

### GetHadMultipleClientsOk

`func (o *RequestQueue) GetHadMultipleClientsOk() (*bool, bool)`

GetHadMultipleClientsOk returns a tuple with the HadMultipleClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadMultipleClients

`func (o *RequestQueue) SetHadMultipleClients(v bool)`

SetHadMultipleClients sets HadMultipleClients field to given value.


### GetConsoleUrl

`func (o *RequestQueue) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *RequestQueue) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *RequestQueue) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.


### GetStats

`func (o *RequestQueue) GetStats() RequestQueueStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *RequestQueue) GetStatsOk() (*RequestQueueStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *RequestQueue) SetStats(v RequestQueueStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *RequestQueue) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetGeneralAccess

`func (o *RequestQueue) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *RequestQueue) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *RequestQueue) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.

### HasGeneralAccess

`func (o *RequestQueue) HasGeneralAccess() bool`

HasGeneralAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


