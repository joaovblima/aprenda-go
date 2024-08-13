package main

import (
	"fmt"
)

type novenove int
var x novenove

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}