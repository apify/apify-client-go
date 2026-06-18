# RunUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACTOR_COMPUTE_UNITS** | Pointer to **NullableFloat32** |  | [optional] 
**DATASET_READS** | Pointer to **NullableInt32** |  | [optional] 
**DATASET_WRITES** | Pointer to **NullableInt32** |  | [optional] 
**KEY_VALUE_STORE_READS** | Pointer to **NullableInt32** |  | [optional] 
**KEY_VALUE_STORE_WRITES** | Pointer to **NullableInt32** |  | [optional] 
**KEY_VALUE_STORE_LISTS** | Pointer to **NullableInt32** |  | [optional] 
**REQUEST_QUEUE_READS** | Pointer to **NullableInt32** |  | [optional] 
**REQUEST_QUEUE_WRITES** | Pointer to **NullableInt32** |  | [optional] 
**DATA_TRANSFER_INTERNAL_GBYTES** | Pointer to **NullableFloat32** |  | [optional] 
**DATA_TRANSFER_EXTERNAL_GBYTES** | Pointer to **NullableFloat32** |  | [optional] 
**PROXY_RESIDENTIAL_TRANSFER_GBYTES** | Pointer to **NullableFloat32** |  | [optional] 
**PROXY_SERPS** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRunUsage

`func NewRunUsage() *RunUsage`

NewRunUsage instantiates a new RunUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunUsageWithDefaults

`func NewRunUsageWithDefaults() *RunUsage`

NewRunUsageWithDefaults instantiates a new RunUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACTOR_COMPUTE_UNITS

`func (o *RunUsage) GetACTOR_COMPUTE_UNITS() float32`

GetACTOR_COMPUTE_UNITS returns the ACTOR_COMPUTE_UNITS field if non-nil, zero value otherwise.

### GetACTOR_COMPUTE_UNITSOk

`func (o *RunUsage) GetACTOR_COMPUTE_UNITSOk() (*float32, bool)`

GetACTOR_COMPUTE_UNITSOk returns a tuple with the ACTOR_COMPUTE_UNITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACTOR_COMPUTE_UNITS

`func (o *RunUsage) SetACTOR_COMPUTE_UNITS(v float32)`

SetACTOR_COMPUTE_UNITS sets ACTOR_COMPUTE_UNITS field to given value.

### HasACTOR_COMPUTE_UNITS

`func (o *RunUsage) HasACTOR_COMPUTE_UNITS() bool`

HasACTOR_COMPUTE_UNITS returns a boolean if a field has been set.

### SetACTOR_COMPUTE_UNITSNil

`func (o *RunUsage) SetACTOR_COMPUTE_UNITSNil(b bool)`

 SetACTOR_COMPUTE_UNITSNil sets the value for ACTOR_COMPUTE_UNITS to be an explicit nil

### UnsetACTOR_COMPUTE_UNITS
`func (o *RunUsage) UnsetACTOR_COMPUTE_UNITS()`

UnsetACTOR_COMPUTE_UNITS ensures that no value is present for ACTOR_COMPUTE_UNITS, not even an explicit nil
### GetDATASET_READS

`func (o *RunUsage) GetDATASET_READS() int32`

GetDATASET_READS returns the DATASET_READS field if non-nil, zero value otherwise.

### GetDATASET_READSOk

`func (o *RunUsage) GetDATASET_READSOk() (*int32, bool)`

GetDATASET_READSOk returns a tuple with the DATASET_READS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATASET_READS

`func (o *RunUsage) SetDATASET_READS(v int32)`

SetDATASET_READS sets DATASET_READS field to given value.

### HasDATASET_READS

`func (o *RunUsage) HasDATASET_READS() bool`

HasDATASET_READS returns a boolean if a field has been set.

### SetDATASET_READSNil

`func (o *RunUsage) SetDATASET_READSNil(b bool)`

 SetDATASET_READSNil sets the value for DATASET_READS to be an explicit nil

### UnsetDATASET_READS
`func (o *RunUsage) UnsetDATASET_READS()`

