package cumtd_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestGetDepartures(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/STOP1/departures": {StatusCode: 200, Body: testutil.MustReadFixture(t, "departures.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	deps, err := c.GetDepartures(context.Background(), "STOP1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(deps) != 2 {
		t.Errorf("len(deps) = %d, want 2", len(deps))
	}
	if deps[0].UniqueID != "DEP1" {
		t.Errorf("UniqueID = %q, want DEP1", deps[0].UniqueID)
	}
}

func TestGetDepartures_routesParam(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/STOP1/departures": {
			StatusCode: 200,
			Body:       testutil.MustReadFixture(t, "departures.json"),
			AssertQuery: func(t *testing.T, q url.Values) {
				if q.Get("routes") != "ILLINI,GOLDLINE" {
					t.Errorf("routes = %q, want ILLINI,GOLDLINE", q.Get("routes"))
				}
			},
		},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetDepartures(context.Background(), "STOP1", &cumtd.GetDeparturesOptions{Routes: "ILLINI,GOLDLINE"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetDepartures_timeParam(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/STOP1/departures": {
			StatusCode: 200,
			Body:       testutil.MustReadFixture(t, "departures.json"),
			AssertQuery: func(t *testing.T, q url.Values) {
				if q.Get("time") != "08:00:00" {
					t.Errorf("time = %q, want 08:00:00", q.Get("time"))
				}
			},
		},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetDepartures(context.Background(), "STOP1", &cumtd.GetDeparturesOptions{Time: "08:00:00"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetDepartures_429(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`rate limited`))
	}))
	t.Cleanup(srv.Close)
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetDepartures(context.Background(), "STOP1", nil)
	var rlErr *cumtd.RateLimitError
	if !errors.As(err, &rlErr) {
		t.Fatalf("expected *RateLimitError, got %v", err)
	}
}

func TestGetDepartures_emptyStopID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetDepartures(context.Background(), "", nil)
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}
