package apify

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListKeysOptions configures [KeyValueStoreClient.ListKeys].
type ListKeysOptions struct {
	// Limit is the maximum number of keys to return.
	Limit *int64
	// ExclusiveStartKey lists keys after this one (for pagination).
	ExclusiveStartKey *string
	// Prefix restricts the listing to keys with this prefix.
	Prefix *string
	// Collection restricts the listing to a named collection of keys.
	Collection *string
	// Signature is a pre-shared URL signature granting access without an API token.
	Signature *string
}

func (o ListKeysOptions) apply(q *QueryParams) {
	q.AddInt("limit", o.Limit).
		AddString("exclusiveStartKey", o.ExclusiveStartKey).
		AddString("prefix", o.Prefix).
		AddString("collection", o.Collection).
		AddString("signature", o.Signature)
}

// GetRecordsOptions configures [KeyValueStoreClient.GetRecords].
type GetRecordsOptions struct {
	// Collection restricts the download to a named collection of records.
	Collection *string
	// Prefix restricts the download to records with this key prefix.
	Prefix *string
	// Signature is a pre-shared URL signature granting access without an API token.
	Signature *string
}

func (o GetRecordsOptions) apply(q *QueryParams) {
	q.AddString("collection", o.Collection).
		AddString("prefix", o.Prefix).
		AddString("signature", o.Signature)
}

// GetRecordOptions configures [KeyValueStoreClient.GetRecordWithOptions].
type GetRecordOptions struct {
	// Attachment, if set, controls the Content-Disposition: attachment behaviour.
	Attachment *bool
	// Signature is a pre-shared URL signature granting access without an API token.
	Signature *string
}

func (o GetRecordOptions) apply(q *QueryParams) {
	q.AddBool("attachment", o.Attachment).AddString("signature", o.Signature)
}

// KeyValueStoreClient is a client for a specific key-value store (and run-nested variants).
type KeyValueStoreClient struct {
	ctx *resourceContext
}

func newKeyValueStoreClient(hc *httpClient, baseURL, resourcePath, id string) *KeyValueStoreClient {
	return &KeyValueStoreClient{ctx: newSingleContext(hc, baseURL, resourcePath, id)}
}

// newKeyValueStoreNestedClient creates a client for a run's default key-value store.
func newKeyValueStoreNestedClient(hc *httpClient, base, subPath string) *KeyValueStoreClient {
	return &KeyValueStoreClient{ctx: newCollectionContext(hc, base, subPath)}
}

// withPublicBase sets the public origin used when building shareable URLs.
func (c *KeyValueStoreClient) withPublicBase(publicBaseURL string) *KeyValueStoreClient {
	c.ctx = c.ctx.withPublicOrigin(publicBaseURL)
	return c
}

