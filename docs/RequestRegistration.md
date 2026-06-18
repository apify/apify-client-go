# RequestRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A unique identifier assigned to the request. | 
**WasAlreadyPresent** | **bool** | Indicates whether a request with the same unique key already existed in the request queue. If true, no new request was created. | 
**WasAlreadyHandled** | **bool** | Indicates whether a request with the same unique key has already been processed by the request queue. | 

## Methods

### NewRequestRegistration

`func NewRequestRegistration(requestId string, wasAlreadyPresent bool, wasAlreadyHandled bool, ) *RequestRegistration`

NewRequestRegistration instantiates a new RequestRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRegistrationWithDefaults

`func NewRequestRegistrationWithDefaults() *RequestRegistration`

NewRequestRegistrationWithDefaults instantiates a new RequestRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *RequestRegistration) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RequestRegistration) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RequestRegistration) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetWasAlreadyPresent

`func (o *RequestRegistration) GetWasAlreadyPresent() bool`

GetWasAlreadyPresent returns the WasAlreadyPresent field if non-nil, zero value otherwise.

### GetWasAlreadyPresentOk

`func (o *RequestRegistration) GetWasAlreadyPresentOk() (*bool, bool)`

GetWasAlreadyPresentOk returns a tuple with the WasAlreadyPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasAlreadyPresent

`func (o *RequestRegistration) SetWasAlreadyPresent(v bool)`

SetWasAlreadyPresent sets WasAlreadyPresent field to given value.


### GetWasAlreadyHandled

`func (o *RequestRegistration) GetWasAlreadyHandled() bool`

GetWasAlreadyHandled returns the WasAlreadyHandled field if non-nil, zero value otherwise.

### GetWasAlreadyHandledOk

`func (o *RequestRegistration) GetWasAlreadyHandledOk() (*bool, bool)`

GetWasAlreadyHandledOk returns a tuple with the WasAlreadyHandled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasAlreadyHandled

`func (o *RequestRegistration) SetWasAlreadyHandled(v bool)`

SetWasAlreadyHandled sets WasAlreadyHandled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


