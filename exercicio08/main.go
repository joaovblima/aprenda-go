package main 

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Printf("decimal: %d, binario: %b, hexadecimal: %#x\n", a, a, a)
	b := a << 1
	fmt.Printf("decimal: %d, binario: %b, hexadecimal: %#x", b, b, b)



}