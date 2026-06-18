# CreateOrUpdateVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionNumber** | Pointer to **NullableString** |  | [optional] 
**SourceType** | Pointer to [**NullableVersionSourceType**](VersionSourceType.md) |  | [optional] 
**EnvVars** | Pointer to [**[]EnvVarRequest**](EnvVarRequest.md) |  | [optional] 
**ApplyEnvVarsToBuild** | Pointer to **NullableBool** |  | [optional] 
**BuildTag** | Pointer to **NullableString** |  | [optional] 
**SourceFiles** | Pointer to [**[]VersionSourceFilesInner**](VersionSourceFilesInner.md) |  | [optional] 
**GitRepoUrl** | Pointer to **NullableString** | URL of the Git repository when sourceType is GIT_REPO. | [optional] 
**TarballUrl** | Pointer to **NullableString** | URL of the tarball when sourceType is TARBALL. | [optional] 
**GitHubGistUrl** | Pointer to **NullableString** | URL of the GitHub Gist when sourceType is GITHUB_GIST. | [optional] 

## Methods

### NewCreateOrUpdateVersionRequest

`func NewCreateOrUpdateVersionRequest() *CreateOrUpdateVersionRequest`

NewCreateOrUpdateVersionRequest instantiates a new CreateOrUpdateVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateVersionRequestWithDefaults

`func NewCreateOrUpdateVersionRequestWithDefaults() *CreateOrUpdateVersionRequest`

NewCreateOrUpdateVersionRequestWithDefaults instantiates a new CreateOrUpdateVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionNumber

`func (o *CreateOrUpdateVersionRequest) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *CreateOrUpdateVersionRequest) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *CreateOrUpdateVersionRequest) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *CreateOrUpdateVersionRequest) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### SetVersionNumberNil

`func (o *CreateOrUpdateVersionRequest) SetVersionNumberNil(b bool)`

 SetVersionNumberNil sets the value for VersionNumber to be an explicit nil

### UnsetVersionNumber
`func (o *CreateOrUpdateVersionRequest) UnsetVersionNumber()`

UnsetVersionNumber ensures that no value is present for VersionNumber, not even an explicit nil
### GetSourceType

`func (o *CreateOrUpdateVersionRequest) GetSourceType() VersionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CreateOrUpdateVersionRequest) GetSourceTypeOk() (*VersionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CreateOrUpdateVersionRequest) SetSourceType(v VersionSourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *CreateOrUpdateVersionRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *CreateOrUpdateVersionRequest) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *CreateOrUpdateVersionRequest) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetEnvVars

`func (o *CreateOrUpdateVersionRequest) GetEnvVars() []EnvVarRequest`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *CreateOrUpdateVersionRequest) GetEnvVarsOk() (*[]EnvVarRequest, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *CreateOrUpdateVersionRequest) SetEnvVars(v []EnvVarRequest)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *CreateOrUpdateVersionRequest) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### SetEnvVarsNil

`func (o *CreateOrUpdateVersionRequest) SetEnvVarsNil(b bool)`

 SetEnvVarsNil sets the value for EnvVars to be an explicit nil

### UnsetEnvVars
`func (o *CreateOrUpdateVersionRequest) UnsetEnvVars()`

UnsetEnvVars ensures that no value is present for EnvVars, not even an explicit nil
### GetApplyEnvVarsToBuild

`func (o *CreateOrUpdateVersionRequest) GetApplyEnvVarsToBuild() bool`

GetApplyEnvVarsToBuild returns the ApplyEnvVarsToBuild field if non-nil, zero value otherwise.

### GetApplyEnvVarsToBuildOk

`func (o *CreateOrUpdateVersionRequest) GetApplyEnvVarsToBuildOk() (*bool, bool)`

GetApplyEnvVarsToBuildOk returns a tuple with the ApplyEnvVarsToBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyEnvVarsToBuild

`func (o *CreateOrUpdateVersionRequest) SetApplyEnvVarsToBuild(v bool)`

SetApplyEnvVarsToBuild sets ApplyEnvVarsToBuild field to given value.

### HasApplyEnvVarsToBuild

`func (o *CreateOrUpdateVersionRequest) HasApplyEnvVarsToBuild() bool`

HasApplyEnvVarsToBuild returns a boolean if a field has been set.

### SetApplyEnvVarsToBuildNil

`func (o *CreateOrUpdateVersionRequest) SetApplyEnvVarsToBuildNil(b bool)`

 SetApplyEnvVarsToBuildNil sets the value for ApplyEnvVarsToBuild to be an explicit nil

### UnsetApplyEnvVarsToBuild
`func (o *CreateOrUpdateVersionRequest) UnsetApplyEnvVarsToBuild()`

UnsetApplyEnvVarsToBuild ensures that no value is present for ApplyEnvVarsToBuild, not even an explicit nil
### GetBuildTag

`func (o *CreateOrUpdateVersionRequest) GetBuildTag() string`

