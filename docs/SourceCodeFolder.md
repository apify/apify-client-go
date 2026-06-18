# SourceCodeFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The folder path relative to the Actor&#39;s root directory. | 
**Folder** | **bool** | Always &#x60;true&#x60; for folders. Used to distinguish folders from files. | 

## Methods

### NewSourceCodeFolder

`func NewSourceCodeFolder(name string, folder bool, ) *SourceCodeFolder`

NewSourceCodeFolder instantiates a new SourceCodeFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceCodeFolderWithDefaults

`func NewSourceCodeFolderWithDefaults() *SourceCodeFolder`

NewSourceCodeFolderWithDefaults instantiates a new SourceCodeFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceCodeFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceCodeFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceCodeFolder) SetName(v string)`

SetName sets Name field to given value.


### GetFolder

`func (o *SourceCodeFolder) GetFolder() bool`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *SourceCodeFolder) GetFolderOk() (*bool, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *SourceCodeFolder) SetFolder(v bool)`

SetFolder sets Folder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


