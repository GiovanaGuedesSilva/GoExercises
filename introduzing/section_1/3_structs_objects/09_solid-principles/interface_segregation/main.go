package main

import "fmt"

// Interfaces específicas
type Leitor interface {
	Ler()
}

type Gravador interface {
	Gravar()
}

type Arquivo struct{}

func (a Arquivo) Ler() {
	fmt.Println("Lendo do arquivo")
}
func (a Arquivo) Gravar() {
	fmt.Println("Gravando no arquivo")
}

func main() {
	var r Leitor = Arquivo{}
	var w Gravador = Arquivo{}

	r.Ler()
	w.Gravar()
}
