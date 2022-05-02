package main

import "fmt"

func recuperarExecuçao() {
	if r := recover(); r != nil {
		fmt.Println("Execução revuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic(("A media e exatamente 6!"))
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 7))
	fmt.Println("Pós execução")

}
