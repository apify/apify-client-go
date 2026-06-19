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
