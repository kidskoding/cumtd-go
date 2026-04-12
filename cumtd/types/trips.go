package types

type Trip struct {
	ID        string         `json:"id"`
	BlockID   string         `json:"blockId"`
	ShapeID   string         `json:"shapeId"`
	Headsign  string         `json:"headsign"`
	Direction *TripDirection `json:"direction"`
	Route     *Route         `json:"route"`
}
