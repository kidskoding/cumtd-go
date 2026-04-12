package types

// RouteGroup is a named collection of related routes displayed together in
// the MTD app (e.g. all Illini route variants belong to the Illini group).
type RouteGroup struct {
	// ID is the unique route group identifier.
	ID string `json:"id"`
	// SortNumber controls display order. Typed as any because the API may
	// return an int or string; use internal/coerce to convert.
	SortNumber     any    `json:"sortNumber"`
	RouteGroupName string `json:"routeGroupName"`
	// Color is the route group brand color as a hex string (e.g. "#FF6600").
	Color string `json:"color"`
	// TextColor is the foreground color for text on Color backgrounds.
	TextColor string  `json:"textColor"`
	Routes    []Route `json:"routes"`
}

// Route is a single route within a [RouteGroup].
type Route struct {
	// ID is the unique route identifier.
	ID string `json:"id"`
	// Number is the public-facing route number. Nil if not assigned.
	Number    *string `json:"number"`
	FirstTrip string  `json:"firstTrip"`
	LastTrip  string  `json:"lastTrip"`
	// LastTripAfterMidnight is true when the last trip ends after midnight.
	LastTripAfterMidnight bool `json:"lastTripAfterMidnight"`
	// DayType may be a string or object depending on context.
	DayType    any    `json:"dayType"`
	GtfsRoutes []any  `json:"gtfsRoutes"`
	RouteGroupID string `json:"routeGroupId"`
}
