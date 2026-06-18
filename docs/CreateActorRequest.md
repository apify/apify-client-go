# CreateActorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **NullableBool** |  | [optional] 
**SeoTitle** | Pointer to **NullableString** |  | [optional] 
**SeoDescription** | Pointer to **NullableString** |  | [optional] 
**RestartOnError** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**[]Version**](Version.md) |  | [optional] 
**PricingInfos** | Pointer to [**[]ActorRunPricingInfo**](ActorRunPricingInfo.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**DefaultRunOptions** | Pointer to [**DefaultRunOptions**](DefaultRunOptions.md) |  | [optional] 
**ActorStandby** | Pointer to [**NullableActorStandby**](ActorStandby.md) |  | [optional] 
**ExampleRunInput** | Pointer to [**NullableExampleRunInput**](ExampleRunInput.md) |  | [optional] 
**IsDeprecated** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateActorRequest

`func NewCreateActorRequest() *CreateActorRequest`

NewCreateActorRequest instantiates a new CreateActorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActorRequestWithDefaults

`func NewCreateActorRequestWithDefaults() *CreateActorRequest`

NewCreateActorRequestWithDefaults instantiates a new CreateActorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateActorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateActorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateActorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateActorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateActorRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateActorRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateActorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateActorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateActorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateActorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateActorRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateActorRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTitle

`func (o *CreateActorRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateActorRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateActorRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateActorRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateActorRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateActorRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetIsPublic

`func (o *CreateActorRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CreateActorRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CreateActorRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CreateActorRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### SetIsPublicNil

`func (o *CreateActorRequest) SetIsPublicNil(b bool)`

 SetIsPublicNil sets the value for IsPublic to be an explicit nil

### UnsetIsPublic
`func (o *CreateActorRequest) UnsetIsPublic()`

UnsetIsPublic ensures that no value is present for IsPublic, not even an explicit nil
### GetSeoTitle

`func (o *CreateActorRequest) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *CreateActorRequest) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *CreateActorRequest) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *CreateActorRequest) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *CreateActorRequest) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *CreateActorRequest) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoDescription

`func (o *CreateActorRequest) GetSeoDescription() string`

GetSeoDescription returns the SeoDescription field if non-nil, zero value otherwise.

### GetSeoDescriptionOk

`func (o *CreateActorRequest) GetSeoDescriptionOk() (*string, bool)`

GetSeoDescriptionOk returns a tuple with the SeoDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoDescription

`func (o *CreateActorRequest) SetSeoDescription(v string)`

SetSeoDescription sets SeoDescription field to given value.

### HasSeoDescription

`func (o *CreateActorRequest) HasSeoDescription() bool`

HasSeoDescription returns a boolean if a field has been set.

### SetSeoDescriptionNil

`func (o *CreateActorRequest) SetSeoDescriptionNil(b bool)`

 SetSeoDescriptionNil sets the value for SeoDescription to be an explicit nil

### UnsetSeoDescription
`func (o *CreateActorRequest) UnsetSeoDescription()`

UnsetSeoDescription ensures that no value is present for SeoDescription, not even an explicit nil
### GetRestartOnError

`func (o *CreateActorRequest) GetRestartOnError() bool`

GetRestartOnError returns the RestartOnError field if non-nil, zero value otherwise.

### GetRestartOnErrorOk

`func (o *CreateActorRequest) GetRestartOnErrorOk() (*bool, bool)`

GetRestartOnErrorOk returns a tuple with the RestartOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnError

`func (o *CreateActorRequest) SetRestartOnError(v bool)`

SetRestartOnError sets RestartOnError field to given value.

### HasRestartOnError

`func (o *CreateActorRequest) HasRestartOnError() bool`

HasRestartOnError returns a boolean if a field has been set.

### GetVersions

`func (o *CreateActorRequest) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CreateActorRequest) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CreateActorRequest) SetVersions(v []Version)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CreateActorRequest) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *CreateActorRequest) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *CreateActorRequest) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil
### GetPricingInfos

`func (o *CreateActorRequest) GetPricingInfos() []ActorRunPricingInfo`

GetPricingInfos returns the PricingInfos field if non-nil, zero value otherwise.

