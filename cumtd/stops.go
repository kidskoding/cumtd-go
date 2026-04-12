package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetStopsOptions controls optional parameters for GetStops.
type GetStopsOptions struct {
	ExcludeBoardingPoints bool
}

// GetStopScheduleOptions controls optional parameters for GetStopSchedule.
type GetStopScheduleOptions struct {
	RouteID string
	Date    string // YYYY-MM-DD
}

// GetStops returns all stops.
func (c *Client) GetStops(ctx context.Context, opts *GetStopsOptions) ([]types.StopBase, error) {
	params := url.Values{}
	if opts != nil && opts.ExcludeBoardingPoints {
		params.Set("excludeBoardingPoints", "true")
	}
	var out []types.StopBase
	return out, c.get(ctx, "/stops", params, &out)
}

// GetStop returns a single stop by ID.
func (c *Client) GetStop(ctx context.Context, stopID string) (*types.StopBase, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	var out types.StopBase
	return &out, c.get(ctx, "/stops/"+url.PathEscape(stopID), nil, &out)
}

// SearchStops searches stops by query string.
func (c *Client) SearchStops(ctx context.Context, query string) ([]types.StopSearchResult, error) {
	params := url.Values{"query": {query}}
	var out []types.StopSearchResult
	return out, c.get(ctx, "/stops/search", params, &out)
}

// GetStopSchedule returns the schedule for a stop.
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

// GetStopTrips returns trips serving a stop.
func (c *Client) GetStopTrips(ctx context.Context, stopID string) ([]types.Trip, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	var out []types.Trip
	return out, c.get(ctx, "/stops/"+url.PathEscape(stopID)+"/trips", nil, &out)
}

// GetStopRouteGroups returns route groups serving a stop.
func (c *Client) GetStopRouteGroups(ctx context.Context, stopID string) ([]types.RouteGroup, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	var out []types.RouteGroup
	return out, c.get(ctx, "/stops/"+url.PathEscape(stopID)+"/route-groups", nil, &out)
}
