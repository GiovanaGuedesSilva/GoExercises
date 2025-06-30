package main

/*
	Teoria — Strings em Go

	- Strings em Go são imutáveis, codificadas em UTF-8.
	- Cada caractere (rune) pode ocupar mais de um byte.
	- O índice de uma string acessa bytes, não necessariamente caracteres completos.

	Declaração:
		var s string = "texto"
		s := "outra forma"

	Principais funções da biblioteca "strings":
		- strings.ToUpper(s), strings.ToLower(s)
		- strings.Contains(s, substr)
		- strings.HasPrefix(s, prefix), strings.HasSuffix(s, suffix)
		- strings.Split(s, sep)
		- strings.Join([]string, sep)
		- strings.ReplaceAll(s, old, new)
		- strings.TrimSpace(s)

	Atenção:
		- Para iterar sobre caracteres corretamente, use `for range`, pois ele trabalha com runes.
*/

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Declaração de strings
	s1 := "Olá, Go!"
	s2 := "Olá, Go!"
	s3 := `Texto
	com múltiplas linhas`
	fmt.Println("s1:", s1)
	fmt.Println("s3 (multilinha):", s3)

	// 2. Comparação
	fmt.Println("s1 == s2:", s1 == s2)

	// 3. Concatenação
	joined := s1 + " Seja bem-vindo."
	fmt.Println("Concatenação:", joined)

	// 4. Tamanho em bytes
	fmt.Println("len(s1):", len(s1)) // quantidade de bytes, não de caracteres

	// 5. Iteração por byte (índice tradicional)
	for i := 0; i < len(s1); i++ {
		fmt.Printf("Byte %d = %x\n", i, s1[i])
	}

	// 6. Iteração por caractere (rune)
	for i, r := range s1 {
		fmt.Printf("Posição %d = %c\n", i, r)
	}

	// 7. Algumas funções da biblioteca strings
	texto := "   Go é ótimo! go go go   "
	fmt.Println("Original:", texto)
	fmt.Println("ToUpper:", strings.ToUpper(texto))
	fmt.Println("ToLower:", strings.ToLower(texto))
	fmt.Println("Contains 'go':", strings.Contains(texto, "go"))
	fmt.Println("HasPrefix 'Go':", strings.HasPrefix(texto, "Go"))
	fmt.Println("TrimSpace:", strings.TrimSpace(texto))

	// 8. Split e Join
	frase := "Golang é incrível"
	palavras := strings.Split(frase, " ")
	fmt.Println("Split:", palavras)
	fmt.Println("Join:", strings.Join(palavras, "-"))

	// 9. Replace
	replaced := strings.ReplaceAll(texto, "go", "GO")
	fmt.Println("ReplaceAll:", replaced)
}
