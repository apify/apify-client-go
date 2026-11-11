package apify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// DatasetListItemsOptions configures listing or downloading dataset items
// (GET /v2/datasets/{datasetId}/items).
type DatasetListItemsOptions struct {
	// Offset is the number of items to skip.
	Offset *int64
	// Limit is the maximum number of items to return.
	Limit *int64
	// Desc returns items newest-first.
	Desc *bool
	// Fields restricts the output to these fields.
	Fields []string
	// OutputFields positionally renames the fields selected by Fields in the output
	// (requires Fields). The i-th name becomes the output name of the i-th Fields entry.
	OutputFields []string
	// Omit excludes these fields from the output.
	Omit []string
	// SkipEmpty skips empty items.
	SkipEmpty *bool
	// SkipHidden skips hidden fields (those starting with "#").
	SkipHidden *bool
	// Clean returns only clean (non-empty, non-hidden) items.
	Clean *bool
	// Unwind expands these fields (each array element becomes a separate item).
	Unwind []string
	// Flatten flattens these nested fields into dot-notation keys.
	Flatten []string
	// View selects a predefined dataset view for field selection.
	View *string
	// Simplified returns simplified (flattened, cleaned) items.
	Simplified *bool
	// SkipFailedPages skips items that come from failed pages.
	SkipFailedPages *bool
	// Signature is a pre-shared URL signature granting access without an API token.
	Signature *string
}

func (o DatasetListItemsOptions) apply(q *QueryParams) {
	q.AddInt("offset", o.Offset).
		AddInt("limit", o.Limit).
		AddBool("desc", o.Desc).
		AddCSV("fields", o.Fields).
		AddCSV("outputFields", o.OutputFields).
		AddCSV("omit", o.Omit).
		AddBool("skipEmpty", o.SkipEmpty).
		AddBool("skipHidden", o.SkipHidden).
		AddBool("clean", o.Clean).
		AddCSV("unwind", o.Unwind).
		AddCSV("flatten", o.Flatten).
		AddString("view", o.View).
		AddBool("simplified", o.Simplified).
		AddBool("skipFailedPages", o.SkipFailedPages).
		AddString("signature", o.Signature)
}

// DownloadItemsFormat is an output format for [DatasetClient.DownloadItems].
type DownloadItemsFormat string

const (
	// FormatJSON serializes items as a JSON array.
	FormatJSON DownloadItemsFormat = "json"
	// FormatJSONL serializes items as newline-delimited JSON.
	FormatJSONL DownloadItemsFormat = "jsonl"
	// FormatCSV serializes items as comma-separated values.
	FormatCSV DownloadItemsFormat = "csv"
	// FormatXLSX serializes items as a Microsoft Excel (XLSX) workbook.
	FormatXLSX DownloadItemsFormat = "xlsx"
	// FormatXML serializes items as XML.
	FormatXML DownloadItemsFormat = "xml"
	// FormatRSS serializes items as an RSS feed.
	FormatRSS DownloadItemsFormat = "rss"
	// FormatHTML serializes items as an HTML table.
	FormatHTML DownloadItemsFormat = "html"
)

// DatasetDownloadOptions adds format-specific options for [DatasetClient.DownloadItems] on
// top of the shared item filtering/projection options.
type DatasetDownloadOptions struct {
	// Items holds the shared filtering/projection options.
	Items DatasetListItemsOptions
	// Attachment sets Content-Disposition: attachment on the response.
	Attachment *bool
	// Bom prepends a UTF-8 BOM (useful for Excel-compatible CSV).
	Bom *bool
	// Delimiter is the CSV field delimiter (default ",").
	Delimiter *string
	// SkipHeaderRow omits the CSV header row.
	SkipHeaderRow *bool
	// XMLRoot is the name of the root XML element (default "items").
	XMLRoot *string
	// XMLRow is the name of the per-item XML element (default "item").
	XMLRow *string
	// FeedTitle is the title used for RSS/Atom feed exports.
	FeedTitle *string
	// FeedDescription is the description used for RSS/Atom feed exports.
	FeedDescription *string
}

func (o DatasetDownloadOptions) apply(q *QueryParams) {
	o.Items.apply(q)
	q.AddBool("attachment", o.Attachment).
		AddBool("bom", o.Bom).
		AddString("delimiter", o.Delimiter).
		AddBool("skipHeaderRow", o.SkipHeaderRow).
		AddString("xmlRoot", o.XMLRoot).
		AddString("xmlRow", o.XMLRow).
		AddString("feedTitle", o.FeedTitle).
		AddString("feedDescription", o.FeedDescription)
}

// DatasetClient is a client for a specific dataset (and run-nested variants).
type DatasetClient struct {
	ctx *resourceContext
}

func newDatasetClient(hc *httpClient, baseURL, resourcePath, id string) *DatasetClient {
	return &DatasetClient{ctx: newSingleContext(hc, baseURL, resourcePath, id)}
}

