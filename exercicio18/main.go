package main 

import (
	"fmt"
)

func main() {
	esporteFavorito := "jiujitsu"
	switch esporteFavorito {
		case "jiujitsu":
			fmt.Println("OSSSSSSSSSSS")
		case "futebol":
			fmt.Println("Torce pro SPFC")
		case "tenis":
			fmt.Println("Fã de Djokovic")
		case "vôlei":
			fmt.Println("Fã de Giba")
		default:
			fmt.Println("vamos nos exercitar guerreiro")
		}
}