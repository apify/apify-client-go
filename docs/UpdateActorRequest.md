# UpdateActorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**ActorPermissionLevel** | Pointer to [**NullableActorPermissionLevel**](ActorPermissionLevel.md) |  | [optional] 
**SeoTitle** | Pointer to **NullableString** |  | [optional] 
**SeoDescription** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**RestartOnError** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**[]CreateOrUpdateVersionRequest**](CreateOrUpdateVersionRequest.md) |  | [optional] 
**PricingInfos** | Pointer to [**[]ActorRunPricingInfo**](ActorRunPricingInfo.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**DefaultRunOptions** | Pointer to [**NullableDefaultRunOptions**](DefaultRunOptions.md) |  | [optional] 
**TaggedBuilds** | Pointer to [**map[string]BuildTag**](BuildTag.md) | An object to modify tags on the Actor&#39;s builds. The key is the tag name (e.g., _latest_), and the value is either an object with a &#x60;buildId&#x60; or &#x60;null&#x60;.  This operation is a patch; any existing tags that you omit from this object will be preserved.  - **To create or reassign a tag**, provide the tag name with a &#x60;buildId&#x60;. e.g., to assign the _latest_ tag:    &amp;nbsp;    &#x60;&#x60;&#x60;json   {     \&quot;latest\&quot;: {       \&quot;buildId\&quot;: \&quot;z2EryhbfhgSyqj6Hn\&quot;     }   }   &#x60;&#x60;&#x60;  - **To remove a tag**, provide the tag name with a &#x60;null&#x60; value. e.g., to remove the _beta_ tag:    &amp;nbsp;    &#x60;&#x60;&#x60;json   {     \&quot;beta\&quot;: null   }   &#x60;&#x60;&#x60;  - **To perform multiple operations**, combine them. The following reassigns _latest_ and removes _beta_, while preserving any other existing tags.    &amp;nbsp;    &#x60;&#x60;&#x60;json   {     \&quot;latest\&quot;: {       \&quot;buildId\&quot;: \&quot;z2EryhbfhgSyqj6Hn\&quot;     },     \&quot;beta\&quot;: null   }   &#x60;&#x60;&#x60;  | [optional] 
**ActorStandby** | Pointer to [**NullableActorStandby**](ActorStandby.md) |  | [optional] 
**ExampleRunInput** | Pointer to [**NullableExampleRunInput**](ExampleRunInput.md) |  | [optional] 
**IsDeprecated** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateActorRequest

`func NewUpdateActorRequest() *UpdateActorRequest`

NewUpdateActorRequest instantiates a new UpdateActorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActorRequestWithDefaults

`func NewUpdateActorRequestWithDefaults() *UpdateActorRequest`

NewUpdateActorRequestWithDefaults instantiates a new UpdateActorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateActorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateActorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateActorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateActorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateActorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateActorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateActorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateActorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateActorRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateActorRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsPublic

`func (o *UpdateActorRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdateActorRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdateActorRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdateActorRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetActorPermissionLevel

`func (o *UpdateActorRequest) GetActorPermissionLevel() ActorPermissionLevel`

GetActorPermissionLevel returns the ActorPermissionLevel field if non-nil, zero value otherwise.

### GetActorPermissionLevelOk

`func (o *UpdateActorRequest) GetActorPermissionLevelOk() (*ActorPermissionLevel, bool)`

GetActorPermissionLevelOk returns a tuple with the ActorPermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorPermissionLevel

`func (o *UpdateActorRequest) SetActorPermissionLevel(v ActorPermissionLevel)`

SetActorPermissionLevel sets ActorPermissionLevel field to given value.

### HasActorPermissionLevel

`func (o *UpdateActorRequest) HasActorPermissionLevel() bool`

HasActorPermissionLevel returns a boolean if a field has been set.

### SetActorPermissionLevelNil

`func (o *UpdateActorRequest) SetActorPermissionLevelNil(b bool)`

 SetActorPermissionLevelNil sets the value for ActorPermissionLevel to be an explicit nil

### UnsetActorPermissionLevel
`func (o *UpdateActorRequest) UnsetActorPermissionLevel()`

UnsetActorPermissionLevel ensures that no value is present for ActorPermissionLevel, not even an explicit nil
### GetSeoTitle

`func (o *UpdateActorRequest) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *UpdateActorRequest) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *UpdateActorRequest) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *UpdateActorRequest) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *UpdateActorRequest) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *UpdateActorRequest) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoDescription

