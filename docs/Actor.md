# Actor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | **string** |  | 
**Username** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**RestartOnError** | Pointer to **bool** |  | [optional] 
**IsPublic** | **bool** |  | 
**ActorPermissionLevel** | Pointer to [**ActorPermissionLevel**](ActorPermissionLevel.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Stats** | [**ActorStats**](ActorStats.md) |  | 
**Versions** | [**[]Version**](Version.md) |  | 
**PricingInfos** | Pointer to [**[]ActorRunPricingInfo**](ActorRunPricingInfo.md) |  | [optional] 
**DefaultRunOptions** | [**DefaultRunOptions**](DefaultRunOptions.md) |  | 
**ExampleRunInput** | Pointer to [**NullableExampleRunInput**](ExampleRunInput.md) |  | [optional] 
**IsDeprecated** | Pointer to **NullableBool** |  | [optional] 
**DeploymentKey** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**TaggedBuilds** | Pointer to [**map[string]TaggedBuildsValue**](TaggedBuildsValue.md) | A dictionary mapping build tag names (e.g., \&quot;latest\&quot;, \&quot;beta\&quot;) to their build information. | [optional] 
**ActorStandby** | Pointer to [**NullableActorStandby**](ActorStandby.md) |  | [optional] 
**ReadmeSummary** | Pointer to **string** | A brief, LLM-generated readme summary | [optional] 
**SeoTitle** | Pointer to **NullableString** |  | [optional] 
**SeoDescription** | Pointer to **NullableString** |  | [optional] 
**PictureUrl** | Pointer to **NullableString** |  | [optional] 
**StandbyUrl** | Pointer to **NullableString** |  | [optional] 
**Notice** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**IsCritical** | Pointer to **bool** |  | [optional] 
**IsGeneric** | Pointer to **bool** |  | [optional] 
**IsSourceCodeHidden** | Pointer to **bool** |  | [optional] 
**HasNoDataset** | Pointer to **bool** |  | [optional] 

## Methods

### NewActor

`func NewActor(id string, userId string, name string, username string, isPublic bool, createdAt time.Time, modifiedAt time.Time, stats ActorStats, versions []Version, defaultRunOptions DefaultRunOptions, ) *Actor`

NewActor instantiates a new Actor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorWithDefaults

`func NewActorWithDefaults() *Actor`

NewActorWithDefaults instantiates a new Actor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Actor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Actor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Actor) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *Actor) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Actor) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Actor) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *Actor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Actor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Actor) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *Actor) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Actor) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Actor) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDescription

`func (o *Actor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Actor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Actor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Actor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Actor) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Actor) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRestartOnError

`func (o *Actor) GetRestartOnError() bool`

GetRestartOnError returns the RestartOnError field if non-nil, zero value otherwise.

### GetRestartOnErrorOk

`func (o *Actor) GetRestartOnErrorOk() (*bool, bool)`

GetRestartOnErrorOk returns a tuple with the RestartOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnError

`func (o *Actor) SetRestartOnError(v bool)`

SetRestartOnError sets RestartOnError field to given value.

### HasRestartOnError

`func (o *Actor) HasRestartOnError() bool`

HasRestartOnError returns a boolean if a field has been set.

### GetIsPublic

`func (o *Actor) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Actor) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Actor) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetActorPermissionLevel

`func (o *Actor) GetActorPermissionLevel() ActorPermissionLevel`

GetActorPermissionLevel returns the ActorPermissionLevel field if non-nil, zero value otherwise.

### GetActorPermissionLevelOk

`func (o *Actor) GetActorPermissionLevelOk() (*ActorPermissionLevel, bool)`

GetActorPermissionLevelOk returns a tuple with the ActorPermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorPermissionLevel

`func (o *Actor) SetActorPermissionLevel(v ActorPermissionLevel)`

SetActorPermissionLevel sets ActorPermissionLevel field to given value.

### HasActorPermissionLevel

`func (o *Actor) HasActorPermissionLevel() bool`

HasActorPermissionLevel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Actor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Actor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Actor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Actor) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Actor) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Actor) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetStats

