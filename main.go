package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	conta1 := contas.ContaCorrente{Titular: clientes.Titular{Nome: "João", CPF: "908.833.149-91", Profissao: "Desenvolvedor"}, NumeroAgencia: 1, NumeroConta: 2}
	clienteMaria := clientes.Titular{Nome: "Maria"}
	conta2 := contas.ContaCorrente{Titular: clienteMaria, NumeroAgencia: 1, NumeroConta: 3}
	
	conta1.Depositar(1000)
	conta2.Depositar(500)
	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println(conta1.ObterSaldo())
	fmt.Println(conta2.ObterSaldo())

	status, valor := conta1.Transferir(500, &conta2)
	fmt.Println(status, valor)

	fmt.Println(conta1.ObterSaldo())
	fmt.Println(conta2.ObterSaldo())

}
