package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// Método com receptor ponteiro (*Vertex), que altera os valores da struct
// Como esse método modifica os campos da struct, usamos *Vertex.
func (v *Vertex) Scale(factor float64) {
	v.X *= factor
	v.Y *= factor
}

func main() {
	v := Vertex{2, 3}
	v.Scale(2)                  // O Go permite chamar com v.Scale(2), mesmo que v seja um valor — ele converte implicitamente para &v.
	fmt.Println("Escalado:", v) // Resultado: {4, 6}
}
