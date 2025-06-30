package main

import (
	"fmt"
)

// Tipo personalizado baseado em float64
// MyFloat é um alias de float64, mas é considerado um tipo distinto
type MyFloat float64

// Método em MyFloat, com receptor de valor
// Métodos podem ser definidos em qualquer tipo declarado localmente, não apenas em structs
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	n := MyFloat(-7.5)
	fmt.Println("Valor absoluto:", n.Abs()) // Resultado: 7.5
}
