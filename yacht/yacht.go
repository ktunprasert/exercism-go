package yacht

func Score(dice []int, category string) int {
	if len(dice) != 5 {
		return -1
	}

	diceCount := make(map[int]int)
	for _, d := range dice {
		diceCount[d]++
	}

	switch category {
	case "ones":
		return diceCount[1]
	case "twos":
		return diceCount[2] * 2
	case "threes":
		return diceCount[3] * 3
	case "fours":
		return diceCount[4] * 4
	case "fives":
		return diceCount[5] * 5
	case "sixes":
		return diceCount[6] * 6
	case "full house":
		sum := 0
		for k, v := range diceCount {
			if v != 2 && v != 3 {
				return 0
			}

			sum += k * v
		}

		return sum
	case "little straight":
		if len(diceCount) == 5 && diceCount[6] == 0 {
			return 30
		}
	case "big straight":
		if len(diceCount) == 5 && diceCount[1] == 0 {
			return 30
		}
	case "choice":
		sum := 0
		for k, v := range diceCount {
			sum += k * v
		}

		return sum
	case "yacht":
		if diceCount[dice[0]] != 5 {
			return 0
		}
		return 50
	case "four of a kind":
		for k, v := range diceCount {
			if v >= 4 {
				return 4 * k
			}
		}
	}

	return 0
}
