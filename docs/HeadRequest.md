# HeadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier assigned to the request. | 
**UniqueKey** | **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | 
**Url** | **string** | The URL of the request. | 
**Method** | Pointer to [**HttpMethod**](HttpMethod.md) |  | [optional] 
**RetryCount** | Pointer to **int32** | The number of times this request has been retried. | [optional] 

## Methods

### NewHeadRequest

`func NewHeadRequest(id string, uniqueKey string, url string, ) *HeadRequest`

NewHeadRequest instantiates a new HeadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadRequestWithDefaults

`func NewHeadRequestWithDefaults() *HeadRequest`

NewHeadRequestWithDefaults instantiates a new HeadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeadRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeadRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeadRequest) SetId(v string)`

SetId sets Id field to given value.


### GetUniqueKey

`func (o *HeadRequest) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *HeadRequest) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *HeadRequest) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.


### GetUrl

`func (o *HeadRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HeadRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HeadRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *HeadRequest) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HeadRequest) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HeadRequest) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HeadRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRetryCount

`func (o *HeadRequest) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *HeadRequest) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *HeadRequest) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *HeadRequest) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


