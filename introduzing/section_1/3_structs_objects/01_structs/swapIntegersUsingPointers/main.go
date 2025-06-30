package main

/*
	Implement a function that swaps the values of two integers using pointers.
*/

import "fmt"

// swap exchanges the values of two integers via their memory addresses.
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 5, 10
	fmt.Printf("Before swap: A = %d, B = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap:  A = %d, B = %d\n", x, y)
}