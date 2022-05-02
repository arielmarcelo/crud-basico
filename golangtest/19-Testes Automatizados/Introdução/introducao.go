package main

import (
	"fmt"
	enderecos "teste-enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida dos patos")
	fmt.Println(tipoEndereco)
}
