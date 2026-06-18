# Proxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**Groups** | [**[]ProxyGroup**](ProxyGroup.md) |  | 

## Methods

### NewProxy

`func NewProxy(password string, groups []ProxyGroup, ) *Proxy`

NewProxy instantiates a new Proxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyWithDefaults

`func NewProxyWithDefaults() *Proxy`

NewProxyWithDefaults instantiates a new Proxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *Proxy) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Proxy) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Proxy) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetGroups

`func (o *Proxy) GetGroups() []ProxyGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Proxy) GetGroupsOk() (*[]ProxyGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Proxy) SetGroups(v []ProxyGroup)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


