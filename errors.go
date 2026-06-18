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
