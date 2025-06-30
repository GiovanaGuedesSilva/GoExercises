package main

import "fmt"

// Abstração
type Repositorio interface {
	Salvar(dado string)
}

// Implementação concreta
type BancoSQL struct{}

func (b BancoSQL) Salvar(dado string) {
	fmt.Println("Salvando em Banco SQL:", dado)
}

// Camada de serviço depende da abstração
type Servico struct {
	repo Repositorio
}

func (s Servico) Processar(dado string) {
	fmt.Println("Processando:", dado)
	s.repo.Salvar(dado)
}

func main() {
	repo := BancoSQL{}
	servico := Servico{repo: repo}
	servico.Processar("Pedido #123")
}
