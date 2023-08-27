package darts

import "math"

func Score(x, y float64) int {
	dist := math.Sqrt(x*x + y*y)

	switch {
	case dist <= 1:
		return 10
	case dist <= 5:
		return 5
	case dist <= 10:
		return 1
	}

	return 0
}
