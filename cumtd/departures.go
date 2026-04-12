package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetDeparturesOptions controls optional parameters for [Client.GetDepartures].
type GetDeparturesOptions struct {
	// Routes filters departures to the given route IDs. Accepts a
	// comma-separated list of route IDs (e.g. "ILLINI,GOLDLINE").
	Routes string
	// Time filters departures to those occurring at or after this time.
	// Format: HH:MM:SS (24-hour). Defaults to now when empty.
	Time string
}

// GetDepartures returns upcoming departures for a stop. This is the primary
// real-time endpoint — it includes both scheduled and live vehicle data.
//
// Use [GetDeparturesOptions] to filter by route or start time:
//
//	deps, err := client.GetDepartures(ctx, "STOP1", &cumtd.GetDeparturesOptions{
//	    Routes: "ILLINI",
//	    Time:   "08:00:00",
//	})
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
