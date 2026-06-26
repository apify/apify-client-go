// Package apify is the official, idiomatic Go client for the Apify API
// (https://docs.apify.com/api/v2).
//
// It provides a resource-oriented interface that mirrors the official JavaScript and Rust
// clients: start from an [ApifyClient], then drill down into resources (Actors, runs,
// datasets, key-value stores, request queues, tasks, schedules, webhooks, the store, users
// and logs).
//
// # Quick start
//
//	client := apify.NewClient("my-api-token")
//
//	// Start an Actor and wait for it to finish.
//	run, err := client.Actor("apify/hello-world").Call(ctx, nil, apify.ActorStartOptions{}, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Read items from the run's default dataset.
//	page, err := client.Dataset(run.DefaultDatasetID).ListItems(ctx, apify.DatasetListItemsOptions{})
//
// # Architecture
//
//   - Public interface: [ApifyClient] and the resource clients it returns.
//   - Replaceable transport: the [HTTPBackend] interface, with a default
//     [DefaultHTTPBackend]. Swap it via [WithHTTPBackend].
//   - Cross-cutting behaviour (auth, User-Agent, retries with exponential backoff,
//     timeouts) lives in the internal HTTP client and is applied to every request.
package apify

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"
)

// Default configuration values, matching the reference clients.
const (
	// defaultBaseURL is the default base URL of the Apify API (without the /v2 suffix).
	defaultBaseURL = "https://api.apify.com"
	// defaultMaxRetries is the default maximum number of retries.
	defaultMaxRetries = 8
	// defaultMinDelayBetweenRetries is the default minimum delay between retries.
	defaultMinDelayBetweenRetries = 500 * time.Millisecond
	// defaultTimeout is the default overall per-request timeout (6 minutes).
	defaultTimeout = 360 * time.Second
	// meUserPlaceholder addresses the current user (/users/me).
	meUserPlaceholder = "me"
)

// ApifyClient is the entry point for interacting with the Apify API.
//
// Construct it with [NewClient] (token-only) or [NewClientWithOptions], then obtain
// resource clients via the accessor methods, e.g. [ApifyClient.Actor], [ApifyClient.Dataset],
// [ApifyClient.Run]. It is safe for concurrent use.
type ApifyClient struct {
	http          *httpClient
	baseURL       string
	publicBaseURL string
}

// clientConfig accumulates the options passed to NewClientWithOptions.
type clientConfig struct {
	token                  string
	baseURL                string
	publicBaseURL          string
	maxRetries             int
	minDelayBetweenRetries time.Duration
	timeout                time.Duration
	userAgentSuffix        string
	backend                HTTPBackend
	isAtHomeFn             func() bool
}

// Option configures an [ApifyClient]. Pass options to [NewClientWithOptions].
type Option func(*clientConfig)

// WithToken sets the API token used for authentication (sent as a Bearer token).
func WithToken(token string) Option {
	return func(c *clientConfig) { c.token = token }
}

// WithBaseURL overrides the base URL of the API. The /v2 suffix is appended automatically.
// Defaults to https://api.apify.com.
func WithBaseURL(baseURL string) Option {
	return func(c *clientConfig) { c.baseURL = baseURL }
}

// WithPublicBaseURL overrides the base URL used when building public, shareable resource
// URLs (e.g. a signed dataset-items URL). Defaults to the API base URL. The /v2 suffix is
// appended automatically.
func WithPublicBaseURL(publicBaseURL string) Option {
	return func(c *clientConfig) { c.publicBaseURL = publicBaseURL }
}

// WithMaxRetries sets the maximum number of retries for failed requests (default 8).
func WithMaxRetries(maxRetries int) Option {
	return func(c *clientConfig) { c.maxRetries = maxRetries }
}

// WithMinDelayBetweenRetries sets the minimum delay between retries (default 500ms).
func WithMinDelayBetweenRetries(d time.Duration) Option {
	return func(c *clientConfig) { c.minDelayBetweenRetries = d }
}

// WithTimeout sets the overall per-request timeout (default 360s).
func WithTimeout(d time.Duration) Option {
	return func(c *clientConfig) { c.timeout = d }
}

// WithUserAgentSuffix appends a custom suffix to the User-Agent header.
func WithUserAgentSuffix(suffix string) Option {
	return func(c *clientConfig) { c.userAgentSuffix = suffix }
}

// WithHTTPBackend replaces the default HTTP backend with a custom implementation. This is
// the seam that makes the transport a replaceable component.
func WithHTTPBackend(backend HTTPBackend) Option {
	return func(c *clientConfig) { c.backend = backend }
}

