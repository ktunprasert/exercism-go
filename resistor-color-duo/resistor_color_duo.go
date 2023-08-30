package resistorcolorduo

var resistanceMap = map[string]int{
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

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) < 2 {
		return -1
	}

	var resistance int
	for i, fact := 0, 10; i < 2; i, fact = i+1, fact/10 {
		val, ok := resistanceMap[colors[i]]
		if !ok {
			return -1
		}

		resistance += fact * val
	}

	return resistance
}
