package main

import(
	"fmt"
)

/*
Crie uma função que receba um parametro variadico de tipo int e retorne a soma de todos os ints
recebidos: [x]
Passe um valor do tipo slice of int como argumento da funcao[]

Crie outra funcao, esta deve receber o valor do tipo slice of int e retornar a soma de todos os 
elementos do slice. []

Passe um valor do tipo slice of int como arguemnto para a funcao.[]


*/
func main() {

	valor1 := []int{12, 14, 5, 19, 32, 94, 15}
	valor2 := []int{17, 54, 123, 452, 512}

	fmt.Println(somaNumeros(valor1...))
	fmt.Println(somaDeNumeros(valor2))
}

func somaNumeros(x ...int) int {
	s :=0
	for _, v := range x {
		s +=v
	}

	return s
}

 
func somaDeNumeros(x []int) int {
	s :=0
	for _, v := range x {
		s +=v
	}

	return s
}