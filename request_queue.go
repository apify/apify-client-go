package apify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ListRequestsOptions configures [RequestQueueClient.ListRequests].
type ListRequestsOptions struct {
	// Limit is the maximum number of requests to return.
	Limit *int64
	// ExclusiveStartID lists requests after this ID.
	ExclusiveStartID *string
	// Cursor is an opaque pagination cursor (alternative to ExclusiveStartID).
	Cursor *string
	// Filter restricts the listing to requests in the given states. Each value must be
	// "locked" or "pending" (see RequestFilterLocked / RequestFilterPending). Multiple
	// values are sent as a comma-separated list and mean the union of those states
	// (requests matching any of them are returned), matching the API.
	Filter []string
}

func (o ListRequestsOptions) apply(q *QueryParams) {
	q.AddInt("limit", o.Limit).
		AddString("exclusiveStartId", o.ExclusiveStartID).
		AddString("cursor", o.Cursor).
		AddCSV("filter", o.Filter)
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
	return getResourceRequired[RequestQueueHead](ctx, c.ctx, "head", params, c.ctx.shortTimeout())
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
	return postWithBody[RequestQueueOperationInfo](ctx, c.ctx, "requests", params, body, contentTypeJSON, c.ctx.shortTimeout())
}

// GetRequest fetches a request by ID, or (nil, false, nil) if it does not exist.
func (c *RequestQueueClient) GetRequest(ctx context.Context, id string) (*RequestQueueRequest, bool, error) {
	safeID, err := safePathSegment(id)
	if err != nil {
		return nil, false, err
	}
	req, present, err := getResourceAlways[RequestQueueRequest](ctx, c.ctx, "requests/"+safeID, NewQueryParams())
	if err != nil || !present {
		return nil, present, err
	}
	return &req, true, nil
}

// UpdateRequest updates an existing request (identified by its ID field) and returns the
// operation info. If forefront is true, the request is moved to the front of the queue.
func (c *RequestQueueClient) UpdateRequest(ctx context.Context, request RequestQueueRequest, forefront bool) (RequestQueueOperationInfo, error) {
	if c.ctx.idErr != nil {
		return RequestQueueOperationInfo{}, c.ctx.idErr
	}
	safeID, err := safePathSegment(request.ID)
	if err != nil {
		return RequestQueueOperationInfo{}, err
	}
	params := NewQueryParams()
	params.AddBool("forefront", &forefront)
	c.withClientKey(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("requests/" + safeID))
	body, err := json.Marshal(request)
	if err != nil {
		return RequestQueueOperationInfo{}, err
	}
	resp, err := c.ctx.http.call(ctx, http.MethodPut, url, body, contentTypeJSON, c.ctx.mediumTimeout())
	if err != nil {
		return RequestQueueOperationInfo{}, err
	}
	return parseDataEnvelope[RequestQueueOperationInfo](resp.body)
}

// DeleteRequest deletes a request by ID.
func (c *RequestQueueClient) DeleteRequest(ctx context.Context, id string) error {
	if c.ctx.idErr != nil {
		return c.ctx.idErr
	}
	safeID, err := safePathSegment(id)
	if err != nil {
		return err
	}
	url := c.ctx.mergedParams(c.withClientKey(NewQueryParams())).applyToURL(c.ctx.subURL("requests/" + safeID))
	_, err = c.ctx.http.call(ctx, http.MethodDelete, url, nil, "", c.ctx.shortTimeout())
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
	return postWithBody[json.RawMessage](ctx, c.ctx, "head/lock", params, nil, "", c.ctx.mediumTimeout())
}

// maxRequestsPerBatchOperation is the API limit on requests per batch call. Larger inputs
// are split into chunks of this size, matching the reference client.
const maxRequestsPerBatchOperation = 25

// maxPayloadSizeBytes is the API's maximum request body size for a batch-requests call,
// matching the published @apify/consts MAX_PAYLOAD_SIZE_BYTES (9 MiB; apify/apify-shared-js).
const maxPayloadSizeBytes = 9437184

// payloadSizeLimitBytes shaves a small safety margin off maxPayloadSizeBytes (0.01%, matching
// the reference client's SAFETY_BUFFER_PERCENT), so a batch this client assembles stays under
// the server's limit even accounting for any rounding differences between the two.
// 9437184 * 0.0001 = 943.7184, rounded up to 944; 9437184 - 944 = 9436240.
const payloadSizeLimitBytes = maxPayloadSizeBytes - 944

