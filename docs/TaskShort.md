# TaskShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**ActId** | **string** |  | 
**ActName** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Username** | Pointer to **NullableString** |  | [optional] 
**ActUsername** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Stats** | Pointer to [**NullableTaskStats**](TaskStats.md) |  | [optional] 

## Methods

### NewTaskShort

`func NewTaskShort(id string, userId string, actId string, name string, createdAt time.Time, modifiedAt time.Time, ) *TaskShort`

NewTaskShort instantiates a new TaskShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskShortWithDefaults

`func NewTaskShortWithDefaults() *TaskShort`

NewTaskShortWithDefaults instantiates a new TaskShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskShort) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *TaskShort) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TaskShort) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TaskShort) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetActId

`func (o *TaskShort) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *TaskShort) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *TaskShort) SetActId(v string)`

SetActId sets ActId field to given value.


### GetActName

`func (o *TaskShort) GetActName() string`

GetActName returns the ActName field if non-nil, zero value otherwise.

### GetActNameOk

`func (o *TaskShort) GetActNameOk() (*string, bool)`

GetActNameOk returns a tuple with the ActName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActName

`func (o *TaskShort) SetActName(v string)`

SetActName sets ActName field to given value.

### HasActName

`func (o *TaskShort) HasActName() bool`

HasActName returns a boolean if a field has been set.

### SetActNameNil

`func (o *TaskShort) SetActNameNil(b bool)`

 SetActNameNil sets the value for ActName to be an explicit nil

### UnsetActName
`func (o *TaskShort) UnsetActName()`

UnsetActName ensures that no value is present for ActName, not even an explicit nil
### GetName

`func (o *TaskShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskShort) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *TaskShort) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskShort) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskShort) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskShort) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskShort) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskShort) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetActUsername

`func (o *TaskShort) GetActUsername() string`

GetActUsername returns the ActUsername field if non-nil, zero value otherwise.

### GetActUsernameOk

`func (o *TaskShort) GetActUsernameOk() (*string, bool)`

GetActUsernameOk returns a tuple with the ActUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActUsername

`func (o *TaskShort) SetActUsername(v string)`

SetActUsername sets ActUsername field to given value.

### HasActUsername

`func (o *TaskShort) HasActUsername() bool`

HasActUsername returns a boolean if a field has been set.

### SetActUsernameNil

`func (o *TaskShort) SetActUsernameNil(b bool)`

 SetActUsernameNil sets the value for ActUsername to be an explicit nil

### UnsetActUsername
`func (o *TaskShort) UnsetActUsername()`

UnsetActUsername ensures that no value is present for ActUsername, not even an explicit nil
### GetCreatedAt

`func (o *TaskShort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskShort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskShort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *TaskShort) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TaskShort) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TaskShort) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetStats

`func (o *TaskShort) GetStats() TaskStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TaskShort) GetStatsOk() (*TaskStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TaskShort) SetStats(v TaskStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *TaskShort) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *TaskShort) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *TaskShort) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


