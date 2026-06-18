# BatchDeleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessedRequests** | [**[]DeletedRequest**](DeletedRequest.md) | Requests that were successfully deleted from the request queue. | 
**UnprocessedRequests** | [**[]RequestDraft**](RequestDraft.md) | Requests that failed to be deleted and can be retried. | 

## Methods

### NewBatchDeleteResult

`func NewBatchDeleteResult(processedRequests []DeletedRequest, unprocessedRequests []RequestDraft, ) *BatchDeleteResult`

NewBatchDeleteResult instantiates a new BatchDeleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchDeleteResultWithDefaults

`func NewBatchDeleteResultWithDefaults() *BatchDeleteResult`

NewBatchDeleteResultWithDefaults instantiates a new BatchDeleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessedRequests

`func (o *BatchDeleteResult) GetProcessedRequests() []DeletedRequest`

GetProcessedRequests returns the ProcessedRequests field if non-nil, zero value otherwise.

### GetProcessedRequestsOk

`func (o *BatchDeleteResult) GetProcessedRequestsOk() (*[]DeletedRequest, bool)`

GetProcessedRequestsOk returns a tuple with the ProcessedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRequests

`func (o *BatchDeleteResult) SetProcessedRequests(v []DeletedRequest)`

SetProcessedRequests sets ProcessedRequests field to given value.


### GetUnprocessedRequests

`func (o *BatchDeleteResult) GetUnprocessedRequests() []RequestDraft`

GetUnprocessedRequests returns the UnprocessedRequests field if non-nil, zero value otherwise.

### GetUnprocessedRequestsOk

`func (o *BatchDeleteResult) GetUnprocessedRequestsOk() (*[]RequestDraft, bool)`

GetUnprocessedRequestsOk returns a tuple with the UnprocessedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessedRequests

`func (o *BatchDeleteResult) SetUnprocessedRequests(v []RequestDraft)`

SetUnprocessedRequests sets UnprocessedRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


