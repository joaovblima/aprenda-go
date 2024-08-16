package main 

import (
	"fmt"
)

const primeira = 10
const segunda float64= 11.0

func main() {
	a := 10 == 5
	b := 50 != 8
	c := 100 <= 101
	d := 90 >= 100
	e := 90 > 100
	f := 100 < 90

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(" ")
	fmt.Printf("%v, %T\n", primeira, primeira)

	fmt.Printf("%v, %T", segunda, segunda)


}

