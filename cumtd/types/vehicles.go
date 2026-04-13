package types

// Vehicle is a bus in the MTD fleet.
type Vehicle struct {
	// ID is the unique vehicle identifier.
	ID string `json:"id"`
	// VehicleConfigurationID references the vehicle's [VehicleConfiguration].
	VehicleConfigurationID string `json:"vehicleConfigurationId"`
	// IsActive is true when the vehicle is currently in service.
	IsActive bool `json:"isActive"`
	// DateInService is the date the vehicle entered service (YYYY-MM-DD).
	// Nil if not recorded.
	DateInService *string `json:"dateInService"`
}

// VehicleLocation is the real-time position and trip assignment of a vehicle.
type VehicleLocation struct {
	// ID is the vehicle identifier.
	ID string `json:"id"`
	// Location is the vehicle's current GPS position. Nil if unavailable.
	Location *Coordinates `json:"location"`
	// LastUpdated is the RFC3339 timestamp of the last GPS update. Nil if unknown.
	LastUpdated *string `json:"lastUpdated"`
	// Trip is the trip the vehicle is currently running. Nil if not on a trip.
	Trip *VehicleLocationTrip `json:"trip"`
	// Route is the route the vehicle is currently serving. Nil if not on a route.
	Route *DepartureRoute `json:"route"`
}

// VehicleLocationTrip is the trip currently assigned to a vehicle.
type VehicleLocationTrip struct {
	// TripID is the active trip identifier. Nil if not yet assigned.
	TripID    *string        `json:"tripId"`
	Direction *TripDirection `json:"direction"`
}

// VehicleConfiguration describes the physical characteristics of a vehicle model.
type VehicleConfiguration struct {
	// ID is the unique configuration identifier (UUID).
	ID string `json:"id"`
	// VehicleType is the vehicle class (e.g. "bus").
	VehicleType VehicleType `json:"vehicleType"`
	// Year is the model year. The API may return an int or string; use internal/coerce to convert.
	Year any `json:"year"`
	// Make is the vehicle manufacturer (e.g. "Nova Bus").
	Make string `json:"make"`
	// Model is the vehicle model name.
	Model string `json:"model"`
	// LengthFeet is the vehicle length in feet. The API may return an int, string, or null.
	LengthFeet any `json:"lengthFeet"`
	// Powertrain is the propulsion technology (e.g. "diesel", "electric", "hybrid").
	Powertrain VehiclePowertrainType `json:"powertrain"`
}

// VehicleType is the class of a vehicle (e.g. "bus").
type VehicleType string

// VehiclePowertrainType is the propulsion technology of a vehicle
// (e.g. "diesel", "electric", "hybrid").
type VehiclePowertrainType string
