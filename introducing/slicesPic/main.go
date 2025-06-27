package main

/*
	Implement Pic. It should return a slice of length dy,
	where each element is a slice of dx unsigned 8-bit integers.
	When you run the program, it will display your image,
	interpreting the integers as grayscale values (or bluescale, actually).

	The image logic is up to you. Interesting functions include:
	x^y, (x+y)/2, and x*y.

	(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
	(Use uint8(intValue) to convert between types.)
*/

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	// Create the image as a slice of slices
	image := make([][]uint8, dy) // height

	for y := 0; y < dy; y++ {
		row := make([]uint8, dx) // width
		for x := 0; x < dx; x++ {
			// You can change the expression below to others, such as:
			// x^y, (x+y)/2, x*y, x*y/(x+y+1), etc.
			row[x] = uint8(x ^ y) // Example using XOR
		}
		image[y] = row
	}

	return image
}

func main() {
	pic := Pic(10, 10) // Generate a 10x10 image
	for _, row := range pic {
		fmt.Println(row)
	}
}

/*
	uint8(x ^ y) → interesting checkered pattern // padrão quadriculado interessante
	uint8((x + y) / 2) → smooth gradient // transição suave
	uint8(x * y) → darkens quickly // escurece rápido
	uint8(x*y/(x+y+1)) → more complex and balanced variation // variação mais complexa e balanceada
*/
