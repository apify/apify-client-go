# Run

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the Actor run. | 
**ActId** | **string** | ID of the Actor that was run. | 
**UserId** | **string** | ID of the user who started the run. | 
**ActorTaskId** | Pointer to **NullableString** | ID of the Actor task, if the run was started from a task. | [optional] 
**StartedAt** | **time.Time** | Time when the Actor run started. | 
**FinishedAt** | Pointer to **NullableTime** | Time when the Actor run finished. | [optional] 
**Status** | [**ActorJobStatus**](ActorJobStatus.md) | Current status of the Actor run. | 
**StatusMessage** | Pointer to **NullableString** | Detailed message about the run status. | [optional] 
**IsStatusMessageTerminal** | Pointer to **NullableBool** | Whether the status message is terminal (final). | [optional] 
**Meta** | [**RunMeta**](RunMeta.md) | Metadata about the Actor run. | 
**PricingInfo** | Pointer to [**ActorRunPricingInfo**](ActorRunPricingInfo.md) | Pricing information for the Actor. | [optional] 
**Stats** | [**RunStats**](RunStats.md) | Statistics of the Actor run. | 
**ChargedEventCounts** | Pointer to **map[string]int32** | A map of charged event types to their counts. The keys are event type identifiers defined by the Actor&#39;s pricing model (pay-per-event), and the values are the number of times each event was charged during this run. | [optional] 
**Options** | [**RunOptions**](RunOptions.md) | Configuration options for the Actor run. | 
**BuildId** | **string** | ID of the Actor build used for this run. | 
**ExitCode** | Pointer to **NullableInt32** | Exit code of the Actor run process. | [optional] 
**GeneralAccess** | [**GeneralAccess**](GeneralAccess.md) | General access level for the Actor run. | 
**DefaultKeyValueStoreId** | **string** | ID of the default key-value store for this run. | 
**DefaultDatasetId** | **string** | ID of the default dataset for this run. | 
**DefaultRequestQueueId** | **string** | ID of the default request queue for this run. | 
**StorageIds** | Pointer to [**RunStorageIds**](RunStorageIds.md) |  | [optional] 
**BuildNumber** | Pointer to **NullableString** | Build number of the Actor build used for this run. | [optional] 
**ContainerUrl** | Pointer to **string** | URL of the container running the Actor. | [optional] 
**IsContainerServerReady** | Pointer to **NullableBool** | Whether the container&#39;s HTTP server is ready to accept requests. | [optional] 
**GitBranchName** | Pointer to **NullableString** | Name of the git branch used for the Actor build. | [optional] 
**Usage** | Pointer to [**NullableRunUsage**](RunUsage.md) | Resource usage statistics for the run. | [optional] 
**UsageTotalUsd** | Pointer to **NullableFloat32** | Total cost in USD for this run. Represents what you actually pay. For run owners: includes platform usage (compute units) and/or event costs depending on the Actor&#39;s pricing model. For run non-owners: only available for Pay-Per-Event Actors (event costs only). Requires authentication token to access. | [optional] 
**UsageUsd** | Pointer to [**NullableRunUsageUsd**](RunUsageUsd.md) | Platform usage costs breakdown in USD. Only present if you own the run AND are paying for platform usage (Pay-Per-Usage, Rental, or Pay-Per-Event with usage costs like standby Actors). Not available for standard Pay-Per-Event Actors. Requires authentication token to access. | [optional] 
**Metamorphs** | Pointer to [**[]Metamorph**](Metamorph.md) | List of metamorph events that occurred during the run. | [optional] 
**PlatformUsageBillingModel** | Pointer to **string** | Indicates which party covers platform usage costs for this run. | [optional] 

## Methods

### NewRun

`func NewRun(id string, actId string, userId string, startedAt time.Time, status ActorJobStatus, meta RunMeta, stats RunStats, options RunOptions, buildId string, generalAccess GeneralAccess, defaultKeyValueStoreId string, defaultDatasetId string, defaultRequestQueueId string, ) *Run`

NewRun instantiates a new Run object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWithDefaults

`func NewRunWithDefaults() *Run`

NewRunWithDefaults instantiates a new Run object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Run) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Run) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Run) SetId(v string)`

SetId sets Id field to given value.


### GetActId

`func (o *Run) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *Run) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *Run) SetActId(v string)`

SetActId sets ActId field to given value.


### GetUserId

`func (o *Run) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Run) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Run) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetActorTaskId

`func (o *Run) GetActorTaskId() string`

GetActorTaskId returns the ActorTaskId field if non-nil, zero value otherwise.

### GetActorTaskIdOk

`func (o *Run) GetActorTaskIdOk() (*string, bool)`

