# ActorStandby

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 
**DesiredRequestsPerActorRun** | Pointer to **NullableInt32** |  | [optional] 
**MaxRequestsPerActorRun** | Pointer to **NullableInt32** |  | [optional] 
**IdleTimeoutSecs** | Pointer to **NullableInt32** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**MemoryMbytes** | Pointer to **NullableInt32** |  | [optional] 
**DisableStandbyFieldsOverride** | Pointer to **NullableBool** |  | [optional] 
**ShouldPassActorInput** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewActorStandby

`func NewActorStandby() *ActorStandby`

NewActorStandby instantiates a new ActorStandby object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorStandbyWithDefaults

`func NewActorStandbyWithDefaults() *ActorStandby`

NewActorStandbyWithDefaults instantiates a new ActorStandby object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ActorStandby) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ActorStandby) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ActorStandby) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ActorStandby) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *ActorStandby) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ActorStandby) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetDesiredRequestsPerActorRun

`func (o *ActorStandby) GetDesiredRequestsPerActorRun() int32`

GetDesiredRequestsPerActorRun returns the DesiredRequestsPerActorRun field if non-nil, zero value otherwise.

### GetDesiredRequestsPerActorRunOk

`func (o *ActorStandby) GetDesiredRequestsPerActorRunOk() (*int32, bool)`

GetDesiredRequestsPerActorRunOk returns a tuple with the DesiredRequestsPerActorRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredRequestsPerActorRun

`func (o *ActorStandby) SetDesiredRequestsPerActorRun(v int32)`

SetDesiredRequestsPerActorRun sets DesiredRequestsPerActorRun field to given value.

### HasDesiredRequestsPerActorRun

`func (o *ActorStandby) HasDesiredRequestsPerActorRun() bool`

HasDesiredRequestsPerActorRun returns a boolean if a field has been set.

### SetDesiredRequestsPerActorRunNil

`func (o *ActorStandby) SetDesiredRequestsPerActorRunNil(b bool)`

 SetDesiredRequestsPerActorRunNil sets the value for DesiredRequestsPerActorRun to be an explicit nil

### UnsetDesiredRequestsPerActorRun
`func (o *ActorStandby) UnsetDesiredRequestsPerActorRun()`

UnsetDesiredRequestsPerActorRun ensures that no value is present for DesiredRequestsPerActorRun, not even an explicit nil
### GetMaxRequestsPerActorRun

`func (o *ActorStandby) GetMaxRequestsPerActorRun() int32`

GetMaxRequestsPerActorRun returns the MaxRequestsPerActorRun field if non-nil, zero value otherwise.

### GetMaxRequestsPerActorRunOk

`func (o *ActorStandby) GetMaxRequestsPerActorRunOk() (*int32, bool)`

GetMaxRequestsPerActorRunOk returns a tuple with the MaxRequestsPerActorRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestsPerActorRun

`func (o *ActorStandby) SetMaxRequestsPerActorRun(v int32)`

SetMaxRequestsPerActorRun sets MaxRequestsPerActorRun field to given value.

### HasMaxRequestsPerActorRun

`func (o *ActorStandby) HasMaxRequestsPerActorRun() bool`

HasMaxRequestsPerActorRun returns a boolean if a field has been set.

### SetMaxRequestsPerActorRunNil

`func (o *ActorStandby) SetMaxRequestsPerActorRunNil(b bool)`

 SetMaxRequestsPerActorRunNil sets the value for MaxRequestsPerActorRun to be an explicit nil

### UnsetMaxRequestsPerActorRun
`func (o *ActorStandby) UnsetMaxRequestsPerActorRun()`

UnsetMaxRequestsPerActorRun ensures that no value is present for MaxRequestsPerActorRun, not even an explicit nil
### GetIdleTimeoutSecs

`func (o *ActorStandby) GetIdleTimeoutSecs() int32`

GetIdleTimeoutSecs returns the IdleTimeoutSecs field if non-nil, zero value otherwise.

### GetIdleTimeoutSecsOk

`func (o *ActorStandby) GetIdleTimeoutSecsOk() (*int32, bool)`

GetIdleTimeoutSecsOk returns a tuple with the IdleTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutSecs

`func (o *ActorStandby) SetIdleTimeoutSecs(v int32)`

SetIdleTimeoutSecs sets IdleTimeoutSecs field to given value.

### HasIdleTimeoutSecs

`func (o *ActorStandby) HasIdleTimeoutSecs() bool`

