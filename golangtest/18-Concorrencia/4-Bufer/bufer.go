package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Ola Mundo!"
	canal <- "Programa em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	
}
