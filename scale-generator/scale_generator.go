package scale

import "strings"

const (
	Sharp = iota + 1
	Flat
)

var (
	sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flats  = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
)

func Scale(tonic, interval string) []string {
	signatureScale := getSignatureScale(tonic)

	i := 0
	for ; !strings.EqualFold(signatureScale[i], tonic); i++ {
	}

	if len(interval) == 0 {
		return append(signatureScale[i:], signatureScale[:i]...)
	}

	chromaticScale := []string{signatureScale[i]}
	for j := 0; j < len(interval); j++ {
		var scalar int
		switch interval[j] {
		case 'M':
			scalar = 2
		case 'A':
			scalar = 3
		default:
			scalar = 1
		}

		i = (i + scalar) % 12
		chromaticScale = append(chromaticScale, signatureScale[i])
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
