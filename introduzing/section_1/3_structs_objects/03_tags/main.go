package main

/*
	Demonstração do uso de tags em structs em Go.

	- Tags são metadados associadas aos campos de uma struct.
	- Usadas por bibliotecas (ex: encoding/json, validações, ORMs).
	- O formato da tag é: `chave:"valor"` (ex: `json:"nome"`).
	- Uma tag pode conter várias chaves separadas por espaço: `json:"id" validate:"required"`

	Uso mais comum:
		- json: define o nome do campo no JSON
		- validate: validações personalizadas com bibliotecas como "validator"
*/

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Struct com tags JSON e tags genéricas
type Produto struct {
	ID     int     `json:"id" validate:"required"`
	Nome   string  `json:"nome" validate:"required"`
	Preco  float64 `json:"preco" validate:"min=0"`
	Sigla  string  `json:"sigla,omitempty"` // omitido se vazio
	Ignore string  `json:"-"`               // campo ignorado no JSON
}

func main() {
	// Instância com valores
	p := Produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 3999.90,
		Sigla: "",
	}

	// Serialização para JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Erro ao serializar:", err)
		return
	}
	fmt.Println("JSON com tags aplicadas:", string(jsonData))

	// Inspeção de tags via reflect
	t := reflect.TypeOf(p)
	fmt.Println("\nTags dos campos da struct:")
	for i := 0; i < t.NumField(); i++ {
		campo := t.Field(i)
		fmt.Printf("Campo: %s, json: '%s', validate: '%s'\n",
			campo.Name,
			campo.Tag.Get("json"),
			campo.Tag.Get("validate"))
	}
}
