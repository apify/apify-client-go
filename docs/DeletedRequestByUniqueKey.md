# DeletedRequestByUniqueKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueKey** | **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | 
**Id** | Pointer to **string** | A unique identifier assigned to the request. | [optional] 

## Methods

### NewDeletedRequestByUniqueKey

`func NewDeletedRequestByUniqueKey(uniqueKey string, ) *DeletedRequestByUniqueKey`

NewDeletedRequestByUniqueKey instantiates a new DeletedRequestByUniqueKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedRequestByUniqueKeyWithDefaults

`func NewDeletedRequestByUniqueKeyWithDefaults() *DeletedRequestByUniqueKey`

NewDeletedRequestByUniqueKeyWithDefaults instantiates a new DeletedRequestByUniqueKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueKey

`func (o *DeletedRequestByUniqueKey) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *DeletedRequestByUniqueKey) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *DeletedRequestByUniqueKey) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.


### GetId

`func (o *DeletedRequestByUniqueKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedRequestByUniqueKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedRequestByUniqueKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeletedRequestByUniqueKey) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


