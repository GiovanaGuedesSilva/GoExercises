package main

/*
	Demonstração de declaração de variáveis e constantes em Go.
	- Tipagem explícita e implícita
	- Zero values
	- Declaração curta (:=)
	- Constantes com e sem tipo
	- Uso de iota para gerar sequências
*/

import (
	"fmt"
)

// Constante de pacote (visível em todo o arquivo)
const pi = 3.1415

// Constantes com tipo explícito
const site string = "google.com"
const idadeMinima int = 18

// Constantes com múltiplas declarações e iota
const (
	// iota começa em 0 e incrementa automaticamente
	Domingo = iota // 0
	Segunda        // 1
	Terça          // 2
	Quarta         // 3
)

func main() {
	// VARIÁVEIS

	// Declaração explícita com tipo
	var hello string = "hello world"
	fmt.Println("Hello:", hello)

	// int explícito
	var age int = 30
	fmt.Println("Age:", age)

	// float64 explícito
	var fb float64 = 30.12345678901234567890
	fmt.Println("Float fb:", fb)

	// Múltiplas atribuições em uma linha
	var hi, word string = "hi", "word"
	fmt.Println("Multiple:", hi, word)

	// Valor padrão (zero-value) para int
	var defaultInt int
	fmt.Println("Zero-value int:", defaultInt) // 0

	// Zero-value para bool
	var defaultBoolean bool
	fmt.Println("Zero-value bool:", defaultBoolean) // false

	// Declaração curta com :=
	great := "great work"
	fmt.Println("Short var declaration:", great)

	// Inferência de tipo
	var b = true
	fmt.Println("Inferência booleana:", b)

	// CONSTANTES

	// Constante local sem tipo explícito
	const nome = "Giovana"

	// Constante com tipo explícito
	const temperatura float64 = 36.6

	fmt.Println("Nome:", nome)
	fmt.Println("Site:", site)
	fmt.Println("PI:", pi)
	fmt.Println("Temperatura:", temperatura)
	fmt.Println("Idade mínima:", idadeMinima)

	// Exibindo valores com iota
	fmt.Println("Dias da semana (via iota):")
	fmt.Println("Domingo =", Domingo)
	fmt.Println("Segunda =", Segunda)
	fmt.Println("Terça =", Terça)
	fmt.Println("Quarta =", Quarta)
}
