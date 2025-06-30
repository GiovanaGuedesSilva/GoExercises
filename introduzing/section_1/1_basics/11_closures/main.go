package main

/*
	Demonstração de closures em Go.

	- Um closure é uma função que “lembra” das variáveis que estavam no escopo onde ela foi criada.
	- Cada vez que `sayHello` ou `seqOfTwo` é chamada, ela retorna uma nova função com seu próprio "estado interno".
*/

import "fmt"

// Closure que acumula frases
func sayHello() func(string) string {
	h := "YYY" // variável capturada pela função interna
	return func(b string) string {
		h = h + " " + b // concatena o novo valor com o anterior
		return h
	}
}

// Closure que gera uma sequência de números de 2 em 2
func seqOfTwo() func() int {
	i := 0 // variável interna que será incrementada
	return func() int {
		i = i + 2
		return i
	}
}

func main() {
	// Criação de duas funções independentes com estado próprio
	a := sayHello()
	b := sayHello()

	// a mantém o histórico próprio de concatenação
	fmt.Println(a("Hello golang")) // YYY Hello golang
	fmt.Println(a("how are you?")) // YYY Hello golang how are you?

	// b começa do zero, separada de a
	fmt.Println(b("Hi!"))      // YYY Hi!
	fmt.Println(b("what up?")) // YYY Hi! what up?

	// Criação de dois contadores separados
	c := seqOfTwo()
	fmt.Println(c()) // 2
	fmt.Println(c()) // 4
	fmt.Println(c()) // 6

	d := seqOfTwo()
	fmt.Println(d()) // 2 (novo contador)
	fmt.Println(d()) // 4
	fmt.Println(d()) // 6
}
