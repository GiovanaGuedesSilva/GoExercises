package main

/*
	Demonstração do uso de JSON em Go.

	- Serialização (marshal): struct → JSON.
	- Desserialização (unmarshal): JSON → struct.
	- Uso de tags JSON para nomes personalizados de campos.
	- Structs aninhadas (nested structs).
	- Ignorando campos na serialização.
*/

import (
	"encoding/json"
	"fmt"
)

// Struct simples com tags JSON
type Pessoa struct {
	Nome     string   `json:"nome"`
	Idade    int      `json:"idade"`
	Senha    string   `json:"-"`        // ignorado no JSON
	Ativo    bool     `json:"ativo"`    // manter o mesmo nome
	Endereco Endereco `json:"endereco"` // struct aninhada
}

// Struct aninhada
type Endereco struct {
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
}

func main() {
	// Criando instância de struct
	p := Pessoa{
		Nome:  "Giovana",
		Idade: 25,
		Senha: "segredo123",
		Ativo: true,
		Endereco: Endereco{
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}

	// Serialização: struct → JSON
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}
	fmt.Println("JSON serializado:", string(jsonBytes))

	// JSON de entrada para desserialização
	entrada := `{"nome":"Lucas","idade":30,"ativo":false,"endereco":{"cidade":"Campinas","estado":"SP"}}`

	// Desserialização: JSON → struct
	var p2 Pessoa
	err = json.Unmarshal([]byte(entrada), &p2)
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		return
	}
	fmt.Println("Struct desserializada:", p2)
	fmt.Println("Nome:", p2.Nome)
	fmt.Println("Cidade:", p2.Endereco.Cidade)
}
