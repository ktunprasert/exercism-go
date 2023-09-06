package perfect

import "errors"

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive error = errors.New("Only positive numbers are allowed")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}

	if n == 1 {
		return ClassificationDeficient, nil
	}

	sum := int64(1)
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	case sum < n:
		return ClassificationDeficient, nil
	}

	return -1, nil
}
