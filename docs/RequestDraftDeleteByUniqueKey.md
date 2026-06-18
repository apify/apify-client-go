# RequestDraftDeleteByUniqueKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier assigned to the request. | [optional] 
**UniqueKey** | **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | 

## Methods

### NewRequestDraftDeleteByUniqueKey

`func NewRequestDraftDeleteByUniqueKey(uniqueKey string, ) *RequestDraftDeleteByUniqueKey`

NewRequestDraftDeleteByUniqueKey instantiates a new RequestDraftDeleteByUniqueKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDraftDeleteByUniqueKeyWithDefaults

`func NewRequestDraftDeleteByUniqueKeyWithDefaults() *RequestDraftDeleteByUniqueKey`

NewRequestDraftDeleteByUniqueKeyWithDefaults instantiates a new RequestDraftDeleteByUniqueKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestDraftDeleteByUniqueKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestDraftDeleteByUniqueKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestDraftDeleteByUniqueKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestDraftDeleteByUniqueKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqueKey

`func (o *RequestDraftDeleteByUniqueKey) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *RequestDraftDeleteByUniqueKey) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *RequestDraftDeleteByUniqueKey) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


