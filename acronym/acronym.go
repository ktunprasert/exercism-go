// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

var wordSeparator = regexp.MustCompile(`[\s-]`)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abbr := strings.Builder{}

	for _, str := range wordSeparator.Split(s, -1) {
		if len(str) == 0 {
			continue
		}

		i := 0
		for !unicode.IsLetter(rune(str[i])) {
			i++
		}

		abbr.WriteRune(unicode.ToUpper(rune(str[i])))
	}

	return abbr.String()
}
