# KeyValueStoreKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Size** | **int32** |  | 
**RecordPublicUrl** | **string** | A public link to access this record directly. | 

## Methods

### NewKeyValueStoreKey

`func NewKeyValueStoreKey(key string, size int32, recordPublicUrl string, ) *KeyValueStoreKey`

NewKeyValueStoreKey instantiates a new KeyValueStoreKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyValueStoreKeyWithDefaults

`func NewKeyValueStoreKeyWithDefaults() *KeyValueStoreKey`

NewKeyValueStoreKeyWithDefaults instantiates a new KeyValueStoreKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KeyValueStoreKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyValueStoreKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyValueStoreKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetSize

`func (o *KeyValueStoreKey) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *KeyValueStoreKey) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *KeyValueStoreKey) SetSize(v int32)`

SetSize sets Size field to given value.


### GetRecordPublicUrl

`func (o *KeyValueStoreKey) GetRecordPublicUrl() string`

GetRecordPublicUrl returns the RecordPublicUrl field if non-nil, zero value otherwise.

### GetRecordPublicUrlOk

`func (o *KeyValueStoreKey) GetRecordPublicUrlOk() (*string, bool)`

GetRecordPublicUrlOk returns a tuple with the RecordPublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordPublicUrl

`func (o *KeyValueStoreKey) SetRecordPublicUrl(v string)`

SetRecordPublicUrl sets RecordPublicUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


