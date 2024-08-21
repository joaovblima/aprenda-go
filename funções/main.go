package main

import (
	"fmt"
)

func main() {
	sliceofints := []int{10, 20, 30,40}

	resultado := soma(sliceofints...)
	fmt.Println(resultado)
}

func soma (x...int) int {
	adiciona := 0
	for _, v := range x {
		adiciona+=v
	}
	return adiciona
}