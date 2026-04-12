package cumtd_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kidskoding/cumtd-go/cumtd"
	"github.com/kidskoding/cumtd-go/internal/testutil"
)

func TestGetShape(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		// plural /shapes/{id} per upstream spec
		"/shapes/SHAPE1": {StatusCode: 200, Body: testutil.MustReadFixture(t, "shape.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	s, err := c.GetShape(context.Background(), "SHAPE1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.ID != "SHAPE1" {
		t.Errorf("ID = %q, want SHAPE1", s.ID)
	}
	if len(s.ShapePoints) != 3 {
		t.Errorf("len(ShapePoints) = %d, want 3", len(s.ShapePoints))
	}
}

func TestGetShape_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetShape(context.Background(), "")
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}

func TestGetShapePolyline(t *testing.T) {
	srv := testutil.NewMockServer(t, map[string]testutil.MockRoute{
		// singular /shape/{id}/polyline per upstream spec (intentional inconsistency)
		"/shape/SHAPE1/polyline": {StatusCode: 200, Body: testutil.MustReadFixture(t, "shape_polyline.json")},
	})
	c := cumtd.New("key", cumtd.WithBaseURL(srv.URL))
	p, err := c.GetShapePolyline(context.Background(), "SHAPE1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.ID != "SHAPE1" {
		t.Errorf("ID = %q, want SHAPE1", p.ID)
	}
	if p.Polyline == "" {
		t.Error("Polyline should not be empty")
	}
}

func TestGetShapePolyline_emptyID(t *testing.T) {
	c := cumtd.New("key")
	_, err := c.GetShapePolyline(context.Background(), "")
	var valErr *cumtd.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *ValidationError, got %v", err)
	}
}
