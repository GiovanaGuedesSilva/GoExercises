package main

/*
	Exemplo de uso polimórfico de interfaces em Go.

	- Interfaces permitem trabalhar com diferentes tipos através de um mesmo ponto de acesso.
	- Podemos passar structs diferentes para a mesma função se elas implementarem a mesma interface.
*/

import "fmt"

// Interface
type Animal interface {
	Falar() string
}

// Tipos diferentes implementando a mesma interface
type Cachorro struct{}
type Gato struct{}
type Pato struct{}

func (Cachorro) Falar() string {
	return "Au au"
}

func (Gato) Falar() string {
	return "Miau"
}

func (Pato) Falar() string {
	return "Quack"
}

// Função que aceita qualquer tipo que implemente Animal
func FazerAnimalFalar(a Animal) {
	fmt.Println(a.Falar())
}

func main() {
	// Lista com múltiplos tipos que implementam a mesma interface
	animais := []Animal{Cachorro{}, Gato{}, Pato{}}

	for _, animal := range animais {
		FazerAnimalFalar(animal)
	}
}