`func (o *UpdateActorRequest) GetSeoDescription() string`

GetSeoDescription returns the SeoDescription field if non-nil, zero value otherwise.

### GetSeoDescriptionOk

`func (o *UpdateActorRequest) GetSeoDescriptionOk() (*string, bool)`

GetSeoDescriptionOk returns a tuple with the SeoDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoDescription

`func (o *UpdateActorRequest) SetSeoDescription(v string)`

SetSeoDescription sets SeoDescription field to given value.

### HasSeoDescription

`func (o *UpdateActorRequest) HasSeoDescription() bool`

HasSeoDescription returns a boolean if a field has been set.

### SetSeoDescriptionNil

`func (o *UpdateActorRequest) SetSeoDescriptionNil(b bool)`

 SetSeoDescriptionNil sets the value for SeoDescription to be an explicit nil

### UnsetSeoDescription
`func (o *UpdateActorRequest) UnsetSeoDescription()`

UnsetSeoDescription ensures that no value is present for SeoDescription, not even an explicit nil
### GetTitle

`func (o *UpdateActorRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateActorRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateActorRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateActorRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateActorRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateActorRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetRestartOnError

`func (o *UpdateActorRequest) GetRestartOnError() bool`

GetRestartOnError returns the RestartOnError field if non-nil, zero value otherwise.

### GetRestartOnErrorOk

`func (o *UpdateActorRequest) GetRestartOnErrorOk() (*bool, bool)`

GetRestartOnErrorOk returns a tuple with the RestartOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnError

`func (o *UpdateActorRequest) SetRestartOnError(v bool)`

SetRestartOnError sets RestartOnError field to given value.

### HasRestartOnError

`func (o *UpdateActorRequest) HasRestartOnError() bool`

HasRestartOnError returns a boolean if a field has been set.

### GetVersions

`func (o *UpdateActorRequest) GetVersions() []CreateOrUpdateVersionRequest`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *UpdateActorRequest) GetVersionsOk() (*[]CreateOrUpdateVersionRequest, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *UpdateActorRequest) SetVersions(v []CreateOrUpdateVersionRequest)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *UpdateActorRequest) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetPricingInfos

`func (o *UpdateActorRequest) GetPricingInfos() []ActorRunPricingInfo`

GetPricingInfos returns the PricingInfos field if non-nil, zero value otherwise.

### GetPricingInfosOk

`func (o *UpdateActorRequest) GetPricingInfosOk() (*[]ActorRunPricingInfo, bool)`

GetPricingInfosOk returns a tuple with the PricingInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingInfos

`func (o *UpdateActorRequest) SetPricingInfos(v []ActorRunPricingInfo)`

SetPricingInfos sets PricingInfos field to given value.

### HasPricingInfos

`func (o *UpdateActorRequest) HasPricingInfos() bool`

HasPricingInfos returns a boolean if a field has been set.

### GetCategories

`func (o *UpdateActorRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *UpdateActorRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *UpdateActorRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *UpdateActorRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *UpdateActorRequest) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *UpdateActorRequest) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetDefaultRunOptions

`func (o *UpdateActorRequest) GetDefaultRunOptions() DefaultRunOptions`

GetDefaultRunOptions returns the DefaultRunOptions field if non-nil, zero value otherwise.

### GetDefaultRunOptionsOk

`func (o *UpdateActorRequest) GetDefaultRunOptionsOk() (*DefaultRunOptions, bool)`

GetDefaultRunOptionsOk returns a tuple with the DefaultRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRunOptions

`func (o *UpdateActorRequest) SetDefaultRunOptions(v DefaultRunOptions)`

SetDefaultRunOptions sets DefaultRunOptions field to given value.

### HasDefaultRunOptions

