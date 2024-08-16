package main 

import (
	"fmt"
)

const (
	ANO = 2024
	a = ANO + iota
	b 
	c
)

func main() {
	fmt.Printf("%v, %v, %v", a, b, c)

}