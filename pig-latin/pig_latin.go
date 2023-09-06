package piglatin

import (
	"regexp"
	"strings"
)

var (
	beginVowelRule = regexp.MustCompile(`^([aeiou]|yt|xr)(.*)$`)
	consonantRule  = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]+)(.*)$`)
	yRule          = regexp.MustCompile(`^([^aeiou]+)(y.*)$`)
)

func Sentence(sentence string) string {
	out := strings.Builder{}

	for _, word := range strings.Fields(sentence) {
		out.WriteString(pigLatinise(word) + " ")
	}

	return out.String()[:out.Len()-1]
}

func pigLatinise(word string) string {
	switch {
	case beginVowelRule.MatchString(word):
		return word + "ay"

	case yRule.MatchString(word):
		parts := yRule.FindStringSubmatch(word)
		return parts[2] + parts[1] + "ay"

	case consonantRule.MatchString(word):
		parts := consonantRule.FindStringSubmatch(word)
		return parts[2] + parts[1] + "ay"
	}

	return word
}
