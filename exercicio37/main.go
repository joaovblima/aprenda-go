package main

import(
	"fmt"
	"math"
)

/*
Criar dois tipos: quadrado e circulo
criar um metodo area para cada tipo que calcule a area da figura.

criar um tipo figura que defina a interface qualquer tipo que tiver o metodo area

criar funcao info que recebe um tipo figura e retorna a area da figura

*/

type quadrado struct {
	lado float64

}

type circulo struct {
	raio float64
}

func (q quadrado) area() {
	result := q.lado * 2
	fmt.Println("A area do quadrado é: ", result)
}

func (c circulo) area() {
	result := math.Pi * (c.raio * c.raio)
	fmt.Println("A area do circulo é: ", result)
}

type figura interface {
	area()

}

func medida(f figura) {
	f.area()
}
func main() {

	q1 := quadrado{
		lado: 4.1,
	}

	c1 := circulo{
		raio: 15,
	}

	medida(q1)
	medida(c1)
	

}