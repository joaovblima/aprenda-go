package main

import (
	"fmt"
)

func main() {

	fmt.Println(retornaOTriplo()(10))
	
}

func retornaOTriplo() func (int) int {
	return func (x int) int {
		return x * 3
	}
}