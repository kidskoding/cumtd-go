package cumtd

import (
	"context"
	"net/url"
	"strconv"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetDeparturesOptions controls optional parameters for [Client.GetDepartures].
type GetDeparturesOptions struct {
	// Routes filters departures to the given route IDs. Accepts a
	// comma-separated list of route IDs (e.g. "ILLINI,GOLDLINE").
	Routes string
	// MinutesAhead limits results to departures within this many minutes.
	// Valid range: 0–120. Defaults to 30 when zero.
	MinutesAhead int
}

// GetDepartures returns upcoming departures for a stop. This is the primary
// real-time endpoint — it includes both scheduled and live vehicle data.
//
// Use [GetDeparturesOptions] to filter by route or time window:
//
//	deps, err := client.GetDepartures(ctx, "STOP1", &cumtd.GetDeparturesOptions{
//	    Routes:       "ILLINI",
//	    MinutesAhead: 60,
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
		if opts.MinutesAhead != 0 {
			params.Set("time", strconv.Itoa(opts.MinutesAhead))
		}
	}
	var out []types.Departure
	return out, c.get(ctx, "/stops/"+url.PathEscape(stopID)+"/departures", params, &out)
}
