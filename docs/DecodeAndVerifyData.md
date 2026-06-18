# DecodeAndVerifyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decoded** | **interface{}** |  | 
**EncodedByUserId** | **NullableString** |  | 
**IsVerifiedUser** | **bool** |  | 

## Methods

### NewDecodeAndVerifyData

`func NewDecodeAndVerifyData(decoded interface{}, encodedByUserId NullableString, isVerifiedUser bool, ) *DecodeAndVerifyData`

NewDecodeAndVerifyData instantiates a new DecodeAndVerifyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodeAndVerifyDataWithDefaults

`func NewDecodeAndVerifyDataWithDefaults() *DecodeAndVerifyData`

NewDecodeAndVerifyDataWithDefaults instantiates a new DecodeAndVerifyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecoded

`func (o *DecodeAndVerifyData) GetDecoded() interface{}`

GetDecoded returns the Decoded field if non-nil, zero value otherwise.

### GetDecodedOk

`func (o *DecodeAndVerifyData) GetDecodedOk() (*interface{}, bool)`

GetDecodedOk returns a tuple with the Decoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoded

`func (o *DecodeAndVerifyData) SetDecoded(v interface{})`

SetDecoded sets Decoded field to given value.


### SetDecodedNil

`func (o *DecodeAndVerifyData) SetDecodedNil(b bool)`

 SetDecodedNil sets the value for Decoded to be an explicit nil

### UnsetDecoded
`func (o *DecodeAndVerifyData) UnsetDecoded()`

UnsetDecoded ensures that no value is present for Decoded, not even an explicit nil
### GetEncodedByUserId

`func (o *DecodeAndVerifyData) GetEncodedByUserId() string`

GetEncodedByUserId returns the EncodedByUserId field if non-nil, zero value otherwise.

### GetEncodedByUserIdOk

`func (o *DecodeAndVerifyData) GetEncodedByUserIdOk() (*string, bool)`

GetEncodedByUserIdOk returns a tuple with the EncodedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedByUserId

`func (o *DecodeAndVerifyData) SetEncodedByUserId(v string)`

SetEncodedByUserId sets EncodedByUserId field to given value.


### SetEncodedByUserIdNil

`func (o *DecodeAndVerifyData) SetEncodedByUserIdNil(b bool)`

 SetEncodedByUserIdNil sets the value for EncodedByUserId to be an explicit nil

### UnsetEncodedByUserId
`func (o *DecodeAndVerifyData) UnsetEncodedByUserId()`

UnsetEncodedByUserId ensures that no value is present for EncodedByUserId, not even an explicit nil
### GetIsVerifiedUser

`func (o *DecodeAndVerifyData) GetIsVerifiedUser() bool`

GetIsVerifiedUser returns the IsVerifiedUser field if non-nil, zero value otherwise.

### GetIsVerifiedUserOk

`func (o *DecodeAndVerifyData) GetIsVerifiedUserOk() (*bool, bool)`

GetIsVerifiedUserOk returns a tuple with the IsVerifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerifiedUser

`func (o *DecodeAndVerifyData) SetIsVerifiedUser(v bool)`

SetIsVerifiedUser sets IsVerifiedUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


