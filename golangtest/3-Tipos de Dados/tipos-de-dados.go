package main

import (
	"errors"
	"fmt"
)

func main() {
	//numeros interios
	var numero int64 = -1000000
	fmt.Println(numero)

	var numero2 int32 = 1000
	fmt.Println(numero2)

	//alias
	//int32 = Rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.5
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000.52
	fmt.Println(numeroReal2)

	//fim dos numeros reais

	//strings
	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'b'
	fmt.Println(char)

	//fim strings

	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