GetActorTaskIdOk returns a tuple with the ActorTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTaskId

`func (o *Run) SetActorTaskId(v string)`

SetActorTaskId sets ActorTaskId field to given value.

### HasActorTaskId

`func (o *Run) HasActorTaskId() bool`

HasActorTaskId returns a boolean if a field has been set.

### SetActorTaskIdNil

`func (o *Run) SetActorTaskIdNil(b bool)`

 SetActorTaskIdNil sets the value for ActorTaskId to be an explicit nil

### UnsetActorTaskId
`func (o *Run) UnsetActorTaskId()`

UnsetActorTaskId ensures that no value is present for ActorTaskId, not even an explicit nil
### GetStartedAt

`func (o *Run) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Run) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Run) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *Run) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Run) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Run) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Run) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *Run) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *Run) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetStatus

`func (o *Run) GetStatus() ActorJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Run) GetStatusOk() (*ActorJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Run) SetStatus(v ActorJobStatus)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *Run) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Run) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Run) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Run) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *Run) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *Run) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetIsStatusMessageTerminal

`func (o *Run) GetIsStatusMessageTerminal() bool`

GetIsStatusMessageTerminal returns the IsStatusMessageTerminal field if non-nil, zero value otherwise.

### GetIsStatusMessageTerminalOk

`func (o *Run) GetIsStatusMessageTerminalOk() (*bool, bool)`

GetIsStatusMessageTerminalOk returns a tuple with the IsStatusMessageTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatusMessageTerminal

`func (o *Run) SetIsStatusMessageTerminal(v bool)`

SetIsStatusMessageTerminal sets IsStatusMessageTerminal field to given value.

### HasIsStatusMessageTerminal

`func (o *Run) HasIsStatusMessageTerminal() bool`

HasIsStatusMessageTerminal returns a boolean if a field has been set.

### SetIsStatusMessageTerminalNil

`func (o *Run) SetIsStatusMessageTerminalNil(b bool)`

 SetIsStatusMessageTerminalNil sets the value for IsStatusMessageTerminal to be an explicit nil

### UnsetIsStatusMessageTerminal
`func (o *Run) UnsetIsStatusMessageTerminal()`

UnsetIsStatusMessageTerminal ensures that no value is present for IsStatusMessageTerminal, not even an explicit nil
### GetMeta

`func (o *Run) GetMeta() RunMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Run) GetMetaOk() (*RunMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Run) SetMeta(v RunMeta)`

SetMeta sets Meta field to given value.


### GetPricingInfo

`func (o *Run) GetPricingInfo() ActorRunPricingInfo`

GetPricingInfo returns the PricingInfo field if non-nil, zero value otherwise.

### GetPricingInfoOk

`func (o *Run) GetPricingInfoOk() (*ActorRunPricingInfo, bool)`

GetPricingInfoOk returns a tuple with the PricingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingInfo

`func (o *Run) SetPricingInfo(v ActorRunPricingInfo)`

SetPricingInfo sets PricingInfo field to given value.

### HasPricingInfo

`func (o *Run) HasPricingInfo() bool`

HasPricingInfo returns a boolean if a field has been set.

### GetStats

`func (o *Run) GetStats() RunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Run) GetStatsOk() (*RunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Run) SetStats(v RunStats)`

SetStats sets Stats field to given value.


### GetChargedEventCounts

`func (o *Run) GetChargedEventCounts() map[string]int32`

GetChargedEventCounts returns the ChargedEventCounts field if non-nil, zero value otherwise.

### GetChargedEventCountsOk

`func (o *Run) GetChargedEventCountsOk() (*map[string]int32, bool)`

GetChargedEventCountsOk returns a tuple with the ChargedEventCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedEventCounts

`func (o *Run) SetChargedEventCounts(v map[string]int32)`

SetChargedEventCounts sets ChargedEventCounts field to given value.

### HasChargedEventCounts

`func (o *Run) HasChargedEventCounts() bool`

HasChargedEventCounts returns a boolean if a field has been set.

### GetOptions

`func (o *Run) GetOptions() RunOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Run) GetOptionsOk() (*RunOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Run) SetOptions(v RunOptions)`

SetOptions sets Options field to given value.


### GetBuildId

`func (o *Run) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *Run) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *Run) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.


### GetExitCode

`func (o *Run) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Run) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Run) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Run) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### SetExitCodeNil

`func (o *Run) SetExitCodeNil(b bool)`

 SetExitCodeNil sets the value for ExitCode to be an explicit nil

### UnsetExitCode
`func (o *Run) UnsetExitCode()`

UnsetExitCode ensures that no value is present for ExitCode, not even an explicit nil
### GetGeneralAccess

