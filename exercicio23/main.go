package main 

import (
	"fmt"
)

func main() {
	x:= []int{41, 42, 43, 44, 45, 46, 47, 48, 49 ,50, 51}
	y:= []int{}
	y = append(y, x...)

	fmt.Println(y)
}