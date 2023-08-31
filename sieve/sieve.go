package sieve

func Sieve(limit int) []int {
	sieve := make(map[int]struct{})
	for i := 2; i <= limit; i++ {
		sieve[i] = struct{}{}
	}

	primes := make([]int, 0)
	for i := 2; i <= limit; i++ {
		if _, ok := sieve[i]; !ok {
			continue
		}

		primes = append(primes, i)

		for j := i + i; j <= limit; j += i {
			delete(sieve, j)
		}
	}

	return primes
}
