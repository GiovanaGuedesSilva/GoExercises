package main

/*
	Teoria — Arrays em Go

	- Um array é uma estrutura de dados com tamanho fixo e valores do mesmo tipo.
	- Todos os elementos são inicializados com o valor "zero" do tipo.
	- O tamanho faz parte do tipo em Go: [3]int e [4]int são tipos diferentes.
	- Arrays podem ser multidimensionais, como [2][3]int (matriz 2x3).
	- Arrays podem ser comparados com == se forem do mesmo tipo e tamanho.

	Criação:
		var a [3]int
		var b = [3]string{"a", "b", "c"}
		var c = [...]int{1, 2, 3} // tamanho inferido

	Iteração:
		for tradicional ou for com range

	Multidimensionais:
		var m [2][2]int → matriz de 2 linhas e 2 colunas
*/

import "fmt"

func main() {
	// 1. Criação de array com valores padrão (zero value)
	var numbers [3]int
	fmt.Println("Array vazio de inteiros:", numbers) // [0 0 0]

	// 2. Inicialização direta
	names := [3]string{"Giovana", "Lucas", "Ana"}
	fmt.Println("Array de strings:", names) // [Giovana Lucas Ana]

	// 3. Inferência de tamanho com ...
	values := [...]float64{3.1, 2.5, 5.0}
	fmt.Println("Array de floats:", values) // [3.1 2.5 5.0]

	// 4. Iteração com for tradicional
	for i := 0; i < len(names); i++ {
		fmt.Println("names[", i, "] =", names[i])
	}

	// 5. Iteração com range
	for i, v := range values {
		fmt.Printf("values[%d] = %.2f\n", i, v)
	}

	// 6. Modificação de valor
	numbers[1] = 99
	fmt.Println("Array modificado:", numbers) // [0 99 0]

	// 7. Matriz (array multidimensional)
	var matrix [2][2]int
	matrix[0][0] = 1
	matrix[1][1] = 2
	fmt.Println("Matriz 2x2:", matrix)

	// Iteração da matriz
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	// 8. Comparação de arrays
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{3, 2, 1}
	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false
}
