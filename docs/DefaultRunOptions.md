# DefaultRunOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to **string** |  | [optional] 
**TimeoutSecs** | Pointer to **int32** |  | [optional] 
**MemoryMbytes** | Pointer to **int32** |  | [optional] 
**RestartOnError** | Pointer to **bool** |  | [optional] 
**MaxItems** | Pointer to **NullableInt32** |  | [optional] 
**ForcePermissionLevel** | Pointer to [**NullableActorPermissionLevel**](ActorPermissionLevel.md) |  | [optional] 

## Methods

### NewDefaultRunOptions

`func NewDefaultRunOptions() *DefaultRunOptions`

NewDefaultRunOptions instantiates a new DefaultRunOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultRunOptionsWithDefaults

`func NewDefaultRunOptionsWithDefaults() *DefaultRunOptions`

NewDefaultRunOptionsWithDefaults instantiates a new DefaultRunOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *DefaultRunOptions) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *DefaultRunOptions) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *DefaultRunOptions) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *DefaultRunOptions) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *DefaultRunOptions) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *DefaultRunOptions) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *DefaultRunOptions) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *DefaultRunOptions) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetMemoryMbytes

`func (o *DefaultRunOptions) GetMemoryMbytes() int32`

GetMemoryMbytes returns the MemoryMbytes field if non-nil, zero value otherwise.

### GetMemoryMbytesOk

`func (o *DefaultRunOptions) GetMemoryMbytesOk() (*int32, bool)`

GetMemoryMbytesOk returns a tuple with the MemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMbytes

`func (o *DefaultRunOptions) SetMemoryMbytes(v int32)`

SetMemoryMbytes sets MemoryMbytes field to given value.

### HasMemoryMbytes

`func (o *DefaultRunOptions) HasMemoryMbytes() bool`

HasMemoryMbytes returns a boolean if a field has been set.

### GetRestartOnError

`func (o *DefaultRunOptions) GetRestartOnError() bool`

GetRestartOnError returns the RestartOnError field if non-nil, zero value otherwise.

### GetRestartOnErrorOk

`func (o *DefaultRunOptions) GetRestartOnErrorOk() (*bool, bool)`

GetRestartOnErrorOk returns a tuple with the RestartOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnError

`func (o *DefaultRunOptions) SetRestartOnError(v bool)`

SetRestartOnError sets RestartOnError field to given value.

### HasRestartOnError

`func (o *DefaultRunOptions) HasRestartOnError() bool`

HasRestartOnError returns a boolean if a field has been set.

### GetMaxItems

`func (o *DefaultRunOptions) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *DefaultRunOptions) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *DefaultRunOptions) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *DefaultRunOptions) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *DefaultRunOptions) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *DefaultRunOptions) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetForcePermissionLevel

`func (o *DefaultRunOptions) GetForcePermissionLevel() ActorPermissionLevel`

GetForcePermissionLevel returns the ForcePermissionLevel field if non-nil, zero value otherwise.

### GetForcePermissionLevelOk

`func (o *DefaultRunOptions) GetForcePermissionLevelOk() (*ActorPermissionLevel, bool)`

GetForcePermissionLevelOk returns a tuple with the ForcePermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePermissionLevel

`func (o *DefaultRunOptions) SetForcePermissionLevel(v ActorPermissionLevel)`

SetForcePermissionLevel sets ForcePermissionLevel field to given value.

### HasForcePermissionLevel

`func (o *DefaultRunOptions) HasForcePermissionLevel() bool`

HasForcePermissionLevel returns a boolean if a field has been set.

### SetForcePermissionLevelNil

`func (o *DefaultRunOptions) SetForcePermissionLevelNil(b bool)`

 SetForcePermissionLevelNil sets the value for ForcePermissionLevel to be an explicit nil

### UnsetForcePermissionLevel
`func (o *DefaultRunOptions) UnsetForcePermissionLevel()`

UnsetForcePermissionLevel ensures that no value is present for ForcePermissionLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


