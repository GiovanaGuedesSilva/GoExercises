package main

/*
	Teoria — Runes em Go

	- Em Go, uma *rune* é um alias para o tipo int32 e representa um caractere Unicode.
	- Strings são sequências de bytes codificadas em UTF-8.
	- Um único caractere pode ocupar mais de um byte, então iterar por índice em uma string pode quebrar caracteres.
	- Para acessar corretamente os caracteres (runes), usamos o `for range`, que interpreta cada rune completo.

	Conversões úteis:
		- []rune(str): converte string em slice de runes
		- string(r): converte rune (int32) para string

	Uso comum:
		- Contar caracteres corretamente
		- Manipular strings multilinguagens
*/

import (
	"fmt"
)

func main() {
	// 1. String com caracteres multibyte
	s := "Olá, 世界"
	fmt.Println("String:", s)

	// 2. Tamanho em bytes (não em caracteres)
	fmt.Println("len(s):", len(s)) // conta bytes

	// 3. Iterando com índice (por byte)
	fmt.Println("Iteração por byte:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte %d = %x\n", i, s[i])
	}

	// 4. Iterando com range (por rune)
	fmt.Println("Iteração por rune:")
	for i, r := range s {
		fmt.Printf("Posição %d = %c (Unicode: %U)\n", i, r, r)
	}

	// 5. Convertendo string para slice de runes
	runes := []rune(s)
	fmt.Println("Total de runes (caracteres reais):", len(runes))

	// 6. Acessando caracteres individuais corretamente
	fmt.Println("Primeira letra (rune):", string(runes[0]))
	fmt.Println("Último caractere (rune):", string(runes[len(runes)-1]))
}
