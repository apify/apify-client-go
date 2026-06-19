package apify

import (
	"encoding/json"
	"time"
)

// Extra is the catch-all map of unmodelled JSON fields. Every resource model carries one
// so that additive changes to the API never break deserialization (forward compatibility).
type Extra = map[string]json.RawMessage

// unmarshalWithExtra unmarshals data into the typed value v (which must be a pointer to a
// struct) and additionally collects any JSON keys not mapped to a struct field into extra.
//
// This gives every model forward compatibility with additive API fields without hand-writing
// an UnmarshalJSON for each one: the model's UnmarshalJSON delegates here.
func unmarshalWithExtra(data []byte, v any, known map[string]struct{}, extra *Extra) error {
	// First, unmarshal the typed fields. v must be an alias that does NOT re-enter this
	// model's UnmarshalJSON (callers pass a type alias to avoid infinite recursion).
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	out := Extra{}
	for k, val := range raw {
		if _, ok := known[k]; !ok {
			out[k] = val
		}
	}
	if len(out) > 0 {
		*extra = out
	}
	return nil
}

// Actor is an Actor on the Apify platform.
type Actor struct {
	// ID is the unique Actor ID.
	ID string `json:"id"`
	// UserID is the ID of the user who owns the Actor.
	UserID string `json:"userId"`
	// Name is the technical name of the Actor (used in API paths).
	Name string `json:"name"`
	// Username is the username of the Actor's owner.
	Username string `json:"username"`
	// Title is the human-readable title shown in the UI.
	Title string `json:"title"`
	// Description describes what the Actor does.
	Description string `json:"description"`
	// IsPublic reports whether the Actor is publicly available in Apify Store.
	IsPublic bool `json:"isPublic"`
	// CreatedAt is when the Actor was created.
	CreatedAt *time.Time `json:"createdAt"`
	// ModifiedAt is when the Actor was last modified.
	ModifiedAt *time.Time `json:"modifiedAt"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (a *Actor) UnmarshalJSON(data []byte) error {
	type alias Actor
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(a), known, &a.Extra)
}

// ActorRun is a single execution of an Actor.
type ActorRun struct {
	// ID is the unique run ID.
	ID string `json:"id"`
	// ActID is the ID of the Actor that produced this run.
	ActID string `json:"actId"`
	// ActorTaskID is the ID of the task that started this run, if any.
	ActorTaskID string `json:"actorTaskId"`
	// UserID is the ID of the user who owns the run.
	UserID string `json:"userId"`
	// Status is the current run status. One of the eight ActorJobStatus values: READY, RUNNING,
	// SUCCEEDED, FAILED, TIMING-OUT, TIMED-OUT, ABORTING, ABORTED.
	Status string `json:"status"`
	// StatusMessage is an optional human-readable status message.
	StatusMessage string `json:"statusMessage"`
	// StartedAt is when the run started.
	StartedAt *time.Time `json:"startedAt"`
	// FinishedAt is when the run finished (absent while still running).
	FinishedAt *time.Time `json:"finishedAt"`
	// BuildID is the ID of the build used for the run.
	BuildID string `json:"buildId"`
	// DefaultDatasetID is the ID of the run's default dataset.
	DefaultDatasetID string `json:"defaultDatasetId"`
	// DefaultKeyValueStoreID is the ID of the run's default key-value store.
	DefaultKeyValueStoreID string `json:"defaultKeyValueStoreId"`
	// DefaultRequestQueueID is the ID of the run's default request queue.
	DefaultRequestQueueID string `json:"defaultRequestQueueId"`
	// ContainerURL is the URL of the run's container (for live access).
	ContainerURL string `json:"containerUrl"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (r *ActorRun) UnmarshalJSON(data []byte) error {
	type alias ActorRun
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(r), known, &r.Extra)
}

// IsTerminal reports whether the run has reached a terminal (finished) status.
func (r *ActorRun) IsTerminal() bool {
	return isTerminalStatus(r.Status)
}

// Build is a single build of an Actor.
type Build struct {
	// ID is the unique build ID.
	ID string `json:"id"`
	// ActID is the ID of the Actor this build belongs to.
	ActID string `json:"actId"`
	// Status is the current build status.
	Status string `json:"status"`
	// StartedAt is when the build started.
	StartedAt *time.Time `json:"startedAt"`
	// FinishedAt is when the build finished (absent while still building).
	FinishedAt *time.Time `json:"finishedAt"`
	// BuildNumber is the human-readable build number (e.g. "0.1.2").
	BuildNumber string `json:"buildNumber"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (b *Build) UnmarshalJSON(data []byte) error {
	type alias Build
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(b), known, &b.Extra)
}

// IsTerminal reports whether the build has reached a terminal (finished) status.
func (b *Build) IsTerminal() bool {
	return isTerminalStatus(b.Status)
}

// Task is a pre-configured Actor run (an Actor task).
type Task struct {
	// ID is the unique task ID.
	ID string `json:"id"`
	// ActID is the ID of the Actor this task runs.
	ActID string `json:"actId"`
	// UserID is the ID of the user who owns the task.
	UserID string `json:"userId"`
	// Name is the technical name of the task.
	Name string `json:"name"`
	// Title is the human-readable title shown in the UI.
	Title string `json:"title"`
	// CreatedAt is when the task was created.
	CreatedAt *time.Time `json:"createdAt"`
	// ModifiedAt is when the task was last modified.
	ModifiedAt *time.Time `json:"modifiedAt"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (t *Task) UnmarshalJSON(data []byte) error {
	type alias Task
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(t), known, &t.Extra)
}

