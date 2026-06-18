# KeyValueStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**AccessedAt** | **time.Time** |  | 
**ActId** | Pointer to **NullableString** |  | [optional] 
**ActRunId** | Pointer to **NullableString** |  | [optional] 
**ConsoleUrl** | Pointer to **string** |  | [optional] 
**KeysPublicUrl** | Pointer to **string** | A public link to access keys of the key-value store directly. | [optional] 
**RecordsPublicUrl** | Pointer to **string** | A public link to access records of the key-value store directly. | [optional] 
**Schema** | Pointer to **map[string]interface{}** | Optional JSON schema describing the keys stored in the key-value store. | [optional] 
**UrlSigningSecretKey** | Pointer to **NullableString** | A secret key for generating signed public URLs. It is only provided to clients with WRITE permission for the key-value store. | [optional] 
**GeneralAccess** | Pointer to [**GeneralAccess**](GeneralAccess.md) |  | [optional] 
**Stats** | Pointer to [**KeyValueStoreStats**](KeyValueStoreStats.md) |  | [optional] 

## Methods

### NewKeyValueStore

`func NewKeyValueStore(id string, createdAt time.Time, modifiedAt time.Time, accessedAt time.Time, ) *KeyValueStore`

NewKeyValueStore instantiates a new KeyValueStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyValueStoreWithDefaults

`func NewKeyValueStoreWithDefaults() *KeyValueStore`

NewKeyValueStoreWithDefaults instantiates a new KeyValueStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyValueStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyValueStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyValueStore) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KeyValueStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyValueStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyValueStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyValueStore) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyValueStore) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyValueStore) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserId

`func (o *KeyValueStore) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KeyValueStore) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KeyValueStore) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KeyValueStore) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *KeyValueStore) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *KeyValueStore) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUsername

`func (o *KeyValueStore) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KeyValueStore) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KeyValueStore) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KeyValueStore) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *KeyValueStore) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *KeyValueStore) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetCreatedAt

`func (o *KeyValueStore) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeyValueStore) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeyValueStore) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *KeyValueStore) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *KeyValueStore) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *KeyValueStore) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetAccessedAt

`func (o *KeyValueStore) GetAccessedAt() time.Time`

GetAccessedAt returns the AccessedAt field if non-nil, zero value otherwise.

### GetAccessedAtOk

`func (o *KeyValueStore) GetAccessedAtOk() (*time.Time, bool)`

GetAccessedAtOk returns a tuple with the AccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedAt

`func (o *KeyValueStore) SetAccessedAt(v time.Time)`

SetAccessedAt sets AccessedAt field to given value.


### GetActId

`func (o *KeyValueStore) GetActId() string`

GetActId returns the ActId field if non-nil, zero value otherwise.

### GetActIdOk

`func (o *KeyValueStore) GetActIdOk() (*string, bool)`

GetActIdOk returns a tuple with the ActId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActId

`func (o *KeyValueStore) SetActId(v string)`

SetActId sets ActId field to given value.

### HasActId

`func (o *KeyValueStore) HasActId() bool`

HasActId returns a boolean if a field has been set.

### SetActIdNil

`func (o *KeyValueStore) SetActIdNil(b bool)`

 SetActIdNil sets the value for ActId to be an explicit nil

### UnsetActId
`func (o *KeyValueStore) UnsetActId()`

UnsetActId ensures that no value is present for ActId, not even an explicit nil
### GetActRunId

`func (o *KeyValueStore) GetActRunId() string`

GetActRunId returns the ActRunId field if non-nil, zero value otherwise.

### GetActRunIdOk

`func (o *KeyValueStore) GetActRunIdOk() (*string, bool)`

GetActRunIdOk returns a tuple with the ActRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActRunId

`func (o *KeyValueStore) SetActRunId(v string)`

SetActRunId sets ActRunId field to given value.

### HasActRunId

`func (o *KeyValueStore) HasActRunId() bool`

HasActRunId returns a boolean if a field has been set.

### SetActRunIdNil

`func (o *KeyValueStore) SetActRunIdNil(b bool)`

 SetActRunIdNil sets the value for ActRunId to be an explicit nil

### UnsetActRunId
`func (o *KeyValueStore) UnsetActRunId()`

UnsetActRunId ensures that no value is present for ActRunId, not even an explicit nil
### GetConsoleUrl

`func (o *KeyValueStore) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *KeyValueStore) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *KeyValueStore) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *KeyValueStore) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetKeysPublicUrl

