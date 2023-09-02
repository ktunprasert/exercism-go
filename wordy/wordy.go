package wordy

import (
	"regexp"
	"strconv"
)

var (
	numbers    = regexp.MustCompile(`-?\d+`)
	operations = regexp.MustCompile(`plus|minus|multiplied by|divided by`)

	valid = regexp.MustCompile(`What is -?\d+( (plus|minus|multiplied by|divided by) -?\d+)*\?`)
)

func Answer(question string) (int, bool) {
	nums := numbers.FindAllString(question, -1)
	ops := operations.FindAllString(question, -1)

	if !valid.MatchString(question) {
		return -1, false
	}

	sum, _ := strconv.Atoi(nums[0])
	for i, op := range ops {
		num, _ := strconv.Atoi(nums[i+1])

		switch op {
		case "plus":
			sum += num
		case "minus":
			sum -= num
		case "multiplied by":
			sum *= num
		case "divided by":
			sum /= num
		}
	}

	return sum, true
}
