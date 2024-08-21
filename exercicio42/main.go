package main

import (
	"fmt"
)

func main() {
	o := plusplus()
	fmt.Println(o())
	fmt.Println(o())
	fmt.Println(o())
}

func plusplus() func () int {
	x := 0
	return func() int {
		x++
		return x
	} 
}