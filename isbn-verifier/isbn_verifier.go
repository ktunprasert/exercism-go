package isbn

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {
	if regexp.MustCompile(`[^0-9X-]`).MatchString(isbn) {
		return false
	}

	isbn = strings.Replace(isbn, "-", "", -1)

	if len(isbn) != 10 {
		return false
	}

	var sum int
	for i, c := range isbn {
		factor := 10 - i

		switch {
		case c == 'X' && i == 9:
			sum += 10 * factor
		case c == 'X' && i != 9:
			return false
		default:
			sum += int(c-'0') * factor
		}
	}

	return sum%11 == 0
}
