# ScheduleInvoked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Level** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewScheduleInvoked

`func NewScheduleInvoked(message string, level string, createdAt time.Time, ) *ScheduleInvoked`

NewScheduleInvoked instantiates a new ScheduleInvoked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleInvokedWithDefaults

`func NewScheduleInvokedWithDefaults() *ScheduleInvoked`

NewScheduleInvokedWithDefaults instantiates a new ScheduleInvoked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ScheduleInvoked) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScheduleInvoked) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScheduleInvoked) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLevel

`func (o *ScheduleInvoked) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ScheduleInvoked) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ScheduleInvoked) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetCreatedAt

`func (o *ScheduleInvoked) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduleInvoked) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduleInvoked) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


