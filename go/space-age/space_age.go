package space

// Planet struct describing each planet in the solar system
type Planet struct {
	name          string
	orbitalperiod float64
}

var planetdata = []Planet{
	Planet{name: "Mercury", orbitalperiod: 0.2408467},
	Planet{name: "Venus", orbitalperiod: 0.61519726},
	Planet{name: "Earth", orbitalperiod: 1.0},
	Planet{name: "Mars", orbitalperiod: 1.8808158},
	Planet{name: "Jupiter", orbitalperiod: 11.862615},
	Planet{name: "Saturn", orbitalperiod: 29.447498},
	Planet{name: "Uranus", orbitalperiod: 84.016846},
	Planet{name: "Neptune", orbitalperiod: 164.79132},
}

func Age(seconds float64, planetname string) float64 {
	for _, planet := range planetdata {
		if planetname == planet.name {
			return seconds / planet.orbitalperiod * 31557600
		}
	}
}
