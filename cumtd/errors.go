package cumtd

import "fmt"

// APIError is returned when the server responds with a non-200 status.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("cumtd: API error %d: %s", e.StatusCode, e.Body)
}

// RateLimitError is returned on HTTP 429.
type RateLimitError struct {
	RetryAfter string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("cumtd: rate limited, retry after %s", e.RetryAfter)
}

// ValidationError is returned when a required parameter is missing.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("cumtd: validation error: %s %s", e.Field, e.Message)
}

// require returns a ValidationError if s is empty.
func require(field, s string) error {
	if s == "" {
		return &ValidationError{Field: field, Message: "must not be empty"}
	}
	return nil
}
