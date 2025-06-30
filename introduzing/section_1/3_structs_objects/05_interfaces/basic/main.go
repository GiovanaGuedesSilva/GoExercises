package main

/*
	Demonstração de interfaces em Go.

	- Uma interface define um conjunto de métodos.
	- Um tipo *satisfaz* uma interface automaticamente se implementar todos os métodos declarados nela.
	- Interfaces são implementadas implicitamente (sem palavras-chave como `implements`).
	- Podem ser usadas para abstração e polimorfismo.

	Declaração:
		type NomeInterface interface {
			Metodo1()
			Metodo2() string
		}
*/

import "fmt"

// Interface com um método
type Desenhavel interface {
	Desenhar()
}

// Tipo que implementa a interface
type Circulo struct {
	Raio float64
}

func (c Circulo) Desenhar() {
	fmt.Println("Desenhando um círculo com raio", c.Raio)
}

// Outro tipo que também implementa a mesma interface
type Quadrado struct {
	Lado float64
}

func (q Quadrado) Desenhar() {
	fmt.Println("Desenhando um quadrado com lado", q.Lado)
}

func main() {
	var d Desenhavel

	d = Circulo{Raio: 5}
	d.Desenhar()

	d = Quadrado{Lado: 10}
	d.Desenhar()
}
