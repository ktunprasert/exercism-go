// Package weather provides current weather condition for a given city.
package weather

// CurrentCondition: Describes current weather condition for a given city.
var CurrentCondition string
// CurrentLocation: Describes current location for a given city.
var CurrentLocation string

// Forecast returns a string describing current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
