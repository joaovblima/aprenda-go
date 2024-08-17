package main 

import (
	"fmt"
)

func main() {
	
	//sabores := []string {"calabresa", "quatro queijos", "nordestina", "frango"}
	//fatia := sabores[2:]

	//fmt.Println(sabores)
	//fmt.Println(fatia)


	sliceUm := []int {1, 2, 3, 4, 5}
	sliceDois := []int {6, 7, 8, 9, 10}
	terceiroSlice := append(sliceUm)
	fmt.Println(sliceUm)
	fmt.Println(sliceDois)
	fmt.Println(terceiroSlice)
	terceiroSlice = append(terceiroSlice, sliceDois...)
	fmt.Println(terceiroSlice)
}
