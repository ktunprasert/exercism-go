package summultiples

func SumMultiples(limit int, divisors ...int) int {
	hash := make(map[int]struct{})

	for _, d := range divisors {
		if d == 0 {
			continue
		}

		for i := 1; d*i < limit; i++ {
			hash[d*i] = struct{}{}
		}
	}

	sum := 0
	for k := range hash {
		sum += k
	}

	return sum
}
