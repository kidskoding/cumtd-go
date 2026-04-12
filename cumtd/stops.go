package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetStopsOptions controls optional parameters for [Client.GetStops].
type GetStopsOptions struct {
	// ExcludeBoardingPoints omits boarding point details from each stop when
	// true. Use this to reduce response size when you only need stop metadata.
	ExcludeBoardingPoints bool
}

// GetStopScheduleOptions controls optional parameters for [Client.GetStopSchedule].
type GetStopScheduleOptions struct {
	// RouteID filters the schedule to a specific route.
	RouteID string
	// Date filters the schedule to a specific service date in YYYY-MM-DD format.
	// Defaults to today when empty.
	Date string
}

// GetStops returns all stops in the MTD system.
// Pass [GetStopsOptions] to exclude boarding point details.
func (c *Client) GetStops(ctx context.Context, opts *GetStopsOptions) ([]types.StopBase, error) {
	params := url.Values{}
	if opts != nil && opts.ExcludeBoardingPoints {
		params.Set("excludeBoardingPoints", "true")
	}
	var out []types.StopBase
	return out, c.get(ctx, "/stops", params, &out)
}

// GetStop returns a single stop by its ID.
// Returns [*APIError] with StatusCode 404 if the stop does not exist.
func (c *Client) GetStop(ctx context.Context, stopID string) (*types.StopBase, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	var out types.StopBase
	return &out, c.get(ctx, "/stops/"+url.PathEscape(stopID), nil, &out)
}

// SearchStops searches for stops by name or code. Returns stops whose name or
// code contains the query string (case-insensitive).
func (c *Client) SearchStops(ctx context.Context, query string) ([]types.StopSearchResult, error) {
	params := url.Values{"query": {query}}
	var out []types.StopSearchResult
	return out, c.get(ctx, "/stops/search", params, &out)
}

// GetStopSchedule returns the scheduled stop times for a stop.
// Use [GetStopScheduleOptions] to filter by route or date.
func (c *Client) GetStopSchedule(ctx context.Context, stopID string, opts *GetStopScheduleOptions) ([]types.StopTime, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	params := url.Values{}
	if opts != nil {
		if opts.RouteID != "" {
			params.Set("routeId", opts.RouteID)
		}
		if opts.Date != "" {
			params.Set("date", opts.Date)
		}
	}
	var out []types.StopTime
	return out, c.get(ctx, "/stops/"+url.PathEscape(stopID)+"/schedule", params, &out)
}

// GetStopTrips returns all trips that serve the given stop.
func (c *Client) GetStopTrips(ctx context.Context, stopID string) ([]types.Trip, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	var out []types.Trip
	return out, c.get(ctx, "/stops/"+url.PathEscape(stopID)+"/trips", nil, &out)
}

// GetStopRouteGroups returns the route groups that serve the given stop.
func (c *Client) GetStopRouteGroups(ctx context.Context, stopID string) ([]types.RouteGroup, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	var out []types.RouteGroup
	return out, c.get(ctx, "/stops/"+url.PathEscape(stopID)+"/route-groups", nil, &out)
}
