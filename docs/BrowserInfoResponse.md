# BrowserInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | HTTP method of the request. | 
**ClientIp** | **NullableString** | IP address of the client. | 
**CountryCode** | **NullableString** | Two-letter country code resolved from the client IP address. | 
**BodyLength** | **int32** | Length of the request body in bytes. | 
**Headers** | Pointer to [**map[string]BrowserInfoResponseHeadersValue**](BrowserInfoResponseHeadersValue.md) | Request headers. Omitted when &#x60;skipHeaders&#x3D;true&#x60;.  | [optional] 
**RawHeaders** | Pointer to **[]string** | Raw request headers as a flat list of alternating name/value strings. Included only when &#x60;rawHeaders&#x3D;true&#x60;.  | [optional] 

## Methods

### NewBrowserInfoResponse

`func NewBrowserInfoResponse(method string, clientIp NullableString, countryCode NullableString, bodyLength int32, ) *BrowserInfoResponse`

NewBrowserInfoResponse instantiates a new BrowserInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserInfoResponseWithDefaults

`func NewBrowserInfoResponseWithDefaults() *BrowserInfoResponse`

NewBrowserInfoResponseWithDefaults instantiates a new BrowserInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *BrowserInfoResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *BrowserInfoResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *BrowserInfoResponse) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetClientIp

`func (o *BrowserInfoResponse) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *BrowserInfoResponse) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *BrowserInfoResponse) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.


### SetClientIpNil

`func (o *BrowserInfoResponse) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *BrowserInfoResponse) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetCountryCode

`func (o *BrowserInfoResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BrowserInfoResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BrowserInfoResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *BrowserInfoResponse) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *BrowserInfoResponse) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetBodyLength

`func (o *BrowserInfoResponse) GetBodyLength() int32`

GetBodyLength returns the BodyLength field if non-nil, zero value otherwise.

### GetBodyLengthOk

`func (o *BrowserInfoResponse) GetBodyLengthOk() (*int32, bool)`

GetBodyLengthOk returns a tuple with the BodyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyLength

`func (o *BrowserInfoResponse) SetBodyLength(v int32)`

SetBodyLength sets BodyLength field to given value.


### GetHeaders

`func (o *BrowserInfoResponse) GetHeaders() map[string]BrowserInfoResponseHeadersValue`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BrowserInfoResponse) GetHeadersOk() (*map[string]BrowserInfoResponseHeadersValue, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BrowserInfoResponse) SetHeaders(v map[string]BrowserInfoResponseHeadersValue)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BrowserInfoResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRawHeaders

`func (o *BrowserInfoResponse) GetRawHeaders() []string`

GetRawHeaders returns the RawHeaders field if non-nil, zero value otherwise.

### GetRawHeadersOk

`func (o *BrowserInfoResponse) GetRawHeadersOk() (*[]string, bool)`

GetRawHeadersOk returns a tuple with the RawHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawHeaders

`func (o *BrowserInfoResponse) SetRawHeaders(v []string)`

SetRawHeaders sets RawHeaders field to given value.

### HasRawHeaders

`func (o *BrowserInfoResponse) HasRawHeaders() bool`

HasRawHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


