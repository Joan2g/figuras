package figuras

import (
	"math"
)

type Circulo struct {
	Radio int
}

func (cir *Circulo) sacarPerimetro() float64 {
	pi := math.Pi
	return 2 * pi * float64(cir.Radio)
}

func (cir *Circulo) sacarArea() float64 {
	pi := math.Pi
	return pi * (float64(cir.Radio) * float64(cir.Radio))
}
