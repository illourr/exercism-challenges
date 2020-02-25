package space

type Planet string

// Mercury orbital period 0.2408467 Earth years
const Mercury = 0.2408467

// Venus orbital period 0.61519726 Earth years
const Venus = 0.61519726

// EarthSeconds orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
const EarthSeconds = 31557600.0

// Mars orbital period 1.8808158 Earth years
const Mars = 1.8808158

// Jupiter orbital period 11.862615 Earth years
const Jupiter = 11.862615

// Saturn orbital period 29.447498 Earth years
const Saturn = 29.447498

// Uranus orbital period 84.016846 Earth years
const Uranus = 84.016846

// Neptune orbital period 164.79132 Earth years
const Neptune = 164.79132

// Age calculates how old someone would be on another planet.
func Age(seconds float64, planet Planet) float64 {
	var years float64

	switch planet {
	case "Mercury":
		years = seconds / (Mercury * EarthSeconds)
	case "Venus":
		years = seconds / (Venus * EarthSeconds)
	case "Earth":
		years = seconds / EarthSeconds
	case "Mars":
		years = seconds / (Mars * EarthSeconds)
	case "Jupiter":
		years = seconds / (Jupiter * EarthSeconds)
	case "Saturn":
		years = seconds / (Saturn * EarthSeconds)
	case "Uranus":
		years = seconds / (Uranus * EarthSeconds)
	case "Neptune":
		years = seconds / (Neptune * EarthSeconds)
	}

	return years
}