// Dataset stores structured results from Actor runs.
type Dataset struct {
	// ID is the unique dataset ID.
	ID string `json:"id"`
	// Name is the dataset name (empty for unnamed datasets).
	Name string `json:"name"`
	// UserID is the ID of the user who owns the dataset.
	UserID string `json:"userId"`
	// CreatedAt is when the dataset was created.
	CreatedAt *time.Time `json:"createdAt"`
	// ModifiedAt is when the dataset was last modified.
	ModifiedAt *time.Time `json:"modifiedAt"`
	// ItemCount is the number of items currently stored.
	ItemCount int64 `json:"itemCount"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (d *Dataset) UnmarshalJSON(data []byte) error {
	type alias Dataset
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(d), known, &d.Extra)
}

// KeyValueStore stores arbitrary data records.
type KeyValueStore struct {
	// ID is the unique store ID.
	ID string `json:"id"`
	// Name is the store name (empty for unnamed stores).
	Name string `json:"name"`
	// UserID is the ID of the user who owns the store.
	UserID string `json:"userId"`
	// CreatedAt is when the store was created.
	CreatedAt *time.Time `json:"createdAt"`
	// ModifiedAt is when the store was last modified.
	ModifiedAt *time.Time `json:"modifiedAt"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (k *KeyValueStore) UnmarshalJSON(data []byte) error {
	type alias KeyValueStore
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(k), known, &k.Extra)
}

// KeyValueStoreKey is a single key listed from a key-value store.
type KeyValueStoreKey struct {
	// Key is the record key.
	Key string `json:"key"`
	// Size is the record size in bytes.
	Size int64 `json:"size"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (k *KeyValueStoreKey) UnmarshalJSON(data []byte) error {
	type alias KeyValueStoreKey
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(k), known, &k.Extra)
}

// KeyValueStoreKeysPage is a page of keys from a key-value store.
type KeyValueStoreKeysPage struct {
	// Limit is the maximum number of keys requested.
	Limit int64 `json:"limit"`
	// IsTruncated reports whether more keys are available.
	IsTruncated bool `json:"isTruncated"`
	// ExclusiveStartKey is the key the listing started after.
	ExclusiveStartKey string `json:"exclusiveStartKey"`
	// NextExclusiveStartKey is the key to pass to fetch the next page.
	NextExclusiveStartKey string `json:"nextExclusiveStartKey"`
	// Items are the listed keys.
	Items []KeyValueStoreKey `json:"items"`
}

// KeyValueStoreRecord is a single record retrieved from a key-value store. Its Value holds
// the raw bytes; callers can decode it according to ContentType.
type KeyValueStoreRecord struct {
	// Key is the record key.
	Key string
	// Value is the raw record bytes.
	Value []byte
	// ContentType is the record's MIME type, as reported by the API.
	ContentType string
}

// RequestQueue stores URLs to be crawled.
type RequestQueue struct {
	// ID is the unique queue ID.
	ID string `json:"id"`
	// Name is the queue name (empty for unnamed queues).
	Name string `json:"name"`
	// UserID is the ID of the user who owns the queue.
	UserID string `json:"userId"`
	// CreatedAt is when the queue was created.
	CreatedAt *time.Time `json:"createdAt"`
	// ModifiedAt is when the queue was last modified.
	ModifiedAt *time.Time `json:"modifiedAt"`
	// TotalRequestCount is the total number of requests ever added.
	TotalRequestCount int64 `json:"totalRequestCount"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (q *RequestQueue) UnmarshalJSON(data []byte) error {
	type alias RequestQueue
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(q), known, &q.Extra)
}

