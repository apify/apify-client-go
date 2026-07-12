# Storages

The platform offers three storage types: datasets, key-value stores, and request queues.
Each is reachable both as a top-level resource and as a run's default storage
(`client.Run(id).Dataset()` etc.).

## Datasets

Collection: `client.Datasets()` — `List(ctx, StorageListOptions) (PaginationList[Dataset], error)`,
`Iterate(StorageListOptions, chunkSize *int64) *ListIterator[Dataset]` (lazy iterator; `Limit` caps total, `chunkSize` is page size),
`GetOrCreate(ctx, name string) (Dataset, error)` (empty name → unnamed dataset).

Single dataset: `client.Dataset(id)`:

| Method | Description |
| --- | --- |
| `Get(ctx) (Dataset, bool, error)` | Fetch the dataset (`false` if it does not exist). |
| `Update(ctx, newFields any) (Dataset, error)` | Update the dataset. |
| `Delete(ctx) error` | Delete the dataset. |
| `ListItems(ctx, DatasetListItemsOptions) (PaginationList[json.RawMessage], error)` | Read items. |
| `IterateItems(DatasetListItemsOptions, chunkSize *int64) *ListIterator[json.RawMessage]` | Lazy iterator over items (`Limit` caps total, `chunkSize` is page size); `IterateDatasetItems[T]` decodes into your type. |
| `PushItems(ctx, items any) error` | Append one item or a slice of items. |
| `DownloadItems(ctx, DownloadItemsFormat, DatasetDownloadOptions) ([]byte, error)` | Export items (JSON, JSONL, CSV, XLSX, XML, RSS, HTML — see the format constants below). |
| `GetStatistics(ctx) (json.RawMessage, bool, error)` | Dataset statistics. |
| `CreateItemsPublicURL(ctx, DatasetListItemsOptions, expiresInSecs *int64) (string, error)` | Signed public items URL. |

The `Dataset` value returned by `Get`/`GetOrCreate`/`Update` and listed by `List`/`Iterate`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique dataset ID. |
| `Name` | `string` | Dataset name (empty for unnamed datasets). |
| `UserID` | `string` | ID of the user who owns the dataset. |
| `CreatedAt` | `*time.Time` | When the dataset was created. |
| `ModifiedAt` | `*time.Time` | When the dataset was last modified. |
| `ItemCount` | `int64` | Number of items currently stored. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

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

For lazy, typed iteration over items across pages use the generic helper
`apify.IterateDatasetItems[T](dataset *DatasetClient, opts DatasetListItemsOptions, chunkSize *int64) *ListIterator[T]`.
It is the typed counterpart of the `IterateItems` method: `Limit` caps the total number of
items yielded and `chunkSize` sets the per-page size, but each item is decoded into your type
`T` instead of `json.RawMessage`. Advance it with `Next(ctx)`, which returns `nil` when the
iteration is exhausted:

```go
// A struct matching the shape of your dataset items.
type Result struct {
	Title string `json:"title"`
}

// Limit caps the total items yielded; the last argument is the per-page size.
it := apify.IterateDatasetItems[Result](client.Dataset("DATASET_ID"),
	apify.DatasetListItemsOptions{Limit: apify.Ptr(int64(1000))},
	apify.Ptr(int64(100)),
)
for {
	item, err := it.Next(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if item == nil {
		break
	}
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

Collection: `client.KeyValueStores()` — `List(ctx, StorageListOptions) (PaginationList[KeyValueStore], error)`, `Iterate(StorageListOptions, chunkSize *int64) *ListIterator[KeyValueStore]` (`Limit` caps total, `chunkSize` is page size), `GetOrCreate(ctx, name string) (KeyValueStore, error)`.

Single store: `client.KeyValueStore(id)`:

| Method | Description |
| --- | --- |
| `Get(ctx) (KeyValueStore, bool, error)` | Fetch the store (`false` if it does not exist). |
| `Update(ctx, newFields any) (KeyValueStore, error)` | Update the store. |
| `Delete(ctx) error` | Delete the store. |
| `ListKeys(ctx, ListKeysOptions) (KeyValueStoreKeysPage, error)` | List one page of keys. |
| `IterateKeys(ListKeysOptions, chunkSize *int64) *KeyValueStoreKeysIterator` | Lazy iterator over keys (cursor-based). `Limit` caps the total yielded; `chunkSize` is the page size; `Prefix`/`Collection`/`Signature` filter every page; `ExclusiveStartKey` sets where to start. |
| `GetRecord(ctx, key) (*KeyValueStoreRecord, bool, error)` | Read a record. |
| `GetRecordWithOptions(ctx, key, GetRecordOptions)` | Read with options. |
| `SetRecordRaw(ctx, key, value []byte, contentType string) error` | Write raw bytes. |
| `SetRecordJSON(ctx, key, value any) error` | Write JSON. |
| `RecordExists(ctx, key) (bool, error)` | HEAD check. |
| `DeleteRecord(ctx, key) error` | Delete a record. |
| `GetRecords(ctx, GetRecordsOptions) ([]byte, error)` | Download all records as a ZIP. |
| `GetRecordPublicURL(ctx, key) (string, error)` | Signed public record URL. |
| `CreateKeysPublicURL(ctx, expiresInSecs *int64) (string, error)` | Signed public key-list URL. |

The `KeyValueStore` value returned by `Get`/`GetOrCreate`/`Update` and listed by `List`/`Iterate`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique store ID. |
| `Name` | `string` | Store name (empty for unnamed stores). |
| `UserID` | `string` | ID of the user who owns the store. |
| `CreatedAt` | `*time.Time` | When the store was created. |
| `ModifiedAt` | `*time.Time` | When the store was last modified. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

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

// Lazily iterate over every key in the store (cursor-based paging is handled internally).
keys := client.KeyValueStore(store.ID).IterateKeys(apify.ListKeysOptions{}, nil)
for {
	key, err := keys.Next(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if key == nil {
		break
	}
	fmt.Println(key.Key, key.Size)
}
```

