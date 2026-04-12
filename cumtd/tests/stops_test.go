package cumtd_test

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestGetStops(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops": {StatusCode: 200, Body: testutil.MustReadFixture(t, "stops.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	stops, err := c.GetStops(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stops) != 2 {
		t.Errorf("len(stops) = %d, want 2", len(stops))
	}
}

func TestGetStops_excludeBoardingPoints(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops": {
			StatusCode: 200,
			Body:       testutil.MustReadFixture(t, "stops.json"),
			AssertQuery: func(t *testing.T, q url.Values) {
				if q.Get("excludeBoardingPoints") != "true" {
					t.Errorf("excludeBoardingPoints = %q, want true", q.Get("excludeBoardingPoints"))
				}
			},
		},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetStops(context.Background(), &cumtd.GetStopsOptions{ExcludeBoardingPoints: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetStop(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/STOP1": {StatusCode: 200, Body: testutil.MustReadFixture(t, "stop.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	s, err := c.GetStop(context.Background(), "STOP1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.ID != "STOP1" {
		t.Errorf("ID = %q, want STOP1", s.ID)
	}
}

func TestGetStop_404(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/NOPE": {StatusCode: 404, Body: []byte(`not found`)},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetStop(context.Background(), "NOPE")
	var apiErr *cumtd.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %v", err)
	}
}

func TestGetStop_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetStop(context.Background(), "")
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}

func TestSearchStops(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/search": {StatusCode: 200, Body: testutil.MustReadFixture(t, "stop_search.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	results, err := c.SearchStops(context.Background(), "main")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("len(results) = %d, want 2", len(results))
	}
}

func TestGetStopSchedule(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/STOP1/schedule": {StatusCode: 200, Body: testutil.MustReadFixture(t, "stop_schedule.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	times, err := c.GetStopSchedule(context.Background(), "STOP1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(times) != 2 {
		t.Errorf("len(times) = %d, want 2", len(times))
	}
}

func TestGetStopSchedule_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetStopSchedule(context.Background(), "", nil)
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}

func TestGetStopTrips(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/STOP1/trips": {StatusCode: 200, Body: testutil.MustReadFixture(t, "stop_trips.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	trips, err := c.GetStopTrips(context.Background(), "STOP1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(trips) != 2 {
		t.Errorf("len(trips) = %d, want 2", len(trips))
	}
}

func TestGetStopRouteGroups(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/stops/STOP1/route-groups": {StatusCode: 200, Body: testutil.MustReadFixture(t, "stop_route_groups.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	groups, err := c.GetStopRouteGroups(context.Background(), "STOP1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(groups) != 2 {
		t.Errorf("len(groups) = %d, want 2", len(groups))
	}
}
