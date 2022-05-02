package main

import "fmt"

func funcao1() {
	fmt.Println("Executando na Tela 1")
}
func funcao2() {
	fmt.Println("Executando na Tela 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado!")
	fmt.Println("Entrando na função para verificar se o aluno esta apto")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
	// defer funcao1()
	// funcao2()

}

//defer ---> significa adiar
