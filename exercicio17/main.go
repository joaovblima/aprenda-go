package main

import (
	"fmt"
)

func main() {
	diaDaSemana := "viagem"

	switch {
	case diaDaSemana == "domingo":
		fmt.Println("hoje nao vou trabalhar")
	case diaDaSemana == "segunda":
		fmt.Println("Hoje vou trabalhar e eh o dia mais dificil")
	case diaDaSemana ==  "terca":
		fmt.Println("Trabalho")
	case  diaDaSemana == "quarta":
		fmt.Println("Trabalho, mas hoje tem Tricolor")
	case diaDaSemana == "quinta":
		fmt.Println("trabalho, mas hoje eh dia de secar o flamengo e o palmeiras")
	case diaDaSemana == "sexta":
		fmt.Println("Trabalho, mas hoje eh sexta-feira!!!!!!!!!!!!!!!!!!!1")
	case diaDaSemana == "sabado":
		fmt.Println("FOLGA MEU PATRAO")
	default:
		fmt.Println("digitiu o dia errado meu patrao")		
	}
}