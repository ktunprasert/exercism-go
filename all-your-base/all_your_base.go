package allyourbase

import "errors"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 || outputBase < 2 {
		return nil, errors.New("base must be >= 2")
	}

	base10 := 0
	for i, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}

		base10 += d * pow(inputBase, len(inputDigits)-i-1)
	}

	outputDigits := make([]int, 0)
	for p := nearestPower(base10, outputBase); p >= 0; p-- {
		d := base10 / pow(outputBase, p)
		outputDigits = append(outputDigits, d)

		base10 -= d * pow(outputBase, p)
	}

	return outputDigits, nil
}

func nearestPower(x, base int) int {
	if x < base {
		return 0
	}

	return 1 + nearestPower(x/base, base)
}

func pow(x, p int) int {
	if p == 0 {
		return 1
	}

	return x * pow(x, p-1)
}
