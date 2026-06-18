# DatasetListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**UserId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**AccessedAt** | **time.Time** |  | 
**ItemCount** | **int32** |  | 
**CleanItemCount** | **int32** |  | 
**ActId** | Pointer to **NullableString** |  | [optional] 
**ActRunId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**GeneralAccess** | Pointer to [**GeneralAccess**](GeneralAccess.md) |  | [optional] 
**Stats** | Pointer to [**DatasetStats**](DatasetStats.md) |  | [optional] 

## Methods

### NewDatasetListItem

`func NewDatasetListItem(id string, name string, userId string, createdAt time.Time, modifiedAt time.Time, accessedAt time.Time, itemCount int32, cleanItemCount int32, ) *DatasetListItem`

NewDatasetListItem instantiates a new DatasetListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetListItemWithDefaults

`func NewDatasetListItemWithDefaults() *DatasetListItem`

NewDatasetListItemWithDefaults instantiates a new DatasetListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatasetListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasetListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasetListItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DatasetListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetListItem) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *DatasetListItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DatasetListItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DatasetListItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *DatasetListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatasetListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatasetListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *DatasetListItem) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DatasetListItem) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DatasetListItem) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetAccessedAt

`func (o *DatasetListItem) GetAccessedAt() time.Time`

GetAccessedAt returns the AccessedAt field if non-nil, zero value otherwise.

### GetAccessedAtOk

`func (o *DatasetListItem) GetAccessedAtOk() (*time.Time, bool)`

GetAccessedAtOk returns a tuple with the AccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedAt

`func (o *DatasetListItem) SetAccessedAt(v time.Time)`

SetAccessedAt sets AccessedAt field to given value.


### GetItemCount

`func (o *DatasetListItem) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *DatasetListItem) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *DatasetListItem) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.


### GetCleanItemCount

`func (o *DatasetListItem) GetCleanItemCount() int32`

GetCleanItemCount returns the CleanItemCount field if non-nil, zero value otherwise.

### GetCleanItemCountOk

`func (o *DatasetListItem) GetCleanItemCountOk() (*int32, bool)`

GetCleanItemCountOk returns a tuple with the CleanItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanItemCount

`func (o *DatasetListItem) SetCleanItemCount(v int32)`

SetCleanItemCount sets CleanItemCount field to given value.


### GetActId

`func (o *DatasetListItem) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *DatasetListItem) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *DatasetListItem) SetActId(v string)`

SetActId sets ActId field to given value.

### HasActId

`func (o *DatasetListItem) HasActId() bool`

HasActId returns a boolean if a field has been set.

### SetActIdNil

`func (o *DatasetListItem) SetActIdNil(b bool)`

 SetActIdNil sets the value for ActId to be an explicit nil

### UnsetActId
`func (o *DatasetListItem) UnsetActId()`

UnsetActId ensures that no value is present for ActId, not even an explicit nil
### GetActRunId

`func (o *DatasetListItem) GetActRunId() string`

GetActRunId returns the ActRunId field if non-nil, zero value otherwise.

### GetActRunIdOk

`func (o *DatasetListItem) GetActRunIdOk() (*string, bool)`

GetActRunIdOk returns a tuple with the ActRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActRunId

`func (o *DatasetListItem) SetActRunId(v string)`

SetActRunId sets ActRunId field to given value.

### HasActRunId

`func (o *DatasetListItem) HasActRunId() bool`

HasActRunId returns a boolean if a field has been set.

### SetActRunIdNil

`func (o *DatasetListItem) SetActRunIdNil(b bool)`

 SetActRunIdNil sets the value for ActRunId to be an explicit nil

### UnsetActRunId
`func (o *DatasetListItem) UnsetActRunId()`

UnsetActRunId ensures that no value is present for ActRunId, not even an explicit nil
### GetTitle

`func (o *DatasetListItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DatasetListItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DatasetListItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DatasetListItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DatasetListItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DatasetListItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUsername

`func (o *DatasetListItem) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DatasetListItem) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DatasetListItem) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DatasetListItem) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetGeneralAccess

`func (o *DatasetListItem) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *DatasetListItem) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *DatasetListItem) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.

### HasGeneralAccess

`func (o *DatasetListItem) HasGeneralAccess() bool`

HasGeneralAccess returns a boolean if a field has been set.

### GetStats

`func (o *DatasetListItem) GetStats() DatasetStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DatasetListItem) GetStatsOk() (*DatasetStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DatasetListItem) SetStats(v DatasetStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *DatasetListItem) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


