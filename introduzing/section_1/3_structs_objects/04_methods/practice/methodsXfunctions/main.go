package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Função normal
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Método
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{6, 8}

	// Ambas funcionam com valor
	fmt.Println("Função:", AbsFunc(v))    // OK
	fmt.Println("Método:", v.AbsMethod()) // OK

	// fmt.Println(AbsFunc(&v))             // ❌ ERRO: espera Vertex, não *Vertex
	// Mas método com valor aceita ponteiro:
	p := &v
	fmt.Println("Método com ponteiro:", p.AbsMethod()) // OK (Go faz conversão automática)

	// Funções não convertem automaticamente de *Vertex para Vertex
	// Métodos fazem essa conversão automaticamente — por isso p.AbsMethod() funciona.
}
