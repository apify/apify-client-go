package apify

import (
	"encoding/json"
	"errors"
	"fmt"
)

// APIError is returned for HTTP requests that reach the Apify API but receive a
// non-success status code.
//
// It mirrors the `ApifyApiError` of the reference JavaScript client and exposes the
// parsed error Type, the human-readable Message, the HTTP StatusCode, the number of
// the final Attempt, and the request HTTPMethod/Path.
type APIError struct {
	// StatusCode is the HTTP status code of the error response.
	StatusCode int
	// Type is the machine-readable error type returned by the API (e.g. "record-not-found").
	Type string
	// Message is the human-readable description of the error returned by the API.
	Message string
	// Attempt is the number of the API call attempt that produced this error (1-based).
	Attempt int
	// HTTPMethod is the HTTP method of the API call (e.g. "GET", "POST").
	HTTPMethod string
	// Path is the full path of the API endpoint (URL excluding origin).
	Path string
	// Data holds additional structured data provided by the API about the error, if any.
	Data map[string]any
}

// HTTP status codes classified by the Is* predicates below.
const (
	statusInvalidRequest = 400
	statusUnauthorized   = 401
	statusForbidden      = 403
	statusConflict       = 409
	statusRateLimited    = 429
	statusServerErrorMin = 500
)

// IsInvalidRequest reports whether this is an HTTP 400 Bad Request, typically because the
// request failed validation.
//
// The reference clients throw a distinct InvalidRequestError subclass for this status; since Go
// has no exception hierarchy to mirror, these Is* predicates are the idiomatic equivalent
// classification on the one APIError type.
func (e *APIError) IsInvalidRequest() bool { return e.StatusCode == statusInvalidRequest }

// IsUnauthorized reports whether this is an HTTP 401 Unauthorized: the token is missing or
// invalid.
func (e *APIError) IsUnauthorized() bool { return e.StatusCode == statusUnauthorized }

// IsForbidden reports whether this is an HTTP 403 Forbidden: the token lacks permission for the
// operation.
func (e *APIError) IsForbidden() bool { return e.StatusCode == statusForbidden }

// IsNotFound reports whether this is an HTTP 404 Not Found.
//
// Most Get-style methods already map a 404 to a false "present" bool rather than an error (see
// e.g. [ActorClient.Get]), so this is mainly useful for methods that return other errors
// directly, e.g. [ActorClient.Start] with a nonexistent Actor ID.
func (e *APIError) IsNotFound() bool { return e.StatusCode == notFoundStatusCode }

// IsConflict reports whether this is an HTTP 409 Conflict.
func (e *APIError) IsConflict() bool { return e.StatusCode == statusConflict }

// IsRateLimited reports whether this is an HTTP 429 Too Many Requests. The client already
// retries these internally (see [WithMaxRetries]), so this surfaces only once retries are
// exhausted.
func (e *APIError) IsRateLimited() bool { return e.StatusCode == statusRateLimited }

// IsServerError reports whether this is an HTTP 5xx status. Like IsRateLimited, the client
// already retries these internally, so this surfaces only once retries are exhausted.
func (e *APIError) IsServerError() bool { return e.StatusCode >= statusServerErrorMin }

// Error implements the error interface.
func (e *APIError) Error() string {
	errType := e.Type
	if errType == "" {
		errType = "unknown"
	}
	return fmt.Sprintf("apify API error (status %d, type %s): %s", e.StatusCode, errType, e.Message)
}

// apiErrorBody is the shape of the `error` object returned by the Apify API on failure:
// `{ "error": { "type": "...", "message": "..." } }`.
type apiErrorBody struct {
	Error struct {
		Type    string         `json:"type"`
		Message string         `json:"message"`
		Data    map[string]any `json:"data"`
	} `json:"error"`
}

// buildAPIError parses an API error response body into an *APIError.
func buildAPIError(status int, body []byte, attempt int, method, path string) *APIError {
	apiErr := &APIError{
		StatusCode: status,
		Attempt:    attempt,
		HTTPMethod: method,
		Path:       path,
	}

	var parsed apiErrorBody
	if err := json.Unmarshal(body, &parsed); err == nil && parsed.Error.Message != "" {
		apiErr.Type = parsed.Error.Type
		apiErr.Message = parsed.Error.Message
		apiErr.Data = parsed.Error.Data
		return apiErr
	}

	if len(body) == 0 {
		apiErr.Message = fmt.Sprintf("unexpected error with status %d", status)
	} else {
		apiErr.Message = fmt.Sprintf("unexpected error: %s", string(body))
	}
	return apiErr
}

// AsAPIError returns the underlying *APIError if err is (or wraps) one, and true;
// otherwise it returns nil and false. It is a convenience wrapper around errors.As.
func AsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}
