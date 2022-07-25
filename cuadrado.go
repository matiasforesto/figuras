package figuras

//Struct Cuadrado
type Cuadrado struct {
	Lado float64
}

func (cua *Cuadrado) area() float64 {
	return cua.Lado * cua.Lado // lado por lado
}

func (cua *Cuadrado) perimetro() float64 {
	return 4 * cua.Lado // suma de sus 4 lados
}
