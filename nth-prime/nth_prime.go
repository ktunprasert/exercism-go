package prime

import (
	"errors"
)

var primes []int = []int{2, 3, 5, 7, 11, 13}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("n must be greater than zero")
	}

	current := primes[len(primes)-1]
	for n >= len(primes) {
		isPrime := true
		for _, p := range primes {
			if current%p == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, current)
		}

		current++
	}

	return primes[n-1], nil
}
