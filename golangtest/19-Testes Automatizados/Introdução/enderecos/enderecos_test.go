package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	endececoIncerido string
	retornoEsperado  string
}

func TestTiposDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Avenida dos patos", "Avenida"},
		{"Rua das galhos", "Rua"},
		{"Avenida dos patros", "Avenida"},
		{"Rua das galinhas", "Rua"},
		{"Avenida dos patios", "Avenida"},
		{"Rua das galinheiros", "Rua"},
		{"Praça das galinheiros", "Tipo Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.endececoIncerido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
