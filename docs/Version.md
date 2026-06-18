# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionNumber** | **string** |  | 
**SourceType** | [**NullableVersionSourceType**](VersionSourceType.md) |  | 
**EnvVars** | Pointer to [**[]EnvVar**](EnvVar.md) |  | [optional] 
**ApplyEnvVarsToBuild** | Pointer to **NullableBool** |  | [optional] 
**BuildTag** | Pointer to **string** |  | [optional] 
**SourceFiles** | Pointer to [**[]VersionSourceFilesInner**](VersionSourceFilesInner.md) |  | [optional] 
**GitRepoUrl** | Pointer to **NullableString** | URL of the Git repository when sourceType is GIT_REPO. | [optional] 
**TarballUrl** | Pointer to **NullableString** | URL of the tarball when sourceType is TARBALL. | [optional] 
**GitHubGistUrl** | Pointer to **NullableString** | URL of the GitHub Gist when sourceType is GITHUB_GIST. | [optional] 

## Methods

### NewVersion

`func NewVersion(versionNumber string, sourceType NullableVersionSourceType, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionNumber

`func (o *Version) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *Version) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *Version) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.


### GetSourceType

`func (o *Version) GetSourceType() VersionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Version) GetSourceTypeOk() (*VersionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Version) SetSourceType(v VersionSourceType)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *Version) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *Version) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetEnvVars

`func (o *Version) GetEnvVars() []EnvVar`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *Version) GetEnvVarsOk() (*[]EnvVar, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *Version) SetEnvVars(v []EnvVar)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *Version) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### SetEnvVarsNil

`func (o *Version) SetEnvVarsNil(b bool)`

 SetEnvVarsNil sets the value for EnvVars to be an explicit nil

### UnsetEnvVars
`func (o *Version) UnsetEnvVars()`

UnsetEnvVars ensures that no value is present for EnvVars, not even an explicit nil
### GetApplyEnvVarsToBuild

`func (o *Version) GetApplyEnvVarsToBuild() bool`

GetApplyEnvVarsToBuild returns the ApplyEnvVarsToBuild field if non-nil, zero value otherwise.

### GetApplyEnvVarsToBuildOk

`func (o *Version) GetApplyEnvVarsToBuildOk() (*bool, bool)`

GetApplyEnvVarsToBuildOk returns a tuple with the ApplyEnvVarsToBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyEnvVarsToBuild

`func (o *Version) SetApplyEnvVarsToBuild(v bool)`

SetApplyEnvVarsToBuild sets ApplyEnvVarsToBuild field to given value.

### HasApplyEnvVarsToBuild

`func (o *Version) HasApplyEnvVarsToBuild() bool`

HasApplyEnvVarsToBuild returns a boolean if a field has been set.

### SetApplyEnvVarsToBuildNil

`func (o *Version) SetApplyEnvVarsToBuildNil(b bool)`

 SetApplyEnvVarsToBuildNil sets the value for ApplyEnvVarsToBuild to be an explicit nil

### UnsetApplyEnvVarsToBuild
`func (o *Version) UnsetApplyEnvVarsToBuild()`

UnsetApplyEnvVarsToBuild ensures that no value is present for ApplyEnvVarsToBuild, not even an explicit nil
### GetBuildTag

`func (o *Version) GetBuildTag() string`

GetBuildTag returns the BuildTag field if non-nil, zero value otherwise.

### GetBuildTagOk

`func (o *Version) GetBuildTagOk() (*string, bool)`

GetBuildTagOk returns a tuple with the BuildTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTag

`func (o *Version) SetBuildTag(v string)`

SetBuildTag sets BuildTag field to given value.

### HasBuildTag

`func (o *Version) HasBuildTag() bool`

HasBuildTag returns a boolean if a field has been set.

### GetSourceFiles

`func (o *Version) GetSourceFiles() []VersionSourceFilesInner`

GetSourceFiles returns the SourceFiles field if non-nil, zero value otherwise.

### GetSourceFilesOk

`func (o *Version) GetSourceFilesOk() (*[]VersionSourceFilesInner, bool)`

GetSourceFilesOk returns a tuple with the SourceFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFiles

`func (o *Version) SetSourceFiles(v []VersionSourceFilesInner)`

SetSourceFiles sets SourceFiles field to given value.

### HasSourceFiles

`func (o *Version) HasSourceFiles() bool`

HasSourceFiles returns a boolean if a field has been set.

### GetGitRepoUrl

`func (o *Version) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *Version) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *Version) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.

### HasGitRepoUrl

`func (o *Version) HasGitRepoUrl() bool`

HasGitRepoUrl returns a boolean if a field has been set.

### SetGitRepoUrlNil

`func (o *Version) SetGitRepoUrlNil(b bool)`

 SetGitRepoUrlNil sets the value for GitRepoUrl to be an explicit nil

### UnsetGitRepoUrl
`func (o *Version) UnsetGitRepoUrl()`

UnsetGitRepoUrl ensures that no value is present for GitRepoUrl, not even an explicit nil
### GetTarballUrl

`func (o *Version) GetTarballUrl() string`

GetTarballUrl returns the TarballUrl field if non-nil, zero value otherwise.

### GetTarballUrlOk

`func (o *Version) GetTarballUrlOk() (*string, bool)`

GetTarballUrlOk returns a tuple with the TarballUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarballUrl

`func (o *Version) SetTarballUrl(v string)`

SetTarballUrl sets TarballUrl field to given value.

### HasTarballUrl

`func (o *Version) HasTarballUrl() bool`

HasTarballUrl returns a boolean if a field has been set.

### SetTarballUrlNil

`func (o *Version) SetTarballUrlNil(b bool)`

 SetTarballUrlNil sets the value for TarballUrl to be an explicit nil

### UnsetTarballUrl
`func (o *Version) UnsetTarballUrl()`

UnsetTarballUrl ensures that no value is present for TarballUrl, not even an explicit nil
### GetGitHubGistUrl

`func (o *Version) GetGitHubGistUrl() string`

GetGitHubGistUrl returns the GitHubGistUrl field if non-nil, zero value otherwise.

### GetGitHubGistUrlOk

`func (o *Version) GetGitHubGistUrlOk() (*string, bool)`

GetGitHubGistUrlOk returns a tuple with the GitHubGistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubGistUrl

`func (o *Version) SetGitHubGistUrl(v string)`

SetGitHubGistUrl sets GitHubGistUrl field to given value.

### HasGitHubGistUrl

`func (o *Version) HasGitHubGistUrl() bool`

HasGitHubGistUrl returns a boolean if a field has been set.

### SetGitHubGistUrlNil

`func (o *Version) SetGitHubGistUrlNil(b bool)`

 SetGitHubGistUrlNil sets the value for GitHubGistUrl to be an explicit nil

### UnsetGitHubGistUrl
`func (o *Version) UnsetGitHubGistUrl()`

UnsetGitHubGistUrl ensures that no value is present for GitHubGistUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


