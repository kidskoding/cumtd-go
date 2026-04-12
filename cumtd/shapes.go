package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetShape returns a shape by ID.
// NOTE: uses /shapes/{id} (plural) per the upstream spec.
func (c *Client) GetShape(ctx context.Context, shapeID string) (*types.Shape, error) {
	if err := require("shapeID", shapeID); err != nil {
		return nil, err
	}
	var out types.Shape
	return &out, c.get(ctx, "/shapes/"+url.PathEscape(shapeID), nil, &out)
}

// GetShapePolyline returns the encoded polyline for a shape.
// NOTE: uses /shape/{id}/polyline (singular — upstream spec inconsistency, matched exactly).
func (c *Client) GetShapePolyline(ctx context.Context, shapeID string) (*types.ShapePolyline, error) {
	if err := require("shapeID", shapeID); err != nil {
		return nil, err
	}
	var out types.ShapePolyline
	return &out, c.get(ctx, "/shape/"+url.PathEscape(shapeID)+"/polyline", nil, &out)
}
