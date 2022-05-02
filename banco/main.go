package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.NewContaCorrente("silvia", 1, 123, 789)
	contaDoGustavo := contas.NewContaCorrente("Gustavo", 2, 2, 456)

	status, saldo := contaDaSilvia.Depositar(500)

	fmt.Println(contaDaSilvia.Titular(), saldo, status)

	status, saldo = contaDoGustavo.Depositar(100)

	fmt.Println(contaDoGustavo.Titular(), saldo, status)

	stats := contaDaSilvia.Transferir(200, &contaDoGustavo)

	fmt.Println(stats)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}
