package main

import "fmt"

func main() {
	// Criação de uma slice com valores iniciais
	// Creating a slice with initial values
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Println("\nSlice original:")
	printSlice(numeros)

	// Acessar elementos específicos
	// Access specific elements
	fmt.Println("\nElemento na posição 2:", numeros[2])

	// Fatiamento (slice de slice)
	// Slicing (creating a subslice)
	sub := numeros[1:4] // índices 1 até 3 // indexes 1 to 3
	fmt.Println("\nSub-slice de índice 1 a 3:")
	printSlice(sub)

	// Reduzir comprimento a zero
	// Reduce length to zero
	numeros = numeros[:0]
	fmt.Println("\nSlice com comprimento 0:")
	printSlice(numeros)

	// Expandir comprimento novamente (respeitando a capacidade)
	// Expand length again (within capacity)
	numeros = append(numeros, 10, 20, 30, 40, 50)
	fmt.Println("\nSlice reexpandido com append:")
	printSlice(numeros)

	// Cortar os primeiros dois elementos
	// Remove the first two elements
	numeros = numeros[2:]
	fmt.Println("\nSlice após remover os dois primeiros:")
	printSlice(numeros)

	// Append de um novo valor
	// Append a new value
	numeros = append(numeros, 60)
	fmt.Println("\nSlice após adicionar novo valor com append:")
	printSlice(numeros)

	// Validação de slice nula
	// Nil slice check
	var sliceNula []int
	if sliceNula == nil {
		fmt.Println("\nA slice está nula (nil) — slice is nil")
	}

	// Criar uma nova slice com `make` (define tamanho e capacidade)
	// Create a new slice using `make` (define length and capacity)
	s := make([]int, 3, 5) // len = 3, cap = 5
	fmt.Println("\nNova slice criada com make:")
	printSlice(s)

	// Copiar slices
	// Copy slices
	copy(s, numeros)
	fmt.Println("\nSlice 's' após copy de 'numeros':")
	printSlice(s)

	// Concatenar duas slices
	// Concatenate two slices
	outra := []int{100, 200}
	s = append(s, outra...)
	fmt.Println("\nSlice após concatenar com outra slice:")
	printSlice(s)

	// Slice de slices (matriz/2D)
	// Slice of slices (2D array-like)
	fmt.Println("\nSlice de slices:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}
	for i, linha := range matrix {
		fmt.Printf("Linha %d: %v\n", i, linha)
	}

	// Uso de range com índice e valor
	// Using range with index and value
	fmt.Println("\nIterando com for range (índice e valor):")
	for i, val := range numeros {
		fmt.Printf("Índice %d, Valor %d\n", i, val)
	}

	// Ignorando o índice com _
	// Ignoring the index using _
	fmt.Println("\nIterando ignorando o índice:")
	for _, val := range numeros {
		fmt.Printf("Valor: %d\n", val)
	}

	// Ignorando o valor com _
	// Ignoring the value using _
	fmt.Println("\nIterando ignorando o valor:")
	for i, _ := range numeros {
		fmt.Printf("Índice: %d\n", i)
	}
}

// Função utilitária para imprimir informações sobre a slice
// Utility function to print slice details
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
