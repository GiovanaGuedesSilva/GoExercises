package main

/*
	Demonstração do uso de números aleatórios em Go.

	- O pacote math/rand permite gerar números pseudoaleatórios.
	- Para valores diferentes a cada execução, usamos rand.Seed com o tempo atual.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 🔹 Inicializa a semente com o tempo atual
	rand.Seed(time.Now().UnixNano())

	// 🔹 Número aleatório entre 0 e 99
	fmt.Println("Número aleatório de 0 a 99:", rand.Intn(100))

	// 🔹 Número aleatório entre 1 e 6 (simulando dado)
	fmt.Println("Valor do dado:", rand.Intn(6)+1)

	// 🔹 Número aleatório decimal entre 0.0 e 1.0
	fmt.Println("Número aleatório (float):", rand.Float64())

	// 🔹 Sorteando item de uma lista
	frutas := []string{"maçã", "banana", "laranja", "uva"}
	indice := rand.Intn(len(frutas))
	fmt.Println("Fruta sorteada:", frutas[indice])
}