`func (o *Actor) GetStats() ActorStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Actor) GetStatsOk() (*ActorStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Actor) SetStats(v ActorStats)`

SetStats sets Stats field to given value.


### GetVersions

`func (o *Actor) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Actor) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Actor) SetVersions(v []Version)`

SetVersions sets Versions field to given value.


### GetPricingInfos

`func (o *Actor) GetPricingInfos() []ActorRunPricingInfo`

GetPricingInfos returns the PricingInfos field if non-nil, zero value otherwise.

### GetPricingInfosOk

`func (o *Actor) GetPricingInfosOk() (*[]ActorRunPricingInfo, bool)`

GetPricingInfosOk returns a tuple with the PricingInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingInfos

`func (o *Actor) SetPricingInfos(v []ActorRunPricingInfo)`

SetPricingInfos sets PricingInfos field to given value.

### HasPricingInfos

`func (o *Actor) HasPricingInfos() bool`

HasPricingInfos returns a boolean if a field has been set.

### GetDefaultRunOptions

`func (o *Actor) GetDefaultRunOptions() DefaultRunOptions`

GetDefaultRunOptions returns the DefaultRunOptions field if non-nil, zero value otherwise.

### GetDefaultRunOptionsOk

`func (o *Actor) GetDefaultRunOptionsOk() (*DefaultRunOptions, bool)`

GetDefaultRunOptionsOk returns a tuple with the DefaultRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRunOptions

`func (o *Actor) SetDefaultRunOptions(v DefaultRunOptions)`

SetDefaultRunOptions sets DefaultRunOptions field to given value.


### GetExampleRunInput

`func (o *Actor) GetExampleRunInput() ExampleRunInput`

GetExampleRunInput returns the ExampleRunInput field if non-nil, zero value otherwise.

### GetExampleRunInputOk

`func (o *Actor) GetExampleRunInputOk() (*ExampleRunInput, bool)`

GetExampleRunInputOk returns a tuple with the ExampleRunInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleRunInput

`func (o *Actor) SetExampleRunInput(v ExampleRunInput)`

SetExampleRunInput sets ExampleRunInput field to given value.

### HasExampleRunInput

`func (o *Actor) HasExampleRunInput() bool`

HasExampleRunInput returns a boolean if a field has been set.

### SetExampleRunInputNil

`func (o *Actor) SetExampleRunInputNil(b bool)`

 SetExampleRunInputNil sets the value for ExampleRunInput to be an explicit nil

### UnsetExampleRunInput
`func (o *Actor) UnsetExampleRunInput()`

UnsetExampleRunInput ensures that no value is present for ExampleRunInput, not even an explicit nil
### GetIsDeprecated

`func (o *Actor) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *Actor) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *Actor) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *Actor) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### SetIsDeprecatedNil

`func (o *Actor) SetIsDeprecatedNil(b bool)`

 SetIsDeprecatedNil sets the value for IsDeprecated to be an explicit nil

### UnsetIsDeprecated
`func (o *Actor) UnsetIsDeprecated()`

UnsetIsDeprecated ensures that no value is present for IsDeprecated, not even an explicit nil
### GetDeploymentKey

`func (o *Actor) GetDeploymentKey() string`

GetDeploymentKey returns the DeploymentKey field if non-nil, zero value otherwise.

### GetDeploymentKeyOk

`func (o *Actor) GetDeploymentKeyOk() (*string, bool)`

GetDeploymentKeyOk returns a tuple with the DeploymentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentKey

`func (o *Actor) SetDeploymentKey(v string)`

SetDeploymentKey sets DeploymentKey field to given value.

### HasDeploymentKey

`func (o *Actor) HasDeploymentKey() bool`

HasDeploymentKey returns a boolean if a field has been set.

### GetTitle

`func (o *Actor) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Actor) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Actor) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Actor) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Actor) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Actor) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTaggedBuilds

`func (o *Actor) GetTaggedBuilds() map[string]TaggedBuildsValue`

GetTaggedBuilds returns the TaggedBuilds field if non-nil, zero value otherwise.

### GetTaggedBuildsOk

