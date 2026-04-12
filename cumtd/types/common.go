package types

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TripDirection struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DayType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
