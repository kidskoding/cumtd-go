package types

// Trip is a single scheduled run of a route from origin to destination.
// Returned by [cumtd.Client.GetTrips], [cumtd.Client.GetTrip],
// [cumtd.Client.GetStopTrips], and embedded in other responses.
type Trip struct {
	// ID is the unique trip identifier.
	ID string `json:"id"`
	// BlockID groups trips served consecutively by the same vehicle.
	BlockID string `json:"blockId"`
	// ShapeID references the geographic shape for this trip's path.
	ShapeID string `json:"shapeId"`
	// Headsign is the destination text displayed on the bus (e.g. "North Terminal").
	Headsign  string         `json:"headsign"`
	Direction *TripDirection `json:"direction"`
	// Route is the route this trip belongs to. Nil in some list contexts.
	Route *Route `json:"route"`
}
