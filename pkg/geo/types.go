package geo

// Point defines the geographic coordinates:
// latitude and longitude.
type Point struct {
	Lat float64
	Lon float64
}

// Valid checks if the geographic coordinates
// of the point belong to the corresponding domains.
func (p Point) Valid() (latok, lonok bool) {
	latok = -90 <= p.Lat && p.Lat <= 90
	lonok = -180 <= p.Lon && p.Lon <= 180
	return
}
