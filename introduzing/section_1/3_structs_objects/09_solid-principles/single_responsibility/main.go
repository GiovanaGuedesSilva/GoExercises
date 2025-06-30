package main

import "fmt"

// Responsabilidade única: separar lógica de usuário e email

type Usuario struct {
	Nome  string
	Email string
}

func (u Usuario) MostrarDados() {
	fmt.Println("Usuário:", u.Nome, "| Email:", u.Email)
}

type ServicoEmail struct{}

func (e ServicoEmail) EnviarEmail(destinatario, mensagem string) {
	fmt.Println("Enviando email para", destinatario)
	fmt.Println("Mensagem:", mensagem)
}

func main() {
	u := Usuario{Nome: "Giovana", Email: "giovana@example.com"}
	email := ServicoEmail{}

	u.MostrarDados()
	email.EnviarEmail(u.Email, "Bem-vinda ao Go!")
}
