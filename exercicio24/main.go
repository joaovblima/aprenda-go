package main 

import (
	"fmt"
)

func main() {
	estados := make([]string, 26, 26)
	estados = []string{ 
		"Acre (Rio Branco)", 
		"Alagoas (Maceió)",
		"Amapá (Macapá)",
		"Amazonas (Manaus)",
		"Bahia (Salvador)",
		"Ceará (Fortaleza)",
		"Distrito Federal (Brasília)",
		"Espírito Santo (Vitória)",
		"Goiás (Goiânia)",
		"Maranhão (São Luís)",
		"Mato Grosso (Cuiabá)",
		"Mato Grosso do Sul (Campo Grande)",
		"Minas Gerais (Belo Horizonte)",
		"Pará (Belém)",
		"Paraíba (João Pessoa)",
		"Paraná (Curitiba)",
		"Pernambuco (Recife)",
		"Piauí (Teresina)",
		"Rio de Janeiro (Rio de Janeiro)",
		"Rio Grande do Norte (Natal)",
		"Rio Grande do Sul (Porto Alegre)",
		"Rondônia (Porto Velho)",
		"Roraima (Boa Vista)",
		"Santa Catarina (Florianópolis)",
		"São Paulo (São Paulo)",
		"Sergipe (Aracaju)",
		"Tocantins (Palmas)"}
	fmt.Println(estados)
	fmt.Println(len(estados), cap(estados))

	for i :=0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}

}