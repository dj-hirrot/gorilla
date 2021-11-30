package usecase

import (
	"github.com/dj-hirrot/gorilla/src/domain/models"
	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	FindAll() (models.Users, error)
	FindById(id uuid.UUID) (models.User, error)
	Store(user models.User) (models.User, error)
	Update(id uuid.UUID, user models.User) (models.User, error)
	DeleteById(id uuid.UUID) error
}
