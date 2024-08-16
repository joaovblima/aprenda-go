package main

import (
	"fmt"
)

func main() {
	for i := 33; i < 122; i++ {
		fmt.Printf("%v - %d - %#x - %#U\n", i, i, i, i)
	}
}