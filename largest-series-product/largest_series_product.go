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

	var maxProduct int64 = 0

	for i := 0; i <= len(digits)-span; i++ {
		seq := digits[i : i+span]
		var product int64 = 1
		for _, n := range seq {
			product *= int64(n - '0')
		}

		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}
