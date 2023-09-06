package resistorcolortrio

import "fmt"

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var metricPrefixes = map[int]string{
	0: "",
	1: "kilo",
	2: "mega",
	3: "giga",
}

func Label(colors []string) string {
	c1, c2, n := colors[0], colors[1], colors[2]

	numZeroes := colorMap[n]

	resistance := 10*colorMap[c1] + colorMap[c2]
	if colorMap[c2] == 0 {
		numZeroes += 1
		resistance /= 10
	}

	resistance *= pow(10, (numZeroes+1)%3)

	return fmt.Sprintf("%d %sohms", resistance, metricPrefixes[numZeroes/3])
}

func pow(x, n int) int {
	if n <= 1 {
		return 1
	}

	return x * pow(x, n-1)
}