// newDatasetNestedClient creates a dataset client for a run's default dataset (no ID;
// nested path only, e.g. /v2/actor-runs/{runId}/dataset).
func newDatasetNestedClient(hc *httpClient, base, subPath string) *DatasetClient {
	return &DatasetClient{ctx: newCollectionContext(hc, base, subPath)}
}

// withPublicBase sets the public origin used when building shareable URLs.
func (c *DatasetClient) withPublicBase(publicBaseURL string) *DatasetClient {
	c.ctx = c.ctx.withPublicOrigin(publicBaseURL)
	return c
}

// Get fetches the dataset metadata. The bool reports whether it exists.
func (c *DatasetClient) Get(ctx context.Context) (Dataset, bool, error) {
	return getResource[Dataset](ctx, c.ctx, "", NewQueryParams())
}

// Update updates the dataset metadata (e.g. name, title) and returns the updated object.
func (c *DatasetClient) Update(ctx context.Context, newFields any) (Dataset, error) {
	return updateResource[Dataset](ctx, c.ctx, "", newFields)
}

// Delete deletes the dataset.
func (c *DatasetClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// ListDatasetItems lists one page of the dataset items, decoding each item into T, for example a
// struct or json.RawMessage. The pagination headers fill the returned [PaginationList].
func ListDatasetItems[T any](ctx context.Context, c *DatasetClient, options DatasetListItemsOptions) (PaginationList[T], error) {
	var result PaginationList[T]
	params := NewQueryParams()
	options.apply(params)
	if c.ctx.idErr != nil {
		return result, c.ctx.idErr
	}
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("items"))
	resp, err := c.ctx.http.call(ctx, http.MethodGet, url, nil, "", c.ctx.longTimeout())
	if err != nil {
		return result, err
	}
	var items []T
	if err := json.Unmarshal(resp.body, &items); err != nil {
		return result, err
	}
	returned := int64(len(items))
	result.Items = items
	result.Total = headerInt(resp.header, "X-Apify-Pagination-Total", returned)
	result.Offset = headerInt(resp.header, "X-Apify-Pagination-Offset", 0)
	result.Limit = headerInt(resp.header, "X-Apify-Pagination-Limit", returned)
	// The API applies Offset and Limit to the rows before the filters and Unwind reshape them, so
	// the items returned can differ in number from the rows scanned. Count is the scanned number,
	// and the iterators advance by it. It is derived from the item count, which lags a fresh push,
	// so the larger of it and the items returned is used.
	result.Count = maxInt64(headerInt(resp.header, "X-Apify-Pagination-Count", returned), returned)
	if options.Desc != nil {
		result.Desc = *options.Desc
	}
	return result, nil
}

// ListItems lists items from the dataset, decoding each into a generic json.RawMessage.
// For typed decoding use [ListDatasetItems].
func (c *DatasetClient) ListItems(ctx context.Context, options DatasetListItemsOptions) (PaginationList[json.RawMessage], error) {
	return ListDatasetItems[json.RawMessage](ctx, c, options)
}

// IterateDatasetItems returns a lazy iterator over the dataset's items, decoding each into T and
// fetching pages on demand. The options' Limit caps the total number of items yielded (unset means
// all); the per-page size is chunkSize (nil for the server default).
func IterateDatasetItems[T any](c *DatasetClient, options DatasetListItemsOptions, chunkSize *int64) *ListIterator[T] {
	return newListIterator(options.Limit, chunkSize, offsetVal(options.Offset), func(ctx context.Context, offset, limit int64) (PaginationList[T], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return ListDatasetItems[T](ctx, c, opts)
	})
}

// IterateItems returns a lazy iterator over the dataset's items, decoding each into a generic
// json.RawMessage. For typed decoding use [IterateDatasetItems]. See [IterateDatasetItems] for
// how the options' Limit (total cap) and chunkSize (page size) are interpreted.
func (c *DatasetClient) IterateItems(options DatasetListItemsOptions, chunkSize *int64) *ListIterator[json.RawMessage] {
	return IterateDatasetItems[json.RawMessage](c, options, chunkSize)
}

