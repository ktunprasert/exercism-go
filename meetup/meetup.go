package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	Last   = -2
	Teenth = -1
	First  = 0
	Second = 1
	Third  = 2
	Fourth = 3
	Fifth  = 4
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	if wSched != Last {
		day := 1
		if wSched == Teenth {
			day = 13
		}

		start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		for start.Weekday() != wDay {
			start = start.AddDate(0, 0, 1)
		}

		schedule := 0
		for schedule < int(wSched) {
			start = start.AddDate(0, 0, 7)
			schedule++
		}

		return start.Day()
	}

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	start = start.AddDate(0, 1, -1)
	for start.Weekday() != wDay {
		start = start.AddDate(0, 0, -1)
	}

	return start.Day()
}