`func (o *Actor) GetTaggedBuildsOk() (*map[string]TaggedBuildsValue, bool)`

GetTaggedBuildsOk returns a tuple with the TaggedBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedBuilds

`func (o *Actor) SetTaggedBuilds(v map[string]TaggedBuildsValue)`

SetTaggedBuilds sets TaggedBuilds field to given value.

### HasTaggedBuilds

`func (o *Actor) HasTaggedBuilds() bool`

HasTaggedBuilds returns a boolean if a field has been set.

### SetTaggedBuildsNil

`func (o *Actor) SetTaggedBuildsNil(b bool)`

 SetTaggedBuildsNil sets the value for TaggedBuilds to be an explicit nil

### UnsetTaggedBuilds
`func (o *Actor) UnsetTaggedBuilds()`

UnsetTaggedBuilds ensures that no value is present for TaggedBuilds, not even an explicit nil
### GetActorStandby

`func (o *Actor) GetActorStandby() ActorStandby`

GetActorStandby returns the ActorStandby field if non-nil, zero value otherwise.

### GetActorStandbyOk

`func (o *Actor) GetActorStandbyOk() (*ActorStandby, bool)`

GetActorStandbyOk returns a tuple with the ActorStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorStandby

`func (o *Actor) SetActorStandby(v ActorStandby)`

SetActorStandby sets ActorStandby field to given value.

### HasActorStandby

`func (o *Actor) HasActorStandby() bool`

HasActorStandby returns a boolean if a field has been set.

### SetActorStandbyNil

`func (o *Actor) SetActorStandbyNil(b bool)`

 SetActorStandbyNil sets the value for ActorStandby to be an explicit nil

### UnsetActorStandby
`func (o *Actor) UnsetActorStandby()`

UnsetActorStandby ensures that no value is present for ActorStandby, not even an explicit nil
### GetReadmeSummary

`func (o *Actor) GetReadmeSummary() string`

GetReadmeSummary returns the ReadmeSummary field if non-nil, zero value otherwise.

### GetReadmeSummaryOk

`func (o *Actor) GetReadmeSummaryOk() (*string, bool)`

GetReadmeSummaryOk returns a tuple with the ReadmeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadmeSummary

`func (o *Actor) SetReadmeSummary(v string)`

SetReadmeSummary sets ReadmeSummary field to given value.

### HasReadmeSummary

`func (o *Actor) HasReadmeSummary() bool`

HasReadmeSummary returns a boolean if a field has been set.

### GetSeoTitle

`func (o *Actor) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *Actor) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *Actor) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *Actor) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *Actor) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *Actor) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoDescription

`func (o *Actor) GetSeoDescription() string`

GetSeoDescription returns the SeoDescription field if non-nil, zero value otherwise.

### GetSeoDescriptionOk

`func (o *Actor) GetSeoDescriptionOk() (*string, bool)`

GetSeoDescriptionOk returns a tuple with the SeoDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoDescription

`func (o *Actor) SetSeoDescription(v string)`

SetSeoDescription sets SeoDescription field to given value.

### HasSeoDescription

`func (o *Actor) HasSeoDescription() bool`

HasSeoDescription returns a boolean if a field has been set.

### SetSeoDescriptionNil

`func (o *Actor) SetSeoDescriptionNil(b bool)`

 SetSeoDescriptionNil sets the value for SeoDescription to be an explicit nil

### UnsetSeoDescription
`func (o *Actor) UnsetSeoDescription()`

UnsetSeoDescription ensures that no value is present for SeoDescription, not even an explicit nil
### GetPictureUrl

`func (o *Actor) GetPictureUrl() string`

GetPictureUrl returns the PictureUrl field if non-nil, zero value otherwise.

### GetPictureUrlOk

`func (o *Actor) GetPictureUrlOk() (*string, bool)`

GetPictureUrlOk returns a tuple with the PictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUrl

`func (o *Actor) SetPictureUrl(v string)`

SetPictureUrl sets PictureUrl field to given value.

### HasPictureUrl

`func (o *Actor) HasPictureUrl() bool`

HasPictureUrl returns a boolean if a field has been set.

### SetPictureUrlNil

`func (o *Actor) SetPictureUrlNil(b bool)`

 SetPictureUrlNil sets the value for PictureUrl to be an explicit nil

### UnsetPictureUrl
`func (o *Actor) UnsetPictureUrl()`

UnsetPictureUrl ensures that no value is present for PictureUrl, not even an explicit nil
### GetStandbyUrl

`func (o *Actor) GetStandbyUrl() string`

GetStandbyUrl returns the StandbyUrl field if non-nil, zero value otherwise.

### GetStandbyUrlOk

`func (o *Actor) GetStandbyUrlOk() (*string, bool)`

GetStandbyUrlOk returns a tuple with the StandbyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyUrl

`func (o *Actor) SetStandbyUrl(v string)`

SetStandbyUrl sets StandbyUrl field to given value.

### HasStandbyUrl

`func (o *Actor) HasStandbyUrl() bool`

HasStandbyUrl returns a boolean if a field has been set.

### SetStandbyUrlNil

`func (o *Actor) SetStandbyUrlNil(b bool)`

 SetStandbyUrlNil sets the value for StandbyUrl to be an explicit nil

### UnsetStandbyUrl
`func (o *Actor) UnsetStandbyUrl()`

UnsetStandbyUrl ensures that no value is present for StandbyUrl, not even an explicit nil
### GetNotice

`func (o *Actor) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *Actor) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *Actor) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *Actor) HasNotice() bool`

HasNotice returns a boolean if a field has been set.

### GetCategories

`func (o *Actor) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Actor) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Actor) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Actor) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetIsCritical

