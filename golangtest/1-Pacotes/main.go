package main

import (
	"fmt"
	"pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbok@gmail.com")
	//checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
