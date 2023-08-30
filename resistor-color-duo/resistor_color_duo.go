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
	for i := 0; i < 2; i++ {

		val, ok := resistanceMap[colors[i]]
		if !ok {
			return -1
		}

		resistance += power(10, 1-i) * val
	}

	return resistance
}

func power[T int](value T, num int) T {
	if num <= 0 {
		return 1
	}

	return value * power(value, num-1)
}
