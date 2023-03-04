package figuras

type Cuadrado struct {
	Lado int
}

func (cua *Cuadrado) sacarPerimetro() float64 {
	return float64(4 * cua.Lado)
}

func (cua *Cuadrado) sacarArea() float64 {
	return float64(cua.Lado * cua.Lado)
}
