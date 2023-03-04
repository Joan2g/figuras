package figuras

import (
	"fmt"
)

type Geometria interface {
	sacarArea() float64
	sacarPerimetro() float64
}

func Medidas(g Geometria) {
	fmt.Println("Medidas:", g)
	fmt.Println("Area:", g.sacarArea())
	fmt.Println("Perimetro:", g.sacarPerimetro())
}