UnsetDATASET_READS ensures that no value is present for DATASET_READS, not even an explicit nil
### GetDATASET_WRITES

`func (o *RunUsage) GetDATASET_WRITES() int32`

GetDATASET_WRITES returns the DATASET_WRITES field if non-nil, zero value otherwise.

### GetDATASET_WRITESOk

`func (o *RunUsage) GetDATASET_WRITESOk() (*int32, bool)`

GetDATASET_WRITESOk returns a tuple with the DATASET_WRITES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATASET_WRITES

`func (o *RunUsage) SetDATASET_WRITES(v int32)`

SetDATASET_WRITES sets DATASET_WRITES field to given value.

### HasDATASET_WRITES

`func (o *RunUsage) HasDATASET_WRITES() bool`

HasDATASET_WRITES returns a boolean if a field has been set.

### SetDATASET_WRITESNil

`func (o *RunUsage) SetDATASET_WRITESNil(b bool)`

 SetDATASET_WRITESNil sets the value for DATASET_WRITES to be an explicit nil

### UnsetDATASET_WRITES
`func (o *RunUsage) UnsetDATASET_WRITES()`

UnsetDATASET_WRITES ensures that no value is present for DATASET_WRITES, not even an explicit nil
### GetKEY_VALUE_STORE_READS

`func (o *RunUsage) GetKEY_VALUE_STORE_READS() int32`

GetKEY_VALUE_STORE_READS returns the KEY_VALUE_STORE_READS field if non-nil, zero value otherwise.

### GetKEY_VALUE_STORE_READSOk

`func (o *RunUsage) GetKEY_VALUE_STORE_READSOk() (*int32, bool)`

GetKEY_VALUE_STORE_READSOk returns a tuple with the KEY_VALUE_STORE_READS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKEY_VALUE_STORE_READS

`func (o *RunUsage) SetKEY_VALUE_STORE_READS(v int32)`

SetKEY_VALUE_STORE_READS sets KEY_VALUE_STORE_READS field to given value.

### HasKEY_VALUE_STORE_READS

`func (o *RunUsage) HasKEY_VALUE_STORE_READS() bool`

HasKEY_VALUE_STORE_READS returns a boolean if a field has been set.

### SetKEY_VALUE_STORE_READSNil

`func (o *RunUsage) SetKEY_VALUE_STORE_READSNil(b bool)`

 SetKEY_VALUE_STORE_READSNil sets the value for KEY_VALUE_STORE_READS to be an explicit nil

### UnsetKEY_VALUE_STORE_READS
`func (o *RunUsage) UnsetKEY_VALUE_STORE_READS()`

UnsetKEY_VALUE_STORE_READS ensures that no value is present for KEY_VALUE_STORE_READS, not even an explicit nil
### GetKEY_VALUE_STORE_WRITES

`func (o *RunUsage) GetKEY_VALUE_STORE_WRITES() int32`

GetKEY_VALUE_STORE_WRITES returns the KEY_VALUE_STORE_WRITES field if non-nil, zero value otherwise.

### GetKEY_VALUE_STORE_WRITESOk

`func (o *RunUsage) GetKEY_VALUE_STORE_WRITESOk() (*int32, bool)`

GetKEY_VALUE_STORE_WRITESOk returns a tuple with the KEY_VALUE_STORE_WRITES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKEY_VALUE_STORE_WRITES

`func (o *RunUsage) SetKEY_VALUE_STORE_WRITES(v int32)`

SetKEY_VALUE_STORE_WRITES sets KEY_VALUE_STORE_WRITES field to given value.

### HasKEY_VALUE_STORE_WRITES

`func (o *RunUsage) HasKEY_VALUE_STORE_WRITES() bool`

HasKEY_VALUE_STORE_WRITES returns a boolean if a field has been set.

### SetKEY_VALUE_STORE_WRITESNil

`func (o *RunUsage) SetKEY_VALUE_STORE_WRITESNil(b bool)`

 SetKEY_VALUE_STORE_WRITESNil sets the value for KEY_VALUE_STORE_WRITES to be an explicit nil

