package main

import (
	"fmt"
)

func main() {
	x:= 1
	y:= 2
	z:=3

	defer fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}