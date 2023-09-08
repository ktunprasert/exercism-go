package yacht

func Score(dice []int, category string) int {
	if len(dice) != 5 {
		return -1
	}

	diceCount := make(map[int]int)
	for _, d := range dice {
		diceCount[d]++
	}

	switch {
	case category == "ones":
		return diceCount[1]
	case category == "twos":
		return diceCount[2] * 2
	case category == "threes":
		return diceCount[3] * 3
	case category == "fours":
		return diceCount[4] * 4
	case category == "fives":
		return diceCount[5] * 5
	case category == "sixes":
		return diceCount[6] * 6
	case category == "full house" && len(diceCount) == 2:
		var condThree, condTwo bool
		sum := 0
		for k, v := range diceCount {
			if v == 3 {
				condThree = true
				sum += k * v
			}

			if v == 2 {
				condTwo = true
				sum += k * v
			}
		}

		if condThree && condTwo {
			return sum
		}
	case category == "little straight" &&
		diceCount[1] > 0 &&
		diceCount[2] > 0 &&
		diceCount[3] > 0 &&
		diceCount[4] > 0 &&
		diceCount[5] > 0:
		return 30
	case category == "big straight" &&
		diceCount[2] > 0 &&
		diceCount[3] > 0 &&
		diceCount[4] > 0 &&
		diceCount[5] > 0 &&
		diceCount[6] > 0:
		return 30
	case category == "choice":
		sum := 0
		for k, v := range diceCount {
			sum += k * v
		}
		return sum
	case len(diceCount) == 1:
		if category == "yacht" {
			return 50
		}
		if category == "four of a kind" {
			return dice[0] * 4
		}
	case len(diceCount) < 3:
		sum := 0
		for k, v := range diceCount {
			if v == 4 {
				sum += k * v
			}
		}
		return sum
	}

	return 0
}
