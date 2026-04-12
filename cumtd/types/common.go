// Package types contains the domain types returned by the MTD API v3.
// All types map directly to the JSON shapes in the API response envelope.
package types

// Coordinates is a geographic lat/lon pair.
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// TripDirection identifies the direction of travel for a trip.
type TripDirection struct {
	// ID is the direction identifier (e.g. "N", "S", "E", "W").
	ID string `json:"id"`
	// Name is the human-readable direction name (e.g. "Northbound").
	Name string `json:"name"`
}

// DayType identifies the service day type for a trip or route.
type DayType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
