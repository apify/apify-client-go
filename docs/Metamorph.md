# Metamorph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Time when the metamorph occurred. | 
**ActorId** | **string** | ID of the Actor that the run was metamorphed to. | 
**BuildId** | **string** | ID of the build used for the metamorphed Actor. | 
**InputKey** | Pointer to **NullableString** | Key of the input record in the key-value store. | [optional] 

## Methods

### NewMetamorph

`func NewMetamorph(createdAt time.Time, actorId string, buildId string, ) *Metamorph`

NewMetamorph instantiates a new Metamorph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetamorphWithDefaults

`func NewMetamorphWithDefaults() *Metamorph`

NewMetamorphWithDefaults instantiates a new Metamorph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Metamorph) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Metamorph) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Metamorph) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActorId

`func (o *Metamorph) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *Metamorph) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *Metamorph) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetBuildId

`func (o *Metamorph) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *Metamorph) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *Metamorph) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.


### GetInputKey

`func (o *Metamorph) GetInputKey() string`

GetInputKey returns the InputKey field if non-nil, zero value otherwise.

### GetInputKeyOk

`func (o *Metamorph) GetInputKeyOk() (*string, bool)`

GetInputKeyOk returns a tuple with the InputKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputKey

`func (o *Metamorph) SetInputKey(v string)`

SetInputKey sets InputKey field to given value.

### HasInputKey

`func (o *Metamorph) HasInputKey() bool`

HasInputKey returns a boolean if a field has been set.

### SetInputKeyNil

`func (o *Metamorph) SetInputKeyNil(b bool)`

 SetInputKeyNil sets the value for InputKey to be an explicit nil

### UnsetInputKey
`func (o *Metamorph) UnsetInputKey()`

UnsetInputKey ensures that no value is present for InputKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


