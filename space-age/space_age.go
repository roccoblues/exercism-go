package space

// Planet is the name of a planet.
type Planet string

const earthYearSeconds = 31557600

var orbitalPeriods = map[Planet]float64{
	"Earth":   earthYearSeconds,
	"Mercury": earthYearSeconds * 0.2408467,
	"Venus":   earthYearSeconds * 0.61519726,
	"Mars":    earthYearSeconds * 1.8808158,
	"Jupiter": earthYearSeconds * 11.862615,
	"Saturn":  earthYearSeconds * 29.447498,
	"Uranus":  earthYearSeconds * 84.016846,
	"Neptune": earthYearSeconds * 164.79132,
}

// Age calculates the age of someone on the given Planet.
func Age(seconds float64, planet Planet) float64 {
	return seconds / orbitalPeriods[planet]
}
