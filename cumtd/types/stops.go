package types

// StopBase is a transit stop with its boarding points and group membership.
// Returned by [cumtd.Client.GetStops] and [cumtd.Client.GetStop].
type StopBase struct {
	// ID is the unique stop identifier (e.g. "MAIN/WRIGHT").
	ID string `json:"id"`
	// Name is the human-readable stop name.
	Name string `json:"name"`
	// Code is the short public-facing stop code. Nil if not assigned.
	Code *string `json:"code"`
	// Location is the stop's GPS position. Nil if not geocoded.
	Location *Coordinates `json:"location"`
	// BoardingPoints are the individual physical boarding locations within
	// this stop (e.g. different corners of an intersection).
	BoardingPoints []BoardingPoint `json:"boardingPoints"`
	// StopGroups are the named groups this stop belongs to.
	StopGroups []StopGroup `json:"stopGroups"`
}

// BoardingPoint is a specific physical location where passengers board within
// a [StopBase].
type BoardingPoint struct {
	// ID is the boarding point identifier.
	ID string `json:"id"`
	// Name describes the boarding point (e.g. "Northbound"). Nil if unnamed.
	Name *string `json:"name"`
	// Location is the GPS position of this boarding point. Nil if not geocoded.
	Location *Coordinates `json:"location"`
}

// StopGroup is a named cluster that a stop belongs to.
type StopGroup struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

// StopSearchResult is a lightweight stop record returned by
// [cumtd.Client.SearchStops].
type StopSearchResult struct {
	// ID is the unique stop identifier.
	ID string `json:"id"`
	// Name is the human-readable stop name.
	Name string `json:"name"`
	// Code is the short public-facing stop code. Nil if not assigned.
	Code *string `json:"code"`
	// Location is the stop's GPS position. Nil if not geocoded.
	Location *Coordinates `json:"location"`
}

// StopTime is a single scheduled arrival/departure at a stop, returned by
// [cumtd.Client.GetStopSchedule].
type StopTime struct {
	StopID      string  `json:"stopId"`
	TripID      string  `json:"tripId"`
	RouteID     *string `json:"routeId"`
	GtfsRouteID *string `json:"gtfsRouteId"`
	Direction   *TripDirection `json:"direction"`
	// StopSequence is the position of this stop within the trip. Typed as any
	// because the API may return an int or string; use internal/coerce to convert.
	StopSequence          any     `json:"stopSequence"`
	// ArrivalTime is the scheduled arrival time (HH:MM:SS).
	ArrivalTime           string  `json:"arrivalTime"`
	// ArrivalPastMidnight is true when ArrivalTime is after midnight on the service day.
	ArrivalPastMidnight   bool    `json:"arrivalPastMidnight"`
	// DepartureTime is the scheduled departure time (HH:MM:SS).
	DepartureTime         string  `json:"departureTime"`
	// DeparturePastMidnight is true when DepartureTime is after midnight on the service day.
	DeparturePastMidnight bool    `json:"departurePastMidnight"`
	// StopHeadsign overrides the trip headsign at this specific stop. Nil if not set.
	StopHeadsign          *string `json:"stopHeadsign"`
}
