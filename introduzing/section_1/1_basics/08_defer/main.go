package main

/*
	Demonstração do uso de `defer` em Go.

	- `defer` adia a execução de uma função até que a função que a contém termine.
	- Usado geralmente para liberar recursos, fechar arquivos ou imprimir ao final.
	- Vários `defer` são empilhados e executados em ordem reversa (LIFO).
*/

import (
	"fmt"
)

func main() {
	fmt.Println("Início do main")

	defer fmt.Println("Defer 1 - será executado por último (dentro do main)")
	defer fmt.Println("Defer 2 - será executado antes do Defer 1")

	executar()

	fmt.Println("Fim do main")
}

func executar() {
	fmt.Println("Início da função executar")

	defer fmt.Println("Defer interno da função executar")

	fmt.Println("Fim da função executar")
}

// Início da função executar
// Fim da função executar
// Defer interno da função executar
// Fim do main
// Defer 2 - será executado antes do Defer 1
// Defer 1 - será executado por último (dentro do main)
