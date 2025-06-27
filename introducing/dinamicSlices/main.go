package main

/*
	Create a Go program that:
	- Asks the user for the size of an array (must be >= 9).
	- Fills the array with the first N natural numbers.
	- Creates three slices from that array:
		Slice A: the first 3 elements.
		Slice B: 3 elements from the middle.
		Slice C: the last 3 elements.
	- Calculates and prints the sum of each slice with detailed output.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter array size (must be at least 9): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	size, err := strconv.Atoi(input)

	if err != nil || size < 9 {
		fmt.Println("Invalid input. Please enter an integer >= 9.")
		return
	}

	// Step 1: Create and fill the array with natural numbers
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i
	}

	fmt.Printf("\nFull array: %v\n", numbers)

	// Step 2: Create slices
	sliceA := numbers[0:3]
	mid := size / 2
	sliceB := numbers[mid-1 : mid+2] // 3 elements centered in the middle
	sliceC := numbers[size-3 : size] // last 3 elements

	// Step 3: Sum each slice
	sumA := sum(sliceA)
	sumB := sum(sliceB)
	sumC := sum(sliceC)

	// Step 4: Print results
	fmt.Println("\n== Slice Summary ==")
	fmt.Printf("Slice A [0:3]        = %v → Sum: %d\n", sliceA, sumA)
	fmt.Printf("Slice B [%d:%d] = %v → Sum: %d\n", mid-1, mid+2, sliceB, sumB)
	fmt.Printf("Slice C [%d:%d]  = %v → Sum: %d\n", size-3, size, sliceC, sumC)
}

// Helper function to sum a slice
func sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}
