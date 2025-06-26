package main

/*
	Create a Go program that:
		- Creates an array with the first 10 natural numbers (0 to 9).
		- Creates three slices from that array:
			Slice A: containing the first 3 elements.
			Slice B: containing the middle elements (4 to 6).
			Slice C: containing the last 3 elements.
		- Sums the values of each slice separately.
		- Prints the sum result of each slice with a message.
*/

import "fmt"

func main() {
	// Step 1: Create the array
	numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Step 2: Create the slices
	sliceA := numbers[0:3]
	sliceB := numbers[4:7]
	sliceC := numbers[7:10]

	// Step 3: Sum the elements
	sumA := sum(sliceA)
	sumB := sum(sliceB)
	sumC := sum(sliceC)

	// Step 4: Print the results
	fmt.Printf("Slice A: %v - Sum: %d\n", sliceA, sumA)
	fmt.Printf("Slice B: %v - Sum: %d\n", sliceB, sumB)
	fmt.Printf("Slice C: %v - Sum: %d\n", sliceC, sumC)
}

// Helper function to sum the elements of a slice
func sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}