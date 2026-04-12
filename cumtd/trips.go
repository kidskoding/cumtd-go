package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetTrips returns all trips.
func (c *Client) GetTrips(ctx context.Context) ([]types.Trip, error) {
	var out []types.Trip
	return out, c.get(ctx, "/trips", nil, &out)
}

// GetTrip returns a single trip by ID.
func (c *Client) GetTrip(ctx context.Context, tripID string) (*types.Trip, error) {
	if err := require("tripID", tripID); err != nil {
		return nil, err
	}
	var out types.Trip
	return &out, c.get(ctx, "/trips/"+url.PathEscape(tripID), nil, &out)
}
