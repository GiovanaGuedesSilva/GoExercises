package delivery

import (
	"clean-arch/internal/usecase"
	"fmt"
)

// UserHandler representa a camada de entrega (por enquanto, CLI)
type UserHandler struct {
	usecase *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) Handle() {
	user, err := h.usecase.GetUser(1)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Usuário encontrado:", user.Name)
}