### UnsetKEY_VALUE_STORE_WRITES
`func (o *RunUsage) UnsetKEY_VALUE_STORE_WRITES()`

UnsetKEY_VALUE_STORE_WRITES ensures that no value is present for KEY_VALUE_STORE_WRITES, not even an explicit nil
### GetKEY_VALUE_STORE_LISTS

`func (o *RunUsage) GetKEY_VALUE_STORE_LISTS() int32`

GetKEY_VALUE_STORE_LISTS returns the KEY_VALUE_STORE_LISTS field if non-nil, zero value otherwise.

### GetKEY_VALUE_STORE_LISTSOk

`func (o *RunUsage) GetKEY_VALUE_STORE_LISTSOk() (*int32, bool)`

GetKEY_VALUE_STORE_LISTSOk returns a tuple with the KEY_VALUE_STORE_LISTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKEY_VALUE_STORE_LISTS

`func (o *RunUsage) SetKEY_VALUE_STORE_LISTS(v int32)`

SetKEY_VALUE_STORE_LISTS sets KEY_VALUE_STORE_LISTS field to given value.

### HasKEY_VALUE_STORE_LISTS

`func (o *RunUsage) HasKEY_VALUE_STORE_LISTS() bool`

HasKEY_VALUE_STORE_LISTS returns a boolean if a field has been set.

### SetKEY_VALUE_STORE_LISTSNil

`func (o *RunUsage) SetKEY_VALUE_STORE_LISTSNil(b bool)`

 SetKEY_VALUE_STORE_LISTSNil sets the value for KEY_VALUE_STORE_LISTS to be an explicit nil

### UnsetKEY_VALUE_STORE_LISTS
`func (o *RunUsage) UnsetKEY_VALUE_STORE_LISTS()`

UnsetKEY_VALUE_STORE_LISTS ensures that no value is present for KEY_VALUE_STORE_LISTS, not even an explicit nil
### GetREQUEST_QUEUE_READS

`func (o *RunUsage) GetREQUEST_QUEUE_READS() int32`

GetREQUEST_QUEUE_READS returns the REQUEST_QUEUE_READS field if non-nil, zero value otherwise.

### GetREQUEST_QUEUE_READSOk

`func (o *RunUsage) GetREQUEST_QUEUE_READSOk() (*int32, bool)`

GetREQUEST_QUEUE_READSOk returns a tuple with the REQUEST_QUEUE_READS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREQUEST_QUEUE_READS

`func (o *RunUsage) SetREQUEST_QUEUE_READS(v int32)`

SetREQUEST_QUEUE_READS sets REQUEST_QUEUE_READS field to given value.

### HasREQUEST_QUEUE_READS

`func (o *RunUsage) HasREQUEST_QUEUE_READS() bool`

HasREQUEST_QUEUE_READS returns a boolean if a field has been set.

### SetREQUEST_QUEUE_READSNil

`func (o *RunUsage) SetREQUEST_QUEUE_READSNil(b bool)`

 SetREQUEST_QUEUE_READSNil sets the value for REQUEST_QUEUE_READS to be an explicit nil

### UnsetREQUEST_QUEUE_READS
`func (o *RunUsage) UnsetREQUEST_QUEUE_READS()`

UnsetREQUEST_QUEUE_READS ensures that no value is present for REQUEST_QUEUE_READS, not even an explicit nil
### GetREQUEST_QUEUE_WRITES

`func (o *RunUsage) GetREQUEST_QUEUE_WRITES() int32`

GetREQUEST_QUEUE_WRITES returns the REQUEST_QUEUE_WRITES field if non-nil, zero value otherwise.

### GetREQUEST_QUEUE_WRITESOk

`func (o *RunUsage) GetREQUEST_QUEUE_WRITESOk() (*int32, bool)`

GetREQUEST_QUEUE_WRITESOk returns a tuple with the REQUEST_QUEUE_WRITES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREQUEST_QUEUE_WRITES

`func (o *RunUsage) SetREQUEST_QUEUE_WRITES(v int32)`

SetREQUEST_QUEUE_WRITES sets REQUEST_QUEUE_WRITES field to given value.

