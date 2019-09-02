package space

type Planet string

const Eop = 31557600

var planets = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(sec float64, planet Planet) float64 {
	planetOp := planets[planet]
	return sec / (planetOp * float64(Eop))
}
