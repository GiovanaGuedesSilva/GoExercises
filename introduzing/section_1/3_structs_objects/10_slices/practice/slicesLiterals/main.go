package main

/*
	Create a Go program that:
		Declares a slice literal with even numbers from 2 to 10.
		Prints the complete slice.
		Creates and prints three new slices:
			The first one containing the first three elements.
			The second one containing the last two elements.
			The third one containing the middle elements (excluding the first and last).
	Use only slice literals, do not use explicit arrays.
*/

import "fmt"

func main() {
	// Slice literal with even numbers from 2 to 10
	pares := []int{2, 4, 6, 8, 10}
	fmt.Println("Complete slice:", pares)

	// First three elements
	first := pares[:3]
	fmt.Println("First 3 elements:", first)

	// Last two elements
	last := pares[len(pares)-2:]
	fmt.Println("Last 2 elements:", last)

	// Middle elements (excluding the first and last)
	middle := pares[1 : len(pares)-1]
	fmt.Println("Middle elements:", middle)
}
