package main 

import (
	"fmt"
)

func main() {

	anoQueNasci := 1996
	anoCorrente := 2024

	for {
		if anoQueNasci == anoCorrente {
			break
		}
		fmt.Println(anoQueNasci)
		anoQueNasci++


	}
}