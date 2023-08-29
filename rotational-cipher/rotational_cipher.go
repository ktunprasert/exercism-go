package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	var out string

	for _, c := range plain {
		if !unicode.IsLetter(c) {
            out += string(c)
			continue
		}

		baseRune := 'a'
		if unicode.IsUpper(c) {
			baseRune = 'A'
		}

		value := int(c - baseRune)
		value = (value + shiftKey) % 26

		out += string(rune(value + int(baseRune)))
	}

	return out
}
