# BuildShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ActId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Status** | [**ActorJobStatus**](ActorJobStatus.md) |  | 
**StartedAt** | **time.Time** |  | 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 
**UsageTotalUsd** | **float32** |  | 
**BuildNumber** | **string** |  | 
**BuildNumberInt** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to [**BuildsMeta**](BuildsMeta.md) |  | [optional] 

## Methods

### NewBuildShort

`func NewBuildShort(id string, status ActorJobStatus, startedAt time.Time, usageTotalUsd float32, buildNumber string, ) *BuildShort`

NewBuildShort instantiates a new BuildShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildShortWithDefaults

`func NewBuildShortWithDefaults() *BuildShort`

NewBuildShortWithDefaults instantiates a new BuildShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BuildShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuildShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuildShort) SetId(v string)`

SetId sets Id field to given value.


### GetActId

`func (o *BuildShort) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *BuildShort) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *BuildShort) SetActId(v string)`

SetActId sets ActId field to given value.

### HasActId

`func (o *BuildShort) HasActId() bool`

HasActId returns a boolean if a field has been set.

### GetUserId

`func (o *BuildShort) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BuildShort) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BuildShort) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BuildShort) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *BuildShort) GetStatus() ActorJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildShort) GetStatusOk() (*ActorJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildShort) SetStatus(v ActorJobStatus)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *BuildShort) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BuildShort) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BuildShort) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *BuildShort) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *BuildShort) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *BuildShort) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *BuildShort) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *BuildShort) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *BuildShort) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetUsageTotalUsd

`func (o *BuildShort) GetUsageTotalUsd() float32`

GetUsageTotalUsd returns the UsageTotalUsd field if non-nil, zero value otherwise.

### GetUsageTotalUsdOk

`func (o *BuildShort) GetUsageTotalUsdOk() (*float32, bool)`

GetUsageTotalUsdOk returns a tuple with the UsageTotalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTotalUsd

`func (o *BuildShort) SetUsageTotalUsd(v float32)`

SetUsageTotalUsd sets UsageTotalUsd field to given value.


### GetBuildNumber

`func (o *BuildShort) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *BuildShort) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *BuildShort) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.


### GetBuildNumberInt

`func (o *BuildShort) GetBuildNumberInt() int32`

GetBuildNumberInt returns the BuildNumberInt field if non-nil, zero value otherwise.

### GetBuildNumberIntOk

`func (o *BuildShort) GetBuildNumberIntOk() (*int32, bool)`

GetBuildNumberIntOk returns a tuple with the BuildNumberInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumberInt

`func (o *BuildShort) SetBuildNumberInt(v int32)`

SetBuildNumberInt sets BuildNumberInt field to given value.

### HasBuildNumberInt

`func (o *BuildShort) HasBuildNumberInt() bool`

HasBuildNumberInt returns a boolean if a field has been set.

### GetMeta

`func (o *BuildShort) GetMeta() BuildsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BuildShort) GetMetaOk() (*BuildsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BuildShort) SetMeta(v BuildsMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BuildShort) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


