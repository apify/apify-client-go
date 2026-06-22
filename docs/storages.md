# Storages

The platform offers three storage types: datasets, key-value stores, and request queues.
Each is reachable both as a top-level resource and as a run's default storage
(`client.Run(id).Dataset()` etc.).

## Datasets

Collection: `client.Datasets()` — `List(ctx, StorageListOptions)`,
`GetOrCreate(ctx, name string)` (empty name → unnamed dataset).

Single dataset: `client.Dataset(id)`:

| Method | Description |
| --- | --- |
| `Get / Update / Delete(ctx)` | CRUD. |
| `ListItems(ctx, DatasetListItemsOptions) (PaginationList[json.RawMessage], error)` | Read items. |
| `PushItems(ctx, items any) error` | Append one item or a slice of items. |
| `DownloadItems(ctx, DownloadItemsFormat, DatasetDownloadOptions) ([]byte, error)` | Export items (JSON, CSV, XLSX, XML, RSS, HTML). |
| `GetStatistics(ctx) (json.RawMessage, bool, error)` | Dataset statistics. |
| `CreateItemsPublicURL(ctx, DatasetListItemsOptions, expiresInSecs *int64) (string, error)` | Signed public items URL. |

`DatasetListItemsOptions` controls filtering, projection, and pagination of the items
(`ListItems`, `ListDatasetItems`, and the `Items` field of `DatasetDownloadOptions`). All
fields are optional; the slice fields default to empty and the pointer fields to unset.