### HasREQUEST_QUEUE_WRITES

`func (o *RunUsage) HasREQUEST_QUEUE_WRITES() bool`

HasREQUEST_QUEUE_WRITES returns a boolean if a field has been set.

### SetREQUEST_QUEUE_WRITESNil

`func (o *RunUsage) SetREQUEST_QUEUE_WRITESNil(b bool)`

 SetREQUEST_QUEUE_WRITESNil sets the value for REQUEST_QUEUE_WRITES to be an explicit nil

### UnsetREQUEST_QUEUE_WRITES
`func (o *RunUsage) UnsetREQUEST_QUEUE_WRITES()`

UnsetREQUEST_QUEUE_WRITES ensures that no value is present for REQUEST_QUEUE_WRITES, not even an explicit nil
### GetDATA_TRANSFER_INTERNAL_GBYTES

`func (o *RunUsage) GetDATA_TRANSFER_INTERNAL_GBYTES() float32`

GetDATA_TRANSFER_INTERNAL_GBYTES returns the DATA_TRANSFER_INTERNAL_GBYTES field if non-nil, zero value otherwise.

### GetDATA_TRANSFER_INTERNAL_GBYTESOk

`func (o *RunUsage) GetDATA_TRANSFER_INTERNAL_GBYTESOk() (*float32, bool)`

GetDATA_TRANSFER_INTERNAL_GBYTESOk returns a tuple with the DATA_TRANSFER_INTERNAL_GBYTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_TRANSFER_INTERNAL_GBYTES

`func (o *RunUsage) SetDATA_TRANSFER_INTERNAL_GBYTES(v float32)`

SetDATA_TRANSFER_INTERNAL_GBYTES sets DATA_TRANSFER_INTERNAL_GBYTES field to given value.

### HasDATA_TRANSFER_INTERNAL_GBYTES

`func (o *RunUsage) HasDATA_TRANSFER_INTERNAL_GBYTES() bool`

HasDATA_TRANSFER_INTERNAL_GBYTES returns a boolean if a field has been set.

### SetDATA_TRANSFER_INTERNAL_GBYTESNil

`func (o *RunUsage) SetDATA_TRANSFER_INTERNAL_GBYTESNil(b bool)`

 SetDATA_TRANSFER_INTERNAL_GBYTESNil sets the value for DATA_TRANSFER_INTERNAL_GBYTES to be an explicit nil

### UnsetDATA_TRANSFER_INTERNAL_GBYTES
`func (o *RunUsage) UnsetDATA_TRANSFER_INTERNAL_GBYTES()`

UnsetDATA_TRANSFER_INTERNAL_GBYTES ensures that no value is present for DATA_TRANSFER_INTERNAL_GBYTES, not even an explicit nil
### GetDATA_TRANSFER_EXTERNAL_GBYTES

`func (o *RunUsage) GetDATA_TRANSFER_EXTERNAL_GBYTES() float32`

GetDATA_TRANSFER_EXTERNAL_GBYTES returns the DATA_TRANSFER_EXTERNAL_GBYTES field if non-nil, zero value otherwise.

### GetDATA_TRANSFER_EXTERNAL_GBYTESOk

`func (o *RunUsage) GetDATA_TRANSFER_EXTERNAL_GBYTESOk() (*float32, bool)`

GetDATA_TRANSFER_EXTERNAL_GBYTESOk returns a tuple with the DATA_TRANSFER_EXTERNAL_GBYTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_TRANSFER_EXTERNAL_GBYTES

`func (o *RunUsage) SetDATA_TRANSFER_EXTERNAL_GBYTES(v float32)`

SetDATA_TRANSFER_EXTERNAL_GBYTES sets DATA_TRANSFER_EXTERNAL_GBYTES field to given value.

### HasDATA_TRANSFER_EXTERNAL_GBYTES

`func (o *RunUsage) HasDATA_TRANSFER_EXTERNAL_GBYTES() bool`

HasDATA_TRANSFER_EXTERNAL_GBYTES returns a boolean if a field has been set.