// BatchAddResult is the typed result of [RequestQueueClient.BatchAddRequests]: the requests
// the API accepted and the ones it could not process.
type BatchAddResult struct {
	// ProcessedRequests are the requests the API successfully added.
	ProcessedRequests []RequestQueueOperationInfo `json:"processedRequests"`
	// UnprocessedRequests are the requests the API did not process.
	UnprocessedRequests []RequestQueueRequest `json:"unprocessedRequests"`
}

// serializedRequest is a request to add, serialized once up front: json is what the body of
// its batch is assembled from (so sending it never re-marshals the request), and original is
// what the result bookkeeping needs.
type serializedRequest struct {
	original RequestQueueRequest
	json     []byte
}

// serializeRequestsForBatch serializes every request once and checks each one against
// payloadSizeLimitBytes before any request is sent, so an oversized request fails the whole
// call up front instead of after the batches before it have already gone out.
func serializeRequestsForBatch(requests []RequestQueueRequest) ([]serializedRequest, error) {
	serialized := make([]serializedRequest, len(requests))
	for i, r := range requests {
		data, err := json.Marshal(r)
		if err != nil {
			return nil, err
		}
		// +2 bytes for the enclosing brackets, which even a batch of one request carries.
		if len(data)+2 > payloadSizeLimitBytes {
			return nil, fmt.Errorf("RequestQueueClient.BatchAddRequests: the request at index %d exceeds the maximum allowed size (%d bytes)", i, payloadSizeLimitBytes)
		}
		serialized[i] = serializedRequest{original: r, json: data}
	}
	return serialized, nil
}

// splitIntoBatches groups serialized requests into consecutive batches of at most maxCount
// requests, each fitting a JSON array body (items joined by commas between brackets) of at
// most maxBytes bytes. A request too large for a batch of its own is rejected beforehand by
// serializeRequestsForBatch, so every request here fits.
func splitIntoBatches(serialized []serializedRequest, maxCount int, maxBytes int) [][]serializedRequest {
	var batches [][]serializedRequest
	var batch []serializedRequest
	byteLen := 1 // the opening bracket
	for _, s := range serialized {
		// Each item adds its own bytes plus one for the following comma or closing bracket.
		if len(batch) > 0 && (len(batch) >= maxCount || byteLen+len(s.json)+1 > maxBytes) {
			batches = append(batches, batch)
			batch = nil
			byteLen = 1
		}
		batch = append(batch, s)
		byteLen += len(s.json) + 1
	}
	if len(batch) > 0 {
		batches = append(batches, batch)
	}
	return batches
}

// joinAsJSONArray joins the requests' pre-serialized JSON into a single JSON array body, so
// sending a batch never re-marshals the requests it contains.
func joinAsJSONArray(batch []serializedRequest) []byte {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, s := range batch {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.Write(s.json)
	}
	buf.WriteByte(']')
	return buf.Bytes()
}

// BatchAddRequests adds multiple requests to the queue. If forefront is true, they are added
// to the front of the queue.
//
// The input is serialized once and split into batches that respect both the API's 25-request
// and (effective, after a small safety margin) 9 MiB payload-size limits; the whole call fails
// before anything is sent if any single request is too large for a batch of its own. The
// per-batch results are merged into a single [BatchAddResult]. Each batch is still subject to
// the client's standard retry policy, but unlike the reference client this does not itself
// retry a batch's unprocessed requests — inspect [BatchAddResult.UnprocessedRequests] and retry
// those explicitly if needed.
func (c *RequestQueueClient) BatchAddRequests(ctx context.Context, requests []RequestQueueRequest, forefront bool) (BatchAddResult, error) {
	var merged BatchAddResult
	serialized, err := serializeRequestsForBatch(requests)
	if err != nil {
		return merged, err
	}
	for _, batch := range splitIntoBatches(serialized, maxRequestsPerBatchOperation, payloadSizeLimitBytes) {
		chunkResult, err := c.batchAddChunk(ctx, batch, forefront)
		if err != nil {
			return merged, err
		}
		merged.ProcessedRequests = append(merged.ProcessedRequests, chunkResult.ProcessedRequests...)
		merged.UnprocessedRequests = append(merged.UnprocessedRequests, chunkResult.UnprocessedRequests...)
	}
	return merged, nil
}