`func (o *Run) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *Run) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *Run) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.


### GetDefaultKeyValueStoreId

`func (o *Run) GetDefaultKeyValueStoreId() string`

GetDefaultKeyValueStoreId returns the DefaultKeyValueStoreId field if non-nil, zero value otherwise.

### GetDefaultKeyValueStoreIdOk

`func (o *Run) GetDefaultKeyValueStoreIdOk() (*string, bool)`

GetDefaultKeyValueStoreIdOk returns a tuple with the DefaultKeyValueStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyValueStoreId

`func (o *Run) SetDefaultKeyValueStoreId(v string)`

SetDefaultKeyValueStoreId sets DefaultKeyValueStoreId field to given value.


### GetDefaultDatasetId

`func (o *Run) GetDefaultDatasetId() string`

GetDefaultDatasetId returns the DefaultDatasetId field if non-nil, zero value otherwise.

### GetDefaultDatasetIdOk

`func (o *Run) GetDefaultDatasetIdOk() (*string, bool)`

GetDefaultDatasetIdOk returns a tuple with the DefaultDatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatasetId

`func (o *Run) SetDefaultDatasetId(v string)`

SetDefaultDatasetId sets DefaultDatasetId field to given value.


### GetDefaultRequestQueueId

`func (o *Run) GetDefaultRequestQueueId() string`

GetDefaultRequestQueueId returns the DefaultRequestQueueId field if non-nil, zero value otherwise.

### GetDefaultRequestQueueIdOk

`func (o *Run) GetDefaultRequestQueueIdOk() (*string, bool)`

GetDefaultRequestQueueIdOk returns a tuple with the DefaultRequestQueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRequestQueueId

`func (o *Run) SetDefaultRequestQueueId(v string)`

SetDefaultRequestQueueId sets DefaultRequestQueueId field to given value.


### GetStorageIds

`func (o *Run) GetStorageIds() RunStorageIds`

GetStorageIds returns the StorageIds field if non-nil, zero value otherwise.

### GetStorageIdsOk

`func (o *Run) GetStorageIdsOk() (*RunStorageIds, bool)`

GetStorageIdsOk returns a tuple with the StorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIds

`func (o *Run) SetStorageIds(v RunStorageIds)`

SetStorageIds sets StorageIds field to given value.

### HasStorageIds

`func (o *Run) HasStorageIds() bool`

HasStorageIds returns a boolean if a field has been set.

### GetBuildNumber

`func (o *Run) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *Run) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *Run) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *Run) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### SetBuildNumberNil

`func (o *Run) SetBuildNumberNil(b bool)`

 SetBuildNumberNil sets the value for BuildNumber to be an explicit nil

### UnsetBuildNumber
`func (o *Run) UnsetBuildNumber()`

UnsetBuildNumber ensures that no value is present for BuildNumber, not even an explicit nil
### GetContainerUrl

`func (o *Run) GetContainerUrl() string`

GetContainerUrl returns the ContainerUrl field if non-nil, zero value otherwise.

### GetContainerUrlOk

`func (o *Run) GetContainerUrlOk() (*string, bool)`

GetContainerUrlOk returns a tuple with the ContainerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerUrl

`func (o *Run) SetContainerUrl(v string)`

SetContainerUrl sets ContainerUrl field to given value.

### HasContainerUrl

`func (o *Run) HasContainerUrl() bool`

HasContainerUrl returns a boolean if a field has been set.

### GetIsContainerServerReady

`func (o *Run) GetIsContainerServerReady() bool`

GetIsContainerServerReady returns the IsContainerServerReady field if non-nil, zero value otherwise.

### GetIsContainerServerReadyOk

`func (o *Run) GetIsContainerServerReadyOk() (*bool, bool)`

GetIsContainerServerReadyOk returns a tuple with the IsContainerServerReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainerServerReady

`func (o *Run) SetIsContainerServerReady(v bool)`

SetIsContainerServerReady sets IsContainerServerReady field to given value.

### HasIsContainerServerReady

`func (o *Run) HasIsContainerServerReady() bool`

HasIsContainerServerReady returns a boolean if a field has been set.

### SetIsContainerServerReadyNil

`func (o *Run) SetIsContainerServerReadyNil(b bool)`

 SetIsContainerServerReadyNil sets the value for IsContainerServerReady to be an explicit nil

### UnsetIsContainerServerReady
`func (o *Run) UnsetIsContainerServerReady()`

UnsetIsContainerServerReady ensures that no value is present for IsContainerServerReady, not even an explicit nil
### GetGitBranchName

`func (o *Run) GetGitBranchName() string`

GetGitBranchName returns the GitBranchName field if non-nil, zero value otherwise.

### GetGitBranchNameOk

