package main

import (
	"fmt"

)

func main() {

	valores := map[string][]string {
		"limajoao": []string {"jiujitsu, futebol, viajar"},
		"tourinho": []string {"comer", "futebol", "passear"},
		"lele": []string {"reclamar", "comer besteira", "dormir"},
	}

	for i, _ := range valores {
		fmt.Println(i)
		for _, v := range valores {
			fmt.Println("\t", " * ", v)

		}
	}

}