# RunShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ActId** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**ActorTaskId** | Pointer to **NullableString** |  | [optional] 
**Status** | [**ActorJobStatus**](ActorJobStatus.md) |  | 
**StartedAt** | **time.Time** |  | 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 
**BuildId** | **string** |  | 
**BuildNumber** | Pointer to **string** |  | [optional] 
**BuildNumberInt** | Pointer to **int32** |  | [optional] 
**Meta** | [**RunMeta**](RunMeta.md) |  | 
**UsageTotalUsd** | **float32** |  | 
**DefaultKeyValueStoreId** | **string** |  | 
**DefaultDatasetId** | **string** |  | 
**DefaultRequestQueueId** | **string** |  | 

## Methods

### NewRunShort

`func NewRunShort(id string, actId string, status ActorJobStatus, startedAt time.Time, buildId string, meta RunMeta, usageTotalUsd float32, defaultKeyValueStoreId string, defaultDatasetId string, defaultRequestQueueId string, ) *RunShort`

NewRunShort instantiates a new RunShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunShortWithDefaults

`func NewRunShortWithDefaults() *RunShort`

NewRunShortWithDefaults instantiates a new RunShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunShort) SetId(v string)`

SetId sets Id field to given value.


### GetActId

`func (o *RunShort) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *RunShort) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *RunShort) SetActId(v string)`

SetActId sets ActId field to given value.


### GetUserId

`func (o *RunShort) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RunShort) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RunShort) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RunShort) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetActorTaskId

`func (o *RunShort) GetActorTaskId() string`

GetActorTaskId returns the ActorTaskId field if non-nil, zero value otherwise.

### GetActorTaskIdOk

`func (o *RunShort) GetActorTaskIdOk() (*string, bool)`

GetActorTaskIdOk returns a tuple with the ActorTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTaskId

`func (o *RunShort) SetActorTaskId(v string)`

SetActorTaskId sets ActorTaskId field to given value.

### HasActorTaskId

`func (o *RunShort) HasActorTaskId() bool`

HasActorTaskId returns a boolean if a field has been set.

### SetActorTaskIdNil

`func (o *RunShort) SetActorTaskIdNil(b bool)`

 SetActorTaskIdNil sets the value for ActorTaskId to be an explicit nil

### UnsetActorTaskId
`func (o *RunShort) UnsetActorTaskId()`

UnsetActorTaskId ensures that no value is present for ActorTaskId, not even an explicit nil
### GetStatus

`func (o *RunShort) GetStatus() ActorJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunShort) GetStatusOk() (*ActorJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunShort) SetStatus(v ActorJobStatus)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *RunShort) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RunShort) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RunShort) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *RunShort) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *RunShort) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *RunShort) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *RunShort) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *RunShort) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *RunShort) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetBuildId

`func (o *RunShort) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *RunShort) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *RunShort) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.


### GetBuildNumber

`func (o *RunShort) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *RunShort) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *RunShort) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *RunShort) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetBuildNumberInt

`func (o *RunShort) GetBuildNumberInt() int32`

GetBuildNumberInt returns the BuildNumberInt field if non-nil, zero value otherwise.

### GetBuildNumberIntOk

`func (o *RunShort) GetBuildNumberIntOk() (*int32, bool)`

GetBuildNumberIntOk returns a tuple with the BuildNumberInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumberInt

`func (o *RunShort) SetBuildNumberInt(v int32)`

SetBuildNumberInt sets BuildNumberInt field to given value.

### HasBuildNumberInt

`func (o *RunShort) HasBuildNumberInt() bool`

HasBuildNumberInt returns a boolean if a field has been set.

### GetMeta

`func (o *RunShort) GetMeta() RunMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RunShort) GetMetaOk() (*RunMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RunShort) SetMeta(v RunMeta)`

SetMeta sets Meta field to given value.


### GetUsageTotalUsd

`func (o *RunShort) GetUsageTotalUsd() float32`

GetUsageTotalUsd returns the UsageTotalUsd field if non-nil, zero value otherwise.

### GetUsageTotalUsdOk

`func (o *RunShort) GetUsageTotalUsdOk() (*float32, bool)`

GetUsageTotalUsdOk returns a tuple with the UsageTotalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTotalUsd

`func (o *RunShort) SetUsageTotalUsd(v float32)`

SetUsageTotalUsd sets UsageTotalUsd field to given value.


### GetDefaultKeyValueStoreId

`func (o *RunShort) GetDefaultKeyValueStoreId() string`

GetDefaultKeyValueStoreId returns the DefaultKeyValueStoreId field if non-nil, zero value otherwise.

### GetDefaultKeyValueStoreIdOk

`func (o *RunShort) GetDefaultKeyValueStoreIdOk() (*string, bool)`

GetDefaultKeyValueStoreIdOk returns a tuple with the DefaultKeyValueStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyValueStoreId

`func (o *RunShort) SetDefaultKeyValueStoreId(v string)`

SetDefaultKeyValueStoreId sets DefaultKeyValueStoreId field to given value.


### GetDefaultDatasetId

`func (o *RunShort) GetDefaultDatasetId() string`

GetDefaultDatasetId returns the DefaultDatasetId field if non-nil, zero value otherwise.

### GetDefaultDatasetIdOk

`func (o *RunShort) GetDefaultDatasetIdOk() (*string, bool)`

GetDefaultDatasetIdOk returns a tuple with the DefaultDatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatasetId

`func (o *RunShort) SetDefaultDatasetId(v string)`

SetDefaultDatasetId sets DefaultDatasetId field to given value.


### GetDefaultRequestQueueId

`func (o *RunShort) GetDefaultRequestQueueId() string`

GetDefaultRequestQueueId returns the DefaultRequestQueueId field if non-nil, zero value otherwise.

### GetDefaultRequestQueueIdOk

`func (o *RunShort) GetDefaultRequestQueueIdOk() (*string, bool)`

GetDefaultRequestQueueIdOk returns a tuple with the DefaultRequestQueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRequestQueueId

`func (o *RunShort) SetDefaultRequestQueueId(v string)`

SetDefaultRequestQueueId sets DefaultRequestQueueId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


