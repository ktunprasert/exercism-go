package prime

func Factors(n int64) []int64 {
	factors := make([]int64, 0)

	divisor := int64(2)
	for n != 1 {
		for ; n%divisor == 0; n /= divisor {
			factors = append(factors, divisor)
		}

		divisor++
	}

	return factors
}
