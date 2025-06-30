package usecase

import "clean-arch/internal/domain"

// UserUseCase representa os casos de uso (lógica de negócio)
type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) GetUser(id int) (*domain.User, error) {
	return u.repo.FindByID(id)
}
