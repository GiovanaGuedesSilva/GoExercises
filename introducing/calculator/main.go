package main

/*
	Implements a simple calculator using a stack structure.
	Numbers are pushed onto the stack, and arithmetic operations
	(addition and subtraction) are performed using pop to manipulate values.
*/

import (
	"fmt"
	"calculator"
)

func main() {
	// Create a new stack
	s := calculator.MakeStack()

	// Push numbers 10, 5, and 3 onto the stack
	s.Push(10)
	s.Push(5)
	s.Push(3)

	// Pop the last two values (3 and 5) and add them
	a := s.Pop().Data // 3
	b := s.Pop().Data // 5
	s.Push(a + b)     // Push(8) back onto the stack

	// Pop the next value (8) and the previous (10), and subtract (10 - 8)
	c := s.Pop().Data // 8
	d := s.Pop().Data // 10
	result := d - c   // 2

	// Print the final result of the operation
	fmt.Println("Result:", result)
}
