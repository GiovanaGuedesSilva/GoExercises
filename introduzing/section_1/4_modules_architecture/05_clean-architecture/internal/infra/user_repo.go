package infra

import (
	"clean-arch/internal/domain"
	"errors"
)

// userMemoryRepo é uma implementação fictícia do UserRepository
type userMemoryRepo struct {
	data map[int]domain.User
}

func NewUserRepository() domain.UserRepository {
	return &userMemoryRepo{
		data: map[int]domain.User{
			1: {ID: 1, Name: "Giovana"},
		},
	}
}

func (r *userMemoryRepo) FindByID(id int) (*domain.User, error) {
	user, ok := r.data[id]
	if !ok {
		return nil, errors.New("usuário não encontrado")
	}
	return &user, nil
}
