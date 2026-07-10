package apify

import (
	"context"
	"encoding/json"
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

// ListDatasetItems lists a single page of items from the dataset, decoding each into T
// (e.g. json.RawMessage or a struct).
//
// The dataset items endpoint returns a bare JSON array (not a data envelope) and reports
// pagination via X-Apify-Pagination-* headers, which are surfaced in the returned
// [PaginationList].
func ListDatasetItems[T any](ctx context.Context, c *DatasetClient, options DatasetListItemsOptions) (PaginationList[T], error) {
	var result PaginationList[T]
	params := NewQueryParams()
	options.apply(params)
	url := params.applyToURL(c.ctx.subURL("items"))
	resp, err := c.ctx.http.call(ctx, http.MethodGet, url, nil, "", defaultRequestTimeout)
	if err != nil {
		return result, err
	}
	var items []T
	if err := json.Unmarshal(resp.body, &items); err != nil {
		return result, err
	}
	count := int64(len(items))
	result.Items = items
	result.Count = count
	result.Total = headerInt(resp.header, "X-Apify-Pagination-Total", count)
	result.Offset = headerInt(resp.header, "X-Apify-Pagination-Offset", 0)
	result.Limit = headerInt(resp.header, "X-Apify-Pagination-Limit", count)
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

// IterateDatasetItems returns a lazy iterator over the dataset's items, decoding each into T
// and fetching pages on demand. The options' Limit caps the total number of items yielded
// (unset means all); the per-page size is chunkSize (nil for the server default). Mirrors the
// reference client's iterable listItems().
func IterateDatasetItems[T any](c *DatasetClient, options DatasetListItemsOptions, chunkSize *int64) *ListIterator[T] {
	return newListIterator(options.Limit, chunkSize, func(ctx context.Context, offset, limit int64) (PaginationList[T], error) {
		opts := options
		opts.Offset = &offset
		opts.Limit = pageLimitPtr(limit)
		return ListDatasetItems[T](ctx, c, opts)
	})
}

// IterateItems returns a lazy iterator over the dataset's items, decoding each into a generic
// json.RawMessage. For typed decoding use [IterateDatasetItems]. See IterateDatasetItems for
// how the options' Limit (total cap) and chunkSize (page size) are interpreted.
func (c *DatasetClient) IterateItems(options DatasetListItemsOptions, chunkSize *int64) *ListIterator[json.RawMessage] {
	return IterateDatasetItems[json.RawMessage](c, options, chunkSize)
}

// DownloadItems downloads dataset items serialized in the given format, returning the raw
// bytes. Unlike ListItems (parsed items), this returns the items already serialized to JSON,
// CSV, XLSX, XML, RSS or HTML — useful for exporting.
func (c *DatasetClient) DownloadItems(ctx context.Context, format DownloadItemsFormat, options DatasetDownloadOptions) ([]byte, error) {
	params := NewQueryParams()
	fmtStr := string(format)
	params.AddString("format", &fmtStr)
	options.apply(params)
	url := params.applyToURL(c.ctx.subURL("items"))
	resp, err := c.ctx.http.call(ctx, http.MethodGet, url, nil, "", defaultRequestTimeout)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

// PushItems pushes one or more items to the dataset. items must serialize to a JSON object
// or an array of objects.
func (c *DatasetClient) PushItems(ctx context.Context, items any) error {
	body, err := json.Marshal(items)
	if err != nil {
		return err
	}
	_, err = c.ctx.http.call(ctx, http.MethodPost, c.ctx.subURL("items"), body, contentTypeJSONCharset, defaultRequestTimeout)
	return err
}

// GetStatistics returns statistical information about the dataset, or (nil, false, nil) if
// unavailable.
func (c *DatasetClient) GetStatistics(ctx context.Context) (json.RawMessage, bool, error) {
	resp, err := getRaw(ctx, c.ctx, "statistics", NewQueryParams())
	if err != nil {
		return nil, false, err
	}
	if resp == nil {
		return nil, false, nil
	}
	data, err := parseDataEnvelope[json.RawMessage](resp.body)
	if err != nil {
		return nil, false, err
	}
	return data, true, nil
}

// CreateItemsPublicURL builds a public URL for downloading this dataset's items.
//
// It mirrors the reference client's createItemsPublicUrl: it fetches the dataset, and if the
// dataset exposes a URL-signing secret key (i.e. it is private), appends an HMAC-SHA256
// signature so the URL grants access without an API token. expiresInSecs optionally bounds
// the validity of a signed URL (nil for non-expiring). The URL is built from the configured
// public base URL.
func (c *DatasetClient) CreateItemsPublicURL(ctx context.Context, options DatasetListItemsOptions, expiresInSecs *int64) (string, error) {
	params := NewQueryParams()
	options.apply(params)

	dataset, present, err := c.Get(ctx)
	if err != nil {
		return "", err
	}
	if present {
		if secret := extractString(dataset.Extra, "urlSigningSecretKey"); secret != "" {
			sig := signStorageContent(secret, dataset.ID, expiresInSecs)
			params.AddString("signature", &sig)
		}
	}
	return params.applyToURL(c.ctx.publicURL("items")), nil
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
