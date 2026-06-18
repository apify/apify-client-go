# BatchAddResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessedRequests** | [**[]AddedRequest**](AddedRequest.md) | Requests that were successfully added to the request queue. | 
**UnprocessedRequests** | [**[]RequestDraft**](RequestDraft.md) | Requests that failed to be added and can be retried. | 

## Methods

### NewBatchAddResult

`func NewBatchAddResult(processedRequests []AddedRequest, unprocessedRequests []RequestDraft, ) *BatchAddResult`

NewBatchAddResult instantiates a new BatchAddResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchAddResultWithDefaults

`func NewBatchAddResultWithDefaults() *BatchAddResult`

NewBatchAddResultWithDefaults instantiates a new BatchAddResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessedRequests

`func (o *BatchAddResult) GetProcessedRequests() []AddedRequest`

GetProcessedRequests returns the ProcessedRequests field if non-nil, zero value otherwise.

### GetProcessedRequestsOk

`func (o *BatchAddResult) GetProcessedRequestsOk() (*[]AddedRequest, bool)`

GetProcessedRequestsOk returns a tuple with the ProcessedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRequests

`func (o *BatchAddResult) SetProcessedRequests(v []AddedRequest)`

SetProcessedRequests sets ProcessedRequests field to given value.


### GetUnprocessedRequests

`func (o *BatchAddResult) GetUnprocessedRequests() []RequestDraft`

GetUnprocessedRequests returns the UnprocessedRequests field if non-nil, zero value otherwise.

### GetUnprocessedRequestsOk

`func (o *BatchAddResult) GetUnprocessedRequestsOk() (*[]RequestDraft, bool)`

GetUnprocessedRequestsOk returns a tuple with the UnprocessedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessedRequests

`func (o *BatchAddResult) SetUnprocessedRequests(v []RequestDraft)`

SetUnprocessedRequests sets UnprocessedRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


