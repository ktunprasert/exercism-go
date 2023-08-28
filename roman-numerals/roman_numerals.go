package romannumerals

import (
	"errors"
	"strings"
)

var numerals = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("input out of range")
	}

	var out string
	var factor int = 1
	for ; input > 0; input /= 10 {
		remainder := input % 10

		switch {
		case remainder < 4:
			out = strings.Repeat(numerals[factor], remainder) + out
		case remainder == 4:
			out = numerals[factor] + numerals[factor*5] + out
		case remainder < 9:
			out = numerals[factor*5] + strings.Repeat(numerals[factor], remainder-5) + out
		case remainder == 9:
			out = numerals[factor] + numerals[factor*10] + out
		}

		factor *= 10
	}

	return out, nil
}