`func (o *Run) GetGitBranchNameOk() (*string, bool)`

GetGitBranchNameOk returns a tuple with the GitBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranchName

`func (o *Run) SetGitBranchName(v string)`

SetGitBranchName sets GitBranchName field to given value.

### HasGitBranchName

`func (o *Run) HasGitBranchName() bool`

HasGitBranchName returns a boolean if a field has been set.

### SetGitBranchNameNil

`func (o *Run) SetGitBranchNameNil(b bool)`

 SetGitBranchNameNil sets the value for GitBranchName to be an explicit nil

### UnsetGitBranchName
`func (o *Run) UnsetGitBranchName()`

UnsetGitBranchName ensures that no value is present for GitBranchName, not even an explicit nil
### GetUsage

`func (o *Run) GetUsage() RunUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Run) GetUsageOk() (*RunUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Run) SetUsage(v RunUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Run) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *Run) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *Run) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetUsageTotalUsd

`func (o *Run) GetUsageTotalUsd() float32`

GetUsageTotalUsd returns the UsageTotalUsd field if non-nil, zero value otherwise.

### GetUsageTotalUsdOk

`func (o *Run) GetUsageTotalUsdOk() (*float32, bool)`

GetUsageTotalUsdOk returns a tuple with the UsageTotalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTotalUsd

`func (o *Run) SetUsageTotalUsd(v float32)`

SetUsageTotalUsd sets UsageTotalUsd field to given value.

### HasUsageTotalUsd

`func (o *Run) HasUsageTotalUsd() bool`

HasUsageTotalUsd returns a boolean if a field has been set.

### SetUsageTotalUsdNil

`func (o *Run) SetUsageTotalUsdNil(b bool)`

 SetUsageTotalUsdNil sets the value for UsageTotalUsd to be an explicit nil

### UnsetUsageTotalUsd
`func (o *Run) UnsetUsageTotalUsd()`

UnsetUsageTotalUsd ensures that no value is present for UsageTotalUsd, not even an explicit nil
### GetUsageUsd

`func (o *Run) GetUsageUsd() RunUsageUsd`

GetUsageUsd returns the UsageUsd field if non-nil, zero value otherwise.

### GetUsageUsdOk

`func (o *Run) GetUsageUsdOk() (*RunUsageUsd, bool)`

GetUsageUsdOk returns a tuple with the UsageUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUsd

`func (o *Run) SetUsageUsd(v RunUsageUsd)`

SetUsageUsd sets UsageUsd field to given value.

### HasUsageUsd

`func (o *Run) HasUsageUsd() bool`

HasUsageUsd returns a boolean if a field has been set.

### SetUsageUsdNil

`func (o *Run) SetUsageUsdNil(b bool)`

 SetUsageUsdNil sets the value for UsageUsd to be an explicit nil

### UnsetUsageUsd
`func (o *Run) UnsetUsageUsd()`

UnsetUsageUsd ensures that no value is present for UsageUsd, not even an explicit nil
### GetMetamorphs

`func (o *Run) GetMetamorphs() []Metamorph`

GetMetamorphs returns the Metamorphs field if non-nil, zero value otherwise.

### GetMetamorphsOk

`func (o *Run) GetMetamorphsOk() (*[]Metamorph, bool)`

GetMetamorphsOk returns a tuple with the Metamorphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetamorphs

`func (o *Run) SetMetamorphs(v []Metamorph)`

SetMetamorphs sets Metamorphs field to given value.

### HasMetamorphs

`func (o *Run) HasMetamorphs() bool`

HasMetamorphs returns a boolean if a field has been set.

### SetMetamorphsNil

`func (o *Run) SetMetamorphsNil(b bool)`

 SetMetamorphsNil sets the value for Metamorphs to be an explicit nil

### UnsetMetamorphs
`func (o *Run) UnsetMetamorphs()`

UnsetMetamorphs ensures that no value is present for Metamorphs, not even an explicit nil
### GetPlatformUsageBillingModel

`func (o *Run) GetPlatformUsageBillingModel() string`

GetPlatformUsageBillingModel returns the PlatformUsageBillingModel field if non-nil, zero value otherwise.

### GetPlatformUsageBillingModelOk

`func (o *Run) GetPlatformUsageBillingModelOk() (*string, bool)`

GetPlatformUsageBillingModelOk returns a tuple with the PlatformUsageBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUsageBillingModel

`func (o *Run) SetPlatformUsageBillingModel(v string)`

SetPlatformUsageBillingModel sets PlatformUsageBillingModel field to given value.

### HasPlatformUsageBillingModel

`func (o *Run) HasPlatformUsageBillingModel() bool`

HasPlatformUsageBillingModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