// Get fetches the store metadata. The bool reports whether it exists.
func (c *KeyValueStoreClient) Get(ctx context.Context) (KeyValueStore, bool, error) {
	return getResource[KeyValueStore](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the store metadata (e.g. name) and returns the updated object.
func (c *KeyValueStoreClient) Update(ctx context.Context, newFields any) (KeyValueStore, error) {
	return updateResource[KeyValueStore](ctx, c.ctx, "", newFields)
}

// Delete deletes the store.
func (c *KeyValueStoreClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// ListKeys lists the keys stored in this key-value store.
func (c *KeyValueStoreClient) ListKeys(ctx context.Context, options ListKeysOptions) (KeyValueStoreKeysPage, error) {
	params := NewQueryParams()
	options.apply(params)
	return getResourceRequired[KeyValueStoreKeysPage](ctx, c.ctx, "keys", params)
}

// IterateKeys returns a lazy iterator over the store's keys, fetching one page at a time on
// demand via the cursor-based keys endpoint (exclusiveStartKey / nextExclusiveStartKey). It
// mirrors the reference client's async-iterable listKeys() and follows this client's iteration
// convention (like the collection Iterate helpers): the options' Limit caps the total number of
// keys yielded across all pages (unset means all keys), the per-page size is the separate
// chunkSize argument (nil for the server default), Prefix/Collection/Signature filter every
// page, and ExclusiveStartKey sets the key to start listing after.
//
// Because keys are cursor-paginated (not offset/limit paginated) it uses its own
// KeyValueStoreKeysIterator rather than the generic ListIterator, sharing the cursor mechanics
// of RequestQueueClient.PaginateRequests.
func (c *KeyValueStoreClient) IterateKeys(options ListKeysOptions, chunkSize *int64) *KeyValueStoreKeysIterator {
	return &KeyValueStoreKeysIterator{client: c, options: options, chunkSize: chunkSize, nextStartKey: options.ExclusiveStartKey}
}

// KeyValueStoreKeysIterator lazily iterates over a key-value store's keys, fetching one page at
// a time via the cursor-based listing endpoint. Obtain one from KeyValueStoreClient.IterateKeys
// and drain it by calling Next until it returns (nil, nil).
type KeyValueStoreKeysIterator struct {
	client    *KeyValueStoreClient
	options   ListKeysOptions
	chunkSize *int64 // per-page size (nil or <=0 means the server default)

	buffer       []KeyValueStoreKey
	pos          int
	nextStartKey *string // cursor for the next page (nil once the API reports no more keys)
	remaining    int64   // total-item cap countdown; <0 means "no cap". Valid once started.
	started      bool
	exhausted    bool
}

// Next returns the next key, or (nil, nil) when the iterator is exhausted (no more keys or the
// total-item cap is reached). It calls the API for another page only when the current in-memory
// page is used up.
func (it *KeyValueStoreKeysIterator) Next(ctx context.Context) (*KeyValueStoreKey, error) {
	for it.pos >= len(it.buffer) {
		if it.exhausted {
			return nil, nil
		}
		if err := it.fetchPage(ctx); err != nil {
			return nil, err
		}
	}
	key := it.buffer[it.pos]
	it.pos++
	return &key, nil
}

// fetchPage loads the next page of keys into the buffer, advancing the cursor and the remaining
// cap. The per-page limit is the smaller of the remaining total cap and the requested chunk
// size, so the final page is never over-fetched.
func (it *KeyValueStoreKeysIterator) fetchPage(ctx context.Context) error {
	opts := it.options
	opts.ExclusiveStartKey = it.nextStartKey

	// capLeft is how many items the total cap still allows (0 = "no cap so far", treated as
	// unset when combined with the chunk size).
	var capLeft int64
	if !it.started {
		if l := it.options.Limit; l != nil && *l > 0 {
			capLeft = *l
		}
	} else if it.remaining > 0 {
		capLeft = it.remaining
	}
	opts.Limit = pageLimitPtr(minForLimitParam(capLeft, it.chunkVal()))

	page, err := it.client.ListKeys(ctx, opts)
	if err != nil {
		return err
	}
	n := int64(len(page.Items))
	it.buffer = page.Items
	it.pos = 0

	if page.NextExclusiveStartKey != "" {
		it.nextStartKey = &page.NextExclusiveStartKey
	} else {
		it.nextStartKey = nil
	}

	if !it.started {
		it.started = true
		if l := it.options.Limit; l != nil && *l > 0 {
			it.remaining = *l - n
		} else {
			it.remaining = -1 // no cap
		}
	} else if it.remaining >= 0 {
		it.remaining -= n
	}

	// Stop when the API returns an empty page, reports no more keys (not truncated / no next
	// cursor), or the total-item cap has been reached.
	if n == 0 || !page.IsTruncated || it.nextStartKey == nil || it.remaining == 0 {
		it.exhausted = true
	}
	return nil
}

// chunkVal returns the per-page size as a plain int64 (0 when unset, i.e. server default).
func (it *KeyValueStoreKeysIterator) chunkVal() int64 {
	if it.chunkSize == nil || *it.chunkSize < 0 {
		return 0
	}
	return *it.chunkSize
}

// GetRecords downloads all records of the store (optionally filtered) as a ZIP archive,
// returning the raw bytes.
func (c *KeyValueStoreClient) GetRecords(ctx context.Context, options GetRecordsOptions) ([]byte, error) {
	params := NewQueryParams()
	options.apply(params)
	url := params.applyToURL(c.ctx.subURL("records"))
	resp, err := c.ctx.http.call(ctx, http.MethodGet, url, nil, "", defaultRequestTimeout)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

// RecordExists reports whether a record with the given key exists.
func (c *KeyValueStoreClient) RecordExists(ctx context.Context, key string) (bool, error) {
	return headExists(ctx, c.ctx, "records/"+encodePathSegment(key), NewQueryParams())
}

// GetRecord fetches a record by key, or (nil, false, nil) if it does not exist. The value
// holds the raw bytes; the content type is reported in the record.
//
// Like the reference client, it requests the record as an attachment (sent on the wire as
// attachment=1, the truthy form this client's bool serializer uses) so the API returns the
// record's raw bytes directly rather than redirecting. Use [GetRecordWithOptions] to override.
func (c *KeyValueStoreClient) GetRecord(ctx context.Context, key string) (*KeyValueStoreRecord, bool, error) {
	return c.GetRecordWithOptions(ctx, key, GetRecordOptions{Attachment: ptrBool(true)})
}

// ptrBool returns a pointer to b. Tiny helper so GetRecord can default attachment=true.
func ptrBool(b bool) *bool { return &b }

// GetRecordWithOptions fetches a record with explicit options (attachment, signature).
func (c *KeyValueStoreClient) GetRecordWithOptions(ctx context.Context, key string, options GetRecordOptions) (*KeyValueStoreRecord, bool, error) {
	params := NewQueryParams()
	options.apply(params)
	resp, err := getRaw(ctx, c.ctx, "records/"+encodePathSegment(key), params)
	if err != nil {
		return nil, false, err
	}
	if resp == nil {
		return nil, false, nil
	}
	return &KeyValueStoreRecord{
		Key:         key,
		Value:       resp.body,
		ContentType: resp.header.Get("Content-Type"),
	}, true, nil
}

// SetRecordRaw stores a record with raw bytes and the given content type.
func (c *KeyValueStoreClient) SetRecordRaw(ctx context.Context, key string, value []byte, contentType string) error {
	return putRaw(ctx, c.ctx, "records/"+encodePathSegment(key), NewQueryParams(), value, contentType)
}

// SetRecordJSON stores a record holding the JSON serialization of value.
func (c *KeyValueStoreClient) SetRecordJSON(ctx context.Context, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.SetRecordRaw(ctx, key, data, contentTypeJSONCharset)
}

// DeleteRecord deletes a record by key.
func (c *KeyValueStoreClient) DeleteRecord(ctx context.Context, key string) error {
	return deleteResource(ctx, c.ctx, "records/"+encodePathSegment(key))
}

// GetRecordPublicURL builds a public URL for fetching the given record.
//
// It mirrors the reference client: it fetches the store, and if the store exposes a
// URL-signing secret key (i.e. it is private), appends an HMAC-SHA256 signature so the URL
// grants access without an API token. The URL is built from the configured public base URL.
func (c *KeyValueStoreClient) GetRecordPublicURL(ctx context.Context, key string) (string, error) {
	params := NewQueryParams()
	store, present, err := c.Get(ctx)
	if err != nil {
		return "", err
	}
	if present {
		if secret := extractString(store.Extra, "urlSigningSecretKey"); secret != "" {
			sig := createHmacSignature(secret, key)
			params.AddString("signature", &sig)
		}
	}
	return params.applyToURL(c.ctx.publicURL("records/" + encodePathSegment(key))), nil
}

// CreateKeysPublicURL builds a public URL for listing this store's keys.
//
// As with GetRecordPublicURL, a signature is appended for private stores. expiresInSecs
// optionally bounds the validity of a signed URL (nil for non-expiring).
func (c *KeyValueStoreClient) CreateKeysPublicURL(ctx context.Context, expiresInSecs *int64) (string, error) {
	params := NewQueryParams()
	store, present, err := c.Get(ctx)
	if err != nil {
		return "", err
	}
	if present {
		if secret := extractString(store.Extra, "urlSigningSecretKey"); secret != "" {
			sig := signStorageContent(secret, store.ID, expiresInSecs)
			params.AddString("signature", &sig)
		}
	}
	return params.applyToURL(c.ctx.publicURL("keys")), nil
}
