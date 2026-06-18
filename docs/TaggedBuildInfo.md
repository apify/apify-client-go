# TaggedBuildInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildId** | Pointer to **string** | The ID of the build associated with this tag. | [optional] 
**BuildNumber** | Pointer to **string** | The build number/version string. | [optional] 
**BuildNumberInt** | Pointer to **int32** | The build number encoded as a single integer. | [optional] 
**FinishedAt** | Pointer to **NullableTime** | The timestamp when the build finished. | [optional] 

## Methods

### NewTaggedBuildInfo

`func NewTaggedBuildInfo() *TaggedBuildInfo`

NewTaggedBuildInfo instantiates a new TaggedBuildInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedBuildInfoWithDefaults

`func NewTaggedBuildInfoWithDefaults() *TaggedBuildInfo`

NewTaggedBuildInfoWithDefaults instantiates a new TaggedBuildInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildId

`func (o *TaggedBuildInfo) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *TaggedBuildInfo) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *TaggedBuildInfo) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *TaggedBuildInfo) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetBuildNumber

`func (o *TaggedBuildInfo) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *TaggedBuildInfo) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *TaggedBuildInfo) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *TaggedBuildInfo) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetBuildNumberInt

`func (o *TaggedBuildInfo) GetBuildNumberInt() int32`

GetBuildNumberInt returns the BuildNumberInt field if non-nil, zero value otherwise.

### GetBuildNumberIntOk

`func (o *TaggedBuildInfo) GetBuildNumberIntOk() (*int32, bool)`

GetBuildNumberIntOk returns a tuple with the BuildNumberInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumberInt

`func (o *TaggedBuildInfo) SetBuildNumberInt(v int32)`

SetBuildNumberInt sets BuildNumberInt field to given value.

### HasBuildNumberInt

`func (o *TaggedBuildInfo) HasBuildNumberInt() bool`

HasBuildNumberInt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *TaggedBuildInfo) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *TaggedBuildInfo) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *TaggedBuildInfo) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *TaggedBuildInfo) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *TaggedBuildInfo) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *TaggedBuildInfo) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


