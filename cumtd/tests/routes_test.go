package cumtd_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestGetRouteGroups(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/routes/groups": {StatusCode: 200, Body: testutil.MustReadFixture(t, "route_groups.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	groups, err := c.GetRouteGroups(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(groups) != 2 {
		t.Errorf("len(groups) = %d, want 2", len(groups))
	}
	if groups[0].ID != "ILLINI" {
		t.Errorf("groups[0].ID = %q, want ILLINI", groups[0].ID)
	}
}

func TestGetRouteGroup(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/routes/groups/ILLINI": {StatusCode: 200, Body: testutil.MustReadFixture(t, "route_group.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	g, err := c.GetRouteGroup(context.Background(), "ILLINI")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.ID != "ILLINI" {
		t.Errorf("ID = %q, want ILLINI", g.ID)
	}
}

func TestGetRouteGroup_404(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/routes/groups/NOPE": {StatusCode: 404, Body: []byte(`not found`)},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	_, err := c.GetRouteGroup(context.Background(), "NOPE")
	var apiErr *cumtd.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %v", err)
	}
}

func TestGetRouteGroup_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetRouteGroup(context.Background(), "")
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}

func TestGetRoute(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		"/routes/ILLINI": {StatusCode: 200, Body: testutil.MustReadFixture(t, "route.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	r, err := c.GetRoute(context.Background(), "ILLINI")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.ID != "ILLINI" {
		t.Errorf("ID = %q, want ILLINI", r.ID)
	}
}

func TestGetRoute_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetRoute(context.Background(), "")
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}
