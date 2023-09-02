package wordy

import (
	"regexp"
	"strconv"
)

var (
	numbers    = regexp.MustCompile(`-?\d+`)
	operations = regexp.MustCompile(`plus|minus|multiplied by|divided by`)

    valid = regexp.MustCompile(`(-?\d+( (plus|minus|multiplied by|divided by) -?\d+){1,})|is -?\d+\?`)
)

func Answer(question string) (int, bool) {
    nums := numbers.FindAllString(question, -1)
    ops := operations.FindAllString(question, -1)

    if !valid.MatchString(question) {
        return -1, false
    }

	if len(nums) == 0 || len(ops) >= len(nums) || len(nums)-1 != len(ops) {
		return -1, false
	}

	for i, op := range ops {
        n1, _ := strconv.Atoi(nums[i])
        n2, _ := strconv.Atoi(nums[i+1])

		switch op {
		case "plus":
			nums[i+1] = strconv.Itoa(n1 + n2)
		case "minus":
			nums[i+1] = strconv.Itoa(n1 - n2)
		case "multiplied by":
			nums[i+1] = strconv.Itoa(n1 * n2)
		case "divided by":
			nums[i+1] = strconv.Itoa(n1 / n2)
		}
	}

	answer, _ := strconv.Atoi(nums[len(nums)-1])

	return answer, true
}
