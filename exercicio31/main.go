package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloDeLuxo bool
}

func main() {

	hilux := caminhonete{veiculo{4, "prata",}, true,}
	
	bmw := sedan{ 
		veiculo: veiculo{
			portas: 4,
			cor: "preto",
	},
		modeloDeLuxo: true,
}
	fmt.Println(hilux)
	fmt.Println(bmw)
	fmt.Println("cor da hilux: ", hilux.cor)
	fmt.Println("cor da bmw: ", bmw.cor)
}