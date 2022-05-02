package main

import "fmt"

func DiaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda feira"
	case 3:
		return "Teca feira"
	default:
		return "Numero invalido"
	}

}
func DiaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda feira"
	case numero == 2:
		return "Terca feira"
	default:
		return "Numero ivalido"

	}

}

func main() {
	fmt.Println("Switch")

	dia := DiaDaSemana(3)
	fmt.Println(dia)

	dia2 := DiaDaSemana2(2)
	fmt.Println(dia2)

}
