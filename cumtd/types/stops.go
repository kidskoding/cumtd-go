package types

type StopBase struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Code           *string         `json:"code"`
	Location       *Coordinates    `json:"location"`
	BoardingPoints []BoardingPoint `json:"boardingPoints"`
	StopGroups     []StopGroup     `json:"stopGroups"`
}

type BoardingPoint struct {
	ID       string       `json:"id"`
	Name     *string      `json:"name"`
	Location *Coordinates `json:"location"`
}

type StopGroup struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

type StopSearchResult struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Code     *string      `json:"code"`
	Location *Coordinates `json:"location"`
}

type StopTime struct {
	StopID                string         `json:"stopId"`
	TripID                string         `json:"tripId"`
	RouteID               *string        `json:"routeId"`
	GtfsRouteID           *string        `json:"gtfsRouteId"`
	Direction             *TripDirection `json:"direction"`
	StopSequence          any            `json:"stopSequence"`
	ArrivalTime           string         `json:"arrivalTime"`
	ArrivalPastMidnight   bool           `json:"arrivalPastMidnight"`
	DepartureTime         string         `json:"departureTime"`
	DeparturePastMidnight bool           `json:"departurePastMidnight"`
	StopHeadsign          *string        `json:"stopHeadsign"`
}
