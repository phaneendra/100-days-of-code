package space

//Planet type string containing name of the planet
type Planet string

// Planets type struct describing each planet in the solar system
type Planets struct {
	name          Planet
	orbitalperiod float64
}

var planetdata = []Planets{
	Planets{name: "Mercury", orbitalperiod: 0.2408467},
	Planets{name: "Venus", orbitalperiod: 0.61519726},
	Planets{name: "Earth", orbitalperiod: 1.0},
	Planets{name: "Mars", orbitalperiod: 1.8808158},
	Planets{name: "Jupiter", orbitalperiod: 11.862615},
	Planets{name: "Saturn", orbitalperiod: 29.447498},
	Planets{name: "Uranus", orbitalperiod: 84.016846},
	Planets{name: "Neptune", orbitalperiod: 164.79132},
}

// Age function computes the age of a person on a planet in the solar system
func Age(seconds float64, planetname Planet) float64 {
	for _, planet := range planetdata {
		if planetname == planet.name {
			return seconds / (planet.orbitalperiod * 31557600)
		}
	}
	return 0
}
