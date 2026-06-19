package apify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// ListRequestsOptions configures [RequestQueueClient.ListRequests].
type ListRequestsOptions struct {
	// Limit is the maximum number of requests to return.
	Limit *int64
	// ExclusiveStartID lists requests after this ID.
	ExclusiveStartID *string
	// Cursor is an opaque pagination cursor (alternative to ExclusiveStartID).
	Cursor *string
	// Filter restricts the listing (e.g. to handled/unhandled requests).
	Filter *string
}

func (o ListRequestsOptions) apply(q *QueryParams) {
	q.AddInt("limit", o.Limit).
		AddString("exclusiveStartId", o.ExclusiveStartID).
		AddString("cursor", o.Cursor).
		AddString("filter", o.Filter)
}

// RequestQueueClient is a client for a specific request queue (and run-nested variants).
type RequestQueueClient struct {
	ctx       *resourceContext
	clientKey string
}

func newRequestQueueClient(hc *httpClient, baseURL, resourcePath, id string) *RequestQueueClient {
	return &RequestQueueClient{ctx: newSingleContext(hc, baseURL, resourcePath, id)}
}

// newRequestQueueNestedClient creates a client for a run's default request queue.
func newRequestQueueNestedClient(hc *httpClient, base, subPath string) *RequestQueueClient {
	return &RequestQueueClient{ctx: newCollectionContext(hc, base, subPath)}
}

// WithClientKey returns a copy of the client that identifies its requests with clientKey.
//
// A stable client key is required to operate on locks the client itself created (e.g. to
// unlock its own requests), and lets the API detect whether multiple clients access a queue.
func (c *RequestQueueClient) WithClientKey(clientKey string) *RequestQueueClient {
	clone := *c
	clone.clientKey = clientKey
	return &clone
}

// withClientKey adds the client key (if set) to the given params.
func (c *RequestQueueClient) withClientKey(params *QueryParams) *QueryParams {
	if c.clientKey != "" {
		params.AddString("clientKey", &c.clientKey)
	}
	return params
}

