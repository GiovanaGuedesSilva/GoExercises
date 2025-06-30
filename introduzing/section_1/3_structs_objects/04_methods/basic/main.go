package main

/*
	Demonstração de métodos em Go.

	- Métodos são funções com um *receiver* (receptor), associado a um tipo (geralmente uma struct).
	- Sintaxe: func (r NomeTipo) NomeMetodo() TipoRetorno
	- Métodos podem usar receiver por valor ou por ponteiro.
	- Receiver por valor: cópia → não altera a struct original.
	- Receiver por ponteiro: referência → altera a struct original.
*/

import "fmt"

// Tipo simples
type Contador struct {
	Valor int
}

// Método com receiver por valor (não altera o original)
func (c Contador) Mostrar() {
	fmt.Println("Valor atual:", c.Valor)
}

// Método com receiver por ponteiro (altera o original)
func (c *Contador) Incrementar() {
	c.Valor++
}

func main() {
	c := Contador{Valor: 10}

	// Método por valor (apenas leitura)
	c.Mostrar()

	// Método por ponteiro (altera o estado interno)
	c.Incrementar()
	c.Incrementar()
	c.Mostrar()
}
