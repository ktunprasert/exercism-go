package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transformed := make(map[string]int)

	for score, group := range in {
		for _, letter := range group {
			transformed[strings.ToLower(letter)] = score
		}
	}

	return transformed
}
