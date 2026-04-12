package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetRouteGroups returns all route groups in the MTD system.
// Route groups are collections of related routes (e.g. the Illini group
// contains all Illini route variants).
func (c *Client) GetRouteGroups(ctx context.Context) ([]types.RouteGroup, error) {
	var out []types.RouteGroup
	return out, c.get(ctx, "/routes/groups", nil, &out)
}

// GetRouteGroup returns a single route group by its ID.
// Returns [*APIError] with StatusCode 404 if the route group does not exist.
func (c *Client) GetRouteGroup(ctx context.Context, routeGroupID string) (*types.RouteGroup, error) {
	if err := require("routeGroupID", routeGroupID); err != nil {
		return nil, err
	}
	var out types.RouteGroup
	return &out, c.get(ctx, "/routes/groups/"+url.PathEscape(routeGroupID), nil, &out)
}

// GetRoute returns a single route by its ID.
// Returns [*APIError] with StatusCode 404 if the route does not exist.
func (c *Client) GetRoute(ctx context.Context, routeID string) (*types.Route, error) {
	if err := require("routeID", routeID); err != nil {
		return nil, err
	}
	var out types.Route
	return &out, c.get(ctx, "/routes/"+url.PathEscape(routeID), nil, &out)
}
