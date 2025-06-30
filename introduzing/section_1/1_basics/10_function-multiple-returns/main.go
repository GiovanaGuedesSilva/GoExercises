package main

/*
	Demonstração de funções com múltiplos retornos em Go.

	- Em Go, uma função pode retornar mais de um valor.
	- Muito comum em operações como divisão, chamadas de sistema e manipulação de erros.
*/

import "fmt"

// Função que retorna quociente e resto de uma divisão inteira
func dividir(a int, b int) (int, int) {
	quociente := a / b
	resto := a % b
	return quociente, resto
}

// Função que retorna nome completo e tamanho total
func nomeCompleto(nome string, sobrenome string) (string, int) {
	completo := nome + " " + sobrenome
	tamanho := len(completo)
	return completo, tamanho
}

func main() {
	// Exemplo 1: Divisão
	div, mod := dividir(10, 3)
	fmt.Println("Divisão inteira de 10 por 3:", div)
	fmt.Println("Resto da divisão:", mod)

	// Exemplo 2: Nome completo
	nome, tamanho := nomeCompleto("Giovana", "Silva")
	fmt.Println("Nome completo:", nome)
	fmt.Println("Número de caracteres:", tamanho)
}
