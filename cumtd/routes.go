package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetRouteGroups returns all route groups.
func (c *Client) GetRouteGroups(ctx context.Context) ([]types.RouteGroup, error) {
	var out []types.RouteGroup
	return out, c.get(ctx, "/routes/groups", nil, &out)
}

// GetRouteGroup returns a single route group by ID.
func (c *Client) GetRouteGroup(ctx context.Context, routeGroupID string) (*types.RouteGroup, error) {
	if err := require("routeGroupID", routeGroupID); err != nil {
		return nil, err
	}
	var out types.RouteGroup
	return &out, c.get(ctx, "/routes/groups/"+url.PathEscape(routeGroupID), nil, &out)
}

// GetRoute returns a single route by ID.
func (c *Client) GetRoute(ctx context.Context, routeID string) (*types.Route, error) {
	if err := require("routeID", routeID); err != nil {
		return nil, err
	}
	var out types.Route
	return &out, c.get(ctx, "/routes/"+url.PathEscape(routeID), nil, &out)
}
