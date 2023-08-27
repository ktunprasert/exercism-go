package isogram

import "unicode"

func IsIsogram(word string) bool {
	count := make(map[rune]struct{})
	for _, r := range word {
		if unicode.IsLetter(r) {
			if _, ok := count[unicode.ToLower(r)]; ok {
				return false
			}

			count[unicode.ToLower(r)] = struct{}{}
		}
	}
	return true
}
