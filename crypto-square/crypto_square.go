package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	normalised := regexp.MustCompile(`\w`).FindAllString(pt, -1)

	c_float := math.Ceil(math.Sqrt(float64(len(normalised))))
	r_float := math.Ceil(float64(len(normalised)) / c_float)

	c, r := int(c_float), int(r_float)

	rectangle := make([][]string, 0)
	for i := 0; i < r; i++ {
		if i == r-1 {
			padding := c - len(normalised[i*c:])
			rectangle = append(rectangle, normalised[i*c:])
			for j := 0; j < padding; j++ {
				rectangle[i] = append(rectangle[i], " ")
			}
			continue
		}
		rectangle = append(rectangle, normalised[i*c:(i+1)*c])
	}

	transpose := make([][]string, c)
	for i := 0; i < len(rectangle); i++ {
		for j := 0; j < len(rectangle[i]); j++ {
			transpose[j] = append(transpose[j], strings.ToLower(rectangle[i][j]))
		}
	}

	cipher := strings.Builder{}
	for i, c := range transpose {
		for _, r := range c {
			if r == "" {
				cipher.WriteString(" ")
				continue
			}
			cipher.WriteString(r)
		}
		if i != len(transpose)-1 {
			cipher.WriteString(" ")
		}
	}

	return cipher.String()
}
