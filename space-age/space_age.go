package space

type Planet string

const (
	MERCURY = 0.2408467
	VENUS   = 0.61519726
	EARTH   = 1.0
	MARS    = 1.8808158
	JUPITER = 11.862615
	SATURN  = 29.447498
	URANUS  = 84.016846
	NEPTUNE = 164.79132

	YEAR_IN_SECONDS = 60 * 60 * 24 * 365.25
)

var planetMap = map[Planet]float64{
	"Mercury": MERCURY,
	"Venus":   VENUS,
	"Earth":   EARTH,
	"Mars":    MARS,
	"Jupiter": JUPITER,
	"Saturn":  SATURN,
	"Uranus":  URANUS,
	"Neptune": NEPTUNE,
}

func Age(seconds float64, planet Planet) float64 {
	if _, ok := planetMap[planet]; !ok {
		return -1.0
	}

	return seconds / (YEAR_IN_SECONDS * planetMap[planet])
}
