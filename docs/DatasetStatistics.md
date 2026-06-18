# DatasetStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldStatistics** | Pointer to [**map[string]DatasetFieldStatistics**](DatasetFieldStatistics.md) | When you configure the dataset [fields schema](https://docs.apify.com/platform/actors/development/actor-definition/dataset-schema/validation), we measure the statistics such as &#x60;min&#x60;, &#x60;max&#x60;, &#x60;nullCount&#x60; and &#x60;emptyCount&#x60; for each field. This property provides statistics for each field from dataset fields schema. &lt;br/&gt;&lt;/br&gt;See dataset field statistics [documentation](https://docs.apify.com/platform/actors/development/actor-definition/dataset-schema/validation#dataset-field-statistics) for more information. | [optional] 

## Methods

### NewDatasetStatistics

`func NewDatasetStatistics() *DatasetStatistics`

NewDatasetStatistics instantiates a new DatasetStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetStatisticsWithDefaults

`func NewDatasetStatisticsWithDefaults() *DatasetStatistics`

NewDatasetStatisticsWithDefaults instantiates a new DatasetStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldStatistics

`func (o *DatasetStatistics) GetFieldStatistics() map[string]DatasetFieldStatistics`

GetFieldStatistics returns the FieldStatistics field if non-nil, zero value otherwise.

### GetFieldStatisticsOk

`func (o *DatasetStatistics) GetFieldStatisticsOk() (*map[string]DatasetFieldStatistics, bool)`

GetFieldStatisticsOk returns a tuple with the FieldStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldStatistics

`func (o *DatasetStatistics) SetFieldStatistics(v map[string]DatasetFieldStatistics)`

SetFieldStatistics sets FieldStatistics field to given value.

### HasFieldStatistics

`func (o *DatasetStatistics) HasFieldStatistics() bool`

HasFieldStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


