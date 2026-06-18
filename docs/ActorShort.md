# ActorShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Username** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**NullableActorStats**](ActorStats.md) |  | [optional] 

## Methods

### NewActorShort

`func NewActorShort(id string, createdAt time.Time, modifiedAt time.Time, name string, username string, ) *ActorShort`

NewActorShort instantiates a new ActorShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorShortWithDefaults

`func NewActorShortWithDefaults() *ActorShort`

NewActorShortWithDefaults instantiates a new ActorShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActorShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActorShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActorShort) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ActorShort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActorShort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActorShort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ActorShort) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ActorShort) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ActorShort) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetName

`func (o *ActorShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActorShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActorShort) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *ActorShort) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActorShort) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActorShort) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetTitle

`func (o *ActorShort) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActorShort) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActorShort) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActorShort) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStats

`func (o *ActorShort) GetStats() ActorStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ActorShort) GetStatsOk() (*ActorStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ActorShort) SetStats(v ActorStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ActorShort) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *ActorShort) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *ActorShort) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