// withIsAtHomeFn overrides how the client decides the isAtHome User-Agent flag (test seam).
func withIsAtHomeFn(fn func() bool) Option {
	return func(c *clientConfig) { c.isAtHomeFn = fn }
}

// Environment variable that signals the client is running on the Apify platform.
//
// Per client_requirements.md the flag is based solely on APIFY_IS_AT_HOME ("based on the
// environment variable `APIFY_IS_AT_HOME`, `False` if env variable is missing"), which also
// matches the JS reference (it reads only APIFY_IS_AT_HOME via @apify/consts).
const envIsAtHome = "APIFY_IS_AT_HOME"

// defaultIsAtHome reports whether the client is running on the Apify platform, by reading the
// APIFY_IS_AT_HOME environment variable (set to any non-empty value on the platform).
func defaultIsAtHome() bool {
	return os.Getenv(envIsAtHome) != ""
}

// NewClient creates a client authenticated with the given API token and default settings.
func NewClient(token string) *ApifyClient {
	return NewClientWithOptions(WithToken(token))
}

// NewClientWithOptions creates a client configured by the given options.
func NewClientWithOptions(opts ...Option) *ApifyClient {
	cfg := &clientConfig{
		baseURL:                defaultBaseURL,
		maxRetries:             defaultMaxRetries,
		minDelayBetweenRetries: defaultMinDelayBetweenRetries,
		timeout:                defaultTimeout,
		isAtHomeFn:             defaultIsAtHome,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	backend := cfg.backend
	if backend == nil {
		backend = NewDefaultHTTPBackend()
	}

	userAgent := BuildUserAgent(cfg.userAgentSuffix, cfg.isAtHomeFn)

	hc := &httpClient{
		backend:   backend,
		token:     cfg.token,
		userAgent: userAgent,
		retry: retryConfig{
			maxRetries:             cfg.maxRetries,
			minDelayBetweenRetries: cfg.minDelayBetweenRetries,
			timeout:                cfg.timeout,
		},
	}

	baseURL := strings.TrimRight(cfg.baseURL, "/") + "/v2"
	publicSource := cfg.publicBaseURL
	if publicSource == "" {
		publicSource = cfg.baseURL
	}
	publicBaseURL := strings.TrimRight(publicSource, "/") + "/v2"

	return &ApifyClient{http: hc, baseURL: baseURL, publicBaseURL: publicBaseURL}
}

// UserAgent returns the User-Agent header value this client sends.
func (c *ApifyClient) UserAgent() string { return c.http.userAgent }

// APIBaseURL returns the fully-qualified API base URL this client targets (including the
// /v2 suffix), e.g. https://api.apify.com/v2. Reflects any base-URL override.
func (c *ApifyClient) APIBaseURL() string { return c.baseURL }

// ----- Actor accessors -----------------------------------------------------

// Actors returns a client for the Actor collection (list & create Actors).
func (c *ApifyClient) Actors() *ActorCollectionClient {
	return newActorCollectionClient(c.http, c.baseURL)
}

// Actor returns a client for a specific Actor, addressed by ID or username~name.
func (c *ApifyClient) Actor(id string) *ActorClient {
	return newActorClient(c, c.http, c.baseURL, id)
}

// ----- Build accessors -----------------------------------------------------

// Builds returns a client for the Actor build collection (list builds).
func (c *ApifyClient) Builds() *BuildCollectionClient {
	return newBuildCollectionClient(c.http, c.baseURL, "actor-builds")
}

// Build returns a client for a specific Actor build.
func (c *ApifyClient) Build(id string) *BuildClient {
	return newBuildClient(c.http, c.baseURL, id)
}

// ----- Run accessors -------------------------------------------------------

// Runs returns a client for the Actor run collection (list runs).
func (c *ApifyClient) Runs() *RunCollectionClient {
	return newRunCollectionClient(c.http, c.baseURL, "actor-runs")
}

// Run returns a client for a specific Actor run.
func (c *ApifyClient) Run(id string) *RunClient {
	return newRunClient(c, c.http, c.baseURL, "actor-runs", id)
}

// ----- Dataset accessors ---------------------------------------------------

// Datasets returns a client for the dataset collection (list & get-or-create datasets).
func (c *ApifyClient) Datasets() *DatasetCollectionClient {
	return newDatasetCollectionClient(c.http, c.baseURL)
}

// Dataset returns a client for a specific dataset, addressed by ID or name.
func (c *ApifyClient) Dataset(id string) *DatasetClient {
	return newDatasetClient(c.http, c.baseURL, "datasets", id).withPublicBase(c.publicBaseURL)
}

// ----- Key-value store accessors -------------------------------------------

// KeyValueStores returns a client for the key-value store collection.
func (c *ApifyClient) KeyValueStores() *KeyValueStoreCollectionClient {
	return newKeyValueStoreCollectionClient(c.http, c.baseURL)
}

// KeyValueStore returns a client for a specific key-value store, addressed by ID or name.
func (c *ApifyClient) KeyValueStore(id string) *KeyValueStoreClient {
	return newKeyValueStoreClient(c.http, c.baseURL, "key-value-stores", id).withPublicBase(c.publicBaseURL)
}

// ----- Request queue accessors ---------------------------------------------

// RequestQueues returns a client for the request queue collection.
func (c *ApifyClient) RequestQueues() *RequestQueueCollectionClient {
	return newRequestQueueCollectionClient(c.http, c.baseURL)
}

// RequestQueue returns a client for a specific request queue, addressed by ID or name.
func (c *ApifyClient) RequestQueue(id string) *RequestQueueClient {
	return newRequestQueueClient(c.http, c.baseURL, "request-queues", id)
}

// ----- Task accessors ------------------------------------------------------

// Tasks returns a client for the Actor task collection (list & create tasks).
func (c *ApifyClient) Tasks() *TaskCollectionClient {
	return newTaskCollectionClient(c.http, c.baseURL)
}

// Task returns a client for a specific Actor task.
func (c *ApifyClient) Task(id string) *TaskClient {
	return newTaskClient(c, c.http, c.baseURL, id)
}

// ----- Schedule accessors --------------------------------------------------

// Schedules returns a client for the schedule collection (list & create schedules).
func (c *ApifyClient) Schedules() *ScheduleCollectionClient {
	return newScheduleCollectionClient(c.http, c.baseURL)
}

// Schedule returns a client for a specific schedule.
func (c *ApifyClient) Schedule(id string) *ScheduleClient {
	return newScheduleClient(c.http, c.baseURL, id)
}

// ----- Webhook accessors ---------------------------------------------------

// Webhooks returns a client for the webhook collection (list & create webhooks).
func (c *ApifyClient) Webhooks() *WebhookCollectionClient {
	return newWebhookCollectionClient(c.http, c.baseURL)
}

// Webhook returns a client for a specific webhook.
func (c *ApifyClient) Webhook(id string) *WebhookClient {
	return newWebhookClient(c.http, c.baseURL, id)
}

// WebhookDispatches returns a client for the webhook dispatch collection.
func (c *ApifyClient) WebhookDispatches() *WebhookDispatchCollectionClient {
	return newWebhookDispatchCollectionClient(c.http, c.baseURL)
}

// WebhookDispatch returns a client for a specific webhook dispatch.
func (c *ApifyClient) WebhookDispatch(id string) *WebhookDispatchClient {
	return newWebhookDispatchClient(c.http, c.baseURL, id)
}

// ----- Misc accessors ------------------------------------------------------

// Store returns a client for browsing the Apify Store.
func (c *ApifyClient) Store() *StoreCollectionClient {
	return newStoreCollectionClient(c.http, c.baseURL)
}

// Log returns a client for accessing a build's or run's log.
func (c *ApifyClient) Log(buildOrRunID string) *LogClient {
	return newLogClient(c.http, c.baseURL, "logs", buildOrRunID)
}

// Me returns a client for the current user (/users/me).
func (c *ApifyClient) Me() *UserClient {
	return newUserClient(c.http, c.baseURL, meUserPlaceholder)
}

// User returns a client for a specific user by ID or username.
func (c *ApifyClient) User(id string) *UserClient {
	return newUserClient(c.http, c.baseURL, id)
}

// SetStatusMessage sets the status message of the current Actor run.
//
// This convenience method updates the run identified by the ACTOR_RUN_ID environment
// variable, so it only works when called from inside an Actor run. If isTerminal is true,
// the message becomes final and won't be overwritten. It returns an error if ACTOR_RUN_ID
// is not set.
func (c *ApifyClient) SetStatusMessage(ctx context.Context, message string, isTerminal bool) (ActorRun, error) {
	runID := os.Getenv("ACTOR_RUN_ID")
	if runID == "" {
		return ActorRun{}, errors.New("ACTOR_RUN_ID environment variable is not set")
	}
	body := map[string]any{
		"statusMessage":           message,
		"isStatusMessageTerminal": isTerminal,
	}
	return c.Run(runID).Update(ctx, body)
}
