# ActorRunFailedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RunFailedErrorDetail**](RunFailedErrorDetail.md) |  | [optional] 

## Methods

### NewActorRunFailedError

`func NewActorRunFailedError() *ActorRunFailedError`

NewActorRunFailedError instantiates a new ActorRunFailedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorRunFailedErrorWithDefaults

`func NewActorRunFailedErrorWithDefaults() *ActorRunFailedError`

NewActorRunFailedErrorWithDefaults instantiates a new ActorRunFailedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ActorRunFailedError) GetError() RunFailedErrorDetail`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ActorRunFailedError) GetErrorOk() (*RunFailedErrorDetail, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ActorRunFailedError) SetError(v RunFailedErrorDetail)`

SetError sets Error field to given value.

### HasError

`func (o *ActorRunFailedError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


