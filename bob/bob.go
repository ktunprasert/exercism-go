// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

var reUpper = regexp.MustCompile(`^[A-Z?]+$`)
var reQuestion = regexp.MustCompile(`\?$`)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if len(strings.TrimSpace(remark)) == 0 {
		return "Fine. Be that way!"
	}

	remark = regexp.MustCompile(`[^A-Za-z?]+`).ReplaceAllString(remark, "")
	isUpper := remark != "?" && reUpper.MatchString(remark)
	isQuestion := reQuestion.MatchString(remark)

	switch {
	case isUpper && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isUpper:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
