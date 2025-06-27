package main

/*
	Create a closure function that returns a counter.
	Each time the function is called, it should return
	the next number in the sequence.
*/

import "fmt"

// Returns a function that increments and returns a count
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	next := counter()

	fmt.Println("Counter:")
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
}