// RequestQueueRequest is a single request stored in a request queue.
type RequestQueueRequest struct {
	// ID is the unique request ID (assigned by the API; omitted on create).
	ID string `json:"id,omitempty"`
	// URL is the request URL.
	URL string `json:"url"`
	// UniqueKey is the deduplication key for the request.
	UniqueKey string `json:"uniqueKey,omitempty"`
	// Method is the HTTP method (e.g. "GET", "POST").
	Method string `json:"method,omitempty"`
	// UserData is arbitrary user-attached metadata.
	UserData json.RawMessage `json:"userData,omitempty"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (r *RequestQueueRequest) UnmarshalJSON(data []byte) error {
	type alias RequestQueueRequest
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(r), known, &r.Extra)
}

// RequestQueueOperationInfo is returned when adding or updating a request.
type RequestQueueOperationInfo struct {
	// RequestID is the ID of the affected request.
	RequestID string `json:"requestId"`
	// WasAlreadyPresent reports whether the request was already in the queue.
	WasAlreadyPresent bool `json:"wasAlreadyPresent"`
	// WasAlreadyHandled reports whether the request had already been handled.
	WasAlreadyHandled bool `json:"wasAlreadyHandled"`
}

// RequestQueueHead is the head (front) of a request queue.
type RequestQueueHead struct {
	// Limit is the maximum number of requests requested.
	Limit int64 `json:"limit"`
	// HadMultipleClients reports whether multiple clients have accessed the queue.
	HadMultipleClients bool `json:"hadMultipleClients"`
	// Items are the requests at the head of the queue.
	Items []RequestQueueRequest `json:"items"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (h *RequestQueueHead) UnmarshalJSON(data []byte) error {
	type alias RequestQueueHead
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(h), known, &h.Extra)
}

// Schedule automatically starts Actor or task runs at specified times.
type Schedule struct {
	// ID is the unique schedule ID.
	ID string `json:"id"`
	// UserID is the ID of the user who owns the schedule.
	UserID string `json:"userId"`
	// Name is the schedule name.
	Name string `json:"name"`
	// CronExpression is the cron expression governing when the schedule fires.
	CronExpression string `json:"cronExpression"`
	// IsEnabled reports whether the schedule is currently active.
	IsEnabled bool `json:"isEnabled"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (s *Schedule) UnmarshalJSON(data []byte) error {
	type alias Schedule
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(s), known, &s.Extra)
}

// Webhook notifies an external service when specific events occur.
type Webhook struct {
	// ID is the unique webhook ID.
	ID string `json:"id"`
	// UserID is the ID of the user who owns the webhook.
	UserID string `json:"userId"`
	// RequestURL is the URL the webhook posts to.
	RequestURL string `json:"requestUrl"`
	// EventTypes are the events that trigger the webhook.
	EventTypes []string `json:"eventTypes"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	type alias Webhook
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(w), known, &w.Extra)
}

// WebhookDispatch is a single invocation of a webhook.
type WebhookDispatch struct {
	// ID is the unique dispatch ID.
	ID string `json:"id"`
	// WebhookID is the ID of the webhook that produced this dispatch.
	WebhookID string `json:"webhookId"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (d *WebhookDispatch) UnmarshalJSON(data []byte) error {
	type alias WebhookDispatch
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(d), known, &d.Extra)
}

// User is an Apify user account.
type User struct {
	// ID is the unique user ID.
	ID string `json:"id"`
	// Username is the user's username.
	Username string `json:"username"`
	// Extra holds any other fields returned by the API (private details for "me").
	Extra Extra `json:"-"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type alias User
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(u), known, &u.Extra)
}

// ActorStoreListItem is an Actor as listed in the Apify Store.
type ActorStoreListItem struct {
	// ID is the unique Actor ID.
	ID string `json:"id"`
	// Name is the technical name of the Actor.
	Name string `json:"name"`
	// Username is the username of the Actor's owner.
	Username string `json:"username"`
	// Title is the human-readable title.
	Title string `json:"title"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (a *ActorStoreListItem) UnmarshalJSON(data []byte) error {
	type alias ActorStoreListItem
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(a), known, &a.Extra)
}

// ActorVersion is a single version of an Actor.
type ActorVersion struct {
	// VersionNumber is the version identifier (e.g. "0.1").
	VersionNumber string `json:"versionNumber"`
	// SourceType is how the version's source is provided (e.g. "SOURCE_FILES").
	SourceType string `json:"sourceType"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (v *ActorVersion) UnmarshalJSON(data []byte) error {
	type alias ActorVersion
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(v), known, &v.Extra)
}

// ActorEnvVar is an environment variable attached to an Actor version.
type ActorEnvVar struct {
	// Name is the environment variable name.
	Name string `json:"name"`
	// Value is the environment variable value.
	Value string `json:"value,omitempty"`
	// IsSecret reports whether the value is stored as a secret.
	IsSecret *bool `json:"isSecret,omitempty"`
	// Extra holds any other fields returned by the API.
	Extra Extra `json:"-"`
}

func (e *ActorEnvVar) UnmarshalJSON(data []byte) error {
	type alias ActorEnvVar
	known := knownJSONKeys(alias{})
	return unmarshalWithExtra(data, (*alias)(e), known, &e.Extra)
}
