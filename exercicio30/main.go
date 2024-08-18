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


	mapaDePessoas := make(map[string]pessoa)
	
	mapaDePessoas["Lima"] = pessoa{
		nome:"Joao",
		sobrenome: "Lima",
		sabores: []string{"Napolitano", "Nutella", "Pistache"},}

	mapaDePessoas["Sofia"] = pessoa{
		nome: "Maria",
		sobrenome: "Sofia",
		sabores: []string{"Nutella", "Ninho Com Oreo", "Diamante Negro"},
	}

	for _, value := range mapaDePessoas {
		fmt.Println("E ai, meu nome é", value.nome, "meus sabores favoritos de sorvete são: ")
		for _, val := range value.sabores {
			fmt.Println("-", val)
		}
	}
	
	

	
}
