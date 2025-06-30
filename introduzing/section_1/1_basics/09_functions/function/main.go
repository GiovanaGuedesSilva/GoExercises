package main

/*
	Create a function that takes another function as a parameter.
	The function should apply the given operation (sum, multiply, etc.)
	to two integers and return the result.
*/

import "fmt"

// Function type that takes two ints and returns an int
type operation func(int, int) int

// Higher-order function that receives an operation
func apply(op operation, a int, b int) int {
	return op(a, b)
}

func main() {
	// Define two simple functions
	sum := func(x, y int) int {
		return x + y
	}

	multiply := func(x, y int) int {
		return x * y
	}

	fmt.Println("Sum:", apply(sum, 3, 4))           // 7
	fmt.Println("Multiply:", apply(multiply, 3, 4)) // 12
}
