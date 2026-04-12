package cumtd_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestGetVehicles(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/vehicles": {StatusCode: 200, Body: testutil.MustReadFixture(t, "vehicles.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	vehs, err := c.GetVehicles(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(vehs) != 2 {
		t.Errorf("len(vehs) = %d, want 2", len(vehs))
	}
}

func TestGetVehicle(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/vehicles/VEH1": {StatusCode: 200, Body: testutil.MustReadFixture(t, "vehicle.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	v, err := c.GetVehicle(context.Background(), "VEH1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.ID != "VEH1" {
		t.Errorf("ID = %q, want VEH1", v.ID)
	}
}

func TestGetVehicle_404(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/vehicles/NOPE": {StatusCode: 404, Body: []byte(`not found`)},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetVehicle(context.Background(), "NOPE")
	var apiErr *cumtd.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %v", err)
	}
}

func TestGetVehicle_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetVehicle(context.Background(), "")
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}

func TestGetVehicleLocations(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/vehicles/locations": {StatusCode: 200, Body: testutil.MustReadFixture(t, "vehicle_locations.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	locs, err := c.GetVehicleLocations(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(locs) != 2 {
		t.Errorf("len(locs) = %d, want 2", len(locs))
	}
}

func TestGetVehicleLocation(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/vehicles/VEH1/location": {StatusCode: 200, Body: testutil.MustReadFixture(t, "vehicle_location.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	loc, err := c.GetVehicleLocation(context.Background(), "VEH1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc.ID != "VEH1" {
		t.Errorf("ID = %q, want VEH1", loc.ID)
	}
}

func TestGetVehicleConfigurations(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/vehicles/configurations": {StatusCode: 200, Body: testutil.MustReadFixture(t, "vehicle_configurations.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	cfgs, err := c.GetVehicleConfigurations(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfgs) != 2 {
		t.Errorf("len(cfgs) = %d, want 2", len(cfgs))
	}
}

func TestGetVehicleConfiguration(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/vehicles/configurations/CONFIG1": {StatusCode: 200, Body: testutil.MustReadFixture(t, "vehicle_configuration.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	cfg, err := c.GetVehicleConfiguration(context.Background(), "CONFIG1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ID != "CONFIG1" {
		t.Errorf("ID = %q, want CONFIG1", cfg.ID)
	}
}
