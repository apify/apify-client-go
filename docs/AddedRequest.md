# AddedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A unique identifier assigned to the request. | 
**UniqueKey** | **string** | A unique key used for request de-duplication. Requests with the same unique key are considered identical. | 
**WasAlreadyPresent** | **bool** | Indicates whether a request with the same unique key already existed in the request queue. If true, no new request was created. | 
**WasAlreadyHandled** | **bool** | Indicates whether a request with the same unique key has already been processed by the request queue. | 

## Methods

### NewAddedRequest

`func NewAddedRequest(requestId string, uniqueKey string, wasAlreadyPresent bool, wasAlreadyHandled bool, ) *AddedRequest`

NewAddedRequest instantiates a new AddedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedRequestWithDefaults

`func NewAddedRequestWithDefaults() *AddedRequest`

NewAddedRequestWithDefaults instantiates a new AddedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *AddedRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AddedRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AddedRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetUniqueKey

`func (o *AddedRequest) GetUniqueKey() string`

GetUniqueKey returns the UniqueKey field if non-nil, zero value otherwise.

### GetUniqueKeyOk

`func (o *AddedRequest) GetUniqueKeyOk() (*string, bool)`

GetUniqueKeyOk returns a tuple with the UniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueKey

`func (o *AddedRequest) SetUniqueKey(v string)`

SetUniqueKey sets UniqueKey field to given value.


### GetWasAlreadyPresent

`func (o *AddedRequest) GetWasAlreadyPresent() bool`

GetWasAlreadyPresent returns the WasAlreadyPresent field if non-nil, zero value otherwise.

### GetWasAlreadyPresentOk

`func (o *AddedRequest) GetWasAlreadyPresentOk() (*bool, bool)`

GetWasAlreadyPresentOk returns a tuple with the WasAlreadyPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasAlreadyPresent

`func (o *AddedRequest) SetWasAlreadyPresent(v bool)`

SetWasAlreadyPresent sets WasAlreadyPresent field to given value.


### GetWasAlreadyHandled

`func (o *AddedRequest) GetWasAlreadyHandled() bool`

GetWasAlreadyHandled returns the WasAlreadyHandled field if non-nil, zero value otherwise.

### GetWasAlreadyHandledOk

`func (o *AddedRequest) GetWasAlreadyHandledOk() (*bool, bool)`

GetWasAlreadyHandledOk returns a tuple with the WasAlreadyHandled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasAlreadyHandled

`func (o *AddedRequest) SetWasAlreadyHandled(v bool)`

SetWasAlreadyHandled sets WasAlreadyHandled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


