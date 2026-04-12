package cumtd_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestGetTrips(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/trips": {StatusCode: 200, Body: testutil.MustReadFixture(t, "trips.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	trips, err := c.GetTrips(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(trips) != 2 {
		t.Errorf("len(trips) = %d, want 2", len(trips))
	}
}

func TestGetTrip(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/trips/TRIP1": {StatusCode: 200, Body: testutil.MustReadFixture(t, "trip.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	trip, err := c.GetTrip(context.Background(), "TRIP1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if trip.ID != "TRIP1" {
		t.Errorf("ID = %q, want TRIP1", trip.ID)
	}
}

func TestGetTrip_404(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/trips/NOPE": {StatusCode: 404, Body: []byte(`not found`)},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetTrip(context.Background(), "NOPE")
	var apiErr *cumtd.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %v", err)
	}
}

func TestGetTrip_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetTrip(context.Background(), "")
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}
