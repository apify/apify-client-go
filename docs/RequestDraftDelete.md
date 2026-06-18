# RequestDraftDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier assigned to the request. | 
**UniqueKey** | **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | 

## Methods

### NewRequestDraftDelete

`func NewRequestDraftDelete(id string, uniqueKey string, ) *RequestDraftDelete`

NewRequestDraftDelete instantiates a new RequestDraftDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDraftDeleteWithDefaults

`func NewRequestDraftDeleteWithDefaults() *RequestDraftDelete`

NewRequestDraftDeleteWithDefaults instantiates a new RequestDraftDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestDraftDelete) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestDraftDelete) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestDraftDelete) SetId(v string)`

SetId sets Id field to given value.


### GetUniqueKey

`func (o *RequestDraftDelete) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *RequestDraftDelete) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *RequestDraftDelete) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


