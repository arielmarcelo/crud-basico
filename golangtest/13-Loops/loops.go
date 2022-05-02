package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0
	// for 1 < 10 {
	// 	i++
	// 	fmt.Println("Icrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementado j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"joao", "davi", "lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Ariel",
		"sobrenome": "felipe",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
