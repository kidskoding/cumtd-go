package types

type RouteGroup struct {
	ID             string  `json:"id"`
	SortNumber     any     `json:"sortNumber"`
	RouteGroupName string  `json:"routeGroupName"`
	Color          string  `json:"color"`
	TextColor      string  `json:"textColor"`
	Routes         []Route `json:"routes"`
}

type Route struct {
	ID                    string  `json:"id"`
	Number                *string `json:"number"`
	FirstTrip             string  `json:"firstTrip"`
	LastTrip              string  `json:"lastTrip"`
	LastTripAfterMidnight bool    `json:"lastTripAfterMidnight"`
	DayType               any     `json:"dayType"`
	GtfsRoutes            []any   `json:"gtfsRoutes"`
	RouteGroupID          string  `json:"routeGroupId"`
}
