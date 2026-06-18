# BuildsMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | [**RunOrigin**](RunOrigin.md) |  | 
**ClientIp** | Pointer to **string** | IP address of the client that started the build. | [optional] 
**UserAgent** | Pointer to **string** | User agent of the client that started the build. | [optional] 

## Methods

### NewBuildsMeta

`func NewBuildsMeta(origin RunOrigin, ) *BuildsMeta`

NewBuildsMeta instantiates a new BuildsMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildsMetaWithDefaults

`func NewBuildsMetaWithDefaults() *BuildsMeta`

NewBuildsMetaWithDefaults instantiates a new BuildsMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *BuildsMeta) GetOrigin() RunOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BuildsMeta) GetOriginOk() (*RunOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BuildsMeta) SetOrigin(v RunOrigin)`

SetOrigin sets Origin field to given value.


### GetClientIp

`func (o *BuildsMeta) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *BuildsMeta) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *BuildsMeta) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *BuildsMeta) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *BuildsMeta) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *BuildsMeta) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *BuildsMeta) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *BuildsMeta) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


