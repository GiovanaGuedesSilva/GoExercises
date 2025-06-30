package main

/*
	Create a function `fibonacci` that returns a closure.
	The closure should return the next Fibonacci number each time it's called.
*/

import "fmt"

// Returns a closure that generates Fibonacci numbers
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		next := a
		a, b = b, a+b
		return next
	}
}

func main() {
	fib := fibonacci()

	fmt.Println("Fibonacci sequence:")
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
