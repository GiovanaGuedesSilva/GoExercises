package main

/*
	Demonstração de encapsulamento em Go.

	- Em Go, o encapsulamento é controlado pela visibilidade de nomes.
	- Se o nome de um campo, função ou método começa com letra **maiúscula**, ele é **exportado** (público).
	- Se começa com letra **minúscula**, ele é **não exportado** (privado ao pacote).

	Go não possui modificadores como "private", "protected" ou "public".
	A visibilidade é definida apenas pela inicial maiúscula ou minúscula.

	- Estrutura recomendada:
		- Campos privados acessados por métodos públicos (get/set)
		- Regras de negócio aplicadas dentro dos métodos
*/

import (
	"fmt"
)

// Struct com campos encapsulados
type Conta struct {
	titular string  // campo privado
	saldo   float64 // campo privado
}

// Construtor (retorna instância com campos privados)
func NovaConta(titular string, saldoInicial float64) Conta {
	return Conta{titular: titular, saldo: saldoInicial}
}

// Método público para acessar o titular
func (c Conta) Titular() string {
	return c.titular
}

// Método público para acessar o saldo
func (c Conta) Saldo() float64 {
	return c.saldo
}

// Método para realizar depósito com validação
func (c *Conta) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
	}
}

// Método para saque com validação
func (c *Conta) Sacar(valor float64) bool {
	if valor <= c.saldo {
		c.saldo -= valor
		return true
	}
	return false
}

func main() {
	// Criando uma conta com construtor
	conta := NovaConta("Giovana", 1000)

	// Acessando dados encapsulados via métodos públicos
	fmt.Println("Titular:", conta.Titular())
	fmt.Println("Saldo inicial:", conta.Saldo())

	// Tentando sacar valor válido
	if conta.Sacar(200) {
		fmt.Println("Saque de R$200 realizado")
	} else {
		fmt.Println("Saldo insuficiente")
	}
	fmt.Println("Saldo atual:", conta.Saldo())

	// Tentando saque inválido
	if !conta.Sacar(10000) {
		fmt.Println("Saque de R$10000 negado")
	}

	// Fazendo depósito
	conta.Depositar(500)
	fmt.Println("Saldo após depósito:", conta.Saldo())
}
