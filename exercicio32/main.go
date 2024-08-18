package main

import (
	"fmt"
)

func main()  {

	y := struct {
		nome string
		idade int
		numeros[]int
		comidas map[string]string
	}{
		nome: "Joao",
		idade: 27,
		numeros: []int{20, 99, 33, 45},
		comidas: map[string]string{
			"ovo": "almoco",
			"paocomqueijo": "lanche",
			"saladadefrutas": "todahora",
		},
	}

	fmt.Println(y)


}