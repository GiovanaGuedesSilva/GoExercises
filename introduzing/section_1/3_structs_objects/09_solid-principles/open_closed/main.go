package main

import "fmt"

// Interface que pode ser estendida
type Pagamento interface {
	Pagar(valor float64)
}

type CartaoCredito struct{}

func (c CartaoCredito) Pagar(valor float64) {
	fmt.Println("Pagando com Cartão de Crédito:", valor)
}

type Boleto struct{}

func (b Boleto) Pagar(valor float64) {
	fmt.Println("Pagando com Boleto:", valor)
}

func processarPagamento(p Pagamento, valor float64) {
	p.Pagar(valor)
}

func main() {
	processarPagamento(CartaoCredito{}, 100.0)
	processarPagamento(Boleto{}, 50.0)
}
