package main 

import (
	"fmt"
)

func main() {
	idade := 15

	if idade > 18 {
		fmt.Println("parabens velhote, ja pode tomar uma breja")
	} else {
		fmt.Printf("muito novinho, volte daqui ha %v anos", (18 - idade))
	}
}