// DownloadItems downloads dataset items serialized in the given format, returning the raw
// bytes. Unlike ListItems (parsed items), this returns the items already serialized to JSON,
// CSV, XLSX, XML, RSS or HTML — useful for exporting. For a large export, see
// [DatasetClient.DownloadItemsStream], which streams the bytes instead of buffering all of them
// in memory.
func (c *DatasetClient) DownloadItems(ctx context.Context, format DownloadItemsFormat, options DatasetDownloadOptions) ([]byte, error) {
	params := NewQueryParams()
	fmtStr := string(format)
	params.AddString("format", &fmtStr)
	options.apply(params)
	url := c.ctx.mergedParams(params).applyToURL(c.ctx.subURL("items"))
	resp, err := c.ctx.http.call(ctx, http.MethodGet, url, nil, "", c.ctx.longTimeout())
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

// DownloadItemsStream is [DatasetClient.DownloadItems], but returns the export unread, so that a
// large export can be processed in chunks. The caller closes it. The request is not retried.
func (c *DatasetClient) DownloadItemsStream(ctx context.Context, format DownloadItemsFormat, options DatasetDownloadOptions) (io.ReadCloser, error) {
	params := NewQueryParams()
	fmtStr := string(format)
	params.AddString("format", &fmtStr)
	options.apply(params)
	stream, _, err := streamRaw(ctx, c.ctx, "items", params, false)
	return stream, err
}

// PushItems pushes items to the dataset. items must serialize to a JSON object or an array of
// objects, otherwise PushItems returns an error before sending a request. A string is serialized
// as a JSON string, so pass a [json.RawMessage] for JSON that is already serialized.
//
// A push that times out is not retried, since the API may have stored the items already.
func (c *DatasetClient) PushItems(ctx context.Context, items any) error {
	body, err := json.Marshal(items)
	if err != nil {
		return err
	}
	if err := checkItemsShape(body); err != nil {
		return err
	}
	if c.ctx.idErr != nil {
		return c.ctx.idErr
	}
	url := c.ctx.mergedParams(NewQueryParams()).applyToURL(c.ctx.subURL("items"))
	_, err = c.ctx.withoutTimeoutRetry().http.call(ctx, http.MethodPost, url, body, contentTypeJSONCharset, c.ctx.mediumTimeout())
	return err
}

// checkItemsShape reports whether body, a JSON document, is an object or an array of objects.
func checkItemsShape(body []byte) error {
	trimmed := bytes.TrimSpace(body)
	switch {
	case bytes.HasPrefix(trimmed, []byte("{")):
		return nil
	case bytes.HasPrefix(trimmed, []byte("[")):
		var elements []json.RawMessage
		if err := json.Unmarshal(trimmed, &elements); err != nil {
			return err
		}
		for i, element := range elements {
			if !bytes.HasPrefix(bytes.TrimSpace(element), []byte("{")) {
				return fmt.Errorf("PushItems: item %d is not a JSON object; the API only accepts objects", i)
			}
		}
		return nil
	default:
		return errors.New("PushItems: items must be a JSON object or an array of objects")
	}
}

// GetStatistics returns statistical information about the dataset.
//
// A 404 here always means the dataset itself is gone (the same way a plain Get would report
// it), so it is returned as an error rather than swallowed: unlike Get, there is no
// presence-reporting bool.
func (c *DatasetClient) GetStatistics(ctx context.Context) (json.RawMessage, error) {
	resp, err := getRawRequired(ctx, c.ctx, "statistics", NewQueryParams(), c.ctx.shortTimeout())
	if err != nil {
		return nil, err
	}
	return parseDataEnvelope[json.RawMessage](resp.body)
}

// CreateItemsPublicURL returns a public URL of the dataset items as JSON. For another format, use
// [DatasetClient.CreateItemsPublicURLWithFormat].
//
// If the dataset has a URL signing secret key, the URL is signed, so it works without a token.
// expiresInSecs is the validity of a signed URL, nil for no expiry, and it must be positive.
func (c *DatasetClient) CreateItemsPublicURL(ctx context.Context, options DatasetListItemsOptions, expiresInSecs *int64) (string, error) {
	return c.CreateItemsPublicURLWithFormat(ctx, options, expiresInSecs, "")
}

// CreateItemsPublicURLWithFormat is like [DatasetClient.CreateItemsPublicURL], but lets the
// caller choose the format the items are served in (JSON, CSV, XLSX, ...) via a `format` query
// parameter on the generated URL. An empty format defaults to "json", matching the endpoint's
// own default. As there, expiresInSecs must be positive when set.
func (c *DatasetClient) CreateItemsPublicURLWithFormat(ctx context.Context, options DatasetListItemsOptions, expiresInSecs *int64, format DownloadItemsFormat) (string, error) {
	if err := validateExpiresInSecs(expiresInSecs); err != nil {
		return "", err
	}
	params := NewQueryParams()
	options.apply(params)
	if format != "" {
		fmtStr := string(format)
		params.AddString("format", &fmtStr)
	}

	// The "long" timeout tier is used for this fetch.
	dataset, present, err := getResourceWithTimeout[Dataset](ctx, c.ctx, "", NewQueryParams(), c.ctx.longTimeout())
	if err != nil {
		return "", err
	}
	if present {
		if secret := extractString(dataset.Extra, urlSigningSecretExtraKey); secret != "" {
			sig := signStorageContent(secret, dataset.ID, expiresInSecs)
			params.AddString("signature", &sig)
		}
	}
	return c.ctx.mergedParams(params).applyToURL(c.ctx.publicURL("items")), nil
}

// headerInt parses an integer HTTP header value, returning fallback if absent or invalid.
func headerInt(h http.Header, name string, fallback int64) int64 {
	v := h.Get(name)
	if v == "" {
		return fallback
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return fallback
	}
	return n
}

// extractString reads a string field from an Extra map, returning "" if absent or not a
// string.
func extractString(extra Extra, key string) string {
	raw, ok := extra[key]
	if !ok {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return ""
	}
	return s
}