`func (o *Actor) GetIsCritical() bool`

GetIsCritical returns the IsCritical field if non-nil, zero value otherwise.

### GetIsCriticalOk

`func (o *Actor) GetIsCriticalOk() (*bool, bool)`

GetIsCriticalOk returns a tuple with the IsCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCritical

`func (o *Actor) SetIsCritical(v bool)`

SetIsCritical sets IsCritical field to given value.

### HasIsCritical

`func (o *Actor) HasIsCritical() bool`

HasIsCritical returns a boolean if a field has been set.

### GetIsGeneric

`func (o *Actor) GetIsGeneric() bool`

GetIsGeneric returns the IsGeneric field if non-nil, zero value otherwise.

### GetIsGenericOk

`func (o *Actor) GetIsGenericOk() (*bool, bool)`

GetIsGenericOk returns a tuple with the IsGeneric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGeneric

`func (o *Actor) SetIsGeneric(v bool)`

SetIsGeneric sets IsGeneric field to given value.

### HasIsGeneric

`func (o *Actor) HasIsGeneric() bool`

HasIsGeneric returns a boolean if a field has been set.

### GetIsSourceCodeHidden

`func (o *Actor) GetIsSourceCodeHidden() bool`

GetIsSourceCodeHidden returns the IsSourceCodeHidden field if non-nil, zero value otherwise.

### GetIsSourceCodeHiddenOk

`func (o *Actor) GetIsSourceCodeHiddenOk() (*bool, bool)`

GetIsSourceCodeHiddenOk returns a tuple with the IsSourceCodeHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceCodeHidden

`func (o *Actor) SetIsSourceCodeHidden(v bool)`

SetIsSourceCodeHidden sets IsSourceCodeHidden field to given value.

### HasIsSourceCodeHidden

`func (o *Actor) HasIsSourceCodeHidden() bool`

HasIsSourceCodeHidden returns a boolean if a field has been set.

### GetHasNoDataset

`func (o *Actor) GetHasNoDataset() bool`

GetHasNoDataset returns the HasNoDataset field if non-nil, zero value otherwise.

### GetHasNoDatasetOk

`func (o *Actor) GetHasNoDatasetOk() (*bool, bool)`

GetHasNoDatasetOk returns a tuple with the HasNoDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNoDataset

`func (o *Actor) SetHasNoDataset(v bool)`

SetHasNoDataset sets HasNoDataset field to given value.

### HasHasNoDataset

`func (o *Actor) HasHasNoDataset() bool`

HasHasNoDataset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


