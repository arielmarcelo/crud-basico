package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	usuario2 := usuario{"Daniel", 5}
	fmt.Println(usuario2)
}