### GetPricingInfosOk

`func (o *CreateActorRequest) GetPricingInfosOk() (*[]ActorRunPricingInfo, bool)`

GetPricingInfosOk returns a tuple with the PricingInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingInfos

`func (o *CreateActorRequest) SetPricingInfos(v []ActorRunPricingInfo)`

SetPricingInfos sets PricingInfos field to given value.

### HasPricingInfos

`func (o *CreateActorRequest) HasPricingInfos() bool`

HasPricingInfos returns a boolean if a field has been set.

### GetCategories

`func (o *CreateActorRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CreateActorRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CreateActorRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CreateActorRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *CreateActorRequest) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *CreateActorRequest) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetDefaultRunOptions

`func (o *CreateActorRequest) GetDefaultRunOptions() DefaultRunOptions`

GetDefaultRunOptions returns the DefaultRunOptions field if non-nil, zero value otherwise.

### GetDefaultRunOptionsOk

`func (o *CreateActorRequest) GetDefaultRunOptionsOk() (*DefaultRunOptions, bool)`

GetDefaultRunOptionsOk returns a tuple with the DefaultRunOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRunOptions

`func (o *CreateActorRequest) SetDefaultRunOptions(v DefaultRunOptions)`

SetDefaultRunOptions sets DefaultRunOptions field to given value.

### HasDefaultRunOptions

`func (o *CreateActorRequest) HasDefaultRunOptions() bool`

HasDefaultRunOptions returns a boolean if a field has been set.

### GetActorStandby

`func (o *CreateActorRequest) GetActorStandby() ActorStandby`

GetActorStandby returns the ActorStandby field if non-nil, zero value otherwise.

### GetActorStandbyOk

`func (o *CreateActorRequest) GetActorStandbyOk() (*ActorStandby, bool)`

GetActorStandbyOk returns a tuple with the ActorStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorStandby

`func (o *CreateActorRequest) SetActorStandby(v ActorStandby)`

SetActorStandby sets ActorStandby field to given value.

### HasActorStandby

`func (o *CreateActorRequest) HasActorStandby() bool`

HasActorStandby returns a boolean if a field has been set.

### SetActorStandbyNil

`func (o *CreateActorRequest) SetActorStandbyNil(b bool)`

 SetActorStandbyNil sets the value for ActorStandby to be an explicit nil

### UnsetActorStandby
`func (o *CreateActorRequest) UnsetActorStandby()`

UnsetActorStandby ensures that no value is present for ActorStandby, not even an explicit nil
### GetExampleRunInput

`func (o *CreateActorRequest) GetExampleRunInput() ExampleRunInput`

GetExampleRunInput returns the ExampleRunInput field if non-nil, zero value otherwise.

### GetExampleRunInputOk

`func (o *CreateActorRequest) GetExampleRunInputOk() (*ExampleRunInput, bool)`

GetExampleRunInputOk returns a tuple with the ExampleRunInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleRunInput

`func (o *CreateActorRequest) SetExampleRunInput(v ExampleRunInput)`

SetExampleRunInput sets ExampleRunInput field to given value.

### HasExampleRunInput

`func (o *CreateActorRequest) HasExampleRunInput() bool`

HasExampleRunInput returns a boolean if a field has been set.

### SetExampleRunInputNil

`func (o *CreateActorRequest) SetExampleRunInputNil(b bool)`

 SetExampleRunInputNil sets the value for ExampleRunInput to be an explicit nil

### UnsetExampleRunInput
`func (o *CreateActorRequest) UnsetExampleRunInput()`

UnsetExampleRunInput ensures that no value is present for ExampleRunInput, not even an explicit nil
### GetIsDeprecated

`func (o *CreateActorRequest) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *CreateActorRequest) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *CreateActorRequest) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *CreateActorRequest) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### SetIsDeprecatedNil

`func (o *CreateActorRequest) SetIsDeprecatedNil(b bool)`

 SetIsDeprecatedNil sets the value for IsDeprecated to be an explicit nil

### UnsetIsDeprecated
`func (o *CreateActorRequest) UnsetIsDeprecated()`

UnsetIsDeprecated ensures that no value is present for IsDeprecated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


