package main 

import (
	"fmt"
)

func main() {

	slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//primeiro ao terceiro
	fmt.Println(slice[:3])
	//quinto ao ultimo
	fmt.Println(slice[4:])
	//do segundo ao setimo
	fmt.Println(slice[1:7])
	//do terceiro ao penultimo
	fmt.Println(slice[2:9])


	//utilizando len
	penultimo := len(slice) -1
	fmt.Println(slice[2:penultimo])

}