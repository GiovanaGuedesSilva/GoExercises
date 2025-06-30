package main

/*
	Demonstração de types básicos e type assertion em Go.

	- Type assertion permite extrair o valor concreto armazenado em uma interface{}.
	- Pode ser feito de forma segura (usando ", ok") ou insegura (sem verificação).
	- Se a asserção falhar sem verificação, o programa dá panic.
*/

import "fmt"

func main() {
	// Interface vazia contendo uma string
	var i interface{} = "hello"

	// Type assertion segura (com verificação de tipo)
	s, ok := i.(string)
	fmt.Println("Com ok (string):", s, ok) // hello true

	// Type assertion insegura (sem verificação)
	// Só use se você tiver certeza do tipo
	s2 := i.(string)
	fmt.Println("Sem ok (string):", s2) // hello

	// Tentando converter para float64 de forma segura
	f, ok := i.(float64)
	fmt.Println("Com ok (float64):", f, ok) // 0 false

	// Tentando converter para float64 sem verificação → panic
	/*
		f2 := i.(float64)
		fmt.Println("Sem ok (float64):", f2) // panic: interface {} is string, not float64
	*/

	// Outro exemplo simples e seguro
	var x interface{} = "sou uma string"
	if str, ok := x.(string); ok {
		fmt.Println("Valor convertido com sucesso:", str)
	} else {
		fmt.Println("Conversão falhou")
	}
}
