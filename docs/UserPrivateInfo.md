# UserPrivateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Profile** | [**Profile**](Profile.md) |  | 
**Email** | **string** |  | 
**Proxy** | [**Proxy**](Proxy.md) |  | 
**Plan** | [**Plan**](Plan.md) |  | 
**EffectivePlatformFeatures** | [**EffectivePlatformFeatures**](EffectivePlatformFeatures.md) |  | 
**CreatedAt** | **time.Time** |  | 
**IsPaying** | **bool** |  | 

## Methods

### NewUserPrivateInfo

`func NewUserPrivateInfo(id string, username string, profile Profile, email string, proxy Proxy, plan Plan, effectivePlatformFeatures EffectivePlatformFeatures, createdAt time.Time, isPaying bool, ) *UserPrivateInfo`

NewUserPrivateInfo instantiates a new UserPrivateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPrivateInfoWithDefaults

`func NewUserPrivateInfoWithDefaults() *UserPrivateInfo`

NewUserPrivateInfoWithDefaults instantiates a new UserPrivateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPrivateInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPrivateInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPrivateInfo) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *UserPrivateInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPrivateInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPrivateInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetProfile

`func (o *UserPrivateInfo) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserPrivateInfo) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserPrivateInfo) SetProfile(v Profile)`

SetProfile sets Profile field to given value.


### GetEmail

`func (o *UserPrivateInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPrivateInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPrivateInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProxy

`func (o *UserPrivateInfo) GetProxy() Proxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *UserPrivateInfo) GetProxyOk() (*Proxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *UserPrivateInfo) SetProxy(v Proxy)`

SetProxy sets Proxy field to given value.


### GetPlan

`func (o *UserPrivateInfo) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UserPrivateInfo) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UserPrivateInfo) SetPlan(v Plan)`

SetPlan sets Plan field to given value.


### GetEffectivePlatformFeatures

`func (o *UserPrivateInfo) GetEffectivePlatformFeatures() EffectivePlatformFeatures`

GetEffectivePlatformFeatures returns the EffectivePlatformFeatures field if non-nil, zero value otherwise.

### GetEffectivePlatformFeaturesOk

`func (o *UserPrivateInfo) GetEffectivePlatformFeaturesOk() (*EffectivePlatformFeatures, bool)`

GetEffectivePlatformFeaturesOk returns a tuple with the EffectivePlatformFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePlatformFeatures

`func (o *UserPrivateInfo) SetEffectivePlatformFeatures(v EffectivePlatformFeatures)`

SetEffectivePlatformFeatures sets EffectivePlatformFeatures field to given value.


### GetCreatedAt

`func (o *UserPrivateInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserPrivateInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserPrivateInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsPaying

`func (o *UserPrivateInfo) GetIsPaying() bool`

GetIsPaying returns the IsPaying field if non-nil, zero value otherwise.

### GetIsPayingOk

`func (o *UserPrivateInfo) GetIsPayingOk() (*bool, bool)`

GetIsPayingOk returns a tuple with the IsPaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaying

`func (o *UserPrivateInfo) SetIsPaying(v bool)`

SetIsPaying sets IsPaying field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