| Field | Type | Meaning |
|---|---|---|
| `Offset` | `*int64` | Number of items to skip. |
| `Limit` | `*int64` | Maximum number of items to return. |
| `Desc` | `*bool` | Return items newest-first. |
| `Fields` | `[]string` | Restrict the output to these fields. |
| `OutputFields` | `[]string` | Positionally rename the fields selected by `Fields` (requires `Fields`). |
| `Omit` | `[]string` | Exclude these fields from the output. |
| `SkipEmpty` | `*bool` | Skip empty items. |
| `SkipHidden` | `*bool` | Skip hidden fields (those starting with `#`). |
| `Clean` | `*bool` | Return only clean (non-empty, non-hidden) items. |
| `Unwind` | `[]string` | Expand these fields (each array element becomes a separate item). |
| `Flatten` | `[]string` | Flatten these nested fields into dot-notation keys. |
| `View` | `*string` | Select a predefined dataset view (a named field-selection/transform defined in the Actor's dataset schema) by name. |
| `Simplified` | `*bool` | Return simplified (flattened, cleaned) items. |
| `SkipFailedPages` | `*bool` | Skip items that come from failed pages. |
| `Signature` | `*string` | Pre-shared URL signature granting access without an API token. |

For typed item decoding use the generic helper
`apify.ListDatasetItems[T](ctx, dataset *DatasetClient, opts DatasetListItemsOptions) (PaginationList[T], error)`.
`ListItems` returns each item as `json.RawMessage`; this helper decodes every item into your
own type `T` instead. Pass the `*DatasetClient` you get from `client.Dataset(id)` as the
`dataset` argument:

```go
// A struct matching the shape of your dataset items.
type Result struct {
	Title string `json:"title"`
}

page, err := apify.ListDatasetItems[Result](ctx, client.Dataset("DATASET_ID"), apify.DatasetListItemsOptions{
	Limit: apify.Ptr(int64(100)),
})
if err != nil {
	log.Fatal(err)
}
for _, item := range page.Items {
	fmt.Println(item.Title)
}
```

`DownloadItems` takes a `DownloadItemsFormat`. The exported constants are:

| Constant | Value |
| --- | --- |
| `apify.FormatJSON` | `json` |
| `apify.FormatJSONL` | `jsonl` |
| `apify.FormatCSV` | `csv` |
| `apify.FormatXLSX` | `xlsx` |
| `apify.FormatXML` | `xml` |
| `apify.FormatRSS` | `rss` |
| `apify.FormatHTML` | `html` |

`DatasetDownloadOptions` adds format-specific export options on top of the shared item
filtering/projection options. All fields are optional:

| Field | Type | Meaning |
|---|---|---|
| `Items` | `DatasetListItemsOptions` | The shared filtering/projection options (see above). |
| `Attachment` | `*bool` | Set `Content-Disposition: attachment` on the response. |
| `Bom` | `*bool` | Prepend a UTF-8 BOM (useful for Excel-compatible CSV). |
| `Delimiter` | `*string` | CSV field delimiter (default `,`). |
| `SkipHeaderRow` | `*bool` | Omit the CSV header row. |
| `XMLRoot` | `*string` | Name of the root XML element (default `items`). |
| `XMLRow` | `*string` | Name of the per-item XML element (default `item`). |
| `FeedTitle` | `*string` | Title used for RSS/Atom feed exports. |
| `FeedDescription` | `*string` | Description used for RSS/Atom feed exports. |

```go
ds, _ := client.Datasets().GetOrCreate(ctx, "")
defer client.Dataset(ds.ID).Delete(ctx)

_ = client.Dataset(ds.ID).PushItems(ctx, []map[string]any{{"x": 1}, {"x": 2}})
page, _ := client.Dataset(ds.ID).ListItems(ctx, apify.DatasetListItemsOptions{Limit: apify.Ptr(int64(100))})
csv, _ := client.Dataset(ds.ID).DownloadItems(ctx, apify.FormatCSV, apify.DatasetDownloadOptions{Bom: apify.Ptr(true)})
```

## Key-value stores

Collection: `client.KeyValueStores()` — `List`, `GetOrCreate`.

Single store: `client.KeyValueStore(id)`:

| Method | Description |
| --- | --- |
| `Get / Update / Delete(ctx)` | CRUD. |
| `ListKeys(ctx, ListKeysOptions) (KeyValueStoreKeysPage, error)` | List keys. |
| `GetRecord(ctx, key) (*KeyValueStoreRecord, bool, error)` | Read a record. |
| `GetRecordWithOptions(ctx, key, GetRecordOptions)` | Read with options. |
| `SetRecordRaw(ctx, key, value []byte, contentType string) error` | Write raw bytes. |
| `SetRecordJSON(ctx, key, value any) error` | Write JSON. |
| `RecordExists(ctx, key) (bool, error)` | HEAD check. |
| `DeleteRecord(ctx, key) error` | Delete a record. |
| `GetRecords(ctx, GetRecordsOptions) ([]byte, error)` | Download all records as a ZIP. |
| `GetRecordPublicURL(ctx, key) (string, error)` | Signed public record URL. |
| `CreateKeysPublicURL(ctx, expiresInSecs *int64) (string, error)` | Signed public key-list URL. |

Option structs (all fields optional):

| Struct | Field | Type | Meaning |
|---|---|---|---|
| `ListKeysOptions` | `Limit` | `*int64` | Maximum number of keys to return. |
| | `ExclusiveStartKey` | `*string` | List keys after this one (pagination). |
| | `Prefix` | `*string` | Restrict the listing to keys with this prefix. |
| | `Collection` | `*string` | Restrict the listing to a named collection of keys. |
| | `Signature` | `*string` | Pre-shared URL signature (access without a token). |
| `GetRecordOptions` | `Attachment` | `*bool` | Control the `Content-Disposition: attachment` behaviour. |
| | `Signature` | `*string` | Pre-shared URL signature (access without a token). |
| `GetRecordsOptions` | `Collection` | `*string` | Restrict the download to a named collection. |
| | `Prefix` | `*string` | Restrict the download to records with this key prefix. |
| | `Signature` | `*string` | Pre-shared URL signature (access without a token). |

Return types:

`GetRecord`/`GetRecordWithOptions` return a `*KeyValueStoreRecord`:

| Field | Type | Meaning |
|---|---|---|
| `Key` | `string` | The record key. |
| `Value` | `[]byte` | The raw record bytes (decode according to `ContentType`). |
| `ContentType` | `string` | The record's MIME type, as reported by the API. |

`ListKeys` returns a `KeyValueStoreKeysPage`:

| Field | Type | Meaning |
|---|---|---|
| `Limit` | `int64` | Maximum number of keys requested. |
| `IsTruncated` | `bool` | Whether more keys are available. |
| `ExclusiveStartKey` | `string` | The key the listing started after. |
| `NextExclusiveStartKey` | `string` | Key to pass to fetch the next page. |
| `Items` | `[]KeyValueStoreKey` | The listed keys. |

Each `KeyValueStoreKey` element has `Key` (`string`) and `Size` (`int64`, the record size in
bytes), plus an `Extra` (`map[string]json.RawMessage`) catch-all.

```go
store, _ := client.KeyValueStores().GetOrCreate(ctx, "")
defer client.KeyValueStore(store.ID).Delete(ctx)

_ = client.KeyValueStore(store.ID).SetRecordJSON(ctx, "OUTPUT", map[string]any{"answer": 42})
rec, ok, _ := client.KeyValueStore(store.ID).GetRecord(ctx, "OUTPUT")
if ok {
	fmt.Println(string(rec.Value))
}
```

## Request queues

Collection: `client.RequestQueues()` — `List`, `GetOrCreate`.

Single queue: `client.RequestQueue(id)`:

| Method | Description |
| --- | --- |
| `Get / Update / Delete(ctx)` | CRUD. |
| `ListHead(ctx, limit *int64) (RequestQueueHead, error)` | Requests at the front. |
| `AddRequest(ctx, RequestQueueRequest, forefront bool) (RequestQueueOperationInfo, error)` | Add a request. |
| `GetRequest(ctx, id) (*RequestQueueRequest, bool, error)` | Read a request. |
| `UpdateRequest(ctx, RequestQueueRequest, forefront bool)` | Update a request (by its `ID`). |
| `DeleteRequest(ctx, id) error` | Delete a request. |
| `ListRequests(ctx, ListRequestsOptions)` | Paginated list. |
| `PaginateRequests(pageLimit *int64) *RequestQueueRequestsIterator` | Lazy iterator over all requests. |
| `BatchAddRequests(ctx, []RequestQueueRequest, forefront) (BatchAddResult, error)` | Add many requests (auto-chunked at 25/call). |
| `BatchDeleteRequests(ctx, requests any)` | Delete many requests. |
| `ListAndLockHead / ProlongRequestLock / DeleteRequestLock / UnlockRequests(ctx, ...)` | Locking. |
| `WithClientKey(key string) *RequestQueueClient` | Pin a stable client key (required to unlock own locks). |

`RequestQueueRequest` is the request payload/record. `URL` is required; `ID` is assigned by the
API (omit it on create):

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique request ID (assigned by the API; omitted on create). |
| `URL` | `string` | The request URL (required). |
| `UniqueKey` | `string` | Deduplication key for the request. |
| `Method` | `string` | HTTP method (e.g. `GET`, `POST`). |
| `UserData` | `json.RawMessage` | Arbitrary user-attached metadata. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

`ListRequestsOptions` (all fields optional):

| Field | Type | Meaning |
|---|---|---|
| `Limit` | `*int64` | Maximum number of requests to return. |
| `ExclusiveStartID` | `*string` | List requests after this ID. |
| `Cursor` | `*string` | Opaque pagination cursor (alternative to `ExclusiveStartID`). |
| `Filter` | `*string` | Restrict the listing. The API accepts only `"locked"` or `"pending"`. |

Return types:

`ListHead`/`ListAndLockHead` return a `RequestQueueHead`:

| Field | Type | Meaning |
|---|---|---|
| `Limit` | `int64` | Maximum number of requests requested. |
| `HadMultipleClients` | `bool` | Whether multiple clients have accessed the queue. |
| `Items` | `[]RequestQueueRequest` | The requests at the head of the queue. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

`AddRequest`/`UpdateRequest` return a `RequestQueueOperationInfo`:

| Field | Type | Meaning |
|---|---|---|
| `RequestID` | `string` | ID of the affected request. |
| `WasAlreadyPresent` | `bool` | Whether the request was already in the queue. |
| `WasAlreadyHandled` | `bool` | Whether the request had already been handled. |

`BatchAddRequests` returns a `BatchAddResult`:

| Field | Type | Meaning |
|---|---|---|
| `ProcessedRequests` | `[]RequestQueueOperationInfo` | Requests the API successfully added. |
| `UnprocessedRequests` | `[]RequestQueueRequest` | Requests the API did not process. |

```go
rq, _ := client.RequestQueues().GetOrCreate(ctx, "")
defer client.RequestQueue(rq.ID).Delete(ctx)

queue := client.RequestQueue(rq.ID)
_, _ = queue.AddRequest(ctx, apify.RequestQueueRequest{URL: "https://example.com", UniqueKey: "example"}, false)

it := queue.PaginateRequests(apify.Ptr(int64(100)))
for {
	req, err := it.Next(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if req == nil {
		break
	}
	fmt.Println(req.URL)
}
```

`StorageListOptions` is shared by the three storage collections' `List` methods. All fields are
optional:

| Field | Type | Meaning |
|---|---|---|
| `Offset` | `*int64` | Number of items to skip from the start of the list. |
| `Limit` | `*int64` | Maximum number of items to return. |
| `Desc` | `*bool` | Return items newest-first. |
| `Unnamed` | `*bool` | Include unnamed storages in the result. |
| `Ownership` | `*string` | Filter by ownership; accepted values are `OWNED` and `ACCESSIBLE`. |
