package types

// Departure is an upcoming bus departure from a stop, combining scheduled and
// real-time data. Returned by [cumtd.Client.GetDepartures].
type Departure struct {
	// StopID is the stop this departure departs from.
	StopID string `json:"stopId"`
	// Headsign is the destination sign shown on the bus. Nil if unavailable.
	Headsign *string `json:"headsign"`
	// Trip is the trip this departure belongs to. Nil if unassigned.
	Trip *DepartureTrip `json:"trip"`
	// BlockID groups trips that are served by the same vehicle in sequence.
	BlockID *string `json:"blockId"`
	// RecordedTime is the RFC3339 timestamp when this prediction was recorded.
	RecordedTime string `json:"recordedTime"`
	// ScheduledDeparture is the timetabled departure time (HH:MM:SS). Nil if unknown.
	ScheduledDeparture *string `json:"scheduledDeparture"`
	// EstimatedDeparture is the real-time predicted departure time (HH:MM:SS).
	// Nil when IsRealTime is false.
	EstimatedDeparture *string `json:"estimatedDeparture"`
	// VehicleID identifies the bus serving this departure. Nil if not yet assigned.
	VehicleID *string `json:"vehicleId"`
	// OriginStopID is the first stop on the trip. Nil if unknown.
	OriginStopID *string `json:"originStopId"`
	// DestinationStopID is the last stop on the trip. Nil if unknown.
	DestinationStopID *string `json:"destinationStopId"`
	// Location is the vehicle's current GPS position. Nil if unavailable.
	Location *Coordinates `json:"location"`
	// ShapeID references the route shape for this departure. Nil if unknown.
	ShapeID *string `json:"shapeId"`
	// MinutesTillDeparture is the number of minutes until departure. Typed as
	// any because the API may return an int or string; use internal/coerce to convert.
	MinutesTillDeparture any `json:"minutesTillDeparture"`
	// IsRealTime is true when EstimatedDeparture is based on live vehicle data.
	IsRealTime bool `json:"isRealTime"`
	// IsHopper is true for on-demand hopper service departures.
	IsHopper bool `json:"isHopper"`
	// Destination is the human-readable destination name. Nil if unavailable.
	Destination *string `json:"destination"`
	// DepartsIn is a human-readable string like "6 min" or "Due".
	DepartsIn string `json:"departsIn"`
	// IsIStop is true when the stop is an iStop (accessibility-enhanced stop).
	IsIStop bool `json:"isIStop"`
	// UniqueID is a stable identifier for this specific departure prediction.
	UniqueID string `json:"uniqueId"`
	// Route is the route this departure runs on. Nil if unknown.
	Route *DepartureRoute `json:"route"`
}

// DepartureTrip is the trip assignment embedded in a [Departure].
type DepartureTrip struct {
	// TripID is the trip identifier. Nil if not yet assigned.
	TripID    *string        `json:"tripId"`
	Direction *TripDirection `json:"direction"`
}

// DepartureRoute is the route summary embedded in a [Departure] or [VehicleLocation].
type DepartureRoute struct {
	// ID is the route identifier.
	ID string `json:"id"`
	// RouteGroupID references the parent [RouteGroup]. Nil if ungrouped.
	RouteGroupID *string `json:"routeGroupId"`
	// GtfsRouteID is the GTFS feed route identifier.
	GtfsRouteID string `json:"gtfsRouteId"`
	// LongName is the full route name. Nil if not provided.
	LongName *string `json:"longName"`
	// ShortName is the abbreviated route name (e.g. "22"). Nil if not provided.
	ShortName *string `json:"shortName"`
	// Color is the route brand color as a hex string. Nil if not assigned.
	Color *string `json:"color"`
	// TextColor is the foreground color for text on Color backgrounds. Nil if not assigned.
	TextColor *string `json:"textColor"`
}
