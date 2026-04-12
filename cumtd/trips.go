package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetTrips returns all trips in the MTD system.
func (c *Client) GetTrips(ctx context.Context) ([]types.Trip, error) {
	var out []types.Trip
	return out, c.get(ctx, "/trips", nil, &out)
}

// GetTrip returns a single trip by its ID.
// Returns [*APIError] with StatusCode 404 if the trip does not exist.
func (c *Client) GetTrip(ctx context.Context, tripID string) (*types.Trip, error) {
	if err := require("tripID", tripID); err != nil {
		return nil, err
	}
	var out types.Trip
	return &out, c.get(ctx, "/trips/"+url.PathEscape(tripID), nil, &out)
}
