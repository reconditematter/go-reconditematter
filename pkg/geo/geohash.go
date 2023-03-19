package geo

import (
	"math"
	"strings"
)

// GeoHash -- computes the geohash of length `n` of the given
// geographic latitude and longitude stored in `point`.
//
// See: https://en.wikipedia.org/wiki/Geohash
//
// The given length `n` is assumed to be 3≤n≤15.
// When n<3, it is set to 3; when n>15, it is set to 15.
func GeoHash(point Point, n uint) string {
	latok, lonok := point.Valid()
	if !(latok && lonok) {
		panic("GeoHash: invalid point")
	}

	if n < 3 {
		n = 3
	}
	if n > 15 {
		n = 15
	}

	const gh = "0123456789bcdefghjkmnpqrstuvwxyz"
	const two45 = 1 << 45
	const lateps = 90.0 / two45
	const loneps = 180.0 / two45

	abs := func(b bool) uint64 {
		if b {
			return 1
		}
		return 0
	}

	lat, lon := point.Lat, point.Lon

	if lat == 90 {
		lat -= lateps / 2
	}
	if lon == 180 {
		lon = -180
	}

	ulat := uint64(math.Floor(lat/lateps) + two45)
	ulon := uint64(math.Floor(lon/loneps) + two45)

	var (
		b uint64
		B strings.Builder
	)
	for i := uint(0); i < 5*n; i++ {
		if i&1 == 0 {
			b = (b << 1) + abs((ulon&two45) != 0)
			ulon <<= 1
		} else {
			b = (b << 1) + abs((ulat&two45) != 0)
			ulat <<= 1
		}
		if (i+1)%5 == 0 {
			B.WriteByte(gh[int(b)])
			b = 0
		}
	}

	return B.String()
}
