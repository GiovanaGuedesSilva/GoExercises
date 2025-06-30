package main

/*
	Entry point da aplicação (camada externa).
	Inicializa dependências e injeta as implementações concretas.
*/

import (
	"clean-arch/internal/delivery"
	"clean-arch/internal/infra"
	"clean-arch/internal/usecase"
)

func main() {
	repo := infra.NewUserRepository()
	usecase := usecase.NewUserUseCase(repo)
	handler := delivery.NewUserHandler(usecase)

	handler.Handle()
}