`func (o *KeyValueStore) GetKeysPublicUrl() string`

GetKeysPublicUrl returns the KeysPublicUrl field if non-nil, zero value otherwise.

### GetKeysPublicUrlOk

`func (o *KeyValueStore) GetKeysPublicUrlOk() (*string, bool)`

GetKeysPublicUrlOk returns a tuple with the KeysPublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysPublicUrl

`func (o *KeyValueStore) SetKeysPublicUrl(v string)`

SetKeysPublicUrl sets KeysPublicUrl field to given value.

### HasKeysPublicUrl

`func (o *KeyValueStore) HasKeysPublicUrl() bool`

HasKeysPublicUrl returns a boolean if a field has been set.

### GetRecordsPublicUrl

`func (o *KeyValueStore) GetRecordsPublicUrl() string`

GetRecordsPublicUrl returns the RecordsPublicUrl field if non-nil, zero value otherwise.

### GetRecordsPublicUrlOk

`func (o *KeyValueStore) GetRecordsPublicUrlOk() (*string, bool)`

GetRecordsPublicUrlOk returns a tuple with the RecordsPublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsPublicUrl

`func (o *KeyValueStore) SetRecordsPublicUrl(v string)`

SetRecordsPublicUrl sets RecordsPublicUrl field to given value.

### HasRecordsPublicUrl

`func (o *KeyValueStore) HasRecordsPublicUrl() bool`

HasRecordsPublicUrl returns a boolean if a field has been set.

### GetSchema

`func (o *KeyValueStore) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *KeyValueStore) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *KeyValueStore) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *KeyValueStore) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *KeyValueStore) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *KeyValueStore) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetUrlSigningSecretKey

`func (o *KeyValueStore) GetUrlSigningSecretKey() string`

GetUrlSigningSecretKey returns the UrlSigningSecretKey field if non-nil, zero value otherwise.

### GetUrlSigningSecretKeyOk

`func (o *KeyValueStore) GetUrlSigningSecretKeyOk() (*string, bool)`

GetUrlSigningSecretKeyOk returns a tuple with the UrlSigningSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSigningSecretKey

`func (o *KeyValueStore) SetUrlSigningSecretKey(v string)`

SetUrlSigningSecretKey sets UrlSigningSecretKey field to given value.

### HasUrlSigningSecretKey

`func (o *KeyValueStore) HasUrlSigningSecretKey() bool`

HasUrlSigningSecretKey returns a boolean if a field has been set.

### SetUrlSigningSecretKeyNil

`func (o *KeyValueStore) SetUrlSigningSecretKeyNil(b bool)`

 SetUrlSigningSecretKeyNil sets the value for UrlSigningSecretKey to be an explicit nil

### UnsetUrlSigningSecretKey
`func (o *KeyValueStore) UnsetUrlSigningSecretKey()`

UnsetUrlSigningSecretKey ensures that no value is present for UrlSigningSecretKey, not even an explicit nil
### GetGeneralAccess

`func (o *KeyValueStore) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *KeyValueStore) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *KeyValueStore) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.

### HasGeneralAccess

`func (o *KeyValueStore) HasGeneralAccess() bool`

HasGeneralAccess returns a boolean if a field has been set.

### GetStats

`func (o *KeyValueStore) GetStats() KeyValueStoreStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *KeyValueStore) GetStatsOk() (*KeyValueStoreStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *KeyValueStore) SetStats(v KeyValueStoreStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *KeyValueStore) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


