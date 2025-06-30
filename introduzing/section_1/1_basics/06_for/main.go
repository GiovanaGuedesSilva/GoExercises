package main

/*
	Demonstração de diferentes formas de usar o `for` em Go.

	Go possui apenas um tipo de laço: o `for`.
	Com ele é possível simular:
	- `for` tradicional com inicialização, condição e incremento
	- `for` tipo `while`
	- `for` infinito com `break`
	- `for range` para percorrer arrays, slices, maps e strings
*/

import "fmt"

func main() {
	// for tradicional (como em outras linguagens)
	for i := 0; i < 5; i++ {
		fmt.Println("Tradicional i:", i)
	}

	// for tipo while (só com condição)
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("While-like sum:", sum)

	// outro exemplo while
	sumb := 1
	for sumb < 1000 {
		sumb += sumb
	}
	fmt.Println("While-like sumb:", sumb)

	// for infinito com break manual
	x := 0
	for {
		x++
		if x >= 10 {
			fmt.Println("Break do loop infinito, x =", x)
			break
		}
	}

	// for range com array
	var numbers = [4]int{2, 4, 6, 8}
	for i, v := range numbers {
		fmt.Println("Array index:", i, "value:", v)
	}

	// for range com slice
	nums := []int{2, 4, 6}
	sumSlice := 0
	for i, v := range nums {
		sumSlice += v
		fmt.Println("Slice index:", i, "running sum:", sumSlice)
	}

	// for range com map
	fruits := map[string]string{"o": "orange", "b": "banana"}
	for k, v := range fruits {
		fmt.Println("Map key:", k, "value:", v)
	}

	// for range com string (caracteres Unicode)
	s := "golang hello world"
	for i, c := range s {
		fmt.Println("String index:", i, "char:", string(c))
	}
}
