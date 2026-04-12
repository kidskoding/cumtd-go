package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetVehicles returns all vehicles in the MTD fleet.
func (c *Client) GetVehicles(ctx context.Context) ([]types.Vehicle, error) {
	var out []types.Vehicle
	return out, c.get(ctx, "/vehicles", nil, &out)
}

// GetVehicle returns a single vehicle by its ID.
// Returns [*APIError] with StatusCode 404 if the vehicle does not exist.
func (c *Client) GetVehicle(ctx context.Context, vehicleID string) (*types.Vehicle, error) {
	if err := require("vehicleID", vehicleID); err != nil {
		return nil, err
	}
	var out types.Vehicle
	return &out, c.get(ctx, "/vehicles/"+url.PathEscape(vehicleID), nil, &out)
}

// GetVehicleLocation returns the current real-time location of a single vehicle.
func (c *Client) GetVehicleLocation(ctx context.Context, vehicleID string) (*types.VehicleLocation, error) {
	if err := require("vehicleID", vehicleID); err != nil {
		return nil, err
	}
	var out types.VehicleLocation
	return &out, c.get(ctx, "/vehicles/"+url.PathEscape(vehicleID)+"/location", nil, &out)
}

// GetVehicleLocations returns real-time locations for all currently active vehicles.
func (c *Client) GetVehicleLocations(ctx context.Context) ([]types.VehicleLocation, error) {
	var out []types.VehicleLocation
	return out, c.get(ctx, "/vehicles/locations", nil, &out)
}

// GetVehicleConfigurations returns all vehicle configurations (bus types, capacity, powertrain).
func (c *Client) GetVehicleConfigurations(ctx context.Context) ([]types.VehicleConfiguration, error) {
	var out []types.VehicleConfiguration
	return out, c.get(ctx, "/vehicles/configurations", nil, &out)
}

// GetVehicleConfiguration returns a single vehicle configuration by its ID.
func (c *Client) GetVehicleConfiguration(ctx context.Context, configID string) (*types.VehicleConfiguration, error) {
	if err := require("configID", configID); err != nil {
		return nil, err
	}
	var out types.VehicleConfiguration
	return &out, c.get(ctx, "/vehicles/configurations/"+url.PathEscape(configID), nil, &out)
}
