package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	h := make(map[rune]struct{})
	for _, char := range strings.ToLower(input) {
		if char >= 'a' && char <= 'z' {
			h[char] = struct{}{}
		}
	}

	return len(h) == 26
}
