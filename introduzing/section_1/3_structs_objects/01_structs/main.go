package main

/*
	Demonstração do uso de structs em Go.

	- Definição de structs com campos nomeados.
	- Atribuição direta e por composição.
	- Acesso e modificação de campos.
	- Structs anônimos (inline).
	- Uso com slices.
	- Structs como retorno de função.
	- Ponteiros para structs.
*/

import "fmt"

// Definição de uma struct chamada Pessoa
type Pessoa struct {
	Nome  string
	Idade int
}

// Struct com campo embutido (composição)
type Funcionario struct {
	Pessoa
	Cargo string
}

// Função que retorna uma struct
func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{Nome: nome, Idade: idade}
}

func main() {
	// Instanciando struct com nome dos campos
	p1 := Pessoa{Nome: "Giovana", Idade: 25}
	fmt.Println("Pessoa 1:", p1)

	// Instanciando struct com valores posicionais
	p2 := Pessoa{"Lucas", 30}
	fmt.Println("Pessoa 2:", p2)

	// Acessando e modificando campos
	fmt.Println("Nome da p2:", p2.Nome)
	p2.Idade = 31
	fmt.Println("Idade atualizada da p2:", p2.Idade)

	// Struct anônima (sem definição de tipo)
	endereco := struct {
		Cidade string
		Estado string
	}{Cidade: "São Paulo", Estado: "SP"}
	fmt.Println("Endereço:", endereco)

	// Uso com slices de structs
	contatos := []Pessoa{
		{Nome: "Ana", Idade: 28},
		{Nome: "Bruno", Idade: 32},
	}
	for _, c := range contatos {
		fmt.Printf("Contato: %s (%d anos)\n", c.Nome, c.Idade)
	}

	// Struct como retorno de função
	nova := NovaPessoa("Clara", 22)
	fmt.Println("Nova pessoa:", nova)

	// Composição (struct dentro de outra struct)
	funcionario := Funcionario{
		Pessoa: Pessoa{Nome: "Daniel", Idade: 40},
		Cargo:  "Desenvolvedor",
	}
	fmt.Println("Funcionário:", funcionario)
	fmt.Println("Nome do funcionário:", funcionario.Nome)

	// Ponteiro para struct
	p3 := &Pessoa{Nome: "Rafael", Idade: 35}
	fmt.Println("Antes:", p3)
	p3.Idade++
	fmt.Println("Depois de incrementar idade:", p3)
}
