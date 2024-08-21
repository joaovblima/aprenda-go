package main

import(
	"fmt"
)

func main() {

	soma := func (x, y int) {
		fmt.Println(x + y)
	} 
	soma(2, 4)

	multi := func (x, y int) {
		fmt.Println(x * y)
	}

	multi(10, 50)

	

}