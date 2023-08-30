package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	sortedSubject := []rune(strings.ToLower(subject))
	slices.Sort(sortedSubject)

	anagrams := make([]string, 0)
	for _, cand := range candidates {
		if strings.EqualFold(subject, cand) {
			continue
		}

		sortedCand := []rune(strings.ToLower(cand))
		slices.Sort(sortedCand)

		if strings.EqualFold(string(sortedSubject), string(sortedCand)) {
			anagrams = append(anagrams, string(cand))
		}
	}

	return anagrams
}
