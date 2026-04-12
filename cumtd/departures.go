package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetDeparturesOptions controls optional parameters for GetDepartures.
type GetDeparturesOptions struct {
	Routes string // comma-separated route IDs
	Time   string // HH:MM:SS
}

// GetDepartures returns upcoming departures for a stop.
// This is the primary real-time endpoint.
func (c *Client) GetDepartures(ctx context.Context, stopID string, opts *GetDeparturesOptions) ([]types.Departure, error) {
	if err := require("stopID", stopID); err != nil {
		return nil, err
	}
	params := url.Values{}
	if opts != nil {
		if opts.Routes != "" {
			params.Set("routes", opts.Routes)
		}
		if opts.Time != "" {
			params.Set("time", opts.Time)
		}
	}
	var out []types.Departure
	return out, c.get(ctx, "/stops/"+url.PathEscape(stopID)+"/departures", params, &out)
}
