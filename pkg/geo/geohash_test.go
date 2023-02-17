package geo

import (
	"testing"
)

func TestGeoHash(t *testing.T) {
	p := Point{Lat: 34.1234, Lon: -117.9876}
	hash := GeoHash(p, 9)
	if hash != "9qh455up1" {
		t.Fatalf("GeoHash %q, want %q\n", hash, "9qh455up1")
	}
}