GetBuildTag returns the BuildTag field if non-nil, zero value otherwise.

### GetBuildTagOk

`func (o *CreateOrUpdateVersionRequest) GetBuildTagOk() (*string, bool)`

GetBuildTagOk returns a tuple with the BuildTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTag

`func (o *CreateOrUpdateVersionRequest) SetBuildTag(v string)`

SetBuildTag sets BuildTag field to given value.

### HasBuildTag

`func (o *CreateOrUpdateVersionRequest) HasBuildTag() bool`

HasBuildTag returns a boolean if a field has been set.

### SetBuildTagNil

`func (o *CreateOrUpdateVersionRequest) SetBuildTagNil(b bool)`

 SetBuildTagNil sets the value for BuildTag to be an explicit nil

### UnsetBuildTag
`func (o *CreateOrUpdateVersionRequest) UnsetBuildTag()`

UnsetBuildTag ensures that no value is present for BuildTag, not even an explicit nil
### GetSourceFiles

`func (o *CreateOrUpdateVersionRequest) GetSourceFiles() []VersionSourceFilesInner`

GetSourceFiles returns the SourceFiles field if non-nil, zero value otherwise.

### GetSourceFilesOk

`func (o *CreateOrUpdateVersionRequest) GetSourceFilesOk() (*[]VersionSourceFilesInner, bool)`

GetSourceFilesOk returns a tuple with the SourceFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFiles

`func (o *CreateOrUpdateVersionRequest) SetSourceFiles(v []VersionSourceFilesInner)`

SetSourceFiles sets SourceFiles field to given value.

### HasSourceFiles

`func (o *CreateOrUpdateVersionRequest) HasSourceFiles() bool`

HasSourceFiles returns a boolean if a field has been set.

### GetGitRepoUrl

`func (o *CreateOrUpdateVersionRequest) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *CreateOrUpdateVersionRequest) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *CreateOrUpdateVersionRequest) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.

### HasGitRepoUrl

`func (o *CreateOrUpdateVersionRequest) HasGitRepoUrl() bool`

HasGitRepoUrl returns a boolean if a field has been set.

### SetGitRepoUrlNil

`func (o *CreateOrUpdateVersionRequest) SetGitRepoUrlNil(b bool)`

 SetGitRepoUrlNil sets the value for GitRepoUrl to be an explicit nil

### UnsetGitRepoUrl
`func (o *CreateOrUpdateVersionRequest) UnsetGitRepoUrl()`

UnsetGitRepoUrl ensures that no value is present for GitRepoUrl, not even an explicit nil
### GetTarballUrl

`func (o *CreateOrUpdateVersionRequest) GetTarballUrl() string`

GetTarballUrl returns the TarballUrl field if non-nil, zero value otherwise.

### GetTarballUrlOk

`func (o *CreateOrUpdateVersionRequest) GetTarballUrlOk() (*string, bool)`

GetTarballUrlOk returns a tuple with the TarballUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarballUrl

`func (o *CreateOrUpdateVersionRequest) SetTarballUrl(v string)`

SetTarballUrl sets TarballUrl field to given value.

### HasTarballUrl

`func (o *CreateOrUpdateVersionRequest) HasTarballUrl() bool`

HasTarballUrl returns a boolean if a field has been set.

### SetTarballUrlNil

`func (o *CreateOrUpdateVersionRequest) SetTarballUrlNil(b bool)`

 SetTarballUrlNil sets the value for TarballUrl to be an explicit nil

### UnsetTarballUrl
`func (o *CreateOrUpdateVersionRequest) UnsetTarballUrl()`

UnsetTarballUrl ensures that no value is present for TarballUrl, not even an explicit nil
### GetGitHubGistUrl

`func (o *CreateOrUpdateVersionRequest) GetGitHubGistUrl() string`

GetGitHubGistUrl returns the GitHubGistUrl field if non-nil, zero value otherwise.

### GetGitHubGistUrlOk

`func (o *CreateOrUpdateVersionRequest) GetGitHubGistUrlOk() (*string, bool)`

GetGitHubGistUrlOk returns a tuple with the GitHubGistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubGistUrl

`func (o *CreateOrUpdateVersionRequest) SetGitHubGistUrl(v string)`

SetGitHubGistUrl sets GitHubGistUrl field to given value.

### HasGitHubGistUrl

`func (o *CreateOrUpdateVersionRequest) HasGitHubGistUrl() bool`

HasGitHubGistUrl returns a boolean if a field has been set.

### SetGitHubGistUrlNil

`func (o *CreateOrUpdateVersionRequest) SetGitHubGistUrlNil(b bool)`

 SetGitHubGistUrlNil sets the value for GitHubGistUrl to be an explicit nil

### UnsetGitHubGistUrl
`func (o *CreateOrUpdateVersionRequest) UnsetGitHubGistUrl()`

UnsetGitHubGistUrl ensures that no value is present for GitHubGistUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


