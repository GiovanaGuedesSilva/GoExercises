package main

/*
	Demonstração de polimorfismo e composição em Go.

	Polimorfismo:
	- Em Go, o polimorfismo é alcançado através de interfaces.
	- Tipos diferentes podem implementar a mesma interface e serem usados de forma intercambiável.

	Composição:
	- Go não possui herança como em linguagens OO tradicionais.
	- Em vez disso, Go usa *composição* — structs podem conter outras structs.
	- Métodos do tipo embutido podem ser acessados diretamente pelo tipo externo.

	Essa abordagem é conhecida como "composição sobre herança".
*/

import "fmt"

// Interface que define comportamento polimórfico
type Trabalhador interface {
	Trabalhar()
}

// Struct base (comum a todos)
type Pessoa struct {
	Nome string
}

// Método comum para todas as pessoas
func (p Pessoa) Identificar() {
	fmt.Println("Olá, meu nome é", p.Nome)
}

// Struct Engenheiro que compõe Pessoa
type Engenheiro struct {
	Pessoa
	Especialidade string
}

// Engenheiro implementa Trabalhador
func (e Engenheiro) Trabalhar() {
	fmt.Println(e.Nome, "está projetando um sistema de", e.Especialidade)
}

// Struct Professor que também compõe Pessoa
type Professor struct {
	Pessoa
	Disciplina string
}

// Professor implementa Trabalhador
func (p Professor) Trabalhar() {
	fmt.Println(p.Nome, "está dando aula de", p.Disciplina)
}

// Função que recebe qualquer Trabalhador (polimórfico)
func IniciarTrabalho(t Trabalhador) {
	t.Trabalhar()
}

func main() {
	// Criando instâncias com composição
	eng := Engenheiro{
		Pessoa:        Pessoa{Nome: "Giovana"},
		Especialidade: "Software",
	}

	prof := Professor{
		Pessoa:     Pessoa{Nome: "Carlos"},
		Disciplina: "Matemática",
	}

	// Chamando métodos da struct embutida
	eng.Identificar()
	prof.Identificar()

	// Polimorfismo em ação
	IniciarTrabalho(eng)
	IniciarTrabalho(prof)

	// Slice de interface com structs diferentes
	funcionarios := []Trabalhador{eng, prof}
	fmt.Println("\nTodos os trabalhadores:")
	for _, t := range funcionarios {
		t.Trabalhar()
	}
}
