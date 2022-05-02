package main

import "fmt"

func main() {
	//Aritimeticos --> + - / * %
	soma := 1 + 1
	subtracao := 1 - 2
	divisao := 1 / 3
	multiplicacao := 1 * 4
	restoDaDivisao := 1 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//Atribuiçao

	var variavel1 string = "String"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	// Operadores Relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos
	fmt.Println("-")
	verdadeiro, false := true, false
	fmt.Println(verdadeiro && false)
	fmt.Println(verdadeiro || false)

	//Operadores Unários
	numero := 10
	numero++     // igual a 10 + 1
	numero += 15 // igual a 10 + 15
	fmt.Println(numero)

	// numero--   // igual a 10 - 1
	// numero-= 15 // igual a 10 - 15

	// Operador ternario
	var texto string
	if numero > 5 {
		texto = "Mairo que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Sprintln(texto)

}
