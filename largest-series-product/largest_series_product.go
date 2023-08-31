package lsproduct

import (
	"errors"
	"regexp"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	switch {
	case !regexp.MustCompile(`^\d+$`).MatchString(digits):
		return 0, errors.New("input must only contain digits")
	case span > len(digits):
		return 0, errors.New("span must be smaller than string length")
	case span < 0:
		return 0, errors.New("span must be greater than zero")
	}

	maxProduct := 0

	for i := 0; i <= len(digits)-span; i++ {
		seq := digits[i : i+span]
		product := 1
		for _, n := range seq {
			product *= int(n - '0')
		}

		maxProduct = max(maxProduct, product)
	}

	return int64(maxProduct), nil
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
