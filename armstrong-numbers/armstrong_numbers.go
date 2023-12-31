package armstrong

import "math"

func IsNumber(n int) bool {
	if n < 10 {
		return true
	}

	digits := math.Floor(math.Log10(float64(n)) + 1)

	cpy := n
	sum := 0
	for ; cpy > 0; cpy /= 10 {
		remainder := cpy % 10
		sum += int(math.Pow(float64(remainder), digits))
	}

	return n == sum
}
