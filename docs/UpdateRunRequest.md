# UpdateRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**IsStatusMessageTerminal** | Pointer to **bool** |  | [optional] 
**GeneralAccess** | Pointer to [**GeneralAccess**](GeneralAccess.md) |  | [optional] 

## Methods

### NewUpdateRunRequest

`func NewUpdateRunRequest() *UpdateRunRequest`

NewUpdateRunRequest instantiates a new UpdateRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRunRequestWithDefaults

`func NewUpdateRunRequestWithDefaults() *UpdateRunRequest`

NewUpdateRunRequestWithDefaults instantiates a new UpdateRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *UpdateRunRequest) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *UpdateRunRequest) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *UpdateRunRequest) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *UpdateRunRequest) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateRunRequest) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateRunRequest) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateRunRequest) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateRunRequest) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetIsStatusMessageTerminal

`func (o *UpdateRunRequest) GetIsStatusMessageTerminal() bool`

GetIsStatusMessageTerminal returns the IsStatusMessageTerminal field if non-nil, zero value otherwise.

### GetIsStatusMessageTerminalOk

`func (o *UpdateRunRequest) GetIsStatusMessageTerminalOk() (*bool, bool)`

GetIsStatusMessageTerminalOk returns a tuple with the IsStatusMessageTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatusMessageTerminal

`func (o *UpdateRunRequest) SetIsStatusMessageTerminal(v bool)`

SetIsStatusMessageTerminal sets IsStatusMessageTerminal field to given value.

### HasIsStatusMessageTerminal

`func (o *UpdateRunRequest) HasIsStatusMessageTerminal() bool`

HasIsStatusMessageTerminal returns a boolean if a field has been set.

### GetGeneralAccess

`func (o *UpdateRunRequest) GetGeneralAccess() GeneralAccess`

GetGeneralAccess returns the GeneralAccess field if non-nil, zero value otherwise.

### GetGeneralAccessOk

`func (o *UpdateRunRequest) GetGeneralAccessOk() (*GeneralAccess, bool)`

GetGeneralAccessOk returns a tuple with the GeneralAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAccess

`func (o *UpdateRunRequest) SetGeneralAccess(v GeneralAccess)`

SetGeneralAccess sets GeneralAccess field to given value.

### HasGeneralAccess

`func (o *UpdateRunRequest) HasGeneralAccess() bool`

HasGeneralAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


