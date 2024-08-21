package main

import (
	"fmt"
)

func main() {
	funcaoQueRecebe(funcaoQueLate)

}


func funcaoQueLate() {
	fmt.Println("au au au au!")
}

func funcaoQueRecebe (x func ()) {
	fmt.Println("la vai")
	x()
}

