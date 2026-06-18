# UserPublicInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 

## Methods

### NewUserPublicInfo

`func NewUserPublicInfo(username string, ) *UserPublicInfo`

NewUserPublicInfo instantiates a new UserPublicInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPublicInfoWithDefaults

`func NewUserPublicInfoWithDefaults() *UserPublicInfo`

NewUserPublicInfoWithDefaults instantiates a new UserPublicInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserPublicInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPublicInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPublicInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetProfile

`func (o *UserPublicInfo) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserPublicInfo) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserPublicInfo) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserPublicInfo) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


