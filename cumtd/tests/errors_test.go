package cumtd_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestAPIError_errorsAs(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/routes/groups/MISSING": {
			StatusCode: 404,
			Body:       []byte(`not found`),
		},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetRouteGroup(context.Background(), "MISSING")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var apiErr *cumtd.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("errors.As(*APIError) = false, err = %v", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("StatusCode = %d, want 404", apiErr.StatusCode)
	}
}

func TestRateLimitError_errorsAs(t *testing.T) {
	// Use a custom server so we can set the Retry-After header.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "60")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`rate limited`))
	}))
	t.Cleanup(srv.Close)

	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetRouteGroups(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var rlErr *cumtd.RateLimitError
	if !errors.As(err, &rlErr) {
		t.Fatalf("errors.As(*RateLimitError) = false, err = %v", err)
	}
	if rlErr.RetryAfter != "60" {
		t.Errorf("RetryAfter = %q, want %q", rlErr.RetryAfter, "60")
	}
}

func TestValidationError_errorsAs(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetRouteGroup(context.Background(), "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("errors.As(*ValidationError) = false, err = %v", err)
	}
	if valErr.Field == "" {
		t.Error("ValidationError.Field should not be empty")
	}
}
