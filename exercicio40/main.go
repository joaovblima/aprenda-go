package main

import (
	"fmt"
)

func main() {
	fmt.Println(dobraValor()(10))
	

}

func dobraValor () func (int) int {
	return func(x int) int {
		return x * 2
	}
}