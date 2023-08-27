package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	var sum int
	secondDigit := false

	for i := len(id) - 1; i >= 0; i-- {
		switch r := rune(id[i]); {
		case !unicode.IsDigit(r):
			return false
		case secondDigit:
			double := int(r-'0') * 2
			if double > 9 {
				double -= 9
			}
			sum += double
		default:
			sum += int(r - '0')
		}

		secondDigit = !secondDigit
	}

	return sum%10 == 0
}
