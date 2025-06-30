package main

import (
	"fmt"
	"math"
)

// Struct simples com dois campos float64
// Em Go, Vertex geralmente é um tipo struct definido para representar um ponto em um plano 2D (com coordenadas X e Y)
type Vertex struct {
	X, Y float64
}

// Esse método não altera o estado de v, então usamos o receptor de valor (v Vertex).
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// A chamada v.Abs() funciona porque v é um valor do tipo Vertex
	fmt.Println("Módulo:", v.Abs())
}
