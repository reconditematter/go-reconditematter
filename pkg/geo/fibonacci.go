package geo

import (
	"math"
)

// FibonacciPoints returns a quasi-uniform sequence
// of 2*n+1 geographic points placed on a Fibonacci grid.
//
// Reference: Swinbank R., Purser R.J., Fibonacci grids: A novel approach
// to global modelling, Q.J.R. Meteorol. Soc., vol.132, no.619, pp.1769-1793 (2006).
func FibonacciPoints(n uint) []Point {
	n21 := float64(n*2 + 1)
	p := make([]Point, 2*n+1)
	for i := -int(n); i <= int(n); i++ {
		fi := float64(i)
		sinφ := 2 * fi / n21
		cosφ := math.Sqrt((1 - sinφ) * (1 + sinφ))
		φ := math.Atan2(sinφ, cosφ)
		λ := (2 * math.Pi / math.Phi) * fi
		sinλ, cosλ := math.Sincos(λ)
		λ = math.Atan2(sinλ, cosλ)
		//
		p[i+int(n)] = Point{Lat: φ * (180 / math.Pi), Lon: λ * (180 / math.Pi)}
	}
	return p
}

// FibonacciCell returns a quasi-uniform sequence
// of _approximately_ n geographic points placed
// on a Fibonacci grid in the geographic cell [lat,lat+1]x[lon,lon+1].
//
// Precondition: -90<=latmin<90 and -180<=lonmin<180;
// otherwise a runtime panic occurs.
func FibonacciCell(n uint, latmin, lonmin int) []Point {
	if !(-90 <= latmin && latmin < 90) {
		panic("FibonacciCell: invalid latmin")
	}
	if !(-180 <= lonmin && lonmin < 180) {
		panic("FibonacciCell: invalid lonmin")
	}
	// compute 1x1 cell area
	const S = 4 * math.Pi
	cellA := (math.Pi / 180) * math.Abs(math.Sin(float64(latmin+1)*(math.Pi/180))-math.Sin(float64(latmin)*(math.Pi/180)))
	nn := int(math.Ceil(float64(n) * S / cellA))
	m := nn / 2
	m21 := float64(m*2 + 1)
	p := make([]Point, 0, n)

	for i := -m; i <= m; i++ {
		fi := float64(i)
		sinφ := 2 * fi / m21
		cosφ := math.Sqrt((1 - sinφ) * (1 + sinφ))
		φdeg := math.Atan2(sinφ, cosφ) * (180 / math.Pi)
		if !(float64(latmin) < φdeg && φdeg < float64(latmin+1)) {
			continue
		}
		λ := (2 * math.Pi / math.Phi) * fi
		sinλ, cosλ := math.Sincos(λ)
		λdeg := math.Atan2(sinλ, cosλ) * (180 / math.Pi)
		if !(float64(lonmin) < λdeg && λdeg < float64(lonmin+1)) {
			continue
		}

		p = append(p, Point{Lat: φdeg, Lon: λdeg})
	}
	return p
}
