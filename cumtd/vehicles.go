package cumtd

import (
	"context"
	"net/url"

	"github.com/kidskoding/cumtd-go/cumtd/types"
)

// GetVehicles returns all vehicles.
func (c *Client) GetVehicles(ctx context.Context) ([]types.Vehicle, error) {
	var out []types.Vehicle
	return out, c.get(ctx, "/vehicles", nil, &out)
}

// GetVehicle returns a single vehicle by ID.
func (c *Client) GetVehicle(ctx context.Context, vehicleID string) (*types.Vehicle, error) {
	if err := require("vehicleID", vehicleID); err != nil {
		return nil, err
	}
	var out types.Vehicle
	return &out, c.get(ctx, "/vehicles/"+url.PathEscape(vehicleID), nil, &out)
}

// GetVehicleLocation returns the current location of a vehicle.
func (c *Client) GetVehicleLocation(ctx context.Context, vehicleID string) (*types.VehicleLocation, error) {
	if err := require("vehicleID", vehicleID); err != nil {
		return nil, err
	}
	var out types.VehicleLocation
	return &out, c.get(ctx, "/vehicles/"+url.PathEscape(vehicleID)+"/location", nil, &out)
}

// GetVehicleLocations returns locations of all active vehicles.
func (c *Client) GetVehicleLocations(ctx context.Context) ([]types.VehicleLocation, error) {
	var out []types.VehicleLocation
	return out, c.get(ctx, "/vehicles/locations", nil, &out)
}

// GetVehicleConfigurations returns all vehicle configurations.
func (c *Client) GetVehicleConfigurations(ctx context.Context) ([]types.VehicleConfiguration, error) {
	var out []types.VehicleConfiguration
	return out, c.get(ctx, "/vehicles/configurations", nil, &out)
}

// GetVehicleConfiguration returns a single vehicle configuration by ID.
func (c *Client) GetVehicleConfiguration(ctx context.Context, configID string) (*types.VehicleConfiguration, error) {
	if err := require("configID", configID); err != nil {
		return nil, err
	}
	var out types.VehicleConfiguration
	return &out, c.get(ctx, "/vehicles/configurations/"+url.PathEscape(configID), nil, &out)
}
