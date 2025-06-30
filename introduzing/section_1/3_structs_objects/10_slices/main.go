package main

/*
	Demonstração do uso de slices em Go.

	- Slice é uma estrutura que representa uma sequência de elementos de mesmo tipo.
	- É mais flexível que array: pode crescer dinamicamente.
	- Internamente, slice é uma "visão" de um array com tamanho e capacidade controlados.

	Principais pontos:
	- Declaração com `[]tipo` e valores iniciais.
	- Uso de `append()` para adicionar elementos.
	- Acesso com `slice[i]` e slicing `slice[start:end]`.
	- Função `len()` retorna o tamanho atual.
	- Função `cap()` retorna a capacidade total do array interno.
*/

import "fmt"

func main() {
	// Declarando um slice com valores
	frutas := []string{"maçã", "banana", "uva"}
	fmt.Println("Frutas:", frutas)

	// Acessando elementos por índice
	fmt.Println("Primeira fruta:", frutas[0])
	fmt.Println("Última fruta:", frutas[len(frutas)-1])

	// Verificando o tamanho (len) e a capacidade (cap)
	fmt.Println("Tamanho:", len(frutas))
	fmt.Println("Capacidade:", cap(frutas))

	// Adicionando elementos com append (retorna um novo slice)
	frutas = append(frutas, "abacaxi")
	fmt.Println("Frutas após append:", frutas)

	// Subslice: pegando parte do slice
	subslice := frutas[1:3] // do índice 1 até 2 (3 não incluso)
	fmt.Println("Subslice [1:3]:", subslice)

	// Criando slice com make (tamanho 3, capacidade 5)
	numeros := make([]int, 3, 5)
	fmt.Println("Slice com make:", numeros)
	fmt.Println("Tamanho:", len(numeros), "Capacidade:", cap(numeros))

	// Preenchendo e expandindo slice
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros = append(numeros, 40, 50)
	fmt.Println("Numeros após append:", numeros)

	// Slice vazio
	var vazio []int
	fmt.Println("Slice vazio:", vazio, "Tamanho:", len(vazio), "Capacidade:", cap(vazio))
}
