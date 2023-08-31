package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	encrypted := strings.Builder{}

	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) {
			encrypted.WriteRune('z' - (c - 'a'))
		} else if unicode.IsDigit(c) {
			encrypted.WriteRune(c)
		}

		if encrypted.Len()%6 == 5 {
			encrypted.WriteRune(' ')
		}
	}

	return strings.TrimSpace(encrypted.String())
}
