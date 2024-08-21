package main

import (
	"fmt"
)

func main() {
	x := 10

	fmt.Println(dobro(x))
}


func dobro(x int) int{
	return x * 2
}