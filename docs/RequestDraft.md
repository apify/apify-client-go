# RequestDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier assigned to the request. | [optional] 
**UniqueKey** | **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | 
**Url** | **string** | The URL of the request. | 
**Method** | Pointer to [**HttpMethod**](HttpMethod.md) |  | [optional] 

## Methods

### NewRequestDraft

`func NewRequestDraft(uniqueKey string, url string, ) *RequestDraft`

NewRequestDraft instantiates a new RequestDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDraftWithDefaults

`func NewRequestDraftWithDefaults() *RequestDraft`

NewRequestDraftWithDefaults instantiates a new RequestDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestDraft) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestDraft) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestDraft) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestDraft) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqueKey

`func (o *RequestDraft) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *RequestDraft) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *RequestDraft) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.


### GetUrl

`func (o *RequestDraft) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestDraft) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestDraft) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *RequestDraft) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RequestDraft) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RequestDraft) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RequestDraft) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


