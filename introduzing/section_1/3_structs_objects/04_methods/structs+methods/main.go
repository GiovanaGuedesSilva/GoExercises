package main

/*
	Métodos associados a structs personalizadas.

	- Métodos permitem encapsular lógica dentro do próprio tipo.
	- Podem ser usados para validação, formatação, operações específicas etc.
	- Ajuda na organização orientada a dados, mesmo sem herança tradicional.
*/

import (
	"fmt"
	"strings"
)

// Struct com campos
type Produto struct {
	Nome  string
	Preco float64
}

// Método para aplicar desconto
func (p *Produto) AplicarDesconto(percentual float64) {
	if percentual > 0 && percentual < 100 {
		desconto := p.Preco * percentual / 100
		p.Preco -= desconto
	}
}

// Método para verificar se é caro
func (p Produto) EhCaro() bool {
	return p.Preco > 1000
}

// Método para representar como string formatada
func (p Produto) Descricao() string {
	return fmt.Sprintf("Produto: %s | Preço: R$ %.2f", strings.ToUpper(p.Nome), p.Preco)
}

func main() {
	prod := Produto{Nome: "Notebook", Preco: 3500}

	fmt.Println(prod.Descricao())
	fmt.Println("É caro?", prod.EhCaro())

	prod.AplicarDesconto(10)
	fmt.Println(prod.Descricao())
}
