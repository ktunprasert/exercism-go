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
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	dayModifier := 1
	switch wSched {
	case Last:
		start = start.AddDate(0, 1, -1)
		dayModifier = -1
	case Teenth:
		start = start.AddDate(0, 0, 12)
	default:
		start = start.AddDate(0, 0, int(wSched)*7)
	}

	for start.Weekday() != wDay {
		start = start.AddDate(0, 0, dayModifier)
	}

	return start.Day()
}
