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

var bottleOnTheWall = "%s green bottle%s hanging on the wall,"
var bottleShouldFall = "And if one green bottle should accidentally fall,"
var bottleLeft = "There'll be %s green bottle%s hanging on the wall."

func Recite(startBottles, takeDown int) []string {
	recital := make([]string, 0)

	for i := 0; i < takeDown; i++ {
		plural := ""
		if startBottles > 1 {
			plural = "s"
		}

		recital = append(recital, []string{
			fmt.Sprintf(bottleOnTheWall, bottleMap[startBottles], plural),
			fmt.Sprintf(bottleOnTheWall, bottleMap[startBottles], plural),
			bottleShouldFall,
		}...)

		if startBottles-1 == 1 {
			plural = ""
		} else {
			plural = "s"
		}

		recital = append(recital, fmt.Sprintf(bottleLeft, strings.ToLower(bottleMap[startBottles-1]), plural))

		if i != takeDown-1 {
			recital = append(recital, "")
		}

		startBottles--
	}

	return recital
}
