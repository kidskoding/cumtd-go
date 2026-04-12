package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetShape returns a shape (sequence of lat/lon points) by its ID.
// Uses the /shapes/{id} path (plural) per the upstream API spec.
func (c *Client) GetShape(ctx context.Context, shapeID string) (*types.Shape, error) {
	if err := require("shapeID", shapeID); err != nil {
		return nil, err
	}
	var out types.Shape
	return &out, c.get(ctx, "/shapes/"+url.PathEscape(shapeID), nil, &out)
}

// GetShapePolyline returns the Google-encoded polyline string for a shape.
// Uses the /shape/{id}/polyline path (singular) — this is an intentional
// inconsistency in the upstream API spec; the path is matched exactly.
func (c *Client) GetShapePolyline(ctx context.Context, shapeID string) (*types.ShapePolyline, error) {
	if err := require("shapeID", shapeID); err != nil {
		return nil, err
	}
	var out types.ShapePolyline
	return &out, c.get(ctx, "/shape/"+url.PathEscape(shapeID)+"/polyline", nil, &out)
}
