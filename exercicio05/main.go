package main

import (
	"fmt"
)

type novenove int
var x novenove
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println("####################")
	
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)





}