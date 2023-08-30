package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subjectArr := strings.Split(strings.ToLower(subject), "")
	sort.Strings(subjectArr)
	sortedSubject := strings.Join(subjectArr, "")

	anagrams := make([]string, 0)
	for _, cand := range candidates {
		if strings.EqualFold(subject, cand) {
			continue
		}

		candArr := strings.Split(strings.ToLower(cand), "")
		sort.Strings(candArr)
		sortedCand := strings.Join(candArr, "")

		if strings.EqualFold(sortedSubject, string(sortedCand)) {
			anagrams = append(anagrams, string(cand))
		}
	}

	return anagrams
}
