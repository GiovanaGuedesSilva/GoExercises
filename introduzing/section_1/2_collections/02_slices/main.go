package main

/*
	Teoria — Slices em Go

	- Slices são estruturas dinâmicas baseadas em arrays.
	- Ao contrário dos arrays, slices não têm tamanho fixo.
	- Um slice contém: ponteiro para o array, tamanho e capacidade.
	- A capacidade representa o número de elementos que podem ser armazenados antes de realocar o array subjacente.

	Criação:
		var s []int
		s := []int{1, 2, 3}
		s := make([]int, tamanho, capacidade)

	Fatiamento:
		slice := array[1:4] // do índice 1 até 3 (exclui o 4)

	Funções úteis:
		len(slice)     → quantidade de elementos
		cap(slice)     → capacidade interna
		append(slice, v) → adiciona valores
		copy(dest, src)  → copia valores

	Comparação:
		Não é possível comparar slices diretamente com ==
*/

import "fmt"

func main() {
	// 1. Criação de slices
	var a []int            // slice vazio
	b := []int{10, 20, 30} // slice com valores
	c := make([]string, 3) // slice com make
	fmt.Println("Slice a (vazio):", a)
	fmt.Println("Slice b (com valores):", b)
	fmt.Println("Slice c (make):", c)

	// 2. Modificando valores
	c[0] = "Giovana"
	c[1] = "Lucas"
	c[2] = "Ana"
	fmt.Println("Slice c modificado:", c)

	// 3. Acessando comprimento e capacidade
	fmt.Println("len(b):", len(b))
	fmt.Println("cap(b):", cap(b))

	// 4. Append - adicionando elementos
	b = append(b, 40)
	b = append(b, 50, 60)
	fmt.Println("Slice b após append:", b)

	// 5. Slicing - fatiamento de slices
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("numeros[1:4]:", numeros[1:4]) // [2 3 4]
	fmt.Println("numeros[:3]:", numeros[:3])   // [1 2 3]
	fmt.Println("numeros[2:]:", numeros[2:])   // [3 4 5]

	// 6. Copy - cópia de slices
	origem := []int{1, 2, 3}
	destino := make([]int, len(origem))
	copy(destino, origem)
	fmt.Println("Origem:", origem)
	fmt.Println("Destino (cópia):", destino)

	// 7. Slice compartilhando mesmo array base
	x := []int{100, 200, 300, 400}
	y := x[1:3]
	y[0] = 999
	fmt.Println("x após modificar y:", x) // x[1] também muda

	// 8. Slice nil vs slice vazio
	var nilSlice []int    // nil
	emptySlice := []int{} // vazio, mas alocado
	fmt.Println("nilSlice == nil:", nilSlice == nil)
	fmt.Println("emptySlice == nil:", emptySlice == nil)
}
