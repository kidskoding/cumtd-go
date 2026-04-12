package types

type Departure struct {
	StopID               string          `json:"stopId"`
	Headsign             *string         `json:"headsign"`
	Trip                 *DepartureTrip  `json:"trip"`
	BlockID              *string         `json:"blockId"`
	RecordedTime         string          `json:"recordedTime"`
	ScheduledDeparture   *string         `json:"scheduledDeparture"`
	EstimatedDeparture   *string         `json:"estimatedDeparture"`
	VehicleID            *string         `json:"vehicleId"`
	OriginStopID         *string         `json:"originStopId"`
	DestinationStopID    *string         `json:"destinationStopId"`
	Location             *Coordinates    `json:"location"`
	ShapeID              *string         `json:"shapeId"`
	MinutesTillDeparture any             `json:"minutesTillDeparture"`
	IsRealTime           bool            `json:"isRealTime"`
	IsHopper             bool            `json:"isHopper"`
	Destination          *string         `json:"destination"`
	DepartsIn            string          `json:"departsIn"`
	IsIStop              bool            `json:"isIStop"`
	UniqueID             string          `json:"uniqueId"`
	Route                *DepartureRoute `json:"route"`
}

type DepartureTrip struct {
	TripID    *string        `json:"tripId"`
	Direction *TripDirection `json:"direction"`
}

type DepartureRoute struct {
	ID           string  `json:"id"`
	RouteGroupID *string `json:"routeGroupId"`
	GtfsRouteID  string  `json:"gtfsRouteId"`
	LongName     *string `json:"longName"`
	ShortName    *string `json:"shortName"`
	Color        *string `json:"color"`
	TextColor    *string `json:"textColor"`
}
