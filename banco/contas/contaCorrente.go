package contas

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
	cpf           int
}

func NewContaCorrente(titular string, numeroAgencia int, numeroConta int, cpf int) ContaCorrente {
	conta := ContaCorrente{
		titular:       titular,
		numeroAgencia: numeroAgencia,
		numeroConta:   numeroConta,
		saldo:         0,
		cpf:           cpf,
	}
	return conta
}

func (c ContaCorrente) Cpf() int {
	return c.cpf
}

func (c ContaCorrente) Titular() string {
	return c.titular
}

func (c ContaCorrente) NumeroAgencia() int {
	return c.numeroAgencia
}

func (c ContaCorrente) NumeroConta() int {
	return c.numeroConta
}

func (c ContaCorrente) Saldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "Saldo insuficiente", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Falha na operação", c.saldo
	}

}
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
