# ActorDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorSpecification** | Pointer to **int32** | The Actor specification version that this Actor follows. This property must be set to 1. | [optional] 
**Name** | Pointer to **string** | The name of the Actor. | [optional] 
**Version** | Pointer to **string** | The version of the Actor, typically a dot-separated sequence of numbers (e.g., &#x60;0.1&#x60;, &#x60;1.0&#x60;, or &#x60;0.0.1&#x60;). | [optional] 
**BuildTag** | Pointer to **string** | The tag name to be applied to a successful build of the Actor. Defaults to &#39;latest&#39; if not specified. | [optional] 
**EnvironmentVariables** | Pointer to **map[string]string** | A map of environment variables to be used during local development and deployment. | [optional] 
**Dockerfile** | Pointer to **string** | The path to the Dockerfile used for building the Actor on the platform. | [optional] 
**DockerContextDir** | Pointer to **string** | The path to the directory used as the Docker context when building the Actor. | [optional] 
**Readme** | Pointer to **string** | The path to the README file for the Actor. | [optional] 
**Input** | Pointer to **map[string]interface{}** | The input schema object, the full specification can be found in [Apify docs](https://docs.apify.com/platform/actors/development/actor-definition/input-schema) | [optional] 
**Changelog** | Pointer to **string** | The path to the CHANGELOG file displayed in the Actor&#39;s information tab. | [optional] 
**Storages** | Pointer to [**ActorDefinitionStorages**](ActorDefinitionStorages.md) |  | [optional] 
**DefaultMemoryMbytes** | Pointer to [**ActorDefinitionDefaultMemoryMbytes**](ActorDefinitionDefaultMemoryMbytes.md) |  | [optional] 
**MinMemoryMbytes** | Pointer to **int32** | Specifies the minimum amount of memory in megabytes required by the Actor. | [optional] 
**MaxMemoryMbytes** | Pointer to **int32** | Specifies the maximum amount of memory in megabytes required by the Actor. | [optional] 
**UsesStandbyMode** | Pointer to **bool** | Specifies whether Standby mode is enabled for the Actor. | [optional] 

## Methods

### NewActorDefinition

`func NewActorDefinition() *ActorDefinition`

NewActorDefinition instantiates a new ActorDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorDefinitionWithDefaults

`func NewActorDefinitionWithDefaults() *ActorDefinition`

NewActorDefinitionWithDefaults instantiates a new ActorDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorSpecification

`func (o *ActorDefinition) GetActorSpecification() int32`

GetActorSpecification returns the ActorSpecification field if non-nil, zero value otherwise.

### GetActorSpecificationOk

`func (o *ActorDefinition) GetActorSpecificationOk() (*int32, bool)`

GetActorSpecificationOk returns a tuple with the ActorSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorSpecification

`func (o *ActorDefinition) SetActorSpecification(v int32)`

SetActorSpecification sets ActorSpecification field to given value.

### HasActorSpecification

`func (o *ActorDefinition) HasActorSpecification() bool`

HasActorSpecification returns a boolean if a field has been set.

### GetName

`func (o *ActorDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActorDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActorDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActorDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *ActorDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ActorDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ActorDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ActorDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuildTag

`func (o *ActorDefinition) GetBuildTag() string`

GetBuildTag returns the BuildTag field if non-nil, zero value otherwise.

### GetBuildTagOk

`func (o *ActorDefinition) GetBuildTagOk() (*string, bool)`

GetBuildTagOk returns a tuple with the BuildTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTag

`func (o *ActorDefinition) SetBuildTag(v string)`

SetBuildTag sets BuildTag field to given value.

### HasBuildTag

`func (o *ActorDefinition) HasBuildTag() bool`

HasBuildTag returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ActorDefinition) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ActorDefinition) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ActorDefinition) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ActorDefinition) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetDockerfile

`func (o *ActorDefinition) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *ActorDefinition) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *ActorDefinition) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *ActorDefinition) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetDockerContextDir

`func (o *ActorDefinition) GetDockerContextDir() string`

GetDockerContextDir returns the DockerContextDir field if non-nil, zero value otherwise.

### GetDockerContextDirOk

`func (o *ActorDefinition) GetDockerContextDirOk() (*string, bool)`

GetDockerContextDirOk returns a tuple with the DockerContextDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContextDir

`func (o *ActorDefinition) SetDockerContextDir(v string)`

SetDockerContextDir sets DockerContextDir field to given value.

