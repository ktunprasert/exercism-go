package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	normalised := regexp.MustCompile(`\w`).FindAllString(strings.ToLower(pt), -1)

	c_float := math.Ceil(math.Sqrt(float64(len(normalised))))
	r_float := math.Ceil(float64(len(normalised)) / c_float)

	c, r := int(c_float), int(r_float)

	cipher := strings.Builder{}

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if i+(j*c) >= len(normalised) {
				cipher.WriteString(" ")
				continue
			}

			cipher.WriteString(normalised[i+(j*c)])
		}

		if i != c-1 {
			cipher.WriteString(" ")
		}
	}

	return cipher.String()
}
