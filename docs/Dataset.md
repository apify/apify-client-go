# Dataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**UserId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**AccessedAt** | **time.Time** |  | 
**ItemCount** | **int32** |  | 
**CleanItemCount** | **int32** |  | 
**ActId** | Pointer to **NullableString** |  | [optional] 
**ActRunId** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Schema** | Pointer to **map[string]interface{}** | Defines the schema of items in your dataset, the full specification can be found in [Apify docs](/platform/actors/development/actor-definition/dataset-schema) | [optional] 
**ConsoleUrl** | **string** |  | 
**ItemsPublicUrl** | Pointer to **string** | A public link to access the dataset items directly. | [optional] 
**UrlSigningSecretKey** | Pointer to **NullableString** | A secret key for generating signed public URLs. It is only provided to clients with WRITE permission for the dataset. | [optional] 
**GeneralAccess** | Pointer to [**GeneralAccess**](GeneralAccess.md) |  | [optional] 
**Stats** | Pointer to [**DatasetStats**](DatasetStats.md) |  | [optional] 

## Methods

### NewDataset

`func NewDataset(id string, userId string, createdAt time.Time, modifiedAt time.Time, accessedAt time.Time, itemCount int32, cleanItemCount int32, consoleUrl string, ) *Dataset`

NewDataset instantiates a new Dataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetWithDefaults

`func NewDatasetWithDefaults() *Dataset`

NewDatasetWithDefaults instantiates a new Dataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dataset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dataset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dataset) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Dataset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dataset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dataset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dataset) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Dataset) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Dataset) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserId

`func (o *Dataset) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Dataset) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Dataset) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *Dataset) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dataset) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dataset) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Dataset) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Dataset) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Dataset) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetAccessedAt

`func (o *Dataset) GetAccessedAt() time.Time`

GetAccessedAt returns the AccessedAt field if non-nil, zero value otherwise.

### GetAccessedAtOk

`func (o *Dataset) GetAccessedAtOk() (*time.Time, bool)`

GetAccessedAtOk returns a tuple with the AccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedAt

`func (o *Dataset) SetAccessedAt(v time.Time)`

SetAccessedAt sets AccessedAt field to given value.


### GetItemCount

`func (o *Dataset) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *Dataset) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *Dataset) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.


### GetCleanItemCount

`func (o *Dataset) GetCleanItemCount() int32`

GetCleanItemCount returns the CleanItemCount field if non-nil, zero value otherwise.

### GetCleanItemCountOk

`func (o *Dataset) GetCleanItemCountOk() (*int32, bool)`

GetCleanItemCountOk returns a tuple with the CleanItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanItemCount

`func (o *Dataset) SetCleanItemCount(v int32)`

SetCleanItemCount sets CleanItemCount field to given value.


### GetActId

`func (o *Dataset) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *Dataset) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *Dataset) SetActId(v string)`

SetActId sets ActId field to given value.

### HasActId

`func (o *Dataset) HasActId() bool`

HasActId returns a boolean if a field has been set.

### SetActIdNil

`func (o *Dataset) SetActIdNil(b bool)`

 SetActIdNil sets the value for ActId to be an explicit nil

### UnsetActId
`func (o *Dataset) UnsetActId()`

UnsetActId ensures that no value is present for ActId, not even an explicit nil
### GetActRunId

`func (o *Dataset) GetActRunId() string`

GetActRunId returns the ActRunId field if non-nil, zero value otherwise.

### GetActRunIdOk

`func (o *Dataset) GetActRunIdOk() (*string, bool)`

GetActRunIdOk returns a tuple with the ActRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActRunId

`func (o *Dataset) SetActRunId(v string)`

SetActRunId sets ActRunId field to given value.

### HasActRunId

`func (o *Dataset) HasActRunId() bool`

HasActRunId returns a boolean if a field has been set.

### SetActRunIdNil

`func (o *Dataset) SetActRunIdNil(b bool)`

 SetActRunIdNil sets the value for ActRunId to be an explicit nil

### UnsetActRunId
`func (o *Dataset) UnsetActRunId()`

UnsetActRunId ensures that no value is present for ActRunId, not even an explicit nil
### GetFields

`func (o *Dataset) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Dataset) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Dataset) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Dataset) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *Dataset) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *Dataset) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetSchema

`func (o *Dataset) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Dataset) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Dataset) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Dataset) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Dataset) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Dataset) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetConsoleUrl

`func (o *Dataset) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *Dataset) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *Dataset) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.


### GetItemsPublicUrl

`func (o *Dataset) GetItemsPublicUrl() string`

GetItemsPublicUrl returns the ItemsPublicUrl field if non-nil, zero value otherwise.

### GetItemsPublicUrlOk

`func (o *Dataset) GetItemsPublicUrlOk() (*string, bool)`

GetItemsPublicUrlOk returns a tuple with the ItemsPublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPublicUrl

`func (o *Dataset) SetItemsPublicUrl(v string)`

SetItemsPublicUrl sets ItemsPublicUrl field to given value.

### HasItemsPublicUrl

`func (o *Dataset) HasItemsPublicUrl() bool`

HasItemsPublicUrl returns a boolean if a field has been set.

### GetUrlSigningSecretKey

`func (o *Dataset) GetUrlSigningSecretKey() string`

GetUrlSigningSecretKey returns the UrlSigningSecretKey field if non-nil, zero value otherwise.

### GetUrlSigningSecretKeyOk

`func (o *Dataset) GetUrlSigningSecretKeyOk() (*string, bool)`

GetUrlSigningSecretKeyOk returns a tuple with the UrlSigningSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSigningSecretKey

`func (o *Dataset) SetUrlSigningSecretKey(v string)`

SetUrlSigningSecretKey sets UrlSigningSecretKey field to given value.

### HasUrlSigningSecretKey

`func (o *Dataset) HasUrlSigningSecretKey() bool`

HasUrlSigningSecretKey returns a boolean if a field has been set.

### SetUrlSigningSecretKeyNil

`func (o *Dataset) SetUrlSigningSecretKeyNil(b bool)`

 SetUrlSigningSecretKeyNil sets the value for UrlSigningSecretKey to be an explicit nil

### UnsetUrlSigningSecretKey
`func (o *Dataset) UnsetUrlSigningSecretKey()`

UnsetUrlSigningSecretKey ensures that no value is present for UrlSigningSecretKey, not even an explicit nil
### GetGeneralAccess

`func (o *Dataset) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *Dataset) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *Dataset) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.

### HasGeneralAccess

`func (o *Dataset) HasGeneralAccess() bool`

HasGeneralAccess returns a boolean if a field has been set.

### GetStats

`func (o *Dataset) GetStats() DatasetStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Dataset) GetStatsOk() (*DatasetStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Dataset) SetStats(v DatasetStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Dataset) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


