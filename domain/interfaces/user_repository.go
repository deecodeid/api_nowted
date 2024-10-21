package interfaces

import "github.com/deecodeid/api_nowted/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) error
	FindById(id string) (*entities.User, error)
}
