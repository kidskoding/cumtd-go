package cumtd

import "fmt"

// APIError is returned when the server responds with a non-2xx status code
// other than 429. Check [APIError.StatusCode] for the HTTP status and
// [APIError.Body] for the raw response body.
//
//	var apiErr *cumtd.APIError
//	if errors.As(err, &apiErr) {
//	    fmt.Println(apiErr.StatusCode, apiErr.Body)
//	}
type APIError struct {
	// StatusCode is the HTTP response status code (e.g. 404, 500).
	StatusCode int
	// Body is the raw response body returned by the server.
	Body string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("cumtd: API error %d: %s", e.StatusCode, e.Body)
}

// RateLimitError is returned when the server responds with HTTP 429 Too Many
// Requests. [RateLimitError.RetryAfter] contains the value of the
// Retry-After response header if the server set one.
//
//	var rlErr *cumtd.RateLimitError
//	if errors.As(err, &rlErr) {
//	    fmt.Println("retry after:", rlErr.RetryAfter)
//	}
type RateLimitError struct {
	// RetryAfter is the value of the Retry-After header, if present.
	// May be a number of seconds or an HTTP-date string.
	RetryAfter string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("cumtd: rate limited, retry after %s", e.RetryAfter)
}

// ValidationError is returned before any HTTP request is made when a required
// method parameter is empty. [ValidationError.Field] names the parameter.
//
//	var valErr *cumtd.ValidationError
//	if errors.As(err, &valErr) {
//	    fmt.Println("missing field:", valErr.Field)
//	}
type ValidationError struct {
	// Field is the name of the parameter that failed validation.
	Field string
	// Message describes why validation failed.
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
