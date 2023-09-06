package scale

import "strings"

var (
	sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flats  = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

    intervals = map[rune]int{'m': 1, 'M': 2, 'A': 3}
)

func Scale(tonic, interval string) []string {
	notes := getSignatureScale(tonic)

	i := 0
	for ; !strings.EqualFold(notes[i], tonic); i++ {
	}

	if len(interval) == 0 {
		return append(notes[i:], notes[:i]...)
	}

	chromaticScale := []string{notes[i]}
	for j := 0; j < len(interval); j++ {
		var scalar int = intervals[rune(interval[j])]

		i = (i + scalar) % 12
		chromaticScale = append(chromaticScale, notes[i])
	}

	return chromaticScale
}

func getSignatureScale(s string) []string {
	switch s {
	case "C", "a":
		return sharps
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		return sharps
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return flats
	}

	return sharps
}
