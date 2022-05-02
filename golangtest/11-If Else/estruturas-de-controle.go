package main

import "fmt"

func main() {
	fmt.Println("estrutura de controle")

	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual aq 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é menor que zero")
	}
}
