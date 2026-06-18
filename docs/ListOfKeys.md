# ListOfKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]KeyValueStoreKey**](KeyValueStoreKey.md) |  | 
**Count** | **int32** |  | 
**Limit** | **int32** |  | 
**ExclusiveStartKey** | Pointer to **NullableString** |  | [optional] 
**IsTruncated** | **bool** |  | 
**NextExclusiveStartKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListOfKeys

`func NewListOfKeys(items []KeyValueStoreKey, count int32, limit int32, isTruncated bool, ) *ListOfKeys`

NewListOfKeys instantiates a new ListOfKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfKeysWithDefaults

`func NewListOfKeysWithDefaults() *ListOfKeys`

NewListOfKeysWithDefaults instantiates a new ListOfKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListOfKeys) GetItems() []KeyValueStoreKey`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfKeys) GetItemsOk() (*[]KeyValueStoreKey, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfKeys) SetItems(v []KeyValueStoreKey)`

SetItems sets Items field to given value.


### GetCount

`func (o *ListOfKeys) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOfKeys) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOfKeys) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLimit

`func (o *ListOfKeys) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOfKeys) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOfKeys) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetExclusiveStartKey

`func (o *ListOfKeys) GetExclusiveStartKey() string`

GetExclusiveStartKey returns the ExclusiveStartKey field if non-nil, zero value otherwise.

### GetExclusiveStartKeyOk

`func (o *ListOfKeys) GetExclusiveStartKeyOk() (*string, bool)`

GetExclusiveStartKeyOk returns a tuple with the ExclusiveStartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveStartKey

`func (o *ListOfKeys) SetExclusiveStartKey(v string)`

SetExclusiveStartKey sets ExclusiveStartKey field to given value.

### HasExclusiveStartKey

`func (o *ListOfKeys) HasExclusiveStartKey() bool`

HasExclusiveStartKey returns a boolean if a field has been set.

### SetExclusiveStartKeyNil

`func (o *ListOfKeys) SetExclusiveStartKeyNil(b bool)`

 SetExclusiveStartKeyNil sets the value for ExclusiveStartKey to be an explicit nil

### UnsetExclusiveStartKey
`func (o *ListOfKeys) UnsetExclusiveStartKey()`

UnsetExclusiveStartKey ensures that no value is present for ExclusiveStartKey, not even an explicit nil
### GetIsTruncated

`func (o *ListOfKeys) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *ListOfKeys) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *ListOfKeys) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.


### GetNextExclusiveStartKey

`func (o *ListOfKeys) GetNextExclusiveStartKey() string`

GetNextExclusiveStartKey returns the NextExclusiveStartKey field if non-nil, zero value otherwise.

### GetNextExclusiveStartKeyOk

`func (o *ListOfKeys) GetNextExclusiveStartKeyOk() (*string, bool)`

GetNextExclusiveStartKeyOk returns a tuple with the NextExclusiveStartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExclusiveStartKey

`func (o *ListOfKeys) SetNextExclusiveStartKey(v string)`

SetNextExclusiveStartKey sets NextExclusiveStartKey field to given value.

### HasNextExclusiveStartKey

`func (o *ListOfKeys) HasNextExclusiveStartKey() bool`

HasNextExclusiveStartKey returns a boolean if a field has been set.

### SetNextExclusiveStartKeyNil

`func (o *ListOfKeys) SetNextExclusiveStartKeyNil(b bool)`

 SetNextExclusiveStartKeyNil sets the value for NextExclusiveStartKey to be an explicit nil

### UnsetNextExclusiveStartKey
`func (o *ListOfKeys) UnsetNextExclusiveStartKey()`

UnsetNextExclusiveStartKey ensures that no value is present for NextExclusiveStartKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