`func (o *UpdateActorRequest) HasDefaultRunOptions() bool`

HasDefaultRunOptions returns a boolean if a field has been set.

### SetDefaultRunOptionsNil

`func (o *UpdateActorRequest) SetDefaultRunOptionsNil(b bool)`

 SetDefaultRunOptionsNil sets the value for DefaultRunOptions to be an explicit nil

### UnsetDefaultRunOptions
`func (o *UpdateActorRequest) UnsetDefaultRunOptions()`

UnsetDefaultRunOptions ensures that no value is present for DefaultRunOptions, not even an explicit nil
### GetTaggedBuilds

`func (o *UpdateActorRequest) GetTaggedBuilds() map[string]BuildTag`

GetTaggedBuilds returns the TaggedBuilds field if non-nil, zero value otherwise.

### GetTaggedBuildsOk

`func (o *UpdateActorRequest) GetTaggedBuildsOk() (*map[string]BuildTag, bool)`

GetTaggedBuildsOk returns a tuple with the TaggedBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedBuilds

`func (o *UpdateActorRequest) SetTaggedBuilds(v map[string]BuildTag)`

SetTaggedBuilds sets TaggedBuilds field to given value.

### HasTaggedBuilds

`func (o *UpdateActorRequest) HasTaggedBuilds() bool`

HasTaggedBuilds returns a boolean if a field has been set.

### GetActorStandby

`func (o *UpdateActorRequest) GetActorStandby() ActorStandby`

GetActorStandby returns the ActorStandby field if non-nil, zero value otherwise.

### GetActorStandbyOk

`func (o *UpdateActorRequest) GetActorStandbyOk() (*ActorStandby, bool)`

GetActorStandbyOk returns a tuple with the ActorStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorStandby

`func (o *UpdateActorRequest) SetActorStandby(v ActorStandby)`

SetActorStandby sets ActorStandby field to given value.

### HasActorStandby

`func (o *UpdateActorRequest) HasActorStandby() bool`

HasActorStandby returns a boolean if a field has been set.

### SetActorStandbyNil

`func (o *UpdateActorRequest) SetActorStandbyNil(b bool)`

 SetActorStandbyNil sets the value for ActorStandby to be an explicit nil

### UnsetActorStandby
`func (o *UpdateActorRequest) UnsetActorStandby()`

UnsetActorStandby ensures that no value is present for ActorStandby, not even an explicit nil
### GetExampleRunInput

`func (o *UpdateActorRequest) GetExampleRunInput() ExampleRunInput`

GetExampleRunInput returns the ExampleRunInput field if non-nil, zero value otherwise.

### GetExampleRunInputOk

`func (o *UpdateActorRequest) GetExampleRunInputOk() (*ExampleRunInput, bool)`

GetExampleRunInputOk returns a tuple with the ExampleRunInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleRunInput

`func (o *UpdateActorRequest) SetExampleRunInput(v ExampleRunInput)`

SetExampleRunInput sets ExampleRunInput field to given value.

### HasExampleRunInput

`func (o *UpdateActorRequest) HasExampleRunInput() bool`

HasExampleRunInput returns a boolean if a field has been set.

### SetExampleRunInputNil

`func (o *UpdateActorRequest) SetExampleRunInputNil(b bool)`

 SetExampleRunInputNil sets the value for ExampleRunInput to be an explicit nil

### UnsetExampleRunInput
`func (o *UpdateActorRequest) UnsetExampleRunInput()`

UnsetExampleRunInput ensures that no value is present for ExampleRunInput, not even an explicit nil
### GetIsDeprecated

`func (o *UpdateActorRequest) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *UpdateActorRequest) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *UpdateActorRequest) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *UpdateActorRequest) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### SetIsDeprecatedNil

`func (o *UpdateActorRequest) SetIsDeprecatedNil(b bool)`

 SetIsDeprecatedNil sets the value for IsDeprecated to be an explicit nil

### UnsetIsDeprecated
`func (o *UpdateActorRequest) UnsetIsDeprecated()`

UnsetIsDeprecated ensures that no value is present for IsDeprecated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


