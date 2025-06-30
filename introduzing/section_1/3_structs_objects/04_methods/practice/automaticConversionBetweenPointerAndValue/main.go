package main

import "fmt"

type Conta struct {
	saldo float64
}

// Receptor ponteiro: modifica o valor
func (c *Conta) Depositar(valor float64) {
	c.saldo += valor
}

// Receptor valor: não modifica
func (c Conta) VerSaldo() float64 {
	return c.saldo
}

func main() {
	c := Conta{}
	c.Depositar(100)                    // Modifica saldo
	fmt.Println("Saldo:", c.VerSaldo()) // OK
}

// Ponteiro para modificar (Depositar)
// Valor para ler (VerSaldo)
// ⚠️ Boa prática: Se um método usa ponteiro, os outros também devem. Evita confusão.