### HasDockerContextDir

`func (o *ActorDefinition) HasDockerContextDir() bool`

HasDockerContextDir returns a boolean if a field has been set.

### GetReadme

`func (o *ActorDefinition) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *ActorDefinition) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *ActorDefinition) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *ActorDefinition) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### GetInput

`func (o *ActorDefinition) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ActorDefinition) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ActorDefinition) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ActorDefinition) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetChangelog

`func (o *ActorDefinition) GetChangelog() string`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *ActorDefinition) GetChangelogOk() (*string, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *ActorDefinition) SetChangelog(v string)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *ActorDefinition) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetStorages

`func (o *ActorDefinition) GetStorages() ActorDefinitionStorages`

GetStorages returns the Storages field if non-nil, zero value otherwise.

### GetStoragesOk

`func (o *ActorDefinition) GetStoragesOk() (*ActorDefinitionStorages, bool)`

GetStoragesOk returns a tuple with the Storages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorages

`func (o *ActorDefinition) SetStorages(v ActorDefinitionStorages)`

SetStorages sets Storages field to given value.

### HasStorages

`func (o *ActorDefinition) HasStorages() bool`

HasStorages returns a boolean if a field has been set.

### GetDefaultMemoryMbytes

`func (o *ActorDefinition) GetDefaultMemoryMbytes() ActorDefinitionDefaultMemoryMbytes`

GetDefaultMemoryMbytes returns the DefaultMemoryMbytes field if non-nil, zero value otherwise.

### GetDefaultMemoryMbytesOk

`func (o *ActorDefinition) GetDefaultMemoryMbytesOk() (*ActorDefinitionDefaultMemoryMbytes, bool)`

GetDefaultMemoryMbytesOk returns a tuple with the DefaultMemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemoryMbytes

`func (o *ActorDefinition) SetDefaultMemoryMbytes(v ActorDefinitionDefaultMemoryMbytes)`

SetDefaultMemoryMbytes sets DefaultMemoryMbytes field to given value.

### HasDefaultMemoryMbytes

`func (o *ActorDefinition) HasDefaultMemoryMbytes() bool`

HasDefaultMemoryMbytes returns a boolean if a field has been set.

### GetMinMemoryMbytes

`func (o *ActorDefinition) GetMinMemoryMbytes() int32`

GetMinMemoryMbytes returns the MinMemoryMbytes field if non-nil, zero value otherwise.

### GetMinMemoryMbytesOk

`func (o *ActorDefinition) GetMinMemoryMbytesOk() (*int32, bool)`

GetMinMemoryMbytesOk returns a tuple with the MinMemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryMbytes

`func (o *ActorDefinition) SetMinMemoryMbytes(v int32)`

SetMinMemoryMbytes sets MinMemoryMbytes field to given value.

### HasMinMemoryMbytes

`func (o *ActorDefinition) HasMinMemoryMbytes() bool`

HasMinMemoryMbytes returns a boolean if a field has been set.

### GetMaxMemoryMbytes

`func (o *ActorDefinition) GetMaxMemoryMbytes() int32`

GetMaxMemoryMbytes returns the MaxMemoryMbytes field if non-nil, zero value otherwise.

### GetMaxMemoryMbytesOk

`func (o *ActorDefinition) GetMaxMemoryMbytesOk() (*int32, bool)`

GetMaxMemoryMbytesOk returns a tuple with the MaxMemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMbytes

`func (o *ActorDefinition) SetMaxMemoryMbytes(v int32)`

SetMaxMemoryMbytes sets MaxMemoryMbytes field to given value.

### HasMaxMemoryMbytes

`func (o *ActorDefinition) HasMaxMemoryMbytes() bool`

HasMaxMemoryMbytes returns a boolean if a field has been set.

### GetUsesStandbyMode

`func (o *ActorDefinition) GetUsesStandbyMode() bool`

GetUsesStandbyMode returns the UsesStandbyMode field if non-nil, zero value otherwise.

### GetUsesStandbyModeOk

`func (o *ActorDefinition) GetUsesStandbyModeOk() (*bool, bool)`

GetUsesStandbyModeOk returns a tuple with the UsesStandbyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesStandbyMode

`func (o *ActorDefinition) SetUsesStandbyMode(v bool)`

SetUsesStandbyMode sets UsesStandbyMode field to given value.

### HasUsesStandbyMode

`func (o *ActorDefinition) HasUsesStandbyMode() bool`

HasUsesStandbyMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


