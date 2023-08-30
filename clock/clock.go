package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	var minutes int = h*60 + m

	if minutes < 0 {
		minutes = 1440 + minutes%1440
	}

	return Clock{hours: minutes / 60 % 24, minutes: minutes % 60}
}

func (c Clock) Add(m int) Clock {
	minutes := c.hours*60 + c.minutes + m

	if minutes < 0 {
		minutes = 1440 + minutes%1440
	}

	return Clock{hours: minutes / 60 % 24, minutes: minutes % 60}
}

func (c Clock) Subtract(m int) Clock {
	minutes := c.hours*60 + c.minutes - m

	if minutes < 0 {
		minutes = 1440 + minutes%1440
	}

	return Clock{hours: minutes / 60 % 24, minutes: minutes % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
