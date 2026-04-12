package types

type Vehicle struct {
	ID                     string  `json:"id"`
	VehicleConfigurationID string  `json:"vehicleConfigurationId"`
	IsActive               bool    `json:"isActive"`
	DateInService          *string `json:"dateInService"`
}

type VehicleLocation struct {
	ID          string               `json:"id"`
	Location    *Coordinates         `json:"location"`
	LastUpdated *string              `json:"lastUpdated"`
	Trip        *VehicleLocationTrip `json:"trip"`
	Route       *DepartureRoute      `json:"route"`
}

type VehicleLocationTrip struct {
	TripID    *string        `json:"tripId"`
	Direction *TripDirection `json:"direction"`
}

type VehicleConfiguration struct {
	ID             string                `json:"id"`
	Type           VehicleType           `json:"type"`
	PowertrainType VehiclePowertrainType `json:"powertrainType"`
	Capacity       int                   `json:"capacity"`
}

type VehicleType string
type VehiclePowertrainType string
