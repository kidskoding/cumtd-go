package types

// Shape is the geographic path of a trip as an ordered sequence of points.
// Returned by [cumtd.Client.GetShape].
type Shape struct {
	// ID is the unique shape identifier.
	ID string `json:"id"`
	// ShapePoints is the ordered sequence of lat/lon points that define the path.
	ShapePoints []ShapePoint `json:"shapePoints"`
}

// ShapePoint is a single lat/lon waypoint in a [Shape].
type ShapePoint struct {
	// Sequence is the 1-based position of this point in the shape.
	Sequence int `json:"sequence"`
	// Lat is the latitude in decimal degrees.
	Lat float64 `json:"lat"`
	// Lon is the longitude in decimal degrees.
	Lon float64 `json:"lon"`
}

// ShapePolyline is a [Shape] encoded as a Google-encoded polyline string.
// Returned by [cumtd.Client.GetShapePolyline]. Decode with any polyline
// library (e.g. github.com/twpayne/go-polyline).
type ShapePolyline struct {
	// ID is the shape identifier.
	ID string `json:"id"`
	// Polyline is the Google-encoded polyline string representing the shape path.
	Polyline string `json:"polyline"`
}
