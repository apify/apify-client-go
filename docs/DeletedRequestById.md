# DeletedRequestById

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueKey** | Pointer to **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | [optional] 
**Id** | **string** | A unique identifier assigned to the request. | 

## Methods

### NewDeletedRequestById

`func NewDeletedRequestById(id string, ) *DeletedRequestById`

NewDeletedRequestById instantiates a new DeletedRequestById object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedRequestByIdWithDefaults

`func NewDeletedRequestByIdWithDefaults() *DeletedRequestById`

NewDeletedRequestByIdWithDefaults instantiates a new DeletedRequestById object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueKey

`func (o *DeletedRequestById) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *DeletedRequestById) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *DeletedRequestById) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.

### HasUniqueKey

`func (o *DeletedRequestById) HasUniqueKey() bool`

HasUniqueKey returns a boolean if a field has been set.

### GetId

`func (o *DeletedRequestById) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedRequestById) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedRequestById) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


