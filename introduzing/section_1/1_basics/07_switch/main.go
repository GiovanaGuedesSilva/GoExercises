package main

/*
	Demonstração do uso do switch em Go.

	- Switch tradicional com múltiplos casos.
	- Casos agrupados (case 4, 5).
	- Switch sem expressão, usado como if/else.
	- Uso de fallthrough para continuar no próximo case.
	- Switch de tipo (type switch) com interface{}.
*/

import (
	"fmt"
	"time"
)

func main() {
	// Switch tradicional com valor fixo
	value := 4
	switch value {
	case 1:
		fmt.Println("Value is one")
	case 2:
		fmt.Println("Value is two")
	case 3:
		fmt.Println("Value is three")
	case 4, 5:
		fmt.Println("Value is four or five")
	default:
		fmt.Println("I do not recognize the value")
	}

	// Switch sem expressão, como if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// Switch com fallthrough (continua no próximo case)
	fmt.Println("Switch com fallthrough:")
	x := 1
	switch x {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2 (executado por fallthrough)")
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default case")
	}

	// 🔹 Type switch (verifica o tipo real da interface)
	fmt.Println("Type switch:")
	var i interface{} = 3.14

	switch v := i.(type) {
	case int:
		fmt.Println("Tipo int com valor:", v)
	case float64:
		fmt.Println("Tipo float64 com valor:", v)
	case string:
		fmt.Println("Tipo string com valor:", v)
	default:
		fmt.Println("Tipo desconhecido")
	}
}
