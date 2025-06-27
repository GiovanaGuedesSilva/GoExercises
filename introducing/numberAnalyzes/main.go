package main

import "fmt"

/*
	Create a Go program that:
		- Asks the user to enter 10 integers (using fmt.Scan).
		- Stores these numbers in a slice.
		- From this slice, creates:
			- A slice with only even numbers.
			- A slice with only odd numbers.
		- Displays:
			- All entered numbers.
			- Even numbers.
			- Odd numbers.
			- The total sum of all numbers.
*/

func main() {
	var numeros []int
	var pares []int
	var impares []int
	var soma int

	fmt.Println("Enter 10 integers:")

	// Read 10 integers from the user and store them in the 'numeros' slice
	for i := 0; i < 10; i++ {
		var num int
		fmt.Print("Number ", i+1, ": ")
		fmt.Scan(&num)
		numeros = append(numeros, num)
	}

	// Separate numbers into evens and odds, and calculate the total sum
	for _, num := range numeros {
		if num%2 == 0 {
			pares = append(pares, num)
		} else {
			impares = append(impares, num)
		}
		soma += num
	}

	// Display the results
	fmt.Println("\nEntered numbers:  ", numeros)
	fmt.Println("Even numbers:     ", pares)
	fmt.Println("Odd numbers:      ", impares)
	fmt.Println("Total sum:        ", soma)
}
