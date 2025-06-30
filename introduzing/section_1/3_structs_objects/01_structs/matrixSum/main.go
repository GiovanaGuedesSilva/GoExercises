package main

import (
	"fmt"
)

func main() {
	/*
		Create a 3x3 matrix of integers.
		- Fill it with fixed values.
		- Calculate and print the sum of all elements.
		- Display the matrix in a table format and the total sum.
	*/

	// Define a 3x3 matrix
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Variable to store the total sum
	var sum int

	// Print the matrix and calculate the sum
	fmt.Println("Matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%2d ", matrix[i][j])
			sum += matrix[i][j]
		}
		fmt.Println()
	}

	// Print the result
	fmt.Printf("\nTotal sum: %d\n", sum)
}