// Get fetches the queue metadata. The bool reports whether it exists.
func (c *RequestQueueClient) Get(ctx context.Context) (RequestQueue, bool, error) {
	return getResource[RequestQueue](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the queue metadata (e.g. name) and returns the updated object.
func (c *RequestQueueClient) Update(ctx context.Context, newFields any) (RequestQueue, error) {
	return updateResource[RequestQueue](ctx, c.ctx, "", newFields)
}

// Delete deletes the queue.
func (c *RequestQueueClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// ListHead returns the requests at the head (front) of the queue, up to limit (nil for the
// server default).
func (c *RequestQueueClient) ListHead(ctx context.Context, limit *int64) (RequestQueueHead, error) {
	params := NewQueryParams()
	params.AddInt("limit", limit)
	c.withClientKey(params)
	return getResourceRequired[RequestQueueHead](ctx, c.ctx, "head", params)
}

// AddRequest adds a request to the queue. If forefront is true, the request is added to the
// front of the queue.
func (c *RequestQueueClient) AddRequest(ctx context.Context, request RequestQueueRequest, forefront bool) (RequestQueueOperationInfo, error) {
	params := NewQueryParams()
	params.AddBool("forefront", &forefront)
	c.withClientKey(params)
	body, err := json.Marshal(request)
	if err != nil {
		return RequestQueueOperationInfo{}, err
	}
	return postWithBody[RequestQueueOperationInfo](ctx, c.ctx, "requests", params, body, contentTypeJSON)
}

// GetRequest fetches a request by ID, or (nil, false, nil) if it does not exist.
func (c *RequestQueueClient) GetRequest(ctx context.Context, id string) (*RequestQueueRequest, bool, error) {
	req, present, err := getResource[RequestQueueRequest](ctx, c.ctx, "requests/"+encodePathSegment(id), NewQueryParams())
	if err != nil || !present {
		return nil, present, err
	}
	return &req, true, nil
}

// UpdateRequest updates an existing request (identified by its ID field) and returns the
// operation info. If forefront is true, the request is moved to the front of the queue.
func (c *RequestQueueClient) UpdateRequest(ctx context.Context, request RequestQueueRequest, forefront bool) (RequestQueueOperationInfo, error) {
	params := NewQueryParams()
	params.AddBool("forefront", &forefront)
	c.withClientKey(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("requests/" + encodePathSegment(request.ID)))
	body, err := json.Marshal(request)
	if err != nil {
		return RequestQueueOperationInfo{}, err
	}
	resp, err := c.ctx.http.call(ctx, "PUT", url, body, contentTypeJSON, defaultRequestTimeout)
	if err != nil {
		return RequestQueueOperationInfo{}, err
	}
	return parseDataEnvelope[RequestQueueOperationInfo](resp.body)
}

// DeleteRequest deletes a request by ID.
func (c *RequestQueueClient) DeleteRequest(ctx context.Context, id string) error {
	url := c.ctx.mergedParams(c.withClientKey(NewQueryParams())).applyToURL(c.ctx.subURL("requests/" + encodePathSegment(id)))
	_, err := c.ctx.http.call(ctx, "DELETE", url, nil, "", defaultRequestTimeout)
	if err != nil && !isNotFound(err) {
		return err
	}
	return nil
}

// ListAndLockHead atomically returns and locks up to limit requests from the head of the
// queue for lockSecs seconds. Returns the raw API response (a locked-head object).
func (c *RequestQueueClient) ListAndLockHead(ctx context.Context, lockSecs int64, limit *int64) (json.RawMessage, error) {
	params := NewQueryParams()
	params.AddInt("lockSecs", &lockSecs).AddInt("limit", limit)
	c.withClientKey(params)
	return postWithBody[json.RawMessage](ctx, c.ctx, "head/lock", params, nil, "")
}

// maxRequestsPerBatchOperation is the API limit on requests per batch call. Larger inputs
// are split into chunks of this size, matching the reference client.
const maxRequestsPerBatchOperation = 25

// BatchAddResult is the typed result of [RequestQueueClient.BatchAddRequests]: the requests
// the API accepted and the ones it could not process.
type BatchAddResult struct {
	// ProcessedRequests are the requests the API successfully added.
	ProcessedRequests []RequestQueueOperationInfo `json:"processedRequests"`
	// UnprocessedRequests are the requests the API did not process.
	UnprocessedRequests []RequestQueueRequest `json:"unprocessedRequests"`
}

// BatchAddRequests adds multiple requests to the queue. If forefront is true, they are added
// to the front of the queue.
//
// The input is automatically split into chunks of at most 25 requests (the API limit), and
// the per-chunk results are merged into a single [BatchAddResult]. Each chunk is still
// subject to the client's standard retry policy.
func (c *RequestQueueClient) BatchAddRequests(ctx context.Context, requests []RequestQueueRequest, forefront bool) (BatchAddResult, error) {
	var merged BatchAddResult
	for start := 0; start < len(requests); start += maxRequestsPerBatchOperation {
		end := start + maxRequestsPerBatchOperation
		if end > len(requests) {
			end = len(requests)
		}
		chunkResult, err := c.batchAddChunk(ctx, requests[start:end], forefront)
		if err != nil {
			return merged, err
		}
		merged.ProcessedRequests = append(merged.ProcessedRequests, chunkResult.ProcessedRequests...)
		merged.UnprocessedRequests = append(merged.UnprocessedRequests, chunkResult.UnprocessedRequests...)
	}
	return merged, nil
}

// batchAddChunk sends a single batch (<= 25 requests) and parses the typed result.
func (c *RequestQueueClient) batchAddChunk(ctx context.Context, requests []RequestQueueRequest, forefront bool) (BatchAddResult, error) {
	params := NewQueryParams()
	params.AddBool("forefront", &forefront)
	c.withClientKey(params)
	body, err := json.Marshal(requests)
	if err != nil {
		return BatchAddResult{}, err
	}
	return postWithBody[BatchAddResult](ctx, c.ctx, "requests/batch", params, body, contentTypeJSON)
}

// BatchDeleteRequests deletes multiple requests in a single call. Each entry identifies a
// request (e.g. by id or uniqueKey). Returns the raw batch result.
func (c *RequestQueueClient) BatchDeleteRequests(ctx context.Context, requests any) (json.RawMessage, error) {
	params := c.withClientKey(NewQueryParams())
	return deleteWithBody[json.RawMessage](ctx, c.ctx, "requests/batch", params, requests)
}

// Allowed values for ListRequestsOptions.Filter, as constrained by the API.
const (
	requestFilterLocked  = "locked"
	requestFilterPending = "pending"
)

// validate checks the listing options for API-level constraints: ExclusiveStartID and Cursor
// are mutually exclusive, and Filter (if set) must be "locked" or "pending".
func (o ListRequestsOptions) validate() error {
	if o.ExclusiveStartID != nil && o.Cursor != nil {
		return errors.New("ListRequestsOptions: ExclusiveStartID and Cursor are mutually exclusive")
	}
	if o.Filter != nil && *o.Filter != requestFilterLocked && *o.Filter != requestFilterPending {
		return fmt.Errorf("ListRequestsOptions: Filter must be %q or %q, got %q", requestFilterLocked, requestFilterPending, *o.Filter)
	}
	return nil
}

// ListRequests lists the queue's requests with pagination.
func (c *RequestQueueClient) ListRequests(ctx context.Context, options ListRequestsOptions) (json.RawMessage, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}
	params := NewQueryParams()
	options.apply(params)
	c.withClientKey(params)
	return getResourceRequired[json.RawMessage](ctx, c.ctx, "requests", params)
}

// ProlongRequestLock extends the lock on a request by lockSecs seconds. If forefront is
// true, the request is moved to the front when its lock expires. Returns the raw response.
func (c *RequestQueueClient) ProlongRequestLock(ctx context.Context, id string, lockSecs int64, forefront bool) (json.RawMessage, error) {
	params := NewQueryParams()
	params.AddInt("lockSecs", &lockSecs).AddBool("forefront", &forefront)
	c.withClientKey(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("requests/" + encodePathSegment(id) + "/lock"))
	resp, err := c.ctx.http.call(ctx, "PUT", url, nil, "", defaultRequestTimeout)
	if err != nil {
		return nil, err
	}
	return parseDataEnvelope[json.RawMessage](resp.body)
}

// DeleteRequestLock releases the lock on a request. If forefront is true, the request is
// moved to the front of the queue.
func (c *RequestQueueClient) DeleteRequestLock(ctx context.Context, id string, forefront bool) error {
	params := NewQueryParams()
	params.AddBool("forefront", &forefront)
	c.withClientKey(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("requests/" + encodePathSegment(id) + "/lock"))
	_, err := c.ctx.http.call(ctx, "DELETE", url, nil, "", defaultRequestTimeout)
	if err != nil && !isNotFound(err) {
		return err
	}
	return nil
}

// UnlockRequests releases all locks the client holds on this queue's requests. Returns the
// raw response.
func (c *RequestQueueClient) UnlockRequests(ctx context.Context) (json.RawMessage, error) {
	params := c.withClientKey(NewQueryParams())
	return postWithBody[json.RawMessage](ctx, c.ctx, "requests/unlock", params, nil, "")
}

// PaginateRequests returns a lazy iterator over all requests in the queue, fetching pages
// of up to pageLimit requests at a time (nil for the server default).
func (c *RequestQueueClient) PaginateRequests(pageLimit *int64) *RequestQueueRequestsIterator {
	return &RequestQueueRequestsIterator{client: c, pageLimit: pageLimit}
}

// RequestQueueRequestsIterator lazily iterates over a request queue's requests, fetching one
// page at a time via the cursor-based listing endpoint.
type RequestQueueRequestsIterator struct {
	client    *RequestQueueClient
	pageLimit *int64

	buffer     []RequestQueueRequest
	pos        int
	nextCursor string
	started    bool
	exhausted  bool
}

// requestsPage is the shape of a paginated requests listing.
type requestsPage struct {
	Items      []RequestQueueRequest `json:"items"`
	NextCursor *string               `json:"nextCursor"`
}

// Next returns the next request, or (nil, nil) when the iterator is exhausted.
func (it *RequestQueueRequestsIterator) Next(ctx context.Context) (*RequestQueueRequest, error) {
	for it.pos >= len(it.buffer) {
		if it.exhausted || (it.started && it.nextCursor == "") {
			return nil, nil
		}
		if err := it.fetchPage(ctx); err != nil {
			return nil, err
		}
	}
	req := it.buffer[it.pos]
	it.pos++
	return &req, nil
}

// fetchPage loads the next page of requests into the buffer.
func (it *RequestQueueRequestsIterator) fetchPage(ctx context.Context) error {
	params := NewQueryParams()
	params.AddInt("limit", it.pageLimit)
	if it.nextCursor != "" {
		params.AddString("cursor", &it.nextCursor)
	}
	it.client.withClientKey(params)

	raw, err := getResourceRequired[requestsPage](ctx, it.client.ctx, "requests", params)
	if err != nil {
		return err
	}
	it.started = true
	it.buffer = raw.Items
	it.pos = 0
	if raw.NextCursor != nil {
		it.nextCursor = *raw.NextCursor
	} else {
		it.nextCursor = ""
	}
	if len(raw.Items) == 0 && it.nextCursor == "" {
		it.exhausted = true
	}
	return nil
}
