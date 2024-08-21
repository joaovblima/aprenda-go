package main 

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int
}

func main() {
	zezinho := pessoa{
		nome: "fi de fiote",
		idade: 21,
	}

	fmt.Println(zezinho.saudacao())

}

func (p pessoa) saudacao() string {
	return "pauuuuu"
}