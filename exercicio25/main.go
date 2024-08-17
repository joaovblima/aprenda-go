package main

import (
	"fmt"
)

func main() {
	info := [][]string{
			  []string{
				"nome",
				"sobrenome", 
				"hobby",
			  }, 
			  []string{
				"pendrive",
				"silva",
				"comer",
			  },
			  []string{
				"JL",
				"BERNAR",
				"jj",
			  },
		
	}
	for _, v := range info {
		fmt.Println(v)
	}
	
}