### SetDATA_TRANSFER_EXTERNAL_GBYTESNil

`func (o *RunUsage) SetDATA_TRANSFER_EXTERNAL_GBYTESNil(b bool)`

 SetDATA_TRANSFER_EXTERNAL_GBYTESNil sets the value for DATA_TRANSFER_EXTERNAL_GBYTES to be an explicit nil

### UnsetDATA_TRANSFER_EXTERNAL_GBYTES
`func (o *RunUsage) UnsetDATA_TRANSFER_EXTERNAL_GBYTES()`

UnsetDATA_TRANSFER_EXTERNAL_GBYTES ensures that no value is present for DATA_TRANSFER_EXTERNAL_GBYTES, not even an explicit nil
### GetPROXY_RESIDENTIAL_TRANSFER_GBYTES

`func (o *RunUsage) GetPROXY_RESIDENTIAL_TRANSFER_GBYTES() float32`

GetPROXY_RESIDENTIAL_TRANSFER_GBYTES returns the PROXY_RESIDENTIAL_TRANSFER_GBYTES field if non-nil, zero value otherwise.

### GetPROXY_RESIDENTIAL_TRANSFER_GBYTESOk

`func (o *RunUsage) GetPROXY_RESIDENTIAL_TRANSFER_GBYTESOk() (*float32, bool)`

GetPROXY_RESIDENTIAL_TRANSFER_GBYTESOk returns a tuple with the PROXY_RESIDENTIAL_TRANSFER_GBYTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROXY_RESIDENTIAL_TRANSFER_GBYTES

`func (o *RunUsage) SetPROXY_RESIDENTIAL_TRANSFER_GBYTES(v float32)`

SetPROXY_RESIDENTIAL_TRANSFER_GBYTES sets PROXY_RESIDENTIAL_TRANSFER_GBYTES field to given value.

### HasPROXY_RESIDENTIAL_TRANSFER_GBYTES

`func (o *RunUsage) HasPROXY_RESIDENTIAL_TRANSFER_GBYTES() bool`

HasPROXY_RESIDENTIAL_TRANSFER_GBYTES returns a boolean if a field has been set.

### SetPROXY_RESIDENTIAL_TRANSFER_GBYTESNil

`func (o *RunUsage) SetPROXY_RESIDENTIAL_TRANSFER_GBYTESNil(b bool)`

 SetPROXY_RESIDENTIAL_TRANSFER_GBYTESNil sets the value for PROXY_RESIDENTIAL_TRANSFER_GBYTES to be an explicit nil

### UnsetPROXY_RESIDENTIAL_TRANSFER_GBYTES
`func (o *RunUsage) UnsetPROXY_RESIDENTIAL_TRANSFER_GBYTES()`

UnsetPROXY_RESIDENTIAL_TRANSFER_GBYTES ensures that no value is present for PROXY_RESIDENTIAL_TRANSFER_GBYTES, not even an explicit nil
### GetPROXY_SERPS

`func (o *RunUsage) GetPROXY_SERPS() int32`

GetPROXY_SERPS returns the PROXY_SERPS field if non-nil, zero value otherwise.

### GetPROXY_SERPSOk

`func (o *RunUsage) GetPROXY_SERPSOk() (*int32, bool)`

GetPROXY_SERPSOk returns a tuple with the PROXY_SERPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROXY_SERPS

`func (o *RunUsage) SetPROXY_SERPS(v int32)`

SetPROXY_SERPS sets PROXY_SERPS field to given value.

### HasPROXY_SERPS

`func (o *RunUsage) HasPROXY_SERPS() bool`

HasPROXY_SERPS returns a boolean if a field has been set.

### SetPROXY_SERPSNil

`func (o *RunUsage) SetPROXY_SERPSNil(b bool)`

 SetPROXY_SERPSNil sets the value for PROXY_SERPS to be an explicit nil

### UnsetPROXY_SERPS
`func (o *RunUsage) UnsetPROXY_SERPS()`

UnsetPROXY_SERPS ensures that no value is present for PROXY_SERPS, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


