package main

/*
	Demonstração do uso de type switch em Go.

	- Um *type switch* permite identificar o tipo concreto armazenado em uma interface{}.
	- Útil para lógica condicional baseada no tipo real do valor em tempo de execução.
	- A sintaxe usa a palavra-chave `.(type)` dentro de um switch.
	- Cada `case` representa um tipo específico, podendo haver também um `default`.

	Cuidados:
	- Deve ser usado com `interface{}`.
	- A variável definida dentro do switch pode ser reutilizada em cada `case`.
*/

import "fmt"

func mostrarTipo(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("É um int:", v)
	case float64:
		fmt.Println("É um float64:", v)
	case string:
		fmt.Println("É uma string:", v)
	case bool:
		fmt.Println("É um booleano:", v)
	default:
		fmt.Println("Tipo desconhecido")
	}
}

func main() {
	var valores = []interface{}{42, 3.14, "golang", true, []int{1, 2, 3}}

	for _, v := range valores {
		mostrarTipo(v)
	}
}
