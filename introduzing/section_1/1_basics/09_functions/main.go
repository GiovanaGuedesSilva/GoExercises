package main

/*
	Demonstração de funções em Go:

	- Funções básicas com parâmetros e retorno
	- Funções com múltiplos retornos
	- Funções com retorno nomeado
	- Funções anônimas
	- Funções como valor (function values)
	- Funções closures
*/

import "fmt"

func main() {
	// Função simples
	saudacao("Giovana")

	// Função com retorno
	soma := somar(5, 3)
	fmt.Println("Soma:", soma)

	// Função com múltiplos retornos
	div, resto := dividir(10, 3)
	fmt.Println("Divisão:", div, "Resto:", resto)

	// Função com retorno nomeado
	total := dobrar(4)
	fmt.Println("Dobro:", total)

	// Função anônima (sem nome, executada imediatamente)
	func(msg string) {
		fmt.Println("Mensagem da função anônima:", msg)
	}("Executando agora")

	// Função atribuída a variável
	multiplicar := func(a, b int) int {
		return a * b
	}
	fmt.Println("Multiplicação:", multiplicar(3, 4))

	// Closure: função que mantém o estado da variável
	contador := criarContador()
	fmt.Println("Contador:", contador())
	fmt.Println("Contador:", contador())
	fmt.Println("Contador:", contador())
}

// Função com parâmetro e sem retorno
func saudacao(nome string) {
	fmt.Println("Olá,", nome)
}

// Função com retorno simples
func somar(a int, b int) int {
	return a + b
}

// Função com múltiplos retornos
func dividir(a int, b int) (int, int) {
	return a / b, a % b
}

// Função com retorno nomeado
func dobrar(n int) (resultado int) {
	resultado = n * 2
	return
}

// Closure: função que mantém contexto da variável `x`
func criarContador() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
