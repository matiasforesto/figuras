package figuras

import "fmt"

//Interface
type Geometrica interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometrica) {
	fmt.Println("Medida: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimetro: ", g.perimetro())
}
