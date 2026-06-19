package apify

import (
	"context"
	"encoding/json"
)

// BuildCollectionClient is a client for a build collection: either the account-wide
// collection (GET /v2/actor-builds) or an Actor's builds (GET /v2/actors/{id}/builds).
type BuildCollectionClient struct {
	ctx *resourceContext
}

func newBuildCollectionClient(hc *httpClient, baseURL, resourcePath string) *BuildCollectionClient {
	return &BuildCollectionClient{ctx: newCollectionContext(hc, baseURL, resourcePath)}
}

// newBuildCollectionClientWithBase builds a collection client nested under another resource
// (e.g. an Actor), using base as the parent URL.
func newBuildCollectionClientWithBase(hc *httpClient, base, resourcePath string) *BuildCollectionClient {
	return &BuildCollectionClient{ctx: newCollectionContext(hc, base, resourcePath)}
}

// List lists builds.
func (c *BuildCollectionClient) List(ctx context.Context, options ListOptions) (PaginationList[Build], error) {
	params := NewQueryParams()
	options.apply(params)
	return listResource[Build](ctx, c.ctx, "", params)
}

// BuildClient is a client for a specific Actor build (/v2/actor-builds/{buildId}).
type BuildClient struct {
	ctx     *resourceContext
	baseURL string
	id      string
}

func newBuildClient(hc *httpClient, baseURL, id string) *BuildClient {
	return &BuildClient{
		ctx:     newSingleContext(hc, baseURL, "actor-builds", id),
		baseURL: baseURL,
		id:      id,
	}
}

// Get fetches the build object. The bool reports whether it exists.
func (c *BuildClient) Get(ctx context.Context) (Build, bool, error) {
	return c.GetWithWait(ctx, nil)
}

// GetWithWait fetches the build, optionally asking the API to wait up to waitForFinishSecs
// seconds (max 60) for the build to finish before responding. Pass nil for an immediate
// fetch. Mirrors the reference client's get({ waitForFinish }).
func (c *BuildClient) GetWithWait(ctx context.Context, waitForFinishSecs *int64) (Build, bool, error) {
	params := NewQueryParams()
	params.AddInt("waitForFinish", waitForFinishSecs)
	return getResource[Build](ctx, c.ctx, "", params)
}

// Abort aborts the build and returns its updated state.
func (c *BuildClient) Abort(ctx context.Context) (Build, error) {
	return postWithBody[Build](ctx, c.ctx, "abort", NewQueryParams(), nil, "")
}

// Delete deletes the build.
func (c *BuildClient) Delete(ctx context.Context) error {
	return deleteResource(ctx, c.ctx, "")
}

// WaitForFinish polls until the build reaches a terminal state or waitSecs elapses (nil
// waits indefinitely). It returns the latest build.
func (c *BuildClient) WaitForFinish(ctx context.Context, waitSecs *int64) (Build, error) {
	return waitForFinish[Build](ctx, c.ctx, waitSecs, "build", func(b *Build) bool { return b.IsTerminal() })
}

// GetOpenAPIDefinition returns the OpenAPI definition generated for the build, or
// (nil, false, nil) if it is not available. The result is the raw OpenAPI document.
func (c *BuildClient) GetOpenAPIDefinition(ctx context.Context) (json.RawMessage, bool, error) {
	resp, err := getRaw(ctx, c.ctx, "openapi.json", NewQueryParams())
	if err != nil {
		return nil, false, err
	}
	if resp == nil {
		return nil, false, nil
	}
	return resp.body, true, nil
}

// Log returns a client for accessing this build's log.
func (c *BuildClient) Log() *LogClient {
	return newNestedLogClient(c.ctx.http, c.ctx.subURL(""))
}
