# ListOfEnvVars

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Items** | [**[]EnvVar**](EnvVar.md) |  | 

## Methods

### NewListOfEnvVars

`func NewListOfEnvVars(total int32, items []EnvVar, ) *ListOfEnvVars`

NewListOfEnvVars instantiates a new ListOfEnvVars object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOfEnvVarsWithDefaults

`func NewListOfEnvVarsWithDefaults() *ListOfEnvVars`

NewListOfEnvVarsWithDefaults instantiates a new ListOfEnvVars object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListOfEnvVars) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListOfEnvVars) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListOfEnvVars) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListOfEnvVars) GetItems() []EnvVar`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOfEnvVars) GetItemsOk() (*[]EnvVar, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOfEnvVars) SetItems(v []EnvVar)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


