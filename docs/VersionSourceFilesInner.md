# VersionSourceFilesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**SourceCodeFileFormat**](SourceCodeFileFormat.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Name** | **string** | The folder path relative to the Actor&#39;s root directory. | 
**Folder** | **bool** | Always &#x60;true&#x60; for folders. Used to distinguish folders from files. | 

## Methods

### NewVersionSourceFilesInner

`func NewVersionSourceFilesInner(name string, folder bool, ) *VersionSourceFilesInner`

NewVersionSourceFilesInner instantiates a new VersionSourceFilesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionSourceFilesInnerWithDefaults

`func NewVersionSourceFilesInnerWithDefaults() *VersionSourceFilesInner`

NewVersionSourceFilesInnerWithDefaults instantiates a new VersionSourceFilesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *VersionSourceFilesInner) GetFormat() SourceCodeFileFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VersionSourceFilesInner) GetFormatOk() (*SourceCodeFileFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VersionSourceFilesInner) SetFormat(v SourceCodeFileFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VersionSourceFilesInner) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetContent

`func (o *VersionSourceFilesInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VersionSourceFilesInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VersionSourceFilesInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *VersionSourceFilesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetName

`func (o *VersionSourceFilesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionSourceFilesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionSourceFilesInner) SetName(v string)`

SetName sets Name field to given value.


### GetFolder

`func (o *VersionSourceFilesInner) GetFolder() bool`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VersionSourceFilesInner) GetFolderOk() (*bool, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VersionSourceFilesInner) SetFolder(v bool)`

SetFolder sets Folder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