HasIdleTimeoutSecs returns a boolean if a field has been set.

### SetIdleTimeoutSecsNil

`func (o *ActorStandby) SetIdleTimeoutSecsNil(b bool)`

 SetIdleTimeoutSecsNil sets the value for IdleTimeoutSecs to be an explicit nil

### UnsetIdleTimeoutSecs
`func (o *ActorStandby) UnsetIdleTimeoutSecs()`

UnsetIdleTimeoutSecs ensures that no value is present for IdleTimeoutSecs, not even an explicit nil
### GetBuild

`func (o *ActorStandby) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ActorStandby) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ActorStandby) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ActorStandby) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *ActorStandby) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *ActorStandby) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetMemoryMbytes

`func (o *ActorStandby) GetMemoryMbytes() int32`

GetMemoryMbytes returns the MemoryMbytes field if non-nil, zero value otherwise.

### GetMemoryMbytesOk

`func (o *ActorStandby) GetMemoryMbytesOk() (*int32, bool)`

GetMemoryMbytesOk returns a tuple with the MemoryMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMbytes

`func (o *ActorStandby) SetMemoryMbytes(v int32)`

SetMemoryMbytes sets MemoryMbytes field to given value.

### HasMemoryMbytes

`func (o *ActorStandby) HasMemoryMbytes() bool`

HasMemoryMbytes returns a boolean if a field has been set.

### SetMemoryMbytesNil

`func (o *ActorStandby) SetMemoryMbytesNil(b bool)`

 SetMemoryMbytesNil sets the value for MemoryMbytes to be an explicit nil

### UnsetMemoryMbytes
`func (o *ActorStandby) UnsetMemoryMbytes()`

UnsetMemoryMbytes ensures that no value is present for MemoryMbytes, not even an explicit nil
### GetDisableStandbyFieldsOverride

`func (o *ActorStandby) GetDisableStandbyFieldsOverride() bool`

GetDisableStandbyFieldsOverride returns the DisableStandbyFieldsOverride field if non-nil, zero value otherwise.

### GetDisableStandbyFieldsOverrideOk

`func (o *ActorStandby) GetDisableStandbyFieldsOverrideOk() (*bool, bool)`

GetDisableStandbyFieldsOverrideOk returns a tuple with the DisableStandbyFieldsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStandbyFieldsOverride

`func (o *ActorStandby) SetDisableStandbyFieldsOverride(v bool)`

SetDisableStandbyFieldsOverride sets DisableStandbyFieldsOverride field to given value.

### HasDisableStandbyFieldsOverride

`func (o *ActorStandby) HasDisableStandbyFieldsOverride() bool`

HasDisableStandbyFieldsOverride returns a boolean if a field has been set.

### SetDisableStandbyFieldsOverrideNil

`func (o *ActorStandby) SetDisableStandbyFieldsOverrideNil(b bool)`

 SetDisableStandbyFieldsOverrideNil sets the value for DisableStandbyFieldsOverride to be an explicit nil

### UnsetDisableStandbyFieldsOverride
`func (o *ActorStandby) UnsetDisableStandbyFieldsOverride()`

UnsetDisableStandbyFieldsOverride ensures that no value is present for DisableStandbyFieldsOverride, not even an explicit nil
### GetShouldPassActorInput

`func (o *ActorStandby) GetShouldPassActorInput() bool`

GetShouldPassActorInput returns the ShouldPassActorInput field if non-nil, zero value otherwise.

### GetShouldPassActorInputOk

`func (o *ActorStandby) GetShouldPassActorInputOk() (*bool, bool)`

GetShouldPassActorInputOk returns a tuple with the ShouldPassActorInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldPassActorInput

`func (o *ActorStandby) SetShouldPassActorInput(v bool)`

SetShouldPassActorInput sets ShouldPassActorInput field to given value.

### HasShouldPassActorInput

`func (o *ActorStandby) HasShouldPassActorInput() bool`

HasShouldPassActorInput returns a boolean if a field has been set.

### SetShouldPassActorInputNil

`func (o *ActorStandby) SetShouldPassActorInputNil(b bool)`

 SetShouldPassActorInputNil sets the value for ShouldPassActorInput to be an explicit nil

### UnsetShouldPassActorInput
`func (o *ActorStandby) UnsetShouldPassActorInput()`

UnsetShouldPassActorInput ensures that no value is present for ShouldPassActorInput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


