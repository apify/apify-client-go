# Build

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ActId** | **string** |  | 
**UserId** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | [**ActorJobStatus**](ActorJobStatus.md) |  | 
**Meta** | [**BuildsMeta**](BuildsMeta.md) |  | 
**Stats** | Pointer to [**NullableBuildStats**](BuildStats.md) |  | [optional] 
**Options** | Pointer to [**NullableBuildOptions**](BuildOptions.md) |  | [optional] 
**Usage** | Pointer to [**NullableBuildUsage**](BuildUsage.md) |  | [optional] 
**UsageTotalUsd** | Pointer to **NullableFloat32** | Total cost in USD for this build. Requires authentication token to access. | [optional] 
**UsageUsd** | Pointer to [**NullableBuildUsage**](BuildUsage.md) | Platform usage costs breakdown in USD for this build. Requires authentication token to access. | [optional] 
**InputSchema** | Pointer to **NullableString** |  | [optional] 
**Readme** | Pointer to **NullableString** |  | [optional] 
**BuildNumber** | **string** |  | 
**ActVersion** | Pointer to [**BuildActVersion**](BuildActVersion.md) |  | [optional] 
**ActorDefinition** | Pointer to [**NullableActorDefinition**](ActorDefinition.md) |  | [optional] 

## Methods

### NewBuild

`func NewBuild(id string, actId string, userId string, startedAt time.Time, status ActorJobStatus, meta BuildsMeta, buildNumber string, ) *Build`

NewBuild instantiates a new Build object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildWithDefaults

`func NewBuildWithDefaults() *Build`

NewBuildWithDefaults instantiates a new Build object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Build) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Build) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Build) SetId(v string)`

SetId sets Id field to given value.


### GetActId

`func (o *Build) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *Build) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *Build) SetActId(v string)`

SetActId sets ActId field to given value.


### GetUserId

`func (o *Build) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Build) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Build) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStartedAt

`func (o *Build) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Build) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Build) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *Build) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Build) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Build) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Build) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *Build) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *Build) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetStatus

`func (o *Build) GetStatus() ActorJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Build) GetStatusOk() (*ActorJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Build) SetStatus(v ActorJobStatus)`

SetStatus sets Status field to given value.


### GetMeta

`func (o *Build) GetMeta() BuildsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Build) GetMetaOk() (*BuildsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Build) SetMeta(v BuildsMeta)`

SetMeta sets Meta field to given value.


### GetStats

`func (o *Build) GetStats() BuildStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Build) GetStatsOk() (*BuildStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Build) SetStats(v BuildStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Build) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *Build) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *Build) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetOptions

`func (o *Build) GetOptions() BuildOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Build) GetOptionsOk() (*BuildOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Build) SetOptions(v BuildOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Build) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *Build) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *Build) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetUsage

`func (o *Build) GetUsage() BuildUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Build) GetUsageOk() (*BuildUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Build) SetUsage(v BuildUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Build) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *Build) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *Build) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetUsageTotalUsd

`func (o *Build) GetUsageTotalUsd() float32`

GetUsageTotalUsd returns the UsageTotalUsd field if non-nil, zero value otherwise.

### GetUsageTotalUsdOk

`func (o *Build) GetUsageTotalUsdOk() (*float32, bool)`

GetUsageTotalUsdOk returns a tuple with the UsageTotalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTotalUsd

`func (o *Build) SetUsageTotalUsd(v float32)`

SetUsageTotalUsd sets UsageTotalUsd field to given value.

### HasUsageTotalUsd

`func (o *Build) HasUsageTotalUsd() bool`

HasUsageTotalUsd returns a boolean if a field has been set.

### SetUsageTotalUsdNil

`func (o *Build) SetUsageTotalUsdNil(b bool)`

 SetUsageTotalUsdNil sets the value for UsageTotalUsd to be an explicit nil

### UnsetUsageTotalUsd
`func (o *Build) UnsetUsageTotalUsd()`

UnsetUsageTotalUsd ensures that no value is present for UsageTotalUsd, not even an explicit nil
### GetUsageUsd

`func (o *Build) GetUsageUsd() BuildUsage`

GetUsageUsd returns the UsageUsd field if non-nil, zero value otherwise.

### GetUsageUsdOk

`func (o *Build) GetUsageUsdOk() (*BuildUsage, bool)`

GetUsageUsdOk returns a tuple with the UsageUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUsd

`func (o *Build) SetUsageUsd(v BuildUsage)`

SetUsageUsd sets UsageUsd field to given value.

### HasUsageUsd

`func (o *Build) HasUsageUsd() bool`

HasUsageUsd returns a boolean if a field has been set.

### SetUsageUsdNil

`func (o *Build) SetUsageUsdNil(b bool)`

 SetUsageUsdNil sets the value for UsageUsd to be an explicit nil

### UnsetUsageUsd
`func (o *Build) UnsetUsageUsd()`

UnsetUsageUsd ensures that no value is present for UsageUsd, not even an explicit nil
### GetInputSchema

`func (o *Build) GetInputSchema() string`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *Build) GetInputSchemaOk() (*string, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *Build) SetInputSchema(v string)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *Build) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### SetInputSchemaNil

`func (o *Build) SetInputSchemaNil(b bool)`

 SetInputSchemaNil sets the value for InputSchema to be an explicit nil

### UnsetInputSchema
`func (o *Build) UnsetInputSchema()`

UnsetInputSchema ensures that no value is present for InputSchema, not even an explicit nil
### GetReadme

`func (o *Build) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *Build) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *Build) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *Build) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### SetReadmeNil

`func (o *Build) SetReadmeNil(b bool)`

 SetReadmeNil sets the value for Readme to be an explicit nil

### UnsetReadme
`func (o *Build) UnsetReadme()`

UnsetReadme ensures that no value is present for Readme, not even an explicit nil
### GetBuildNumber

`func (o *Build) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *Build) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *Build) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.


### GetActVersion

`func (o *Build) GetActVersion() BuildActVersion`

GetActVersion returns the ActVersion field if non-nil, zero value otherwise.

### GetActVersionOk

`func (o *Build) GetActVersionOk() (*BuildActVersion, bool)`

GetActVersionOk returns a tuple with the ActVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActVersion

`func (o *Build) SetActVersion(v BuildActVersion)`

SetActVersion sets ActVersion field to given value.

### HasActVersion

`func (o *Build) HasActVersion() bool`

HasActVersion returns a boolean if a field has been set.

### GetActorDefinition

`func (o *Build) GetActorDefinition() ActorDefinition`

GetActorDefinition returns the ActorDefinition field if non-nil, zero value otherwise.

### GetActorDefinitionOk

`func (o *Build) GetActorDefinitionOk() (*ActorDefinition, bool)`

GetActorDefinitionOk returns a tuple with the ActorDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDefinition

`func (o *Build) SetActorDefinition(v ActorDefinition)`

SetActorDefinition sets ActorDefinition field to given value.

### HasActorDefinition

`func (o *Build) HasActorDefinition() bool`

HasActorDefinition returns a boolean if a field has been set.

### SetActorDefinitionNil

`func (o *Build) SetActorDefinitionNil(b bool)`

 SetActorDefinitionNil sets the value for ActorDefinition to be an explicit nil

### UnsetActorDefinition
`func (o *Build) UnsetActorDefinition()`

UnsetActorDefinition ensures that no value is present for ActorDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


