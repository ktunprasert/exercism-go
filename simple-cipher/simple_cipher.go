package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	distance int
}
type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}

	return shift{distance}
}

func (c shift) Encode(input string) string {
	encoded := strings.Builder{}
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			v := (int(r-'a') + c.distance + 26) % 26
			encoded.WriteRune(rune(v) + 'a')
		}
	}

	return encoded.String()
}

func (c shift) Decode(input string) string {
	decoded := strings.Builder{}
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			v := (int(r-'a') - c.distance + 26) % 26
			decoded.WriteRune(rune(v) + 'a')
		}
	}

	return decoded.String()
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	if strings.ContainsFunc(key, func(r rune) bool { return !unicode.IsLetter(r) || unicode.IsUpper(r) }) {
		return nil
	}

	if regexp.MustCompile(`^a+$`).MatchString(key) {
		return nil
	}

	return vigenere{strings.ToLower(key)}
}

func (v vigenere) Encode(input string) string {
	encoded := strings.Builder{}
	i := 0
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			distance := int(v.key[i%len(v.key)] - 'a')
			val := (int(r-'a') + distance + 26) % 26
			encoded.WriteRune(rune(val) + 'a')
			i += 1
		}
	}

	return encoded.String()
}

func (v vigenere) Decode(input string) string {
	decoded := strings.Builder{}
	i := 0
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			distance := int(v.key[i%len(v.key)] - 'a')
			val := (int(r-'a') - distance + 26) % 26
			decoded.WriteRune(rune(val) + 'a')
			i += 1
		}
	}

	return decoded.String()
}
