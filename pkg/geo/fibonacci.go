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
