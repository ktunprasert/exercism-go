package anagram

import (
	"fmt"
	"strings"
)

func hashmapString(str string) map[string]int {
	m := make(map[string]int)
	for _, v := range str {
		m[string(v)]++
	}

	return m
}

func Detect(subject string, candidates []string) []string {
	subjectMap := hashmapString(strings.ToLower(subject))

	anagrams := make([]string, 0)
	for _, cand := range candidates {
		if strings.EqualFold(subject, cand) {
			continue
		}

		candMap := hashmapString(strings.ToLower(cand))

		if fmt.Sprintf("%v", subjectMap) == fmt.Sprintf("%v", candMap) {
			anagrams = append(anagrams, string(cand))
		}
	}

	return anagrams
}
