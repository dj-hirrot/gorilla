package usecase

import (
	"github.com/dj-hirrot/gorilla/src/domain/entities"
	"github.com/dj-hirrot/gorilla/src/domain/models"
)

type UserRepository interface {
	FindAll() (models.Users, error)
	FindById(id int) (models.User, error)
	Store(user models.User) (models.User, error)
	Update(id int, user entities.UserParams) (models.User, error)
	DeleteById(user models.User) error
}
