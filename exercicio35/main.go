package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) mostraDados() {
	fmt.Println("Nome: ", p.nome)
	fmt.Println("Sobrenome: ", p.sobrenome)
	fmt.Println("Idade: ", p.idade)
}

func main() {

	pessoa1 := pessoa{
		nome: "Joao",
		sobrenome: "Lima",
		idade: 27,
	}

	pessoa1.mostraDados()
}