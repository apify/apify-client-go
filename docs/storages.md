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

For typed item decoding use the generic helper `apify.ListDatasetItems[T](ctx, dataset, opts)`.

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

```go
rq, _ := client.RequestQueues().GetOrCreate(ctx, "")
defer client.RequestQueue(rq.ID).Delete(ctx)

queue := client.RequestQueue(rq.ID)
_, _ = queue.AddRequest(ctx, apify.RequestQueueRequest{URL: "https://example.com", UniqueKey: "example"}, false)

it := queue.PaginateRequests(apify.Ptr(int64(100)))
for {
	req, err := it.Next(ctx)
	if err != nil { log.Fatal(err) }
	if req == nil { break }
	fmt.Println(req.URL)
}
```

`StorageListOptions` (shared by the three collections): `Offset`, `Limit`, `Desc`,
`Unnamed`, `Ownership`.