// batchAddChunk sends a single, pre-serialized and pre-sized batch and parses the typed result.
func (c *RequestQueueClient) batchAddChunk(ctx context.Context, batch []serializedRequest, forefront bool) (BatchAddResult, error) {
	params := NewQueryParams()
	params.AddBool("forefront", &forefront)
	c.withClientKey(params)
	body := joinAsJSONArray(batch)
	return postWithBody[BatchAddResult](ctx, c.ctx, "requests/batch", params, body, contentTypeJSON, c.ctx.mediumTimeout())
}

// BatchDeleteRequests deletes multiple requests in a single call. Each entry identifies a
// request (e.g. by id or uniqueKey). Returns the raw batch result.
func (c *RequestQueueClient) BatchDeleteRequests(ctx context.Context, requests any) (json.RawMessage, error) {
	params := c.withClientKey(NewQueryParams())
	return deleteWithBody[json.RawMessage](ctx, c.ctx, "requests/batch", params, requests)
}

// Allowed values for entries in ListRequestsOptions.Filter, as constrained by the API.
const (
	// RequestFilterLocked filters the listing to currently locked requests.
	RequestFilterLocked = "locked"
	// RequestFilterPending filters the listing to pending (not-yet-handled) requests.
	RequestFilterPending = "pending"
)

// validate checks the listing options for API-level constraints: ExclusiveStartID and Cursor
// are mutually exclusive, and every Filter entry (if any) must be "locked" or "pending".
func (o ListRequestsOptions) validate() error {
	if o.ExclusiveStartID != nil && o.Cursor != nil {
		return errors.New("ListRequestsOptions: ExclusiveStartID and Cursor are mutually exclusive")
	}
	for _, f := range o.Filter {
		if f != RequestFilterLocked && f != RequestFilterPending {
			return fmt.Errorf("ListRequestsOptions: Filter entries must be %q or %q, got %q", RequestFilterLocked, RequestFilterPending, f)
		}
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
	return getResourceRequired[json.RawMessage](ctx, c.ctx, "requests", params, c.ctx.mediumTimeout())
}

// ProlongRequestLock extends the lock on a request by lockSecs seconds. If forefront is
// true, the request is moved to the front when its lock expires. Returns the raw response.
func (c *RequestQueueClient) ProlongRequestLock(ctx context.Context, id string, lockSecs int64, forefront bool) (json.RawMessage, error) {
	if c.ctx.idErr != nil {
		return nil, c.ctx.idErr
	}
	safeID, err := safePathSegment(id)
	if err != nil {
		return nil, err
	}
	params := NewQueryParams()
	params.AddInt("lockSecs", &lockSecs).AddBool("forefront", &forefront)
	c.withClientKey(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("requests/" + safeID + "/lock"))
	resp, err := c.ctx.http.call(ctx, http.MethodPut, url, nil, "", c.ctx.mediumTimeout())
	if err != nil {
		return nil, err
	}
	return parseDataEnvelope[json.RawMessage](resp.body)
}

// DeleteRequestLock releases the lock on a request. If forefront is true, the request is
// moved to the front of the queue.
func (c *RequestQueueClient) DeleteRequestLock(ctx context.Context, id string, forefront bool) error {
	if c.ctx.idErr != nil {
		return c.ctx.idErr
	}
	safeID, err := safePathSegment(id)
	if err != nil {
		return err
	}
	params := NewQueryParams()
	params.AddBool("forefront", &forefront)
	c.withClientKey(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("requests/" + safeID + "/lock"))
	_, err = c.ctx.http.call(ctx, http.MethodDelete, url, nil, "", c.ctx.shortTimeout())
	if err != nil && !isNotFound(err) {
		return err
	}
	return nil
}

// UnlockRequests releases all locks the client holds on this queue's requests. Returns the
// raw response.
func (c *RequestQueueClient) UnlockRequests(ctx context.Context) (json.RawMessage, error) {
	params := c.withClientKey(NewQueryParams())
	return postWithBody[json.RawMessage](ctx, c.ctx, "requests/unlock", params, nil, "", c.ctx.longTimeout())
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

	raw, err := getResourceRequired[requestsPage](ctx, it.client.ctx, "requests", params, it.client.ctx.mediumTimeout())
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