## Request queues

Collection: `client.RequestQueues()` — `List(ctx, StorageListOptions) (PaginationList[RequestQueue], error)`, `Iterate(StorageListOptions, chunkSize *int64) *ListIterator[RequestQueue]` (`Limit` caps total, `chunkSize` is page size), `GetOrCreate(ctx, name string) (RequestQueue, error)`.

Single queue: `client.RequestQueue(id)`:

| Method | Description |
| --- | --- |
| `Get(ctx) (RequestQueue, bool, error)` | Fetch the queue (`false` if it does not exist). |
| `Update(ctx, newFields any) (RequestQueue, error)` | Update the queue. |
| `Delete(ctx) error` | Delete the queue. |
| `ListHead(ctx, limit *int64) (RequestQueueHead, error)` | Requests at the front. |
| `AddRequest(ctx, RequestQueueRequest, forefront bool) (RequestQueueOperationInfo, error)` | Add a request. |
| `GetRequest(ctx, id) (*RequestQueueRequest, bool, error)` | Read a request. |
| `UpdateRequest(ctx, RequestQueueRequest, forefront bool) (RequestQueueOperationInfo, error)` | Update a request (by its `ID`). |
| `DeleteRequest(ctx, id) error` | Delete a request. |
| `ListRequests(ctx, ListRequestsOptions) (json.RawMessage, error)` | Paginated list. Returns the raw API response (`{ "items": [...], "nextCursor": ... }`); unmarshal into your own type, or use `PaginateRequests` for typed iteration. |
| `PaginateRequests(pageLimit *int64) *RequestQueueRequestsIterator` | Lazy iterator over all requests (yields `*RequestQueueRequest`). |
| `BatchAddRequests(ctx, []RequestQueueRequest, forefront) (BatchAddResult, error)` | Add many requests (auto-chunked at 25/call). |
| `BatchDeleteRequests(ctx, requests any) (json.RawMessage, error)` | Delete many requests in one call. `requests` is the JSON-marshalable batch payload — a slice in which each element identifies one request by **either** its `id` **or** its `uniqueKey` (e.g. `[]map[string]string{{"uniqueKey": "..."}}` or a slice of `RequestQueueRequest`), matching the reference client's `batchDeleteRequests`. Returns the raw API response. |
| `ListAndLockHead(ctx, lockSecs int64, limit *int64) (json.RawMessage, error)` | Fetch the head of the queue and lock the returned requests for `lockSecs` seconds. Returns the raw API response. |
| `ProlongRequestLock(ctx, id string, lockSecs int64, forefront bool) (json.RawMessage, error)` | Extend the lock on a request by `lockSecs` seconds; `forefront` controls re-queue position when the lock expires. |
| `DeleteRequestLock(ctx, id string, forefront bool) error` | Release the lock on a request without modifying it. |
| `UnlockRequests(ctx) (json.RawMessage, error)` | Release all locks held by this client (see `WithClientKey`). |
| `WithClientKey(key string) *RequestQueueClient` | Pin a stable client key (required to unlock own locks). |

The `RequestQueue` value returned by `Get`/`GetOrCreate`/`Update` and listed by `List`/`Iterate`:

| Field | Type | Meaning |
|---|---|---|
| `ID` | `string` | Unique queue ID. |
| `Name` | `string` | Queue name (empty for unnamed queues). |
| `UserID` | `string` | ID of the user who owns the queue. |
| `CreatedAt` | `*time.Time` | When the queue was created. |
| `ModifiedAt` | `*time.Time` | When the queue was last modified. |
| `TotalRequestCount` | `int64` | Total number of requests ever added. |
| `Extra` | `map[string]json.RawMessage` | Any other fields returned by the API. |

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
| `ExclusiveStartID` | `*string` | List requests after this ID (deprecated; prefer `Cursor`). |
| `Cursor` | `*string` | Opaque pagination cursor: the `string` returned as `nextCursor` in the previous `ListRequests` response. Pass it back to fetch the next page. Mutually exclusive with the deprecated `ExclusiveStartID` (setting both returns an error). |
| `Filter` | `[]string` | Restrict the listing to requests in the given states; sent as a comma-separated list whose semantics are the union of the states (requests matching any are returned). Each value must be `"locked"` or `"pending"` — use the exported `apify.RequestFilterLocked` / `apify.RequestFilterPending` constants. |

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

// List only pending or locked requests (the Filter values are sent comma-joined, union semantics).
_, _ = queue.ListRequests(ctx, apify.ListRequestsOptions{
	Filter: []string{apify.RequestFilterPending, apify.RequestFilterLocked},
})

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
