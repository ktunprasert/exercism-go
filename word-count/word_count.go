package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordcount := make(Frequency)
	s := strings.Builder{}

	for i, c := range strings.ToLower(phrase) {
		switch {
		case unicode.IsLetter(c), unicode.IsDigit(c):
			s.WriteRune(c)
		case c == '\'' && i > 0 && i < len(phrase)-1:
			if unicode.IsLetter(rune(phrase[i-1])) && unicode.IsLetter(rune(phrase[i+1])) {
				s.WriteRune(c)
			}
		default:
			if s.Len() > 0 {
				wordcount[s.String()]++
			}
			s.Reset()
		}
	}

	if s.Len() > 0 {
		wordcount[s.String()]++
	}

	return wordcount
}
