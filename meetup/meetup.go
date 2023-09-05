package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First = iota * 7
	Second
	Third
	Fourth
	Fifth

	Teenth = 12
	Last   = -1
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	dayModifier := 1

	switch wSched {
	case Last:
		start = start.AddDate(0, 1, -1)
		dayModifier = -1
	default:
		start = start.AddDate(0, 0, int(wSched))
	}

	for start.Weekday() != wDay {
		start = start.AddDate(0, 0, dayModifier)
	}

	return start.Day()
}
