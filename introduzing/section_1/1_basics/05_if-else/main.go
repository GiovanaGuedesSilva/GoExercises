package main

/*
	Demonstração de estruturas condicionais em Go:

	- if simples
	- if/else
	- if/else if/else
	- if com declaração de variável dentro da condição
	- if sem else (bloco único)

	Go não usa parênteses nas condições e exige chaves obrigatórias.
*/

import (
	"fmt"
)

func main() {
	// if / else if / else comum
	s := "hi"
	if s == "hi" {
		fmt.Println("Hello")
	} else if s == "bye" {
		fmt.Println("See you")
	} else {
		fmt.Println("How’s it going")
	}

	// if com operação embutida (declaração de variável no if)
	x := 10
	n := 180
	if v := x + n; v <= 100 {
		fmt.Println("Value is less or equal than 100:", v)
	} else if v <= 200 {
		fmt.Println("Value is less or equal than 200:", v)
	} else {
		fmt.Println("The value is greater than 200:", v)
	}

	// if sem else
	i := 8
	j := 4
	if i%j == 0 {
		fmt.Println("It is divisible")
	}
}
