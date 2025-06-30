package main

/*
	Implement a version of Newton's method to find the square root of a number in Go.
*/

import (
	"fmt"
	"math"
)

// Sqrt calculates the square root of x using Newton's method.
func Sqrt(x float64) float64 {
	currentApprox := 1.0
	minPrecision := 1e-10

	for {
		nextApprox := currentApprox - (currentApprox*currentApprox - x) / (2 * currentApprox)
		if math.Abs(nextApprox - currentApprox) < minPrecision {
			break
		}
		currentApprox = nextApprox
	}

	return currentApprox
}

func main() {
	for i := 1.0; i <= 5.0; i++ {
		fmt.Printf("SquareRoot(%.0f) = %.2f | math.Sqrt = %.2f\n", i, Sqrt(i), math.Sqrt(i))
	}
}
