# BuildActVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | Pointer to [**VersionSourceType**](VersionSourceType.md) |  | [optional] 
**BuildTag** | Pointer to **string** |  | [optional] 
**VersionNumber** | Pointer to **string** |  | [optional] 
**GitRepoUrl** | Pointer to **string** | URL of the git repository, present when sourceType is GIT_REPO. | [optional] 
**SourceFiles** | Pointer to [**[]SourceCodeFile**](SourceCodeFile.md) | Inline source files, present when sourceType is SOURCE_FILES. | [optional] 

## Methods

### NewBuildActVersion

`func NewBuildActVersion() *BuildActVersion`

NewBuildActVersion instantiates a new BuildActVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildActVersionWithDefaults

`func NewBuildActVersionWithDefaults() *BuildActVersion`

NewBuildActVersionWithDefaults instantiates a new BuildActVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *BuildActVersion) GetSourceType() VersionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BuildActVersion) GetSourceTypeOk() (*VersionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BuildActVersion) SetSourceType(v VersionSourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *BuildActVersion) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetBuildTag

`func (o *BuildActVersion) GetBuildTag() string`

GetBuildTag returns the BuildTag field if non-nil, zero value otherwise.

### GetBuildTagOk

`func (o *BuildActVersion) GetBuildTagOk() (*string, bool)`

GetBuildTagOk returns a tuple with the BuildTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTag

`func (o *BuildActVersion) SetBuildTag(v string)`

SetBuildTag sets BuildTag field to given value.

### HasBuildTag

`func (o *BuildActVersion) HasBuildTag() bool`

HasBuildTag returns a boolean if a field has been set.

### GetVersionNumber

`func (o *BuildActVersion) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *BuildActVersion) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *BuildActVersion) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *BuildActVersion) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetGitRepoUrl

`func (o *BuildActVersion) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *BuildActVersion) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *BuildActVersion) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.

### HasGitRepoUrl

`func (o *BuildActVersion) HasGitRepoUrl() bool`

HasGitRepoUrl returns a boolean if a field has been set.

### GetSourceFiles

`func (o *BuildActVersion) GetSourceFiles() []SourceCodeFile`

GetSourceFiles returns the SourceFiles field if non-nil, zero value otherwise.

### GetSourceFilesOk

`func (o *BuildActVersion) GetSourceFilesOk() (*[]SourceCodeFile, bool)`

GetSourceFilesOk returns a tuple with the SourceFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFiles

`func (o *BuildActVersion) SetSourceFiles(v []SourceCodeFile)`

SetSourceFiles sets SourceFiles field to given value.

### HasSourceFiles

`func (o *BuildActVersion) HasSourceFiles() bool`

HasSourceFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


