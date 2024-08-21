package main

import (
	"fmt"

)

/*Em Go,valores podem ter mais de um tipo.

Uma interface permite que um valor tenha mais que um tipo.

Declaração: keyword identifier type -> type x interface

Apos declaração da interface, deve-se definir os metodos necessarios
para implementar essa interface.

Se um tipo possuir todos os metodos necessarios (que, no caso da interface{} pode ser nenhum)
entao esse tipo implicitamente implementa a interface

Esse tipo sera o teu tipo e *tambem* o tipo da interface.

Exemplo:

Os tipos *profissao1* e *profissao2* contem o tipo *pessoa*. Cada um tem seu metodo *oibomdia()*, e podem dar oi
utilizandi pessoa.oibomdia()
Implementem a interface *gente*
Ambos podem acessar o metodo *serhumano* que chama o metodo *oibomdia* de cada *gente*
Tambem podemos no metodo *serhumano()* tomar decisoes diferentes tipo:
switch pessoa.(type) {case profissao1: fmt.Println("qualquer coisa")...}
*/

/*
type pessoa struct {
	nome string 
	idade int
}

type professor struct {
	pessoa
	salario int
	
}

type engenheiro struct {
	pessoa 
	salario int

}

func (x professor ) oibomdia() {
	fmt.Println("E ai, meu nome é", x.nome, "bom dia!")
}
func (x engenheiro) oibomdia() {
	fmt.Println("E ai, meu nome é", x.nome, "bom dia!")
}


type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
		case professor:
			fmt.Println("E ai, eu sou", g.(professor).nome, "estou aqui para ensinar coisa nova." )
		
		case engenheiro:
			fmt.Println("Oi, eu sou", g.(engenheiro).nome , "sou engenheiro e vou construir sua casa")
	}
}



func main() {

	toninho := professor {
		pessoa: pessoa{
			nome: "toninho",
			idade: 56,
		},
		salario: 3500,
	}

	

	lele := engenheiro{
		pessoa: pessoa{
			nome: "Leticia",
			idade: 22,
		},
		salario: 10000,
	}


	serhumano(lele)
	serhumano(toninho)

}
*/
type Impressora interface {
	Imprimir()
}

type ImpressoraDeTexto struct {
	texto string
}

func (i ImpressoraDeTexto) Imprimir() {
	fmt.Println(i.texto)
}

type ImpressoraDeNumero struct {
	numero int
}

func (i ImpressoraDeNumero) Imprimir(){
	fmt.Println(i.numero)
}

func exibirImpressao(i Impressora) {
	i.Imprimir()
}

func main() {
	texto := ImpressoraDeTexto{texto: "E ai fi dos cabrunco"}
	numero := ImpressoraDeNumero{numero: 1231241231}

	exibirImpressao(texto)
	exibirImpressao(numero)
}



