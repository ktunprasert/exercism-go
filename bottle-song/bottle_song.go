package bottlesong

import (
	"fmt"
	"strings"
)

var bottleMap = map[int]string{
	10: "Ten",
	9:  "Nine",
	8:  "Eight",
	7:  "Seven",
	6:  "Six",
	5:  "Five",
	4:  "Four",
	3:  "Three",
	2:  "Two",
	1:  "One",
	0:  "No",
}

var (
	bottleOnTheWall  = "%s green bottle%s hanging on the wall,"
	bottleShouldFall = "And if one green bottle should accidentally fall,"
	bottleLeft       = "There'll be %s green bottle%s hanging on the wall."
)

func Recite(startBottles, takeDown int) []string {
	recital := make([]string, 0)

	for i := 0; i < takeDown; i++ {
		recital = append(recital, []string{
			fmt.Sprintf(bottleOnTheWall, bottleMap[startBottles], plural(startBottles)),
			fmt.Sprintf(bottleOnTheWall, bottleMap[startBottles], plural(startBottles)),
			bottleShouldFall,
			fmt.Sprintf(bottleLeft, strings.ToLower(bottleMap[startBottles-1]), plural(startBottles-1)),
		}...)

		if i != takeDown-1 {
			recital = append(recital, "")
		}

		startBottles--
	}

	return recital
}

func plural(n int) string {
	if n != 1 {
		return "s"
	}

	return ""
}
