package main

import (
	"fmt"
)

func main() {
	t := []int{10, 11, 12, 13, 14}
	b := somentePares(soma, []int {50, 51, 52, 53, 54, 55}...)

	fmt.Println(soma(t...))

	fmt.Println(somaPares(t...))

	fmt.Println(somaImpares(t...))

	fmt.Println(somentePares(soma, t...))
	fmt.Println(b)

	fmt.Println(somenteImpares(soma, t...))

}

func soma (x... int) int {
	n:=0
	for _, v := range x {
		n += v
	}
	return n
}

func somaPares(x...int) int {
	n := 0
	for _, v := range x {
		if v % 2 ==0 {
			n+= v
		}

	}
	return n
}

func somaImpares(x... int) int {
	n := 0 
	for _, v := range x {
		if v % 2 ==1 {
			n +=v
		}
	}

	return n
}

func somentePares (f func (x ...int) int, y ... int) int{
	var slice []int
	for _, v := range y {
		if v % 2 == 0{
			slice = append(slice, v)
		}
				
	}
	total := f(slice...)
	return total
}

func somenteImpares (f func (x ... int) int, y... int) int {
	var slice [] int
	for _, v := range y {
		if v % 2!=0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}