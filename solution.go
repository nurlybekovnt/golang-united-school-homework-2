package square

import "math"

// Sides represents two-dimensional shape.
type Sides int

// Two-dimensional shapes.
const (
	SidesCircle Sides = iota
	_
	_
	SidesTriangle
	SidesSquare
)

// CalcSquare calculates area of three-dimensional object. sideLen is the side
// length of regular polygon. In the case of cicrle it is the radius length. For
// unsupported shapes it returns 0.
func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	if sideLen <= 0 {
		return 0
	}
	switch sidesNum {
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	case SidesTriangle:
		return sideLen * sideLen / 4 * math.Sqrt(3)
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}
}
