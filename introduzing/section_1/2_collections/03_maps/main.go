package main

/*
	Teoria — Maps em Go

	- Maps são coleções de pares chave-valor, semelhantes a dicionários em outras linguagens.
	- A chave deve ser de um tipo comparável (string, int, bool, etc.).
	- A criação pode ser feita com make ou de forma literal.
	- O valor zero de um map é nil.

	Criação:
		var m map[string]int             // nil map (não pode ser usado até ser inicializado)
		m := make(map[string]int)       // map vazio
		m := map[string]string{"a": "A"}

	Acesso e modificação:
		valor := m["chave"]
		m["nova"] = valor
		delete(m, "chave")

	Testar existência:
		valor, ok := m["chave"]
		ok será false se a chave não existir

	Maps não garantem ordem.
*/

import "fmt"

func main() {
	// 1. Criação de map com make
	idades := make(map[string]int)
	idades["Giovana"] = 25
	idades["Lucas"] = 30
	idades["Ana"] = 22

	fmt.Println("Idades:", idades)

	// 2. Criação literal
	frutas := map[string]string{
		"a": "abacaxi",
		"b": "banana",
		"c": "caju",
	}
	fmt.Println("Frutas:", frutas)

	// 3. Acesso a valores
	fmt.Println("Idade da Giovana:", idades["Giovana"])
	fmt.Println("Fruta com chave 'b':", frutas["b"])

	// 4. Verificando existência de chave
	valor, ok := frutas["x"]
	if ok {
		fmt.Println("Fruta com chave 'x':", valor)
	} else {
		fmt.Println("Chave 'x' não encontrada no map")
	}

	// 5. Deletando uma chave
	delete(frutas, "b")
	fmt.Println("Frutas após deletar chave 'b':", frutas)

	// 6. Iterando sobre um map
	for chave, valor := range idades {
		fmt.Println("Nome:", chave, "Idade:", valor)
	}

	// 7. Map vazio vs nil map
	var m1 map[string]string  // nil
	m2 := map[string]string{} // map vazio mas válido
	fmt.Println("m1 == nil:", m1 == nil)
	fmt.Println("m2 == nil:", m2 == nil)
}
