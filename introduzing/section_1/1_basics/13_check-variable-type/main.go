package main

/*
	Demonstração de como verificar o tipo de variáveis em tempo de execução em Go.

	- Type Assertion: permite afirmar e extrair o tipo real de uma interface{}.
	- Type Switch: estrutura que permite testar vários tipos possíveis.

	Muito útil quando lidamos com interfaces genéricas.
*/

import "fmt"

func main() {
	var i interface{}

	// 🔹 Exemplo 1: atribuindo diferentes valores a uma interface{}
	i = "texto"
	checkType(i)

	i = 123
	checkType(i)

	i = true
	checkType(i)

	i = 3.14
	checkType(i)

	// 🔹 Exemplo 2: usando type assertion diretamente
	var x interface{} = "Hello, world"
	s, ok := x.(string)
	if ok {
		fmt.Println("Type assertion bem-sucedida:", s)
	} else {
		fmt.Println("Falha na type assertion")
	}

	// Type assertion incorreta (irá causar panic se não tratada)
	// f := x.(float64) // descomente para ver o panic
}

func checkType(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("É um int:", t)
	case string:
		fmt.Println("É uma string:", t)
	case bool:
		fmt.Println("É um bool:", t)
	case float64:
		fmt.Println("É um float64:", t)
	default:
		fmt.Println("Tipo desconhecido")
	}
}
