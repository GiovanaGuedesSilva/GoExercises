package domain

// User representa a entidade central da aplicação
type User struct {
	ID   int
	Name string
}

// UserRepository define o contrato da camada de persistência
type UserRepository interface {
	FindByID(id int) (*User, error)
}
