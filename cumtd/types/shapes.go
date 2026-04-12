package types

type Shape struct {
	ID          string       `json:"id"`
	ShapePoints []ShapePoint `json:"shapePoints"`
}

type ShapePoint struct {
	Sequence int     `json:"sequence"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

type ShapePolyline struct {
	ID       string `json:"id"`
	Polyline string `json:"polyline"`
}
