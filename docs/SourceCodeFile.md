# SourceCodeFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**SourceCodeFileFormat**](SourceCodeFileFormat.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewSourceCodeFile

`func NewSourceCodeFile(name string, ) *SourceCodeFile`

NewSourceCodeFile instantiates a new SourceCodeFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceCodeFileWithDefaults

`func NewSourceCodeFileWithDefaults() *SourceCodeFile`

NewSourceCodeFileWithDefaults instantiates a new SourceCodeFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *SourceCodeFile) GetFormat() SourceCodeFileFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SourceCodeFile) GetFormatOk() (*SourceCodeFileFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SourceCodeFile) SetFormat(v SourceCodeFileFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SourceCodeFile) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetContent

`func (o *SourceCodeFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SourceCodeFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SourceCodeFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SourceCodeFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetName

`func (o *SourceCodeFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceCodeFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceCodeFile) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


