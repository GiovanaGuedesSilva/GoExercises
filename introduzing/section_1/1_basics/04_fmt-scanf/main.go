package main

/*
	Demonstração de leitura de dados com fmt.Scanf e fmt.Scanln em Go.

	- fmt.Scanf permite formatar a leitura como em C (com %s, %d, etc).
	- fmt.Scanln lê até a primeira quebra de linha (enter).
	- Todas as variáveis devem ser passadas por referência (&).
*/

import "fmt"

func main() {
	var nome string
	var idade int

	// Leitura com Scanf: exige formato exato e separação com espaço
	fmt.Print("Digite seu nome e idade (ex: Maria 30): ")
	fmt.Scanf("%s %d", &nome, &idade)
	fmt.Println("Scanf → Nome:", nome, "Idade:", idade)

	// Leitura com Scanln: quebra ao pressionar Enter
	var cidade string
	fmt.Print("Digite sua cidade: ")
	fmt.Scanln(&cidade)
	fmt.Println("Scanln → Cidade:", cidade)

	// Lendo múltiplos valores com Scan
	var a, b int
	fmt.Print("Digite dois números (ex: 10 20): ")
	fmt.Scan(&a, &b)
	fmt.Println("Scan → Soma:", a+b)
}
