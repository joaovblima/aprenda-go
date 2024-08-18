package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	sabores []string
}

func main() {

	pessoa1 := pessoa{
		nome:"Joao",
		sobrenome: "Lima",
		sabores: []string{"Napolitano", "Nutella", "Pistache"},}

	pessoa2 := pessoa{
		nome: "Maria",
		sobrenome: "Sofia",
		sabores: []string{"Nutella", "Ninho Com Oreo", "Diamante Negro"},
	}
	

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for i, v := range pessoa1.sabores{
		fmt.Println(i, "-", v)
	}
	fmt.Println(" ")

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for i , v := range pessoa2.sabores{
		fmt.Println(i, "-", v)
	}

	
}
