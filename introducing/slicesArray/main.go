package main

import "fmt"

func main() {
	// 🇧🇷 Criamos um array base com 5 elementos
	// 🇺🇸 We create a base array with 5 elements
	arr := [5]int{10, 20, 30, 40, 50}

	/*
		🇧🇷 Agora criamos dois slices que apontam para partes do mesmo array.
		🇺🇸 Now we create two slices that point to parts of the same array.

		🇧🇷 Slice A pega os elementos dos índices 1 a 3 (não inclui o 4):
		      => sliceA = [20, 30, 40]   // corresponde a arr[1:4]
		🇺🇸 Slice A gets elements from index 1 to 3 (does not include 4):
		      => sliceA = [20, 30, 40]   // corresponds to arr[1:4]

		🇧🇷 Slice B pega os elementos dos índices 2 a 4:
		      => sliceB = [30, 40, 50]   // corresponde a arr[2:5]
		🇺🇸 Slice B gets elements from index 2 to 4:
		      => sliceB = [30, 40, 50]   // corresponds to arr[2:5]

		🇧🇷 IMPORTANTE: Slices não armazenam dados próprios.
		🇺🇸 IMPORTANT: Slices do not store their own data.

		🇧🇷 Eles apenas "enxergam" partes do array original (matriz subjacente).
		🇺🇸 They only "view" parts of the original array (underlying array).
	*/

	sliceA := arr[1:4]
	sliceB := arr[2:5]

	// 🇧🇷 Mostra os valores antes da modificação
	// 🇺🇸 Show values before modification
	fmt.Println("ANTES DA ALTERAÇÃO / BEFORE MODIFICATION:")
	fmt.Println("Array original / Original array:", arr)
	fmt.Println("Slice A:", sliceA)
	fmt.Println("Slice B:", sliceB)

	/*
		🇧🇷 Vamos modificar sliceA[1], que é o valor 30 (índice 2 no array).
		🇺🇸 Let's modify sliceA[1], which is value 30 (index 2 in the array).

		🇧🇷 Isso altera arr[2], que também é visível em sliceB.
		🇺🇸 This changes arr[2], which is also visible in sliceB.
	*/
	sliceA[1] = 999

	// 🇧🇷 Mostra os valores após a modificação
	// 🇺🇸 Show values after the modification
	fmt.Println("\nDEPOIS DA ALTERAÇÃO / AFTER MODIFICATION:")
	fmt.Println("Array original / Original array:", arr)
	fmt.Println("Slice A:", sliceA)
	fmt.Println("Slice B:", sliceB)

	/*
		🇧🇷 CONCLUSÃO:
		🇺🇸 CONCLUSION:

		🇧🇷 - Alterar slices afeta o array original.
		🇺🇸 - Modifying slices affects the original array.

		🇧🇷 - Outros slices que compartilham o mesmo array também veem as mudanças.
		🇺🇸 - Other slices sharing the same array also see the changes.

		🇧🇷 - Slices são como janelas para arrays. Não armazenam os dados, apenas apontam.
		🇺🇸 - Slices are like windows into arrays. They don't store data, they just point to it.

		🇧🇷 ⚠️ Use com cuidado para evitar efeitos colaterais.
		🇺🇸 ⚠️ Use with caution to avoid side effects.
	*/
}
