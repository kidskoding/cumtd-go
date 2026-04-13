package cumtd_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestTransport_authHeader(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/routes/groups": {
			StatusCode: 200,
			Body:       []byte(`{"result":[],"error":null}`),
		},
	})

	// Use a custom transport to capture the header.
	var capturedReq *http.Request
	transport := &captureTransport{
		base: http.DefaultTransport,
		capture: func(r *http.Request) {
			capturedReq = r.Clone(r.Context())
		},
	}
	hc := &http.Client{Transport: transport}

	// Point at our mock server.
	c := cumtd.New("test-key-123",
		cumtd.WithBaseURL(srv.URL),
		cumtd.WithHTTPClient(hc),
	)
	_, _ = c.GetRouteGroups(context.Background())

	if capturedReq == nil {
		t.Fatal("no request captured")
	}
	if got := capturedReq.Header.Get("X-ApiKey"); got != "test-key-123" {
		t.Errorf("X-ApiKey = %q, want %q", got, "test-key-123")
	}
	if got := capturedReq.Header.Get("Accept"); got != "application/json" {
		t.Errorf("Accept = %q, want application/json", got)
	}
	if got := capturedReq.Header.Get("User-Agent"); got == "" {
		t.Error("User-Agent header missing")
	}
}

func TestTransport_contextCancel(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/routes/groups": {
			StatusCode: 200,
			Body:       []byte(`{"result":[],"error":null}`),
		},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately
	_, err := c.GetRouteGroups(ctx)
	if err == nil {
		t.Error("expected error for cancelled context, got nil")
	}
}

// captureTransport wraps an http.RoundTripper and calls capture before sending.
type captureTransport struct {
	base    http.RoundTripper
	capture func(*http.Request)
}

func (t *captureTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	t.capture(r)
	return t.base.RoundTrip(r)
}
